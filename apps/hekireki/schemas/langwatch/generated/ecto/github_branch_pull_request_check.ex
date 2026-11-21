defmodule LangWatch.GithubBranchPullRequestCheck do
  @moduledoc """
  The per-branch bookkeeping behind pull-request mapping: when a branch was
  last looked up on GitHub, how many pull requests it resolved to, and when it
  may be asked about again. A branch with no pull request is the common case,
  so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
  costing a GitHub call on every read.
  Spec: specs/integrations/github-connection.feature.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          repository_host: String.t(),
          repository_full_name: String.t(),
          head_branch: String.t(),
          last_checked_at: DateTime.t(),
          pr_count: integer(),
          not_found_at: DateTime.t() | nil,
          recheck_after: DateTime.t() | nil,
          attempts: integer(),
          last_requested_at: DateTime.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil
        }

  schema "GithubBranchPullRequestCheck" do
    field(:id, :string, primary_key: true)
    field(:repository_host, :string, source: :repositoryHost)
    # "owner/repo", stored LOWERCASED, matching GithubPullRequest.
    field(:repository_full_name, :string, source: :repositoryFullName)
    field(:head_branch, :string, source: :headBranch)
    field(:last_checked_at, LangWatch.PrismaDateTime, source: :lastCheckedAt)
    field(:pr_count, :integer, default: 0, source: :prCount)
    # Set while the branch resolves to no pull request; cleared once one is
    # found.
    field(:not_found_at, LangWatch.PrismaDateTime, source: :notFoundAt)
    # The earliest instant GitHub may be asked about this branch again.
    field(:recheck_after, LangWatch.PrismaDateTime, source: :recheckAfter)
    # Consecutive lookups that found nothing, backing off the recheck.
    field(:attempts, :integer, default: 0)
    # When a reader last asked about this branch, whether or not GitHub was
    # called.
    field(:last_requested_at, LangWatch.PrismaDateTime, source: :lastRequestedAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end