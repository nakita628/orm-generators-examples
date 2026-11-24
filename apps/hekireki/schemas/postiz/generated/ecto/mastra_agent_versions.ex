defmodule Postiz.MastraAgentVersions do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          agent_id: String.t(),
          version_number: integer(),
          name: String.t(),
          description: String.t() | nil,
          instructions: String.t(),
          model: map() | list() | String.t() | number() | boolean(),
          tools: map() | list() | String.t() | number() | boolean() | nil,
          default_options: map() | list() | String.t() | number() | boolean() | nil,
          workflows: map() | list() | String.t() | number() | boolean() | nil,
          agents: map() | list() | String.t() | number() | boolean() | nil,
          integration_tools: map() | list() | String.t() | number() | boolean() | nil,
          input_processors: map() | list() | String.t() | number() | boolean() | nil,
          output_processors: map() | list() | String.t() | number() | boolean() | nil,
          memory: map() | list() | String.t() | number() | boolean() | nil,
          scorers: map() | list() | String.t() | number() | boolean() | nil,
          mcp_clients: map() | list() | String.t() | number() | boolean() | nil,
          request_context_schema: map() | list() | String.t() | number() | boolean() | nil,
          workspace: map() | list() | String.t() | number() | boolean() | nil,
          skills: map() | list() | String.t() | number() | boolean() | nil,
          skills_format: String.t() | nil,
          changed_fields: map() | list() | String.t() | number() | boolean() | nil,
          change_message: String.t() | nil,
          created_at: DateTime.t(),
          created_at_z: DateTime.t() | nil,
          durable: map() | list() | String.t() | number() | boolean() | nil,
          browser: map() | list() | String.t() | number() | boolean() | nil,
          tool_providers: map() | list() | String.t() | number() | boolean() | nil
        }

  schema "mastra_agent_versions" do
    field(:id, :string, primary_key: true)
    field(:agent_id, :string, source: :agentId)
    field(:version_number, :integer, source: :versionNumber)
    field(:name, :string)
    field(:description, :string)
    field(:instructions, :string)
    field(:model, Postiz.PrismaJson)
    field(:tools, Postiz.PrismaJson)
    field(:default_options, Postiz.PrismaJson, source: :defaultOptions)
    field(:workflows, Postiz.PrismaJson)
    field(:agents, Postiz.PrismaJson)
    field(:integration_tools, Postiz.PrismaJson, source: :integrationTools)
    field(:input_processors, Postiz.PrismaJson, source: :inputProcessors)
    field(:output_processors, Postiz.PrismaJson, source: :outputProcessors)
    field(:memory, Postiz.PrismaJson)
    field(:scorers, Postiz.PrismaJson)
    field(:mcp_clients, Postiz.PrismaJson, source: :mcpClients)
    field(:request_context_schema, Postiz.PrismaJson, source: :requestContextSchema)
    field(:workspace, Postiz.PrismaJson)
    field(:skills, Postiz.PrismaJson)
    field(:skills_format, :string, source: :skillsFormat)
    field(:changed_fields, Postiz.PrismaJson, source: :changedFields)
    field(:change_message, :string, source: :changeMessage)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
    field(:durable, Postiz.PrismaJson)
    field(:browser, Postiz.PrismaJson)
    field(:tool_providers, Postiz.PrismaJson, source: :toolProviders)
  end
end