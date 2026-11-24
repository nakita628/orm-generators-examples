defmodule LangWatch.AgentStopTargetPriorState do
  @moduledoc """
  Whether a target was already stopped before a stop first touched it, recorded once before the
  stop's platform call, so a retried call after a success still knows the change is the stop's own.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          stop_id: String.t(),
          target_id: String.t(),
          was_stopped: boolean(),
          recorded_at: DateTime.t()
        }

  schema "AgentStopTargetPriorState" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:stop_id, :string, primary_key: true, source: :stopId)
    field(:target_id, :string, primary_key: true, source: :targetId)
    field(:was_stopped, :boolean, source: :wasStopped)
    field(:recorded_at, LangWatch.PrismaDateTime, source: :recordedAt)
  end
end