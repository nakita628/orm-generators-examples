#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum KitStatus
{
    [PgName("AVAILABLE")]
    Available,

    [PgName("IN_CUSTODY")]
    InCustody,

    [PgName("CHECKED_OUT")]
    CheckedOut,
}
