# The per-branch bookkeeping behind pull-request mapping: when a branch was
# last looked up on GitHub, how many pull requests it resolved to, and when it
# may be asked about again. A branch with no pull request is the common case,
# so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
# costing a GitHub call on every read.
# Spec: specs/integrations/github-connection.feature.
#
# repositoryFullName:: "owner/repo", stored LOWERCASED, matching GithubPullRequest.
# notFoundAt:: Set while the branch resolves to no pull request; cleared once one is
#   found.
# recheckAfter:: The earliest instant GitHub may be asked about this branch again.
# attempts:: Consecutive lookups that found nothing, backing off the recheck.
# lastRequestedAt:: When a reader last asked about this branch, whether or not GitHub was
#   called.
class GithubBranchPullRequestCheck < ApplicationRecord
  self.table_name = "GithubBranchPullRequestCheck"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :prCount, default: 0
  attribute :attempts, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :repositoryHost, presence: true,
                             uniqueness: { scope: [ :organizationId, :repositoryFullName, :headBranch ] }
  validates :repositoryFullName, presence: true
  validates :headBranch, presence: true
  validates :lastCheckedAt, presence: true
  validates :lastRequestedAt, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :github_branch_pr_checks
end
