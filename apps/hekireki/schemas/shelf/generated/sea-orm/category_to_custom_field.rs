use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "_CategoryToCustomField")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false, column_name = "A")]
    pub category_id: String,
    #[sea_orm(primary_key, auto_increment = false, column_name = "B")]
    pub custom_field_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::category::Entity",
        from = "Column::CategoryId",
        to = "super::category::Column::Id"
    )]
    Category,
    #[sea_orm(
        belongs_to = "super::custom_field::Entity",
        from = "Column::CustomFieldId",
        to = "super::custom_field::Column::Id"
    )]
    CustomField,
}

impl ActiveModelBehavior for ActiveModel {}
