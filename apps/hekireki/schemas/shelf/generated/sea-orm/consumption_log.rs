use super::consumption_category::ConsumptionCategory;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// *
/// * Audit trail for all quantity changes on quantity-tracked assets.
/// * Every event records who performed it, who received/returned items,
/// * which booking it relates to, and the category of change.
/// * Quantity is always positive — direction is determined by category.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ConsumptionLog")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "assetId", column_type = "Text")]
    pub asset_id: String,
    pub category: ConsumptionCategory,
    pub quantity: i32,
    #[sea_orm(column_type = "Text")]
    pub note: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "bookingId", column_type = "Text")]
    pub booking_id: Option<String>,
    #[sea_orm(column_name = "custodianId", column_type = "Text")]
    pub custodian_id: Option<String>,
    #[sea_orm(column_name = "bookingAssetId", column_type = "Text")]
    pub booking_asset_id: Option<String>,
    #[sea_orm(column_name = "locationId", column_type = "Text")]
    pub location_id: Option<String>,
    #[sea_orm(column_name = "stockChange")]
    pub stock_change: Option<i32>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "ConsumptionLog_assetId_fkey"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        fk_name = "ConsumptionLog_userId_fkey"
    )]
    PerformedBy,
    #[sea_orm(
        belongs_to = "super::booking::Entity",
        from = "Column::BookingId",
        to = "super::booking::Column::Id",
        fk_name = "ConsumptionLog_bookingId_fkey"
    )]
    Booking,
    #[sea_orm(
        belongs_to = "super::team_member::Entity",
        from = "Column::CustodianId",
        to = "super::team_member::Column::Id",
        fk_name = "ConsumptionLog_custodianId_fkey"
    )]
    Custodian,
    #[sea_orm(
        belongs_to = "super::location::Entity",
        from = "Column::LocationId",
        to = "super::location::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "ConsumptionLog_locationId_fkey"
    )]
    Location,
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Asset.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PerformedBy.def()
    }
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Booking.def()
    }
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Custodian.def()
    }
}

impl Related<super::location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Location.def()
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
        Ok(self)
    }
}
