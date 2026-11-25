use super::action_type::ActionType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Action")]
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
    #[sea_orm(column_name = "ruleId", column_type = "Text")]
    pub rule_id: String,
    #[sea_orm(column_name = "emailAccountId", column_type = "Text")]
    pub email_account_id: String,
    #[sea_orm(column_name = "messagingChannelId", column_type = "Text")]
    pub messaging_channel_id: Option<String>,
    #[sea_orm(column_name = "messagingChannelEmailAccountId", column_type = "Text")]
    pub messaging_channel_email_account_id: Option<String>,
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
    #[sea_orm(column_name = "delayInMinutes")]
    pub delay_in_minutes: Option<i32>,
    #[sea_orm(column_name = "staticAttachments", column_type = "JsonBinary")]
    pub static_attachments: Option<Json>,
    #[sea_orm(column_name = "integrationName", column_type = "Text")]
    pub integration_name: Option<String>,
    #[sea_orm(column_name = "integrationToolName", column_type = "Text")]
    pub integration_tool_name: Option<String>,
    #[sea_orm(column_name = "integrationArgs", column_type = "JsonBinary")]
    pub integration_args: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::rule::Entity",
        from = "(Column::RuleId, Column::EmailAccountId)",
        to = "(super::rule::Column::Id, super::rule::Column::EmailAccountId)",
        on_delete = "Cascade",
        fk_name = "Action_ruleId_emailAccountId_fkey"
    )]
    Rule,
    #[sea_orm(
        belongs_to = "super::email_account::Entity",
        from = "Column::EmailAccountId",
        to = "super::email_account::Column::Id",
        on_delete = "Cascade",
        fk_name = "Action_emailAccountId_fkey"
    )]
    EmailAccount,
    #[sea_orm(
        belongs_to = "super::messaging_channel::Entity",
        from = "(Column::MessagingChannelId, Column::MessagingChannelEmailAccountId)",
        to = "(super::messaging_channel::Column::Id, super::messaging_channel::Column::EmailAccountId)",
        on_delete = "SetNull",
        fk_name = "Action_messagingChannelId_messagingChannelEmailAccountId_fkey"
    )]
    MessagingChannel,
}

impl Related<super::rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Rule.def()
    }
}

impl Related<super::email_account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmailAccount.def()
    }
}

impl Related<super::messaging_channel::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::MessagingChannel.def()
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
