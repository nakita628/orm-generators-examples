# A person's opt-out of the digest. No row, or no `unsubscribedAt`: subscribed.
#
# unsubscribedFromSendId:: The send whose unsubscribe link or one-click header the person used.
class DigestSubscription < ApplicationRecord
  self.table_name = "DigestSubscription"
  self.primary_key = "userId"

  alias_attribute :updated_at, :updatedAt
end
