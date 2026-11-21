use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Persisted anomaly detection. Inserted by the anomaly-detection
/// reactor (event-sourcing pipeline; modeled on PR #3351's
/// alertTrigger pattern) when a rule's threshold trips while
/// processing a freshly-arrived ActivityEventReceived event.
/// Read by api.activityMonitor.recentAnomalies for the admin dashboard.
///
/// Dedup: (ruleId, triggerWindowStart) - re-evaluating the same
/// window updates the existing row rather than creating a duplicate.
/// Acknowledgement / resolution are admin actions on the row.
///
/// Spec: specs/ai-gateway/governance/anomaly-detection.feature
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AnomalyAlert")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "ruleId")]
    pub rule_id: String,
    /// Denormalised from rule for fast WHERE filtering on the
    /// dashboard query (no join needed for the common path).
    pub severity: String,
    #[sea_orm(column_name = "ruleName")]
    pub rule_name: String,
    #[sea_orm(column_name = "ruleType")]
    pub rule_type: String,
    /// Window the trigger refers to (UTC).
    #[sea_orm(column_name = "triggerWindowStart")]
    pub trigger_window_start: DateTime,
    #[sea_orm(column_name = "triggerWindowEnd")]
    pub trigger_window_end: DateTime,
    /// Numeric trigger context. Per-rule-type detail in `detail` JSONB.
    #[sea_orm(
        column_name = "triggerSpendUsd",
        column_type = "Decimal(Some((18, 6)))"
    )]
    pub trigger_spend_usd: Option<Decimal>,
    #[sea_orm(column_name = "triggerEventCount")]
    pub trigger_event_count: Option<i32>,
    /// Per-rule-type extra context: {topActors: [...], baseline: ...,
    /// observed: ..., timezone: ..., etc}.
    #[sea_orm(default_value = "{}")]
    pub detail: Json,
    /// 'open' | 'acknowledged' | 'resolved'.
    #[sea_orm(default_value = "open")]
    pub state: String,
    /// Per-destination dispatch outcome:
    /// { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
    /// slack: ..., pagerduty: ..., siem: ..., email: ... }
    #[sea_orm(column_name = "destinationStatus", default_value = "{}")]
    pub destination_status: Json,
    #[sea_orm(column_name = "detectedAt")]
    pub detected_at: DateTime,
    #[sea_orm(column_name = "acknowledgedAt")]
    pub acknowledged_at: Option<DateTime>,
    #[sea_orm(column_name = "resolvedAt")]
    pub resolved_at: Option<DateTime>,
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
        belongs_to = "super::anomaly_rule::Entity",
        from = "Column::RuleId",
        to = "super::anomaly_rule::Column::Id",
        on_delete = "Cascade"
    )]
    Rule,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::anomaly_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Rule.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.detected_at.is_not_set() {
            self.detected_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
