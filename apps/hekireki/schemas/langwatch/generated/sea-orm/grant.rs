use super::grant_principal_type::GrantPrincipalType;
use super::grant_scope_type::GrantScopeType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Grant")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "principalType")]
    pub principal_type: GrantPrincipalType,
    #[sea_orm(column_name = "principalId", column_type = "Text")]
    pub principal_id: Option<String>,
    #[sea_orm(column_name = "roleKey", column_type = "Text")]
    pub role_key: Option<String>,
    #[sea_orm(column_name = "legacyRole", column_type = "Text")]
    pub legacy_role: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub source: String,
    #[sea_orm(column_name = "scopeType")]
    pub scope_type: GrantScopeType,
    #[sea_orm(column_name = "scopeId", column_type = "Text")]
    pub scope_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub token: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub permission: Option<String>,
    #[sea_orm(column_name = "resourceKind", column_type = "Text")]
    pub resource_kind: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_name = "createdByUserId", column_type = "Text")]
    pub created_by_user_id: Option<String>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_name = "maxViews")]
    pub max_views: Option<i32>,
    #[sea_orm(column_type = "JsonBinary")]
    pub condition: Option<Json>,
    #[sea_orm(column_name = "replacesGrantId", column_type = "Text")]
    pub replaces_grant_id: Option<String>,
    #[sea_orm(column_name = "occurredAt", column_type = "custom(\"timestamp(3)\")")]
    pub occurred_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "revokedAt", column_type = "custom(\"timestamp(3)\")")]
    pub revoked_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedReason", column_type = "Text")]
    pub revoked_reason: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
