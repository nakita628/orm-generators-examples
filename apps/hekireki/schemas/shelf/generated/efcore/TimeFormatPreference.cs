#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum TimeFormatPreference
{
    [PgName("H12")]
    H12,

    [PgName("H24")]
    H24,
}
