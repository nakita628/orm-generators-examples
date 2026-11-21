defmodule Shelf.PartialBookingCheckout do
  @moduledoc """
  A progressive checkout SESSION: one scan batch, covering N assets via the
  positional `assetIds` / `quantities` / `bookingAssetIds` arrays.

  Records sessions and the units they claimed — NOT whether an asset is out.
  The all-at-once checkout (`checkoutBooking`, the Check out button) flips
  statuses and writes NO row here, so **absence proves nothing** and must never
  be read as "was not checked out". Reading it that way is what let a single
  later scan mark every button-checked-out asset on a booking as
  "never checked out". `BookingAsset.checkedOutAt` answers that question.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          asset_ids: [String.t()],
          quantities: [integer()],
          booking_asset_ids: [String.t()],
          checkout_count: integer(),
          checkout_timestamp: DateTime.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          booking: Shelf.Booking.t() | nil,
          checked_out_by: Shelf.User.t() | nil
        }

  schema "PartialBookingCheckout" do
    field(:id, :string, primary_key: true)
    field(:asset_ids, {:array, :string}, source: :assetIds)
    field(:quantities, {:array, :integer})
    field(:booking_asset_ids, {:array, :string}, source: :bookingAssetIds)
    field(:checkout_count, :integer, source: :checkoutCount)
    field(:checkout_timestamp, Shelf.PrismaDateTime, autogenerate: true, source: :checkoutTimestamp)
    field(:booking_id, :string, source: :bookingId)
    field(:checked_out_by_id, :string, source: :checkedOutById)
    belongs_to(:booking, Shelf.Booking, foreign_key: :booking_id, define_field: false, type: :string)
    belongs_to(:checked_out_by, Shelf.User, foreign_key: :checked_out_by_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end