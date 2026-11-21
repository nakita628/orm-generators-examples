use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Nurturing's copy of each project that received a real trace, from trace's and project's facts.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingTracedProject")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
