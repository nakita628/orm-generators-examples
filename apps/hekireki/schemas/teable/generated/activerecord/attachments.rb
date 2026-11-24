class Attachments < ApplicationRecord
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
  alias_attribute :hash_value, :hash

  validates :token, presence: true, uniqueness: true
  validates :hash_value, presence: true
  validates :size, presence: true
  validates :mimetype, presence: true
  validates :path, presence: true
  validates :created_by, presence: true
end
