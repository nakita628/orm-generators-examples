defmodule Shelf.TeamMemberNote do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          content: String.t(),
          type: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          team_member: Shelf.TeamMember.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "TeamMemberNote" do
    field(:id, :string, primary_key: true)
    field(:content, :string)
    field(:type, Ecto.Enum, values: [:COMMENT, :UPDATE], default: :COMMENT)
    field(:user_id, :string, source: :userId)
    field(:team_member_id, :string, source: :teamMemberId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:team_member, Shelf.TeamMember, foreign_key: :team_member_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end