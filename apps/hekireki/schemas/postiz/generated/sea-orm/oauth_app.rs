use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthApp")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "pictureId", column_type = "Text")]
    pub picture_id: Option<String>,
    #[sea_orm(column_name = "redirectUrl", column_type = "Text")]
    pub redirect_url: String,
    #[sea_orm(column_name = "redirectUris", column_type = "Text")]
    pub redirect_uris: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "clientId", column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_name = "clientSecret", column_type = "Text")]
    pub client_secret: Option<String>,
    #[sea_orm(default_value = false)]
    pub dynamic: bool,
    #[sea_orm(column_name = "tokenEndpointAuthMethod", column_type = "Text")]
    pub token_endpoint_auth_method: Option<String>,
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
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "OAuthApp_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::media::Entity",
        from = "Column::PictureId",
        to = "super::media::Column::Id",
        fk_name = "OAuthApp_pictureId_fkey"
    )]
    Picture,
    #[sea_orm(has_many = "super::oauth_authorization::Entity")]
    Authorizations,
    #[sea_orm(has_many = "super::oauth_self_hosted_authorization::Entity")]
    SelfHostedAuthorizations,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::media::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Picture.def()
    }
}

impl Related<super::oauth_authorization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Authorizations.def()
    }
}

impl Related<super::oauth_self_hosted_authorization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SelfHostedAuthorizations.def()
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
