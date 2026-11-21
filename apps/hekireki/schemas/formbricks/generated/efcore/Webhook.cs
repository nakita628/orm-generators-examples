#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a webhook endpoint for receiving survey-related events.
/// Webhooks can be configured to receive notifications about response creation, updates, and completion.
///
/// @property id - Unique identifier for the webhook
/// @property name - Optional display name for the webhook
/// @property url - The endpoint URL where events will be sent
/// @property source - Origin of the webhook (user, zapier, make, etc.)
/// @property workspace - Associated workspace
/// @property triggers - Types of events that trigger this webhook
/// @property surveyIds - List of surveys this webhook is monitoring
/// </summary>
public partial class Webhook
{
    public string Id { get; set; } = null!;

    public string? Name { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Url { get; set; } = null!;

    public WebhookSource Source { get; set; } = WebhookSource.User;

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public List<PipelineTriggers>? Triggers { get; set; }

    public List<string>? SurveyIds { get; set; }

    public string? Secret { get; set; }
}
