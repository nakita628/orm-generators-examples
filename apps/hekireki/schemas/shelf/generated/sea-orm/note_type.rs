use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "NoteType")]
pub enum NoteType {
    #[sea_orm(string_value = "COMMENT")]
    Comment,
    #[sea_orm(string_value = "UPDATE")]
    Update,
}
