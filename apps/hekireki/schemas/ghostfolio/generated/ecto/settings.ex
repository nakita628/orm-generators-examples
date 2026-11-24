defmodule Ghostfolio.Settings do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Ghostfolio.PrismaDateTime, autogenerate: {Ghostfolio.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          user_id: String.t(),
          settings: map() | list() | String.t() | number() | boolean() | nil,
          updated_at: DateTime.t(),
          user: Ghostfolio.User.t() | nil
        }

  schema "Settings" do
    field(:settings, Ghostfolio.PrismaJson)
    field(:user_id, :string, primary_key: true, source: :userId)
    belongs_to(:user, Ghostfolio.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end