#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// The operating system a companion-app push token was issued for.
/// </summary>
public enum DevicePlatform
{
    [PgName("IOS")]
    Ios,

    [PgName("ANDROID")]
    Android,
}
