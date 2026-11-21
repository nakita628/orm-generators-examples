defmodule Shelf.ConsumptionLog do
  @moduledoc """
  *
  * Audit trail for all quantity changes on quantity-tracked assets.
  * Every event records who performed it, who received/returned items,
  * which booking it relates to, and the category of change.
  * Quantity is always positive — direction is determined by category.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          category: atom(),
          quantity: integer(),
          note: String.t() | nil,
          booking_asset_id: String.t() | nil,
          stock_change: integer() | nil,
          inserted_at: DateTime.t(),
          asset: Shelf.Asset.t() | nil,
          performed_by: Shelf.User.t() | nil,
          booking: Shelf.Booking.t() | nil,
          custodian: Shelf.TeamMember.t() | nil,
          location: Shelf.Location.t() | nil
        }

  schema "ConsumptionLog" do
    field(:id, :string, primary_key: true)
    field(:category, Ecto.Enum, values: [:CHECKOUT, :RETURN, :RESTOCK, :ADJUSTMENT, :LOSS, :CONSUME, :DAMAGE, :MOVE, :INITIAL, :OPENING_BALANCE])
    field(:quantity, :integer)
    field(:note, :string)
    field(:booking_asset_id, :string, source: :bookingAssetId)
    field(:stock_change, :integer, source: :stockChange)
    field(:asset_id, :string, source: :assetId)
    field(:user_id, :string, source: :userId)
    field(:booking_id, :string, source: :bookingId)
    field(:custodian_id, :string, source: :custodianId)
    field(:location_id, :string, source: :locationId)
    belongs_to(:asset, Shelf.Asset, foreign_key: :asset_id, define_field: false, type: :string)
    belongs_to(:performed_by, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:booking, Shelf.Booking, foreign_key: :booking_id, define_field: false, type: :string)
    belongs_to(:custodian, Shelf.TeamMember, foreign_key: :custodian_id, define_field: false, type: :string)
    belongs_to(:location, Shelf.Location, foreign_key: :location_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end