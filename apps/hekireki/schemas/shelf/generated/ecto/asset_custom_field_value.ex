defmodule Shelf.AssetCustomFieldValue do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          value: map(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          asset: Shelf.Asset.t() | nil,
          custom_field: Shelf.CustomField.t() | nil
        }

  schema "AssetCustomFieldValue" do
    field(:id, :string, primary_key: true)
    field(:value, :map)
    field(:asset_id, :string, source: :assetId)
    field(:custom_field_id, :string, source: :customFieldId)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:custom_field, Shelf.CustomField, foreign_key: :custom_field_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end