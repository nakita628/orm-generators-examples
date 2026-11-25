defmodule Langfuse.GatewayGuardrailAssignment do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Langfuse.PrismaDateTime, autogenerate: {Langfuse.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          scope: atom(),
          inserted_at: DateTime.t(),
          guardrail: Langfuse.GatewayGuardrail.t() | nil,
          organization: Langfuse.Organization.t() | nil,
          api_key: Langfuse.ApiKey.t() | nil,
          user: Langfuse.User.t() | nil
        }

  schema "gateway_guardrail_assignments" do
    field(:id, :string, primary_key: true)
    field(:scope, Ecto.Enum, values: [:ORGANIZATION, :API_KEY, :USER])
    belongs_to(:guardrail, Langfuse.GatewayGuardrail, foreign_key: :guardrail_id, type: :string)
    belongs_to(:organization, Langfuse.Organization, foreign_key: :organization_id, type: :string)
    belongs_to(:api_key, Langfuse.ApiKey, foreign_key: :api_key_id, type: :string)
    belongs_to(:user, Langfuse.User, foreign_key: :user_id, type: :string)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end