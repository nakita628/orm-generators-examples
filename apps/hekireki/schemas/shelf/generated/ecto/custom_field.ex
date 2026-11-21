defmodule Shelf.CustomField do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          help_text: String.t() | nil,
          required: boolean(),
          active: boolean(),
          type: atom(),
          options: [String.t()],
          format: atom(),
          deleted_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil,
          created_by: Shelf.User.t() | nil,
          asset_custom_fields_values: [Shelf.AssetCustomFieldValue.t()],
          categories: [Shelf.Category.t()]
        }

  schema "CustomField" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:help_text, :string, source: :helpText)
    field(:required, :boolean, default: false)
    field(:active, :boolean, default: true)
    field(:type, Ecto.Enum, values: [:TEXT, :OPTION, :BOOLEAN, :DATE, :MULTILINE_TEXT, :AMOUNT, :NUMBER], default: :TEXT)
    field(:options, {:array, :string})
    # How a TEXT field's values are read. LINK requires a web link in every
    # value; any other type stays PLAIN (enforced by a CHECK constraint).
    field(:format, Ecto.Enum, values: [:PLAIN, :LINK], default: :PLAIN)
    field(:deleted_at, Shelf.PrismaDateTime, source: :deletedAt)
    field(:organization_id, :string, source: :organizationId)
    field(:user_id, :string, source: :userId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:created_by, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    has_many(:asset_custom_fields_values, Shelf.AssetCustomFieldValue, foreign_key: :custom_field_id, references: :id)
    many_to_many(:categories, Shelf.Category, join_through: "_CategoryToCustomField", join_keys: [B: :id, A: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end