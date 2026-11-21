#nullable enable

namespace Teable.Models;

/// <summary>
/// Actual output-value evidence; independent from continuation progress.
/// </summary>
public partial class ComputedUpdateChangeFrontier
{
    public string ScopeId { get; set; } = null!;

    public string Kind { get; set; } = null!;

    public string TableId { get; set; } = null!;

    public string RecordId { get; set; } = null!;

    public string FieldId { get; set; } = null!;
}
