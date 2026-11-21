defmodule LangWatch.SelfHostedInstanceReport do
  @moduledoc """
  One usage report, exactly as it arrived. Kept so a number on an install's
  row can be traced back to the report that set it, and so a question asked
  later can be answered over the history rather than over the latest snapshot.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          instance_id: String.t(),
          received_at: DateTime.t(),
          version: String.t() | nil,
          report_schema_version: integer() | nil,
          unknown_fields: integer(),
          payload: map()
        }

  schema "SelfHostedInstanceReport" do
    field(:id, :string, primary_key: true)
    # The `SelfHostedInstance.instanceId` this report came from. A plain column
    # ref for the same reason the instance row's refs are.
    field(:instance_id, :string, source: :instanceId)
    field(:received_at, LangWatch.PrismaDateTime, autogenerate: true, source: :receivedAt)
    field(:version, :string)
    field(:report_schema_version, :integer, source: :reportSchemaVersion)
    # Fields this release had no name for. Counted, never stored by name.
    field(:unknown_fields, :integer, default: 0, source: :unknownFields)
    # The validated report, with every unknown field already dropped.
    field(:payload, :map)
  end
end