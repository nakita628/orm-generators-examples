class Model < ApplicationRecord
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

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :model_name_value, :model_name

  enum :scope, { langfuse: "LANGFUSE", organization: "ORGANIZATION", project: "PROJECT" }, validate: true

  validates :model_name_value, presence: true,
                               uniqueness: { scope: %i[project_id start_date unit], conditions: -> { where.not(project_id: nil).where.not(start_date: nil).where.not(unit: nil) } }
  validates :match_pattern, presence: true

  belongs_to :project, optional: true, inverse_of: :model
  belongs_to :organization, optional: true, inverse_of: :models
  has_many :price, dependent: :destroy
  has_many :pricing_tiers, dependent: :destroy
end
