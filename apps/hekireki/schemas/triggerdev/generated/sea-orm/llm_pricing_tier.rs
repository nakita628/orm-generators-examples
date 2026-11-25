use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// A pricing tier for a model (supports volume-based or conditional pricing)
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "llm_pricing_tiers")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub model_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(default_value = true)]
    pub is_default: bool,
    #[sea_orm(default_value = 0)]
    pub priority: i32,
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub conditions: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::llm_model::Entity",
        from = "Column::ModelId",
        to = "super::llm_model::Column::Id",
        on_delete = "Cascade",
        fk_name = "llm_pricing_tiers_model_id_fkey"
    )]
    Model,
    #[sea_orm(has_many = "super::llm_price::Entity")]
    Prices,
}

impl Related<super::llm_model::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Model.def()
    }
}

impl Related<super::llm_price::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Prices.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
