#nullable enable

using System;
using System.Collections.Generic;

namespace Langfuse.Models;

public partial class GatewayGuardrail
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string Name { get; set; } = null!;

    public GatewayGuardrailType Type { get; set; } = GatewayGuardrailType.Regex;

    public GatewayGuardrailAction Action { get; set; }

    public string Config { get; set; } = "{}";

    public string? EncryptedCredential { get; set; }

    public string? CreatedBy { get; set; }

    public virtual User? CreatedByUser { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<GatewayGuardrailAssignment> Assignments { get; set; } = new List<GatewayGuardrailAssignment>();
}
