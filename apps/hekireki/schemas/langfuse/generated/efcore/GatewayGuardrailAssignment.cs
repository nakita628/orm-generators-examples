#nullable enable

using System;

namespace Langfuse.Models;

public partial class GatewayGuardrailAssignment
{
    public string Id { get; set; } = null!;

    public string GuardrailId { get; set; } = null!;

    public virtual GatewayGuardrail Guardrail { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public GatewayGuardrailAssignmentScope Scope { get; set; }

    public string? ApiKeyId { get; set; }

    public virtual ApiKey? ApiKey { get; set; }

    public string? UserId { get; set; }

    public virtual User? User { get; set; }

    public DateTime CreatedAt { get; set; }
}
