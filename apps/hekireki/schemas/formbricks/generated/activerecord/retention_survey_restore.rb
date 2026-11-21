# When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
# and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
# so retention adds nothing to the survey row.
class RetentionSurveyRestore < ApplicationRecord
  self.table_name = "RetentionSurveyRestore"
  self.primary_key = "surveyId"

  validates :restoredAt, presence: true

  belongs_to :survey, foreign_key: "surveyId", inverse_of: :retention_restore
end
