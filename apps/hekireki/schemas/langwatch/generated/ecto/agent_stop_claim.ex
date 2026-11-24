defmodule LangWatch.AgentStopClaim do
  @moduledoc """
  The one open stop an agent may have: the api takes this row before it sends a stop's request,
  so two requests at once cannot both open one. A later request takes it over once the stop it
  names has closed.
  `sentAt` is set once the request command was sent: from then on the claim passes on only after
  its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
  `decisionTakenAt` is set while that decision is in flight and cleared once its command was
  sent; a decision left in flight is taken over once abandoned.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          agent_id: String.t(),
          stop_id: String.t(),
          claimed_at: DateTime.t(),
          sent_at: DateTime.t() | nil,
          decision: String.t() | nil,
          decision_taken_at: DateTime.t() | nil
        }

  schema "AgentStopClaim" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:agent_id, :string, primary_key: true, source: :agentId)
    field(:stop_id, :string, source: :stopId)
    field(:claimed_at, LangWatch.PrismaDateTime, source: :claimedAt)
    field(:sent_at, LangWatch.PrismaDateTime, source: :sentAt)
    field(:decision, :string)
    field(:decision_taken_at, LangWatch.PrismaDateTime, source: :decisionTakenAt)
  end
end