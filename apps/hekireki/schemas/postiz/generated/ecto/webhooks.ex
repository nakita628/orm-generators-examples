defmodule Postiz.Webhooks do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          url: String.t(),
          deleted_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Postiz.Organization.t() | nil,
          integrations: [Postiz.IntegrationsWebhooks.t()]
        }

  schema "Webhooks" do
    field(:name, :string)
    field(:url, :string)
    field(:deleted_at, Postiz.PrismaDateTime, source: :deletedAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Postiz.Organization, foreign_key: :organization_id, define_field: false)
    has_many(:integrations, Postiz.IntegrationsWebhooks, foreign_key: :webhook_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end