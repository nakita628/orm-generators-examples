defmodule LangWatch.AgentStopPlan do
  @moduledoc """
  One agent's stop plan version: bumped by every save, so a save over a plan someone else saved
  since it was read is refused. No row means the plan was never saved (version 0).
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          organization_id: String.t(),
          agent_id: String.t(),
          version: integer(),
          updated_at: DateTime.t()
        }

  schema "AgentStopPlan" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    # The AgentRegistryEntry the plan belongs to.
    field(:agent_id, :string, primary_key: true, source: :agentId)
    field(:version, :integer)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end