#nullable enable

using NpgsqlTypes;

namespace Trigger.Models;

public enum TaskExecutionMode
{
    [PgName("STANDARD")]
    Standard,

    [PgName("SANDBOX")]
    Sandbox,
}
