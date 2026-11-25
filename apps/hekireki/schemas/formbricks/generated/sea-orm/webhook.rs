use super::pipeline_triggers::PipelineTriggers;
use super::webhook_source::WebhookSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a webhook endpoint for receiving survey-related events.
/// Webhooks can be configured to receive notifications about response creation, updates, and completion.
///
/// @property id - Unique identifier for the webhook
/// @property name - Optional display name for the webhook
/// @property url - The endpoint URL where events will be sent
/// @property source - Origin of the webhook (user, zapier, make, etc.)
/// @property workspace - Associated workspace
/// @property triggers - Types of events that trigger this webhook
/// @property surveyIds - List of surveys this webhook is monitoring
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Webhook")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub url: String,
    #[sea_orm(default_value = "user")]
    pub source: WebhookSource,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    pub triggers: Vec<PipelineTriggers>,
    #[sea_orm(
        column_name = "surveyIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub survey_ids: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub secret: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "Webhook_workspaceId_fkey"
    )]
    Workspace,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
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
