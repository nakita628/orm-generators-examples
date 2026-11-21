#nullable enable

namespace Trigger.Models;

/// <summary>
/// deprecated, we hadn't included the project id in the unique constraint
/// </summary>
public partial class TaskRunCounter
{
    public string TaskIdentifier { get; set; } = null!;

    public int LastNumber { get; set; } = 0;
}
