defmodule Shelf.WorkingHours do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          enabled: boolean(),
          weekly_schedule: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil,
          overrides: [Shelf.WorkingHoursOverride.t()]
        }

  schema "WorkingHours" do
    field(:id, :string, primary_key: true)
    field(:enabled, :boolean, default: false)
    field(:weekly_schedule, Shelf.PrismaJson, default: %{"0" => %{"isOpen" => false}, "1" => %{"isOpen" => true, "openTime" => "09:00", "closeTime" => "17:00"}, "2" => %{"isOpen" => true, "openTime" => "09:00", "closeTime" => "17:00"}, "3" => %{"isOpen" => true, "openTime" => "09:00", "closeTime" => "17:00"}, "4" => %{"isOpen" => true, "openTime" => "09:00", "closeTime" => "17:00"}, "5" => %{"isOpen" => true, "openTime" => "09:00", "closeTime" => "17:00"}, "6" => %{"isOpen" => false}}, source: :weeklySchedule)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:overrides, Shelf.WorkingHoursOverride, foreign_key: :working_hours_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end