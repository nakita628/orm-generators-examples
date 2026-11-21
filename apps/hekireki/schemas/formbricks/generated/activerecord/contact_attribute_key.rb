# Defines the possible attributes that can be assigned to contacts.
# Acts as a schema for contact attributes within a workspace.
#
# @property id - Unique identifier for the attribute key
# @property isUnique - Whether the attribute must have unique values across contacts
# @property key - The attribute identifier used in the system
# @property name - Display name for the attribute
# @property type - Whether this is a default or custom attribute
# @property dataType - The data type of the attribute (string, number, date)
# @property workspace - The workspace this attribute belongs to
class ContactAttributeKey < ApplicationRecord
  self.table_name = "ContactAttributeKey"
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :isUnique, default: false

  enum :type, { default: "default", custom: "custom" }, default: :custom, validate: true
  enum :dataType, { string: "string", number: "number", date: "date" }, default: :string, validate: true

  validates :key, presence: true, uniqueness: { scope: :workspaceId }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :contact_attribute_keys
  has_many :contact_attributes, foreign_key: "attributeKeyId", inverse_of: :attribute_key, dependent: :destroy
  has_many :attribute_filters, class_name: "SurveyAttributeFilter", foreign_key: "attributeKeyId",
                               inverse_of: :attribute_key, dependent: :destroy
end
