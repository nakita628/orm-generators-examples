class Image < ApplicationRecord
  self.table_name = "Image"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :contentType, presence: true
  validates :blob, exclusion: { in: [ nil ] }

  belongs_to :owner_org, class_name: "Organization", foreign_key: "ownerOrgId", inverse_of: :images
  belongs_to :user, foreign_key: "userId", inverse_of: :images
  has_one :location, foreign_key: "imageId", inverse_of: :image, dependent: :nullify
  has_one :organization, foreign_key: "imageId", inverse_of: :image, dependent: :nullify
end
