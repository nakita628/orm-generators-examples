defmodule Shelf.KitCustody do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          custodian: Shelf.TeamMember.t() | nil,
          kit: Shelf.Kit.t() | nil,
          inherited_custody: [Shelf.Custody.t()]
        }

  schema "KitCustody" do
    field(:id, :string, primary_key: true)
    field(:custodian_id, :string, source: :custodianId)
    field(:kit_id, :string, source: :kitId)
    belongs_to(:custodian, Shelf.TeamMember, foreign_key: :custodian_id, define_field: false, type: :string)
    belongs_to(:kit, Shelf.Kit, foreign_key: :kit_id, define_field: false, type: :string)
    has_many(:inherited_custody, Shelf.Custody, foreign_key: :kit_custody_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end