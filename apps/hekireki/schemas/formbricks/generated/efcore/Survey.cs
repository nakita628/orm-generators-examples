#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a complete survey configuration including questions, styling, and display rules.
/// Core model for the survey functionality in Formbricks.
///
/// @property id - Unique identifier for the survey
/// @property name - Display name of the survey
/// @property type - Survey delivery method (link, web, website, app)
/// @property status - Current state of the survey (draft, active, completed, etc)
/// @property workspace - The workspace this survey belongs to
/// @property questions - JSON array containing survey questions configuration
/// @property displayOption - Rules for how often the survey can be shown
/// @property triggers - Actions that can trigger this survey
/// @property attributeFilters - Rules for targeting specific contacts
/// @property customHeadScripts - Survey-specific custom HTML scripts (self-hosted only)
/// @property customHeadScriptsMode - "add" (merge with workspace) or "replace" (override workspace)
/// </summary>
public partial class Survey
{
    /// <summary>
    /// [SurveyHiddenFields]
    /// </summary>
    public string HiddenFields { get; set; } = "{\"enabled\": false}";

    /// <summary>
    /// [SurveyVariables]
    /// </summary>
    public string Variables { get; set; } = "[]";

    public bool OwnershipInitialized { get; set; } = false;

    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public string? RedirectUrl { get; set; }

    public SurveyType Type { get; set; } = SurveyType.App;

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual User? Creator { get; set; }

    public string? CreatedBy { get; set; }

    /// <summary>
    /// ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
    /// the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
    /// </summary>
    public SurveyVisibility Visibility { get; set; } = SurveyVisibility.Workspace;

    public virtual User? Owner { get; set; }

    public string? OwnerId { get; set; }

    /// <summary>
    /// Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
    /// with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
    /// projection, decided from PostgreSQL facts until the projector acknowledges it.
    /// </summary>
    public int VisibilityVersion { get; set; } = 0;

    /// <summary>
    /// The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was.
    /// </summary>
    public int VisibilityProjectedVersion { get; set; } = 0;

    /// <summary>
    /// "visibilityVersion" &lt;&gt; "visibilityProjectedVersion", kept by the `survey_visibility_pending`
    /// trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
    /// in create/update data.
    /// </summary>
    public bool VisibilityPending { get; set; } = false;

    public DateTime? VisibilityChangedAt { get; set; }

    public string? VisibilityChangedById { get; set; }

    public SurveyStatus Status { get; set; } = SurveyStatus.Draft;

    /// <summary>
    /// [SurveyWelcomeCard]
    /// </summary>
    public string WelcomeCard { get; set; } = "{\"enabled\": false}";

    /// <summary>
    /// [SurveyQuestions]
    /// </summary>
    public string Questions { get; set; } = "[]";

    /// <summary>
    /// [SurveyBlocks]
    /// </summary>
    public List<string>? Blocks { get; set; } = new List<string>();

    /// <summary>
    /// [SurveyEnding]
    /// </summary>
    public List<string>? Endings { get; set; } = new List<string>();

    public virtual ICollection<Response> Responses { get; set; } = new List<Response>();

    public virtual ICollection<SurveyQuota> Quotas { get; set; } = new List<SurveyQuota>();

    public DisplayOptions DisplayOption { get; set; } = DisplayOptions.DisplayOnce;

    public int? RecontactDays { get; set; }

    public int? DisplayLimit { get; set; }

    public virtual ICollection<SurveyTrigger> Triggers { get; set; } = new List<SurveyTrigger>();

    /// <summary>
    /// [SurveyInlineTriggers]
    /// </summary>
    public string? InlineTriggers { get; set; }

    public virtual ICollection<SurveyAttributeFilter> AttributeFilters { get; set; } = new List<SurveyAttributeFilter>();

    public virtual ICollection<Display> Displays { get; set; } = new List<Display>();

    public int? AutoClose { get; set; }

    public bool KeepEndingCardOpen { get; set; } = false;

    public int? AutoComplete { get; set; }

    public int Delay { get; set; } = 0;

    public DateTime? PublishOn { get; set; }

    public DateTime? CloseOn { get; set; }

    public DateTime? ArchivedAt { get; set; }

    /// <summary>
    /// [SurveyClosedMessage]
    /// </summary>
    public string? SurveyClosedMessage { get; set; }

    public string? SegmentId { get; set; }

    public virtual Segment? Segment { get; set; }

    /// <summary>
    /// [SurveyWorkspaceOverwrites]
    /// </summary>
    public string? WorkspaceOverwrites { get; set; }

    /// <summary>
    /// [SurveyStyling]
    /// </summary>
    public string? Styling { get; set; }

    /// <summary>
    /// [SurveySingleUse]
    /// </summary>
    public string? SingleUse { get; set; } = "{\"enabled\": false, \"isEncrypted\": true}";

    public bool IsVerifyEmailEnabled { get; set; } = false;

    public bool IsBackButtonHidden { get; set; } = false;

    public bool IsAutoProgressingEnabled { get; set; } = false;

    public bool IsCaptureIpEnabled { get; set; } = false;

    public bool IsAnonymizeResponsesEnabled { get; set; } = false;

    public string? Pin { get; set; }

    public decimal? DisplayPercentage { get; set; }

    public virtual ICollection<SurveyLanguage> Languages { get; set; } = new List<SurveyLanguage>();

    public virtual ICollection<EmbeddedData> EmbeddedData { get; set; } = new List<EmbeddedData>();

    public virtual ICollection<SurveyEmbeddedData> EmbeddedDataLinks { get; set; } = new List<SurveyEmbeddedData>();

    public bool? ShowLanguageSwitch { get; set; }

    /// <summary>
    /// Open the survey in the respondent's browser language when it matches an enabled survey
    /// language. Null on surveys that predate the setting, which reads as off.
    /// </summary>
    public bool? AutoSelectLanguage { get; set; }

    public virtual ICollection<SurveyFollowUp> FollowUps { get; set; } = new List<SurveyFollowUp>();

    /// <summary>
    /// [SurveyRecaptcha]
    /// </summary>
    public string? Recaptcha { get; set; } = "{\"enabled\": false, \"threshold\":0.1}";

    /// <summary>
    /// [SurveyLinkMetadata]
    /// </summary>
    public string Metadata { get; set; } = "{}";

    public virtual ICollection<FeedbackSourceFormbricksMapping> FeedbackSourceMappings { get; set; } = new List<FeedbackSourceFormbricksMapping>();

    public virtual ICollection<RetentionExemption> RetentionExemptions { get; set; } = new List<RetentionExemption>();

    public virtual ICollection<RetentionNotice> RetentionNotices { get; set; } = new List<RetentionNotice>();

    public virtual RetentionSurveyRestore? RetentionRestore { get; set; }

    public string? Slug { get; set; }

    public string? CustomHeadScripts { get; set; }

    public SurveyScriptMode? CustomHeadScriptsMode { get; set; } = SurveyScriptMode.Add;

    /// <summary>
    /// [CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
    /// by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
    /// </summary>
    public string? CustomCss { get; set; }

    /// <summary>
    /// Creatives whose Destination is this survey (ENG-3993).
    /// </summary>
    public virtual ICollection<Creative> Creatives { get; set; } = new List<Creative>();

    /// <summary>
    /// Background attachment ZIP exports of this survey (ENG-6862).
    /// </summary>
    public virtual ICollection<AttachmentExport> AttachmentExports { get; set; } = new List<AttachmentExport>();

    /// <summary>
    /// [SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
    /// (ENG-3873). null = no conditions, which is what every existing survey keeps.
    /// </summary>
    public string? ShowConditions { get; set; }

    /// <summary>
    /// [SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
    /// survey on load. null = none, which is what every existing survey keeps.
    /// </summary>
    public string? StartLogic { get; set; }
}
