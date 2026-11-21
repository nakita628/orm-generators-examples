#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// Defines permission levels for API keys.
/// Controls what operations an API key can perform.
/// </summary>
public enum ApiKeyPermission
{
    [PgName("read")]
    Read,

    [PgName("write")]
    Write,

    [PgName("manage")]
    Manage,
}
