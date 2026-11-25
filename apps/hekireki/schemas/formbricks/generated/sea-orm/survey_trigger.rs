use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Links surveys to specific trigger actions.
/// Defines which user actions should cause a survey to be displayed.
/// This is the connection table between Surveys and ActionClasses that determines
/// when and under what conditions a survey should be triggered.
///
/// @property id - Unique identifier for the trigger
/// @property survey - The survey to be triggered
/// @property actionClass - The action that triggers the survey
/// @property createdAt - When the trigger was created
/// @property updatedAt - When the trigger was last modified
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SurveyTrigger")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: String,
    #[sea_orm(column_name = "actionClassId", column_type = "Text")]
    pub action_class_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "Column::SurveyId",
        to = "super::survey::Column::Id",
        on_delete = "Cascade",
        fk_name = "SurveyTrigger_surveyId_fkey"
    )]
    Survey,
    #[sea_orm(
        belongs_to = "super::action_class::Entity",
        from = "Column::ActionClassId",
        to = "super::action_class::Column::Id",
        on_delete = "Cascade",
        fk_name = "SurveyTrigger_actionClassId_fkey"
    )]
    ActionClass,
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
    }
}

impl Related<super::action_class::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ActionClass.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
