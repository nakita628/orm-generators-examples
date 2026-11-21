defmodule LangWatch.WebPushSubscription do
  @moduledoc """
  One browser a person subscribed to Web Push, owned by notification. A
  person holds one row per device; the endpoint names the browser, so a
  browser that subscribes again replaces its own row.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          user_id: String.t(),
          endpoint: String.t(),
          p256dh: String.t(),
          auth: String.t(),
          user_agent: String.t() | nil,
          last_success_at: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "WebPushSubscription" do
    field(:id, :string, primary_key: true)
    # The person the pushes are for. A plain column: the row leaves with the
    # person through notification's own subscriber on the user lifecycle.
    field(:user_id, :string, source: :userId)
    # The push service URL the browser handed out. Unique per browser.
    field(:endpoint, :string)
    # The browser's P-256 public key and auth secret, base64url, as the
    # browser reports them. Payloads are encrypted to these.
    field(:p256dh, :string)
    field(:auth, :string)
    field(:user_agent, :string, source: :userAgent)
    # When the push service last accepted a push for this browser.
    field(:last_success_at, LangWatch.PrismaDateTime, source: :lastSuccessAt)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end