defmodule LangWatch.AuthzOrganizationPolicy do
  @moduledoc """
  ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          organization_id: String.t(),
          braids_organization_admins_only: boolean(),
          braids_enabled: boolean(),
          updated_by_user_id: String.t(),
          updated_at: DateTime.t()
        }

  schema "AuthzOrganizationPolicy" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:braids_organization_admins_only, :boolean, default: false, source: :braidsOrganizationAdminsOnly)
    field(:braids_enabled, :boolean, default: false, source: :braidsEnabled)
    field(:updated_by_user_id, :string, source: :updatedByUserId)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end