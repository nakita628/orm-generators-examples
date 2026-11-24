defmodule Ghostfolio.Access do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Ghostfolio.PrismaDateTime, autogenerate: {Ghostfolio.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          alias: String.t() | nil,
          expires_at: DateTime.t(),
          last_used_at: DateTime.t() | nil,
          scopes: [String.t()],
          settings: map() | list() | String.t() | number() | boolean(),
          type: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          grantee_user: Ghostfolio.User.t() | nil,
          user: Ghostfolio.User.t() | nil
        }

  schema "Access" do
    field(:alias, :string)
    field(:expires_at, Ghostfolio.PrismaDateTime, source: :expiresAt)
    field(:last_used_at, Ghostfolio.PrismaDateTime, source: :lastUsedAt)
    field(:scopes, {:array, :string})
    field(:settings, Ghostfolio.PrismaJson, default: %{})
    field(:type, Ecto.Enum, values: [:MCP, :PRIVATE, :PUBLIC], default: :PRIVATE)
    field(:grantee_user_id, :string, source: :granteeUserId)
    field(:user_id, :string, source: :userId)
    belongs_to(:grantee_user, Ghostfolio.User, foreign_key: :grantee_user_id, define_field: false)
    belongs_to(:user, Ghostfolio.User, foreign_key: :user_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end