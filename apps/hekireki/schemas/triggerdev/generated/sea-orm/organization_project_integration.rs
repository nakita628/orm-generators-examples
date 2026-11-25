use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganizationProjectIntegration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationIntegrationId", column_type = "Text")]
    pub organization_integration_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// Identifier for webhooks, for example Vercel's projectId
    #[sea_orm(column_name = "externalEntityId", column_type = "Text")]
    pub external_entity_id: String,
    /// Save useful data like config or external entity name
    #[sea_orm(column_name = "integrationData", column_type = "JsonBinary")]
    pub integration_data: Json,
    /// UserId who installed the integration
    #[sea_orm(column_name = "installedBy", column_type = "Text")]
    pub installed_by: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "deletedAt", column_type = "custom(\"timestamp(3)\")")]
    pub deleted_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization_integration::Entity",
        from = "Column::OrganizationIntegrationId",
        to = "super::organization_integration::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "OrganizationProjectIntegration_organizationIntegrationId_fkey"
    )]
    OrganizationIntegration,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "OrganizationProjectIntegration_projectId_fkey"
    )]
    Project,
}

impl Related<super::organization_integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrganizationIntegration.def()
    }
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
