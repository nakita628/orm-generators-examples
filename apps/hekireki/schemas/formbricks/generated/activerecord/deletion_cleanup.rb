# Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
# storage files that can't be removed inside the database transaction. A row is written in the same
# transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
# queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
#
# organizationId:: Scoping and logs. Not foreign keys: the rows they name are usually gone already.
# tenantIds:: Hub kinds: the organisation's feedback directories, read inside the delete's transaction.
# responseIds:: `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response.
# fileKeys:: `storageFiles`: the object keys or URLs to delete.
# nextAttemptAt:: Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup.
# lastError:: The last failure: a short code and message, never record or file content.
class DeletionCleanup < ApplicationRecord
  self.table_name = "DeletionCleanup"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid2.call }
  attribute :tenantIds, default: -> { [] }
  attribute :responseIds, default: -> { [] }
  attribute :fileKeys, default: -> { [] }
  attribute :attempts, default: 0

  enum :kind, { hub_survey: "hubSurvey", hub_responses: "hubResponses", storage_files: "storageFiles", storage_survey_folder: "storageSurveyFolder", hub_response_outcome: "hubResponseOutcome" },
       validate: true

  validates :organizationId, presence: true
  validates :workspaceId, presence: true
  validates :surveyId, presence: true
end
