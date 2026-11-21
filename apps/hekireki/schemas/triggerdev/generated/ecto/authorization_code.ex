defmodule Trigger.AuthorizationCode do
  @moduledoc """
  Used to generate PersonalAccessTokens, they're one-time use
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          code: String.t(),
          requester_ip: String.t() | nil,
          code_challenge: String.t() | nil,
          callback_port: integer() | nil,
          approval_code_hash: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          personal_access_token: Trigger.PersonalAccessToken.t() | nil
        }

  schema "AuthorizationCode" do
    field(:id, :string, primary_key: true)
    field(:code, :string)
    # Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires.
    field(:requester_ip, :string, source: :requesterIp)
    # S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it.
    field(:code_challenge, :string, source: :codeChallenge)
    # Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it.
    field(:callback_port, :integer, source: :callbackPort)
    # sha256 of the approval code, which only the approving browser ever sees.
    field(:approval_code_hash, :string, source: :approvalCodeHash)
    field(:personal_access_token_id, :string, source: :personalAccessTokenId)
    belongs_to(:personal_access_token, Trigger.PersonalAccessToken, foreign_key: :personal_access_token_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end