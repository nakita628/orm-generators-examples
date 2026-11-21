#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum InviteStatuses
{
    [PgName("PENDING")]
    Pending,

    [PgName("ACCEPTED")]
    Accepted,

    [PgName("REJECTED")]
    Rejected,

    [PgName("INVALIDATED")]
    Invalidated,
}
