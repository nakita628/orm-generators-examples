defmodule Postiz.Sets do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          content: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Postiz.Organization.t() | nil
        }

  schema "Sets" do
    field(:name, :string)
    field(:content, :string)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Postiz.Organization, foreign_key: :organization_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end