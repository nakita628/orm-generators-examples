class UserPassword < ApplicationRecord
  self.table_name = "UserPassword"
  self.primary_key = "userId"

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

  validates :hash_value, presence: true
  validates :userId, uniqueness: true

  belongs_to :user, foreign_key: "userId", inverse_of: :password
end
