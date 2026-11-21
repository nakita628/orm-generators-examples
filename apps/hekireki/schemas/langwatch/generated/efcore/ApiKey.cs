#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

public partial class ApiKey
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public string LookupId { get; set; } = null!;

    public string HashedSecret { get; set; } = null!;

    public string PermissionMode { get; set; } = "all";

    public string? UserId { get; set; }

    public virtual User? User { get; set; }

    public string? CreatedByUserId { get; set; }

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public virtual ICollection<RoleBinding> RoleBindings { get; set; } = new List<RoleBinding>();

    public DateTime? ExpiresAt { get; set; }

    public DateTime? LastUsedAt { get; set; }

    public DateTime? RevokedAt { get; set; }

    /// <summary>
    /// Why the key was revoked, set with `revokedAt`: "user" when a person
    /// revoked it (the API-keys page, the REST API, the devices tab), "rotation"
    /// when a mint replaced it, "session" when the CLI login key it was parented
    /// to was revoked, "expired" when that session ran out, "offboarded" when the
    /// person's membership of the organization ended. "cap" survives on rows the
    /// retired per-tool cap revoked. The CLI reads it to decide whether a dead
    /// ingest key may be re-minted: a key a person revoked stays dead.
    /// NULL for live keys and for keys revoked before the cause was recorded.
    /// </summary>
    public string? RevocationCause { get; set; }

    /// <summary>
    /// The CLI login key of the device session that minted this ingestion key.
    /// Set by the CLI's personal mint, NULL for every other key. Revoking the
    /// parent (logout, the devices tab, a re-login from the same device, session
    /// expiry) revokes the children with it. Column ref only, no FK and no
    /// index: lookups filter by the indexed `userId` first and match the parent
    /// in memory over a person's few live ingest keys.
    /// </summary>
    public string? ParentApiKeyId { get; set; }

    /// <summary>
    /// Non-null marks this key as an "ingestion key": a project-scoped,
    /// ingest-only ApiKey the unified `langwatch &lt;tool&gt;` CLI (and template
    /// installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
    /// opencode / claude_cowork) is stamped as `langwatch.source` provenance on
    /// every span the key authorizes. NULL for normal API / service keys.
    /// </summary>
    public string? IngestSourceType { get; set; }

    /// <summary>
    /// Set only for ingestion keys derived from an admin IngestionTemplate
    /// (e.g. claude_cowork). Drives receiver OTTL application + the
    /// `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
    /// (which have no template) and for normal keys. Column ref only (no FK,
    /// per the app-layer-integrity convention); resolved by the ingest-key service.
    /// </summary>
    public string? IngestionTemplateId { get; set; }

    /// <summary>
    /// Human label of the CLI device session that minted this ingestion key
    /// (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
    /// flow's `client_info`. Denormalized so the org-wide API-keys settings page
    /// can show which device a key came from without resolving another user's
    /// Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
    /// a CLI old enough to predate device metadata.
    /// </summary>
    public string? CreatedByDeviceLabel { get; set; }

    /// <summary>
    /// True for a key LangWatch minted for itself (a run, the gateway's trace
    /// export, a Langy session). Hides it from listings, refuses customer edits
    /// and revokes, and lets an ownerless run key act as the system. Never set
    /// by a customer create; rows before the column stay false.
    /// </summary>
    public bool IsSystemManaged { get; set; } = false;

    /// <summary>
    /// The browser session that minted a docs session key (ADR-185), so signing
    /// out of that session revokes it. NULL for every other key. Unindexed:
    /// sign-out filters by the indexed `userId` first.
    /// </summary>
    public string? CreatedByBrowserSessionId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
