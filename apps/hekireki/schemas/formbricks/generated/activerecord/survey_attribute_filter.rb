# Defines targeting rules for surveys based on contact attributes.
# Used to show surveys only to contacts matching specific criteria.
#
# @property id - Unique identifier for the filter
# @property attributeKey - The contact attribute to filter on
# @property survey - The survey being filtered
# @property condition - The comparison operator to use
# @property value - The value to compare against
class SurveyAttributeFilter < ApplicationRecord
  self.table_name = "SurveyAttributeFilter"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  enum :condition, { equals: "equals", not_equals: "notEquals" }, validate: true

  validates :surveyId, uniqueness: { scope: :attributeKeyId }
  validates :value, presence: true

  belongs_to :attribute_key, class_name: "ContactAttributeKey", foreign_key: "attributeKeyId",
                             inverse_of: :attribute_filters
  belongs_to :survey, foreign_key: "surveyId", inverse_of: :attribute_filters
end
