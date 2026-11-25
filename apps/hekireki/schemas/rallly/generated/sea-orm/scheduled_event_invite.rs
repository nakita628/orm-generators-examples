use super::scheduled_event_invite_status::ScheduledEventInviteStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "scheduled_event_invites")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub uid: String,
    #[sea_orm(column_type = "Text")]
    pub scheduled_event_id: String,
    #[sea_orm(column_type = "Text")]
    pub invitee_name: String,
    #[sea_orm(column_type = "custom(\"citext\")")]
    pub invitee_email: String,
    #[sea_orm(column_type = "Text")]
    pub invitee_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub invitee_time_zone: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub invitee_locale: Option<String>,
    #[sea_orm(default_value = "pending")]
    pub status: ScheduledEventInviteStatus,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::scheduled_event::Entity",
        from = "Column::ScheduledEventId",
        to = "super::scheduled_event::Column::Id",
        on_delete = "Cascade",
        fk_name = "scheduled_event_invites_scheduled_event_id_fkey"
    )]
    ScheduledEvent,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::InviteeId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "scheduled_event_invites_invitee_id_fkey"
    )]
    User,
}

impl Related<super::scheduled_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScheduledEvent.def()
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
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
