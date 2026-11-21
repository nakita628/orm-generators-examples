#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// Who can use a Slack connection (ADR-093 §5a): every project in the
/// organization, or one project. Chosen per connection.
/// </summary>
public enum SlackIntegrationScopeType
{
    [PgName("ORGANIZATION")]
    Organization,

    [PgName("PROJECT")]
    Project,
}
