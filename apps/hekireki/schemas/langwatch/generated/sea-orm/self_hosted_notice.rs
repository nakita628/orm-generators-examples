use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
/// report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
/// tenant: every install receives every published notice that has not expired.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SelfHostedNotice")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// `info`, `warning` or `security`.
    pub severity: String,
    pub title: String,
    pub body: String,
    /// An https link to read more; null where the notice names none.
    pub url: Option<String>,
    /// The release that fixes what the notice is about; an install on it or newer hides it.
    #[sea_orm(column_name = "fixedIn")]
    pub fixed_in: Option<String>,
    /// `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent.
    pub status: String,
    #[sea_orm(column_name = "expiresAt")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_name = "publishedAt")]
    pub published_at: Option<DateTime>,
    /// The staff member who published it, by user id.
    #[sea_orm(column_name = "publishedBy")]
    pub published_by: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
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
