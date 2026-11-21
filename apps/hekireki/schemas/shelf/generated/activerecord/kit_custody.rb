class KitCustody < ApplicationRecord
  self.table_name = "KitCustody"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :kitId, uniqueness: true

  belongs_to :custodian, class_name: "TeamMember", foreign_key: "custodianId", inverse_of: :kit_custodies
  belongs_to :kit, foreign_key: "kitId", inverse_of: :custody
  has_many :inherited_custody, class_name: "Custody", foreign_key: "kitCustodyId", inverse_of: :kit_custody,
                               dependent: :destroy
end
