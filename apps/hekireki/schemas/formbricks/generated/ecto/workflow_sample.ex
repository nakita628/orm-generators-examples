defmodule Formbricks.WorkflowSample do
  @moduledoc """
  Per-call data kept beside a workflow (ADR 0048): one row per source, `trigger` for the Webhook
  trigger's last payload (plus the last call's time and IP), or an HTTP request step id for its last
  Test request response. Written on every call, so it lives here and never bumps `Workflow.updatedAt`.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          workflow_id: String.t(),
          workspace_id: String.t(),
          source: String.t(),
          sample: map() | list() | String.t() | number() | boolean(),
          last_called_at: DateTime.t() | nil,
          last_caller_ip: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "WorkflowSample" do
    field(:id, :string, primary_key: true)
    field(:workflow_id, :string, source: :workflowId)
    field(:workspace_id, :string, source: :workspaceId)
    field(:source, :string)
    # [WorkflowSampleData]
    field(:sample, Formbricks.PrismaJson)
    field(:last_called_at, Formbricks.PrismaDateTime, source: :lastCalledAt)
    field(:last_caller_ip, :string, source: :lastCallerIp)
    timestamps(inserted_at_source: :created_at)
  end
end