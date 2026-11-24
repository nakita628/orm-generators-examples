defmodule Trigger.AlertEvent do
  @moduledoc """
  RANGE-partitioned by month on `at` (see migration).
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          alert_id: String.t(),
          type: String.t(),
          at: DateTime.t(),
          facts: map() | list() | String.t() | number() | boolean(),
          actor_id: String.t() | nil,
          source_event_id: String.t() | nil
        }

  schema "AlertEvent" do
    field(:id, :string, primary_key: true)
    field(:alert_id, :string, source: :alertId)
    field(:type, :string)
    field(:at, Trigger.PrismaDateTime, primary_key: true, autogenerate: true)
    field(:facts, Trigger.PrismaJson, default: %{})
    field(:actor_id, :string, source: :actorId)
    field(:source_event_id, :string, source: :sourceEventId)
  end
end