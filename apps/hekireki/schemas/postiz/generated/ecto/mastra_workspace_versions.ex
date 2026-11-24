defmodule Postiz.MastraWorkspaceVersions do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          workspace_id: String.t(),
          version_number: integer(),
          name: String.t(),
          description: String.t() | nil,
          filesystem: map() | list() | String.t() | number() | boolean() | nil,
          sandbox: map() | list() | String.t() | number() | boolean() | nil,
          mounts: map() | list() | String.t() | number() | boolean() | nil,
          search: map() | list() | String.t() | number() | boolean() | nil,
          skills: map() | list() | String.t() | number() | boolean() | nil,
          tools: map() | list() | String.t() | number() | boolean() | nil,
          auto_sync: boolean() | nil,
          operation_timeout: integer() | nil,
          changed_fields: map() | list() | String.t() | number() | boolean() | nil,
          change_message: String.t() | nil,
          created_at: DateTime.t(),
          created_at_z: DateTime.t() | nil
        }

  schema "mastra_workspace_versions" do
    field(:id, :string, primary_key: true)
    field(:workspace_id, :string, source: :workspaceId)
    field(:version_number, :integer, source: :versionNumber)
    field(:name, :string)
    field(:description, :string)
    field(:filesystem, Postiz.PrismaJson)
    field(:sandbox, Postiz.PrismaJson)
    field(:mounts, Postiz.PrismaJson)
    field(:search, Postiz.PrismaJson)
    field(:skills, Postiz.PrismaJson)
    field(:tools, Postiz.PrismaJson)
    field(:auto_sync, :boolean, source: :autoSync)
    field(:operation_timeout, :integer, source: :operationTimeout)
    field(:changed_fields, Postiz.PrismaJson, source: :changedFields)
    field(:change_message, :string, source: :changeMessage)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
  end
end