defmodule Formbricks.DeletionCleanup do
  @moduledoc """
  Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
  storage files that can't be removed inside the database transaction. A row is written in the same
  transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
  queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          kind: atom(),
          organization_id: String.t(),
          workspace_id: String.t(),
          survey_id: String.t(),
          tenant_ids: [String.t()],
          response_ids: [String.t()],
          file_keys: [String.t()],
          attempts: integer(),
          next_attempt_at: DateTime.t(),
          last_error: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "DeletionCleanup" do
    field(:id, :string, primary_key: true)
    field(:kind, Ecto.Enum, values: [:hubSurvey, :hubResponses, :storageFiles, :storageSurveyFolder, :hubResponseOutcome])
    # Scoping and logs. Not foreign keys: the rows they name are usually gone already.
    field(:organization_id, :string, source: :organizationId)
    field(:workspace_id, :string, source: :workspaceId)
    field(:survey_id, :string, source: :surveyId)
    # Hub kinds: the organisation's feedback directories, read inside the delete's transaction.
    field(:tenant_ids, {:array, :string}, source: :tenantIds)
    # `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response.
    field(:response_ids, {:array, :string}, source: :responseIds)
    # `storageFiles`: the object keys or URLs to delete.
    field(:file_keys, {:array, :string}, source: :fileKeys)
    field(:attempts, :integer, default: 0)
    # Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup.
    field(:next_attempt_at, Formbricks.PrismaDateTime, autogenerate: true, source: :nextAttemptAt)
    # The last failure: a short code and message, never record or file content.
    field(:last_error, :string, source: :lastError)
    timestamps(inserted_at_source: :created_at)
  end
end