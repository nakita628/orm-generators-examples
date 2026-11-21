defmodule Shelf.UserScimExternalId do
  @moduledoc """
  *
  * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
  * object within a specific organization.  A single Shelf user can appear in
  * multiple organizations, each managed by a different identity provider, so the
  * external ID must be scoped to the (user, organization) pair.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          scim_external_id: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "UserScimExternalId" do
    field(:id, :string, primary_key: true)
    field(:scim_external_id, :string, source: :scimExternalId)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end