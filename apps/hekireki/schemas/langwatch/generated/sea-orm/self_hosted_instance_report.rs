use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One usage report, exactly as it arrived. Kept so a number on an install's
/// row can be traced back to the report that set it, and so a question asked
/// later can be answered over the history rather than over the latest snapshot.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SelfHostedInstanceReport")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The `SelfHostedInstance.instanceId` this report came from. A plain column
    /// ref for the same reason the instance row's refs are.
    #[sea_orm(column_name = "instanceId", column_type = "Text")]
    pub instance_id: String,
    #[sea_orm(
        column_name = "receivedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub received_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub version: Option<String>,
    #[sea_orm(column_name = "reportSchemaVersion")]
    pub report_schema_version: Option<i32>,
    /// Fields this release had no name for. Counted, never stored by name.
    #[sea_orm(column_name = "unknownFields", default_value = 0)]
    pub unknown_fields: i32,
    /// The validated report, with every unknown field already dropped.
    #[sea_orm(column_type = "JsonBinary")]
    pub payload: Json,
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
