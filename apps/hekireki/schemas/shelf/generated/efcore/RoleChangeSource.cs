#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// How a membership's role came to change.
/// </summary>
public enum RoleChangeSource
{
    [PgName("MANUAL")]
    Manual,

    [PgName("SSO")]
    Sso,
}
