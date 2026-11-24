defmodule Calcom.AttributeSyncFieldMapping do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Calcom.PrismaDateTime, autogenerate: {Calcom.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          integration_field_name: String.t(),
          enabled: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          attribute: Calcom.Attribute.t() | nil,
          integration_attribute_sync: Calcom.IntegrationAttributeSync.t() | nil
        }

  schema "AttributeSyncFieldMapping" do
    field(:integration_field_name, :string, source: :integrationFieldName)
    field(:enabled, :boolean)
    field(:attribute_id, :string, source: :attributeId)
    field(:integration_attribute_sync_id, :string, source: :integrationAttributeSyncId)
    belongs_to(:attribute, Calcom.Attribute, foreign_key: :attribute_id, define_field: false)
    belongs_to(:integration_attribute_sync, Calcom.IntegrationAttributeSync, foreign_key: :integration_attribute_sync_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end