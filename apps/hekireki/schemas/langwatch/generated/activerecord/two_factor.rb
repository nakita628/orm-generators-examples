# The two-factor plugin's own table (D06). Protocol state, and row-truth for
# good (ADR-101 R12): the shared secret and the backup codes live here and
# in no event, ever. Nothing in this repo reads or writes these columns -
# the plugin owns them end to end, including the lockout counters, which is
# why `failedVerificationCount` and `lockedUntil` are here rather than
# rebuilt on the projection.
#
# `backupCodes` holds an ENCRYPTED blob, not hashes. A backup code has to be
# compared against what somebody types, and the plugin decrypts and compares;
# hashing them would leave it unable to verify its own codes.
class TwoFactor < ApplicationRecord
  self.table_name = "TwoFactor"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :verified, default: true
  attribute :failedVerificationCount, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :secret, presence: true
  validates :backupCodes, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :two_factors
end
