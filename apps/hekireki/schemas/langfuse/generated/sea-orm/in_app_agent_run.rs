use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "in_app_agent_runs")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub conversation_id: String,
    #[sea_orm(column_type = "Text")]
    pub triggered_by_user_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub model: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub mcp_api_key_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub error_code: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub error_message: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub finished_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub status: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub request: Option<Json>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub claimed_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub heartbeat_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub cancel_requested_at: Option<DateTime>,
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
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "in_app_agent_runs_project_id_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::in_app_agent_conversation::Entity",
        from = "(Column::ConversationId, Column::ProjectId)",
        to = "(super::in_app_agent_conversation::Column::Id, super::in_app_agent_conversation::Column::ProjectId)",
        on_delete = "Cascade",
        fk_name = "in_app_agent_runs_conversation_id_project_id_fkey"
    )]
    Conversation,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::TriggeredByUserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "in_app_agent_runs_triggered_by_user_id_fkey"
    )]
    TriggeredByUser,
    #[sea_orm(has_many = "super::in_app_agent_event::Entity")]
    Events,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::in_app_agent_conversation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Conversation.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TriggeredByUser.def()
    }
}

impl Related<super::in_app_agent_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Events.def()
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
