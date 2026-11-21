defmodule Formbricks.VerificationToken do
  @moduledoc """
  Stores one-time verification tokens used by Auth.js adapter flows.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          identifier: String.t(),
          token: String.t(),
          expires: DateTime.t()
        }

  schema "VerificationToken" do
    field(:identifier, :string, primary_key: true)
    field(:token, :string, primary_key: true)
    field(:expires, Formbricks.PrismaDateTime)
  end
end