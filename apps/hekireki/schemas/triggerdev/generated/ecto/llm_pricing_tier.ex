defmodule Trigger.LlmPricingTier do
  @moduledoc """
  A pricing tier for a model (supports volume-based or conditional pricing)
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          is_default: boolean(),
          priority: integer(),
          conditions: map() | list() | String.t() | number() | boolean(),
          model: Trigger.LlmModel.t() | nil,
          prices: [Trigger.LlmPrice.t()]
        }

  schema "llm_pricing_tiers" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:is_default, :boolean, default: true)
    field(:priority, :integer, default: 0)
    field(:conditions, Trigger.PrismaJson, default: [])
    belongs_to(:model, Trigger.LlmModel, foreign_key: :model_id, type: :string)
    has_many(:prices, Trigger.LlmPrice, foreign_key: :pricing_tier_id, references: :id)
  end
end