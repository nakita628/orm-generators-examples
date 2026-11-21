# Historical record of every run a Session has owned. Append-only —
# rows are inserted on each `ensureRunForSession` claim, never updated.
# Lets us reconstruct the run timeline of a chat for debugging /
# dashboard surfaces. The relation cascades on Session delete (tied to
# the session lifecycle) but `runId` is a plain string column with no
# FK to TaskRun so run pruning is independent.
#
# runId:: TaskRun.id (no FK — runs may be archived independently of session history)
# reason:: One of: "initial" | "continuation" | "upgrade" | "manual".
#   Plain string for forward-compat with future trigger reasons.
class SessionRun < ApplicationRecord
  self.table_name = "SessionRun"

  attribute :id, default: -> { Cuid.generate }

  validates :runId, presence: true, uniqueness: true
  validates :reason, presence: true

  belongs_to :session, foreign_key: "sessionId", inverse_of: :runs
end
