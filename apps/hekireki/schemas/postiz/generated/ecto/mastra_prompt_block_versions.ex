defmodule Postiz.MastraPromptBlockVersions do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          block_id: String.t(),
          version_number: integer(),
          name: String.t(),
          description: String.t() | nil,
          content: String.t(),
          rules: map() | list() | String.t() | number() | boolean() | nil,
          request_context_schema: map() | list() | String.t() | number() | boolean() | nil,
          changed_fields: map() | list() | String.t() | number() | boolean() | nil,
          change_message: String.t() | nil,
          created_at: DateTime.t(),
          created_at_z: DateTime.t() | nil
        }

  schema "mastra_prompt_block_versions" do
    field(:id, :string, primary_key: true)
    field(:block_id, :string, source: :blockId)
    field(:version_number, :integer, source: :versionNumber)
    field(:name, :string)
    field(:description, :string)
    field(:content, :string)
    field(:rules, Postiz.PrismaJson)
    field(:request_context_schema, Postiz.PrismaJson, source: :requestContextSchema)
    field(:changed_fields, Postiz.PrismaJson, source: :changedFields)
    field(:change_message, :string, source: :changeMessage)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
  end
end