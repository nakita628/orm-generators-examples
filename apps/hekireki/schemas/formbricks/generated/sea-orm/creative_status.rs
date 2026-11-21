use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The author's intent only. Scheduled, Active, "Paused with its survey" and Closed after the
/// end time are derived (`apps/web/modules/ee/creatives/lib/display-status.ts` and its SQL twin); nothing
/// writes this column when a schedule ends or a survey pauses.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "CreativeStatus")]
pub enum CreativeStatus {
    #[sea_orm(string_value = "draft")]
    Draft,
    #[sea_orm(string_value = "inProgress")]
    Inprogress,
    #[sea_orm(string_value = "paused")]
    Paused,
    #[sea_orm(string_value = "completed")]
    Completed,
}
