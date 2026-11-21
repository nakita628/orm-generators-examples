defmodule LangWatch.NurturingOrganizationCreated do
  @moduledoc """
  When an organization was created, from organization's created event: nurturing's own copy for
  the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
  NurturingOrganization row.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          created_at: DateTime.t()
        }

  schema "NurturingOrganizationCreated" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:created_at, LangWatch.PrismaDateTime, source: :createdAt)
  end
end