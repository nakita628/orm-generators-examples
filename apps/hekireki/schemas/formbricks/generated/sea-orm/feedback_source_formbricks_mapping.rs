use super::hub_field_type::HubFieldType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Maps survey elements to Hub FeedbackRecords for Formbricks feedback sources.
/// Each row represents one element that will create FeedbackRecords when answered.
///
/// @property id - Unique identifier for the mapping
/// @property feedbackSource - The parent feedback source
/// @property survey - The survey containing the element
/// @property elementId - The element ID within the survey (from blocks[].elements[].id)
/// @property hubFieldType - The field_type to use in Hub (text, nps, rating, etc.)
/// @property customFieldLabel - Optional override for the element headline as field_label in Hub
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "FeedbackSourceFormbricksMapping")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub feedback_source_id: String,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: String,
    #[sea_orm(column_name = "elementId", column_type = "Text")]
    pub element_id: String,
    #[sea_orm(column_name = "hubFieldType")]
    pub hub_field_type: HubFieldType,
    #[sea_orm(column_type = "Text")]
    pub custom_field_label: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::feedback_source::Entity",
        from = "(Column::FeedbackSourceId, Column::WorkspaceId)",
        to = "(super::feedback_source::Column::Id, super::feedback_source::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "FeedbackSourceFormbricksMapping_feedbackSourceId_workspaceId_fk"
    )]
    FeedbackSource,
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "(Column::SurveyId, Column::WorkspaceId)",
        to = "(super::survey::Column::Id, super::survey::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "FeedbackSourceFormbricksMapping_surveyId_workspaceId_fkey"
    )]
    Survey,
}

impl Related<super::feedback_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackSource.def()
    }
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
        Ok(self)
    }
}
