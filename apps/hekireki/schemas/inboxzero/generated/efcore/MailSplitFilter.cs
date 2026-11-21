#nullable enable

namespace InboxZero.Models;

/// <summary>
/// One condition row in the split builder. A split may hold several of the same
/// kind (two senders, two labels), so kind is deliberately not unique per split.
/// </summary>
public partial class MailSplitFilter
{
    public string Id { get; set; } = null!;

    public MailSplitFilterKind Kind { get; set; }

    public string? Value { get; set; }

    public int Order { get; set; }

    public string MailSplitId { get; set; } = null!;

    public virtual MailSplit MailSplit { get; set; } = null!;
}
