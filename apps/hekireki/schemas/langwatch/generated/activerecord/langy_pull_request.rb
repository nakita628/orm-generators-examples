# One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
# checks look it up by repository and number to wake the conversation; the daily pull request
# cap counts a person's rows by day. Rebuilt by a projection replay.
#
# repositoryFullName:: `owner/repo`, lowercased: GitHub matches repository names without case.
# userId:: The conversation's owner, who opened it through Langy.
# openedAt:: When the answer that opened it was recorded (epoch ms).
# sourceEventId:: The agent_responded event it came from.
class LangyPullRequest < ApplicationRecord
  self.table_name = "LangyPullRequest"

  attribute :id, default: -> { Nanoid.generate }

  validates :projectId, presence: true
  validates :repositoryFullName, presence: true, uniqueness: { scope: :prNumber }
  validates :prNumber, presence: true
  validates :conversationId, presence: true
  validates :userId, presence: true
  validates :openedAt, presence: true
  validates :sourceEventId, presence: true
end
