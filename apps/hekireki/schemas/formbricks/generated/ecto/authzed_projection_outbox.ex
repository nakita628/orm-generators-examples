defmodule Formbricks.AuthzedProjectionOutbox do
  @moduledoc """
  Durable PostgreSQL queue for rebuilding SpiceDB relationships after an authorization source changes.
  Source-table triggers insert these rows in the same transaction as the mutation. Identifiers are
  operationally sensitive and must never be included in logs, metrics, or CLI output.

  This model declares NO indexes on purpose. All three live in
  `migration/20260818120000_add_authzed_projection_outbox/migration.sql`, because each one is
  partial (`WHERE "processedAt" IS NULL AND ...`) and Prisma's `@@index` has no `where`. Declaring
  approximate non-partial copies here would be worse than declaring none: `prisma db push` would
  see a name match with a different definition, drop and recreate them without the predicate, and
  the migration's `CREATE INDEX IF NOT EXISTS` would then skip them — leaving a silently wrong
  shape that carries a week of delivered rows on every hot-path lookup. Deployments are unaffected
  either way (`prisma migrate deploy` never reads this file); a dev who runs `db push` restores
  them by rerunning the migration, which is the same contract its triggers already live under.
  """

  use Ecto.Schema

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          target_type: String.t(),
          primary_id: String.t(),
          secondary_id: String.t() | nil,
          is_revocation: boolean(),
          attempts: integer(),
          permanent_failures: integer(),
          available_at: DateTime.t(),
          leased_at: DateTime.t() | nil,
          lease_expires_at: DateTime.t() | nil,
          lease_owner: String.t() | nil,
          processed_at: DateTime.t() | nil,
          dead_lettered_at: DateTime.t() | nil,
          last_attempt_at: DateTime.t() | nil,
          last_error_code: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "AuthzedProjectionOutbox" do
    field(:target_type, :string, source: :targetType)
    field(:primary_id, :string, source: :primaryId)
    field(:secondary_id, :string, source: :secondaryId)
    field(:is_revocation, :boolean, default: false, source: :isRevocation)
    # Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried.
    field(:attempts, :integer, default: 0)
    # Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
    # that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
    field(:permanent_failures, :integer, default: 0, source: :permanentFailures)
    field(:available_at, Formbricks.PrismaDateTime, autogenerate: true, source: :availableAt)
    field(:leased_at, Formbricks.PrismaDateTime, source: :leasedAt)
    field(:lease_expires_at, Formbricks.PrismaDateTime, source: :leaseExpiresAt)
    field(:lease_owner, :string, source: :leaseOwner)
    field(:processed_at, Formbricks.PrismaDateTime, source: :processedAt)
    field(:dead_lettered_at, Formbricks.PrismaDateTime, source: :deadLetteredAt)
    field(:last_attempt_at, Formbricks.PrismaDateTime, source: :lastAttemptAt)
    field(:last_error_code, :string, source: :lastErrorCode)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end