use super::integration_service::IntegrationService;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganizationIntegration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    pub service: IntegrationService,
    /// Identifier for external, integration's organization (e.g. Vercel's team)
    #[sea_orm(column_name = "externalOrganizationId", column_type = "Text")]
    pub external_organization_id: Option<String>,
    #[sea_orm(column_name = "integrationData", column_type = "JsonBinary")]
    pub integration_data: Json,
    #[sea_orm(column_name = "tokenReferenceId", column_type = "Text")]
    pub token_reference_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
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
        belongs_to = "super::secret_reference::Entity",
        from = "Column::TokenReferenceId",
        to = "super::secret_reference::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "OrganizationIntegration_tokenReferenceId_fkey"
    )]
    TokenReference,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "OrganizationIntegration_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::project_alert_channel::Entity")]
    AlertChannels,
    #[sea_orm(has_many = "super::organization_project_integration::Entity")]
    OrganizationProjectIntegration,
}

impl Related<super::secret_reference::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TokenReference.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::project_alert_channel::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AlertChannels.def()
    }
}

impl Related<super::organization_project_integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrganizationProjectIntegration.def()
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
