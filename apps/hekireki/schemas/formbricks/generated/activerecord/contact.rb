# Represents a person or user who can receive and respond to surveys.
# Contacts are workspace-specific and can have multiple attributes and responses.
#
# @property id - Unique identifier for the contact
# @property workspace - The workspace this contact belongs to
# @property responses - Survey responses from this contact
# @property attributes - Custom attributes associated with this contact
# @property displays - Record of surveys shown to this contact
class Contact < ApplicationRecord
  self.table_name = "Contact"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :contacts
  has_many :responses, foreign_key: "contactId", inverse_of: :contact, dependent: :destroy
  has_many :contact_attributes, foreign_key: "contactId", inverse_of: :contact, dependent: :destroy
  has_many :displays, foreign_key: "contactId", inverse_of: :contact, dependent: :destroy
end
