package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var NamingStrategy = schema.NamingStrategy{SingularTable: true, NoLowerCase: true}

type User struct {
	ID    string `gorm:"column:id;primaryKey" json:"id"`
	Email string `gorm:"column:email;uniqueIndex;not null" json:"email"`
	// Address awaiting confirmation; becomes `email` only after a link sent to it is opened
	PendingEmail              *string        `gorm:"column:pendingEmail" json:"pendingEmail"`
	AuthenticationMethod      string         `gorm:"column:authenticationMethod;not null" json:"authenticationMethod"`
	AuthenticationProfile     datatypes.JSON `gorm:"column:authenticationProfile" json:"authenticationProfile"`
	AuthenticationExtraParams datatypes.JSON `gorm:"column:authenticationExtraParams" json:"authenticationExtraParams"`
	AuthIdentifier            *string        `gorm:"column:authIdentifier;uniqueIndex" json:"authIdentifier"`
	DisplayName               *string        `gorm:"column:displayName" json:"displayName"`
	Name                      *string        `gorm:"column:name" json:"name"`
	AvatarURL                 *string        `gorm:"column:avatarUrl" json:"avatarUrl"`
	Admin                     bool           `gorm:"column:admin;default:false;not null" json:"admin"`
	// Preferences for the dashboard
	DashboardPreferences datatypes.JSON `gorm:"column:dashboardPreferences" json:"dashboardPreferences"`
	CreatedAt            DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// @deprecated
	IsOnCloudWaitlist bool `gorm:"column:isOnCloudWaitlist;default:false;not null" json:"isOnCloudWaitlist"`
	// @deprecated
	FeatureCloud bool `gorm:"column:featureCloud;default:false;not null" json:"featureCloud"`
	// @deprecated
	IsOnHostedRepoWaitlist bool           `gorm:"column:isOnHostedRepoWaitlist;default:false;not null" json:"isOnHostedRepoWaitlist"`
	MarketingEmails        *bool          `gorm:"column:marketingEmails;default:true;not null" json:"marketingEmails"`
	ConfirmedBasicDetails  bool           `gorm:"column:confirmedBasicDetails;default:false;not null" json:"confirmedBasicDetails"`
	ReferralSource         *string        `gorm:"column:referralSource" json:"referralSource"`
	OnboardingData         datatypes.JSON `gorm:"column:onboardingData" json:"onboardingData"`
	MfaEnabledAt           *DateTime      `gorm:"column:mfaEnabledAt" json:"mfaEnabledAt"`
	MfaSecretReferenceID   *string        `gorm:"column:mfaSecretReferenceId" json:"mfaSecretReferenceId"`
	// Hash of the last used code to prevent replay attacks
	MfaLastUsedCode *string `gorm:"column:mfaLastUsedCode" json:"mfaLastUsedCode"`
	// Maximum session lifetime in seconds for this user. Default = 1 year (Gregorian).
	// May be further restricted by Organization.maxSessionDuration on any of the user's orgs.
	SessionDuration *int `gorm:"column:sessionDuration;default:31556952;not null" json:"sessionDuration"`
	// Absolute timestamp at which the user's current session must end. Written
	// at login (and any time the effective duration is recomputed) and shortened
	// when an admin lowers an org cap. Read on requireUser / getUser; null is
	// treated as "no enforced deadline" (legacy sessions from before this
	// shipped). Read-on-demand has zero per-request DB cost — the value
	// piggybacks on the User fetch that requireUser already does.
	NextSessionEnd                   *DateTime        `gorm:"column:nextSessionEnd" json:"nextSessionEnd"`
	InvitationCodeID                 *string          `gorm:"column:invitationCodeId" json:"invitationCodeId"`
	MfaSecretReference               *SecretReference `gorm:"foreignKey:MfaSecretReferenceID"`
	InvitationCode                   *InvitationCode
	OrgMemberships                   []OrgMember                       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SentInvites                      []OrgMemberInvite                 `gorm:"foreignKey:InviterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PersonalAccessTokens             []PersonalAccessToken             `gorm:"foreignKey:UserID"`
	CreatedAPIKeys                   []APIKey                          `gorm:"foreignKey:CreatedByUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Deployments                      []WorkerDeployment                `gorm:"foreignKey:TriggeredByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	BackupCodes                      []MfaBackupCode                   `gorm:"foreignKey:UserID"`
	BulkActions                      []BulkActionGroup                 `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ImpersonationsPerformed          []ImpersonationAuditLog           `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ImpersonationsReceived           []ImpersonationAuditLog           `gorm:"foreignKey:TargetID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SupportAccessRequestsMade        []SupportAccessRequest            `gorm:"foreignKey:RequestedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SupportAccessRequestsApproved    []SupportAccessRequest            `gorm:"foreignKey:ApprovedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CustomerQueries                  []CustomerQuery                   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	MetricsDashboards                []MetricsDashboard                `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PlatformNotifications            []PlatformNotification            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlatformNotificationInteractions []PlatformNotificationInteraction `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "User"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*User) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type MfaBackupCode struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// Hash of the actual code
	Code      string    `gorm:"column:code;uniqueIndex:MfaBackupCode_userId_code_key;not null" json:"code"`
	UserID    string    `gorm:"column:userId;uniqueIndex:MfaBackupCode_userId_code_key;not null" json:"userId"`
	UsedAt    *DateTime `gorm:"column:usedAt" json:"usedAt"`
	CreatedAt DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	User      *User
}

func (MfaBackupCode) TableName() string {
	return "MfaBackupCode"
}

func (m *MfaBackupCode) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*MfaBackupCode) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type InvitationCode struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Code      string   `gorm:"column:code;uniqueIndex;not null" json:"code"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Users     []User   `gorm:"foreignKey:InvitationCodeID"`
}

func (InvitationCode) TableName() string {
	return "InvitationCode"
}

func (m *InvitationCode) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Used to generate PersonalAccessTokens, they're one-time use
type AuthorizationCode struct {
	ID                    string  `gorm:"column:id;primaryKey" json:"id"`
	Code                  string  `gorm:"column:code;uniqueIndex;not null" json:"code"`
	PersonalAccessTokenID *string `gorm:"column:personalAccessTokenId" json:"personalAccessTokenId"`
	// Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires.
	RequesterIP *string `gorm:"column:requesterIp" json:"requesterIp"`
	// S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it.
	CodeChallenge *string `gorm:"column:codeChallenge" json:"codeChallenge"`
	// Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it.
	CallbackPort *int `gorm:"column:callbackPort" json:"callbackPort"`
	// sha256 of the approval code, which only the approving browser ever sees.
	ApprovalCodeHash    *string  `gorm:"column:approvalCodeHash" json:"approvalCodeHash"`
	CreatedAt           DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	PersonalAccessToken *PersonalAccessToken
}

func (AuthorizationCode) TableName() string {
	return "AuthorizationCode"
}

func (m *AuthorizationCode) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AuthorizationCode) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PersonalAccessToken struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// If generated by the CLI this will be "cli", otherwise user-provided
	Name string `gorm:"column:name;not null" json:"name"`
	// This is the token encrypted using the ENCRYPTION_KEY
	EncryptedToken datatypes.JSON `gorm:"column:encryptedToken;not null" json:"encryptedToken"`
	// This is shown in the UI, with ********
	ObfuscatedToken string `gorm:"column:obfuscatedToken;not null" json:"obfuscatedToken"`
	// This is used to find the token in the database
	HashedToken    string    `gorm:"column:hashedToken;uniqueIndex;not null" json:"hashedToken"`
	UserID         string    `gorm:"column:userId;index:PersonalAccessToken_userId_idx;not null" json:"userId"`
	RevokedAt      *DateTime `gorm:"column:revokedAt" json:"revokedAt"`
	LastAccessedAt *DateTime `gorm:"column:lastAccessedAt" json:"lastAccessedAt"`
	// After this the token no longer authenticates. Null means it never expires.
	ExpiresAt *DateTime `gorm:"column:expiresAt" json:"expiresAt"`
	// Set on CLI login tokens: they stop authenticating this many seconds after their last use
	// (or creation, if never used). Null means no idle limit.
	IdleTimeoutSeconds *int `gorm:"column:idleTimeoutSeconds" json:"idleTimeoutSeconds"`
	// When the "about to expire" email was sent. Set before sending, so it's sent at most once.
	ExpiryReminderSentAt *DateTime `gorm:"column:expiryReminderSentAt" json:"expiryReminderSentAt"`
	// When the "has expired" email was sent. Set before sending, so it's sent at most once.
	ExpiredNoticeSentAt *DateTime `gorm:"column:expiredNoticeSentAt" json:"expiredNoticeSentAt"`
	CreatedAt           DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	User                *User
	AuthorizationCodes  []AuthorizationCode `gorm:"foreignKey:PersonalAccessTokenID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (PersonalAccessToken) TableName() string {
	return "PersonalAccessToken"
}

func (m *PersonalAccessToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PersonalAccessToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganizationAccessToken struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// User-provided name for the token
	Name string `gorm:"column:name;not null" json:"name"`
	// Used to differentiate between user-generated and system-generated tokens
	Type string `gorm:"column:type;index:OrganizationAccessToken_organizationId_type_createdAt_idx;default:'USER';not null" json:"type"`
	// This is used to find the token in the database
	HashedToken    string `gorm:"column:hashedToken;uniqueIndex;not null" json:"hashedToken"`
	OrganizationID string `gorm:"column:organizationId;index:OrganizationAccessToken_organizationId_type_createdAt_idx;not null" json:"organizationId"`
	// Optional expiration date for the token
	ExpiresAt      *DateTime `gorm:"column:expiresAt" json:"expiresAt"`
	RevokedAt      *DateTime `gorm:"column:revokedAt" json:"revokedAt"`
	LastAccessedAt *DateTime `gorm:"column:lastAccessedAt" json:"lastAccessedAt"`
	CreatedAt      DateTime  `gorm:"column:createdAt;index:OrganizationAccessToken_organizationId_type_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization
}

func (OrganizationAccessToken) TableName() string {
	return "OrganizationAccessToken"
}

func (m *OrganizationAccessToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganizationAccessToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Organization struct {
	ID                             string `gorm:"column:id;primaryKey" json:"id"`
	Slug                           string `gorm:"column:slug;uniqueIndex;not null" json:"slug"`
	Title                          string `gorm:"column:title;not null" json:"title"`
	MaximumExecutionTimePerRunInMs *int   `gorm:"column:maximumExecutionTimePerRunInMs;default:900000;not null" json:"maximumExecutionTimePerRunInMs"`
	MaximumConcurrencyLimit        *int   `gorm:"column:maximumConcurrencyLimit;default:10;not null" json:"maximumConcurrencyLimit"`
	// This is deprecated and will be removed in the future
	MaximumSchedulesLimit    *int           `gorm:"column:maximumSchedulesLimit;default:5;not null" json:"maximumSchedulesLimit"`
	MaximumDevQueueSize      *int           `gorm:"column:maximumDevQueueSize" json:"maximumDevQueueSize"`
	MaximumDeployedQueueSize *int           `gorm:"column:maximumDeployedQueueSize" json:"maximumDeployedQueueSize"`
	CreatedAt                DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt                *DateTime      `gorm:"column:deletedAt" json:"deletedAt"`
	CompanySize              *string        `gorm:"column:companySize" json:"companySize"`
	OnboardingData           datatypes.JSON `gorm:"column:onboardingData" json:"onboardingData"`
	Avatar                   datatypes.JSON `gorm:"column:avatar" json:"avatar"`
	RunsEnabled              *bool          `gorm:"column:runsEnabled;default:true;not null" json:"runsEnabled"`
	IsActivated              bool           `gorm:"column:v3Enabled;default:false;not null" json:"v3Enabled"`
	// @deprecated
	V2Enabled bool `gorm:"column:v2Enabled;default:false;not null" json:"v2Enabled"`
	// @deprecated
	V2MarqsEnabled bool `gorm:"column:v2MarqsEnabled;default:false;not null" json:"v2MarqsEnabled"`
	// @deprecated
	HasRequestedV3              bool           `gorm:"column:hasRequestedV3;default:false;not null" json:"hasRequestedV3"`
	APIRateLimiterConfig        datatypes.JSON `gorm:"column:apiRateLimiterConfig" json:"apiRateLimiterConfig"`
	RealtimeRateLimiterConfig   datatypes.JSON `gorm:"column:realtimeRateLimiterConfig" json:"realtimeRateLimiterConfig"`
	BatchRateLimitConfig        datatypes.JSON `gorm:"column:batchRateLimitConfig" json:"batchRateLimitConfig"`
	BatchQueueConcurrencyConfig datatypes.JSON `gorm:"column:batchQueueConcurrencyConfig" json:"batchQueueConcurrencyConfig"`
	// Webhook limits from the org's plan or the back office: { maxWaitersPerEnvironment, maxWaitersPerEndpoint, concurrency }. Null uses the env defaults.
	WebhookLimitsConfig datatypes.JSON `gorm:"column:webhookLimitsConfig" json:"webhookLimitsConfig"`
	FeatureFlags        datatypes.JSON `gorm:"column:featureFlags" json:"featureFlags"`
	MaximumProjectCount *int           `gorm:"column:maximumProjectCount;default:25;not null" json:"maximumProjectCount"`
	// Maximum session lifetime in seconds for members of this organization.
	// When set, caps each member's User.sessionDuration. The most restrictive cap across
	// all of a user's orgs (excluding nulls) wins.
	MaxSessionDuration *int `gorm:"column:maxSessionDuration" json:"maxSessionDuration"`
	// S2 basin that holds this org's realtime streams. Null until the
	// per-org basin has been provisioned (OSS / s2-lite installs leave
	// it null forever; reads fall back to the global basin env var).
	// Set once at provisioning time; retention is reconfigured in-place
	// when the org's plan changes.
	StreamBasinName              *string                   `gorm:"column:streamBasinName" json:"streamBasinName"`
	SupportAccessMode            string                    `gorm:"column:supportAccessMode;default:'ALLOW';not null" json:"supportAccessMode"`
	SupportAccessModeUpdatedAt   *DateTime                 `gorm:"column:supportAccessModeUpdatedAt" json:"supportAccessModeUpdatedAt"`
	SupportAccessModeUpdatedByID *string                   `gorm:"column:supportAccessModeUpdatedById" json:"supportAccessModeUpdatedById"`
	Environments                 []RuntimeEnvironment      `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Projects                     []Project                 `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Members                      []OrgMember               `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites                      []OrgMemberInvite         `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationIntegrations     []OrganizationIntegration `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationAccessTokens     []OrganizationAccessToken `gorm:"foreignKey:OrganizationID"`
	WorkerGroups                 []WorkerInstanceGroup     `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkerInstances              []WorkerInstance          `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GithubAppInstallations       []GithubAppInstallation   `gorm:"foreignKey:OrganizationID"`
	CustomerQueries              []CustomerQuery           `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MetricsDashboards            []MetricsDashboard        `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Prompts                      []Prompt                  `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlatformNotifications        []PlatformNotification    `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ErrorGroupStates             []ErrorGroupState         `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SupportAccessRequests        []SupportAccessRequest    `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Organization) TableName() string {
	return "Organization"
}

func (m *Organization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Organization) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrgMember struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string   `gorm:"column:organizationId;uniqueIndex:OrgMember_organizationId_userId_key;not null" json:"organizationId"`
	UserID         string   `gorm:"column:userId;uniqueIndex:OrgMember_organizationId_userId_key;index:OrgMember_userId_idx;not null" json:"userId"`
	Role           string   `gorm:"column:role;default:'MEMBER';not null" json:"role"`
	CreatedAt      DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization
	User           *User
	Environments   []RuntimeEnvironment `gorm:"foreignKey:OrgMemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (OrgMember) TableName() string {
	return "OrgMember"
}

func (m *OrgMember) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrgMember) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrgMemberInvite struct {
	ID    string `gorm:"column:id;primaryKey" json:"id"`
	Token string `gorm:"column:token;uniqueIndex;not null" json:"token"`
	Email string `gorm:"column:email;uniqueIndex:OrgMemberInvite_organizationId_email_key;not null" json:"email"`
	Role  string `gorm:"column:role;default:'MEMBER';not null" json:"role"`
	// Optional RBAC role to assign on invite acceptance. When set, the
	// accept-invite flow calls the loaded RBAC plugin's setUserRole with
	// this id after creating the OrgMember. Null = legacy behaviour, the
	// runtime fallback derives the role from `role` above.
	//
	// Plain text (not an FK) — the RBAC plugin's RbacRole table lives on
	// a separate schema (Drizzle, not Prisma) so we can't model the FK
	// here. Validation happens at write time (action) and read time
	// (acceptInvite).
	RbacRoleID     *string  `gorm:"column:rbacRoleId" json:"rbacRoleId"`
	OrganizationID string   `gorm:"column:organizationId;uniqueIndex:OrgMemberInvite_organizationId_email_key;not null" json:"organizationId"`
	InviterID      string   `gorm:"column:inviterId;not null" json:"inviterId"`
	CreatedAt      DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization
	Inviter        *User `gorm:"foreignKey:InviterID"`
}

func (OrgMemberInvite) TableName() string {
	return "OrgMemberInvite"
}

func (m *OrgMemberInvite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.Token == "" {
		m.Token = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrgMemberInvite) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RuntimeEnvironment struct {
	ID     string `gorm:"column:id;primaryKey" json:"id"`
	Slug   string `gorm:"column:slug;uniqueIndex:RuntimeEnvironment_projectId_slug_orgMemberId_key;not null" json:"slug"`
	APIKey string `gorm:"column:apiKey;uniqueIndex;not null" json:"apiKey"`
	// Non-null prevents the root API key from being returned to users.
	RootAPIKeyHiddenAt *DateTime `gorm:"column:rootApiKeyHiddenAt" json:"rootApiKeyHiddenAt"`
	// @deprecated was for v2
	PkAPIKey string `gorm:"column:pkApiKey;uniqueIndex;not null" json:"pkApiKey"`
	Type     string `gorm:"column:type;default:'DEVELOPMENT';not null" json:"type"`
	// If true, this environment has branches and is treated differently in the dashboard/API
	// NB: this flag is NOT used for Development branches, instead (type, parentEnvironmentId) = (DEVELOPMENT, NULL) is used
	IsBranchableEnvironment bool    `gorm:"column:isBranchableEnvironment;default:false;not null" json:"isBranchableEnvironment"`
	BranchName              *string `gorm:"column:branchName" json:"branchName"`
	ParentEnvironmentID     *string `gorm:"column:parentEnvironmentId;index:RuntimeEnvironment_parentEnvironmentId_idx" json:"parentEnvironmentId"`
	// Only configured on the PREVIEW root. Null disables automatic archiving.
	PreviewAutoArchiveAfterDays        *int     `gorm:"column:previewAutoArchiveAfterDays" json:"previewAutoArchiveAfterDays"`
	PreviewAutoArchiveExcludedBranches []string `gorm:"column:previewAutoArchiveExcludedBranches;serializer:json;not null" json:"previewAutoArchiveExcludedBranches"`
	// Durable per-project cleanup progress. Null next-check disables scheduling.
	PreviewAutoArchiveNextCheckAt     *DateTime      `gorm:"column:previewAutoArchiveNextCheckAt" json:"previewAutoArchiveNextCheckAt"`
	PreviewAutoArchiveCursorCreatedAt *DateTime      `gorm:"column:previewAutoArchiveCursorCreatedAt" json:"previewAutoArchiveCursorCreatedAt"`
	PreviewAutoArchiveCursorID        *string        `gorm:"column:previewAutoArchiveCursorId" json:"previewAutoArchiveCursorId"`
	Git                               datatypes.JSON `gorm:"column:git" json:"git"`
	// When set API calls will fail
	ArchivedAt *DateTime `gorm:"column:archivedAt" json:"archivedAt"`
	// A memorable code for the environment
	Shortcode                   string   `gorm:"column:shortcode;uniqueIndex:RuntimeEnvironment_projectId_shortcode_key;not null" json:"shortcode"`
	MaximumConcurrencyLimit     *int     `gorm:"column:maximumConcurrencyLimit;default:5;not null" json:"maximumConcurrencyLimit"`
	ConcurrencyLimitBurstFactor *float64 `gorm:"column:concurrencyLimitBurstFactor;type:decimal(4,2);default:2;not null" json:"concurrencyLimitBurstFactor"`
	Paused                      bool     `gorm:"column:paused;default:false;not null" json:"paused"`
	PauseSource                 *string  `gorm:"column:pauseSource" json:"pauseSource"`
	AutoEnableInternalSources   *bool    `gorm:"column:autoEnableInternalSources;default:true;not null" json:"autoEnableInternalSources"`
	OrganizationID              string   `gorm:"column:organizationId;index:RuntimeEnvironment_organizationId_idx;not null" json:"organizationId"`
	ProjectID                   string   `gorm:"column:projectId;uniqueIndex:RuntimeEnvironment_projectId_slug_orgMemberId_key;uniqueIndex:RuntimeEnvironment_projectId_shortcode_key;index:RuntimeEnvironment_projectId_idx;not null" json:"projectId"`
	OrgMemberID                 *string  `gorm:"column:orgMemberId;uniqueIndex:RuntimeEnvironment_projectId_slug_orgMemberId_key" json:"orgMemberId"`
	CreatedAt                   DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                   DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// Allows us to customize the built-in environment variables for a specific environment, like TRIGGER_OTEL_SPAN_ATTRIBUTE_VALUE_LENGTH_LIMIT
	BuiltInEnvironmentVariableOverrides datatypes.JSON      `gorm:"column:builtInEnvironmentVariableOverrides" json:"builtInEnvironmentVariableOverrides"`
	TunnelID                            *string             `gorm:"column:tunnelId" json:"tunnelId"`
	CurrentSessionID                    *string             `gorm:"column:currentSessionId" json:"currentSessionId"`
	ParentEnvironment                   *RuntimeEnvironment `gorm:"foreignKey:ParentEnvironmentID"`
	Organization                        *Organization
	Project                             *Project
	OrgMember                           *OrgMember
	CurrentSession                      *RuntimeEnvironmentSession   `gorm:"foreignKey:CurrentSessionID"`
	EncryptionKeys                      []EnvironmentEncryptionKey   `gorm:"foreignKey:EnvironmentID;constraint:OnDelete:CASCADE"`
	ChildEnvironments                   []RuntimeEnvironment         `gorm:"foreignKey:ParentEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BackgroundWorkers                   []BackgroundWorker           `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BackgroundWorkerTasks               []BackgroundWorkerTask       `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskRuns                            []TaskRun                    `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskQueues                          []TaskQueue                  `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	EnvironmentVariableValues           []EnvironmentVariableValue   `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Checkpoints                         []Checkpoint                 `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkerDeployments                   []WorkerDeployment           `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkerDeploymentPromotions          []WorkerDeploymentPromotion  `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskRunAttempts                     []TaskRunAttempt             `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CheckpointRestoreEvent              []CheckpointRestoreEvent     `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskScheduleInstances               []TaskScheduleInstance       `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Alerts                              []ProjectAlert               `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions                            []RuntimeEnvironmentSession  `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskRunNumberCounter                []TaskRunNumberCounter       `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskRunCheckpoints                  []TaskRunCheckpoint          `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Waitpoints                          []Waitpoint                  `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkerInstances                     []WorkerInstance             `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	WaitpointTags                       []WaitpointTag               `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BulkActionGroup                     []BulkActionGroup            `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomerQueries                     []CustomerQuery              `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Prompts                             []Prompt                     `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlaygroundConversations             []PlaygroundConversation     `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ErrorGroupStates                    []ErrorGroupState            `gorm:"foreignKey:EnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskIdentifiers                     []TaskIdentifier             `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeys                             []APIKey                     `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RevokedAPIKeys                      []RevokedAPIKey              `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	EncryptionPolicy                    *EnvironmentEncryptionPolicy `gorm:"foreignKey:EnvironmentID;constraint:OnDelete:CASCADE"`
	TrustedDeployGrant                  *TrustedDeployEnvironment    `gorm:"foreignKey:RuntimeEnvironmentID;constraint:OnDelete:CASCADE"`
}

func (RuntimeEnvironment) TableName() string {
	return "RuntimeEnvironment"
}

func (m *RuntimeEnvironment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*RuntimeEnvironment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type APIKey struct {
	ID       string `gorm:"column:id;primaryKey" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	KeyHash  string `gorm:"column:key_hash;uniqueIndex;not null" json:"key_hash"`
	LastFour string `gorm:"column:last_four;not null" json:"last_four"`
	// Optional one-line note the creator adds to say what the key is for.
	Description          *string  `gorm:"column:description" json:"description"`
	RuntimeEnvironmentID string   `gorm:"column:runtime_environment_id;index:api_keys_runtime_environment_id_revoked_at_created_at_idx;index:api_keys_env_issued_by_revoked_at_created_at_idx;not null" json:"runtime_environment_id"`
	CreatedByUserID      *string  `gorm:"column:created_by_user_id" json:"created_by_user_id"`
	IssuedBy             string   `gorm:"column:issued_by;index:api_keys_env_issued_by_revoked_at_created_at_idx;default:'USER';not null" json:"issued_by"`
	PresetID             *string  `gorm:"column:preset_id" json:"preset_id"`
	Scopes               []string `gorm:"column:scopes;serializer:json;not null" json:"scopes"`
	// CIDR ranges this key may be used from. Empty means any IP.
	AllowedIPRanges    []string  `gorm:"column:allowed_ip_ranges;serializer:json;not null" json:"allowed_ip_ranges"`
	LastUsedAt         *DateTime `gorm:"column:last_used_at" json:"last_used_at"`
	LastUsedIP         *string   `gorm:"column:last_used_ip" json:"last_used_ip"`
	RevokedAt          *DateTime `gorm:"column:revoked_at;index:api_keys_runtime_environment_id_revoked_at_created_at_idx;index:api_keys_env_issued_by_revoked_at_created_at_idx" json:"revoked_at"`
	ExpiresAt          *DateTime `gorm:"column:expires_at" json:"expires_at"`
	UpdatedAt          DateTime  `gorm:"column:updated_at;autoUpdateTime:false;not null" json:"updated_at"`
	CreatedAt          DateTime  `gorm:"column:created_at;index:api_keys_runtime_environment_id_revoked_at_created_at_idx;index:api_keys_env_issued_by_revoked_at_created_at_idx;autoCreateTime:false;not null" json:"created_at"`
	RuntimeEnvironment *RuntimeEnvironment
	CreatedBy          *User `gorm:"foreignKey:CreatedByUserID"`
}

func (APIKey) TableName() string {
	return "api_keys"
}

func (m *APIKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

func (*APIKey) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Records of previously-valid API keys that are still accepted for authentication
// during a grace window after rotation. Extend or end the grace period by updating `expiresAt`.
type RevokedAPIKey struct {
	ID                   string   `gorm:"column:id;primaryKey" json:"id"`
	APIKey               string   `gorm:"column:apiKey;index:RevokedApiKey_apiKey_idx;not null" json:"apiKey"`
	RuntimeEnvironmentID string   `gorm:"column:runtimeEnvironmentId;index:RevokedApiKey_runtimeEnvironmentId_idx;not null" json:"runtimeEnvironmentId"`
	ExpiresAt            DateTime `gorm:"column:expiresAt;not null" json:"expiresAt"`
	CreatedAt            DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	RuntimeEnvironment   *RuntimeEnvironment
}

func (RevokedAPIKey) TableName() string {
	return "RevokedApiKey"
}

func (m *RevokedAPIKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Project struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`
	Slug           string    `gorm:"column:slug;uniqueIndex;not null" json:"slug"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	ExternalRef    string    `gorm:"column:externalRef;uniqueIndex;not null" json:"externalRef"`
	OrganizationID string    `gorm:"column:organizationId;not null" json:"organizationId"`
	CreatedAt      DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt      *DateTime `gorm:"column:deletedAt" json:"deletedAt"`
	// Set the first time the CLI `init` command completes against this project. Drives the dev onboarding progress.
	InitializedAt *DateTime `gorm:"column:initializedAt" json:"initializedAt"`
	// Runtime used when a deployment config does not specify one. Null preserves the legacy Node 20 fallback.
	DefaultRuntime       *string `gorm:"column:defaultRuntime" json:"defaultRuntime"`
	Version              string  `gorm:"column:version;default:'V2';not null" json:"version"`
	Engine               string  `gorm:"column:engine;default:'V1';not null" json:"engine"`
	BuilderProjectID     *string `gorm:"column:builderProjectId" json:"builderProjectId"`
	DefaultWorkerGroupID *string `gorm:"column:defaultWorkerGroupId" json:"defaultWorkerGroupId"`
	// The master queues they are allowed to use (impacts what they can set as default and trigger runs with)
	AllowedWorkerQueues            []string       `gorm:"column:allowedMasterQueues;serializer:json;not null" json:"allowedMasterQueues"`
	BuildSettings                  datatypes.JSON `gorm:"column:buildSettings" json:"buildSettings"`
	OnboardingData                 datatypes.JSON `gorm:"column:onboardingData" json:"onboardingData"`
	Organization                   *Organization
	DefaultWorkerGroup             *WorkerInstanceGroup             `gorm:"foreignKey:DefaultWorkerGroupID"`
	WorkerGroups                   []WorkerInstanceGroup            `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Workers                        []WorkerInstance                 `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Environments                   []RuntimeEnvironment             `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BackgroundWorkers              []BackgroundWorker               `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BackgroundWorkerTasks          []BackgroundWorkerTask           `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskRuns                       []TaskRun                        `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RunTags                        []TaskRunTag                     `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskQueues                     []TaskQueue                      `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	EnvironmentVariables           []EnvironmentVariable            `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Checkpoints                    []Checkpoint                     `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkerDeployment               []WorkerDeployment               `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CheckpointRestoreEvent         []CheckpointRestoreEvent         `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskSchedules                  []TaskSchedule                   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AlertChannels                  []ProjectAlertChannel            `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Alerts                         []ProjectAlert                   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AlertStorages                  []ProjectAlertStorage            `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BulkActionGroups               []BulkActionGroup                `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BackgroundWorkerFile           []BackgroundWorkerFile           `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Waitpoints                     []Waitpoint                      `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskRunWaitpoints              []TaskRunWaitpoint               `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskRunCheckpoints             []TaskRunCheckpoint              `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	WaitpointTags                  []WaitpointTag                   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationProjectIntegration []OrganizationProjectIntegration `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomerQueries                []CustomerQuery                  `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlaygroundConversations        []PlaygroundConversation         `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskScheduleInstances          []TaskScheduleInstance           `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MetricsDashboards              []MetricsDashboard               `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LlmModels                      []LlmModel                       `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Prompts                        []Prompt                         `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlatformNotifications          []PlatformNotification           `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ErrorGroupStates               []ErrorGroupState                `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskIdentifiers                []TaskIdentifier                 `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConnectedGithubRepository      *ConnectedGithubRepository       `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TrustedDeployRepository        *TrustedDeployRepository         `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}

func (Project) TableName() string {
	return "Project"
}

func (m *Project) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Project) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EnvironmentEncryptionPolicy struct {
	EnvironmentID        string                    `gorm:"column:environmentId;primaryKey" json:"environmentId"`
	Revision             int                       `gorm:"column:revision;default:0;not null" json:"revision"`
	WrappingProvider     *string                   `gorm:"column:wrappingProvider" json:"wrappingProvider"`
	WrappingKeyReference *string                   `gorm:"column:wrappingKeyReference" json:"wrappingKeyReference"`
	ActiveKeyID          *string                   `gorm:"column:activeKeyId" json:"activeKeyId"`
	Environment          *RuntimeEnvironment       `gorm:"foreignKey:EnvironmentID"`
	ActiveKey            *EnvironmentEncryptionKey `gorm:"foreignKey:ActiveKeyID,EnvironmentID;references:ID,EnvironmentID"`
}

func (EnvironmentEncryptionPolicy) TableName() string {
	return "EnvironmentEncryptionPolicy"
}

type EnvironmentEncryptionKey struct {
	ID                   string                        `gorm:"column:id;primaryKey;uniqueIndex:EnvironmentEncryptionKey_id_environmentId_key" json:"id"`
	EnvironmentID        string                        `gorm:"column:environmentId;uniqueIndex:EnvironmentEncryptionKey_id_environmentId_key;index:EnvironmentEncryptionKey_environmentId_idx;not null" json:"environmentId"`
	WrappedKey           datatypes.JSON                `gorm:"column:wrappedKey;not null" json:"wrappedKey"`
	WrappingProvider     string                        `gorm:"column:wrappingProvider;not null" json:"wrappingProvider"`
	WrappingKeyReference string                        `gorm:"column:wrappingKeyReference;not null" json:"wrappingKeyReference"`
	Revision             int                           `gorm:"column:revision;default:0;not null" json:"revision"`
	CreatedAt            DateTime                      `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Environment          *RuntimeEnvironment           `gorm:"foreignKey:EnvironmentID"`
	ActiveFor            []EnvironmentEncryptionPolicy `gorm:"foreignKey:ActiveKeyID,EnvironmentID;references:ID,EnvironmentID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}

func (EnvironmentEncryptionKey) TableName() string {
	return "EnvironmentEncryptionKey"
}

func (m *EnvironmentEncryptionKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SecretReference struct {
	ID                        string                     `gorm:"column:id;primaryKey" json:"id"`
	Key                       string                     `gorm:"column:key;uniqueIndex;not null" json:"key"`
	Provider                  string                     `gorm:"column:provider;default:'DATABASE';not null" json:"provider"`
	CreatedAt                 DateTime                   `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime                   `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	EnvironmentVariableValues []EnvironmentVariableValue `gorm:"foreignKey:ValueReferenceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	OrganizationIntegration   []OrganizationIntegration  `gorm:"foreignKey:TokenReferenceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User                      []User                     `gorm:"foreignKey:MfaSecretReferenceID"`
}

func (SecretReference) TableName() string {
	return "SecretReference"
}

func (m *SecretReference) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SecretReference) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SecretStore struct {
	Key       string         `gorm:"column:key;primaryKey;uniqueIndex;index:SecretStore_key_idx" json:"key"`
	Value     datatypes.JSON `gorm:"column:value;not null" json:"value"`
	Version   *string        `gorm:"column:version;default:'1';not null" json:"version"`
	CreatedAt DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SecretStore) TableName() string {
	return "SecretStore"
}

func (m *SecretStore) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SecretStore) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OidcSigningKey struct {
	ID      string `gorm:"column:id;primaryKey" json:"id"`
	Kid     string `gorm:"column:kid;uniqueIndex;not null" json:"kid"`
	Purpose string `gorm:"column:purpose;default:'EXTERNAL_OIDC';not null" json:"purpose"`
	// At most one STAGED and one ACTIVE key per purpose: enforced by a partial unique index in the migration.
	Status            string         `gorm:"column:status;not null" json:"status"`
	PublicJwk         datatypes.JSON `gorm:"column:publicJwk;not null" json:"publicJwk"`
	ReadyAt           DateTime       `gorm:"column:readyAt;not null" json:"readyAt"`
	PublicUntil       *DateTime      `gorm:"column:publicUntil" json:"publicUntil"`
	GeneratedByUserID string         `gorm:"column:generatedByUserId;not null" json:"generatedByUserId"`
	ActivatedAt       *DateTime      `gorm:"column:activatedAt" json:"activatedAt"`
	ActivatedByUserID *string        `gorm:"column:activatedByUserId" json:"activatedByUserId"`
	RetiredAt         *DateTime      `gorm:"column:retiredAt" json:"retiredAt"`
	RetiredByUserID   *string        `gorm:"column:retiredByUserId" json:"retiredByUserId"`
	CreatedAt         DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (OidcSigningKey) TableName() string {
	return "OidcSigningKey"
}

func (m *OidcSigningKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OidcSigningKey) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DataMigration struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name;uniqueIndex;not null" json:"name"`
	CreatedAt   DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	CompletedAt *DateTime `gorm:"column:completedAt" json:"completedAt"`
}

func (DataMigration) TableName() string {
	return "DataMigration"
}

func (m *DataMigration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DataMigration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BackgroundWorker struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID           string         `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Engine               string         `gorm:"column:engine;default:'V1';not null" json:"engine"`
	ContentHash          string         `gorm:"column:contentHash;not null" json:"contentHash"`
	SdkVersion           *string        `gorm:"column:sdkVersion;default:'unknown';not null" json:"sdkVersion"`
	CliVersion           *string        `gorm:"column:cliVersion;default:'unknown';not null" json:"cliVersion"`
	ProjectID            string         `gorm:"column:projectId;uniqueIndex:BackgroundWorker_projectId_runtimeEnvironmentId_version_key;not null" json:"projectId"`
	RuntimeEnvironmentID string         `gorm:"column:runtimeEnvironmentId;uniqueIndex:BackgroundWorker_projectId_runtimeEnvironmentId_version_key;index:BackgroundWorker_runtimeEnvironmentId_idx;index:BackgroundWorker_runtimeEnvironmentId_createdAt_idx;not null" json:"runtimeEnvironmentId"`
	Version              string         `gorm:"column:version;uniqueIndex:BackgroundWorker_projectId_runtimeEnvironmentId_version_key;not null" json:"version"`
	Metadata             datatypes.JSON `gorm:"column:metadata;not null" json:"metadata"`
	Runtime              *string        `gorm:"column:runtime" json:"runtime"`
	RuntimeVersion       *string        `gorm:"column:runtimeVersion" json:"runtimeVersion"`
	CreatedAt            DateTime       `gorm:"column:createdAt;index:BackgroundWorker_runtimeEnvironmentId_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	WorkerGroupID        *string        `gorm:"column:workerGroupId" json:"workerGroupId"`
	SupportsLazyAttempts bool           `gorm:"column:supportsLazyAttempts;default:false;not null" json:"supportsLazyAttempts"`
	Project              *Project
	RuntimeEnvironment   *RuntimeEnvironment
	WorkerGroup          *WorkerInstanceGroup   `gorm:"foreignKey:WorkerGroupID"`
	Tasks                []BackgroundWorkerTask `gorm:"foreignKey:WorkerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attempts             []TaskRunAttempt       `gorm:"foreignKey:BackgroundWorkerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LockedRuns           []TaskRun              `gorm:"foreignKey:LockedToVersionID"`
	PromptVersions       []PromptVersion        `gorm:"foreignKey:WorkerID"`
	TaskIdentifiers      []TaskIdentifier       `gorm:"foreignKey:CurrentWorkerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Deployment           *WorkerDeployment      `gorm:"foreignKey:WorkerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Files                []BackgroundWorkerFile `gorm:"many2many:_BackgroundWorkerToBackgroundWorkerFile;joinForeignKey:A;joinReferences:B"`
	Queues               []TaskQueue            `gorm:"many2many:_BackgroundWorkerToTaskQueue;joinForeignKey:A;joinReferences:B"`
}

func (BackgroundWorker) TableName() string {
	return "BackgroundWorker"
}

func (m *BackgroundWorker) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BackgroundWorker) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BackgroundWorkerFile struct {
	ID                string   `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID        string   `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	FilePath          string   `gorm:"column:filePath;not null" json:"filePath"`
	ContentHash       string   `gorm:"column:contentHash;uniqueIndex:BackgroundWorkerFile_projectId_contentHash_key;not null" json:"contentHash"`
	Contents          []byte   `gorm:"column:contents;not null" json:"contents"`
	ProjectID         string   `gorm:"column:projectId;uniqueIndex:BackgroundWorkerFile_projectId_contentHash_key;not null" json:"projectId"`
	CreatedAt         DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Project           *Project
	Tasks             []BackgroundWorkerTask `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BackgroundWorkers []BackgroundWorker     `gorm:"many2many:_BackgroundWorkerToBackgroundWorkerFile;joinForeignKey:B;joinReferences:A"`
}

func (BackgroundWorkerFile) TableName() string {
	return "BackgroundWorkerFile"
}

func (m *BackgroundWorkerFile) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Prompt struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID           string         `gorm:"column:friendly_id;uniqueIndex;not null" json:"friendly_id"`
	Slug                 string         `gorm:"column:slug;uniqueIndex:prompts_projectId_runtimeEnvironmentId_slug_key;not null" json:"slug"`
	Description          *string        `gorm:"column:description" json:"description"`
	Type                 *string        `gorm:"column:type;default:'text';not null" json:"type"`
	OrganizationID       string         `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID            string         `gorm:"column:projectId;uniqueIndex:prompts_projectId_runtimeEnvironmentId_slug_key;index:prompts_projectId_idx;not null" json:"projectId"`
	RuntimeEnvironmentID string         `gorm:"column:runtimeEnvironmentId;uniqueIndex:prompts_projectId_runtimeEnvironmentId_slug_key;index:prompts_runtimeEnvironmentId_idx;not null" json:"runtimeEnvironmentId"`
	FilePath             *string        `gorm:"column:filePath" json:"filePath"`
	ExportName           *string        `gorm:"column:exportName" json:"exportName"`
	VariableSchema       datatypes.JSON `gorm:"column:variableSchema" json:"variableSchema"`
	DefaultModel         *string        `gorm:"column:defaultModel" json:"defaultModel"`
	DefaultConfig        datatypes.JSON `gorm:"column:defaultConfig" json:"defaultConfig"`
	Tags                 []string       `gorm:"column:tags;serializer:json;not null" json:"tags"`
	ArchivedAt           *DateTime      `gorm:"column:archivedAt" json:"archivedAt"`
	CreatedAt            DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization         *Organization
	Project              *Project
	RuntimeEnvironment   *RuntimeEnvironment
	Versions             []PromptVersion `gorm:"foreignKey:PromptID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Prompt) TableName() string {
	return "prompts"
}

func (m *Prompt) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Prompt) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PromptVersion struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	Version       int            `gorm:"column:version;uniqueIndex:prompt_versions_promptId_version_key;not null" json:"version"`
	PromptID      string         `gorm:"column:promptId;uniqueIndex:prompt_versions_promptId_version_key;index:prompt_versions_promptId_idx;not null" json:"promptId"`
	TextContent   *string        `gorm:"column:textContent" json:"textContent"`
	ChatContent   datatypes.JSON `gorm:"column:chatContent" json:"chatContent"`
	Model         *string        `gorm:"column:model" json:"model"`
	Config        datatypes.JSON `gorm:"column:config" json:"config"`
	Source        string         `gorm:"column:source;not null" json:"source"`
	CommitMessage *string        `gorm:"column:commitMessage" json:"commitMessage"`
	ContentHash   string         `gorm:"column:contentHash;index:prompt_versions_contentHash_idx;not null" json:"contentHash"`
	Labels        []string       `gorm:"column:labels;serializer:json;not null" json:"labels"`
	WorkerID      *string        `gorm:"column:workerId" json:"workerId"`
	CreatedBy     *string        `gorm:"column:createdBy" json:"createdBy"`
	CreatedAt     DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Prompt        *Prompt
	Worker        *BackgroundWorker `gorm:"foreignKey:WorkerID"`
}

func (PromptVersion) TableName() string {
	return "prompt_versions"
}

func (m *PromptVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type BackgroundWorkerTask struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	Slug                 string         `gorm:"column:slug;uniqueIndex:BackgroundWorkerTask_workerId_slug_key;index:BackgroundWorkerTask_projectId_slug_idx;index:BackgroundWorkerTask_projectId_slug_createdAt_idx;index:BackgroundWorkerTask_runtimeEnvironmentId_slug_triggerSource_idx;not null" json:"slug"`
	Description          *string        `gorm:"column:description" json:"description"`
	FriendlyID           string         `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	FilePath             string         `gorm:"column:filePath;not null" json:"filePath"`
	ExportName           *string        `gorm:"column:exportName" json:"exportName"`
	WorkerID             string         `gorm:"column:workerId;uniqueIndex:BackgroundWorkerTask_workerId_slug_key;not null" json:"workerId"`
	ProjectID            string         `gorm:"column:projectId;index:BackgroundWorkerTask_projectId_slug_idx;index:BackgroundWorkerTask_projectId_slug_createdAt_idx;index:BackgroundWorkerTask_runtimeEnvironmentId_projectId_idx;not null" json:"projectId"`
	FileID               *string        `gorm:"column:fileId" json:"fileId"`
	RuntimeEnvironmentID string         `gorm:"column:runtimeEnvironmentId;index:BackgroundWorkerTask_runtimeEnvironmentId_projectId_idx;index:BackgroundWorkerTask_runtimeEnvironmentId_slug_triggerSource_idx;not null" json:"runtimeEnvironmentId"`
	CreatedAt            DateTime       `gorm:"column:createdAt;index:BackgroundWorkerTask_projectId_slug_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	QueueConfig          datatypes.JSON `gorm:"column:queueConfig" json:"queueConfig"`
	RetryConfig          datatypes.JSON `gorm:"column:retryConfig" json:"retryConfig"`
	MachineConfig        datatypes.JSON `gorm:"column:machineConfig" json:"machineConfig"`
	// Gates declared on the task: other queues its runs must also hold a concurrency
	// slot in while executing. Shape: [{ queue: string, concurrencyKey?: string }]
	Gates                datatypes.JSON `gorm:"column:gates" json:"gates"`
	QueueID              *string        `gorm:"column:queueId" json:"queueId"`
	MaxDurationInSeconds *int           `gorm:"column:maxDurationInSeconds" json:"maxDurationInSeconds"`
	TTL                  *string        `gorm:"column:ttl" json:"ttl"`
	TriggerSource        string         `gorm:"column:triggerSource;index:BackgroundWorkerTask_runtimeEnvironmentId_slug_triggerSource_idx;default:'STANDARD';not null" json:"triggerSource"`
	// Extra task configuration JSON. Shape depends on triggerSource.
	// AGENT: { type: "ai-sdk-chat" }
	Config             datatypes.JSON    `gorm:"column:config" json:"config"`
	PayloadSchema      datatypes.JSON    `gorm:"column:payloadSchema" json:"payloadSchema"`
	Worker             *BackgroundWorker `gorm:"foreignKey:WorkerID"`
	Project            *Project
	File               *BackgroundWorkerFile `gorm:"foreignKey:FileID"`
	RuntimeEnvironment *RuntimeEnvironment
	Queue              *TaskQueue       `gorm:"foreignKey:QueueID"`
	Attempts           []TaskRunAttempt `gorm:"foreignKey:BackgroundWorkerTaskID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Runs               []TaskRun        `gorm:"foreignKey:LockedByID"`
}

func (BackgroundWorkerTask) TableName() string {
	return "BackgroundWorkerTask"
}

func (m *BackgroundWorkerTask) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BackgroundWorkerTask) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WebhookEndpoint struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID string `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	// URL key. CSPRNG, NOT a friendlyId: crypto.randomBytes(16) base64url.
	// Generated by the create path (deploy-sync / P2 API), NOT by @default.
	OpaqueID             string `gorm:"column:opaqueId;uniqueIndex;not null" json:"opaqueId"`
	OrganizationID       string `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID            string `gorm:"column:projectId;index:WebhookEndpoint_projectId_createdAt_idx;not null" json:"projectId"`
	RuntimeEnvironmentID string `gorm:"column:runtimeEnvironmentId;uniqueIndex:WebhookEndpoint_runtimeEnvironmentId_declaredId_endpointTenantId_endpointExternalRef_key;index:WebhookEndpoint_runtimeEnvironmentId_source_idx;not null" json:"runtimeEnvironmentId"`
	EnvironmentType      string `gorm:"column:environmentType;not null" json:"environmentType"`
	EndpointTenantID     string `gorm:"column:endpointTenantId;uniqueIndex:WebhookEndpoint_runtimeEnvironmentId_declaredId_endpointTenantId_endpointExternalRef_key;default:'';not null" json:"endpointTenantId"`
	EndpointExternalRef  string `gorm:"column:endpointExternalRef;uniqueIndex:WebhookEndpoint_runtimeEnvironmentId_declaredId_endpointTenantId_endpointExternalRef_key;default:'';not null" json:"endpointExternalRef"`
	Source               string `gorm:"column:source;index:WebhookEndpoint_runtimeEnvironmentId_source_idx;not null" json:"source"`
	// The declared webhooks.endpoint.define id (string ref, no relation).
	DeclaredID string `gorm:"column:declaredId;uniqueIndex:WebhookEndpoint_runtimeEnvironmentId_declaredId_endpointTenantId_endpointExternalRef_key;not null" json:"declaredId"`
	// Subscribers of this endpoint: StoredWebhookRoutingTarget[] ({ id, type: "task" | "session", ..., filterAst }).
	RoutingTargets     datatypes.JSON `gorm:"column:routingTargets;default:'[]';not null" json:"routingTargets"`
	VerifierArtifact   datatypes.JSON `gorm:"column:verifierArtifact;not null" json:"verifierArtifact"`
	Metadata           datatypes.JSON `gorm:"column:metadata;default:'{}';not null" json:"metadata"`
	SecretProvisioning *string        `gorm:"column:secretProvisioning;default:'either';not null" json:"secretProvisioning"`
	// The source's instructions for an AI agent registering this endpoint with the provider.
	SetupPrompt *string `gorm:"column:setupPrompt" json:"setupPrompt"`
	// SecretReference.key string. Plain String, NO @relation -> no FK to SecretReference.
	SigningSecretKey *string `gorm:"column:signingSecretKey" json:"signingSecretKey"`
	Status           string  `gorm:"column:status;default:'ACTIVE';not null" json:"status"`
	// When an operator disabled the endpoint via the dashboard/API. Null means the declarative sync
	// owns the status: a redeploy that re-declares a previously-removed (auto-deactivated) webhook
	// reactivates it. Non-null means the operator disabled it, so the sync leaves the status alone.
	ManuallyDeactivatedAt *DateTime `gorm:"column:manuallyDeactivatedAt" json:"manuallyDeactivatedAt"`
	CreatedAt             DateTime  `gorm:"column:createdAt;index:WebhookEndpoint_projectId_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (WebhookEndpoint) TableName() string {
	return "WebhookEndpoint"
}

func (m *WebhookEndpoint) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WebhookEndpoint) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WebhookDelivery struct {
	ID                   string  `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID           string  `gorm:"column:friendlyId;not null" json:"friendlyId"`
	WebhookEndpointID    string  `gorm:"column:webhookEndpointId;index:WebhookDelivery_webhookEndpointId_createdAt_idx;not null" json:"webhookEndpointId"`
	OrganizationID       string  `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID            string  `gorm:"column:projectId;not null" json:"projectId"`
	RuntimeEnvironmentID string  `gorm:"column:runtimeEnvironmentId;not null" json:"runtimeEnvironmentId"`
	EnvironmentType      string  `gorm:"column:environmentType;not null" json:"environmentType"`
	ExternalDeliveryID   string  `gorm:"column:externalDeliveryId;not null" json:"externalDeliveryId"`
	IdempotencyKey       string  `gorm:"column:idempotencyKey;not null" json:"idempotencyKey"`
	RunID                *string `gorm:"column:runId" json:"runId"`
	Status               string  `gorm:"column:status;default:'PENDING';not null" json:"status"`
	// Set from the x-trigger-test ingress header; marks console/test-send deliveries so the list can filter them.
	IsTest       bool           `gorm:"column:isTest;default:false;not null" json:"isTest"`
	ParsedEvent  datatypes.JSON `gorm:"column:parsedEvent" json:"parsedEvent"`
	Headers      datatypes.JSON `gorm:"column:headers" json:"headers"`
	RawBodyHash  *string        `gorm:"column:rawBodyHash" json:"rawBodyHash"`
	ErrorMessage *string        `gorm:"column:errorMessage" json:"errorMessage"`
	FilterReason *string        `gorm:"column:filterReason" json:"filterReason"`
	// Per-target results: WebhookDeliveryTargetResult[] ({ id, type, status, reason?, runId?, error? }).
	Targets     datatypes.JSON `gorm:"column:targets;default:'[]';not null" json:"targets"`
	CreatedAt   DateTime       `gorm:"column:createdAt;primaryKey;index:WebhookDelivery_webhookEndpointId_createdAt_idx;autoCreateTime:false" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	ProcessedAt *DateTime      `gorm:"column:processedAt" json:"processedAt"`
	// Retention class in days: the LIST partition key, stamped at ingest from the org's plan and encoded in the id.
	RetentionDays int `gorm:"column:retentionDays;primaryKey;default:30" json:"retentionDays"`
}

func (WebhookDelivery) TableName() string {
	return "WebhookDelivery"
}

func (m *WebhookDelivery) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WebhookDelivery) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Alert struct {
	ID                         string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID             string         `gorm:"column:organizationId;index:Alert_organizationId_idx;not null" json:"organizationId"`
	ProjectID                  string         `gorm:"column:projectId;index:Alert_projectId_idx;not null" json:"projectId"`
	EnvironmentID              string         `gorm:"column:environmentId;index:Alert_environmentId_lifecycle_idx;not null" json:"environmentId"`
	CapabilityKey              string         `gorm:"column:capabilityKey;not null" json:"capabilityKey"`
	Params                     datatypes.JSON `gorm:"column:params;not null" json:"params"`
	Condition                  datatypes.JSON `gorm:"column:condition;not null" json:"condition"`
	Query                      *string        `gorm:"column:query" json:"query"`
	ExpiresAt                  *DateTime      `gorm:"column:expiresAt;type:timestamptz(6)" json:"expiresAt"`
	Lifecycle                  *string        `gorm:"column:lifecycle;index:Alert_environmentId_lifecycle_idx;index:Alert_lifecycle_evaluationLeaseUntil_idx;default:'active';not null" json:"lifecycle"`
	Attention                  *string        `gorm:"column:attention;default:'clear';not null" json:"attention"`
	ObservedCondition          *string        `gorm:"column:observedCondition;default:'unknown';not null" json:"observedCondition"`
	ObservedConditionChangedAt *DateTime      `gorm:"column:observedConditionChangedAt;type:timestamptz(6)" json:"observedConditionChangedAt"`
	PendingMetSince            *DateTime      `gorm:"column:pendingMetSince;type:timestamptz(6)" json:"pendingMetSince"`
	EvaluatorState             datatypes.JSON `gorm:"column:evaluatorState" json:"evaluatorState"`
	EvaluationLeaseID          *string        `gorm:"column:evaluationLeaseId" json:"evaluationLeaseId"`
	EvaluationLeaseUntil       *DateTime      `gorm:"column:evaluationLeaseUntil;index:Alert_lifecycle_evaluationLeaseUntil_idx;type:timestamptz(6)" json:"evaluationLeaseUntil"`
	LastEvaluatedAt            *DateTime      `gorm:"column:lastEvaluatedAt;type:timestamptz(6)" json:"lastEvaluatedAt"`
	NextEvaluationAt           *DateTime      `gorm:"column:nextEvaluationAt;type:timestamptz(6)" json:"nextEvaluationAt"`
	Notifiers                  datatypes.JSON `gorm:"column:notifiers;default:'[]';not null" json:"notifiers"`
	Name                       *string        `gorm:"column:name" json:"name"`
	CreatedBy                  string         `gorm:"column:createdBy;not null" json:"createdBy"`
	CreatedAt                  DateTime       `gorm:"column:createdAt;type:timestamptz(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	ResolvedAt                 *DateTime      `gorm:"column:resolvedAt;type:timestamptz(6)" json:"resolvedAt"`
	ResolvedBy                 *string        `gorm:"column:resolvedBy" json:"resolvedBy"`
	StoppedAt                  *DateTime      `gorm:"column:stoppedAt;type:timestamptz(6)" json:"stoppedAt"`
	StoppedBy                  *string        `gorm:"column:stoppedBy" json:"stoppedBy"`
	LastOccurrenceAt           *DateTime      `gorm:"column:lastOccurrenceAt;type:timestamptz(6)" json:"lastOccurrenceAt"`
	OccurrencesSinceRaise      int            `gorm:"column:occurrencesSinceRaise;default:0;not null" json:"occurrencesSinceRaise"`
	OccurrencesTotal           int64          `gorm:"column:occurrencesTotal;default:0;not null" json:"occurrencesTotal"`
	GroupID                    *string        `gorm:"column:groupId;index:Alert_groupId_idx" json:"groupId"`
}

func (Alert) TableName() string {
	return "Alert"
}

func (m *Alert) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

// RANGE-partitioned by month on `at` (see migration).
type AlertEvent struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	AlertID       string         `gorm:"column:alertId;index:AlertEvent_alertId_at_idx;not null" json:"alertId"`
	Type          string         `gorm:"column:type;not null" json:"type"`
	At            DateTime       `gorm:"column:at;primaryKey;index:AlertEvent_alertId_at_idx;type:timestamptz(6)" json:"at"`
	Facts         datatypes.JSON `gorm:"column:facts;default:'{}';not null" json:"facts"`
	ActorID       *string        `gorm:"column:actorId" json:"actorId"`
	SourceEventID *string        `gorm:"column:sourceEventId" json:"sourceEventId"`
}

func (AlertEvent) TableName() string {
	return "AlertEvent"
}

func (m *AlertEvent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.At.IsZero() {
		m.At = DateTime{Time: now}
	}
	return nil
}

type AlertOccurrenceSeen struct {
	AlertID       string   `gorm:"column:alertId;primaryKey" json:"alertId"`
	SourceEventID string   `gorm:"column:sourceEventId;primaryKey" json:"sourceEventId"`
	RowTime       DateTime `gorm:"column:rowTime;type:timestamptz(6);not null" json:"rowTime"`
	SeenAt        DateTime `gorm:"column:seenAt;type:timestamptz(6);not null" json:"seenAt"`
	ExpiresAt     DateTime `gorm:"column:expiresAt;index:AlertOccurrenceSeen_expiresAt_idx;type:timestamptz(6);not null" json:"expiresAt"`
}

func (AlertOccurrenceSeen) TableName() string {
	return "AlertOccurrenceSeen"
}

func (m *AlertOccurrenceSeen) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.SeenAt.IsZero() {
		m.SeenAt = DateTime{Time: now}
	}
	return nil
}

// RANGE-partitioned by month on `eventAt` (see migration).
type AlertDelivery struct {
	ID               string         `gorm:"column:id;primaryKey" json:"id"`
	AlertEventID     string         `gorm:"column:alertEventId;uniqueIndex:AlertDelivery_alertEventId_notifierId_eventAt_key;not null" json:"alertEventId"`
	EventAt          DateTime       `gorm:"column:eventAt;primaryKey;uniqueIndex:AlertDelivery_alertEventId_notifierId_eventAt_key;type:timestamptz(6)" json:"eventAt"`
	NotifierID       string         `gorm:"column:notifierId;uniqueIndex:AlertDelivery_alertEventId_notifierId_eventAt_key;not null" json:"notifierId"`
	Notifier         datatypes.JSON `gorm:"column:notifier;not null" json:"notifier"`
	Status           *string        `gorm:"column:status;index:AlertDelivery_status_createdAt_idx;index:AlertDelivery_status_claimedAt_idx;default:'pending';not null" json:"status"`
	ClaimID          *string        `gorm:"column:claimId" json:"claimId"`
	ClaimedAt        *DateTime      `gorm:"column:claimedAt;index:AlertDelivery_status_claimedAt_idx;type:timestamptz(6)" json:"claimedAt"`
	DeliveredAt      *DateTime      `gorm:"column:deliveredAt;type:timestamptz(6)" json:"deliveredAt"`
	FailureReason    *string        `gorm:"column:failureReason" json:"failureReason"`
	ProviderMetadata datatypes.JSON `gorm:"column:providerMetadata" json:"providerMetadata"`
	Attempts         int            `gorm:"column:attempts;default:0;not null" json:"attempts"`
	FirstAttemptAt   *DateTime      `gorm:"column:firstAttemptAt;type:timestamptz(6)" json:"firstAttemptAt"`
	CreatedAt        DateTime       `gorm:"column:createdAt;index:AlertDelivery_status_createdAt_idx;type:timestamptz(6);autoCreateTime:false;not null" json:"createdAt"`
}

func (AlertDelivery) TableName() string {
	return "AlertDelivery"
}

func (m *AlertDelivery) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type PlaygroundConversation struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The chat session ID used by the transport
	ChatID string `gorm:"column:chatId;uniqueIndex:PlaygroundConversation_chatId_runtimeEnvironmentId_key;not null" json:"chatId"`
	// User-editable conversation title (auto-generated from first message)
	Title *string `gorm:"column:title;default:'New conversation';not null" json:"title"`
	// Which agent this conversation is with
	AgentSlug string `gorm:"column:agentSlug;index:PlaygroundConversation_runtimeEnvironmentId_agentSlug_updatedAt_idx;not null" json:"agentSlug"`
	// The current active run backing this conversation (null if no run yet)
	RunID *string `gorm:"column:runId" json:"runId"`
	// The client data JSON used for this conversation
	ClientData datatypes.JSON `gorm:"column:clientData" json:"clientData"`
	// Accumulated UIMessages from completed turns (for resume without stream replay)
	Messages datatypes.JSON `gorm:"column:messages" json:"messages"`
	// Last SSE event ID — resume from this position to avoid replaying old turns
	LastEventID          *string `gorm:"column:lastEventId" json:"lastEventId"`
	ProjectID            string  `gorm:"column:projectId;not null" json:"projectId"`
	RuntimeEnvironmentID string  `gorm:"column:runtimeEnvironmentId;uniqueIndex:PlaygroundConversation_chatId_runtimeEnvironmentId_key;index:PlaygroundConversation_runtimeEnvironmentId_agentSlug_updatedAt_idx;index:PlaygroundConversation_userId_runtimeEnvironmentId_idx;not null" json:"runtimeEnvironmentId"`
	// The user who started this conversation
	UserID             string   `gorm:"column:userId;index:PlaygroundConversation_userId_runtimeEnvironmentId_idx;not null" json:"userId"`
	CreatedAt          DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime `gorm:"column:updatedAt;index:PlaygroundConversation_runtimeEnvironmentId_agentSlug_updatedAt_idx;autoUpdateTime:false;not null" json:"updatedAt"`
	Run                *TaskRun `gorm:"foreignKey:RunID"`
	Project            *Project
	RuntimeEnvironment *RuntimeEnvironment
}

func (PlaygroundConversation) TableName() string {
	return "PlaygroundConversation"
}

func (m *PlaygroundConversation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PlaygroundConversation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
// The row is essentially static — no status, no counters, no pointers. No
// foreign keys: project/runtimeEnvironment/organization ids are plain
// scalar columns (matches TaskRun pattern). List-style queries are served
// from ClickHouse, not Postgres, so only point-lookup indexes live here.
type Session struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID string `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	// User-supplied identifier scoped to the environment. Used for
	// idempotent upsert and for resolving sessions via the public API.
	ExternalID *string `gorm:"column:externalId;uniqueIndex:Session_runtimeEnvironmentId_externalId_key" json:"externalId"`
	// Plain string — intentionally not an enum.
	Type string `gorm:"column:type;not null" json:"type"`
	// Denormalized scoping columns — no FK relations.
	ProjectID            string `gorm:"column:projectId;not null" json:"projectId"`
	RuntimeEnvironmentID string `gorm:"column:runtimeEnvironmentId;uniqueIndex:Session_runtimeEnvironmentId_externalId_key;not null" json:"runtimeEnvironmentId"`
	EnvironmentType      string `gorm:"column:environmentType;not null" json:"environmentType"`
	OrganizationID       string `gorm:"column:organizationId;not null" json:"organizationId"`
	// Task this session triggers runs against. Required — Sessions are
	// task-bound: creating a session also triggers its first run, and
	// every subsequent re-trigger uses this same identifier.
	TaskIdentifier string `gorm:"column:taskIdentifier;not null" json:"taskIdentifier"`
	// Trigger config used for every run this session schedules. Shape
	// (validated at the route layer, opaque to the DB):
	// { basePayload: object, machine?: string, queue?: string,
	// tags?: string[], maxAttempts?: number,
	// idleTimeoutInSeconds?: number }
	// `basePayload` carries the customer's client-data; runtime fields
	// (chatId, messages, trigger) are merged at trigger time.
	TriggerConfig datatypes.JSON `gorm:"column:triggerConfig;not null" json:"triggerConfig"`
	// Whether this session was created from the Test/playground UI rather
	// than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
	IsTest   bool           `gorm:"column:isTest;default:false;not null" json:"isTest"`
	Tags     []string       `gorm:"column:tags;serializer:json;not null" json:"tags"`
	Metadata datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	// Live run pointer — non-FK so run deletion never cascades. Can lag
	// reality; the `.in/append` handler re-checks the snapshot status
	// before reusing it.
	CurrentRunID *string `gorm:"column:currentRunId;index:Session_currentRunId_idx" json:"currentRunId"`
	// Monotonic counter used for optimistic locking on `currentRunId`
	// swaps. Bumped atomically alongside any update that changes
	// `currentRunId`.
	CurrentRunVersion int `gorm:"column:currentRunVersion;default:0;not null" json:"currentRunVersion"`
	// Terminal markers — written once, never flipped back.
	ClosedAt     *DateTime `gorm:"column:closedAt" json:"closedAt"`
	ClosedReason *string   `gorm:"column:closedReason" json:"closedReason"`
	ExpiresAt    *DateTime `gorm:"column:expiresAt;index:Session_expiresAt_idx" json:"expiresAt"`
	CreatedAt    DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// S2 basin where this session's stream pair lives. Stamped at create
	// time from `Organization.streamBasinName` so reads can resolve the
	// basin without joining org. Null when the org has no per-org basin
	// (OSS, or pre-backfill); reads fall back to the global basin.
	StreamBasinName *string `gorm:"column:streamBasinName" json:"streamBasinName"`
	// Storage URI (with protocol prefix) for this session's chat.agent
	// snapshot blob. Set on first snapshot write. Null = pre-column session,
	// fall back to computed default path.
	ChatSnapshotStoragePath *string      `gorm:"column:chatSnapshotStoragePath" json:"chatSnapshotStoragePath"`
	Runs                    []SessionRun `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE"`
}

func (Session) TableName() string {
	return "Session"
}

func (m *Session) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Session) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Historical record of every run a Session has owned. Append-only —
// rows are inserted on each `ensureRunForSession` claim, never updated.
// Lets us reconstruct the run timeline of a chat for debugging /
// dashboard surfaces. The relation cascades on Session delete (tied to
// the session lifecycle) but `runId` is a plain string column with no
// FK to TaskRun so run pruning is independent.
type SessionRun struct {
	ID        string `gorm:"column:id;primaryKey;index:SessionRun_sessionId_triggeredAt_id_idx" json:"id"`
	SessionID string `gorm:"column:sessionId;index:SessionRun_sessionId_idx;index:SessionRun_sessionId_triggeredAt_id_idx;not null" json:"sessionId"`
	// TaskRun.id (no FK — runs may be archived independently of session history)
	RunID string `gorm:"column:runId;uniqueIndex;not null" json:"runId"`
	// One of: "initial" | "continuation" | "upgrade" | "manual".
	// Plain string for forward-compat with future trigger reasons.
	Reason      string   `gorm:"column:reason;not null" json:"reason"`
	TriggeredAt DateTime `gorm:"column:triggeredAt;index:SessionRun_sessionId_triggeredAt_id_idx;not null" json:"triggeredAt"`
	Session     *Session
}

func (SessionRun) TableName() string {
	return "SessionRun"
}

func (m *SessionRun) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.TriggeredAt.IsZero() {
		m.TriggeredAt = DateTime{Time: now}
	}
	return nil
}

type TaskRun struct {
	ID                      string    `gorm:"column:id;primaryKey" json:"id"`
	Number                  int       `gorm:"column:number;default:0;not null" json:"number"`
	FriendlyID              string    `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Engine                  string    `gorm:"column:engine;default:'V1';not null" json:"engine"`
	Status                  string    `gorm:"column:status;default:'PENDING';not null" json:"status"`
	StatusReason            *string   `gorm:"column:statusReason" json:"statusReason"`
	IdempotencyKey          *string   `gorm:"column:idempotencyKey;uniqueIndex:TaskRun_runtimeEnvironmentId_taskIdentifier_idempotencyKey_key" json:"idempotencyKey"`
	IdempotencyKeyExpiresAt *DateTime `gorm:"column:idempotencyKeyExpiresAt" json:"idempotencyKeyExpiresAt"`
	// Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" }
	IdempotencyKeyOptions datatypes.JSON `gorm:"column:idempotencyKeyOptions" json:"idempotencyKeyOptions"`
	// Debounce options: { key: string, delay: string, createdAt: Date }
	Debounce             datatypes.JSON `gorm:"column:debounce" json:"debounce"`
	TaskIdentifier       string         `gorm:"column:taskIdentifier;uniqueIndex:TaskRun_runtimeEnvironmentId_taskIdentifier_idempotencyKey_key;not null" json:"taskIdentifier"`
	IsTest               bool           `gorm:"column:isTest;default:false;not null" json:"isTest"`
	Payload              string         `gorm:"column:payload;not null" json:"payload"`
	PayloadType          *string        `gorm:"column:payloadType;default:'application/json';not null" json:"payloadType"`
	Context              datatypes.JSON `gorm:"column:context" json:"context"`
	TraceContext         datatypes.JSON `gorm:"column:traceContext" json:"traceContext"`
	TraceID              string         `gorm:"column:traceId;not null" json:"traceId"`
	SpanID               string         `gorm:"column:spanId;index:TaskRun_spanId_idx;not null" json:"spanId"`
	RuntimeEnvironmentID string         `gorm:"column:runtimeEnvironmentId;uniqueIndex:TaskRun_runtimeEnvironmentId_taskIdentifier_idempotencyKey_key;index:TaskRun_runtimeEnvironmentId_batchId_idx;index:TaskRun_runtimeEnvironmentId_createdAt_idx;not null" json:"runtimeEnvironmentId"`
	EnvironmentType      *string        `gorm:"column:environmentType" json:"environmentType"`
	ProjectID            string         `gorm:"column:projectId;not null" json:"projectId"`
	OrganizationID       *string        `gorm:"column:organizationId" json:"organizationId"`
	Queue                string         `gorm:"column:queue;not null" json:"queue"`
	LockedQueueID        *string        `gorm:"column:lockedQueueId" json:"lockedQueueId"`
	// The main queue that this run is part of
	WorkerQueue *string `gorm:"column:masterQueue;default:'main';not null" json:"masterQueue"`
	// User-facing geo region, stamped at trigger; workerQueue is where it actually ran.
	Region *string `gorm:"column:region" json:"region"`
	// @deprecated
	SecondaryMasterQueue *string `gorm:"column:secondaryMasterQueue" json:"secondaryMasterQueue"`
	// From engine v2+ this will be defined after a run has been dequeued (starting at 1)
	AttemptNumber *int     `gorm:"column:attemptNumber" json:"attemptNumber"`
	CreatedAt     DateTime `gorm:"column:createdAt;index:TaskRun_runtimeEnvironmentId_createdAt_idx;index:TaskRun_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// Denormized column that holds the raw tags
	RunTags []string `gorm:"column:runTags;index:TaskRun_runTags_idx;serializer:json;not null" json:"runTags"`
	// Denormalized version of the background worker task
	TaskVersion *string `gorm:"column:taskVersion" json:"taskVersion"`
	SdkVersion  *string `gorm:"column:sdkVersion" json:"sdkVersion"`
	CliVersion  *string `gorm:"column:cliVersion" json:"cliVersion"`
	// startedAt marks the point at which a run is dequeued from MarQS
	StartedAt *DateTime `gorm:"column:startedAt" json:"startedAt"`
	// executedAt is set when the first attempt is about to execute
	ExecutedAt        *DateTime `gorm:"column:executedAt" json:"executedAt"`
	CompletedAt       *DateTime `gorm:"column:completedAt" json:"completedAt"`
	MachinePreset     *string   `gorm:"column:machinePreset" json:"machinePreset"`
	UsageDurationMs   int       `gorm:"column:usageDurationMs;default:0;not null" json:"usageDurationMs"`
	CostInCents       float64   `gorm:"column:costInCents;default:0;not null" json:"costInCents"`
	BaseCostInCents   float64   `gorm:"column:baseCostInCents;default:0;not null" json:"baseCostInCents"`
	LockedAt          *DateTime `gorm:"column:lockedAt" json:"lockedAt"`
	LockedByID        *string   `gorm:"column:lockedById" json:"lockedById"`
	LockedToVersionID *string   `gorm:"column:lockedToVersionId" json:"lockedToVersionId"`
	// The "priority" of the run. This is just a negative offset in ms for the queue timestamp
	// E.g. a value of 60_000 would put the run into the queue 60s ago.
	PriorityMs     int     `gorm:"column:priorityMs;default:0;not null" json:"priorityMs"`
	ConcurrencyKey *string `gorm:"column:concurrencyKey" json:"concurrencyKey"`
	// Gates for this run: other queues it must also hold a concurrency slot in while
	// executing. Shape: [{ queue: string, concurrencyKey?: string }]
	Gates             datatypes.JSON `gorm:"column:gates" json:"gates"`
	DelayUntil        *DateTime      `gorm:"column:delayUntil" json:"delayUntil"`
	QueuedAt          *DateTime      `gorm:"column:queuedAt" json:"queuedAt"`
	TTL               *string        `gorm:"column:ttl" json:"ttl"`
	ExpiredAt         *DateTime      `gorm:"column:expiredAt" json:"expiredAt"`
	MaxAttempts       *int           `gorm:"column:maxAttempts" json:"maxAttempts"`
	LockedRetryConfig datatypes.JSON `gorm:"column:lockedRetryConfig" json:"lockedRetryConfig"`
	// optional token that can be used to authenticate the task run
	OneTimeUseToken *string `gorm:"column:oneTimeUseToken;uniqueIndex" json:"oneTimeUseToken"`
	// Where the logs are stored
	TaskEventStore                *string   `gorm:"column:taskEventStore;default:'taskEvent';not null" json:"taskEventStore"`
	QueueTimestamp                *DateTime `gorm:"column:queueTimestamp" json:"queueTimestamp"`
	ScheduleInstanceID            *string   `gorm:"column:scheduleInstanceId" json:"scheduleInstanceId"`
	ScheduleID                    *string   `gorm:"column:scheduleId" json:"scheduleId"`
	BulkActionGroupIds            []string  `gorm:"column:bulkActionGroupIds;serializer:json;not null" json:"bulkActionGroupIds"`
	LogsDeletedAt                 *DateTime `gorm:"column:logsDeletedAt" json:"logsDeletedAt"`
	ReplayedFromTaskRunFriendlyID *string   `gorm:"column:replayedFromTaskRunFriendlyId" json:"replayedFromTaskRunFriendlyId"`
	RootTaskRunID                 *string   `gorm:"column:rootTaskRunId" json:"rootTaskRunId"`
	ParentTaskRunID               *string   `gorm:"column:parentTaskRunId;index:TaskRun_parentTaskRunId_idx" json:"parentTaskRunId"`
	ParentTaskRunAttemptID        *string   `gorm:"column:parentTaskRunAttemptId" json:"parentTaskRunAttemptId"`
	BatchID                       *string   `gorm:"column:batchId;index:TaskRun_runtimeEnvironmentId_batchId_idx" json:"batchId"`
	// whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait
	ResumeParentOnCompletion bool `gorm:"column:resumeParentOnCompletion;default:false;not null" json:"resumeParentOnCompletion"`
	// The depth of this task run in the task run hierarchy
	Depth int `gorm:"column:depth;default:0;not null" json:"depth"`
	// The span ID of the "trigger" span in the parent task run
	ParentSpanID *string `gorm:"column:parentSpanId;index:TaskRun_parentSpanId_idx" json:"parentSpanId"`
	// Holds the state of the run chain for deadlock detection
	RunChainState datatypes.JSON `gorm:"column:runChainState" json:"runChainState"`
	// seed run metadata
	SeedMetadata     *string `gorm:"column:seedMetadata" json:"seedMetadata"`
	SeedMetadataType *string `gorm:"column:seedMetadataType;default:'application/json';not null" json:"seedMetadataType"`
	// Run metadata
	Metadata        *string `gorm:"column:metadata" json:"metadata"`
	MetadataType    *string `gorm:"column:metadataType;default:'application/json';not null" json:"metadataType"`
	MetadataVersion *int    `gorm:"column:metadataVersion;default:1;not null" json:"metadataVersion"`
	// Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId
	Annotations datatypes.JSON `gorm:"column:annotations" json:"annotations"`
	// Whether the latest attempt was a warm start. Null until first attempt starts.
	IsWarmStart *bool `gorm:"column:isWarmStart" json:"isWarmStart"`
	// Run output
	Output     *string `gorm:"column:output" json:"output"`
	OutputType *string `gorm:"column:outputType;default:'application/json';not null" json:"outputType"`
	// Run error
	Error datatypes.JSON `gorm:"column:error" json:"error"`
	// Organization's billing plan type (cached for fallback when billing API fails)
	PlanType             *string `gorm:"column:planType" json:"planType"`
	MaxDurationInSeconds *int    `gorm:"column:maxDurationInSeconds" json:"maxDurationInSeconds"`
	// The version of the realtime streams implementation used by the run
	RealtimeStreamsVersion *string `gorm:"column:realtimeStreamsVersion;default:'v1';not null" json:"realtimeStreamsVersion"`
	// Store the stream keys that are being used by the run
	RealtimeStreams []string `gorm:"column:realtimeStreams;serializer:json;not null" json:"realtimeStreams"`
	// S2 basin where this run's realtime streams live. Stamped at create
	// time from `Organization.streamBasinName` so reads can resolve the
	// basin without joining org. Null when the org has no per-org basin
	// (OSS, or pre-backfill); reads fall back to the global basin.
	StreamBasinName    *string `gorm:"column:streamBasinName" json:"streamBasinName"`
	RuntimeEnvironment *RuntimeEnvironment
	Project            *Project
	LockedBy           *BackgroundWorkerTask `gorm:"foreignKey:LockedByID"`
	LockedToVersion    *BackgroundWorker     `gorm:"foreignKey:LockedToVersionID"`
	// This represents the original task that that was triggered outside of a Trigger.dev task
	RootTaskRun *TaskRun `gorm:"foreignKey:RootTaskRunID"`
	// The immediate parent run of this task run
	ParentTaskRun *TaskRun `gorm:"foreignKey:ParentTaskRunID"`
	// The immediate parent attempt of this task run
	ParentTaskRunAttempt *TaskRunAttempt `gorm:"foreignKey:ParentTaskRunAttemptID"`
	// The batch run that this task run is a part of
	Batch       *BatchTaskRun    `gorm:"foreignKey:BatchID"`
	Attempts    []TaskRunAttempt `gorm:"foreignKey:TaskRunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Checkpoints []Checkpoint     `gorm:"foreignKey:RunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// If there are any blocked waitpoints, the run won't be executed
	BlockedByWaitpoints        []TaskRunWaitpoint         `gorm:"foreignKey:TaskRunID"`
	BatchItems                 []BatchTaskRunItem         `gorm:"foreignKey:TaskRunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CheckpointRestoreEvent     []CheckpointRestoreEvent   `gorm:"foreignKey:RunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExecutionSnapshots         []TaskRunExecutionSnapshot `gorm:"foreignKey:RunID"`
	SourceBulkActionItems      []BulkActionItem           `gorm:"foreignKey:SourceRunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DestinationBulkActionItems []BulkActionItem           `gorm:"foreignKey:DestinationRunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// The root run will have a list of all the descendant runs, children, grand children, etc.
	DescendantRuns []TaskRun `gorm:"foreignKey:RootTaskRunID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
	// The immediate child runs of this task run
	ChildRuns               []TaskRun                `gorm:"foreignKey:ParentTaskRunID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
	PlaygroundConversations []PlaygroundConversation `gorm:"foreignKey:RunID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	// When this run is finished, the waitpoint will be marked as completed
	AssociatedWaitpoint *Waitpoint         `gorm:"foreignKey:CompletedByTaskRunID;constraint:OnDelete:SET NULL"`
	Dependency          *TaskRunDependency `gorm:"foreignKey:TaskRunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// All waitpoints that blocked this run at some point, used for display purposes
	ConnectedWaitpoints []Waitpoint `gorm:"many2many:_WaitpointRunConnections;joinForeignKey:A;joinReferences:B"`
}

func (TaskRun) TableName() string {
	return "TaskRun"
}

func (m *TaskRun) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskRun) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TaskRunTemplate struct {
	ID                 string   `gorm:"column:id;primaryKey" json:"id"`
	TaskSlug           string   `gorm:"column:taskSlug;index:TaskRunTemplate_projectId_taskSlug_triggerSource_createdAt_idx;not null" json:"taskSlug"`
	TriggerSource      string   `gorm:"column:triggerSource;index:TaskRunTemplate_projectId_taskSlug_triggerSource_createdAt_idx;not null" json:"triggerSource"`
	Label              string   `gorm:"column:label;not null" json:"label"`
	Payload            *string  `gorm:"column:payload" json:"payload"`
	PayloadType        *string  `gorm:"column:payloadType;default:'application/json';not null" json:"payloadType"`
	Metadata           *string  `gorm:"column:metadata" json:"metadata"`
	MetadataType       *string  `gorm:"column:metadataType;default:'application/json';not null" json:"metadataType"`
	Queue              *string  `gorm:"column:queue" json:"queue"`
	ConcurrencyKey     *string  `gorm:"column:concurrencyKey" json:"concurrencyKey"`
	TTLSeconds         *int     `gorm:"column:ttlSeconds" json:"ttlSeconds"`
	MaxAttempts        *int     `gorm:"column:maxAttempts" json:"maxAttempts"`
	MaxDurationSeconds *int     `gorm:"column:maxDurationSeconds" json:"maxDurationSeconds"`
	Tags               []string `gorm:"column:tags;serializer:json;not null" json:"tags"`
	MachinePreset      *string  `gorm:"column:machinePreset" json:"machinePreset"`
	ProjectID          string   `gorm:"column:projectId;index:TaskRunTemplate_projectId_taskSlug_triggerSource_createdAt_idx;not null" json:"projectId"`
	OrganizationID     string   `gorm:"column:organizationId;not null" json:"organizationId"`
	CreatedAt          DateTime `gorm:"column:createdAt;index:TaskRunTemplate_projectId_taskSlug_triggerSource_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (TaskRunTemplate) TableName() string {
	return "TaskRunTemplate"
}

func (m *TaskRunTemplate) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskRunTemplate) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Used by the RunEngine during TaskRun execution
// It has the required information to transactionally progress a run through states,
// and prevent side effects like heartbeats failing a run that has progressed.
// It is optimised for performance and is designed to be cleared at some point,
// so there are no cascading relationships to other models.
type TaskRunExecutionSnapshot struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// This should always be 2+ (V1 didn't use the run engine or snapshots)
	Engine string `gorm:"column:engine;default:'V2';not null" json:"engine"`
	// The execution status
	ExecutionStatus string `gorm:"column:executionStatus;not null" json:"executionStatus"`
	// For debugging
	Description string `gorm:"column:description;not null" json:"description"`
	// We store invalid snapshots as a record of the run state when we tried to move
	IsValid *bool   `gorm:"column:isValid;index:TaskRunExecutionSnapshot_runId_isValid_createdAt_idx;default:true;not null" json:"isValid"`
	Error   *string `gorm:"column:error" json:"error"`
	// The previous snapshot ID
	PreviousSnapshotID *string `gorm:"column:previousSnapshotId" json:"previousSnapshotId"`
	// Run
	RunID     string  `gorm:"column:runId;index:TaskRunExecutionSnapshot_runId_isValid_createdAt_idx;not null" json:"runId"`
	RunStatus string  `gorm:"column:runStatus;not null" json:"runStatus"`
	BatchID   *string `gorm:"column:batchId" json:"batchId"`
	// This is the current run attempt number. Users can define how many attempts they want for a run.
	AttemptNumber *int `gorm:"column:attemptNumber" json:"attemptNumber"`
	// Environment
	EnvironmentID   string `gorm:"column:environmentId;not null" json:"environmentId"`
	EnvironmentType string `gorm:"column:environmentType;not null" json:"environmentType"`
	ProjectID       string `gorm:"column:projectId;not null" json:"projectId"`
	OrganizationID  string `gorm:"column:organizationId;not null" json:"organizationId"`
	// An array of waitpoint IDs in the correct order, used for batches
	CompletedWaitpointOrder []string `gorm:"column:completedWaitpointOrder;serializer:json;not null" json:"completedWaitpointOrder"`
	// Checkpoint
	CheckpointID *string `gorm:"column:checkpointId" json:"checkpointId"`
	// Worker
	WorkerID        *string   `gorm:"column:workerId" json:"workerId"`
	RunnerID        *string   `gorm:"column:runnerId" json:"runnerId"`
	CreatedAt       DateTime  `gorm:"column:createdAt;index:TaskRunExecutionSnapshot_runId_isValid_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	LastHeartbeatAt *DateTime `gorm:"column:lastHeartbeatAt" json:"lastHeartbeatAt"`
	// Metadata used by various systems in the run engine
	Metadata   datatypes.JSON     `gorm:"column:metadata" json:"metadata"`
	Run        *TaskRun           `gorm:"foreignKey:RunID"`
	Checkpoint *TaskRunCheckpoint `gorm:"foreignKey:CheckpointID"`
	// Waitpoints that have been completed for this execution
	CompletedWaitpoints []Waitpoint `gorm:"many2many:_completedWaitpoints;joinForeignKey:A;joinReferences:B"`
}

func (TaskRunExecutionSnapshot) TableName() string {
	return "TaskRunExecutionSnapshot"
}

func (m *TaskRunExecutionSnapshot) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskRunExecutionSnapshot) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TaskRunCheckpoint struct {
	ID                   string   `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID           string   `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Type                 string   `gorm:"column:type;not null" json:"type"`
	Location             string   `gorm:"column:location;not null" json:"location"`
	ImageRef             *string  `gorm:"column:imageRef" json:"imageRef"`
	Reason               *string  `gorm:"column:reason" json:"reason"`
	Metadata             *string  `gorm:"column:metadata" json:"metadata"`
	ProjectID            string   `gorm:"column:projectId;not null" json:"projectId"`
	RuntimeEnvironmentID string   `gorm:"column:runtimeEnvironmentId;not null" json:"runtimeEnvironmentId"`
	CreatedAt            DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project              *Project
	RuntimeEnvironment   *RuntimeEnvironment
	ExecutionSnapshot    []TaskRunExecutionSnapshot `gorm:"foreignKey:CheckpointID"`
}

func (TaskRunCheckpoint) TableName() string {
	return "TaskRunCheckpoint"
}

func (m *TaskRunCheckpoint) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskRunCheckpoint) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A Waitpoint blocks a run from continuing until it's completed
// If there's a waitpoint blocking a run, it shouldn't be in the queue
type Waitpoint struct {
	ID          string    `gorm:"column:id;primaryKey;index:Waitpoint_environmentId_type_id_idx" json:"id"`
	FriendlyID  string    `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Type        string    `gorm:"column:type;index:Waitpoint_environmentId_type_createdAt_idx;index:Waitpoint_environmentId_type_status_idx;index:Waitpoint_environmentId_type_id_idx;not null" json:"type"`
	Status      string    `gorm:"column:status;index:Waitpoint_environmentId_type_status_idx;default:'PENDING';not null" json:"status"`
	CompletedAt *DateTime `gorm:"column:completedAt" json:"completedAt"`
	// If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type
	IdempotencyKey string `gorm:"column:idempotencyKey;uniqueIndex:Waitpoint_environmentId_idempotencyKey_key;not null" json:"idempotencyKey"`
	// If this is true then we can show it in the dashboard/return it from the SDK
	UserProvidedIdempotencyKey bool `gorm:"column:userProvidedIdempotencyKey;not null" json:"userProvidedIdempotencyKey"`
	// If there's a user provided idempotency key, this is the time it expires at
	IdempotencyKeyExpiresAt *DateTime `gorm:"column:idempotencyKeyExpiresAt" json:"idempotencyKeyExpiresAt"`
	// If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
	// Clearing an idempotencyKey is useful for debounce or cancelling child runs.
	// This is a workaround because Prisma doesn't support partial indexes.
	InactiveIdempotencyKey *string `gorm:"column:inactiveIdempotencyKey" json:"inactiveIdempotencyKey"`
	// If it's a RUN type waitpoint, this is the associated run
	CompletedByTaskRunID *string `gorm:"column:completedByTaskRunId;uniqueIndex" json:"completedByTaskRunId"`
	// If it's a DATETIME type waitpoint, this is the date.
	// If it's a MANUAL waitpoint, this can be set as the `timeout`.
	CompletedAfter *DateTime `gorm:"column:completedAfter" json:"completedAfter"`
	// If it's a BATCH type waitpoint, this is the associated batch
	CompletedByBatchID *string `gorm:"column:completedByBatchId;index:Waitpoint_completedByBatchId_idx" json:"completedByBatchId"`
	// When completed, an output can be stored here
	Output        *string  `gorm:"column:output" json:"output"`
	OutputType    *string  `gorm:"column:outputType;default:'application/json';not null" json:"outputType"`
	OutputIsError bool     `gorm:"column:outputIsError;default:false;not null" json:"outputIsError"`
	ProjectID     string   `gorm:"column:projectId;not null" json:"projectId"`
	EnvironmentID string   `gorm:"column:environmentId;uniqueIndex:Waitpoint_environmentId_idempotencyKey_key;index:Waitpoint_environmentId_type_createdAt_idx;index:Waitpoint_environmentId_type_status_idx;index:Waitpoint_environmentId_type_id_idx;not null" json:"environmentId"`
	CreatedAt     DateTime `gorm:"column:createdAt;index:Waitpoint_environmentId_type_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// Denormized column that holds the raw tags
	// Denormalized column that holds the raw tags
	Tags               []string      `gorm:"column:tags;serializer:json;not null" json:"tags"`
	CompletedByTaskRun *TaskRun      `gorm:"foreignKey:CompletedByTaskRunID"`
	CompletedByBatch   *BatchTaskRun `gorm:"foreignKey:CompletedByBatchID"`
	Project            *Project
	Environment        *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
	// The runs this waitpoint is blocking
	BlockingTaskRuns []TaskRunWaitpoint `gorm:"foreignKey:WaitpointID"`
	// All runs that have ever been blocked by this waitpoint, used for display purposes
	ConnectedRuns []TaskRun `gorm:"many2many:_WaitpointRunConnections;joinForeignKey:B;joinReferences:A"`
	// When a waitpoint is complete
	CompletedExecutionSnapshots []TaskRunExecutionSnapshot `gorm:"many2many:_completedWaitpoints;joinForeignKey:B;joinReferences:A"`
}

func (Waitpoint) TableName() string {
	return "Waitpoint"
}

func (m *Waitpoint) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Waitpoint) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TaskRunWaitpoint struct {
	ID          string `gorm:"column:id;primaryKey" json:"id"`
	TaskRunID   string `gorm:"column:taskRunId;uniqueIndex:TaskRunWaitpoint_taskRunId_waitpointId_batchIndex_key;index:TaskRunWaitpoint_taskRunId_idx;not null" json:"taskRunId"`
	WaitpointID string `gorm:"column:waitpointId;uniqueIndex:TaskRunWaitpoint_taskRunId_waitpointId_batchIndex_key;index:TaskRunWaitpoint_waitpointId_idx;not null" json:"waitpointId"`
	ProjectID   string `gorm:"column:projectId;not null" json:"projectId"`
	// This span id is completed when the waitpoint is completed. This is used with cached runs (idempotent)
	SpanIDToComplete *string  `gorm:"column:spanIdToComplete" json:"spanIdToComplete"`
	BatchID          *string  `gorm:"column:batchId" json:"batchId"`
	BatchIndex       *int     `gorm:"column:batchIndex;uniqueIndex:TaskRunWaitpoint_taskRunId_waitpointId_batchIndex_key" json:"batchIndex"`
	CreatedAt        DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	TaskRun          *TaskRun
	Waitpoint        *Waitpoint
	Project          *Project
	Batch            *BatchTaskRun `gorm:"foreignKey:BatchID"`
}

func (TaskRunWaitpoint) TableName() string {
	return "TaskRunWaitpoint"
}

func (m *TaskRunWaitpoint) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskRunWaitpoint) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WaitpointTag struct {
	ID            string              `gorm:"column:id;primaryKey" json:"id"`
	Name          string              `gorm:"column:name;uniqueIndex:WaitpointTag_environmentId_name_key;not null" json:"name"`
	EnvironmentID string              `gorm:"column:environmentId;uniqueIndex:WaitpointTag_environmentId_name_key;not null" json:"environmentId"`
	ProjectID     string              `gorm:"column:projectId;not null" json:"projectId"`
	CreatedAt     DateTime            `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Environment   *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
	Project       *Project
}

func (WaitpointTag) TableName() string {
	return "WaitpointTag"
}

func (m *WaitpointTag) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type FeatureFlag struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	Key       string         `gorm:"column:key;uniqueIndex;not null" json:"key"`
	Value     datatypes.JSON `gorm:"column:value" json:"value"`
	CreatedAt DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (FeatureFlag) TableName() string {
	return "FeatureFlag"
}

func (m *FeatureFlag) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*FeatureFlag) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkerInstance struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// For example "worker-1"
	Name string `gorm:"column:name;not null" json:"name"`
	// If managed, it will default to the name, e.g. "worker-1"
	// If unmanged, it will be prefixed with the deployment ID e.g. "deploy-123-worker-1"
	ResourceIdentifier string               `gorm:"column:resourceIdentifier;uniqueIndex:WorkerInstance_workerGroupId_resourceIdentifier_key;not null" json:"resourceIdentifier"`
	Metadata           datatypes.JSON       `gorm:"column:metadata" json:"metadata"`
	WorkerGroupID      string               `gorm:"column:workerGroupId;uniqueIndex:WorkerInstance_workerGroupId_resourceIdentifier_key;not null" json:"workerGroupId"`
	OrganizationID     *string              `gorm:"column:organizationId" json:"organizationId"`
	ProjectID          *string              `gorm:"column:projectId" json:"projectId"`
	EnvironmentID      *string              `gorm:"column:environmentId" json:"environmentId"`
	DeploymentID       *string              `gorm:"column:deploymentId" json:"deploymentId"`
	CreatedAt          DateTime             `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime             `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	LastDequeueAt      *DateTime            `gorm:"column:lastDequeueAt" json:"lastDequeueAt"`
	LastHeartbeatAt    *DateTime            `gorm:"column:lastHeartbeatAt" json:"lastHeartbeatAt"`
	WorkerGroup        *WorkerInstanceGroup `gorm:"foreignKey:WorkerGroupID"`
	Organization       *Organization
	Project            *Project
	Environment        *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
	Deployment         *WorkerDeployment   `gorm:"foreignKey:DeploymentID"`
}

func (WorkerInstance) TableName() string {
	return "WorkerInstance"
}

func (m *WorkerInstance) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkerInstance) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkerInstanceGroup struct {
	ID   string `gorm:"column:id;primaryKey" json:"id"`
	Type string `gorm:"column:type;not null" json:"type"`
	// For example "us-east-1"
	Name string `gorm:"column:name;not null" json:"name"`
	// If managed, it will default to the name, e.g. "us-east-1"
	// If unmanged, it will be prefixed with the project ID e.g. "project_1-us-east-1"
	MasterQueue string `gorm:"column:masterQueue;uniqueIndex;not null" json:"masterQueue"`
	// "N. Virginia, USA", "Frankfurt, Germany", etc. Used for display purposes
	Description    *string `gorm:"column:description" json:"description"`
	Hidden         bool    `gorm:"column:hidden;default:false;not null" json:"hidden"`
	TokenID        string  `gorm:"column:tokenId;uniqueIndex;not null" json:"tokenId"`
	OrganizationID *string `gorm:"column:organizationId" json:"organizationId"`
	ProjectID      *string `gorm:"column:projectId" json:"projectId"`
	CloudProvider  *string `gorm:"column:cloudProvider" json:"cloudProvider"`
	// "usa", "europe", etc. Used like a pseudo enum for things like flags
	Location     *string `gorm:"column:location" json:"location"`
	StaticIPs    *string `gorm:"column:staticIPs" json:"staticIPs"`
	WorkloadType string  `gorm:"column:workloadType;default:'CONTAINER';not null" json:"workloadType"`
	// Geo region; container + MICROVM groups for one geo share it. Set-once once it has runs.
	Region *string `gorm:"column:region" json:"region"`
	// When true, runs enqueued to this worker queue may skip the intermediate queue
	// and be pushed directly to the worker queue when concurrency is available.
	EnableFastPath     bool              `gorm:"column:enableFastPath;default:false;not null" json:"enableFastPath"`
	CreatedAt          DateTime          `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime          `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Token              *WorkerGroupToken `gorm:"foreignKey:TokenID"`
	Organization       *Organization
	Project            *Project
	Workers            []WorkerInstance   `gorm:"foreignKey:WorkerGroupID"`
	BackgroundWorkers  []BackgroundWorker `gorm:"foreignKey:WorkerGroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	DefaultForProjects []Project          `gorm:"foreignKey:DefaultWorkerGroupID"`
}

func (WorkerInstanceGroup) TableName() string {
	return "WorkerInstanceGroup"
}

func (m *WorkerInstanceGroup) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkerInstanceGroup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkerGroupToken struct {
	ID                      string               `gorm:"column:id;primaryKey" json:"id"`
	TokenHash               string               `gorm:"column:tokenHash;uniqueIndex;not null" json:"tokenHash"`
	PreviousTokenHash       *string              `gorm:"column:previousTokenHash;uniqueIndex" json:"previousTokenHash"`
	PreviousTokenValidUntil *DateTime            `gorm:"column:previousTokenValidUntil" json:"previousTokenValidUntil"`
	CreatedAt               DateTime             `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime             `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	WorkerGroup             *WorkerInstanceGroup `gorm:"foreignKey:TokenID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (WorkerGroupToken) TableName() string {
	return "WorkerGroupToken"
}

func (m *WorkerGroupToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkerGroupToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TaskRunTag struct {
	ID         string   `gorm:"column:id;primaryKey;index:TaskRunTag_name_id_idx" json:"id"`
	Name       string   `gorm:"column:name;uniqueIndex:TaskRunTag_projectId_name_key;index:TaskRunTag_name_id_idx;not null" json:"name"`
	FriendlyID string   `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	ProjectID  string   `gorm:"column:projectId;uniqueIndex:TaskRunTag_projectId_name_key;not null" json:"projectId"`
	CreatedAt  DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Project    *Project
}

func (TaskRunTag) TableName() string {
	return "TaskRunTag"
}

func (m *TaskRunTag) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// This is used for triggerAndWait and batchTriggerAndWait. The taskRun is the child task, it points at a parent attempt or a batch
type TaskRunDependency struct {
	ID                  string    `gorm:"column:id;primaryKey" json:"id"`
	TaskRunID           string    `gorm:"column:taskRunId;uniqueIndex;not null" json:"taskRunId"`
	CheckpointEventID   *string   `gorm:"column:checkpointEventId;uniqueIndex" json:"checkpointEventId"`
	DependentAttemptID  *string   `gorm:"column:dependentAttemptId;index:TaskRunDependency_dependentAttemptId_idx" json:"dependentAttemptId"`
	DependentBatchRunID *string   `gorm:"column:dependentBatchRunId;index:TaskRunDependency_dependentBatchRunId_idx" json:"dependentBatchRunId"`
	CreatedAt           DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	ResumedAt           *DateTime `gorm:"column:resumedAt" json:"resumedAt"`
	// The child run
	TaskRun         *TaskRun
	CheckpointEvent *CheckpointRestoreEvent `gorm:"foreignKey:CheckpointEventID"`
	// An attempt that is dependent on this task run.
	DependentAttempt *TaskRunAttempt `gorm:"foreignKey:DependentAttemptID"`
	// A batch run that is dependent on this task run
	DependentBatchRun *BatchTaskRun `gorm:"foreignKey:DependentBatchRunID"`
}

func (TaskRunDependency) TableName() string {
	return "TaskRunDependency"
}

func (m *TaskRunDependency) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskRunDependency) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// deprecated, we hadn't included the project id in the unique constraint
type TaskRunCounter struct {
	TaskIdentifier string `gorm:"column:taskIdentifier;primaryKey" json:"taskIdentifier"`
	LastNumber     int    `gorm:"column:lastNumber;default:0;not null" json:"lastNumber"`
}

func (TaskRunCounter) TableName() string {
	return "TaskRunCounter"
}

// Used for the TaskRun number (e.g. #1,421)
type TaskRunNumberCounter struct {
	ID             string              `gorm:"column:id;primaryKey" json:"id"`
	TaskIdentifier string              `gorm:"column:taskIdentifier;uniqueIndex:TaskRunNumberCounter_taskIdentifier_environmentId_key;not null" json:"taskIdentifier"`
	EnvironmentID  string              `gorm:"column:environmentId;uniqueIndex:TaskRunNumberCounter_taskIdentifier_environmentId_key;not null" json:"environmentId"`
	LastNumber     int                 `gorm:"column:lastNumber;default:0;not null" json:"lastNumber"`
	Environment    *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
}

func (TaskRunNumberCounter) TableName() string {
	return "TaskRunNumberCounter"
}

func (m *TaskRunNumberCounter) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// This is not used from engine v2+, attempts use the TaskRunExecutionSnapshot and TaskRun
type TaskRunAttempt struct {
	ID                     string         `gorm:"column:id;primaryKey" json:"id"`
	Number                 int            `gorm:"column:number;uniqueIndex:TaskRunAttempt_taskRunId_number_key;default:0;not null" json:"number"`
	FriendlyID             string         `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	TaskRunID              string         `gorm:"column:taskRunId;uniqueIndex:TaskRunAttempt_taskRunId_number_key;index:TaskRunAttempt_taskRunId_idx;not null" json:"taskRunId"`
	BackgroundWorkerID     string         `gorm:"column:backgroundWorkerId;not null" json:"backgroundWorkerId"`
	BackgroundWorkerTaskID string         `gorm:"column:backgroundWorkerTaskId;not null" json:"backgroundWorkerTaskId"`
	RuntimeEnvironmentID   string         `gorm:"column:runtimeEnvironmentId;not null" json:"runtimeEnvironmentId"`
	QueueID                string         `gorm:"column:queueId;not null" json:"queueId"`
	Status                 string         `gorm:"column:status;default:'PENDING';not null" json:"status"`
	CreatedAt              DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt              DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	StartedAt              *DateTime      `gorm:"column:startedAt" json:"startedAt"`
	CompletedAt            *DateTime      `gorm:"column:completedAt" json:"completedAt"`
	UsageDurationMs        int            `gorm:"column:usageDurationMs;default:0;not null" json:"usageDurationMs"`
	Error                  datatypes.JSON `gorm:"column:error" json:"error"`
	Output                 *string        `gorm:"column:output" json:"output"`
	OutputType             *string        `gorm:"column:outputType;default:'application/json';not null" json:"outputType"`
	TaskRun                *TaskRun
	BackgroundWorker       *BackgroundWorker
	BackgroundWorkerTask   *BackgroundWorkerTask
	RuntimeEnvironment     *RuntimeEnvironment
	Queue                  *TaskQueue               `gorm:"foreignKey:QueueID"`
	Dependencies           []TaskRunDependency      `gorm:"foreignKey:DependentAttemptID"`
	BatchDependencies      []BatchTaskRun           `gorm:"foreignKey:DependentTaskAttemptID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Checkpoints            []Checkpoint             `gorm:"foreignKey:AttemptID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BatchTaskRunItems      []BatchTaskRunItem       `gorm:"foreignKey:TaskRunAttemptID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CheckpointRestoreEvent []CheckpointRestoreEvent `gorm:"foreignKey:AttemptID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChildRuns              []TaskRun                `gorm:"foreignKey:ParentTaskRunAttemptID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
}

func (TaskRunAttempt) TableName() string {
	return "TaskRunAttempt"
}

func (m *TaskRunAttempt) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskRunAttempt) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// This is the unified otel span/log model that will eventually be replaced by clickhouse
type TaskEvent struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// This matches the span name for a trace event, or the log body for a log event
	Message     string  `gorm:"column:message;not null" json:"message"`
	TraceID     string  `gorm:"column:traceId;index:TaskEvent_traceId_idx;not null" json:"traceId"`
	SpanID      string  `gorm:"column:spanId;index:TaskEvent_spanId_idx;not null" json:"spanId"`
	ParentID    *string `gorm:"column:parentId" json:"parentId"`
	Tracestate  *string `gorm:"column:tracestate" json:"tracestate"`
	IsError     bool    `gorm:"column:isError;default:false;not null" json:"isError"`
	IsPartial   bool    `gorm:"column:isPartial;default:false;not null" json:"isPartial"`
	IsCancelled bool    `gorm:"column:isCancelled;default:false;not null" json:"isCancelled"`
	// deprecated: don't use this, moving this to properties, this now uses TaskEventKind.LOG
	IsDebug          bool           `gorm:"column:isDebug;default:false;not null" json:"isDebug"`
	ServiceName      string         `gorm:"column:serviceName;not null" json:"serviceName"`
	ServiceNamespace string         `gorm:"column:serviceNamespace;not null" json:"serviceNamespace"`
	Level            string         `gorm:"column:level;default:'TRACE';not null" json:"level"`
	Kind             string         `gorm:"column:kind;default:'INTERNAL';not null" json:"kind"`
	Status           string         `gorm:"column:status;default:'UNSET';not null" json:"status"`
	Links            datatypes.JSON `gorm:"column:links" json:"links"`
	Events           datatypes.JSON `gorm:"column:events" json:"events"`
	// This is the time the event started in nanoseconds since the epoch
	StartTime int64 `gorm:"column:startTime;not null" json:"startTime"`
	// This is the duration of the event in nanoseconds
	Duration        int64   `gorm:"column:duration;default:0;not null" json:"duration"`
	AttemptID       *string `gorm:"column:attemptId" json:"attemptId"`
	AttemptNumber   *int    `gorm:"column:attemptNumber" json:"attemptNumber"`
	EnvironmentID   string  `gorm:"column:environmentId;not null" json:"environmentId"`
	EnvironmentType string  `gorm:"column:environmentType;not null" json:"environmentType"`
	OrganizationID  string  `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID       string  `gorm:"column:projectId;not null" json:"projectId"`
	ProjectRef      string  `gorm:"column:projectRef;not null" json:"projectRef"`
	RunID           string  `gorm:"column:runId;index:TaskEvent_runId_idx;not null" json:"runId"`
	RunIsTest       bool    `gorm:"column:runIsTest;default:false;not null" json:"runIsTest"`
	IdempotencyKey  *string `gorm:"column:idempotencyKey" json:"idempotencyKey"`
	TaskSlug        string  `gorm:"column:taskSlug;not null" json:"taskSlug"`
	TaskPath        *string `gorm:"column:taskPath" json:"taskPath"`
	TaskExportName  *string `gorm:"column:taskExportName" json:"taskExportName"`
	WorkerID        *string `gorm:"column:workerId" json:"workerId"`
	WorkerVersion   *string `gorm:"column:workerVersion" json:"workerVersion"`
	QueueID         *string `gorm:"column:queueId" json:"queueId"`
	QueueName       *string `gorm:"column:queueName" json:"queueName"`
	BatchID         *string `gorm:"column:batchId" json:"batchId"`
	// This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering
	Properties datatypes.JSON `gorm:"column:properties;not null" json:"properties"`
	// This represents all span attributes in the $metadata namespace, like $metadata.payload
	Metadata datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	// This represents all span attributes in the $style namespace, like $style
	Style datatypes.JSON `gorm:"column:style" json:"style"`
	// This represents all span attributes in the $output namespace, like $output
	Output datatypes.JSON `gorm:"column:output" json:"output"`
	// This represents the mimetype of the output, such as application/json or application/super+json
	OutputType              *string        `gorm:"column:outputType" json:"outputType"`
	Payload                 datatypes.JSON `gorm:"column:payload" json:"payload"`
	PayloadType             *string        `gorm:"column:payloadType" json:"payloadType"`
	CreatedAt               DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UsageDurationMs         int            `gorm:"column:usageDurationMs;default:0;not null" json:"usageDurationMs"`
	UsageCostInCents        float64        `gorm:"column:usageCostInCents;default:0;not null" json:"usageCostInCents"`
	MachinePreset           *string        `gorm:"column:machinePreset" json:"machinePreset"`
	MachinePresetCPU        *float64       `gorm:"column:machinePresetCpu" json:"machinePresetCpu"`
	MachinePresetMemory     *float64       `gorm:"column:machinePresetMemory" json:"machinePresetMemory"`
	MachinePresetCentsPerMs *float64       `gorm:"column:machinePresetCentsPerMs" json:"machinePresetCentsPerMs"`
}

func (TaskEvent) TableName() string {
	return "TaskEvent"
}

func (m *TaskEvent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type TaskQueue struct {
	ID            string  `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID    string  `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Name          string  `gorm:"column:name;uniqueIndex:TaskQueue_runtimeEnvironmentId_name_key;not null" json:"name"`
	Type          string  `gorm:"column:type;default:'VIRTUAL';not null" json:"type"`
	Version       string  `gorm:"column:version;default:'V1';not null" json:"version"`
	OrderableName *string `gorm:"column:orderableName" json:"orderableName"`
	// How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
	// (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
	ConcurrencyVersion string `gorm:"column:concurrencyVersion;default:'V1';not null" json:"concurrencyVersion"`
	// QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
	// excluded from every queue-facing surface.
	Role                 string `gorm:"column:role;default:'QUEUE';not null" json:"role"`
	ProjectID            string `gorm:"column:projectId;not null" json:"projectId"`
	RuntimeEnvironmentID string `gorm:"column:runtimeEnvironmentId;uniqueIndex:TaskQueue_runtimeEnvironmentId_name_key;not null" json:"runtimeEnvironmentId"`
	// Represents the current concurrency limit for the queue
	ConcurrencyLimit *int `gorm:"column:concurrencyLimit" json:"concurrencyLimit"`
	// When the concurrency limit was overridden
	ConcurrencyLimitOverriddenAt *DateTime `gorm:"column:concurrencyLimitOverriddenAt" json:"concurrencyLimitOverriddenAt"`
	// Who overrode the concurrency limit (will be null if overridden via the API)
	ConcurrencyLimitOverriddenBy *string `gorm:"column:concurrencyLimitOverriddenBy" json:"concurrencyLimitOverriddenBy"`
	// If concurrencyLimit is overridden, this is the overridden value
	ConcurrencyLimitBase *int `gorm:"column:concurrencyLimitBase" json:"concurrencyLimitBase"`
	// If the override was expressed as a percentage of the environment limit, this stores that
	// percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
	// Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
	ConcurrencyLimitOverridePercent *float64 `gorm:"column:concurrencyLimitOverridePercent;type:decimal(5,2)" json:"concurrencyLimitOverridePercent"`
	// Caps total concurrent runs across ALL concurrencyKey values of this queue
	// (concurrencyLimit applies per key value). Null = no total cap.
	TotalConcurrencyLimit *int `gorm:"column:totalConcurrencyLimit" json:"totalConcurrencyLimit"`
	// When the total concurrency limit was overridden
	TotalConcurrencyLimitOverriddenAt *DateTime `gorm:"column:totalConcurrencyLimitOverriddenAt" json:"totalConcurrencyLimitOverriddenAt"`
	// Who overrode the total concurrency limit (null when overridden via the API)
	TotalConcurrencyLimitOverriddenBy *string `gorm:"column:totalConcurrencyLimitOverriddenBy" json:"totalConcurrencyLimitOverriddenBy"`
	// If totalConcurrencyLimit is overridden, the declared value it reverts to on reset
	TotalConcurrencyLimitBase *int           `gorm:"column:totalConcurrencyLimitBase" json:"totalConcurrencyLimitBase"`
	RateLimit                 datatypes.JSON `gorm:"column:rateLimit" json:"rateLimit"`
	Paused                    bool           `gorm:"column:paused;default:false;not null" json:"paused"`
	// Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths.
	ArchivedAt         *DateTime `gorm:"column:archivedAt" json:"archivedAt"`
	CreatedAt          DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project            *Project
	RuntimeEnvironment *RuntimeEnvironment
	Attempts           []TaskRunAttempt       `gorm:"foreignKey:QueueID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tasks              []BackgroundWorkerTask `gorm:"foreignKey:QueueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Workers            []BackgroundWorker     `gorm:"many2many:_BackgroundWorkerToTaskQueue;joinForeignKey:B;joinReferences:A"`
}

func (TaskQueue) TableName() string {
	return "TaskQueue"
}

func (m *TaskQueue) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskQueue) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BatchTaskRun struct {
	ID                      string    `gorm:"column:id;primaryKey;index:BatchTaskRun_runtimeEnvironmentId_id_idx;index:BatchTaskRun_runtimeEnvironmentId_createdAt_id_idx" json:"id"`
	FriendlyID              string    `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	IdempotencyKey          *string   `gorm:"column:idempotencyKey;uniqueIndex:BatchTaskRun_runtimeEnvironmentId_idempotencyKey_key" json:"idempotencyKey"`
	IdempotencyKeyExpiresAt *DateTime `gorm:"column:idempotencyKeyExpiresAt" json:"idempotencyKeyExpiresAt"`
	Status                  string    `gorm:"column:status;default:'PENDING';not null" json:"status"`
	RuntimeEnvironmentID    string    `gorm:"column:runtimeEnvironmentId;uniqueIndex:BatchTaskRun_runtimeEnvironmentId_idempotencyKey_key;index:BatchTaskRun_runtimeEnvironmentId_id_idx;index:BatchTaskRun_runtimeEnvironmentId_createdAt_id_idx;not null" json:"runtimeEnvironmentId"`
	CreatedAt               DateTime  `gorm:"column:createdAt;index:BatchTaskRun_runtimeEnvironmentId_createdAt_id_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// Friendly IDs
	RunIds       []string       `gorm:"column:runIds;serializer:json;not null" json:"runIds"`
	RunCount     int            `gorm:"column:runCount;default:0;not null" json:"runCount"`
	Payload      *string        `gorm:"column:payload" json:"payload"`
	PayloadType  *string        `gorm:"column:payloadType;default:'application/json';not null" json:"payloadType"`
	Options      datatypes.JSON `gorm:"column:options" json:"options"`
	BatchVersion *string        `gorm:"column:batchVersion;default:'v1';not null" json:"batchVersion"`
	// sealed is set to true once no more items can be added to the batch
	Sealed   bool      `gorm:"column:sealed;default:false;not null" json:"sealed"`
	SealedAt *DateTime `gorm:"column:sealedAt" json:"sealedAt"`
	// this is the expected number of items in the batch
	ExpectedCount int `gorm:"column:expectedCount;default:0;not null" json:"expectedCount"`
	// this is the completed number of items in the batch. once this reaches expectedCount, and the batch is sealed, the batch is considered completed
	CompletedCount int       `gorm:"column:completedCount;default:0;not null" json:"completedCount"`
	CompletedAt    *DateTime `gorm:"column:completedAt" json:"completedAt"`
	ResumedAt      *DateTime `gorm:"column:resumedAt" json:"resumedAt"`
	// this is used to be able to "seal" this BatchTaskRun when all of the runs have been triggered asynchronously, and using the "parallel" processing strategy
	ProcessingJobsCount         int `gorm:"column:processingJobsCount;default:0;not null" json:"processingJobsCount"`
	ProcessingJobsExpectedCount int `gorm:"column:processingJobsExpectedCount;default:0;not null" json:"processingJobsExpectedCount"`
	// optional token that can be used to authenticate the task run
	OneTimeUseToken *string `gorm:"column:oneTimeUseToken;uniqueIndex" json:"oneTimeUseToken"`
	// When processing started (status changed to PROCESSING)
	ProcessingStartedAt *DateTime `gorm:"column:processingStartedAt" json:"processingStartedAt"`
	// When processing completed (all items processed)
	ProcessingCompletedAt *DateTime `gorm:"column:processingCompletedAt" json:"processingCompletedAt"`
	// Count of successfully created runs
	SuccessfulRunCount *int `gorm:"column:successfulRunCount" json:"successfulRunCount"`
	// Count of failed run creations
	FailedRunCount         *int                    `gorm:"column:failedRunCount" json:"failedRunCount"`
	TaskIdentifier         *string                 `gorm:"column:taskIdentifier" json:"taskIdentifier"`
	CheckpointEventID      *string                 `gorm:"column:checkpointEventId;uniqueIndex" json:"checkpointEventId"`
	DependentTaskAttemptID *string                 `gorm:"column:dependentTaskAttemptId;index:BatchTaskRun_dependentTaskAttemptId_idx" json:"dependentTaskAttemptId"`
	CheckpointEvent        *CheckpointRestoreEvent `gorm:"foreignKey:CheckpointEventID"`
	DependentTaskAttempt   *TaskRunAttempt         `gorm:"foreignKey:DependentTaskAttemptID"`
	// This only includes new runs, not idempotent runs.
	Runs []TaskRun `gorm:"foreignKey:BatchID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
	// Specific run blockers,
	RunsBlocked []TaskRunWaitpoint `gorm:"foreignKey:BatchID"`
	// Waitpoints that are blocked by this batch.
	// When a Batch is created it blocks execution of the associated parent run (for andWait)
	Waitpoints []Waitpoint `gorm:"foreignKey:CompletedByBatchID;constraint:OnDelete:SET NULL"`
	// Detailed failure records
	Errors []BatchTaskRunError `gorm:"foreignKey:BatchTaskRunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// all the below properties are engine v1 only
	Items           []BatchTaskRunItem  `gorm:"foreignKey:BatchTaskRunID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RunDependencies []TaskRunDependency `gorm:"foreignKey:DependentBatchRunID"`
}

func (BatchTaskRun) TableName() string {
	return "BatchTaskRun"
}

func (m *BatchTaskRun) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BatchTaskRun) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Used in engine V1 only
type BatchTaskRunItem struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	Status           string    `gorm:"column:status;default:'PENDING';not null" json:"status"`
	BatchTaskRunID   string    `gorm:"column:batchTaskRunId;uniqueIndex:BatchTaskRunItem_batchTaskRunId_taskRunId_key;not null" json:"batchTaskRunId"`
	TaskRunID        string    `gorm:"column:taskRunId;uniqueIndex:BatchTaskRunItem_batchTaskRunId_taskRunId_key;index:idx_batchtaskrunitem_taskrun;not null" json:"taskRunId"`
	TaskRunAttemptID *string   `gorm:"column:taskRunAttemptId;index:idx_batchtaskrunitem_taskrunattempt" json:"taskRunAttemptId"`
	CreatedAt        DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	CompletedAt      *DateTime `gorm:"column:completedAt" json:"completedAt"`
	BatchTaskRun     *BatchTaskRun
	TaskRun          *TaskRun
	TaskRunAttempt   *TaskRunAttempt
}

func (BatchTaskRunItem) TableName() string {
	return "BatchTaskRunItem"
}

func (m *BatchTaskRunItem) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BatchTaskRunItem) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Track individual run creation failures in batch processing (Run Engine v2)
type BatchTaskRunError struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	BatchTaskRunID string `gorm:"column:batchTaskRunId;uniqueIndex:BatchTaskRunError_batchTaskRunId_index_key;index:BatchTaskRunError_batchTaskRunId_idx;not null" json:"batchTaskRunId"`
	// Which item in the batch (0-based index)
	Index int `gorm:"column:index;uniqueIndex:BatchTaskRunError_batchTaskRunId_index_key;not null" json:"index"`
	// The task identifier that was being triggered
	TaskIdentifier string `gorm:"column:taskIdentifier;not null" json:"taskIdentifier"`
	// The payload that failed (JSON, may be truncated)
	Payload *string `gorm:"column:payload" json:"payload"`
	// The options that were used
	Options datatypes.JSON `gorm:"column:options" json:"options"`
	// Error message
	Error string `gorm:"column:error;not null" json:"error"`
	// Error code if available
	ErrorCode    *string  `gorm:"column:errorCode" json:"errorCode"`
	CreatedAt    DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	BatchTaskRun *BatchTaskRun
}

func (BatchTaskRunError) TableName() string {
	return "BatchTaskRunError"
}

func (m *BatchTaskRunError) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type EnvironmentVariable struct {
	ID         string   `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID string   `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Key        string   `gorm:"column:key;uniqueIndex:EnvironmentVariable_projectId_key_key;not null" json:"key"`
	ProjectID  string   `gorm:"column:projectId;uniqueIndex:EnvironmentVariable_projectId_key_key;not null" json:"projectId"`
	CreatedAt  DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project    *Project
	Values     []EnvironmentVariableValue `gorm:"foreignKey:VariableID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (EnvironmentVariable) TableName() string {
	return "EnvironmentVariable"
}

func (m *EnvironmentVariable) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EnvironmentVariable) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EnvironmentVariableValue struct {
	ID               string  `gorm:"column:id;primaryKey" json:"id"`
	ValueReferenceID *string `gorm:"column:valueReferenceId;index:EnvironmentVariableValue_valueReferenceId_idx" json:"valueReferenceId"`
	VariableID       string  `gorm:"column:variableId;uniqueIndex:EnvironmentVariableValue_variableId_environmentId_key;index:EnvironmentVariableValue_environmentId_updatedAt_variableId_idx;not null" json:"variableId"`
	EnvironmentID    string  `gorm:"column:environmentId;uniqueIndex:EnvironmentVariableValue_variableId_environmentId_key;index:EnvironmentVariableValue_environmentId_idx;index:EnvironmentVariableValue_environmentId_updatedAt_variableId_idx;not null" json:"environmentId"`
	// If true, the value is secret and cannot be revealed
	IsSecret       bool                 `gorm:"column:isSecret;default:false;not null" json:"isSecret"`
	CreatedAt      DateTime             `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime             `gorm:"column:updatedAt;index:EnvironmentVariableValue_environmentId_updatedAt_variableId_idx;autoUpdateTime:false;not null" json:"updatedAt"`
	Version        *int                 `gorm:"column:version;default:1;not null" json:"version"`
	LastUpdatedBy  datatypes.JSON       `gorm:"column:lastUpdatedBy" json:"lastUpdatedBy"`
	ValueReference *SecretReference     `gorm:"foreignKey:ValueReferenceID"`
	Variable       *EnvironmentVariable `gorm:"foreignKey:VariableID"`
	Environment    *RuntimeEnvironment  `gorm:"foreignKey:EnvironmentID"`
}

func (EnvironmentVariableValue) TableName() string {
	return "EnvironmentVariableValue"
}

func (m *EnvironmentVariableValue) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EnvironmentVariableValue) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Checkpoint struct {
	ID                   string          `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID           string          `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Type                 string          `gorm:"column:type;not null" json:"type"`
	Location             string          `gorm:"column:location;not null" json:"location"`
	ImageRef             string          `gorm:"column:imageRef;not null" json:"imageRef"`
	Reason               *string         `gorm:"column:reason" json:"reason"`
	Metadata             *string         `gorm:"column:metadata" json:"metadata"`
	RunID                string          `gorm:"column:runId;index:Checkpoint_runId_idx;not null" json:"runId"`
	AttemptID            string          `gorm:"column:attemptId;index:Checkpoint_attemptId_idx;not null" json:"attemptId"`
	AttemptNumber        *int            `gorm:"column:attemptNumber" json:"attemptNumber"`
	ProjectID            string          `gorm:"column:projectId;not null" json:"projectId"`
	RuntimeEnvironmentID string          `gorm:"column:runtimeEnvironmentId;not null" json:"runtimeEnvironmentId"`
	CreatedAt            DateTime        `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime        `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Run                  *TaskRun        `gorm:"foreignKey:RunID"`
	Attempt              *TaskRunAttempt `gorm:"foreignKey:AttemptID"`
	Project              *Project
	RuntimeEnvironment   *RuntimeEnvironment
	Events               []CheckpointRestoreEvent `gorm:"foreignKey:CheckpointID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Checkpoint) TableName() string {
	return "Checkpoint"
}

func (m *Checkpoint) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Checkpoint) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CheckpointRestoreEvent struct {
	ID                     string   `gorm:"column:id;primaryKey" json:"id"`
	Type                   string   `gorm:"column:type;not null" json:"type"`
	Reason                 *string  `gorm:"column:reason" json:"reason"`
	Metadata               *string  `gorm:"column:metadata" json:"metadata"`
	CheckpointID           string   `gorm:"column:checkpointId;index:CheckpointRestoreEvent_checkpointId_idx;not null" json:"checkpointId"`
	RunID                  string   `gorm:"column:runId;index:CheckpointRestoreEvent_runId_idx;not null" json:"runId"`
	AttemptID              string   `gorm:"column:attemptId;not null" json:"attemptId"`
	ProjectID              string   `gorm:"column:projectId;not null" json:"projectId"`
	RuntimeEnvironmentID   string   `gorm:"column:runtimeEnvironmentId;not null" json:"runtimeEnvironmentId"`
	CreatedAt              DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt              DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Checkpoint             *Checkpoint
	Run                    *TaskRun        `gorm:"foreignKey:RunID"`
	Attempt                *TaskRunAttempt `gorm:"foreignKey:AttemptID"`
	Project                *Project
	RuntimeEnvironment     *RuntimeEnvironment
	TaskRunDependency      *TaskRunDependency `gorm:"foreignKey:CheckpointEventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BatchTaskRunDependency *BatchTaskRun      `gorm:"foreignKey:CheckpointEventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (CheckpointRestoreEvent) TableName() string {
	return "CheckpointRestoreEvent"
}

func (m *CheckpointRestoreEvent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CheckpointRestoreEvent) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkerDeployment struct {
	ID             string  `gorm:"column:id;primaryKey;index:WorkerDeployment_environmentId_status_id_idx" json:"id"`
	ContentHash    string  `gorm:"column:contentHash;not null" json:"contentHash"`
	FriendlyID     string  `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	ShortCode      string  `gorm:"column:shortCode;uniqueIndex:WorkerDeployment_projectId_shortCode_key;not null" json:"shortCode"`
	Version        string  `gorm:"column:version;uniqueIndex:WorkerDeployment_environmentId_version_key;not null" json:"version"`
	Runtime        *string `gorm:"column:runtime" json:"runtime"`
	RuntimeVersion *string `gorm:"column:runtimeVersion" json:"runtimeVersion"`
	// CLI version that initiated the deploy, stamped at initialization
	CliVersion     *string `gorm:"column:cliVersion" json:"cliVersion"`
	ImageReference *string `gorm:"column:imageReference" json:"imageReference"`
	ImagePlatform  *string `gorm:"column:imagePlatform;default:'linux/amd64';not null" json:"imagePlatform"`
	// Key of the microVM template prebuilt from this deployment's image, set when the builder succeeds
	TemplateKey *string `gorm:"column:templateKey" json:"templateKey"`
	// Object key of the full index snapshot in the deployment index store. Set only after the upload succeeds.
	IndexObjectKey *string `gorm:"column:indexObjectKey" json:"indexObjectKey"`
	// sha256 hex of the snapshot's uncompressed JSON
	IndexSha256         *string        `gorm:"column:indexSha256" json:"indexSha256"`
	IndexFormatVersion  *int           `gorm:"column:indexFormatVersion" json:"indexFormatVersion"`
	ExternalBuildData   datatypes.JSON `gorm:"column:externalBuildData" json:"externalBuildData"`
	BuildServerMetadata datatypes.JSON `gorm:"column:buildServerMetadata" json:"buildServerMetadata"`
	// Encrypted build-time env vars for pre-bundled (fromBundle) deploys, as an
	// EncryptedSecretValue envelope. Kept for the redeploy window, then cleared by a sweep.
	BuildEnvVars           datatypes.JSON `gorm:"column:buildEnvVars" json:"buildEnvVars"`
	Status                 string         `gorm:"column:status;index:WorkerDeployment_environmentId_status_id_idx;default:'PENDING';not null" json:"status"`
	Type                   string         `gorm:"column:type;default:'V1';not null" json:"type"`
	ProjectID              string         `gorm:"column:projectId;uniqueIndex:WorkerDeployment_projectId_shortCode_key;not null" json:"projectId"`
	EnvironmentID          string         `gorm:"column:environmentId;uniqueIndex:WorkerDeployment_environmentId_version_key;index:WorkerDeployment_environmentId_createdAt_idx;index:WorkerDeployment_environmentId_status_id_idx;index:WorkerDeployment_environmentId_externalId_idx;not null" json:"environmentId"`
	WorkerID               *string        `gorm:"column:workerId;uniqueIndex" json:"workerId"`
	TriggeredByID          *string        `gorm:"column:triggeredById" json:"triggeredById"`
	TriggeredVia           *string        `gorm:"column:triggeredVia" json:"triggeredVia"`
	CommitSHA              *string        `gorm:"column:commitSHA;index:WorkerDeployment_commitSHA_idx" json:"commitSHA"`
	ExternalID             *string        `gorm:"column:externalId;index:WorkerDeployment_environmentId_externalId_idx" json:"externalId"`
	StartedAt              *DateTime      `gorm:"column:startedAt" json:"startedAt"`
	InstalledAt            *DateTime      `gorm:"column:installedAt" json:"installedAt"`
	BuiltAt                *DateTime      `gorm:"column:builtAt" json:"builtAt"`
	DeployedAt             *DateTime      `gorm:"column:deployedAt" json:"deployedAt"`
	FailedAt               *DateTime      `gorm:"column:failedAt" json:"failedAt"`
	ErrorData              datatypes.JSON `gorm:"column:errorData" json:"errorData"`
	CanceledAt             *DateTime      `gorm:"column:canceledAt" json:"canceledAt"`
	CanceledReason         *string        `gorm:"column:canceledReason" json:"canceledReason"`
	Git                    datatypes.JSON `gorm:"column:git" json:"git"`
	CreatedAt              DateTime       `gorm:"column:createdAt;index:WorkerDeployment_environmentId_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt              DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project                *Project
	Environment            *RuntimeEnvironment         `gorm:"foreignKey:EnvironmentID"`
	Worker                 *BackgroundWorker           `gorm:"foreignKey:WorkerID"`
	TriggeredBy            *User                       `gorm:"foreignKey:TriggeredByID"`
	Promotions             []WorkerDeploymentPromotion `gorm:"foreignKey:DeploymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Alerts                 []ProjectAlert              `gorm:"foreignKey:WorkerDeploymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkerInstance         []WorkerInstance            `gorm:"foreignKey:DeploymentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	IntegrationDeployments []IntegrationDeployment     `gorm:"foreignKey:DeploymentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (WorkerDeployment) TableName() string {
	return "WorkerDeployment"
}

func (m *WorkerDeployment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkerDeployment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkerDeploymentPromotion struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// This is the promotion label, e.g. "current"
	Label         string              `gorm:"column:label;uniqueIndex:WorkerDeploymentPromotion_environmentId_label_key;not null" json:"label"`
	DeploymentID  string              `gorm:"column:deploymentId;not null" json:"deploymentId"`
	EnvironmentID string              `gorm:"column:environmentId;uniqueIndex:WorkerDeploymentPromotion_environmentId_label_key;not null" json:"environmentId"`
	Deployment    *WorkerDeployment   `gorm:"foreignKey:DeploymentID"`
	Environment   *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
}

func (WorkerDeploymentPromotion) TableName() string {
	return "WorkerDeploymentPromotion"
}

func (m *WorkerDeploymentPromotion) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Schedules can be attached to tasks to trigger them on a schedule
type TaskSchedule struct {
	ID   string `gorm:"column:id;primaryKey" json:"id"`
	Type string `gorm:"column:type;default:'IMPERATIVE';not null" json:"type"`
	// users see this as `id`. They start with schedule_
	FriendlyID string `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	// a reference to a task (not a foreign key because it's across versions)
	TaskIdentifier string `gorm:"column:taskIdentifier;not null" json:"taskIdentifier"`
	// can be provided and we won't create another with the same key
	DeduplicationKey             string `gorm:"column:deduplicationKey;uniqueIndex:TaskSchedule_projectId_deduplicationKey_key;not null" json:"deduplicationKey"`
	UserProvidedDeduplicationKey bool   `gorm:"column:userProvidedDeduplicationKey;default:false;not null" json:"userProvidedDeduplicationKey"`
	// the CRON pattern
	GeneratorExpression  string `gorm:"column:generatorExpression;not null" json:"generatorExpression"`
	GeneratorDescription string `gorm:"column:generatorDescription;default:'';not null" json:"generatorDescription"`
	GeneratorType        string `gorm:"column:generatorType;default:'CRON';not null" json:"generatorType"`
	// These are IANA format string, or the default "UTC". E.g. "America/New_York"
	Timezone                     *string `gorm:"column:timezone;default:'UTC';not null" json:"timezone"`
	WindowDurationSeconds        *int    `gorm:"column:windowDurationSeconds" json:"windowDurationSeconds"`
	WindowPercentage             *int    `gorm:"column:windowPercentage" json:"windowPercentage"`
	DefaultWindowDurationSeconds *int    `gorm:"column:defaultWindowDurationSeconds" json:"defaultWindowDurationSeconds"`
	MinimumWindowDurationSeconds *int    `gorm:"column:minimumWindowDurationSeconds" json:"minimumWindowDurationSeconds"`
	// Can be provided by the user then accessed inside a run
	ExternalID *string `gorm:"column:externalId" json:"externalId"`
	// @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up.
	LastRunTriggeredAt *DateTime `gorm:"column:lastRunTriggeredAt" json:"lastRunTriggeredAt"`
	ProjectID          string    `gorm:"column:projectId;uniqueIndex:TaskSchedule_projectId_deduplicationKey_key;index:TaskSchedule_projectId_idx;index:TaskSchedule_projectId_createdAt_idx;not null" json:"projectId"`
	CreatedAt          DateTime  `gorm:"column:createdAt;index:TaskSchedule_projectId_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Active             *bool     `gorm:"column:active;default:true;not null" json:"active"`
	Project            *Project
	// Instances of the schedule that are active
	Instances []TaskScheduleInstance `gorm:"foreignKey:TaskScheduleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TaskSchedule) TableName() string {
	return "TaskSchedule"
}

func (m *TaskSchedule) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.DeduplicationKey == "" {
		m.DeduplicationKey = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskSchedule) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// An instance links a schedule with an environment
type TaskScheduleInstance struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	TaskScheduleID string   `gorm:"column:taskScheduleId;uniqueIndex:TaskScheduleInstance_taskScheduleId_environmentId_key;not null" json:"taskScheduleId"`
	EnvironmentID  string   `gorm:"column:environmentId;uniqueIndex:TaskScheduleInstance_taskScheduleId_environmentId_key;index:TaskScheduleInstance_environmentId_idx;not null" json:"environmentId"`
	ProjectID      string   `gorm:"column:projectId;index:TaskScheduleInstance_projectId_active_idx;not null" json:"projectId"`
	SchedulePhase  *int     `gorm:"column:schedulePhase" json:"schedulePhase"`
	CreatedAt      DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Active         *bool    `gorm:"column:active;index:TaskScheduleInstance_projectId_active_idx;default:true;not null" json:"active"`
	// @deprecated stop writing 2026-04-30; engine derives from cron + exactScheduleTime. Drop in follow-up.
	LastScheduledTimestamp *DateTime `gorm:"column:lastScheduledTimestamp" json:"lastScheduledTimestamp"`
	// @deprecated stop writing 2026-04-30; engine derives from cron + now(). Drop in follow-up.
	NextScheduledTimestamp *DateTime `gorm:"column:nextScheduledTimestamp" json:"nextScheduledTimestamp"`
	TaskSchedule           *TaskSchedule
	Environment            *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
	Project                *Project
}

func (TaskScheduleInstance) TableName() string {
	return "TaskScheduleInstance"
}

func (m *TaskScheduleInstance) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TaskScheduleInstance) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RuntimeEnvironmentSession struct {
	ID                  string               `gorm:"column:id;primaryKey" json:"id"`
	IPAddress           string               `gorm:"column:ipAddress;not null" json:"ipAddress"`
	EnvironmentID       string               `gorm:"column:environmentId;not null" json:"environmentId"`
	CreatedAt           DateTime             `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime             `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	DisconnectedAt      *DateTime            `gorm:"column:disconnectedAt" json:"disconnectedAt"`
	Environment         *RuntimeEnvironment  `gorm:"foreignKey:EnvironmentID"`
	CurrentEnvironments []RuntimeEnvironment `gorm:"foreignKey:CurrentSessionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (RuntimeEnvironmentSession) TableName() string {
	return "RuntimeEnvironmentSession"
}

func (m *RuntimeEnvironmentSession) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*RuntimeEnvironmentSession) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ProjectAlertChannel struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID string `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	// can be provided and we won't create another with the same key
	DeduplicationKey             string                   `gorm:"column:deduplicationKey;uniqueIndex:ProjectAlertChannel_projectId_deduplicationKey_key;not null" json:"deduplicationKey"`
	UserProvidedDeduplicationKey bool                     `gorm:"column:userProvidedDeduplicationKey;default:false;not null" json:"userProvidedDeduplicationKey"`
	IntegrationID                *string                  `gorm:"column:integrationId" json:"integrationId"`
	Enabled                      *bool                    `gorm:"column:enabled;default:true;not null" json:"enabled"`
	Type                         string                   `gorm:"column:type;not null" json:"type"`
	Name                         string                   `gorm:"column:name;not null" json:"name"`
	Properties                   datatypes.JSON           `gorm:"column:properties;not null" json:"properties"`
	AlertTypes                   []string                 `gorm:"column:alertTypes;serializer:json;not null" json:"alertTypes"`
	EnvironmentTypes             []string                 `gorm:"column:environmentTypes;serializer:json;not null" json:"environmentTypes"`
	ErrorAlertConfig             datatypes.JSON           `gorm:"column:errorAlertConfig" json:"errorAlertConfig"`
	ProjectID                    string                   `gorm:"column:projectId;uniqueIndex:ProjectAlertChannel_projectId_deduplicationKey_key;not null" json:"projectId"`
	CreatedAt                    DateTime                 `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                    DateTime                 `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Integration                  *OrganizationIntegration `gorm:"foreignKey:IntegrationID"`
	Project                      *Project
	Alerts                       []ProjectAlert        `gorm:"foreignKey:ChannelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AlertStorages                []ProjectAlertStorage `gorm:"foreignKey:AlertChannelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ProjectAlertChannel) TableName() string {
	return "ProjectAlertChannel"
}

func (m *ProjectAlertChannel) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.DeduplicationKey == "" {
		m.DeduplicationKey = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ProjectAlertChannel) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ProjectAlert struct {
	ID                 string   `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID         string   `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	ProjectID          string   `gorm:"column:projectId;not null" json:"projectId"`
	EnvironmentID      string   `gorm:"column:environmentId;not null" json:"environmentId"`
	ChannelID          string   `gorm:"column:channelId;index:ProjectAlert_channelId_idx;not null" json:"channelId"`
	Status             string   `gorm:"column:status;default:'PENDING';not null" json:"status"`
	Type               string   `gorm:"column:type;not null" json:"type"`
	TaskRunAttemptID   *string  `gorm:"column:taskRunAttemptId" json:"taskRunAttemptId"`
	TaskRunID          *string  `gorm:"column:taskRunId" json:"taskRunId"`
	WorkerDeploymentID *string  `gorm:"column:workerDeploymentId" json:"workerDeploymentId"`
	CreatedAt          DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project            *Project
	Environment        *RuntimeEnvironment  `gorm:"foreignKey:EnvironmentID"`
	Channel            *ProjectAlertChannel `gorm:"foreignKey:ChannelID"`
	WorkerDeployment   *WorkerDeployment
}

func (ProjectAlert) TableName() string {
	return "ProjectAlert"
}

func (m *ProjectAlert) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ProjectAlert) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ProjectAlertStorage struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID      string         `gorm:"column:projectId;not null" json:"projectId"`
	AlertChannelID string         `gorm:"column:alertChannelId;index:ProjectAlertStorage_alertChannelId_alertType_storageId_idx;not null" json:"alertChannelId"`
	AlertType      string         `gorm:"column:alertType;index:ProjectAlertStorage_alertChannelId_alertType_storageId_idx;not null" json:"alertType"`
	StorageID      string         `gorm:"column:storageId;index:ProjectAlertStorage_alertChannelId_alertType_storageId_idx;not null" json:"storageId"`
	StorageData    datatypes.JSON `gorm:"column:storageData;not null" json:"storageData"`
	CreatedAt      DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project        *Project
	AlertChannel   *ProjectAlertChannel `gorm:"foreignKey:AlertChannelID"`
}

func (ProjectAlertStorage) TableName() string {
	return "ProjectAlertStorage"
}

func (m *ProjectAlertStorage) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ProjectAlertStorage) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganizationIntegration struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID string `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Service    string `gorm:"column:service;not null" json:"service"`
	// Identifier for external, integration's organization (e.g. Vercel's team)
	ExternalOrganizationID         *string          `gorm:"column:externalOrganizationId;index:OrganizationIntegration_externalOrganizationId_idx" json:"externalOrganizationId"`
	IntegrationData                datatypes.JSON   `gorm:"column:integrationData;not null" json:"integrationData"`
	TokenReferenceID               string           `gorm:"column:tokenReferenceId;not null" json:"tokenReferenceId"`
	OrganizationID                 string           `gorm:"column:organizationId;not null" json:"organizationId"`
	CreatedAt                      DateTime         `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                      DateTime         `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt                      *DateTime        `gorm:"column:deletedAt" json:"deletedAt"`
	TokenReference                 *SecretReference `gorm:"foreignKey:TokenReferenceID"`
	Organization                   *Organization
	AlertChannels                  []ProjectAlertChannel            `gorm:"foreignKey:IntegrationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	OrganizationProjectIntegration []OrganizationProjectIntegration `gorm:"foreignKey:OrganizationIntegrationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OrganizationIntegration) TableName() string {
	return "OrganizationIntegration"
}

func (m *OrganizationIntegration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganizationIntegration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganizationProjectIntegration struct {
	ID                        string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationIntegrationID string `gorm:"column:organizationIntegrationId;index:OrganizationProjectIntegration_projectId_organizationIntegrationId_idx;not null" json:"organizationIntegrationId"`
	ProjectID                 string `gorm:"column:projectId;index:OrganizationProjectIntegration_projectId_idx;index:OrganizationProjectIntegration_projectId_organizationIntegrationId_idx;not null" json:"projectId"`
	// Identifier for webhooks, for example Vercel's projectId
	ExternalEntityID string `gorm:"column:externalEntityId;index:OrganizationProjectIntegration_externalEntityId_idx;not null" json:"externalEntityId"`
	// Save useful data like config or external entity name
	IntegrationData datatypes.JSON `gorm:"column:integrationData;not null" json:"integrationData"`
	// UserId who installed the integration
	InstalledBy             *string   `gorm:"column:installedBy" json:"installedBy"`
	CreatedAt               DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt               *DateTime `gorm:"column:deletedAt" json:"deletedAt"`
	OrganizationIntegration *OrganizationIntegration
	Project                 *Project
}

func (OrganizationProjectIntegration) TableName() string {
	return "OrganizationProjectIntegration"
}

func (m *OrganizationProjectIntegration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganizationProjectIntegration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Bulk actions, like canceling and replaying runs
type BulkActionGroup struct {
	ID            string  `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID    string  `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	ProjectID     string  `gorm:"column:projectId;not null" json:"projectId"`
	EnvironmentID *string `gorm:"column:environmentId;index:BulkActionGroup_environmentId_createdAt_idx;index:BulkActionGroup_environmentId_type_dedupeKey_idx;index:BulkActionGroup_environmentId_status_type_idx" json:"environmentId"`
	Type          string  `gorm:"column:type;index:BulkActionGroup_environmentId_type_dedupeKey_idx;index:BulkActionGroup_environmentId_status_type_idx;not null" json:"type"`
	// When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed.
	Status string `gorm:"column:status;index:BulkActionGroup_environmentId_status_type_idx;default:'PENDING';not null" json:"status"`
	// The query that will be executed to get the runs to process (if null, we are passing run ids directly)
	QueryName *string `gorm:"column:queryName" json:"queryName"`
	// The params that will be passed to the query
	Params datatypes.JSON `gorm:"column:params" json:"params"`
	// Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions.
	DedupeKey *string `gorm:"column:dedupeKey;index:BulkActionGroup_environmentId_type_dedupeKey_idx" json:"dedupeKey"`
	// The cursor that will be passed to the query (null for the first page)
	Cursor datatypes.JSON `gorm:"column:cursor" json:"cursor"`
	// The number of runs that have been processed successfully
	SuccessCount int `gorm:"column:successCount;default:0;not null" json:"successCount"`
	// The number of runs that have been failed
	FailureCount int `gorm:"column:failureCount;default:0;not null" json:"failureCount"`
	// The total number of runs that will be processed
	TotalCount             int     `gorm:"column:totalCount;default:0;not null" json:"totalCount"`
	CompletionNotification string  `gorm:"column:completionNotification;default:'NONE';not null" json:"completionNotification"`
	UserID                 *string `gorm:"column:userId" json:"userId"`
	// The name of the bulk action
	Name *string `gorm:"column:name" json:"name"`
	// The time the bulk action was completed
	CompletedAt *DateTime `gorm:"column:completedAt" json:"completedAt"`
	CreatedAt   DateTime  `gorm:"column:createdAt;index:BulkActionGroup_environmentId_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project     *Project
	// If this is set then it's a V2 Bulk Action that supports queries
	Environment *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
	// The userId who did the bulk action
	User *User
}

func (BulkActionGroup) TableName() string {
	return "BulkActionGroup"
}

func (m *BulkActionGroup) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BulkActionGroup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BulkActionItem struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @deprecated not used in new BulkActions
	FriendlyID *string `gorm:"column:friendlyId" json:"friendlyId"`
	GroupID    string  `gorm:"column:groupId;not null" json:"groupId"`
	Type       string  `gorm:"column:type;not null" json:"type"`
	// When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed.
	Status           string   `gorm:"column:status;default:'PENDING';not null" json:"status"`
	SourceRunID      string   `gorm:"column:sourceRunId;not null" json:"sourceRunId"`
	DestinationRunID *string  `gorm:"column:destinationRunId" json:"destinationRunId"`
	Error            *string  `gorm:"column:error" json:"error"`
	CreatedAt        DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// The run that is the source of the action, e.g. when replaying this is the original run
	SourceRun *TaskRun `gorm:"foreignKey:SourceRunID"`
	// The run that's a result of the action, this will be set when the run has been created
	DestinationRun *TaskRun `gorm:"foreignKey:DestinationRunID"`
}

func (BulkActionItem) TableName() string {
	return "BulkActionItem"
}

func (m *BulkActionItem) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BulkActionItem) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RealtimeStreamChunk struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Key       string   `gorm:"column:key;not null" json:"key"`
	Value     string   `gorm:"column:value;not null" json:"value"`
	Sequence  int      `gorm:"column:sequence;not null" json:"sequence"`
	RunID     string   `gorm:"column:runId;index:RealtimeStreamChunk_runId_idx;not null" json:"runId"`
	CreatedAt DateTime `gorm:"column:createdAt;index:RealtimeStreamChunk_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
}

func (RealtimeStreamChunk) TableName() string {
	return "RealtimeStreamChunk"
}

func (m *RealtimeStreamChunk) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// This is the unified otel span/log model that will eventually be replaced by clickhouse
type TaskEventPartitioned struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// This matches the span name for a trace event, or the log body for a log event
	Message          string         `gorm:"column:message;not null" json:"message"`
	TraceID          string         `gorm:"column:traceId;index:TaskEventPartitioned_traceId_idx;not null" json:"traceId"`
	SpanID           string         `gorm:"column:spanId;index:TaskEventPartitioned_spanId_idx;not null" json:"spanId"`
	ParentID         *string        `gorm:"column:parentId" json:"parentId"`
	Tracestate       *string        `gorm:"column:tracestate" json:"tracestate"`
	IsError          bool           `gorm:"column:isError;default:false;not null" json:"isError"`
	IsPartial        bool           `gorm:"column:isPartial;default:false;not null" json:"isPartial"`
	IsCancelled      bool           `gorm:"column:isCancelled;default:false;not null" json:"isCancelled"`
	ServiceName      string         `gorm:"column:serviceName;not null" json:"serviceName"`
	ServiceNamespace string         `gorm:"column:serviceNamespace;not null" json:"serviceNamespace"`
	Level            string         `gorm:"column:level;default:'TRACE';not null" json:"level"`
	Kind             string         `gorm:"column:kind;default:'INTERNAL';not null" json:"kind"`
	Status           string         `gorm:"column:status;default:'UNSET';not null" json:"status"`
	Links            datatypes.JSON `gorm:"column:links" json:"links"`
	Events           datatypes.JSON `gorm:"column:events" json:"events"`
	// This is the time the event started in nanoseconds since the epoch
	StartTime int64 `gorm:"column:startTime;not null" json:"startTime"`
	// This is the duration of the event in nanoseconds
	Duration        int64   `gorm:"column:duration;default:0;not null" json:"duration"`
	AttemptID       *string `gorm:"column:attemptId" json:"attemptId"`
	AttemptNumber   *int    `gorm:"column:attemptNumber" json:"attemptNumber"`
	EnvironmentID   string  `gorm:"column:environmentId;not null" json:"environmentId"`
	EnvironmentType string  `gorm:"column:environmentType;not null" json:"environmentType"`
	OrganizationID  string  `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID       string  `gorm:"column:projectId;not null" json:"projectId"`
	ProjectRef      string  `gorm:"column:projectRef;not null" json:"projectRef"`
	RunID           string  `gorm:"column:runId;index:TaskEventPartitioned_runId_idx;not null" json:"runId"`
	RunIsTest       bool    `gorm:"column:runIsTest;default:false;not null" json:"runIsTest"`
	IdempotencyKey  *string `gorm:"column:idempotencyKey" json:"idempotencyKey"`
	TaskSlug        string  `gorm:"column:taskSlug;not null" json:"taskSlug"`
	TaskPath        *string `gorm:"column:taskPath" json:"taskPath"`
	TaskExportName  *string `gorm:"column:taskExportName" json:"taskExportName"`
	WorkerID        *string `gorm:"column:workerId" json:"workerId"`
	WorkerVersion   *string `gorm:"column:workerVersion" json:"workerVersion"`
	QueueID         *string `gorm:"column:queueId" json:"queueId"`
	QueueName       *string `gorm:"column:queueName" json:"queueName"`
	BatchID         *string `gorm:"column:batchId" json:"batchId"`
	// This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering
	Properties datatypes.JSON `gorm:"column:properties;not null" json:"properties"`
	// This represents all span attributes in the $metadata namespace, like $metadata.payload
	Metadata datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	// This represents all span attributes in the $style namespace, like $style
	Style datatypes.JSON `gorm:"column:style" json:"style"`
	// This represents all span attributes in the $output namespace, like $output
	Output datatypes.JSON `gorm:"column:output" json:"output"`
	// This represents the mimetype of the output, such as application/json or application/super+json
	OutputType              *string        `gorm:"column:outputType" json:"outputType"`
	Payload                 datatypes.JSON `gorm:"column:payload" json:"payload"`
	PayloadType             *string        `gorm:"column:payloadType" json:"payloadType"`
	CreatedAt               DateTime       `gorm:"column:createdAt;primaryKey;autoCreateTime:false" json:"createdAt"`
	UsageDurationMs         int            `gorm:"column:usageDurationMs;default:0;not null" json:"usageDurationMs"`
	UsageCostInCents        float64        `gorm:"column:usageCostInCents;default:0;not null" json:"usageCostInCents"`
	MachinePreset           *string        `gorm:"column:machinePreset" json:"machinePreset"`
	MachinePresetCPU        *float64       `gorm:"column:machinePresetCpu" json:"machinePresetCpu"`
	MachinePresetMemory     *float64       `gorm:"column:machinePresetMemory" json:"machinePresetMemory"`
	MachinePresetCentsPerMs *float64       `gorm:"column:machinePresetCentsPerMs" json:"machinePresetCentsPerMs"`
}

func (TaskEventPartitioned) TableName() string {
	return "TaskEventPartitioned"
}

func (m *TaskEventPartitioned) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type GithubAppInstallation struct {
	ID                  string         `gorm:"column:id;primaryKey" json:"id"`
	AppInstallationID   int64          `gorm:"column:appInstallationId;uniqueIndex;not null" json:"appInstallationId"`
	TargetID            int64          `gorm:"column:targetId;not null" json:"targetId"`
	TargetType          string         `gorm:"column:targetType;not null" json:"targetType"`
	AccountHandle       string         `gorm:"column:accountHandle;not null" json:"accountHandle"`
	Permissions         datatypes.JSON `gorm:"column:permissions" json:"permissions"`
	RepositorySelection string         `gorm:"column:repositorySelection;not null" json:"repositorySelection"`
	InstalledBy         *string        `gorm:"column:installedBy" json:"installedBy"`
	InstalledByUserID   *string        `gorm:"column:installedByUserId" json:"installedByUserId"`
	OrganizationID      string         `gorm:"column:organizationId;index:GithubAppInstallation_organizationId_idx;not null" json:"organizationId"`
	DeletedAt           *DateTime      `gorm:"column:deletedAt" json:"deletedAt"`
	SuspendedAt         *DateTime      `gorm:"column:suspendedAt" json:"suspendedAt"`
	CreatedAt           DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization        *Organization
	Repositories        []GithubRepository `gorm:"foreignKey:InstallationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (GithubAppInstallation) TableName() string {
	return "GithubAppInstallation"
}

func (m *GithubAppInstallation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GithubAppInstallation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GithubRepository struct {
	ID                        string                      `gorm:"column:id;primaryKey" json:"id"`
	GithubID                  int64                       `gorm:"column:githubId;uniqueIndex:GithubRepository_installationId_githubId_key;not null" json:"githubId"`
	Name                      string                      `gorm:"column:name;not null" json:"name"`
	FullName                  string                      `gorm:"column:fullName;not null" json:"fullName"`
	HTMLURL                   string                      `gorm:"column:htmlUrl;not null" json:"htmlUrl"`
	Private                   bool                        `gorm:"column:private;not null" json:"private"`
	DefaultBranch             string                      `gorm:"column:defaultBranch;not null" json:"defaultBranch"`
	InstallationID            string                      `gorm:"column:installationId;uniqueIndex:GithubRepository_installationId_githubId_key;index:GithubRepository_installationId_idx;not null" json:"installationId"`
	CreatedAt                 DateTime                    `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime                    `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Installation              *GithubAppInstallation      `gorm:"foreignKey:InstallationID"`
	ConnectedGithubRepository []ConnectedGithubRepository `gorm:"foreignKey:RepositoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (GithubRepository) TableName() string {
	return "GithubRepository"
}

func (m *GithubRepository) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GithubRepository) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TrustedDeployRepository struct {
	ID                string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID         string   `gorm:"column:projectId;uniqueIndex;not null" json:"projectId"`
	Issuer            *string  `gorm:"column:issuer;default:'https://token.actions.githubusercontent.com';not null" json:"issuer"`
	RepositoryOwner   string   `gorm:"column:repositoryOwner;not null" json:"repositoryOwner"`
	RepositoryName    string   `gorm:"column:repositoryName;not null" json:"repositoryName"`
	RepositoryID      *string  `gorm:"column:repositoryId" json:"repositoryId"`
	RepositoryOwnerID *string  `gorm:"column:repositoryOwnerId" json:"repositoryOwnerId"`
	CreatedAt         DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project           *Project
	Environments      []TrustedDeployEnvironment `gorm:"foreignKey:TrustedDeployRepositoryID;constraint:OnDelete:CASCADE"`
}

func (TrustedDeployRepository) TableName() string {
	return "TrustedDeployRepository"
}

func (m *TrustedDeployRepository) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TrustedDeployRepository) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TrustedDeployEnvironment struct {
	ID                        string         `gorm:"column:id;primaryKey" json:"id"`
	TrustedDeployRepositoryID string         `gorm:"column:trustedDeployRepositoryId;index:TrustedDeployEnvironment_trustedDeployRepositoryId_idx;not null" json:"trustedDeployRepositoryId"`
	RuntimeEnvironmentID      string         `gorm:"column:runtimeEnvironmentId;uniqueIndex;not null" json:"runtimeEnvironmentId"`
	Enabled                   bool           `gorm:"column:enabled;default:false;not null" json:"enabled"`
	Scope                     string         `gorm:"column:scope;not null" json:"scope"`
	ScopeValue                *string        `gorm:"column:scopeValue" json:"scopeValue"`
	AdditionalClaims          datatypes.JSON `gorm:"column:additionalClaims;default:'{}';not null" json:"additionalClaims"`
	CreatedAt                 DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	TrustedDeployRepository   *TrustedDeployRepository
	RuntimeEnvironment        *RuntimeEnvironment
}

func (TrustedDeployEnvironment) TableName() string {
	return "TrustedDeployEnvironment"
}

func (m *TrustedDeployEnvironment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TrustedDeployEnvironment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ConnectedGithubRepository struct {
	ID                        string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID                 string         `gorm:"column:projectId;uniqueIndex;not null" json:"projectId"`
	RepositoryID              string         `gorm:"column:repositoryId;index:ConnectedGithubRepository_repositoryId_idx;not null" json:"repositoryId"`
	BranchTracking            datatypes.JSON `gorm:"column:branchTracking;not null" json:"branchTracking"`
	PreviewDeploymentsEnabled *bool          `gorm:"column:previewDeploymentsEnabled;default:true;not null" json:"previewDeploymentsEnabled"`
	BuildFilters              datatypes.JSON `gorm:"column:buildFilters" json:"buildFilters"`
	CreatedAt                 DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project                   *Project
	Repository                *GithubRepository `gorm:"foreignKey:RepositoryID"`
}

func (ConnectedGithubRepository) TableName() string {
	return "ConnectedGithubRepository"
}

func (m *ConnectedGithubRepository) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ConnectedGithubRepository) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ImpersonationAuditLog struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Action    string   `gorm:"column:action;not null" json:"action"`
	AdminID   string   `gorm:"column:adminId;index:ImpersonationAuditLog_adminId_idx;not null" json:"adminId"`
	TargetID  string   `gorm:"column:targetId;index:ImpersonationAuditLog_targetId_idx;not null" json:"targetId"`
	IPAddress *string  `gorm:"column:ipAddress" json:"ipAddress"`
	CreatedAt DateTime `gorm:"column:createdAt;index:ImpersonationAuditLog_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	// The admin user who initiated/ended the impersonation
	Admin *User `gorm:"foreignKey:AdminID"`
	// The user being impersonated
	Target *User `gorm:"foreignKey:TargetID"`
}

func (ImpersonationAuditLog) TableName() string {
	return "ImpersonationAuditLog"
}

func (m *ImpersonationAuditLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SupportAccessRequest struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string    `gorm:"column:organizationId;index:SupportAccessRequest_organizationId_status_expiresAt_idx;not null" json:"organizationId"`
	RequestedByID  *string   `gorm:"column:requestedById;index:SupportAccessRequest_requestedById_idx" json:"requestedById"`
	Reason         string    `gorm:"column:reason;not null" json:"reason"`
	Status         string    `gorm:"column:status;index:SupportAccessRequest_organizationId_status_expiresAt_idx;default:'PENDING';not null" json:"status"`
	ApprovedByID   *string   `gorm:"column:approvedById;index:SupportAccessRequest_approvedById_idx" json:"approvedById"`
	ApprovedAt     *DateTime `gorm:"column:approvedAt" json:"approvedAt"`
	ExpiresAt      *DateTime `gorm:"column:expiresAt;index:SupportAccessRequest_organizationId_status_expiresAt_idx" json:"expiresAt"`
	CancelledByID  *string   `gorm:"column:cancelledById" json:"cancelledById"`
	CancelledAt    *DateTime `gorm:"column:cancelledAt" json:"cancelledAt"`
	CreatedAt      DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization
	RequestedBy    *User `gorm:"foreignKey:RequestedByID"`
	ApprovedBy     *User `gorm:"foreignKey:ApprovedByID"`
}

func (SupportAccessRequest) TableName() string {
	return "SupportAccessRequest"
}

func (m *SupportAccessRequest) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SupportAccessRequest) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CustomerQuery struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The TSQL query text that was executed
	Query string `gorm:"column:query;not null" json:"query"`
	// The scope of the query (determines which tenant IDs were used)
	Scope string `gorm:"column:scope;not null" json:"scope"`
	// Query execution statistics from ClickHouse
	Stats datatypes.JSON `gorm:"column:stats;not null" json:"stats"`
	// AI-generated title summarizing the query
	Title *string `gorm:"column:title" json:"title"`
	// Where the query originated from
	Source         string  `gorm:"column:source;default:'DASHBOARD';not null" json:"source"`
	OrganizationID string  `gorm:"column:organizationId;index:CustomerQuery_organizationId_createdAt_idx;not null" json:"organizationId"`
	ProjectID      *string `gorm:"column:projectId" json:"projectId"`
	EnvironmentID  *string `gorm:"column:environmentId" json:"environmentId"`
	UserID         *string `gorm:"column:userId" json:"userId"`
	// Time filter settings used when the query was executed
	// Period like "7d", "24h", etc.
	FilterPeriod *string `gorm:"column:filterPeriod" json:"filterPeriod"`
	// Custom start date (ISO 8601)
	FilterFrom *DateTime `gorm:"column:filterFrom" json:"filterFrom"`
	// Custom end date (ISO 8601)
	FilterTo     *DateTime `gorm:"column:filterTo" json:"filterTo"`
	CreatedAt    DateTime  `gorm:"column:createdAt;index:CustomerQuery_organizationId_createdAt_idx;index:CustomerQuery_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	Organization *Organization
	Project      *Project
	Environment  *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
	// Optional user who executed the query (null for API calls)
	User *User
}

func (CustomerQuery) TableName() string {
	return "CustomerQuery"
}

func (m *CustomerQuery) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type IntegrationDeployment struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// For example Vercel
	IntegrationName string `gorm:"column:integrationName;not null" json:"integrationName"`
	// External ID
	IntegrationDeploymentID string  `gorm:"column:integrationDeploymentId;not null" json:"integrationDeploymentId"`
	CommitSHA               string  `gorm:"column:commitSHA;index:IntegrationDeployment_commitSHA_idx;not null" json:"commitSHA"`
	DeploymentID            *string `gorm:"column:deploymentId;index:IntegrationDeployment_deploymentId_idx" json:"deploymentId"`
	// External deployment status
	Status           *string           `gorm:"column:status" json:"status"`
	CreatedAt        DateTime          `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime          `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	WorkerDeployment *WorkerDeployment `gorm:"foreignKey:DeploymentID"`
}

func (IntegrationDeployment) TableName() string {
	return "IntegrationDeployment"
}

func (m *IntegrationDeployment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*IntegrationDeployment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A user-defined metrics dashboard
type MetricsDashboard struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID     string   `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	Title          string   `gorm:"column:title;not null" json:"title"`
	Description    string   `gorm:"column:description;default:'';not null" json:"description"`
	OrganizationID string   `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID      *string  `gorm:"column:projectId;index:MetricsDashboard_projectId_createdAt_idx" json:"projectId"`
	OwnerID        *string  `gorm:"column:ownerId" json:"ownerId"`
	CreatedAt      DateTime `gorm:"column:createdAt;index:MetricsDashboard_projectId_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// JSON that defines the config, queries, layout and config of all widgets.
	// There will be a version field for the format.
	Layout       string `gorm:"column:layout;not null" json:"layout"`
	Organization *Organization
	Project      *Project
	// Who created the dashboard
	Owner *User `gorm:"foreignKey:OwnerID"`
}

func (MetricsDashboard) TableName() string {
	return "MetricsDashboard"
}

func (m *MetricsDashboard) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*MetricsDashboard) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A known LLM model or model pattern for cost tracking
type LlmModel struct {
	ID              string    `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID      string    `gorm:"column:friendly_id;uniqueIndex;not null" json:"friendly_id"`
	ProjectID       *string   `gorm:"column:project_id;uniqueIndex:llm_models_project_id_model_name_start_date_key;index:llm_models_project_id_idx" json:"project_id"`
	ModelName       string    `gorm:"column:model_name;uniqueIndex:llm_models_project_id_model_name_start_date_key;not null" json:"model_name"`
	MatchPattern    string    `gorm:"column:match_pattern;not null" json:"match_pattern"`
	StartDate       *DateTime `gorm:"column:start_date;uniqueIndex:llm_models_project_id_model_name_start_date_key" json:"start_date"`
	Source          *string   `gorm:"column:source;default:'default';not null" json:"source"`
	CreatedAt       DateTime  `gorm:"column:created_at;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt       DateTime  `gorm:"column:updated_at;autoUpdateTime:false;not null" json:"updated_at"`
	Provider        *string   `gorm:"column:provider" json:"provider"`
	Description     *string   `gorm:"column:description" json:"description"`
	ContextWindow   *int      `gorm:"column:context_window" json:"context_window"`
	MaxOutputTokens *int      `gorm:"column:max_output_tokens" json:"max_output_tokens"`
	Capabilities    []string  `gorm:"column:capabilities;serializer:json;not null" json:"capabilities"`
	IsHidden        bool      `gorm:"column:is_hidden;default:false;not null" json:"is_hidden"`
	BaseModelName   *string   `gorm:"column:base_model_name" json:"base_model_name"`
	PricingUnit     *string   `gorm:"column:pricing_unit" json:"pricing_unit"`
	Project         *Project
	PricingTiers    []LlmPricingTier `gorm:"foreignKey:ModelID;constraint:OnDelete:CASCADE"`
	Prices          []LlmPrice       `gorm:"foreignKey:ModelID;constraint:OnDelete:CASCADE"`
}

func (LlmModel) TableName() string {
	return "llm_models"
}

func (m *LlmModel) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*LlmModel) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A pricing tier for a model (supports volume-based or conditional pricing)
type LlmPricingTier struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	ModelID    string         `gorm:"column:model_id;uniqueIndex:llm_pricing_tiers_model_id_priority_key;uniqueIndex:llm_pricing_tiers_model_id_name_key;not null" json:"model_id"`
	Name       string         `gorm:"column:name;uniqueIndex:llm_pricing_tiers_model_id_name_key;not null" json:"name"`
	IsDefault  *bool          `gorm:"column:is_default;default:true;not null" json:"is_default"`
	Priority   int            `gorm:"column:priority;uniqueIndex:llm_pricing_tiers_model_id_priority_key;default:0;not null" json:"priority"`
	Conditions datatypes.JSON `gorm:"column:conditions;type:jsonb;default:'[]';not null" json:"conditions"`
	Model      *LlmModel      `gorm:"foreignKey:ModelID"`
	Prices     []LlmPrice     `gorm:"foreignKey:PricingTierID;constraint:OnDelete:CASCADE"`
}

func (LlmPricingTier) TableName() string {
	return "llm_pricing_tiers"
}

func (m *LlmPricingTier) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// A price point for a usage type within a pricing tier
type LlmPrice struct {
	ID            string          `gorm:"column:id;primaryKey" json:"id"`
	ModelID       string          `gorm:"column:model_id;uniqueIndex:llm_prices_model_id_usage_type_pricing_tier_id_key;not null" json:"model_id"`
	PricingTierID string          `gorm:"column:pricing_tier_id;uniqueIndex:llm_prices_model_id_usage_type_pricing_tier_id_key;not null" json:"pricing_tier_id"`
	UsageType     string          `gorm:"column:usage_type;uniqueIndex:llm_prices_model_id_usage_type_pricing_tier_id_key;not null" json:"usage_type"`
	Price         float64         `gorm:"column:price;type:decimal(20,12);not null" json:"price"`
	Model         *LlmModel       `gorm:"foreignKey:ModelID"`
	PricingTier   *LlmPricingTier `gorm:"foreignKey:PricingTierID"`
}

func (LlmPrice) TableName() string {
	return "llm_prices"
}

func (m *LlmPrice) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Admin-created notification definitions
type PlatformNotification struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`
	FriendlyID string `gorm:"column:friendlyId;uniqueIndex;not null" json:"friendlyId"`
	// Admin-facing title for identification in admin tools
	Title string `gorm:"column:title;not null" json:"title"`
	// Versioned JSON rendering content (see payload schema in spec)
	Payload        datatypes.JSON `gorm:"column:payload;not null" json:"payload"`
	Surface        string         `gorm:"column:surface;index:PlatformNotification_surface_scope_startsAt_idx;not null" json:"surface"`
	Scope          string         `gorm:"column:scope;index:PlatformNotification_surface_scope_startsAt_idx;not null" json:"scope"`
	UserID         *string        `gorm:"column:userId" json:"userId"`
	OrganizationID *string        `gorm:"column:organizationId" json:"organizationId"`
	ProjectID      *string        `gorm:"column:projectId" json:"projectId"`
	// When notification becomes active
	StartsAt DateTime `gorm:"column:startsAt;index:PlatformNotification_surface_scope_startsAt_idx;not null" json:"startsAt"`
	// When notification expires
	EndsAt DateTime `gorm:"column:endsAt;not null" json:"endsAt"`
	// CLI: auto-expire N days after user first saw it
	CliMaxDaysAfterFirstSeen *int `gorm:"column:cliMaxDaysAfterFirstSeen" json:"cliMaxDaysAfterFirstSeen"`
	// CLI: auto-expire after shown N times to user
	CliMaxShowCount *int `gorm:"column:cliMaxShowCount" json:"cliMaxShowCount"`
	// CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…)
	CliShowEvery *int `gorm:"column:cliShowEvery" json:"cliShowEvery"`
	// Ordering within same scope level (higher = more important)
	Priority int `gorm:"column:priority;default:0;not null" json:"priority"`
	// Staged-but-unpublished. While true the notification is hidden from all
	// user-facing reads regardless of startsAt/endsAt; publishing sets real
	// dates and flips this to false.
	IsDraft bool `gorm:"column:isDraft;default:false;not null" json:"isDraft"`
	// Soft delete
	ArchivedAt *DateTime `gorm:"column:archivedAt" json:"archivedAt"`
	CreatedAt  DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// Set when scope = USER
	User *User
	// Set when scope = ORGANIZATION
	Organization *Organization
	// Set when scope = PROJECT
	Project      *Project
	Interactions []PlatformNotificationInteraction `gorm:"foreignKey:NotificationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (PlatformNotification) TableName() string {
	return "PlatformNotification"
}

func (m *PlatformNotification) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.FriendlyID == "" {
		m.FriendlyID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PlatformNotification) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Per-user tracking of notification views and dismissals
type PlatformNotificationInteraction struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	NotificationID string `gorm:"column:notificationId;uniqueIndex:PlatformNotificationInteraction_notificationId_userId_key;not null" json:"notificationId"`
	UserID         string `gorm:"column:userId;uniqueIndex:PlatformNotificationInteraction_notificationId_userId_key;not null" json:"userId"`
	// Set by beacon/CLI GET on first impression
	FirstSeenAt DateTime `gorm:"column:firstSeenAt;not null" json:"firstSeenAt"`
	// Times shown (incremented per beacon/CLI view)
	ShowCount *int `gorm:"column:showCount;default:1;not null" json:"showCount"`
	// User dismissed in webapp
	WebappDismissedAt *DateTime `gorm:"column:webappDismissedAt" json:"webappDismissedAt"`
	// User clicked a link in webapp
	WebappClickedAt *DateTime `gorm:"column:webappClickedAt" json:"webappClickedAt"`
	// Auto-dismissed or expired in CLI
	CliDismissedAt *DateTime             `gorm:"column:cliDismissedAt" json:"cliDismissedAt"`
	CreatedAt      DateTime              `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime              `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Notification   *PlatformNotification `gorm:"foreignKey:NotificationID"`
	User           *User
}

func (PlatformNotificationInteraction) TableName() string {
	return "PlatformNotificationInteraction"
}

func (m *PlatformNotificationInteraction) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.FirstSeenAt.IsZero() {
		m.FirstSeenAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PlatformNotificationInteraction) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// *
// * Error group state is used to track when a user has interacted with an error (ignored/resolved)
// * The actual error data is in ClickHouse.
type ErrorGroupState struct {
	ID               string  `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID   string  `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID        string  `gorm:"column:projectId;not null" json:"projectId"`
	EnvironmentID    *string `gorm:"column:environmentId;uniqueIndex:ErrorGroupState_environmentId_taskIdentifier_errorFingerprint_key;index:ErrorGroupState_environmentId_status_idx" json:"environmentId"`
	TaskIdentifier   string  `gorm:"column:taskIdentifier;uniqueIndex:ErrorGroupState_environmentId_taskIdentifier_errorFingerprint_key;not null" json:"taskIdentifier"`
	ErrorFingerprint string  `gorm:"column:errorFingerprint;uniqueIndex:ErrorGroupState_environmentId_taskIdentifier_errorFingerprint_key;not null" json:"errorFingerprint"`
	Status           string  `gorm:"column:status;index:ErrorGroupState_environmentId_status_idx;default:'UNRESOLVED';not null" json:"status"`
	// *
	// * Error is ignored until this date
	IgnoredUntil *DateTime `gorm:"column:ignoredUntil" json:"ignoredUntil"`
	// *
	// * Error is ignored until this occurrence rate
	IgnoredUntilOccurrenceRate *int `gorm:"column:ignoredUntilOccurrenceRate" json:"ignoredUntilOccurrenceRate"`
	// *
	// * Error is ignored until this total occurrences
	IgnoredUntilTotalOccurrences *int `gorm:"column:ignoredUntilTotalOccurrences" json:"ignoredUntilTotalOccurrences"`
	// Total occurrence count at the time the error was ignored (from ClickHouse).
	// Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
	IgnoredAtOccurrenceCount *int64 `gorm:"column:ignoredAtOccurrenceCount" json:"ignoredAtOccurrenceCount"`
	// *
	// * Error was ignored at this date
	IgnoredAt *DateTime `gorm:"column:ignoredAt" json:"ignoredAt"`
	// *
	// * Reason for ignoring the error
	IgnoredReason *string `gorm:"column:ignoredReason" json:"ignoredReason"`
	// *
	// * User who ignored the error
	IgnoredByUserID *string `gorm:"column:ignoredByUserId" json:"ignoredByUserId"`
	// *
	// * Error was resolved at this date
	ResolvedAt *DateTime `gorm:"column:resolvedAt" json:"resolvedAt"`
	// *
	// * Error was resolved in this version
	ResolvedInVersion *string `gorm:"column:resolvedInVersion" json:"resolvedInVersion"`
	// *
	// * User who resolved the error
	ResolvedBy   *string  `gorm:"column:resolvedBy" json:"resolvedBy"`
	CreatedAt    DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization *Organization
	Project      *Project
	// *
	// * You can ignore/resolve an error across all environments, or specific ones
	Environment *RuntimeEnvironment `gorm:"foreignKey:EnvironmentID"`
}

func (ErrorGroupState) TableName() string {
	return "ErrorGroupState"
}

func (m *ErrorGroupState) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ErrorGroupState) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TaskIdentifier struct {
	ID                   string   `gorm:"column:id;primaryKey" json:"id"`
	RuntimeEnvironmentID string   `gorm:"column:runtimeEnvironmentId;uniqueIndex:TaskIdentifier_runtimeEnvironmentId_slug_key;index:TaskIdentifier_runtimeEnvironmentId_isInLatestDeployment_idx;not null" json:"runtimeEnvironmentId"`
	ProjectID            string   `gorm:"column:projectId;not null" json:"projectId"`
	Slug                 string   `gorm:"column:slug;uniqueIndex:TaskIdentifier_runtimeEnvironmentId_slug_key;not null" json:"slug"`
	CurrentTriggerSource string   `gorm:"column:currentTriggerSource;default:'STANDARD';not null" json:"currentTriggerSource"`
	CurrentWorkerID      *string  `gorm:"column:currentWorkerId" json:"currentWorkerId"`
	FirstSeenAt          DateTime `gorm:"column:firstSeenAt;not null" json:"firstSeenAt"`
	LastSeenAt           DateTime `gorm:"column:lastSeenAt;not null" json:"lastSeenAt"`
	IsInLatestDeployment *bool    `gorm:"column:isInLatestDeployment;index:TaskIdentifier_runtimeEnvironmentId_isInLatestDeployment_idx;default:true;not null" json:"isInLatestDeployment"`
	RuntimeEnvironment   *RuntimeEnvironment
	Project              *Project
	CurrentWorker        *BackgroundWorker `gorm:"foreignKey:CurrentWorkerID"`
}

func (TaskIdentifier) TableName() string {
	return "TaskIdentifier"
}

func (m *TaskIdentifier) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.FirstSeenAt.IsZero() {
		m.FirstSeenAt = DateTime{Time: now}
	}
	if m.LastSeenAt.IsZero() {
		m.LastSeenAt = DateTime{Time: now}
	}
	return nil
}

// Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
// Multiple organizations can share a single data store row via organizationIds.
type OrganizationDataStore struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// Human-readable unique key (e.g. "hipaa-clickhouse-us-east")
	Key string `gorm:"column:key;uniqueIndex;not null" json:"key"`
	// Organization IDs that use this data store
	OrganizationIds []string `gorm:"column:organizationIds;serializer:json;not null" json:"organizationIds"`
	Kind            string   `gorm:"column:kind;index:OrganizationDataStore_kind_idx;not null" json:"kind"`
	// Versioned config JSON. Structure is discriminated by the top-level `version` field.
	Config    datatypes.JSON `gorm:"column:config;not null" json:"config"`
	CreatedAt DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (OrganizationDataStore) TableName() string {
	return "OrganizationDataStore"
}

func (m *OrganizationDataStore) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganizationDataStore) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DateTime struct{ time.Time }

func (DateTime) GormDataType() string {
	return "time"
}

func (dateTime DateTime) Value() (driver.Value, error) {
	return dateTime.UTC().Truncate(time.Millisecond), nil
}

func (dateTime *DateTime) Scan(src any) error {
	read, err := readPrismaTime(src)
	dateTime.Time = read
	return err
}

var prismaTimeLayouts = []string{
	"2006-01-02T15:04:05.999999999Z07:00",
	"2006-01-02 15:04:05.999999999Z07:00",
	"2006-01-02 15:04:05.999999999Z07",
	"2006-01-02T15:04:05.999999999",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02",
}

func readPrismaTime(src any) (time.Time, error) {
	switch value := src.(type) {
	case nil:
		return time.Time{}, nil
	case time.Time:
		return value.UTC().Truncate(time.Millisecond), nil
	case int64:
		return time.UnixMilli(value).UTC(), nil
	case []byte:
		return readPrismaTime(string(value))
	case string:
		if millis, err := strconv.ParseInt(value, 10, 64); err == nil {
			return time.UnixMilli(millis).UTC(), nil
		}
		for _, layout := range prismaTimeLayouts {
			if read, err := time.Parse(layout, value); err == nil {
				return read.UTC().Truncate(time.Millisecond), nil
			}
		}
		return time.Time{}, fmt.Errorf("a DateTime column held %q", value)
	}
	return time.Time{}, fmt.Errorf("a DateTime column held %T", src)
}
