use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_AppCustomerMemberTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub app_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub customer_member_id: String,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    #[sea_orm(column_type = "Text")]
    pub spend_reset_window: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub budget_override_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub budget_id: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub budget_override_expires_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        on_delete = "Restrict",
        fk_name = "LiteLLM_AppCustomerMemberTable_budget_id_fkey"
    )]
    Budget,
    #[sea_orm(
        belongs_to = "super::lite_llm_app_table::Entity",
        from = "Column::AppId",
        to = "super::lite_llm_app_table::Column::AppId",
        on_delete = "Cascade",
        fk_name = "LiteLLM_AppCustomerMemberTable_app_id_fkey"
    )]
    App,
    #[sea_orm(
        belongs_to = "super::lite_llm_customer_member_table::Entity",
        from = "Column::CustomerMemberId",
        to = "super::lite_llm_customer_member_table::Column::CustomerMemberId",
        on_delete = "Cascade",
        fk_name = "LiteLLM_AppCustomerMemberTable_customer_member_id_fkey"
    )]
    CustomerMember,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetOverrideId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        fk_name = "LiteLLM_AppCustomerMemberTable_budget_override_id_fkey"
    )]
    BudgetOverride,
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

impl Related<super::lite_llm_customer_member_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomerMember.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
