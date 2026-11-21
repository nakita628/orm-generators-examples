class Asset < ApplicationRecord
  self.table_name = "Asset"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :availableToBook, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { available: "AVAILABLE", in_custody: "IN_CUSTODY", checked_out: "CHECKED_OUT" }, default: :available,
                                                                                                  validate: true
  enum :type, { individual: "INDIVIDUAL", quantity_tracked: "QUANTITY_TRACKED" }, default: :individual, validate: true
  enum :consumptionType, { one_way: "ONE_WAY", two_way: "TWO_WAY" }, validate: { allow_nil: true }

  validates :title, presence: true
  validates :sequentialId, uniqueness: { scope: :organizationId, allow_nil: true }
  validates :preferredBarcodeId, uniqueness: { allow_nil: true }

  belongs_to :asset_model, foreign_key: "assetModelId", optional: true, inverse_of: :assets
  belongs_to :user, foreign_key: "userId", inverse_of: :assets
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :assets
  belongs_to :category, foreign_key: "categoryId", optional: true, inverse_of: :assets
  belongs_to :preferred_barcode, class_name: "Barcode", foreign_key: "preferredBarcodeId", optional: true,
                                 inverse_of: :preferred_by_asset
  has_many :custody, foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_many :notes, foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_many :qr_codes, class_name: "Qr", foreign_key: "assetId", inverse_of: :asset, dependent: :nullify
  has_many :barcodes, foreign_key: "assetId", inverse_of: :asset, dependent: :nullify
  has_many :reports, class_name: "ReportFound", foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_many :custom_fields, class_name: "AssetCustomFieldValue", foreign_key: "assetId", inverse_of: :asset,
                           dependent: :destroy
  has_many :booking_assets, foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_many :consumption_logs, foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_many :reminders, class_name: "AssetReminder", foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_many :audit_entries, class_name: "AuditAsset", foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_many :audit_scans, foreign_key: "assetId", inverse_of: :asset, dependent: :nullify
  has_many :scans, foreign_key: "assetId", inverse_of: :asset, dependent: :nullify
  has_many :asset_kits, foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_many :asset_locations, foreign_key: "assetId", inverse_of: :asset, dependent: :destroy
  has_and_belongs_to_many :tags, join_table: "_AssetToTag", foreign_key: "A", association_foreign_key: "B"
end
