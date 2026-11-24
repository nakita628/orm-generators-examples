use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_AppTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub app_id: String,
    pub app_group_id: Option<String>,
    pub resource_type: Option<String>,
    pub resource_id: Option<String>,
    #[sea_orm(unique)]
    pub alias: String,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    pub spend_reset_window: Option<String>,
    pub budget_id: Option<String>,
    pub default_tier_budget_id: Option<String>,
    pub default_customer_member_budget_id: Option<String>,
    #[sea_orm(default_value = true)]
    pub auto_create_customers: bool,
    pub customer_margin: Option<f64>,
    pub created_at: DateTime,
    pub created_by: String,
    pub updated_at: DateTime,
    pub updated_by: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_app_group_table::Entity",
        from = "Column::AppGroupId",
        to = "super::lite_llm_app_group_table::Column::AppGroupId",
        on_delete = "Restrict"
    )]
    AppGroup,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId"
    )]
    Budget,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::DefaultTierBudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId"
    )]
    DefaultTierBudget,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::DefaultCustomerMemberBudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId"
    )]
    DefaultCustomerMemberBudget,
    #[sea_orm(has_many = "super::lite_llm_app_customer_org_table::Entity")]
    CustomerOrganizations,
    #[sea_orm(has_many = "super::lite_llm_app_customer_member_table::Entity")]
    CustomerMembers,
}

impl Related<super::lite_llm_app_group_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AppGroup.def()
    }
}

impl Related<super::lite_llm_budget_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Budget.def()
    }
}

impl Related<super::lite_llm_app_customer_org_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomerOrganizations.def()
    }
}

impl Related<super::lite_llm_app_customer_member_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomerMembers.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            app_id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
