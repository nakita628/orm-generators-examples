use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
/// Maps source fields to Hub FeedbackRecord fields.
///
/// @property id - Unique identifier for the mapping
/// @property feedbackSource - The parent feedback source
/// @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
/// @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
/// @property staticValue - If set, use this value instead of reading from sourceFieldId
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "FeedbackSourceFieldMapping")]
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
    #[sea_orm(column_type = "Text")]
    pub source_field_id: String,
    #[sea_orm(column_type = "Text")]
    pub target_field_id: String,
    #[sea_orm(column_type = "Text")]
    pub static_value: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::feedback_source::Entity",
        from = "(Column::FeedbackSourceId, Column::WorkspaceId)",
        to = "(super::feedback_source::Column::Id, super::feedback_source::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "FeedbackSourceFieldMapping_feedbackSourceId_workspaceId_fkey"
    )]
    FeedbackSource,
}

impl Related<super::feedback_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackSource.def()
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
