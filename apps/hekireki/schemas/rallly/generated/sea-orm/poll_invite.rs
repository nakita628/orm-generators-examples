use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "poll_invites")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub poll_id: String,
    pub email: String,
    /// Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access
    #[sea_orm(unique)]
    pub token: String,
    pub opened_at: Option<DateTime>,
    pub reminded_at: Option<DateTime>,
    /// Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token
    pub revoked_at: Option<DateTime>,
    /// Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite
    #[sea_orm(unique)]
    pub participant_id: Option<String>,
    pub created_at: DateTime,
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::poll::Entity",
        from = "Column::PollId",
        to = "super::poll::Column::Id",
        on_delete = "Cascade"
    )]
    Poll,
    #[sea_orm(
        belongs_to = "super::participant::Entity",
        from = "Column::ParticipantId",
        to = "super::participant::Column::Id",
        on_delete = "SetNull"
    )]
    Participant,
}

impl Related<super::poll::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Poll.def()
    }
}

impl Related<super::participant::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Participant.def()
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
