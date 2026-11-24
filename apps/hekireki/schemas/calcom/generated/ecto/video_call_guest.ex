defmodule Calcom.VideoCallGuest do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Calcom.PrismaDateTime, autogenerate: {Calcom.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          booking_uid: String.t(),
          email: String.t(),
          name: String.t(),
          joined_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "VideoCallGuest" do
    field(:booking_uid, :string, source: :bookingUid)
    field(:email, :string)
    field(:name, :string)
    field(:joined_at, Calcom.PrismaDateTime, source: :joinedAt)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end