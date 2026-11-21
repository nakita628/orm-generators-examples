use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// What a reviewer decided about a response. One value today; null on `Response.outcome` means none.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "ResponseOutcome")]
pub enum ResponseOutcome {
    #[sea_orm(string_value = "lowQuality")]
    Lowquality,
}
