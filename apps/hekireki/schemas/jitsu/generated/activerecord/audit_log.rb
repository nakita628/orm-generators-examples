# timestamp:: @zod.custom(z.coerce.date())
class AuditLog < ApplicationRecord
  self.table_name = "AuditLog"
  self.inheritance_column = nil

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
  attribute :severity, default: "info"
  alias_attribute :changes_value, :changes

  validates :type, presence: true
end
