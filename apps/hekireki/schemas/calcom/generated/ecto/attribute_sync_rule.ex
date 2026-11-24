defmodule Calcom.AttributeSyncRule do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Calcom.PrismaDateTime, autogenerate: {Calcom.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          rule: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          integration_attribute_sync: Calcom.IntegrationAttributeSync.t() | nil
        }

  schema "AttributeSyncRule" do
    field(:rule, Calcom.PrismaJson)
    field(:integration_attribute_sync_id, :string, source: :integrationAttributeSyncId)
    belongs_to(:integration_attribute_sync, Calcom.IntegrationAttributeSync, foreign_key: :integration_attribute_sync_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end