#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a user in the Formbricks system.
/// Central model for user authentication and profile management.
///
/// @property id - Unique identifier for the user
/// @property name - Display name of the user
/// @property email - User's email address
/// @property twoFactorEnabled - Whether 2FA is active
/// @property memberships - Organizations the user belongs to
/// @property notificationSettings - User's notification preferences
/// </summary>
public partial class User
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public string Email { get; set; } = null!;

    public bool EmailVerified { get; set; } = false;

    public string? TwoFactorSecret { get; set; }

    public bool TwoFactorEnabled { get; set; } = false;

    public string? BackupCodes { get; set; }

    public string? Password { get; set; }

    public IdentityProvider IdentityProvider { get; set; } = IdentityProvider.Email;

    public string? IdentityProviderAccountId { get; set; }

    public virtual ICollection<Membership> Memberships { get; set; } = new List<Membership>();

    public virtual ICollection<Account> Accounts { get; set; } = new List<Account>();

    public virtual ICollection<Session> Sessions { get; set; } = new List<Session>();

    public virtual TwoFactor? TwoFactor { get; set; }

    public virtual ICollection<OauthClient> OauthClients { get; set; } = new List<OauthClient>();

    public virtual ICollection<OauthAccessToken> OauthAccessTokens { get; set; } = new List<OauthAccessToken>();

    public virtual ICollection<OauthRefreshToken> OauthRefreshTokens { get; set; } = new List<OauthRefreshToken>();

    public virtual ICollection<OauthConsent> OauthConsents { get; set; } = new List<OauthConsent>();

    public virtual PasswordResetToken? PasswordResetToken { get; set; }

    public string? GroupId { get; set; }

    public virtual ICollection<Invite> InvitesCreated { get; set; } = new List<Invite>();

    public virtual ICollection<Invite> InvitesAccepted { get; set; } = new List<Invite>();

    /// <summary>
    /// [UserNotificationSettings]
    /// </summary>
    public string NotificationSettings { get; set; } = "{}";

    /// <summary>
    /// [Locale]
    /// </summary>
    public string Locale { get; set; } = "en-US";

    public virtual ICollection<Survey> Surveys { get; set; } = new List<Survey>();

    public virtual ICollection<Survey> OwnedSurveys { get; set; } = new List<Survey>();

    public virtual ICollection<Chart> Charts { get; set; } = new List<Chart>();

    public virtual ICollection<Dashboard> Dashboards { get; set; } = new List<Dashboard>();

    public virtual ICollection<Creative> Creatives { get; set; } = new List<Creative>();

    public virtual ICollection<FeedbackSource> FeedbackSources { get; set; } = new List<FeedbackSource>();

    public virtual ICollection<TeamUser> TeamUsers { get; set; } = new List<TeamUser>();

    public virtual ICollection<Workflow> Workflows { get; set; } = new List<Workflow>();

    public virtual ICollection<WorkflowVersion> WorkflowVersionsPublished { get; set; } = new List<WorkflowVersion>();

    public virtual ICollection<RetentionPolicy> RetentionPoliciesUpdated { get; set; } = new List<RetentionPolicy>();

    public virtual ICollection<RetentionExemption> ExemptionsCreated { get; set; } = new List<RetentionExemption>();

    public virtual ICollection<RetentionExemption> ExemptionsRevoked { get; set; } = new List<RetentionExemption>();

    public virtual ICollection<RetentionNotice> RetentionNotices { get; set; } = new List<RetentionNotice>();

    public virtual ICollection<AttachmentExport> AttachmentExports { get; set; } = new List<AttachmentExport>();

    public DateTime? LastLoginAt { get; set; }

    public bool IsActive { get; set; } = true;

    /// <summary>
    /// When an admin last reactivated this user. Restarts the members retention clock (ENG-3713)
    /// without touching `lastLoginAt`, which the member list, usage counts and APIs report as a real
    /// sign-in.
    /// </summary>
    public DateTime? ReactivatedAt { get; set; }

    /// <summary>
    /// When the user last completed a sign-in (`lastActivePlugin`, after any second factor) or renewed a
    /// session (Better Auth `databaseHooks.session.update`), on the database clock and never moved
    /// backwards. The members retention clock reads it: sessions are
    /// deleted on sign-out and expiry, so their own timestamps can't serve as a record of activity.
    /// </summary>
    public DateTime? LastActiveAt { get; set; }

    public BootstrapAdminMarker? IsBootstrapAdmin { get; set; }
}
