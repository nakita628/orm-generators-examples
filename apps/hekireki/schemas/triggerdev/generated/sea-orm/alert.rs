use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Alert")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "environmentId", column_type = "Text")]
    pub environment_id: String,
    #[sea_orm(column_name = "capabilityKey", column_type = "Text")]
    pub capability_key: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub params: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub condition: Json,
    #[sea_orm(column_type = "Text")]
    pub query: Option<String>,
    #[sea_orm(column_name = "expiresAt", column_type = "TimestampWithTimeZone")]
    pub expires_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_type = "Text", default_value = "active")]
    pub lifecycle: String,
    #[sea_orm(column_type = "Text", default_value = "clear")]
    pub attention: String,
    #[sea_orm(
        column_name = "observedCondition",
        column_type = "Text",
        default_value = "unknown"
    )]
    pub observed_condition: String,
    #[sea_orm(
        column_name = "observedConditionChangedAt",
        column_type = "TimestampWithTimeZone"
    )]
    pub observed_condition_changed_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "pendingMetSince", column_type = "TimestampWithTimeZone")]
    pub pending_met_since: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "evaluatorState", column_type = "JsonBinary")]
    pub evaluator_state: Option<Json>,
    #[sea_orm(column_name = "evaluationLeaseId", column_type = "Text")]
    pub evaluation_lease_id: Option<String>,
    #[sea_orm(
        column_name = "evaluationLeaseUntil",
        column_type = "TimestampWithTimeZone"
    )]
    pub evaluation_lease_until: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "lastEvaluatedAt", column_type = "TimestampWithTimeZone")]
    pub last_evaluated_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "nextEvaluationAt",
        column_type = "TimestampWithTimeZone"
    )]
    pub next_evaluation_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub notifiers: Json,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTimeWithTimeZone,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTimeWithTimeZone,
    #[sea_orm(column_name = "resolvedAt", column_type = "TimestampWithTimeZone")]
    pub resolved_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "resolvedBy", column_type = "Text")]
    pub resolved_by: Option<String>,
    #[sea_orm(column_name = "stoppedAt", column_type = "TimestampWithTimeZone")]
    pub stopped_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "stoppedBy", column_type = "Text")]
    pub stopped_by: Option<String>,
    #[sea_orm(
        column_name = "lastOccurrenceAt",
        column_type = "TimestampWithTimeZone"
    )]
    pub last_occurrence_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "occurrencesSinceRaise", default_value = 0)]
    pub occurrences_since_raise: i32,
    #[sea_orm(column_name = "occurrencesTotal", default_value = "0")]
    pub occurrences_total: i64,
    #[sea_orm(column_name = "groupId", column_type = "Text")]
    pub group_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.fixed_offset());
        }
        if insert && self.updated_at.is_not_set() {
            self.updated_at = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
