defmodule TypeHero.Tag do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          tag: atom(),
          challenges: [TypeHero.Challenge.t()]
        }

  schema "Tag" do
    field(:tag, Ecto.Enum, values: [:POPULAR, :NEWEST])
    many_to_many(:challenges, TypeHero.Challenge, join_through: "_ChallengeToTag", join_keys: [B: :id, A: :id])
  end
end