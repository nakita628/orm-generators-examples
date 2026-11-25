use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ConversationComment")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "conversationId", column_type = "Text")]
    pub conversation_id: String,
    #[sea_orm(column_name = "authorMemberId", column_type = "Text")]
    pub author_member_id: Option<String>,
    #[sea_orm(column_name = "clientMutationId", column_type = "Text")]
    pub client_mutation_id: String,
    #[sea_orm(column_type = "Text")]
    pub body: String,
    pub revision: i32,
    #[sea_orm(column_name = "deletedAt", column_type = "custom(\"timestamp(3)\")")]
    pub deleted_at: Option<DateTime>,
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
        fk_name = "ConversationComment_conversationId_fkey"
    )]
    Conversation,
    #[sea_orm(
        belongs_to = "super::member::Entity",
        from = "Column::AuthorMemberId",
        to = "super::member::Column::Id",
        on_delete = "SetNull",
        fk_name = "ConversationComment_authorMemberId_fkey"
    )]
    Author,
    #[sea_orm(has_many = "super::conversation_activity::Entity")]
    Activity,
}

impl Related<super::conversation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Conversation.def()
    }
}

impl Related<super::member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Author.def()
    }
}

impl Related<super::conversation_activity::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Activity.def()
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
