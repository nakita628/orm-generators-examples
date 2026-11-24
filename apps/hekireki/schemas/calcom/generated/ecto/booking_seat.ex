defmodule Calcom.BookingSeat do
  use Ecto.Schema

  @primary_key {:id, :id, autogenerate: true}

  @type t :: %__MODULE__{
          id: integer(),
          reference_uid: String.t(),
          data: map() | list() | String.t() | number() | boolean() | nil,
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          booking: Calcom.Booking.t() | nil,
          attendee: Calcom.Attendee.t() | nil
        }

  schema "BookingSeat" do
    field(:reference_uid, :string, source: :referenceUid)
    # @zod.import(["import { bookingSeatDataSchema } from '../../zod-utils'"]).custom.use(bookingSeatDataSchema)
    field(:data, Calcom.PrismaJson)
    field(:metadata, Calcom.PrismaJson)
    field(:booking_id, :id, source: :bookingId)
    field(:attendee_id, :id, source: :attendeeId)
    belongs_to(:booking, Calcom.Booking, foreign_key: :booking_id, define_field: false)
    belongs_to(:attendee, Calcom.Attendee, foreign_key: :attendee_id, define_field: false)
  end
end