use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "llm_api_keys")]
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
    pub provider: String,
    #[sea_orm(column_type = "Text")]
    pub adapter: String,
    #[sea_orm(column_type = "Text")]
    pub display_secret_key: String,
    #[sea_orm(column_type = "Text")]
    pub secret_key: String,
    #[sea_orm(column_type = "Text")]
    pub base_url: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub custom_models: Vec<String>,
    #[sea_orm(default_value = true)]
    pub with_default_models: bool,
    #[sea_orm(column_type = "Text")]
    pub extra_headers: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub extra_header_keys: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub config: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub external_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub organization_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "llm_api_keys_project_id_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "llm_api_keys_organization_id_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::default_llm_model::Entity")]
    DefaultLlmModel,
    #[sea_orm(
        has_many = "super::topic_config::Entity",
        from = "Column::Id",
        to = "super::topic_config::Column::SummaryLlmApiKeyId"
    )]
    TopicSummaryConfigs,
    #[sea_orm(
        has_many = "super::topic_config::Entity",
        from = "Column::Id",
        to = "super::topic_config::Column::EmbeddingLlmApiKeyId"
    )]
    TopicEmbeddingConfigs,
    #[sea_orm(
        has_many = "super::topic_config::Entity",
        from = "Column::Id",
        to = "super::topic_config::Column::ClusteringLlmApiKeyId"
    )]
    TopicClusteringConfigs,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::default_llm_model::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DefaultLlmModel.def()
    }
}

impl Related<super::topic_config::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TopicSummaryConfigs.def()
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
