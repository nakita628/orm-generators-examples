defmodule Shelf.RoleChangeLog do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          previous_role: atom(),
          new_role: atom(),
          source: atom(),
          inserted_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          changed_by: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "RoleChangeLog" do
    field(:id, :string, primary_key: true)
    field(:previous_role, Ecto.Enum, values: [:ADMIN, :BASE, :OWNER, :SELF_SERVICE, :MANAGER], source: :previousRole)
    field(:new_role, Ecto.Enum, values: [:ADMIN, :BASE, :OWNER, :SELF_SERVICE, :MANAGER], source: :newRole)
    # How the change was made. SSO rows name the member themself as `changedBy`.
    field(:source, Ecto.Enum, values: [:MANUAL, :SSO], default: :MANUAL)
    field(:user_id, :string, source: :userId)
    field(:changed_by_id, :string, source: :changedById)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:changed_by, Shelf.User, foreign_key: :changed_by_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end