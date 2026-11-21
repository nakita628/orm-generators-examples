# Represents an attribute value associated with a contact.
# Used to store custom properties and characteristics of contacts.
#
# @property id - Unique identifier for the attribute
# @property attributeKey - Reference to the attribute definition
# @property contact - The contact this attribute belongs to
# @property value - The string value of the attribute (used for string type + backwards compatibility)
# @property valueNumber - Native numeric storage for number type attributes
# @property valueDate - Native date storage for date type attributes
class ContactAttribute < ApplicationRecord
  self.table_name = "ContactAttribute"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :contactId, uniqueness: { scope: :attributeKeyId }
  validates :value, presence: true

  belongs_to :attribute_key, class_name: "ContactAttributeKey", foreign_key: "attributeKeyId",
                             inverse_of: :contact_attributes
  belongs_to :contact, foreign_key: "contactId", inverse_of: :contact_attributes
end
