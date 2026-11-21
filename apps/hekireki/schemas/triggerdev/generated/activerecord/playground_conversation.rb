# chatId:: The chat session ID used by the transport
# title:: User-editable conversation title (auto-generated from first message)
# agentSlug:: Which agent this conversation is with
# runId:: The current active run backing this conversation (null if no run yet)
# clientData:: The client data JSON used for this conversation
# messages:: Accumulated UIMessages from completed turns (for resume without stream replay)
# lastEventId:: Last SSE event ID — resume from this position to avoid replaying old turns
# userId:: The user who started this conversation
class PlaygroundConversation < ApplicationRecord
  self.table_name = "PlaygroundConversation"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :title, default: "New conversation"
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :chatId, presence: true, uniqueness: { scope: :runtimeEnvironmentId }
  validates :agentSlug, presence: true
  validates :userId, presence: true

  belongs_to :run, class_name: "TaskRun", foreign_key: "runId", optional: true, inverse_of: :playground_conversations
  belongs_to :project, foreign_key: "projectId", inverse_of: :playground_conversations
  belongs_to :runtime_environment, foreign_key: "runtimeEnvironmentId", inverse_of: :playground_conversations
end
