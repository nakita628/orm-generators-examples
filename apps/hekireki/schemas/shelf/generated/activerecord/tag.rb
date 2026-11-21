class Tag < ApplicationRecord
  self.table_name = "Tag"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :tags
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :tags
  has_and_belongs_to_many :assets, join_table: "_AssetToTag", foreign_key: "B", association_foreign_key: "A"
  has_and_belongs_to_many :bookings, join_table: "_BookingToTag", foreign_key: "B", association_foreign_key: "A"
end
