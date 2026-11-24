class MastraTraces < ApplicationRecord
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

  alias_attribute :attributes_value, :attributes
  alias_attribute :created_at, :createdAt

  validates :name, presence: true
  validates :traceId, presence: true
  validates :scope, presence: true
  validates :kind, presence: true
  validates :startTime, presence: true
  validates :endTime, presence: true
end
