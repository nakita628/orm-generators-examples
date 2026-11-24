#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
/// its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
/// a token is now bound to the resource it was approved for instead of to a static allow-list.
/// `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
/// What one user allows one MCP client to do (GLOSSARY.md "Grant", ADR-0041): its organization &amp; workspace
/// access and an access level per permission. One per OAuth connection, keyed by user and client rather
/// than by `oauthConsent` row, because the provider can write duplicate consent rows for one connection
/// and settles them later. Every MCP request reads it; without one the request is refused. The client's
/// OAuth scopes stay on the consent and cap it.
/// </summary>
public partial class McpGrant
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string ClientId { get; set; } = null!;

    public virtual OauthClient Client { get; set; } = null!;

    /// <summary>
    /// Which organizations and workspaces the grant reaches (ENG-7153): All access, only the organizations
    /// in `organizations` (and every workspace in them, later ones included), or only the workspaces in
    /// `workspaces`. Enforced in SpiceDB through `can()` (ADR-0042); a value no row backs reaches nothing.
    /// </summary>
    public McpGrantAccess Access { get; set; } = McpGrantAccess.All;

    /// <summary>
    /// The access level per permission of `apps/web/modules/mcp/permissions.ts`: `{ "surveys": "write" }`.
    /// A permission that is absent has No access.
    /// </summary>
    public string Permissions { get; set; } = "{}";

    /// <summary>
    /// The selections behind `organizations` and `workspaces` access (ENG-7153).
    /// </summary>
    public virtual ICollection<McpGrantOrganization> Organizations { get; set; } = new List<McpGrantOrganization>();

    public virtual ICollection<McpGrantWorkspace> Workspaces { get; set; } = new List<McpGrantWorkspace>();
}
