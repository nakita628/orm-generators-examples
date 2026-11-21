# Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
# Maps source fields to Hub FeedbackRecord fields.
#
# @property id - Unique identifier for the mapping
# @property feedbackSource - The parent feedback source
# @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
# @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
# @property staticValue - If set, use this value instead of reading from sourceFieldId
class FeedbackSourceFieldMapping < ApplicationRecord
  self.table_name = "FeedbackSourceFieldMapping"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :source_field_id, presence: true,
                              uniqueness: { scope: [ :workspaceId, :feedback_source_id, :target_field_id ] }
  validates :target_field_id, presence: true

  belongs_to :feedback_source, foreign_key: [ :feedback_source_id, :workspaceId ], primary_key: [ :id, :workspaceId ],
                               inverse_of: :field_mappings
end
