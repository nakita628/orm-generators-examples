defmodule Formbricks.TwoFactor do
  @moduledoc """
  Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
  User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
  at rest by Better Auth using BETTER_AUTH_SECRET.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          secret: String.t(),
          backup_codes: String.t(),
          verified: boolean(),
          failed_verification_count: integer(),
          locked_until: DateTime.t() | nil,
          user: Formbricks.User.t() | nil
        }

  schema "TwoFactor" do
    field(:id, :string, primary_key: true)
    field(:secret, :string)
    field(:backup_codes, :string, source: :backupCodes)
    field(:verified, :boolean, default: true)
    # Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts.
    field(:failed_verification_count, :integer, default: 0, source: :failedVerificationCount)
    # Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit.
    field(:locked_until, Formbricks.PrismaDateTime, source: :lockedUntil)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Formbricks.User, foreign_key: :user_id, define_field: false, type: :string)
  end
end