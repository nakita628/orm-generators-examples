#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A pull request discovered through the organization's GitHub connection and
/// mapped to the branch it was opened from, so a coding-agent session that
/// knows only its branch can name the pull request it produced.
/// Spec: specs/integrations/github-connection.feature.
///
/// The stored title, state, draft and timestamp columns are SNAPSHOTS taken
/// when the mapping was last refreshed. They are fallback labels only: the
/// frontend reads live status from GitHub, so a stale row never presents
/// itself as current truth.
/// </summary>
public partial class GithubPullRequest
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    /// <summary>
    /// The GitHub host the repository lives on, e.g. "github.com".
    /// </summary>
    public string RepositoryHost { get; set; } = null!;

    /// <summary>
    /// "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
    /// session reported.
    /// </summary>
    public string RepositoryFullName { get; set; } = null!;

    /// <summary>
    /// The branch the pull request was opened from.
    /// </summary>
    public string HeadBranch { get; set; } = null!;

    public int PrNumber { get; set; }

    public string HtmlUrl { get; set; } = null!;

    public string Title { get; set; } = null!;

    /// <summary>
    /// GitHub's own state string: "open" or "closed".
    /// </summary>
    public string State { get; set; } = null!;

    public bool IsDraft { get; set; } = false;

    public string? AuthorLogin { get; set; }

    public DateTime PrCreatedAt { get; set; }

    public DateTime? PrClosedAt { get; set; }

    public DateTime? PrMergedAt { get; set; }

    /// <summary>
    /// GitHub's own `updated_at` for the snapshot this row holds, and the
    /// freshness guard on every write to it: a snapshot older than the stored one
    /// is skipped. GitHub may deliver webhooks out of order, and without the
    /// guard a late `opened` lands after `closed` and reopens a merged pull
    /// request. NULL means the row predates the column, so any write is accepted.
    /// </summary>
    public DateTime? PrUpdatedAt { get; set; }

    /// <summary>
    /// When this row was first mapped to its branch.
    /// </summary>
    public DateTime MappedAt { get; set; }

    /// <summary>
    /// When the snapshot columns were last refreshed from GitHub.
    /// </summary>
    public DateTime LastCheckedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
