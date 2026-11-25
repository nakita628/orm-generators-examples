use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "_CategoryToCustomField")]
pub struct Model {
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "A",
        column_type = "Text"
    )]
    pub category_id: String,
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "B",
        column_type = "Text"
    )]
    pub custom_field_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::category::Entity",
        from = "Column::CategoryId",
        to = "super::category::Column::Id",
        fk_name = "_CategoryToCustomField_A_fkey"
    )]
    Category,
    #[sea_orm(
        belongs_to = "super::custom_field::Entity",
        from = "Column::CustomFieldId",
        to = "super::custom_field::Column::Id",
        fk_name = "_CategoryToCustomField_B_fkey"
    )]
    CustomField,
}

impl ActiveModelBehavior for ActiveModel {}
