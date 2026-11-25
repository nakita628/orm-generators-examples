use super::email_source_type::EmailSourceType;
use super::email_status::EmailStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "emails")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "contactId", column_type = "Text")]
    pub contact_id: String,
    #[sea_orm(column_name = "toName", column_type = "Text")]
    pub to_name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub subject: String,
    #[sea_orm(column_type = "Text")]
    pub body: String,
    #[sea_orm(column_type = "Text")]
    pub from: String,
    #[sea_orm(column_name = "fromName", column_type = "Text")]
    pub from_name: Option<String>,
    #[sea_orm(column_name = "replyTo", column_type = "Text")]
    pub reply_to: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub headers: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub attachments: Option<Json>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "messageId", column_type = "Text")]
    pub message_id: Option<String>,
    #[sea_orm(column_name = "sourceType")]
    pub source_type: EmailSourceType,
    #[sea_orm(column_name = "templateId", column_type = "Text")]
    pub template_id: Option<String>,
    #[sea_orm(column_name = "campaignId", column_type = "Text")]
    pub campaign_id: Option<String>,
    #[sea_orm(column_name = "workflowExecutionId", column_type = "Text")]
    pub workflow_execution_id: Option<String>,
    #[sea_orm(column_name = "workflowStepExecutionId", column_type = "Text")]
    pub workflow_step_execution_id: Option<String>,
    #[sea_orm(default_value = "PENDING")]
    pub status: EmailStatus,
    /// When a transactional email sent with `delivery.at` is due. Only meaningful while the
    /// status is SCHEDULED; the delayed queue job is the timer and the scheduled-email sweep
    /// re-queues any row still SCHEDULED past this time. The partial index on SCHEDULED rows
    /// is built by the background index builder, not a migration (see background-index-builder.ts).
    #[sea_orm(column_name = "scheduledFor", column_type = "custom(\"timestamp(3)\")")]
    pub scheduled_for: Option<DateTime>,
    #[sea_orm(column_name = "sentAt", column_type = "custom(\"timestamp(3)\")")]
    pub sent_at: Option<DateTime>,
    #[sea_orm(column_name = "deliveredAt", column_type = "custom(\"timestamp(3)\")")]
    pub delivered_at: Option<DateTime>,
    #[sea_orm(column_name = "openedAt", column_type = "custom(\"timestamp(3)\")")]
    pub opened_at: Option<DateTime>,
    #[sea_orm(column_name = "clickedAt", column_type = "custom(\"timestamp(3)\")")]
    pub clicked_at: Option<DateTime>,
    #[sea_orm(column_name = "bouncedAt", column_type = "custom(\"timestamp(3)\")")]
    pub bounced_at: Option<DateTime>,
    #[sea_orm(column_name = "complainedAt", column_type = "custom(\"timestamp(3)\")")]
    pub complained_at: Option<DateTime>,
    /// When the recipient opted out from this email (unsubscribe or snooze). Only set for
    /// opt-outs attributed to a campaign email and counted in `Campaign.unsubscribedCount`,
    /// in the same transaction, so the stamped rows are the list behind that figure. Indexed
    /// by a partial index built in the background (see background-index-builder.ts).
    #[sea_orm(
        column_name = "unsubscribedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub unsubscribed_at: Option<DateTime>,
    #[sea_orm(default_value = 0)]
    pub opens: i32,
    #[sea_orm(default_value = 0)]
    pub clicks: i32,
    /// The address handed to SES belonged to the mailbox simulator
    /// (`@simulator.amazonses.com`), stamped at send time from the address actually used --
    /// which is the recipient override when one is set, not necessarily `contact.email`.
    ///
    /// AWS itself excludes simulator mail from bounce rates, complaint rates, sending quotas
    /// and reputation metrics, so Plunk's security rates exclude it from both the numerator
    /// and the denominator. The events themselves are untouched: the webhook fires, the
    /// activity feed shows the bounce, the contact is still suppressed.
    #[sea_orm(default_value = false)]
    pub simulated: bool,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::contact::Entity",
        from = "Column::ContactId",
        to = "super::contact::Column::Id",
        on_delete = "Cascade",
        fk_name = "emails_contactId_fkey"
    )]
    Contact,
    #[sea_orm(
        belongs_to = "super::template::Entity",
        from = "Column::TemplateId",
        to = "super::template::Column::Id",
        fk_name = "emails_templateId_fkey"
    )]
    Template,
    #[sea_orm(
        belongs_to = "super::campaign::Entity",
        from = "Column::CampaignId",
        to = "super::campaign::Column::Id",
        on_delete = "Cascade",
        fk_name = "emails_campaignId_fkey"
    )]
    Campaign,
    #[sea_orm(
        belongs_to = "super::workflow_execution::Entity",
        from = "Column::WorkflowExecutionId",
        to = "super::workflow_execution::Column::Id",
        on_delete = "Cascade",
        fk_name = "emails_workflowExecutionId_fkey"
    )]
    WorkflowExecution,
    #[sea_orm(
        belongs_to = "super::workflow_step_execution::Entity",
        from = "Column::WorkflowStepExecutionId",
        to = "super::workflow_step_execution::Column::Id",
        on_delete = "Cascade",
        fk_name = "emails_workflowStepExecutionId_fkey"
    )]
    WorkflowStepExecution,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "emails_projectId_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::event::Entity")]
    Events,
}

impl Related<super::contact::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Contact.def()
    }
}

impl Related<super::template::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Template.def()
    }
}

impl Related<super::campaign::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Campaign.def()
    }
}

impl Related<super::workflow_execution::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkflowExecution.def()
    }
}

impl Related<super::workflow_step_execution::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkflowStepExecution.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Events.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
