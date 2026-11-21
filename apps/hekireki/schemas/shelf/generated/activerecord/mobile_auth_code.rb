# Native-app SSO web-delegated login — single-use session-handoff codes.
# @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
class MobileAuthCode < ApplicationRecord
  self.table_name = "MobileAuthCode"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  validates :codeHash, presence: true, uniqueness: true
  validates :expiresAt, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :mobile_auth_codes
end
