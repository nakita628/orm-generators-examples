use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Admin-defined anomaly detection rule. Evaluated server-side at
/// ingest time (Option C - not in this slice yet) against
/// gateway_activity_events. When triggered, an alert is dispatched
/// to the destinations configured on the rule (Slack / SIEM / webhook
/// / PagerDuty / email).
///
/// This is a configuration entity - the actual rule evaluation +
/// alert state lives in a future ClickHouse table
/// (gateway_anomaly_alerts) and a server-side worker. Spec:
/// specs/ai-gateway/governance/anomaly-rules.feature
/// (Alexis lane B C-UI; mocked-v0 today, real persistence is THIS slice).
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AnomalyRule")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// 'organization' | 'team' | 'project' | 'source_type' | 'source'.
    /// String not enum so new scopes ship without migrations.
    pub scope: String,
    /// Org id / team id / project id / source-type string / source id
    /// - narrows what the rule applies to.
    #[sea_orm(column_name = "scopeId")]
    pub scope_id: String,
    pub name: String,
    pub description: Option<String>,
    /// 'critical' | 'warning' | 'info'. Drives default destination
    /// routing (e.g. critical → PagerDuty, warning → Slack only).
    pub severity: String,
    /// Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
    /// 'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
    /// | 'custom'. New types ship without migrations - evaluator
    /// picks a per-type implementation from a code registry.
    #[sea_orm(column_name = "ruleType")]
    pub rule_type: String,
    /// Per-rule-type threshold shape. Examples:
    /// spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
    /// rate_limit: { windowSec, requestsThreshold }
    /// after_hours: { startHour, endHour, timezone, requestsThreshold }
    #[sea_orm(column_name = "thresholdConfig", default_value = "{}")]
    pub threshold_config: Json,
    /// Alert routing per destination. Shape:
    /// { slack?: { channelId }, siem?: { endpointUrl, authHeader },
    /// webhook?: { url, authHeader }, pagerduty?: { routingKey },
    /// email?: { to: string[] } }
    /// Empty map = use org defaults (when defaults exist; today no-op).
    #[sea_orm(column_name = "destinationConfig", default_value = "{}")]
    pub destination_config: Json,
    /// 'active' (default - evaluator picks it up) | 'disabled' (admin
    /// paused; rule rows kept for audit). 'archived' uses archivedAt
    /// timestamp instead so we can filter `archivedAt IS NULL` cheaply.
    #[sea_orm(default_value = "active")]
    pub status: String,
    #[sea_orm(column_name = "archivedAt")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull"
    )]
    CreatedBy,
    #[sea_orm(has_many = "super::anomaly_alert::Entity")]
    Alerts,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::anomaly_alert::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Alerts.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
