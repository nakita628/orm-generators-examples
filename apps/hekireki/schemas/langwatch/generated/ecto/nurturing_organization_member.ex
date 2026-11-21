defmodule LangWatch.NurturingOrganizationMember do
  @moduledoc """
  Nurturing's own copy of each organization's members, from the owners' member facts (§9):
  organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
  and member_disabled leave, member_enabled returns. Each column keeps its latest time.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          user_id: String.t(),
          joined_at: DateTime.t() | nil,
          left_at: DateTime.t() | nil,
          returned_at: DateTime.t() | nil
        }

  schema "NurturingOrganizationMember" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:user_id, :string, primary_key: true, source: :userId)
    field(:joined_at, LangWatch.PrismaDateTime, source: :joinedAt)
    field(:left_at, LangWatch.PrismaDateTime, source: :leftAt)
    field(:returned_at, LangWatch.PrismaDateTime, source: :returnedAt)
  end
end