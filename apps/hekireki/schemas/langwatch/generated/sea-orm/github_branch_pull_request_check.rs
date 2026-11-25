use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The per-branch bookkeeping behind pull-request mapping: when a branch was
/// last looked up on GitHub, how many pull requests it resolved to, and when it
/// may be asked about again. A branch with no pull request is the common case,
/// so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
/// costing a GitHub call on every read.
/// Spec: specs/integrations/github-connection.feature.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "GithubBranchPullRequestCheck")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "repositoryHost", column_type = "Text")]
    pub repository_host: String,
    /// "owner/repo", stored LOWERCASED, matching GithubPullRequest.
    #[sea_orm(column_name = "repositoryFullName", column_type = "Text")]
    pub repository_full_name: String,
    #[sea_orm(column_name = "headBranch", column_type = "Text")]
    pub head_branch: String,
    #[sea_orm(
        column_name = "lastCheckedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_checked_at: DateTime,
    #[sea_orm(column_name = "prCount", default_value = 0)]
    pub pr_count: i32,
    /// Set while the branch resolves to no pull request; cleared once one is
    /// found.
    #[sea_orm(column_name = "notFoundAt", column_type = "custom(\"timestamp(3)\")")]
    pub not_found_at: Option<DateTime>,
    /// The earliest instant GitHub may be asked about this branch again.
    #[sea_orm(column_name = "recheckAfter", column_type = "custom(\"timestamp(3)\")")]
    pub recheck_after: Option<DateTime>,
    /// Consecutive lookups that found nothing, backing off the recheck.
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    /// When a reader last asked about this branch, whether or not GitHub was
    /// called.
    #[sea_orm(
        column_name = "lastRequestedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_requested_at: DateTime,
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
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "GithubBranchPullRequestCheck_organizationId_fkey"
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
