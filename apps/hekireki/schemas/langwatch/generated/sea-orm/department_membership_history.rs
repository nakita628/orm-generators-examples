use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One dated person → department link (ADR-128 §13, #7882).
///
/// `OrganizationUser.departmentId` stays the CURRENT assignment every screen
/// reads; these rows are its history, appended by the same `assignUser` write
/// so the two cannot drift. They exist for the cost reads: January's spend
/// stays with January's department, so a read resolves against the link that
/// was open on the cost's day, not against today's pointer. Members with no
/// link on a day resolve to "unassigned", never to an error.
///
/// Same shape as IdentityMatch's dated links, same one-open-link guard: the
/// partial unique index in the migration rejects a second open row per
/// (organizationId, userId) with SQLSTATE 23505.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DepartmentMembershipHistory")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "departmentId", column_type = "Text")]
    pub department_id: String,
    #[sea_orm(column_name = "validFrom", column_type = "custom(\"timestamp(3)\")")]
    pub valid_from: DateTime,
    /// Null means the link is open. Reassignment and clearing close it.
    #[sea_orm(column_name = "validTo", column_type = "custom(\"timestamp(3)\")")]
    pub valid_to: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
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
