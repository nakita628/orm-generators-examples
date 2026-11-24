class DomainEventInbox < ApplicationRecord
  self.table_name = "domain_event_inbox"
  self.primary_key = %w[consumer_id event_id]
end
