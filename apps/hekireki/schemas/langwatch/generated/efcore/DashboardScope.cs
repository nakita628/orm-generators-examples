#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// Who sees a dashboard, named as the prompt scope is, with one more level.
/// </summary>
public enum DashboardScope
{
    [PgName("PRIVATE")]
    Private,

    [PgName("PROJECT")]
    Project,

    [PgName("ORGANIZATION")]
    Organization,
}
