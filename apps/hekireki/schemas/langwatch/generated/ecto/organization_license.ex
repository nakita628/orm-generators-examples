defmodule LangWatch.OrganizationLicense do
  @moduledoc """
  The licence an organization activated and its dates, owned by licensing
  (round 37 D6). A row whose key is null is a cleared licence. An organization
  with no row is read from its Organization columns until the copy step runs.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          organization_id: String.t(),
          license_key: String.t() | nil,
          expires_at: DateTime.t() | nil,
          validated_at: DateTime.t() | nil,
          updated_at: DateTime.t()
        }

  schema "OrganizationLicense" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:license_key, :string, source: :licenseKey)
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    field(:validated_at, LangWatch.PrismaDateTime, source: :validatedAt)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end