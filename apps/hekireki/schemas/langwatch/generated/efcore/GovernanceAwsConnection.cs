#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// An AWS credential a Bedrock kill switch target signs with: a static key pair, a role to assume,
/// or both. The secret access key and the external id are stored sealed.
/// </summary>
public partial class GovernanceAwsConnection
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? AccessKeyId { get; set; }

    public string? EncryptedSecretAccessKey { get; set; }

    public string? RoleArn { get; set; }

    public string? EncryptedExternalId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
