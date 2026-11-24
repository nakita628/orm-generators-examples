defmodule LiteLLM.LiteLLMSubscriptionAccountTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :subscription_account_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          custom_llm_provider: String.t(),
          account_id: String.t(),
          label: String.t() | nil,
          monthly_fee: float(),
          currency: String.t(),
          billing_period_start: String.t(),
          created_by: String.t(),
          updated_by: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_SubscriptionAccountTable" do
    field(:custom_llm_provider, :string)
    field(:account_id, :string)
    field(:label, :string)
    field(:monthly_fee, :float)
    field(:currency, :string)
    field(:billing_period_start, :string)
    field(:created_by, :string)
    field(:updated_by, :string)
    timestamps(inserted_at_source: :created_at)
  end
end