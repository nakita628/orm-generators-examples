use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// How a membership's role came to change.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "RoleChangeSource")]
pub enum RoleChangeSource {
    #[sea_orm(string_value = "MANUAL")]
    Manual,
    #[sea_orm(string_value = "SSO")]
    Sso,
}
