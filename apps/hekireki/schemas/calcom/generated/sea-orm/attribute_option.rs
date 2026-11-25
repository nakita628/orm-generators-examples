use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AttributeOption")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "attributeId", column_type = "Text")]
    pub attribute_id: String,
    #[sea_orm(column_type = "Text")]
    pub value: String,
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_name = "isGroup", default_value = false)]
    pub is_group: bool,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub contains: Vec<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::attribute::Entity",
        from = "Column::AttributeId",
        to = "super::attribute::Column::Id",
        on_delete = "Cascade",
        fk_name = "AttributeOption_attributeId_fkey"
    )]
    Attribute,
    #[sea_orm(has_many = "super::attribute_to_user::Entity")]
    AssignedUsers,
}

impl Related<super::attribute::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Attribute.def()
    }
}

impl Related<super::attribute_to_user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssignedUsers.def()
    }
}

impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }
}
