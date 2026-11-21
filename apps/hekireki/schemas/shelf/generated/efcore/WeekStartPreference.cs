#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum WeekStartPreference
{
    [PgName("MONDAY")]
    Monday,

    [PgName("SUNDAY")]
    Sunday,

    [PgName("SATURDAY")]
    Saturday,
}
