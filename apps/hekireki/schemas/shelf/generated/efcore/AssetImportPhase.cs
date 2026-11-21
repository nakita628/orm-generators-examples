#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// Which step of the background import runs next (or failed).
/// </summary>
public enum AssetImportPhase
{
    [PgName("PREPARE")]
    Prepare,

    [PgName("ROWS")]
    Rows,

    [PgName("FINALIZE")]
    Finalize,
}
