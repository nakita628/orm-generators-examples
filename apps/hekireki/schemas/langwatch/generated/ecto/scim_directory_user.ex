defmodule LangWatch.ScimDirectoryUser do
  @moduledoc """
  A connection's managed users, including resources without an externalId.
  Scalar keys only; readers exclude users that no longer exist.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          connection_id: String.t(),
          user_id: String.t()
        }

  schema "ScimDirectoryUser" do
    field(:organization_id, :string, source: :organizationId)
    field(:connection_id, :string, primary_key: true, source: :connectionId)
    field(:user_id, :string, primary_key: true, source: :userId)
  end
end