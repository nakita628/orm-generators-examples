use super::mcp_grant_access::McpGrantAccess;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
/// its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
/// a token is now bound to the resource it was approved for instead of to a static allow-list.
/// `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
/// What one user allows one MCP client to do (GLOSSARY.md "Grant", ADR-0041): its organization & workspace
/// access and an access level per permission. One per OAuth connection, keyed by user and client rather
/// than by `oauthConsent` row, because the provider can write duplicate consent rows for one connection
/// and settles them later. Every MCP request reads it; without one the request is refused. The client's
/// OAuth scopes stay on the consent and cap it.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "McpGrant")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    #[sea_orm(column_name = "clientId")]
    pub client_id: String,
    /// Which organizations and workspaces the grant reaches (ENG-7153): All access, only the organizations
    /// in `organizations` (and every workspace in them, later ones included), or only the workspaces in
    /// `workspaces`. Enforced in SpiceDB through `can()` (ADR-0042); a value no row backs reaches nothing.
    #[sea_orm(default_value = "all")]
    pub access: McpGrantAccess,
    /// The access level per permission of `apps/web/modules/mcp/permissions.ts`: `{ "surveys": "write" }`.
    /// A permission that is absent has No access.
    #[sea_orm(default_value = "{}")]
    pub permissions: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::ClientId",
        to = "super::oauth_client::Column::ClientId",
        on_delete = "Cascade"
    )]
    Client,
    #[sea_orm(has_many = "super::mcp_grant_organization::Entity")]
    Organizations,
    #[sea_orm(has_many = "super::mcp_grant_workspace::Entity")]
    Workspaces,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::oauth_client::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Client.def()
    }
}

impl Related<super::mcp_grant_organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organizations.def()
    }
}

impl Related<super::mcp_grant_workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspaces.def()
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
