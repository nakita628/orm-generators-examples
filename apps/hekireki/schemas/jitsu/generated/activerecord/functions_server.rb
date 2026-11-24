class FunctionsServer < ApplicationRecord
  self.table_name = "FunctionsServer"
  self.primary_key = %w[workspaceId class]

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

  attribute :connections, default: -> { [] }
  attribute :emptyConnections, default: -> { [] }
  attribute :profileBuilders, default: -> { [] }
  alias_attribute :class_value, :class
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :deploymentId, presence: true
end
