use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_CustomerTeamTable")]
pub struct Model {
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    #[sea_orm(column_type = "Text")]
    pub spend_reset_window: Option<String>,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub customer_team_id: String,
    #[sea_orm(column_type = "Text")]
    pub alias: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub parent_org_id: String,
    #[sea_orm(column_type = "Text")]
    pub budget_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub budget_override_id: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub budget_override_expires_at: Option<DateTime>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub updated_by: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_customer_organization_table::Entity",
        from = "Column::ParentOrgId",
        to = "super::lite_llm_customer_organization_table::Column::CustomerOrgId",
        fk_name = "LiteLLM_CustomerTeamTable_parent_org_id_fkey"
    )]
    ParentOrg,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        fk_name = "LiteLLM_CustomerTeamTable_budget_id_fkey"
    )]
    Budget,
    #[sea_orm(
        belongs_to = "super::lite_llm_budget_table::Entity",
        from = "Column::BudgetOverrideId",
        to = "super::lite_llm_budget_table::Column::BudgetId",
        on_delete = "Restrict",
        fk_name = "LiteLLM_CustomerTeamTable_budget_override_id_fkey"
    )]
    BudgetOverride,
}

impl Related<super::lite_llm_customer_organization_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ParentOrg.def()
    }
}

impl Related<super::lite_llm_budget_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Budget.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            customer_team_id: Set(uuid::Uuid::new_v4().to_string()),
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
