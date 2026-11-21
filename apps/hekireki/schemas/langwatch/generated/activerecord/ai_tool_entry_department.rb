# Per-department scope binding for AiToolEntry. Empty set = org-wide.
# Non-empty = entry visible only to members of those departments (an
# OrganizationUser whose `departmentId` is in the set). Mirrors
# AiToolEntryTeam's shape. `departmentId` is a plain column ref
# (relationMode="prisma"), matching how Department is referenced
# elsewhere; the service validates the id belongs to the org.
class AiToolEntryDepartment < ApplicationRecord
  self.table_name = "AiToolEntryDepartment"

  attribute :id, default: -> { Nanoid.generate }

  validates :departmentId, presence: true, uniqueness: { scope: :entryId }

  belongs_to :entry, class_name: "AiToolEntry", foreign_key: "entryId", inverse_of: :departments
end
