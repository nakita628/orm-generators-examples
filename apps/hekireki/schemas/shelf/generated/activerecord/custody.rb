class Custody < ApplicationRecord
  self.table_name = "Custody"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :quantity, default: 1
  attribute :sourceUnknown, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  belongs_to :custodian, class_name: "TeamMember", foreign_key: "teamMemberId", inverse_of: :custodies
  belongs_to :asset, foreign_key: "assetId", inverse_of: :custody
  belongs_to :kit_custody, foreign_key: "kitCustodyId", optional: true, inverse_of: :inherited_custody
  belongs_to :location, foreign_key: "locationId", optional: true, inverse_of: :custodies
end
