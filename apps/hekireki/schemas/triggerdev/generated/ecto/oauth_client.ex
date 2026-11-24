defmodule Trigger.OAuthClient do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          kind: atom(),
          client_id: String.t(),
          metadata_url: String.t() | nil,
          display_name: String.t() | nil,
          self_reported_name: String.t() | nil,
          redirect_uris: [String.t()],
          dedupe_key: String.t() | nil,
          metadata_expires_at: DateTime.t() | nil,
          last_used_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          authorization_codes: [Trigger.AuthorizationCode.t()],
          personal_access_tokens: [Trigger.PersonalAccessToken.t()],
          refresh_tokens: [Trigger.OAuthRefreshToken.t()]
        }

  schema "OAuthClient" do
    field(:id, :string, primary_key: true)
    field(:kind, Ecto.Enum, values: [:PREREGISTERED, :CIMD, :DCR])
    field(:client_id, :string, source: :clientId)
    field(:metadata_url, :string, source: :metadataUrl)
    field(:display_name, :string, source: :displayName)
    field(:self_reported_name, :string, source: :selfReportedName)
    field(:redirect_uris, {:array, :string}, source: :redirectUris)
    field(:dedupe_key, :string, source: :dedupeKey)
    field(:metadata_expires_at, Trigger.PrismaDateTime, source: :metadataExpiresAt)
    field(:last_used_at, Trigger.PrismaDateTime, source: :lastUsedAt)
    has_many(:authorization_codes, Trigger.AuthorizationCode, foreign_key: :oauth_client_id, references: :id)
    has_many(:personal_access_tokens, Trigger.PersonalAccessToken, foreign_key: :oauth_client_id, references: :id)
    has_many(:refresh_tokens, Trigger.OAuthRefreshToken, foreign_key: :oauth_client_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end