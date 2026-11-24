defmodule Postiz.MastraMcpClientVersions do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          mcp_client_id: String.t(),
          version_number: integer(),
          name: String.t(),
          description: String.t() | nil,
          servers: map() | list() | String.t() | number() | boolean(),
          changed_fields: map() | list() | String.t() | number() | boolean() | nil,
          change_message: String.t() | nil,
          created_at: DateTime.t(),
          created_at_z: DateTime.t() | nil
        }

  schema "mastra_mcp_client_versions" do
    field(:id, :string, primary_key: true)
    field(:mcp_client_id, :string, source: :mcpClientId)
    field(:version_number, :integer, source: :versionNumber)
    field(:name, :string)
    field(:description, :string)
    field(:servers, Postiz.PrismaJson)
    field(:changed_fields, Postiz.PrismaJson, source: :changedFields)
    field(:change_message, :string, source: :changeMessage)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
  end
end