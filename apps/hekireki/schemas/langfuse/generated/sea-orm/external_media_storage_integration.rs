use super::external_media_storage_integration_type::ExternalMediaStorageIntegrationType;
use super::storage_integration_auth_type::StorageIntegrationAuthType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "external_media_storage_integrations")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
    pub r#type: ExternalMediaStorageIntegrationType,
    #[sea_orm(column_type = "Text")]
    pub bucket_name: String,
    #[sea_orm(column_type = "Text")]
    pub prefix: Option<String>,
    #[sea_orm(default_value = "ACCESS_KEYS")]
    pub auth_type: StorageIntegrationAuthType,
    #[sea_orm(column_type = "Text")]
    pub access_key_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub secret_access_key: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub role_arn: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub external_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub region: String,
    #[sea_orm(column_type = "Text")]
    pub endpoint: Option<String>,
    pub force_path_style: bool,
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
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "external_media_storage_integrations_project_id_fkey"
    )]
    Project,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
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
