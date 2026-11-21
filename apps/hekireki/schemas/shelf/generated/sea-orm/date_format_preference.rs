use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "DateFormatPreference"
)]
pub enum DateFormatPreference {
    #[sea_orm(string_value = "DD_MM_YYYY")]
    DdMmYyyy,
    #[sea_orm(string_value = "MM_DD_YYYY")]
    MmDdYyyy,
    #[sea_orm(string_value = "YYYY_MM_DD")]
    YyyyMmDd,
    #[sea_orm(string_value = "MMM_DD_YYYY")]
    MmmDdYyyy,
    #[sea_orm(string_value = "DD_MMM_YYYY")]
    DdMmmYyyy,
}
