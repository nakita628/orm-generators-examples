use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A pull request discovered through the organization's GitHub connection and
/// mapped to the branch it was opened from, so a coding-agent session that
/// knows only its branch can name the pull request it produced.
/// Spec: specs/integrations/github-connection.feature.
///
/// The stored title, state, draft and timestamp columns are SNAPSHOTS taken
/// when the mapping was last refreshed. They are fallback labels only: the
/// frontend reads live status from GitHub, so a stale row never presents
/// itself as current truth.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "GithubPullRequest")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// The GitHub host the repository lives on, e.g. "github.com".
    #[sea_orm(column_name = "repositoryHost")]
    pub repository_host: String,
    /// "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
    /// session reported.
    #[sea_orm(column_name = "repositoryFullName")]
    pub repository_full_name: String,
    /// The branch the pull request was opened from.
    #[sea_orm(column_name = "headBranch")]
    pub head_branch: String,
    #[sea_orm(column_name = "prNumber")]
    pub pr_number: i32,
    #[sea_orm(column_name = "htmlUrl")]
    pub html_url: String,
    pub title: String,
    /// GitHub's own state string: "open" or "closed".
    pub state: String,
    #[sea_orm(column_name = "isDraft", default_value = false)]
    pub is_draft: bool,
    #[sea_orm(column_name = "authorLogin")]
    pub author_login: Option<String>,
    #[sea_orm(column_name = "prCreatedAt")]
    pub pr_created_at: DateTime,
    #[sea_orm(column_name = "prClosedAt")]
    pub pr_closed_at: Option<DateTime>,
    #[sea_orm(column_name = "prMergedAt")]
    pub pr_merged_at: Option<DateTime>,
    /// GitHub's own `updated_at` for the snapshot this row holds, and the
    /// freshness guard on every write to it: a snapshot older than the stored one
    /// is skipped. GitHub may deliver webhooks out of order, and without the
    /// guard a late `opened` lands after `closed` and reopens a merged pull
    /// request. NULL means the row predates the column, so any write is accepted.
    #[sea_orm(column_name = "prUpdatedAt")]
    pub pr_updated_at: Option<DateTime>,
    /// When this row was first mapped to its branch.
    #[sea_orm(column_name = "mappedAt")]
    pub mapped_at: DateTime,
    /// When the snapshot columns were last refreshed from GitHub.
    #[sea_orm(column_name = "lastCheckedAt")]
    pub last_checked_at: DateTime,
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
        if insert && self.mapped_at.is_not_set() {
            self.mapped_at = Set(now.naive_utc());
        }
        if insert && self.last_checked_at.is_not_set() {
            self.last_checked_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
