#nullable enable

using System;

namespace LangWatch.Models;

public partial class AuthzBraidWrite
{
    public string OrganizationId { get; set; } = null!;

    public string CommandId { get; set; } = null!;

    public string GrantId { get; set; } = null!;

    public string ReaderProjectId { get; set; } = null!;

    public string OwnerProjectId { get; set; } = null!;

    public string TermsKey { get; set; } = null!;

    public string Condition { get; set; } = null!;

    public DateTime? ExpiresAt { get; set; }

    public string? ReplacesGrantId { get; set; }

    public string? CreatedByUserId { get; set; }

    public DateTime CreatedAt { get; set; }
}
