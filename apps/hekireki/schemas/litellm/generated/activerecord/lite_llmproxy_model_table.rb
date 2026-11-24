class LiteLLMProxyModelTable < ApplicationRecord
  self.table_name = "LiteLLM_ProxyModelTable"
  self.primary_key = "model_id"
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

  attribute :model_id, default: -> { SecureRandom.uuid }
  attribute :blocked, default: false
  alias_attribute :model_name_value, :model_name

  validates :model_name_value, presence: true
  validates :litellm_params, exclusion: { in: [ nil ] }
  validates :created_by, presence: true
  validates :updated_by, presence: true
end
