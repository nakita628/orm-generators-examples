defmodule Shelf.MobileAuthCode do
  @moduledoc """
  Native-app SSO web-delegated login — single-use session-handoff codes.
  @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          code_hash: String.t(),
          code_challenge: String.t() | nil,
          session_ciphertext: String.t() | nil,
          consumed_at: DateTime.t() | nil,
          expires_at: DateTime.t(),
          inserted_at: DateTime.t(),
          user: Shelf.User.t() | nil
        }

  schema "MobileAuthCode" do
    field(:id, :string, primary_key: true)
    field(:code_hash, :string, source: :codeHash)
    field(:code_challenge, :string, source: :codeChallenge)
    field(:session_ciphertext, :string, source: :sessionCiphertext)
    field(:consumed_at, Shelf.PrismaDateTime, source: :consumedAt)
    field(:expires_at, Shelf.PrismaDateTime, source: :expiresAt)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end