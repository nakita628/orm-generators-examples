use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "KitStatus")]
pub enum KitStatus {
    #[sea_orm(string_value = "AVAILABLE")]
    Available,
    #[sea_orm(string_value = "IN_CUSTODY")]
    InCustody,
    #[sea_orm(string_value = "CHECKED_OUT")]
    CheckedOut,
}
