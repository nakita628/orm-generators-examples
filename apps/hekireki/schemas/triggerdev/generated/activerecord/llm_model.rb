# A known LLM model or model pattern for cost tracking
class LlmModel < ApplicationRecord
  self.implicit_order_column = "created_at"

  def self.instance_method_already_implemented?(method_name)
    dangerous_attribute_method?(method_name) || super
  end
  private_class_method :instance_method_already_implemented?

  def read_attribute_for_serialization(name)
    if has_attribute?(name) && self.class.dangerous_attribute_method?(name)
      _read_attribute(name.to_s)
    else
      super
    end
  end

  attribute :id, default: -> { Cuid.generate }
  attribute :source, default: "default"
  attribute :capabilities, default: -> { [] }
  attribute :is_hidden, default: false
  alias_attribute :model_name_value, :model_name

  validates :friendly_id, presence: true, uniqueness: true
  validates :model_name_value, presence: true,
                               uniqueness: { scope: %i[project_id start_date], conditions: -> { where.not(project_id: nil).where.not(start_date: nil) } }
  validates :match_pattern, presence: true

  belongs_to :project, optional: true, inverse_of: :llm_models
  has_many :pricing_tiers, class_name: "LlmPricingTier", foreign_key: "model_id", inverse_of: :model,
                           dependent: :destroy
  has_many :prices, class_name: "LlmPrice", foreign_key: "model_id", inverse_of: :model, dependent: :destroy
end
