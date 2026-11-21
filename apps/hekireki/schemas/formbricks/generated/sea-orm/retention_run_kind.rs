use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// What wrote a `RetentionRun`: the nightly sweep, or the archive purge recording the deletion of surveys
/// the surveys policy archived. Only sweep runs count as the sweep having run (its lease, its warning
/// restart after a gap, where its scan resumes, the health check).
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "RetentionRunKind")]
pub enum RetentionRunKind {
    #[sea_orm(string_value = "sweep")]
    Sweep,
    #[sea_orm(string_value = "purge")]
    Purge,
}
