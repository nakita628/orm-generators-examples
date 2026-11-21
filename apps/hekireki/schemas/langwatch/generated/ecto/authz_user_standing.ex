defmodule LangWatch.AuthzUserStanding do
  @moduledoc """
  Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant").
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          user_id: String.t(),
          deactivated_at: DateTime.t() | nil,
          erased_at: DateTime.t() | nil,
          standing_changed_at: DateTime.t()
        }

  schema "AuthzUserStanding" do
    field(:user_id, :string, primary_key: true, source: :userId)
    field(:deactivated_at, LangWatch.PrismaDateTime, source: :deactivatedAt)
    field(:erased_at, LangWatch.PrismaDateTime, source: :erasedAt)
    field(:standing_changed_at, LangWatch.PrismaDateTime, source: :standingChangedAt)
  end
end