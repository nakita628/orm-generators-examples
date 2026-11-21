defmodule LangWatch.UserErasure do
  @moduledoc """
  One row per erased account, holding only its creation time: no id, name or address. It keeps
  the install's first-created account knowable after an erasure deletes that row.
  Spec: modules/user/specs/user.feature
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          account_created_at: DateTime.t(),
          erased_at: DateTime.t()
        }

  schema "UserErasure" do
    field(:id, :string, primary_key: true)
    field(:account_created_at, LangWatch.PrismaDateTime, source: :accountCreatedAt)
    field(:erased_at, LangWatch.PrismaDateTime, autogenerate: true, source: :erasedAt)
  end
end