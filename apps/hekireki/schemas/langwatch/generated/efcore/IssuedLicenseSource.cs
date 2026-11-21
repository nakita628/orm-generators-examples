#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// Where a license in the registry came from.
/// </summary>
public enum IssuedLicenseSource
{
    [PgName("BACKOFFICE")]
    Backoffice,

    [PgName("PURCHASE")]
    Purchase,

    [PgName("SCRIPT")]
    Script,

    [PgName("LEGACY_IMPORT")]
    LegacyImport,
}
