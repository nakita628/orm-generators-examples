use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// RANGE-partitioned by month on `at` (see migration).
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AlertEvent")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "alertId", column_type = "Text")]
    pub alert_id: String,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub at: DateTimeWithTimeZone,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub facts: Json,
    #[sea_orm(column_name = "actorId", column_type = "Text")]
    pub actor_id: Option<String>,
    #[sea_orm(column_name = "sourceEventId", column_type = "Text")]
    pub source_event_id: Option<String>,
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
        if insert && self.at.is_not_set() {
            self.at = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
