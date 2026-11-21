defmodule Shelf.BookingSettings do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          buffer_start_time: integer(),
          tags_required: boolean(),
          max_booking_length: integer() | nil,
          max_booking_length_skip_closed_days: boolean(),
          auto_archive_bookings: boolean(),
          auto_archive_days: integer(),
          auto_archive_expired_reservations: boolean(),
          require_explicit_checkin_for_admin: boolean(),
          require_explicit_checkin_for_self_service: boolean(),
          require_explicit_checkout_for_admin: boolean(),
          require_explicit_checkout_for_self_service: boolean(),
          count_kits_as_single_unit: boolean(),
          notify_booking_creator: boolean(),
          notify_admins_on_new_booking: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil,
          always_notify_team_members: [Shelf.TeamMember.t()]
        }

  schema "BookingSettings" do
    field(:id, :string, primary_key: true)
    field(:buffer_start_time, :integer, default: 0, source: :bufferStartTime)
    field(:tags_required, :boolean, default: false, source: :tagsRequired)
    field(:max_booking_length, :integer, source: :maxBookingLength)
    field(:max_booking_length_skip_closed_days, :boolean, default: false, source: :maxBookingLengthSkipClosedDays)
    field(:auto_archive_bookings, :boolean, default: false, source: :autoArchiveBookings)
    field(:auto_archive_days, :integer, default: 2, source: :autoArchiveDays)
    field(:auto_archive_expired_reservations, :boolean, default: false, source: :autoArchiveExpiredReservations)
    field(:require_explicit_checkin_for_admin, :boolean, default: false, source: :requireExplicitCheckinForAdmin)
    field(:require_explicit_checkin_for_self_service, :boolean, default: false, source: :requireExplicitCheckinForSelfService)
    field(:require_explicit_checkout_for_admin, :boolean, default: false, source: :requireExplicitCheckoutForAdmin)
    field(:require_explicit_checkout_for_self_service, :boolean, default: false, source: :requireExplicitCheckoutForSelfService)
    field(:count_kits_as_single_unit, :boolean, default: false, source: :countKitsAsSingleUnit)
    field(:notify_booking_creator, :boolean, default: true, source: :notifyBookingCreator)
    field(:notify_admins_on_new_booking, :boolean, default: true, source: :notifyAdminsOnNewBooking)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    many_to_many(:always_notify_team_members, Shelf.TeamMember, join_through: "_BookingSettingsAlwaysNotify", join_keys: [A: :id, B: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end