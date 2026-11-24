class Trending < ApplicationRecord
  self.table_name = "Trending"
  self.implicit_order_column = "createdAt"

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

  attribute :id, default: -> { SecureRandom.uuid }
  alias_attribute :hash_value, :hash
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :trendingList, presence: true
  validates :language, uniqueness: { allow_nil: true }
  validates :hash_value, presence: true
  validates :date, presence: true
end
