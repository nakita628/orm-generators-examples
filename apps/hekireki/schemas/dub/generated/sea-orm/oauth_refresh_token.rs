use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthRefreshToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(
        column_name = "installationId",
        column_type = "String(StringLen::N(191))"
    )]
    pub installation_id: String,
    #[sea_orm(
        column_name = "accessTokenId",
        column_type = "String(StringLen::N(191))"
    )]
    pub access_token_id: String,
    #[sea_orm(unique)]
    #[sea_orm(
        column_name = "hashedRefreshToken",
        column_type = "String(StringLen::N(191))"
    )]
    pub hashed_refresh_token: String,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"datetime(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::restricted_token::Entity",
        from = "Column::AccessTokenId",
        to = "super::restricted_token::Column::Id",
        on_delete = "Cascade",
        fk_name = "OAuthRefreshToken_accessTokenId_fkey"
    )]
    AccessToken,
    #[sea_orm(
        belongs_to = "super::installed_integration::Entity",
        from = "Column::InstallationId",
        to = "super::installed_integration::Column::Id",
        on_delete = "Cascade",
        fk_name = "OAuthRefreshToken_installationId_fkey"
    )]
    InstalledIntegration,
}

impl Related<super::restricted_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AccessToken.def()
    }
}

impl Related<super::installed_integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InstalledIntegration.def()
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
        Ok(self)
    }
}
