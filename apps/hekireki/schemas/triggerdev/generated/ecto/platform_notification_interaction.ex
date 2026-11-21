defmodule Trigger.PlatformNotificationInteraction do
  @moduledoc """
  Per-user tracking of notification views and dismissals
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          first_seen_at: DateTime.t(),
          show_count: integer(),
          webapp_dismissed_at: DateTime.t() | nil,
          webapp_clicked_at: DateTime.t() | nil,
          cli_dismissed_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          notification: Trigger.PlatformNotification.t() | nil,
          user: Trigger.User.t() | nil
        }

  schema "PlatformNotificationInteraction" do
    field(:id, :string, primary_key: true)
    # Set by beacon/CLI GET on first impression
    field(:first_seen_at, Trigger.PrismaDateTime, autogenerate: true, source: :firstSeenAt)
    # Times shown (incremented per beacon/CLI view)
    field(:show_count, :integer, default: 1, source: :showCount)
    # User dismissed in webapp
    field(:webapp_dismissed_at, Trigger.PrismaDateTime, source: :webappDismissedAt)
    # User clicked a link in webapp
    field(:webapp_clicked_at, Trigger.PrismaDateTime, source: :webappClickedAt)
    # Auto-dismissed or expired in CLI
    field(:cli_dismissed_at, Trigger.PrismaDateTime, source: :cliDismissedAt)
    field(:notification_id, :string, source: :notificationId)
    field(:user_id, :string, source: :userId)
    belongs_to(:notification, Trigger.PlatformNotification, foreign_key: :notification_id, define_field: false, type: :string)
    belongs_to(:user, Trigger.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end