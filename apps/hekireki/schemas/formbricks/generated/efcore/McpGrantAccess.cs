#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

public enum McpGrantAccess
{
    [PgName("all")]
    All,

    [PgName("organizations")]
    Organizations,

    [PgName("workspaces")]
    Workspaces,
}
