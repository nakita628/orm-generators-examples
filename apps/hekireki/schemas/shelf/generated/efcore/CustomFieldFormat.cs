#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// The format of a TEXT custom field. Unlike the type, it can change after the
/// field exists.
/// </summary>
public enum CustomFieldFormat
{
    [PgName("PLAIN")]
    Plain,

    [PgName("LINK")]
    Link,
}
