#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// What a reviewer decided about a response. One value today; null on `Response.outcome` means none.
/// </summary>
public enum ResponseOutcome
{
    [PgName("lowQuality")]
    LowQuality,
}
