use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Which step of the background import runs next (or failed).
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "AssetImportPhase")]
pub enum AssetImportPhase {
    #[sea_orm(string_value = "PREPARE")]
    Prepare,
    #[sea_orm(string_value = "ROWS")]
    Rows,
    #[sea_orm(string_value = "FINALIZE")]
    Finalize,
}
