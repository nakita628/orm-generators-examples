use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "BarcodeType")]
pub enum BarcodeType {
    #[sea_orm(string_value = "Code128")]
    Code128,
    #[sea_orm(string_value = "Code39")]
    Code39,
    #[sea_orm(string_value = "DataMatrix")]
    Datamatrix,
    #[sea_orm(string_value = "ExternalQR")]
    Externalqr,
    #[sea_orm(string_value = "EAN13")]
    Ean13,
}
