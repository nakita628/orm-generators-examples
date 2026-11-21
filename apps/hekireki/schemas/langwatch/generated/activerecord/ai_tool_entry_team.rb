# @deprecated Per-team scope binding for AiToolEntry. Superseded by
# AiToolEntryDepartment; kept so the rename migration is non-destructive.
# No new rows are written here.
class AiToolEntryTeam < ApplicationRecord
  self.table_name = "AiToolEntryTeam"

  attribute :id, default: -> { Nanoid.generate }

  validates :entryId, uniqueness: { scope: :teamId }

  belongs_to :entry, class_name: "AiToolEntry", foreign_key: "entryId", inverse_of: :teams
  belongs_to :team, foreign_key: "teamId", inverse_of: :ai_tool_entry_teams
end
