defmodule Shelf.WorkingHoursOverride do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          date: Date.t(),
          is_open: boolean(),
          open_time: String.t() | nil,
          close_time: String.t() | nil,
          reason: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          working_hours: Shelf.WorkingHours.t() | nil
        }

  schema "WorkingHoursOverride" do
    field(:id, :string, primary_key: true)
    field(:date, Shelf.PrismaDate)
    field(:is_open, :boolean, default: false, source: :isOpen)
    field(:open_time, :string, source: :openTime)
    field(:close_time, :string, source: :closeTime)
    field(:reason, :string)
    field(:working_hours_id, :string, source: :workingHoursId)
    belongs_to(:working_hours, Shelf.WorkingHours, foreign_key: :working_hours_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end