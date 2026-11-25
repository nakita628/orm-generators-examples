defmodule LangWatch.NurturingErasedUser do
  @moduledoc """
  Each account nurturing erased, as a SHA-256 of its id and never the id: a gateway key fact
  naming it, however late it lands, stores no creator or holder. Platform-wide, like the erasure.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          user_id_hash: String.t(),
          erased_at: DateTime.t()
        }

  schema "NurturingErasedUser" do
    field(:user_id_hash, :string, primary_key: true, source: :userIdHash)
    field(:erased_at, LangWatch.PrismaDateTime, autogenerate: true, source: :erasedAt)
  end
end