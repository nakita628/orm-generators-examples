# A short code a fresh install pastes instead of a license blob (ADR-139,
# section 5). Redeeming it mints the license the code describes and returns
# it signed, which is what identifies a new customer without anybody emailing
# a license key around.
#
# The row holds a hash of the code, never the code. Every lookup hashes what
# the caller presented first, so the stored value is a step away from anything
# that could be presented: read access to this table is not credential access.
#
# codeHash:: SHA-256 of the normalised code. Looked up on every redemption.
# codeHint:: The last four characters, so an operator tells two codes apart in a list
#   without the list holding either one.
# organizationId:: The customer this code belongs to, and the terms it mints.
# maxMembersLite:: The lite seats the minted license carries: null when unlimited. Never a stand-in number.
# liteSeatsUnlimited:: True when the minted license carries unlimited lite seats; maxMembersLite is then null.
# maxMembersDeveloper:: The Developer seats the minted license carries: null when not stated, which sells none.
# licenseTermDays:: How long the license this code mints runs for.
# services:: Hosted services the minted license may call.
# expiresAt:: When the code stops working, which is not the term of what it mints.
# reusable:: A code a customer may redeem on more than one install, for a rollout.
# redeemedAt:: Set by the redemption that won, which is what makes a single-use code
#   single use: the write itself carries the condition.
# issuedLicenseId:: The license the winning redemption minted.
# redemptionCount:: How many installs have redeemed it, which only moves on a reusable code.
class ActivationCode < ApplicationRecord
  self.table_name = "ActivationCode"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :maxMembersLite, default: 0
  attribute :liteSeatsUnlimited, default: false
  attribute :services, default: -> { [] }
  attribute :reusable, default: false
  attribute :redemptionCount, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :codeHash, presence: true, uniqueness: true
  validates :codeHint, presence: true
  validates :organizationId, presence: true
  validates :organizationName, presence: true
  validates :email, presence: true
  validates :planType, presence: true
  validates :maxMembers, presence: true
  validates :licenseTermDays, presence: true
  validates :expiresAt, presence: true
end
