use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The surface that recorded a `Scan`.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "ScanSource")]
pub enum ScanSource {
    #[sea_orm(string_value = "QR_LINK")]
    QrLink,
    #[sea_orm(string_value = "WEB_SCANNER")]
    WebScanner,
    #[sea_orm(string_value = "WEB_DRAWER")]
    WebDrawer,
    #[sea_orm(string_value = "AUDIT")]
    Audit,
    #[sea_orm(string_value = "COMPANION")]
    Companion,
    #[sea_orm(string_value = "GPS_UPDATE")]
    GpsUpdate,
}
