use super::project_alert_channel_type::ProjectAlertChannelType;
use super::project_alert_type::ProjectAlertType;
use super::runtime_environment_type::RuntimeEnvironmentType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ProjectAlertChannel")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    /// can be provided and we won't create another with the same key
    #[sea_orm(column_name = "deduplicationKey", column_type = "Text")]
    pub deduplication_key: String,
    #[sea_orm(column_name = "userProvidedDeduplicationKey", default_value = false)]
    pub user_provided_deduplication_key: bool,
    #[sea_orm(column_name = "integrationId", column_type = "Text")]
    pub integration_id: Option<String>,
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    pub r#type: ProjectAlertChannelType,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub properties: Json,
    #[sea_orm(column_name = "alertTypes")]
    pub alert_types: Vec<ProjectAlertType>,
    #[sea_orm(column_name = "environmentTypes")]
    pub environment_types: Vec<RuntimeEnvironmentType>,
    #[sea_orm(column_name = "errorAlertConfig", column_type = "JsonBinary")]
    pub error_alert_config: Option<Json>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization_integration::Entity",
        from = "Column::IntegrationId",
        to = "super::organization_integration::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "ProjectAlertChannel_integrationId_fkey"
    )]
    Integration,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "ProjectAlertChannel_projectId_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::project_alert::Entity")]
    Alerts,
    #[sea_orm(has_many = "super::project_alert_storage::Entity")]
    AlertStorages,
}

impl Related<super::organization_integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Integration.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::project_alert::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Alerts.def()
    }
}

impl Related<super::project_alert_storage::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AlertStorages.def()
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
