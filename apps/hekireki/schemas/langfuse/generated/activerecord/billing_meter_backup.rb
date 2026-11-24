class BillingMeterBackup < ApplicationRecord
  self.primary_key = %w[stripe_customer_id meter_id start_time end_time]

  validates :stripe_customer_id, presence: true, uniqueness: { scope: %i[meter_id start_time end_time] }
  validates :meter_id, presence: true
  validates :start_time, presence: true
  validates :end_time, presence: true
  validates :aggregated_value, presence: true
  validates :event_name, presence: true
  validates :org_id, presence: true
end
