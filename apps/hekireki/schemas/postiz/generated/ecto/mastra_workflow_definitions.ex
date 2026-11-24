defmodule Postiz.MastraWorkflowDefinitions do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          description: String.t() | nil,
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          input_schema: map() | list() | String.t() | number() | boolean(),
          output_schema: map() | list() | String.t() | number() | boolean(),
          state_schema: map() | list() | String.t() | number() | boolean() | nil,
          request_context_schema: map() | list() | String.t() | number() | boolean() | nil,
          graph: map() | list() | String.t() | number() | boolean(),
          schedule: map() | list() | String.t() | number() | boolean() | nil,
          status: String.t(),
          source: String.t(),
          author_id: String.t() | nil,
          created_at: DateTime.t(),
          updated_at: DateTime.t(),
          created_at_z: DateTime.t() | nil,
          updated_at_z: DateTime.t() | nil
        }

  schema "mastra_workflow_definitions" do
    field(:id, :string, primary_key: true)
    field(:description, :string)
    field(:metadata, Postiz.PrismaJson)
    field(:input_schema, Postiz.PrismaJson, source: :inputSchema)
    field(:output_schema, Postiz.PrismaJson, source: :outputSchema)
    field(:state_schema, Postiz.PrismaJson, source: :stateSchema)
    field(:request_context_schema, Postiz.PrismaJson, source: :requestContextSchema)
    field(:graph, Postiz.PrismaJson)
    field(:schedule, Postiz.PrismaJson)
    field(:status, :string)
    field(:source, :string)
    field(:author_id, :string, source: :authorId)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:updated_at, Postiz.PrismaDateTime, source: :updatedAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
    field(:updated_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :updatedAtZ)
  end
end