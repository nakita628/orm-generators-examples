#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class User
{
    public string Id { get; set; } = null!;

    public string Email { get; set; } = null!;

    /// <summary>
    /// Address awaiting confirmation; becomes `email` only after a link sent to it is opened
    /// </summary>
    public string? PendingEmail { get; set; }

    public AuthenticationMethod AuthenticationMethod { get; set; }

    public string? AuthenticationProfile { get; set; }

    public string? AuthenticationExtraParams { get; set; }

    public string? AuthIdentifier { get; set; }

    public string? DisplayName { get; set; }

    public string? Name { get; set; }

    public string? AvatarUrl { get; set; }

    public bool Admin { get; set; } = false;

    /// <summary>
    /// Preferences for the dashboard
    /// </summary>
    public string? DashboardPreferences { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// @deprecated
    /// </summary>
    public bool IsOnCloudWaitlist { get; set; } = false;

    /// <summary>
    /// @deprecated
    /// </summary>
    public bool FeatureCloud { get; set; } = false;

    /// <summary>
    /// @deprecated
    /// </summary>
    public bool IsOnHostedRepoWaitlist { get; set; } = false;

    public bool MarketingEmails { get; set; } = true;

    public bool ConfirmedBasicDetails { get; set; } = false;

    public string? ReferralSource { get; set; }

    public string? OnboardingData { get; set; }

    public virtual ICollection<OrgMember> OrgMemberships { get; set; } = new List<OrgMember>();

    public virtual ICollection<OrgMemberInvite> SentInvites { get; set; } = new List<OrgMemberInvite>();

    public DateTime? MfaEnabledAt { get; set; }

    public virtual SecretReference? MfaSecretReference { get; set; }

    public string? MfaSecretReferenceId { get; set; }

    /// <summary>
    /// Hash of the last used code to prevent replay attacks
    /// </summary>
    public string? MfaLastUsedCode { get; set; }

    /// <summary>
    /// Maximum session lifetime in seconds for this user. Default = 1 year (Gregorian).
    /// May be further restricted by Organization.maxSessionDuration on any of the user's orgs.
    /// </summary>
    public int SessionDuration { get; set; } = 31556952;

    /// <summary>
    /// Absolute timestamp at which the user's current session must end. Written
    /// at login (and any time the effective duration is recomputed) and shortened
    /// when an admin lowers an org cap. Read on requireUser / getUser; null is
    /// treated as "no enforced deadline" (legacy sessions from before this
    /// shipped). Read-on-demand has zero per-request DB cost — the value
    /// piggybacks on the User fetch that requireUser already does.
    /// </summary>
    public DateTime? NextSessionEnd { get; set; }

    public virtual InvitationCode? InvitationCode { get; set; }

    public string? InvitationCodeId { get; set; }

    public virtual ICollection<PersonalAccessToken> PersonalAccessTokens { get; set; } = new List<PersonalAccessToken>();

    public virtual ICollection<ApiKey> CreatedApiKeys { get; set; } = new List<ApiKey>();

    public virtual ICollection<WorkerDeployment> Deployments { get; set; } = new List<WorkerDeployment>();

    public virtual ICollection<MfaBackupCode> BackupCodes { get; set; } = new List<MfaBackupCode>();

    public virtual ICollection<BulkActionGroup> BulkActions { get; set; } = new List<BulkActionGroup>();

    public virtual ICollection<ImpersonationAuditLog> ImpersonationsPerformed { get; set; } = new List<ImpersonationAuditLog>();

    public virtual ICollection<ImpersonationAuditLog> ImpersonationsReceived { get; set; } = new List<ImpersonationAuditLog>();

    public virtual ICollection<SupportAccessRequest> SupportAccessRequestsMade { get; set; } = new List<SupportAccessRequest>();

    public virtual ICollection<SupportAccessRequest> SupportAccessRequestsApproved { get; set; } = new List<SupportAccessRequest>();

    public virtual ICollection<CustomerQuery> CustomerQueries { get; set; } = new List<CustomerQuery>();

    public virtual ICollection<MetricsDashboard> MetricsDashboards { get; set; } = new List<MetricsDashboard>();

    public virtual ICollection<PlatformNotification> PlatformNotifications { get; set; } = new List<PlatformNotification>();

    public virtual ICollection<PlatformNotificationInteraction> PlatformNotificationInteractions { get; set; } = new List<PlatformNotificationInteraction>();
}
