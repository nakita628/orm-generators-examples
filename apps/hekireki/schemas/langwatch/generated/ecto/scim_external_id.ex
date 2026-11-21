defmodule LangWatch.ScimExternalId do
  @moduledoc """
  D08 — who the directory means, per connection.

  The identity provider's own identifier for a person, which is what
  survives their email changing: an address is not identity, people change
  theirs, and the same person carries different identifiers on two different
  connections. So the PAIR is the key and `externalId` alone never is.

  No relation to `SsoConnection` (that projection carries no foreign keys);
  the `User` relation is real, so a deleted user takes their directory
  identities with them rather than stranding a mapping that resolves to
  nobody.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          organization_id: String.t(),
          id: String.t(),
          connection_id: String.t(),
          external_id: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: LangWatch.User.t() | nil
        }

  schema "ScimExternalId" do
    field(:organization_id, :string, source: :organizationId)
    field(:id, :string, primary_key: true)
    field(:connection_id, :string, source: :connectionId)
    field(:external_id, :string, source: :externalId)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, LangWatch.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end