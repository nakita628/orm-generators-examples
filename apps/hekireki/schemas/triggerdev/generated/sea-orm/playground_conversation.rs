use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PlaygroundConversation")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The chat session ID used by the transport
    #[sea_orm(column_name = "chatId", column_type = "Text")]
    pub chat_id: String,
    /// User-editable conversation title (auto-generated from first message)
    #[sea_orm(column_type = "Text", default_value = "New conversation")]
    pub title: String,
    /// Which agent this conversation is with
    #[sea_orm(column_name = "agentSlug", column_type = "Text")]
    pub agent_slug: String,
    /// The current active run backing this conversation (null if no run yet)
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: Option<String>,
    /// The client data JSON used for this conversation
    #[sea_orm(column_name = "clientData", column_type = "JsonBinary")]
    pub client_data: Option<Json>,
    /// Accumulated UIMessages from completed turns (for resume without stream replay)
    #[sea_orm(column_type = "JsonBinary")]
    pub messages: Option<Json>,
    /// Last SSE event ID — resume from this position to avoid replaying old turns
    #[sea_orm(column_name = "lastEventId", column_type = "Text")]
    pub last_event_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "runtimeEnvironmentId", column_type = "Text")]
    pub runtime_environment_id: String,
    /// The user who started this conversation
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
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
        belongs_to = "super::task_run::Entity",
        from = "Column::RunId",
        to = "super::task_run::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "PlaygroundConversation_runId_fkey"
    )]
    Run,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "PlaygroundConversation_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::RuntimeEnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "PlaygroundConversation_runtimeEnvironmentId_fkey"
    )]
    RuntimeEnvironment,
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Run.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RuntimeEnvironment.def()
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
