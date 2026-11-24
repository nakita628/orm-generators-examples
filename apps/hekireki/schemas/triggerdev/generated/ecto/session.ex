defmodule Trigger.Session do
  @moduledoc """
  Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
  The row is essentially static — no status, no counters, no pointers. No
  foreign keys: project/runtimeEnvironment/organization ids are plain
  scalar columns (matches TaskRun pattern). List-style queries are served
  from ClickHouse, not Postgres, so only point-lookup indexes live here.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          friendly_id: String.t(),
          external_id: String.t() | nil,
          type: String.t(),
          project_id: String.t(),
          runtime_environment_id: String.t(),
          environment_type: atom(),
          organization_id: String.t(),
          task_identifier: String.t(),
          trigger_config: map() | list() | String.t() | number() | boolean(),
          is_test: boolean(),
          tags: [String.t()],
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          current_run_id: String.t() | nil,
          current_run_version: integer(),
          closed_at: DateTime.t() | nil,
          closed_reason: String.t() | nil,
          expires_at: DateTime.t() | nil,
          stream_basin_name: String.t() | nil,
          chat_snapshot_storage_path: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          runs: [Trigger.SessionRun.t()]
        }

  schema "Session" do
    field(:id, :string, primary_key: true)
    field(:friendly_id, :string, source: :friendlyId)
    # User-supplied identifier scoped to the environment. Used for
    # idempotent upsert and for resolving sessions via the public API.
    field(:external_id, :string, source: :externalId)
    # Plain string — intentionally not an enum.
    field(:type, :string)
    # Denormalized scoping columns — no FK relations.
    field(:project_id, :string, source: :projectId)
    field(:runtime_environment_id, :string, source: :runtimeEnvironmentId)
    field(:environment_type, Ecto.Enum, values: [:PRODUCTION, :STAGING, :DEVELOPMENT, :PREVIEW], source: :environmentType)
    field(:organization_id, :string, source: :organizationId)
    # Task this session triggers runs against. Required — Sessions are
    # task-bound: creating a session also triggers its first run, and
    # every subsequent re-trigger uses this same identifier.
    field(:task_identifier, :string, source: :taskIdentifier)
    # Trigger config used for every run this session schedules. Shape
    # (validated at the route layer, opaque to the DB):
    # { basePayload: object, machine?: string, queue?: string,
    # tags?: string[], maxAttempts?: number,
    # idleTimeoutInSeconds?: number }
    # `basePayload` carries the customer's client-data; runtime fields
    # (chatId, messages, trigger) are merged at trigger time.
    field(:trigger_config, Trigger.PrismaJson, source: :triggerConfig)
    # Whether this session was created from the Test/playground UI rather
    # than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
    field(:is_test, :boolean, default: false, source: :isTest)
    field(:tags, {:array, :string})
    field(:metadata, Trigger.PrismaJson)
    # Live run pointer — non-FK so run deletion never cascades. Can lag
    # reality; the `.in/append` handler re-checks the snapshot status
    # before reusing it.
    field(:current_run_id, :string, source: :currentRunId)
    # Monotonic counter used for optimistic locking on `currentRunId`
    # swaps. Bumped atomically alongside any update that changes
    # `currentRunId`.
    field(:current_run_version, :integer, default: 0, source: :currentRunVersion)
    # Terminal markers — written once, never flipped back.
    field(:closed_at, Trigger.PrismaDateTime, source: :closedAt)
    field(:closed_reason, :string, source: :closedReason)
    field(:expires_at, Trigger.PrismaDateTime, source: :expiresAt)
    # S2 basin where this session's stream pair lives. Stamped at create
    # time from `Organization.streamBasinName` so reads can resolve the
    # basin without joining org. Null when the org has no per-org basin
    # (OSS, or pre-backfill); reads fall back to the global basin.
    field(:stream_basin_name, :string, source: :streamBasinName)
    # Storage URI (with protocol prefix) for this session's chat.agent
    # snapshot blob. Set on first snapshot write. Null = pre-column session,
    # fall back to computed default path.
    field(:chat_snapshot_storage_path, :string, source: :chatSnapshotStoragePath)
    has_many(:runs, Trigger.SessionRun, foreign_key: :session_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end