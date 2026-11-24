defmodule Rallly.SpaceWebhook do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Rallly.PrismaDateTime, autogenerate: {Rallly.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          url: String.t(),
          secret: String.t(),
          events: [String.t()],
          enabled: boolean(),
          version: String.t(),
          cursor: DateTime.t(),
          last_delivered_at: DateTime.t() | nil,
          last_failed_at: DateTime.t() | nil,
          consecutive_failures: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          space: Rallly.Space.t() | nil,
          deliveries: [Rallly.WebhookDelivery.t()]
        }

  schema "space_webhooks" do
    field(:id, :string, primary_key: true)
    field(:url, :string)
    # AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation
    field(:secret, :string)
    # Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations
    field(:events, {:array, :string})
    field(:enabled, :boolean, default: true)
    # The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape
    field(:version, :string)
    # Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history
    field(:cursor, Rallly.PrismaDateTime, autogenerate: true)
    field(:last_delivered_at, Rallly.PrismaDateTime)
    field(:last_failed_at, Rallly.PrismaDateTime)
    # Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap
    field(:consecutive_failures, :integer, default: 0)
    belongs_to(:space, Rallly.Space, foreign_key: :space_id, type: :string)
    has_many(:deliveries, Rallly.WebhookDelivery, foreign_key: :webhook_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end