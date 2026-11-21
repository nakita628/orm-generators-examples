# Represents a feedback directory (Hub tenant) owned by an organization.
# Directories group feedback data and are assigned to workspaces for access control.
#
# @property id - Unique identifier for the directory
# @property name - Display name of the directory
# @property isArchived - Soft delete flag
# @property organization - The parent organization
# @property workspaces - Workspaces assigned to this directory
class FeedbackDirectory < ApplicationRecord
  self.table_name = "FeedbackDirectory"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :isArchived, default: false

  validates :name, presence: true, uniqueness: { scope: :organizationId }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :feedback_directories
  has_many :workspaces, class_name: "FeedbackDirectoryWorkspace", foreign_key: "feedbackDirectoryId",
                        inverse_of: :feedback_directory, dependent: :destroy
  has_many :feedback_sources, foreign_key: "feedbackDirectoryId", inverse_of: :feedback_directory, dependent: :destroy
  has_many :charts, foreign_key: "feedbackDirectoryId", inverse_of: :feedback_directory, dependent: :restrict_with_error
end
