use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "ActivityEntity")]
pub enum ActivityEntity {
    #[sea_orm(string_value = "ASSET")]
    Asset,
    #[sea_orm(string_value = "BOOKING")]
    Booking,
    #[sea_orm(string_value = "AUDIT")]
    Audit,
    #[sea_orm(string_value = "KIT")]
    Kit,
    #[sea_orm(string_value = "LOCATION")]
    Location,
    #[sea_orm(string_value = "TEAM_MEMBER")]
    TeamMember,
    #[sea_orm(string_value = "CUSTODY")]
    Custody,
    #[sea_orm(string_value = "USER")]
    User,
    #[sea_orm(string_value = "ORGANIZATION")]
    Organization,
}
