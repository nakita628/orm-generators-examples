defmodule Teable.DomainEventInbox do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Teable.PrismaDateTime, autogenerate: {Teable.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          consumer_id: String.t(),
          event_id: String.t(),
          inserted_at: DateTime.t()
        }

  schema "domain_event_inbox" do
    field(:consumer_id, :string, primary_key: true)
    field(:event_id, :string, primary_key: true)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end