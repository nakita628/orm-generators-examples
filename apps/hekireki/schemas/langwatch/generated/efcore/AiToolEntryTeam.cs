#nullable enable

namespace LangWatch.Models;

/// <summary>
/// @deprecated Per-team scope binding for AiToolEntry. Superseded by
/// AiToolEntryDepartment; kept so the rename migration is non-destructive.
/// No new rows are written here.
/// </summary>
public partial class AiToolEntryTeam
{
    public string Id { get; set; } = null!;

    public string EntryId { get; set; } = null!;

    public string TeamId { get; set; } = null!;

    public virtual AiToolEntry Entry { get; set; } = null!;

    public virtual Team Team { get; set; } = null!;
}
