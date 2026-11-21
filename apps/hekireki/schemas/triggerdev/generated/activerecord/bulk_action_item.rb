# friendlyId:: @deprecated not used in new BulkActions
# status:: When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed.
# source_run:: The run that is the source of the action, e.g. when replaying this is the original run
# destination_run:: The run that's a result of the action, this will be set when the run has been created
class BulkActionItem < ApplicationRecord
  self.table_name = "BulkActionItem"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { cancel: "CANCEL", replay: "REPLAY" }, validate: true
  enum :status, { pending: "PENDING", completed: "COMPLETED", failed: "FAILED" }, default: :pending, validate: true

  validates :groupId, presence: true

  belongs_to :source_run, class_name: "TaskRun", foreign_key: "sourceRunId", inverse_of: :source_bulk_action_items
  belongs_to :destination_run, class_name: "TaskRun", foreign_key: "destinationRunId", optional: true,
                               inverse_of: :destination_bulk_action_items
end
