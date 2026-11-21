#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// How the last fetch of a link's page went.
/// </summary>
public enum LinkPreviewOutcome
{
    [PgName("TITLE")]
    Title,

    [PgName("NO_TITLE")]
    NoTitle,

    [PgName("SIGN_IN")]
    SignIn,

    [PgName("BLOCKED")]
    Blocked,

    [PgName("UNREACHABLE")]
    Unreachable,
}
