defmodule LiteLLM.LiteLLMTeamMembership do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          rollover_credit: float() | nil,
          user_id: String.t(),
          team_id: String.t(),
          spend: float(),
          total_spend: float(),
          litellm_budget_table: LiteLLM.LiteLLMBudgetTable.t() | nil
        }

  schema "LiteLLM_TeamMembership" do
    field(:rollover_credit, :float)
    field(:user_id, :string, primary_key: true)
    field(:team_id, :string, primary_key: true)
    field(:spend, :float, default: 0.0)
    field(:total_spend, :float, default: 0.0)
    belongs_to(:litellm_budget_table, LiteLLM.LiteLLMBudgetTable, foreign_key: :budget_id, type: :string)
  end
end