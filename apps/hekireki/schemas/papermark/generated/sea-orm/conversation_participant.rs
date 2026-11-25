use super::participant_role::ParticipantRole;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ConversationParticipant")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "conversationId", column_type = "Text")]
    pub conversation_id: String,
    #[sea_orm(default_value = "PARTICIPANT")]
    pub role: ParticipantRole,
    #[sea_orm(column_name = "viewerId", column_type = "Text")]
    pub viewer_id: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "receiveNotifications", default_value = false)]
    pub receive_notifications: bool,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::conversation::Entity",
        from = "Column::ConversationId",
        to = "super::conversation::Column::Id",
        on_delete = "Cascade",
        fk_name = "ConversationParticipant_conversationId_fkey"
    )]
    Conversation,
    #[sea_orm(
        belongs_to = "super::viewer::Entity",
        from = "Column::ViewerId",
        to = "super::viewer::Column::Id",
        on_delete = "SetNull",
        fk_name = "ConversationParticipant_viewerId_fkey"
    )]
    Viewer,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "ConversationParticipant_userId_fkey"
    )]
    User,
}

impl Related<super::conversation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Conversation.def()
    }
}

impl Related<super::viewer::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Viewer.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
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
