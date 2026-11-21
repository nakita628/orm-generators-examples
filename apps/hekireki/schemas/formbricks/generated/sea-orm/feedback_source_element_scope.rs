use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Whether a formbricks_survey source tracks every supported element of the surveys it maps, or an
/// explicit subset the operator curated. Ignored by csv sources, which map fields instead.
///
/// Orthogonal to FeedbackSourceImportMode: that decides which responses are imported, this decides
/// which questions a response is imported for.
///
/// This exists because a question added to a survey after the source was connected is
/// indistinguishable, from stored state alone, from a question the operator deliberately excluded —
/// survey elements carry no timestamps. Recording the intent at write time (where the full candidate
/// set IS known) is the same approach webhooks take with `surveyIds: []` meaning "all surveys", and
/// segment survey-interaction filters with `surveyScope: "any" | "specific"`.
///
/// Defaults to `specific` so sources that predate the column keep their current behaviour: they gain
/// no mappings on their own until the operator next saves the source, at which point the scope is
/// derived from the selection they submit.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "FeedbackSourceElementScope"
)]
pub enum FeedbackSourceElementScope {
    #[sea_orm(string_value = "all")]
    All,
    #[sea_orm(string_value = "specific")]
    Specific,
}
