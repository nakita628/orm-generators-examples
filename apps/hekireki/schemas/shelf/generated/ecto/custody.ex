defmodule Shelf.Custody do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          quantity: integer(),
          source_unknown: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          custodian: Shelf.TeamMember.t() | nil,
          asset: Shelf.Asset.t() | nil,
          kit_custody: Shelf.KitCustody.t() | nil,
          location: Shelf.Location.t() | nil
        }

  schema "Custody" do
    field(:id, :string, primary_key: true)
    field(:quantity, :integer, default: 1)
    field(:source_unknown, :boolean, default: false, source: :sourceUnknown)
    field(:team_member_id, :string, source: :teamMemberId)
    field(:asset_id, :string, source: :assetId)
    field(:kit_custody_id, :string, source: :kitCustodyId)
    field(:location_id, :string, source: :locationId)
    belongs_to(:custodian, Shelf.TeamMember, foreign_key: :team_member_id, define_field: false, type: :string)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:kit_custody, Shelf.KitCustody, foreign_key: :kit_custody_id, define_field: false, type: :string)
    belongs_to(:location, Shelf.Location, foreign_key: :location_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end