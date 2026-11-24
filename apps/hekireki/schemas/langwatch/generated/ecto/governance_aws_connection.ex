defmodule LangWatch.GovernanceAwsConnection do
  @moduledoc """
  An AWS credential a Bedrock kill switch target signs with: a static key pair, a role to assume,
  or both. The secret access key and the external id are stored sealed.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          name: String.t(),
          access_key_id: String.t() | nil,
          encrypted_secret_access_key: String.t() | nil,
          role_arn: String.t() | nil,
          encrypted_external_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "GovernanceAwsConnection" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:name, :string)
    field(:access_key_id, :string, source: :accessKeyId)
    field(:encrypted_secret_access_key, :string, source: :encryptedSecretAccessKey)
    field(:role_arn, :string, source: :roleArn)
    field(:encrypted_external_id, :string, source: :encryptedExternalId)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end