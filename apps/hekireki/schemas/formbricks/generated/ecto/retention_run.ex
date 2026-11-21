defmodule Formbricks.RetentionRun do
  @moduledoc """
  One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
  while the run is in progress, or if it never finished.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          entity: atom(),
          kind: atom(),
          started_at: DateTime.t(),
          finished_at: DateTime.t() | nil,
          notified_count: integer(),
          archived_count: integer(),
          deleted_count: integer(),
          skipped_count: integer(),
          has_changes: boolean(),
          scan_cursor: String.t() | nil,
          organization: Formbricks.Organization.t() | nil,
          items: [Formbricks.RetentionRunItem.t()]
        }

  schema "RetentionRun" do
    field(:id, :string, primary_key: true)
    field(:entity, Ecto.Enum, values: [:responses, :surveys, :members])
    field(:kind, Ecto.Enum, values: [:sweep, :purge], default: :sweep)
    field(:started_at, Formbricks.PrismaDateTime, autogenerate: true, source: :startedAt)
    field(:finished_at, Formbricks.PrismaDateTime, source: :finishedAt)
    field(:notified_count, :integer, default: 0, source: :notifiedCount)
    # Archived surveys, or deactivated members.
    field(:archived_count, :integer, default: 0, source: :archivedCount)
    # Deleted responses, or, on a purge run, deleted surveys.
    field(:deleted_count, :integer, default: 0, source: :deletedCount)
    field(:skipped_count, :integer, default: 0, source: :skippedCount)
    # False when the run changed nothing. History hides those by default.
    field(:has_changes, :boolean, default: false, source: :hasChanges)
    # Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
    # of rescanning the same first targets every night. Null when the scan reached the end.
    field(:scan_cursor, :string, source: :scanCursor)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Formbricks.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:items, Formbricks.RetentionRunItem, foreign_key: :run_id, references: :id)
  end
end