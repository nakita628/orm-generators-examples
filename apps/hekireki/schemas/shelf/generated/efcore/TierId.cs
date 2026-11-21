#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum TierId
{
    [PgName("free")]
    Free,

    [PgName("tier_1")]
    Tier1,

    [PgName("tier_2")]
    Tier2,

    [PgName("custom")]
    Custom,
}
