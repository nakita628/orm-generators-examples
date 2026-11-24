class LiteLLMSpendLogs < ApplicationRecord
  self.table_name = "LiteLLM_SpendLogs"
  self.primary_key = "request_id"

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

  attribute :api_key, default: ""
  attribute :spend, default: 0
  attribute :total_tokens, default: 0
  attribute :prompt_tokens, default: 0
  attribute :completion_tokens, default: 0
  attribute :model, default: ""
  attribute :model_id, default: ""
  attribute :model_group, default: ""
  attribute :custom_llm_provider, default: ""
  attribute :api_base, default: ""
  attribute :user, default: ""
  attribute :metadata, default: -> { JSON.parse("{}") }
  attribute :cache_hit, default: ""
  attribute :cache_key, default: ""
  attribute :request_tags, default: -> { JSON.parse("[]") }
  attribute :messages, default: -> { JSON.parse("{}") }
  attribute :response, default: -> { JSON.parse("{}") }
  attribute :proxy_server_request, default: -> { JSON.parse("{}") }
  alias_attribute :cache_key_value, :cache_key

  validates :call_type, presence: true
  validates :startTime, presence: true
  validates :endTime, presence: true
end
