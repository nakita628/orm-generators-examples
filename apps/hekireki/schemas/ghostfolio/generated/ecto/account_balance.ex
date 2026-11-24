defmodule Ghostfolio.AccountBalance do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Ghostfolio.PrismaDateTime, autogenerate: {Ghostfolio.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          account_id: String.t(),
          comment: String.t() | nil,
          date: DateTime.t(),
          user_id: String.t(),
          value: float(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "AccountBalance" do
    field(:account_id, :string, source: :accountId)
    field(:comment, :string)
    field(:date, Ghostfolio.PrismaDateTime, autogenerate: true)
    field(:user_id, :string, source: :userId)
    field(:value, :float)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end