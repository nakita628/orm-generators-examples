use super::queued_email_status::QueuedEmailStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// An email waiting to be sent by the queue, one row per recipient. The
/// email is built when it is sent, from the row its subject points at.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "queued_emails")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// Which handler builds the email. A string rather than an enum so a new
    /// kind needs no migration; the claim only takes kinds the build knows.
    pub kind: String,
    /// The row the email is about, interpreted by the kind. Not a foreign key:
    /// kinds point at different tables, and a subject deleted before the send
    /// makes its handler skip the email.
    pub subject_id: String,
    /// Groups the emails one action queued, so the action can send its own
    /// first batch at once and its progress can be read back.
    pub batch_id: String,
    /// The account whose action queued the email. Banning it skips whatever
    /// is still pending.
    pub user_id: String,
    #[sea_orm(default_value = "pending")]
    pub status: QueuedEmailStatus,
    /// Not claimed before this time. Unset sends at the next run.
    pub send_after: Option<DateTime>,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    pub claimed_at: Option<DateTime>,
    pub sent_at: Option<DateTime>,
    pub last_error: Option<String>,
    pub created_at: DateTime,
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade"
    )]
    User,
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
