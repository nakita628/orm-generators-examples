defmodule LangWatch.AuthzBraidWrite do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          organization_id: String.t(),
          command_id: String.t(),
          grant_id: String.t(),
          reader_project_id: String.t(),
          owner_project_id: String.t(),
          terms_key: String.t(),
          condition: map() | list() | String.t() | number() | boolean(),
          expires_at: DateTime.t() | nil,
          replaces_grant_id: String.t() | nil,
          created_by_user_id: String.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "AuthzBraidWrite" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:command_id, :string, primary_key: true, source: :commandId)
    field(:grant_id, :string, source: :grantId)
    field(:reader_project_id, :string, source: :readerProjectId)
    field(:owner_project_id, :string, source: :ownerProjectId)
    field(:terms_key, :string, source: :termsKey)
    field(:condition, LangWatch.PrismaJson)
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    field(:replaces_grant_id, :string, source: :replacesGrantId)
    field(:created_by_user_id, :string, source: :createdByUserId)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end