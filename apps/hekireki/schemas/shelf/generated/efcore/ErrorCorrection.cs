#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum ErrorCorrection
{
    [PgName("L")]
    L,

    [PgName("M")]
    M,

    [PgName("Q")]
    Q,

    [PgName("H")]
    H,
}
