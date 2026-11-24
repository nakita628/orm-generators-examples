defmodule Trigger.PlatformNotification do
  @moduledoc """
  Admin-created notification definitions
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          friendly_id: String.t(),
          title: String.t(),
          payload: map() | list() | String.t() | number() | boolean(),
          surface: atom(),
          scope: atom(),
          starts_at: DateTime.t(),
          ends_at: DateTime.t(),
          cli_max_days_after_first_seen: integer() | nil,
          cli_max_show_count: integer() | nil,
          cli_show_every: integer() | nil,
          priority: integer(),
          is_draft: boolean(),
          archived_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Trigger.User.t() | nil,
          organization: Trigger.Organization.t() | nil,
          project: Trigger.Project.t() | nil,
          interactions: [Trigger.PlatformNotificationInteraction.t()]
        }

  schema "PlatformNotification" do
    field(:id, :string, primary_key: true)
    field(:friendly_id, :string, source: :friendlyId)
    # Admin-facing title for identification in admin tools
    field(:title, :string)
    # Versioned JSON rendering content (see payload schema in spec)
    field(:payload, Trigger.PrismaJson)
    field(:surface, Ecto.Enum, values: [:WEBAPP, :CLI])
    field(:scope, Ecto.Enum, values: [:USER, :PROJECT, :ORGANIZATION, :GLOBAL])
    # When notification becomes active
    field(:starts_at, Trigger.PrismaDateTime, source: :startsAt)
    # When notification expires
    field(:ends_at, Trigger.PrismaDateTime, source: :endsAt)
    # CLI: auto-expire N days after user first saw it
    field(:cli_max_days_after_first_seen, :integer, source: :cliMaxDaysAfterFirstSeen)
    # CLI: auto-expire after shown N times to user
    field(:cli_max_show_count, :integer, source: :cliMaxShowCount)
    # CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…)
    field(:cli_show_every, :integer, source: :cliShowEvery)
    # Ordering within same scope level (higher = more important)
    field(:priority, :integer, default: 0)
    # Staged-but-unpublished. While true the notification is hidden from all
    # user-facing reads regardless of startsAt/endsAt; publishing sets real
    # dates and flips this to false.
    field(:is_draft, :boolean, default: false, source: :isDraft)
    # Soft delete
    field(:archived_at, Trigger.PrismaDateTime, source: :archivedAt)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    field(:project_id, :string, source: :projectId)
    # Set when scope = USER
    belongs_to(:user, Trigger.User, foreign_key: :user_id, define_field: false, type: :string)
    # Set when scope = ORGANIZATION
    belongs_to(:organization, Trigger.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    # Set when scope = PROJECT
    belongs_to(:project, Trigger.Project, foreign_key: :project_id, define_field: false, type: :string)
    has_many(:interactions, Trigger.PlatformNotificationInteraction, foreign_key: :notification_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end