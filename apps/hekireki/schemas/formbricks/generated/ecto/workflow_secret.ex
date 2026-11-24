defmodule Formbricks.WorkflowSecret do
  @moduledoc """
  Encrypted value of an HTTP request secret header (ENG-7168, ADR 0043). The definition and every version
  store only `{name, secret: true, secretId}`; Replace overwrites this row in place, so the enabled
  version uses the new value at once. Rows nothing references are deleted on Enable; all are deleted
  with the workflow. A value is only sent to the origin it was set for (ADR 0051).
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          workflow_id: String.t(),
          workspace_id: String.t(),
          value: String.t(),
          bound_origin: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "WorkflowSecret" do
    field(:id, :string, primary_key: true)
    field(:workflow_id, :string, source: :workflowId)
    field(:workspace_id, :string, source: :workspaceId)
    # `symmetricEncrypt` ciphertext (AES-256-GCM). Never logged, never returned.
    field(:value, :string)
    # Origin (scheme, host, port) of the step's URL when the value was set or replaced; the value is only
    # sent there (ADR 0051). Null for a value set before host binding: held to the step's saved URL
    # until it is set again.
    field(:bound_origin, :string, source: :boundOrigin)
    timestamps(inserted_at_source: :created_at)
  end
end