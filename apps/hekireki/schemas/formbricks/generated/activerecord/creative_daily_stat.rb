# A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
# only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
# running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
# the creative deletes its counts.
#
# date:: UTC day.
class CreativeDailyStat < ApplicationRecord
  self.table_name = "CreativeDailyStat"
  self.primary_key = %w[creativeId date deviceType]

  attribute :impressions, default: 0
  attribute :clicks, default: 0

  enum :deviceType, { phone: "phone", tablet: "tablet", desktop: "desktop" }, validate: true

  belongs_to :creative, foreign_key: "creativeId", inverse_of: :daily_stats
end
