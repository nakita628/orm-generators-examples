use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Consecutive failed sign-ins, and the lock they earn (GAC-09).
/// Spec: specs/identity/org-account-lockout.feature.
///
/// KEYED ON THE ADDRESS THAT WAS TYPED, NOT ON A USER, and that is the whole
/// design. We already answer "no such account" and "wrong password" with one
/// refusal so that sign-in cannot be used to discover who has an account. A
/// lock that existed only for real accounts would put that back: five wrong
/// guesses saying "locked" for one address and "incorrect" for another tells
/// an attacker which addresses are worth attacking. So an address with
/// nothing behind it gets a row, a count and a lock exactly like one that
/// resolves, and both answer identically.
///
/// `identifierHash` is a KEYED hash of the normalised address, never the
/// address. The key is the deployment's own secret, so this table cannot be
/// lifted and reversed into the list of every address anybody has ever tried
/// to sign in as - which is what it would otherwise become, including every
/// typo and every address that has no account here at all.
///
/// Row-truth, deliberately (ADR-101 R12): this is protocol state read and
/// written on the sign-in path, like TwoFactor's own counters next door. The
/// EVIDENCE an auditor reads is the identity event emitted when a lock is
/// taken, which is durable; these rows are reaped once they stop meaning
/// anything.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SignInAttemptLock")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// Keyed hash of the normalised address. Never the address itself.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "identifierHash", column_type = "Text")]
    pub identifier_hash: String,
    /// Consecutive failures since the last success. Reset by a success, not by
    /// time - "five consecutive", not five ever.
    #[sea_orm(column_name = "failedCount", default_value = 0)]
    pub failed_count: i32,
    /// When the current lock lifts. Null when not locked.
    #[sea_orm(column_name = "lockedUntil", column_type = "custom(\"timestamp(3)\")")]
    pub locked_until: Option<DateTime>,
    /// Locks taken since the last success. At five this stops lifting by
    /// itself and becomes an incident - see `heldForReview`.
    #[sea_orm(column_name = "consecutiveLockouts", default_value = 0)]
    pub consecutive_lockouts: i32,
    /// Held past its lock because the count of locks reached the ceiling the
    /// control names. Cleared by an administrator, or by the owner proving
    /// mailbox control through a password reset - the second is what keeps a
    /// held account from being a denial of service anyone can inflict by
    /// typing a colleague's address wrong often enough.
    #[sea_orm(column_name = "heldForReview", default_value = false)]
    pub held_for_review: bool,
    /// Whose account this resolved to at the time, when it resolved to one.
    /// Null for an address with no account - those rows are the reason this is
    /// nullable, and the reason this table is not hung off User.
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
