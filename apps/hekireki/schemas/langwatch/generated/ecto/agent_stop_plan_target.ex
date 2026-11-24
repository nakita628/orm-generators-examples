defmodule LangWatch.AgentStopPlanTarget do
  @moduledoc """
  One target of a registry agent's kill switch stop plan (enterprise/modules/governance/specs/
  agent-kill-switch.feature): a Databricks serving endpoint, a Bedrock agent alias or a gateway
  virtual key. A plan write replaces every row of the agent at once.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          agent_id: String.t(),
          position: integer(),
          label: String.t(),
          target: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "AgentStopPlanTarget" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # The AgentRegistryEntry the plan belongs to.
    field(:agent_id, :string, source: :agentId)
    field(:position, :integer)
    field(:label, :string)
    field(:target, LangWatch.PrismaJson)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end