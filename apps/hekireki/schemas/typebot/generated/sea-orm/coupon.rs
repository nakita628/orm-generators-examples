use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Coupon")]
pub struct Model {
    #[sea_orm(column_name = "userPropertiesToUpdate", column_type = "JsonBinary")]
    pub user_properties_to_update: Json,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub code: String,
    #[sea_orm(column_name = "dateRedeemed", column_type = "custom(\"timestamp(3)\")")]
    pub date_redeemed: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
