defmodule Umami.TeamUser do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          role: String.t(),
          inserted_at: DateTime.t() | nil,
          updated_at: DateTime.t() | nil,
          team: Umami.Team.t() | nil,
          user: Umami.User.t() | nil
        }

  schema "team_user" do
    field(:id, Ecto.UUID, primary_key: true, source: :team_user_id)
    field(:role, :string)
    belongs_to(:team, Umami.Team, foreign_key: :team_id, type: Ecto.UUID)
    belongs_to(:user, Umami.User, foreign_key: :user_id, type: Ecto.UUID)
    timestamps(inserted_at_source: :created_at)
  end
end