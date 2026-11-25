use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_OrganizationMembership")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub user_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_type = "Text")]
    pub user_role: Option<String>,
    #[sea_orm(default_value = 0)]
    pub spend: Option<f64>,
    #[sea_orm(column_type = "Text")]
    pub budget_id: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: Option<DateTime>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_user_table::Entity",
        from = "Column::UserId",
        to = "super::lite_llm_user_table::Column::UserId",
        fk_name = "LiteLLM_OrganizationMembership_user_id_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::lite_llm_organization_table::Entity",
        from = "Column::OrganizationId",
        to = "super::lite_llm_organization_table::Column::OrganizationId",
        fk_name = "LiteLLM_OrganizationMembership_organization_id_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        fk_name = "LiteLLM_OrganizationMembership_budget_id_fkey"
    )]
    LitellmBudgetTable,
}

impl Related<super::lite_llm_user_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::lite_llm_organization_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::lite_llm_budget_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LitellmBudgetTable.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(Some(now.naive_utc()));
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
