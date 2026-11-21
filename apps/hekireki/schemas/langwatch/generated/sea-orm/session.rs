use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Session")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "sessionToken")]
    pub session_token: String,
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    pub expires: DateTime,
    #[sea_orm(column_name = "ipAddress")]
    pub ip_address: Option<String>,
    #[sea_orm(column_name = "userAgent", column_type = "Text")]
    pub user_agent: Option<String>,
    /// Which of the user's sign-in methods minted this session (D06). NULL on
    /// every session that predates the column and on any mint we could not
    /// attribute - an ordinary value, not an error. Nothing reads it to decide
    /// whether the session is valid, so no session was ended to add it.
    #[sea_orm(column_name = "identifierId")]
    pub identifier_id: Option<String>,
    /// What the sign-in that minted this session PROVED, as RFC 8176 method
    /// references (D06). Empty on every session minted before the column, and
    /// on any ordinary password sign-in. An organization that requires a second
    /// factor reads this when a member reaches its data; it is never read at
    /// mint and never ends a session.
    pub amr: Vec<String>,
    /// Minted while the account had never signed in: the session its sign-up
    /// started, which skips the secure-account offer. False on every later one
    /// and on every session that predates the column (specs/identity/passkeys.feature).
    #[sea_orm(column_name = "startedBySignUp", default_value = false)]
    pub started_by_sign_up: bool,
    /// The legacy impersonation payload. NOTHING READS OR WRITES IT any more -
    /// the {actor, subject} columns below replaced it, and the migration beside
    /// them ended every session that still carried a value.
    ///
    /// It stays for ONE release, and that is the whole reason it is still here.
    /// `prisma migrate deploy` runs at container start, so the first pod of a
    /// new release drops a column while the previous release's pods are still
    /// serving - and those pods select this column on EVERY authenticated
    /// request and fail closed when the row does not come back. Dropping it in
    /// the same release that stops reading it signs everybody out for the length
    /// of the rollout. Expand now, contract next release.
    pub impersonating: Option<Json>,
    /// The impersonation claims (D06), aligned with the authz Principal
    /// {actor, subject}. The actor is the operator who really made the request
    /// and the subject is the person whose access they are borrowing; both are
    /// recorded on every authorization decision, so the audit trail can answer
    /// who really did it. These replace the legacy `impersonating` JSON blob,
    /// which is dropped in the release AFTER this one.
    ///
    /// NULL on every ordinary session, which is all of them: an impersonation
    /// is a handful of rows at a time and nothing else reads these columns.
    #[sea_orm(column_name = "actorUserId")]
    pub actor_user_id: Option<String>,
    #[sea_orm(column_name = "subjectUserId")]
    pub subject_user_id: Option<String>,
    /// Why the operator started it. Required to start one
    /// (specs/features/backoffice-user-impersonation-reason.feature) and kept
    /// beside both people rather than only in the audit log, so a request under
    /// impersonation carries its own justification.
    #[sea_orm(column_name = "impersonationReason", column_type = "Text")]
    pub impersonation_reason: Option<String>,
    /// When the borrowed access lapses. A session outlives its impersonation
    /// window: past this the claims are ignored and the operator is themselves
    /// again, with nothing ended.
    #[sea_orm(column_name = "impersonationExpiresAt")]
    pub impersonation_expires_at: Option<DateTime>,
    /// When this session was last used, for the idle timeout an organization
    /// may set (GAC-10, specs/identity/org-session-lifetime.feature).
    ///
    /// A COLUMN OF OUR OWN rather than `updatedAt`, because better-auth rolls a
    /// live session's expiry once per `updateAge` - a day - so `updatedAt` is
    /// activity to the nearest day and cannot tell an idle hour from an idle
    /// minute. Written only for sessions whose owner belongs to an organization
    /// that has set a window, and then at most once per quarter-window, so a
    /// deployment where nobody has turned this on writes it never.
    ///
    /// NULL on every session minted before the column and on every session
    /// under no window at all. A null READS AS `updatedAt`, which better-auth
    /// keeps rolled to within a day of the last use. That fallback is
    /// conservative in the direction that matters - `updatedAt` is never
    /// earlier than the real last use, so it can only ever judge a session MORE
    /// active than it was, never less - and it is what lets a window turned on
    /// today end the sessions that have been idle since February instead of
    /// waiting a day for a stamp they have never had.
    #[sea_orm(column_name = "lastSeenAt")]
    pub last_seen_at: Option<DateTime>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
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
