use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "HashedLink")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub link: String,
    #[sea_orm(column_name = "eventTypeId")]
    pub event_type_id: i32,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_name = "maxUsageCount", default_value = 1)]
    pub max_usage_count: i32,
    #[sea_orm(column_name = "usageCount", default_value = 0)]
    pub usage_count: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::event_type::Entity",
        from = "Column::EventTypeId",
        to = "super::event_type::Column::Id",
        on_delete = "Cascade",
        fk_name = "HashedLink_eventTypeId_fkey"
    )]
    EventType,
}

impl Related<super::event_type::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EventType.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
