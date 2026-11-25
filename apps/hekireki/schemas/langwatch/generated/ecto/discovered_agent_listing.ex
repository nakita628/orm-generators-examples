defmodule LangWatch.DiscoveredAgentListing do
  @moduledoc """
  One source's presence fact for one discovered agent. An agent is gone only when every
  source that listed it marked it gone; an agent no source listed (older sightings) never is.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          ingestion_source_id: String.t(),
          discovered_agent_id: String.t(),
          last_listed_at: DateTime.t(),
          gone_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "DiscoveredAgentListing" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:ingestion_source_id, :string, source: :ingestionSourceId)
    field(:discovered_agent_id, :string, source: :discoveredAgentId)
    field(:last_listed_at, LangWatch.PrismaDateTime, source: :lastListedAt)
    # Set when this source's listing stopped naming the agent; cleared when it names it again.
    field(:gone_at, LangWatch.PrismaDateTime, source: :goneAt)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end