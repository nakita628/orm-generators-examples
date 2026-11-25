use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// An organisation's agent registry settings. No row reads as the defaults.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AgentRegistrySettings")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// Days without traffic after which a listed agent reads as dormant.
    #[sea_orm(column_name = "dormantAfterDays", default_value = 30)]
    pub dormant_after_days: i32,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, _insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
