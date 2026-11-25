use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrderItems")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "orderId", column_type = "Text")]
    pub order_id: String,
    #[sea_orm(column_name = "integrationId", column_type = "Text")]
    pub integration_id: String,
    pub quantity: i32,
    pub price: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::integration::Entity",
        from = "Column::IntegrationId",
        to = "super::integration::Column::Id",
        fk_name = "OrderItems_integrationId_fkey"
    )]
    Integration,
    #[sea_orm(
        belongs_to = "super::orders::Entity",
        from = "Column::OrderId",
        to = "super::orders::Column::Id",
        fk_name = "OrderItems_orderId_fkey"
    )]
    Order,
}

impl Related<super::integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Integration.def()
    }
}

impl Related<super::orders::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Order.def()
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
