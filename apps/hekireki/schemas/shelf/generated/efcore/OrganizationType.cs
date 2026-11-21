#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum OrganizationType
{
    [PgName("PERSONAL")]
    Personal,

    [PgName("TEAM")]
    Team,
}
