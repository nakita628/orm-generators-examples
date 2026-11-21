defmodule Trigger.BulkActionGroup do
  @moduledoc """
  Bulk actions, like canceling and replaying runs
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          friendly_id: String.t(),
          type: atom(),
          status: atom(),
          query_name: String.t() | nil,
          params: map() | nil,
          dedupe_key: String.t() | nil,
          cursor: map() | nil,
          success_count: integer(),
          failure_count: integer(),
          total_count: integer(),
          completion_notification: atom(),
          name: String.t() | nil,
          completed_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: Trigger.Project.t() | nil,
          environment: Trigger.RuntimeEnvironment.t() | nil,
          user: Trigger.User.t() | nil
        }

  schema "BulkActionGroup" do
    field(:id, :string, primary_key: true)
    field(:friendly_id, :string, source: :friendlyId)
    field(:type, Ecto.Enum, values: [:CANCEL, :REPLAY])
    # When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed.
    field(:status, Ecto.Enum, values: [:PENDING, :COMPLETED, :ABORTED], default: :PENDING)
    # The query that will be executed to get the runs to process (if null, we are passing run ids directly)
    field(:query_name, :string, source: :queryName)
    # The params that will be passed to the query
    field(:params, :map)
    # Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions.
    field(:dedupe_key, :string, source: :dedupeKey)
    # The cursor that will be passed to the query (null for the first page)
    field(:cursor, :map)
    # The number of runs that have been processed successfully
    field(:success_count, :integer, default: 0, source: :successCount)
    # The number of runs that have been failed
    field(:failure_count, :integer, default: 0, source: :failureCount)
    # The total number of runs that will be processed
    field(:total_count, :integer, default: 0, source: :totalCount)
    field(:completion_notification, Ecto.Enum, values: [:NONE, :EMAIL], default: :NONE, source: :completionNotification)
    # The name of the bulk action
    field(:name, :string)
    # The time the bulk action was completed
    field(:completed_at, Trigger.PrismaDateTime, source: :completedAt)
    field(:project_id, :string, source: :projectId)
    field(:environment_id, :string, source: :environmentId)
    field(:user_id, :string, source: :userId)
    belongs_to(:project, Trigger.Project, foreign_key: :project_id, define_field: false, type: :string)
    # If this is set then it's a V2 Bulk Action that supports queries
    belongs_to(:environment, Trigger.RuntimeEnvironment, foreign_key: :environment_id, define_field: false, type: :string)
    # The userId who did the bulk action
    belongs_to(:user, Trigger.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end