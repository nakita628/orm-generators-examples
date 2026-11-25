use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Chat")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub title: Option<String>,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
    #[sea_orm(column_name = "documentId", column_type = "Text")]
    pub document_id: Option<String>,
    #[sea_orm(column_name = "dataroomId", column_type = "Text")]
    pub dataroom_id: Option<String>,
    #[sea_orm(column_name = "linkId", column_type = "Text")]
    pub link_id: Option<String>,
    #[sea_orm(column_name = "viewId", column_type = "Text")]
    pub view_id: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "viewerId", column_type = "Text")]
    pub viewer_id: Option<String>,
    #[sea_orm(column_name = "vectorStoreId", column_type = "Text")]
    pub vector_store_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(
        column_name = "lastMessageAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_message_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "Chat_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::document::Entity",
        from = "Column::DocumentId",
        to = "super::document::Column::Id",
        on_delete = "Cascade",
        fk_name = "Chat_documentId_fkey"
    )]
    Document,
    #[sea_orm(
        belongs_to = "super::dataroom::Entity",
        from = "Column::DataroomId",
        to = "super::dataroom::Column::Id",
        on_delete = "Cascade",
        fk_name = "Chat_dataroomId_fkey"
    )]
    Dataroom,
    #[sea_orm(
        belongs_to = "super::link::Entity",
        from = "Column::LinkId",
        to = "super::link::Column::Id",
        on_delete = "Cascade",
        fk_name = "Chat_linkId_fkey"
    )]
    Link,
    #[sea_orm(
        belongs_to = "super::view::Entity",
        from = "Column::ViewId",
        to = "super::view::Column::Id",
        on_delete = "Cascade",
        fk_name = "Chat_viewId_fkey"
    )]
    View,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Chat_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::viewer::Entity",
        from = "Column::ViewerId",
        to = "super::viewer::Column::Id",
        on_delete = "Cascade",
        fk_name = "Chat_viewerId_fkey"
    )]
    Viewer,
    #[sea_orm(has_many = "super::chat_message::Entity")]
    Messages,
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::document::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Document.def()
    }
}

impl Related<super::dataroom::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dataroom.def()
    }
}

impl Related<super::link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Link.def()
    }
}

impl Related<super::view::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::View.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::viewer::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Viewer.def()
    }
}

impl Related<super::chat_message::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Messages.def()
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
