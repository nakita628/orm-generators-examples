#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Stores third-party authentication account information.
/// Enables OAuth and other external authentication methods.
///
/// @property id - Unique identifier for the account
/// @property user - The Formbricks user who owns this account
/// @property provider - The authentication provider (GitHub, Google, etc.)
/// @property providerAccountId - User ID from the provider
/// @property access_token - OAuth access token
/// @property refresh_token - OAuth refresh token
/// </summary>
public partial class Account
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User? User { get; set; }

    public string UserId { get; set; } = null!;

    public string? Type { get; set; }

    public string Provider { get; set; } = null!;

    public string ProviderAccountId { get; set; } = null!;

    public string? AccessToken { get; set; }

    public string? RefreshToken { get; set; }

    public int? ExpiresAt { get; set; }

    public int? ExtExpiresIn { get; set; }

    public string? TokenType { get; set; }

    public string? Scope { get; set; }

    public string? IdToken { get; set; }

    public string? SessionState { get; set; }

    public string? Password { get; set; }

    public DateTime? AccessTokenExpiresAt { get; set; }

    public DateTime? RefreshTokenExpiresAt { get; set; }

    /// <summary>
    /// Better Auth 1.7.0 (ENG-2343): the account's identity namespace. 1.7.0 keyed accounts on
    /// (issuer, accountId) and filtered every lookup on it — findAccountByKey, findAccountOwnerByKey,
    /// findCredentialAccount and updatePassword — so on 1.7.0 a null here locks the user out.
    /// 1.7.3 dropped the field and 1.7.7 keys accounts on (provider, providerAccountId) again
    /// (ENG-3880). We keep writing it (see `account.create.before` in better-auth-hooks.ts) so a
    /// rollback to 1.7.0 still finds every row; drop it once that rollback is off the table.
    ///
    /// NULLABLE on purpose. A NOT NULL column would fail every INSERT from a 1.6 pod still serving
    /// during a rolling deploy. The migration backfills every existing row; see the migration SQL for
    /// the value derivation.
    /// </summary>
    public string? Issuer { get; set; }
}
