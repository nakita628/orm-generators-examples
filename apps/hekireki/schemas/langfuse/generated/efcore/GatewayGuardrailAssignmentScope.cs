#nullable enable

using NpgsqlTypes;

namespace Langfuse.Models;

public enum GatewayGuardrailAssignmentScope
{
    [PgName("ORGANIZATION")]
    Organization,

    [PgName("API_KEY")]
    ApiKey,

    [PgName("USER")]
    User,
}
