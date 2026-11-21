defmodule LangWatch.AgentRegistryEvent do
  @moduledoc """
  The registry's append-only change feed: one row per registered, versioned, removed or
  tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
  fan-out as recorded in governance's outbox.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          sequence: integer(),
          type: String.t(),
          entry_id: String.t(),
          from_version: integer() | nil,
          to_version: integer() | nil,
          payload: map(),
          occurred_at: DateTime.t(),
          delivery_recorded_at: DateTime.t() | nil
        }

  schema "AgentRegistryEvent" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:sequence, :integer)
    field(:type, :string)
    field(:entry_id, :string, source: :entryId)
    field(:from_version, :integer, source: :fromVersion)
    field(:to_version, :integer, source: :toVersion)
    field(:payload, :map)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:delivery_recorded_at, LangWatch.PrismaDateTime, source: :deliveryRecordedAt)
  end
end