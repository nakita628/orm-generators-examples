use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant").
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuthzUserStanding")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(
        column_name = "deactivatedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub deactivated_at: Option<DateTime>,
    #[sea_orm(column_name = "erasedAt", column_type = "custom(\"timestamp(3)\")")]
    pub erased_at: Option<DateTime>,
    #[sea_orm(
        column_name = "standingChangedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub standing_changed_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
