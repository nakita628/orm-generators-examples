use super::email_reminder_status::EmailReminderStatus;
use super::scheduled_email_status::ScheduledEmailStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ScheduledEmail")]
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
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "emailAccountId", column_type = "Text")]
    pub email_account_id: String,
    #[sea_orm(column_name = "clientMutationId", column_type = "Text")]
    pub client_mutation_id: String,
    #[sea_orm(column_name = "payloadHash", column_type = "Text")]
    pub payload_hash: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub payload: Json,
    #[sea_orm(column_name = "threadId", column_type = "Text")]
    pub thread_id: Option<String>,
    #[sea_orm(column_name = "sendAt", column_type = "custom(\"timestamp(3)\")")]
    pub send_at: DateTime,
    #[sea_orm(default_value = "PENDING")]
    pub status: ScheduledEmailStatus,
    #[sea_orm(
        column_name = "processingStartedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub processing_started_at: Option<DateTime>,
    #[sea_orm(
        column_name = "executionQueuedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub execution_queued_at: Option<DateTime>,
    #[sea_orm(column_name = "sentAt", column_type = "custom(\"timestamp(3)\")")]
    pub sent_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "remindAt", column_type = "custom(\"timestamp(3)\")")]
    pub remind_at: Option<DateTime>,
    #[sea_orm(column_name = "reminderStatus", default_value = "NONE")]
    pub reminder_status: EmailReminderStatus,
    #[sea_orm(
        column_name = "reminderStartedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub reminder_started_at: Option<DateTime>,
    #[sea_orm(column_name = "heldForUndo", default_value = false)]
    pub held_for_undo: bool,
    #[sea_orm(
        column_name = "draftMessageIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub draft_message_ids: Vec<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::email_account::Entity",
        from = "Column::EmailAccountId",
        to = "super::email_account::Column::Id",
        on_delete = "Cascade",
        fk_name = "ScheduledEmail_emailAccountId_fkey"
    )]
    EmailAccount,
}

impl Related<super::email_account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmailAccount.def()
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
