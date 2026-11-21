# Public half of the user's capture key. The row outlives switching the
# feature off: `disabledAt` stops a stale device re-registering that key.
#
# public_key:: Base64 SPKI DER of an ECDH P-256 public key.
class CaptureKey < ApplicationRecord
  self.primary_key = "user_id"

  validates :public_key, presence: true
  validates :fingerprint, presence: true
  validates :created_at, presence: true

  belongs_to :user
end
