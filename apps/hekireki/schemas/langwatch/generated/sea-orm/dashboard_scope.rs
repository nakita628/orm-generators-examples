use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Who sees a dashboard, named as the prompt scope is, with one more level.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "DashboardScope")]
pub enum DashboardScope {
    #[sea_orm(string_value = "PRIVATE")]
    Private,
    #[sea_orm(string_value = "PROJECT")]
    Project,
    #[sea_orm(string_value = "ORGANIZATION")]
    Organization,
}
