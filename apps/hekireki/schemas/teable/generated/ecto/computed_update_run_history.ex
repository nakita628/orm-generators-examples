defmodule Teable.ComputedUpdateRunHistory do
  @moduledoc """
  Completion ledger for computed-update outbox tasks (lineage / latency
  tracing). Successful tasks are hard-deleted from the outbox on markDone,
  so this table keeps a retention-bounded trail per completed task: the
  run/stage chain (runId, stageDepth, predecessorTaskId) and the timeline
  from source mutation to converged write (sourceChangedAt → completedAt).
  Observability only — the worker never reads it; rows are pruned after a
  retention window.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          task_id: String.t(),
          base_id: String.t(),
          seed_table_id: String.t(),
          change_type: String.t(),
          run_id: String.t(),
          origin_run_ids: [String.t()],
          steps: map() | nil,
          edges: map() | nil,
          affected_table_ids: [String.t()],
          affected_field_ids: [String.t()],
          source_field_ids: [String.t()],
          seed_record_count: integer(),
          stage_depth: integer(),
          predecessor_task_id: String.t() | nil,
          run_total_steps: integer(),
          run_completed_steps_before: integer(),
          sync_max_level: integer() | nil,
          estimated_complexity: integer(),
          attempts: integer(),
          outcome: String.t(),
          source_changed_at: DateTime.t() | nil,
          enqueued_at: DateTime.t(),
          started_at: DateTime.t() | nil,
          completed_at: DateTime.t(),
          duration_ms: integer()
        }

  schema "computed_update_run_history" do
    field(:task_id, :string, primary_key: true)
    field(:base_id, :string)
    field(:seed_table_id, :string)
    field(:change_type, :string)
    field(:run_id, :string)
    field(:origin_run_ids, {:array, :string})
    field(:steps, :map)
    field(:edges, :map)
    field(:affected_table_ids, {:array, :string})
    field(:affected_field_ids, {:array, :string})
    # Source (changed) field ids; populated for seed tasks, where
    # affected_field_ids carries the trigger fields rather than targets.
    field(:source_field_ids, {:array, :string})
    field(:seed_record_count, :integer, default: 0)
    field(:stage_depth, :integer, default: 0)
    field(:predecessor_task_id, :string)
    field(:run_total_steps, :integer, default: 0)
    field(:run_completed_steps_before, :integer, default: 0)
    field(:sync_max_level, :integer)
    field(:estimated_complexity, :integer, default: 0)
    field(:attempts, :integer, default: 0)
    field(:outcome, :string)
    field(:source_changed_at, Teable.PrismaDateTime)
    field(:enqueued_at, Teable.PrismaDateTime)
    field(:started_at, Teable.PrismaDateTime)
    field(:completed_at, Teable.PrismaDateTime)
    field(:duration_ms, :integer, default: 0)
  end
end