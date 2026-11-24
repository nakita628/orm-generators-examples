defmodule LangWatch.NurturingProjectUsageView do
  @moduledoc """
  Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
  and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
  organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          project_id: String.t(),
          owner: String.t(),
          folded: map() | list() | String.t() | number() | boolean(),
          applied_event_ids: [String.t()],
          projection_version: String.t(),
          updated_at: DateTime.t()
        }

  schema "NurturingProjectUsageView" do
    field(:project_id, :string, primary_key: true, source: :projectId)
    # The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
    # or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
    field(:owner, :string, primary_key: true)
    # The project's usage as the fold folded it (NurturingProjectUsageState).
    field(:folded, LangWatch.PrismaJson)
    # The event ids of the last delivery folded in, so a redelivery folds once.
    field(:applied_event_ids, {:array, :string}, source: :appliedEventIds)
    # A row of another version reads as undecodable and is re-folded from the log.
    field(:projection_version, :string, source: :projectionVersion)
    # Event-derived, so a replayed row equals the row it rebuilds.
    field(:updated_at, LangWatch.PrismaDateTime, source: :updatedAt)
  end
end