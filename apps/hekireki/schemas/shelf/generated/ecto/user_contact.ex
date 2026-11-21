defmodule Shelf.UserContact do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          phone: String.t() | nil,
          street: String.t() | nil,
          city: String.t() | nil,
          state_province: String.t() | nil,
          zip_postal_code: String.t() | nil,
          country_region: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil
        }

  schema "UserContact" do
    field(:id, :string, primary_key: true)
    field(:phone, :string)
    field(:street, :string)
    field(:city, :string)
    field(:state_province, :string, source: :stateProvince)
    field(:zip_postal_code, :string, source: :zipPostalCode)
    field(:country_region, :string, source: :countryRegion)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end