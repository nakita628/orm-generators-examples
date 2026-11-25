use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RateLimitCounter")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "bucketKey", column_type = "Text")]
    pub bucket_key: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(
        column_name = "windowStart",
        column_type = "custom(\"timestamptz(3)\")"
    )]
    pub window_start: DateTimeWithTimeZone,
    #[sea_orm(default_value = 0)]
    pub count: i32,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamptz(3)\")")]
    pub expires_at: DateTimeWithTimeZone,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
