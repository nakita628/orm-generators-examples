#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// The surveys policy's conditions. Every ticked one must hold, and they share the policy's period.
/// </summary>
public enum RetentionSurveyCondition
{
    [PgName("noResponse")]
    NoResponse,

    [PgName("noChange")]
    NoChange,

    [PgName("createdBefore")]
    CreatedBefore,
}
