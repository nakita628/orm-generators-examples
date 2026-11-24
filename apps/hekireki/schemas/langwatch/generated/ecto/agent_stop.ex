defmodule LangWatch.AgentStop do
  @moduledoc """
  The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
  every target's outcome, folded from the stop's events.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          agent_id: String.t(),
          agent_name: String.t(),
          status: String.t(),
          severity: String.t(),
          approval: String.t(),
          trigger: String.t(),
          reason: String.t(),
          requested_by_user_id: String.t() | nil,
          requested_by_api_key_id: String.t() | nil,
          decided_by_user_id: String.t() | nil,
          decided_by_api_key_id: String.t() | nil,
          restored_by_user_id: String.t() | nil,
          restored_by_api_key_id: String.t() | nil,
          targets: map() | list() | String.t() | number() | boolean(),
          requested_at: DateTime.t(),
          decided_at: DateTime.t() | nil,
          quarantined_at: DateTime.t() | nil,
          restore_requested_at: DateTime.t() | nil,
          restored_at: DateTime.t() | nil,
          last_applied_event_id: String.t(),
          last_accepted_at: integer(),
          last_occurred_at: integer(),
          projection_version: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "AgentStop" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:agent_id, :string, source: :agentId)
    field(:agent_name, :string, source: :agentName)
    field(:status, :string)
    field(:severity, :string)
    field(:approval, :string)
    field(:trigger, :string)
    field(:reason, :string)
    field(:requested_by_user_id, :string, source: :requestedByUserId)
    field(:requested_by_api_key_id, :string, source: :requestedByApiKeyId)
    field(:decided_by_user_id, :string, source: :decidedByUserId)
    field(:decided_by_api_key_id, :string, source: :decidedByApiKeyId)
    field(:restored_by_user_id, :string, source: :restoredByUserId)
    field(:restored_by_api_key_id, :string, source: :restoredByApiKeyId)
    field(:targets, LangWatch.PrismaJson)
    field(:requested_at, LangWatch.PrismaDateTime, source: :requestedAt)
    field(:decided_at, LangWatch.PrismaDateTime, source: :decidedAt)
    field(:quarantined_at, LangWatch.PrismaDateTime, source: :quarantinedAt)
    field(:restore_requested_at, LangWatch.PrismaDateTime, source: :restoreRequestedAt)
    field(:restored_at, LangWatch.PrismaDateTime, source: :restoredAt)
    field(:last_applied_event_id, :string, source: :lastAppliedEventId)
    field(:last_accepted_at, :integer, source: :lastAcceptedAt)
    field(:last_occurred_at, :integer, source: :lastOccurredAt)
    field(:projection_version, :string, source: :projectionVersion)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end