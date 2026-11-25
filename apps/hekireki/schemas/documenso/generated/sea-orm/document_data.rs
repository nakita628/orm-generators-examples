use super::document_data_type::DocumentDataType;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DocumentData")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    pub r#type: DocumentDataType,
    #[sea_orm(column_type = "Text")]
    pub data: String,
    #[sea_orm(column_name = "initialData", column_type = "Text")]
    pub initial_data: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_one = "super::envelope_item::Entity")]
    EnvelopeItem,
}

impl Related<super::envelope_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EnvelopeItem.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
