#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum OrganizationRoles
{
    [PgName("ADMIN")]
    Admin,

    [PgName("BASE")]
    Base,

    [PgName("OWNER")]
    Owner,

    [PgName("SELF_SERVICE")]
    SelfService,

    [PgName("MANAGER")]
    Manager,
}
