#nullable enable

using NpgsqlTypes;

namespace Langfuse.Models;

public enum GatewayGuardrailAction
{
    [PgName("BLOCK")]
    Block,

    [PgName("LOG")]
    Log,

    [PgName("REDACT")]
    Redact,
}
