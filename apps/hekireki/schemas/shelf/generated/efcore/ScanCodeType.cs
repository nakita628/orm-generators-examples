#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// The kind of code a `Scan` recorded.
/// </summary>
public enum ScanCodeType
{
    [PgName("QR")]
    Qr,

    [PgName("BARCODE")]
    Barcode,

    [PgName("SAM_ID")]
    SamId,
}
