use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Each project and UTC day with a completed evaluation, from evaluation's lifecycle_completed,
/// written insert-ignore. Replaces the `evaluation` rows of NurturingProjectUsageView.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingProjectEvaluationDay")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// UTC `YYYY-MM-DD`.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub day: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
