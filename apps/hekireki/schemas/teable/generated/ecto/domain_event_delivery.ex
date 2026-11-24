defmodule Teable.DomainEventDelivery do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Teable.PrismaDateTime, autogenerate: {Teable.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          event_id: String.t(),
          consumer_id: String.t(),
          status: String.t(),
          attempts: integer(),
          max_attempts: integer(),
          lease_token: String.t() | nil,
          lease_expires_at: DateTime.t() | nil,
          next_attempt_at: DateTime.t(),
          last_error: String.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "domain_event_delivery" do
    field(:id, :string, primary_key: true)
    field(:event_id, :string)
    field(:consumer_id, :string)
    field(:status, :string)
    field(:attempts, :integer, default: 0)
    field(:max_attempts, :integer, default: 12)
    field(:lease_token, :string)
    field(:lease_expires_at, Teable.PrismaDateTime)
    field(:next_attempt_at, Teable.PrismaDateTime, autogenerate: true)
    field(:last_error, :string)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end