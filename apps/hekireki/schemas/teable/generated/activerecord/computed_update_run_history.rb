# Completion ledger for computed-update outbox tasks (lineage / latency
# tracing). Successful tasks are hard-deleted from the outbox on markDone,
# so this table keeps a retention-bounded trail per completed task: the
# run/stage chain (runId, stageDepth, predecessorTaskId) and the timeline
# from source mutation to converged write (sourceChangedAt → completedAt).
# Observability only — the worker never reads it; rows are pruned after a
# retention window.
#
# source_field_ids:: Source (changed) field ids; populated for seed tasks, where
#   affected_field_ids carries the trigger fields rather than targets.
class ComputedUpdateRunHistory < ApplicationRecord
  self.table_name = "computed_update_run_history"
  self.primary_key = "task_id"

  attribute :origin_run_ids, default: -> { [] }
  attribute :affected_table_ids, default: -> { [] }
  attribute :affected_field_ids, default: -> { [] }
  attribute :source_field_ids, default: -> { [] }
  attribute :seed_record_count, default: 0
  attribute :stage_depth, default: 0
  attribute :run_total_steps, default: 0
  attribute :run_completed_steps_before, default: 0
  attribute :estimated_complexity, default: 0
  attribute :attempts, default: 0
  attribute :duration_ms, default: 0

  validates :base_id, presence: true
  validates :seed_table_id, presence: true
  validates :change_type, presence: true
  validates :run_id, presence: true
  validates :outcome, presence: true
  validates :enqueued_at, presence: true
  validates :completed_at, presence: true
end
