use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "McpGrantAccess")]
pub enum McpGrantAccess {
    #[sea_orm(string_value = "all")]
    All,
    #[sea_orm(string_value = "organizations")]
    Organizations,
    #[sea_orm(string_value = "workspaces")]
    Workspaces,
}
