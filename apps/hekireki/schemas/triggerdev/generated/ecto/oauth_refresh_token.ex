defmodule Trigger.OAuthRefreshToken do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          hashed_token: String.t(),
          scope: [String.t()],
          resource: String.t(),
          used_at: DateTime.t() | nil,
          replaced_by_id: String.t() | nil,
          idle_expires_at: DateTime.t(),
          inserted_at: DateTime.t(),
          personal_access_token: Trigger.PersonalAccessToken.t() | nil,
          oauth_client: Trigger.OAuthClient.t() | nil
        }

  schema "OAuthRefreshToken" do
    field(:id, :string, primary_key: true)
    field(:hashed_token, :string, source: :hashedToken)
    field(:scope, {:array, :string})
    field(:resource, :string)
    field(:used_at, Trigger.PrismaDateTime, source: :usedAt)
    field(:replaced_by_id, :string, source: :replacedById)
    field(:idle_expires_at, Trigger.PrismaDateTime, source: :idleExpiresAt)
    field(:personal_access_token_id, :string, source: :personalAccessTokenId)
    field(:oauth_client_id, :string, source: :oauthClientId)
    belongs_to(:personal_access_token, Trigger.PersonalAccessToken, foreign_key: :personal_access_token_id, define_field: false, type: :string)
    belongs_to(:oauth_client, Trigger.OAuthClient, foreign_key: :oauth_client_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end