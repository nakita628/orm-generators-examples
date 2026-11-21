defmodule Shelf.PartialBookingCheckin do
  @moduledoc """
  A progressive check-in SESSION, mirroring PartialBookingCheckout.

  `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
  QUANTITY_TRACKED only once its remaining units hit zero (partial returns
  live in `ConsumptionLog`). Completing a booking outright writes no row here,
  so the same warning applies — absence is not evidence of anything.
  `BookingAsset.checkedInAt` is the per-slice fact.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          asset_ids: [String.t()],
          checkin_count: integer(),
          checkin_timestamp: DateTime.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          booking: Shelf.Booking.t() | nil,
          checked_in_by: Shelf.User.t() | nil
        }

  schema "PartialBookingCheckin" do
    field(:id, :string, primary_key: true)
    field(:asset_ids, {:array, :string}, source: :assetIds)
    field(:checkin_count, :integer, source: :checkinCount)
    field(:checkin_timestamp, Shelf.PrismaDateTime, autogenerate: true, source: :checkinTimestamp)
    field(:booking_id, :string, source: :bookingId)
    field(:checked_in_by_id, :string, source: :checkedInById)
    belongs_to(:booking, Shelf.Booking, foreign_key: :booking_id, define_field: false, type: :string)
    belongs_to(:checked_in_by, Shelf.User, foreign_key: :checked_in_by_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end