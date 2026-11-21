#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum CustomFieldType
{
    [PgName("TEXT")]
    Text,

    [PgName("OPTION")]
    Option,

    [PgName("BOOLEAN")]
    Boolean,

    [PgName("DATE")]
    Date,

    [PgName("MULTILINE_TEXT")]
    MultilineText,

    [PgName("AMOUNT")]
    Amount,

    [PgName("NUMBER")]
    Number,
}
