defmodule Formbricks.AttachmentExport do
  @moduledoc """
  A survey's response attachments built into one ZIP in object storage by a background job and
  downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
  Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
  survey deletes it with the survey's storage prefix.

  A partial unique index in the migration allows one queued or running export per requester and
  survey, so a double click returns the export already in flight.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          survey_id: String.t(),
          status: atom(),
          filter_criteria: map() | nil,
          time_zone: String.t(),
          dispatch_count: integer(),
          file_count: integer(),
          response_count: integer(),
          included_file_count: integer() | nil,
          total_bytes: integer() | nil,
          truncated: boolean(),
          storage_key: String.t() | nil,
          error_code: String.t() | nil,
          started_at: DateTime.t() | nil,
          finished_at: DateTime.t() | nil,
          expires_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          workspace: Formbricks.Workspace.t() | nil,
          requested_by: Formbricks.User.t() | nil
        }

  schema "AttachmentExport" do
    field(:id, :string, primary_key: true)
    field(:survey_id, :string, source: :surveyId)
    field(:status, Ecto.Enum, values: [:queued, :running, :completed, :failed, :expired], default: :queued)
    # The response filter at request time, a `TResponseFilterCriteria`; null exports every response.
    field(:filter_criteria, :map, source: :filterCriteria)
    # IANA zone the archive's folder names are stamped in: the org's display zone at request time.
    field(:time_zone, :string, source: :timeZone)
    # Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
    # older pod dropped as unknown during a rolling upgrade.
    field(:dispatch_count, :integer, default: 0, source: :dispatchCount)
    # The pre-flight's counts, shown while the archive builds.
    field(:file_count, :integer, source: :fileCount)
    field(:response_count, :integer, source: :responseCount)
    # What the archive actually holds, once built.
    field(:included_file_count, :integer, source: :includedFileCount)
    field(:total_bytes, :integer, source: :totalBytes)
    field(:truncated, :boolean, default: false)
    field(:storage_key, :string, source: :storageKey)
    # A short code (`stale`, `storage_upload_failed`, …), never a message or file name.
    field(:error_code, :string, source: :errorCode)
    field(:started_at, Formbricks.PrismaDateTime, source: :startedAt)
    field(:finished_at, Formbricks.PrismaDateTime, source: :finishedAt)
    field(:expires_at, Formbricks.PrismaDateTime, source: :expiresAt)
    field(:workspace_id, :string, source: :workspaceId)
    field(:requested_by_id, :string, source: :requestedById)
    belongs_to(:workspace, Formbricks.Workspace, foreign_key: :workspace_id, define_field: false, type: :string)
    belongs_to(:requested_by, Formbricks.User, foreign_key: :requested_by_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end