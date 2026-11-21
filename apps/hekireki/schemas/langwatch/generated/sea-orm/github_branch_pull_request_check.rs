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
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "repositoryHost")]
    pub repository_host: String,
    /// "owner/repo", stored LOWERCASED, matching GithubPullRequest.
    #[sea_orm(column_name = "repositoryFullName")]
    pub repository_full_name: String,
    #[sea_orm(column_name = "headBranch")]
    pub head_branch: String,
    #[sea_orm(column_name = "lastCheckedAt")]
    pub last_checked_at: DateTime,
    #[sea_orm(column_name = "prCount", default_value = 0)]
    pub pr_count: i32,
    /// Set while the branch resolves to no pull request; cleared once one is
    /// found.
    #[sea_orm(column_name = "notFoundAt")]
    pub not_found_at: Option<DateTime>,
    /// The earliest instant GitHub may be asked about this branch again.
    #[sea_orm(column_name = "recheckAfter")]
    pub recheck_after: Option<DateTime>,
    /// Consecutive lookups that found nothing, backing off the recheck.
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    /// When a reader last asked about this branch, whether or not GitHub was
    /// called.
    #[sea_orm(column_name = "lastRequestedAt")]
    pub last_requested_at: DateTime,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
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
