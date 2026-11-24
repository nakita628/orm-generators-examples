defmodule Postiz.MastraDatasets do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          input_schema: map() | list() | String.t() | number() | boolean() | nil,
          ground_truth_schema: map() | list() | String.t() | number() | boolean() | nil,
          request_context_schema: map() | list() | String.t() | number() | boolean() | nil,
          tags: map() | list() | String.t() | number() | boolean() | nil,
          target_type: String.t() | nil,
          target_ids: map() | list() | String.t() | number() | boolean() | nil,
          scorer_ids: map() | list() | String.t() | number() | boolean() | nil,
          version: integer(),
          created_at: DateTime.t(),
          updated_at: DateTime.t(),
          created_at_z: DateTime.t() | nil,
          updated_at_z: DateTime.t() | nil,
          organization_id: String.t() | nil,
          project_id: String.t() | nil,
          candidate_key: String.t() | nil,
          candidate_id: String.t() | nil
        }

  schema "mastra_datasets" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:metadata, Postiz.PrismaJson)
    field(:input_schema, Postiz.PrismaJson, source: :inputSchema)
    field(:ground_truth_schema, Postiz.PrismaJson, source: :groundTruthSchema)
    field(:request_context_schema, Postiz.PrismaJson, source: :requestContextSchema)
    field(:tags, Postiz.PrismaJson)
    field(:target_type, :string, source: :targetType)
    field(:target_ids, Postiz.PrismaJson, source: :targetIds)
    field(:scorer_ids, Postiz.PrismaJson, source: :scorerIds)
    field(:version, :integer)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:updated_at, Postiz.PrismaDateTime, source: :updatedAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
    field(:updated_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :updatedAtZ)
    field(:organization_id, :string, source: :organizationId)
    field(:project_id, :string, source: :projectId)
    field(:candidate_key, :string, source: :candidateKey)
    field(:candidate_id, :string, source: :candidateId)
  end
end