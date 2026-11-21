#nullable enable

namespace LangWatch.Models;

/// <summary>
/// A connection currently entitled to use an organization-owned verified
/// domain. Usually there is one holder. During an Auth0 migration there are
/// exactly two: the grandfathered predecessor and its direct replacement.
///
/// The projection repository enforces that relationship before inserting a
/// second holder. The flattened migration also installs partial unique live
/// legacy/direct slots on `SsoConnection`, closing concurrent third-writer
/// races that Prisma cannot describe declaratively.
/// </summary>
public partial class SsoVerifiedDomainHolder
{
    public string Domain { get; set; } = null!;

    public string ConnectionId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual SsoVerifiedDomain Ownership { get; set; } = null!;
}
