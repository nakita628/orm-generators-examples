#nullable enable

namespace Formbricks.Models;

/// <summary>
/// Represents third-party service integrations.
/// Enables data flow between Formbricks and external services.
///
/// @property id - Unique identifier for the integration
/// @property type - The service being integrated (Google Sheets, Notion, etc.)
/// @property workspace - The workspace this integration belongs to
/// @property config - Service-specific configuration details
/// </summary>
public partial class Integration
{
    public string Id { get; set; } = null!;

    public IntegrationType Type { get; set; }

    /// <summary>
    /// [IntegrationConfig]
    /// </summary>
    public string Config { get; set; } = null!;

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;
}
