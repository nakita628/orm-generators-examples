class LiteLLMHealthCheckTable < ApplicationRecord
  self.table_name = "LiteLLM_HealthCheckTable"
  self.primary_key = "health_check_id"
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

  attribute :health_check_id, default: -> { SecureRandom.uuid }
  attribute :healthy_count, default: 0
  attribute :unhealthy_count, default: 0
  alias_attribute :model_name_value, :model_name

  validates :model_name_value, presence: true
  validates :status, presence: true
end
