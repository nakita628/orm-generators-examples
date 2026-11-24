defmodule Postiz.MastraSkillVersions do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          skill_id: String.t(),
          version_number: integer(),
          name: String.t(),
          description: String.t(),
          instructions: String.t(),
          license: String.t() | nil,
          compatibility: map() | list() | String.t() | number() | boolean() | nil,
          source: map() | list() | String.t() | number() | boolean() | nil,
          references: map() | list() | String.t() | number() | boolean() | nil,
          scripts: map() | list() | String.t() | number() | boolean() | nil,
          assets: map() | list() | String.t() | number() | boolean() | nil,
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          tree: map() | list() | String.t() | number() | boolean() | nil,
          changed_fields: map() | list() | String.t() | number() | boolean() | nil,
          change_message: String.t() | nil,
          created_at: DateTime.t(),
          created_at_z: DateTime.t() | nil,
          files: map() | list() | String.t() | number() | boolean() | nil
        }

  schema "mastra_skill_versions" do
    field(:id, :string, primary_key: true)
    field(:skill_id, :string, source: :skillId)
    field(:version_number, :integer, source: :versionNumber)
    field(:name, :string)
    field(:description, :string)
    field(:instructions, :string)
    field(:license, :string)
    field(:compatibility, Postiz.PrismaJson)
    field(:source, Postiz.PrismaJson)
    field(:references, Postiz.PrismaJson)
    field(:scripts, Postiz.PrismaJson)
    field(:assets, Postiz.PrismaJson)
    field(:metadata, Postiz.PrismaJson)
    field(:tree, Postiz.PrismaJson)
    field(:changed_fields, Postiz.PrismaJson, source: :changedFields)
    field(:change_message, :string, source: :changeMessage)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
    field(:files, Postiz.PrismaJson)
  end
end