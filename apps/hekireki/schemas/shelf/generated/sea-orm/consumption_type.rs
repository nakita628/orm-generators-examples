use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// * Controls what happens after checkout for quantity-tracked assets.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "ConsumptionType")]
pub enum ConsumptionType {
    #[sea_orm(string_value = "ONE_WAY")]
    OneWay,
    #[sea_orm(string_value = "TWO_WAY")]
    TwoWay,
}
