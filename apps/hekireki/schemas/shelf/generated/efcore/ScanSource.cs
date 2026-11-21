#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// The surface that recorded a `Scan`.
/// </summary>
public enum ScanSource
{
    [PgName("QR_LINK")]
    QrLink,

    [PgName("WEB_SCANNER")]
    WebScanner,

    [PgName("WEB_DRAWER")]
    WebDrawer,

    [PgName("AUDIT")]
    Audit,

    [PgName("COMPANION")]
    Companion,

    [PgName("GPS_UPDATE")]
    GpsUpdate,
}
