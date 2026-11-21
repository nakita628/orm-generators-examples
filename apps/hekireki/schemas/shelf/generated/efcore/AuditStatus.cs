#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum AuditStatus
{
    [PgName("PENDING")]
    Pending,

    [PgName("ACTIVE")]
    Active,

    [PgName("COMPLETED")]
    Completed,

    [PgName("CANCELLED")]
    Cancelled,

    [PgName("ARCHIVED")]
    Archived,
}
