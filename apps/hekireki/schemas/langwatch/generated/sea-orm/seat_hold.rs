use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A full seat still counted after the person who took it stopped holding it (removed,
/// disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
/// One per organization and person; a return within the hold takes the seat back.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SeatHold")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    /// Display copy for the Users list, kept because a removed person has no membership row.
    #[sea_orm(column_type = "Text")]
    pub email: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    /// When the person took the seat this hold keeps.
    #[sea_orm(column_name = "seatTakenAt", column_type = "custom(\"timestamp(3)\")")]
    pub seat_taken_at: DateTime,
    /// seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly.
    #[sea_orm(column_name = "heldUntil", column_type = "custom(\"timestamp(3)\")")]
    pub held_until: DateTime,
    /// REMOVED | DISABLED | DEMOTED
    #[sea_orm(column_type = "Text")]
    pub reason: String,
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
