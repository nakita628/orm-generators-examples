# Represents a label that can be applied to survey responses.
# Used for categorizing and organizing responses within a workspace.
#
# @property id - Unique identifier for the tag
# @property name - Display name of the tag
# @property responses - Survey responses tagged with this label
# @property workspace - The workspace this tag belongs to
class Tag < ApplicationRecord
  self.table_name = "Tag"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :name, presence: true, uniqueness: { scope: :workspaceId }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :tags
  has_many :responses, class_name: "TagsOnResponses", foreign_key: "tagId", inverse_of: :tag, dependent: :destroy
end
