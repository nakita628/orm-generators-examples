use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ApiKey")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "lookupId", column_type = "Text")]
    pub lookup_id: String,
    #[sea_orm(column_name = "hashedSecret", column_type = "Text")]
    pub hashed_secret: String,
    #[sea_orm(
        column_name = "permissionMode",
        column_type = "Text",
        default_value = "all"
    )]
    pub permission_mode: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "createdByUserId", column_type = "Text")]
    pub created_by_user_id: Option<String>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_name = "lastUsedAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedAt", column_type = "custom(\"timestamp(3)\")")]
    pub revoked_at: Option<DateTime>,
    /// Why the key was revoked, set with `revokedAt`: "user" when a person
    /// revoked it (the API-keys page, the REST API, the devices tab), "rotation"
    /// when a mint replaced it, "session" when the CLI login key it was parented
    /// to was revoked, "expired" when that session ran out, "offboarded" when the
    /// person's membership of the organization ended. "cap" survives on rows the
    /// retired per-tool cap revoked. The CLI reads it to decide whether a dead
    /// ingest key may be re-minted: a key a person revoked stays dead.
    /// NULL for live keys and for keys revoked before the cause was recorded.
    #[sea_orm(column_name = "revocationCause", column_type = "Text")]
    pub revocation_cause: Option<String>,
    /// The CLI login key of the device session that minted this ingestion key.
    /// Set by the CLI's personal mint, NULL for every other key. Revoking the
    /// parent (logout, the devices tab, a re-login from the same device, session
    /// expiry) revokes the children with it. Column ref only, no FK and no
    /// index: lookups filter by the indexed `userId` first and match the parent
    /// in memory over a person's few live ingest keys.
    #[sea_orm(column_name = "parentApiKeyId", column_type = "Text")]
    pub parent_api_key_id: Option<String>,
    /// Non-null marks this key as an "ingestion key": a project-scoped,
    /// ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
    /// installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
    /// opencode / claude_cowork) is stamped as `langwatch.source` provenance on
    /// every span the key authorizes. NULL for normal API / service keys.
    #[sea_orm(column_name = "ingestSourceType", column_type = "Text")]
    pub ingest_source_type: Option<String>,
    /// Set only for ingestion keys derived from an admin IngestionTemplate
    /// (e.g. claude_cowork). Drives receiver OTTL application + the
    /// `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
    /// (which have no template) and for normal keys. Column ref only (no FK,
    /// per the app-layer-integrity convention); resolved by the ingest-key service.
    #[sea_orm(column_name = "ingestionTemplateId", column_type = "Text")]
    pub ingestion_template_id: Option<String>,
    /// Human label of the CLI device session that minted this ingestion key
    /// (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
    /// flow's `client_info`. Denormalized so the org-wide API-keys settings page
    /// can show which device a key came from without resolving another user's
    /// Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
    /// a CLI old enough to predate device metadata.
    #[sea_orm(column_name = "createdByDeviceLabel", column_type = "Text")]
    pub created_by_device_label: Option<String>,
    /// True for a key LangWatch minted for itself (a run, the gateway's trace
    /// export, a Langy session). Hides it from listings, refuses customer edits
    /// and revokes, and lets an ownerless run key act as the system. Never set
    /// by a customer create; rows before the column stay false.
    #[sea_orm(column_name = "isSystemManaged", default_value = false)]
    pub is_system_managed: bool,
    /// The browser session that minted a docs session key (ADR-185), so signing
    /// out of that session revokes it. NULL for every other key. Unindexed:
    /// sign-out filters by the indexed `userId` first.
    #[sea_orm(column_name = "createdByBrowserSessionId", column_type = "Text")]
    pub created_by_browser_session_id: Option<String>,
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
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Restrict",
        fk_name = "ApiKey_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Restrict",
        fk_name = "ApiKey_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::role_binding::Entity")]
    RoleBindings,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::role_binding::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RoleBindings.def()
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
