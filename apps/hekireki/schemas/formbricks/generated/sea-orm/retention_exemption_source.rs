use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Who made an exemption: a person (with a reason), or a restore from the archive under a surveys policy
/// whose only condition is "created before" (`restartSurveyRetention`), which carries no reason.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "RetentionExemptionSource"
)]
pub enum RetentionExemptionSource {
    #[sea_orm(string_value = "manual")]
    Manual,
    #[sea_orm(string_value = "restore")]
    Restore,
}
