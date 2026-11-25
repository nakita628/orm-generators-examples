use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
/// and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
/// so retention adds nothing to the survey row.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RetentionSurveyRestore")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: String,
    #[sea_orm(column_name = "restoredAt", column_type = "custom(\"timestamp(3)\")")]
    pub restored_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "Column::SurveyId",
        to = "super::survey::Column::Id",
        on_delete = "Cascade",
        fk_name = "RetentionSurveyRestore_surveyId_fkey"
    )]
    Survey,
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
