use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "EnvironmentVariableIgnoreReason"
)]
pub enum EnvironmentVariableIgnoreReason {
    #[sea_orm(string_value = "EMPTY")]
    Empty,
    #[sea_orm(string_value = "MISSING")]
    Missing,
    #[sea_orm(string_value = "OIDC")]
    Oidc,
}
