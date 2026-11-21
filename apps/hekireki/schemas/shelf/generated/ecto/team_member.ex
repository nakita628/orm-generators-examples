defmodule Shelf.TeamMember do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          deleted_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil,
          user: Shelf.User.t() | nil,
          custodies: [Shelf.Custody.t()],
          received_invites: [Shelf.Invite.t()],
          kit_custodies: [Shelf.KitCustody.t()],
          bookings: [Shelf.Booking.t()],
          consumption_logs: [Shelf.ConsumptionLog.t()],
          team_member_notes: [Shelf.TeamMemberNote.t()],
          asset_reminders: [Shelf.AssetReminder.t()],
          always_notify_booking_settings: [Shelf.BookingSettings.t()],
          booking_notifications: [Shelf.Booking.t()]
        }

  schema "TeamMember" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:deleted_at, Shelf.PrismaDateTime, source: :deletedAt)
    field(:organization_id, :string, source: :organizationId)
    field(:user_id, :string, source: :userId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    has_many(:custodies, Shelf.Custody, foreign_key: :team_member_id, references: :id)
    has_many(:received_invites, Shelf.Invite, foreign_key: :team_member_id, references: :id)
    has_many(:kit_custodies, Shelf.KitCustody, foreign_key: :custodian_id, references: :id)
    has_many(:bookings, Shelf.Booking, foreign_key: :custodian_team_member_id, references: :id)
    has_many(:consumption_logs, Shelf.ConsumptionLog, foreign_key: :custodian_id, references: :id)
    has_many(:team_member_notes, Shelf.TeamMemberNote, foreign_key: :team_member_id, references: :id)
    many_to_many(:asset_reminders, Shelf.AssetReminder, join_through: "_AssetReminderToTeamMember", join_keys: [B: :id, A: :id])
    many_to_many(:always_notify_booking_settings, Shelf.BookingSettings, join_through: "_BookingSettingsAlwaysNotify", join_keys: [B: :id, A: :id])
    many_to_many(:booking_notifications, Shelf.Booking, join_through: "_BookingNotificationRecipients", join_keys: [B: :id, A: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end