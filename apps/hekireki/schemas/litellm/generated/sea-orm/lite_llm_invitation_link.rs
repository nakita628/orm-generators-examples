use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_InvitationLink")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub user_id: String,
    #[sea_orm(default_value = false)]
    pub is_accepted: bool,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub accepted_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub updated_by: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_user_table::Entity",
        from = "Column::UserId",
        to = "super::lite_llm_user_table::Column::UserId",
        fk_name = "LiteLLM_InvitationLink_user_id_fkey"
    )]
    LiteLLMUserTableUser,
    #[sea_orm(
        belongs_to = "super::lite_llm_user_table::Entity",
        from = "Column::CreatedBy",
        to = "super::lite_llm_user_table::Column::UserId",
        fk_name = "LiteLLM_InvitationLink_created_by_fkey"
    )]
    LiteLLMUserTableCreated,
    #[sea_orm(
        belongs_to = "super::lite_llm_user_table::Entity",
        from = "Column::UpdatedBy",
        to = "super::lite_llm_user_table::Column::UserId",
        fk_name = "LiteLLM_InvitationLink_updated_by_fkey"
    )]
    LiteLLMUserTableUpdated,
}

impl Related<super::lite_llm_user_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LiteLLMUserTableUser.def()
    }
}

impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }
}
