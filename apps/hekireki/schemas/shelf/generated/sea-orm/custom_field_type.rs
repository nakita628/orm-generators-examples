use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "CustomFieldType")]
pub enum CustomFieldType {
    #[sea_orm(string_value = "TEXT")]
    Text,
    #[sea_orm(string_value = "OPTION")]
    Option,
    #[sea_orm(string_value = "BOOLEAN")]
    Boolean,
    #[sea_orm(string_value = "DATE")]
    Date,
    #[sea_orm(string_value = "MULTILINE_TEXT")]
    MultilineText,
    #[sea_orm(string_value = "AMOUNT")]
    Amount,
    #[sea_orm(string_value = "NUMBER")]
    Number,
}
