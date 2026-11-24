defmodule Langfuse.GatewaySpendLimitAlert do
  @moduledoc """
  One notification per limit, subject, alert type, threshold, counting start and window, so the gateway may report a limit any number of times.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Langfuse.PrismaDateTime, autogenerate: {Langfuse.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          subject_id: String.t(),
          type: atom(),
          threshold_percent: integer(),
          counting_started_at: DateTime.t(),
          window_start: DateTime.t(),
          inserted_at: DateTime.t(),
          spend_limit: Langfuse.GatewaySpendLimit.t() | nil
        }

  schema "gateway_spend_limit_alerts" do
    field(:id, :string, primary_key: true)
    field(:subject_id, :string, default: "")
    field(:type, Ecto.Enum, values: [:WARNING, :EXCEEDED])
    field(:threshold_percent, :integer, default: 100)
    field(:counting_started_at, Langfuse.PrismaDateTime)
    field(:window_start, Langfuse.PrismaDateTime)
    belongs_to(:spend_limit, Langfuse.GatewaySpendLimit, foreign_key: :spend_limit_id, type: :string)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end