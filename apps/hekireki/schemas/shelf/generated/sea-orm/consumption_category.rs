use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// * Classifies the type of quantity change in a ConsumptionLog entry.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "ConsumptionCategory"
)]
pub enum ConsumptionCategory {
    #[sea_orm(string_value = "CHECKOUT")]
    Checkout,
    #[sea_orm(string_value = "RETURN")]
    Return,
    #[sea_orm(string_value = "RESTOCK")]
    Restock,
    #[sea_orm(string_value = "ADJUSTMENT")]
    Adjustment,
    #[sea_orm(string_value = "LOSS")]
    Loss,
    #[sea_orm(string_value = "CONSUME")]
    Consume,
    #[sea_orm(string_value = "DAMAGE")]
    Damage,
    #[sea_orm(string_value = "MOVE")]
    Move,
    #[sea_orm(string_value = "INITIAL")]
    Initial,
    #[sea_orm(string_value = "OPENING_BALANCE")]
    OpeningBalance,
}
