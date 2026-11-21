use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A progressive checkout SESSION: one scan batch, covering N assets via the
/// positional `assetIds` / `quantities` / `bookingAssetIds` arrays.
///
/// Records sessions and the units they claimed — NOT whether an asset is out.
/// The all-at-once checkout (`checkoutBooking`, the Check out button) flips
/// statuses and writes NO row here, so **absence proves nothing** and must never
/// be read as "was not checked out". Reading it that way is what let a single
/// later scan mark every button-checked-out asset on a booking as
/// "never checked out". `BookingAsset.checkedOutAt` answers that question.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PartialBookingCheckout")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "assetIds")]
    pub asset_ids: Vec<String>,
    pub quantities: Vec<i32>,
    #[sea_orm(column_name = "bookingAssetIds")]
    pub booking_asset_ids: Vec<String>,
    #[sea_orm(column_name = "checkoutCount")]
    pub checkout_count: i32,
    #[sea_orm(
        column_name = "checkoutTimestamp",
        column_type = "TimestampWithTimeZone"
    )]
    pub checkout_timestamp: DateTimeWithTimeZone,
    #[sea_orm(column_name = "bookingId")]
    pub booking_id: String,
    #[sea_orm(column_name = "checkedOutById")]
    pub checked_out_by_id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::booking::Entity",
        from = "Column::BookingId",
        to = "super::booking::Column::Id",
        on_delete = "Cascade"
    )]
    Booking,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CheckedOutById",
        to = "super::user::Column::Id"
    )]
    CheckedOutBy,
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Booking.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CheckedOutBy.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.checkout_timestamp.is_not_set() {
            self.checkout_timestamp = Set(now.fixed_offset());
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
