use super::virtual_key_purpose::VirtualKeyPurpose;
use super::virtual_key_routing_mode::VirtualKeyRoutingMode;
use super::virtual_key_status::VirtualKeyStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "VirtualKey")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(default_value = "ACTIVE")]
    pub status: VirtualKeyStatus,
    #[sea_orm(default_value = "USER")]
    pub purpose: VirtualKeyPurpose,
    #[sea_orm(column_name = "externalId", column_type = "Text")]
    pub external_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub metadata: Json,
    #[sea_orm(column_name = "disabledAt", column_type = "custom(\"timestamp(3)\")")]
    pub disabled_at: Option<DateTime>,
    #[sea_orm(column_name = "disabledReason", column_type = "Text")]
    pub disabled_reason: Option<String>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "hashedSecret", column_type = "Text")]
    pub hashed_secret: String,
    #[sea_orm(column_name = "displayPrefix", column_type = "Text")]
    pub display_prefix: String,
    #[sea_orm(column_name = "principalUserId", column_type = "Text")]
    pub principal_user_id: Option<String>,
    #[sea_orm(column_name = "traceProjectId", column_type = "Text")]
    pub trace_project_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub config: Json,
    #[sea_orm(default_value = "0")]
    pub revision: i64,
    #[sea_orm(column_name = "previousHashedSecret", column_type = "Text")]
    pub previous_hashed_secret: Option<String>,
    #[sea_orm(
        column_name = "previousSecretValidUntil",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub previous_secret_valid_until: Option<DateTime>,
    #[sea_orm(column_name = "revokedAt", column_type = "custom(\"timestamp(3)\")")]
    pub revoked_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedById", column_type = "Text")]
    pub revoked_by_id: Option<String>,
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
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: String,
    #[sea_orm(column_name = "lastUsedAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: Option<DateTime>,
    #[sea_orm(column_name = "routingPolicyId", column_type = "Text")]
    pub routing_policy_id: Option<String>,
    #[sea_orm(column_name = "routingMode", default_value = "NONE")]
    pub routing_mode: VirtualKeyRoutingMode,
    #[sea_orm(
        column_name = "connectServices",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub connect_services: Vec<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "licenseTokenHash", column_type = "Text")]
    pub license_token_hash: Option<String>,
    #[sea_orm(column_name = "licenseInstanceId", column_type = "Text")]
    pub license_instance_id: Option<String>,
    #[sea_orm(
        column_name = "licenseExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub license_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "licenseId", column_type = "Text")]
    pub license_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::PrincipalUserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "VirtualKey_principalUserId_fkey"
    )]
    PrincipalUser,
    #[sea_orm(
        belongs_to = "super::routing_policy::Entity",
        from = "Column::RoutingPolicyId",
        to = "super::routing_policy::Column::Id",
        on_delete = "SetNull",
        fk_name = "VirtualKey_routingPolicyId_fkey"
    )]
    RoutingPolicy,
    #[sea_orm(has_many = "super::virtual_key_scope::Entity")]
    Scopes,
    #[sea_orm(has_many = "super::gateway_budget_ledger::Entity")]
    LedgerEntries,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PrincipalUser.def()
    }
}

impl Related<super::routing_policy::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RoutingPolicy.def()
    }
}

impl Related<super::virtual_key_scope::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scopes.def()
    }
}

impl Related<super::gateway_budget_ledger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LedgerEntries.def()
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
