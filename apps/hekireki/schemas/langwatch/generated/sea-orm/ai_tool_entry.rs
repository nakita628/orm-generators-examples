use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Org-scoped catalog of AI tools surfaced on the personal home portal
/// (/me) and managed by org admins at /governance/tool-catalog.
///
/// Three first-class tile types via `type` discriminator:
/// - "coding_assistant" - click expands to setup helper (e.g. `langwatch claude`)
/// - "model_provider"   - click expands to inline personal-VK creation
/// - "external_tool"    - click expands to admin-attached markdown + external link
///
/// Scoping: organization-default entries can be OVERRIDDEN at the team
/// tier (same `slug` + `scope=team` + matching `scopeId`) so admins can
/// hide or re-skin a globally-listed tool for a specific team. The
/// user-facing `list` endpoint resolves this hierarchy at read time.
///
/// `config` is a discriminated JSON union per `type`:
/// coding_assistant: { setupCommand, setupDocsUrl, helperText? }
/// model_provider:   { providerKey, suggestedRoutingPolicyId?, defaultLabel?, projectSuggestionText? }
/// external_tool:    { descriptionMarkdown, linkUrl, ctaLabel? }
///
/// Spec: specs/ai-governance/personal-portal/tool-catalog-*.feature
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AiToolEntry")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// "organization" | "team". String not enum so new scopes ship
    /// without migrations.
    pub scope: String,
    /// organizationId or teamId, narrowing what scope this entry applies
    /// to. Always equals organizationId when scope=organization.
    #[sea_orm(column_name = "scopeId")]
    pub scope_id: String,
    /// "coding_assistant" | "model_provider" | "external_tool".
    /// Discriminator drives the per-tile UX shape and validates
    /// the `config` JSON column at the app layer.
    pub r#type: String,
    #[sea_orm(column_name = "displayName")]
    pub display_name: String,
    /// Stable slug for team-overrides-org resolution (a team-scoped
    /// entry with the same slug as an org-scoped one hides the org
    /// default for users in that team). Server-owned: generated from
    /// displayName + nanoid(6) on create, immutable on update.
    pub slug: String,
    /// @deprecated Use `iconAsset` instead. Kept for back-compat with
    /// pre-refactor rows; new writes should populate iconAsset.
    #[sea_orm(column_name = "iconKey")]
    pub icon_key: Option<String>,
    /// Icon source - single column, prefix-discriminated:
    /// "preset:claude_code" / "preset:codex" / ...   (built-in)
    /// "data:image/svg+xml;base64,..."               (admin-uploaded)
    /// null                                          → fall back to
    /// type-default in UI
    #[sea_orm(column_name = "iconAsset", column_type = "Text")]
    pub icon_asset: Option<String>,
    /// Admin-curated sort within the user-facing portal grid.
    #[sea_orm(default_value = 0)]
    pub order: i32,
    /// When false the entry is hidden from the user-facing list but
    /// remains in the admin list (so admins can re-enable without
    /// losing per-entry config). Disabled is "soft-hide", archived is
    /// "soft-delete" (preserves audit trail + scope binding history).
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    /// Per-type config JSON. Validated via zod discriminated union
    /// at the service layer (see AiToolEntryService).
    #[sea_orm(default_value = "{}")]
    pub config: Json,
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
    #[sea_orm(has_many = "super::ai_tool_entry_team::Entity")]
    Teams,
    #[sea_orm(has_many = "super::ai_tool_entry_department::Entity")]
    Departments,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::ai_tool_entry_team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Teams.def()
    }
}

impl Related<super::ai_tool_entry_department::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Departments.def()
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
