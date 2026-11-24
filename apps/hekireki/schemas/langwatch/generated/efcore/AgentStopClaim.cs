#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The one open stop an agent may have: the api takes this row before it sends a stop's request,
/// so two requests at once cannot both open one. A later request takes it over once the stop it
/// names has closed.
/// `sentAt` is set once the request command was sent: from then on the claim passes on only after
/// its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
/// `decisionTakenAt` is set while that decision is in flight and cleared once its command was
/// sent; a decision left in flight is taken over once abandoned.
/// </summary>
public partial class AgentStopClaim
{
    public string OrganizationId { get; set; } = null!;

    public string AgentId { get; set; } = null!;

    public string StopId { get; set; } = null!;

    public DateTime ClaimedAt { get; set; }

    public DateTime? SentAt { get; set; }

    public string? Decision { get; set; }

    public DateTime? DecisionTakenAt { get; set; }
}
