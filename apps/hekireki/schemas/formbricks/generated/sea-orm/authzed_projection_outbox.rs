use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Durable PostgreSQL queue for rebuilding SpiceDB relationships after an authorization source changes.
/// Source-table triggers insert these rows in the same transaction as the mutation. Identifiers are
/// operationally sensitive and must never be included in logs, metrics, or CLI output.
///
/// This model declares NO indexes on purpose. All three live in
/// `migration/20260818120000_add_authzed_projection_outbox/migration.sql`, because each one is
/// partial (`WHERE "processedAt" IS NULL AND ...`) and Prisma's `@@index` has no `where`. Declaring
/// approximate non-partial copies here would be worse than declaring none: `prisma db push` would
/// see a name match with a different definition, drop and recreate them without the predicate, and
/// the migration's `CREATE INDEX IF NOT EXISTS` would then skip them — leaving a silently wrong
/// shape that carries a week of delivered rows on every hot-path lookup. Deployments are unaffected
/// either way (`prisma migrate deploy` never reads this file); a dev who runs `db push` restores
/// them by rerunning the migration, which is the same contract its triggers already live under.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuthzedProjectionOutbox")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "targetType", column_type = "Text")]
    pub target_type: String,
    #[sea_orm(column_name = "primaryId", column_type = "Text")]
    pub primary_id: String,
    #[sea_orm(column_name = "secondaryId", column_type = "Text")]
    pub secondary_id: Option<String>,
    #[sea_orm(column_name = "isRevocation", default_value = false)]
    pub is_revocation: bool,
    /// Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried.
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    /// Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
    /// that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
    #[sea_orm(column_name = "permanentFailures", default_value = 0)]
    pub permanent_failures: i32,
    #[sea_orm(
        column_name = "availableAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub available_at: DateTime,
    #[sea_orm(column_name = "leasedAt", column_type = "custom(\"timestamp(3)\")")]
    pub leased_at: Option<DateTime>,
    #[sea_orm(
        column_name = "leaseExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub lease_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "leaseOwner", column_type = "Text")]
    pub lease_owner: Option<String>,
    #[sea_orm(column_name = "processedAt", column_type = "custom(\"timestamp(3)\")")]
    pub processed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "deadLetteredAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub dead_lettered_at: Option<DateTime>,
    #[sea_orm(
        column_name = "lastAttemptAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_attempt_at: Option<DateTime>,
    #[sea_orm(column_name = "lastErrorCode", column_type = "Text")]
    pub last_error_code: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.available_at.is_not_set() {
            self.available_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
