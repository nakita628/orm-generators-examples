defmodule Documenso.WebhookCall do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Documenso.PrismaDateTime, autogenerate: {Documenso.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          status: atom(),
          url: String.t(),
          event: atom(),
          request_body: map() | list() | String.t() | number() | boolean(),
          response_code: integer(),
          response_headers: map() | list() | String.t() | number() | boolean() | nil,
          response_body: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          webhook: Documenso.Webhook.t() | nil
        }

  schema "WebhookCall" do
    field(:id, :string, primary_key: true)
    field(:status, Ecto.Enum, values: [:SUCCESS, :FAILED])
    field(:url, :string)
    field(:event, Ecto.Enum, values: [:DOCUMENT_CREATED, :DOCUMENT_SENT, :DOCUMENT_OPENED, :DOCUMENT_SIGNED, :DOCUMENT_COMPLETED, :DOCUMENT_REJECTED, :DOCUMENT_CANCELLED, :RECIPIENT_EXPIRED, :DOCUMENT_RECIPIENT_COMPLETED, :DOCUMENT_REMINDER_SENT, :TEMPLATE_CREATED, :TEMPLATE_UPDATED, :TEMPLATE_DELETED, :TEMPLATE_USED])
    field(:request_body, Documenso.PrismaJson, source: :requestBody)
    field(:response_code, :integer, source: :responseCode)
    field(:response_headers, Documenso.PrismaJson, source: :responseHeaders)
    field(:response_body, Documenso.PrismaJson, source: :responseBody)
    field(:webhook_id, :string, source: :webhookId)
    belongs_to(:webhook, Documenso.Webhook, foreign_key: :webhook_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end