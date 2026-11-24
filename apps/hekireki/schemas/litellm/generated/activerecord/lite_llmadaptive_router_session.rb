class LiteLLMAdaptiveRouterSession < ApplicationRecord
  self.table_name = "LiteLLM_AdaptiveRouterSession"
  self.primary_key = %w[session_id router_name model_name]

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

  attribute :misalignment_count, default: 0
  attribute :stagnation_count, default: 0
  attribute :disengagement_count, default: 0
  attribute :satisfaction_count, default: 0
  attribute :failure_count, default: 0
  attribute :loop_count, default: 0
  attribute :exhaustion_count, default: 0
  attribute :tool_call_history, default: -> { JSON.parse("[]") }
  attribute :pending_tool_calls, default: -> { JSON.parse("{}") }
  attribute :turn_count, default: 0
  attribute :last_processed_turn, default: -1
  attribute :clean_credit_awarded, default: false
  alias_attribute :model_name_value, :model_name
  alias_attribute :updated_at, :last_activity_at

  validates :classified_type, presence: true
end
