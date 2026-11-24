defmodule LiteLLM.LiteLLMRetiredAgentIdentity do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :binding_id}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          provider: String.t(),
          issuer: String.t(),
          tenant_id: String.t(),
          client_id: String.t(),
          agent: LiteLLM.LiteLLMAgentsTable.t() | nil
        }

  schema "LiteLLM_RetiredAgentIdentity" do
    field(:provider, :string)
    field(:issuer, :string)
    field(:tenant_id, :string)
    field(:client_id, :string)
    belongs_to(:agent, LiteLLM.LiteLLMAgentsTable, foreign_key: :agent_id)
  end
end