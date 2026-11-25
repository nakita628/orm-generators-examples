#nullable enable

using NpgsqlTypes;

namespace Langfuse.Models;

public enum GatewayGuardrailType
{
    [PgName("REGEX")]
    Regex,

    [PgName("BEDROCK")]
    Bedrock,
}
