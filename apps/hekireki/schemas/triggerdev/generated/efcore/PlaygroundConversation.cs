#nullable enable

using System;

namespace Trigger.Models;

public partial class PlaygroundConversation
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The chat session ID used by the transport
    /// </summary>
    public string ChatId { get; set; } = null!;

    /// <summary>
    /// User-editable conversation title (auto-generated from first message)
    /// </summary>
    public string Title { get; set; } = "New conversation";

    /// <summary>
    /// Which agent this conversation is with
    /// </summary>
    public string AgentSlug { get; set; } = null!;

    /// <summary>
    /// The current active run backing this conversation (null if no run yet)
    /// </summary>
    public string? RunId { get; set; }

    public virtual TaskRun? Run { get; set; }

    /// <summary>
    /// The client data JSON used for this conversation
    /// </summary>
    public string? ClientData { get; set; }

    /// <summary>
    /// Accumulated UIMessages from completed turns (for resume without stream replay)
    /// </summary>
    public string? Messages { get; set; }

    /// <summary>
    /// Last SSE event ID — resume from this position to avoid replaying old turns
    /// </summary>
    public string? LastEventId { get; set; }

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual RuntimeEnvironment RuntimeEnvironment { get; set; } = null!;

    public string RuntimeEnvironmentId { get; set; } = null!;

    /// <summary>
    /// The user who started this conversation
    /// </summary>
    public string UserId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
