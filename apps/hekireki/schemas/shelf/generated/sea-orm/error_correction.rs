use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "ErrorCorrection")]
pub enum ErrorCorrection {
    #[sea_orm(string_value = "L")]
    L,
    #[sea_orm(string_value = "M")]
    M,
    #[sea_orm(string_value = "Q")]
    Q,
    #[sea_orm(string_value = "H")]
    H,
}
