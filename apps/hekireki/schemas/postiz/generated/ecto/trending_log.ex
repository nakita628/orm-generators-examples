defmodule Postiz.TrendingLog do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          language: String.t() | nil,
          date: DateTime.t()
        }

  schema "TrendingLog" do
    field(:language, :string)
    field(:date, Postiz.PrismaDateTime)
  end
end