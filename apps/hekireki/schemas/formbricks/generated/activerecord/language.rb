# Represents a supported language in the system.
# Used for multilingual survey support.
#
# @property id - Unique identifier for the language
# @property code - Language code (e.g., 'en-US')
# @property alias - Optional friendly name
# @property workspace - The workspace this language is enabled for
class Language < ApplicationRecord
  self.table_name = "Language"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :code, presence: true, uniqueness: { scope: :workspaceId }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :languages
  has_many :survey_languages, foreign_key: "languageId", inverse_of: :language, dependent: :destroy
  has_many :creative_languages, foreign_key: "languageId", inverse_of: :language, dependent: :destroy
end
