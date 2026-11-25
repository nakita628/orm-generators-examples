use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SelectedSlots")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(column_name = "eventTypeId")]
    pub event_type_id: i32,
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(
        column_name = "slotUtcStartDate",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub slot_utc_start_date: DateTime,
    #[sea_orm(
        column_name = "slotUtcEndDate",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub slot_utc_end_date: DateTime,
    #[sea_orm(column_type = "Text")]
    pub uid: String,
    #[sea_orm(column_name = "releaseAt", column_type = "custom(\"timestamp(3)\")")]
    pub release_at: DateTime,
    #[sea_orm(column_name = "isSeat", default_value = false)]
    pub is_seat: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
