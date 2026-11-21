use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// When an organization was created, from organization's created event: nurturing's own copy for
/// the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
/// NurturingOrganization row.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingOrganizationCreated")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
