use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "_OrderToTag")]
pub struct Model {
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "A",
        column_type = "Text"
    )]
    pub order_id: String,
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "B",
        column_type = "Text"
    )]
    pub tag_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::order::Entity",
        from = "Column::OrderId",
        to = "super::order::Column::Id",
        fk_name = "_OrderToTag_A_fkey"
    )]
    Order,
    #[sea_orm(
        belongs_to = "super::tag::Entity",
        from = "Column::TagId",
        to = "super::tag::Column::Id",
        fk_name = "_OrderToTag_B_fkey"
    )]
    Tag,
}

impl ActiveModelBehavior for ActiveModel {}
