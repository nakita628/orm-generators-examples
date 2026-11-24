defmodule LiteLLM.LiteLLMManagedObjectTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          unified_object_id: String.t(),
          model_object_id: String.t(),
          file_object: map() | list() | String.t() | number() | boolean(),
          file_purpose: String.t(),
          status: String.t() | nil,
          batch_processed: boolean(),
          created_by: String.t() | nil,
          team_id: String.t() | nil,
          org_id: String.t() | nil,
          api_key: String.t() | nil,
          request_tags: map() | list() | String.t() | number() | boolean() | nil,
          end_user_id: String.t() | nil,
          creator_billing_context: map() | list() | String.t() | number() | boolean() | nil,
          updated_by: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_ManagedObjectTable" do
    field(:unified_object_id, :string)
    field(:model_object_id, :string)
    field(:file_object, LiteLLM.PrismaJson)
    field(:file_purpose, :string)
    field(:status, :string)
    field(:batch_processed, :boolean, default: false)
    field(:created_by, :string)
    field(:team_id, :string)
    field(:org_id, :string)
    field(:api_key, :string)
    field(:request_tags, LiteLLM.PrismaJson, default: [])
    field(:end_user_id, :string)
    field(:creator_billing_context, LiteLLM.PrismaJson)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end