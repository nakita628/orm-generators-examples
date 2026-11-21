use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "OrganizationRoles")]
pub enum OrganizationRoles {
    #[sea_orm(string_value = "ADMIN")]
    Admin,
    #[sea_orm(string_value = "BASE")]
    Base,
    #[sea_orm(string_value = "OWNER")]
    Owner,
    #[sea_orm(string_value = "SELF_SERVICE")]
    SelfService,
    #[sea_orm(string_value = "MANAGER")]
    Manager,
}
