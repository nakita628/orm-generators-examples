use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "TierId")]
pub enum TierId {
    #[sea_orm(string_value = "free")]
    Free,
    #[sea_orm(string_value = "tier_1")]
    Tier1,
    #[sea_orm(string_value = "tier_2")]
    Tier2,
    #[sea_orm(string_value = "custom")]
    Custom,
}
