defmodule LangWatch.SsoAuthenticationActivity do
  @moduledoc """
  A successful callback through one SSO connection. Identifier.lastUsedAt
  also advances during routing and resolution, so it cannot prove that a
  member actually authenticated through the replacement before cutover.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          connection_id: String.t(),
          user_id: String.t(),
          authenticated_at: DateTime.t(),
          provider_account_id: String.t() | nil
        }

  schema "SsoAuthenticationActivity" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:connection_id, :string, source: :connectionId)
    field(:user_id, :string, source: :userId)
    field(:authenticated_at, LangWatch.PrismaDateTime, source: :authenticatedAt)
    # The provider-side account the assertion named; null before #7633's arrival door recorded it.
    field(:provider_account_id, :string, source: :providerAccountId)
  end
end