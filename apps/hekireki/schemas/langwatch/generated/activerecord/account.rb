# issuer:: WHO asserted `providerAccountId`. Better Auth keys accounts by
#   `(provider, providerAccountId)`; auth's account create hook writes this
#   (modules/auth/specs/account-issuer.feature). Nullable only so the column
#   could be added to a populated table; the migration backfilled every row.
class Account < ApplicationRecord
  self.table_name = "Account"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :type, default: "oauth"
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :provider, presence: true, uniqueness: { scope: :providerAccountId }
  validates :providerAccountId, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :accounts
end
