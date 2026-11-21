use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The dated match table: "`user_abc123` on Anthropic is platform user #42,
/// valid from March 3", closed when it stops being true.
///
/// Dating is what makes a re-issued email address survivable: the leaver's link
/// closes at offboarding, the new hire gets a new link, and January's spend
/// stays with January's person. A correction closes the old row and opens a new
/// one; nothing is ever silently rewritten.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "IdentityMatch")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// The provider-side identity this link is about.
    #[sea_orm(column_name = "discoveredPersonId")]
    pub discovered_person_id: String,
    /// The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
    /// the row and its dates remain, so the history of the link is not rewritten
    /// by the erasure either.
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<String>,
    /// directory_id | verified_email | human_confirmed. Exact verified-email
    /// equality is the primary evidence; a directory id corroborates but never
    /// stands alone (ADR-128 §12).
    #[sea_orm(column_name = "evidenceKind")]
    pub evidence_kind: String,
    #[sea_orm(column_name = "validFrom")]
    pub valid_from: DateTime,
    /// Null means the link is open. Offboarding and corrections close it.
    #[sea_orm(column_name = "validTo")]
    pub valid_to: Option<DateTime>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
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
        Ok(self)
    }
}
