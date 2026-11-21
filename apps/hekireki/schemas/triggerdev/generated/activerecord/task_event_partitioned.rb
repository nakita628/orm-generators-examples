# This is the unified otel span/log model that will eventually be replaced by clickhouse
#
# message:: This matches the span name for a trace event, or the log body for a log event
# startTime:: This is the time the event started in nanoseconds since the epoch
# duration:: This is the duration of the event in nanoseconds
# properties:: This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering
# metadata:: This represents all span attributes in the $metadata namespace, like $metadata.payload
# style:: This represents all span attributes in the $style namespace, like $style
# output:: This represents all span attributes in the $output namespace, like $output
# outputType:: This represents the mimetype of the output, such as application/json or application/super+json
class TaskEventPartitioned < ApplicationRecord
  self.table_name = "TaskEventPartitioned"
  self.primary_key = %w[id createdAt]

  attribute :id, default: -> { Cuid.generate }
  attribute :isError, default: false
  attribute :isPartial, default: false
  attribute :isCancelled, default: false
  attribute :duration, default: 0
  attribute :runIsTest, default: false
  attribute :usageDurationMs, default: 0
  attribute :usageCostInCents, default: 0
  alias_attribute :created_at, :createdAt

  enum :level, { trace: "TRACE", debug: "DEBUG", log: "LOG", info: "INFO", warn: "WARN", error: "ERROR" },
       default: :trace, prefix: true, validate: true
  enum :kind, { unspecified: "UNSPECIFIED", internal: "INTERNAL", server: "SERVER", client: "CLIENT", producer: "PRODUCER", consumer: "CONSUMER", unrecognized: "UNRECOGNIZED", log: "LOG" },
       default: :internal, validate: true
  enum :status, { unset: "UNSET", ok: "OK", error: "ERROR", unrecognized: "UNRECOGNIZED" }, default: :unset,
                                                                                            prefix: true, validate: true
  enum :environmentType, { production: "PRODUCTION", staging: "STAGING", development: "DEVELOPMENT", preview: "PREVIEW" },
       validate: true

  validates :message, presence: true
  validates :traceId, presence: true
  validates :spanId, presence: true
  validates :serviceName, presence: true
  validates :serviceNamespace, presence: true
  validates :startTime, presence: true
  validates :environmentId, presence: true
  validates :organizationId, presence: true
  validates :projectId, presence: true
  validates :projectRef, presence: true
  validates :runId, presence: true
  validates :taskSlug, presence: true
  validates :properties, exclusion: { in: [ nil ] }
end
