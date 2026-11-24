defmodule Umami.EventData do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          data_key: String.t(),
          string_value: String.t() | nil,
          number_value: Decimal.t() | nil,
          date_value: DateTime.t() | nil,
          data_type: integer(),
          inserted_at: DateTime.t() | nil,
          website: Umami.Website.t() | nil,
          website_event: Umami.WebsiteEvent.t() | nil
        }

  schema "event_data" do
    field(:id, Ecto.UUID, primary_key: true, source: :event_data_id)
    field(:data_key, :string)
    field(:string_value, :string)
    field(:number_value, :decimal)
    field(:date_value, Umami.PrismaDateTime)
    field(:data_type, :integer)
    belongs_to(:website, Umami.Website, foreign_key: :website_id, type: Ecto.UUID)
    belongs_to(:website_event, Umami.WebsiteEvent, foreign_key: :website_event_id, type: Ecto.UUID)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end