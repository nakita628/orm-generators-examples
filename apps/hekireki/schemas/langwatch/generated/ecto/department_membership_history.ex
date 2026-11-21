defmodule LangWatch.DepartmentMembershipHistory do
  @moduledoc """
  One dated person → department link (ADR-128 §13, #7882).

  `OrganizationUser.departmentId` stays the CURRENT assignment every screen
  reads; these rows are its history, appended by the same `assignUser` write
  so the two cannot drift. They exist for the cost reads: January's spend
  stays with January's department, so a read resolves against the link that
  was open on the cost's day, not against today's pointer. Members with no
  link on a day resolve to "unassigned", never to an error.

  Same shape as IdentityMatch's dated links, same one-open-link guard: the
  partial unique index in the migration rejects a second open row per
  (organizationId, userId) with SQLSTATE 23505.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          user_id: String.t(),
          department_id: String.t(),
          valid_from: DateTime.t(),
          valid_to: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "DepartmentMembershipHistory" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:user_id, :string, source: :userId)
    field(:department_id, :string, source: :departmentId)
    field(:valid_from, LangWatch.PrismaDateTime, source: :validFrom)
    # Null means the link is open. Reassignment and clearing close it.
    field(:valid_to, LangWatch.PrismaDateTime, source: :validTo)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end