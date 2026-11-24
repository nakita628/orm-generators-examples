class Account < ApplicationRecord
  self.table_name = "Account"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  def self.timestamp_attributes_for_create
    super - %w[created_at] + %w[createdAt]
  end
  private_class_method :timestamp_attributes_for_create

  validates :type, presence: true
  validates :provider, presence: true, uniqueness: { scope: :providerAccountId }
  validates :providerAccountId, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :accounts
end
