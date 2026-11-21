#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// What an invoice of a connected customer is for.
/// </summary>
public enum ConnectedInvoiceKind
{
    [PgName("annual")]
    Annual,

    [PgName("seat_change")]
    SeatChange,

    [PgName("usage")]
    Usage,
}
