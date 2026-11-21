use super::organization_user_role::OrganizationUserRole;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganizationUser")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    pub role: OrganizationUserRole,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    /// Department assignment for spend attribution (the people lens).
    /// Single-valued, org-scoped, accounting-only (never gates access).
    /// Plain column ref (relationMode="prisma"); the service validates the
    /// id belongs to this org. See departments.feature.
    #[sea_orm(column_name = "departmentId")]
    pub department_id: Option<String>,
    /// Set when an admin disables this membership to stay within the licensed
    /// seat count. The person loses access to this organization only, keeps
    /// their role, department and everything they did, and frees their seat.
    /// Reversible: re-enabling restores access if a seat is available.
    /// Seat counting excludes disabled memberships. See
    /// seat-reconciliation.feature.
    #[sea_orm(column_name = "disabledAt")]
    pub disabled_at: Option<DateTime>,
    /// Stable grant intent for an unfinished automatic SSO admission.
    /// Cleared after the authoritative grant is observed; never inferred from missing access.
    #[sea_orm(column_name = "pendingSsoGrantId")]
    pub pending_sso_grant_id: Option<String>,
    /// A new generation for each membership lifetime. Stamped USER grant facts
    /// must match the current active membership before they can be projected;
    /// removing and rejoining therefore cannot revive a delayed old attach.
    #[sea_orm(column_name = "membershipStamp")]
    pub membership_stamp: String,
    /// When this membership started holding the full seat it holds now; null means createdAt.
    /// Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
    #[sea_orm(column_name = "fullSeatSince")]
    pub full_seat_since: Option<DateTime>,
    /// Set when admission wrote this membership disabled because no seat was free: the person
    /// is waiting for one. Cleared when the membership is enabled.
    #[sea_orm(column_name = "seatPendingSince")]
    pub seat_pending_since: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id"
    )]
    Organization,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
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
