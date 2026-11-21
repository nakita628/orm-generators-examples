use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "WeekStartPreference"
)]
pub enum WeekStartPreference {
    #[sea_orm(string_value = "MONDAY")]
    Monday,
    #[sea_orm(string_value = "SUNDAY")]
    Sunday,
    #[sea_orm(string_value = "SATURDAY")]
    Saturday,
}
