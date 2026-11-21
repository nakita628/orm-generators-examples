#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// * Controls what happens after checkout for quantity-tracked assets.
/// </summary>
public enum ConsumptionType
{
    [PgName("ONE_WAY")]
    OneWay,

    [PgName("TWO_WAY")]
    TwoWay,
}
