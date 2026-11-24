use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_AppCustomerOrgTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub app_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    pub customer_org_id: String,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    pub spend_reset_window: Option<String>,
    pub tier_budget_id: Option<String>,
    pub budget_id: Option<String>,
    pub budget_override_id: Option<String>,
    pub budget_override_expires_at: Option<DateTime>,
    pub customer_member_budget_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        on_delete = "Restrict"
    )]
    Budget,
    #[sea_orm(
        belongs_to = "super::lite_llm_app_table::Entity",
        from = "Column::AppId",
        to = "super::lite_llm_app_table::Column::AppId",
        on_delete = "Cascade"
    )]
    App,
    #[sea_orm(
        belongs_to = "super::lite_llm_customer_organization_table::Entity",
        from = "Column::CustomerOrgId",
        to = "super::lite_llm_customer_organization_table::Column::CustomerOrgId",
        on_delete = "Cascade"
    )]
    CustomerOrg,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::TierBudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId"
    )]
    TierBudget,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetOverrideId",
        to = "super::lite_llm_budget_table::Column::BudgetId"
    )]
    BudgetOverride,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::CustomerMemberBudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId"
    )]
    CustomerMemberBudget,
}

impl Related<super::lite_llm_budget_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Budget.def()
    }
}

impl Related<super::lite_llm_app_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::App.def()
    }
}

impl Related<super::lite_llm_customer_organization_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomerOrg.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
