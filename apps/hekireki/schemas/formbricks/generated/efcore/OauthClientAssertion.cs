#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
/// assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
/// table, so omitting it risks an adapter error on a code path we do not control.
/// </summary>
public partial class OauthClientAssertion
{
    public string Id { get; set; } = null!;

    public DateTime ExpiresAt { get; set; }
}
