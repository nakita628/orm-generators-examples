use super::invitation_status::InvitationStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ViewerInvitation")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "viewerId", column_type = "Text")]
    pub viewer_id: String,
    #[sea_orm(column_name = "linkId", column_type = "Text")]
    pub link_id: String,
    #[sea_orm(column_name = "groupId", column_type = "Text")]
    pub group_id: Option<String>,
    #[sea_orm(column_name = "invitedBy", column_type = "Text")]
    pub invited_by: String,
    #[sea_orm(column_name = "customMessage", column_type = "Text")]
    pub custom_message: Option<String>,
    #[sea_orm(
        column_name = "sentAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub sent_at: DateTime,
    #[sea_orm(default_value = "SENT")]
    pub status: InvitationStatus,
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
        belongs_to = "super::viewer::Entity",
        from = "Column::ViewerId",
        to = "super::viewer::Column::Id",
        on_delete = "Cascade",
        fk_name = "ViewerInvitation_viewerId_fkey"
    )]
    Viewer,
    #[sea_orm(
        belongs_to = "super::link::Entity",
        from = "Column::LinkId",
        to = "super::link::Column::Id",
        on_delete = "Cascade",
        fk_name = "ViewerInvitation_linkId_fkey"
    )]
    Link,
    #[sea_orm(
        belongs_to = "super::viewer_group::Entity",
        from = "Column::GroupId",
        to = "super::viewer_group::Column::Id",
        on_delete = "SetNull",
        fk_name = "ViewerInvitation_groupId_fkey"
    )]
    Group,
}

impl Related<super::viewer::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Viewer.def()
    }
}

impl Related<super::link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Link.def()
    }
}

impl Related<super::viewer_group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Group.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.sent_at.is_not_set() {
            self.sent_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
