# One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
# while the run is in progress, or if it never finished.
#
# archivedCount:: Archived surveys, or deactivated members.
# deletedCount:: Deleted responses, or, on a purge run, deleted surveys.
# hasChanges:: False when the run changed nothing. History hides those by default.
# scanCursor:: Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
#   of rescanning the same first targets every night. Null when the scan reached the end.
class RetentionRun < ApplicationRecord
  self.table_name = "RetentionRun"

  attribute :id, default: -> { Cuid2.call }
  attribute :notifiedCount, default: 0
  attribute :archivedCount, default: 0
  attribute :deletedCount, default: 0
  attribute :skippedCount, default: 0
  attribute :hasChanges, default: false

  enum :entity, { responses: "responses", surveys: "surveys", members: "members" }, validate: true
  enum :kind, { sweep: "sweep", purge: "purge" }, default: :sweep, validate: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :retention_runs
  has_many :items, class_name: "RetentionRunItem", foreign_key: "runId", inverse_of: :run, dependent: :destroy
end
