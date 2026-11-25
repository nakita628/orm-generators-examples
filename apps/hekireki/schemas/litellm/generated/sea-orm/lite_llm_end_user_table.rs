use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_EndUserTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "Text")]
    pub alias: Option<String>,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    #[sea_orm(column_type = "Text")]
    pub spend_reset_window: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub allowed_model_region: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub default_model: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub models: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub budget_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub object_permission_id: Option<String>,
    #[sea_orm(default_value = false)]
    pub blocked: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        fk_name = "LiteLLM_EndUserTable_budget_id_fkey"
    )]
    LitellmBudgetTable,
    #[sea_orm(
        belongs_to = "super::lite_llm_object_permission_table::Entity",
        from = "Column::ObjectPermissionId",
        to = "super::lite_llm_object_permission_table::Column::ObjectPermissionId",
        fk_name = "LiteLLM_EndUserTable_object_permission_id_fkey"
    )]
    ObjectPermission,
}

impl Related<super::lite_llm_budget_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LitellmBudgetTable.def()
    }
}

impl Related<super::lite_llm_object_permission_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ObjectPermission.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
