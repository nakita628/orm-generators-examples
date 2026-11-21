#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The organization's GitHub connection: a GitHub App installation mapped to a
/// LangWatch organization. Short-lived (1h) installation access tokens are
/// minted against these rows on demand, by Langy for writes (bot-authored pull
/// requests, scoped to the installation's repositories) and by pull-request
/// linkage for reads. No user token and no secret are stored here, the crown
/// jewel is the App private key, held only in the control-plane env.
/// Spec: specs/integrations/github-connection.feature.
///
/// installationId is globally unique (one GitHub installation, one row); a
/// single LangWatch organization may own MANY installations (different GitHub
/// accounts), so the org to installation relationship is one-to-many.
/// </summary>
public partial class GithubInstallation
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// GitHub installation id (numeric; stored as string). Globally unique.
    /// </summary>
    public string InstallationId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    /// <summary>
    /// The GitHub account (organization or user) the App is installed on.
    /// </summary>
    public string AccountLogin { get; set; } = null!;

    /// <summary>
    /// "Organization" or "User".
    /// </summary>
    public string AccountType { get; set; } = null!;

    /// <summary>
    /// GitHub numeric account id (stable across renames).
    /// </summary>
    public string AccountId { get; set; } = null!;

    /// <summary>
    /// "all" or "selected" — whether the App can see every repo on the account
    /// or only a chosen subset.
    /// </summary>
    public string RepositorySelection { get; set; } = null!;

    /// <summary>
    /// Cached repository list [{ id, fullName }] when repositorySelection =
    /// "selected". Null when "all" (the full set is resolved live at mint time).
    /// </summary>
    public string? Repositories { get; set; }

    /// <summary>
    /// Set when GitHub suspends the installation; cleared on unsuspend. A
    /// suspended installation cannot mint tokens.
    /// </summary>
    public DateTime? SuspendedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
