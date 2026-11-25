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
    #[sea_orm(column_type = "Text")]
    pub name: String,
    /// [SurveyFollowUpTrigger]
    #[sea_orm(column_type = "JsonBinary")]
    pub trigger: Json,
    /// [SurveyFollowUpAction]
    #[sea_orm(column_type = "JsonBinary")]
    pub action: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "Column::SurveyId",
        to = "super::survey::Column::Id",
        on_delete = "Cascade",
        fk_name = "SurveyFollowUp_surveyId_fkey"
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
