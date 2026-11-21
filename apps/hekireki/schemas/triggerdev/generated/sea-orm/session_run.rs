use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Historical record of every run a Session has owned. Append-only —
/// rows are inserted on each `ensureRunForSession` claim, never updated.
/// Lets us reconstruct the run timeline of a chat for debugging /
/// dashboard surfaces. The relation cascades on Session delete (tied to
/// the session lifecycle) but `runId` is a plain string column with no
/// FK to TaskRun so run pruning is independent.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SessionRun")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "sessionId")]
    pub session_id: String,
    /// TaskRun.id (no FK — runs may be archived independently of session history)
    #[sea_orm(unique)]
    #[sea_orm(column_name = "runId")]
    pub run_id: String,
    /// One of: "initial" | "continuation" | "upgrade" | "manual".
    /// Plain string for forward-compat with future trigger reasons.
    pub reason: String,
    #[sea_orm(column_name = "triggeredAt")]
    pub triggered_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::session::Entity",
        from = "Column::SessionId",
        to = "super::session::Column::Id",
        on_delete = "Cascade"
    )]
    Session,
}

impl Related<super::session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Session.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.triggered_at.is_not_set() {
            self.triggered_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
