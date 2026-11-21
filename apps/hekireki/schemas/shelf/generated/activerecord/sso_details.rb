class SsoDetails < ApplicationRecord
  self.table_name = "SsoDetails"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :requireSsoLogin, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :domain, presence: true

  has_many :organizations, foreign_key: "ssoDetailsId", inverse_of: :sso_details, dependent: :nullify
end
