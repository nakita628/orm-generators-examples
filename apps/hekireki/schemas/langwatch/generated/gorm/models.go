package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lucsky/cuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Account struct {
	ID       string  `gorm:"column:id;primaryKey" json:"id"`
	UserID   string  `gorm:"column:userId;index:Account_userId_idx;not null" json:"userId"`
	Type     *string `gorm:"column:type;default:'oauth';not null" json:"type"`
	Provider string  `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key,priority:1;not null" json:"provider"`
	// WHO asserted `providerAccountId`. Better Auth keys accounts by
	// `(provider, providerAccountId)`; auth's account create hook writes this
	// (modules/auth/specs/account-issuer.feature). Nullable only so the column
	// could be added to a populated table; the migration backfilled every row.
	Issuer            *string   `gorm:"column:issuer;index:Account_issuer_providerAccountId_idx,priority:1" json:"issuer"`
	ProviderAccountID string    `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key,priority:2;index:Account_issuer_providerAccountId_idx,priority:2;not null" json:"providerAccountId"`
	RefreshToken      *string   `gorm:"column:refresh_token;type:text" json:"refresh_token"`
	AccessToken       *string   `gorm:"column:access_token;type:text" json:"access_token"`
	ExpiresAt         *DateTime `gorm:"column:expires_at;type:timestamp(3)" json:"expires_at"`
	ExtExpiresIn      *int      `gorm:"column:ext_expires_in;size:32" json:"ext_expires_in"`
	TokenType         *string   `gorm:"column:token_type" json:"token_type"`
	Scope             *string   `gorm:"column:scope" json:"scope"`
	IDToken           *string   `gorm:"column:id_token;type:text" json:"id_token"`
	SessionState      *string   `gorm:"column:session_state" json:"session_state"`
	Password          *string   `gorm:"column:password;type:text" json:"password"`
	CreatedAt         DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User              *User     `gorm:"belongsTo"`
}

func (Account) TableName() string {
	return "Account"
}

func (m *Account) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Account) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AccountCredential struct {
	ID                    string    `gorm:"column:id;primaryKey" json:"id"`
	UserID                string    `gorm:"column:userId;index:AccountCredential_userId_idx;not null" json:"userId"`
	Provider              string    `gorm:"column:provider;not null" json:"provider"`
	Password              *string   `gorm:"column:password;type:text" json:"password"`
	AccessToken           *string   `gorm:"column:accessToken;type:text" json:"accessToken"`
	RefreshToken          *string   `gorm:"column:refreshToken;type:text" json:"refreshToken"`
	IDToken               *string   `gorm:"column:idToken;type:text" json:"idToken"`
	AccessTokenExpiresAt  *DateTime `gorm:"column:accessTokenExpiresAt;type:timestamp(3)" json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt *DateTime `gorm:"column:refreshTokenExpiresAt;type:timestamp(3)" json:"refreshTokenExpiresAt"`
	Scope                 *string   `gorm:"column:scope" json:"scope"`
	CreatedAt             DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User                  *User     `gorm:"belongsTo"`
}

func (AccountCredential) TableName() string {
	return "AccountCredential"
}

func (m *AccountCredential) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AccountCredential) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Session struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	SessionToken string   `gorm:"column:sessionToken;uniqueIndex:Session_sessionToken_key;not null" json:"sessionToken"`
	UserID       string   `gorm:"column:userId;index:Session_userId_idx;not null" json:"userId"`
	Expires      DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	IPAddress    *string  `gorm:"column:ipAddress" json:"ipAddress"`
	UserAgent    *string  `gorm:"column:userAgent;type:text" json:"userAgent"`
	// Which of the user's sign-in methods minted this session (D06). NULL on
	// every session that predates the column and on any mint we could not
	// attribute - an ordinary value, not an error. Nothing reads it to decide
	// whether the session is valid, so no session was ended to add it.
	IdentifierID *string `gorm:"column:identifierId;index:Session_identifierId_idx" json:"identifierId"`
	// What the sign-in that minted this session PROVED, as RFC 8176 method
	// references (D06). Empty on every session minted before the column, and
	// on any ordinary password sign-in. An organization that requires a second
	// factor reads this when a member reaches its data; it is never read at
	// mint and never ends a session.
	Amr ScalarList[string] `gorm:"column:amr;type:text[]" json:"amr"`
	// Minted while the account had never signed in: the session its sign-up
	// started, which skips the secure-account offer. False on every later one
	// and on every session that predates the column (specs/identity/passkeys.feature).
	StartedBySignUp bool `gorm:"column:startedBySignUp;default:false;not null" json:"startedBySignUp"`
	// The legacy impersonation payload. NOTHING READS OR WRITES IT any more -
	// the {actor, subject} columns below replaced it, and the migration beside
	// them ended every session that still carried a value.
	//
	// It stays for ONE release, and that is the whole reason it is still here.
	// `prisma migrate deploy` runs at container start, so the first pod of a
	// new release drops a column while the previous release's pods are still
	// serving - and those pods select this column on EVERY authenticated
	// request and fail closed when the row does not come back. Dropping it in
	// the same release that stops reading it signs everybody out for the length
	// of the rollout. Expand now, contract next release.
	Impersonating datatypes.JSON `gorm:"column:impersonating" json:"impersonating"`
	// The impersonation claims (D06), aligned with the authz Principal
	// {actor, subject}. The actor is the operator who really made the request
	// and the subject is the person whose access they are borrowing; both are
	// recorded on every authorization decision, so the audit trail can answer
	// who really did it. These replace the legacy `impersonating` JSON blob,
	// which is dropped in the release AFTER this one.
	//
	// NULL on every ordinary session, which is all of them: an impersonation
	// is a handful of rows at a time and nothing else reads these columns.
	ActorUserID   *string `gorm:"column:actorUserId" json:"actorUserId"`
	SubjectUserID *string `gorm:"column:subjectUserId" json:"subjectUserId"`
	// Why the operator started it. Required to start one
	// (specs/features/backoffice-user-impersonation-reason.feature) and kept
	// beside both people rather than only in the audit log, so a request under
	// impersonation carries its own justification.
	ImpersonationReason *string `gorm:"column:impersonationReason;type:text" json:"impersonationReason"`
	// When the borrowed access lapses. A session outlives its impersonation
	// window: past this the claims are ignored and the operator is themselves
	// again, with nothing ended.
	ImpersonationExpiresAt *DateTime `gorm:"column:impersonationExpiresAt;type:timestamp(3)" json:"impersonationExpiresAt"`
	// When this session was last used, for the idle timeout an organization
	// may set (GAC-10, specs/identity/org-session-lifetime.feature).
	//
	// A COLUMN OF OUR OWN rather than `updatedAt`, because better-auth rolls a
	// live session's expiry once per `updateAge` - a day - so `updatedAt` is
	// activity to the nearest day and cannot tell an idle hour from an idle
	// minute. Written only for sessions whose owner belongs to an organization
	// that has set a window, and then at most once per quarter-window, so a
	// deployment where nobody has turned this on writes it never.
	//
	// NULL on every session minted before the column and on every session
	// under no window at all. A null READS AS `updatedAt`, which better-auth
	// keeps rolled to within a day of the last use. That fallback is
	// conservative in the direction that matters - `updatedAt` is never
	// earlier than the real last use, so it can only ever judge a session MORE
	// active than it was, never less - and it is what lets a window turned on
	// today end the sessions that have been idle since February instead of
	// waiting a day for a stamp they have never had.
	LastSeenAt *DateTime `gorm:"column:lastSeenAt;type:timestamp(3)" json:"lastSeenAt"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User       *User     `gorm:"belongsTo"`
}

func (Session) TableName() string {
	return "Session"
}

func (m *Session) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
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

// One row per erased account, holding only its creation time: no id, name or address. It keeps
// the install's first-created account knowable after an erasure deletes that row.
// Spec: modules/user/specs/user.feature
type UserErasure struct {
	ID               string   `gorm:"column:id;primaryKey" json:"id"`
	AccountCreatedAt DateTime `gorm:"column:accountCreatedAt;index:UserErasure_accountCreatedAt_idx;type:timestamp(3);not null" json:"accountCreatedAt"`
	ErasedAt         DateTime `gorm:"column:erasedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"erasedAt"`
}

func (UserErasure) TableName() string {
	return "UserErasure"
}

func (m *UserErasure) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.ErasedAt.IsZero() {
		m.ErasedAt = DateTime{Time: now}
	}
	return nil
}

type User struct {
	ID                        string  `gorm:"column:id;primaryKey" json:"id"`
	Name                      *string `gorm:"column:name" json:"name"`
	Email                     *string `gorm:"column:email;uniqueIndex:User_email_key" json:"email"`
	EmailVerified             bool    `gorm:"column:emailVerified;default:false;not null" json:"emailVerified"`
	SignupConfirmationPending bool    `gorm:"column:signupConfirmationPending;default:false;not null" json:"signupConfirmationPending"`
	PasskeySignupClaimHash    *string `gorm:"column:passkeySignupClaimHash" json:"passkeySignupClaimHash"`
	Image                     *string `gorm:"column:image;type:text" json:"image"`
	PendingSsoSetup           bool    `gorm:"column:pendingSsoSetup;default:false;not null" json:"pendingSsoSetup"`
	UserHashKey               *string `gorm:"column:userHashKey" json:"userHashKey"`
	// Owned by the two-factor plugin (D06), which is why it is a plain
	// boolean rather than a state: the plugin flips it when a setup is
	// confirmed and when it is turned off, and challenges every sign-in while
	// it is true. The lifecycle - and the history - lives on MfaEnrollment.
	TwoFactorEnabled              bool      `gorm:"column:twoFactorEnabled;default:false;not null" json:"twoFactorEnabled"`
	CreatedAt                     DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                     DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	LastLoginAt                   *DateTime `gorm:"column:lastLoginAt;type:timestamp(3)" json:"lastLoginAt"`
	DeactivatedAt                 *DateTime `gorm:"column:deactivatedAt;type:timestamp(3)" json:"deactivatedAt"`
	LastHomePath                  *string   `gorm:"column:lastHomePath" json:"lastHomePath"`
	TracesExplorerTourDismissedAt *DateTime `gorm:"column:tracesExplorerTourDismissedAt;type:timestamp(3)" json:"tracesExplorerTourDismissedAt"`
	// How Langy should reach this person's code when a task needs a change
	// (ADR-129). "github" means they chose the pull request path and asked to
	// be remembered; NULL means ask. A shared local folder is never stored
	// here, because a folder has to be shared again each time.
	// Spec: specs/langy/langy-code-access.feature
	LangyCodeAccessPreference *string `gorm:"column:langyCodeAccessPreference" json:"langyCodeAccessPreference"`
	// When this person last said "not now" to being offered a passkey
	// (ADR-120). On the account rather than in browser storage for two
	// reasons: a per-browser record forgets on a new device, which is exactly
	// where the offer should be MORE eager, and an interval nobody can enforce
	// is not an interval. NULL means never asked.
	PasskeyNudgeDismissedAt *DateTime `gorm:"column:passkeyNudgeDismissedAt;type:timestamp(3)" json:"passkeyNudgeDismissedAt"`
	// The domains this person has said "no thanks" to being offered their
	// colleagues on (D12). Per domain rather than one flag: somebody who
	// changes jobs and verifies a second address should be offered their new
	// employer's organization, not silenced forever by a decision they made
	// about the old one. Empty means never dismissed anything.
	JoinOfferDismissedDomains ScalarList[string] `gorm:"column:joinOfferDismissedDomains;type:text[]" json:"joinOfferDismissedDomains"`
	// This person's answer about browser notifications, one entry per topic:
	// { "langy": "enabled" | "declined" }. A topic that is absent was never
	// answered. The browser's own permission is not stored here.
	// Spec: specs/langy/langy-notifications.feature
	NotificationPreferences datatypes.JSON `gorm:"column:notificationPreferences;default:'{}';not null" json:"notificationPreferences"`
	// When an operator marked this person's password as no longer trusted. A
	// password sign-in opens no session while it is set; any write of a new
	// password hash clears it. Spec: modules/auth/specs/forced-password-reset.feature
	PasswordResetRequiredAt  *DateTime                `gorm:"column:passwordResetRequiredAt;type:timestamp(3)" json:"passwordResetRequiredAt"`
	Accounts                 []Account                `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccountCredentials       []AccountCredential      `gorm:"foreignKey:UserID;constraint:AccountCredential_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions                 []Session                `gorm:"foreignKey:UserID;constraint:Session_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TwoFactors               []TwoFactor              `gorm:"foreignKey:UserID;constraint:TwoFactor_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Passkeys                 []Passkey                `gorm:"foreignKey:UserID;constraint:Passkey_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TeamMemberships          []TeamUser               `gorm:"foreignKey:UserID;constraint:TeamUser_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OrgMemberships           []OrganizationUser       `gorm:"foreignKey:UserID;constraint:OrganizationUser_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ScimUserResources        []ScimUserResource       `gorm:"foreignKey:UserID;constraint:ScimUserResource_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Annotation               []Annotation             `gorm:"foreignKey:UserID;constraint:Annotation_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ShareLinks               []ShareLink              `gorm:"foreignKey:UserID;constraint:ShareLink_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Workflow                 []Workflow               `gorm:"foreignKey:PublishedByID;constraint:Workflow_publishedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	WorkflowVersion          []WorkflowVersion        `gorm:"foreignKey:AuthorID;constraint:WorkflowVersion_authorId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AnnotationQueues         []AnnotationQueueMembers `gorm:"foreignKey:UserID;constraint:AnnotationQueueMembers_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AssignedQueueItems       []AnnotationQueueItem    `gorm:"foreignKey:UserID;constraint:AnnotationQueueItem_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedQueueItems        []AnnotationQueueItem    `gorm:"foreignKey:CreatedByUserID;constraint:AnnotationQueueItem_createdByUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	LlmPromptConfigVersions  []LlmPromptConfigVersion `gorm:"foreignKey:AuthorID;constraint:LlmPromptConfigVersion_authorId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Scenarios                []Scenario               `gorm:"foreignKey:LastUpdatedByID;constraint:Scenario_lastUpdatedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	InviteRequests           []OrganizationInvite     `gorm:"foreignKey:RequestedBy;constraint:OrganizationInvite_requestedBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	SecretsCreated           []ProjectSecret          `gorm:"foreignKey:CreatedByID;constraint:ProjectSecret_createdById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SecretsUpdated           []ProjectSecret          `gorm:"foreignKey:UpdatedByID;constraint:ProjectSecret_updatedById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SavedViews               []SavedView              `gorm:"foreignKey:UserID;constraint:SavedView_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	VersionTagsCreated       []PromptTagAssignment    `gorm:"foreignKey:CreatedByID;constraint:PromptTagAssignment_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	VersionTagsUpdated       []PromptTagAssignment    `gorm:"foreignKey:UpdatedByID;constraint:PromptTagAssignment_updatedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PromptTagsCreated        []PromptTag              `gorm:"foreignKey:CreatedByID;constraint:PromptTag_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PromptTagsUpdated        []PromptTag              `gorm:"foreignKey:UpdatedByID;constraint:PromptTag_updatedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	GroupMemberships         []GroupMembership        `gorm:"foreignKey:UserID;constraint:GroupMembership_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleBindings             []RoleBinding            `gorm:"foreignKey:UserID;constraint:RoleBinding_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ScimExternalIds          []ScimExternalID         `gorm:"foreignKey:UserID;constraint:ScimExternalId_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeys                  []APIKey                 `gorm:"foreignKey:UserID;constraint:ApiKey_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	GatewayBudgetsCreated    []GatewayBudget          `gorm:"foreignKey:CreatedByID;constraint:GatewayBudget_createdById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PrincipalVirtualKeys     []VirtualKey             `gorm:"foreignKey:PrincipalUserID;constraint:VirtualKey_principalUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PersonalTeams            []Team                   `gorm:"foreignKey:OwnerUserID;constraint:Team_ownerUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PersonalProjects         []Project                `gorm:"foreignKey:OwnerUserID;constraint:Project_ownerUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	RoutingPoliciesCreated   []RoutingPolicy          `gorm:"foreignKey:CreatedByID;constraint:RoutingPolicy_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	RoutingPoliciesUpdated   []RoutingPolicy          `gorm:"foreignKey:UpdatedByID;constraint:RoutingPolicy_updatedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	GatewayGuardrailsCreated []GatewayGuardrail       `gorm:"foreignKey:CreatedByID;constraint:GatewayGuardrail_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	GatewayGuardrailsUpdated []GatewayGuardrail       `gorm:"foreignKey:UpdatedByID;constraint:GatewayGuardrail_updatedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	IngestionSourcesCreated  []IngestionSource        `gorm:"foreignKey:CreatedByID;constraint:IngestionSource_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AnomalyRulesCreated      []AnomalyRule            `gorm:"foreignKey:CreatedByID;constraint:AnomalyRule_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PinnedTraces             []PinnedTrace            `gorm:"foreignKey:UserID;constraint:PinnedTrace_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TraceEditOverlaysCreated []TraceEditOverlay       `gorm:"foreignKey:CreatedByID;constraint:TraceEditOverlay_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TraceEditOverlaysUpdated []TraceEditOverlay       `gorm:"foreignKey:UpdatedByID;constraint:TraceEditOverlay_updatedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (User) TableName() string {
	return "User"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
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

type VerificationToken struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`
	Identifier string `gorm:"column:identifier;uniqueIndex:VerificationToken_identifier_token_key,priority:1;not null" json:"identifier"`
	// Not unique alone: better-auth stores plain values here (a user id, a counter). See WEB-19201.
	Token     string   `gorm:"column:token;uniqueIndex:VerificationToken_identifier_token_key,priority:2;not null" json:"token"`
	Expires   DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

func (m *VerificationToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*VerificationToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Identifier struct {
	ID                string    `gorm:"column:id;primaryKey" json:"id"`
	UserID            string    `gorm:"column:userId;index:Identifier_userId_idx;not null" json:"userId"`
	Provider          string    `gorm:"column:provider;index:Identifier_provider_providerAccountId_idx,priority:1;not null" json:"provider"`
	Value             *string   `gorm:"column:value;index:Identifier_value_idx" json:"value"`
	Domain            *string   `gorm:"column:domain;index:Identifier_domain_idx" json:"domain"`
	IdentifierHash    *string   `gorm:"column:identifierHash;index:Identifier_identifierHash_idx" json:"identifierHash"`
	AccountID         *string   `gorm:"column:accountId;index:Identifier_accountId_idx" json:"accountId"`
	ProviderID        *string   `gorm:"column:providerId" json:"providerId"`
	Issuer            *string   `gorm:"column:issuer;index:Identifier_issuer_providerAccountId_idx,priority:1" json:"issuer"`
	ProviderAccountID *string   `gorm:"column:providerAccountId;index:Identifier_provider_providerAccountId_idx,priority:2;index:Identifier_issuer_providerAccountId_idx,priority:2" json:"providerAccountId"`
	State             string    `gorm:"column:state;not null" json:"state"`
	ConnectionID      *string   `gorm:"column:connectionId" json:"connectionId"`
	VerifiedAt        *DateTime `gorm:"column:verifiedAt;type:timestamp(3)" json:"verifiedAt"`
	AttachedAt        DateTime  `gorm:"column:attachedAt;type:timestamp(3);not null" json:"attachedAt"`
	DetachedAt        *DateTime `gorm:"column:detachedAt;type:timestamp(3)" json:"detachedAt"`
	LastUsedAt        *DateTime `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	CreatedAt         DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Identifier) TableName() string {
	return "Identifier"
}

func (m *Identifier) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Identifier) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Whether a user signs in on the identifier path: identity's own switch, set when the
// identifier backfill proves the user and pinned off by an operator's rollback. No row is the
// legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
type IdentitySignInSwitch struct {
	UserID             string   `gorm:"column:userId;primaryKey" json:"userId"`
	OnIdentifiers      bool     `gorm:"column:onIdentifiers;not null" json:"onIdentifiers"`
	Pinned             bool     `gorm:"column:pinned;default:false;not null" json:"pinned"`
	SwitchedAt         DateTime `gorm:"column:switchedAt;type:timestamp(3);not null" json:"switchedAt"`
	LastAppliedEventID *string  `gorm:"column:lastAppliedEventId" json:"lastAppliedEventId"`
	UpdatedAt          DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (IdentitySignInSwitch) TableName() string {
	return "IdentitySignInSwitch"
}

func (m *IdentitySignInSwitch) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*IdentitySignInSwitch) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type IdentifierReservation struct {
	NormalizedValue string   `gorm:"column:normalizedValue;primaryKey" json:"normalizedValue"`
	UserID          string   `gorm:"column:userId;index:IdentifierReservation_userId_idx;not null" json:"userId"`
	IdentifierID    string   `gorm:"column:identifierId;index:IdentifierReservation_identifierId_idx;not null" json:"identifierId"`
	CommandID       string   `gorm:"column:commandId;not null" json:"commandId"`
	CreatedAt       DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (IdentifierReservation) TableName() string {
	return "IdentifierReservation"
}

func (m *IdentifierReservation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type IdentityProjectionCursor struct {
	UserID            string   `gorm:"column:userId;primaryKey" json:"userId"`
	LastEventID       string   `gorm:"column:lastEventId;not null" json:"lastEventId"`
	AcceptedAt        DateTime `gorm:"column:acceptedAt;type:timestamp(3);not null" json:"acceptedAt"`
	OccurredAt        DateTime `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	ProjectionVersion string   `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	CreatedAt         DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (IdentityProjectionCursor) TableName() string {
	return "IdentityProjectionCursor"
}

func (m *IdentityProjectionCursor) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*IdentityProjectionCursor) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The two-factor plugin's own table (D06). Protocol state, and row-truth for
// good (ADR-101 R12): the shared secret and the backup codes live here and
// in no event, ever. Nothing in this repo reads or writes these columns -
// the plugin owns them end to end, including the lockout counters, which is
// why `failedVerificationCount` and `lockedUntil` are here rather than
// rebuilt on the projection.
//
// `backupCodes` holds an ENCRYPTED blob, not hashes. A backup code has to be
// compared against what somebody types, and the plugin decrypts and compares;
// hashing them would leave it unable to verify its own codes.
type TwoFactor struct {
	ID                      string    `gorm:"column:id;primaryKey" json:"id"`
	UserID                  string    `gorm:"column:userId;index:TwoFactor_userId_idx;not null" json:"userId"`
	Secret                  string    `gorm:"column:secret;index:TwoFactor_secret_idx;type:text;not null" json:"secret"`
	BackupCodes             string    `gorm:"column:backupCodes;type:text;not null" json:"backupCodes"`
	Verified                *bool     `gorm:"column:verified;default:true;not null" json:"verified"`
	FailedVerificationCount int       `gorm:"column:failedVerificationCount;size:32;default:0;not null" json:"failedVerificationCount"`
	LockedUntil             *DateTime `gorm:"column:lockedUntil;type:timestamp(3)" json:"lockedUntil"`
	CreatedAt               DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User                    *User     `gorm:"belongsTo"`
}

func (TwoFactor) TableName() string {
	return "TwoFactor"
}

func (m *TwoFactor) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TwoFactor) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Consecutive failed sign-ins, and the lock they earn (GAC-09).
// Spec: specs/identity/org-account-lockout.feature.
//
// KEYED ON THE ADDRESS THAT WAS TYPED, NOT ON A USER, and that is the whole
// design. We already answer "no such account" and "wrong password" with one
// refusal so that sign-in cannot be used to discover who has an account. A
// lock that existed only for real accounts would put that back: five wrong
// guesses saying "locked" for one address and "incorrect" for another tells
// an attacker which addresses are worth attacking. So an address with
// nothing behind it gets a row, a count and a lock exactly like one that
// resolves, and both answer identically.
//
// `identifierHash` is a KEYED hash of the normalised address, never the
// address. The key is the deployment's own secret, so this table cannot be
// lifted and reversed into the list of every address anybody has ever tried
// to sign in as - which is what it would otherwise become, including every
// typo and every address that has no account here at all.
//
// Row-truth, deliberately (ADR-101 R12): this is protocol state read and
// written on the sign-in path, like TwoFactor's own counters next door. The
// EVIDENCE an auditor reads is the identity event emitted when a lock is
// taken, which is durable; these rows are reaped once they stop meaning
// anything.
type SignInAttemptLock struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// Keyed hash of the normalised address. Never the address itself.
	IdentifierHash string `gorm:"column:identifierHash;uniqueIndex:SignInAttemptLock_identifierHash_key;not null" json:"identifierHash"`
	// Consecutive failures since the last success. Reset by a success, not by
	// time - "five consecutive", not five ever.
	FailedCount int `gorm:"column:failedCount;size:32;default:0;not null" json:"failedCount"`
	// When the current lock lifts. Null when not locked.
	LockedUntil *DateTime `gorm:"column:lockedUntil;index:SignInAttemptLock_lockedUntil_idx;type:timestamp(3)" json:"lockedUntil"`
	// Locks taken since the last success. At five this stops lifting by
	// itself and becomes an incident - see `heldForReview`.
	ConsecutiveLockouts int `gorm:"column:consecutiveLockouts;size:32;default:0;not null" json:"consecutiveLockouts"`
	// Held past its lock because the count of locks reached the ceiling the
	// control names. Cleared by an administrator, or by the owner proving
	// mailbox control through a password reset - the second is what keeps a
	// held account from being a denial of service anyone can inflict by
	// typing a colleague's address wrong often enough.
	HeldForReview bool `gorm:"column:heldForReview;default:false;not null" json:"heldForReview"`
	// Whose account this resolved to at the time, when it resolved to one.
	// Null for an address with no account - those rows are the reason this is
	// nullable, and the reason this table is not hung off User.
	UserID    *string  `gorm:"column:userId;index:SignInAttemptLock_userId_idx" json:"userId"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SignInAttemptLock) TableName() string {
	return "SignInAttemptLock"
}

func (m *SignInAttemptLock) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SignInAttemptLock) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
// public key and the signature counter are the plugin's to read and write.
// A passkey also becomes an `Identifier` row with provider `passkey`, which
// is what makes it a sign-in method like any other - this table is the
// credential, that one is the identity.
//
// `aaguid` names the authenticator MODEL, never a device or a person, and is
// only present at registration. Privacy-preserving platforms report all
// zeroes, so a label derived from it is best-effort by construction.
type Passkey struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	Name         *string  `gorm:"column:name" json:"name"`
	PublicKey    string   `gorm:"column:publicKey;type:text;not null" json:"publicKey"`
	UserID       string   `gorm:"column:userId;index:Passkey_userId_idx;not null" json:"userId"`
	CredentialID string   `gorm:"column:credentialID;index:Passkey_credentialID_idx;not null" json:"credentialID"`
	Counter      int      `gorm:"column:counter;size:32;default:0;not null" json:"counter"`
	DeviceType   string   `gorm:"column:deviceType;not null" json:"deviceType"`
	BackedUp     bool     `gorm:"column:backedUp;default:false;not null" json:"backedUp"`
	Transports   *string  `gorm:"column:transports" json:"transports"`
	Aaguid       *string  `gorm:"column:aaguid" json:"aaguid"`
	CreatedAt    DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User         *User    `gorm:"belongsTo"`
}

func (Passkey) TableName() string {
	return "Passkey"
}

func (m *Passkey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Passkey) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The two-step verification pipeline's projection (D06 - see
// specs/identity/mfa-and-session-shape.feature). One row per PERSON, because
// one person has one setup however many organizations they belong to.
//
// A pure event-truth head like `SsoConnection`: every column is fold-written,
// rows are never deleted (DISABLED is a tombstone, so history survives being
// turned off), and a replay rebuilds it whole-row from the event log. The
// cursor columns are inline because the fold produces exactly one row per
// aggregate - the `Identifier` fan-out needed its own cursor table, this
// does not.
//
// Every column here is a lifecycle fact. There is deliberately nowhere in
// this model a secret or a backup code could be put: `consumedBackupCodeIndexes`
// holds POSITIONS, which is what lets "how many are left" be answered
// without the projection ever knowing a code.
type MfaEnrollment struct {
	// The user. Also the aggregate id and the tenant - one identity history
	// per person (ADR-029 §4).
	UserID string `gorm:"column:userId;primaryKey" json:"userId"`
	// The current enrollment, or NULL for somebody who never started one.
	EnrollmentID *string `gorm:"column:enrollmentId" json:"enrollmentId"`
	// totp. Never sms.
	Method *string `gorm:"column:method" json:"method"`
	// NONE | PENDING | ENABLED | EXPIRED | DISABLED
	State       string    `gorm:"column:state;index:MfaEnrollment_state_idx;not null" json:"state"`
	EnrolledAt  *DateTime `gorm:"column:enrolledAt;type:timestamp(3)" json:"enrolledAt"`
	ConfirmedAt *DateTime `gorm:"column:confirmedAt;type:timestamp(3)" json:"confirmedAt"`
	ExpiredAt   *DateTime `gorm:"column:expiredAt;type:timestamp(3)" json:"expiredAt"`
	DisabledAt  *DateTime `gorm:"column:disabledAt;type:timestamp(3)" json:"disabledAt"`
	// password+totp | org-admin - who ended it, and on whose authority.
	DisabledVia     *string `gorm:"column:disabledVia" json:"disabledVia"`
	BackupCodeCount int     `gorm:"column:backupCodeCount;size:32;default:0;not null" json:"backupCodeCount"`
	// Positions already spent, ascending. Never codes.
	ConsumedBackupCodeIndexes ScalarList[int] `gorm:"column:consumedBackupCodeIndexes;type:integer[]" json:"consumedBackupCodeIndexes"`
	FailedCount               int             `gorm:"column:failedCount;size:32;default:0;not null" json:"failedCount"`
	OccurredAt                DateTime        `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	LastEventID               string          `gorm:"column:lastEventId;not null" json:"lastEventId"`
	AcceptedAt                DateTime        `gorm:"column:acceptedAt;type:timestamp(3);not null" json:"acceptedAt"`
	ProjectionVersion         string          `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	CreatedAt                 DateTime        `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime        `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (MfaEnrollment) TableName() string {
	return "MfaEnrollment"
}

func (m *MfaEnrollment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*MfaEnrollment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The SSO connection pipeline's projection (ADR-117 §5, D04 - see
// specs/identity/sso-connection-lifecycle.feature). A pure event-truth head
// like `Identifier`: every column is fold-written, rows are never deleted
// (TORN_DOWN is a tombstone), and a replay rebuilds it whole-row from the
// event log.
//
// The fold cursor lives on the row rather than in a sibling table, because
// one row IS one aggregate here - a connection - so there is nothing for a
// separate cursor table to key on that this row does not already key on.
//
// No relation to Organization on purpose, matching `Identifier`: a
// projection is rebuilt from events, and a foreign key would make a replay
// depend on relational state the log does not carry. `idpMetadata` holds
// endpoints and credential REFERENCES; no secret is ever stored here or in
// any event.
type SsoConnection struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:SsoConnection_organizationId_idx;index:SsoConnection_organizationId_migrationPhase_idx,priority:1;not null" json:"organizationId"`
	// oidc | saml
	Type string `gorm:"column:type;not null" json:"type"`
	// DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
	// VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
	// TEARDOWN_PENDING | TORN_DOWN
	State          string             `gorm:"column:state;index:SsoConnection_state_idx;not null" json:"state"`
	ClaimedDomains ScalarList[string] `gorm:"column:claimedDomains;type:text[]" json:"claimedDomains"`
	// [{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
	// decidedByActorId, authority, waitedMs, note }] - every claim this
	// connection has made and how long each waited to be decided. The tier-3
	// approval queue is a read of this column, and the epic's Open Q2 wants
	// that latency measured from the day the queue exists.
	DomainClaims    datatypes.JSON     `gorm:"column:domainClaims;default:'[]';not null" json:"domainClaims"`
	ApprovedDomains ScalarList[string] `gorm:"column:approvedDomains;type:text[]" json:"approvedDomains"`
	// The only domains that ever route, and only while state is ACTIVE.
	VerifiedDomains ScalarList[string] `gorm:"column:verifiedDomains;index:SsoConnection_verifiedDomains_idx,type:gin;type:text[]" json:"verifiedDomains"`
	// The subset of `verifiedDomains` whose published record stayed missing
	// through its grace window (ADR-123). They still ROUTE - everyone already
	// here signs in unchanged - and they stop vouching for anybody NEW: no
	// provisioning on first sign-in, no walking in by domain. A column rather
	// than a read of the JSON below, because both questions are asked on a
	// sign-in path and have to be one indexed predicate.
	LapsedDomains ScalarList[string] `gorm:"column:lapsedDomains;index:SsoConnection_lapsedDomains_idx,type:gin;type:text[]" json:"lapsedDomains"`
	// [{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
	// graceEndsAtMs }] - what proved each verified domain, who, and whether
	// that evidence is still there. Kept on the row rather than only in the
	// log so that an operator-attested domain can never be read as one the
	// customer proved, and a wavering one never as one nothing is doubting.
	DomainVerifications datatypes.JSON `gorm:"column:domainVerifications;default:'[]';not null" json:"domainVerifications"`
	// { domain, method, tokenHash } while a ceremony is in flight.
	PendingVerification datatypes.JSON `gorm:"column:pendingVerification" json:"pendingVerification"`
	IdpMetadata         datatypes.JSON `gorm:"column:idpMetadata;not null" json:"idpMetadata"`
	// admit | request | refuse - what this connection does with somebody who
	// signs in through it and is not a member yet (ADR-117 3). Stated at
	// registration and changed by the setup journey, so it is never absent
	// and no reader has to decide what absence means. This was two fields
	// for a while - a boolean `allowsJit` beside it, kept in step by the
	// fold - and they disagreed about the middle answer.
	ArrivalPolicy *string `gorm:"column:arrivalPolicy;default:'refuse';not null" json:"arrivalPolicy"`
	// DEAD, AND STAYING FOR ONE RELEASE. Nothing reads or writes this; the
	// routing repository derives the same answer from `arrivalPolicy`. It is
	// declared so the schema still describes the table, and left in place so
	// the PREVIOUS release's pods — which select it by name on every
	// connection read — keep working through the rollout. Dropping it in the
	// release that stopped reading it would stop single sign-on routing and
	// the connection projection for the length of the deploy. Dropped by
	// `20260827120000_drop_sso_connection_allows_jit` in the next one.
	AllowsJit bool `gorm:"column:allowsJit;default:false;not null" json:"allowsJit"`
	// When somebody CHOSE the policy above, and null while the registration
	// default stands. A different fact from the policy rather than a second
	// copy of it: a connection always has a behaviour, and separately
	// somebody has or has not decided on it. Going live waits for the
	// deciding - "turn everybody away" is a decision too.
	ArrivalPolicyDecidedAt *DateTime `gorm:"column:arrivalPolicyDecidedAt;type:timestamp(3)" json:"arrivalPolicyDecidedAt"`
	// self-serve | legacy-grandfathered
	Source             string  `gorm:"column:source;not null" json:"source"`
	TestLoginAccountID *string `gorm:"column:testLoginAccountId" json:"testLoginAccountId"`
	// The grandfathered connection this direct connection replaces. Null for
	// ordinary connections and for the grandfathered predecessor itself.
	//
	// Partial uniqueness for nonterminal replacements is installed by the
	// flattened migration; Prisma cannot express that predicate here.
	ReplacesConnectionID *string `gorm:"column:replacesConnectionId" json:"replacesConnectionId"`
	// SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
	// Null means this connection is not a legacy-to-direct replacement.
	MigrationPhase *string   `gorm:"column:migrationPhase;index:SsoConnection_organizationId_migrationPhase_idx,priority:2" json:"migrationPhase"`
	GraceStartedAt *DateTime `gorm:"column:graceStartedAt;type:timestamp(3)" json:"graceStartedAt"`
	// When an administrator last selected the normal route. This is audit
	// evidence only: routing is derived from `migrationPhase`, never recency.
	RouteChangedAt          *DateTime `gorm:"column:routeChangedAt;type:timestamp(3)" json:"routeChangedAt"`
	FinalizationRequestedAt *DateTime `gorm:"column:finalizationRequestedAt;type:timestamp(3)" json:"finalizationRequestedAt"`
	FinalizedAt             *DateTime `gorm:"column:finalizedAt;type:timestamp(3)" json:"finalizedAt"`
	// { domain, note } - the last ops rejection, read back on a re-claim.
	Rejection datatypes.JSON `gorm:"column:rejection" json:"rejection"`
	CreatedBy *string        `gorm:"column:createdBy" json:"createdBy"`
	// While TEARDOWN_PENDING: when the process manager's wake completes it.
	TearDownAfter     *DateTime `gorm:"column:tearDownAfter;type:timestamp(3)" json:"tearDownAfter"`
	OccurredAt        DateTime  `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	LastEventID       string    `gorm:"column:lastEventId;not null" json:"lastEventId"`
	AcceptedAt        DateTime  `gorm:"column:acceptedAt;type:timestamp(3);not null" json:"acceptedAt"`
	ProjectionVersion string    `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	// Event-derived, not database-managed: `@default(now())` / `@updatedAt`
	// would make a replayed row differ from the row it rebuilds, which is
	// exactly the parity this projection has to keep.
	CreatedAt     DateTime                    `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime                    `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	ReproofCursor *SsoConnectionReproofCursor `gorm:"foreignKey:ConnectionID;constraint:SsoConnectionReproofCursor_connectionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (SsoConnection) TableName() string {
	return "SsoConnection"
}

// The one-statement pre-event registration lock. `SsoConnection` remains
// event truth; this row only decides which command may append that truth.
// A slot whose projected connection is terminal may be atomically replaced.
type SsoConnectionRegistrationSlot struct {
	OrganizationID string `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	// legacy | direct
	Kind                 string   `gorm:"column:kind;primaryKey" json:"kind"`
	ConnectionID         string   `gorm:"column:connectionId;uniqueIndex:SsoConnectionRegistrationSlot_connectionId_key;not null" json:"connectionId"`
	ReplacesConnectionID *string  `gorm:"column:replacesConnectionId" json:"replacesConnectionId"`
	CommandID            string   `gorm:"column:commandId;not null" json:"commandId"`
	CreatedAt            DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SsoConnectionRegistrationSlot) TableName() string {
	return "SsoConnectionRegistrationSlot"
}

func (m *SsoConnectionRegistrationSlot) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SsoConnectionRegistrationSlot) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Operational round-robin position for ADR-123's domain re-proof sweep.
//
// This is deliberately not a column on `SsoConnection`: that row is a
// whole-row event projection and its fold is its only writer, while a
// healthy re-read states no fact. Keeping the scheduling cursor beside the
// projection preserves both rules and still lets a fixed-size sweep revisit
// every connection, including one whose proof has started wavering.
type SsoConnectionReproofCursor struct {
	ConnectionID  string         `gorm:"column:connectionId;primaryKey" json:"connectionId"`
	LastReproofAt DateTime       `gorm:"column:lastReproofAt;index:SsoConnectionReproofCursor_lastReproofAt_idx;type:timestamp(3);not null" json:"lastReproofAt"`
	Connection    *SsoConnection `gorm:"foreignKey:ConnectionID;belongsTo"`
}

func (SsoConnectionReproofCursor) TableName() string {
	return "SsoConnectionReproofCursor"
}

// The values an `SsoConnection`'s credential references point AT (D09 - see
// specs/identity/sso-idp-termination.feature).
//
// The aggregate has always carried `clientIdRef` / `secretRef` / `certRefs`
// rather than values, because a fact may not carry a secret (ADR-101 §4).
// Until D09 there was nothing behind the references and every writer passed
// null. This is the thing behind them.
//
// `ciphertext` is AES-256-GCM under the deployment's `CREDENTIALS_SECRET`,
// the same key every other credential in the product is kept under. Nothing
// reads a row except the projection that builds the engine's dialing
// configuration; no surface renders one back.
//
// Written by the command path rather than the fold, and that split is
// deliberate: the VALUE is not event truth and cannot be, so the log records
// WHEN a credential changed (a new reference in a new fact) and never what
// it was. A replay rebuilds every reference from the log and finds the value
// still here.
type SsoCredential struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:SsoCredential_organizationId_connectionId_idx,priority:1;not null" json:"organizationId"`
	// The connection whose registration minted this. Not a relation, matching
	// `SsoConnection` itself: a projection rebuilt from events must not depend
	// on relational state the log does not carry.
	ConnectionID string `gorm:"column:connectionId;index:SsoCredential_organizationId_connectionId_idx,priority:2;not null" json:"connectionId"`
	// oidc-client-id | oidc-client-secret | saml-idp-config
	Kind       string   `gorm:"column:kind;not null" json:"kind"`
	Ciphertext string   `gorm:"column:ciphertext;not null" json:"ciphertext"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SsoCredential) TableName() string {
	return "SsoCredential"
}

func (m *SsoCredential) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SsoCredential) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// better-auth's single sign-on plugin's provider table (D09), and a
// PROJECTION of the connection log rather than a second source of truth.
//
// The plugin owns the columns and the reads: it is the only thing that
// speaks OpenID Connect discovery or reads a SAML assertion, and it dials
// what it finds here. What it never does is decide that a row should exist —
// that is folded from `SsoConnection`'s events by the same projection that
// writes the connection head, in the same transaction, so a replay rebuilds
// both together and a row can never outlive the connection that justified
// it. Nothing else writes here.
//
// `providerId` is the CONNECTION ID rather than the name the customer typed.
// The plugin's column is globally unique, and two organizations both calling
// their provider `okta` must both be able to.
//
// `oidcConfig` and `samlConfig` are the plugin's own JSON documents, rebuilt
// on every fold from `SsoCredential` — the vault is where the material
// lives, and this row is a derived copy the engine can read synchronously.
//
// SEALED AT REST, under the same key the vault uses. The document carries
// the client secret the derivation reads OUT of the vault, so a plaintext
// copy here would hand every customer's live identity-provider credential
// to anybody holding a database dump — the exact threat `SsoCredential`
// exists to defeat, reopened one table over. A stored value announces its
// form (`enc:v1:`), so a row written before the seal existed still opens;
// the one reader is the identity storage adapter, on the path that dials.
// The reader tolerates both forms, so this needs no backfill — but a pod
// without it cannot read a sealed row, so it must not be rolled back past.
// WHO OWNS A PROVED DOMAIN, as a row the database can refuse twice.
//
// "First verifier owns" was enforced by reading `findDomainOwner` and then
// writing the fact — two different aggregates, so the per-connection queue
// never ordered them against each other, and `SsoConnection.verifiedDomains`
// is a `String[]` with only a GIN index behind it, which enforces nothing.
// Two organizations proving `acme.com` at the same moment both read "nobody
// owns it" and both committed; the sign-in router then picked between two
// live connections by planner order, so a member of one company could be
// handed to another company's identity provider — and with `arrivalPolicy`
// on `admit`, provisioned into it.
//
// The legacy single-domain model had this: `Organization.ssoDomain` is
// `@unique`. The multi-domain rewrite dropped the guarantee, and self-serve
// proof is what made the race reachable by customers rather than by two
// operators working at once.
//
// `domain` is the primary key, so a second organization loses on the
// constraint rather than on a check it raced. Connection holders are
// separate because an exact predecessor/replacement pair must coexist
// during migration. Both are projected with the connection head.
type SsoVerifiedDomain struct {
	Domain         string                    `gorm:"column:domain;primaryKey;uniqueIndex:SsoVerifiedDomain_domain_organizationId_key,priority:1" json:"domain"`
	OrganizationID string                    `gorm:"column:organizationId;uniqueIndex:SsoVerifiedDomain_domain_organizationId_key,priority:2;index:SsoVerifiedDomain_organizationId_idx;not null" json:"organizationId"`
	Holders        []SsoVerifiedDomainHolder `gorm:"foreignKey:Domain,OrganizationID;references:Domain,OrganizationID;constraint:SsoVerifiedDomainHolder_domain_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (SsoVerifiedDomain) TableName() string {
	return "SsoVerifiedDomain"
}

// A connection currently entitled to use an organization-owned verified
// domain. Usually there is one holder. During an Auth0 migration there are
// exactly two: the grandfathered predecessor and its direct replacement.
//
// The projection repository enforces that relationship before inserting a
// second holder. The flattened migration also installs partial unique live
// legacy/direct slots on `SsoConnection`, closing concurrent third-writer
// races that Prisma cannot describe declaratively.
type SsoVerifiedDomainHolder struct {
	Domain         string             `gorm:"column:domain;primaryKey" json:"domain"`
	ConnectionID   string             `gorm:"column:connectionId;primaryKey;index:SsoVerifiedDomainHolder_connectionId_idx" json:"connectionId"`
	OrganizationID string             `gorm:"column:organizationId;index:SsoVerifiedDomainHolder_organizationId_idx;not null" json:"organizationId"`
	Ownership      *SsoVerifiedDomain `gorm:"foreignKey:Domain,OrganizationID;references:Domain,OrganizationID;belongsTo"`
}

func (SsoVerifiedDomainHolder) TableName() string {
	return "SsoVerifiedDomainHolder"
}

// A successful callback through one SSO connection. Identifier.lastUsedAt
// also advances during routing and resolution, so it cannot prove that a
// member actually authenticated through the replacement before cutover.
type SsoAuthenticationActivity struct {
	ID              string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID  string   `gorm:"column:organizationId;index:SsoAuthenticationActivity_organizationId_connectionId_authe_idx,priority:1;index:SsoAuthenticationActivity_organizationId_userId_idx,priority:1;not null" json:"organizationId"`
	ConnectionID    string   `gorm:"column:connectionId;index:SsoAuthenticationActivity_organizationId_connectionId_authe_idx,priority:2;not null" json:"connectionId"`
	UserID          string   `gorm:"column:userId;index:SsoAuthenticationActivity_organizationId_userId_idx,priority:2;not null" json:"userId"`
	AuthenticatedAt DateTime `gorm:"column:authenticatedAt;index:SsoAuthenticationActivity_organizationId_connectionId_authe_idx,priority:3;type:timestamp(3);not null" json:"authenticatedAt"`
	// The provider-side account the assertion named; null before #7633's arrival door recorded it.
	ProviderAccountID *string `gorm:"column:providerAccountId" json:"providerAccountId"`
}

func (SsoAuthenticationActivity) TableName() string {
	return "SsoAuthenticationActivity"
}

type SsoProvider struct {
	ID             string  `gorm:"column:id;primaryKey" json:"id"`
	Issuer         string  `gorm:"column:issuer;not null" json:"issuer"`
	OidcConfig     *string `gorm:"column:oidcConfig" json:"oidcConfig"`
	SamlConfig     *string `gorm:"column:samlConfig" json:"samlConfig"`
	UserID         *string `gorm:"column:userId" json:"userId"`
	ProviderID     string  `gorm:"column:providerId;uniqueIndex:SsoProvider_providerId_key;not null" json:"providerId"`
	OrganizationID *string `gorm:"column:organizationId;index:SsoProvider_organizationId_idx" json:"organizationId"`
	Domain         string  `gorm:"column:domain;index:SsoProvider_domain_idx;not null" json:"domain"`
}

func (SsoProvider) TableName() string {
	return "SsoProvider"
}

// A way back in (D05 - see specs/identity/sso-onboarding-tiers.feature).
//
// One named person who may still sign in without the identity provider
// while their organization's sign-in belongs to one, and the date on which
// that stops. Activation asks whether such a person exists, because the
// failure it prevents is the only one that cannot be recovered from inside
// the product: a misconfigured connection goes live, every member is handed
// to an identity provider that will not authenticate them, and nobody is
// left who can turn it off.
//
// Rows are immutable except for `supersededAt` and `warnedDays`. A renewal
// writes a NEW row naming the one it replaced, so the date a way in
// previously ended is still readable in the history - which is what makes a
// renewal auditable rather than a column somebody moved.
//
// No relation to Organization or User, matching the identity heads above:
// the binding is the record of who could get in, and a cascade would erase
// exactly the history a lockout post-mortem reads.
type SsoBreakGlassBinding struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:SsoBreakGlassBinding_organizationId_expiresAt_idx,priority:1;not null" json:"organizationId"`
	// Who holds the way in.
	UserID string `gorm:"column:userId;index:SsoBreakGlassBinding_userId_idx;not null" json:"userId"`
	// Who granted it. A way back in is never self-served.
	GrantedByUserID string   `gorm:"column:grantedByUserId;not null" json:"grantedByUserId"`
	GrantedAt       DateTime `gorm:"column:grantedAt;type:timestamp(3);not null" json:"grantedAt"`
	ExpiresAt       DateTime `gorm:"column:expiresAt;index:SsoBreakGlassBinding_organizationId_expiresAt_idx,priority:2;type:timestamp(3);not null" json:"expiresAt"`
	// Set when a renewal replaced this row; the row itself is never edited.
	SupersededAt *DateTime `gorm:"column:supersededAt;type:timestamp(3)" json:"supersededAt"`
	// The binding this one renewed, so the previous end date stays readable.
	RenewedFromID *string `gorm:"column:renewedFromId" json:"renewedFromId"`
	// Which of the fourteen, seven and one day warnings have been sent, so a
	// sweep that runs twice in a day never says the same thing twice.
	WarnedDays ScalarList[int] `gorm:"column:warnedDays;type:integer[]" json:"warnedDays"`
	CreatedAt  DateTime        `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (SsoBreakGlassBinding) TableName() string {
	return "SsoBreakGlassBinding"
}

func (m *SsoBreakGlassBinding) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// A durable claim on one live recovery path while activation is between its
// guard and projected ACTIVE fact. It has no expiry: only projection of the
// matching activation, or an explicit cancellation, may consume it.
type SsoActivationRecoveryReservation struct {
	CommandID      string   `gorm:"column:commandId;primaryKey" json:"commandId"`
	OrganizationID string   `gorm:"column:organizationId;uniqueIndex:SsoActivationRecoveryReservation_organizationId_connectionI_key,priority:1;index:SsoActivationRecoveryReservation_organizationId_idx;not null" json:"organizationId"`
	ConnectionID   string   `gorm:"column:connectionId;uniqueIndex:SsoActivationRecoveryReservation_organizationId_connectionI_key,priority:2;not null" json:"connectionId"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
}

func (SsoActivationRecoveryReservation) TableName() string {
	return "SsoActivationRecoveryReservation"
}

// The join-request pipeline's projection (ADR-117, D12 - see
// specs/identity/join-requests.feature). A pure event-truth head like
// `Identifier` and `SsoConnection`: every column is fold-written, rows are
// never deleted (a terminal state is a tombstone the admin panel stops
// listing, not an absence), and a replay rebuilds it whole-row from the
// event log.
//
// The fold cursor lives on the row rather than in a sibling table, because
// one row IS one aggregate here - a request.
//
// No relation to Organization or User on purpose, matching the two heads
// above: a projection is rebuilt from events, and a foreign key would make a
// replay depend on relational state the log does not carry.
//
// Membership is NOT here. An approval dispatches an attach on the grants
// ledger, exactly as accepting an invitation does, and that ledger fact
// carries `source: "join-request"` as its provenance - which is what puts a
// join approval on the customer's audit page beside every other membership
// change rather than in a history of its own.
//
// The requester's address is not here either: the DOMAIN is the fact, and it
// is the only part of an address a join decision may depend on.
type JoinRequest struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	UserID         string `gorm:"column:userId;index:JoinRequest_userId_state_idx,priority:1;index:JoinRequest_userId_organizationId_idx,priority:1;not null" json:"userId"`
	OrganizationID string `gorm:"column:organizationId;index:JoinRequest_organizationId_state_idx,priority:1;index:JoinRequest_userId_organizationId_idx,priority:2;not null" json:"organizationId"`
	// The normalized domain the match was made on.
	Domain string `gorm:"column:domain;index:JoinRequest_domain_idx;not null" json:"domain"`
	// PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN
	State string `gorm:"column:state;index:JoinRequest_organizationId_state_idx,priority:2;index:JoinRequest_userId_state_idx,priority:2;not null" json:"state"`
	// How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`.
	MatchedVia string `gorm:"column:matchedVia;not null" json:"matchedVia"`
	// Where the request was made: `web` or `cli` (ADR-171 v6). A request from
	// the terminal lands a Developer seat whatever the joiner seat says.
	Origin *string `gorm:"column:origin;default:'web';not null" json:"origin"`
	// While PENDING: when the day-14 wake lapses it. Null once resolved.
	ExpiresAt  *DateTime `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	ResolvedAt *DateTime `gorm:"column:resolvedAt;type:timestamp(3)" json:"resolvedAt"`
	// user | policy | invite - who ended it. `policy` is domain auto-join's
	// own principal, deliberately not the one SSO auto-join uses.
	ResolvedByType *string `gorm:"column:resolvedByType" json:"resolvedByType"`
	// The admin's user id, `domain-auto`, or the invitation id.
	ResolvedByID *string `gorm:"column:resolvedById" json:"resolvedById"`
	// user | invite-accepted, on a withdrawal.
	WithdrawalCause *string `gorm:"column:withdrawalCause" json:"withdrawalCause"`
	// The SSO connection a single sign-on arrival was admitted through; null otherwise.
	ConnectionID      *string  `gorm:"column:connectionId" json:"connectionId"`
	OccurredAt        DateTime `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	LastEventID       string   `gorm:"column:lastEventId;not null" json:"lastEventId"`
	AcceptedAt        DateTime `gorm:"column:acceptedAt;type:timestamp(3);not null" json:"acceptedAt"`
	ProjectionVersion string   `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	// Event-derived, not database-managed: `@default(now())` / `@updatedAt`
	// would make a replayed row differ from the row it rebuilds, which is
	// exactly the parity this projection has to keep.
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (JoinRequest) TableName() string {
	return "JoinRequest"
}

type TeamUser struct {
	UserID         string      `gorm:"column:userId;primaryKey;index:TeamUser_userId_idx" json:"userId"`
	TeamID         string      `gorm:"column:teamId;primaryKey;index:TeamUser_teamId_idx" json:"teamId"`
	Role           string      `gorm:"column:role;not null" json:"role"`
	AssignedRoleID *string     `gorm:"column:assignedRoleId;index:TeamUser_assignedRoleId_idx" json:"assignedRoleId"`
	CreatedAt      DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime    `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	AssignedRole   *CustomRole `gorm:"foreignKey:AssignedRoleID;belongsTo"`
	User           *User       `gorm:"belongsTo"`
	Team           *Team       `gorm:"belongsTo"`
}

func (TeamUser) TableName() string {
	return "TeamUser"
}

func (m *TeamUser) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TeamUser) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganizationUser struct {
	UserID         string   `gorm:"column:userId;primaryKey;index:OrganizationUser_userId_idx" json:"userId"`
	OrganizationID string   `gorm:"column:organizationId;primaryKey;index:OrganizationUser_organizationId_idx" json:"organizationId"`
	Role           string   `gorm:"column:role;not null" json:"role"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	// Department assignment for spend attribution (the people lens).
	// Single-valued, org-scoped, accounting-only (never gates access).
	// Plain column ref (relationMode="prisma"); the service validates the
	// id belongs to this org. See departments.feature.
	DepartmentID *string `gorm:"column:departmentId;index:OrganizationUser_departmentId_idx" json:"departmentId"`
	// Set when an admin disables this membership to stay within the licensed
	// seat count. The person loses access to this organization only, keeps
	// their role, department and everything they did, and frees their seat.
	// Reversible: re-enabling restores access if a seat is available.
	// Seat counting excludes disabled memberships. See
	// seat-reconciliation.feature.
	DisabledAt *DateTime `gorm:"column:disabledAt;type:timestamp(3)" json:"disabledAt"`
	// Stable grant intent for an unfinished automatic SSO admission.
	// Cleared after the authoritative grant is observed; never inferred from missing access.
	PendingSsoGrantID *string `gorm:"column:pendingSsoGrantId" json:"pendingSsoGrantId"`
	// A new generation for each membership lifetime. Stamped USER grant facts
	// must match the current active membership before they can be projected;
	// removing and rejoining therefore cannot revive a delayed old attach.
	MembershipStamp string `gorm:"column:membershipStamp;default:gen_random_uuid()::text;not null" json:"membershipStamp"`
	// When this membership started holding the full seat it holds now; null means createdAt.
	// Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
	FullSeatSince *DateTime `gorm:"column:fullSeatSince;type:timestamp(3)" json:"fullSeatSince"`
	// Set when admission wrote this membership disabled because no seat was free: the person
	// is waiting for one. Cleared when the membership is enabled.
	SeatPendingSince *DateTime     `gorm:"column:seatPendingSince;type:timestamp(3)" json:"seatPendingSince"`
	User             *User         `gorm:"belongsTo"`
	Organization     *Organization `gorm:"belongsTo"`
}

func (OrganizationUser) TableName() string {
	return "OrganizationUser"
}

func (m *OrganizationUser) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganizationUser) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A full seat still counted after the person who took it stopped holding it (removed,
// disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
// One per organization and person; a return within the hold takes the seat back.
type SeatHold struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:SeatHold_organizationId_userId_key,priority:1;not null" json:"organizationId"`
	UserID         string `gorm:"column:userId;uniqueIndex:SeatHold_organizationId_userId_key,priority:2;not null" json:"userId"`
	// Display copy for the Users list, kept because a removed person has no membership row.
	Email *string `gorm:"column:email" json:"email"`
	Name  *string `gorm:"column:name" json:"name"`
	// When the person took the seat this hold keeps.
	SeatTakenAt DateTime `gorm:"column:seatTakenAt;type:timestamp(3);not null" json:"seatTakenAt"`
	// seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly.
	HeldUntil DateTime `gorm:"column:heldUntil;index:SeatHold_heldUntil_idx;type:timestamp(3);not null" json:"heldUntil"`
	// REMOVED | DISABLED | DEMOTED
	Reason    string   `gorm:"column:reason;not null" json:"reason"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (SeatHold) TableName() string {
	return "SeatHold"
}

func (m *SeatHold) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// A person asked for a full seat the organization had no room for (modules/organization/specs/
// seat-requests.feature). One row per organization and address; reopening reuses the row.
type SeatRequest struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:SeatRequest_organizationId_email_key,priority:1;index:SeatRequest_organizationId_status_idx,priority:1;not null" json:"organizationId"`
	// Lowercased address of the person asking.
	Email string `gorm:"column:email;uniqueIndex:SeatRequest_organizationId_email_key,priority:2;not null" json:"email"`
	// Null when only the address is known (a sign-in the identity provider refused).
	UserID *string `gorm:"column:userId;index:SeatRequest_userId_idx" json:"userId"`
	Name   *string `gorm:"column:name" json:"name"`
	// OPEN | APPROVED | DECLINED
	Status string `gorm:"column:status;index:SeatRequest_organizationId_status_idx,priority:2;not null" json:"status"`
	// pending | lite | idp_refused: how the person was left when the request opened.
	Source string `gorm:"column:source;not null" json:"source"`
	// The identity provider also refused this person's sign-in.
	IdpRefused bool `gorm:"column:idpRefused;default:false;not null" json:"idpRefused"`
	// okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml
	IdpKind *string `gorm:"column:idpKind" json:"idpKind"`
	// When sales were told about this request; reset when it reopens.
	SalesNotifiedAt  *DateTime `gorm:"column:salesNotifiedAt;type:timestamp(3)" json:"salesNotifiedAt"`
	OpenedAt         DateTime  `gorm:"column:openedAt;type:timestamp(3);not null" json:"openedAt"`
	ResolvedAt       *DateTime `gorm:"column:resolvedAt;type:timestamp(3)" json:"resolvedAt"`
	ResolvedByUserID *string   `gorm:"column:resolvedByUserId" json:"resolvedByUserId"`
	// Set when an approval claims the request, cleared once its seat is given; an approval
	// still pending long after is one whose seat never landed, and is reopened.
	SeatPendingSince *DateTime `gorm:"column:seatPendingSince;type:timestamp(3)" json:"seatPendingSince"`
	CreatedAt        DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SeatRequest) TableName() string {
	return "SeatRequest"
}

func (m *SeatRequest) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SeatRequest) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Held by the one seat-taking write running on an organization (seat-requests.feature). A
// row past expiresAt is free to take; a holder deletes its own row when its write is done.
type OrganizationSeatLease struct {
	OrganizationID string `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	// Random per hold, so a release deletes only its own lease.
	Holder    string   `gorm:"column:holder;not null" json:"holder"`
	ExpiresAt DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (OrganizationSeatLease) TableName() string {
	return "OrganizationSeatLease"
}

func (m *OrganizationSeatLease) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// The mail log for seat requests and refused sign-ins: at most one row per organization,
// kind, address and UTC day, each mailed at once.
type OrganizationAccessNotice struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:OrganizationAccessNotice_organizationId_kind_email_dayKey_key,priority:1;not null" json:"organizationId"`
	// seat_request | idp_refused
	Kind  string  `gorm:"column:kind;uniqueIndex:OrganizationAccessNotice_organizationId_kind_email_dayKey_key,priority:2;not null" json:"kind"`
	Email string  `gorm:"column:email;uniqueIndex:OrganizationAccessNotice_organizationId_kind_email_dayKey_key,priority:3;not null" json:"email"`
	Name  *string `gorm:"column:name" json:"name"`
	Idp   *string `gorm:"column:idp" json:"idp"`
	// UTC day the notice was recorded on, YYYY-MM-DD.
	DayKey string `gorm:"column:dayKey;uniqueIndex:OrganizationAccessNotice_organizationId_kind_email_dayKey_key,priority:4;not null" json:"dayKey"`
	// Whether the seat request it announces was also refused by the identity provider.
	IdpRefused bool     `gorm:"column:idpRefused;default:false;not null" json:"idpRefused"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// Set once the notice was handed to the mailer; null if a send never started.
	MailedAt *DateTime `gorm:"column:mailedAt;type:timestamp(3)" json:"mailedAt"`
}

func (OrganizationAccessNotice) TableName() string {
	return "OrganizationAccessNotice"
}

func (m *OrganizationAccessNotice) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Team struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	Slug           string    `gorm:"column:slug;uniqueIndex:Team_slug_key;not null" json:"slug"`
	OrganizationID string    `gorm:"column:organizationId;index:Team_organizationId_idx;not null" json:"organizationId"`
	CreatedAt      DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	ArchivedAt     *DateTime `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	IsPersonal     bool      `gorm:"column:isPersonal;index:Team_isPersonal_idx;default:false;not null" json:"isPersonal"`
	OwnerUserID    *string   `gorm:"column:ownerUserId;type:varchar(255)" json:"ownerUserId"`
	// Department assignment inherited as the default for members and
	// projects that have none of their own. Accounting-only.
	DepartmentID     *string           `gorm:"column:departmentId;index:Team_departmentId_idx" json:"departmentId"`
	Organization     *Organization     `gorm:"belongsTo"`
	OwnerUser        *User             `gorm:"foreignKey:OwnerUserID;belongsTo"`
	Members          []TeamUser        `gorm:"foreignKey:TeamID;constraint:TeamUser_teamId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Projects         []Project         `gorm:"foreignKey:TeamID;constraint:Project_teamId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	IngestionSources []IngestionSource `gorm:"foreignKey:TeamID;constraint:IngestionSource_teamId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AiToolEntryTeams []AiToolEntryTeam `gorm:"foreignKey:TeamID;constraint:AiToolEntryTeam_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Team) TableName() string {
	return "Team"
}

func (m *Team) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Team) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Organization struct {
	ID   string `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
	// `customer` for every organization a person founds or an operator provisions; `system` for
	// the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
	// Spec: modules/ops/specs/checkup-canary-project.feature.
	Kind                  *string  `gorm:"column:kind;default:'customer';not null" json:"kind"`
	PhoneNumber           *string  `gorm:"column:phoneNumber" json:"phoneNumber"`
	Slug                  string   `gorm:"column:slug;uniqueIndex:Organization_slug_key;not null" json:"slug"`
	CreatedAt             DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	UsageSpendingMaxLimit *int     `gorm:"column:usageSpendingMaxLimit;size:32" json:"usageSpendingMaxLimit"`
	// The largest file a dataset image or file cell accepts for this
	// organization, in mebibytes. Null means the default. Set by an operator
	// from the back office only; every other dataset size limit scales with it.
	DatasetAttachmentMaxMb *int `gorm:"column:datasetAttachmentMaxMb;size:32" json:"datasetAttachmentMaxMb"`
	// The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
	// organization module the first time it is seen. Once set, a missing route refuses the
	// organization instead of serving it from the shared ClickHouse. Null: never private.
	// Spec: specs/private-dataplane/clickhouse-routing.feature.
	PrivateClickHouseLabel *string `gorm:"column:privateClickHouseLabel" json:"privateClickHouseLabel"`
	// The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature.
	PrivateS3Label *string `gorm:"column:privateS3Label" json:"privateS3Label"`
	// When either label was first recorded.
	PrivateDataplaneSeenAt *DateTime `gorm:"column:privateDataplaneSeenAt;type:timestamp(3)" json:"privateDataplaneSeenAt"`
	// Phase 8 - admin-configurable max lifetime for CLI device-flow
	// sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
	// gh-style flows). Companies that need finite session lifetime
	// (eg. SOC 2 / ISO 27001 customers with mandated session
	// expiration) can set a positive value. The /api/auth/cli/refresh
	// endpoint rejects refresh attempts where now - issued_at exceeds
	// this value, regardless of the underlying refresh-token TTL.
	// Spec: specs/ai-governance/sessions/session-ttl.feature.
	MaxSessionDurationDays int `gorm:"column:maxSessionDurationDays;size:32;default:0;not null" json:"maxSessionDurationDays"`
	// Every member of this organization can prove a second factor (D06). A
	// MEMBERSHIP CONDITION, not a session policy: it is asked when a member
	// reaches this organization's data, it holds the ones who cannot prove
	// one at an enrollment gate for THIS organization alone, and it ends no
	// session, ever. Turning it on signs nobody out.
	// Spec: specs/identity/mfa-and-session-shape.feature.
	MfaRequired bool `gorm:"column:mfaRequired;default:false;not null" json:"mfaRequired"`
	// Whether members may connect the documentation playground to their own access with a
	// short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
	// no key is minted and live ones stop authenticating. The pasted-key path is untouched.
	// Spec: modules/api-key/specs/docs-session-keys.feature.
	DocsPlaygroundSessionsEnabled *bool `gorm:"column:docsPlaygroundSessionsEnabled;default:true;not null" json:"docsPlaygroundSessionsEnabled"`
	// When the docs playground was last turned off, written in the same update as the switch. A
	// docs session key minted at or before it never authenticates again, nor any key minted under
	// one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
	DocsPlaygroundSessionsDisabledAt *DateTime `gorm:"column:docsPlaygroundSessionsDisabledAt;type:timestamp(3)" json:"docsPlaygroundSessionsDisabledAt"`
	// How many consecutive failed sign-ins lock an account (GAC-09).
	// 0 = never lock, which is the default and what every organization that
	// has not asked for this keeps. Counts EVERY credential together - a
	// password, a one-time code, a backup code, a passkey - because five per
	// factor hands an attacker one budget per factor.
	// Spec: specs/identity/org-account-lockout.feature.
	LockoutAfterFailedAttempts int `gorm:"column:lockoutAfterFailedAttempts;size:32;default:0;not null" json:"lockoutAfterFailedAttempts"`
	// How long a lock lasts, in minutes. Ignored while the threshold is 0.
	LockoutMinutes *int `gorm:"column:lockoutMinutes;size:32;default:30;not null" json:"lockoutMinutes"`
	// How long a BROWSER session may sit idle before it ends, in minutes
	// (GAC-10). 0 = no idle timeout. Activity rolls the window, which is the
	// "refreshed every hour" half of the control.
	//
	// A day for organizations created from now on; organizations that predate
	// the column keep 0 and are unchanged, because the migration adds the
	// column at zero and moves the default afterwards.
	//
	// NOT `maxSessionDurationDays`, which bounds CLI and device sessions and
	// is enforced when the CLI refreshes. Two surfaces, two lifetimes, two
	// decisions an administrator makes separately.
	// Spec: specs/identity/org-session-lifetime.feature.
	SessionIdleTimeoutMinutes *int `gorm:"column:sessionIdleTimeoutMinutes;size:32;default:1440;not null" json:"sessionIdleTimeoutMinutes"`
	// The ceiling on a browser session from the moment it was minted, in
	// minutes, which no amount of activity extends. 0 = no ceiling, the
	// default. Offered beside the idle timeout because a session kept alive
	// forever by a polling page is not really bounded.
	SessionMaxLifetimeMinutes int            `gorm:"column:sessionMaxLifetimeMinutes;size:32;default:0;not null" json:"sessionMaxLifetimeMinutes"`
	SignupData                datatypes.JSON `gorm:"column:signupData" json:"signupData"`
	SignedDPA                 bool           `gorm:"column:signedDPA;default:false;not null" json:"signedDPA"`
	ElasticsearchNodeURL      *string        `gorm:"column:elasticsearchNodeUrl" json:"elasticsearchNodeUrl"`
	ElasticsearchAPIKey       *string        `gorm:"column:elasticsearchApiKey" json:"elasticsearchApiKey"`
	UseCustomElasticsearch    bool           `gorm:"column:useCustomElasticsearch;default:false;not null" json:"useCustomElasticsearch"`
	S3Endpoint                *string        `gorm:"column:s3Endpoint" json:"s3Endpoint"`
	S3AccessKeyID             *string        `gorm:"column:s3AccessKeyId" json:"s3AccessKeyId"`
	S3SecretAccessKey         *string        `gorm:"column:s3SecretAccessKey" json:"s3SecretAccessKey"`
	S3Bucket                  *string        `gorm:"column:s3Bucket" json:"s3Bucket"`
	UseCustomS3               bool           `gorm:"column:useCustomS3;default:false;not null" json:"useCustomS3"`
	SentPlanLimitAlert        *DateTime      `gorm:"column:sentPlanLimitAlert;type:timestamp(3)" json:"sentPlanLimitAlert"`
	SsoDomain                 *string        `gorm:"column:ssoDomain;uniqueIndex:Organization_ssoDomain_key" json:"ssoDomain"`
	SsoProvider               *string        `gorm:"column:ssoProvider;index:Organization_ssoProvider_idx" json:"ssoProvider"`
	// How colleagues on a matching domain get in (D12, ADR-117):
	// `off` (invisible to them entirely) | `request` (they ask, an admin
	// approves) | `auto` (a verified colleague walks straight in).
	//
	// `request` is the default for a self-serve organization; `auto` is never
	// a default and never inferred - an administrator turns it on, names the
	// domain in `joinDomains` while doing it, and the licence gate holds that
	// value while letting `request` through.
	DomainJoin *string `gorm:"column:domainJoin;default:'request';not null" json:"domainJoin"`
	// The organisation role handed to a person admitted without an invitation:
	// a self-service domain join (`domainJoin`) or an SSO-admitted login.
	// `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
	// organisation settings page only; the join flow never asks, and an
	// invitation always names its own role.
	JoinerRole string `gorm:"column:joinerRole;default:'MEMBER';not null" json:"joinerRole"`
	// The domains an administrator NAMED when turning automatic joining on.
	// Empty means automatic joining admits nobody whatever `domainJoin` says -
	// which is what keeps a setting left on from admitting a domain nobody
	// ever agreed to.
	JoinDomains     ScalarList[string] `gorm:"column:joinDomains;type:text[]" json:"joinDomains"`
	PresenceEnabled *bool              `gorm:"column:presenceEnabled;default:true;not null" json:"presenceEnabled"`
	// Org-level trace-sharing kill switch. Off = trace share links cannot be
	// created in ANY project of this org, and existing links stop resolving —
	// the org toggle is the global switch, the per-project
	// `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
	// `presenceEnabled`). Effective sharing = org AND project. ADR-039.
	TraceSharingEnabled *bool `gorm:"column:traceSharingEnabled;default:true;not null" json:"traceSharingEnabled"`
	// Org-configurable contact surfaced to users in CLI "contact your
	// admin" messages and the in-app budget-exceeded banner. Free text:
	// accepts an email, a URL (internal ticketing system), or any other
	// instruction. When null, callers fall back to the first ADMIN
	// member's email (`resolveSupportContact` helper).
	SupportContact *string `gorm:"column:supportContact" json:"supportContact"`
	// Signup intent, editable later as the org "Primary use" setting.
	// AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
	// NULL means intent unset (every pre-fork org and any org created
	// outside onboarding) — the home resolver keeps legacy behavior.
	// ADR-038.
	PrimaryIntent          *string   `gorm:"column:primaryIntent" json:"primaryIntent"`
	PromoCode              *string   `gorm:"column:promoCode" json:"promoCode"`
	StripeCustomerID       *string   `gorm:"column:stripeCustomerId;uniqueIndex:Organization_stripeCustomerId_key" json:"stripeCustomerId"`
	Currency               string    `gorm:"column:currency;default:'EUR';not null" json:"currency"`
	PricingModel           string    `gorm:"column:pricingModel;default:'TIERED';not null" json:"pricingModel"`
	License                *string   `gorm:"column:license;type:text" json:"license"`
	LicenseExpiresAt       *DateTime `gorm:"column:licenseExpiresAt;type:timestamp(3)" json:"licenseExpiresAt"`
	LicenseLastValidatedAt *DateTime `gorm:"column:licenseLastValidatedAt;type:timestamp(3)" json:"licenseLastValidatedAt"`
	// A customer that runs LangWatch self-hosted. The hosted usage, budgets and
	// invoices of its connected installs attach to this organization (ADR-139).
	SelfHostedCustomer bool `gorm:"column:selfHostedCustomer;default:false;not null" json:"selfHostedCustomer"`
	// LangWatch-hosted services this organization has switched OFF, named the
	// way its license names them ("instant_evals"). A service the license is
	// entitled to is on unless it is listed here, so a customer who bought
	// hosted judging gets it without finding a settings page first, and an
	// administrator who switched one off keeps it off across license changes
	// (ADR-139, section 2). A license that names no service reaches nothing,
	// whatever this column holds.
	ConnectServicesDisabled ScalarList[string] `gorm:"column:connectServicesDisabled;type:text[]" json:"connectServicesDisabled"`
	// When a member of this organization switched Instant Evals on from the
	// search bar, which is the organization's consent to have judged text sent
	// to the judge's provider. Null until then: a self-serve organization is
	// never opted in on its behalf, and an enterprise one is switched on by an
	// operator through the release flag after asking, never here.
	InstantEvalsEnabledAt *DateTime `gorm:"column:instantEvalsEnabledAt;type:timestamp(3)" json:"instantEvalsEnabledAt"`
	// The member who switched them on, so the audit trail names a person and
	// not just a moment.
	InstantEvalsEnabledByUserID *string `gorm:"column:instantEvalsEnabledByUserId" json:"instantEvalsEnabledByUserId"`
	// When the last license sync succeeded.
	ConnectLastSyncAt *DateTime `gorm:"column:connectLastSyncAt;type:timestamp(3)" json:"connectLastSyncAt"`
	// The code of the failure the last sync ended on, cleared on the next
	// success. Settings, Connect renders the copy for it.
	ConnectLastSyncError *string                        `gorm:"column:connectLastSyncError" json:"connectLastSyncError"`
	Members              []OrganizationUser             `gorm:"foreignKey:OrganizationID;constraint:OrganizationUser_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Teams                []Team                         `gorm:"foreignKey:OrganizationID;constraint:Team_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	WebhookEndpoints     []WebhookEndpoint              `gorm:"foreignKey:OrganizationID;constraint:WebhookEndpoint_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationInvite   []OrganizationInvite           `gorm:"foreignKey:OrganizationID;constraint:OrganizationInvite_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CustomRoles          []CustomRole                   `gorm:"foreignKey:OrganizationID;constraint:CustomRole_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Notifications        []Notification                 `gorm:"foreignKey:OrganizationID;constraint:Notification_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subscriptions        []Subscription                 `gorm:"foreignKey:OrganizationID;constraint:Subscription_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ScimTokens           []ScimToken                    `gorm:"foreignKey:OrganizationID;constraint:ScimToken_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ScimRequestLog       []ScimRequestLog               `gorm:"foreignKey:OrganizationID;constraint:ScimRequestLog_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PromptTags           []PromptTag                    `gorm:"foreignKey:OrganizationID;constraint:PromptTag_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Groups               []Group                        `gorm:"foreignKey:OrganizationID;constraint:Group_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleBindings         []RoleBinding                  `gorm:"foreignKey:OrganizationID;constraint:RoleBinding_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeys              []APIKey                       `gorm:"foreignKey:OrganizationID;constraint:ApiKey_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	IssuedLicenses       []IssuedLicense                `gorm:"foreignKey:OrganizationID;constraint:IssuedLicense_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	GatewayChangeEvents  []GatewayChangeEvent           `gorm:"foreignKey:OrganizationID;constraint:GatewayChangeEvent_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoutingPolicies      []RoutingPolicy                `gorm:"foreignKey:OrganizationID;constraint:RoutingPolicy_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	IngestionSources     []IngestionSource              `gorm:"foreignKey:OrganizationID;constraint:IngestionSource_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AnomalyRules         []AnomalyRule                  `gorm:"foreignKey:OrganizationID;constraint:AnomalyRule_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AnomalyAlerts        []AnomalyAlert                 `gorm:"foreignKey:OrganizationID;constraint:AnomalyAlert_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AiToolEntries        []AiToolEntry                  `gorm:"foreignKey:OrganizationID;constraint:AiToolEntry_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	IngestionTemplates   []IngestionTemplate            `gorm:"foreignKey:OrganizationID;constraint:IngestionTemplate_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	GithubInstallations  []GithubInstallation           `gorm:"foreignKey:OrganizationID;constraint:GithubInstallation_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	GithubPullRequests   []GithubPullRequest            `gorm:"foreignKey:OrganizationID;constraint:GithubPullRequest_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	GithubBranchPrChecks []GithubBranchPullRequestCheck `gorm:"foreignKey:OrganizationID;constraint:GithubBranchPullRequestCheck_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConnectedBilling     *ConnectedBillingAccount       `gorm:"foreignKey:OrganizationID;constraint:ConnectedBillingAccount_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Organization) TableName() string {
	return "Organization"
}

func (m *Organization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
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

// The registry of licenses LangWatch has issued (ADR-139). It lives on
// LangWatch Cloud; a self-hosted install has the table and no rows.
//
// The row never holds the token an install presents. `tokenHash` is the
// SHA-256 of that token, so read access to this table is not credential
// access. Status is derived: revoked, superseded, past `expiresAt`, else active.
type IssuedLicense struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The `licenseId` inside the signed license (`lic-<uuid>`).
	LicenseID string `gorm:"column:licenseId;uniqueIndex:IssuedLicense_licenseId_key;not null" json:"licenseId"`
	// SHA-256 of the `lwl_` token. Looked up on every credential resolution.
	TokenHash string `gorm:"column:tokenHash;uniqueIndex:IssuedLicense_tokenHash_key;not null" json:"tokenHash"`
	// The customer organization on LangWatch Cloud. Null for a license that is
	// recorded but not linked to a customer yet, such as a self-serve purchase:
	// it resolves to nothing until an operator links it.
	OrganizationID *string `gorm:"column:organizationId;index:IssuedLicense_organizationId_idx" json:"organizationId"`
	// As signed into the license.
	OrganizationName string `gorm:"column:organizationName;not null" json:"organizationName"`
	Email            string `gorm:"column:email;not null" json:"email"`
	PlanType         string `gorm:"column:planType;not null" json:"planType"`
	MaxMembers       int    `gorm:"column:maxMembers;size:32;not null" json:"maxMembers"`
	// The lite seats sold: null when not stated or when unlimited. Never a stand-in number.
	MaxMembersLite *int `gorm:"column:maxMembersLite;size:32" json:"maxMembersLite"`
	// True when the license sells unlimited lite seats; maxMembersLite is then null.
	LiteSeatsUnlimited bool `gorm:"column:liteSeatsUnlimited;default:false;not null" json:"liteSeatsUnlimited"`
	// The Developer seats sold: null when not stated, which sells none.
	MaxMembersDeveloper *int     `gorm:"column:maxMembersDeveloper;size:32" json:"maxMembersDeveloper"`
	IssuedAt            DateTime `gorm:"column:issuedAt;type:timestamp(3);not null" json:"issuedAt"`
	ExpiresAt           DateTime `gorm:"column:expiresAt;index:IssuedLicense_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
	Source              string   `gorm:"column:source;not null" json:"source"`
	// The operator who issued or registered it. Null for the purchase flow and the script.
	IssuedByID    *string   `gorm:"column:issuedById" json:"issuedById"`
	RevokedAt     *DateTime `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	RevokedByID   *string   `gorm:"column:revokedById" json:"revokedById"`
	RevokedReason *string   `gorm:"column:revokedReason" json:"revokedReason"`
	// Set when the install first presents the license that replaced this one.
	SupersededAt *DateTime `gorm:"column:supersededAt;type:timestamp(3)" json:"supersededAt"`
	// The id of the IssuedLicense this one replaces, when it was reissued. A
	// plain column ref (relationMode="prisma" cannot declare a self-relation on
	// Postgres); the registry service validates it. Unique, so a license is
	// replaced by one license at most.
	ReplacesID *string `gorm:"column:replacesId;uniqueIndex:IssuedLicense_replacesId_key" json:"replacesId"`
	// A reissued license waiting to be delivered over sync, encrypted. Erased
	// when the install first presents it. The only license text the registry holds.
	PendingDeliveryLicense *string `gorm:"column:pendingDeliveryLicense;type:text" json:"pendingDeliveryLicense"`
	// Hosted services this license may call, such as `instant_evals`.
	Services ScalarList[string] `gorm:"column:services;type:text[]" json:"services"`
	// The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats.
	SeatRateCents *int    `gorm:"column:seatRateCents;size:32" json:"seatRateCents"`
	SeatCurrency  *string `gorm:"column:seatCurrency" json:"seatCurrency"`
	// The prepaid usage commit for the term, in USD cents.
	CommitUsdCents int `gorm:"column:commitUsdCents;size:32;default:0;not null" json:"commitUsdCents"`
	// Whether usage may continue past the commit, up to `overageMaxUsdCents`.
	OverageEnabled     bool `gorm:"column:overageEnabled;default:false;not null" json:"overageEnabled"`
	OverageMaxUsdCents *int `gorm:"column:overageMaxUsdCents;size:32" json:"overageMaxUsdCents"`
	// The install this license is bound to. Set on first use, cleared by an operator.
	InstanceID      *string   `gorm:"column:instanceId" json:"instanceId"`
	InstanceBoundAt *DateTime `gorm:"column:instanceBoundAt;type:timestamp(3)" json:"instanceBoundAt"`
	// The last license sync this install reported: when it arrived, the
	// LangWatch version that sent it, and the seats in use at that moment.
	// Null until the install syncs for the first time.
	LastSyncAt          *DateTime `gorm:"column:lastSyncAt;type:timestamp(3)" json:"lastSyncAt"`
	LastSyncVersion     *string   `gorm:"column:lastSyncVersion" json:"lastSyncVersion"`
	ReportedMembers     *int      `gorm:"column:reportedMembers;size:32" json:"reportedMembers"`
	ReportedMembersLite *int      `gorm:"column:reportedMembersLite;size:32" json:"reportedMembersLite"`
	// The managed gateway key (purpose CONNECT) that hosted-service traffic of
	// this license runs under. One per license, so spend rows name the install.
	// Created when the license first resolves. A plain column, not a relation.
	VirtualKeyID *string `gorm:"column:virtualKeyId;uniqueIndex:IssuedLicense_virtualKeyId_key" json:"virtualKeyId"`
	// The seats of the license this one replaced, set only when a seat change
	// raised a linked license. Billing reads it through licensing to invoice
	// the added seats, so the fact commits with the license or not at all.
	SeatsRaisedFrom *int          `gorm:"column:seatsRaisedFrom;index:IssuedLicense_seatsRaisedFrom_idx;size:32" json:"seatsRaisedFrom"`
	CreatedAt       DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization    *Organization `gorm:"belongsTo"`
}

func (IssuedLicense) TableName() string {
	return "IssuedLicense"
}

func (m *IssuedLicense) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*IssuedLicense) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The licence an organization activated and its dates, owned by licensing
// (round 37 D6). A row whose key is null is a cleared licence. An organization
// with no row is read from its Organization columns until the copy step runs.
type OrganizationLicense struct {
	OrganizationID string    `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	LicenseKey     *string   `gorm:"column:licenseKey;type:text" json:"licenseKey"`
	ExpiresAt      *DateTime `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	ValidatedAt    *DateTime `gorm:"column:validatedAt;type:timestamp(3)" json:"validatedAt"`
	UpdatedAt      DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (OrganizationLicense) TableName() string {
	return "OrganizationLicense"
}

func (m *OrganizationLicense) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganizationLicense) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A short code a fresh install pastes instead of a license blob (ADR-139,
// section 5). Redeeming it mints the license the code describes and returns
// it signed, which is what identifies a new customer without anybody emailing
// a license key around.
//
// The row holds a hash of the code, never the code. Every lookup hashes what
// the caller presented first, so the stored value is a step away from anything
// that could be presented: read access to this table is not credential access.
type ActivationCode struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// SHA-256 of the normalised code. Looked up on every redemption.
	CodeHash string `gorm:"column:codeHash;uniqueIndex:ActivationCode_codeHash_key;not null" json:"codeHash"`
	// The last four characters, so an operator tells two codes apart in a list
	// without the list holding either one.
	CodeHint string `gorm:"column:codeHint;not null" json:"codeHint"`
	// The customer this code belongs to, and the terms it mints.
	OrganizationID   string `gorm:"column:organizationId;index:ActivationCode_organizationId_idx;not null" json:"organizationId"`
	OrganizationName string `gorm:"column:organizationName;not null" json:"organizationName"`
	Email            string `gorm:"column:email;not null" json:"email"`
	PlanType         string `gorm:"column:planType;not null" json:"planType"`
	MaxMembers       int    `gorm:"column:maxMembers;size:32;not null" json:"maxMembers"`
	// The lite seats the minted license carries: null when unlimited. Never a stand-in number.
	MaxMembersLite *int `gorm:"column:maxMembersLite;size:32;default:0" json:"maxMembersLite"`
	// True when the minted license carries unlimited lite seats; maxMembersLite is then null.
	LiteSeatsUnlimited bool `gorm:"column:liteSeatsUnlimited;default:false;not null" json:"liteSeatsUnlimited"`
	// The Developer seats the minted license carries: null when not stated, which sells none.
	MaxMembersDeveloper *int `gorm:"column:maxMembersDeveloper;size:32" json:"maxMembersDeveloper"`
	// How long the license this code mints runs for.
	LicenseTermDays int `gorm:"column:licenseTermDays;size:32;not null" json:"licenseTermDays"`
	// Hosted services the minted license may call.
	Services ScalarList[string] `gorm:"column:services;type:text[]" json:"services"`
	// When the code stops working, which is not the term of what it mints.
	ExpiresAt DateTime `gorm:"column:expiresAt;index:ActivationCode_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
	// A code a customer may redeem on more than one install, for a rollout.
	Reusable bool `gorm:"column:reusable;default:false;not null" json:"reusable"`
	// Set by the redemption that won, which is what makes a single-use code
	// single use: the write itself carries the condition.
	RedeemedAt           *DateTime `gorm:"column:redeemedAt;type:timestamp(3)" json:"redeemedAt"`
	RedeemedByInstanceID *string   `gorm:"column:redeemedByInstanceId" json:"redeemedByInstanceId"`
	// The license the winning redemption minted.
	IssuedLicenseID *string `gorm:"column:issuedLicenseId" json:"issuedLicenseId"`
	// How many installs have redeemed it, which only moves on a reusable code.
	RedemptionCount int       `gorm:"column:redemptionCount;size:32;default:0;not null" json:"redemptionCount"`
	RevokedAt       *DateTime `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	RevokedByID     *string   `gorm:"column:revokedById" json:"revokedById"`
	CreatedByID     *string   `gorm:"column:createdById" json:"createdById"`
	CreatedAt       DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ActivationCode) TableName() string {
	return "ActivationCode"
}

func (m *ActivationCode) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ActivationCode) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// How a connected customer is invoiced (ADR-139, section 7): one payment
// provider customer, one quarterly usage subscription in USD, and the
// commercial terms the annual invoice was raised from.
//
// One row per organization. Every write through it is addressed by a stored
// provider id, so running onboarding or a renewal twice changes nothing.
type ConnectedBillingAccount struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:ConnectedBillingAccount_organizationId_key;not null" json:"organizationId"`
	// The customer at the payment provider. Invoiced, never charged.
	StripeCustomerID string `gorm:"column:stripeCustomerId;uniqueIndex:ConnectedBillingAccount_stripeCustomerId_key;not null" json:"stripeCustomerId"`
	// The one metered subscription, and the item hosted usage is reported on.
	// Null between the customer being created and the subscription existing,
	// which is where a half-finished onboarding resumes from.
	UsageSubscriptionID     *string  `gorm:"column:usageSubscriptionId" json:"usageSubscriptionId"`
	UsageSubscriptionItemID *string  `gorm:"column:usageSubscriptionItemId" json:"usageSubscriptionItemId"`
	TermStartsAt            DateTime `gorm:"column:termStartsAt;type:timestamp(3);not null" json:"termStartsAt"`
	TermEndsAt              DateTime `gorm:"column:termEndsAt;type:timestamp(3);not null" json:"termEndsAt"`
	// Prepaid usage for the term, in USD cents. The credit grants add up to it.
	CommitUsdCents int `gorm:"column:commitUsdCents;size:32;default:0;not null" json:"commitUsdCents"`
	// The seat contract, which may be in another currency than hosted usage.
	SeatCurrency  string `gorm:"column:seatCurrency;not null" json:"seatCurrency"`
	SeatRateCents int    `gorm:"column:seatRateCents;size:32;not null" json:"seatRateCents"`
	Seats         int    `gorm:"column:seats;size:32;not null" json:"seats"`
	// `us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
	// virtual bank account at the provider. Null when it wires to LangWatch's
	// own account and finance marks the invoice paid.
	BankTransferType    *string `gorm:"column:bankTransferType" json:"bankTransferType"`
	BankTransferCountry *string `gorm:"column:bankTransferCountry" json:"bankTransferCountry"`
	BillingEmail        string  `gorm:"column:billingEmail;not null" json:"billingEmail"`
	// A renewal whose credit is waiting for the last usage invoice of the old
	// term to be finalized: the new term's end and commit, and who asked.
	PendingRenewal datatypes.JSON `gorm:"column:pendingRenewal" json:"pendingRenewal"`
	CreatedByID    *string        `gorm:"column:createdById" json:"createdById"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization  `gorm:"belongsTo"`
}

func (ConnectedBillingAccount) TableName() string {
	return "ConnectedBillingAccount"
}

func (m *ConnectedBillingAccount) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ConnectedBillingAccount) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A paid credit at the payment provider, which only metered usage draws down.
// Kept here so a grant is created once per reason and the backoffice can show
// what the customer holds without asking the provider.
type ConnectedCreditGrant struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The `ConnectedBillingAccount` row. A plain column ref, like
	// `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
	AccountID string `gorm:"column:accountId;index:ConnectedCreditGrant_accountId_idx;not null" json:"accountId"`
	// The grant at the payment provider. Unique, so a retry records one row.
	StripeCreditGrantID string `gorm:"column:stripeCreditGrantId;uniqueIndex:ConnectedCreditGrant_stripeCreditGrantId_key;not null" json:"stripeCreditGrantId"`
	AmountUsdCents      int    `gorm:"column:amountUsdCents;size:32;not null" json:"amountUsdCents"`
	Kind                string `gorm:"column:kind;not null" json:"kind"`
	// The term this credit was granted for, and when the provider lets it lapse:
	// after the last invoice of that term has been finalized, never on the day
	// the term ends.
	TermEndsAt DateTime `gorm:"column:termEndsAt;type:timestamp(3);not null" json:"termEndsAt"`
	ExpiresAt  DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (ConnectedCreditGrant) TableName() string {
	return "ConnectedCreditGrant"
}

func (m *ConnectedCreditGrant) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// An invoice LangWatch raised for a connected customer, as the backoffice
// reads it.
type ConnectedInvoice struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The `ConnectedBillingAccount` row, as a plain column ref.
	AccountID       string `gorm:"column:accountId;index:ConnectedInvoice_accountId_idx;not null" json:"accountId"`
	StripeInvoiceID string `gorm:"column:stripeInvoiceId;uniqueIndex:ConnectedInvoice_stripeInvoiceId_key;not null" json:"stripeInvoiceId"`
	Kind            string `gorm:"column:kind;not null" json:"kind"`
	Currency        string `gorm:"column:currency;not null" json:"currency"`
	AmountCents     int    `gorm:"column:amountCents;size:32;not null" json:"amountCents"`
	// The provider's own status: `draft`, `open`, `paid`, `void` or `uncollectible`.
	Status          string    `gorm:"column:status;not null" json:"status"`
	PaidOutOfBandAt *DateTime `gorm:"column:paidOutOfBandAt;type:timestamp(3)" json:"paidOutOfBandAt"`
	TermStartsAt    *DateTime `gorm:"column:termStartsAt;type:timestamp(3)" json:"termStartsAt"`
	CreatedAt       DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ConnectedInvoice) TableName() string {
	return "ConnectedInvoice"
}

func (m *ConnectedInvoice) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ConnectedInvoice) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Every SCIM request we could attribute, and what we answered (ADR-126).
//
// EVIDENCE, NOT TRUTH. A request authors nothing — an HTTP round trip that
// was refused decided nothing — so this is a table and not an event: nothing
// downstream derives from it, rebuilding the world from the log must not
// depend on how often a provider retried a GET, and rows age out of a
// retention window an event log would never have.
//
// NO ROW FOR A REQUEST WE CANNOT ATTRIBUTE. A token that does not verify is
// answered 401 and recorded nowhere: we do not know whose organization to
// file it under, and a table unauthenticated traffic can write is a table
// anybody on the internet can fill. What answers that case instead is
// `ScimToken.lastUsedAt` being null. The 403 IS recorded — a lapsed plan is
// a credential we recognize.
//
// THE PAYLOAD RULE HOLDS. A method, a resource name, a status, a reason slug
// and our own sentence. Never a token, never a header, never a provider's
// raw message — the same reason `scim_apply_failed` carries a code and not a
// message.
type ScimRequestLog struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:ScimRequestLog_organizationId_occurredAt_idx,priority:1;not null" json:"organizationId"`
	// Null when the refusal resolved an organization and stopped before the
	// connection — a lapsed plan is answered off the token alone.
	ConnectionID *string `gorm:"column:connectionId;index:ScimRequestLog_connectionId_occurredAt_idx,priority:1" json:"connectionId"`
	Method       string  `gorm:"column:method;not null" json:"method"`
	// The SCIM resource asked for ("Users", "Users/:id", "Groups"), never the
	// raw path: a path carries ids and query strings, and this is read by
	// people rather than matched by machines.
	Resource string `gorm:"column:resource;not null" json:"resource"`
	Status   int    `gorm:"column:status;size:32;not null" json:"status"`
	// A stable slug, only on a refusal, so a reader branches on this and never
	// on the sentence beside it.
	Reason *string `gorm:"column:reason" json:"reason"`
	// Our own short sentence, customer-safe.
	Detail       *string       `gorm:"column:detail" json:"detail"`
	OccurredAt   DateTime      `gorm:"column:occurredAt;index:ScimRequestLog_organizationId_occurredAt_idx,priority:2;index:ScimRequestLog_connectionId_occurredAt_idx,priority:2;index:ScimRequestLog_occurredAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"occurredAt"`
	Organization *Organization `gorm:"belongsTo"`
}

func (ScimRequestLog) TableName() string {
	return "ScimRequestLog"
}

func (m *ScimRequestLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.OccurredAt.IsZero() {
		m.OccurredAt = DateTime{Time: now}
	}
	return nil
}

type ScimToken struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:ScimToken_organizationId_idx;not null" json:"organizationId"`
	// D08: the SSO connection this token was issued for, and its ENTIRE write
	// authority — a token minted for one connection can never touch another
	// connection's people. No relation, matching `SsoConnection` itself: that
	// projection carries no foreign keys, so neither may a row pointing at it.
	//
	// Nullable only for the tokens that predate the column. They were minted
	// when "which organization" was the whole answer, so they keep the
	// organization-wide authority they were sold with; the backfill points
	// each at its organization's first connection where there is one. Minting
	// a NEW token without a connection is refused (`scim_connection_required`).
	ConnectionID *string `gorm:"column:connectionId;index:ScimToken_connectionId_idx" json:"connectionId"`
	// UNIQUE, and that is a tenancy boundary rather than tidiness.
	//
	// The lookup that turns a bearer token into an organization is keyed on
	// this column alone — it has to be, since the token is the only thing the
	// caller presents. While every token was 32 random bytes a collision was
	// impossible and the missing constraint cost nothing. An administrator may
	// now supply a value they already hold, so two organizations CAN choose the
	// same string, and without this the lookup would answer with whichever row
	// the planner reached first: one customer's directory provisioning and
	// deleting another customer's people.
	HashedToken string `gorm:"column:hashedToken;uniqueIndex:ScimToken_hashedToken_key;not null" json:"hashedToken"`
	// How `hashedToken` was derived, so the two schemes can coexist while
	// tokens minted under the old one are still in use.
	//
	// `sha256` is a bare digest, which is correct for 32 random bytes and wrong
	// for a value a person chose — a database dump of those is a wordlist away
	// from live credentials. `hmac-sha256` is keyed on the deployment's own
	// secret, so a dump on its own is inert. Every new row is the latter.
	HashScheme   *string       `gorm:"column:hashScheme;default:'sha256';not null" json:"hashScheme"`
	Description  *string       `gorm:"column:description" json:"description"`
	CreatedAt    DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	LastUsedAt   *DateTime     `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	Organization *Organization `gorm:"belongsTo"`
}

func (ScimToken) TableName() string {
	return "ScimToken"
}

func (m *ScimToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// D08 — who the directory means, per connection.
//
// The identity provider's own identifier for a person, which is what
// survives their email changing: an address is not identity, people change
// theirs, and the same person carries different identifiers on two different
// connections. So the PAIR is the key and `externalId` alone never is.
//
// No relation to `SsoConnection` (that projection carries no foreign keys);
// the `User` relation is real, so a deleted user takes their directory
// identities with them rather than stranding a mapping that resolves to
// nobody.
type ScimExternalID struct {
	OrganizationID string   `gorm:"column:organizationId;index:ScimExternalId_organizationId_userId_idx,priority:1;not null" json:"organizationId"`
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	ConnectionID   string   `gorm:"column:connectionId;uniqueIndex:ScimExternalId_connectionId_externalId_key,priority:1;index:ScimExternalId_connectionId_idx;not null" json:"connectionId"`
	ExternalID     string   `gorm:"column:externalId;uniqueIndex:ScimExternalId_connectionId_externalId_key,priority:2;not null" json:"externalId"`
	UserID         string   `gorm:"column:userId;index:ScimExternalId_organizationId_userId_idx,priority:2;index:ScimExternalId_userId_idx;not null" json:"userId"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User           *User    `gorm:"belongsTo"`
}

func (ScimExternalID) TableName() string {
	return "ScimExternalId"
}

func (m *ScimExternalID) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ScimExternalID) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Directory profile and lifecycle are tenant facts, independent of the shared account.
type ScimUserResource struct {
	OrganizationID string    `gorm:"column:organizationId;primaryKey;index:ScimUserResource_organizationId_userName_idx,priority:1" json:"organizationId"`
	UserID         string    `gorm:"column:userId;primaryKey;index:ScimUserResource_userId_idx" json:"userId"`
	UserName       string    `gorm:"column:userName;index:ScimUserResource_organizationId_userName_idx,priority:2;not null" json:"userName"`
	Name           *string   `gorm:"column:name" json:"name"`
	Active         *bool     `gorm:"column:active;default:true;not null" json:"active"`
	DeletedAt      *DateTime `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	CreatedAt      DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User           *User     `gorm:"belongsTo"`
}

func (ScimUserResource) TableName() string {
	return "ScimUserResource"
}

func (m *ScimUserResource) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ScimUserResource) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A connection's managed users, including resources without an externalId.
// Scalar keys only; readers exclude users that no longer exist.
type ScimDirectoryUser struct {
	OrganizationID string `gorm:"column:organizationId;index:ScimDirectoryUser_organizationId_userId_idx,priority:1;not null" json:"organizationId"`
	ConnectionID   string `gorm:"column:connectionId;primaryKey" json:"connectionId"`
	UserID         string `gorm:"column:userId;primaryKey;index:ScimDirectoryUser_organizationId_userId_idx,priority:2" json:"userId"`
}

func (ScimDirectoryUser) TableName() string {
	return "ScimDirectoryUser"
}

// D08 — one connection's directory sync, as the `scim-sync` pipeline's fold
// writes it.
//
// Pure event truth: every column is fold-written, rows are never deleted,
// and REVOKED is a tombstone rather than an absence — a dead letter that
// vanished when a connection was torn down would be a removal nobody could
// check afterwards.
//
// No relation to Organization or SsoConnection, matching `SsoConnection`
// itself: a projection is rebuilt from events, and a foreign key would make
// a replay depend on relational state the log does not carry.
type ScimSyncState struct {
	// The aggregate: the connection's id, which IS its sync's id.
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	ConnectionID   string `gorm:"column:connectionId;index:ScimSyncState_connectionId_idx;not null" json:"connectionId"`
	OrganizationID string `gorm:"column:organizationId;index:ScimSyncState_organizationId_idx;not null" json:"organizationId"`
	// TOKEN_ISSUED | SYNCING | ERROR | REVOKED
	State        string    `gorm:"column:state;index:ScimSyncState_state_idx;not null" json:"state"`
	LastPushedAt *DateTime `gorm:"column:lastPushedAt;type:timestamp(3)" json:"lastPushedAt"`
	// The failure standing right now: { op, errorCode, attempts, retiredAtMs,
	// userId, occurredAtMs }. A reason CODE and ids — never a token, never a
	// secret, never an internal hostname.
	LastFailure datatypes.JSON `gorm:"column:lastFailure" json:"lastFailure"`
	// Every failure retired without ever being applied, newest last. The dead
	// letters an administrator has to see; retiring never drops one.
	DeadLetters datatypes.JSON `gorm:"column:deadLetters;default:'[]';not null" json:"deadLetters"`
	// revoke | teardown, once the sync has ended.
	RevokedCause      *string  `gorm:"column:revokedCause" json:"revokedCause"`
	OccurredAt        DateTime `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	LastEventID       string   `gorm:"column:lastEventId;not null" json:"lastEventId"`
	AcceptedAt        DateTime `gorm:"column:acceptedAt;type:timestamp(3);not null" json:"acceptedAt"`
	ProjectionVersion string   `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	// Event-derived, not database-managed: `@default(now())` / `@updatedAt`
	// would make a replayed row differ from the row it rebuilds, which is
	// exactly the parity this projection has to keep.
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ScimSyncState) TableName() string {
	return "ScimSyncState"
}

// SCIM's own copy of an organization's SSO connections: a peer fold over
// identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
// reducer, one row per connection. SCIM reads this instead of asking identity.
// Rebuilt from identity's event log by a projection replay, so no relation:
// a foreign key would make a replay depend on state the log does not carry.
type ScimSsoConnectionView struct {
	// The aggregate: identity's connection id.
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:ScimSsoConnectionView_organizationId_createdAt_idx,priority:1;not null" json:"organizationId"`
	// The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt).
	Folded datatypes.JSON `gorm:"column:folded;not null" json:"folded"`
	// The event ids of the last delivery folded in, so a redelivery folds once.
	AppliedEventIds ScalarList[string] `gorm:"column:appliedEventIds;type:text[]" json:"appliedEventIds"`
	// A row of another version reads as undecodable and is re-folded from the log.
	ProjectionVersion string `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	// Event-derived, so a replayed row equals the row it rebuilds.
	CreatedAt DateTime `gorm:"column:createdAt;index:ScimSsoConnectionView_organizationId_createdAt_idx,priority:2;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ScimSsoConnectionView) TableName() string {
	return "ScimSsoConnectionView"
}

// Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
// A read model, one row per organization: rebuildable by replaying the event log, so no relation.
// Epochs in milliseconds.
// Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
type OrganizationSeatPeak struct {
	// The organization, which is also the aggregate id.
	OrganizationID string `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	// When the opening snapshot was taken, epoch milliseconds.
	OpenedAt       float64 `gorm:"column:openedAt;type:double precision;not null" json:"openedAt"`
	FullSeats      int     `gorm:"column:fullSeats;size:32;not null" json:"fullSeats"`
	LiteSeats      int     `gorm:"column:liteSeats;size:32;not null" json:"liteSeats"`
	DeveloperSeats int     `gorm:"column:developerSeats;size:32;not null" json:"developerSeats"`
	// The latest change time applied to any holder; invitations expire up to it.
	LatestChangeAt float64 `gorm:"column:latestChangeAt;type:double precision;not null" json:"latestChangeAt"`
	// Changes whose `from` disagreed with the seat the fold held.
	Mismatches int `gorm:"column:mismatches;size:32;not null" json:"mismatches"`
}

func (OrganizationSeatPeak) TableName() string {
	return "OrganizationSeatPeak"
}

// One member's or pending invitation's seat as the seat peak last applied it (ADR-002).
type OrganizationSeatPeakHolder struct {
	OrganizationID string `gorm:"column:organizationId;primaryKey;index:OrganizationSeatPeakHolder_organizationId_holderKind_seat_e_idx,priority:1;index:OrganizationSeatPeakHolder_organizationId_suppressedBy_idx,priority:1" json:"organizationId"`
	// `member`, `invitation` or `hold` (a seat held after somebody left).
	HolderKind string `gorm:"column:holderKind;primaryKey;index:OrganizationSeatPeakHolder_organizationId_holderKind_seat_e_idx,priority:2" json:"holderKind"`
	// The member's or the hold's userId, or the invitation's id.
	HolderID string `gorm:"column:holderId;primaryKey" json:"holderId"`
	// `full`, `lite`, `developer` or `none`.
	Seat          string  `gorm:"column:seat;index:OrganizationSeatPeakHolder_organizationId_holderKind_seat_e_idx,priority:3;not null" json:"seat"`
	LastChangedAt float64 `gorm:"column:lastChangedAt;type:double precision;not null" json:"lastChangedAt"`
	// Absent for a seat the opening snapshot recorded.
	LastChangeID *string `gorm:"column:lastChangeId" json:"lastChangeId"`
	// An invitation's expiry; null for a member or an invitation that never expires.
	ExpiresAt *float64 `gorm:"column:expiresAt;index:OrganizationSeatPeakHolder_organizationId_holderKind_seat_e_idx,priority:4;type:double precision" json:"expiresAt"`
	// The fold gave the invitation's seat back at `expiresAt`.
	ExpiredByFold bool `gorm:"column:expiredByFold;default:false;not null" json:"expiredByFold"`
	// A hold whose seat a live full invitation to its address took: that invitation's id.
	SuppressedBy *string `gorm:"column:suppressedBy;index:OrganizationSeatPeakHolder_organizationId_suppressedBy_idx,priority:2" json:"suppressedBy"`
	// The fold gave a suppressed hold its seat back when that invitation expired.
	RestoredByFold bool `gorm:"column:restoredByFold;default:false;not null" json:"restoredByFold"`
}

func (OrganizationSeatPeakHolder) TableName() string {
	return "OrganizationSeatPeakHolder"
}

// One calendar month's (UTC) seat peaks for an organization (ADR-002).
type OrganizationSeatPeakMonth struct {
	OrganizationID string `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	// `YYYY-MM`.
	Month      string  `gorm:"column:month;primaryKey" json:"month"`
	FullPeak   int     `gorm:"column:fullPeak;size:32;not null" json:"fullPeak"`
	FullPeakAt float64 `gorm:"column:fullPeakAt;type:double precision;not null" json:"fullPeakAt"`
	LitePeak   int     `gorm:"column:litePeak;size:32;not null" json:"litePeak"`
	LitePeakAt float64 `gorm:"column:litePeakAt;type:double precision;not null" json:"litePeakAt"`
	FullAtEnd  int     `gorm:"column:fullAtEnd;size:32;not null" json:"fullAtEnd"`
	LiteAtEnd  int     `gorm:"column:liteAtEnd;size:32;not null" json:"liteAtEnd"`
}

func (OrganizationSeatPeakMonth) TableName() string {
	return "OrganizationSeatPeakMonth"
}

type Project struct {
	ID     string `gorm:"column:id;primaryKey" json:"id"`
	Name   string `gorm:"column:name;not null" json:"name"`
	Slug   string `gorm:"column:slug;uniqueIndex:Project_slug_key;not null" json:"slug"`
	APIKey string `gorm:"column:apiKey;uniqueIndex:Project_apiKey_key;not null" json:"apiKey"`
	// When the legacy apiKey last authenticated a request, written at most once a
	// minute per project and process. Spec: modules/api-key/specs/request-credential-type.feature
	LegacyAPIKeyLastUsedAt *DateTime `gorm:"column:legacyApiKeyLastUsedAt;type:timestamp(3)" json:"legacyApiKeyLastUsedAt"`
	// Random per-project secret whose SHA-256 digest is this project's LangWatchQL
	// SQL tenant capability in ClickHouse. Deliberately NOT the API key: the
	// credential a caller authenticates with and the capability the database
	// resolves a tenant from stay decoupled, so rotating one never invalidates
	// the other, and a leak of the ClickHouse key map's input discloses nothing
	// a caller could authenticate with. Generated by the database on insert.
	LwqlKey   string `gorm:"column:lwqlKey;uniqueIndex:Project_lwqlKey_key;default:(gen_random_uuid())::text;not null" json:"lwqlKey"`
	TeamID    string `gorm:"column:teamId;index:Project_teamId_idx;index:Project_teamId_kind_idx,priority:1;not null" json:"teamId"`
	Language  string `gorm:"column:language;not null" json:"language"`
	Framework string `gorm:"column:framework;not null" json:"framework"`
	// Project kind:
	// - "application": user-facing workspace where developers send their
	// own app traces. The default for everything created via the
	// onboarding/quickstart flow.
	// - "internal_governance": hidden per-org routing/tenancy artifact
	// for IngestionSource ingest data. Auto-created lazily on the
	// first IngestionSource mint via ensureHiddenGovernanceProject().
	// MUST be filtered out of every user-visible Project surface
	// (project picker, /api/v1/projects, billing exports, RBAC
	// pickers, etc.) - the regression-test invariant in
	// specs/ai-gateway/governance/ui-contract.feature codifies this.
	// - "aggregate": a project that owns no traces and reads its member
	// projects through shared project-reader grants (ADR-177). Hidden
	// from every "send traces here" picker and from ingestion; visible in
	// the project switcher to organisation admins only.
	// Stored as a free-form string (matching IngestionSource.sourceType
	// pattern) so future routing kinds can ship without schema migrations.
	Kind *string `gorm:"column:kind;index:Project_teamId_kind_idx,priority:2;default:'application';not null" json:"kind"`
	// The scope rule an aggregate project materialises into grants. Null on
	// every other kind; decided in ADR-177. Shapes: { kind: "all-personal" } |
	// { kind: "personal-by-department", departmentId } |
	// { kind: "explicit", projectIds: string[] }.
	AggregateRule       datatypes.JSON `gorm:"column:aggregateRule" json:"aggregateRule"`
	FirstMessage        bool           `gorm:"column:firstMessage;default:false;not null" json:"firstMessage"`
	Integrated          bool           `gorm:"column:integrated;default:false;not null" json:"integrated"`
	CreatedAt           DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	UserLinkTemplate    *string        `gorm:"column:userLinkTemplate" json:"userLinkTemplate"`
	TraceSharingEnabled *bool          `gorm:"column:traceSharingEnabled;default:true;not null" json:"traceSharingEnabled"`
	PresenceEnabled     *bool          `gorm:"column:presenceEnabled;default:true;not null" json:"presenceEnabled"`
	S3Endpoint          *string        `gorm:"column:s3Endpoint" json:"s3Endpoint"`
	S3AccessKeyID       *string        `gorm:"column:s3AccessKeyId" json:"s3AccessKeyId"`
	S3SecretAccessKey   *string        `gorm:"column:s3SecretAccessKey" json:"s3SecretAccessKey"`
	S3Bucket            *string        `gorm:"column:s3Bucket" json:"s3Bucket"`
	ArchivedAt          *DateTime      `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	IsPersonal          bool           `gorm:"column:isPersonal;index:Project_isPersonal_idx;default:false;not null" json:"isPersonal"`
	OwnerUserID         *string        `gorm:"column:ownerUserId;index:Project_ownerUserId_idx;type:varchar(255)" json:"ownerUserId"`
	// Per-feature unlock flags for personal projects only. Shape:
	// `{ evaluations: bool, datasets: bool, annotations: bool, automations: bool }`.
	// Empty `{}` (default) reads as all-false at the service layer -
	// existing rows pre-migration require no backfill. The bundle is
	// a UI/nav predicate, NOT an auth gate (the underlying tRPC
	// routers stay open even when the bundle is off; admin tooling +
	// offline migration scripts can still operate on the data).
	// Team / org-scope projects are not feature-gated - this column
	// is meaningful only when `isPersonal=true`.
	PersonalFeatures datatypes.JSON `gorm:"column:personalFeatures;default:'{}';not null" json:"personalFeatures"`
	// Department assignment for spend attribution (the project / agent
	// lens, used when a trace has no principal user). Accounting-only.
	DepartmentID *string `gorm:"column:departmentId;index:Project_departmentId_idx" json:"departmentId"`
	// Per-project Langy egress allow-list (ADR-076). A nullable JSON array of
	// host patterns (`["registry.npmjs.org", "*.internal.acme.com"]`) that the
	// project's Langy workers may reach outbound. The *presence* of the list is
	// the enforcement mode: null/absent ⇒ monitor-only (watch, never block);
	// non-empty ⇒ the worker's egress adapter restricts outbound to
	// floor ∪ this list. Enforced in the agent pod's egress adapter — NOT the
	// gateway — so it lives on the project, not the VirtualKey. Resolved by
	// LangyCredentialService.getEgressAllowlist and threaded through the
	// per-request credentials envelope. See specs/langy/langy-egress-enforcement.feature.
	LangyEgressAllowlist datatypes.JSON `gorm:"column:langyEgressAllowlist" json:"langyEgressAllowlist"`
	// When a coding-agent session was last folded for this project, and when a
	// pull request was last mapped for a branch one of those sessions ran on.
	// Both are instrumentation signals, not counters: the project sidebar shows
	// its Sessions and Pull requests destinations only while the matching column
	// is inside the recency window, so a project that never sends coding-agent
	// telemetry never grows the two links, and one that stops sending loses
	// them again. Written through the throttled touch methods on
	// ProjectService, which skip the write while the stored value is still
	// recent, so a busy project costs at most one update an hour per column.
	// No index: every read is by project id, which is the primary key.
	// See specs/coding-agent/project-menu-links.feature.
	LastCodingAgentSessionAt     *DateTime                 `gorm:"column:lastCodingAgentSessionAt;type:timestamp(3)" json:"lastCodingAgentSessionAt"`
	LastCodingAgentPullRequestAt *DateTime                 `gorm:"column:lastCodingAgentPullRequestAt;type:timestamp(3)" json:"lastCodingAgentPullRequestAt"`
	Team                         *Team                     `gorm:"belongsTo"`
	OwnerUser                    *User                     `gorm:"foreignKey:OwnerUserID;belongsTo"`
	Checks                       []Monitor                 `gorm:"foreignKey:ProjectID;constraint:Monitor_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Costs                        []Cost                    `gorm:"foreignKey:ProjectID;constraint:Cost_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Topics                       []Topic                   `gorm:"foreignKey:ProjectID;constraint:Topic_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Datasets                     []Dataset                 `gorm:"foreignKey:ProjectID;constraint:Dataset_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DatasetRecords               []DatasetRecord           `gorm:"foreignKey:ProjectID;constraint:DatasetRecord_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CustomGraphs                 []CustomGraph             `gorm:"foreignKey:ProjectID;constraint:CustomGraph_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Dashboards                   []Dashboard               `gorm:"foreignKey:ProjectID;constraint:Dashboard_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	BatchEvaluations             []BatchEvaluation         `gorm:"foreignKey:ProjectID;constraint:BatchEvaluation_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Triggers                     []Trigger                 `gorm:"foreignKey:ProjectID;constraint:Trigger_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Experiments                  []Experiment              `gorm:"foreignKey:ProjectID;constraint:Experiment_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Annotations                  []Annotation              `gorm:"foreignKey:ProjectID;constraint:Annotation_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ProjectSecrets               []ProjectSecret           `gorm:"foreignKey:ProjectID;constraint:ProjectSecret_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	TriggerSent                  []TriggerSent             `gorm:"foreignKey:ProjectID;constraint:TriggerSent_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	WebhookDeliveries            []WebhookEndpointDelivery `gorm:"foreignKey:ProjectID;constraint:WebhookEndpointDelivery_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AnnotationScores             []AnnotationScore         `gorm:"foreignKey:ProjectID;constraint:AnnotationScore_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ShareLinks                   []ShareLink               `gorm:"foreignKey:ProjectID;constraint:ShareLink_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Workflows                    []Workflow                `gorm:"foreignKey:ProjectID;constraint:Workflow_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	WorkflowVersion              []WorkflowVersion         `gorm:"foreignKey:ProjectID;constraint:WorkflowVersion_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AnnotationQueue              []AnnotationQueue         `gorm:"foreignKey:ProjectID;constraint:AnnotationQueue_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AnnotationQueueItem          []AnnotationQueueItem     `gorm:"foreignKey:ProjectID;constraint:AnnotationQueueItem_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	LlmPromptConfigs             []LlmPromptConfig         `gorm:"foreignKey:ProjectID;constraint:LlmPromptConfig_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Analytics                    []Analytics               `gorm:"foreignKey:ProjectID;constraint:Analytics_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Notifications                []Notification            `gorm:"foreignKey:ProjectID;constraint:Notification_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Agents                       []Agent                   `gorm:"foreignKey:ProjectID;constraint:Agent_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Evaluators                   []Evaluator               `gorm:"foreignKey:ProjectID;constraint:Evaluator_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Scenarios                    []Scenario                `gorm:"foreignKey:ProjectID;constraint:Scenario_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SuiteConfigurations          []SimulationSuite         `gorm:"foreignKey:ProjectID;constraint:SimulationSuite_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SavedViews                   []SavedView               `gorm:"foreignKey:ProjectID;constraint:SavedView_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	GatewayGuardrails            []GatewayGuardrail        `gorm:"foreignKey:ProjectID;constraint:GatewayGuardrail_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	GatewayChangeEvents          []GatewayChangeEvent      `gorm:"foreignKey:ProjectID;constraint:GatewayChangeEvent_projectId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PinnedTraces                 []PinnedTrace             `gorm:"foreignKey:ProjectID;constraint:PinnedTrace_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	TraceEditOverlays            []TraceEditOverlay        `gorm:"foreignKey:ProjectID;constraint:TraceEditOverlay_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	EmailSuppressions            []EmailSuppression        `gorm:"foreignKey:ProjectID;constraint:EmailSuppression_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Project) TableName() string {
	return "Project"
}

func (m *Project) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
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

// A department is a single-valued accounting dimension, separate from
// RBAC, for rolling up AI spend by org unit (Engineering, Marketing,
// ...). Assigned to OrganizationUser, Team, and Project via their
// `departmentId` column. Pure accounting: never grants or restricts
// access. See specs/ai-gateway/governance/departments.feature.
type Department struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string    `gorm:"column:organizationId;index:Department_organizationId_idx;not null" json:"organizationId"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	CreatedAt      DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	ArchivedAt     *DateTime `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
}

func (Department) TableName() string {
	return "Department"
}

func (m *Department) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Department) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One dated person → department link (ADR-128 §13, #7882).
//
// `OrganizationUser.departmentId` stays the CURRENT assignment every screen
// reads; these rows are its history, appended by the same `assignUser` write
// so the two cannot drift. They exist for the cost reads: January's spend
// stays with January's department, so a read resolves against the link that
// was open on the cost's day, not against today's pointer. Members with no
// link on a day resolve to "unassigned", never to an error.
//
// Same shape as IdentityMatch's dated links, same one-open-link guard: the
// partial unique index in the migration rejects a second open row per
// (organizationId, userId) with SQLSTATE 23505.
type DepartmentMembershipHistory struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string   `gorm:"column:organizationId;index:DepartmentMembershipHistory_organizationId_userId_validFrom_idx,priority:1;not null" json:"organizationId"`
	UserID         string   `gorm:"column:userId;index:DepartmentMembershipHistory_organizationId_userId_validFrom_idx,priority:2;not null" json:"userId"`
	DepartmentID   string   `gorm:"column:departmentId;not null" json:"departmentId"`
	ValidFrom      DateTime `gorm:"column:validFrom;index:DepartmentMembershipHistory_organizationId_userId_validFrom_idx,priority:3;type:timestamp(3);not null" json:"validFrom"`
	// Null means the link is open. Reassignment and clearing close it.
	ValidTo   *DateTime `gorm:"column:validTo;type:timestamp(3)" json:"validTo"`
	CreatedAt DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (DepartmentMembershipHistory) TableName() string {
	return "DepartmentMembershipHistory"
}

func (m *DepartmentMembershipHistory) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Per-(organization, tool) override for which `langwatch <tool>` paths the
// CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
// `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
// row resolves to the hardcoded defaults (both true for claude/codex/gemini/
// opencode; cursor allows the gateway path only), so the policy is purely
// additive: disabling a path is an explicit admin choice. The CLI caches the
// resolved map at login and gates path selection on it.
// See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
type PlatformToolPolicy struct {
	ID              string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID  string   `gorm:"column:organizationId;uniqueIndex:PlatformToolPolicy_organizationId_toolSlug_key,priority:1;index:PlatformToolPolicy_organizationId_idx;not null" json:"organizationId"`
	ToolSlug        string   `gorm:"column:toolSlug;uniqueIndex:PlatformToolPolicy_organizationId_toolSlug_key,priority:2;not null" json:"toolSlug"`
	AllowVk         *bool    `gorm:"column:allowVk;default:true;not null" json:"allowVk"`
	AllowOtelDirect *bool    `gorm:"column:allowOtelDirect;default:true;not null" json:"allowOtelDirect"`
	CreatedAt       DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (PlatformToolPolicy) TableName() string {
	return "PlatformToolPolicy"
}

func (m *PlatformToolPolicy) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PlatformToolPolicy) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganizationInvite struct {
	ID                      string         `gorm:"column:id;primaryKey" json:"id"`
	Email                   string         `gorm:"column:email;not null" json:"email"`
	InviteCode              string         `gorm:"column:inviteCode;uniqueIndex:OrganizationInvite_inviteCode_key;not null" json:"inviteCode"`
	Expiration              *DateTime      `gorm:"column:expiration;type:timestamp(3)" json:"expiration"`
	Status                  string         `gorm:"column:status;default:'PENDING';not null" json:"status"`
	OrganizationID          string         `gorm:"column:organizationId;index:OrganizationInvite_organizationId_idx;not null" json:"organizationId"`
	TeamIds                 string         `gorm:"column:teamIds;not null" json:"teamIds"`
	TeamAssignments         datatypes.JSON `gorm:"column:teamAssignments" json:"teamAssignments"`
	Role                    string         `gorm:"column:role;not null" json:"role"`
	RequestedBy             *string        `gorm:"column:requestedBy;index:OrganizationInvite_requestedBy_idx" json:"requestedBy"`
	SubscriptionID          *string        `gorm:"column:subscriptionId" json:"subscriptionId"`
	AcceptedByUserID        *string        `gorm:"column:acceptedByUserId" json:"acceptedByUserId"`
	AcceptedViaIdentifierID *string        `gorm:"column:acceptedViaIdentifierId" json:"acceptedViaIdentifierId"`
	CreatedAt               DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization            *Organization  `gorm:"belongsTo"`
	RequestedByUser         *User          `gorm:"foreignKey:RequestedBy;belongsTo"`
}

func (OrganizationInvite) TableName() string {
	return "OrganizationInvite"
}

func (m *OrganizationInvite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganizationInvite) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Monitor struct {
	ID                string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID         string         `gorm:"column:projectId;uniqueIndex:Monitor_projectId_slug_key,priority:1;index:Monitor_projectId_idx;not null" json:"projectId"`
	ExperimentID      *string        `gorm:"column:experimentId;uniqueIndex:Monitor_experimentId_key" json:"experimentId"`
	EvaluatorID       *string        `gorm:"column:evaluatorId;index:Monitor_evaluatorId_idx" json:"evaluatorId"`
	CheckType         string         `gorm:"column:checkType;not null" json:"checkType"`
	Name              string         `gorm:"column:name;not null" json:"name"`
	Slug              string         `gorm:"column:slug;uniqueIndex:Monitor_projectId_slug_key,priority:2;not null" json:"slug"`
	ExecutionMode     string         `gorm:"column:executionMode;default:'ON_MESSAGE';not null" json:"executionMode"`
	Enabled           *bool          `gorm:"column:enabled;default:true;not null" json:"enabled"`
	Preconditions     datatypes.JSON `gorm:"column:preconditions;not null" json:"preconditions"`
	Parameters        datatypes.JSON `gorm:"column:parameters;not null" json:"parameters"`
	Mappings          datatypes.JSON `gorm:"column:mappings" json:"mappings"`
	Sample            *float64       `gorm:"column:sample;type:double precision;default:1;not null" json:"sample"`
	Level             *string        `gorm:"column:level;default:'trace';not null" json:"level"`
	ThreadIdleTimeout *int           `gorm:"column:threadIdleTimeout;size:32" json:"threadIdleTimeout"`
	CreatedAt         DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	AutoPausedAt      *DateTime      `gorm:"column:autoPausedAt;type:timestamp(3)" json:"autoPausedAt"`
	AutoPausedReason  *string        `gorm:"column:autoPausedReason" json:"autoPausedReason"`
	EvaluatorLabel    *string        `gorm:"column:evaluatorLabel" json:"evaluatorLabel"`
	CreatedByID       *string        `gorm:"column:createdById" json:"createdById"`
	UpdatedByID       *string        `gorm:"column:updatedById" json:"updatedById"`
	Project           *Project       `gorm:"belongsTo"`
	Experiment        *Experiment    `gorm:"belongsTo"`
	Evaluator         *Evaluator     `gorm:"belongsTo"`
}

func (Monitor) TableName() string {
	return "Monitor"
}

func (m *Monitor) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Monitor) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Cost struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID     string         `gorm:"column:projectId;index:Cost_projectId_idx;index:Cost_projectId_createdAt_idx,priority:1;not null" json:"projectId"`
	CostType      string         `gorm:"column:costType;index:Cost_costType_idx;not null" json:"costType"`
	CostName      *string        `gorm:"column:costName" json:"costName"`
	ReferenceType string         `gorm:"column:referenceType;index:Cost_referenceType_referenceId_idx,priority:1;not null" json:"referenceType"`
	ReferenceID   string         `gorm:"column:referenceId;index:Cost_referenceType_referenceId_idx,priority:2;not null" json:"referenceId"`
	Amount        float64        `gorm:"column:amount;type:double precision;not null" json:"amount"`
	Currency      string         `gorm:"column:currency;not null" json:"currency"`
	CreatedAt     DateTime       `gorm:"column:createdAt;index:Cost_projectId_createdAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	ExtraInfo     datatypes.JSON `gorm:"column:extraInfo" json:"extraInfo"`
	Project       *Project       `gorm:"belongsTo"`
}

func (Cost) TableName() string {
	return "Cost"
}

func (m *Cost) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Cost) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// PROJECTION of the topic-clustering event stream (ADR-051): topics are
// `topics_recorded` events; this table is written ONLY by
// TopicModelFoldProjection (cursor: TopicModelProjection) and is
// rebuildable by replay. Kept named `Topic` because it is the domain
// entity every surface reads — the `*Projection` suffix is reserved for
// cursor/read-model tables.
type Topic struct {
	ID                     string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID              string         `gorm:"column:projectId;index:Topic_projectId_idx;not null" json:"projectId"`
	Name                   string         `gorm:"column:name;not null" json:"name"`
	ParentID               *string        `gorm:"column:parentId;index:Topic_parentId_idx" json:"parentId"`
	EmbeddingsModel        string         `gorm:"column:embeddings_model;not null" json:"embeddings_model"`
	Centroid               datatypes.JSON `gorm:"column:centroid;not null" json:"centroid"`
	P95Distance            float64        `gorm:"column:p95Distance;type:double precision;not null" json:"p95Distance"`
	AutomaticallyGenerated *bool          `gorm:"column:automaticallyGenerated;default:true;not null" json:"automaticallyGenerated"`
	// The `topics_recorded` event that recorded this topic (its id in the
	// event log). Null means the row predates event ownership and has not
	// been seeded onto the stream yet; the projection sets it on every row
	// it writes.
	LastEventID *string  `gorm:"column:lastEventId" json:"lastEventId"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project     *Project `gorm:"belongsTo"`
	Parent      *Topic   `gorm:"foreignKey:ParentID;belongsTo"`
	Subtopics   []Topic  `gorm:"foreignKey:ParentID;constraint:Topic_parentId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (Topic) TableName() string {
	return "Topic"
}

func (m *Topic) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Topic) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Dataset struct {
	ID                   string            `gorm:"column:id;primaryKey" json:"id"`
	ProjectID            string            `gorm:"column:projectId;uniqueIndex:Dataset_projectId_slug_key,priority:1;not null" json:"projectId"`
	Name                 string            `gorm:"column:name;not null" json:"name"`
	Slug                 string            `gorm:"column:slug;uniqueIndex:Dataset_projectId_slug_key,priority:2;not null" json:"slug"`
	ColumnTypes          datatypes.JSON    `gorm:"column:columnTypes;not null" json:"columnTypes"`
	CreatedAt            DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime          `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	ArchivedAt           *DateTime         `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	Mapping              datatypes.JSON    `gorm:"column:mapping" json:"mapping"`
	UseS3                bool              `gorm:"column:useS3;default:false;not null" json:"useS3"`
	S3RecordCount        *int              `gorm:"column:s3RecordCount;size:32" json:"s3RecordCount"`
	ContentLayout        *string           `gorm:"column:contentLayout;default:'postgres';not null" json:"contentLayout"`
	Status               *string           `gorm:"column:status;default:'ready';not null" json:"status"`
	StatusError          *string           `gorm:"column:statusError" json:"statusError"`
	StagingKey           *string           `gorm:"column:stagingKey" json:"stagingKey"`
	SourceStoredObjectID *string           `gorm:"column:sourceStoredObjectId" json:"sourceStoredObjectId"`
	UploadFilename       *string           `gorm:"column:uploadFilename" json:"uploadFilename"`
	RowCount             *int              `gorm:"column:rowCount;size:32" json:"rowCount"`
	SizeBytes            *int64            `gorm:"column:sizeBytes" json:"sizeBytes"`
	ChunkCount           *int              `gorm:"column:chunkCount;size:32" json:"chunkCount"`
	ChunkOffsets         datatypes.JSON    `gorm:"column:chunkOffsets" json:"chunkOffsets"`
	Project              *Project          `gorm:"belongsTo"`
	DatasetRecords       []DatasetRecord   `gorm:"foreignKey:DatasetID;constraint:DatasetRecord_datasetId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	BatchEvaluations     []BatchEvaluation `gorm:"foreignKey:DatasetID;constraint:BatchEvaluation_datasetId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Dataset) TableName() string {
	return "Dataset"
}

func (m *Dataset) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Dataset) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DatasetRecord struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	DatasetID string         `gorm:"column:datasetId;index:DatasetRecord_datasetId_idx;index:DatasetRecord_projectId_datasetId_idx,priority:2;not null" json:"datasetId"`
	ProjectID string         `gorm:"column:projectId;index:DatasetRecord_projectId_idx;index:DatasetRecord_projectId_datasetId_idx,priority:1;not null" json:"projectId"`
	Entry     datatypes.JSON `gorm:"column:entry;not null" json:"entry"`
	CreatedAt DateTime       `gorm:"column:createdAt;index:DatasetRecord_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Dataset   *Dataset       `gorm:"belongsTo"`
	Project   *Project       `gorm:"belongsTo"`
}

func (DatasetRecord) TableName() string {
	return "DatasetRecord"
}

func (m *DatasetRecord) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DatasetRecord) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage.
type DatasetVersion struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string `gorm:"column:projectId;index:DatasetVersion_projectId_datasetId_idx,priority:1;not null" json:"projectId"`
	DatasetID string `gorm:"column:datasetId;uniqueIndex:DatasetVersion_datasetId_version_key,priority:1;index:DatasetVersion_projectId_datasetId_idx,priority:2;not null" json:"datasetId"`
	Version   int    `gorm:"column:version;uniqueIndex:DatasetVersion_datasetId_version_key,priority:2;size:32;not null" json:"version"`
	// A name given by hand; a named version is never thinned by retention.
	Name *string `gorm:"column:name" json:"name"`
	// "autosave", "manual", "experiment_run", "restore" or "initial".
	Trigger string `gorm:"column:trigger;not null" json:"trigger"`
	// SHA-256 over the columns and the rows in order; equal content saves no new version.
	ContentHash         string         `gorm:"column:contentHash;not null" json:"contentHash"`
	ColumnTypes         datatypes.JSON `gorm:"column:columnTypes;not null" json:"columnTypes"`
	RowCount            int            `gorm:"column:rowCount;size:32;not null" json:"rowCount"`
	SizeBytes           int64          `gorm:"column:sizeBytes;not null" json:"sizeBytes"`
	ChunkCount          int            `gorm:"column:chunkCount;size:32;not null" json:"chunkCount"`
	RestoredFromVersion *int           `gorm:"column:restoredFromVersion;size:32" json:"restoredFromVersion"`
	AuthorID            *string        `gorm:"column:authorId" json:"authorId"`
	// "user", "api_key" or "system": who saved the version when authorId is null.
	AuthorKind *string `gorm:"column:authorKind" json:"authorKind"`
	// Set when an experiment pins this version; a pinned version is never thinned.
	PinnedAt  *DateTime `gorm:"column:pinnedAt;type:timestamp(3)" json:"pinnedAt"`
	CreatedAt DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (DatasetVersion) TableName() string {
	return "DatasetVersion"
}

func (m *DatasetVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// An experiment run that read one dataset version; a version a run read is never thinned.
type DatasetVersionRun struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID    string   `gorm:"column:projectId;uniqueIndex:DatasetVersionRun_projectId_runId_datasetId_key,priority:1;not null" json:"projectId"`
	DatasetID    string   `gorm:"column:datasetId;uniqueIndex:DatasetVersionRun_projectId_runId_datasetId_key,priority:3;not null" json:"datasetId"`
	VersionID    string   `gorm:"column:versionId;index:DatasetVersionRun_versionId_idx;not null" json:"versionId"`
	RunID        string   `gorm:"column:runId;uniqueIndex:DatasetVersionRun_projectId_runId_datasetId_key,priority:2;not null" json:"runId"`
	ExperimentID *string  `gorm:"column:experimentId" json:"experimentId"`
	CreatedAt    DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (DatasetVersionRun) TableName() string {
	return "DatasetVersionRun"
}

func (m *DatasetVersionRun) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Dashboard struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The project that owns the board; only there can it be edited.
	ProjectID string `gorm:"column:projectId;index:Dashboard_projectId_idx;index:Dashboard_projectId_order_idx,priority:1;not null" json:"projectId"`
	Name      string `gorm:"column:name;not null" json:"name"`
	// Legacy board sort position. The Dashboards v2 app orders boards by a
	// member's favourites and never writes this; kept because the legacy
	// Analytics reports section and the REST reorder endpoint still write it,
	// and a rolling deploy's older image still reads it.
	Order int `gorm:"column:order;index:Dashboard_projectId_order_idx,priority:2;size:32;default:0;not null" json:"order"`
	// Free text under the board's name; null until a member writes one.
	// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
	Description *string `gorm:"column:description" json:"description"`
	// The member who created it; null for boards that predate the column.
	// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
	CreatedByID *string `gorm:"column:createdById" json:"createdById"`
	// Who sees it. Only the author changes it; a board with no author stays PROJECT.
	// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
	Scope string `gorm:"column:scope;index:Dashboard_organizationId_scope_idx,priority:2;default:'PROJECT';not null" json:"scope"`
	// The owning project's organization, stamped when the board is first set to
	// ORGANIZATION; null before. No relation, per the database rule.
	// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
	OrganizationID *string       `gorm:"column:organizationId;index:Dashboard_organizationId_scope_idx,priority:1" json:"organizationId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project        *Project      `gorm:"belongsTo"`
	Graphs         []CustomGraph `gorm:"foreignKey:DashboardID;constraint:CustomGraph_dashboardId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Dashboard) TableName() string {
	return "Dashboard"
}

func (m *Dashboard) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Dashboard) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A member's star, in their own order within a project: a board or a From
// LangWatch template. Personal: one member's stars never show to another. No
// foreign key (per the database rule): the dashboard service removes a board's
// favourites when it deletes the board, so a star never outlives its dashboard.
type DashboardFavourite struct {
	ID     string `gorm:"column:id;primaryKey" json:"id"`
	UserID string `gorm:"column:userId;uniqueIndex:DashboardFavourite_userId_dashboardId_key,priority:1;uniqueIndex:DashboardFavourite_userId_projectId_templateId_key,priority:1;index:DashboardFavourite_userId_projectId_position_idx,priority:1;not null" json:"userId"`
	// A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
	// the star is a template.
	DashboardID *string `gorm:"column:dashboardId;uniqueIndex:DashboardFavourite_userId_dashboardId_key,priority:2;index:DashboardFavourite_dashboardId_idx" json:"dashboardId"`
	// The starred From LangWatch template id; null when the star is a board.
	TemplateID *string `gorm:"column:templateId;uniqueIndex:DashboardFavourite_userId_projectId_templateId_key,priority:3" json:"templateId"`
	ProjectID  string  `gorm:"column:projectId;uniqueIndex:DashboardFavourite_userId_projectId_templateId_key,priority:2;index:DashboardFavourite_userId_projectId_position_idx,priority:2;not null" json:"projectId"`
	// The member's drag order among their own stars in the project.
	Position  int      `gorm:"column:position;index:DashboardFavourite_userId_projectId_position_idx,priority:3;size:32;not null" json:"position"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (DashboardFavourite) TableName() string {
	return "DashboardFavourite"
}

func (m *DashboardFavourite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
// made before the split were copied here with their ids, so stars and insight
// pointers still name them. No relation, per the database rule.
type DashboardsV2 struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The project that owns the board; only there can it be edited.
	ProjectID string `gorm:"column:projectId;index:DashboardsV2_projectId_idx;not null" json:"projectId"`
	Name      string `gorm:"column:name;not null" json:"name"`
	// Free text under the board's name; null until a member writes one.
	Description *string `gorm:"column:description" json:"description"`
	// The member who created it; null for boards that predate authorship.
	CreatedByID *string `gorm:"column:createdById" json:"createdById"`
	// Who sees it. Only the author changes it; a board with no author stays PROJECT.
	Scope string `gorm:"column:scope;index:DashboardsV2_organizationId_scope_idx,priority:2;default:'PROJECT';not null" json:"scope"`
	// The owning project's organization, stamped when the board is first set to
	// ORGANIZATION; null before.
	OrganizationID *string `gorm:"column:organizationId;index:DashboardsV2_organizationId_scope_idx,priority:1" json:"organizationId"`
	// Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
	// 2026-10-28); dropped in a later migration, once no serving release reads it.
	ExcludedOrigins datatypes.JSON `gorm:"column:excludedOrigins" json:"excludedOrigins"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (DashboardsV2) TableName() string {
	return "DashboardsV2"
}

func (m *DashboardsV2) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DashboardsV2) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
// unplaced. Widgets made before the split were copied from CustomGraph with their
// ids. No relation, per the database rule.
type DashboardWidget struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string `gorm:"column:projectId;index:DashboardWidget_projectId_dashboardId_idx,priority:1;not null" json:"projectId"`
	// The DashboardsV2 board it sits on; null while unplaced.
	DashboardID *string `gorm:"column:dashboardId;index:DashboardWidget_projectId_dashboardId_idx,priority:2" json:"dashboardId"`
	Name        string  `gorm:"column:name;not null" json:"name"`
	// The widget's srcdoc definition.
	Definition datatypes.JSON `gorm:"column:definition;not null" json:"definition"`
	GridColumn int            `gorm:"column:gridColumn;size:32;default:0;not null" json:"gridColumn"`
	GridRow    int            `gorm:"column:gridRow;size:32;default:0;not null" json:"gridRow"`
	ColSpan    *int           `gorm:"column:colSpan;size:32;default:1;not null" json:"colSpan"`
	RowSpan    *int           `gorm:"column:rowSpan;size:32;default:1;not null" json:"rowSpan"`
	CreatedAt  DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (DashboardWidget) TableName() string {
	return "DashboardWidget"
}

func (m *DashboardWidget) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DashboardWidget) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One insight in its owner's inbox in a project, folded from its `insight` aggregate
// (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
// milliseconds.
type InsightProjection struct {
	// The insight id, which is also its aggregate id.
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string `gorm:"column:projectId;index:InsightProjection_projectId_filedAt_idx,priority:1;index:InsightProjection_projectId_ownerUserId_filedAt_idx,priority:1;not null" json:"projectId"`
	Title     string `gorm:"column:title;not null" json:"title"`
	Body      string `gorm:"column:body;not null" json:"body"`
	// bad | watch | good
	Tone                 string  `gorm:"column:tone;not null" json:"tone"`
	Topic                *string `gorm:"column:topic" json:"topic"`
	ValidDays            int     `gorm:"column:validDays;size:32;not null" json:"validDays"`
	Lwql                 *string `gorm:"column:lwql" json:"lwql"`
	SourceConversationID *string `gorm:"column:sourceConversationId" json:"sourceConversationId"`
	SourceMessageID      *string `gorm:"column:sourceMessageId" json:"sourceMessageId"`
	FiledByUserID        *string `gorm:"column:filedByUserId" json:"filedByUserId"`
	FiledAt              float64 `gorm:"column:filedAt;index:InsightProjection_projectId_filedAt_idx,priority:2;index:InsightProjection_projectId_ownerUserId_filedAt_idx,priority:3;type:double precision;not null" json:"filedAt"`
	// Set by a later run that found the insight still true.
	RenewedAt         *float64 `gorm:"column:renewedAt;type:double precision" json:"renewedAt"`
	CreatedAt         float64  `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt         float64  `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	OccurredAt        float64  `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	AcceptedAt        float64  `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string   `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string   `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	// The board and widget it came from. A pointer only, never a relation: the row reads the
	// same once either is deleted, so the names are kept as filed.
	BoardID    *string `gorm:"column:boardId" json:"boardId"`
	BoardName  *string `gorm:"column:boardName" json:"boardName"`
	WidgetID   *string `gorm:"column:widgetId" json:"widgetId"`
	WidgetName *string `gorm:"column:widgetName" json:"widgetName"`
	// chat | run
	FiledVia *string `gorm:"column:filedVia;default:'chat';not null" json:"filedVia"`
	// The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step.
	ReplayStart              *float64 `gorm:"column:replayStart;type:double precision" json:"replayStart"`
	ReplayEnd                *float64 `gorm:"column:replayEnd;type:double precision" json:"replayEnd"`
	ReplayGranularitySeconds *int     `gorm:"column:replayGranularitySeconds;size:32" json:"replayGranularitySeconds"`
	// The values in force when it was filed: `{ period, parameters }`.
	ReplayContext datatypes.JSON `gorm:"column:replayContext" json:"replayContext"`
	// Whose insight it is: the one person who reads it. NULL on a row folded before the
	// column existed, which is owned by whoever filed it (`filedByUserId`).
	OwnerUserID *string `gorm:"column:ownerUserId;index:InsightProjection_projectId_ownerUserId_filedAt_idx,priority:2" json:"ownerUserId"`
	// The stable key a daily run matches a repeated finding by, so it renews the insight
	// and files no twin. NULL for an insight saved from a chat and on a row folded before
	// the column existed.
	FindingKey *string `gorm:"column:findingKey" json:"findingKey"`
}

func (InsightProjection) TableName() string {
	return "InsightProjection"
}

// One reader's own state on one insight (seen, done, kept), folded from the same stream.
// An insight has one reader today, its owner; a later forward may add more.
type InsightReaderProjection struct {
	ID                string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID         string   `gorm:"column:projectId;uniqueIndex:InsightReaderProjection_projectId_insightId_userId_key,priority:1;index:InsightReaderProjection_projectId_userId_idx,priority:1;not null" json:"projectId"`
	InsightID         string   `gorm:"column:insightId;uniqueIndex:InsightReaderProjection_projectId_insightId_userId_key,priority:2;not null" json:"insightId"`
	UserID            string   `gorm:"column:userId;uniqueIndex:InsightReaderProjection_projectId_insightId_userId_key,priority:3;index:InsightReaderProjection_projectId_userId_idx,priority:2;not null" json:"userId"`
	SeenAt            *float64 `gorm:"column:seenAt;type:double precision" json:"seenAt"`
	ArchivedAt        *float64 `gorm:"column:archivedAt;type:double precision" json:"archivedAt"`
	KeptAt            *float64 `gorm:"column:keptAt;type:double precision" json:"keptAt"`
	CreatedAt         float64  `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt         float64  `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	OccurredAt        float64  `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	AcceptedAt        float64  `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string   `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string   `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (InsightReaderProjection) TableName() string {
	return "InsightReaderProjection"
}

// One person's daily insights run on one board in a project, folded from its
// `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
// replaying the event log. Epochs in milliseconds.
type InsightDailyScheduleProjection struct {
	// The schedule id, which is also its aggregate id: one per project, person and board.
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string `gorm:"column:projectId;uniqueIndex:InsightDailySchedule_projectId_userId_board_key,priority:1;not null" json:"projectId"`
	UserID    string `gorm:"column:userId;uniqueIndex:InsightDailySchedule_projectId_userId_board_key,priority:2;not null" json:"userId"`
	// dashboard | template. A pointer only, never a relation: a run checks the board is there.
	BoardKind string `gorm:"column:boardKind;uniqueIndex:InsightDailySchedule_projectId_userId_board_key,priority:3;not null" json:"boardKind"`
	BoardID   string `gorm:"column:boardId;uniqueIndex:InsightDailySchedule_projectId_userId_board_key,priority:4;not null" json:"boardId"`
	BoardName string `gorm:"column:boardName;not null" json:"boardName"`
	// on | off | undecided.
	State *string `gorm:"column:state;default:'undecided';not null" json:"state"`
	// board_deleted | skipped_days | member_removed | project_archived. Why the system turned
	// it off. NULL while it is on, and when the person turned it off.
	OffReason   *string `gorm:"column:offReason" json:"offReason"`
	Hour        *int    `gorm:"column:hour;size:32" json:"hour"`
	Timezone    *string `gorm:"column:timezone" json:"timezone"`
	MaxInsights *int    `gorm:"column:maxInsights;size:32" json:"maxInsights"`
	// The last run that recorded an outcome. NULL until one did.
	LastRunID *string  `gorm:"column:lastRunId" json:"lastRunId"`
	LastRunAt *float64 `gorm:"column:lastRunAt;type:double precision" json:"lastRunAt"`
	// filed | nothing | failed | skipped
	LastRunOutcome *string `gorm:"column:lastRunOutcome" json:"lastRunOutcome"`
	LastRunReason  *string `gorm:"column:lastRunReason" json:"lastRunReason"`
	LastRunFiled   *int    `gorm:"column:lastRunFiled;size:32" json:"lastRunFiled"`
	// How many insights the run found still true and renewed.
	LastRunRenewed        *int    `gorm:"column:lastRunRenewed;size:32" json:"lastRunRenewed"`
	LastRunConversationID *string `gorm:"column:lastRunConversationId" json:"lastRunConversationId"`
	// The scheduled days in a row whose run ended skipped, and the slot of the newest counted.
	ConsecutiveSkips  int      `gorm:"column:consecutiveSkips;size:32;default:0;not null" json:"consecutiveSkips"`
	SkipStreakSlot    *float64 `gorm:"column:skipStreakSlot;type:double precision" json:"skipStreakSlot"`
	CreatedAt         float64  `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt         float64  `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	OccurredAt        float64  `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	AcceptedAt        float64  `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string   `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string   `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (InsightDailyScheduleProjection) TableName() string {
	return "InsightDailyScheduleProjection"
}

type SavedView struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string         `gorm:"column:projectId;index:SavedView_projectId_idx;index:SavedView_projectId_order_idx,priority:1;index:SavedView_projectId_kind_order_idx,priority:1;not null" json:"projectId"`
	UserID    *string        `gorm:"column:userId;index:SavedView_userId_idx" json:"userId"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	Filters   datatypes.JSON `gorm:"column:filters;default:'{}';not null" json:"filters"`
	Query     *string        `gorm:"column:query" json:"query"`
	Period    datatypes.JSON `gorm:"column:period" json:"period"`
	Order     int            `gorm:"column:order;index:SavedView_projectId_order_idx,priority:2;index:SavedView_projectId_kind_order_idx,priority:3;size:32;default:0;not null" json:"order"`
	// Discriminator for the two storage shapes that share this table:
	// - "v1-traces-filter" (default) - legacy /messages page filter map saved views
	// - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
	// Each client filters by its own kind on read so neither sees the other's rows.
	Kind      *string  `gorm:"column:kind;index:SavedView_projectId_kind_order_idx,priority:2;default:'v1-traces-filter';not null" json:"kind"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project   *Project `gorm:"belongsTo"`
	User      *User    `gorm:"belongsTo"`
}

func (SavedView) TableName() string {
	return "SavedView"
}

func (m *SavedView) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SavedView) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CustomGraph struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string `gorm:"column:projectId;index:CustomGraph_projectId_idx;index:CustomGraph_projectId_kind_idx,priority:1;not null" json:"projectId"`
	Name      string `gorm:"column:name;not null" json:"name"`
	// The chart's definition, in the shape its `kind` names.
	Graph   datatypes.JSON `gorm:"column:graph;not null" json:"graph"`
	Filters datatypes.JSON `gorm:"column:filters" json:"filters"`
	// Discriminator for the chart shapes that share this table:
	// - "builder" (default) - the chart builder's own `graph` payload
	// - "workbench_sql"     - a saved LangWatchQL workbench chart, whose
	// `graph` holds the versioned { sql, parameters, vegaLiteSpec }
	// definition (see server/analytics/saved-workbench-charts)
	// - "dashboard_srcdoc"  - frozen legacy copies of Dashboards v2 widgets, which
	// now live in DashboardWidget; nothing new writes this kind
	// Each reader filters by its own kind, so neither sees the other's rows and
	// a builder chart can never be interpreted as a workbench definition.
	Kind        *string       `gorm:"column:kind;index:CustomGraph_projectId_kind_idx,priority:2;default:'builder';not null" json:"kind"`
	CreatedAt   DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	DashboardID *string       `gorm:"column:dashboardId;index:CustomGraph_dashboardId_idx" json:"dashboardId"`
	GridColumn  int           `gorm:"column:gridColumn;size:32;default:0;not null" json:"gridColumn"`
	GridRow     int           `gorm:"column:gridRow;size:32;default:0;not null" json:"gridRow"`
	ColSpan     *int          `gorm:"column:colSpan;size:32;default:1;not null" json:"colSpan"`
	RowSpan     *int          `gorm:"column:rowSpan;size:32;default:1;not null" json:"rowSpan"`
	Project     *Project      `gorm:"belongsTo"`
	Dashboard   *Dashboard    `gorm:"belongsTo"`
	TriggerSent []TriggerSent `gorm:"foreignKey:CustomGraphID;constraint:TriggerSent_customGraphId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Trigger     *Trigger      `gorm:"foreignKey:CustomGraphID;constraint:Trigger_customGraphId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (CustomGraph) TableName() string {
	return "CustomGraph"
}

func (m *CustomGraph) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CustomGraph) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BatchEvaluation struct {
	ID           string         `gorm:"column:id;primaryKey;uniqueIndex:BatchEvaluation_id_key" json:"id"`
	ExperimentID string         `gorm:"column:experimentId;index:BatchEvaluation_experimentId_idx;not null" json:"experimentId"`
	ProjectID    string         `gorm:"column:projectId;index:BatchEvaluation_projectId_idx;not null" json:"projectId"`
	Data         datatypes.JSON `gorm:"column:data;default:'{}';not null" json:"data"`
	Status       string         `gorm:"column:status;not null" json:"status"`
	Score        float64        `gorm:"column:score;type:double precision;not null" json:"score"`
	Label        *string        `gorm:"column:label" json:"label"`
	Passed       bool           `gorm:"column:passed;not null" json:"passed"`
	Details      string         `gorm:"column:details;not null" json:"details"`
	Cost         float64        `gorm:"column:cost;type:double precision;not null" json:"cost"`
	DatasetSlug  string         `gorm:"column:datasetSlug;not null" json:"datasetSlug"`
	DatasetID    string         `gorm:"column:datasetId;index:BatchEvaluation_datasetId_idx;not null" json:"datasetId"`
	Evaluation   string         `gorm:"column:evaluation;not null" json:"evaluation"`
	CreatedAt    DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Experiment   *Experiment    `gorm:"belongsTo"`
	Project      *Project       `gorm:"belongsTo"`
	Dataset      *Dataset       `gorm:"belongsTo"`
}

func (BatchEvaluation) TableName() string {
	return "BatchEvaluation"
}

func (m *BatchEvaluation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BatchEvaluation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Trigger struct {
	ID                   string                    `gorm:"column:id;primaryKey" json:"id"`
	Name                 string                    `gorm:"column:name;not null" json:"name"`
	ProjectID            string                    `gorm:"column:projectId;index:Trigger_projectId_idx;not null" json:"projectId"`
	Action               string                    `gorm:"column:action;not null" json:"action"`
	TriggerKind          string                    `gorm:"column:triggerKind;default:'AUTOMATION';not null" json:"triggerKind"`
	ActionParams         datatypes.JSON            `gorm:"column:actionParams;not null" json:"actionParams"`
	Filters              datatypes.JSON            `gorm:"column:filters;not null" json:"filters"`
	FilterQuery          *string                   `gorm:"column:filterQuery" json:"filterQuery"`
	LastRunAt            float64                   `gorm:"column:lastRunAt;type:double precision;default:0;not null" json:"lastRunAt"`
	CreatedAt            DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime                  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Active               *bool                     `gorm:"column:active;default:true;not null" json:"active"`
	PausedReason         *string                   `gorm:"column:pausedReason" json:"pausedReason"`
	PausedAt             *DateTime                 `gorm:"column:pausedAt;type:timestamptz(3)" json:"pausedAt"`
	Message              *string                   `gorm:"column:message" json:"message"`
	Deleted              bool                      `gorm:"column:deleted;default:false;not null" json:"deleted"`
	AlertType            *string                   `gorm:"column:alertType" json:"alertType"`
	SlackTemplateType    *string                   `gorm:"column:slackTemplateType" json:"slackTemplateType"`
	SlackTemplate        *string                   `gorm:"column:slackTemplate" json:"slackTemplate"`
	EmailSubjectTemplate *string                   `gorm:"column:emailSubjectTemplate" json:"emailSubjectTemplate"`
	EmailBodyTemplate    *string                   `gorm:"column:emailBodyTemplate" json:"emailBodyTemplate"`
	NotificationCadence  *string                   `gorm:"column:notificationCadence;default:'immediate';not null" json:"notificationCadence"`
	TraceDebounceMs      *int                      `gorm:"column:traceDebounceMs;size:32;default:30000;not null" json:"traceDebounceMs"`
	CustomGraphID        *string                   `gorm:"column:customGraphId;uniqueIndex:Trigger_customGraphId_key" json:"customGraphId"`
	Project              *Project                  `gorm:"belongsTo"`
	CustomGraph          *CustomGraph              `gorm:"belongsTo"`
	TriggerSent          []TriggerSent             `gorm:"foreignKey:TriggerID;constraint:TriggerSent_triggerId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	WebhookDeliveries    []WebhookEndpointDelivery `gorm:"foreignKey:TriggerID;constraint:WebhookEndpointDelivery_triggerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	LatestEvaluation     *TriggerLatestEvaluation  `gorm:"foreignKey:TriggerID;constraint:TriggerLatestEvaluation_triggerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Trigger) TableName() string {
	return "Trigger"
}

func (m *Trigger) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Trigger) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One named Slack connection (ADR-093 §5a); an organization holds as many as
// it likes, each scoped to the organization or to one project. The secret is
// AES-256-GCM ciphertext and never returned to a client. `secretFingerprint`
// is an HMAC of the plaintext, so one secret is one connection per scope
// without decrypting anything. No foreign key: `scopeId` is
// polymorphic (ADR-021), as on RetentionPolicy and CustomLLMModelCost.
type SlackIntegration struct {
	ID                  string   `gorm:"column:id;primaryKey" json:"id"`
	Name                string   `gorm:"column:name;not null" json:"name"`
	Kind                string   `gorm:"column:kind;not null" json:"kind"`
	ScopeType           string   `gorm:"column:scopeType;uniqueIndex:SlackIntegration_organizationId_scopeType_scopeId_secretFin_key,priority:2;index:SlackIntegration_scopeType_scopeId_idx,priority:1;not null" json:"scopeType"`
	ScopeID             string   `gorm:"column:scopeId;uniqueIndex:SlackIntegration_organizationId_scopeType_scopeId_secretFin_key,priority:3;index:SlackIntegration_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	OrganizationID      string   `gorm:"column:organizationId;uniqueIndex:SlackIntegration_organizationId_scopeType_scopeId_secretFin_key,priority:1;not null" json:"organizationId"`
	BotTokenEncrypted   *string  `gorm:"column:botTokenEncrypted" json:"botTokenEncrypted"`
	WebhookURLEncrypted *string  `gorm:"column:webhookUrlEncrypted" json:"webhookUrlEncrypted"`
	SecretFingerprint   string   `gorm:"column:secretFingerprint;uniqueIndex:SlackIntegration_organizationId_scopeType_scopeId_secretFin_key,priority:4;not null" json:"secretFingerprint"`
	SecretHint          string   `gorm:"column:secretHint;not null" json:"secretHint"`
	SlackTeamID         *string  `gorm:"column:slackTeamId" json:"slackTeamId"`
	SlackTeamName       *string  `gorm:"column:slackTeamName" json:"slackTeamName"`
	CreatedByID         string   `gorm:"column:createdById;not null" json:"createdById"`
	UpdatedByID         string   `gorm:"column:updatedById;not null" json:"updatedById"`
	CreatedAt           DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SlackIntegration) TableName() string {
	return "SlackIntegration"
}

func (m *SlackIntegration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SlackIntegration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// An automation holding a Slack connection in use (ARCHITECTURE.md §3, the
// claims ruling): a claimed connection is never deleted. `projectId` is the
// claimant's project; `organizationId` is the connection's tenancy anchor.
// No foreign key: slack's service refuses to delete a claimed connection.
type SlackConnectionClaim struct {
	ConnectionID   string   `gorm:"column:connectionId;primaryKey" json:"connectionId"`
	ClaimantID     string   `gorm:"column:claimantId;primaryKey" json:"claimantId"`
	ClaimantLabel  string   `gorm:"column:claimantLabel;not null" json:"claimantLabel"`
	OrganizationID string   `gorm:"column:organizationId;index:slack_connection_claim_organizationId_idx;not null" json:"organizationId"`
	ProjectID      string   `gorm:"column:projectId;not null" json:"projectId"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SlackConnectionClaim) TableName() string {
	return "slack_connection_claim"
}

func (m *SlackConnectionClaim) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SlackConnectionClaim) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Org-anchored outbound webhook endpoint (the webhook platform, enterprise
// only). The signing secret is AES-256-GCM encrypted at rest (shared
// encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
// once, at create or roll time. `enabledEvents` holds registry event types
// or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
// 72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
type WebhookEndpoint struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:WebhookEndpoint_organizationId_idempotencyKey_key,priority:1;index:WebhookEndpoint_organizationId_status_idx,priority:1;not null" json:"organizationId"`
	// Where this endpoint delivers. Everything above the last hop is shared:
	// same batching, same retry ladder, same delivery log, same signature over
	// the same bytes.
	DestinationKind string `gorm:"column:destinationKind;default:'http';not null" json:"destinationKind"`
	// The receiver URL. Present on `http` endpoints and null on the others;
	// the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
	URL *string `gorm:"column:url" json:"url"`
	// The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
	// region and the owning account are read off it rather than configured
	// separately, so they can never disagree with the queue.
	SqsQueueURL *string `gorm:"column:sqsQueueUrl" json:"sqsQueueUrl"`
	// The role we assume to write to the queue, and the ExternalId that proves
	// the request came from us rather than from anyone who learned the role's
	// name. The recommended credential mode.
	SqsRoleArn    *string `gorm:"column:sqsRoleArn" json:"sqsRoleArn"`
	SqsExternalID *string `gorm:"column:sqsExternalId" json:"sqsExternalId"`
	// A static key pair for the queue, encrypted at rest exactly like the
	// signing secret. Both halves are present or both are absent: a half-filled
	// pair is what the AWS SDK reads as "we have credentials" and stops looking
	// with, so the CHECK constraint refuses it.
	SqsAccessKeyID              *string `gorm:"column:sqsAccessKeyId" json:"sqsAccessKeyId"`
	SqsSecretAccessKeyEncrypted *string `gorm:"column:sqsSecretAccessKeyEncrypted" json:"sqsSecretAccessKeyEncrypted"`
	SecretEncrypted             string  `gorm:"column:secretEncrypted;not null" json:"secretEncrypted"`
	// How deliveries are formatted and signed. Null: the batch envelope signed
	// `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
	// only for endpoints migrated from governance anomaly destinations.
	SignatureScheme *string `gorm:"column:signatureScheme" json:"signatureScheme"`
	// Set only by governance's anomaly destination migration: a create that
	// repeats it in the same organization answers the endpoint it made. An
	// archived endpoint gives its key up, so a retry after cleanup creates anew.
	IdempotencyKey *string `gorm:"column:idempotencyKey;uniqueIndex:WebhookEndpoint_organizationId_idempotencyKey_key,priority:2" json:"idempotencyKey"`
	// The secret this endpoint signed with before the last roll. Deliveries
	// carry a second `v1=` computed from it until it expires, so a receiver
	// swaps secrets on its own schedule instead of dropping deliveries in the
	// gap. Null once the window closes or if the secret was never rolled.
	PreviousSecretEncrypted *string `gorm:"column:previousSecretEncrypted" json:"previousSecretEncrypted"`
	// When the previous secret stops being signed with and accepted.
	PreviousSecretExpiresAt *DateTime          `gorm:"column:previousSecretExpiresAt;type:timestamp(3)" json:"previousSecretExpiresAt"`
	EnabledEvents           ScalarList[string] `gorm:"column:enabledEvents;type:text[]" json:"enabledEvents"`
	Status                  string             `gorm:"column:status;index:WebhookEndpoint_organizationId_status_idx,priority:2;default:'ACTIVE';not null" json:"status"`
	// Why the endpoint is disabled: "auto_failures_72h" or "manual".
	DisabledReason *string   `gorm:"column:disabledReason" json:"disabledReason"`
	DisabledAt     *DateTime `gorm:"column:disabledAt;type:timestamp(3)" json:"disabledAt"`
	// Start of the current consecutive-failure streak; cleared on any 2xx.
	// 72 hours of unbroken failures from this instant auto-disables.
	FailingSince *DateTime `gorm:"column:failingSince;type:timestamp(3)" json:"failingSince"`
	// Delivery controls, each server-bounded (see webhookDeliveryControls
	// constants): envelopes per POST, coalescing window before a partial
	// batch ships, and concurrent POSTs to this receiver.
	MaxBatchSize    *int `gorm:"column:maxBatchSize;size:32;default:100;not null" json:"maxBatchSize"`
	MaxBatchDelayMs *int `gorm:"column:maxBatchDelayMs;size:32;default:250;not null" json:"maxBatchDelayMs"`
	MaxInFlight     *int `gorm:"column:maxInFlight;size:32;default:4;not null" json:"maxInFlight"`
	// Self-hosted only: deliveries accept the receiver's self-signed TLS
	// certificate. Off by default; the hosted product ignores it.
	AllowSelfSignedCertificate bool                      `gorm:"column:allowSelfSignedCertificate;default:false;not null" json:"allowSelfSignedCertificate"`
	LastSuccessAt              *DateTime                 `gorm:"column:lastSuccessAt;type:timestamp(3)" json:"lastSuccessAt"`
	LastFailureAt              *DateTime                 `gorm:"column:lastFailureAt;type:timestamp(3)" json:"lastFailureAt"`
	ArchivedAt                 *DateTime                 `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	CreatedAt                  DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime                  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization               *Organization             `gorm:"belongsTo"`
	Deliveries                 []WebhookEndpointDelivery `gorm:"foreignKey:EndpointID;constraint:WebhookEndpointDelivery_endpointId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (WebhookEndpoint) TableName() string {
	return "WebhookEndpoint"
}

func (m *WebhookEndpoint) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

// The per-ATTEMPT delivery log for BOTH webhook channels: the endpoints
// platform's batches and the automations channel's trigger fires. One row per
// HTTP attempt; `dispatchId` groups every attempt of one logical send (the
// value of the X-LangWatch-Delivery-Id or X-LangWatch-Event-Id header). A slim
// facts table: outcome, status, latency, a capped error message, and a
// truncated copy of a failure response. Our request content (URL, headers,
// body) is never stored. Pruned after 30 days.
//
// The two channels used to keep separate tables with the same columns and two
// prune implementations. They are one table with a discriminator now, so the
// tenancy columns of the channel that did not write a row are null.
type WebhookEndpointDelivery struct {
	ID      string `gorm:"column:id;primaryKey" json:"id"`
	Channel string `gorm:"column:channel;default:'platform';not null" json:"channel"`
	// Platform tenancy. Null on automations rows.
	OrganizationID *string `gorm:"column:organizationId;index:WebhookEndpointDelivery_organizationId_firedAt_idx,priority:1" json:"organizationId"`
	EndpointID     *string `gorm:"column:endpointId;index:WebhookEndpointDelivery_endpointId_firedAt_idx,priority:1;index:WebhookEndpointDelivery_endpointId_dispatchId_idx,priority:1" json:"endpointId"`
	// Automations tenancy. Null on platform rows.
	ProjectID  *string `gorm:"column:projectId;index:WebhookEndpointDelivery_projectId_triggerId_firedAt_idx,priority:1;index:WebhookEndpointDelivery_projectId_triggerId_dispatchId_idx,priority:1" json:"projectId"`
	TriggerID  *string `gorm:"column:triggerId;index:WebhookEndpointDelivery_projectId_triggerId_firedAt_idx,priority:2;index:WebhookEndpointDelivery_projectId_triggerId_dispatchId_idx,priority:2" json:"triggerId"`
	DispatchID string  `gorm:"column:dispatchId;index:WebhookEndpointDelivery_endpointId_dispatchId_idx,priority:2;index:WebhookEndpointDelivery_projectId_triggerId_dispatchId_idx,priority:3;not null" json:"dispatchId"`
	// 1-based attempt and the batch's envelope count. Both null on automations
	// rows, which record neither: one trigger fire is one event, and the
	// channel's sender does not thread its attempt number into the log.
	Attempt        *int             `gorm:"column:attempt;size:32" json:"attempt"`
	EventCount     *int             `gorm:"column:eventCount;size:32" json:"eventCount"`
	ResponseStatus *int             `gorm:"column:responseStatus;size:32" json:"responseStatus"`
	LatencyMs      *int             `gorm:"column:latencyMs;size:32" json:"latencyMs"`
	Error          *string          `gorm:"column:error" json:"error"`
	Response       datatypes.JSON   `gorm:"column:response" json:"response"`
	Outcome        string           `gorm:"column:outcome;not null" json:"outcome"`
	FiredAt        DateTime         `gorm:"column:firedAt;index:WebhookEndpointDelivery_endpointId_firedAt_idx,priority:2;index:WebhookEndpointDelivery_organizationId_firedAt_idx,priority:2;index:WebhookEndpointDelivery_firedAt_idx;index:WebhookEndpointDelivery_projectId_triggerId_firedAt_idx,priority:3;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"firedAt"`
	CreatedAt      DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Endpoint       *WebhookEndpoint `gorm:"foreignKey:EndpointID;belongsTo"`
	Project        *Project         `gorm:"belongsTo"`
	Trigger        *Trigger         `gorm:"belongsTo"`
}

func (WebhookEndpointDelivery) TableName() string {
	return "WebhookEndpointDelivery"
}

func (m *WebhookEndpointDelivery) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.FiredAt.IsZero() {
		m.FiredAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Experiment struct {
	ID               string            `gorm:"column:id;primaryKey" json:"id"`
	Name             *string           `gorm:"column:name" json:"name"`
	Type             string            `gorm:"column:type;not null" json:"type"`
	Slug             string            `gorm:"column:slug;uniqueIndex:Experiment_projectId_slug_key,priority:2;not null" json:"slug"`
	ProjectID        string            `gorm:"column:projectId;uniqueIndex:Experiment_projectId_slug_key,priority:1;index:Experiment_projectId_idx;index:Experiment_projectId_archivedAt_idx,priority:1;not null" json:"projectId"`
	WorkflowID       *string           `gorm:"column:workflowId;index:Experiment_workflowId_idx" json:"workflowId"`
	CreatedAt        DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime          `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	ArchivedAt       *DateTime         `gorm:"column:archivedAt;index:Experiment_projectId_archivedAt_idx,priority:2;type:timestamp(3)" json:"archivedAt"`
	WorkbenchState   datatypes.JSON    `gorm:"column:workbenchState" json:"workbenchState"`
	WorkbenchVersion int               `gorm:"column:workbenchVersion;size:32;default:0;not null" json:"workbenchVersion"`
	Project          *Project          `gorm:"belongsTo"`
	Workflow         *Workflow         `gorm:"belongsTo"`
	BatchEvaluations []BatchEvaluation `gorm:"foreignKey:ExperimentID;constraint:BatchEvaluation_experimentId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Monitor          *Monitor          `gorm:"foreignKey:ExperimentID;constraint:Monitor_experimentId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (Experiment) TableName() string {
	return "Experiment"
}

func (m *Experiment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Experiment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ExperimentVersion struct {
	ID           string `gorm:"column:id;primaryKey" json:"id"`
	ExperimentID string `gorm:"column:experimentId;uniqueIndex:ExperimentVersion_experimentId_version_key,priority:1;index:ExperimentVersion_projectId_experimentId_autoSaved_idx,priority:2;index:ExperimentVersion_projectId_experimentId_version_idx,priority:2;index:ExperimentVersion_projectId_experimentId_counterVersion_idx,priority:2;not null" json:"experimentId"`
	ProjectID    string `gorm:"column:projectId;index:ExperimentVersion_projectId_experimentId_autoSaved_idx,priority:1;index:ExperimentVersion_projectId_experimentId_version_idx,priority:1;index:ExperimentVersion_projectId_experimentId_counterVersion_idx,priority:1;not null" json:"projectId"`
	// The number the version is addressed and displayed by. Numbered rows hold
	// a gap-free sequence. The rolling autosave row rides on the counter
	// instead, which is what keeps it clear of that sequence.
	Version int `gorm:"column:version;uniqueIndex:ExperimentVersion_experimentId_version_key,priority:2;index:ExperimentVersion_projectId_experimentId_version_idx,priority:3;size:32;not null" json:"version"`
	// `Experiment.workbenchVersion` at the moment this row was last written.
	// The rolling row is rewritten in place, so this is the only column that
	// says how recent its content is, and the version list sorts and pages by
	// it. It also joins a row to the live state: the row whose `counterVersion`
	// equals the experiment's counter holds what the workbench shows now.
	CounterVersion int     `gorm:"column:counterVersion;index:ExperimentVersion_projectId_experimentId_counterVersion_idx,priority:3;size:32;not null" json:"counterVersion"`
	AutoSaved      bool    `gorm:"column:autoSaved;index:ExperimentVersion_projectId_experimentId_autoSaved_idx,priority:3;default:false;not null" json:"autoSaved"`
	CommitMessage  *string `gorm:"column:commitMessage" json:"commitMessage"`
	AuthorID       *string `gorm:"column:authorId" json:"authorId"`
	// Who wrote this version: "user", "langy" or "api".
	AuthorLabel string `gorm:"column:authorLabel;not null" json:"authorLabel"`
	// The run that wrote this version, when a run wrote it.
	//
	// A run writes its cells into the workbench state, which advances the
	// counter. Without this, the page that started that run reads the bump as
	// somebody else's write, refuses to save, and asks the reader to reload
	// over their own unsaved edits. Named here, the page adopts the version and
	// carries on.
	RunID *string        `gorm:"column:runId" json:"runId"`
	State datatypes.JSON `gorm:"column:state;not null" json:"state"`
	// Schema version of `state`, so a later shape change can be read forward.
	SchemaVersion *string  `gorm:"column:schemaVersion;default:'1';not null" json:"schemaVersion"`
	CreatedAt     DateTime `gorm:"column:createdAt;index:ExperimentVersion_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ExperimentVersion) TableName() string {
	return "ExperimentVersion"
}

func (m *ExperimentVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ExperimentVersion) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Annotation struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID      string         `gorm:"column:projectId;index:Annotation_projectId_idx;index:Annotation_projectId_traceId_idx,priority:1;index:Annotation_projectId_traceId_anchorKind_idx,priority:1;not null" json:"projectId"`
	TraceID        string         `gorm:"column:traceId;index:Annotation_traceId_idx;index:Annotation_projectId_traceId_idx,priority:2;index:Annotation_projectId_traceId_anchorKind_idx,priority:2;not null" json:"traceId"`
	Comment        *string        `gorm:"column:comment" json:"comment"`
	IsThumbsUp     *bool          `gorm:"column:isThumbsUp" json:"isThumbsUp"`
	UserID         *string        `gorm:"column:userId;index:Annotation_userId_idx" json:"userId"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Email          *string        `gorm:"column:email" json:"email"`
	ScoreOptions   datatypes.JSON `gorm:"column:scoreOptions" json:"scoreOptions"`
	ExpectedOutput *string        `gorm:"column:expectedOutput" json:"expectedOutput"`
	AnchorKind     *string        `gorm:"column:anchorKind;index:Annotation_projectId_traceId_anchorKind_idx,priority:3" json:"anchorKind"`
	AnchorID       *string        `gorm:"column:anchorId" json:"anchorId"`
	AnchorPath     *string        `gorm:"column:anchorPath" json:"anchorPath"`
	Project        *Project       `gorm:"belongsTo"`
	User           *User          `gorm:"belongsTo"`
}

func (Annotation) TableName() string {
	return "Annotation"
}

func (m *Annotation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Annotation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ModelProvider struct {
	ID                         string               `gorm:"column:id;primaryKey" json:"id"`
	Name                       string               `gorm:"column:name;not null" json:"name"`
	Provider                   string               `gorm:"column:provider;not null" json:"provider"`
	RoutingHandle              *string              `gorm:"column:routingHandle;uniqueIndex:ModelProvider_organizationId_routingHandle_key,priority:2" json:"routingHandle"`
	Enabled                    bool                 `gorm:"column:enabled;not null" json:"enabled"`
	CustomKeys                 datatypes.JSON       `gorm:"column:customKeys" json:"customKeys"`
	ExtraHeaders               datatypes.JSON       `gorm:"column:extraHeaders" json:"extraHeaders"`
	CustomModels               datatypes.JSON       `gorm:"column:customModels" json:"customModels"`
	CustomEmbeddingsModels     datatypes.JSON       `gorm:"column:customEmbeddingsModels" json:"customEmbeddingsModels"`
	DeploymentMapping          datatypes.JSON       `gorm:"column:deploymentMapping" json:"deploymentMapping"`
	RateLimitRpm               *int                 `gorm:"column:rateLimitRpm;size:32" json:"rateLimitRpm"`
	RateLimitTpm               *int                 `gorm:"column:rateLimitTpm;size:32" json:"rateLimitTpm"`
	RateLimitRpd               *int                 `gorm:"column:rateLimitRpd;size:32" json:"rateLimitRpd"`
	RotationPolicy             string               `gorm:"column:rotationPolicy;default:'MANUAL';not null" json:"rotationPolicy"`
	ProviderConfig             datatypes.JSON       `gorm:"column:providerConfig" json:"providerConfig"`
	FallbackPriorityGlobal     *int                 `gorm:"column:fallbackPriorityGlobal;index:ModelProvider_fallbackPriorityGlobal_idx;size:32" json:"fallbackPriorityGlobal"`
	LangySkipPermissionsModels datatypes.JSON       `gorm:"column:langySkipPermissionsModels" json:"langySkipPermissionsModels"`
	HealthStatus               string               `gorm:"column:healthStatus;index:ModelProvider_healthStatus_idx;default:'UNKNOWN';not null" json:"healthStatus"`
	CircuitOpenedAt            *DateTime            `gorm:"column:circuitOpenedAt;type:timestamp(3)" json:"circuitOpenedAt"`
	LastHealthCheckAt          *DateTime            `gorm:"column:lastHealthCheckAt;type:timestamp(3)" json:"lastHealthCheckAt"`
	DisabledAt                 *DateTime            `gorm:"column:disabledAt;type:timestamp(3)" json:"disabledAt"`
	CreatedAt                  DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime             `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	OrganizationID             string               `gorm:"column:organizationId;uniqueIndex:ModelProvider_organizationId_routingHandle_key,priority:1;index:ModelProvider_organizationId_idx;not null" json:"organizationId"`
	Scopes                     []ModelProviderScope `gorm:"foreignKey:ModelProviderID;constraint:ModelProviderScope_modelProviderId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ModelProvider) TableName() string {
	return "ModelProvider"
}

func (m *ModelProvider) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ModelProvider) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ModelProviderScope struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	ModelProviderID string         `gorm:"column:modelProviderId;uniqueIndex:ModelProviderScope_modelProviderId_scopeType_scopeId_key,priority:1;index:ModelProviderScope_modelProviderId_idx;not null" json:"modelProviderId"`
	ScopeType       string         `gorm:"column:scopeType;uniqueIndex:ModelProviderScope_modelProviderId_scopeType_scopeId_key,priority:2;index:ModelProviderScope_scopeType_scopeId_idx,priority:1;not null" json:"scopeType"`
	ScopeID         string         `gorm:"column:scopeId;uniqueIndex:ModelProviderScope_modelProviderId_scopeType_scopeId_key,priority:3;index:ModelProviderScope_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ModelProvider   *ModelProvider `gorm:"belongsTo"`
}

func (ModelProviderScope) TableName() string {
	return "ModelProviderScope"
}

func (m *ModelProviderScope) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type ModelDefaultConfig struct {
	ID             string                    `gorm:"column:id;primaryKey" json:"id"`
	Config         datatypes.JSON            `gorm:"column:config;not null" json:"config"`
	AuthorID       *string                   `gorm:"column:authorId" json:"authorId"`
	CreatedAt      DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime                  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OrganizationID string                    `gorm:"column:organizationId;index:ModelDefaultConfig_organizationId_idx;not null" json:"organizationId"`
	Scopes         []ModelDefaultConfigScope `gorm:"foreignKey:ConfigID;constraint:ModelDefaultConfigScope_configId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ModelDefaultConfig) TableName() string {
	return "ModelDefaultConfig"
}

func (m *ModelDefaultConfig) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ModelDefaultConfig) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ModelDefaultConfigScope struct {
	ID        string              `gorm:"column:id;primaryKey" json:"id"`
	ConfigID  string              `gorm:"column:configId;uniqueIndex:ModelDefaultConfigScope_configId_scopeType_scopeId_key,priority:1;index:ModelDefaultConfigScope_configId_idx;not null" json:"configId"`
	ScopeType string              `gorm:"column:scopeType;uniqueIndex:ModelDefaultConfigScope_configId_scopeType_scopeId_key,priority:2;index:ModelDefaultConfigScope_scopeType_scopeId_idx,priority:1;not null" json:"scopeType"`
	ScopeID   string              `gorm:"column:scopeId;uniqueIndex:ModelDefaultConfigScope_configId_scopeType_scopeId_key,priority:3;index:ModelDefaultConfigScope_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	CreatedAt DateTime            `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Config    *ModelDefaultConfig `gorm:"foreignKey:ConfigID;belongsTo"`
}

func (ModelDefaultConfigScope) TableName() string {
	return "ModelDefaultConfigScope"
}

func (m *ModelDefaultConfigScope) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// The organization's GitHub connection: a GitHub App installation mapped to a
// LangWatch organization. Short-lived (1h) installation access tokens are
// minted against these rows on demand, by Langy for writes (bot-authored pull
// requests, scoped to the installation's repositories) and by pull-request
// linkage for reads. No user token and no secret are stored here, the crown
// jewel is the App private key, held only in the control-plane env.
// Spec: specs/integrations/github-connection.feature.
//
// installationId is globally unique (one GitHub installation, one row); a
// single LangWatch organization may own MANY installations (different GitHub
// accounts), so the org to installation relationship is one-to-many.
type GithubInstallation struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// GitHub installation id (numeric; stored as string). Globally unique.
	InstallationID string `gorm:"column:installationId;uniqueIndex:GithubInstallation_installationId_key;not null" json:"installationId"`
	OrganizationID string `gorm:"column:organizationId;index:GithubInstallation_organizationId_idx;not null" json:"organizationId"`
	// The GitHub account (organization or user) the App is installed on.
	AccountLogin string `gorm:"column:accountLogin;not null" json:"accountLogin"`
	// "Organization" or "User".
	AccountType string `gorm:"column:accountType;not null" json:"accountType"`
	// GitHub numeric account id (stable across renames).
	AccountID string `gorm:"column:accountId;not null" json:"accountId"`
	// "all" or "selected" — whether the App can see every repo on the account
	// or only a chosen subset.
	RepositorySelection string `gorm:"column:repositorySelection;not null" json:"repositorySelection"`
	// Cached repository list [{ id, fullName }] when repositorySelection =
	// "selected". Null when "all" (the full set is resolved live at mint time).
	Repositories datatypes.JSON `gorm:"column:repositories" json:"repositories"`
	// Set when GitHub suspends the installation; cleared on unsuspend. A
	// suspended installation cannot mint tokens.
	SuspendedAt  *DateTime     `gorm:"column:suspendedAt;type:timestamp(3)" json:"suspendedAt"`
	CreatedAt    DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization *Organization `gorm:"belongsTo"`
}

func (GithubInstallation) TableName() string {
	return "GithubInstallation"
}

func (m *GithubInstallation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GithubInstallation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A pull request discovered through the organization's GitHub connection and
// mapped to the branch it was opened from, so a coding-agent session that
// knows only its branch can name the pull request it produced.
// Spec: specs/integrations/github-connection.feature.
//
// The stored title, state, draft and timestamp columns are SNAPSHOTS taken
// when the mapping was last refreshed. They are fallback labels only: the
// frontend reads live status from GitHub, so a stale row never presents
// itself as current truth.
type GithubPullRequest struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:GithubPullRequest_organizationId_repositoryHost_repositoryF_key,priority:1;index:GithubPullRequest_organizationId_repositoryHost_repositoryF_idx,priority:1;not null" json:"organizationId"`
	// The GitHub host the repository lives on, e.g. "github.com".
	RepositoryHost string `gorm:"column:repositoryHost;uniqueIndex:GithubPullRequest_organizationId_repositoryHost_repositoryF_key,priority:2;index:GithubPullRequest_organizationId_repositoryHost_repositoryF_idx,priority:2;not null" json:"repositoryHost"`
	// "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
	// session reported.
	RepositoryFullName string `gorm:"column:repositoryFullName;uniqueIndex:GithubPullRequest_organizationId_repositoryHost_repositoryF_key,priority:3;index:GithubPullRequest_organizationId_repositoryHost_repositoryF_idx,priority:3;not null" json:"repositoryFullName"`
	// The branch the pull request was opened from.
	HeadBranch string `gorm:"column:headBranch;index:GithubPullRequest_organizationId_repositoryHost_repositoryF_idx,priority:4;not null" json:"headBranch"`
	PrNumber   int    `gorm:"column:prNumber;uniqueIndex:GithubPullRequest_organizationId_repositoryHost_repositoryF_key,priority:4;size:32;not null" json:"prNumber"`
	HTMLURL    string `gorm:"column:htmlUrl;not null" json:"htmlUrl"`
	Title      string `gorm:"column:title;not null" json:"title"`
	// GitHub's own state string: "open" or "closed".
	State       string    `gorm:"column:state;not null" json:"state"`
	IsDraft     bool      `gorm:"column:isDraft;default:false;not null" json:"isDraft"`
	AuthorLogin *string   `gorm:"column:authorLogin" json:"authorLogin"`
	PrCreatedAt DateTime  `gorm:"column:prCreatedAt;type:timestamp(3);not null" json:"prCreatedAt"`
	PrClosedAt  *DateTime `gorm:"column:prClosedAt;type:timestamp(3)" json:"prClosedAt"`
	PrMergedAt  *DateTime `gorm:"column:prMergedAt;type:timestamp(3)" json:"prMergedAt"`
	// GitHub's own `updated_at` for the snapshot this row holds, and the
	// freshness guard on every write to it: a snapshot older than the stored one
	// is skipped. GitHub may deliver webhooks out of order, and without the
	// guard a late `opened` lands after `closed` and reopens a merged pull
	// request. NULL means the row predates the column, so any write is accepted.
	PrUpdatedAt *DateTime `gorm:"column:prUpdatedAt;type:timestamp(3)" json:"prUpdatedAt"`
	// When this row was first mapped to its branch.
	MappedAt DateTime `gorm:"column:mappedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"mappedAt"`
	// When the snapshot columns were last refreshed from GitHub.
	LastCheckedAt DateTime      `gorm:"column:lastCheckedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastCheckedAt"`
	CreatedAt     DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization  *Organization `gorm:"belongsTo"`
}

func (GithubPullRequest) TableName() string {
	return "GithubPullRequest"
}

func (m *GithubPullRequest) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.MappedAt.IsZero() {
		m.MappedAt = DateTime{Time: now}
	}
	if m.LastCheckedAt.IsZero() {
		m.LastCheckedAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GithubPullRequest) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The per-branch bookkeeping behind pull-request mapping: when a branch was
// last looked up on GitHub, how many pull requests it resolved to, and when it
// may be asked about again. A branch with no pull request is the common case,
// so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
// costing a GitHub call on every read.
// Spec: specs/integrations/github-connection.feature.
type GithubBranchPullRequestCheck struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:GithubBranchPullRequestCheck_organizationId_repositoryHost__key,priority:1;not null" json:"organizationId"`
	RepositoryHost string `gorm:"column:repositoryHost;uniqueIndex:GithubBranchPullRequestCheck_organizationId_repositoryHost__key,priority:2;not null" json:"repositoryHost"`
	// "owner/repo", stored LOWERCASED, matching GithubPullRequest.
	RepositoryFullName string   `gorm:"column:repositoryFullName;uniqueIndex:GithubBranchPullRequestCheck_organizationId_repositoryHost__key,priority:3;not null" json:"repositoryFullName"`
	HeadBranch         string   `gorm:"column:headBranch;uniqueIndex:GithubBranchPullRequestCheck_organizationId_repositoryHost__key,priority:4;not null" json:"headBranch"`
	LastCheckedAt      DateTime `gorm:"column:lastCheckedAt;type:timestamp(3);not null" json:"lastCheckedAt"`
	PrCount            int      `gorm:"column:prCount;size:32;default:0;not null" json:"prCount"`
	// Set while the branch resolves to no pull request; cleared once one is
	// found.
	NotFoundAt *DateTime `gorm:"column:notFoundAt;index:GithubBranchPullRequestCheck_recheckAfter_lastRequestedAt_n_idx,priority:3;type:timestamp(3)" json:"notFoundAt"`
	// The earliest instant GitHub may be asked about this branch again.
	RecheckAfter *DateTime `gorm:"column:recheckAfter;index:GithubBranchPullRequestCheck_recheckAfter_lastRequestedAt_n_idx,priority:1;type:timestamp(3)" json:"recheckAfter"`
	// Consecutive lookups that found nothing, backing off the recheck.
	Attempts int `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	// When a reader last asked about this branch, whether or not GitHub was
	// called.
	LastRequestedAt DateTime      `gorm:"column:lastRequestedAt;index:GithubBranchPullRequestCheck_recheckAfter_lastRequestedAt_n_idx,priority:2;type:timestamp(3);not null" json:"lastRequestedAt"`
	CreatedAt       DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization    *Organization `gorm:"belongsTo"`
}

func (GithubBranchPullRequestCheck) TableName() string {
	return "GithubBranchPullRequestCheck"
}

func (m *GithubBranchPullRequestCheck) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GithubBranchPullRequestCheck) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Rebuildable operational state folded from Langy's canonical event stream.
// The source of truth remains the ClickHouse event log.
type LangyConversationProjection struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	ProjectID      string `gorm:"column:projectId;uniqueIndex:LangyConversationProjection_projectId_conversationId_key,priority:1;index:LangyConversationProjection_projectId_userId_conversationId_idx,priority:1;index:LangyConversationProjection_projectId_userId_lastActivityAt_idx,priority:1;not null" json:"projectId"`
	ConversationID string `gorm:"column:conversationId;uniqueIndex:LangyConversationProjection_projectId_conversationId_key,priority:2;index:LangyConversationProjection_projectId_userId_conversationId_idx,priority:3;not null" json:"conversationId"`
	// Owner; first-writer-wins from the first message's userId.
	UserID string  `gorm:"column:userId;index:LangyConversationProjection_projectId_userId_conversationId_idx,priority:2;index:LangyConversationProjection_projectId_userId_lastActivityAt_idx,priority:2;not null" json:"userId"`
	Title  *string `gorm:"column:title" json:"title"`
	// derived | auto | user — governs auto-title precedence (user is sticky).
	TitleSource string `gorm:"column:titleSource;not null" json:"titleSource"`
	// Who started the conversation: `interactive` for the person in the panel, `run` for a
	// module acting for them with nobody at the keyboard. A row folded before the column
	// existed reads as interactive.
	Origin         *string  `gorm:"column:origin;default:'interactive';not null" json:"origin"`
	Status         string   `gorm:"column:status;not null" json:"status"`
	IsShared       bool     `gorm:"column:isShared;default:false;not null" json:"isShared"`
	SharedAt       *float64 `gorm:"column:sharedAt;type:double precision" json:"sharedAt"`
	SharedByID     *string  `gorm:"column:sharedById" json:"sharedById"`
	MessageCount   int      `gorm:"column:messageCount;size:32;default:0;not null" json:"messageCount"`
	LastActivityAt *float64 `gorm:"column:lastActivityAt;index:LangyConversationProjection_projectId_userId_lastActivityAt_idx,priority:3;type:double precision" json:"lastActivityAt"`
	// The turn in flight, or null when idle (busy guard reads this).
	CurrentTurnID *string `gorm:"column:currentTurnId" json:"currentTurnId"`
	LastError     *string `gorm:"column:lastError" json:"lastError"`
	// The model the latest accepted turn ran on; reopening seeds the picker from it.
	LastModel *string `gorm:"column:lastModel" json:"lastModel"`
	// ADR-048 shutdown-handoff resume token (opaque, worker-authored).
	PendingHandoffToken  *string `gorm:"column:pendingHandoffToken" json:"pendingHandoffToken"`
	PendingHandoffTurnID *string `gorm:"column:pendingHandoffTurnId" json:"pendingHandoffTurnId"`
	// Per-conversation HMAC key for worker stream frames; server-side only.
	RunToken   *string  `gorm:"column:runToken" json:"runToken"`
	ArchivedAt *float64 `gorm:"column:archivedAt;type:double precision" json:"archivedAt"`
	// Messages CI started on a pull request since the person last wrote.
	CiFollowUpsInARow int     `gorm:"column:ciFollowUpsInARow;size:32;default:0;not null" json:"ciFollowUpsInARow"`
	CreatedAt         float64 `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt         float64 `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	// Business time of the latest applied event.
	OccurredAt float64 `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	// Event-log acceptance time. Together with lastEventId this is the cursor.
	AcceptedAt        float64 `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string  `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string  `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (LangyConversationProjection) TableName() string {
	return "LangyConversationProjection"
}

func (m *LangyConversationProjection) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Durable idempotency receipt for one logical Langy send. A transport retry
// reuses the same requestId and receives the original conversation/turn ids
// instead of minting a second message or turn.
type LangyTurnRequest struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID      string   `gorm:"column:projectId;uniqueIndex:LangyTurnRequest_projectId_userId_requestId_key,priority:1;index:LangyTurnRequest_projectId_conversationId_idx,priority:1;not null" json:"projectId"`
	UserID         string   `gorm:"column:userId;uniqueIndex:LangyTurnRequest_projectId_userId_requestId_key,priority:2;not null" json:"userId"`
	RequestID      string   `gorm:"column:requestId;uniqueIndex:LangyTurnRequest_projectId_userId_requestId_key,priority:3;not null" json:"requestId"`
	ConversationID string   `gorm:"column:conversationId;index:LangyTurnRequest_projectId_conversationId_idx,priority:2;not null" json:"conversationId"`
	TurnID         string   `gorm:"column:turnId;not null" json:"turnId"`
	Status         string   `gorm:"column:status;not null" json:"status"`
	LeaseOwner     string   `gorm:"column:leaseOwner;not null" json:"leaseOwner"`
	LeaseExpiresAt DateTime `gorm:"column:leaseExpiresAt;type:timestamp(3);not null" json:"leaseExpiresAt"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (LangyTurnRequest) TableName() string {
	return "LangyTurnRequest"
}

func (m *LangyTurnRequest) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*LangyTurnRequest) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Postgres-authoritative one-turn admission for a conversation. PREPARING
// claims are leased so a crash before the durable start event can recover;
// COMMITTED claims are released only by the matching terminal event.
type LangyActiveTurn struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID      string   `gorm:"column:projectId;uniqueIndex:LangyActiveTurn_projectId_conversationId_key,priority:1;index:LangyActiveTurn_projectId_userId_idx,priority:1;not null" json:"projectId"`
	ConversationID string   `gorm:"column:conversationId;uniqueIndex:LangyActiveTurn_projectId_conversationId_key,priority:2;not null" json:"conversationId"`
	TurnID         string   `gorm:"column:turnId;not null" json:"turnId"`
	RequestID      string   `gorm:"column:requestId;not null" json:"requestId"`
	UserID         string   `gorm:"column:userId;index:LangyActiveTurn_projectId_userId_idx,priority:2;not null" json:"userId"`
	Status         string   `gorm:"column:status;not null" json:"status"`
	LeaseOwner     string   `gorm:"column:leaseOwner;not null" json:"leaseOwner"`
	LeaseExpiresAt DateTime `gorm:"column:leaseExpiresAt;type:timestamp(3);not null" json:"leaseExpiresAt"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (LangyActiveTurn) TableName() string {
	return "LangyActiveTurn"
}

func (m *LangyActiveTurn) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*LangyActiveTurn) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Rebuildable per-turn operational render document.
type LangyConversationTurnProjection struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	ProjectID      string `gorm:"column:projectId;uniqueIndex:LangyConversationTurnProjection_projectId_conversationId_tu_key,priority:1;index:LangyConversationTurnProjection_projectId_conversationId_up_idx,priority:1;not null" json:"projectId"`
	ConversationID string `gorm:"column:conversationId;uniqueIndex:LangyConversationTurnProjection_projectId_conversationId_tu_key,priority:2;index:LangyConversationTurnProjection_projectId_conversationId_up_idx,priority:2;not null" json:"conversationId"`
	TurnID         string `gorm:"column:turnId;uniqueIndex:LangyConversationTurnProjection_projectId_conversationId_tu_key,priority:3;not null" json:"turnId"`
	// LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
	// is a PROJECTION — one writer (the fold), rebuildable from the event log,
	// and nothing queries or indexes on it. An enum bought no selectivity here,
	// only typo protection, and charged a migration for every new status that
	// Postgres then could not remove. The one definition lives in
	// `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
	Status string `gorm:"column:status;not null" json:"status"`
	// LangyMessagePart[] — the user question that opened the turn.
	QuestionParts datatypes.JSON `gorm:"column:questionParts;not null" json:"questionParts"`
	// LangyMessagePart[] — the agent's whole final answer.
	AnswerParts datatypes.JSON `gorm:"column:answerParts;not null" json:"answerParts"`
	// LangyTurnToolCall[] — tool-call lifecycle audit in initiation order.
	ToolCalls datatypes.JSON `gorm:"column:toolCalls;not null" json:"toolCalls"`
	// LangyPlanItemData[] | null — null means the turn never kept a plan.
	Plan              datatypes.JSON `gorm:"column:plan" json:"plan"`
	Error             *string        `gorm:"column:error" json:"error"`
	StartedAt         *float64       `gorm:"column:startedAt;type:double precision" json:"startedAt"`
	EndedAt           *float64       `gorm:"column:endedAt;type:double precision" json:"endedAt"`
	CreatedAt         float64        `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt         float64        `gorm:"column:updatedAt;index:LangyConversationTurnProjection_projectId_conversationId_up_idx,priority:3;type:double precision;not null" json:"updatedAt"`
	OccurredAt        float64        `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	AcceptedAt        float64        `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string         `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string         `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (LangyConversationTurnProjection) TableName() string {
	return "LangyConversationTurnProjection"
}

func (m *LangyConversationTurnProjection) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
// checks look it up by repository and number to wake the conversation; the daily pull request
// cap counts a person's rows by day. Rebuilt by a projection replay.
type LangyPullRequest struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string `gorm:"column:projectId;not null" json:"projectId"`
	// `owner/repo`, lowercased: GitHub matches repository names without case.
	RepositoryFullName string `gorm:"column:repositoryFullName;uniqueIndex:LangyPullRequest_repositoryFullName_prNumber_key,priority:1;not null" json:"repositoryFullName"`
	PrNumber           int    `gorm:"column:prNumber;uniqueIndex:LangyPullRequest_repositoryFullName_prNumber_key,priority:2;size:32;not null" json:"prNumber"`
	ConversationID     string `gorm:"column:conversationId;not null" json:"conversationId"`
	// The conversation's owner, who opened it through Langy.
	UserID string `gorm:"column:userId;index:LangyPullRequest_userId_openedAt_idx,priority:1;not null" json:"userId"`
	// When the answer that opened it was recorded (epoch ms).
	OpenedAt float64 `gorm:"column:openedAt;index:LangyPullRequest_userId_openedAt_idx,priority:2;type:double precision;not null" json:"openedAt"`
	// The agent_responded event it came from.
	SourceEventID string `gorm:"column:sourceEventId;not null" json:"sourceEventId"`
}

func (LangyPullRequest) TableName() string {
	return "LangyPullRequest"
}

func (m *LangyPullRequest) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Rebuildable operational message fact. One source event produces one row.
type LangyMessageProjection struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID      string         `gorm:"column:projectId;uniqueIndex:LangyMessageProjection_projectId_conversationId_messageId_key,priority:1;uniqueIndex:LangyMessageProjection_projectId_sourceEventId_key,priority:1;index:LangyMessageProjection_projectId_conversationId_createdAt_idx,priority:1;not null" json:"projectId"`
	ConversationID string         `gorm:"column:conversationId;uniqueIndex:LangyMessageProjection_projectId_conversationId_messageId_key,priority:2;index:LangyMessageProjection_projectId_conversationId_createdAt_idx,priority:2;not null" json:"conversationId"`
	MessageID      string         `gorm:"column:messageId;uniqueIndex:LangyMessageProjection_projectId_conversationId_messageId_key,priority:3;not null" json:"messageId"`
	Role           string         `gorm:"column:role;not null" json:"role"`
	Parts          datatypes.JSON `gorm:"column:parts;not null" json:"parts"`
	SourceEventID  string         `gorm:"column:sourceEventId;uniqueIndex:LangyMessageProjection_projectId_sourceEventId_key,priority:2;not null" json:"sourceEventId"`
	OccurredAt     float64        `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	AcceptedAt     float64        `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	CreatedAt      float64        `gorm:"column:createdAt;index:LangyMessageProjection_projectId_conversationId_createdAt_idx,priority:3;type:double precision;not null" json:"createdAt"`
	UpdatedAt      float64        `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
}

func (LangyMessageProjection) TableName() string {
	return "LangyMessageProjection"
}

func (m *LangyMessageProjection) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

type ProjectSecret struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID      string   `gorm:"column:projectId;uniqueIndex:ProjectSecret_projectId_name_key,priority:1;index:ProjectSecret_projectId_idx;not null" json:"projectId"`
	Name           string   `gorm:"column:name;uniqueIndex:ProjectSecret_projectId_name_key,priority:2;not null" json:"name"`
	EncryptedValue string   `gorm:"column:encryptedValue;not null" json:"encryptedValue"`
	CreatedByID    string   `gorm:"column:createdById;not null" json:"createdById"`
	UpdatedByID    string   `gorm:"column:updatedById;not null" json:"updatedById"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project        *Project `gorm:"belongsTo"`
	CreatedBy      *User    `gorm:"foreignKey:CreatedByID;belongsTo"`
	UpdatedBy      *User    `gorm:"foreignKey:UpdatedByID;belongsTo"`
}

func (ProjectSecret) TableName() string {
	return "ProjectSecret"
}

func (m *ProjectSecret) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ProjectSecret) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TriggerSent struct {
	ID              string       `gorm:"column:id;primaryKey;index:TriggerSent_projectId_triggerId_createdAt_id_idx,priority:4" json:"id"`
	TriggerID       string       `gorm:"column:triggerId;uniqueIndex:TriggerSent_triggerId_traceId_key,priority:1;index:TriggerSent_triggerId_idx;index:TriggerSent_projectId_triggerId_createdAt_id_idx,priority:2;not null" json:"triggerId"`
	TraceID         *string      `gorm:"column:traceId;uniqueIndex:TriggerSent_triggerId_traceId_key,priority:2" json:"traceId"`
	CustomGraphID   *string      `gorm:"column:customGraphId;index:TriggerSent_customGraphId_idx" json:"customGraphId"`
	ProjectID       string       `gorm:"column:projectId;index:TriggerSent_projectId_idx;index:TriggerSent_projectId_triggerId_createdAt_id_idx,priority:1;not null" json:"projectId"`
	ResolvedAt      *DateTime    `gorm:"column:resolvedAt;index:TriggerSent_resolvedAt_idx;type:timestamp(3)" json:"resolvedAt"`
	OpenIncidentKey *string      `gorm:"column:openIncidentKey;uniqueIndex:TriggerSent_openIncidentKey_key" json:"openIncidentKey"`
	CreatedAt       DateTime     `gorm:"column:createdAt;index:TriggerSent_projectId_triggerId_createdAt_id_idx,priority:3;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime     `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CustomGraph     *CustomGraph `gorm:"belongsTo"`
	Project         *Project     `gorm:"belongsTo"`
	Trigger         *Trigger     `gorm:"belongsTo"`
}

func (TriggerSent) TableName() string {
	return "TriggerSent"
}

func (m *TriggerSent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TriggerSent) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TriggerLatestEvaluation struct {
	TriggerID         string   `gorm:"column:triggerId;primaryKey" json:"triggerId"`
	ProjectID         string   `gorm:"column:projectId;index:TriggerLatestEvaluation_projectId_idx;not null" json:"projectId"`
	EvaluatedAt       DateTime `gorm:"column:evaluatedAt;type:timestamptz(3);not null" json:"evaluatedAt"`
	Verdict           string   `gorm:"column:verdict;not null" json:"verdict"`
	ObservedValue     *float64 `gorm:"column:observedValue;type:double precision" json:"observedValue"`
	Threshold         *float64 `gorm:"column:threshold;type:double precision" json:"threshold"`
	Operator          *string  `gorm:"column:operator" json:"operator"`
	TimePeriodMinutes *int     `gorm:"column:timePeriodMinutes;size:32" json:"timePeriodMinutes"`
	SkipCode          *string  `gorm:"column:skipCode" json:"skipCode"`
	UpdatedAt         DateTime `gorm:"column:updatedAt;type:timestamptz(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Trigger           *Trigger `gorm:"belongsTo"`
}

func (TriggerLatestEvaluation) TableName() string {
	return "TriggerLatestEvaluation"
}

func (m *TriggerLatestEvaluation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TriggerLatestEvaluation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EmailSuppression struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string   `gorm:"column:projectId;uniqueIndex:EmailSuppression_projectId_email_triggerId_key,priority:1;index:EmailSuppression_projectId_idx;not null" json:"projectId"`
	Email     string   `gorm:"column:email;uniqueIndex:EmailSuppression_projectId_email_triggerId_key,priority:2;not null" json:"email"`
	TriggerID *string  `gorm:"column:triggerId;uniqueIndex:EmailSuppression_projectId_email_triggerId_key,priority:3" json:"triggerId"`
	Reason    *string  `gorm:"column:reason;default:'unsubscribe';not null" json:"reason"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Project   *Project `gorm:"belongsTo"`
}

func (EmailSuppression) TableName() string {
	return "EmailSuppression"
}

func (m *EmailSuppression) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type AnnotationScore struct {
	ID                    string                  `gorm:"column:id;primaryKey" json:"id"`
	Name                  string                  `gorm:"column:name;not null" json:"name"`
	ProjectID             string                  `gorm:"column:projectId;index:AnnotationScore_projectId_idx;not null" json:"projectId"`
	CreatedAt             DateTime                `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime                `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt             *DateTime               `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	Description           *string                 `gorm:"column:description" json:"description"`
	Active                *bool                   `gorm:"column:active;default:true;not null" json:"active"`
	DataType              *string                 `gorm:"column:dataType" json:"dataType"`
	Options               datatypes.JSON          `gorm:"column:options" json:"options"`
	DefaultValue          datatypes.JSON          `gorm:"column:defaultValue" json:"defaultValue"`
	Global                bool                    `gorm:"column:global;default:false;not null" json:"global"`
	Project               *Project                `gorm:"belongsTo"`
	AnnotationQueueScores []AnnotationQueueScores `gorm:"foreignKey:AnnotationScoreID;constraint:AnnotationQueueScores_annotationScoreId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (AnnotationScore) TableName() string {
	return "AnnotationScore"
}

func (m *AnnotationScore) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AnnotationScore) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AnnotationQueue struct {
	ID                    string                   `gorm:"column:id;primaryKey" json:"id"`
	Name                  string                   `gorm:"column:name;not null" json:"name"`
	Slug                  string                   `gorm:"column:slug;uniqueIndex:AnnotationQueue_projectId_slug_key,priority:2;not null" json:"slug"`
	ProjectID             string                   `gorm:"column:projectId;uniqueIndex:AnnotationQueue_projectId_slug_key,priority:1;index:AnnotationQueue_projectId_idx;not null" json:"projectId"`
	Description           *string                  `gorm:"column:description" json:"description"`
	CreatedAt             DateTime                 `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime                 `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project               *Project                 `gorm:"belongsTo"`
	Members               []AnnotationQueueMembers `gorm:"foreignKey:AnnotationQueueID;constraint:AnnotationQueueMembers_annotationQueueId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AnnotationQueueScores []AnnotationQueueScores  `gorm:"foreignKey:AnnotationQueueID;constraint:AnnotationQueueScores_annotationQueueId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AnnotationQueueItems  []AnnotationQueueItem    `gorm:"foreignKey:AnnotationQueueID;constraint:AnnotationQueueItem_annotationQueueId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (AnnotationQueue) TableName() string {
	return "AnnotationQueue"
}

func (m *AnnotationQueue) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AnnotationQueue) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AnnotationQueueMembers struct {
	AnnotationQueueID string           `gorm:"column:annotationQueueId;primaryKey;index:AnnotationQueueMembers_annotationQueueId_idx" json:"annotationQueueId"`
	UserID            string           `gorm:"column:userId;primaryKey;index:AnnotationQueueMembers_userId_idx" json:"userId"`
	User              *User            `gorm:"belongsTo"`
	AnnotationQueue   *AnnotationQueue `gorm:"belongsTo"`
}

func (AnnotationQueueMembers) TableName() string {
	return "AnnotationQueueMembers"
}

type AnnotationQueueScores struct {
	AnnotationQueueID string           `gorm:"column:annotationQueueId;primaryKey;index:AnnotationQueueScores_annotationQueueId_idx" json:"annotationQueueId"`
	AnnotationScoreID string           `gorm:"column:annotationScoreId;primaryKey;index:AnnotationQueueScores_annotationScoreId_idx" json:"annotationScoreId"`
	AnnotationScore   *AnnotationScore `gorm:"belongsTo"`
	AnnotationQueue   *AnnotationQueue `gorm:"belongsTo"`
}

func (AnnotationQueueScores) TableName() string {
	return "AnnotationQueueScores"
}

type AnnotationQueueItem struct {
	ID                string    `gorm:"column:id;primaryKey" json:"id"`
	AnnotationQueueID *string   `gorm:"column:annotationQueueId;uniqueIndex:AnnotationQueueItem_traceId_annotationQueueId_projectId_key,priority:2;index:AnnotationQueueItem_annotationQueueId_idx;index:AnnotationQueueItem_annotationQueueId_doneAt_idx,priority:1" json:"annotationQueueId"`
	UserID            *string   `gorm:"column:userId;uniqueIndex:AnnotationQueueItem_traceId_userId_projectId_key,priority:2;index:AnnotationQueueItem_userId_idx;index:AnnotationQueueItem_projectId_userId_doneAt_idx,priority:2" json:"userId"`
	CreatedAt         DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByUserID   *string   `gorm:"column:createdByUserId" json:"createdByUserId"`
	TraceID           string    `gorm:"column:traceId;uniqueIndex:AnnotationQueueItem_traceId_userId_projectId_key,priority:1;uniqueIndex:AnnotationQueueItem_traceId_annotationQueueId_projectId_key,priority:1;not null" json:"traceId"`
	ProjectID         string    `gorm:"column:projectId;uniqueIndex:AnnotationQueueItem_traceId_userId_projectId_key,priority:3;uniqueIndex:AnnotationQueueItem_traceId_annotationQueueId_projectId_key,priority:3;index:AnnotationQueueItem_projectId_idx;index:AnnotationQueueItem_projectId_doneAt_idx,priority:1;index:AnnotationQueueItem_projectId_userId_doneAt_idx,priority:1;not null" json:"projectId"`
	DoneAt            *DateTime `gorm:"column:doneAt;index:AnnotationQueueItem_projectId_doneAt_idx,priority:2;index:AnnotationQueueItem_projectId_userId_doneAt_idx,priority:3;index:AnnotationQueueItem_annotationQueueId_doneAt_idx,priority:2;type:timestamp(3)" json:"doneAt"`
	// Set while the annotator wants this item's trace added to a dataset once
	// the queue walk is finished. Persisted rather than kept in the browser so
	// the selection survives a refresh and stays visible to teammates.
	MarkedForDatasetAt *DateTime        `gorm:"column:markedForDatasetAt;type:timestamp(3)" json:"markedForDatasetAt"`
	AnnotationQueue    *AnnotationQueue `gorm:"belongsTo"`
	User               *User            `gorm:"foreignKey:UserID;belongsTo"`
	CreatedByUser      *User            `gorm:"foreignKey:CreatedByUserID;belongsTo"`
	Project            *Project         `gorm:"belongsTo"`
}

func (AnnotationQueueItem) TableName() string {
	return "AnnotationQueueItem"
}

func (m *AnnotationQueueItem) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AnnotationQueueItem) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ShareLink struct {
	ID           string    `gorm:"column:id;primaryKey" json:"id"`
	Token        string    `gorm:"column:token;uniqueIndex:ShareLink_token_key;not null" json:"token"`
	ResourceType string    `gorm:"column:resourceType;index:ShareLink_projectId_resourceType_resourceId_idx,priority:2;not null" json:"resourceType"`
	ResourceID   string    `gorm:"column:resourceId;index:ShareLink_projectId_resourceType_resourceId_idx,priority:3;not null" json:"resourceId"`
	ThreadID     *string   `gorm:"column:threadId" json:"threadId"`
	ProjectID    string    `gorm:"column:projectId;index:ShareLink_projectId_resourceType_resourceId_idx,priority:1;not null" json:"projectId"`
	UserID       *string   `gorm:"column:userId;index:ShareLink_userId_idx" json:"userId"`
	Visibility   string    `gorm:"column:visibility;default:'PUBLIC';not null" json:"visibility"`
	ExpiresAt    *DateTime `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	MaxViews     *int      `gorm:"column:maxViews;size:32" json:"maxViews"`
	ViewCount    int       `gorm:"column:viewCount;size:32;default:0;not null" json:"viewCount"`
	CreatedAt    DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project      *Project  `gorm:"belongsTo"`
	User         *User     `gorm:"belongsTo"`
}

func (ShareLink) TableName() string {
	return "ShareLink"
}

func (m *ShareLink) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ShareLink) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PinnedTrace struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string   `gorm:"column:projectId;uniqueIndex:PinnedTrace_projectId_traceId_key,priority:1;index:PinnedTrace_projectId_idx;not null" json:"projectId"`
	TraceID   string   `gorm:"column:traceId;uniqueIndex:PinnedTrace_projectId_traceId_key,priority:2;not null" json:"traceId"`
	UserID    *string  `gorm:"column:userId" json:"userId"`
	Source    string   `gorm:"column:source;default:'manual';not null" json:"source"`
	Reason    *string  `gorm:"column:reason" json:"reason"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Project   *Project `gorm:"belongsTo"`
	User      *User    `gorm:"belongsTo"`
}

func (PinnedTrace) TableName() string {
	return "PinnedTrace"
}

func (m *PinnedTrace) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type TraceEditOverlay struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID   string         `gorm:"column:projectId;uniqueIndex:TraceEditOverlay_projectId_traceId_key,priority:1;index:TraceEditOverlay_projectId_idx;not null" json:"projectId"`
	TraceID     string         `gorm:"column:traceId;uniqueIndex:TraceEditOverlay_projectId_traceId_key,priority:2;not null" json:"traceId"`
	Patch       datatypes.JSON `gorm:"column:patch;not null" json:"patch"`
	CreatedByID *string        `gorm:"column:createdById" json:"createdById"`
	UpdatedByID *string        `gorm:"column:updatedById" json:"updatedById"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project     *Project       `gorm:"belongsTo"`
	CreatedBy   *User          `gorm:"foreignKey:CreatedByID;belongsTo"`
	UpdatedBy   *User          `gorm:"foreignKey:UpdatedByID;belongsTo"`
}

func (TraceEditOverlay) TableName() string {
	return "TraceEditOverlay"
}

func (m *TraceEditOverlay) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TraceEditOverlay) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TraceIngestSourceBilling struct {
	OrganizationID string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	SourceType     string   `gorm:"column:sourceType;primaryKey" json:"sourceType"`
	Billed         bool     `gorm:"column:billed;not null" json:"billed"`
	RecordedAt     DateTime `gorm:"column:recordedAt;type:timestamp(3);not null" json:"recordedAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (TraceIngestSourceBilling) TableName() string {
	return "TraceIngestSourceBilling"
}

func (m *TraceIngestSourceBilling) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TraceIngestSourceBilling) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RetentionPolicy struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string   `gorm:"column:organizationId;index:RetentionPolicy_organizationId_idx;not null" json:"organizationId"`
	ScopeType      string   `gorm:"column:scopeType;uniqueIndex:RetentionPolicy_scopeType_scopeId_category_key,priority:1;index:RetentionPolicy_scopeType_scopeId_idx,priority:1;not null" json:"scopeType"`
	ScopeID        string   `gorm:"column:scopeId;uniqueIndex:RetentionPolicy_scopeType_scopeId_category_key,priority:2;index:RetentionPolicy_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	Category       string   `gorm:"column:category;uniqueIndex:RetentionPolicy_scopeType_scopeId_category_key,priority:3;not null" json:"category"`
	RetentionDays  int      `gorm:"column:retentionDays;size:32;not null" json:"retentionDays"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (RetentionPolicy) TableName() string {
	return "RetentionPolicy"
}

func (m *RetentionPolicy) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*RetentionPolicy) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RetentionApplyStatus struct {
	ID                string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID    string   `gorm:"column:organizationId;uniqueIndex:RetentionApplyStatus_organizationId_scopeType_scopeId_key,priority:1;not null" json:"organizationId"`
	ScopeType         string   `gorm:"column:scopeType;uniqueIndex:RetentionApplyStatus_organizationId_scopeType_scopeId_key,priority:2;not null" json:"scopeType"`
	ScopeID           string   `gorm:"column:scopeId;uniqueIndex:RetentionApplyStatus_organizationId_scopeType_scopeId_key,priority:3;not null" json:"scopeId"`
	ChangedAt         *float64 `gorm:"column:changedAt;type:double precision" json:"changedAt"`
	AppliedAt         *float64 `gorm:"column:appliedAt;type:double precision" json:"appliedAt"`
	RaisedAt          *float64 `gorm:"column:raisedAt;type:double precision" json:"raisedAt"`
	ShortenedAt       *float64 `gorm:"column:shortenedAt;type:double precision" json:"shortenedAt"`
	CreatedAt         float64  `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt         float64  `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	OccurredAt        float64  `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	AcceptedAt        float64  `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string   `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string   `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (RetentionApplyStatus) TableName() string {
	return "RetentionApplyStatus"
}

type DataPrivacyPolicy struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string         `gorm:"column:organizationId;index:DataPrivacyPolicy_organizationId_idx;not null" json:"organizationId"`
	ScopeType      string         `gorm:"column:scopeType;uniqueIndex:DataPrivacyPolicy_scopeType_scopeId_personalOnly_key,priority:1;index:DataPrivacyPolicy_scopeType_scopeId_idx,priority:1;not null" json:"scopeType"`
	ScopeID        string         `gorm:"column:scopeId;uniqueIndex:DataPrivacyPolicy_scopeType_scopeId_personalOnly_key,priority:2;index:DataPrivacyPolicy_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	PersonalOnly   bool           `gorm:"column:personalOnly;uniqueIndex:DataPrivacyPolicy_scopeType_scopeId_personalOnly_key,priority:3;default:false;not null" json:"personalOnly"`
	Config         datatypes.JSON `gorm:"column:config;not null" json:"config"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (DataPrivacyPolicy) TableName() string {
	return "DataPrivacyPolicy"
}

func (m *DataPrivacyPolicy) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DataPrivacyPolicy) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type InstantEvalJudgeProject struct {
	ProjectID      string   `gorm:"column:projectId;primaryKey" json:"projectId"`
	OrganizationID string   `gorm:"column:organizationId;not null" json:"organizationId"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
}

func (InstantEvalJudgeProject) TableName() string {
	return "InstantEvalJudgeProject"
}

type InstantEvalJudgeUsageBilling struct {
	OrganizationID string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	UsageBilled    bool     `gorm:"column:usageBilled;not null" json:"usageBilled"`
	OccurredAt     DateTime `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	FromCatchUp    bool     `gorm:"column:fromCatchUp;not null" json:"fromCatchUp"`
}

func (InstantEvalJudgeUsageBilling) TableName() string {
	return "InstantEvalJudgeUsageBilling"
}

type InstantEvalJudgeSpend struct {
	OrganizationID string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	RequestID      string   `gorm:"column:requestId;primaryKey" json:"requestId"`
	SpendNanoUsd   int64    `gorm:"column:spendNanoUsd;not null" json:"spendNanoUsd"`
	OccurredAt     DateTime `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
}

func (InstantEvalJudgeSpend) TableName() string {
	return "InstantEvalJudgeSpend"
}

type CustomLLMModelCost struct {
	ID                          string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID              string   `gorm:"column:organizationId;index:CustomLLMModelCost_organizationId_idx;not null" json:"organizationId"`
	ScopeType                   string   `gorm:"column:scopeType;index:CustomLLMModelCost_scopeType_scopeId_idx,priority:1;default:'PROJECT';not null" json:"scopeType"`
	ScopeID                     string   `gorm:"column:scopeId;index:CustomLLMModelCost_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	ProjectID                   *string  `gorm:"column:projectId;index:CustomLLMModelCost_projectId_idx" json:"projectId"`
	Model                       string   `gorm:"column:model;not null" json:"model"`
	Regex                       string   `gorm:"column:regex;not null" json:"regex"`
	InputCostPerToken           *float64 `gorm:"column:inputCostPerToken;type:double precision" json:"inputCostPerToken"`
	OutputCostPerToken          *float64 `gorm:"column:outputCostPerToken;type:double precision" json:"outputCostPerToken"`
	CacheReadCostPerToken       *float64 `gorm:"column:cacheReadCostPerToken;type:double precision" json:"cacheReadCostPerToken"`
	CacheCreationCostPerToken   *float64 `gorm:"column:cacheCreationCostPerToken;type:double precision" json:"cacheCreationCostPerToken"`
	CacheCreation1hCostPerToken *float64 `gorm:"column:cacheCreation1hCostPerToken;type:double precision" json:"cacheCreation1hCostPerToken"`
	CreatedAt                   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                   DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (CustomLLMModelCost) TableName() string {
	return "CustomLLMModelCost"
}

func (m *CustomLLMModelCost) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CustomLLMModelCost) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Workflow struct {
	ID                   string            `gorm:"column:id;primaryKey" json:"id"`
	ProjectID            string            `gorm:"column:projectId;index:Workflow_projectId_idx;not null" json:"projectId"`
	Name                 string            `gorm:"column:name;not null" json:"name"`
	Icon                 string            `gorm:"column:icon;not null" json:"icon"`
	Description          string            `gorm:"column:description;not null" json:"description"`
	CreatedAt            DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime          `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	LatestVersionID      *string           `gorm:"column:latestVersionId;index:Workflow_latestVersionId_idx" json:"latestVersionId"`
	CurrentVersionID     *string           `gorm:"column:currentVersionId;index:Workflow_currentVersionId_idx" json:"currentVersionId"`
	PublishedID          *string           `gorm:"column:publishedId" json:"publishedId"`
	PublishedByID        *string           `gorm:"column:publishedById;index:Workflow_publishedById_idx" json:"publishedById"`
	ArchivedAt           *DateTime         `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	IsEvaluator          bool              `gorm:"column:isEvaluator;default:false;not null" json:"isEvaluator"`
	IsComponent          bool              `gorm:"column:isComponent;default:false;not null" json:"isComponent"`
	CopiedFromWorkflowID *string           `gorm:"column:copiedFromWorkflowId;index:Workflow_copiedFromWorkflowId_idx" json:"copiedFromWorkflowId"`
	Project              *Project          `gorm:"belongsTo"`
	LatestVersion        *WorkflowVersion  `gorm:"foreignKey:LatestVersionID;belongsTo"`
	CurrentVersion       *WorkflowVersion  `gorm:"foreignKey:CurrentVersionID;belongsTo"`
	PublishedBy          *User             `gorm:"foreignKey:PublishedByID;belongsTo"`
	CopiedFrom           *Workflow         `gorm:"foreignKey:CopiedFromWorkflowID;belongsTo"`
	Versions             []WorkflowVersion `gorm:"foreignKey:WorkflowID;constraint:WorkflowVersion_workflowId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Experiments          []Experiment      `gorm:"foreignKey:WorkflowID;constraint:Experiment_workflowId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	CopiedWorkflows      []Workflow        `gorm:"foreignKey:CopiedFromWorkflowID;constraint:Workflow_copiedFromWorkflowId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Agents               []Agent           `gorm:"foreignKey:WorkflowID;constraint:Agent_workflowId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Evaluators           []Evaluator       `gorm:"foreignKey:WorkflowID;constraint:Evaluator_workflowId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Workflow) TableName() string {
	return "Workflow"
}

func (m *Workflow) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Workflow) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkflowVersion struct {
	ID                       string            `gorm:"column:id;primaryKey" json:"id"`
	Version                  string            `gorm:"column:version;uniqueIndex:WorkflowVersion_workflowId_version_key,priority:2;not null" json:"version"`
	CommitMessage            string            `gorm:"column:commitMessage;not null" json:"commitMessage"`
	AuthorID                 string            `gorm:"column:authorId;index:WorkflowVersion_authorId_idx;not null" json:"authorId"`
	ProjectID                string            `gorm:"column:projectId;index:WorkflowVersion_projectId_idx;not null" json:"projectId"`
	WorkflowID               string            `gorm:"column:workflowId;uniqueIndex:WorkflowVersion_workflowId_version_key,priority:1;index:WorkflowVersion_workflowId_idx;not null" json:"workflowId"`
	ParentID                 *string           `gorm:"column:parentId;index:WorkflowVersion_parentId_idx" json:"parentId"`
	AutoSaved                bool              `gorm:"column:autoSaved;default:false;not null" json:"autoSaved"`
	CreatedAt                DateTime          `gorm:"column:createdAt;index:WorkflowVersion_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime          `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Dsl                      datatypes.JSON    `gorm:"column:dsl;not null" json:"dsl"`
	Author                   *User             `gorm:"foreignKey:AuthorID;belongsTo"`
	Project                  *Project          `gorm:"belongsTo"`
	Workflow                 *Workflow         `gorm:"belongsTo"`
	Parent                   *WorkflowVersion  `gorm:"foreignKey:ParentID;belongsTo"`
	Children                 []WorkflowVersion `gorm:"foreignKey:ParentID;constraint:WorkflowVersion_parentId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	WorkflowAsLatestVersion  []Workflow        `gorm:"foreignKey:LatestVersionID;constraint:Workflow_latestVersionId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	WorkflowAsCurrentVersion []Workflow        `gorm:"foreignKey:CurrentVersionID;constraint:Workflow_currentVersionId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (WorkflowVersion) TableName() string {
	return "WorkflowVersion"
}

func (m *WorkflowVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkflowVersion) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AuditLog struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime       `gorm:"column:createdAt;index:AuditLog_createdAt_idx;index:AuditLog_organizationId_createdAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UserID         *string        `gorm:"column:userId;index:AuditLog_userId_idx" json:"userId"`
	ProjectID      *string        `gorm:"column:projectId;index:AuditLog_projectId_idx" json:"projectId"`
	OrganizationID *string        `gorm:"column:organizationId;index:AuditLog_organizationId_createdAt_idx,priority:1" json:"organizationId"`
	Action         string         `gorm:"column:action;index:AuditLog_action_idx;not null" json:"action"`
	Args           datatypes.JSON `gorm:"column:args" json:"args"`
	Error          *string        `gorm:"column:error" json:"error"`
	IPAddress      *string        `gorm:"column:ipAddress;index:AuditLog_ipAddress_idx" json:"ipAddress"`
	UserAgent      *string        `gorm:"column:userAgent" json:"userAgent"`
	Metadata       datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	// WHO REALLY DID IT, when that is not `userId`.
	//
	// Under an impersonation the session's `user.id` is the SUBJECT — the
	// customer whose access an operator is borrowing — so a row naming only
	// `userId` filed the operator's act against the customer. An operator
	// granting a permanent way into somebody's organization appeared, in the
	// only durable record of it, to be that somebody.
	//
	// Null on an ordinary request, where the two would be the same person and
	// saying so twice adds nothing. ADR-118 / D06: every decision names both.
	ActorUserID *string        `gorm:"column:actorUserId;index:AuditLog_actorUserId_idx" json:"actorUserId"`
	TargetKind  *string        `gorm:"column:targetKind;index:AuditLog_targetKind_targetId_idx,priority:1" json:"targetKind"`
	TargetID    *string        `gorm:"column:targetId;index:AuditLog_targetKind_targetId_idx,priority:2" json:"targetId"`
	Before      datatypes.JSON `gorm:"column:before" json:"before"`
	After       datatypes.JSON `gorm:"column:after" json:"after"`
	// A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2).
	IdempotencyKey *string `gorm:"column:idempotencyKey;uniqueIndex:AuditLog_idempotencyKey_key" json:"idempotencyKey"`
}

func (AuditLog) TableName() string {
	return "AuditLog"
}

func (m *AuditLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type LlmPromptConfig struct {
	ID                 string                   `gorm:"column:id;primaryKey" json:"id"`
	Handle             *string                  `gorm:"column:handle;uniqueIndex:LlmPromptConfig_handle_key" json:"handle"`
	Name               string                   `gorm:"column:name;not null" json:"name"`
	ProjectID          string                   `gorm:"column:projectId;index:LlmPromptConfig_projectId_idx;not null" json:"projectId"`
	CreatedAt          DateTime                 `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime                 `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt          *DateTime                `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	OrganizationID     string                   `gorm:"column:organizationId;index:LlmPromptConfig_organizationId_scope_idx,priority:1;not null" json:"organizationId"`
	Scope              string                   `gorm:"column:scope;index:LlmPromptConfig_organizationId_scope_idx,priority:2;default:'PROJECT';not null" json:"scope"`
	CopiedFromPromptID *string                  `gorm:"column:copiedFromPromptId;index:LlmPromptConfig_copiedFromPromptId_idx" json:"copiedFromPromptId"`
	Project            *Project                 `gorm:"belongsTo"`
	CopiedFrom         *LlmPromptConfig         `gorm:"foreignKey:CopiedFromPromptID;belongsTo"`
	CopiedPrompts      []LlmPromptConfig        `gorm:"foreignKey:CopiedFromPromptID;constraint:LlmPromptConfig_copiedFromPromptId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Versions           []LlmPromptConfigVersion `gorm:"foreignKey:ConfigID;constraint:LlmPromptConfigVersion_configId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	VersionTags        []PromptTagAssignment    `gorm:"foreignKey:ConfigID;constraint:PromptTagAssignment_configId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (LlmPromptConfig) TableName() string {
	return "LlmPromptConfig"
}

func (m *LlmPromptConfig) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*LlmPromptConfig) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type LlmPromptConfigVersion struct {
	ID                string                `gorm:"column:id;primaryKey" json:"id"`
	Version           int                   `gorm:"column:version;uniqueIndex:LlmPromptConfigVersion_configId_version_key,priority:2;size:32;not null" json:"version"`
	CommitMessage     *string               `gorm:"column:commitMessage" json:"commitMessage"`
	AuthorID          *string               `gorm:"column:authorId;index:LlmPromptConfigVersion_authorId_idx" json:"authorId"`
	ConfigID          string                `gorm:"column:configId;uniqueIndex:LlmPromptConfigVersion_configId_version_key,priority:1;index:LlmPromptConfigVersion_configId_idx;not null" json:"configId"`
	ConfigData        datatypes.JSON        `gorm:"column:config;not null" json:"config"`
	SchemaVersion     string                `gorm:"column:schemaVersion;not null" json:"schemaVersion"`
	RuntimeParameters datatypes.JSON        `gorm:"column:runtimeParameters;default:'{}';not null" json:"runtimeParameters"`
	CreatedAt         DateTime              `gorm:"column:createdAt;index:LlmPromptConfigVersion_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ProjectID         string                `gorm:"column:projectId;not null" json:"projectId"`
	Author            *User                 `gorm:"foreignKey:AuthorID;belongsTo"`
	Config            *LlmPromptConfig      `gorm:"foreignKey:ConfigID;belongsTo"`
	VersionTags       []PromptTagAssignment `gorm:"foreignKey:VersionID;constraint:PromptTagAssignment_versionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (LlmPromptConfigVersion) TableName() string {
	return "LlmPromptConfigVersion"
}

func (m *LlmPromptConfigVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type PromptTagAssignment struct {
	ID          string                  `gorm:"column:id;primaryKey" json:"id"`
	ConfigID    string                  `gorm:"column:configId;uniqueIndex:PromptTagAssignment_configId_tagId_key,priority:1;index:PromptTagAssignment_configId_idx;not null" json:"configId"`
	VersionID   string                  `gorm:"column:versionId;index:PromptTagAssignment_versionId_idx;not null" json:"versionId"`
	TagID       string                  `gorm:"column:tagId;uniqueIndex:PromptTagAssignment_configId_tagId_key,priority:2;index:PromptTagAssignment_tagId_idx;not null" json:"tagId"`
	ProjectID   string                  `gorm:"column:projectId;index:PromptTagAssignment_projectId_idx;not null" json:"projectId"`
	CreatedAt   DateTime                `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime                `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID *string                 `gorm:"column:createdById" json:"createdById"`
	UpdatedByID *string                 `gorm:"column:updatedById" json:"updatedById"`
	Config      *LlmPromptConfig        `gorm:"foreignKey:ConfigID;belongsTo"`
	Version     *LlmPromptConfigVersion `gorm:"foreignKey:VersionID;belongsTo"`
	PromptTag   *PromptTag              `gorm:"foreignKey:TagID;belongsTo"`
	CreatedBy   *User                   `gorm:"foreignKey:CreatedByID;belongsTo"`
	UpdatedBy   *User                   `gorm:"foreignKey:UpdatedByID;belongsTo"`
}

func (PromptTagAssignment) TableName() string {
	return "PromptTagAssignment"
}

func (m *PromptTagAssignment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PromptTagAssignment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PromptTag struct {
	ID             string                `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string                `gorm:"column:organizationId;uniqueIndex:PromptTag_organizationId_name_key,priority:1;index:PromptTag_organizationId_idx;not null" json:"organizationId"`
	Name           string                `gorm:"column:name;uniqueIndex:PromptTag_organizationId_name_key,priority:2;not null" json:"name"`
	CreatedAt      DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime              `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID    *string               `gorm:"column:createdById" json:"createdById"`
	UpdatedByID    *string               `gorm:"column:updatedById" json:"updatedById"`
	Organization   *Organization         `gorm:"belongsTo"`
	CreatedBy      *User                 `gorm:"foreignKey:CreatedByID;belongsTo"`
	UpdatedBy      *User                 `gorm:"foreignKey:UpdatedByID;belongsTo"`
	Assignments    []PromptTagAssignment `gorm:"foreignKey:TagID;constraint:PromptTagAssignment_tagId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (PromptTag) TableName() string {
	return "PromptTag"
}

func (m *PromptTag) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PromptTag) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Analytics struct {
	ID           string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID    string         `gorm:"column:projectId;index:Analytics_createdAt_key_projectId_numericValue_idx,priority:3;index:Analytics_projectId_idx;not null" json:"projectId"`
	Key          string         `gorm:"column:key;index:Analytics_createdAt_key_projectId_numericValue_idx,priority:2;not null" json:"key"`
	Value        datatypes.JSON `gorm:"column:value" json:"value"`
	NumericValue *float64       `gorm:"column:numericValue;index:Analytics_createdAt_key_projectId_numericValue_idx,priority:4;type:double precision" json:"numericValue"`
	StringValue  *string        `gorm:"column:stringValue" json:"stringValue"`
	BoolValue    *bool          `gorm:"column:boolValue" json:"boolValue"`
	CreatedAt    DateTime       `gorm:"column:createdAt;index:Analytics_createdAt_key_projectId_numericValue_idx,priority:1;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Project      *Project       `gorm:"belongsTo"`
}

func (Analytics) TableName() string {
	return "Analytics"
}

func (m *Analytics) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type CustomRole struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string         `gorm:"column:organizationId;uniqueIndex:CustomRole_organizationId_name_key,priority:1;index:CustomRole_organizationId_idx;not null" json:"organizationId"`
	Name           string         `gorm:"column:name;uniqueIndex:CustomRole_organizationId_name_key,priority:2;not null" json:"name"`
	Description    *string        `gorm:"column:description" json:"description"`
	Permissions    datatypes.JSON `gorm:"column:permissions;not null" json:"permissions"`
	Kind           *string        `gorm:"column:kind;default:'custom';not null" json:"kind"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization  `gorm:"belongsTo"`
	AssignedUsers  []TeamUser     `gorm:"foreignKey:AssignedRoleID;constraint:TeamUser_assignedRoleId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	RoleBindings   []RoleBinding  `gorm:"foreignKey:CustomRoleID;constraint:RoleBinding_customRoleId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (CustomRole) TableName() string {
	return "CustomRole"
}

func (m *CustomRole) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CustomRole) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Notification struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID *string        `gorm:"column:organizationId;index:Notification_organizationId_idx" json:"organizationId"`
	ProjectID      *string        `gorm:"column:projectId;index:Notification_projectId_idx" json:"projectId"`
	Metadata       datatypes.JSON `gorm:"column:metadata;not null" json:"metadata"`
	CreatedAt      DateTime       `gorm:"column:createdAt;index:Notification_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	SentAt         DateTime       `gorm:"column:sentAt;index:Notification_sentAt_idx;type:timestamp(3);not null" json:"sentAt"`
	Organization   *Organization  `gorm:"belongsTo"`
	Project        *Project       `gorm:"belongsTo"`
}

func (Notification) TableName() string {
	return "Notification"
}

func (m *Notification) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Notification) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Agent struct {
	ID                string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID         string         `gorm:"column:projectId;uniqueIndex:Agent_projectId_identityKey_key,priority:1;index:Agent_projectId_idx;index:Agent_projectId_lastSeenAt_idx,priority:1;not null" json:"projectId"`
	Name              string         `gorm:"column:name;not null" json:"name"`
	Type              string         `gorm:"column:type;not null" json:"type"`
	Config            datatypes.JSON `gorm:"column:config;not null" json:"config"`
	WorkflowID        *string        `gorm:"column:workflowId;index:Agent_workflowId_idx" json:"workflowId"`
	CopiedFromAgentID *string        `gorm:"column:copiedFromAgentId;index:Agent_copiedFromAgentId_idx" json:"copiedFromAgentId"`
	Environment       *string        `gorm:"column:environment" json:"environment"`
	OwnerUserID       *string        `gorm:"column:ownerUserId" json:"ownerUserId"`
	HostLabel         *string        `gorm:"column:hostLabel" json:"hostLabel"`
	IdentityKey       *string        `gorm:"column:identityKey;uniqueIndex:Agent_projectId_identityKey_key,priority:2" json:"identityKey"`
	LastSeenAt        *DateTime      `gorm:"column:lastSeenAt;index:Agent_projectId_lastSeenAt_idx,priority:2;type:timestamp(3)" json:"lastSeenAt"`
	ArchivedAt        *DateTime      `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	CreatedAt         DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID       *string        `gorm:"column:createdById" json:"createdById"`
	UpdatedByID       *string        `gorm:"column:updatedById" json:"updatedById"`
	Project           *Project       `gorm:"belongsTo"`
	Workflow          *Workflow      `gorm:"belongsTo"`
	CopiedFrom        *Agent         `gorm:"foreignKey:CopiedFromAgentID;belongsTo"`
	CopiedAgents      []Agent        `gorm:"foreignKey:CopiedFromAgentID;constraint:Agent_copiedFromAgentId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (Agent) TableName() string {
	return "Agent"
}

func (m *Agent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Agent) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One immutable snapshot of an agent; the Agent row is always the latest.
type AgentVersion struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID     string         `gorm:"column:projectId;index:AgentVersion_projectId_idx;not null" json:"projectId"`
	AgentID       string         `gorm:"column:agentId;uniqueIndex:AgentVersion_agentId_version_key,priority:1;not null" json:"agentId"`
	Version       int            `gorm:"column:version;uniqueIndex:AgentVersion_agentId_version_key,priority:2;size:32;not null" json:"version"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	Type          string         `gorm:"column:type;not null" json:"type"`
	Config        datatypes.JSON `gorm:"column:config;not null" json:"config"`
	WorkflowID    *string        `gorm:"column:workflowId" json:"workflowId"`
	CommitMessage *string        `gorm:"column:commitMessage" json:"commitMessage"`
	AuthorID      *string        `gorm:"column:authorId" json:"authorId"`
	// "user", "api_key" or "system": who wrote the version when authorId is null.
	AuthorKind *string  `gorm:"column:authorKind" json:"authorKind"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (AgentVersion) TableName() string {
	return "AgentVersion"
}

func (m *AgentVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// A deploy label ("production", "staging") pointing at one version of an agent.
type AgentVersionLabel struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID   string   `gorm:"column:projectId;index:AgentVersionLabel_projectId_idx;not null" json:"projectId"`
	AgentID     string   `gorm:"column:agentId;uniqueIndex:AgentVersionLabel_agentId_label_key,priority:1;not null" json:"agentId"`
	Label       string   `gorm:"column:label;uniqueIndex:AgentVersionLabel_agentId_label_key,priority:2;not null" json:"label"`
	VersionID   string   `gorm:"column:versionId;index:AgentVersionLabel_versionId_idx;not null" json:"versionId"`
	CreatedByID *string  `gorm:"column:createdById" json:"createdById"`
	UpdatedByID *string  `gorm:"column:updatedById" json:"updatedById"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AgentVersionLabel) TableName() string {
	return "AgentVersionLabel"
}

func (m *AgentVersionLabel) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AgentVersionLabel) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Evaluator struct {
	ID                    string             `gorm:"column:id;primaryKey" json:"id"`
	ProjectID             string             `gorm:"column:projectId;uniqueIndex:Evaluator_projectId_slug_key,priority:1;index:Evaluator_projectId_idx;not null" json:"projectId"`
	Name                  string             `gorm:"column:name;not null" json:"name"`
	Slug                  *string            `gorm:"column:slug;uniqueIndex:Evaluator_projectId_slug_key,priority:2" json:"slug"`
	Type                  string             `gorm:"column:type;not null" json:"type"`
	Config                datatypes.JSON     `gorm:"column:config;not null" json:"config"`
	WorkflowID            *string            `gorm:"column:workflowId;index:Evaluator_workflowId_idx" json:"workflowId"`
	CopiedFromEvaluatorID *string            `gorm:"column:copiedFromEvaluatorId;index:Evaluator_copiedFromEvaluatorId_idx" json:"copiedFromEvaluatorId"`
	ArchivedAt            *DateTime          `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	CreatedAt             DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID           *string            `gorm:"column:createdById" json:"createdById"`
	UpdatedByID           *string            `gorm:"column:updatedById" json:"updatedById"`
	Project               *Project           `gorm:"belongsTo"`
	Workflow              *Workflow          `gorm:"belongsTo"`
	CopiedFrom            *Evaluator         `gorm:"foreignKey:CopiedFromEvaluatorID;belongsTo"`
	CopiedEvaluators      []Evaluator        `gorm:"foreignKey:CopiedFromEvaluatorID;constraint:Evaluator_copiedFromEvaluatorId_fkey,OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Monitors              []Monitor          `gorm:"foreignKey:EvaluatorID;constraint:Monitor_evaluatorId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	GatewayGuardrails     []GatewayGuardrail `gorm:"foreignKey:EvaluatorID;constraint:GatewayGuardrail_evaluatorId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Evaluator) TableName() string {
	return "Evaluator"
}

func (m *Evaluator) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Evaluator) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One immutable snapshot of an evaluator; the Evaluator row is always the latest.
type EvaluatorVersion struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID     string         `gorm:"column:projectId;index:EvaluatorVersion_projectId_idx;not null" json:"projectId"`
	EvaluatorID   string         `gorm:"column:evaluatorId;uniqueIndex:EvaluatorVersion_evaluatorId_version_key,priority:1;not null" json:"evaluatorId"`
	Version       int            `gorm:"column:version;uniqueIndex:EvaluatorVersion_evaluatorId_version_key,priority:2;size:32;not null" json:"version"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	Type          string         `gorm:"column:type;not null" json:"type"`
	Config        datatypes.JSON `gorm:"column:config;not null" json:"config"`
	WorkflowID    *string        `gorm:"column:workflowId" json:"workflowId"`
	CommitMessage *string        `gorm:"column:commitMessage" json:"commitMessage"`
	AuthorID      *string        `gorm:"column:authorId" json:"authorId"`
	// "user", "api_key" or "system": who wrote the version when authorId is null.
	AuthorKind *string  `gorm:"column:authorKind" json:"authorKind"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (EvaluatorVersion) TableName() string {
	return "EvaluatorVersion"
}

func (m *EvaluatorVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// A deploy label ("production", "staging") pointing at one version of an evaluator.
type EvaluatorVersionLabel struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID   string   `gorm:"column:projectId;index:EvaluatorVersionLabel_projectId_idx;not null" json:"projectId"`
	EvaluatorID string   `gorm:"column:evaluatorId;uniqueIndex:EvaluatorVersionLabel_evaluatorId_label_key,priority:1;not null" json:"evaluatorId"`
	Label       string   `gorm:"column:label;uniqueIndex:EvaluatorVersionLabel_evaluatorId_label_key,priority:2;not null" json:"label"`
	VersionID   string   `gorm:"column:versionId;index:EvaluatorVersionLabel_versionId_idx;not null" json:"versionId"`
	CreatedByID *string  `gorm:"column:createdById" json:"createdById"`
	UpdatedByID *string  `gorm:"column:updatedById" json:"updatedById"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (EvaluatorVersionLabel) TableName() string {
	return "EvaluatorVersionLabel"
}

func (m *EvaluatorVersionLabel) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EvaluatorVersionLabel) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Scenario struct {
	ID                 string             `gorm:"column:id;primaryKey" json:"id"`
	ProjectID          string             `gorm:"column:projectId;index:Scenario_projectId_idx;index:Scenario_projectId_testSuiteId_idx,priority:1;not null" json:"projectId"`
	Name               string             `gorm:"column:name;not null" json:"name"`
	Situation          string             `gorm:"column:situation;not null" json:"situation"`
	Criteria           ScalarList[string] `gorm:"column:criteria;type:text[]" json:"criteria"`
	Labels             ScalarList[string] `gorm:"column:labels;type:text[]" json:"labels"`
	Parameters         datatypes.JSON     `gorm:"column:parameters" json:"parameters"`
	SimulatorModel     *string            `gorm:"column:simulatorModel" json:"simulatorModel"`
	JudgeModel         *string            `gorm:"column:judgeModel" json:"judgeModel"`
	TraceQuietPeriodMs *int               `gorm:"column:traceQuietPeriodMs;size:32" json:"traceQuietPeriodMs"`
	TraceWaitTimeoutMs *int               `gorm:"column:traceWaitTimeoutMs;size:32" json:"traceWaitTimeoutMs"`
	MaxTurns           *int               `gorm:"column:maxTurns;size:32" json:"maxTurns"`
	MinTurns           *int               `gorm:"column:minTurns;size:32" json:"minTurns"`
	Fields             datatypes.JSON     `gorm:"column:fields" json:"fields"`
	CallerVoice        datatypes.JSON     `gorm:"column:callerVoice" json:"callerVoice"`
	TestSuiteID        *string            `gorm:"column:testSuiteId;index:Scenario_projectId_testSuiteId_idx,priority:2" json:"testSuiteId"`
	Kind               *string            `gorm:"column:kind;default:'standard';not null" json:"kind"`
	RedTeam            datatypes.JSON     `gorm:"column:redTeam" json:"redTeam"`
	SourceTemplateID   *string            `gorm:"column:sourceTemplateId" json:"sourceTemplateId"`
	Version            *int               `gorm:"column:version;size:32;default:1;not null" json:"version"`
	LastUpdatedByID    *string            `gorm:"column:lastUpdatedById" json:"lastUpdatedById"`
	ArchivedAt         *DateTime          `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	CreatedAt          DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project            *Project           `gorm:"belongsTo"`
	TestSuite          *SimulationSuite   `gorm:"foreignKey:TestSuiteID;belongsTo"`
	LastUpdatedBy      *User              `gorm:"foreignKey:LastUpdatedByID;belongsTo"`
	Versions           []ScenarioVersion  `gorm:"foreignKey:ScenarioID;constraint:ScenarioVersion_scenarioId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Scenario) TableName() string {
	return "Scenario"
}

func (m *Scenario) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Scenario) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ScenarioVersion struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`
	ScenarioID string `gorm:"column:scenarioId;uniqueIndex:ScenarioVersion_scenarioId_version_key,priority:1;index:ScenarioVersion_projectId_scenarioId_version_idx,priority:2;not null" json:"scenarioId"`
	// The project of the parent case, carried here so every read filters by
	// tenant without a join. The service is what keeps the pair honest: no
	// composite foreign key can, under relationMode="prisma".
	ProjectID string  `gorm:"column:projectId;index:ScenarioVersion_projectId_scenarioId_version_idx,priority:1;not null" json:"projectId"`
	Version   int     `gorm:"column:version;uniqueIndex:ScenarioVersion_scenarioId_version_key,priority:2;index:ScenarioVersion_projectId_scenarioId_version_idx,priority:3;size:32;not null" json:"version"`
	AuthorID  *string `gorm:"column:authorId" json:"authorId"`
	// Who wrote this version: "user", "api", "cli" or "langy".
	AuthorLabel       string  `gorm:"column:authorLabel;not null" json:"authorLabel"`
	ChangeDescription *string `gorm:"column:changeDescription" json:"changeDescription"`
	// Envelope { schemaVersion, fields, changedFields } holding the editable
	// field set as saved and the names of the fields this save changed.
	Snapshot datatypes.JSON `gorm:"column:snapshot;not null" json:"snapshot"`
	// Schema version of `snapshot`, so a later shape change can be read forward.
	SchemaVersion *int     `gorm:"column:schemaVersion;size:32;default:1;not null" json:"schemaVersion"`
	CreatedAt     DateTime `gorm:"column:createdAt;index:ScenarioVersion_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// Deleting the case takes its history with it. Under relationMode="prisma"
	// the cascade is the client's, so a raw SQL delete still has to clear these
	// rows itself.
	Scenario *Scenario `gorm:"belongsTo"`
}

func (ScenarioVersion) TableName() string {
	return "ScenarioVersion"
}

func (m *ScenarioVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SimulationSuite struct {
	ID                 string             `gorm:"column:id;primaryKey" json:"id"`
	ProjectID          string             `gorm:"column:projectId;uniqueIndex:SimulationSuite_projectId_slug_key,priority:1;index:SimulationSuite_projectId_idx;index:SimulationSuite_projectId_kind_idx,priority:1;not null" json:"projectId"`
	Name               string             `gorm:"column:name;not null" json:"name"`
	Slug               string             `gorm:"column:slug;uniqueIndex:SimulationSuite_projectId_slug_key,priority:2;not null" json:"slug"`
	Description        *string            `gorm:"column:description" json:"description"`
	ScenarioIds        ScalarList[string] `gorm:"column:scenarioIds;type:text[]" json:"scenarioIds"`
	Targets            datatypes.JSON     `gorm:"column:targets;not null" json:"targets"`
	RepeatCount        *int               `gorm:"column:repeatCount;size:32;default:1;not null" json:"repeatCount"`
	Labels             ScalarList[string] `gorm:"column:labels;type:text[]" json:"labels"`
	SimulatorModel     *string            `gorm:"column:simulatorModel" json:"simulatorModel"`
	JudgeModel         *string            `gorm:"column:judgeModel" json:"judgeModel"`
	TraceQuietPeriodMs *int               `gorm:"column:traceQuietPeriodMs;size:32" json:"traceQuietPeriodMs"`
	TraceWaitTimeoutMs *int               `gorm:"column:traceWaitTimeoutMs;size:32" json:"traceWaitTimeoutMs"`
	Kind               *string            `gorm:"column:kind;index:SimulationSuite_projectId_kind_idx,priority:2;default:'run_plan';not null" json:"kind"`
	Scope              datatypes.JSON     `gorm:"column:scope" json:"scope"`
	Fields             datatypes.JSON     `gorm:"column:fields" json:"fields"`
	Evaluators         datatypes.JSON     `gorm:"column:evaluators" json:"evaluators"`
	TemplateID         *string            `gorm:"column:templateId" json:"templateId"`
	TemplateVersion    *int               `gorm:"column:templateVersion;size:32" json:"templateVersion"`
	ArchivedAt         *DateTime          `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	CreatedAt          DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project            *Project           `gorm:"belongsTo"`
	TestSuiteScenarios []Scenario         `gorm:"foreignKey:TestSuiteID;constraint:Scenario_testSuiteId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (SimulationSuite) TableName() string {
	return "SimulationSuite"
}

func (m *SimulationSuite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SimulationSuite) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Subscription struct {
	ID                    string    `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID        string    `gorm:"column:organizationId;index:Subscription_organizationId_idx;not null" json:"organizationId"`
	Plan                  string    `gorm:"column:plan;not null" json:"plan"`
	StripeSubscriptionID  *string   `gorm:"column:stripeSubscriptionId;uniqueIndex:Subscription_stripeSubscriptionId_key" json:"stripeSubscriptionId"`
	Status                string    `gorm:"column:status;not null" json:"status"`
	StartDate             *DateTime `gorm:"column:startDate;type:timestamp(3)" json:"startDate"`
	EndDate               *DateTime `gorm:"column:endDate;type:timestamp(3)" json:"endDate"`
	LastPaymentFailedDate *DateTime `gorm:"column:lastPaymentFailedDate;type:timestamp(3)" json:"lastPaymentFailedDate"`
	CreatedAt             DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	MaxMembers            *int      `gorm:"column:maxMembers;size:32" json:"maxMembers"`
	MaxProjects           *int      `gorm:"column:maxProjects;size:32" json:"maxProjects"`
	MaxMessagesPerMonth   *int      `gorm:"column:maxMessagesPerMonth;size:32" json:"maxMessagesPerMonth"`
	MaxWorkflows          *int      `gorm:"column:maxWorkflows;size:32" json:"maxWorkflows"`
	MaxRetentionDays      *int      `gorm:"column:maxRetentionDays;size:32" json:"maxRetentionDays"`
	MaxMembersLite        *int      `gorm:"column:maxMembersLite;size:32" json:"maxMembersLite"`
	// Developer seats the organization pays for; null falls to the plan's own cap.
	MaxMembersDeveloper  *int `gorm:"column:maxMembersDeveloper;size:32" json:"maxMembersDeveloper"`
	MaxTeams             *int `gorm:"column:maxTeams;size:32" json:"maxTeams"`
	MaxPrompts           *int `gorm:"column:maxPrompts;size:32" json:"maxPrompts"`
	MaxEvaluators        *int `gorm:"column:maxEvaluators;size:32" json:"maxEvaluators"`
	MaxScenarios         *int `gorm:"column:maxScenarios;size:32" json:"maxScenarios"`
	MaxAgents            *int `gorm:"column:maxAgents;size:32" json:"maxAgents"`
	MaxExperiments       *int `gorm:"column:maxExperiments;size:32" json:"maxExperiments"`
	MaxOnlineEvaluations *int `gorm:"column:maxOnlineEvaluations;size:32" json:"maxOnlineEvaluations"`
	MaxDatasets          *int `gorm:"column:maxDatasets;size:32" json:"maxDatasets"`
	MaxDashboards        *int `gorm:"column:maxDashboards;size:32" json:"maxDashboards"`
	MaxCustomGraphs      *int `gorm:"column:maxCustomGraphs;size:32" json:"maxCustomGraphs"`
	MaxAutomations       *int `gorm:"column:maxAutomations;size:32" json:"maxAutomations"`
	// Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil.
	SeatsKept      *int      `gorm:"column:seatsKept;size:32" json:"seatsKept"`
	SeatsKeptUntil *DateTime `gorm:"column:seatsKeptUntil;type:timestamp(3)" json:"seatsKeptUntil"`
	// The seat quantity the last finished seat change set the seat line to; provider events never write it.
	SeatLineQuantity   *int      `gorm:"column:seatLineQuantity;size:32" json:"seatLineQuantity"`
	SeatLineQuantityAt *DateTime `gorm:"column:seatLineQuantityAt;type:timestamp(3)" json:"seatLineQuantityAt"`
	// A Stripe seat line write recorded before it is sent and cleared once it is known to have
	// finished under the lease; one left behind is settled by the seat line repair.
	// Every seat quantity written while this record was pending.
	SeatLineWriteQuantities ScalarList[int] `gorm:"column:seatLineWriteQuantities;type:integer[]" json:"seatLineWriteQuantities"`
	// The first write still pending; a later holder's write keeps it.
	SeatLineWriteSince *DateTime `gorm:"column:seatLineWriteSince;type:timestamp(3)" json:"seatLineWriteSince"`
	// The latest write sent; the repair waits a grace window after it.
	SeatLineWriteAt        *DateTime `gorm:"column:seatLineWriteAt;type:timestamp(3)" json:"seatLineWriteAt"`
	SeatLineWriteAttempts  *int      `gorm:"column:seatLineWriteAttempts;size:32" json:"seatLineWriteAttempts"`
	SeatLineWriteFlaggedAt *DateTime `gorm:"column:seatLineWriteFlaggedAt;type:timestamp(3)" json:"seatLineWriteFlaggedAt"`
	// The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled.
	SeatRepairQuantity *int `gorm:"column:seatRepairQuantity;size:32" json:"seatRepairQuantity"`
	// The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it.
	SeatRepairLiftedQuantity *int      `gorm:"column:seatRepairLiftedQuantity;size:32" json:"seatRepairLiftedQuantity"`
	SeatRepairSince          *DateTime `gorm:"column:seatRepairSince;type:timestamp(3)" json:"seatRepairSince"`
	SeatRepairAttempts       *int      `gorm:"column:seatRepairAttempts;size:32" json:"seatRepairAttempts"`
	// Set once the seat line repair has failed often enough that Ops should look at it.
	SeatRepairFlaggedAt *DateTime `gorm:"column:seatRepairFlaggedAt;type:timestamp(3)" json:"seatRepairFlaggedAt"`
	// The lease a seat change or seat line repair holds while it reads and moves the seat line.
	SeatLineLeaseToken *string   `gorm:"column:seatLineLeaseToken" json:"seatLineLeaseToken"`
	SeatLineLeaseUntil *DateTime `gorm:"column:seatLineLeaseUntil;type:timestamp(3)" json:"seatLineLeaseUntil"`
	// An open failed payment recovery case: the unpaid renewal invoice, day 0 (its first failed
	// charge) and the last notice sent (3, 7, 14 or 21). Cleared when the case closes.
	PaymentRecoveryInvoiceID *string   `gorm:"column:paymentRecoveryInvoiceId" json:"paymentRecoveryInvoiceId"`
	PaymentRecoveryStartedAt *DateTime `gorm:"column:paymentRecoveryStartedAt;type:timestamp(3)" json:"paymentRecoveryStartedAt"`
	PaymentRecoveryNoticeDay *int      `gorm:"column:paymentRecoveryNoticeDay;size:32" json:"paymentRecoveryNoticeDay"`
	// The over-plan mark: traces recorded from overPlanSince are withheld while it is open and the
	// organization is on Free with more members than Free allows; ended once that no longer holds.
	OverPlanSince   *DateTime     `gorm:"column:overPlanSince;type:timestamp(3)" json:"overPlanSince"`
	OverPlanEndedAt *DateTime     `gorm:"column:overPlanEndedAt;type:timestamp(3)" json:"overPlanEndedAt"`
	Organization    *Organization `gorm:"belongsTo"`
	Invoices        []Invoice     `gorm:"foreignKey:SubscriptionID;constraint:Invoice_subscriptionId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Subscription) TableName() string {
	return "Subscription"
}

func (m *Subscription) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Subscription) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Invoice struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	SubscriptionID string        `gorm:"column:subscriptionId;index:Invoice_subscriptionId_idx;not null" json:"subscriptionId"`
	TotalAmount    float64       `gorm:"column:totalAmount;type:double precision;not null" json:"totalAmount"`
	Currency       string        `gorm:"column:currency;not null" json:"currency"`
	DueDate        DateTime      `gorm:"column:dueDate;type:timestamp(3);not null" json:"dueDate"`
	Paid           bool          `gorm:"column:paid;default:false;not null" json:"paid"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Subscription   *Subscription `gorm:"belongsTo"`
	LineItems      []InvoiceItem `gorm:"foreignKey:InvoiceID;constraint:InvoiceItem_invoiceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Invoice) TableName() string {
	return "Invoice"
}

func (m *Invoice) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Invoice) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type InvoiceItem struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	InvoiceID   string   `gorm:"column:invoiceId;index:InvoiceItem_invoiceId_idx;not null" json:"invoiceId"`
	Description string   `gorm:"column:description;not null" json:"description"`
	Amount      float64  `gorm:"column:amount;type:double precision;not null" json:"amount"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Invoice     *Invoice `gorm:"belongsTo"`
}

func (InvoiceItem) TableName() string {
	return "InvoiceItem"
}

func (m *InvoiceItem) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*InvoiceItem) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Group struct {
	ID             string  `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string  `gorm:"column:organizationId;uniqueIndex:Group_organizationId_slug_key,priority:1;index:Group_organizationId_idx;not null" json:"organizationId"`
	Name           string  `gorm:"column:name;not null" json:"name"`
	Slug           string  `gorm:"column:slug;uniqueIndex:Group_organizationId_slug_key,priority:2;not null" json:"slug"`
	ExternalID     *string `gorm:"column:externalId;uniqueIndex:Group_scimConnectionId_externalId_key,priority:2" json:"externalId"`
	ScimSource     *string `gorm:"column:scimSource" json:"scimSource"`
	// D08: which SSO connection provisioned this group. Null for manually
	// created groups and for the SCIM groups that predate connection scoping.
	ScimConnectionID *string           `gorm:"column:scimConnectionId;uniqueIndex:Group_scimConnectionId_externalId_key,priority:1;index:Group_scimConnectionId_idx" json:"scimConnectionId"`
	CreatedAt        DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime          `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization     *Organization     `gorm:"belongsTo"`
	Members          []GroupMembership `gorm:"foreignKey:GroupID;constraint:GroupMembership_groupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleBindings     []RoleBinding     `gorm:"foreignKey:GroupID;constraint:RoleBinding_groupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Group) TableName() string {
	return "Group"
}

func (m *Group) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Group) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GroupMembership struct {
	UserID    string   `gorm:"column:userId;primaryKey;index:GroupMembership_userId_idx" json:"userId"`
	GroupID   string   `gorm:"column:groupId;primaryKey;index:GroupMembership_groupId_idx" json:"groupId"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	User      *User    `gorm:"belongsTo"`
	Group     *Group   `gorm:"belongsTo"`
}

func (GroupMembership) TableName() string {
	return "GroupMembership"
}

func (m *GroupMembership) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant").
type AuthzUserStanding struct {
	UserID            string    `gorm:"column:userId;primaryKey" json:"userId"`
	DeactivatedAt     *DateTime `gorm:"column:deactivatedAt;type:timestamp(3)" json:"deactivatedAt"`
	ErasedAt          *DateTime `gorm:"column:erasedAt;type:timestamp(3)" json:"erasedAt"`
	StandingChangedAt DateTime  `gorm:"column:standingChangedAt;type:timestamp(3);not null" json:"standingChangedAt"`
}

func (AuthzUserStanding) TableName() string {
	return "AuthzUserStanding"
}

// Whether an organization's permission checks run on the authorization engine: authz's own
// switch, set when the engine migration proves the organization and pinned off by an operator's
// rollback. No row is the legacy path. Spec: modules/authz/specs/engine-switch.feature.
type AuthzEngineSwitch struct {
	OrganizationID     string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	OnEngine           bool     `gorm:"column:onEngine;not null" json:"onEngine"`
	Pinned             bool     `gorm:"column:pinned;default:false;not null" json:"pinned"`
	SwitchedAt         DateTime `gorm:"column:switchedAt;type:timestamp(3);not null" json:"switchedAt"`
	LastAppliedEventID *string  `gorm:"column:lastAppliedEventId" json:"lastAppliedEventId"`
	UpdatedAt          DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AuthzEngineSwitch) TableName() string {
	return "AuthzEngineSwitch"
}

func (m *AuthzEngineSwitch) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AuthzEngineSwitch) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot.
type AuthzOrganizationPolicy struct {
	OrganizationID               string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	BraidsOrganizationAdminsOnly bool     `gorm:"column:braidsOrganizationAdminsOnly;default:false;not null" json:"braidsOrganizationAdminsOnly"`
	BraidsEnabled                bool     `gorm:"column:braidsEnabled;default:false;not null" json:"braidsEnabled"`
	UpdatedByUserID              string   `gorm:"column:updatedByUserId;not null" json:"updatedByUserId"`
	UpdatedAt                    DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AuthzOrganizationPolicy) TableName() string {
	return "AuthzOrganizationPolicy"
}

func (m *AuthzOrganizationPolicy) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AuthzOrganizationPolicy) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AuthzBraidWrite struct {
	OrganizationID  string         `gorm:"column:organizationId;primaryKey;index:AuthzBraidWrite_organizationId_readerProjectId_createdAt_idx,priority:1" json:"organizationId"`
	CommandID       string         `gorm:"column:commandId;primaryKey" json:"commandId"`
	GrantID         string         `gorm:"column:grantId;not null" json:"grantId"`
	ReaderProjectID string         `gorm:"column:readerProjectId;index:AuthzBraidWrite_organizationId_readerProjectId_createdAt_idx,priority:2;not null" json:"readerProjectId"`
	OwnerProjectID  string         `gorm:"column:ownerProjectId;not null" json:"ownerProjectId"`
	TermsKey        string         `gorm:"column:termsKey;not null" json:"termsKey"`
	Condition       datatypes.JSON `gorm:"column:condition;not null" json:"condition"`
	ExpiresAt       *DateTime      `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	ReplacesGrantID *string        `gorm:"column:replacesGrantId" json:"replacesGrantId"`
	CreatedByUserID *string        `gorm:"column:createdByUserId" json:"createdByUserId"`
	CreatedAt       DateTime       `gorm:"column:createdAt;index:AuthzBraidWrite_organizationId_readerProjectId_createdAt_idx,priority:3;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (AuthzBraidWrite) TableName() string {
	return "AuthzBraidWrite"
}

func (m *AuthzBraidWrite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type RoleBinding struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string        `gorm:"column:organizationId;index:RoleBinding_organizationId_idx;not null" json:"organizationId"`
	UserID         *string       `gorm:"column:userId;index:RoleBinding_userId_idx" json:"userId"`
	GroupID        *string       `gorm:"column:groupId;index:RoleBinding_groupId_idx" json:"groupId"`
	APIKeyID       *string       `gorm:"column:apiKeyId;index:RoleBinding_apiKeyId_idx" json:"apiKeyId"`
	Role           string        `gorm:"column:role;not null" json:"role"`
	CustomRoleID   *string       `gorm:"column:customRoleId" json:"customRoleId"`
	ScopeType      string        `gorm:"column:scopeType;index:RoleBinding_scopeType_scopeId_idx,priority:1;not null" json:"scopeType"`
	ScopeID        string        `gorm:"column:scopeId;index:RoleBinding_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization `gorm:"belongsTo"`
	User           *User         `gorm:"belongsTo"`
	Group          *Group        `gorm:"belongsTo"`
	APIKey         *APIKey       `gorm:"belongsTo"`
	CustomRole     *CustomRole   `gorm:"belongsTo"`
}

func (RoleBinding) TableName() string {
	return "RoleBinding"
}

func (m *RoleBinding) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*RoleBinding) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Grant struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID  string         `gorm:"column:organizationId;index:Grant_organizationId_scopeType_scopeId_idx,priority:1;not null" json:"organizationId"`
	PrincipalType   string         `gorm:"column:principalType;not null" json:"principalType"`
	PrincipalID     *string        `gorm:"column:principalId" json:"principalId"`
	RoleKey         *string        `gorm:"column:roleKey" json:"roleKey"`
	LegacyRole      *string        `gorm:"column:legacyRole" json:"legacyRole"`
	Source          string         `gorm:"column:source;not null" json:"source"`
	ScopeType       string         `gorm:"column:scopeType;index:Grant_organizationId_scopeType_scopeId_idx,priority:2;not null" json:"scopeType"`
	ScopeID         string         `gorm:"column:scopeId;index:Grant_organizationId_scopeType_scopeId_idx,priority:3;index:Grant_projectId_resourceKind_scopeId_idx,priority:3;not null" json:"scopeId"`
	Token           *string        `gorm:"column:token;uniqueIndex:Grant_token_key" json:"token"`
	Permission      *string        `gorm:"column:permission" json:"permission"`
	ResourceKind    *string        `gorm:"column:resourceKind;index:Grant_projectId_resourceKind_scopeId_idx,priority:2" json:"resourceKind"`
	ProjectID       *string        `gorm:"column:projectId;index:Grant_projectId_resourceKind_scopeId_idx,priority:1" json:"projectId"`
	CreatedByUserID *string        `gorm:"column:createdByUserId" json:"createdByUserId"`
	ExpiresAt       *DateTime      `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	MaxViews        *int           `gorm:"column:maxViews;size:32" json:"maxViews"`
	Condition       datatypes.JSON `gorm:"column:condition" json:"condition"`
	ReplacesGrantID *string        `gorm:"column:replacesGrantId" json:"replacesGrantId"`
	OccurredAt      DateTime       `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	RevokedAt       *DateTime      `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	RevokedReason   *string        `gorm:"column:revokedReason" json:"revokedReason"`
}

func (Grant) TableName() string {
	return "Grant"
}

func (m *Grant) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Grant) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GrantUsage struct {
	GrantID        string    `gorm:"column:grantId;primaryKey" json:"grantId"`
	OrganizationID string    `gorm:"column:organizationId;index:GrantUsage_organizationId_idx;not null" json:"organizationId"`
	ProjectID      string    `gorm:"column:projectId;index:GrantUsage_projectId_idx;not null" json:"projectId"`
	ViewCount      int       `gorm:"column:viewCount;size:32;default:0;not null" json:"viewCount"`
	LastViewedAt   *DateTime `gorm:"column:lastViewedAt;type:timestamp(3)" json:"lastViewedAt"`
	CreatedAt      DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (GrantUsage) TableName() string {
	return "GrantUsage"
}

func (m *GrantUsage) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GrantUsage) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Role struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string         `gorm:"column:organizationId;index:Role_organizationId_idx;not null" json:"organizationId"`
	Name           string         `gorm:"column:name;not null" json:"name"`
	Description    *string        `gorm:"column:description" json:"description"`
	Permissions    datatypes.JSON `gorm:"column:permissions;not null" json:"permissions"`
	Kind           *string        `gorm:"column:kind;default:'custom';not null" json:"kind"`
	OccurredAt     DateTime       `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt      *DateTime      `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
}

func (Role) TableName() string {
	return "Role"
}

func (m *Role) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Role) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type APIKey struct {
	ID              string    `gorm:"column:id;primaryKey" json:"id"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Description     *string   `gorm:"column:description" json:"description"`
	LookupID        string    `gorm:"column:lookupId;uniqueIndex:ApiKey_lookupId_key;index:ApiKey_lookupId_idx;not null" json:"lookupId"`
	HashedSecret    string    `gorm:"column:hashedSecret;not null" json:"hashedSecret"`
	PermissionMode  *string   `gorm:"column:permissionMode;default:'all';not null" json:"permissionMode"`
	UserID          *string   `gorm:"column:userId;index:ApiKey_userId_idx" json:"userId"`
	CreatedByUserID *string   `gorm:"column:createdByUserId" json:"createdByUserId"`
	OrganizationID  string    `gorm:"column:organizationId;index:ApiKey_organizationId_idx;not null" json:"organizationId"`
	ExpiresAt       *DateTime `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	LastUsedAt      *DateTime `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	RevokedAt       *DateTime `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	// Why the key was revoked, set with `revokedAt`: "user" when a person
	// revoked it (the API-keys page, the REST API, the devices tab), "rotation"
	// when a mint replaced it, "session" when the CLI login key it was parented
	// to was revoked, "expired" when that session ran out, "offboarded" when the
	// person's membership of the organization ended. "cap" survives on rows the
	// retired per-tool cap revoked. The CLI reads it to decide whether a dead
	// ingest key may be re-minted: a key a person revoked stays dead.
	// NULL for live keys and for keys revoked before the cause was recorded.
	RevocationCause *string `gorm:"column:revocationCause" json:"revocationCause"`
	// The CLI login key of the device session that minted this ingestion key.
	// Set by the CLI's personal mint, NULL for every other key. Revoking the
	// parent (logout, the devices tab, a re-login from the same device, session
	// expiry) revokes the children with it. Column ref only, no FK and no
	// index: lookups filter by the indexed `userId` first and match the parent
	// in memory over a person's few live ingest keys.
	ParentAPIKeyID *string `gorm:"column:parentApiKeyId" json:"parentApiKeyId"`
	// Non-null marks this key as an "ingestion key": a project-scoped,
	// ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
	// installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
	// opencode / claude_cowork) is stamped as `langwatch.source` provenance on
	// every span the key authorizes. NULL for normal API / service keys.
	IngestSourceType *string `gorm:"column:ingestSourceType;index:ApiKey_ingestSourceType_idx" json:"ingestSourceType"`
	// Set only for ingestion keys derived from an admin IngestionTemplate
	// (e.g. claude_cowork). Drives receiver OTTL application + the
	// `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
	// (which have no template) and for normal keys. Column ref only (no FK,
	// per the app-layer-integrity convention); resolved by the ingest-key service.
	IngestionTemplateID *string `gorm:"column:ingestionTemplateId;index:ApiKey_ingestionTemplateId_idx" json:"ingestionTemplateId"`
	// Human label of the CLI device session that minted this ingestion key
	// (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
	// flow's `client_info`. Denormalized so the org-wide API-keys settings page
	// can show which device a key came from without resolving another user's
	// Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
	// a CLI old enough to predate device metadata.
	CreatedByDeviceLabel *string `gorm:"column:createdByDeviceLabel" json:"createdByDeviceLabel"`
	// True for a key LangWatch minted for itself (a run, the gateway's trace
	// export, a Langy session). Hides it from listings, refuses customer edits
	// and revokes, and lets an ownerless run key act as the system. Never set
	// by a customer create; rows before the column stay false.
	IsSystemManaged bool `gorm:"column:isSystemManaged;default:false;not null" json:"isSystemManaged"`
	// The browser session that minted a docs session key (ADR-185), so signing
	// out of that session revokes it. NULL for every other key. Unindexed:
	// sign-out filters by the indexed `userId` first.
	CreatedByBrowserSessionID *string       `gorm:"column:createdByBrowserSessionId" json:"createdByBrowserSessionId"`
	CreatedAt                 DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User                      *User         `gorm:"belongsTo"`
	Organization              *Organization `gorm:"belongsTo"`
	RoleBindings              []RoleBinding `gorm:"foreignKey:APIKeyID;constraint:RoleBinding_apiKeyId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (APIKey) TableName() string {
	return "ApiKey"
}

func (m *APIKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
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

// ADR-201: a project's key for signing the tags on its calls. Owned by the API key module.
type SigningKey struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID   string    `gorm:"column:organizationId;index:SigningKey_organizationId_idx;not null" json:"organizationId"`
	ProjectID        string    `gorm:"column:projectId;index:SigningKey_projectId_idx;not null" json:"projectId"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	PublicID         string    `gorm:"column:publicId;uniqueIndex:SigningKey_publicId_key;not null" json:"publicId"`
	SecretCiphertext string    `gorm:"column:secretCiphertext;not null" json:"secretCiphertext"`
	CreatedByUserID  string    `gorm:"column:createdByUserId;not null" json:"createdByUserId"`
	CreatedAt        DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	RotatedAt        *DateTime `gorm:"column:rotatedAt;type:timestamp(3)" json:"rotatedAt"`
	VerifiesUntil    *DateTime `gorm:"column:verifiesUntil;type:timestamp(3)" json:"verifiesUntil"`
	RevokedAt        *DateTime `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	LastUsedAt       *DateTime `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
}

func (SigningKey) TableName() string {
	return "SigningKey"
}

func (m *SigningKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type ProjectActiveDay struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string   `gorm:"column:projectId;uniqueIndex:ProjectActiveDay_projectId_day_key,priority:1;not null" json:"projectId"`
	Day       string   `gorm:"column:day;uniqueIndex:ProjectActiveDay_projectId_day_key,priority:2;not null" json:"day"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (ProjectActiveDay) TableName() string {
	return "ProjectActiveDay"
}

func (m *ProjectActiveDay) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Nurturing's read model: which organization a project belongs to, from project's created event.
// The per-email usage places a project and lists an organization's projects from these rows only.
type NurturingProject struct {
	ProjectID      string `gorm:"column:projectId;primaryKey" json:"projectId"`
	OrganizationID string `gorm:"column:organizationId;index:NurturingProject_organizationId_idx;not null" json:"organizationId"`
}

func (NurturingProject) TableName() string {
	return "NurturingProject"
}

// A project nurturing saw traces of with no placement copy, and when it first did: every worker
// retries its trace batches only within one window from then, across restarts. Deleted when the
// copy lands.
type NurturingUncopiedProject struct {
	ProjectID   string   `gorm:"column:projectId;primaryKey" json:"projectId"`
	FirstSeenAt DateTime `gorm:"column:firstSeenAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"firstSeenAt"`
}

func (NurturingUncopiedProject) TableName() string {
	return "NurturingUncopiedProject"
}

func (m *NurturingUncopiedProject) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.FirstSeenAt.IsZero() {
		m.FirstSeenAt = DateTime{Time: now}
	}
	return nil
}

// Nurturing's own copy of each organization's members, from the owners' member facts (§9):
// organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
// and member_disabled leave, member_enabled returns. Each column keeps its latest time.
type NurturingOrganizationMember struct {
	OrganizationID string    `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	UserID         string    `gorm:"column:userId;primaryKey;index:NurturingOrganizationMember_userId_idx" json:"userId"`
	JoinedAt       *DateTime `gorm:"column:joinedAt;type:timestamp(3)" json:"joinedAt"`
	LeftAt         *DateTime `gorm:"column:leftAt;type:timestamp(3)" json:"leftAt"`
	ReturnedAt     *DateTime `gorm:"column:returnedAt;type:timestamp(3)" json:"returnedAt"`
}

func (NurturingOrganizationMember) TableName() string {
	return "NurturingOrganizationMember"
}

// When an organization was created, from organization's created event: nurturing's own copy for
// the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
// NurturingOrganization row.
type NurturingOrganizationCreated struct {
	OrganizationID string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	CreatedAt      DateTime `gorm:"column:createdAt;index:NurturingOrganizationCreated_createdAt_idx;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
}

func (NurturingOrganizationCreated) TableName() string {
	return "NurturingOrganizationCreated"
}

// Nurturing's read model per organization: its admin and its milestone counts.
// `seeded` marks an organization learned from the project backfill: its first_* already happened.
type NurturingOrganization struct {
	OrganizationID       string             `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	AdminUserID          *string            `gorm:"column:adminUserId" json:"adminUserId"`
	Seeded               bool               `gorm:"column:seeded;default:false;not null" json:"seeded"`
	EvaluationCount      int                `gorm:"column:evaluationCount;size:32;default:0;not null" json:"evaluationCount"`
	SimulationRunCount   int                `gorm:"column:simulationRunCount;size:32;default:0;not null" json:"simulationRunCount"`
	CodingFirstSessionAt *DateTime          `gorm:"column:codingFirstSessionAt;type:timestamp(3)" json:"codingFirstSessionAt"`
	CodingAgentTypes     ScalarList[string] `gorm:"column:codingAgentTypes;type:text[]" json:"codingAgentTypes"`
	CodingActiveDays     int                `gorm:"column:codingActiveDays;size:32;default:0;not null" json:"codingActiveDays"`
	CodingLastActiveDay  *string            `gorm:"column:codingLastActiveDay" json:"codingLastActiveDay"`
	CodingTerminalTypes  ScalarList[string] `gorm:"column:codingTerminalTypes;type:text[]" json:"codingTerminalTypes"`
	CodingRepositoryKeys ScalarList[string] `gorm:"column:codingRepositoryKeys;type:text[]" json:"codingRepositoryKeys"`
	CodingFirstDays      ScalarList[int]    `gorm:"column:codingFirstDays;type:integer[]" json:"codingFirstDays"`
	LastTraceAt          *DateTime          `gorm:"column:lastTraceAt;type:timestamp(3)" json:"lastTraceAt"`
	UpdatedAt            DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (NurturingOrganization) TableName() string {
	return "NurturingOrganization"
}

func (m *NurturingOrganization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*NurturingOrganization) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
// and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
// organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
type NurturingProjectUsageView struct {
	ProjectID string `gorm:"column:projectId;primaryKey" json:"projectId"`
	// The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
	// or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
	Owner string `gorm:"column:owner;primaryKey" json:"owner"`
	// The project's usage as the fold folded it (NurturingProjectUsageState).
	Folded datatypes.JSON `gorm:"column:folded;not null" json:"folded"`
	// The event ids of the last delivery folded in, so a redelivery folds once.
	AppliedEventIds ScalarList[string] `gorm:"column:appliedEventIds;type:text[]" json:"appliedEventIds"`
	// A row of another version reads as undecodable and is re-folded from the log.
	ProjectionVersion string `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	// Event-derived, so a replayed row equals the row it rebuilds.
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (NurturingProjectUsageView) TableName() string {
	return "NurturingProjectUsageView"
}

// Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
// root span being one trace, one row per project and UTC day of append time, incremented in place
// by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
type NurturingProjectTraceDay struct {
	ProjectID string `gorm:"column:projectId;primaryKey" json:"projectId"`
	// UTC `YYYY-MM-DD` of the root spans' append time.
	Day        string `gorm:"column:day;primaryKey" json:"day"`
	TraceCount int    `gorm:"column:traceCount;size:32;default:0;not null" json:"traceCount"`
	// Traces trace recorded failed (lw.trace.trace_failed).
	ErrorCount int `gorm:"column:errorCount;size:32;default:0;not null" json:"errorCount"`
	// Append time of the newest root span.
	LastTraceAt *DateTime `gorm:"column:lastTraceAt;type:timestamp(3)" json:"lastTraceAt"`
	// Append time of the newest trace_failed.
	LastFailedTraceAt *DateTime `gorm:"column:lastFailedTraceAt;type:timestamp(3)" json:"lastFailedTraceAt"`
}

func (NurturingProjectTraceDay) TableName() string {
	return "NurturingProjectTraceDay"
}

// Nurturing's copy of each personal project's owner, from organization's and project's facts.
// The first owner recorded stays: a personal project's owner never changes.
type NurturingPersonalProject struct {
	ProjectID      string `gorm:"column:projectId;primaryKey" json:"projectId"`
	OrganizationID string `gorm:"column:organizationId;not null" json:"organizationId"`
	OwnerUserID    string `gorm:"column:ownerUserId;not null" json:"ownerUserId"`
}

func (NurturingPersonalProject) TableName() string {
	return "NurturingPersonalProject"
}

// Nurturing's copy of each project that received a real trace, from trace's and project's facts.
type NurturingTracedProject struct {
	ProjectID      string `gorm:"column:projectId;primaryKey" json:"projectId"`
	OrganizationID string `gorm:"column:organizationId;index:NurturingTracedProject_organizationId_idx;not null" json:"organizationId"`
}

func (NurturingTracedProject) TableName() string {
	return "NurturingTracedProject"
}

type BillingMeterCheckpoint struct {
	ID                   string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID       string   `gorm:"column:organizationId;uniqueIndex:BillingMeterCheckpoint_organizationId_billingMonth_meter_key,priority:1;not null" json:"organizationId"`
	BillingMonth         string   `gorm:"column:billingMonth;uniqueIndex:BillingMeterCheckpoint_organizationId_billingMonth_meter_key,priority:2;not null" json:"billingMonth"`
	Meter                *string  `gorm:"column:meter;uniqueIndex:BillingMeterCheckpoint_organizationId_billingMonth_meter_key,priority:3;default:'langwatch_billable_events';not null" json:"meter"`
	LastReportedTotal    int      `gorm:"column:lastReportedTotal;size:32;default:0;not null" json:"lastReportedTotal"`
	PendingReportedTotal *int     `gorm:"column:pendingReportedTotal;size:32" json:"pendingReportedTotal"`
	ConsecutiveFailures  int      `gorm:"column:consecutiveFailures;size:32;default:0;not null" json:"consecutiveFailures"`
	LastCountedEventID   *string  `gorm:"column:lastCountedEventId" json:"lastCountedEventId"`
	UpdatedAt            DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (BillingMeterCheckpoint) TableName() string {
	return "BillingMeterCheckpoint"
}

func (m *BillingMeterCheckpoint) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BillingMeterCheckpoint) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type VirtualKey struct {
	ID                       string                `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID           string                `gorm:"column:organizationId;uniqueIndex:VirtualKey_organizationId_externalId_key,priority:1;index:VirtualKey_organizationId_idx;index:VirtualKey_organizationId_purpose_idx,priority:1;not null" json:"organizationId"`
	Name                     string                `gorm:"column:name;not null" json:"name"`
	Description              *string               `gorm:"column:description" json:"description"`
	Status                   string                `gorm:"column:status;index:VirtualKey_status_idx;default:'ACTIVE';not null" json:"status"`
	Purpose                  string                `gorm:"column:purpose;index:VirtualKey_organizationId_purpose_idx,priority:2;default:'USER';not null" json:"purpose"`
	ExternalID               *string               `gorm:"column:externalId;uniqueIndex:VirtualKey_organizationId_externalId_key,priority:2" json:"externalId"`
	Metadata                 datatypes.JSON        `gorm:"column:metadata;default:'{}';not null" json:"metadata"`
	DisabledAt               *DateTime             `gorm:"column:disabledAt;type:timestamp(3)" json:"disabledAt"`
	DisabledReason           *string               `gorm:"column:disabledReason" json:"disabledReason"`
	ExpiresAt                *DateTime             `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	HashedSecret             string                `gorm:"column:hashedSecret;uniqueIndex:VirtualKey_hashedSecret_key;not null" json:"hashedSecret"`
	DisplayPrefix            string                `gorm:"column:displayPrefix;index:VirtualKey_displayPrefix_idx;not null" json:"displayPrefix"`
	PrincipalUserID          *string               `gorm:"column:principalUserId;index:VirtualKey_principalUserId_idx" json:"principalUserId"`
	TraceProjectID           *string               `gorm:"column:traceProjectId" json:"traceProjectId"`
	Config                   datatypes.JSON        `gorm:"column:config;default:'{}';not null" json:"config"`
	Revision                 int64                 `gorm:"column:revision;index:VirtualKey_revision_idx;default:0;not null" json:"revision"`
	PreviousHashedSecret     *string               `gorm:"column:previousHashedSecret" json:"previousHashedSecret"`
	PreviousSecretValidUntil *DateTime             `gorm:"column:previousSecretValidUntil;type:timestamp(3)" json:"previousSecretValidUntil"`
	RevokedAt                *DateTime             `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	RevokedByID              *string               `gorm:"column:revokedById" json:"revokedById"`
	CreatedAt                DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime              `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID              string                `gorm:"column:createdById;not null" json:"createdById"`
	LastUsedAt               *DateTime             `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	RoutingPolicyID          *string               `gorm:"column:routingPolicyId;index:VirtualKey_routingPolicyId_idx" json:"routingPolicyId"`
	RoutingMode              string                `gorm:"column:routingMode;default:'NONE';not null" json:"routingMode"`
	ConnectServices          ScalarList[string]    `gorm:"column:connectServices;type:text[]" json:"connectServices"`
	LicenseTokenHash         *string               `gorm:"column:licenseTokenHash;uniqueIndex:VirtualKey_licenseTokenHash_key" json:"licenseTokenHash"`
	LicenseInstanceID        *string               `gorm:"column:licenseInstanceId" json:"licenseInstanceId"`
	LicenseExpiresAt         *DateTime             `gorm:"column:licenseExpiresAt;type:timestamp(3)" json:"licenseExpiresAt"`
	LicenseID                *string               `gorm:"column:licenseId" json:"licenseId"`
	PrincipalUser            *User                 `gorm:"foreignKey:PrincipalUserID;belongsTo"`
	RoutingPolicy            *RoutingPolicy        `gorm:"belongsTo"`
	Scopes                   []VirtualKeyScope     `gorm:"foreignKey:VirtualKeyID;constraint:VirtualKeyScope_virtualKeyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	LedgerEntries            []GatewayBudgetLedger `gorm:"foreignKey:VirtualKeyID;constraint:GatewayBudgetLedger_virtualKeyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (VirtualKey) TableName() string {
	return "VirtualKey"
}

func (m *VirtualKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*VirtualKey) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type VirtualKeyScope struct {
	ID           string      `gorm:"column:id;primaryKey" json:"id"`
	VirtualKeyID string      `gorm:"column:virtualKeyId;uniqueIndex:VirtualKeyScope_virtualKeyId_scopeType_scopeId_key,priority:1;index:VirtualKeyScope_virtualKeyId_idx;not null" json:"virtualKeyId"`
	ScopeType    string      `gorm:"column:scopeType;uniqueIndex:VirtualKeyScope_virtualKeyId_scopeType_scopeId_key,priority:2;index:VirtualKeyScope_scopeType_scopeId_idx,priority:1;not null" json:"scopeType"`
	ScopeID      string      `gorm:"column:scopeId;uniqueIndex:VirtualKeyScope_virtualKeyId_scopeType_scopeId_key,priority:3;index:VirtualKeyScope_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	CreatedAt    DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	VirtualKey   *VirtualKey `gorm:"belongsTo"`
}

func (VirtualKeyScope) TableName() string {
	return "VirtualKeyScope"
}

func (m *VirtualKeyScope) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type RoutingPolicy struct {
	ID               string               `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID   string               `gorm:"column:organizationId;index:RoutingPolicy_organizationId_idx;not null" json:"organizationId"`
	Name             string               `gorm:"column:name;not null" json:"name"`
	Description      *string              `gorm:"column:description" json:"description"`
	ModelProviderIds datatypes.JSON       `gorm:"column:modelProviderIds;default:'[]';not null" json:"modelProviderIds"`
	ModelAliases     datatypes.JSON       `gorm:"column:modelAliases;default:'{}';not null" json:"modelAliases"`
	DefaultModel     *string              `gorm:"column:defaultModel" json:"defaultModel"`
	PolicyRules      datatypes.JSON       `gorm:"column:policyRules;default:'{}';not null" json:"policyRules"`
	IsDefault        bool                 `gorm:"column:isDefault;default:false;not null" json:"isDefault"`
	CreatedAt        DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime             `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID      *string              `gorm:"column:createdById" json:"createdById"`
	UpdatedByID      *string              `gorm:"column:updatedById" json:"updatedById"`
	Organization     *Organization        `gorm:"belongsTo"`
	CreatedBy        *User                `gorm:"foreignKey:CreatedByID;belongsTo"`
	UpdatedBy        *User                `gorm:"foreignKey:UpdatedByID;belongsTo"`
	VirtualKeys      []VirtualKey         `gorm:"foreignKey:RoutingPolicyID;constraint:VirtualKey_routingPolicyId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Scopes           []RoutingPolicyScope `gorm:"foreignKey:RoutingPolicyID;constraint:RoutingPolicyScope_routingPolicyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RoutingPolicy) TableName() string {
	return "RoutingPolicy"
}

func (m *RoutingPolicy) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*RoutingPolicy) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RoutingPolicyScope struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	RoutingPolicyID string         `gorm:"column:routingPolicyId;uniqueIndex:RoutingPolicyScope_routingPolicyId_scopeType_scopeId_key,priority:1;index:RoutingPolicyScope_routingPolicyId_idx;not null" json:"routingPolicyId"`
	ScopeType       string         `gorm:"column:scopeType;uniqueIndex:RoutingPolicyScope_routingPolicyId_scopeType_scopeId_key,priority:2;index:RoutingPolicyScope_scopeType_scopeId_idx,priority:1;not null" json:"scopeType"`
	ScopeID         string         `gorm:"column:scopeId;uniqueIndex:RoutingPolicyScope_routingPolicyId_scopeType_scopeId_key,priority:3;index:RoutingPolicyScope_scopeType_scopeId_idx,priority:2;not null" json:"scopeId"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	RoutingPolicy   *RoutingPolicy `gorm:"belongsTo"`
}

func (RoutingPolicyScope) TableName() string {
	return "RoutingPolicyScope"
}

func (m *RoutingPolicyScope) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type GatewayGuardrail struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	ProjectID   string     `gorm:"column:projectId;index:GatewayGuardrail_projectId_idx;index:GatewayGuardrail_projectId_archivedAt_idx,priority:1;not null" json:"projectId"`
	Name        string     `gorm:"column:name;not null" json:"name"`
	Description *string    `gorm:"column:description" json:"description"`
	EvaluatorID string     `gorm:"column:evaluatorId;index:GatewayGuardrail_evaluatorId_idx;not null" json:"evaluatorId"`
	Direction   string     `gorm:"column:direction;not null" json:"direction"`
	FailureMode string     `gorm:"column:failureMode;default:'FAIL_CLOSED';not null" json:"failureMode"`
	ArchivedAt  *DateTime  `gorm:"column:archivedAt;index:GatewayGuardrail_projectId_archivedAt_idx,priority:2;type:timestamp(3)" json:"archivedAt"`
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime   `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID *string    `gorm:"column:createdById" json:"createdById"`
	UpdatedByID *string    `gorm:"column:updatedById" json:"updatedById"`
	Project     *Project   `gorm:"belongsTo"`
	Evaluator   *Evaluator `gorm:"belongsTo"`
	CreatedBy   *User      `gorm:"foreignKey:CreatedByID;belongsTo"`
	UpdatedBy   *User      `gorm:"foreignKey:UpdatedByID;belongsTo"`
}

func (GatewayGuardrail) TableName() string {
	return "GatewayGuardrail"
}

func (m *GatewayGuardrail) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GatewayGuardrail) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Per-platform fleet configuration that connects a closed SaaS
// product's audit / OTel / S3 stream to LangWatch's Activity
// Monitor (cf. specs/ai-gateway/governance/activity-monitor.feature
// + ingestion-sources.feature, docs/ai-gateway/governance/architecture.md).
//
// One IngestionSource = one upstream fleet (e.g. "Miro Cowork",
// "Acme Workato production"). The runtime ingests events on whatever
// cadence the source supports - push for OTel/webhook/S3 (Tier C/D),
// pull for poll-based admin APIs (Tier C). Per-source-type adapters
// normalise platform-specific shapes into OCSF + AOS before they
// land in trace_summaries with SourceType + SourceId tags.
type IngestionSource struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:IngestionSource_organizationId_name_key,priority:1;index:IngestionSource_organizationId_idx;not null" json:"organizationId"`
	// Optional team-scope. null = org-wide.
	TeamID *string `gorm:"column:teamId;index:IngestionSource_teamId_idx" json:"teamId"`
	// 'otel_generic' | 'claude_cowork' | 'workato' | 'copilot_studio' |
	// 'openai_compliance' | 'claude_compliance' | 's3_custom' (extensible).
	// String not enum so adapters can be added without migrations - the
	// supported set lives in code (activity-monitor adapter registry).
	SourceType string `gorm:"column:sourceType;index:IngestionSource_sourceType_idx;not null" json:"sourceType"`
	// Human-friendly label for the admin UI. Unique per-org so two
	// "Cowork" sources can coexist with different names.
	Name        string  `gorm:"column:name;uniqueIndex:IngestionSource_organizationId_name_key,priority:2;not null" json:"name"`
	Description *string `gorm:"column:description" json:"description"`
	// HMAC bearer token the upstream operator pastes into the platform
	// (or LangWatch sends in pull mode). Stored encrypted at rest via
	// the same pepper-derived KEK as VirtualKey secrets. Auto-generated
	// on create, rotatable via the admin UI with a 24h grace window.
	IngestSecretHash string `gorm:"column:ingestSecretHash;not null" json:"ingestSecretHash"`
	// Per-source-type configuration. Shape is platform-specific:
	// - otel_generic: { allowedSourceTypeLabel?: string }
	// - claude_cowork: { workspaceId: string }
	// - workato: { sharedSecretLastFour: string }
	// - copilot_studio: { tenantId, clientId, clientSecretHash, pollEverySec }
	// - openai_compliance: { bucket, prefix, roleArn, pollEverySec }
	// - claude_compliance: { workspaceApiKeyHash, pollEverySec }
	// - s3_custom: { bucket, prefix, roleArn, parserDsl, pollEverySec }
	ParserConfig datatypes.JSON `gorm:"column:parserConfig;default:'{}';not null" json:"parserConfig"`
	// For pull-based sources, persisted between polls so the next poll
	// resumes where the last left off (timestamp / cursor / file marker).
	PollerCursor datatypes.JSON `gorm:"column:pollerCursor" json:"pollerCursor"`
	// Phase 10 - number of consecutive errors observed by the puller
	// worker. Reset to 0 on a successful run. Drives admin-UI banners
	// ("Microsoft authentication failed - re-authorize") + back-off
	// scheduling when the source has been broken for a while.
	// Spec: specs/ai-governance/puller-framework/puller-adapter-contract.feature
	ErrorCount int `gorm:"column:errorCount;size:32;default:0;not null" json:"errorCount"`
	// Phase 10 - cron schedule for pull-mode sources. Validated by the
	// adapter's pullConfig schema; mirrored here for fast worker
	// dispatch. Null for push-mode sources.
	PullSchedule *string `gorm:"column:pullSchedule" json:"pullSchedule"`
	// Status flag. 'active' (default), 'disabled' (admin paused), or
	// 'awaiting_first_event' (created but never received).
	Status *string `gorm:"column:status;default:'awaiting_first_event';not null" json:"status"`
	// ADR-088 v7: optional trace DESTINATION for conversation-bearing pulls
	// (Databricks Genie). Null = don't route. Same shape as
	// VirtualKey.traceProjectId: a destination is NOT a scope row and grants
	// no access to the source; it only decides where routed conversation
	// traces land. Must be a project of this org (validated at write time).
	// Freely editable; already-routed history stays where it landed.
	// NOT the deferred cost-attribution projectId (ADR-088 Decision 4).
	TraceProjectID *string `gorm:"column:traceProjectId" json:"traceProjectId"`
	// The account the provider itself named when this connection was saved -
	// an OpenAI or Anthropic organization id, read once per save and compared
	// so a second connection cannot read one account's spend twice. NOT
	// derived from the administrator key: nothing worked out from a customer
	// credential is stored here or anywhere else, because the account id is
	// not a secret and cannot be turned back into a key.
	// Nullable because every connection that predates the guard has none, and
	// because Azure and Power Platform name what they read in parserConfig
	// instead. "Unknown account" has to be a storable state.
	// Spec: specs/ai-gateway/governance/ingestion-sources.feature
	ProviderAccountID *string `gorm:"column:providerAccountId" json:"providerAccountId"`
	// Last event timestamp observed (any direction). Surfaces in the
	// admin UI as "last event 2 minutes ago".
	LastEventAt *DateTime `gorm:"column:lastEventAt;type:timestamp(3)" json:"lastEventAt"`
	// ADR-128 - when a pull run last reached the provider without erroring,
	// mirrored from the ingestionPullRunStatus projection. Deliberately NOT
	// lastEventAt: an empty period is a working puller, and a run that
	// delivered data and then failed is not. Drives the health badge and
	// the "no data since" cut-off, below which a day's spend is unknown
	// rather than zero.
	// Spec: specs/governance/ingestion-source-health.feature
	LastSuccessAt *DateTime `gorm:"column:lastSuccessAt;type:timestamp(3)" json:"lastSuccessAt"`
	// ADR-128 - how far the last run actually read, and whether it reached the
	// end. Mirrored from the ingestionPullRunStatus projection, in the same way
	// and for the same reason as lastSuccessAt above: every reader of this
	// surface selects from this row, and two of them render a LIST of sources,
	// so reaching a per-run projection would be one query per row on a screen
	// built to show many at once.
	//
	// These two are the only evidence separating a source that truncated once
	// and then finished from one that is stuck. A stuck source's run clock
	// advances on every attempt, so the clock cannot be that evidence and is
	// deliberately not stored here - only the point the run REACHED.
	//
	// Completeness is "complete" or "truncated". Null on both is unknown, which
	// is what every row written before this existed reads as, and unknown must
	// never render as either answer.
	// Spec: specs/governance/ingestion-source-health.feature
	LastReadThroughAt   *DateTime `gorm:"column:lastReadThroughAt;type:timestamp(3)" json:"lastReadThroughAt"`
	LastRunCompleteness *string   `gorm:"column:lastRunCompleteness" json:"lastRunCompleteness"`
	// ADR-088 - the window this source read while pulled cost recording was
	// off for the organization, so its spend was dropped rather than priced.
	// Both null = nothing dropped. The pull cursor advances whether or not the
	// money path is live (audit-only is a supported mode), so without this the
	// loss is unrecoverable AND invisible: those days read as no spend rather
	// than as unknown. Widened by any run that drops a priced item, cleared
	// only when a later run re-reads back across the whole window.
	// Spec: specs/governance/pulled-usage-cost-gap.feature
	UnpricedUsageSince   *DateTime     `gorm:"column:unpricedUsageSince;type:timestamp(3)" json:"unpricedUsageSince"`
	UnpricedUsageThrough *DateTime     `gorm:"column:unpricedUsageThrough;type:timestamp(3)" json:"unpricedUsageThrough"`
	ArchivedAt           *DateTime     `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	CreatedAt            DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID          *string       `gorm:"column:createdById" json:"createdById"`
	Organization         *Organization `gorm:"belongsTo"`
	Team                 *Team         `gorm:"belongsTo"`
	CreatedBy            *User         `gorm:"foreignKey:CreatedByID;belongsTo"`
}

func (IngestionSource) TableName() string {
	return "IngestionSource"
}

func (m *IngestionSource) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*IngestionSource) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Admin-defined anomaly detection rule. Evaluated server-side at
// ingest time (Option C - not in this slice yet) against
// gateway_activity_events. When triggered, an alert is dispatched
// to the destinations configured on the rule (Slack / SIEM / webhook
// / PagerDuty / email).
//
// This is a configuration entity - the actual rule evaluation +
// alert state lives in a future ClickHouse table
// (gateway_anomaly_alerts) and a server-side worker. Spec:
// specs/ai-gateway/governance/anomaly-rules.feature
// (Alexis lane B C-UI; mocked-v0 today, real persistence is THIS slice).
type AnomalyRule struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:AnomalyRule_organizationId_name_key,priority:1;index:AnomalyRule_organizationId_idx;not null" json:"organizationId"`
	// 'organization' | 'team' | 'project' | 'source_type' | 'source'.
	// String not enum so new scopes ship without migrations.
	Scope string `gorm:"column:scope;index:AnomalyRule_scope_scopeId_idx,priority:1;not null" json:"scope"`
	// Org id / team id / project id / source-type string / source id
	// - narrows what the rule applies to.
	ScopeID     string  `gorm:"column:scopeId;index:AnomalyRule_scope_scopeId_idx,priority:2;not null" json:"scopeId"`
	Name        string  `gorm:"column:name;uniqueIndex:AnomalyRule_organizationId_name_key,priority:2;not null" json:"name"`
	Description *string `gorm:"column:description" json:"description"`
	// 'critical' | 'warning' | 'info'. Drives default destination
	// routing (e.g. critical → PagerDuty, warning → Slack only).
	Severity string `gorm:"column:severity;not null" json:"severity"`
	// Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
	// 'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
	// | 'custom'. New types ship without migrations - evaluator
	// picks a per-type implementation from a code registry.
	RuleType string `gorm:"column:ruleType;index:AnomalyRule_ruleType_idx;not null" json:"ruleType"`
	// Per-rule-type threshold shape. Examples:
	// spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
	// rate_limit: { windowSec, requestsThreshold }
	// after_hours: { startHour, endHour, timezone, requestsThreshold }
	ThresholdConfig datatypes.JSON `gorm:"column:thresholdConfig;default:'{}';not null" json:"thresholdConfig"`
	// Alert routing per destination. Shape:
	// { slack?: { channelId }, siem?: { endpointUrl, authHeader },
	// webhook?: { url, authHeader }, pagerduty?: { routingKey },
	// email?: { to: string[] } }
	// Empty map = use org defaults (when defaults exist; today no-op).
	DestinationConfig datatypes.JSON `gorm:"column:destinationConfig;default:'{}';not null" json:"destinationConfig"`
	// 'active' (default - evaluator picks it up) | 'disabled' (admin
	// paused; rule rows kept for audit). 'archived' uses archivedAt
	// timestamp instead so we can filter `archivedAt IS NULL` cheaply.
	Status       *string        `gorm:"column:status;index:AnomalyRule_status_idx;default:'active';not null" json:"status"`
	ArchivedAt   *DateTime      `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	CreatedAt    DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID  *string        `gorm:"column:createdById" json:"createdById"`
	Organization *Organization  `gorm:"belongsTo"`
	CreatedBy    *User          `gorm:"foreignKey:CreatedByID;belongsTo"`
	Alerts       []AnomalyAlert `gorm:"foreignKey:RuleID;constraint:AnomalyAlert_ruleId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AnomalyRule) TableName() string {
	return "AnomalyRule"
}

func (m *AnomalyRule) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AnomalyRule) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Persisted anomaly detection. Inserted by the anomaly-detection
// reactor (event-sourcing pipeline; modeled on PR #3351's
// alertTrigger pattern) when a rule's threshold trips while
// processing a freshly-arrived ActivityEventReceived event.
// Read by api.activityMonitor.recentAnomalies for the admin dashboard.
//
// Dedup: (ruleId, triggerWindowStart) - re-evaluating the same
// window updates the existing row rather than creating a duplicate.
// Acknowledgement / resolution are admin actions on the row.
//
// Spec: specs/ai-gateway/governance/anomaly-detection.feature
type AnomalyAlert struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:AnomalyAlert_organizationId_detectedAt_idx,priority:1;not null" json:"organizationId"`
	RuleID         string `gorm:"column:ruleId;uniqueIndex:AnomalyAlert_ruleId_triggerWindowStart_key,priority:1;index:AnomalyAlert_ruleId_idx;not null" json:"ruleId"`
	// Denormalised from rule for fast WHERE filtering on the
	// dashboard query (no join needed for the common path).
	Severity string `gorm:"column:severity;not null" json:"severity"`
	RuleName string `gorm:"column:ruleName;not null" json:"ruleName"`
	RuleType string `gorm:"column:ruleType;not null" json:"ruleType"`
	// Window the trigger refers to (UTC).
	TriggerWindowStart DateTime `gorm:"column:triggerWindowStart;uniqueIndex:AnomalyAlert_ruleId_triggerWindowStart_key,priority:2;type:timestamp(3);not null" json:"triggerWindowStart"`
	TriggerWindowEnd   DateTime `gorm:"column:triggerWindowEnd;type:timestamp(3);not null" json:"triggerWindowEnd"`
	// Numeric trigger context. Per-rule-type detail in `detail` JSONB.
	TriggerSpendUsd   *float64 `gorm:"column:triggerSpendUsd;type:decimal(18,6)" json:"triggerSpendUsd"`
	TriggerEventCount *int     `gorm:"column:triggerEventCount;size:32" json:"triggerEventCount"`
	// Per-rule-type extra context: {topActors: [...], baseline: ...,
	// observed: ..., timezone: ..., etc}.
	Detail datatypes.JSON `gorm:"column:detail;default:'{}';not null" json:"detail"`
	// 'open' | 'acknowledged' | 'resolved'.
	State *string `gorm:"column:state;index:AnomalyAlert_state_idx;default:'open';not null" json:"state"`
	// Per-destination dispatch outcome:
	// { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
	// slack: ..., pagerduty: ..., siem: ..., email: ... }
	DestinationStatus datatypes.JSON `gorm:"column:destinationStatus;default:'{}';not null" json:"destinationStatus"`
	DetectedAt        DateTime       `gorm:"column:detectedAt;index:AnomalyAlert_organizationId_detectedAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"detectedAt"`
	AcknowledgedAt    *DateTime      `gorm:"column:acknowledgedAt;type:timestamp(3)" json:"acknowledgedAt"`
	ResolvedAt        *DateTime      `gorm:"column:resolvedAt;type:timestamp(3)" json:"resolvedAt"`
	Organization      *Organization  `gorm:"belongsTo"`
	Rule              *AnomalyRule   `gorm:"foreignKey:RuleID;belongsTo"`
}

func (AnomalyAlert) TableName() string {
	return "AnomalyAlert"
}

func (m *AnomalyAlert) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.DetectedAt.IsZero() {
		m.DetectedAt = DateTime{Time: now}
	}
	return nil
}

// Org-scoped catalog of AI tools surfaced on the personal home portal
// (/me) and managed by org admins at /governance/tool-catalog.
//
// Three first-class tile types via `type` discriminator:
// - "coding_assistant" - click expands to setup helper (e.g. `langwatch claude`)
// - "model_provider"   - click expands to inline personal-VK creation
// - "external_tool"    - click expands to admin-attached markdown + external link
//
// Scoping: organization-default entries can be OVERRIDDEN at the team
// tier (same `slug` + `scope=team` + matching `scopeId`) so admins can
// hide or re-skin a globally-listed tool for a specific team. The
// user-facing `list` endpoint resolves this hierarchy at read time.
//
// `config` is a discriminated JSON union per `type`:
// coding_assistant: { setupCommand, setupDocsUrl, helperText? }
// model_provider:   { providerKey, suggestedRoutingPolicyId?, defaultLabel?, projectSuggestionText? }
// external_tool:    { descriptionMarkdown, linkUrl, ctaLabel? }
//
// Spec: specs/ai-governance/personal-portal/tool-catalog-*.feature
type AiToolEntry struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:AiToolEntry_organizationId_scope_scopeId_idx,priority:1;index:AiToolEntry_organizationId_enabled_archivedAt_idx,priority:1;index:AiToolEntry_organizationId_type_idx,priority:1;not null" json:"organizationId"`
	// "organization" | "team". String not enum so new scopes ship
	// without migrations.
	Scope string `gorm:"column:scope;index:AiToolEntry_organizationId_scope_scopeId_idx,priority:2;not null" json:"scope"`
	// organizationId or teamId, narrowing what scope this entry applies
	// to. Always equals organizationId when scope=organization.
	ScopeID string `gorm:"column:scopeId;index:AiToolEntry_organizationId_scope_scopeId_idx,priority:3;not null" json:"scopeId"`
	// "coding_assistant" | "model_provider" | "external_tool".
	// Discriminator drives the per-tile UX shape and validates
	// the `config` JSON column at the app layer.
	Type        string `gorm:"column:type;index:AiToolEntry_organizationId_type_idx,priority:2;not null" json:"type"`
	DisplayName string `gorm:"column:displayName;not null" json:"displayName"`
	// Stable slug for team-overrides-org resolution (a team-scoped
	// entry with the same slug as an org-scoped one hides the org
	// default for users in that team). Server-owned: generated from
	// displayName + nanoid(6) on create, immutable on update.
	Slug string `gorm:"column:slug;not null" json:"slug"`
	// @deprecated Use `iconAsset` instead. Kept for back-compat with
	// pre-refactor rows; new writes should populate iconAsset.
	IconKey *string `gorm:"column:iconKey" json:"iconKey"`
	// Icon source - single column, prefix-discriminated:
	// "preset:claude_code" / "preset:codex" / ...   (built-in)
	// "data:image/svg+xml;base64,..."               (admin-uploaded)
	// null                                          → fall back to
	// type-default in UI
	IconAsset *string `gorm:"column:iconAsset;type:text" json:"iconAsset"`
	// Admin-curated sort within the user-facing portal grid.
	Order int `gorm:"column:order;size:32;default:0;not null" json:"order"`
	// When false the entry is hidden from the user-facing list but
	// remains in the admin list (so admins can re-enable without
	// losing per-entry config). Disabled is "soft-hide", archived is
	// "soft-delete" (preserves audit trail + scope binding history).
	Enabled *bool `gorm:"column:enabled;index:AiToolEntry_organizationId_enabled_archivedAt_idx,priority:2;default:true;not null" json:"enabled"`
	// Per-type config JSON. Validated via zod discriminated union
	// at the service layer (see AiToolEntryService).
	Config       datatypes.JSON `gorm:"column:config;default:'{}';not null" json:"config"`
	ArchivedAt   *DateTime      `gorm:"column:archivedAt;index:AiToolEntry_organizationId_enabled_archivedAt_idx,priority:3;type:timestamp(3)" json:"archivedAt"`
	CreatedAt    DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID  *string        `gorm:"column:createdById" json:"createdById"`
	UpdatedByID  *string        `gorm:"column:updatedById" json:"updatedById"`
	Organization *Organization  `gorm:"belongsTo"`
	// @deprecated Multi-team scope. Superseded by `departments` - the
	// "Visible to" model is now WHOLE ORGANIZATION or a set of
	// DEPARTMENTS. Rows are no longer written here (the table is kept
	// for now to avoid a destructive drop); readers prefer `departments`.
	Teams []AiToolEntryTeam `gorm:"foreignKey:EntryID;constraint:AiToolEntryTeam_entryId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Department scope. Empty = org-wide (visible to every member).
	// Non-empty = visible only to members whose `departmentId` is in
	// this set. Mirrored onto the legacy scope/scopeId pair for
	// back-compat reads (scope='organization' when empty,
	// scope='department' otherwise).
	Departments []AiToolEntryDepartment `gorm:"foreignKey:EntryID;constraint:AiToolEntryDepartment_entryId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AiToolEntry) TableName() string {
	return "AiToolEntry"
}

func (m *AiToolEntry) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AiToolEntry) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// @deprecated Per-team scope binding for AiToolEntry. Superseded by
// AiToolEntryDepartment; kept so the rename migration is non-destructive.
// No new rows are written here.
type AiToolEntryTeam struct {
	ID      string       `gorm:"column:id;primaryKey" json:"id"`
	EntryID string       `gorm:"column:entryId;uniqueIndex:AiToolEntryTeam_entryId_teamId_key,priority:1;not null" json:"entryId"`
	TeamID  string       `gorm:"column:teamId;uniqueIndex:AiToolEntryTeam_entryId_teamId_key,priority:2;index:AiToolEntryTeam_teamId_idx;not null" json:"teamId"`
	Entry   *AiToolEntry `gorm:"foreignKey:EntryID;belongsTo"`
	Team    *Team        `gorm:"belongsTo"`
}

func (AiToolEntryTeam) TableName() string {
	return "AiToolEntryTeam"
}

func (m *AiToolEntryTeam) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Per-department scope binding for AiToolEntry. Empty set = org-wide.
// Non-empty = entry visible only to members of those departments (an
// OrganizationUser whose `departmentId` is in the set). Mirrors
// AiToolEntryTeam's shape. `departmentId` is a plain column ref
// (relationMode="prisma"), matching how Department is referenced
// elsewhere; the service validates the id belongs to the org.
type AiToolEntryDepartment struct {
	ID           string       `gorm:"column:id;primaryKey" json:"id"`
	EntryID      string       `gorm:"column:entryId;uniqueIndex:AiToolEntryDepartment_entryId_departmentId_key,priority:1;not null" json:"entryId"`
	DepartmentID string       `gorm:"column:departmentId;uniqueIndex:AiToolEntryDepartment_entryId_departmentId_key,priority:2;index:AiToolEntryDepartment_departmentId_idx;not null" json:"departmentId"`
	Entry        *AiToolEntry `gorm:"foreignKey:EntryID;belongsTo"`
}

func (AiToolEntryDepartment) TableName() string {
	return "AiToolEntryDepartment"
}

func (m *AiToolEntryDepartment) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Admin/platform-curated catalog row for upstream OTLP normalization.
// Owns the OTTL transform that reshapes a specific upstream tool's
// span shape into canonical gen_ai conventions. Users connect via an
// ingestion key (an ingest-only ApiKey carrying ingestionTemplateId) to
// land traces in their personal project.
//
// `organizationId` NULL = platform-published default (always-visible
// across every org). NOT NULL = org-authored (visible only within
// that org). v1 ships only platform-published; org-authored authoring
// UI lands v2 (the column ships now to avoid a future migration).
//
// Spec: specs/ai-gateway/governance/ingestion-templates-catalog.feature
type IngestionTemplate struct {
	ID             string  `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID *string `gorm:"column:organizationId;uniqueIndex:IngestionTemplate_organizationId_slug_key,priority:1;index:IngestionTemplate_organizationId_archivedAt_idx,priority:1" json:"organizationId"`
	// Stable slug - unique within (organizationId, slug) scope. Generated
	// from displayName + nanoid(6) on create for org-authored; canonical
	// values for platform-published (`claude_code` / `cursor` /
	// `claude_cowork` / `raw_otlp_advanced`).
	Slug string `gorm:"column:slug;uniqueIndex:IngestionTemplate_organizationId_slug_key,priority:2;not null" json:"slug"`
	// Mirrors the slug for v1 platform-published rows. Kept as a separate
	// column so a future v2 may publish two templates that share an
	// upstream sourceType (e.g. claude_code-personal vs claude_code-team)
	// without a migration. String not enum - extensibility.
	SourceType  string  `gorm:"column:sourceType;index:IngestionTemplate_sourceType_idx;not null" json:"sourceType"`
	DisplayName string  `gorm:"column:displayName;not null" json:"displayName"`
	Description *string `gorm:"column:description;type:text" json:"description"`
	// Icon source - single column, prefix-discriminated, mirrors
	// AiToolEntry.iconAsset:
	// `preset:<kind>`                 (built-in slot)
	// `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
	// null                            → fall back to slug-default in UI
	IconAsset *string `gorm:"column:iconAsset;type:text" json:"iconAsset"`
	// Discriminator for what (if any) extra credential metadata the
	// install drawer captures from the user:
	// null              → otlp_token-only (drawer auto-issues the
	// ingestion key, no input form)
	// `static_api_key`  → drawer captures the user's upstream tool API key
	// `agent_id`        → drawer captures an agent identifier (e.g.
	// Cursor agentId)
	// String not enum so v2 shapes (oauth_redirect / s3_bucket) ship
	// without a migration.
	CredentialSchema *string `gorm:"column:credentialSchema" json:"credentialSchema"`
	// Admin/platform-authored OTTL statements. Applied at receive time
	// AFTER the receiver re-stamps the binding's authoritative principal
	// + provenance keys (see `protectedTemplateAttributeKeys` 19-key
	// closed list). Templates write canonical OTTL outputs (gen_ai.*,
	// cost.usd, response.model) - they CANNOT write attribution or
	// provenance keys; the receiver guard restores those post-OTTL.
	OttlRules string `gorm:"column:ottlRules;type:text;default:'';not null" json:"ottlRules"`
	// True for the v1 platform-published default templates
	// (organizationId IS NULL). Future org-authored rows write false.
	// Drives the admin UI "READ-ONLY view + Request a custom template"
	// surface - admins consume + opt-out only in v1.
	PlatformPublished bool `gorm:"column:platformPublished;default:false;not null" json:"platformPublished"`
	// Admin/platform enable. False = hidden from /me catalog. Existing
	// ingestion keys continue to land traces; new installs are blocked.
	Enabled      *bool         `gorm:"column:enabled;default:true;not null" json:"enabled"`
	ArchivedAt   *DateTime     `gorm:"column:archivedAt;index:IngestionTemplate_organizationId_archivedAt_idx,priority:2;type:timestamp(3)" json:"archivedAt"`
	CreatedAt    DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID  *string       `gorm:"column:createdById" json:"createdById"`
	UpdatedByID  *string       `gorm:"column:updatedById" json:"updatedById"`
	Organization *Organization `gorm:"belongsTo"`
}

func (IngestionTemplate) TableName() string {
	return "IngestionTemplate"
}

func (m *IngestionTemplate) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*IngestionTemplate) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GatewayA2ATarget struct {
	ID                       string             `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID           string             `gorm:"column:organizationId;index:GatewayA2ATarget_organizationId_projectId_idx,priority:1;not null" json:"organizationId"`
	ProjectID                *string            `gorm:"column:projectId;index:GatewayA2ATarget_organizationId_projectId_idx,priority:2" json:"projectId"`
	Name                     string             `gorm:"column:name;not null" json:"name"`
	AgentName                string             `gorm:"column:agentName;not null" json:"agentName"`
	CardURL                  string             `gorm:"column:cardUrl;not null" json:"cardUrl"`
	UpstreamURL              string             `gorm:"column:upstreamUrl;not null" json:"upstreamUrl"`
	EncryptedUpstreamHeaders *string            `gorm:"column:encryptedUpstreamHeaders" json:"encryptedUpstreamHeaders"`
	RestrictToAllowList      bool               `gorm:"column:restrictToAllowList;default:false;not null" json:"restrictToAllowList"`
	AllowedVirtualKeyIds     ScalarList[string] `gorm:"column:allowedVirtualKeyIds;type:text[]" json:"allowedVirtualKeyIds"`
	Blocked                  bool               `gorm:"column:blocked;default:false;not null" json:"blocked"`
	BlockReason              *string            `gorm:"column:blockReason" json:"blockReason"`
	CreatedByID              string             `gorm:"column:createdById;not null" json:"createdById"`
	CreatedAt                DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	ArchivedAt               *DateTime          `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
}

func (GatewayA2ATarget) TableName() string {
	return "GatewayA2ATarget"
}

func (m *GatewayA2ATarget) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GatewayA2ATarget) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GatewayConnectUpstream struct {
	OrganizationID string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	BaseURL        string   `gorm:"column:baseUrl;not null" json:"baseUrl"`
	EncryptedToken string   `gorm:"column:encryptedToken;not null" json:"encryptedToken"`
	InstanceID     string   `gorm:"column:instanceId;not null" json:"instanceId"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (GatewayConnectUpstream) TableName() string {
	return "GatewayConnectUpstream"
}

func (m *GatewayConnectUpstream) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GatewayConnectUpstream) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GatewayTraceExportKey struct {
	ProjectID      string   `gorm:"column:projectId;primaryKey" json:"projectId"`
	APIKeyID       string   `gorm:"column:apiKeyId;not null" json:"apiKeyId"`
	EncryptedToken string   `gorm:"column:encryptedToken;not null" json:"encryptedToken"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (GatewayTraceExportKey) TableName() string {
	return "GatewayTraceExportKey"
}

func (m *GatewayTraceExportKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type GatewayBudget struct {
	ID                     string                        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID         string                        `gorm:"column:organizationId;uniqueIndex:GatewayBudget_organizationId_externalId_key,priority:1;index:GatewayBudget_organizationId_idx;not null" json:"organizationId"`
	ScopeType              string                        `gorm:"column:scopeType;index:GatewayBudget_scopeType_scopeId_providerKey_idx,priority:1;not null" json:"scopeType"`
	ScopeID                string                        `gorm:"column:scopeId;index:GatewayBudget_scopeType_scopeId_providerKey_idx,priority:2;not null" json:"scopeId"`
	ProviderKey            *string                       `gorm:"column:providerKey;index:GatewayBudget_scopeType_scopeId_providerKey_idx,priority:3" json:"providerKey"`
	Name                   string                        `gorm:"column:name;not null" json:"name"`
	Description            *string                       `gorm:"column:description" json:"description"`
	Window                 string                        `gorm:"column:window;not null" json:"window"`
	LimitUsd               float64                       `gorm:"column:limitUsd;type:decimal(18,6);not null" json:"limitUsd"`
	OnBreach               string                        `gorm:"column:onBreach;default:'BLOCK';not null" json:"onBreach"`
	Timezone               *string                       `gorm:"column:timezone" json:"timezone"`
	ExternalID             *string                       `gorm:"column:externalId;uniqueIndex:GatewayBudget_organizationId_externalId_key,priority:2" json:"externalId"`
	Metadata               datatypes.JSON                `gorm:"column:metadata;default:'{}';not null" json:"metadata"`
	SpentUsd               float64                       `gorm:"column:spentUsd;type:decimal(18,6);default:0;not null" json:"spentUsd"`
	CurrentPeriodStartedAt DateTime                      `gorm:"column:currentPeriodStartedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"currentPeriodStartedAt"`
	ResetsAt               DateTime                      `gorm:"column:resetsAt;index:GatewayBudget_resetsAt_idx;type:timestamp(3);not null" json:"resetsAt"`
	LastResetAt            *DateTime                     `gorm:"column:lastResetAt;type:timestamp(3)" json:"lastResetAt"`
	CycleAnchorAt          *DateTime                     `gorm:"column:cycleAnchorAt;type:timestamp(3)" json:"cycleAnchorAt"`
	ArchivedAt             *DateTime                     `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	CreatedAt              DateTime                      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt              DateTime                      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID            string                        `gorm:"column:createdById;not null" json:"createdById"`
	ManagedByVirtualKeyID  *string                       `gorm:"column:managedByVirtualKeyId;index:GatewayBudget_managedByVirtualKeyId_idx" json:"managedByVirtualKeyId"`
	CreatedBy              *User                         `gorm:"foreignKey:CreatedByID;belongsTo"`
	LedgerEntries          []GatewayBudgetLedger         `gorm:"foreignKey:BudgetID;constraint:GatewayBudgetLedger_budgetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BucketBoundaries       []GatewayBudgetBucketBoundary `gorm:"foreignKey:BudgetID;constraint:GatewayBudgetBucketBoundary_budgetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (GatewayBudget) TableName() string {
	return "GatewayBudget"
}

func (m *GatewayBudget) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CurrentPeriodStartedAt.IsZero() {
		m.CurrentPeriodStartedAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GatewayBudget) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GatewayBudgetBucketBoundary struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID  string         `gorm:"column:organizationId;index:GatewayBudgetBucketBoundary_organizationId_idx;not null" json:"organizationId"`
	BudgetID        string         `gorm:"column:budgetId;uniqueIndex:GatewayBudgetBucketBoundary_budgetId_bucketScopeId_key,priority:1;not null" json:"budgetId"`
	BucketScopeID   string         `gorm:"column:bucketScopeId;uniqueIndex:GatewayBudgetBucketBoundary_budgetId_bucketScopeId_key,priority:2;not null" json:"bucketScopeId"`
	PeriodStartedAt DateTime       `gorm:"column:periodStartedAt;type:timestamp(3);not null" json:"periodStartedAt"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Budget          *GatewayBudget `gorm:"foreignKey:BudgetID;belongsTo"`
}

func (GatewayBudgetBucketBoundary) TableName() string {
	return "GatewayBudgetBucketBoundary"
}

func (m *GatewayBudgetBucketBoundary) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GatewayBudgetBucketBoundary) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GatewayBudgetLedger struct {
	ID               string         `gorm:"column:id;primaryKey" json:"id"`
	BudgetID         string         `gorm:"column:budgetId;uniqueIndex:GatewayBudgetLedger_budgetId_gatewayRequestId_key,priority:1;not null" json:"budgetId"`
	VirtualKeyID     string         `gorm:"column:virtualKeyId;index:GatewayBudgetLedger_virtualKeyId_idx;not null" json:"virtualKeyId"`
	ModelProviderID  *string        `gorm:"column:modelProviderId" json:"modelProviderId"`
	ProviderKey      *string        `gorm:"column:providerKey" json:"providerKey"`
	GatewayRequestID string         `gorm:"column:gatewayRequestId;uniqueIndex:GatewayBudgetLedger_budgetId_gatewayRequestId_key,priority:2;index:GatewayBudgetLedger_gatewayRequestId_idx;not null" json:"gatewayRequestId"`
	AmountUsd        float64        `gorm:"column:amountUsd;type:decimal(18,6);not null" json:"amountUsd"`
	TokensInput      int            `gorm:"column:tokensInput;size:32;default:0;not null" json:"tokensInput"`
	TokensOutput     int            `gorm:"column:tokensOutput;size:32;default:0;not null" json:"tokensOutput"`
	TokensCacheRead  int            `gorm:"column:tokensCacheRead;size:32;default:0;not null" json:"tokensCacheRead"`
	TokensCacheWrite int            `gorm:"column:tokensCacheWrite;size:32;default:0;not null" json:"tokensCacheWrite"`
	Model            string         `gorm:"column:model;not null" json:"model"`
	DurationMs       *int           `gorm:"column:durationMs;size:32" json:"durationMs"`
	Status           string         `gorm:"column:status;default:'SUCCESS';not null" json:"status"`
	OccurredAt       DateTime       `gorm:"column:occurredAt;index:GatewayBudgetLedger_occurredAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"occurredAt"`
	Budget           *GatewayBudget `gorm:"foreignKey:BudgetID;belongsTo"`
	VirtualKey       *VirtualKey    `gorm:"belongsTo"`
}

func (GatewayBudgetLedger) TableName() string {
	return "GatewayBudgetLedger"
}

func (m *GatewayBudgetLedger) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.OccurredAt.IsZero() {
		m.OccurredAt = DateTime{Time: now}
	}
	return nil
}

type GatewayChangeEvent struct {
	Revision        int64          `gorm:"column:revision;primaryKey;autoIncrement;index:GatewayChangeEvent_organizationId_revision_idx,priority:2;index:GatewayChangeEvent_projectId_revision_idx,priority:2" json:"revision"`
	OrganizationID  string         `gorm:"column:organizationId;index:GatewayChangeEvent_organizationId_revision_idx,priority:1;not null" json:"organizationId"`
	Kind            string         `gorm:"column:kind;not null" json:"kind"`
	VirtualKeyID    *string        `gorm:"column:virtualKeyId;index:GatewayChangeEvent_virtualKeyId_idx" json:"virtualKeyId"`
	BudgetID        *string        `gorm:"column:budgetId;index:GatewayChangeEvent_budgetId_idx" json:"budgetId"`
	ModelProviderID *string        `gorm:"column:modelProviderId" json:"modelProviderId"`
	ProjectID       *string        `gorm:"column:projectId;index:GatewayChangeEvent_projectId_revision_idx,priority:1" json:"projectId"`
	Payload         datatypes.JSON `gorm:"column:payload" json:"payload"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Organization    *Organization  `gorm:"belongsTo"`
	Project         *Project       `gorm:"belongsTo"`
}

func (GatewayChangeEvent) TableName() string {
	return "GatewayChangeEvent"
}

func (m *GatewayChangeEvent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type GatewayCacheRule struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string         `gorm:"column:organizationId;index:GatewayCacheRule_organizationId_archivedAt_priority_idx,priority:1;index:GatewayCacheRule_organizationId_enabled_priority_idx,priority:1;not null" json:"organizationId"`
	Name           string         `gorm:"column:name;not null" json:"name"`
	Description    *string        `gorm:"column:description" json:"description"`
	Priority       *int           `gorm:"column:priority;index:GatewayCacheRule_organizationId_archivedAt_priority_idx,priority:3;index:GatewayCacheRule_organizationId_enabled_priority_idx,priority:3;size:32;default:100;not null" json:"priority"`
	Enabled        *bool          `gorm:"column:enabled;index:GatewayCacheRule_organizationId_enabled_priority_idx,priority:2;default:true;not null" json:"enabled"`
	Matchers       datatypes.JSON `gorm:"column:matchers;not null" json:"matchers"`
	Action         datatypes.JSON `gorm:"column:action;not null" json:"action"`
	ModeEnum       string         `gorm:"column:modeEnum;not null" json:"modeEnum"`
	ArchivedAt     *DateTime      `gorm:"column:archivedAt;index:GatewayCacheRule_organizationId_archivedAt_priority_idx,priority:2;type:timestamp(3)" json:"archivedAt"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID    string         `gorm:"column:createdById;not null" json:"createdById"`
}

func (GatewayCacheRule) TableName() string {
	return "GatewayCacheRule"
}

func (m *GatewayCacheRule) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GatewayCacheRule) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type FeatureFlag struct {
	Key          string         `gorm:"column:key;primaryKey" json:"key"`
	Enabled      bool           `gorm:"column:enabled;not null" json:"enabled"`
	Rules        datatypes.JSON `gorm:"column:rules" json:"rules"`
	LastEditedBy *string        `gorm:"column:lastEditedBy" json:"lastEditedBy"`
	CreatedAt    DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (FeatureFlag) TableName() string {
	return "FeatureFlag"
}

func (m *FeatureFlag) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

// One experiment setting per (flag, subject). A USER row is that person's
// own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
// for that exact scope. No row means inherit.
type FeatureFlagExperimentSetting struct {
	FlagKey         string   `gorm:"column:flagKey;primaryKey" json:"flagKey"`
	SubjectType     string   `gorm:"column:subjectType;primaryKey;index:FeatureFlagExperimentSetting_subjectType_subjectId_idx,priority:1" json:"subjectType"`
	SubjectID       string   `gorm:"column:subjectId;primaryKey;index:FeatureFlagExperimentSetting_subjectType_subjectId_idx,priority:2" json:"subjectId"`
	Enabled         bool     `gorm:"column:enabled;not null" json:"enabled"`
	ChangedByUserID *string  `gorm:"column:changedByUserId" json:"changedByUserId"`
	CreatedAt       DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (FeatureFlagExperimentSetting) TableName() string {
	return "FeatureFlagExperimentSetting"
}

func (m *FeatureFlagExperimentSetting) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*FeatureFlagExperimentSetting) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Durable state for a long-running event-sourced process.
type ProcessManagerInstance struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	ProcessName string         `gorm:"column:processName;uniqueIndex:ProcessManagerInstance_processName_projectId_processKey_key,priority:1;not null" json:"processName"`
	ProjectID   string         `gorm:"column:projectId;uniqueIndex:ProcessManagerInstance_processName_projectId_processKey_key,priority:2;not null" json:"projectId"`
	ProcessKey  string         `gorm:"column:processKey;uniqueIndex:ProcessManagerInstance_processName_projectId_processKey_key,priority:3;not null" json:"processKey"`
	TenantID    string         `gorm:"column:tenantId;not null" json:"tenantId"`
	UserID      *string        `gorm:"column:userId" json:"userId"`
	State       datatypes.JSON `gorm:"column:state;not null" json:"state"`
	Revision    int            `gorm:"column:revision;size:32;not null" json:"revision"`
	NextWakeAt  *DateTime      `gorm:"column:nextWakeAt;index:ProcessManagerInstance_nextWakeAt_idx;type:timestamptz(3)" json:"nextWakeAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamptz(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ProcessManagerInstance) TableName() string {
	return "ProcessManagerInstance"
}

func (m *ProcessManagerInstance) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Idempotency marker for a source event consumed by a process definition.
type ProcessManagerInbox struct {
	ID          string `gorm:"column:id;primaryKey" json:"id"`
	ProcessName string `gorm:"column:processName;uniqueIndex:ProcessManagerInbox_processName_projectId_sourceEventKey_key,priority:1;index:ProcessManagerInbox_processName_projectId_processKey_consum_idx,priority:1;not null" json:"processName"`
	ProjectID   string `gorm:"column:projectId;uniqueIndex:ProcessManagerInbox_processName_projectId_sourceEventKey_key,priority:2;index:ProcessManagerInbox_processName_projectId_processKey_consum_idx,priority:2;not null" json:"projectId"`
	ProcessKey  string `gorm:"column:processKey;index:ProcessManagerInbox_processName_projectId_processKey_consum_idx,priority:3;not null" json:"processKey"`
	TenantID    string `gorm:"column:tenantId;not null" json:"tenantId"`
	// The raw source event id, kept UNINDEXED for diagnostics. A caller composes
	// this and its length is that domain's business, so it must never carry a
	// btree constraint — see `sourceEventKey`.
	SourceEventID string `gorm:"column:sourceEventId;not null" json:"sourceEventId"`
	// The fixed-width digest of `sourceEventId` the uniqueness is enforced on,
	// derived by the store (`stores/inboxKey.ts`). Postgres cannot index a btree
	// row past ~2704 bytes; keying on a digest makes the constraint the same
	// width no matter what any pipeline concatenates into an idempotency key.
	//
	// Nullable ONLY for the rolling deploy that introduces it — pods running the
	// previous build write rows without it, and a NOT NULL would fail their
	// commits outright. The store always supplies it.
	//
	// The trade, stated plainly: a NULL is not constrained by the unique index
	// (btree treats NULLs as distinct), so while an old writer is serving, the
	// one case that loses its guard is the same source event reaching two
	// different process keys at once — the store's own read-check plus its
	// per-processKey advisory lock cover everything else. A follow-up migration
	// adds NOT NULL once the fleet has cycled and closes it.
	SourceEventKey *string  `gorm:"column:sourceEventKey;uniqueIndex:ProcessManagerInbox_processName_projectId_sourceEventKey_key,priority:3" json:"sourceEventKey"`
	ConsumedAt     DateTime `gorm:"column:consumedAt;index:ProcessManagerInbox_processName_projectId_processKey_consum_idx,priority:4;index:ProcessManagerInbox_consumedAt_idx;type:timestamptz(3);not null" json:"consumedAt"`
}

func (ProcessManagerInbox) TableName() string {
	return "ProcessManagerInbox"
}

func (m *ProcessManagerInbox) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Transactional intents emitted by a durable process transition.
type ProcessManagerOutbox struct {
	ID            string                        `gorm:"column:id;primaryKey" json:"id"`
	ProcessName   string                        `gorm:"column:processName;uniqueIndex:ProcessManagerOutbox_processName_projectId_messageKey_key,priority:1;index:ProcessManagerOutbox_processName_projectId_processKey_creat_idx,priority:1;index:ProcessManagerOutbox_status_processName_nextAttemptAt_idx,priority:2;not null" json:"processName"`
	ProjectID     string                        `gorm:"column:projectId;uniqueIndex:ProcessManagerOutbox_processName_projectId_messageKey_key,priority:2;index:ProcessManagerOutbox_processName_projectId_processKey_creat_idx,priority:2;not null" json:"projectId"`
	ProcessKey    string                        `gorm:"column:processKey;index:ProcessManagerOutbox_processName_projectId_processKey_creat_idx,priority:3;not null" json:"processKey"`
	TenantID      string                        `gorm:"column:tenantId;not null" json:"tenantId"`
	UserID        *string                       `gorm:"column:userId" json:"userId"`
	MessageKey    string                        `gorm:"column:messageKey;uniqueIndex:ProcessManagerOutbox_processName_projectId_messageKey_key,priority:3;not null" json:"messageKey"`
	IntentType    string                        `gorm:"column:intentType;not null" json:"intentType"`
	Payload       datatypes.JSON                `gorm:"column:payload;not null" json:"payload"`
	TraceCarrier  datatypes.JSON                `gorm:"column:traceCarrier;not null" json:"traceCarrier"`
	SourceEventID *string                       `gorm:"column:sourceEventId" json:"sourceEventId"`
	Status        string                        `gorm:"column:status;index:ProcessManagerOutbox_status_nextAttemptAt_leasedUntil_idx,priority:1;index:ProcessManagerOutbox_status_processName_nextAttemptAt_idx,priority:1;default:'pending';not null" json:"status"`
	Attempts      int                           `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	NextAttemptAt DateTime                      `gorm:"column:nextAttemptAt;index:ProcessManagerOutbox_status_nextAttemptAt_leasedUntil_idx,priority:2;index:ProcessManagerOutbox_status_processName_nextAttemptAt_idx,priority:3;type:timestamptz(3);not null" json:"nextAttemptAt"`
	LeasedUntil   *DateTime                     `gorm:"column:leasedUntil;index:ProcessManagerOutbox_status_nextAttemptAt_leasedUntil_idx,priority:3;type:timestamptz(3)" json:"leasedUntil"`
	LeaseToken    *string                       `gorm:"column:leaseToken;uniqueIndex:ProcessManagerOutbox_leaseToken_key" json:"leaseToken"`
	DispatchedAt  *DateTime                     `gorm:"column:dispatchedAt;type:timestamptz(3)" json:"dispatchedAt"`
	CreatedAt     DateTime                      `gorm:"column:createdAt;index:ProcessManagerOutbox_processName_projectId_processKey_creat_idx,priority:4;type:timestamptz(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime                      `gorm:"column:updatedAt;type:timestamptz(3);autoUpdateTime:false;not null" json:"updatedAt"`
	AttemptLog    []ProcessManagerOutboxAttempt `gorm:"foreignKey:OutboxID;constraint:ProcessManagerOutboxAttempt_outboxId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ProcessManagerOutbox) TableName() string {
	return "ProcessManagerOutbox"
}

func (m *ProcessManagerOutbox) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Why an outbox delivery failed, one row per FAILED attempt
// (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
// table grows with trouble, not with traffic. Rows die with their message:
// the cascade delete is the retention story, riding the existing sweep.
type ProcessManagerOutboxAttempt struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	OutboxID  string `gorm:"column:outboxId;index:ProcessManagerOutboxAttempt_outboxId_attempt_idx,priority:1;not null" json:"outboxId"`
	ProjectID string `gorm:"column:projectId;not null" json:"projectId"`
	// 1-based attempt number, as the dispatcher counts it.
	Attempt    int      `gorm:"column:attempt;index:ProcessManagerOutboxAttempt_outboxId_attempt_idx,priority:2;size:32;not null" json:"attempt"`
	OccurredAt DateTime `gorm:"column:occurredAt;type:timestamptz(3);not null" json:"occurredAt"`
	Outcome    string   `gorm:"column:outcome;not null" json:"outcome"`
	ErrorType  string   `gorm:"column:errorType;not null" json:"errorType"`
	// The safe failure diagnostic the dispatcher already computes for its
	// span — never a raw provider body.
	ErrorMessage string                `gorm:"column:errorMessage;not null" json:"errorMessage"`
	RetryAfterMs *int                  `gorm:"column:retryAfterMs;size:32" json:"retryAfterMs"`
	Outbox       *ProcessManagerOutbox `gorm:"foreignKey:OutboxID;belongsTo"`
}

func (ProcessManagerOutboxAttempt) TableName() string {
	return "ProcessManagerOutboxAttempt"
}

func (m *ProcessManagerOutboxAttempt) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Per-project topic clustering run status (ADR-051): the rebuildable
// operational read model behind the settings page. One row per project;
// folded from run_completed/run_failed events, never written directly.
type TopicClusteringRunProjection struct {
	// KSUID (`topicrun_…`), generated by the repository on first upsert.
	ID                 string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectID          string   `gorm:"column:projectId;uniqueIndex:TopicClusteringRunProjection_projectId_key;not null" json:"projectId"`
	LastRequestedAt    *float64 `gorm:"column:lastRequestedAt;type:double precision" json:"lastRequestedAt"`
	LastRequestTrigger *string  `gorm:"column:lastRequestTrigger" json:"lastRequestTrigger"`
	LastRunAt          *float64 `gorm:"column:lastRunAt;type:double precision" json:"lastRunAt"`
	// completed | skipped | failed
	LastRunOutcome       *string `gorm:"column:lastRunOutcome" json:"lastRunOutcome"`
	LastRunMode          *string `gorm:"column:lastRunMode" json:"lastRunMode"`
	LastRunSkippedReason *string `gorm:"column:lastRunSkippedReason" json:"lastRunSkippedReason"`
	LastRunError         *string `gorm:"column:lastRunError" json:"lastRunError"`
	LastRunErrorCode     *string `gorm:"column:lastRunErrorCode" json:"lastRunErrorCode"`
	// True when the customer can resolve the failure (credentials, quota).
	LastRunErrorUserActionable bool    `gorm:"column:lastRunErrorUserActionable;default:false;not null" json:"lastRunErrorUserActionable"`
	LastRunTracesProcessed     int     `gorm:"column:lastRunTracesProcessed;size:32;default:0;not null" json:"lastRunTracesProcessed"`
	LastRunTopicsCount         int     `gorm:"column:lastRunTopicsCount;size:32;default:0;not null" json:"lastRunTopicsCount"`
	LastRunSubtopicsCount      int     `gorm:"column:lastRunSubtopicsCount;size:32;default:0;not null" json:"lastRunSubtopicsCount"`
	LastRunPages               int     `gorm:"column:lastRunPages;size:32;default:0;not null" json:"lastRunPages"`
	InProgressRunID            *string `gorm:"column:inProgressRunId" json:"inProgressRunId"`
	InProgressTraces           int     `gorm:"column:inProgressTraces;size:32;default:0;not null" json:"inProgressTraces"`
	InProgressPages            int     `gorm:"column:inProgressPages;size:32;default:0;not null" json:"inProgressPages"`
	// Business time the in-progress run opened; bounds how long the read
	// model may report it as running (mirrors the process's stale-run guard).
	InProgressStartedAt *float64 `gorm:"column:inProgressStartedAt;type:double precision" json:"inProgressStartedAt"`
	CreatedAt           float64  `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt           float64  `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	// Business time of the latest applied event.
	OccurredAt float64 `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	// Event-log acceptance time. Together with lastEventId this is the cursor.
	AcceptedAt        float64 `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string  `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string  `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (TopicClusteringRunProjection) TableName() string {
	return "TopicClusteringRunProjection"
}

// Cursor row for the topic-model projection (ADR-051). The model itself is
// the `Topic` table, written through by the projection; this row only
// records how far the fold has applied the event log.
type TopicModelProjection struct {
	// KSUID (`topicmodel_…`), generated by the repository on first upsert.
	ID        string  `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string  `gorm:"column:projectId;uniqueIndex:TopicModelProjection_projectId_key;not null" json:"projectId"`
	CreatedAt float64 `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt float64 `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	// Business time of the latest applied event.
	OccurredAt float64 `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	// Event-log acceptance time. Together with lastEventId this is the cursor.
	AcceptedAt        float64 `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string  `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string  `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (TopicModelProjection) TableName() string {
	return "TopicModelProjection"
}

// Per-project topic clustering run history (audit read model, ADR-051).
// One row per project; `runs` is the bounded newest-first entry list.
// Rebuildable by replaying the event log.
type TopicClusteringRunHistoryProjection struct {
	// KSUID (`topicrunhist_…`), generated by the repository on first upsert.
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID string         `gorm:"column:projectId;uniqueIndex:TopicClusteringRunHistoryProjection_projectId_key;not null" json:"projectId"`
	Runs      datatypes.JSON `gorm:"column:runs;not null" json:"runs"`
	CreatedAt float64        `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt float64        `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	// Business time of the latest applied event.
	OccurredAt float64 `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	// Event-log acceptance time. Together with lastEventId this is the cursor.
	AcceptedAt        float64 `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID       string  `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion string  `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (TopicClusteringRunHistoryProjection) TableName() string {
	return "TopicClusteringRunHistoryProjection"
}

// Rebuildable per-source pull checkpoint and operator-facing run status.
type IngestionPullRunProjection struct {
	ID                string   `gorm:"column:id;primaryKey" json:"id"`
	SourceID          string   `gorm:"column:sourceId;uniqueIndex:IngestionPullRunProjection_sourceId_key;not null" json:"sourceId"`
	ProjectID         string   `gorm:"column:projectId;index:IngestionPullRunProjection_projectId_idx;not null" json:"projectId"`
	Enabled           bool     `gorm:"column:enabled;default:false;not null" json:"enabled"`
	Cron              *string  `gorm:"column:cron" json:"cron"`
	Cursor            *string  `gorm:"column:cursor" json:"cursor"`
	LastRunAt         *float64 `gorm:"column:lastRunAt;type:double precision" json:"lastRunAt"`
	LastRunOutcome    *string  `gorm:"column:lastRunOutcome" json:"lastRunOutcome"`
	LastRunEventCount int      `gorm:"column:lastRunEventCount;size:32;default:0;not null" json:"lastRunEventCount"`
	LastRunError      *string  `gorm:"column:lastRunError" json:"lastRunError"`
	LastRunErrorCode  *string  `gorm:"column:lastRunErrorCode" json:"lastRunErrorCode"`
	ConsecutiveErrors int      `gorm:"column:consecutiveErrors;size:32;default:0;not null" json:"consecutiveErrors"`
	// When a run last completed without erroring. Backfilled by replay, which
	// is why the projection version moved with this column.
	LastSuccessAt *float64 `gorm:"column:lastSuccessAt;type:double precision" json:"lastSuccessAt"`
	// Identity of the run whose outcome this row reflects. Outcomes from a
	// superseded run arrive with a smaller value and are ignored, so a late
	// completion cannot regress Cursor or the status fields.
	LastRunScheduledFor *float64 `gorm:"column:lastRunScheduledFor;type:double precision" json:"lastRunScheduledFor"`
	// ADR-128 - how far the last run read, in epoch milliseconds, and whether
	// it reached the end. The fold writes both; the IngestionSource row mirrors
	// them for the screens. Null read-through means no run has reported one.
	// Null completeness reads as unknown, never as either answer, which is what
	// every row written before this existed holds.
	//
	// Nullable on purpose, so replay is not required to make this migration
	// safe: a row that has never seen a run carrying these stays null and every
	// reader already handles unknown.
	// Spec: specs/governance/ingestion-source-health.feature
	LastReadThroughAt   *float64 `gorm:"column:lastReadThroughAt;type:double precision" json:"lastReadThroughAt"`
	LastRunCompleteness *string  `gorm:"column:lastRunCompleteness" json:"lastRunCompleteness"`
	// The last time somebody asked this source to enumerate its agents, and how
	// that ask ended. A listing is a different scope on the same connection as
	// the pull above: a credential that cannot enumerate a directory still pulls
	// cost perfectly, so no health, coverage or completeness reader may treat a
	// refusal here as the source failing to pull.
	//
	// Count is set ONLY when the outcome is listed, and stays null on a refusal.
	// Zero is a real answer meaning the provider returned an empty list, so a
	// zero default would manufacture that answer for a source that refused.
	// Reason is set only on a refusal and is a plain string, never an enum and
	// never a foreign key, because the log outlives the vocabulary. Status is
	// nullable even when refused, because a refusal we generated ourselves has
	// no HTTP status, and it is never rendered to a customer.
	//
	// The two `...ListingAt` columns hold the instant the ASK was accepted, not
	// the instant its answer was recorded, and the fold orders outcomes by them
	// so a superseded listing cannot overwrite the one that replaced it. The
	// name predates that job and is kept because renaming a column is a
	// migration; the fold's own field docs are the statement of record.
	LastAgentsListingAt      *float64 `gorm:"column:lastAgentsListingAt;type:double precision" json:"lastAgentsListingAt"`
	LastAgentsListingOutcome *string  `gorm:"column:lastAgentsListingOutcome" json:"lastAgentsListingOutcome"`
	LastAgentsListingCount   *int     `gorm:"column:lastAgentsListingCount;size:32" json:"lastAgentsListingCount"`
	LastAgentsListingReason  *string  `gorm:"column:lastAgentsListingReason" json:"lastAgentsListingReason"`
	LastAgentsListingStatus  *int     `gorm:"column:lastAgentsListingStatus;size:32" json:"lastAgentsListingStatus"`
	// The same for people, in its own set of columns rather than sharing the
	// agents set with a kind discriminator. A source can be refused for people
	// and fine for agents, and one shared set holds only the most recent listing
	// of either kind, so syncing people would erase an agents refusal and the
	// agents page would go back to not knowing after briefly knowing.
	//
	// LastPeopleDirectoryCount is everyone the provider's directory named.
	// LastPeopleWithheldCount is how many of those this deployment does not hold
	// because erasure suppression removed them: a SUBSET of the directory count,
	// never an addition to it. Adding the two counts the same people twice.
	// Subtracting is the valid arithmetic, and the surviving total is that
	// subtraction rather than a third column that could disagree.
	//
	// The withheld count is safe to show as a CURRENT figure and never as a
	// series: it moving from zero to one at a known moment says an erasure
	// happened then, which on a small tenant identifies the person as surely as
	// a name would. So no trend line, no history drawer, no per-run export, and
	// never beside a per-person list.
	LastPeopleListingAt      *float64 `gorm:"column:lastPeopleListingAt;type:double precision" json:"lastPeopleListingAt"`
	LastPeopleListingOutcome *string  `gorm:"column:lastPeopleListingOutcome" json:"lastPeopleListingOutcome"`
	LastPeopleDirectoryCount *int     `gorm:"column:lastPeopleDirectoryCount;size:32" json:"lastPeopleDirectoryCount"`
	LastPeopleWithheldCount  *int     `gorm:"column:lastPeopleWithheldCount;size:32" json:"lastPeopleWithheldCount"`
	LastPeopleListingReason  *string  `gorm:"column:lastPeopleListingReason" json:"lastPeopleListingReason"`
	LastPeopleListingStatus  *int     `gorm:"column:lastPeopleListingStatus;size:32" json:"lastPeopleListingStatus"`
	CreatedAt                float64  `gorm:"column:createdAt;type:double precision;not null" json:"createdAt"`
	UpdatedAt                float64  `gorm:"column:updatedAt;type:double precision;not null" json:"updatedAt"`
	OccurredAt               float64  `gorm:"column:occurredAt;type:double precision;not null" json:"occurredAt"`
	AcceptedAt               float64  `gorm:"column:acceptedAt;type:double precision;not null" json:"acceptedAt"`
	LastEventID              string   `gorm:"column:lastEventId;not null" json:"lastEventId"`
	ProjectionVersion        string   `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
}

func (IngestionPullRunProjection) TableName() string {
	return "IngestionPullRunProjection"
}

type ScheduledJob struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	ProjectID   string    `gorm:"column:projectId;index:ScheduledJob_projectId_idx;not null" json:"projectId"`
	TargetType  string    `gorm:"column:targetType;uniqueIndex:ScheduledJob_targetType_targetId_key,priority:1;not null" json:"targetType"`
	TargetID    string    `gorm:"column:targetId;uniqueIndex:ScheduledJob_targetType_targetId_key,priority:2;not null" json:"targetId"`
	Cron        string    `gorm:"column:cron;not null" json:"cron"`
	Timezone    string    `gorm:"column:timezone;not null" json:"timezone"`
	NextRunAt   DateTime  `gorm:"column:nextRunAt;index:ScheduledJob_active_nextRunAt_idx,priority:2;type:timestamp(3);not null" json:"nextRunAt"`
	LastSlot    *DateTime `gorm:"column:lastSlot;type:timestamp(3)" json:"lastSlot"`
	CurrentSlot *DateTime `gorm:"column:currentSlot;type:timestamp(3)" json:"currentSlot"`
	Attempts    int       `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	LastError   *string   `gorm:"column:lastError" json:"lastError"`
	Active      *bool     `gorm:"column:active;index:ScheduledJob_active_nextRunAt_idx,priority:1;default:true;not null" json:"active"`
	CreatedAt   DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ScheduledJob) TableName() string {
	return "ScheduledJob"
}

func (m *ScheduledJob) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ScheduledJob) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BugReport struct {
	ID               string         `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt        DateTime       `gorm:"column:createdAt;index:BugReport_createdAt_idx;type:timestamptz(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Source           string         `gorm:"column:source;not null" json:"source"`
	Kind             string         `gorm:"column:kind;not null" json:"kind"`
	Title            string         `gorm:"column:title;not null" json:"title"`
	Summary          *string        `gorm:"column:summary" json:"summary"`
	SessionData      *string        `gorm:"column:sessionData" json:"sessionData"`
	SessionTruncated bool           `gorm:"column:sessionTruncated;default:false;not null" json:"sessionTruncated"`
	Agent            *string        `gorm:"column:agent" json:"agent"`
	ContactEmail     *string        `gorm:"column:contactEmail" json:"contactEmail"`
	CliVersion       *string        `gorm:"column:cliVersion" json:"cliVersion"`
	LinkedProjectID  *string        `gorm:"column:linkedProjectId" json:"linkedProjectId"`
	Metadata         datatypes.JSON `gorm:"column:metadata" json:"metadata"`
}

func (BugReport) TableName() string {
	return "BugReport"
}

func (m *BugReport) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type IdempotencyReceipt struct {
	ID                 string   `gorm:"column:id;primaryKey" json:"id"`
	ScopeID            string   `gorm:"column:scopeId;uniqueIndex:IdempotencyReceipt_scopeId_key_key,priority:1;not null" json:"scopeId"`
	Key                string   `gorm:"column:key;uniqueIndex:IdempotencyReceipt_scopeId_key_key,priority:2;not null" json:"key"`
	ClaimID            string   `gorm:"column:claimId;not null" json:"claimId"`
	RequestFingerprint string   `gorm:"column:requestFingerprint;not null" json:"requestFingerprint"`
	ResponseStatus     *int     `gorm:"column:responseStatus;size:32" json:"responseStatus"`
	ResponseBody       *string  `gorm:"column:responseBody;type:text" json:"responseBody"`
	CreatedAt          DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	HeartbeatAt        DateTime `gorm:"column:heartbeatAt;type:timestamp(3);not null" json:"heartbeatAt"`
	ExpiresAt          DateTime `gorm:"column:expiresAt;index:IdempotencyReceipt_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
}

func (IdempotencyReceipt) TableName() string {
	return "IdempotencyReceipt"
}

func (m *IdempotencyReceipt) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type GatewayRealtimeSession struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	ProjectID            string         `gorm:"column:projectId;index:GatewayRealtimeSession_projectId_status_mintedAt_idx,priority:1;not null" json:"projectId"`
	OrganizationID       string         `gorm:"column:organizationId;not null" json:"organizationId"`
	VirtualKeyID         string         `gorm:"column:virtualKeyId;index:GatewayRealtimeSession_virtualKeyId_status_mintedAt_idx,priority:1;not null" json:"virtualKeyId"`
	ModelProviderID      string         `gorm:"column:modelProviderId;not null" json:"modelProviderId"`
	Vendor               string         `gorm:"column:vendor;uniqueIndex:GatewayRealtimeSession_vendor_vendorConversationId_key,priority:1;not null" json:"vendor"`
	AgentID              *string        `gorm:"column:agentId" json:"agentId"`
	Model                string         `gorm:"column:model;not null" json:"model"`
	RequestedModel       *string        `gorm:"column:requestedModel" json:"requestedModel"`
	VendorConversationID *string        `gorm:"column:vendorConversationId;uniqueIndex:GatewayRealtimeSession_vendor_vendorConversationId_key,priority:2" json:"vendorConversationId"`
	Status               string         `gorm:"column:status;index:GatewayRealtimeSession_virtualKeyId_status_mintedAt_idx,priority:2;index:GatewayRealtimeSession_projectId_status_mintedAt_idx,priority:2;index:GatewayRealtimeSession_status_metering_mintedAt_idx,priority:1;default:'OPEN';not null" json:"status"`
	MintedAt             DateTime       `gorm:"column:mintedAt;index:GatewayRealtimeSession_virtualKeyId_status_mintedAt_idx,priority:3;index:GatewayRealtimeSession_projectId_status_mintedAt_idx,priority:3;index:GatewayRealtimeSession_status_metering_mintedAt_idx,priority:3;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"mintedAt"`
	ClosedAt             *DateTime      `gorm:"column:closedAt;type:timestamp(3)" json:"closedAt"`
	CloseReason          *string        `gorm:"column:closeReason" json:"closeReason"`
	TraceID              *string        `gorm:"column:traceId" json:"traceId"`
	VendorCostRaw        datatypes.JSON `gorm:"column:vendorCostRaw" json:"vendorCostRaw"`
	Kind                 *string        `gorm:"column:kind" json:"kind"`
	Metering             *string        `gorm:"column:metering;index:GatewayRealtimeSession_status_metering_mintedAt_idx,priority:2" json:"metering"`
	CredentialExpiresAt  *DateTime      `gorm:"column:credentialExpiresAt;type:timestamp(3)" json:"credentialExpiresAt"`
	TranscriptionModel   *string        `gorm:"column:transcriptionModel" json:"transcriptionModel"`
	EndUserID            *string        `gorm:"column:endUserId" json:"endUserId"`
	LastReportAt         *DateTime      `gorm:"column:lastReportAt;type:timestamp(3)" json:"lastReportAt"`
	ReportedCostNanoUsd  int64          `gorm:"column:reportedCostNanoUsd;default:0;not null" json:"reportedCostNanoUsd"`
	ReportCount          int            `gorm:"column:reportCount;size:32;default:0;not null" json:"reportCount"`
	CreatedAt            DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (GatewayRealtimeSession) TableName() string {
	return "GatewayRealtimeSession"
}

func (m *GatewayRealtimeSession) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.MintedAt.IsZero() {
		m.MintedAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GatewayRealtimeSession) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GatewayRealtimeSessionReport struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	SessionID   string         `gorm:"column:sessionId;uniqueIndex:GatewayRealtimeSessionReport_sessionId_reportKey_key,priority:1;not null" json:"sessionId"`
	ReportKey   string         `gorm:"column:reportKey;uniqueIndex:GatewayRealtimeSessionReport_sessionId_reportKey_key,priority:2;not null" json:"reportKey"`
	ProjectID   string         `gorm:"column:projectId;index:GatewayRealtimeSessionReport_projectId_idx;not null" json:"projectId"`
	Model       string         `gorm:"column:model;not null" json:"model"`
	Usage       datatypes.JSON `gorm:"column:usage;not null" json:"usage"`
	CostNanoUsd int64          `gorm:"column:costNanoUsd;not null" json:"costNanoUsd"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (GatewayRealtimeSessionReport) TableName() string {
	return "GatewayRealtimeSessionReport"
}

func (m *GatewayRealtimeSessionReport) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SystemMigrationTenantState struct {
	MigrationName string         `gorm:"column:migrationName;primaryKey;index:SystemMigrationTenantState_migrationName_status_idx,priority:1" json:"migrationName"`
	TenantID      string         `gorm:"column:tenantId;primaryKey" json:"tenantId"`
	Status        string         `gorm:"column:status;index:SystemMigrationTenantState_migrationName_status_idx,priority:2;not null" json:"status"`
	Report        datatypes.JSON `gorm:"column:report" json:"report"`
	OccurredAt    DateTime       `gorm:"column:occurredAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"occurredAt"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SystemMigrationTenantState) TableName() string {
	return "SystemMigrationTenantState"
}

func (m *SystemMigrationTenantState) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.OccurredAt.IsZero() {
		m.OccurredAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SystemMigrationTenantState) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SystemMigrationEnrollment struct {
	OrganizationID   string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	MigrationName    string   `gorm:"column:migrationName;primaryKey;index:SystemMigrationEnrollment_migrationName_idx" json:"migrationName"`
	EnrolledByUserID string   `gorm:"column:enrolledByUserId;not null" json:"enrolledByUserId"`
	CreatedAt        DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (SystemMigrationEnrollment) TableName() string {
	return "SystemMigrationEnrollment"
}

func (m *SystemMigrationEnrollment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type StoredObject struct {
	TenantID                  string         `gorm:"column:tenantId;primaryKey;index:StoredObject_tenantId_status_id_idx,priority:1" json:"tenantId"`
	ID                        string         `gorm:"column:id;primaryKey;index:StoredObject_tenantId_status_id_idx,priority:3" json:"id"`
	Status                    string         `gorm:"column:status;index:StoredObject_tenantId_status_id_idx,priority:2;index:StoredObject_ownerKind_ownerId_status_idx,priority:3;index:StoredObject_status_expiresAt_idx,priority:1;not null" json:"status"`
	Purpose                   string         `gorm:"column:purpose;not null" json:"purpose"`
	OwnerKind                 string         `gorm:"column:ownerKind;index:StoredObject_ownerKind_ownerId_status_idx,priority:1;not null" json:"ownerKind"`
	OwnerID                   string         `gorm:"column:ownerId;index:StoredObject_ownerKind_ownerId_status_idx,priority:2;not null" json:"ownerId"`
	Filename                  string         `gorm:"column:filename;not null" json:"filename"`
	Sha256                    string         `gorm:"column:sha256;not null" json:"sha256"`
	SizeBytes                 int64          `gorm:"column:sizeBytes;not null" json:"sizeBytes"`
	MediaType                 string         `gorm:"column:mediaType;not null" json:"mediaType"`
	MediaTypeVerified         bool           `gorm:"column:mediaTypeVerified;default:false;not null" json:"mediaTypeVerified"`
	StorageProvider           *string        `gorm:"column:storageProvider" json:"storageProvider"`
	StorageDestinationID      *string        `gorm:"column:storageDestinationId" json:"storageDestinationId"`
	StorageProviderRelativeID *string        `gorm:"column:storageProviderRelativeId;type:text" json:"storageProviderRelativeId"`
	Generation                int            `gorm:"column:generation;size:32;default:0;not null" json:"generation"`
	Audiences                 datatypes.JSON `gorm:"column:audiences;not null" json:"audiences"`
	ExpiresAt                 *DateTime      `gorm:"column:expiresAt;index:StoredObject_status_expiresAt_idx,priority:2;type:timestamp(3)" json:"expiresAt"`
	AvailableAt               *DateTime      `gorm:"column:availableAt;type:timestamp(3)" json:"availableAt"`
	DeletedAt                 *DateTime      `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	Source                    string         `gorm:"column:source;not null" json:"source"`
	LegacyFingerprint         *string        `gorm:"column:legacyFingerprint" json:"legacyFingerprint"`
	CreatedAt                 DateTime       `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (StoredObject) TableName() string {
	return "StoredObject"
}

func (m *StoredObject) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*StoredObject) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A name-text seen on a provider's cost or audit rows: "the provider told us
// someone called m.silva exists". One row per provider-scoped identity.
//
// This is the unit of GDPR erasure (ADR-128 §9), and deliberately so: most of
// these people have no LangWatch login at all - contractors, seat holders who
// never signed in, anyone whose email a provider put on a cost row - so there
// is no platform user to key on and a user-deletion event can never fire for
// them.
type DiscoveredPerson struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:DiscoveredPerson_organizationId_provider_rawActorId_key,priority:1;index:DiscoveredPerson_organizationId_lastSeenAt_idx,priority:1;not null" json:"organizationId"`
	// "anthropic" | "openai" | "databricks" | ... - the provider that named them.
	Provider string `gorm:"column:provider;uniqueIndex:DiscoveredPerson_organizationId_provider_rawActorId_key,priority:2;not null" json:"provider"`
	// What the provider calls them, verbatim. The join key onto the money rows'
	// RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
	// the row addressable for spend attribution without keeping the identifier.
	RawActorID string `gorm:"column:rawActorId;uniqueIndex:DiscoveredPerson_organizationId_provider_rawActorId_key,priority:3;not null" json:"rawActorId"`
	// Name or email text as seen. Pseudonymized alongside rawActorId on erasure.
	DisplayText string `gorm:"column:displayText;not null" json:"displayText"`
	// person | service_account - determined deterministically per ADR-128 §10
	// (a bare UUID in Databricks `executed_by` is a service account, never an
	// agent), never guessed from the shape of a name.
	Kind string `gorm:"column:kind;not null" json:"kind"`
	// The department the provider's directory filed this person under, verbatim.
	// An OBSERVATION, not the `Department` entity: it is free text a tenant owns
	// on their side, it exists for people who hold no LangWatch account, and
	// nothing rolls spend up by it. Spend attribution stays on `Department`, via
	// the membership the match engine proves (directoryDepartmentSync.service).
	// Null means no directory has named one - never "no department", because a
	// blank field in the tenant's directory and a person no directory lists are
	// the same null here.
	Department  *string  `gorm:"column:department" json:"department"`
	FirstSeenAt DateTime `gorm:"column:firstSeenAt;type:timestamp(3);not null" json:"firstSeenAt"`
	LastSeenAt  DateTime `gorm:"column:lastSeenAt;index:DiscoveredPerson_organizationId_lastSeenAt_idx,priority:2;type:timestamp(3);not null" json:"lastSeenAt"`
	// Set when this row's identity has been erased. The row survives so its
	// spend stays attributable to *someone*; the identifier does not.
	ErasedAt *DateTime `gorm:"column:erasedAt;type:timestamp(3)" json:"erasedAt"`
	// Set alongside `erasedAt` and cleared once the daily cost rows have been
	// removed and their rebuild requested. Non-null means the ClickHouse half
	// of the erasure did not finish — a crash between the delete and the
	// replay leaves days short by the erased amount, and without this marker a
	// re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
	// did no work. The re-run resumes at that step instead.
	MoneyRowsPendingAt *DateTime `gorm:"column:moneyRowsPendingAt;type:timestamp(3)" json:"moneyRowsPendingAt"`
	// The earliest day to rebuild from, recorded BEFORE the rows are deleted.
	// Once they are gone there is nothing left to ask which days they were on,
	// so a resume with no record of this could only replay everything or
	// nothing. Null means nothing was rebuildable — either no day carried the
	// identifier, or every one of them is older than the event log keeps.
	MoneyRebuildSince *string `gorm:"column:moneyRebuildSince" json:"moneyRebuildSince"`
	// ADR-128 §12 conflict rule: automatic linking has stopped for this
	// identity because evidence pointed at two candidates, or contradicted an
	// existing link. Stored HERE rather than on the suggestion rows, because a
	// halt on auto-linking that a recompute or a restart clears is not a halt.
	SuspendedAt *DateTime `gorm:"column:suspendedAt;type:timestamp(3)" json:"suspendedAt"`
	// What tripped the suspension, for the human who reviews it.
	SuspendedReason *string  `gorm:"column:suspendedReason" json:"suspendedReason"`
	CreatedAt       DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (DiscoveredPerson) TableName() string {
	return "DiscoveredPerson"
}

func (m *DiscoveredPerson) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DiscoveredPerson) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The same shape as DiscoveredPerson for non-human actors - Genie spaces,
// Copilot bots, OpenAI projects - plus whatever provider-native metadata the
// source published. Discovering an agent grants it nothing either.
type DiscoveredAgent struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:DiscoveredAgent_organizationId_provider_rawAgentId_key,priority:1;index:DiscoveredAgent_organizationId_lastSeenAt_idx,priority:1;not null" json:"organizationId"`
	Provider       string `gorm:"column:provider;uniqueIndex:DiscoveredAgent_organizationId_provider_rawAgentId_key,priority:2;not null" json:"provider"`
	// Genie space / Copilot bot / OpenAI project id, verbatim.
	RawAgentID  string `gorm:"column:rawAgentId;uniqueIndex:DiscoveredAgent_organizationId_provider_rawAgentId_key,priority:3;not null" json:"rawAgentId"`
	DisplayText string `gorm:"column:displayText;not null" json:"displayText"`
	// Workspace, plan, and kin - provider-native descriptive fields only. Never
	// credentials: this row is read by admin screens, not by any pull path.
	Metadata    datatypes.JSON `gorm:"column:metadata;not null" json:"metadata"`
	FirstSeenAt DateTime       `gorm:"column:firstSeenAt;type:timestamp(3);not null" json:"firstSeenAt"`
	LastSeenAt  DateTime       `gorm:"column:lastSeenAt;index:DiscoveredAgent_organizationId_lastSeenAt_idx,priority:2;type:timestamp(3);not null" json:"lastSeenAt"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (DiscoveredAgent) TableName() string {
	return "DiscoveredAgent"
}

func (m *DiscoveredAgent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DiscoveredAgent) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One agent in the organisation's governance registry, from any source governance reads
// (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
// machine-readable snapshot; `version` grows by one on every material change.
type AgentRegistryEntry struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:AgentRegistryEntry_organizationId_sourceKey_key,priority:1;index:AgentRegistryEntry_organizationId_updatedAt_idx,priority:1;not null" json:"organizationId"`
	// Where the agent came from and its id there, e.g. `langwatch:<agentId>`.
	SourceKey   string         `gorm:"column:sourceKey;uniqueIndex:AgentRegistryEntry_organizationId_sourceKey_key,priority:2;not null" json:"sourceKey"`
	Source      string         `gorm:"column:source;not null" json:"source"`
	Status      string         `gorm:"column:status;not null" json:"status"`
	Version     int            `gorm:"column:version;size:32;not null" json:"version"`
	Fingerprint string         `gorm:"column:fingerprint;not null" json:"fingerprint"`
	Entry       datatypes.JSON `gorm:"column:entry;not null" json:"entry"`
	FirstSeenAt DateTime       `gorm:"column:firstSeenAt;type:timestamp(3);not null" json:"firstSeenAt"`
	LastSeenAt  *DateTime      `gorm:"column:lastSeenAt;type:timestamp(3)" json:"lastSeenAt"`
	RemovedAt   *DateTime      `gorm:"column:removedAt;type:timestamp(3)" json:"removedAt"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;index:AgentRegistryEntry_organizationId_updatedAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AgentRegistryEntry) TableName() string {
	return "AgentRegistryEntry"
}

func (m *AgentRegistryEntry) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AgentRegistryEntry) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The registry's append-only change feed: one row per registered, versioned, removed or
// tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
// fan-out as recorded in governance's outbox.
type AgentRegistryEvent struct {
	ID                 string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID     string         `gorm:"column:organizationId;uniqueIndex:AgentRegistryEvent_organizationId_sequence_key,priority:1;index:AgentRegistryEvent_organizationId_deliveryRecordedAt_idx,priority:1;index:AgentRegistryEvent_organizationId_entryId_idx,priority:1;not null" json:"organizationId"`
	Sequence           int64          `gorm:"column:sequence;uniqueIndex:AgentRegistryEvent_organizationId_sequence_key,priority:2;not null" json:"sequence"`
	Type               string         `gorm:"column:type;not null" json:"type"`
	EntryID            string         `gorm:"column:entryId;index:AgentRegistryEvent_organizationId_entryId_idx,priority:2;not null" json:"entryId"`
	FromVersion        *int           `gorm:"column:fromVersion;size:32" json:"fromVersion"`
	ToVersion          *int           `gorm:"column:toVersion;size:32" json:"toVersion"`
	Payload            datatypes.JSON `gorm:"column:payload;not null" json:"payload"`
	OccurredAt         DateTime       `gorm:"column:occurredAt;type:timestamp(3);not null" json:"occurredAt"`
	DeliveryRecordedAt *DateTime      `gorm:"column:deliveryRecordedAt;index:AgentRegistryEvent_organizationId_deliveryRecordedAt_idx,priority:2;type:timestamp(3)" json:"deliveryRecordedAt"`
}

func (AgentRegistryEvent) TableName() string {
	return "AgentRegistryEvent"
}

func (m *AgentRegistryEvent) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// An organisation's registry baseline: its first reconcile, written as already pushed so the
// agents it already had are not announced. Every reconcile after it pushes its changes.
type AgentRegistryBaseline struct {
	OrganizationID string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	BaselinedAt    DateTime `gorm:"column:baselinedAt;type:timestamp(3);not null" json:"baselinedAt"`
}

func (AgentRegistryBaseline) TableName() string {
	return "AgentRegistryBaseline"
}

// One target of a registry agent's kill switch stop plan (enterprise/modules/governance/specs/
// agent-kill-switch.feature): a Databricks serving endpoint, a Bedrock agent alias or a gateway
// virtual key. A plan write replaces every row of the agent at once.
type AgentStopPlanTarget struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:AgentStopPlanTarget_organizationId_agentId_idx,priority:1;not null" json:"organizationId"`
	// The AgentRegistryEntry the plan belongs to.
	AgentID   string         `gorm:"column:agentId;index:AgentStopPlanTarget_organizationId_agentId_idx,priority:2;not null" json:"agentId"`
	Position  int            `gorm:"column:position;size:32;not null" json:"position"`
	Label     string         `gorm:"column:label;not null" json:"label"`
	Target    datatypes.JSON `gorm:"column:target;not null" json:"target"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AgentStopPlanTarget) TableName() string {
	return "AgentStopPlanTarget"
}

func (m *AgentStopPlanTarget) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AgentStopPlanTarget) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One agent's stop plan version: bumped by every save, so a save over a plan someone else saved
// since it was read is refused. No row means the plan was never saved (version 0).
type AgentStopPlan struct {
	OrganizationID string `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	// The AgentRegistryEntry the plan belongs to.
	AgentID   string   `gorm:"column:agentId;primaryKey" json:"agentId"`
	Version   int      `gorm:"column:version;size:32;not null" json:"version"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AgentStopPlan) TableName() string {
	return "AgentStopPlan"
}

func (m *AgentStopPlan) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AgentStopPlan) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// An AWS credential a Bedrock kill switch target signs with: a static key pair, a role to assume,
// or both. The secret access key and the external id are stored sealed.
type GovernanceAwsConnection struct {
	ID                       string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID           string   `gorm:"column:organizationId;index:GovernanceAwsConnection_organizationId_idx;not null" json:"organizationId"`
	Name                     string   `gorm:"column:name;not null" json:"name"`
	AccessKeyID              *string  `gorm:"column:accessKeyId" json:"accessKeyId"`
	EncryptedSecretAccessKey *string  `gorm:"column:encryptedSecretAccessKey" json:"encryptedSecretAccessKey"`
	RoleArn                  *string  `gorm:"column:roleArn" json:"roleArn"`
	EncryptedExternalID      *string  `gorm:"column:encryptedExternalId" json:"encryptedExternalId"`
	CreatedAt                DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (GovernanceAwsConnection) TableName() string {
	return "GovernanceAwsConnection"
}

func (m *GovernanceAwsConnection) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GovernanceAwsConnection) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
// every target's outcome, folded from the stop's events.
type AgentStop struct {
	ID                  string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID      string         `gorm:"column:organizationId;index:AgentStop_organizationId_agentId_idx,priority:1;index:AgentStop_organizationId_status_idx,priority:1;not null" json:"organizationId"`
	AgentID             string         `gorm:"column:agentId;index:AgentStop_organizationId_agentId_idx,priority:2;not null" json:"agentId"`
	AgentName           string         `gorm:"column:agentName;not null" json:"agentName"`
	Status              string         `gorm:"column:status;index:AgentStop_organizationId_status_idx,priority:2;not null" json:"status"`
	Severity            string         `gorm:"column:severity;not null" json:"severity"`
	Approval            string         `gorm:"column:approval;not null" json:"approval"`
	Trigger             string         `gorm:"column:trigger;not null" json:"trigger"`
	Reason              string         `gorm:"column:reason;not null" json:"reason"`
	RequestedByUserID   *string        `gorm:"column:requestedByUserId" json:"requestedByUserId"`
	RequestedByAPIKeyID *string        `gorm:"column:requestedByApiKeyId" json:"requestedByApiKeyId"`
	DecidedByUserID     *string        `gorm:"column:decidedByUserId" json:"decidedByUserId"`
	DecidedByAPIKeyID   *string        `gorm:"column:decidedByApiKeyId" json:"decidedByApiKeyId"`
	RestoredByUserID    *string        `gorm:"column:restoredByUserId" json:"restoredByUserId"`
	RestoredByAPIKeyID  *string        `gorm:"column:restoredByApiKeyId" json:"restoredByApiKeyId"`
	Targets             datatypes.JSON `gorm:"column:targets;not null" json:"targets"`
	RequestedAt         DateTime       `gorm:"column:requestedAt;type:timestamp(3);not null" json:"requestedAt"`
	DecidedAt           *DateTime      `gorm:"column:decidedAt;type:timestamp(3)" json:"decidedAt"`
	QuarantinedAt       *DateTime      `gorm:"column:quarantinedAt;type:timestamp(3)" json:"quarantinedAt"`
	RestoreRequestedAt  *DateTime      `gorm:"column:restoreRequestedAt;type:timestamp(3)" json:"restoreRequestedAt"`
	RestoredAt          *DateTime      `gorm:"column:restoredAt;type:timestamp(3)" json:"restoredAt"`
	LastAppliedEventID  string         `gorm:"column:lastAppliedEventId;not null" json:"lastAppliedEventId"`
	LastAcceptedAt      int64          `gorm:"column:lastAcceptedAt;not null" json:"lastAcceptedAt"`
	LastOccurredAt      int64          `gorm:"column:lastOccurredAt;not null" json:"lastOccurredAt"`
	ProjectionVersion   string         `gorm:"column:projectionVersion;not null" json:"projectionVersion"`
	CreatedAt           DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AgentStop) TableName() string {
	return "AgentStop"
}

func (m *AgentStop) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AgentStop) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The one open stop an agent may have: the api takes this row before it sends a stop's request,
// so two requests at once cannot both open one. A later request takes it over once the stop it
// names has closed.
// `sentAt` is set once the request command was sent: from then on the claim passes on only after
// its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
// `decisionTakenAt` is set while that decision is in flight and cleared once its command was
// sent; a decision left in flight is taken over once abandoned.
type AgentStopClaim struct {
	OrganizationID  string    `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	AgentID         string    `gorm:"column:agentId;primaryKey" json:"agentId"`
	StopID          string    `gorm:"column:stopId;not null" json:"stopId"`
	ClaimedAt       DateTime  `gorm:"column:claimedAt;type:timestamp(3);not null" json:"claimedAt"`
	SentAt          *DateTime `gorm:"column:sentAt;type:timestamp(3)" json:"sentAt"`
	Decision        *string   `gorm:"column:decision" json:"decision"`
	DecisionTakenAt *DateTime `gorm:"column:decisionTakenAt;type:timestamp(3)" json:"decisionTakenAt"`
}

func (AgentStopClaim) TableName() string {
	return "AgentStopClaim"
}

// Whether a target was already stopped before a stop first touched it, recorded once before the
// stop's platform call, so a retried call after a success still knows the change is the stop's own.
type AgentStopTargetPriorState struct {
	OrganizationID string   `gorm:"column:organizationId;primaryKey" json:"organizationId"`
	StopID         string   `gorm:"column:stopId;primaryKey" json:"stopId"`
	TargetID       string   `gorm:"column:targetId;primaryKey" json:"targetId"`
	WasStopped     bool     `gorm:"column:wasStopped;not null" json:"wasStopped"`
	RecordedAt     DateTime `gorm:"column:recordedAt;type:timestamp(3);not null" json:"recordedAt"`
}

func (AgentStopTargetPriorState) TableName() string {
	return "AgentStopTargetPriorState"
}

// The dated match table: "`user_abc123` on Anthropic is platform user #42,
// valid from March 3", closed when it stops being true.
//
// Dating is what makes a re-issued email address survivable: the leaver's link
// closes at offboarding, the new hire gets a new link, and January's spend
// stays with January's person. A correction closes the old row and opens a new
// one; nothing is ever silently rewritten.
type IdentityMatch struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;index:IdentityMatch_organizationId_discoveredPersonId_idx,priority:1;index:IdentityMatch_organizationId_userId_idx,priority:1;not null" json:"organizationId"`
	// The provider-side identity this link is about.
	DiscoveredPersonID string `gorm:"column:discoveredPersonId;index:IdentityMatch_organizationId_discoveredPersonId_idx,priority:2;not null" json:"discoveredPersonId"`
	// The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
	// the row and its dates remain, so the history of the link is not rewritten
	// by the erasure either.
	UserID *string `gorm:"column:userId;index:IdentityMatch_organizationId_userId_idx,priority:2" json:"userId"`
	// directory_id | verified_email | human_confirmed. Exact verified-email
	// equality is the primary evidence; a directory id corroborates but never
	// stands alone (ADR-128 §12).
	EvidenceKind string   `gorm:"column:evidenceKind;not null" json:"evidenceKind"`
	ValidFrom    DateTime `gorm:"column:validFrom;type:timestamp(3);not null" json:"validFrom"`
	// Null means the link is open. Offboarding and corrections close it.
	ValidTo   *DateTime `gorm:"column:validTo;type:timestamp(3)" json:"validTo"`
	CreatedAt DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (IdentityMatch) TableName() string {
	return "IdentityMatch"
}

func (m *IdentityMatch) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Every ClickHouse `TenantId` this organization has ever written governance
// rows under, appended on first use and never pruned.
//
// The live `resolveGovProjectId` lookup cannot be this translation: it filters
// `archivedAt: null` and so returns null forever once someone archives the
// governance project, while the write path re-reads the same project by slug
// with no such filter and keeps landing rows under the old tenant. The
// organization would map to zero tenants on read and one on write - a
// permanent split-brain in which the cost screen reports "no governance data"
// and an erasure job erases nothing and reports success.
//
// Reads resolve against the whole history; erasure walks all of it, because
// personal data does not stop existing in a tenant that stopped being current.
type GovernanceTenantHistory struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string `gorm:"column:organizationId;uniqueIndex:GovernanceTenantHistory_organizationId_tenantId_key,priority:1;index:GovernanceTenantHistory_organizationId_idx;not null" json:"organizationId"`
	// A hidden governance project id this organization has written rows under.
	TenantID    string   `gorm:"column:tenantId;uniqueIndex:GovernanceTenantHistory_organizationId_tenantId_key,priority:2;not null" json:"tenantId"`
	FirstUsedAt DateTime `gorm:"column:firstUsedAt;type:timestamp(3);not null" json:"firstUsedAt"`
	LastUsedAt  DateTime `gorm:"column:lastUsedAt;type:timestamp(3);not null" json:"lastUsedAt"`
}

func (GovernanceTenantHistory) TableName() string {
	return "GovernanceTenantHistory"
}

func (m *GovernanceTenantHistory) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Hashes of erased identifiers - the do-not-reimport list.
//
// Without it the erasure is undone by the pipeline that produced it: the
// pullers look 30 days back, so the next day's pull re-ingests the same email
// and re-writes what we erased. Every write path carrying an identifier
// consults this list and drops a suppressed one (ADR-128 §9 step 1).
//
// It stores hashes and never identifiers, so it is not itself a copy of the
// data it exists to keep out. The stored `identifierHash` and the pseudonym
// written in place of the erased value are the same function of the same
// input - HMAC-SHA256(secret, identifier) - so a write path computes that digest
// once and uses it twice: as the membership test, and, on a hit, as the
// replacement value. That is why no original-to-pseudonym mapping table
// exists; one would hold the erased identifier in plaintext forever.
//
// Scoped to the provider as well as the organization: the same string can be a
// different person under a different provider.
type ErasedIdentifierSuppression struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string   `gorm:"column:organizationId;uniqueIndex:ErasedIdentifierSuppression_organizationId_provider_identif_key,priority:1;index:ErasedIdentifierSuppression_organizationId_idx;not null" json:"organizationId"`
	Provider       string   `gorm:"column:provider;uniqueIndex:ErasedIdentifierSuppression_organizationId_provider_identif_key,priority:2;not null" json:"provider"`
	IdentifierHash string   `gorm:"column:identifierHash;uniqueIndex:ErasedIdentifierSuppression_organizationId_provider_identif_key,priority:3;not null" json:"identifierHash"`
	ErasedAt       DateTime `gorm:"column:erasedAt;type:timestamp(3);not null" json:"erasedAt"`
}

func (ErasedIdentifierSuppression) TableName() string {
	return "ErasedIdentifierSuppression"
}

func (m *ErasedIdentifierSuppression) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// Match candidates a background job computed and a human has yet to confirm
// (ADR-128 §12).
//
// Stored rather than computed when the review screen asks, and the reason is
// measured rather than stylistic: fuzzy matching is quadratic, there is no
// database route to it here, and at this document's own example size - 2,000
// discovered people against 500 accounts - plain edit distance over the
// 1,000,000 pairs blocked the event loop for 2.9 seconds. Per page load,
// uncached, stalling every other request on the instance. A stored row also
// makes a pending count answerable without paying for the whole sweep.
//
// Rows are written only by the job, which prefilters before it scores and
// recomputes when its inputs change. Confirming one opens an `IdentityMatch`
// and deletes the suggestion.
//
// Suspension deliberately does NOT live here - it is on `DiscoveredPerson`,
// because a halt on automatic linking that the next recompute clears is not a
// halt.
type IdentityMatchSuggestion struct {
	ID                 string `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID     string `gorm:"column:organizationId;uniqueIndex:IdentityMatchSuggestion_organizationId_discoveredPersonId_u_key,priority:1;index:IdentityMatchSuggestion_organizationId_score_idx,priority:1;not null" json:"organizationId"`
	DiscoveredPersonID string `gorm:"column:discoveredPersonId;uniqueIndex:IdentityMatchSuggestion_organizationId_discoveredPersonId_u_key,priority:2;not null" json:"discoveredPersonId"`
	// The platform user this might be. Not nullable: a suggestion with no
	// candidate is not a suggestion.
	UserID string `gorm:"column:userId;uniqueIndex:IdentityMatchSuggestion_organizationId_discoveredPersonId_u_key,priority:3;not null" json:"userId"`
	// Similarity of the two name texts, after the prefilter, in [0, 1]. A
	// display hint and an ordering key - never a threshold anything auto-links
	// on, which is the whole distinction this table exists to hold.
	Score      float64  `gorm:"column:score;index:IdentityMatchSuggestion_organizationId_score_idx,priority:2;type:double precision;not null" json:"score"`
	ComputedAt DateTime `gorm:"column:computedAt;type:timestamp(3);not null" json:"computedAt"`
}

func (IdentityMatchSuggestion) TableName() string {
	return "IdentityMatchSuggestion"
}

func (m *IdentityMatchSuggestion) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// What a mid-term seat change of one connected license owes (ADR-141,
// section 7): the seats added on the reissued license, invoiced prorated to
// the end of the term.
//
// The row is written before the payment provider is called and completed
// after it, so a call that failed halfway is retried by the next seat
// invoicing pass without inventing a second invoice. One row per reissued license.
type ConnectedSeatChange struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The reissued `IssuedLicense` row that carries the new seat count. A
	// plain column ref, like `replacesId`: relationMode="prisma" carries no
	// foreign keys.
	LicenseID string `gorm:"column:licenseId;uniqueIndex:ConnectedSeatChange_licenseId_key;not null" json:"licenseId"`
	// The `ConnectedBillingAccount` row the invoice is raised against. Null
	// when the customer had no billing account (`not_onboarded`).
	AccountID *string `gorm:"column:accountId;index:ConnectedSeatChange_accountId_idx" json:"accountId"`
	// The customer organization the license belongs to.
	OrganizationID *string `gorm:"column:organizationId;index:ConnectedSeatChange_organizationId_idx" json:"organizationId"`
	// When the seats changed, which is the day the proration counts from.
	ChangedAt DateTime `gorm:"column:changedAt;type:timestamp(3);not null" json:"changedAt"`
	// Seats on the new license above the seats on the one it replaced.
	AddedSeats int `gorm:"column:addedSeats;size:32;not null" json:"addedSeats"`
	// What one added seat costs for the rest of the term, in cents of `currency`.
	UnitAmountCents int `gorm:"column:unitAmountCents;size:32;not null" json:"unitAmountCents"`
	// `unitAmountCents` times `addedSeats`. Zero when nothing is owed.
	AmountCents int `gorm:"column:amountCents;size:32;not null" json:"amountCents"`
	// Null when no billing account named one (`not_onboarded`).
	Currency *string `gorm:"column:currency" json:"currency"`
	// The one-off invoice the added seats were billed on.
	StripeInvoiceID *string  `gorm:"column:stripeInvoiceId" json:"stripeInvoiceId"`
	State           string   `gorm:"column:state;index:ConnectedSeatChange_state_idx;not null" json:"state"`
	CreatedAt       DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ConnectedSeatChange) TableName() string {
	return "ConnectedSeatChange"
}

func (m *ConnectedSeatChange) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ConnectedSeatChange) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The monthly usage statement one connected customer was sent (ADR-139,
// section 7). One row per account per month keeps the statement to one send,
// whatever runs the tick.
type ConnectedStatement struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The `ConnectedBillingAccount` row. A plain column ref.
	AccountID string `gorm:"column:accountId;uniqueIndex:ConnectedStatement_accountId_month_key,priority:1;index:ConnectedStatement_accountId_idx;not null" json:"accountId"`
	// First day of the month the statement covers, in UTC.
	Month  DateTime `gorm:"column:month;uniqueIndex:ConnectedStatement_accountId_month_key,priority:2;type:timestamp(3);not null" json:"month"`
	SentAt DateTime `gorm:"column:sentAt;type:timestamp(3);not null" json:"sentAt"`
}

func (ConnectedStatement) TableName() string {
	return "ConnectedStatement"
}

func (m *ConnectedStatement) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	return nil
}

// One renewed licence per paid Stripe invoice of a self-serve licence
// subscription (specs/licensing/self-serving-license-purchase.feature). A
// redelivered invoice event finds the row and mints nothing. A claim never
// delivered is taken over once `claimedAt` is older than the claim lease.
type LicenseRenewalClaim struct {
	StripeInvoiceID      string `gorm:"column:stripeInvoiceId;primaryKey" json:"stripeInvoiceId"`
	StripeSubscriptionID string `gorm:"column:stripeSubscriptionId;not null" json:"stripeSubscriptionId"`
	// Fences completion and release to the delivery that holds the claim.
	ClaimToken string   `gorm:"column:claimToken;not null" json:"claimToken"`
	ClaimedAt  DateTime `gorm:"column:claimedAt;type:timestamp(3);not null" json:"claimedAt"`
	// The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered.
	IssuedLicense datatypes.JSON `gorm:"column:issuedLicense" json:"issuedLicense"`
	DeliveredAt   *DateTime      `gorm:"column:deliveredAt;type:timestamp(3)" json:"deliveredAt"`
	// The `licenseId` inside the delivered licence.
	LicenseID *string `gorm:"column:licenseId" json:"licenseId"`
}

func (LicenseRenewalClaim) TableName() string {
	return "LicenseRenewalClaim"
}

// The identity this install presents to LangWatch (ADR-139, section 2).
//
// Exactly one row, minted on first use and never rewritten. It survives
// restarts, backups and hostname changes, and an install restored into
// another database mints another one, which is what the registry refuses as
// the wrong instance.
//
// It replaced an identity built from the organization name and id, which
// leaked the customer's name in cleartext, named one organization rather than
// one install, and could not be joined to a license row without splitting a
// string.
type InstanceIdentity struct {
	// Fixed, so this table holds one row. A second insert loses the race on the
	// primary key and re-reads the winner's id.
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// A random UUID. Carries nothing about the customer.
	InstanceID string   `gorm:"column:instanceId;uniqueIndex:InstanceIdentity_instanceId_key;not null" json:"instanceId"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// When the usage report last reached LangWatch, and why the last attempt
	// did not. Both are what the checkup page reads to say whether reporting
	// works, which nothing recorded before: the sender logged success as soon
	// as `fetch` resolved and never looked at the response.
	LastReportAt    *DateTime `gorm:"column:lastReportAt;type:timestamp(3)" json:"lastReportAt"`
	LastReportError *string   `gorm:"column:lastReportError" json:"lastReportError"`
	// The optional category of the usage report, switched off from the checkup
	// page. With it off the report still says which release runs and how many
	// projects the install carries, and says nothing about what anyone does
	// with it (ADR-139, section 10).
	OptionalMetricsOptOut bool `gorm:"column:optionalMetricsOptOut;default:false;not null" json:"optionalMetricsOptOut"`
	// Hostname, which names the customer's own network and so has a switch of
	// its own rather than travelling with its category.
	HostnameOptOut bool `gorm:"column:hostnameOptOut;default:false;not null" json:"hostnameOptOut"`
	// Sending a failed upgrade step's diagnostics to LangWatch on its own,
	// switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
	UpgradeDiagnosticsOptOut bool `gorm:"column:upgradeDiagnosticsOptOut;default:false;not null" json:"upgradeDiagnosticsOptOut"`
}

func (InstanceIdentity) TableName() string {
	return "InstanceIdentity"
}

func (m *InstanceIdentity) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// One self-hosted install, as its daily usage report describes it (ADR-139,
// section 10). The row is created by the first report that names an instance
// id and is updated by every report after that, so the backoffice reads one
// row per install rather than folding a history on every page load.
//
// The report carries no credential, so it identifies an install and never a
// customer. `organizationId` and `issuedLicenseId` are resolved from the
// license bound to this instance, which the install proved by presenting its
// token to license sync. A field in the posted body never reaches them.
type SelfHostedInstance struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The UUID the install minted for itself. The join key for everything else.
	InstanceID  string   `gorm:"column:instanceId;uniqueIndex:SelfHostedInstance_instanceId_key;not null" json:"instanceId"`
	FirstSeenAt DateTime `gorm:"column:firstSeenAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"firstSeenAt"`
	LastSeenAt  DateTime `gorm:"column:lastSeenAt;index:SelfHostedInstance_lastSeenAt_idx;type:timestamp(3);not null" json:"lastSeenAt"`
	// What the install says about itself.
	Version       *string `gorm:"column:version" json:"version"`
	InstallMethod *string `gorm:"column:installMethod" json:"installMethod"`
	ChartVersion  *string `gorm:"column:chartVersion" json:"chartVersion"`
	// The host it answers on, when the operator left hostname reporting on.
	Hostname    *string `gorm:"column:hostname" json:"hostname"`
	Environment *string `gorm:"column:environment" json:"environment"`
	// The day the install itself says it first came up, which predates the
	// first report on an install that switched reporting on later.
	InstalledAt         *DateTime `gorm:"column:installedAt;type:timestamp(3)" json:"installedAt"`
	ReportSchemaVersion *int      `gorm:"column:reportSchemaVersion;size:32" json:"reportSchemaVersion"`
	// The customer, and the license that named them. Plain column refs rather
	// than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
	// carries no foreign keys, and the receiver resolves both from the license
	// registry rather than from the report.
	OrganizationID  *string `gorm:"column:organizationId;index:SelfHostedInstance_organizationId_idx" json:"organizationId"`
	IssuedLicenseID *string `gorm:"column:issuedLicenseId;index:SelfHostedInstance_issuedLicenseId_idx" json:"issuedLicenseId"`
	// Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
	// install that switched the optional category off.
	UserEmailDomains datatypes.JSON `gorm:"column:userEmailDomains" json:"userEmailDomains"`
	// The same domains as the keys of `userEmailDomains`, flattened so an
	// operator can search for one. A JSON object's keys are the domains
	// themselves, and no query language addresses a key it does not already
	// know the name of.
	UserDomains ScalarList[string] `gorm:"column:userDomains;type:text[]" json:"userDomains"`
	// What the last report carried, so the list and the drawer read one row:
	// the onboarding ladder, the usage counts and the operational block, under
	// the same keys the dictionary declares.
	LatestReport datatypes.JSON `gorm:"column:latestReport" json:"latestReport"`
	// What the last report said the operator had switched off, read from which
	// categories arrived rather than from a field the install could assert.
	OptionalMetricsReported *bool `gorm:"column:optionalMetricsReported;default:true;not null" json:"optionalMetricsReported"`
	HostnameReported        *bool `gorm:"column:hostnameReported;default:true;not null" json:"hostnameReported"`
	// How many reports have landed, and how many fields the last one carried
	// that this release has no name for. A count above zero says the senders
	// are ahead of the receiver.
	ReportCount       int `gorm:"column:reportCount;size:32;default:0;not null" json:"reportCount"`
	LastUnknownFields int `gorm:"column:lastUnknownFields;size:32;default:0;not null" json:"lastUnknownFields"`
	// The lead signals already raised for this install. A signal on this list
	// is never raised again, which is what keeps a large install from
	// reporting that it is large every morning.
	RaisedSignals ScalarList[string] `gorm:"column:raisedSignals;type:text[]" json:"raisedSignals"`
	CreatedAt     DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SelfHostedInstance) TableName() string {
	return "SelfHostedInstance"
}

func (m *SelfHostedInstance) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
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

func (*SelfHostedInstance) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One upgrade diagnostics bundle a self-hosted install sent over Connect, on
// LangWatch Cloud. The bytes live in object storage under
// `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
// points at them and feeds Cloud admin and the Slack signal. Plain column
// refs, no relation, as the usage report rows.
type SelfHostedInstanceDiagnostic struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The `SelfHostedInstance.instanceId` the licence caller presented.
	InstanceID      string  `gorm:"column:instanceId;index:SelfHostedInstanceDiagnostic_instanceId_receivedAt_idx,priority:1;not null" json:"instanceId"`
	StepID          string  `gorm:"column:stepId;not null" json:"stepId"`
	StepDescription *string `gorm:"column:stepDescription" json:"stepDescription"`
	Release         *string `gorm:"column:release" json:"release"`
	Attempt         int     `gorm:"column:attempt;size:32;not null" json:"attempt"`
	// `operator` or `automatic`.
	Trigger    string   `gorm:"column:trigger;not null" json:"trigger"`
	ByteLength int      `gorm:"column:byteLength;size:32;not null" json:"byteLength"`
	ObjectKey  string   `gorm:"column:objectKey;not null" json:"objectKey"`
	ReceivedAt DateTime `gorm:"column:receivedAt;index:SelfHostedInstanceDiagnostic_instanceId_receivedAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"receivedAt"`
	// When the lifecycle rule removes the bundle; the row is listed until then.
	ExpiresAt DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	// Set on the report that posted this step and release's one Slack message.
	SlackAnnouncedAt *DateTime `gorm:"column:slackAnnouncedAt;type:timestamp(3)" json:"slackAnnouncedAt"`
}

func (SelfHostedInstanceDiagnostic) TableName() string {
	return "SelfHostedInstanceDiagnostic"
}

func (m *SelfHostedInstanceDiagnostic) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ReceivedAt.IsZero() {
		m.ReceivedAt = DateTime{Time: now}
	}
	return nil
}

// A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
// report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
// tenant: every install receives every published notice that has not expired.
type SelfHostedNotice struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// `info`, `warning` or `security`.
	Severity string `gorm:"column:severity;not null" json:"severity"`
	Title    string `gorm:"column:title;not null" json:"title"`
	Body     string `gorm:"column:body;not null" json:"body"`
	// An https link to read more; null where the notice names none.
	URL *string `gorm:"column:url" json:"url"`
	// The release that fixes what the notice is about; an install on it or newer hides it.
	FixedIn *string `gorm:"column:fixedIn" json:"fixedIn"`
	// `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent.
	Status      string    `gorm:"column:status;not null" json:"status"`
	ExpiresAt   *DateTime `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	PublishedAt *DateTime `gorm:"column:publishedAt;type:timestamp(3)" json:"publishedAt"`
	// The staff member who published it, by user id.
	PublishedBy *string  `gorm:"column:publishedBy" json:"publishedBy"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SelfHostedNotice) TableName() string {
	return "SelfHostedNotice"
}

func (m *SelfHostedNotice) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SelfHostedNotice) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One usage report, exactly as it arrived. Kept so a number on an install's
// row can be traced back to the report that set it, and so a question asked
// later can be answered over the history rather than over the latest snapshot.
type SelfHostedInstanceReport struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The `SelfHostedInstance.instanceId` this report came from. A plain column
	// ref for the same reason the instance row's refs are.
	InstanceID          string   `gorm:"column:instanceId;index:SelfHostedInstanceReport_instanceId_receivedAt_idx,priority:1;not null" json:"instanceId"`
	ReceivedAt          DateTime `gorm:"column:receivedAt;index:SelfHostedInstanceReport_instanceId_receivedAt_idx,priority:2;index:SelfHostedInstanceReport_receivedAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"receivedAt"`
	Version             *string  `gorm:"column:version" json:"version"`
	ReportSchemaVersion *int     `gorm:"column:reportSchemaVersion;size:32" json:"reportSchemaVersion"`
	// Fields this release had no name for. Counted, never stored by name.
	UnknownFields int `gorm:"column:unknownFields;size:32;default:0;not null" json:"unknownFields"`
	// The validated report, with every unknown field already dropped.
	Payload datatypes.JSON `gorm:"column:payload;not null" json:"payload"`
}

func (SelfHostedInstanceReport) TableName() string {
	return "SelfHostedInstanceReport"
}

func (m *SelfHostedInstanceReport) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.ReceivedAt.IsZero() {
		m.ReceivedAt = DateTime{Time: now}
	}
	return nil
}

// One browser a person subscribed to Web Push, owned by notification. A
// person holds one row per device; the endpoint names the browser, so a
// browser that subscribes again replaces its own row.
type WebPushSubscription struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The person the pushes are for. A plain column: the row leaves with the
	// person through notification's own subscriber on the user lifecycle.
	UserID string `gorm:"column:userId;index:WebPushSubscription_userId_idx;not null" json:"userId"`
	// The push service URL the browser handed out. Unique per browser.
	Endpoint string `gorm:"column:endpoint;uniqueIndex:WebPushSubscription_endpoint_key;not null" json:"endpoint"`
	// The browser's P-256 public key and auth secret, base64url, as the
	// browser reports them. Payloads are encrypted to these.
	P256dh    string   `gorm:"column:p256dh;not null" json:"p256dh"`
	Auth      string   `gorm:"column:auth;not null" json:"auth"`
	UserAgent *string  `gorm:"column:userAgent" json:"userAgent"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// When the push service last accepted a push for this browser.
	LastSuccessAt *DateTime `gorm:"column:lastSuccessAt;type:timestamp(3)" json:"lastSuccessAt"`
}

func (WebPushSubscription) TableName() string {
	return "WebPushSubscription"
}

func (m *WebPushSubscription) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = gonanoid.Must()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// This installation's VAPID key pair, generated on first use so Web Push works
// without configuration. Fixed id, so the table holds one row: a second
// insert loses the race on the primary key and reads the winner's pair.
type WebPushVapidKey struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The uncompressed P-256 public key, base64url. The browser subscribes with it.
	PublicKey string `gorm:"column:publicKey;not null" json:"publicKey"`
	// The private key, base64url, encrypted with the process encryption key.
	PrivateKeyEncrypted string   `gorm:"column:privateKeyEncrypted;not null" json:"privateKeyEncrypted"`
	CreatedAt           DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (WebPushVapidKey) TableName() string {
	return "WebPushVapidKey"
}

func (m *WebPushVapidKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SignInProviderMount struct {
	// Fixed, so this table holds one row: licensing's copy of sso's sign-in
	// provider fact (ADR-179).
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// The configured provider's name; null when the deployment signs in by email.
	Provider *string `gorm:"column:provider" json:"provider"`
	Mounted  bool    `gorm:"column:mounted;not null" json:"mounted"`
	// The last sso event applied: the row's idempotency guard and its version.
	LastAppliedEventID string   `gorm:"column:lastAppliedEventId;not null" json:"lastAppliedEventId"`
	UpdatedAt          DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (SignInProviderMount) TableName() string {
	return "SignInProviderMount"
}

func (m *SignInProviderMount) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SignInProviderMount) BeforeUpdate(tx *gorm.DB) error {
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

type ScalarList[T string | int | int64 | float64 | bool] []T

func (list ScalarList[T]) Value() (driver.Value, error) {
	items := make([]string, len(list))
	for i, item := range list {
		switch value := any(item).(type) {
		case string:
			items[i] = `"` + strings.NewReplacer(`\`, `\\`, `"`, `\"`).Replace(value) + `"`
		case int:
			items[i] = strconv.FormatInt(int64(value), 10)
		case int64:
			items[i] = strconv.FormatInt(value, 10)
		case float64:
			items[i] = strings.NewReplacer("+Inf", "Infinity", "-Inf", "-Infinity").Replace(strconv.FormatFloat(value, 'g', -1, 64))
		case bool:
			items[i] = strconv.FormatBool(value)
		}
	}
	return "{" + strings.Join(items, ",") + "}", nil
}

func (list *ScalarList[T]) Scan(src any) error {
	var text string
	switch value := src.(type) {
	case nil:
		*list = ScalarList[T]{}
		return nil
	case []byte:
		text = string(value)
	case string:
		text = value
	default:
		return fmt.Errorf("an array column held %T", src)
	}
	items, err := arrayItems(text)
	if err != nil {
		return err
	}
	read := make(ScalarList[T], len(items))
	for i, item := range items {
		var value any
		var parseErr error
		switch any(read[i]).(type) {
		case string:
			value = item
		case int:
			value, parseErr = strconv.Atoi(item)
		case int64:
			value, parseErr = strconv.ParseInt(item, 10, 64)
		case float64:
			value, parseErr = strconv.ParseFloat(item, 64)
		case bool:
			switch item {
			case "t", "true":
				value = true
			case "f", "false":
				value = false
			default:
				parseErr = fmt.Errorf("%q is not a boolean", item)
			}
		}
		if parseErr != nil {
			return fmt.Errorf("an array column held %q: %w", text, parseErr)
		}
		read[i] = value.(T)
	}
	*list = read
	return nil
}

// arrayItems splits the text form of a one-dimensional PostgreSQL array into its elements,
// unquoted. A NULL element (an unquoted NULL) is an error: a Prisma list holds none, and the
// string "NULL" is written quoted.
func arrayItems(text string) ([]string, error) {
	// An array whose lower bound is not 1 is printed with its bounds first: [0:1]={x,a}.
	if strings.HasPrefix(text, "[") {
		if eq := strings.Index(text, "="); eq >= 0 {
			text = text[eq+1:]
		}
	}
	if len(text) < 2 || text[0] != '{' || text[len(text)-1] != '}' {
		return nil, fmt.Errorf("an array column held %q", text)
	}
	body := text[1 : len(text)-1]
	items := []string{}
	if body == "" {
		return items, nil
	}
	var item strings.Builder
	quoted, escaped, wasQuoted := false, false, false
	end := func() error {
		if !wasQuoted && item.String() == "NULL" {
			return fmt.Errorf("an array column held a NULL element: %q", text)
		}
		items = append(items, item.String())
		item.Reset()
		wasQuoted = false
		return nil
	}
	// Every delimiter is ASCII, so the text is walked byte by byte and any other byte kept as it is.
	for i := 0; i < len(body); i++ {
		r := body[i]
		switch {
		case escaped:
			item.WriteByte(r)
			escaped = false
		case r == '\\':
			escaped = true
		case r == '"':
			quoted = !quoted
			wasQuoted = true
		case r == ',' && !quoted:
			if err := end(); err != nil {
				return nil, err
			}
		default:
			item.WriteByte(r)
		}
	}
	if err := end(); err != nil {
		return nil, err
	}
	return items, nil
}
