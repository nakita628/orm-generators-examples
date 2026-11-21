defmodule Calcom.WorkspacePlatform do
  use Ecto.Schema

  @primary_key {:id, :id, autogenerate: true}
  @timestamps_opts [type: Calcom.PrismaDateTime, autogenerate: {Calcom.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: integer(),
          slug: String.t(),
          name: String.t(),
          description: String.t(),
          default_service_account_key: map(),
          enabled: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          delegation_credentials: [Calcom.DelegationCredential.t()]
        }

  schema "WorkspacePlatform" do
    # @zod.string.min(1)
    field(:slug, :string)
    # @zod.string.min(1)
    field(:name, :string)
    field(:description, :string)
    # @zod.import(["import { serviceAccountKeySchema } from '../../zod-utils'"]).custom.use(serviceAccountKeySchema)
    field(:default_service_account_key, :map, source: :defaultServiceAccountKey)
    field(:enabled, :boolean, default: false)
    has_many(:delegation_credentials, Calcom.DelegationCredential, foreign_key: :workspace_platform_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end