#nullable enable

using NpgsqlTypes;

namespace Langfuse.Models;

public enum GatewaySpendLimitScope
{
    [PgName("API_KEY")]
    ApiKey,

    [PgName("USER")]
    User,

    [PgName("API_KEY_DEFAULT")]
    ApiKeyDefault,

    [PgName("USER_DEFAULT")]
    UserDefault,

    [PgName("ORGANIZATION")]
    Organization,
}
