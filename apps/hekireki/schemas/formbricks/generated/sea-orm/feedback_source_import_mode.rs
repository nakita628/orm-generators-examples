use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Which survey responses a Formbricks feedback source imports into Hub.
/// completedOnly matches the live ingestion path, which only runs on responseFinished.
/// all additionally imports partial responses — answers a respondent typed but never submitted.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "FeedbackSourceImportMode"
)]
pub enum FeedbackSourceImportMode {
    #[sea_orm(string_value = "completedOnly")]
    Completedonly,
    #[sea_orm(string_value = "all")]
    All,
}
