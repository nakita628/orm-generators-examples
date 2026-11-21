#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// What a `RetentionRunItem` row points at. Response deletions are grouped into one `survey` row per
/// survey, with the number of responses in `count`.
/// </summary>
public enum RetentionTargetType
{
    [PgName("survey")]
    Survey,

    [PgName("user")]
    User,
}
