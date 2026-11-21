#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// The declared type of an Embedded Data field's value.
/// </summary>
public enum EmbeddedDataType
{
    [PgName("string")]
    String,

    [PgName("number")]
    Number,

    [PgName("boolean")]
    Boolean,

    [PgName("date")]
    Date,
}
