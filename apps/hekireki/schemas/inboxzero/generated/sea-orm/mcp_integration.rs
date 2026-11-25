use super::mcp_auth_type::McpAuthType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "McpIntegration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "registeredServerUrl", column_type = "Text")]
    pub registered_server_url: Option<String>,
    #[sea_orm(column_name = "registeredAuthorizationUrl", column_type = "Text")]
    pub registered_authorization_url: Option<String>,
    #[sea_orm(column_name = "registeredTokenUrl", column_type = "Text")]
    pub registered_token_url: Option<String>,
    #[sea_orm(column_name = "oauthClientId", column_type = "Text")]
    pub oauth_client_id: Option<String>,
    #[sea_orm(column_name = "oauthClientSecret", column_type = "Text")]
    pub oauth_client_secret: Option<String>,
    #[sea_orm(column_name = "displayName", column_type = "Text")]
    pub display_name: Option<String>,
    #[sea_orm(column_name = "serverUrl", column_type = "Text")]
    pub server_url: Option<String>,
    #[sea_orm(column_name = "authType")]
    pub auth_type: Option<McpAuthType>,
    #[sea_orm(column_name = "emailAccountId", column_type = "Text")]
    pub email_account_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::email_account::Entity",
        from = "Column::EmailAccountId",
        to = "super::email_account::Column::Id",
        on_delete = "Cascade",
        fk_name = "McpIntegration_emailAccountId_fkey"
    )]
    EmailAccount,
    #[sea_orm(has_many = "super::mcp_connection::Entity")]
    Connections,
}

impl Related<super::email_account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmailAccount.def()
    }
}

impl Related<super::mcp_connection::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Connections.def()
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
