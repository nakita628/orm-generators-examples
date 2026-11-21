defmodule LangWatch.MfaEnrollment do
  @moduledoc """
  The two-step verification pipeline's projection (D06 - see
  specs/identity/mfa-and-session-shape.feature). One row per PERSON, because
  one person has one setup however many organizations they belong to.

  A pure event-truth head like `SsoConnection`: every column is fold-written,
  rows are never deleted (DISABLED is a tombstone, so history survives being
  turned off), and a replay rebuilds it whole-row from the event log. The
  cursor columns are inline because the fold produces exactly one row per
  aggregate - the `Identifier` fan-out needed its own cursor table, this
  does not.

  Every column here is a lifecycle fact. There is deliberately nowhere in
  this model a secret or a backup code could be put: `consumedBackupCodeIndexes`
  holds POSITIONS, which is what lets "how many are left" be answered
  without the projection ever knowing a code.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          user_id: String.t(),
          enrollment_id: String.t() | nil,
          method: String.t() | nil,
          state: String.t(),
          enrolled_at: DateTime.t() | nil,
          confirmed_at: DateTime.t() | nil,
          expired_at: DateTime.t() | nil,
          disabled_at: DateTime.t() | nil,
          disabled_via: String.t() | nil,
          backup_code_count: integer(),
          consumed_backup_code_indexes: [integer()],
          failed_count: integer(),
          occurred_at: DateTime.t(),
          last_event_id: String.t(),
          accepted_at: DateTime.t(),
          projection_version: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "MfaEnrollment" do
    # The user. Also the aggregate id and the tenant - one identity history
    # per person (ADR-029 §4).
    field(:user_id, :string, primary_key: true, source: :userId)
    # The current enrollment, or NULL for somebody who never started one.
    field(:enrollment_id, :string, source: :enrollmentId)
    # totp. Never sms.
    field(:method, :string)
    # NONE | PENDING | ENABLED | EXPIRED | DISABLED
    field(:state, :string)
    field(:enrolled_at, LangWatch.PrismaDateTime, source: :enrolledAt)
    field(:confirmed_at, LangWatch.PrismaDateTime, source: :confirmedAt)
    field(:expired_at, LangWatch.PrismaDateTime, source: :expiredAt)
    field(:disabled_at, LangWatch.PrismaDateTime, source: :disabledAt)
    # password+totp | org-admin - who ended it, and on whose authority.
    field(:disabled_via, :string, source: :disabledVia)
    field(:backup_code_count, :integer, default: 0, source: :backupCodeCount)
    # Positions already spent, ascending. Never codes.
    field(:consumed_backup_code_indexes, {:array, :integer}, source: :consumedBackupCodeIndexes)
    field(:failed_count, :integer, default: 0, source: :failedCount)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:accepted_at, LangWatch.PrismaDateTime, source: :acceptedAt)
    field(:projection_version, :string, source: :projectionVersion)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end