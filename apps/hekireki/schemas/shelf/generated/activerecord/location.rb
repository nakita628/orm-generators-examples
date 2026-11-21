class Location < ApplicationRecord
  self.table_name = "Location"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true
  validates :imageId, uniqueness: { allow_nil: true }

  belongs_to :image, foreign_key: "imageId", optional: true, inverse_of: :location
  belongs_to :user, foreign_key: "userId", inverse_of: :locations
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :locations
  belongs_to :parent, class_name: "Location", foreign_key: "parentId", optional: true, inverse_of: :children
  has_many :asset_locations, foreign_key: "locationId", inverse_of: :location, dependent: :destroy
  has_many :kits, foreign_key: "locationId", inverse_of: :location, dependent: :nullify
  has_many :notes, class_name: "LocationNote", foreign_key: "locationId", inverse_of: :location, dependent: :destroy
  has_many :custodies, foreign_key: "locationId", inverse_of: :location, dependent: :nullify
  has_many :consumption_logs, foreign_key: "locationId", inverse_of: :location, dependent: :nullify
  has_many :children, class_name: "Location", foreign_key: "parentId", inverse_of: :parent, dependent: :nullify
end
