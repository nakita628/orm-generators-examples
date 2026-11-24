defmodule Umami.Revenue do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          event_id: Ecto.UUID.t(),
          event_name: String.t(),
          currency: String.t(),
          revenue: Decimal.t() | nil,
          inserted_at: DateTime.t() | nil,
          website: Umami.Website.t() | nil,
          session: Umami.Session.t() | nil
        }

  schema "revenue" do
    field(:id, Ecto.UUID, primary_key: true, source: :revenue_id)
    field(:event_id, Ecto.UUID)
    field(:event_name, :string)
    field(:currency, :string)
    field(:revenue, :decimal)
    belongs_to(:website, Umami.Website, foreign_key: :website_id, type: Ecto.UUID)
    belongs_to(:session, Umami.Session, foreign_key: :session_id, type: Ecto.UUID)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end