use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Defines follow-up actions for survey responses.
/// Enables automated actions based on specific survey response conditions.
///
/// @property id - Unique identifier for the follow-up
/// @property survey - The associated survey
/// @property name - Display name for the follow-up
/// @property trigger - Conditions that activate the follow-up
/// @property action - Actions to take when triggered
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SurveyFollowUp")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    #[sea_orm(column_name = "surveyId")]
    pub survey_id: String,
    pub name: String,
    /// [SurveyFollowUpTrigger]
    pub trigger: Json,
    /// [SurveyFollowUpAction]
    pub action: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "Column::SurveyId",
        to = "super::survey::Column::Id",
        on_delete = "Cascade"
    )]
    Survey,
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
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
