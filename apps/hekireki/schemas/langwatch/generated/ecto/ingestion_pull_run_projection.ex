defmodule LangWatch.IngestionPullRunProjection do
  @moduledoc """
  Rebuildable per-source pull checkpoint and operator-facing run status.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          source_id: String.t(),
          project_id: String.t(),
          enabled: boolean(),
          cron: String.t() | nil,
          cursor: String.t() | nil,
          last_run_at: float() | nil,
          last_run_outcome: String.t() | nil,
          last_run_event_count: integer(),
          last_run_error: String.t() | nil,
          last_run_error_code: String.t() | nil,
          consecutive_errors: integer(),
          last_success_at: float() | nil,
          last_run_scheduled_for: float() | nil,
          last_read_through_at: float() | nil,
          last_run_completeness: String.t() | nil,
          last_agents_listing_at: float() | nil,
          last_agents_listing_outcome: String.t() | nil,
          last_agents_listing_count: integer() | nil,
          last_agents_listing_reason: String.t() | nil,
          last_agents_listing_status: integer() | nil,
          last_people_listing_at: float() | nil,
          last_people_listing_outcome: String.t() | nil,
          last_people_directory_count: integer() | nil,
          last_people_withheld_count: integer() | nil,
          last_people_listing_reason: String.t() | nil,
          last_people_listing_status: integer() | nil,
          created_at: float(),
          updated_at: float(),
          occurred_at: float(),
          accepted_at: float(),
          last_event_id: String.t(),
          projection_version: String.t()
        }

  schema "IngestionPullRunProjection" do
    field(:id, :string, primary_key: true)
    field(:source_id, :string, source: :sourceId)
    field(:project_id, :string, source: :projectId)
    field(:enabled, :boolean, default: false)
    field(:cron, :string)
    field(:cursor, :string)
    field(:last_run_at, :float, source: :lastRunAt)
    field(:last_run_outcome, :string, source: :lastRunOutcome)
    field(:last_run_event_count, :integer, default: 0, source: :lastRunEventCount)
    field(:last_run_error, :string, source: :lastRunError)
    field(:last_run_error_code, :string, source: :lastRunErrorCode)
    field(:consecutive_errors, :integer, default: 0, source: :consecutiveErrors)
    # When a run last completed without erroring. Backfilled by replay, which
    # is why the projection version moved with this column.
    field(:last_success_at, :float, source: :lastSuccessAt)
    # Identity of the run whose outcome this row reflects. Outcomes from a
    # superseded run arrive with a smaller value and are ignored, so a late
    # completion cannot regress Cursor or the status fields.
    field(:last_run_scheduled_for, :float, source: :lastRunScheduledFor)
    # ADR-128 - how far the last run read, in epoch milliseconds, and whether
    # it reached the end. The fold writes both; the IngestionSource row mirrors
    # them for the screens. Null read-through means no run has reported one.
    # Null completeness reads as unknown, never as either answer, which is what
    # every row written before this existed holds.
    #
    # Nullable on purpose, so replay is not required to make this migration
    # safe: a row that has never seen a run carrying these stays null and every
    # reader already handles unknown.
    # Spec: specs/governance/ingestion-source-health.feature
    field(:last_read_through_at, :float, source: :lastReadThroughAt)
    field(:last_run_completeness, :string, source: :lastRunCompleteness)
    # The last time somebody asked this source to enumerate its agents, and how
    # that ask ended. A listing is a different scope on the same connection as
    # the pull above: a credential that cannot enumerate a directory still pulls
    # cost perfectly, so no health, coverage or completeness reader may treat a
    # refusal here as the source failing to pull.
    #
    # Count is set ONLY when the outcome is listed, and stays null on a refusal.
    # Zero is a real answer meaning the provider returned an empty list, so a
    # zero default would manufacture that answer for a source that refused.
    # Reason is set only on a refusal and is a plain string, never an enum and
    # never a foreign key, because the log outlives the vocabulary. Status is
    # nullable even when refused, because a refusal we generated ourselves has
    # no HTTP status, and it is never rendered to a customer.
    #
    # The two `...ListingAt` columns hold the instant the ASK was accepted, not
    # the instant its answer was recorded, and the fold orders outcomes by them
    # so a superseded listing cannot overwrite the one that replaced it. The
    # name predates that job and is kept because renaming a column is a
    # migration; the fold's own field docs are the statement of record.
    field(:last_agents_listing_at, :float, source: :lastAgentsListingAt)
    field(:last_agents_listing_outcome, :string, source: :lastAgentsListingOutcome)
    field(:last_agents_listing_count, :integer, source: :lastAgentsListingCount)
    field(:last_agents_listing_reason, :string, source: :lastAgentsListingReason)
    field(:last_agents_listing_status, :integer, source: :lastAgentsListingStatus)
    # The same for people, in its own set of columns rather than sharing the
    # agents set with a kind discriminator. A source can be refused for people
    # and fine for agents, and one shared set holds only the most recent listing
    # of either kind, so syncing people would erase an agents refusal and the
    # agents page would go back to not knowing after briefly knowing.
    #
    # LastPeopleDirectoryCount is everyone the provider's directory named.
    # LastPeopleWithheldCount is how many of those this deployment does not hold
    # because erasure suppression removed them: a SUBSET of the directory count,
    # never an addition to it. Adding the two counts the same people twice.
    # Subtracting is the valid arithmetic, and the surviving total is that
    # subtraction rather than a third column that could disagree.
    #
    # The withheld count is safe to show as a CURRENT figure and never as a
    # series: it moving from zero to one at a known moment says an erasure
    # happened then, which on a small tenant identifies the person as surely as
    # a name would. So no trend line, no history drawer, no per-run export, and
    # never beside a per-person list.
    field(:last_people_listing_at, :float, source: :lastPeopleListingAt)
    field(:last_people_listing_outcome, :string, source: :lastPeopleListingOutcome)
    field(:last_people_directory_count, :integer, source: :lastPeopleDirectoryCount)
    field(:last_people_withheld_count, :integer, source: :lastPeopleWithheldCount)
    field(:last_people_listing_reason, :string, source: :lastPeopleListingReason)
    field(:last_people_listing_status, :integer, source: :lastPeopleListingStatus)
    field(:created_at, :float, source: :createdAt)
    field(:updated_at, :float, source: :updatedAt)
    field(:occurred_at, :float, source: :occurredAt)
    field(:accepted_at, :float, source: :acceptedAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:projection_version, :string, source: :projectionVersion)
  end
end