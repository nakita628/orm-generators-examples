# A survey's response attachments built into one ZIP in object storage by a background job and
# downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
# Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
# survey deletes it with the survey's storage prefix.
#
# A partial unique index in the migration allows one queued or running export per requester and
# survey, so a double click returns the export already in flight.
#
# filterCriteria:: The response filter at request time, a `TResponseFilterCriteria`; null exports every response.
# timeZone:: IANA zone the archive's folder names are stamped in: the org's display zone at request time.
# dispatchCount:: Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
#   older pod dropped as unknown during a rolling upgrade.
# fileCount:: The pre-flight's counts, shown while the archive builds.
# includedFileCount:: What the archive actually holds, once built.
# errorCode:: A short code (`stale`, `storage_upload_failed`, …), never a message or file name.
class AttachmentExport < ApplicationRecord
  self.table_name = "AttachmentExport"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid2.call }
  attribute :dispatchCount, default: 0
  attribute :truncated, default: false

  enum :status, { queued: "queued", running: "running", completed: "completed", failed: "failed", expired: "expired" },
       default: :queued, validate: true

  validates :timeZone, presence: true
  validates :fileCount, presence: true
  validates :responseCount, presence: true

  belongs_to :survey, foreign_key: [ :surveyId, :workspaceId ], primary_key: [ :id, :workspaceId ],
                      inverse_of: :attachment_exports
  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :attachment_exports
  belongs_to :requested_by, class_name: "User", foreign_key: "requestedById", optional: true,
                            inverse_of: :attachment_exports
end
