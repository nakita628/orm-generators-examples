use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "OAuthClientKind")]
pub enum OAuthClientKind {
    #[sea_orm(string_value = "PREREGISTERED")]
    Preregistered,
    #[sea_orm(string_value = "CIMD")]
    Cimd,
    #[sea_orm(string_value = "DCR")]
    Dcr,
}
