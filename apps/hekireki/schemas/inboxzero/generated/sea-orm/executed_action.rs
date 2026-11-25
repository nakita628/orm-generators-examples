use super::action_type::ActionType;
use super::draft_email_status::DraftEmailStatus;
use super::executed_action_status::ExecutedActionStatus;
use super::messaging_message_status::MessagingMessageStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ExecutedAction")]
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
    pub r#type: ActionType,
    #[sea_orm(column_name = "executionStatus")]
    pub execution_status: Option<ExecutedActionStatus>,
    #[sea_orm(
        column_name = "executionStartedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub execution_started_at: Option<DateTime>,
    #[sea_orm(column_name = "executedAt", column_type = "custom(\"timestamp(3)\")")]
    pub executed_at: Option<DateTime>,
    #[sea_orm(column_name = "executionError", column_type = "JsonBinary")]
    pub execution_error: Option<Json>,
    #[sea_orm(
        column_name = "sentMessageIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub sent_message_ids: Vec<String>,
    #[sea_orm(column_name = "executedRuleId", column_type = "Text")]
    pub executed_rule_id: String,
    #[sea_orm(column_name = "messagingChannelId", column_type = "Text")]
    pub messaging_channel_id: Option<String>,
    #[sea_orm(column_name = "messagingMessageId", column_type = "Text")]
    pub messaging_message_id: Option<String>,
    #[sea_orm(
        column_name = "messagingMessageSentAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub messaging_message_sent_at: Option<DateTime>,
    #[sea_orm(column_name = "messagingMessageStatus")]
    pub messaging_message_status: Option<MessagingMessageStatus>,
    #[sea_orm(column_type = "Text")]
    pub label: Option<String>,
    #[sea_orm(column_name = "labelId", column_type = "Text")]
    pub label_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub subject: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub content: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub to: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub cc: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub bcc: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub url: Option<String>,
    #[sea_orm(column_name = "folderName", column_type = "Text")]
    pub folder_name: Option<String>,
    #[sea_orm(column_name = "folderId", column_type = "Text")]
    pub folder_id: Option<String>,
    #[sea_orm(column_name = "staticAttachments", column_type = "JsonBinary")]
    pub static_attachments: Option<Json>,
    #[sea_orm(column_name = "selectedAttachments", column_type = "JsonBinary")]
    pub selected_attachments: Option<Json>,
    #[sea_orm(column_name = "integrationName", column_type = "Text")]
    pub integration_name: Option<String>,
    #[sea_orm(column_name = "integrationToolName", column_type = "Text")]
    pub integration_tool_name: Option<String>,
    #[sea_orm(column_name = "integrationArgs", column_type = "JsonBinary")]
    pub integration_args: Option<Json>,
    #[sea_orm(column_name = "draftId", column_type = "Text")]
    pub draft_id: Option<String>,
    #[sea_orm(column_name = "draftStatus")]
    pub draft_status: Option<DraftEmailStatus>,
    #[sea_orm(column_name = "wasDraftSent")]
    pub was_draft_sent: Option<bool>,
    #[sea_orm(column_name = "draftModelProvider", column_type = "Text")]
    pub draft_model_provider: Option<String>,
    #[sea_orm(column_name = "draftModelName", column_type = "Text")]
    pub draft_model_name: Option<String>,
    #[sea_orm(column_name = "draftPipelineVersion")]
    pub draft_pipeline_version: Option<i32>,
    #[sea_orm(column_name = "draftContextMetadata", column_type = "JsonBinary")]
    pub draft_context_metadata: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::executed_rule::Entity",
        from = "Column::ExecutedRuleId",
        to = "super::executed_rule::Column::Id",
        on_delete = "Cascade",
        fk_name = "ExecutedAction_executedRuleId_fkey"
    )]
    ExecutedRule,
    #[sea_orm(
        belongs_to = "super::messaging_channel::Entity",
        from = "Column::MessagingChannelId",
        to = "super::messaging_channel::Column::Id",
        on_delete = "SetNull",
        fk_name = "ExecutedAction_messagingChannelId_fkey"
    )]
    MessagingChannel,
    #[sea_orm(has_many = "super::digest_item::Entity")]
    DigestItems,
    #[sea_orm(has_one = "super::draft_send_log::Entity")]
    DraftSendLog,
    #[sea_orm(has_one = "super::scheduled_action::Entity")]
    ScheduledAction,
}

impl Related<super::executed_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ExecutedRule.def()
    }
}

impl Related<super::messaging_channel::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::MessagingChannel.def()
    }
}

impl Related<super::digest_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DigestItems.def()
    }
}

impl Related<super::draft_send_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DraftSendLog.def()
    }
}

impl Related<super::scheduled_action::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScheduledAction.def()
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
