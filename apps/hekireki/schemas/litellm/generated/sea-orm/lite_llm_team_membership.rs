use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_TeamMembership")]
pub struct Model {
    pub rollover_credit: Option<f64>,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub user_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub team_id: String,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    #[sea_orm(default_value = 0)]
    pub total_spend: f64,
    #[sea_orm(column_type = "Text")]
    pub budget_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        fk_name = "LiteLLM_TeamMembership_budget_id_fkey"
    )]
    LitellmBudgetTable,
}

impl Related<super::lite_llm_budget_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LitellmBudgetTable.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
