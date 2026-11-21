use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "BookingStatus")]
pub enum BookingStatus {
    #[sea_orm(string_value = "DRAFT")]
    Draft,
    #[sea_orm(string_value = "RESERVED")]
    Reserved,
    #[sea_orm(string_value = "ONGOING")]
    Ongoing,
    #[sea_orm(string_value = "OVERDUE")]
    Overdue,
    #[sea_orm(string_value = "COMPLETE")]
    Complete,
    #[sea_orm(string_value = "ARCHIVED")]
    Archived,
    #[sea_orm(string_value = "CANCELLED")]
    Cancelled,
}
