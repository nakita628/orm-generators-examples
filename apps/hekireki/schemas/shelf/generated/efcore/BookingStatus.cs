#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum BookingStatus
{
    [PgName("DRAFT")]
    Draft,

    [PgName("RESERVED")]
    Reserved,

    [PgName("ONGOING")]
    Ongoing,

    [PgName("OVERDUE")]
    Overdue,

    [PgName("COMPLETE")]
    Complete,

    [PgName("ARCHIVED")]
    Archived,

    [PgName("CANCELLED")]
    Cancelled,
}
