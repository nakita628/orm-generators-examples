# Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
# The row is essentially static — no status, no counters, no pointers. No
# foreign keys: project/runtimeEnvironment/organization ids are plain
# scalar columns (matches TaskRun pattern). List-style queries are served
# from ClickHouse, not Postgres, so only point-lookup indexes live here.
#
# externalId:: User-supplied identifier scoped to the environment. Used for
#   idempotent upsert and for resolving sessions via the public API.
# type:: Plain string — intentionally not an enum.
# projectId:: Denormalized scoping columns — no FK relations.
# taskIdentifier:: Task this session triggers runs against. Required — Sessions are
#   task-bound: creating a session also triggers its first run, and
#   every subsequent re-trigger uses this same identifier.
# triggerConfig:: Trigger config used for every run this session schedules. Shape
#   (validated at the route layer, opaque to the DB):
#   { basePayload: object, machine?: string, queue?: string,
#   tags?: string[], maxAttempts?: number,
#   idleTimeoutInSeconds?: number }
#   `basePayload` carries the customer's client-data; runtime fields
#   (chatId, messages, trigger) are merged at trigger time.
# isTest:: Whether this session was created from the Test/playground UI rather
#   than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
# currentRunId:: Live run pointer — non-FK so run deletion never cascades. Can lag
#   reality; the `.in/append` handler re-checks the snapshot status
#   before reusing it.
# currentRunVersion:: Monotonic counter used for optimistic locking on `currentRunId`
#   swaps. Bumped atomically alongside any update that changes
#   `currentRunId`.
# closedAt:: Terminal markers — written once, never flipped back.
# streamBasinName:: S2 basin where this session's stream pair lives. Stamped at create
#   time from `Organization.streamBasinName` so reads can resolve the
#   basin without joining org. Null when the org has no per-org basin
#   (OSS, or pre-backfill); reads fall back to the global basin.
# chatSnapshotStoragePath:: Storage URI (with protocol prefix) for this session's chat.agent
#   snapshot blob. Set on first snapshot write. Null = pre-column session,
#   fall back to computed default path.
class Session < ApplicationRecord
  self.table_name = "Session"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :isTest, default: false
  attribute :tags, default: -> { [] }
  attribute :currentRunVersion, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :environmentType, { production: "PRODUCTION", staging: "STAGING", development: "DEVELOPMENT", preview: "PREVIEW" },
       validate: true

  validates :friendlyId, presence: true, uniqueness: true
  validates :type, presence: true
  validates :projectId, presence: true
  validates :runtimeEnvironmentId, presence: true,
                                   uniqueness: { scope: :externalId, conditions: -> { where.not(externalId: nil) } }
  validates :organizationId, presence: true
  validates :taskIdentifier, presence: true
  validates :triggerConfig, exclusion: { in: [ nil ] }

  has_many :runs, class_name: "SessionRun", foreign_key: "sessionId", inverse_of: :session, dependent: :destroy
end
