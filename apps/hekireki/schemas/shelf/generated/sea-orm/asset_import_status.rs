use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Lifecycle of a background asset import.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "AssetImportStatus")]
pub enum AssetImportStatus {
    #[sea_orm(string_value = "DRAFT")]
    Draft,
    #[sea_orm(string_value = "QUEUED")]
    Queued,
    #[sea_orm(string_value = "RUNNING")]
    Running,
    #[sea_orm(string_value = "COMPLETED")]
    Completed,
    #[sea_orm(string_value = "FAILED")]
    Failed,
}
