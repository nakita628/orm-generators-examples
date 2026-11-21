#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Per-(organization, tool) override for which `langwatch &lt;tool&gt;` paths the
/// CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
/// `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
/// row resolves to the hardcoded defaults (both true for claude/codex/gemini/
/// opencode; cursor allows the gateway path only), so the policy is purely
/// additive: disabling a path is an explicit admin choice. The CLI caches the
/// resolved map at login and gates path selection on it.
/// See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
/// </summary>
public partial class PlatformToolPolicy
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string ToolSlug { get; set; } = null!;

    public bool AllowVk { get; set; } = true;

    public bool AllowOtelDirect { get; set; } = true;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
