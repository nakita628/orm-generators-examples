use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "ReminderRecurrenceUnit"
)]
pub enum ReminderRecurrenceUnit {
    #[sea_orm(string_value = "DAY")]
    Day,
    #[sea_orm(string_value = "WEEK")]
    Week,
    #[sea_orm(string_value = "MONTH")]
    Month,
    #[sea_orm(string_value = "YEAR")]
    Year,
}
