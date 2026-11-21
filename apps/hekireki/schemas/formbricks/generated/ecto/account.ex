defmodule Formbricks.Account do
  @moduledoc """
  Stores third-party authentication account information.
  Enables OAuth and other external authentication methods.

  @property id - Unique identifier for the account
  @property user - The Formbricks user who owns this account
  @property provider - The authentication provider (GitHub, Google, etc.)
  @property providerAccountId - User ID from the provider
  @property access_token - OAuth access token
  @property refresh_token - OAuth refresh token
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          type: String.t() | nil,
          provider: String.t(),
          provider_account_id: String.t(),
          access_token: String.t() | nil,
          refresh_token: String.t() | nil,
          expires_at: integer() | nil,
          ext_expires_in: integer() | nil,
          token_type: String.t() | nil,
          scope: String.t() | nil,
          id_token: String.t() | nil,
          session_state: String.t() | nil,
          password: String.t() | nil,
          access_token_expires_at: DateTime.t() | nil,
          refresh_token_expires_at: DateTime.t() | nil,
          issuer: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Formbricks.User.t() | nil
        }

  schema "Account" do
    field(:id, :string, primary_key: true)
    field(:type, :string)
    field(:provider, :string)
    field(:provider_account_id, :string, source: :providerAccountId)
    field(:access_token, :string)
    field(:refresh_token, :string)
    field(:expires_at, :integer)
    field(:ext_expires_in, :integer)
    field(:token_type, :string)
    field(:scope, :string)
    field(:id_token, :string)
    field(:session_state, :string)
    field(:password, :string)
    field(:access_token_expires_at, Formbricks.PrismaDateTime, source: :accessTokenExpiresAt)
    field(:refresh_token_expires_at, Formbricks.PrismaDateTime, source: :refreshTokenExpiresAt)
    # Better Auth 1.7.0 (ENG-2343): the account's identity namespace. 1.7.0 keyed accounts on
    # (issuer, accountId) and filtered every lookup on it — findAccountByKey, findAccountOwnerByKey,
    # findCredentialAccount and updatePassword — so on 1.7.0 a null here locks the user out.
    # 1.7.3 dropped the field and 1.7.7 keys accounts on (provider, providerAccountId) again
    # (ENG-3880). We keep writing it (see `account.create.before` in better-auth-hooks.ts) so a
    # rollback to 1.7.0 still finds every row; drop it once that rollback is off the table.
    #
    # NULLABLE on purpose. A NOT NULL column would fail every INSERT from a 1.6 pod still serving
    # during a rolling deploy. The migration backfills every existing row; see the migration SQL for
    # the value derivation.
    field(:issuer, :string)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Formbricks.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end