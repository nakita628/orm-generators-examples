use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// * Determines how an asset is tracked throughout the system. Immutable after creation.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "AssetType")]
pub enum AssetType {
    #[sea_orm(string_value = "INDIVIDUAL")]
    Individual,
    #[sea_orm(string_value = "QUANTITY_TRACKED")]
    QuantityTracked,
}
