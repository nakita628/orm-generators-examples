# Bulk actions, like canceling and replaying runs
#
# environment:: If this is set then it's a V2 Bulk Action that supports queries
# status:: When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed.
# queryName:: The query that will be executed to get the runs to process (if null, we are passing run ids directly)
# params:: The params that will be passed to the query
# dedupeKey:: Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions.
# cursor:: The cursor that will be passed to the query (null for the first page)
# successCount:: The number of runs that have been processed successfully
# failureCount:: The number of runs that have been failed
# totalCount:: The total number of runs that will be processed
# user:: The userId who did the bulk action
# name:: The name of the bulk action
# completedAt:: The time the bulk action was completed
class BulkActionGroup < ApplicationRecord
  self.table_name = "BulkActionGroup"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :successCount, default: 0
  attribute :failureCount, default: 0
  attribute :totalCount, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { cancel: "CANCEL", replay: "REPLAY" }, validate: true
  enum :status, { pending: "PENDING", completed: "COMPLETED", aborted: "ABORTED" }, default: :pending, validate: true
  enum :completionNotification, { none: "NONE", email: "EMAIL" }, default: :none, prefix: true, validate: true

  validates :friendlyId, presence: true, uniqueness: true

  belongs_to :project, foreign_key: "projectId", inverse_of: :bulk_action_groups
  belongs_to :environment, class_name: "RuntimeEnvironment", foreign_key: "environmentId", optional: true,
                           inverse_of: :bulk_action_group
  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :bulk_actions
end
