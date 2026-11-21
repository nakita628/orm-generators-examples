defmodule Shelf.ReportFound do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          email: String.t(),
          content: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          asset: Shelf.Asset.t() | nil,
          kit: Shelf.Kit.t() | nil
        }

  schema "ReportFound" do
    field(:id, :string, primary_key: true)
    field(:email, :string)
    field(:content, :string)
    field(:asset_id, :string, source: :assetId)
    field(:kit_id, :string, source: :kitId)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:kit, Shelf.Kit, foreign_key: :kit_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end