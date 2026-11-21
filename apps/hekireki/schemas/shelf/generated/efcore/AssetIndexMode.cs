#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum AssetIndexMode
{
    [PgName("SIMPLE")]
    Simple,

    [PgName("ADVANCED")]
    Advanced,
}
