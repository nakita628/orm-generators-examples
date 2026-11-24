class MastraSkillBlobs < ApplicationRecord
  self.primary_key = "hash"

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

  alias_attribute :hash_value, :hash
  alias_attribute :created_at, :createdAt

  validates :content, presence: true
  validates :size, presence: true
end
