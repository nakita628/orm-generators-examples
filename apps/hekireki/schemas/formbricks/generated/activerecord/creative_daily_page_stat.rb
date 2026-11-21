# A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
# `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
#
# date:: UTC day.
# page:: Host + path, without query string or hash, or `(other)`.
class CreativeDailyPageStat < ApplicationRecord
  self.table_name = "CreativeDailyPageStat"
  self.primary_key = %w[creativeId date page]

  attribute :impressions, default: 0
  attribute :clicks, default: 0

  belongs_to :creative, foreign_key: "creativeId", inverse_of: :daily_page_stats
end
