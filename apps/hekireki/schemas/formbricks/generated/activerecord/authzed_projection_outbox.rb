# Durable PostgreSQL queue for rebuilding SpiceDB relationships after an authorization source changes.
# Source-table triggers insert these rows in the same transaction as the mutation. Identifiers are
# operationally sensitive and must never be included in logs, metrics, or CLI output.
#
# This model declares NO indexes on purpose. All three live in
# `migration/20260818120000_add_authzed_projection_outbox/migration.sql`, because each one is
# partial (`WHERE "processedAt" IS NULL AND ...`) and Prisma's `@@index` has no `where`. Declaring
# approximate non-partial copies here would be worse than declaring none: `prisma db push` would
# see a name match with a different definition, drop and recreate them without the predicate, and
# the migration's `CREATE INDEX IF NOT EXISTS` would then skip them — leaving a silently wrong
# shape that carries a week of delivered rows on every hot-path lookup. Deployments are unaffected
# either way (`prisma migrate deploy` never reads this file); a dev who runs `db push` restores
# them by rerunning the migration, which is the same contract its triggers already live under.
#
# attempts:: Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried.
# permanentFailures:: Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
#   that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
class AuthzedProjectionOutbox < ApplicationRecord
  self.table_name = "AuthzedProjectionOutbox"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { SecureRandom.uuid }
  attribute :isRevocation, default: false
  attribute :attempts, default: 0
  attribute :permanentFailures, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :targetType, presence: true
  validates :primaryId, presence: true
end
