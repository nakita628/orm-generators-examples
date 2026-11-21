use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// *
/// * Explicit pivot table between Booking and Asset.
/// * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BookingAsset")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "bookingId")]
    pub booking_id: String,
    #[sea_orm(column_name = "assetId")]
    pub asset_id: String,
    #[sea_orm(default_value = 1)]
    pub quantity: i32,
    #[sea_orm(column_name = "bookingModelRequestId")]
    pub booking_model_request_id: Option<String>,
    #[sea_orm(column_name = "assetKitId")]
    pub asset_kit_id: Option<String>,
    #[sea_orm(column_name = "sourceKitId")]
    pub source_kit_id: Option<String>,
    #[sea_orm(column_name = "checkedOutAt", column_type = "TimestampWithTimeZone")]
    pub checked_out_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "checkedOutById")]
    pub checked_out_by_id: Option<String>,
    #[sea_orm(column_name = "checkedInAt", column_type = "TimestampWithTimeZone")]
    pub checked_in_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "checkedInById")]
    pub checked_in_by_id: Option<String>,
    #[sea_orm(column_name = "checkedOutQuantity", default_value = 0)]
    pub checked_out_quantity: i32,
    #[sea_orm(column_name = "sourceLocationId")]
    pub source_location_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::booking::Entity",
        from = "Column::BookingId",
        to = "super::booking::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Booking,
    #[sea_orm(
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Asset,
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Booking.def()
    }
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Asset.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
