defmodule LangWatch.TwoFactor do
  @moduledoc """
  The two-factor plugin's own table (D06). Protocol state, and row-truth for
  good (ADR-101 R12): the shared secret and the backup codes live here and
  in no event, ever. Nothing in this repo reads or writes these columns -
  the plugin owns them end to end, including the lockout counters, which is
  why `failedVerificationCount` and `lockedUntil` are here rather than
  rebuilt on the projection.

  `backupCodes` holds an ENCRYPTED blob, not hashes. A backup code has to be
  compared against what somebody types, and the plugin decrypts and compares;
  hashing them would leave it unable to verify its own codes.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          secret: String.t(),
          backup_codes: String.t(),
          verified: boolean(),
          failed_verification_count: integer(),
          locked_until: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: LangWatch.User.t() | nil
        }

  schema "TwoFactor" do
    field(:id, :string, primary_key: true)
    field(:secret, :string)
    field(:backup_codes, :string, source: :backupCodes)
    field(:verified, :boolean, default: true)
    field(:failed_verification_count, :integer, default: 0, source: :failedVerificationCount)
    field(:locked_until, LangWatch.PrismaDateTime, source: :lockedUntil)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, LangWatch.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end