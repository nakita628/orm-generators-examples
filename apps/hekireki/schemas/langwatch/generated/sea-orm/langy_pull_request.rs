use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
/// checks look it up by repository and number to wake the conversation; the daily pull request
/// cap counts a person's rows by day. Rebuilt by a projection replay.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LangyPullRequest")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// `owner/repo`, lowercased: GitHub matches repository names without case.
    #[sea_orm(column_name = "repositoryFullName", column_type = "Text")]
    pub repository_full_name: String,
    #[sea_orm(column_name = "prNumber")]
    pub pr_number: i32,
    #[sea_orm(column_name = "conversationId", column_type = "Text")]
    pub conversation_id: String,
    /// The conversation's owner, who opened it through Langy.
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    /// When the answer that opened it was recorded (epoch ms).
    #[sea_orm(column_name = "openedAt")]
    pub opened_at: f64,
    /// The agent_responded event it came from.
    #[sea_orm(column_name = "sourceEventId", column_type = "Text")]
    pub source_event_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
