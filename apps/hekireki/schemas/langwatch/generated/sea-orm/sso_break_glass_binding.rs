use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A way back in (D05 - see specs/identity/sso-onboarding-tiers.feature).
///
/// One named person who may still sign in without the identity provider
/// while their organization's sign-in belongs to one, and the date on which
/// that stops. Activation asks whether such a person exists, because the
/// failure it prevents is the only one that cannot be recovered from inside
/// the product: a misconfigured connection goes live, every member is handed
/// to an identity provider that will not authenticate them, and nobody is
/// left who can turn it off.
///
/// Rows are immutable except for `supersededAt` and `warnedDays`. A renewal
/// writes a NEW row naming the one it replaced, so the date a way in
/// previously ended is still readable in the history - which is what makes a
/// renewal auditable rather than a column somebody moved.
///
/// No relation to Organization or User, matching the identity heads above:
/// the binding is the record of who could get in, and a cascade would erase
/// exactly the history a lockout post-mortem reads.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SsoBreakGlassBinding")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// Who holds the way in.
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    /// Who granted it. A way back in is never self-served.
    #[sea_orm(column_name = "grantedByUserId")]
    pub granted_by_user_id: String,
    #[sea_orm(column_name = "grantedAt")]
    pub granted_at: DateTime,
    #[sea_orm(column_name = "expiresAt")]
    pub expires_at: DateTime,
    /// Set when a renewal replaced this row; the row itself is never edited.
    #[sea_orm(column_name = "supersededAt")]
    pub superseded_at: Option<DateTime>,
    /// The binding this one renewed, so the previous end date stays readable.
    #[sea_orm(column_name = "renewedFromId")]
    pub renewed_from_id: Option<String>,
    /// Which of the fourteen, seven and one day warnings have been sent, so a
    /// sweep that runs twice in a day never says the same thing twice.
    #[sea_orm(column_name = "warnedDays")]
    pub warned_days: Vec<i32>,
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
