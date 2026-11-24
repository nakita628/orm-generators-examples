defmodule Blinko.Config do
  use Ecto.Schema

  @primary_key {:id, :id, autogenerate: true}

  @type t :: %__MODULE__{
          id: integer(),
          key: String.t(),
          config: map() | list() | String.t() | number() | boolean() | nil,
          user: Blinko.Accounts.t() | nil
        }

  schema "config" do
    field(:key, :string, default: "")
    field(:config, Blinko.PrismaJson)
    field(:user_id, :id, source: :userId)
    belongs_to(:user, Blinko.Accounts, foreign_key: :user_id, define_field: false)
  end
end