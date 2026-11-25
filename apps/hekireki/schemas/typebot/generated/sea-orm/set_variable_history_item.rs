use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SetVariableHistoryItem")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "resultId", column_type = "Text")]
    pub result_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    pub index: i32,
    #[sea_orm(column_name = "variableId", column_type = "Text")]
    pub variable_id: String,
    #[sea_orm(column_name = "blockId", column_type = "Text")]
    pub block_id: String,
    #[sea_orm(column_name = "blockIndex")]
    pub block_index: Option<i32>,
    #[sea_orm(column_type = "JsonBinary")]
    pub value: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::result::Entity",
        from = "Column::ResultId",
        to = "super::result::Column::Id",
        on_delete = "Cascade",
        fk_name = "SetVariableHistoryItem_resultId_fkey"
    )]
    Result,
}

impl Related<super::result::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Result.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
