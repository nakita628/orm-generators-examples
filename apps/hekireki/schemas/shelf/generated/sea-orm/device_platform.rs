use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The operating system a companion-app push token was issued for.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "DevicePlatform")]
pub enum DevicePlatform {
    #[sea_orm(string_value = "IOS")]
    Ios,
    #[sea_orm(string_value = "ANDROID")]
    Android,
}
