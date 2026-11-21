defmodule Formbricks.User do
  @moduledoc """
  Represents a user in the Formbricks system.
  Central model for user authentication and profile management.

  @property id - Unique identifier for the user
  @property name - Display name of the user
  @property email - User's email address
  @property twoFactorEnabled - Whether 2FA is active
  @property memberships - Organizations the user belongs to
  @property notificationSettings - User's notification preferences
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          email: String.t(),
          email_verified: boolean(),
          two_factor_secret: String.t() | nil,
          two_factor_enabled: boolean(),
          backup_codes: String.t() | nil,
          password: String.t() | nil,
          identity_provider: atom(),
          identity_provider_account_id: String.t() | nil,
          group_id: String.t() | nil,
          notification_settings: map(),
          locale: String.t(),
          last_login_at: DateTime.t() | nil,
          is_active: boolean(),
          reactivated_at: DateTime.t() | nil,
          last_active_at: DateTime.t() | nil,
          is_bootstrap_admin: atom() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          two_factor: Formbricks.TwoFactor.t() | nil,
          password_reset_token: Formbricks.PasswordResetToken.t() | nil,
          memberships: [Formbricks.Membership.t()],
          accounts: [Formbricks.Account.t()],
          sessions: [Formbricks.Session.t()],
          oauth_clients: [Formbricks.OauthClient.t()],
          oauth_access_tokens: [Formbricks.OauthAccessToken.t()],
          oauth_refresh_tokens: [Formbricks.OauthRefreshToken.t()],
          oauth_consents: [Formbricks.OauthConsent.t()],
          invites_created: [Formbricks.Invite.t()],
          invites_accepted: [Formbricks.Invite.t()],
          surveys: [Formbricks.Survey.t()],
          owned_surveys: [Formbricks.Survey.t()],
          charts: [Formbricks.Chart.t()],
          dashboards: [Formbricks.Dashboard.t()],
          creatives: [Formbricks.Creative.t()],
          feedback_sources: [Formbricks.FeedbackSource.t()],
          team_users: [Formbricks.TeamUser.t()],
          workflows: [Formbricks.Workflow.t()],
          workflow_versions_published: [Formbricks.WorkflowVersion.t()],
          retention_policies_updated: [Formbricks.RetentionPolicy.t()],
          exemptions_created: [Formbricks.RetentionExemption.t()],
          exemptions_revoked: [Formbricks.RetentionExemption.t()],
          retention_notices: [Formbricks.RetentionNotice.t()],
          attachment_exports: [Formbricks.AttachmentExport.t()]
        }

  schema "User" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:email, :string)
    field(:email_verified, :boolean, default: false)
    field(:two_factor_secret, :string, source: :twoFactorSecret)
    field(:two_factor_enabled, :boolean, default: false, source: :twoFactorEnabled)
    field(:backup_codes, :string, source: :backupCodes)
    field(:password, :string)
    field(:identity_provider, Ecto.Enum, values: [:email, :github, :google, :azuread, :openid, :saml], default: :email, source: :identityProvider)
    field(:identity_provider_account_id, :string, source: :identityProviderAccountId)
    field(:group_id, :string, source: :groupId)
    # [UserNotificationSettings]
    field(:notification_settings, :map, default: %{}, source: :notificationSettings)
    # [Locale]
    field(:locale, :string, default: "en-US")
    field(:last_login_at, Formbricks.PrismaDateTime, source: :lastLoginAt)
    field(:is_active, :boolean, default: true, source: :isActive)
    # When an admin last reactivated this user. Restarts the members retention clock (ENG-3713)
    # without touching `lastLoginAt`, which the member list, usage counts and APIs report as a real
    # sign-in.
    field(:reactivated_at, Formbricks.PrismaDateTime, source: :reactivatedAt)
    # When the user last completed a sign-in (`lastActivePlugin`, after any second factor) or renewed a
    # session (Better Auth `databaseHooks.session.update`), on the database clock and never moved
    # backwards. The members retention clock reads it: sessions are
    # deleted on sign-out and expiry, so their own timestamps can't serve as a record of activity.
    field(:last_active_at, Formbricks.PrismaDateTime, source: :lastActiveAt)
    field(:is_bootstrap_admin, Ecto.Enum, values: [:bootstrapAdmin], source: :isBootstrapAdmin)
    has_one(:two_factor, Formbricks.TwoFactor, foreign_key: :user_id, references: :id)
    has_one(:password_reset_token, Formbricks.PasswordResetToken, foreign_key: :user_id, references: :id)
    has_many(:memberships, Formbricks.Membership, foreign_key: :user_id, references: :id)
    has_many(:accounts, Formbricks.Account, foreign_key: :user_id, references: :id)
    has_many(:sessions, Formbricks.Session, foreign_key: :user_id, references: :id)
    has_many(:oauth_clients, Formbricks.OauthClient, foreign_key: :user_id, references: :id)
    has_many(:oauth_access_tokens, Formbricks.OauthAccessToken, foreign_key: :user_id, references: :id)
    has_many(:oauth_refresh_tokens, Formbricks.OauthRefreshToken, foreign_key: :user_id, references: :id)
    has_many(:oauth_consents, Formbricks.OauthConsent, foreign_key: :user_id, references: :id)
    has_many(:invites_created, Formbricks.Invite, foreign_key: :creator_id, references: :id)
    has_many(:invites_accepted, Formbricks.Invite, foreign_key: :acceptor_id, references: :id)
    has_many(:surveys, Formbricks.Survey, foreign_key: :created_by, references: :id)
    has_many(:owned_surveys, Formbricks.Survey, foreign_key: :owner_id, references: :id)
    has_many(:charts, Formbricks.Chart, foreign_key: :created_by, references: :id)
    has_many(:dashboards, Formbricks.Dashboard, foreign_key: :created_by, references: :id)
    has_many(:creatives, Formbricks.Creative, foreign_key: :created_by, references: :id)
    has_many(:feedback_sources, Formbricks.FeedbackSource, foreign_key: :created_by, references: :id)
    has_many(:team_users, Formbricks.TeamUser, foreign_key: :user_id, references: :id)
    has_many(:workflows, Formbricks.Workflow, foreign_key: :created_by, references: :id)
    has_many(:workflow_versions_published, Formbricks.WorkflowVersion, foreign_key: :published_by, references: :id)
    has_many(:retention_policies_updated, Formbricks.RetentionPolicy, foreign_key: :updated_by_id, references: :id)
    has_many(:exemptions_created, Formbricks.RetentionExemption, foreign_key: :created_by_id, references: :id)
    has_many(:exemptions_revoked, Formbricks.RetentionExemption, foreign_key: :revoked_by_id, references: :id)
    has_many(:retention_notices, Formbricks.RetentionNotice, foreign_key: :user_id, references: :id)
    has_many(:attachment_exports, Formbricks.AttachmentExport, foreign_key: :requested_by_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end