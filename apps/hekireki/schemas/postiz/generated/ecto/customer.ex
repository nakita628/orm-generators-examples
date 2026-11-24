defmodule Postiz.Customer do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          deleted_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Postiz.Organization.t() | nil,
          integrations: [Postiz.Integration.t()]
        }

  schema "Customer" do
    field(:name, :string)
    field(:deleted_at, Postiz.PrismaDateTime, source: :deletedAt)
    field(:org_id, :string, source: :orgId)
    belongs_to(:organization, Postiz.Organization, foreign_key: :org_id, define_field: false)
    has_many(:integrations, Postiz.Integration, foreign_key: :customer_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end