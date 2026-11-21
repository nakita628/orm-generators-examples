# Represents third-party service integrations.
# Enables data flow between Formbricks and external services.
#
# @property id - Unique identifier for the integration
# @property type - The service being integrated (Google Sheets, Notion, etc.)
# @property workspace - The workspace this integration belongs to
# @property config - Service-specific configuration details
#
# config:: [IntegrationConfig]
class Integration < ApplicationRecord
  self.table_name = "Integration"
  self.inheritance_column = nil

  attribute :id, default: -> { Cuid.generate }

  enum :type, { google_sheets: "googleSheets", notion: "notion", airtable: "airtable", slack: "slack" }, validate: true

  validates :config, exclusion: { in: [ nil ] }
  validates :workspaceId, uniqueness: { scope: :type }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :integrations
end
