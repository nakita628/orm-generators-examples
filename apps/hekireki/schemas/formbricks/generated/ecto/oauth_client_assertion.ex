defmodule Formbricks.OauthClientAssertion do
  @moduledoc """
  Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
  assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
  table, so omitting it risks an adapter error on a code path we do not control.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          expires_at: DateTime.t()
        }

  schema "oauthClientAssertion" do
    field(:id, :string, primary_key: true)
    field(:expires_at, Formbricks.PrismaDateTime, source: :expiresAt)
  end
end