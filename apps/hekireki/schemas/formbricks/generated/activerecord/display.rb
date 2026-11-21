# Records when a survey is shown to a user.
# Tracks survey display history and response status.
#
# @property id - Unique identifier for the display event
# @property survey - The survey that was displayed
# @property contact - The contact who saw the survey
# @property response - The associated response if one exists
# @property creativeId - The creative that opened the survey. No foreign key, so the record survives the creative's deletion.
class Display < ApplicationRecord
  self.table_name = "Display"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  belongs_to :survey, foreign_key: "surveyId", inverse_of: :displays
  belongs_to :contact, foreign_key: "contactId", optional: true, inverse_of: :displays
  has_one :response, foreign_key: "displayId", inverse_of: :display, dependent: :nullify
end
