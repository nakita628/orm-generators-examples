# Defines a segment of contacts based on attributes.
# Used for targeting surveys to specific user groups.
#
# @property id - Unique identifier for the segment
# @property title - Display name of the segment
# @property filters - Rules defining the segment
# @property isPrivate - Whether the segment is private
# @property workspace - The workspace this segment belongs to
#
# filters:: [SegmentFilter]
class Segment < ApplicationRecord
  self.table_name = "Segment"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :isPrivate, default: true
  attribute :filters, default: -> { JSON.parse("[]") }

  validates :title, presence: true, uniqueness: { scope: :workspaceId }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :segments
  has_many :surveys, foreign_key: "segmentId", inverse_of: :segment, dependent: :nullify
end
