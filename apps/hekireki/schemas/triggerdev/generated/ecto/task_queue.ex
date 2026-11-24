defmodule Trigger.TaskQueue do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          friendly_id: String.t(),
          name: String.t(),
          type: atom(),
          version: atom(),
          orderable_name: String.t() | nil,
          concurrency_version: atom(),
          role: atom(),
          concurrency_limit: integer() | nil,
          concurrency_limit_overridden_at: DateTime.t() | nil,
          concurrency_limit_overridden_by: String.t() | nil,
          concurrency_limit_base: integer() | nil,
          concurrency_limit_override_percent: Decimal.t() | nil,
          total_concurrency_limit: integer() | nil,
          total_concurrency_limit_overridden_at: DateTime.t() | nil,
          total_concurrency_limit_overridden_by: String.t() | nil,
          total_concurrency_limit_base: integer() | nil,
          rate_limit: map() | list() | String.t() | number() | boolean() | nil,
          paused: boolean(),
          archived_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: Trigger.Project.t() | nil,
          runtime_environment: Trigger.RuntimeEnvironment.t() | nil,
          attempts: [Trigger.TaskRunAttempt.t()],
          tasks: [Trigger.BackgroundWorkerTask.t()],
          workers: [Trigger.BackgroundWorker.t()]
        }

  schema "TaskQueue" do
    field(:id, :string, primary_key: true)
    field(:friendly_id, :string, source: :friendlyId)
    field(:name, :string)
    field(:type, Ecto.Enum, values: [:VIRTUAL, :NAMED], default: :VIRTUAL)
    field(:version, Ecto.Enum, values: [:V1, :V2], default: :V1)
    field(:orderable_name, :string, source: :orderableName)
    # How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
    # (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
    field(:concurrency_version, Ecto.Enum, values: [:V1, :V2], default: :V1, source: :concurrencyVersion)
    # QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
    # excluded from every queue-facing surface.
    field(:role, Ecto.Enum, values: [:QUEUE, :LIMIT], default: :QUEUE)
    # Represents the current concurrency limit for the queue
    field(:concurrency_limit, :integer, source: :concurrencyLimit)
    # When the concurrency limit was overridden
    field(:concurrency_limit_overridden_at, Trigger.PrismaDateTime, source: :concurrencyLimitOverriddenAt)
    # Who overrode the concurrency limit (will be null if overridden via the API)
    field(:concurrency_limit_overridden_by, :string, source: :concurrencyLimitOverriddenBy)
    # If concurrencyLimit is overridden, this is the overridden value
    field(:concurrency_limit_base, :integer, source: :concurrencyLimitBase)
    # If the override was expressed as a percentage of the environment limit, this stores that
    # percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
    # Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
    field(:concurrency_limit_override_percent, :decimal, source: :concurrencyLimitOverridePercent)
    # Caps total concurrent runs across ALL concurrencyKey values of this queue
    # (concurrencyLimit applies per key value). Null = no total cap.
    field(:total_concurrency_limit, :integer, source: :totalConcurrencyLimit)
    # When the total concurrency limit was overridden
    field(:total_concurrency_limit_overridden_at, Trigger.PrismaDateTime, source: :totalConcurrencyLimitOverriddenAt)
    # Who overrode the total concurrency limit (null when overridden via the API)
    field(:total_concurrency_limit_overridden_by, :string, source: :totalConcurrencyLimitOverriddenBy)
    # If totalConcurrencyLimit is overridden, the declared value it reverts to on reset
    field(:total_concurrency_limit_base, :integer, source: :totalConcurrencyLimitBase)
    field(:rate_limit, Trigger.PrismaJson, source: :rateLimit)
    field(:paused, :boolean, default: false)
    # Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths.
    field(:archived_at, Trigger.PrismaDateTime, source: :archivedAt)
    field(:project_id, :string, source: :projectId)
    field(:runtime_environment_id, :string, source: :runtimeEnvironmentId)
    belongs_to(:project, Trigger.Project, foreign_key: :project_id, define_field: false, type: :string)
    belongs_to(:runtime_environment, Trigger.RuntimeEnvironment, foreign_key: :runtime_environment_id, define_field: false, type: :string)
    has_many(:attempts, Trigger.TaskRunAttempt, foreign_key: :queue_id, references: :id)
    has_many(:tasks, Trigger.BackgroundWorkerTask, foreign_key: :queue_id, references: :id)
    many_to_many(:workers, Trigger.BackgroundWorker, join_through: "_BackgroundWorkerToTaskQueue", join_keys: [B: :id, A: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end