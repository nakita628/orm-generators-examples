#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// Organization signup intent (ADR-038): which product the org came for.
/// </summary>
public enum OrganizationIntent
{
    [PgName("AGENT_GOVERNANCE")]
    AgentGovernance,

    [PgName("LLM_OPS")]
    LlmOps,
}
