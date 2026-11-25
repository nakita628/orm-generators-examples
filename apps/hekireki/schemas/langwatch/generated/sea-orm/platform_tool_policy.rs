use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Per-(organization, tool) override for which `langwatch <tool>` paths the
/// CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
/// `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
/// row resolves to the hardcoded defaults (both true for claude/codex/gemini/
/// opencode; cursor allows the gateway path only), so the policy is purely
/// additive: disabling a path is an explicit admin choice. The CLI caches the
/// resolved map at login and gates path selection on it.
/// See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PlatformToolPolicy")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "toolSlug", column_type = "Text")]
    pub tool_slug: String,
    #[sea_orm(column_name = "allowVk", default_value = true)]
    pub allow_vk: bool,
    #[sea_orm(column_name = "allowOtelDirect", default_value = true)]
    pub allow_otel_direct: bool,
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
