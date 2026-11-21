defmodule LangWatch.SeatHold do
  @moduledoc """
  A full seat still counted after the person who took it stopped holding it (removed,
  disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
  One per organization and person; a return within the hold takes the seat back.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          user_id: String.t(),
          email: String.t() | nil,
          name: String.t() | nil,
          seat_taken_at: DateTime.t(),
          held_until: DateTime.t(),
          reason: String.t(),
          inserted_at: DateTime.t()
        }

  schema "SeatHold" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:user_id, :string, source: :userId)
    # Display copy for the Users list, kept because a removed person has no membership row.
    field(:email, :string)
    field(:name, :string)
    # When the person took the seat this hold keeps.
    field(:seat_taken_at, LangWatch.PrismaDateTime, source: :seatTakenAt)
    # seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly.
    field(:held_until, LangWatch.PrismaDateTime, source: :heldUntil)
    # REMOVED | DISABLED | DEMOTED
    field(:reason, :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end