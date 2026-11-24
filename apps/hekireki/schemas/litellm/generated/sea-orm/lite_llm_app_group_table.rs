use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_AppGroupTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub app_group_id: String,
    #[sea_orm(unique)]
    pub name: String,
    pub description: Option<String>,
    pub metadata: Option<Json>,
    pub budget_id: Option<String>,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    pub spend_reset_window: Option<String>,
    pub created_at: DateTime,
    pub created_by: String,
    pub updated_at: DateTime,
    pub updated_by: String,
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
    #[sea_orm(has_many = "super::lite_llm_app_table::Entity")]
    Apps,
    #[sea_orm(has_many = "super::lite_llm_app_group_customer_org_table::Entity")]
    CustomerOrganizations,
}

impl Related<super::lite_llm_budget_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Budget.def()
    }
}

impl Related<super::lite_llm_app_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Apps.def()
    }
}

impl Related<super::lite_llm_app_group_customer_org_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomerOrganizations.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            app_group_id: Set(uuid::Uuid::new_v4().to_string()),
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
