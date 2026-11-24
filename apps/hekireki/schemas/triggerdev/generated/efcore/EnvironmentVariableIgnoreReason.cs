#nullable enable

using NpgsqlTypes;

namespace Trigger.Models;

public enum EnvironmentVariableIgnoreReason
{
    [PgName("EMPTY")]
    Empty,

    [PgName("MISSING")]
    Missing,

    [PgName("OIDC")]
    Oidc,
}
