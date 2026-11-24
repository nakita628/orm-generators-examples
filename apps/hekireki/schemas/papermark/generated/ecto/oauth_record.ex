defmodule Papermark.OAuthRecord do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Papermark.PrismaDateTime, autogenerate: {Papermark.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          kind: String.t(),
          payload: map() | list() | String.t() | number() | boolean(),
          grant_id: String.t() | nil,
          user_code: String.t() | nil,
          uid: String.t() | nil,
          expires_at: DateTime.t() | nil,
          consumed_at: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "OAuthRecord" do
    field(:id, :string, primary_key: true)
    field(:kind, :string)
    field(:payload, Papermark.PrismaJson)
    field(:grant_id, :string, source: :grantId)
    field(:user_code, :string, source: :userCode)
    field(:uid, :string)
    field(:expires_at, Papermark.PrismaDateTime, source: :expiresAt)
    field(:consumed_at, Papermark.PrismaDateTime, source: :consumedAt)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end