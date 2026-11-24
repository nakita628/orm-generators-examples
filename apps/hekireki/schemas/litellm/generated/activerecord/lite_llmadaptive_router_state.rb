class LiteLLMAdaptiveRouterState < ApplicationRecord
  self.table_name = "LiteLLM_AdaptiveRouterState"
  self.primary_key = %w[router_name request_type model_name]

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

  attribute :total_samples, default: 0
  alias_attribute :model_name_value, :model_name
  alias_attribute :updated_at, :last_updated_at

  validates :alpha, presence: true
  validates :beta, presence: true
end
