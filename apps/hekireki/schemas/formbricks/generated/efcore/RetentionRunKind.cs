#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// What wrote a `RetentionRun`: the nightly sweep, or the archive purge recording the deletion of surveys
/// the surveys policy archived. Only sweep runs count as the sweep having run (its lease, its warning
/// restart after a gap, where its scan resumes, the health check).
/// </summary>
public enum RetentionRunKind
{
    [PgName("sweep")]
    Sweep,

    [PgName("purge")]
    Purge,
}
