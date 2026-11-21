defmodule LangWatch.WebPushVapidKey do
  @moduledoc """
  This installation's VAPID key pair, generated on first use so Web Push works
  without configuration. Fixed id, so the table holds one row: a second
  insert loses the race on the primary key and reads the winner's pair.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          public_key: String.t(),
          private_key_encrypted: String.t(),
          inserted_at: DateTime.t()
        }

  schema "WebPushVapidKey" do
    field(:id, :string, primary_key: true, default: "self")
    # The uncompressed P-256 public key, base64url. The browser subscribes with it.
    field(:public_key, :string, source: :publicKey)
    # The private key, base64url, encrypted with the process encryption key.
    field(:private_key_encrypted, :string, source: :privateKeyEncrypted)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end