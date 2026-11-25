use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DraftSendLog")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "executedActionId", column_type = "Text")]
    pub executed_action_id: String,
    #[sea_orm(column_name = "sentMessageId", column_type = "Text")]
    pub sent_message_id: String,
    #[sea_orm(column_name = "similarityScore")]
    pub similarity_score: f64,
    #[sea_orm(column_name = "bodySimilarityScore")]
    pub body_similarity_score: Option<f64>,
    #[sea_orm(column_name = "bodySimilarityStatus", column_type = "Text")]
    pub body_similarity_status: Option<String>,
    #[sea_orm(column_name = "similarityMetadata", column_type = "JsonBinary")]
    pub similarity_metadata: Option<Json>,
    #[sea_orm(column_name = "sentText", column_type = "Text")]
    pub sent_text: Option<String>,
    #[sea_orm(column_name = "replyMemorySentText", column_type = "Text")]
    pub reply_memory_sent_text: Option<String>,
    #[sea_orm(column_name = "replyMemoryAttemptCount", default_value = 0)]
    pub reply_memory_attempt_count: i32,
    #[sea_orm(
        column_name = "replyMemoryProcessedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub reply_memory_processed_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::executed_action::Entity",
        from = "Column::ExecutedActionId",
        to = "super::executed_action::Column::Id",
        on_delete = "Cascade",
        fk_name = "DraftSendLog_executedActionId_fkey"
    )]
    ExecutedAction,
    #[sea_orm(has_many = "super::reply_memory_source::Entity")]
    ReplyMemorySources,
}

impl Related<super::executed_action::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ExecutedAction.def()
    }
}

impl Related<super::reply_memory_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ReplyMemorySources.def()
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
