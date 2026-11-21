defmodule Trigger.OrgMemberInvite do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          token: String.t(),
          email: String.t(),
          role: atom(),
          rbac_role_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Trigger.Organization.t() | nil,
          inviter: Trigger.User.t() | nil
        }

  schema "OrgMemberInvite" do
    field(:id, :string, primary_key: true)
    field(:token, :string)
    field(:email, :string)
    field(:role, Ecto.Enum, values: [:ADMIN, :MEMBER], default: :MEMBER)
    # Optional RBAC role to assign on invite acceptance. When set, the
    # accept-invite flow calls the loaded RBAC plugin's setUserRole with
    # this id after creating the OrgMember. Null = legacy behaviour, the
    # runtime fallback derives the role from `role` above.
    #
    # Plain text (not an FK) — the RBAC plugin's RbacRole table lives on
    # a separate schema (Drizzle, not Prisma) so we can't model the FK
    # here. Validation happens at write time (action) and read time
    # (acceptInvite).
    field(:rbac_role_id, :string, source: :rbacRoleId)
    field(:organization_id, :string, source: :organizationId)
    field(:inviter_id, :string, source: :inviterId)
    belongs_to(:organization, Trigger.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:inviter, Trigger.User, foreign_key: :inviter_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end