#nullable enable

using NpgsqlTypes;

namespace Trigger.Models;

public enum OAuthClientKind
{
    [PgName("PREREGISTERED")]
    Preregistered,

    [PgName("CIMD")]
    Cimd,

    [PgName("DCR")]
    Dcr,
}
