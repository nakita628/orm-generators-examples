defmodule OpenSaas.ProcessedWebhookEvent do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: OpenSaas.PrismaDateTime, autogenerate: {OpenSaas.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          payment_processor_event_id: String.t(),
          event_type: String.t(),
          inserted_at: DateTime.t()
        }

  schema "ProcessedWebhookEvent" do
    field(:payment_processor_event_id, :string, source: :paymentProcessorEventId)
    field(:event_type, :string, source: :eventType)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end