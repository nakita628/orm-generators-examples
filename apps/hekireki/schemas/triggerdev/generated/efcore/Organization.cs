#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class Organization
{
    public string Id { get; set; } = null!;

    public string Slug { get; set; } = null!;

    public string Title { get; set; } = null!;

    public int MaximumExecutionTimePerRunInMs { get; set; } = 900000;

    public int MaximumConcurrencyLimit { get; set; } = 10;

    /// <summary>
    /// This is deprecated and will be removed in the future
    /// </summary>
    public int MaximumSchedulesLimit { get; set; } = 5;

    public int? MaximumDevQueueSize { get; set; }

    public int? MaximumDeployedQueueSize { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? DeletedAt { get; set; }

    public string? CompanySize { get; set; }

    public string? OnboardingData { get; set; }

    public string? Avatar { get; set; }

    public bool RunsEnabled { get; set; } = true;

    public bool IsActivated { get; set; } = false;

    /// <summary>
    /// @deprecated
    /// </summary>
    public bool V2Enabled { get; set; } = false;

    /// <summary>
    /// @deprecated
    /// </summary>
    public bool V2MarqsEnabled { get; set; } = false;

    /// <summary>
    /// @deprecated
    /// </summary>
    public bool HasRequestedV3 { get; set; } = false;

    public virtual ICollection<RuntimeEnvironment> Environments { get; set; } = new List<RuntimeEnvironment>();

    public string? ApiRateLimiterConfig { get; set; }

    public string? RealtimeRateLimiterConfig { get; set; }

    public string? BatchRateLimitConfig { get; set; }

    public string? BatchQueueConcurrencyConfig { get; set; }

    /// <summary>
    /// Webhook limits from the org's plan or the back office: { maxWaitersPerEnvironment, maxWaitersPerEndpoint, concurrency }. Null uses the env defaults.
    /// </summary>
    public string? WebhookLimitsConfig { get; set; }

    public string? FeatureFlags { get; set; }

    public int MaximumProjectCount { get; set; } = 25;

    /// <summary>
    /// Maximum session lifetime in seconds for members of this organization.
    /// When set, caps each member's User.sessionDuration. The most restrictive cap across
    /// all of a user's orgs (excluding nulls) wins.
    /// </summary>
    public int? MaxSessionDuration { get; set; }

    public virtual ICollection<Project> Projects { get; set; } = new List<Project>();

    public virtual ICollection<OrgMember> Members { get; set; } = new List<OrgMember>();

    public virtual ICollection<OrgMemberInvite> Invites { get; set; } = new List<OrgMemberInvite>();

    public virtual ICollection<OrganizationIntegration> OrganizationIntegrations { get; set; } = new List<OrganizationIntegration>();

    public virtual ICollection<OrganizationAccessToken> OrganizationAccessTokens { get; set; } = new List<OrganizationAccessToken>();

    public virtual ICollection<WorkerInstanceGroup> WorkerGroups { get; set; } = new List<WorkerInstanceGroup>();

    public virtual ICollection<WorkerInstance> WorkerInstances { get; set; } = new List<WorkerInstance>();

    public virtual ICollection<GithubAppInstallation> GithubAppInstallations { get; set; } = new List<GithubAppInstallation>();

    public virtual ICollection<CustomerQuery> CustomerQueries { get; set; } = new List<CustomerQuery>();

    public virtual ICollection<MetricsDashboard> MetricsDashboards { get; set; } = new List<MetricsDashboard>();

    public virtual ICollection<Prompt> Prompts { get; set; } = new List<Prompt>();

    public virtual ICollection<PlatformNotification> PlatformNotifications { get; set; } = new List<PlatformNotification>();

    public virtual ICollection<ErrorGroupState> ErrorGroupStates { get; set; } = new List<ErrorGroupState>();

    /// <summary>
    /// S2 basin that holds this org's realtime streams. Null until the
    /// per-org basin has been provisioned (OSS / s2-lite installs leave
    /// it null forever; reads fall back to the global basin env var).
    /// Set once at provisioning time; retention is reconfigured in-place
    /// when the org's plan changes.
    /// </summary>
    public string? StreamBasinName { get; set; }

    public SupportAccessMode SupportAccessMode { get; set; } = SupportAccessMode.Allow;

    public DateTime? SupportAccessModeUpdatedAt { get; set; }

    public string? SupportAccessModeUpdatedById { get; set; }

    public virtual ICollection<SupportAccessRequest> SupportAccessRequests { get; set; } = new List<SupportAccessRequest>();
}
