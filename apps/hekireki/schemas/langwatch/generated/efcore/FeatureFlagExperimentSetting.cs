#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One experiment setting per (flag, subject). A USER row is that person's
/// own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
/// for that exact scope. No row means inherit.
/// </summary>
public partial class FeatureFlagExperimentSetting
{
    public string FlagKey { get; set; } = null!;

    public FeatureFlagExperimentSubjectType SubjectType { get; set; }

    public string SubjectId { get; set; } = null!;

    public bool Enabled { get; set; }

    public string? ChangedByUserId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
