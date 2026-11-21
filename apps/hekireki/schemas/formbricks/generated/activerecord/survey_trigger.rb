# Links surveys to specific trigger actions.
# Defines which user actions should cause a survey to be displayed.
# This is the connection table between Surveys and ActionClasses that determines
# when and under what conditions a survey should be triggered.
#
# @property id - Unique identifier for the trigger
# @property survey - The survey to be triggered
# @property actionClass - The action that triggers the survey
# @property createdAt - When the trigger was created
# @property updatedAt - When the trigger was last modified
class SurveyTrigger < ApplicationRecord
  self.table_name = "SurveyTrigger"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :surveyId, uniqueness: { scope: :actionClassId }

  belongs_to :survey, foreign_key: "surveyId", inverse_of: :triggers
  belongs_to :action_class, foreign_key: "actionClassId", inverse_of: :survey_triggers
end
