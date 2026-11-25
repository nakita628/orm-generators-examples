use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "GatewayGuardrailAction"
)]
pub enum GatewayGuardrailAction {
    #[sea_orm(string_value = "BLOCK")]
    Block,
    #[sea_orm(string_value = "LOG")]
    Log,
    #[sea_orm(string_value = "REDACT")]
    Redact,
}
