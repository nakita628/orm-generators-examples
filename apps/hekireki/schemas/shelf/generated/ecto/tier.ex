defmodule Shelf.Tier do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: atom(),
          name: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          tier_limit: Shelf.TierLimit.t() | nil,
          subscribers: [Shelf.User.t()]
        }

  schema "Tier" do
    field(:id, Ecto.Enum, values: [:free, :tier_1, :tier_2, :custom])
    field(:name, :string)
    field(:tier_limit_id, :string, source: :tierLimitId)
    belongs_to(:tier_limit, Shelf.TierLimit, foreign_key: :tier_limit_id, define_field: false, type: :string)
    has_many(:subscribers, Shelf.User, foreign_key: :tier_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end