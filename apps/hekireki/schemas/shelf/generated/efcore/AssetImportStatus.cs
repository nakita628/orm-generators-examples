#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// Lifecycle of a background asset import.
/// </summary>
public enum AssetImportStatus
{
    [PgName("DRAFT")]
    Draft,

    [PgName("QUEUED")]
    Queued,

    [PgName("RUNNING")]
    Running,

    [PgName("COMPLETED")]
    Completed,

    [PgName("FAILED")]
    Failed,
}
