class Category < ApplicationRecord
  self.table_name = "Category"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true
  validates :color, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :categories
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :categories
  has_many :assets, foreign_key: "categoryId", inverse_of: :category, dependent: :nullify
  has_many :kits, foreign_key: "categoryId", inverse_of: :category, dependent: :nullify
  has_many :asset_model_defaults, class_name: "AssetModel", foreign_key: "defaultCategoryId",
                                  inverse_of: :default_category, dependent: :nullify
  has_and_belongs_to_many :custom_fields, join_table: "_CategoryToCustomField", foreign_key: "A",
                                          association_foreign_key: "B"
end
