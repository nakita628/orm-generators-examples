use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "TaskExecutionMode")]
pub enum TaskExecutionMode {
    #[sea_orm(string_value = "STANDARD")]
    Standard,
    #[sea_orm(string_value = "SANDBOX")]
    Sandbox,
}
