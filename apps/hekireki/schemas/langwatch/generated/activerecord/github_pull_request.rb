# A pull request discovered through the organization's GitHub connection and
# mapped to the branch it was opened from, so a coding-agent session that
# knows only its branch can name the pull request it produced.
# Spec: specs/integrations/github-connection.feature.
#
# The stored title, state, draft and timestamp columns are SNAPSHOTS taken
# when the mapping was last refreshed. They are fallback labels only: the
# frontend reads live status from GitHub, so a stale row never presents
# itself as current truth.
#
# repositoryHost:: The GitHub host the repository lives on, e.g. "github.com".
# repositoryFullName:: "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
#   session reported.
# headBranch:: The branch the pull request was opened from.
# state:: GitHub's own state string: "open" or "closed".
# prUpdatedAt:: GitHub's own `updated_at` for the snapshot this row holds, and the
#   freshness guard on every write to it: a snapshot older than the stored one
#   is skipped. GitHub may deliver webhooks out of order, and without the
#   guard a late `opened` lands after `closed` and reopens a merged pull
#   request. NULL means the row predates the column, so any write is accepted.
# mappedAt:: When this row was first mapped to its branch.
# lastCheckedAt:: When the snapshot columns were last refreshed from GitHub.
class GithubPullRequest < ApplicationRecord
  self.table_name = "GithubPullRequest"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :isDraft, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :repositoryHost, presence: true, uniqueness: { scope: [ :organizationId, :repositoryFullName, :prNumber ] }
  validates :repositoryFullName, presence: true
  validates :headBranch, presence: true
  validates :prNumber, presence: true
  validates :htmlUrl, presence: true
  validates :title, presence: true
  validates :state, presence: true
  validates :prCreatedAt, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :github_pull_requests
end
