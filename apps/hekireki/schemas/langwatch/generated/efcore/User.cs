#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

public partial class User
{
    public string Id { get; set; } = null!;

    public string? Name { get; set; }

    public string? Email { get; set; }

    public bool EmailVerified { get; set; } = false;

    public bool SignupConfirmationPending { get; set; } = false;

    public string? PasskeySignupClaimHash { get; set; }

    public string? Image { get; set; }

    public bool PendingSsoSetup { get; set; } = false;

    public string? UserHashKey { get; set; }

    /// <summary>
    /// Owned by the two-factor plugin (D06), which is why it is a plain
    /// boolean rather than a state: the plugin flips it when a setup is
    /// confirmed and when it is turned off, and challenges every sign-in while
    /// it is true. The lifecycle - and the history - lives on MfaEnrollment.
    /// </summary>
    public bool TwoFactorEnabled { get; set; } = false;

    public virtual ICollection<Account> Accounts { get; set; } = new List<Account>();

    public virtual ICollection<AccountCredential> AccountCredentials { get; set; } = new List<AccountCredential>();

    public virtual ICollection<Session> Sessions { get; set; } = new List<Session>();

    public virtual ICollection<TwoFactor> TwoFactors { get; set; } = new List<TwoFactor>();

    public virtual ICollection<Passkey> Passkeys { get; set; } = new List<Passkey>();

    public virtual ICollection<TeamUser> TeamMemberships { get; set; } = new List<TeamUser>();

    public virtual ICollection<OrganizationUser> OrgMemberships { get; set; } = new List<OrganizationUser>();

    public virtual ICollection<ScimUserResource> ScimUserResources { get; set; } = new List<ScimUserResource>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? LastLoginAt { get; set; }

    public DateTime? DeactivatedAt { get; set; }

    public string? LastHomePath { get; set; }

    public DateTime? TracesExplorerTourDismissedAt { get; set; }

    /// <summary>
    /// How Langy should reach this person's code when a task needs a change
    /// (ADR-129). "github" means they chose the pull request path and asked to
    /// be remembered; NULL means ask. A shared local folder is never stored
    /// here, because a folder has to be shared again each time.
    /// Spec: specs/langy/langy-code-access.feature
    /// </summary>
    public string? LangyCodeAccessPreference { get; set; }

    /// <summary>
    /// When this person last said "not now" to being offered a passkey
    /// (ADR-120). On the account rather than in browser storage for two
    /// reasons: a per-browser record forgets on a new device, which is exactly
    /// where the offer should be MORE eager, and an interval nobody can enforce
    /// is not an interval. NULL means never asked.
    /// </summary>
    public DateTime? PasskeyNudgeDismissedAt { get; set; }

    /// <summary>
    /// The domains this person has said "no thanks" to being offered their
    /// colleagues on (D12). Per domain rather than one flag: somebody who
    /// changes jobs and verifies a second address should be offered their new
    /// employer's organization, not silenced forever by a decision they made
    /// about the old one. Empty means never dismissed anything.
    /// </summary>
    public List<string>? JoinOfferDismissedDomains { get; set; } = new List<string>();

    /// <summary>
    /// This person's answer about browser notifications, one entry per topic:
    /// { "langy": "enabled" | "declined" }. A topic that is absent was never
    /// answered. The browser's own permission is not stored here.
    /// Spec: specs/langy/langy-notifications.feature
    /// </summary>
    public string NotificationPreferences { get; set; } = "{}";

    /// <summary>
    /// When an operator marked this person's password as no longer trusted. A
    /// password sign-in opens no session while it is set; any write of a new
    /// password hash clears it. Spec: modules/auth/specs/forced-password-reset.feature
    /// </summary>
    public DateTime? PasswordResetRequiredAt { get; set; }

    public virtual ICollection<Annotation> Annotation { get; set; } = new List<Annotation>();

    public virtual ICollection<ShareLink> ShareLinks { get; set; } = new List<ShareLink>();

    public virtual ICollection<Workflow> Workflow { get; set; } = new List<Workflow>();

    public virtual ICollection<WorkflowVersion> WorkflowVersion { get; set; } = new List<WorkflowVersion>();

    public virtual ICollection<AnnotationQueueMembers> AnnotationQueues { get; set; } = new List<AnnotationQueueMembers>();

    public virtual ICollection<AnnotationQueueItem> AssignedQueueItems { get; set; } = new List<AnnotationQueueItem>();

    public virtual ICollection<AnnotationQueueItem> CreatedQueueItems { get; set; } = new List<AnnotationQueueItem>();

    public virtual ICollection<LlmPromptConfigVersion> LlmPromptConfigVersions { get; set; } = new List<LlmPromptConfigVersion>();

    public virtual ICollection<Scenario> Scenarios { get; set; } = new List<Scenario>();

    public virtual ICollection<OrganizationInvite> InviteRequests { get; set; } = new List<OrganizationInvite>();

    public virtual ICollection<ProjectSecret> SecretsCreated { get; set; } = new List<ProjectSecret>();

    public virtual ICollection<ProjectSecret> SecretsUpdated { get; set; } = new List<ProjectSecret>();

    public virtual ICollection<SavedView> SavedViews { get; set; } = new List<SavedView>();

    public virtual ICollection<PromptTagAssignment> VersionTagsCreated { get; set; } = new List<PromptTagAssignment>();

    public virtual ICollection<PromptTagAssignment> VersionTagsUpdated { get; set; } = new List<PromptTagAssignment>();

    public virtual ICollection<PromptTag> PromptTagsCreated { get; set; } = new List<PromptTag>();

    public virtual ICollection<PromptTag> PromptTagsUpdated { get; set; } = new List<PromptTag>();

    public virtual ICollection<GroupMembership> GroupMemberships { get; set; } = new List<GroupMembership>();

    public virtual ICollection<RoleBinding> RoleBindings { get; set; } = new List<RoleBinding>();

    public virtual ICollection<ScimExternalId> ScimExternalIds { get; set; } = new List<ScimExternalId>();

    public virtual ICollection<ApiKey> ApiKeys { get; set; } = new List<ApiKey>();

    public virtual ICollection<GatewayBudget> GatewayBudgetsCreated { get; set; } = new List<GatewayBudget>();

    public virtual ICollection<VirtualKey> PrincipalVirtualKeys { get; set; } = new List<VirtualKey>();

    public virtual ICollection<Team> PersonalTeams { get; set; } = new List<Team>();

    public virtual ICollection<Project> PersonalProjects { get; set; } = new List<Project>();

    public virtual ICollection<RoutingPolicy> RoutingPoliciesCreated { get; set; } = new List<RoutingPolicy>();

    public virtual ICollection<RoutingPolicy> RoutingPoliciesUpdated { get; set; } = new List<RoutingPolicy>();

    public virtual ICollection<GatewayGuardrail> GatewayGuardrailsCreated { get; set; } = new List<GatewayGuardrail>();

    public virtual ICollection<GatewayGuardrail> GatewayGuardrailsUpdated { get; set; } = new List<GatewayGuardrail>();

    public virtual ICollection<IngestionSource> IngestionSourcesCreated { get; set; } = new List<IngestionSource>();

    public virtual ICollection<AnomalyRule> AnomalyRulesCreated { get; set; } = new List<AnomalyRule>();

    public virtual ICollection<PinnedTrace> PinnedTraces { get; set; } = new List<PinnedTrace>();

    public virtual ICollection<TraceEditOverlay> TraceEditOverlaysCreated { get; set; } = new List<TraceEditOverlay>();

    public virtual ICollection<TraceEditOverlay> TraceEditOverlaysUpdated { get; set; } = new List<TraceEditOverlay>();
}
