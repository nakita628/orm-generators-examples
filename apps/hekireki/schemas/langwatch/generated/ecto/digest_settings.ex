defmodule LangWatch.DigestSettings do
  @moduledoc """
  The digest's one settings row. `armed` is the weekly send's master switch; it starts off.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          armed: boolean(),
          last_planned_week: String.t() | nil,
          updated_at: DateTime.t()
        }

  schema "DigestSettings" do
    field(:id, :string, primary_key: true, default: "singleton")
    field(:armed, :boolean, default: false)
    field(:last_planned_week, :string, source: :lastPlannedWeek)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end