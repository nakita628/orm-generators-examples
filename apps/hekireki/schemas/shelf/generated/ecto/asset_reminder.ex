defmodule Shelf.AssetReminder do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          message: String.t(),
          alert_date_time: DateTime.t(),
          active_scheduler_reference: String.t() | nil,
          recurrence_unit: atom() | nil,
          recurrence_interval: integer() | nil,
          recurrence_timezone: String.t() | nil,
          recurrence_ends_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil,
          asset: Shelf.Asset.t() | nil,
          created_by: Shelf.User.t() | nil,
          team_members: [Shelf.TeamMember.t()]
        }

  schema "AssetReminder" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:message, :string)
    field(:alert_date_time, Shelf.PrismaDateTime, source: :alertDateTime)
    field(:active_scheduler_reference, :string, source: :activeSchedulerReference)
    field(:recurrence_unit, Ecto.Enum, values: [:DAY, :WEEK, :MONTH, :YEAR], source: :recurrenceUnit)
    field(:recurrence_interval, :integer, source: :recurrenceInterval)
    field(:recurrence_timezone, :string, source: :recurrenceTimezone)
    field(:recurrence_ends_at, Shelf.PrismaDateTime, source: :recurrenceEndsAt)
    field(:organization_id, :string, source: :organizationId)
    field(:asset_id, :string, source: :assetId)
    field(:created_by_id, :string, source: :createdById)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:created_by, Shelf.User, foreign_key: :created_by_id, define_field: false, type: :string)
    many_to_many(:team_members, Shelf.TeamMember, join_through: "_AssetReminderToTeamMember", join_keys: [A: :id, B: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end