use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// How the last fetch of a link's page went.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "LinkPreviewOutcome")]
pub enum LinkPreviewOutcome {
    #[sea_orm(string_value = "TITLE")]
    Title,
    #[sea_orm(string_value = "NO_TITLE")]
    NoTitle,
    #[sea_orm(string_value = "SIGN_IN")]
    SignIn,
    #[sea_orm(string_value = "BLOCKED")]
    Blocked,
    #[sea_orm(string_value = "UNREACHABLE")]
    Unreachable,
}
