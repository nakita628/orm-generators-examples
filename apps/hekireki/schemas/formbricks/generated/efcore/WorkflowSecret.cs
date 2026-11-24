#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Encrypted value of an HTTP request secret header (ENG-7168, ADR 0043). The definition and every version
/// store only `{name, secret: true, secretId}`; Replace overwrites this row in place, so the enabled
/// version uses the new value at once. Rows nothing references are deleted on Enable; all are deleted
/// with the workflow. A value is only sent to the origin it was set for (ADR 0051).
/// </summary>
public partial class WorkflowSecret
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Workflow Workflow { get; set; } = null!;

    public string WorkflowId { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    /// <summary>
    /// `symmetricEncrypt` ciphertext (AES-256-GCM). Never logged, never returned.
    /// </summary>
    public string Value { get; set; } = null!;

    /// <summary>
    /// Origin (scheme, host, port) of the step's URL when the value was set or replaced; the value is only
    /// sent there (ADR 0051). Null for a value set before host binding: held to the step's saved URL
    /// until it is set again.
    /// </summary>
    public string? BoundOrigin { get; set; }
}
