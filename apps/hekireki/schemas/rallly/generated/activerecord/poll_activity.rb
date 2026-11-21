# type:: Validated by a zod union in the feature, not a DB enum: the activity vocabulary grows without migrations
# user_id:: Subject ids are soft references (no FKs): activity must outlive its subjects — including account deletion — so refs persist as identity and the payload carries the display snapshot; whether a ref still resolves is the renderer's question
class PollActivity < ApplicationRecord
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :payload, default: -> { JSON.parse("{}") }

  validates :type, presence: true

  belongs_to :poll, inverse_of: :activities
end
