defmodule Postiz.UsedCodes do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          code: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Postiz.Organization.t() | nil
        }

  schema "UsedCodes" do
    field(:code, :string)
    field(:org_id, :string, source: :orgId)
    belongs_to(:organization, Postiz.Organization, foreign_key: :org_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end