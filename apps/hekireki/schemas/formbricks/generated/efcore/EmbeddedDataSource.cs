#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// Where an Embedded Data field's value comes from.
/// `computed` is set by survey logic, `ingested` arrives from a URL param or the SDK, and
/// `reserved` is auto-captured system metadata. Reserved fields are never stored as rows —
/// they are a static catalog in code — but the value exists so every layer shares one type.
/// </summary>
public enum EmbeddedDataSource
{
    [PgName("computed")]
    Computed,

    [PgName("ingested")]
    Ingested,

    [PgName("reserved")]
    Reserved,
}
