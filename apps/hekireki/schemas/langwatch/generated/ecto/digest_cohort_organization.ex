defmodule LangWatch.DigestCohortOrganization do
  @moduledoc """
  An organization enrolled in the weekly digest. No row: not enrolled, so nothing sends.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          enrolled_by: String.t(),
          enrolled_at: DateTime.t()
        }

  schema "DigestCohortOrganization" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:enrolled_by, :string, source: :enrolledBy)
    field(:enrolled_at, LangWatch.PrismaDateTime, autogenerate: true, source: :enrolledAt)
  end
end