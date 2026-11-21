# Schedules can be attached to tasks to trigger them on a schedule
#
# friendlyId:: users see this as `id`. They start with schedule_
# taskIdentifier:: a reference to a task (not a foreign key because it's across versions)
# deduplicationKey:: can be provided and we won't create another with the same key
# generatorExpression:: the CRON pattern
# timezone:: These are IANA format string, or the default "UTC". E.g. "America/New_York"
# externalId:: Can be provided by the user then accessed inside a run
# instances:: Instances of the schedule that are active
# lastRunTriggeredAt:: @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up.
class TaskSchedule < ApplicationRecord
  self.table_name = "TaskSchedule"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :deduplicationKey, default: -> { Cuid.generate }
  attribute :userProvidedDeduplicationKey, default: false
  attribute :generatorDescription, default: ""
  attribute :timezone, default: "UTC"
  attribute :active, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { declarative: "DECLARATIVE", imperative: "IMPERATIVE" }, default: :imperative, validate: true
  enum :generatorType, { cron: "CRON" }, default: :cron, validate: true

  validates :friendlyId, presence: true, uniqueness: true
  validates :taskIdentifier, presence: true
  validates :deduplicationKey, uniqueness: { scope: :projectId }
  validates :generatorExpression, presence: true

  belongs_to :project, foreign_key: "projectId", inverse_of: :task_schedules
  has_many :instances, class_name: "TaskScheduleInstance", foreign_key: "taskScheduleId", inverse_of: :task_schedule,
                       dependent: :destroy
end
