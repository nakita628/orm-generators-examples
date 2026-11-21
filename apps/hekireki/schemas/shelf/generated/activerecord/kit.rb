class Kit < ApplicationRecord
  self.table_name = "Kit"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { available: "AVAILABLE", in_custody: "IN_CUSTODY", checked_out: "CHECKED_OUT" }, default: :available,
                                                                                                  validate: true

  validates :name, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :kits
  belongs_to :created_by, class_name: "User", foreign_key: "createdById", inverse_of: :created_kits
  belongs_to :category, foreign_key: "categoryId", optional: true, inverse_of: :kits
  belongs_to :location, foreign_key: "locationId", optional: true, inverse_of: :kits
  has_one :custody, class_name: "KitCustody", foreign_key: "kitId", inverse_of: :kit, dependent: :destroy
  has_many :asset_kits, foreign_key: "kitId", inverse_of: :kit, dependent: :destroy
  has_many :qr_codes, class_name: "Qr", foreign_key: "kitId", inverse_of: :kit, dependent: :nullify
  has_many :barcodes, foreign_key: "kitId", inverse_of: :kit, dependent: :nullify
  has_many :reports, class_name: "ReportFound", foreign_key: "kitId", inverse_of: :kit, dependent: :destroy
  has_many :scans, foreign_key: "kitId", inverse_of: :kit, dependent: :nullify
end
