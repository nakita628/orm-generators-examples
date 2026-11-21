defmodule Shelf.ScimToken do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          token_hash: String.t(),
          label: String.t(),
          created_by_id: String.t(),
          last_used_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil
        }

  schema "ScimToken" do
    field(:id, :string, primary_key: true)
    field(:token_hash, :string, source: :tokenHash)
    field(:label, :string)
    field(:created_by_id, :string, source: :createdById)
    field(:last_used_at, Shelf.PrismaDateTime, source: :lastUsedAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end