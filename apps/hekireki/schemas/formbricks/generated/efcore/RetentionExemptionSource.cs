#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// Who made an exemption: a person (with a reason), or a restore from the archive under a surveys policy
/// whose only condition is "created before" (`restartSurveyRetention`), which carries no reason.
/// </summary>
public enum RetentionExemptionSource
{
    [PgName("manual")]
    Manual,

    [PgName("restore")]
    Restore,
}
