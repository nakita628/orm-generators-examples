use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "InviteStatuses")]
pub enum InviteStatuses {
    #[sea_orm(string_value = "PENDING")]
    Pending,
    #[sea_orm(string_value = "ACCEPTED")]
    Accepted,
    #[sea_orm(string_value = "REJECTED")]
    Rejected,
    #[sea_orm(string_value = "INVALIDATED")]
    Invalidated,
}
