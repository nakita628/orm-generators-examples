#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum DateFormatPreference
{
    [PgName("DD_MM_YYYY")]
    DdMmYyyy,

    [PgName("MM_DD_YYYY")]
    MmDdYyyy,

    [PgName("YYYY_MM_DD")]
    YyyyMmDd,

    [PgName("MMM_DD_YYYY")]
    MmmDdYyyy,

    [PgName("DD_MMM_YYYY")]
    DdMmmYyyy,
}
