use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "GatewayGuardrailType"
)]
pub enum GatewayGuardrailType {
    #[sea_orm(string_value = "REGEX")]
    Regex,
    #[sea_orm(string_value = "BEDROCK")]
    Bedrock,
}
