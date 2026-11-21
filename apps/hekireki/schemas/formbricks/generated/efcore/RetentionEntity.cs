#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// Data retention (ENG-3697). The kind of data a policy governs.
/// </summary>
public enum RetentionEntity
{
    [PgName("responses")]
    Responses,

    [PgName("surveys")]
    Surveys,

    [PgName("members")]
    Members,
}
