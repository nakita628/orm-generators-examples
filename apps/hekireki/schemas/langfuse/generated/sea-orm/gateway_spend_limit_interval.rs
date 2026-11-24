use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "GatewaySpendLimitInterval"
)]
pub enum GatewaySpendLimitInterval {
    #[sea_orm(string_value = "DAILY")]
    Daily,
    #[sea_orm(string_value = "WEEKLY")]
    Weekly,
    #[sea_orm(string_value = "MONTHLY")]
    Monthly,
    #[sea_orm(string_value = "TOTAL")]
    Total,
}
