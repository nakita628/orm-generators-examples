use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Book-by-Model: intent row attached to a booking for a quantity of
/// assets of a given AssetModel, without yet picking specific assets.
/// Concrete `BookingAsset` rows are only created at
/// scan-to-assign time; when the last unit is assigned, the request
/// row is deleted. `quantity` is the number of units still needing
/// assignment (starts at the full requested total, decrements on scan).
///
/// Keeping model requests separate from `BookingAsset` means the ~40
/// code paths that assume `BookingAsset.assetId` points to a real
/// asset stay untouched — model-only bookings become concrete the
/// moment they're checked out.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BookingModelRequest")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "bookingId", column_type = "Text")]
    pub booking_id: String,
    #[sea_orm(column_name = "assetModelId", column_type = "Text")]
    pub asset_model_id: String,
    /// Total units originally reserved. Mutated only by explicit
    /// operator edits (upsert) — never decremented on materialisation.
    /// This preserves the audit trail of "how many were planned".
    pub quantity: i32,
    /// Units already assigned to concrete `BookingAsset` rows via
    /// `materializeModelRequestForAsset`. Scans increment it.
    ///
    /// It DOES decrease: `removeAssets` (booking service) decrements it by the
    /// number of removed `BookingAsset` rows whose `bookingModelRequestId`
    /// points at this request, floored at 0, and clears `fulfilledAt` when the
    /// result drops below `quantity`.
    ///
    /// Counted from that column, never from a shared `assetModelId`: an asset
    /// added directly alongside a reservation shares the model without having
    /// discharged anything, so grouping by model re-opened reservations that
    /// were still correctly satisfied.
    #[sea_orm(column_name = "fulfilledQuantity", default_value = 0)]
    pub fulfilled_quantity: i32,
    /// Timestamp of the moment the reservation became fully fulfilled — set
    /// when `fulfilledQuantity` catches up to `quantity`, and preserved (not
    /// re-stamped) if an unchanged quantity is saved again. Cleared back to
    /// null when outstanding units reappear: the operator edits `quantity`
    /// upwards, or `removeAssets` drops `fulfilledQuantity` below it.
    ///
    /// The invariant that makes this the primary "is this outstanding?" filter:
    /// non-null IFF `fulfilledQuantity === quantity`, so `fulfilledAt IS NULL`
    /// means units remain to materialise. `fulfilledQuantity` can never exceed
    /// `quantity` (the claim in `materializeModelRequestForAsset` is a single
    /// statement guarded by `WHERE "fulfilledQuantity" < "quantity"`; upsert
    /// rejects a shrink below it), so any decrement re-opens the row.
    #[sea_orm(column_name = "fulfilledAt", column_type = "custom(\"timestamp(3)\")")]
    pub fulfilled_at: Option<DateTime>,
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
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BookingModelRequest_bookingId_fkey"
    )]
    Booking,
    #[sea_orm(
        belongs_to = "super::asset_model::Entity",
        from = "Column::AssetModelId",
        to = "super::asset_model::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BookingModelRequest_assetModelId_fkey"
    )]
    AssetModel,
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Booking.def()
    }
}

impl Related<super::asset_model::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetModel.def()
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
