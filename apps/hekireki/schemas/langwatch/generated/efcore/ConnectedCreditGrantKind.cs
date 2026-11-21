#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// Why a credit was granted. The amount is always prepaid, never promotional.
/// </summary>
public enum ConnectedCreditGrantKind
{
    [PgName("commit")]
    Commit,

    [PgName("added")]
    Added,

    [PgName("renewal")]
    Renewal,
}
