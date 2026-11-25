use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_AppGroupCustomerOrgTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub app_group_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub customer_org_id: String,
    #[sea_orm(column_type = "Text")]
    pub budget_id: Option<String>,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    #[sea_orm(column_type = "Text")]
    pub spend_reset_window: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        on_delete = "Restrict",
        fk_name = "LiteLLM_AppGroupCustomerOrgTable_budget_id_fkey"
    )]
    Budget,
    #[sea_orm(
        belongs_to = "super::lite_llm_app_group_table::Entity",
        from = "Column::AppGroupId",
        to = "super::lite_llm_app_group_table::Column::AppGroupId",
        on_delete = "Cascade",
        fk_name = "LiteLLM_AppGroupCustomerOrgTable_app_group_id_fkey"
    )]
    AppGroup,
    #[sea_orm(
        belongs_to = "super::lite_llm_customer_organization_table::Entity",
        from = "Column::CustomerOrgId",
        to = "super::lite_llm_customer_organization_table::Column::CustomerOrgId",
        on_delete = "Cascade",
        fk_name = "LiteLLM_AppGroupCustomerOrgTable_customer_org_id_fkey"
    )]
    CustomerOrg,
}

impl Related<super::lite_llm_budget_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Budget.def()
    }
}

impl Related<super::lite_llm_app_group_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AppGroup.def()
    }
}

impl Related<super::lite_llm_customer_organization_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomerOrg.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
