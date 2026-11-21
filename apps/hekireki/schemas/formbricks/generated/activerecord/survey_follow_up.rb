# Defines follow-up actions for survey responses.
# Enables automated actions based on specific survey response conditions.
#
# @property id - Unique identifier for the follow-up
# @property survey - The associated survey
# @property name - Display name for the follow-up
# @property trigger - Conditions that activate the follow-up
# @property action - Actions to take when triggered
#
# trigger:: [SurveyFollowUpTrigger]
# action:: [SurveyFollowUpAction]
class SurveyFollowUp < ApplicationRecord
  self.table_name = "SurveyFollowUp"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :name, presence: true
  validates :trigger, exclusion: { in: [ nil ] }
  validates :action, exclusion: { in: [ nil ] }

  belongs_to :survey, foreign_key: "surveyId", inverse_of: :follow_ups
end
