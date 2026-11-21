# Represents a user action that can trigger surveys.
# Used to define points in the user journey where surveys can be shown.
#
# @property id - Unique identifier for the action
# @property name - Display name of the action
# @property type - Whether this is a code or no-code action
# @property key - Unique identifier used in code implementation
# @property noCodeConfig - Configuration for no-code setup
# @property workspace - The workspace this action belongs to
#
# noCodeConfig:: [ActionClassNoCodeConfig]
class ActionClass < ApplicationRecord
  self.table_name = "ActionClass"
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  enum :type, { code: "code", no_code: "noCode" }, validate: true

  validates :name, presence: true, uniqueness: { scope: :workspaceId }
  validates :key, uniqueness: { scope: :workspaceId, allow_nil: true }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :action_classes
  has_many :survey_triggers, foreign_key: "actionClassId", inverse_of: :action_class, dependent: :destroy
  has_many :creative_triggers, foreign_key: "actionClassId", inverse_of: :action_class, dependent: :destroy
end
