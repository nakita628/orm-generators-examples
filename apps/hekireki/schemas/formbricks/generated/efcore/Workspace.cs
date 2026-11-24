#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Main grouping mechanism for resources in Formbricks.
/// Each organization can have multiple workspaces to separate different applications or products.
///
/// @property id - Unique identifier for the workspace
/// @property name - Display name of the workspace
/// @property organization - Reference to parent organization
/// @property surveys - Collection of surveys in this workspace
/// @property styling - Workspace-wide styling configuration
/// @property config - Workspace-specific configuration
/// @property recontactDays - Default recontact delay for surveys
/// @property placement - Default widget placement for in-app surveys
/// </summary>
public partial class Workspace
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public string? LegacyEnvironmentId { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// [Styling]
    /// </summary>
    public string Styling { get; set; } = "{\"allowStyleOverwrite\":true}";

    /// <summary>
    /// [WorkspaceConfig]
    /// </summary>
    public string Config { get; set; } = "{}";

    public int RecontactDays { get; set; } = 7;

    public bool LinkSurveyBranding { get; set; } = true;

    public bool InAppSurveyBranding { get; set; } = true;

    public WidgetPlacement Placement { get; set; } = WidgetPlacement.BottomRight;

    public bool ClickOutsideClose { get; set; } = true;

    public SurveyOverlay Overlay { get; set; } = SurveyOverlay.None;

    /// <summary>
    /// Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`.
    /// </summary>
    public string? OverlayColor { get; set; }

    public int? OverlayOpacity { get; set; }

    public virtual ICollection<Language> Languages { get; set; } = new List<Language>();

    public virtual ICollection<EmbeddedData> EmbeddedData { get; set; } = new List<EmbeddedData>();

    /// <summary>
    /// [Logo]
    /// </summary>
    public string? Logo { get; set; }

    public virtual ICollection<WorkspaceTeam> WorkspaceTeams { get; set; } = new List<WorkspaceTeam>();

    public bool AppSetupCompleted { get; set; } = false;

    public string? CustomHeadScripts { get; set; }

    /// <summary>
    /// [CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service.
    /// </summary>
    public string? CustomCss { get; set; }

    /// <summary>
    /// [CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
    /// submitting its source again through the current processor and authorization checks.
    /// </summary>
    public string? CustomCssPrevious { get; set; }

    public virtual ICollection<FeedbackDirectoryWorkspace> FeedbackDirectoryWorkspaces { get; set; } = new List<FeedbackDirectoryWorkspace>();

    public virtual ICollection<Chart> Charts { get; set; } = new List<Chart>();

    public virtual ICollection<Dashboard> Dashboards { get; set; } = new List<Dashboard>();

    public virtual ICollection<EmbeddedFieldLabel> EmbeddedFieldLabels { get; set; } = new List<EmbeddedFieldLabel>();

    public virtual ICollection<Creative> Creatives { get; set; } = new List<Creative>();

    public virtual ICollection<AttachmentExport> AttachmentExports { get; set; } = new List<AttachmentExport>();

    public virtual ICollection<Survey> Surveys { get; set; } = new List<Survey>();

    public virtual ICollection<Contact> Contacts { get; set; } = new List<Contact>();

    public virtual ICollection<ActionClass> ActionClasses { get; set; } = new List<ActionClass>();

    public virtual ICollection<ContactAttributeKey> ContactAttributeKeys { get; set; } = new List<ContactAttributeKey>();

    public virtual ICollection<Webhook> Webhooks { get; set; } = new List<Webhook>();

    public virtual ICollection<Tag> Tags { get; set; } = new List<Tag>();

    public virtual ICollection<Segment> Segments { get; set; } = new List<Segment>();

    public virtual ICollection<Integration> Integrations { get; set; } = new List<Integration>();

    public virtual ICollection<ApiKeyWorkspace> ApiKeyWorkspaces { get; set; } = new List<ApiKeyWorkspace>();

    public virtual ICollection<FeedbackSource> FeedbackSources { get; set; } = new List<FeedbackSource>();

    public virtual ICollection<Workflow> Workflows { get; set; } = new List<Workflow>();

    public virtual ICollection<WorkflowRun> WorkflowRuns { get; set; } = new List<WorkflowRun>();

    public virtual ICollection<McpGrantWorkspace> McpGrants { get; set; } = new List<McpGrantWorkspace>();
}
