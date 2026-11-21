#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum AuditAssetStatus
{
    [PgName("PENDING")]
    Pending,

    [PgName("FOUND")]
    Found,

    [PgName("MISSING")]
    Missing,

    [PgName("UNEXPECTED")]
    Unexpected,
}
