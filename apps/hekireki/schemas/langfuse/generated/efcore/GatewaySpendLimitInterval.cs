#nullable enable

using NpgsqlTypes;

namespace Langfuse.Models;

public enum GatewaySpendLimitInterval
{
    [PgName("DAILY")]
    Daily,

    [PgName("WEEKLY")]
    Weekly,

    [PgName("MONTHLY")]
    Monthly,

    [PgName("TOTAL")]
    Total,
}
