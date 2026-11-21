# Links surveys to their supported languages.
# Manages which languages are available for each survey.
#
# @property language - The supported language
# @property survey - The associated survey
# @property default - Whether this is the default language
# @property enabled - Whether this language is currently active
class SurveyLanguage < ApplicationRecord
  self.table_name = "SurveyLanguage"
  self.primary_key = %w[languageId surveyId]

  attribute :default, default: false
  attribute :enabled, default: true

  belongs_to :language, foreign_key: "languageId", inverse_of: :survey_languages
  belongs_to :survey, foreign_key: "surveyId", inverse_of: :languages
end
