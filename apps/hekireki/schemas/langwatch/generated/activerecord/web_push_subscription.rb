# One browser a person subscribed to Web Push, owned by notification. A
# person holds one row per device; the endpoint names the browser, so a
# browser that subscribes again replaces its own row.
#
# userId:: The person the pushes are for. A plain column: the row leaves with the
#   person through notification's own subscriber on the user lifecycle.
# endpoint:: The push service URL the browser handed out. Unique per browser.
# p256dh:: The browser's P-256 public key and auth secret, base64url, as the
#   browser reports them. Payloads are encrypted to these.
# lastSuccessAt:: When the push service last accepted a push for this browser.
class WebPushSubscription < ApplicationRecord
  self.table_name = "WebPushSubscription"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt

  validates :userId, presence: true
  validates :endpoint, presence: true, uniqueness: true
  validates :p256dh, presence: true
  validates :auth, presence: true
end
