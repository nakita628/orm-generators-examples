use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Admin/platform-curated catalog row for upstream OTLP normalization.
/// Owns the OTTL transform that reshapes a specific upstream tool's
/// span shape into canonical gen_ai conventions. Users connect via an
/// ingestion key (an ingest-only ApiKey carrying ingestionTemplateId) to
/// land traces in their personal project.
///
/// `organizationId` NULL = platform-published default (always-visible
/// across every org). NOT NULL = org-authored (visible only within
/// that org). v1 ships only platform-published; org-authored authoring
/// UI lands v2 (the column ships now to avoid a future migration).
///
/// Spec: specs/ai-gateway/governance/ingestion-templates-catalog.feature
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "IngestionTemplate")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: Option<String>,
    /// Stable slug - unique within (organizationId, slug) scope. Generated
    /// from displayName + nanoid(6) on create for org-authored; canonical
    /// values for platform-published (`claude_code` / `cursor` /
    /// `claude_cowork` / `raw_otlp_advanced`).
    pub slug: String,
    /// Mirrors the slug for v1 platform-published rows. Kept as a separate
    /// column so a future v2 may publish two templates that share an
    /// upstream sourceType (e.g. claude_code-personal vs claude_code-team)
    /// without a migration. String not enum - extensibility.
    #[sea_orm(column_name = "sourceType")]
    pub source_type: String,
    #[sea_orm(column_name = "displayName")]
    pub display_name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    /// Icon source - single column, prefix-discriminated, mirrors
    /// AiToolEntry.iconAsset:
    /// `preset:<kind>`                 (built-in slot)
    /// `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
    /// null                            → fall back to slug-default in UI
    #[sea_orm(column_name = "iconAsset", column_type = "Text")]
    pub icon_asset: Option<String>,
    /// Discriminator for what (if any) extra credential metadata the
    /// install drawer captures from the user:
    /// null              → otlp_token-only (drawer auto-issues the
    /// ingestion key, no input form)
    /// `static_api_key`  → drawer captures the user's upstream tool API key
    /// `agent_id`        → drawer captures an agent identifier (e.g.
    /// Cursor agentId)
    /// String not enum so v2 shapes (oauth_redirect / s3_bucket) ship
    /// without a migration.
    #[sea_orm(column_name = "credentialSchema")]
    pub credential_schema: Option<String>,
    /// Admin/platform-authored OTTL statements. Applied at receive time
    /// AFTER the receiver re-stamps the binding's authoritative principal
    /// + provenance keys (see `protectedTemplateAttributeKeys` 19-key
    /// closed list). Templates write canonical OTTL outputs (gen_ai.*,
    /// cost.usd, response.model) - they CANNOT write attribution or
    /// provenance keys; the receiver guard restores those post-OTTL.
    #[sea_orm(column_name = "ottlRules", column_type = "Text", default_value = "")]
    pub ottl_rules: String,
    /// True for the v1 platform-published default templates
    /// (organizationId IS NULL). Future org-authored rows write false.
    /// Drives the admin UI "READ-ONLY view + Request a custom template"
    /// surface - admins consume + opt-out only in v1.
    #[sea_orm(column_name = "platformPublished", default_value = false)]
    pub platform_published: bool,
    /// Admin/platform enable. False = hidden from /me catalog. Existing
    /// ingestion keys continue to land traces; new installs are blocked.
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    #[sea_orm(column_name = "archivedAt")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: Option<String>,
    #[sea_orm(column_name = "updatedById")]
    pub updated_by_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade"
    )]
    Organization,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
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
