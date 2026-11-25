defmodule Langfuse.GatewayGuardrail do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Langfuse.PrismaDateTime, autogenerate: {Langfuse.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          type: atom(),
          action: atom(),
          config: map() | list() | String.t() | number() | boolean(),
          encrypted_credential: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Langfuse.Organization.t() | nil,
          created_by_user: Langfuse.User.t() | nil,
          assignments: [Langfuse.GatewayGuardrailAssignment.t()]
        }

  schema "gateway_guardrails" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:type, Ecto.Enum, values: [:REGEX, :BEDROCK], default: :REGEX)
    field(:action, Ecto.Enum, values: [:BLOCK, :LOG, :REDACT])
    field(:config, Langfuse.PrismaJson, default: %{})
    field(:encrypted_credential, :string, source: :encrypted_credentials)
    belongs_to(:organization, Langfuse.Organization, foreign_key: :organization_id, type: :string)
    belongs_to(:created_by_user, Langfuse.User, foreign_key: :created_by, type: :string)
    has_many(:assignments, Langfuse.GatewayGuardrailAssignment, foreign_key: :guardrail_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end