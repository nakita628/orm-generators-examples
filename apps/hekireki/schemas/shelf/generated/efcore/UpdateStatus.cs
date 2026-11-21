#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum UpdateStatus
{
    [PgName("DRAFT")]
    Draft,

    [PgName("PUBLISHED")]
    Published,
}
