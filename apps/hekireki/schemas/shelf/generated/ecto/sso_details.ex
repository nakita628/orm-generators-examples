defmodule Shelf.SsoDetails do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          domain: String.t(),
          base_user_group_id: String.t() | nil,
          self_service_group_id: String.t() | nil,
          admin_group_id: String.t() | nil,
          manager_group_id: String.t() | nil,
          require_sso_login: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organizations: [Shelf.Organization.t()]
        }

  schema "SsoDetails" do
    field(:id, :string, primary_key: true)
    field(:domain, :string)
    field(:base_user_group_id, :string, source: :baseUserGroupId)
    field(:self_service_group_id, :string, source: :selfServiceGroupId)
    field(:admin_group_id, :string, source: :adminGroupId)
    field(:manager_group_id, :string, source: :managerGroupId)
    field(:require_sso_login, :boolean, default: true, source: :requireSsoLogin)
    has_many(:organizations, Shelf.Organization, foreign_key: :sso_details_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end