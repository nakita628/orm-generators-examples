defmodule Umami.Share do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          entity_id: Ecto.UUID.t(),
          name: String.t(),
          share_type: integer(),
          slug: String.t(),
          parameters: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t() | nil,
          updated_at: DateTime.t() | nil
        }

  schema "share" do
    field(:id, Ecto.UUID, primary_key: true, source: :share_id)
    field(:entity_id, Ecto.UUID)
    field(:name, :string)
    field(:share_type, :integer)
    field(:slug, :string)
    field(:parameters, Umami.PrismaJson)
    timestamps(inserted_at_source: :created_at)
  end
end