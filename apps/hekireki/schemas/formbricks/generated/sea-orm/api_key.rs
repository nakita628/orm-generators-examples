use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents enhanced API authentication keys with organization-level ownership.
/// Used for authenticating API requests to Formbricks with more granular permissions.
///
/// @property id - Unique identifier for the API key
/// @property label - Optional descriptive name for the key
/// @property hashedKey - Securely stored API key
/// @property organization - The organization this key belongs to
/// @property createdBy - User ID who created this key
/// @property lastUsedAt - Timestamp of last usage
/// @property apiKeyWorkspaces - Workspaces this key has access to
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ApiKey")]
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
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(column_name = "lastUsedAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub label: String,
    #[sea_orm(column_name = "hashedKey", column_type = "Text")]
    pub hashed_key: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "lookupHash", column_type = "Text")]
    pub lookup_hash: Option<String>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// [OrganizationAccess]
    #[sea_orm(
        column_name = "organizationAccess",
        column_type = "JsonBinary",
        default_value = "{}"
    )]
    pub organization_access: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "ApiKey_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::api_key_workspace::Entity")]
    ApiKeyWorkspaces,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::api_key_workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ApiKeyWorkspaces.def()
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
