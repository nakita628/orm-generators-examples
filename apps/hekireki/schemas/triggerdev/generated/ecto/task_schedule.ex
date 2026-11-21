defmodule Trigger.TaskSchedule do
  @moduledoc """
  Schedules can be attached to tasks to trigger them on a schedule
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          type: atom(),
          friendly_id: String.t(),
          task_identifier: String.t(),
          deduplication_key: String.t(),
          user_provided_deduplication_key: boolean(),
          generator_expression: String.t(),
          generator_description: String.t(),
          generator_type: atom(),
          timezone: String.t(),
          window_duration_seconds: integer() | nil,
          window_percentage: integer() | nil,
          default_window_duration_seconds: integer() | nil,
          minimum_window_duration_seconds: integer() | nil,
          external_id: String.t() | nil,
          last_run_triggered_at: DateTime.t() | nil,
          active: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: Trigger.Project.t() | nil,
          instances: [Trigger.TaskScheduleInstance.t()]
        }

  schema "TaskSchedule" do
    field(:id, :string, primary_key: true)
    field(:type, Ecto.Enum, values: [:DECLARATIVE, :IMPERATIVE], default: :IMPERATIVE)
    # users see this as `id`. They start with schedule_
    field(:friendly_id, :string, source: :friendlyId)
    # a reference to a task (not a foreign key because it's across versions)
    field(:task_identifier, :string, source: :taskIdentifier)
    # can be provided and we won't create another with the same key
    field(:deduplication_key, :string, source: :deduplicationKey)
    field(:user_provided_deduplication_key, :boolean, default: false, source: :userProvidedDeduplicationKey)
    # the CRON pattern
    field(:generator_expression, :string, source: :generatorExpression)
    field(:generator_description, :string, default: "", source: :generatorDescription)
    field(:generator_type, Ecto.Enum, values: [:CRON], default: :CRON, source: :generatorType)
    # These are IANA format string, or the default "UTC". E.g. "America/New_York"
    field(:timezone, :string, default: "UTC")
    field(:window_duration_seconds, :integer, source: :windowDurationSeconds)
    field(:window_percentage, :integer, source: :windowPercentage)
    field(:default_window_duration_seconds, :integer, source: :defaultWindowDurationSeconds)
    field(:minimum_window_duration_seconds, :integer, source: :minimumWindowDurationSeconds)
    # Can be provided by the user then accessed inside a run
    field(:external_id, :string, source: :externalId)
    # @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up.
    field(:last_run_triggered_at, Trigger.PrismaDateTime, source: :lastRunTriggeredAt)
    field(:active, :boolean, default: true)
    field(:project_id, :string, source: :projectId)
    belongs_to(:project, Trigger.Project, foreign_key: :project_id, define_field: false, type: :string)
    # Instances of the schedule that are active
    has_many(:instances, Trigger.TaskScheduleInstance, foreign_key: :task_schedule_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end