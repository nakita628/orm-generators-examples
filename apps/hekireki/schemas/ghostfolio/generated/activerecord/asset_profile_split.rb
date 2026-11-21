# A stock split of an asset profile. The numerator is the number of shares
# held after the split, the denominator the number of shares held before, for
# example 4 and 1 for a 4:1 split or 1 and 10 for a 1:10 reverse split.
#
# The ratio is stored as two integers instead of the resulting factor, so that
# it stays exact, for example 1/3 for a 1:3 reverse split.
#
# Activities before the date are adjusted by the ratio: the quantity is
# multiplied by it and the unit price divided by it, so that the total value
# of an activity stays the same.
#
# Market data is already split-adjusted by the data providers and must not be
# adjusted again.
class AssetProfileSplit < ApplicationRecord
  self.table_name = "AssetProfileSplit"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { SecureRandom.uuid }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :date, presence: true, uniqueness: { scope: :symbolProfileId }
  validates :denominator, presence: true
  validates :numerator, presence: true

  belongs_to :symbol_profile, foreign_key: "symbolProfileId", inverse_of: :asset_profile_splits
end
