# *
# * Error group state is used to track when a user has interacted with an error (ignored/resolved)
# * The actual error data is in ClickHouse.
#
# environment:: *
#   * You can ignore/resolve an error across all environments, or specific ones
# ignoredUntil:: *
#   * Error is ignored until this date
# ignoredUntilOccurrenceRate:: *
#   * Error is ignored until this occurrence rate
# ignoredUntilTotalOccurrences:: *
#   * Error is ignored until this total occurrences
# ignoredAtOccurrenceCount:: Total occurrence count at the time the error was ignored (from ClickHouse).
#   Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
# ignoredAt:: *
#   * Error was ignored at this date
# ignoredReason:: *
#   * Reason for ignoring the error
# ignoredByUserId:: *
#   * User who ignored the error
# resolvedAt:: *
#   * Error was resolved at this date
# resolvedInVersion:: *
#   * Error was resolved in this version
# resolvedBy:: *
#   * User who resolved the error
class ErrorGroupState < ApplicationRecord
  self.table_name = "ErrorGroupState"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { unresolved: "UNRESOLVED", resolved: "RESOLVED", ignored: "IGNORED" }, default: :unresolved,
                                                                                        validate: true

  validates :taskIdentifier, presence: true,
                             uniqueness: { scope: [ :environmentId, :errorFingerprint ], conditions: -> { where.not(environmentId: nil) } }
  validates :errorFingerprint, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :error_group_states
  belongs_to :project, foreign_key: "projectId", inverse_of: :error_group_states
  belongs_to :environment, class_name: "RuntimeEnvironment", foreign_key: "environmentId", optional: true,
                           inverse_of: :error_group_states
end
