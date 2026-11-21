# What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
# never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
# `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
#
# count:: 1, except for grouped response deletions, where it is the number of responses deleted.
# recipient:: Who was emailed, for `notified` rows.
class RetentionRunItem < ApplicationRecord
  self.table_name = "RetentionRunItem"

  attribute :id, default: -> { Cuid2.call }
  attribute :count, default: 1

  enum :targetType, { survey: "survey", user: "user" }, validate: true
  enum :action, { notified: "notified", archived: "archived", deactivated: "deactivated", deleted: "deleted", skipped: "skipped" },
       validate: true
  enum :skipReason, { exempt: "exempt", last_owner: "lastOwner", other_organization: "otherOrganization", no_recipient: "noRecipient" },
       validate: { allow_nil: true }

  validates :targetId, presence: true

  belongs_to :run, class_name: "RetentionRun", foreign_key: "runId", inverse_of: :items
end
