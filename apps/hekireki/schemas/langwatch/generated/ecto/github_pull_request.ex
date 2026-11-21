defmodule LangWatch.GithubPullRequest do
  @moduledoc """
  A pull request discovered through the organization's GitHub connection and
  mapped to the branch it was opened from, so a coding-agent session that
  knows only its branch can name the pull request it produced.
  Spec: specs/integrations/github-connection.feature.

  The stored title, state, draft and timestamp columns are SNAPSHOTS taken
  when the mapping was last refreshed. They are fallback labels only: the
  frontend reads live status from GitHub, so a stale row never presents
  itself as current truth.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          repository_host: String.t(),
          repository_full_name: String.t(),
          head_branch: String.t(),
          pr_number: integer(),
          html_url: String.t(),
          title: String.t(),
          state: String.t(),
          is_draft: boolean(),
          author_login: String.t() | nil,
          pr_created_at: DateTime.t(),
          pr_closed_at: DateTime.t() | nil,
          pr_merged_at: DateTime.t() | nil,
          pr_updated_at: DateTime.t() | nil,
          mapped_at: DateTime.t(),
          last_checked_at: DateTime.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil
        }

  schema "GithubPullRequest" do
    field(:id, :string, primary_key: true)
    # The GitHub host the repository lives on, e.g. "github.com".
    field(:repository_host, :string, source: :repositoryHost)
    # "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
    # session reported.
    field(:repository_full_name, :string, source: :repositoryFullName)
    # The branch the pull request was opened from.
    field(:head_branch, :string, source: :headBranch)
    field(:pr_number, :integer, source: :prNumber)
    field(:html_url, :string, source: :htmlUrl)
    field(:title, :string)
    # GitHub's own state string: "open" or "closed".
    field(:state, :string)
    field(:is_draft, :boolean, default: false, source: :isDraft)
    field(:author_login, :string, source: :authorLogin)
    field(:pr_created_at, LangWatch.PrismaDateTime, source: :prCreatedAt)
    field(:pr_closed_at, LangWatch.PrismaDateTime, source: :prClosedAt)
    field(:pr_merged_at, LangWatch.PrismaDateTime, source: :prMergedAt)
    # GitHub's own `updated_at` for the snapshot this row holds, and the
    # freshness guard on every write to it: a snapshot older than the stored one
    # is skipped. GitHub may deliver webhooks out of order, and without the
    # guard a late `opened` lands after `closed` and reopens a merged pull
    # request. NULL means the row predates the column, so any write is accepted.
    field(:pr_updated_at, LangWatch.PrismaDateTime, source: :prUpdatedAt)
    # When this row was first mapped to its branch.
    field(:mapped_at, LangWatch.PrismaDateTime, autogenerate: true, source: :mappedAt)
    # When the snapshot columns were last refreshed from GitHub.
    field(:last_checked_at, LangWatch.PrismaDateTime, autogenerate: true, source: :lastCheckedAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end