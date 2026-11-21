defmodule LangWatch.DiscoveredAgent do
  @moduledoc """
  The same shape as DiscoveredPerson for non-human actors - Genie spaces,
  Copilot bots, OpenAI projects - plus whatever provider-native metadata the
  source published. Discovering an agent grants it nothing either.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          provider: String.t(),
          raw_agent_id: String.t(),
          display_text: String.t(),
          metadata: map(),
          first_seen_at: DateTime.t(),
          last_seen_at: DateTime.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "DiscoveredAgent" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:provider, :string)
    # Genie space / Copilot bot / OpenAI project id, verbatim.
    field(:raw_agent_id, :string, source: :rawAgentId)
    field(:display_text, :string, source: :displayText)
    # Workspace, plan, and kin - provider-native descriptive fields only. Never
    # credentials: this row is read by admin screens, not by any pull path.
    field(:metadata, :map)
    field(:first_seen_at, LangWatch.PrismaDateTime, source: :firstSeenAt)
    field(:last_seen_at, LangWatch.PrismaDateTime, source: :lastSeenAt)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end