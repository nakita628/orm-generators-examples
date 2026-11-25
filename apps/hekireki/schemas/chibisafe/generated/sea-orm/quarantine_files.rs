use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "quarantineFiles")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(unique)]
    #[sea_orm(column_name = "fileId", column_type = "Text")]
    pub file_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub reason: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::files::Entity",
        from = "Column::FileId",
        to = "super::files::Column::Uuid",
        fk_name = "quarantineFiles_fileId_fkey"
    )]
    File,
}

impl Related<super::files::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::File.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
