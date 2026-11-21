#nullable enable

namespace LangWatch.Models;

/// <summary>
/// A connection's managed users, including resources without an externalId.
/// Scalar keys only; readers exclude users that no longer exist.
/// </summary>
public partial class ScimDirectoryUser
{
    public string OrganizationId { get; set; } = null!;

    public string ConnectionId { get; set; } = null!;

    public string UserId { get; set; } = null!;
}
