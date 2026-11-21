defmodule Shelf.Invite do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          invitee_email: String.t(),
          status: atom(),
          invite_code: String.t(),
          roles: atom(),
          invite_message: String.t() | nil,
          expires_at: DateTime.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          inviter: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil,
          invitee_user: Shelf.User.t() | nil,
          invitee_team_member: Shelf.TeamMember.t() | nil
        }

  schema "Invite" do
    field(:id, :string, primary_key: true)
    field(:invitee_email, :string, source: :inviteeEmail)
    field(:status, Ecto.Enum, values: [:PENDING, :ACCEPTED, :REJECTED, :INVALIDATED], default: :PENDING)
    field(:invite_code, :string, source: :inviteCode)
    field(:roles, {:array, Ecto.Enum}, values: [:ADMIN, :BASE, :OWNER, :SELF_SERVICE, :MANAGER])
    field(:invite_message, :string, source: :inviteMessage)
    field(:expires_at, Shelf.PrismaDateTime, source: :expiresAt)
    field(:inviter_id, :string, source: :inviterId)
    field(:organization_id, :string, source: :organizationId)
    field(:invitee_user_id, :string, source: :inviteeUserId)
    field(:team_member_id, :string, source: :teamMemberId)
    belongs_to(:inviter, Shelf.User, foreign_key: :inviter_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:invitee_user, Shelf.User, foreign_key: :invitee_user_id, define_field: false, type: :string)
    belongs_to(:invitee_team_member, Shelf.TeamMember, foreign_key: :team_member_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end