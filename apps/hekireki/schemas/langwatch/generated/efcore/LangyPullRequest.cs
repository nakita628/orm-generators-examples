#nullable enable

namespace LangWatch.Models;

/// <summary>
/// One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
/// checks look it up by repository and number to wake the conversation; the daily pull request
/// cap counts a person's rows by day. Rebuilt by a projection replay.
/// </summary>
public partial class LangyPullRequest
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// `owner/repo`, lowercased: GitHub matches repository names without case.
    /// </summary>
    public string RepositoryFullName { get; set; } = null!;

    public int PrNumber { get; set; }

    public string ConversationId { get; set; } = null!;

    /// <summary>
    /// The conversation's owner, who opened it through Langy.
    /// </summary>
    public string UserId { get; set; } = null!;

    /// <summary>
    /// When the answer that opened it was recorded (epoch ms).
    /// </summary>
    public double OpenedAt { get; set; }

    /// <summary>
    /// The agent_responded event it came from.
    /// </summary>
    public string SourceEventId { get; set; } = null!;
}
