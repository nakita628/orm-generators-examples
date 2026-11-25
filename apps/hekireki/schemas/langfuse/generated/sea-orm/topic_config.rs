use super::evaluator_block_reason::EvaluatorBlockReason;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "topic_configs")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub updated_by: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
    #[sea_orm(default_value = false)]
    pub enabled: bool,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub blocked_at: Option<DateTime>,
    pub block_reason: Option<EvaluatorBlockReason>,
    #[sea_orm(column_type = "Text")]
    pub block_message: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub summary_llm_api_key_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub summary_model: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub embedding_llm_api_key_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub embedding_model: Option<String>,
    #[sea_orm(default_value = 1024)]
    pub embedding_dimensions: i32,
    #[sea_orm(column_type = "Text")]
    pub clustering_llm_api_key_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub clustering_model: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "topic_configs_created_by_fkey"
    )]
    CreatedByUser,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UpdatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "topic_configs_updated_by_fkey"
    )]
    UpdatedByUser,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "topic_configs_project_id_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::llm_api_keys::Entity",
        from = "Column::SummaryLlmApiKeyId",
        to = "super::llm_api_keys::Column::Id",
        on_delete = "SetNull",
        fk_name = "topic_configs_summary_llm_api_key_id_fkey"
    )]
    SummaryLlmApiKey,
    #[sea_orm(
        belongs_to = "super::llm_api_keys::Entity",
        from = "Column::EmbeddingLlmApiKeyId",
        to = "super::llm_api_keys::Column::Id",
        on_delete = "SetNull",
        fk_name = "topic_configs_embedding_llm_api_key_id_fkey"
    )]
    EmbeddingLlmApiKey,
    #[sea_orm(
        belongs_to = "super::llm_api_keys::Entity",
        from = "Column::ClusteringLlmApiKeyId",
        to = "super::llm_api_keys::Column::Id",
        on_delete = "SetNull",
        fk_name = "topic_configs_clustering_llm_api_key_id_fkey"
    )]
    ClusteringLlmApiKey,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedByUser.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::llm_api_keys::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SummaryLlmApiKey.def()
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
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
