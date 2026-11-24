defmodule Shelf.AssetModelCustomFieldDefault do
  @moduledoc """
  A custom field value an AssetModel gives the assets that join it.

  `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
  asset is a plain copy. Rows cascade away with their model or field.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          value: map() | list() | String.t() | number() | boolean(),
          custom_field_id: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          asset_model: Shelf.AssetModel.t() | nil
        }

  schema "AssetModelCustomFieldDefault" do
    field(:id, :string, primary_key: true)
    field(:value, Shelf.PrismaJson)
    field(:custom_field_id, :string, source: :customFieldId)
    field(:asset_model_id, :string, source: :assetModelId)
    belongs_to(:asset_model, Shelf.AssetModel, foreign_key: :asset_model_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end