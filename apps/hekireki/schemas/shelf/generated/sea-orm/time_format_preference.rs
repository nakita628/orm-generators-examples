use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "TimeFormatPreference"
)]
pub enum TimeFormatPreference {
    #[sea_orm(string_value = "H12")]
    H12,
    #[sea_orm(string_value = "H24")]
    H24,
}
