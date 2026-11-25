use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "program_coaches")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "programId", column_type = "Text")]
    pub program_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub image: String,
    #[sea_orm(default_value = 0)]
    pub order: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::program::Entity",
        from = "Column::ProgramId",
        to = "super::program::Column::Id",
        on_delete = "Cascade",
        fk_name = "program_coaches_programId_fkey"
    )]
    Program,
}

impl Related<super::program::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Program.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
