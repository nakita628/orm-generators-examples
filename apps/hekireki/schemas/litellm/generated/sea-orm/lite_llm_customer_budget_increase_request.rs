use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_CustomerBudgetIncreaseRequest")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub request_id: String,
    #[sea_orm(column_type = "Text")]
    pub customer_member_id: String,
    pub requested_max_budget: f64,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub reason: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "pending_review")]
    pub status: String,
    #[sea_orm(column_type = "Text")]
    pub submitted_by: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub submitted_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub reviewed_by: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub reviewed_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub review_notes: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_customer_member_table::Entity",
        from = "Column::CustomerMemberId",
        to = "super::lite_llm_customer_member_table::Column::CustomerMemberId",
        on_delete = "Cascade",
        fk_name = "LiteLLM_CustomerBudgetIncreaseRequest_customer_member_id_fkey"
    )]
    CustomerMember,
}

impl Related<super::lite_llm_customer_member_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomerMember.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            request_id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.submitted_at.is_not_set() {
            self.submitted_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
