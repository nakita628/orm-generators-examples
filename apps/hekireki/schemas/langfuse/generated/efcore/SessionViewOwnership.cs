#nullable enable

using NpgsqlTypes;

namespace Langfuse.Models;

public enum SessionViewOwnership
{
    [PgName("PRIVATE")]
    Private,

    [PgName("PROJECT")]
    Project,
}
