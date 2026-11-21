#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// The author's intent only. Scheduled, Active, "Paused with its survey" and Closed after the
/// end time are derived (`apps/web/modules/ee/creatives/lib/display-status.ts` and its SQL twin); nothing
/// writes this column when a schedule ends or a survey pauses.
/// </summary>
public enum CreativeStatus
{
    [PgName("draft")]
    Draft,

    [PgName("inProgress")]
    InProgress,

    [PgName("paused")]
    Paused,

    [PgName("completed")]
    Completed,
}
