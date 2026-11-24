class DomainEventOutbox < ApplicationRecord
  self.table_name = "domain_event_outbox"

  attribute :unpublished, default: true

  validates :base_id, presence: true
  validates :message_name, presence: true
  validates :schema_version, presence: true
  validates :payload, exclusion: { in: [ nil ] }
  validates :payload_bytes, presence: true
  validates :catalog_generation, presence: true
  validates :required_consumers, exclusion: { in: [ nil ] }
end
