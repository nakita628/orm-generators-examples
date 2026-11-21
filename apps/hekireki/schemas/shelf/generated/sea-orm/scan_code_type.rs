use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The kind of code a `Scan` recorded.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "ScanCodeType")]
pub enum ScanCodeType {
    #[sea_orm(string_value = "QR")]
    Qr,
    #[sea_orm(string_value = "BARCODE")]
    Barcode,
    #[sea_orm(string_value = "SAM_ID")]
    SamId,
}
