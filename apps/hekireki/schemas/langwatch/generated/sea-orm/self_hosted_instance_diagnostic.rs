use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One upgrade diagnostics bundle a self-hosted install sent over Connect, on
/// LangWatch Cloud. The bytes live in object storage under
/// `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
/// points at them and feeds Cloud admin and the Slack signal. Plain column
/// refs, no relation, as the usage report rows.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SelfHostedInstanceDiagnostic")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The `SelfHostedInstance.instanceId` the licence caller presented.
    #[sea_orm(column_name = "instanceId", column_type = "Text")]
    pub instance_id: String,
    #[sea_orm(column_name = "stepId", column_type = "Text")]
    pub step_id: String,
    #[sea_orm(column_name = "stepDescription", column_type = "Text")]
    pub step_description: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub release: Option<String>,
    pub attempt: i32,
    /// `operator` or `automatic`.
    #[sea_orm(column_type = "Text")]
    pub trigger: String,
    #[sea_orm(column_name = "byteLength")]
    pub byte_length: i32,
    #[sea_orm(column_name = "objectKey", column_type = "Text")]
    pub object_key: String,
    #[sea_orm(
        column_name = "receivedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub received_at: DateTime,
    /// When the lifecycle rule removes the bundle; the row is listed until then.
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    /// Set on the report that posted this step and release's one Slack message.
    #[sea_orm(
        column_name = "slackAnnouncedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub slack_announced_at: Option<DateTime>,
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
        if insert && self.received_at.is_not_set() {
            self.received_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
