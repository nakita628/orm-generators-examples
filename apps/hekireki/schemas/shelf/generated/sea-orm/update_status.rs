use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "UpdateStatus")]
pub enum UpdateStatus {
    #[sea_orm(string_value = "DRAFT")]
    Draft,
    #[sea_orm(string_value = "PUBLISHED")]
    Published,
}
