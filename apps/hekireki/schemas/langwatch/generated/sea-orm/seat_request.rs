use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A person asked for a full seat the organization had no room for (modules/organization/specs/
/// seat-requests.feature). One row per organization and address; reopening reuses the row.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SeatRequest")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// Lowercased address of the person asking.
    pub email: String,
    /// Null when only the address is known (a sign-in the identity provider refused).
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<String>,
    pub name: Option<String>,
    /// OPEN | APPROVED | DECLINED
    pub status: String,
    /// pending | lite | idp_refused: how the person was left when the request opened.
    pub source: String,
    /// The identity provider also refused this person's sign-in.
    #[sea_orm(column_name = "idpRefused", default_value = false)]
    pub idp_refused: bool,
    /// okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml
    #[sea_orm(column_name = "idpKind")]
    pub idp_kind: Option<String>,
    /// When sales were told about this request; reset when it reopens.
    #[sea_orm(column_name = "salesNotifiedAt")]
    pub sales_notified_at: Option<DateTime>,
    #[sea_orm(column_name = "openedAt")]
    pub opened_at: DateTime,
    #[sea_orm(column_name = "resolvedAt")]
    pub resolved_at: Option<DateTime>,
    #[sea_orm(column_name = "resolvedByUserId")]
    pub resolved_by_user_id: Option<String>,
    /// Set when an approval claims the request, cleared once its seat is given; an approval
    /// still pending long after is one whose seat never landed, and is reopened.
    #[sea_orm(column_name = "seatPendingSince")]
    pub seat_pending_since: Option<DateTime>,
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
