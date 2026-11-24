#nullable enable

using NpgsqlTypes;

namespace Langfuse.Models;

public enum ModelScope
{
    [PgName("LANGFUSE")]
    Langfuse,

    [PgName("ORGANIZATION")]
    Organization,

    [PgName("PROJECT")]
    Project,
}
