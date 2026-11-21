#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The per-branch bookkeeping behind pull-request mapping: when a branch was
/// last looked up on GitHub, how many pull requests it resolved to, and when it
/// may be asked about again. A branch with no pull request is the common case,
/// so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
/// costing a GitHub call on every read.
/// Spec: specs/integrations/github-connection.feature.
/// </summary>
public partial class GithubBranchPullRequestCheck
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string RepositoryHost { get; set; } = null!;

    /// <summary>
    /// "owner/repo", stored LOWERCASED, matching GithubPullRequest.
    /// </summary>
    public string RepositoryFullName { get; set; } = null!;

    public string HeadBranch { get; set; } = null!;

    public DateTime LastCheckedAt { get; set; }

    public int PrCount { get; set; } = 0;

    /// <summary>
    /// Set while the branch resolves to no pull request; cleared once one is
    /// found.
    /// </summary>
    public DateTime? NotFoundAt { get; set; }

    /// <summary>
    /// The earliest instant GitHub may be asked about this branch again.
    /// </summary>
    public DateTime? RecheckAfter { get; set; }

    /// <summary>
    /// Consecutive lookups that found nothing, backing off the recheck.
    /// </summary>
    public int Attempts { get; set; } = 0;

    /// <summary>
    /// When a reader last asked about this branch, whether or not GitHub was
    /// called.
    /// </summary>
    public DateTime LastRequestedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
