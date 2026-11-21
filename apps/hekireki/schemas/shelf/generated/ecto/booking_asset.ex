defmodule Shelf.BookingAsset do
  @moduledoc """
  *
  * Explicit pivot table between Booking and Asset.
  * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          quantity: integer(),
          booking_model_request_id: String.t() | nil,
          asset_kit_id: String.t() | nil,
          source_kit_id: String.t() | nil,
          checked_out_at: DateTime.t() | nil,
          checked_out_by_id: String.t() | nil,
          checked_in_at: DateTime.t() | nil,
          checked_in_by_id: String.t() | nil,
          checked_out_quantity: integer(),
          source_location_id: String.t() | nil,
          booking: Shelf.Booking.t() | nil,
          asset: Shelf.Asset.t() | nil
        }

  schema "BookingAsset" do
    field(:id, :string, primary_key: true)
    field(:quantity, :integer, default: 1)
    field(:booking_model_request_id, :string, source: :bookingModelRequestId)
    field(:asset_kit_id, :string, source: :assetKitId)
    field(:source_kit_id, :string, source: :sourceKitId)
    field(:checked_out_at, Shelf.PrismaDateTime, source: :checkedOutAt)
    field(:checked_out_by_id, :string, source: :checkedOutById)
    field(:checked_in_at, Shelf.PrismaDateTime, source: :checkedInAt)
    field(:checked_in_by_id, :string, source: :checkedInById)
    field(:checked_out_quantity, :integer, default: 0, source: :checkedOutQuantity)
    field(:source_location_id, :string, source: :sourceLocationId)
    field(:booking_id, :string, source: :bookingId)
    field(:asset_id, :string, source: :assetId)
    belongs_to(:booking, Shelf.Booking, foreign_key: :booking_id, define_field: false, type: :string)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
  end
end