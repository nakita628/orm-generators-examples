class FieldPeriodicRefreshState < ApplicationRecord
  self.table_name = "field_periodic_refresh_state"
  self.primary_key = "field_id"

  attribute :enabled, default: true
  attribute :time_zone, default: "UTC"
  attribute :contains_now_precision, default: false
  attribute :last_status, default: "idle"
  alias_attribute :updated_at, :last_modified_time

  validates :table_id, presence: true
  validates :base_id, presence: true
  validates :time_zone, length: { maximum: 64 }
  validates :last_refreshed_local_date, length: { maximum: 16 }
  validates :last_range_hash, length: { maximum: 128 }
  validates :last_status, length: { maximum: 32 }
  validates :pending_task_id, length: { maximum: 64 }
  validates :pending_local_date, length: { maximum: 16 }
  validates :pending_range_hash, length: { maximum: 128 }

  belongs_to :field, inverse_of: :periodic_refresh_state
end
