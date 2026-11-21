#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum TagUseFor
{
    [PgName("ASSET")]
    Asset,

    [PgName("BOOKING")]
    Booking,
}
