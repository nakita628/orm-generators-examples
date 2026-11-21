#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum BarcodeType
{
    [PgName("Code128")]
    Code128,

    [PgName("Code39")]
    Code39,

    [PgName("DataMatrix")]
    DataMatrix,

    [PgName("ExternalQR")]
    ExternalQR,

    [PgName("EAN13")]
    Ean13,
}
