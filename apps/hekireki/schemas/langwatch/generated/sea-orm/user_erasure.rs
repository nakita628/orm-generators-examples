use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One row per erased account, holding only its creation time: no id, name or address. It keeps
/// the install's first-created account knowable after an erasure deletes that row.
/// Spec: modules/user/specs/user.feature
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "UserErasure")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "accountCreatedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub account_created_at: DateTime,
    #[sea_orm(
        column_name = "erasedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub erased_at: DateTime,
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
        if insert && self.erased_at.is_not_set() {
            self.erased_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
