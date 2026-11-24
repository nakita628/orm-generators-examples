# definition:: [WorkflowDefinition]
# webhookTokenHash:: Webhook Secret URL token (ADRs 0048, 0057): SHA-256 hash to check calls, and an encrypted copy so
#   the editor can show the URL again. Never part of the definition, its versions or its copies.
class Workflow < ApplicationRecord
  self.table_name = "Workflow"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid2.call }

  enum :status, { draft: "draft", enabled: "enabled", disabled: "disabled", archived: "archived" }, default: :draft,
                                                                                                    validate: true

  validates :id, uniqueness: { scope: :workspaceId }
  validates :name, presence: true
  validates :definition, exclusion: { in: [ nil ] }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :workflows
  belongs_to :creator, class_name: "User", foreign_key: "createdBy", optional: true, inverse_of: :workflows
  has_many :versions, class_name: "WorkflowVersion", foreign_key: [ :workflowId, :workspaceId ],
                      primary_key: [ :id, :workspaceId ], inverse_of: :workflow, dependent: :destroy
  has_many :runs, class_name: "WorkflowRun", foreign_key: [ :workflowId, :workspaceId ],
                  primary_key: [ :id, :workspaceId ], inverse_of: :workflow, dependent: :destroy
  has_many :secrets, class_name: "WorkflowSecret", foreign_key: [ :workflowId, :workspaceId ],
                     primary_key: [ :id, :workspaceId ], inverse_of: :workflow, dependent: :destroy
  has_many :samples, class_name: "WorkflowSample", foreign_key: [ :workflowId, :workspaceId ],
                     primary_key: [ :id, :workspaceId ], inverse_of: :workflow, dependent: :destroy
end
