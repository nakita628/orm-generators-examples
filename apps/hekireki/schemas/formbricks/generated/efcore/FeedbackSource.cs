#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// A feedback source — connects external data into Unify Feedback for record creation.
/// Naming note: "FeedbackSource" is the Unify Feedback domain term; the word "Connector"
/// is reserved for the Workflows domain (see docs/development/standards/practices/naming-conventions.mdx).
///
/// @property id - Unique identifier for the feedback source
/// @property name - Display name for the feedback source
/// @property type - Type of feedback source (formbricks_survey, csv, …)
/// @property status - Current state of the feedback source (active, paused, error)
/// @property workspace - The workspace this feedback source belongs to
/// @property feedbackDirectory - Hub directory the feedback source writes into
/// @property formbricksMappings - Element mappings for Formbricks feedback sources
/// @property fieldMappings - Field mappings for other feedback source types
/// </summary>
public partial class FeedbackSource
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public FeedbackSourceType Type { get; set; }

    public FeedbackSourceStatus Status { get; set; } = FeedbackSourceStatus.Active;

    /// <summary>
    /// Defaults to completedOnly so a new source agrees with the live ingestion path, which only
    /// fires on responseFinished. Existing rows take the default; nothing is re-imported.
    /// </summary>
    public FeedbackSourceImportMode ImportMode { get; set; } = FeedbackSourceImportMode.CompletedOnly;

    public FeedbackSourceElementScope ElementScope { get; set; } = FeedbackSourceElementScope.Specific;

    public string WorkspaceId { get; set; } = null!;

    public virtual Workspace Workspace { get; set; } = null!;

    public string FeedbackDirectoryId { get; set; } = null!;

    public virtual FeedbackDirectory FeedbackDirectory { get; set; } = null!;

    public virtual ICollection<FeedbackSourceFormbricksMapping> FormbricksMappings { get; set; } = new List<FeedbackSourceFormbricksMapping>();

    public virtual ICollection<FeedbackSourceFieldMapping> FieldMappings { get; set; } = new List<FeedbackSourceFieldMapping>();

    public DateTime? LastSyncAt { get; set; }

    public string? CreatedBy { get; set; }

    public virtual User? Creator { get; set; }

    public virtual FeedbackDirectoryWorkspace FeedbackDirectoryWorkspace { get; set; } = null!;
}
