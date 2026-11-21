#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum ReminderRecurrenceUnit
{
    [PgName("DAY")]
    Day,

    [PgName("WEEK")]
    Week,

    [PgName("MONTH")]
    Month,

    [PgName("YEAR")]
    Year,
}
