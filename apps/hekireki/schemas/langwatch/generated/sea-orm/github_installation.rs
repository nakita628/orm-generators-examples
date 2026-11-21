use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The organization's GitHub connection: a GitHub App installation mapped to a
/// LangWatch organization. Short-lived (1h) installation access tokens are
/// minted against these rows on demand, by Langy for writes (bot-authored pull
/// requests, scoped to the installation's repositories) and by pull-request
/// linkage for reads. No user token and no secret are stored here, the crown
/// jewel is the App private key, held only in the control-plane env.
/// Spec: specs/integrations/github-connection.feature.
///
/// installationId is globally unique (one GitHub installation, one row); a
/// single LangWatch organization may own MANY installations (different GitHub
/// accounts), so the org to installation relationship is one-to-many.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "GithubInstallation")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// GitHub installation id (numeric; stored as string). Globally unique.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "installationId")]
    pub installation_id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// The GitHub account (organization or user) the App is installed on.
    #[sea_orm(column_name = "accountLogin")]
    pub account_login: String,
    /// "Organization" or "User".
    #[sea_orm(column_name = "accountType")]
    pub account_type: String,
    /// GitHub numeric account id (stable across renames).
    #[sea_orm(column_name = "accountId")]
    pub account_id: String,
    /// "all" or "selected" — whether the App can see every repo on the account
    /// or only a chosen subset.
    #[sea_orm(column_name = "repositorySelection")]
    pub repository_selection: String,
    /// Cached repository list [{ id, fullName }] when repositorySelection =
    /// "selected". Null when "all" (the full set is resolved live at mint time).
    pub repositories: Option<Json>,
    /// Set when GitHub suspends the installation; cleared on unsuspend. A
    /// suspended installation cannot mint tokens.
    #[sea_orm(column_name = "suspendedAt")]
    pub suspended_at: Option<DateTime>,
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
