# Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
# User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
# at rest by Better Auth using BETTER_AUTH_SECRET.
#
# failedVerificationCount:: Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts.
# lockedUntil:: Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit.
class TwoFactor < ApplicationRecord
  self.table_name = "TwoFactor"

  attribute :id, default: -> { Cuid.generate }
  attribute :verified, default: true
  attribute :failedVerificationCount, default: 0

  validates :secret, presence: true
  validates :backupCodes, presence: true
  validates :userId, uniqueness: true
  validates :userId, uniqueness: true

  belongs_to :user, foreign_key: "userId", inverse_of: :two_factor
end
