use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "AssetIndexMode")]
pub enum AssetIndexMode {
    #[sea_orm(string_value = "SIMPLE")]
    Simple,
    #[sea_orm(string_value = "ADVANCED")]
    Advanced,
}
