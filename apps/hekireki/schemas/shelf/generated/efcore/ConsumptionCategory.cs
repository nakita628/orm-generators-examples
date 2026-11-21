#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// * Classifies the type of quantity change in a ConsumptionLog entry.
/// </summary>
public enum ConsumptionCategory
{
    [PgName("CHECKOUT")]
    Checkout,

    [PgName("RETURN")]
    Return,

    [PgName("RESTOCK")]
    Restock,

    [PgName("ADJUSTMENT")]
    Adjustment,

    [PgName("LOSS")]
    Loss,

    [PgName("CONSUME")]
    Consume,

    [PgName("DAMAGE")]
    Damage,

    [PgName("MOVE")]
    Move,

    [PgName("INITIAL")]
    Initial,

    [PgName("OPENING_BALANCE")]
    OpeningBalance,
}
