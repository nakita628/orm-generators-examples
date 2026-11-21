#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A project nurturing saw traces of with no placement copy, and when it first did: every worker
/// retries its trace batches only within one window from then, across restarts. Deleted when the
/// copy lands.
/// </summary>
public partial class NurturingUncopiedProject
{
    public string ProjectId { get; set; } = null!;

    public DateTime FirstSeenAt { get; set; }
}
