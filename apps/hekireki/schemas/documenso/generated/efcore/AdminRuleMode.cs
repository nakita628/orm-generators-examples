#nullable enable

using NpgsqlTypes;

namespace Documenso.Models;

public enum AdminRuleMode
{
    [PgName("ENFORCE")]
    Enforce,

    [PgName("MONITOR")]
    Monitor,

    [PgName("DISABLED")]
    Disabled,
}
