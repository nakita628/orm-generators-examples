use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A progressive check-in SESSION, mirroring PartialBookingCheckout.
///
/// `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
/// QUANTITY_TRACKED only once its remaining units hit zero (partial returns
/// live in `ConsumptionLog`). Completing a booking outright writes no row here,
/// so the same warning applies — absence is not evidence of anything.
/// `BookingAsset.checkedInAt` is the per-slice fact.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PartialBookingCheckin")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "assetIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub asset_ids: Vec<String>,
    #[sea_orm(column_name = "checkinCount")]
    pub checkin_count: i32,
    #[sea_orm(
        column_name = "checkinTimestamp",
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub checkin_timestamp: DateTimeWithTimeZone,
    #[sea_orm(column_name = "bookingId", column_type = "Text")]
    pub booking_id: String,
    #[sea_orm(column_name = "checkedInById", column_type = "Text")]
    pub checked_in_by_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::booking::Entity",
        from = "Column::BookingId",
        to = "super::booking::Column::Id",
        on_delete = "Cascade",
        fk_name = "PartialBookingCheckin_bookingId_fkey"
    )]
    Booking,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CheckedInById",
        to = "super::user::Column::Id",
        fk_name = "PartialBookingCheckin_checkedInById_fkey"
    )]
    CheckedInBy,
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Booking.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CheckedInBy.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.checkin_timestamp.is_not_set() {
            self.checkin_timestamp = Set(now.fixed_offset());
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
