use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "_AssetToTag")]
pub struct Model {
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "A",
        column_type = "Text"
    )]
    pub asset_id: String,
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
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        fk_name = "_AssetToTag_A_fkey"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::tag::Entity",
        from = "Column::TagId",
        to = "super::tag::Column::Id",
        fk_name = "_AssetToTag_B_fkey"
    )]
    Tag,
}

impl ActiveModelBehavior for ActiveModel {}
