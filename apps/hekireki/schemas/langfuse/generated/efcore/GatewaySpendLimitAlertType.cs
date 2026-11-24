#nullable enable

using NpgsqlTypes;

namespace Langfuse.Models;

public enum GatewaySpendLimitAlertType
{
    [PgName("WARNING")]
    Warning,

    [PgName("EXCEEDED")]
    Exceeded,
}
