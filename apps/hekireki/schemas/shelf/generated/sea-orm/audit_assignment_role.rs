use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "AuditAssignmentRole"
)]
pub enum AuditAssignmentRole {
    #[sea_orm(string_value = "LEAD")]
    Lead,
    #[sea_orm(string_value = "PARTICIPANT")]
    Participant,
}
