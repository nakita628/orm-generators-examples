use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Nurturing's copy of each personal project's owner, from organization's and project's facts.
/// The first owner recorded stays: a personal project's owner never changes.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingPersonalProject")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "ownerUserId", column_type = "Text")]
    pub owner_user_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
