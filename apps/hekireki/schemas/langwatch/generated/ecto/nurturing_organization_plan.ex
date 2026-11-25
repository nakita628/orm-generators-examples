defmodule LangWatch.NurturingOrganizationPlan do
  @moduledoc """
  Each organization's plan, folded from billing's subscription_changed facts: nurturing's own
  copy. The newest fact wins, and a real fact wins a tie over a catch-up; Customer.io hears a
  fact only when it changes this row.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          plan: String.t() | nil,
          occurred_at: DateTime.t(),
          from_catch_up: boolean(),
          applied_event_ids: [String.t()],
          winner_event_id: String.t() | nil
        }

  schema "NurturingOrganizationPlan" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:plan, :string)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:from_catch_up, :boolean, source: :fromCatchUp)
    # The event ids folded in, so a redelivery folds once.
    field(:applied_event_ids, {:array, :string}, source: :appliedEventIds)
    # The fact that last changed the row; only it may tell Customer.io.
    field(:winner_event_id, :string, source: :winnerEventId)
  end
end