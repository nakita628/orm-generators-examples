#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// * Determines how an asset is tracked throughout the system. Immutable after creation.
/// </summary>
public enum AssetType
{
    [PgName("INDIVIDUAL")]
    Individual,

    [PgName("QUANTITY_TRACKED")]
    QuantityTracked,
}
