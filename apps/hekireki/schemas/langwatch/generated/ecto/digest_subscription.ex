defmodule LangWatch.DigestSubscription do
  @moduledoc """
  A person's opt-out of the digest. No row, or no `unsubscribedAt`: subscribed.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          user_id: String.t(),
          unsubscribed_at: DateTime.t() | nil,
          unsubscribed_from_send_id: String.t() | nil,
          updated_at: DateTime.t()
        }

  schema "DigestSubscription" do
    field(:user_id, :string, primary_key: true, source: :userId)
    field(:unsubscribed_at, LangWatch.PrismaDateTime, source: :unsubscribedAt)
    # The send whose unsubscribe link or one-click header the person used.
    field(:unsubscribed_from_send_id, :string, source: :unsubscribedFromSendId)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end