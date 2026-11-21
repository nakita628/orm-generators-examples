# This installation's VAPID key pair, generated on first use so Web Push works
# without configuration. Fixed id, so the table holds one row: a second
# insert loses the race on the primary key and reads the winner's pair.
#
# publicKey:: The uncompressed P-256 public key, base64url. The browser subscribes with it.
# privateKeyEncrypted:: The private key, base64url, encrypted with the process encryption key.
class WebPushVapidKey < ApplicationRecord
  self.table_name = "WebPushVapidKey"

  attribute :id, default: "self"
  alias_attribute :created_at, :createdAt

  validates :publicKey, presence: true
  validates :privateKeyEncrypted, presence: true
end
