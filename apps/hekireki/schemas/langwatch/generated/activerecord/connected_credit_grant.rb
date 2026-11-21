# A paid credit at the payment provider, which only metered usage draws down.
# Kept here so a grant is created once per reason and the backoffice can show
# what the customer holds without asking the provider.
#
# accountId:: The `ConnectedBillingAccount` row. A plain column ref, like
#   `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
# stripeCreditGrantId:: The grant at the payment provider. Unique, so a retry records one row.
# termEndsAt:: The term this credit was granted for, and when the provider lets it lapse:
#   after the last invoice of that term has been finalized, never on the day
#   the term ends.
class ConnectedCreditGrant < ApplicationRecord
  self.table_name = "ConnectedCreditGrant"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt

  enum :kind, { commit: "commit", added: "added", renewal: "renewal" }, validate: true

  validates :accountId, presence: true
  validates :stripeCreditGrantId, presence: true, uniqueness: true
  validates :amountUsdCents, presence: true
  validates :termEndsAt, presence: true
  validates :expiresAt, presence: true
end
