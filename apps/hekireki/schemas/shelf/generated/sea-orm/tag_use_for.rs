use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "TagUseFor")]
pub enum TagUseFor {
    #[sea_orm(string_value = "ASSET")]
    Asset,
    #[sea_orm(string_value = "BOOKING")]
    Booking,
}
