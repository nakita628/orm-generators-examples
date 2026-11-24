class DomainEventDelivery < ApplicationRecord
  self.table_name = "domain_event_delivery"

  attribute :attempts, default: 0
  attribute :max_attempts, default: 12

  validates :event_id, presence: true, uniqueness: { scope: :consumer_id }
  validates :consumer_id, presence: true
  validates :status, presence: true
end
