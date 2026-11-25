use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_JWTKeyMapping")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub jwt_issuer: String,
    #[sea_orm(column_type = "Text")]
    pub jwt_claim_name: String,
    #[sea_orm(column_type = "Text")]
    pub jwt_claim_value: String,
    #[sea_orm(column_type = "Text")]
    pub token: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(default_value = true)]
    pub is_active: bool,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub updated_by: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_verification_token::Entity",
        from = "Column::Token",
        to = "super::lite_llm_verification_token::Column::Token",
        on_delete = "Cascade",
        fk_name = "LiteLLM_JWTKeyMapping_token_fkey"
    )]
    LitellmVerificationToken,
}

impl Related<super::lite_llm_verification_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LitellmVerificationToken.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
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
