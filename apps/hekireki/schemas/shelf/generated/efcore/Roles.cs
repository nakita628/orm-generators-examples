#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum Roles
{
    [PgName("USER")]
    User,

    [PgName("ADMIN")]
    Admin,
}
