#nullable enable

namespace Calcom.Models;

public partial class EventTypeCustomInput
{
    public int Id { get; set; }

    public int EventTypeId { get; set; }

    public virtual EventType EventType { get; set; } = null!;

    public string Label { get; set; } = null!;

    public EventTypeCustomInputType Type { get; set; }

    /// <summary>
    /// @zod.import(["import { customInputOptionSchema } from '../../zod-utils'"]).custom.use(customInputOptionSchema)
    /// </summary>
    public string? Options { get; set; }

    public bool Required { get; set; }

    public string Placeholder { get; set; } = "";
}
