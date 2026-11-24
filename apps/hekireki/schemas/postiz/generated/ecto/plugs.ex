defmodule Postiz.Plugs do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          plug_function: String.t(),
          data: String.t(),
          activated: boolean(),
          integration: Postiz.Integration.t() | nil,
          organization: Postiz.Organization.t() | nil
        }

  schema "Plugs" do
    field(:plug_function, :string, source: :plugFunction)
    field(:data, :string)
    field(:activated, :boolean, default: true)
    field(:integration_id, :string, source: :integrationId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:integration, Postiz.Integration, foreign_key: :integration_id, define_field: false)
    belongs_to(:organization, Postiz.Organization, foreign_key: :organization_id, define_field: false)
  end
end