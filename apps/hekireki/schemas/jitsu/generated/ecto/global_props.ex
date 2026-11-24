defmodule Jitsu.GlobalProps do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Jitsu.PrismaDateTime, autogenerate: {Jitsu.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          value: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "GlobalProps" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:value, Jitsu.PrismaJson)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end