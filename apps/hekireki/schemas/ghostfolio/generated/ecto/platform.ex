defmodule Ghostfolio.Platform do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t() | nil,
          url: String.t(),
          accounts: [Ghostfolio.Account.t()]
        }

  schema "Platform" do
    field(:name, :string)
    field(:url, :string)
    has_many(:accounts, Ghostfolio.Account, foreign_key: :platform_id)
  end
end