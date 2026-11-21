# format:: How a TEXT field's values are read. LINK requires a web link in every
#   value; any other type stays PLAIN (enforced by a CHECK constraint).
class CustomField < ApplicationRecord
  self.table_name = "CustomField"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :required, default: false
  attribute :active, default: true
  attribute :options, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { text: "TEXT", option: "OPTION", boolean: "BOOLEAN", date: "DATE", multiline_text: "MULTILINE_TEXT", amount: "AMOUNT", number: "NUMBER" },
       default: :text, validate: true
  enum :format, { plain: "PLAIN", link: "LINK" }, default: :plain, validate: true

  validates :name, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :custom_fields
  belongs_to :created_by, class_name: "User", foreign_key: "userId", inverse_of: :custom_fields
  has_many :asset_custom_fields_values, class_name: "AssetCustomFieldValue", foreign_key: "customFieldId",
                                        inverse_of: :custom_field, dependent: :destroy
  has_and_belongs_to_many :categories, join_table: "_CategoryToCustomField", foreign_key: "B",
                                       association_foreign_key: "A"
end
