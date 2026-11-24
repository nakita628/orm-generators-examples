defmodule Umami.ErrorRateLimit do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          website_id: Ecto.UUID.t(),
          bucket: DateTime.t(),
          count: integer()
        }

  schema "error_rate_limit" do
    field(:website_id, Ecto.UUID, primary_key: true)
    field(:bucket, Umami.PrismaDateTime)
    field(:count, :integer)
  end
end