defmodule Shelf.UserOrganization do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          roles: atom(),
          calendar_token_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "UserOrganization" do
    field(:id, :string, primary_key: true)
    field(:roles, {:array, Ecto.Enum}, values: [:ADMIN, :BASE, :OWNER, :SELF_SERVICE, :MANAGER])
    # Secret, unguessable token for this member's iCal calendar-feed subscription
    # in this workspace. Null until the user generates a feed; regenerating it
    # rotates the value (invalidating any previously-shared subscription URL).
    # `@unique` doubles as the lookup index for the public feed route.
    field(:calendar_token_id, :string, source: :calendarTokenId)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end