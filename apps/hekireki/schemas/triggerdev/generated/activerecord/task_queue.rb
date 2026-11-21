# concurrencyVersion:: How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
#   (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
# role:: QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
#   excluded from every queue-facing surface.
# concurrencyLimit:: Represents the current concurrency limit for the queue
# concurrencyLimitOverriddenAt:: When the concurrency limit was overridden
# concurrencyLimitOverriddenBy:: Who overrode the concurrency limit (will be null if overridden via the API)
# concurrencyLimitBase:: If concurrencyLimit is overridden, this is the overridden value
# concurrencyLimitOverridePercent:: If the override was expressed as a percentage of the environment limit, this stores that
#   percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
#   Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
# totalConcurrencyLimit:: Caps total concurrent runs across ALL concurrencyKey values of this queue
#   (concurrencyLimit applies per key value). Null = no total cap.
# totalConcurrencyLimitOverriddenAt:: When the total concurrency limit was overridden
# totalConcurrencyLimitOverriddenBy:: Who overrode the total concurrency limit (null when overridden via the API)
# totalConcurrencyLimitBase:: If totalConcurrencyLimit is overridden, the declared value it reverts to on reset
# archivedAt:: Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths.
class TaskQueue < ApplicationRecord
  self.table_name = "TaskQueue"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :paused, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { virtual: "VIRTUAL", named: "NAMED" }, default: :virtual, validate: true
  enum :version, { v1: "V1", v2: "V2" }, default: :v1, validate: true
  enum :concurrencyVersion, { v1: "V1", v2: "V2" }, default: :v1, prefix: true, validate: true
  enum :role, { queue: "QUEUE", limit: "LIMIT" }, default: :queue, prefix: true, validate: true

  validates :friendlyId, presence: true, uniqueness: true
  validates :name, presence: true, uniqueness: { scope: :runtimeEnvironmentId }

  belongs_to :project, foreign_key: "projectId", inverse_of: :task_queues
  belongs_to :runtime_environment, foreign_key: "runtimeEnvironmentId", inverse_of: :task_queues
  has_many :attempts, class_name: "TaskRunAttempt", foreign_key: "queueId", inverse_of: :queue, dependent: :destroy
  has_many :tasks, class_name: "BackgroundWorkerTask", foreign_key: "queueId", inverse_of: :queue, dependent: :nullify
  has_and_belongs_to_many :workers, class_name: "BackgroundWorker", join_table: "_BackgroundWorkerToTaskQueue",
                                    foreign_key: "B", association_foreign_key: "A"
end
