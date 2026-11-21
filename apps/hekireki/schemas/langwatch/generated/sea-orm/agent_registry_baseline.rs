use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// An organisation's registry baseline: its first reconcile, written as already pushed so the
/// agents it already had are not announced. Every reconcile after it pushes its changes.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AgentRegistryBaseline")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "baselinedAt")]
    pub baselined_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
