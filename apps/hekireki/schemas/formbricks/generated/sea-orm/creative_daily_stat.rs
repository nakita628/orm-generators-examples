use super::creative_device_type::CreativeDeviceType;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
/// only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
/// running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
/// the creative deletes its counts.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "CreativeDailyStat")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "creativeId", column_type = "Text")]
    pub creative_id: String,
    /// UTC day.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Date")]
    pub date: Date,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "deviceType")]
    pub device_type: CreativeDeviceType,
    #[sea_orm(default_value = 0)]
    pub impressions: i32,
    #[sea_orm(default_value = 0)]
    pub clicks: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::creative::Entity",
        from = "Column::CreativeId",
        to = "super::creative::Column::Id",
        on_delete = "Cascade",
        fk_name = "CreativeDailyStat_creativeId_fkey"
    )]
    Creative,
}

impl Related<super::creative::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creative.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
