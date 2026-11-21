#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// What makes a creative appear. A creative left with no action never shows; it never falls back to
/// every page view.
/// </summary>
public enum CreativeShowOn
{
    [PgName("pageView")]
    PageView,

    [PgName("action")]
    Action,
}
