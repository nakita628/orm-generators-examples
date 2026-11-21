# The monthly usage statement one connected customer was sent (ADR-139,
# section 7). One row per account per month keeps the statement to one send,
# whatever runs the tick.
#
# accountId:: The `ConnectedBillingAccount` row. A plain column ref.
# month:: First day of the month the statement covers, in UTC.
class ConnectedStatement < ApplicationRecord
  self.table_name = "ConnectedStatement"

  attribute :id, default: -> { Nanoid.generate }

  validates :accountId, presence: true, uniqueness: { scope: :month }
  validates :month, presence: true
  validates :sentAt, presence: true
end
