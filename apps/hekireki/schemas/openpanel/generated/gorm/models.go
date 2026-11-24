package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var NamingStrategy = schema.NamingStrategy{SingularTable: true, NoLowerCase: true}

type CodeMigration struct {
	ID        string   `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string   `gorm:"column:name;uniqueIndex:__code_migrations_name_key;not null" json:"name"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (CodeMigration) TableName() string {
	return "__code_migrations"
}

func (m *CodeMigration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CodeMigration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Conversation struct {
	ID             string        `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	ProjectID      string        `gorm:"column:projectId;index:conversations_projectId_updatedAt_idx;not null" json:"projectId"`
	OrganizationID string        `gorm:"column:organizationId;not null" json:"organizationId"`
	UserID         string        `gorm:"column:userId;index:conversations_userId_updatedAt_idx;not null" json:"userId"`
	Title          *string       `gorm:"column:title" json:"title"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;index:conversations_projectId_updatedAt_idx;index:conversations_userId_updatedAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project        *Project      `gorm:"belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
	User           *User         `gorm:"belongsTo"`
	Messages       []ChatMessage `gorm:"foreignKey:ConversationID;constraint:chat_messages_conversationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Conversation) TableName() string {
	return "conversations"
}

func (m *Conversation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Conversation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ChatMessage struct {
	ID             string `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	ConversationID string `gorm:"column:conversationId;index:chat_messages_conversationId_createdAt_idx;not null" json:"conversationId"`
	Role           string `gorm:"column:role;not null" json:"role"`
	// [IPrismaUIMessageParts]
	Parts datatypes.JSON `gorm:"column:parts;not null" json:"parts"`
	// [IPrismaChatMessageMetadata]
	Metadata     datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt    DateTime       `gorm:"column:createdAt;index:chat_messages_conversationId_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Conversation *Conversation  `gorm:"belongsTo"`
}

func (ChatMessage) TableName() string {
	return "chat_messages"
}

func (m *ChatMessage) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Organization struct {
	ID              string  `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name            string  `gorm:"column:name;not null" json:"name"`
	Slug            *string `gorm:"column:slug;uniqueIndex:organizations_slug_key" json:"slug"`
	CreatedByUserID *string `gorm:"column:createdByUserId" json:"createdByUserId"`
	Timezone        *string `gorm:"column:timezone" json:"timezone"`
	Onboarding      *string `gorm:"column:onboarding;default:'completed'" json:"onboarding"`
	// [IPrismaOrganizationAttribution]
	Attribution            datatypes.JSON `gorm:"column:attribution" json:"attribution"`
	SubscriptionID         *string        `gorm:"column:subscriptionId" json:"subscriptionId"`
	SubscriptionCustomerID *string        `gorm:"column:subscriptionCustomerId" json:"subscriptionCustomerId"`
	SubscriptionPriceID    *string        `gorm:"column:subscriptionPriceId" json:"subscriptionPriceId"`
	SubscriptionProductID  *string        `gorm:"column:subscriptionProductId" json:"subscriptionProductId"`
	// [IPrismaSubscriptionStatus]
	SubscriptionStatus                      *string   `gorm:"column:subscriptionStatus" json:"subscriptionStatus"`
	SubscriptionStartsAt                    *DateTime `gorm:"column:subscriptionStartsAt;type:timestamp(3)" json:"subscriptionStartsAt"`
	SubscriptionEndsAt                      *DateTime `gorm:"column:subscriptionEndsAt;type:timestamp(3)" json:"subscriptionEndsAt"`
	SubscriptionCanceledAt                  *DateTime `gorm:"column:subscriptionCanceledAt;type:timestamp(3)" json:"subscriptionCanceledAt"`
	SubscriptionCreatedByUserID             *string   `gorm:"column:subscriptionCreatedByUserId" json:"subscriptionCreatedByUserId"`
	SubscriptionPeriodEventsCount           int       `gorm:"column:subscriptionPeriodEventsCount;size:32;default:0;not null" json:"subscriptionPeriodEventsCount"`
	SubscriptionPeriodEventsCountExceededAt *DateTime `gorm:"column:subscriptionPeriodEventsCountExceededAt;type:timestamp(3)" json:"subscriptionPeriodEventsCountExceededAt"`
	SubscriptionPeriodEventsLimit           int       `gorm:"column:subscriptionPeriodEventsLimit;size:32;default:0;not null" json:"subscriptionPeriodEventsLimit"`
	SubscriptionInterval                    *string   `gorm:"column:subscriptionInterval" json:"subscriptionInterval"`
	// [IPrismaCancellationReason]
	SubscriptionCancelReason          *string   `gorm:"column:subscriptionCancelReason" json:"subscriptionCancelReason"`
	SubscriptionCancelComment         *string   `gorm:"column:subscriptionCancelComment" json:"subscriptionCancelComment"`
	SubscriptionSaveDiscountAppliedAt *DateTime `gorm:"column:subscriptionSaveDiscountAppliedAt;type:timestamp(3)" json:"subscriptionSaveDiscountAppliedAt"`
	SubscriptionPauseAtPeriodEnd      bool      `gorm:"column:subscriptionPauseAtPeriodEnd;default:false;not null" json:"subscriptionPauseAtPeriodEnd"`
	SubscriptionResumesAt             *DateTime `gorm:"column:subscriptionResumesAt;type:timestamp(3)" json:"subscriptionResumesAt"`
	SubscriptionFirstStartedAt        *DateTime `gorm:"column:subscriptionFirstStartedAt;type:timestamp(3)" json:"subscriptionFirstStartedAt"`
	UsageWarningSentAt                *DateTime `gorm:"column:usageWarningSentAt;type:timestamp(3)" json:"usageWarningSentAt"`
	UsageExceededSentAt               *DateTime `gorm:"column:usageExceededSentAt;type:timestamp(3)" json:"usageExceededSentAt"`
	// [IPrismaSubscriptionDiscount]
	SubscriptionDiscount datatypes.JSON `gorm:"column:subscriptionDiscount" json:"subscriptionDiscount"`
	WindDownStartedAt    *DateTime      `gorm:"column:windDownStartedAt;index:organizations_windDownStartedAt_idx;type:timestamp(3)" json:"windDownStartedAt"`
	// [IPrismaWindDownStep]
	WindDownStep          *string          `gorm:"column:windDownStep" json:"windDownStep"`
	DeleteAt              *DateTime        `gorm:"column:deleteAt;type:timestamp(3)" json:"deleteAt"`
	CreatedAt             DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime         `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedBy             *User            `gorm:"foreignKey:CreatedByUserID;belongsTo"`
	SubscriptionCreatedBy *User            `gorm:"foreignKey:SubscriptionCreatedByUserID;belongsTo"`
	Projects              []Project        `gorm:"foreignKey:OrganizationID;constraint:projects_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Members               []Member         `gorm:"foreignKey:OrganizationID;constraint:members_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Client                []Client         `gorm:"foreignKey:OrganizationID;constraint:clients_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Dashboard             []Dashboard      `gorm:"foreignKey:OrganizationID;constraint:dashboards_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShareOverview         []ShareOverview  `gorm:"foreignKey:OrganizationID;constraint:shares_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShareDashboard        []ShareDashboard `gorm:"foreignKey:OrganizationID;constraint:share_dashboards_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShareReport           []ShareReport    `gorm:"foreignKey:OrganizationID;constraint:share_reports_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShareWidget           []ShareWidget    `gorm:"foreignKey:OrganizationID;constraint:share_widgets_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Integrations          []Integration    `gorm:"foreignKey:OrganizationID;constraint:integrations_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites               []Invite         `gorm:"foreignKey:OrganizationID;constraint:invites_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Conversations         []Conversation   `gorm:"foreignKey:OrganizationID;constraint:conversations_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Organization) TableName() string {
	return "organizations"
}

func (m *Organization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type User struct {
	ID                   string               `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Email                string               `gorm:"column:email;uniqueIndex:users_email_key;not null" json:"email"`
	FirstName            *string              `gorm:"column:firstName" json:"firstName"`
	LastName             *string              `gorm:"column:lastName" json:"lastName"`
	Name                 *string              `gorm:"column:name" json:"name"`
	EmailVerified        bool                 `gorm:"column:emailVerified;default:false;not null" json:"emailVerified"`
	Image                *string              `gorm:"column:image" json:"image"`
	TwoFactorEnabled     bool                 `gorm:"column:twoFactorEnabled;default:false;not null" json:"twoFactorEnabled"`
	CreatedAt            DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime             `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt            *DateTime            `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	CreatedOrganizations []Organization       `gorm:"foreignKey:CreatedByUserID;constraint:organizations_createdByUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Subscriptions        []Organization       `gorm:"foreignKey:SubscriptionCreatedByUserID;constraint:organizations_subscriptionCreatedByUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Membership           []Member             `gorm:"foreignKey:UserID;constraint:members_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SentInvites          []Member             `gorm:"foreignKey:InvitedByID;constraint:members_invitedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Sessions             []Session            `gorm:"foreignKey:UserID;constraint:sessions_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Accounts             []Account            `gorm:"foreignKey:UserID;constraint:accounts_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites              []Invite             `gorm:"foreignKey:CreatedByID;constraint:invites_createdById_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Conversations        []Conversation       `gorm:"foreignKey:UserID;constraint:conversations_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TwoFactorChallenges  []TwoFactorChallenge `gorm:"foreignKey:UserID;constraint:two_factor_challenges_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TwoFactor            *TwoFactor           `gorm:"foreignKey:UserID;constraint:two_factor_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Totp                 *UserTotp            `gorm:"foreignKey:UserID;constraint:user_totp_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "users"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type UserTotp struct {
	ID            string             `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID        string             `gorm:"column:userId;uniqueIndex:user_totp_userId_key;not null" json:"userId"`
	Secret        string             `gorm:"column:secret;not null" json:"secret"`
	RecoveryCodes ScalarList[string] `gorm:"column:recoveryCodes;type:text[]" json:"recoveryCodes"`
	EnabledAt     *DateTime          `gorm:"column:enabledAt;type:timestamp(3)" json:"enabledAt"`
	CreatedAt     DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User          *User              `gorm:"belongsTo"`
}

func (UserTotp) TableName() string {
	return "user_totp"
}

func (m *UserTotp) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*UserTotp) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TwoFactor struct {
	ID                      string    `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID                  string    `gorm:"column:userId;uniqueIndex:two_factor_userId_key;not null" json:"userId"`
	Secret                  string    `gorm:"column:secret;not null" json:"secret"`
	BackupCodes             string    `gorm:"column:backupCodes;not null" json:"backupCodes"`
	Verified                *bool     `gorm:"column:verified;default:true;not null" json:"verified"`
	FailedVerificationCount int       `gorm:"column:failedVerificationCount;size:32;default:0;not null" json:"failedVerificationCount"`
	LockedUntil             *DateTime `gorm:"column:lockedUntil;type:timestamp(3)" json:"lockedUntil"`
	CreatedAt               DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User                    *User     `gorm:"belongsTo"`
}

func (TwoFactor) TableName() string {
	return "two_factor"
}

func (m *TwoFactor) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type TwoFactorChallenge struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	UserID    string   `gorm:"column:userId;index:two_factor_challenges_userId_idx;not null" json:"userId"`
	ExpiresAt DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	User      *User    `gorm:"belongsTo"`
}

func (TwoFactorChallenge) TableName() string {
	return "two_factor_challenges"
}

func (m *TwoFactorChallenge) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Account struct {
	ID                    string          `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID                string          `gorm:"column:userId;not null" json:"userId"`
	Email                 *string         `gorm:"column:email" json:"email"`
	Provider              string          `gorm:"column:provider;not null" json:"provider"`
	ProviderID            *string         `gorm:"column:providerId" json:"providerId"`
	AccountID             *string         `gorm:"column:accountId" json:"accountId"`
	IDToken               *string         `gorm:"column:idToken" json:"idToken"`
	AccessToken           *string         `gorm:"column:accessToken" json:"accessToken"`
	RefreshToken          *string         `gorm:"column:refreshToken" json:"refreshToken"`
	AccessTokenExpiresAt  *DateTime       `gorm:"column:accessTokenExpiresAt;type:timestamp(3)" json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt *DateTime       `gorm:"column:refreshTokenExpiresAt;type:timestamp(3)" json:"refreshTokenExpiresAt"`
	Scope                 *string         `gorm:"column:scope" json:"scope"`
	Password              *string         `gorm:"column:password" json:"password"`
	CreatedAt             DateTime        `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime        `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User                  *User           `gorm:"belongsTo"`
	ResetPasswords        []ResetPassword `gorm:"foreignKey:AccountID;constraint:reset_password_accountId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Account) TableName() string {
	return "accounts"
}

func (m *Account) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type Session struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Token     *string  `gorm:"column:token;uniqueIndex:sessions_token_key" json:"token"`
	UserID    string   `gorm:"column:userId;not null" json:"userId"`
	ExpiresAt DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	IPAddress *string  `gorm:"column:ipAddress" json:"ipAddress"`
	UserAgent *string  `gorm:"column:userAgent" json:"userAgent"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User      *User    `gorm:"belongsTo"`
}

func (Session) TableName() string {
	return "sessions"
}

func (m *Session) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type Verification struct {
	ID         string   `gorm:"column:id;primaryKey" json:"id"`
	Identifier string   `gorm:"column:identifier;index:verifications_identifier_idx;not null" json:"identifier"`
	Value      string   `gorm:"column:value;not null" json:"value"`
	ExpiresAt  DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Verification) TableName() string {
	return "verifications"
}

func (m *Verification) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Verification) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Member struct {
	ID             string          `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Role           string          `gorm:"column:role;not null" json:"role"`
	Email          string          `gorm:"column:email;not null" json:"email"`
	UserID         *string         `gorm:"column:userId;uniqueIndex:members_organizationId_userId_key" json:"userId"`
	InvitedByID    *string         `gorm:"column:invitedById" json:"invitedById"`
	OrganizationID string          `gorm:"column:organizationId;uniqueIndex:members_organizationId_userId_key;not null" json:"organizationId"`
	ProjectScope   string          `gorm:"column:projectScope;default:'all';not null" json:"projectScope"`
	Meta           datatypes.JSON  `gorm:"column:meta" json:"meta"`
	CreatedAt      DateTime        `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime        `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User           *User           `gorm:"foreignKey:UserID;belongsTo"`
	InvitedBy      *User           `gorm:"foreignKey:InvitedByID;belongsTo"`
	Organization   *Organization   `gorm:"belongsTo"`
	Projects       []MemberProject `gorm:"foreignKey:MemberID;constraint:member_projects_memberId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Member) TableName() string {
	return "members"
}

func (m *Member) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Member) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type MemberProject struct {
	MemberID  string   `gorm:"column:memberId;primaryKey;type:char(36)" json:"memberId"`
	ProjectID string   `gorm:"column:projectId;primaryKey;index:member_projects_projectId_idx" json:"projectId"`
	Level     *string  `gorm:"column:level" json:"level"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Member    *Member  `gorm:"belongsTo"`
	Project   *Project `gorm:"belongsTo"`
}

func (MemberProject) TableName() string {
	return "member_projects"
}

func (m *MemberProject) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Invite struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	Email          string `gorm:"column:email;not null" json:"email"`
	CreatedByID    string `gorm:"column:createdById;not null" json:"createdById"`
	OrganizationID string `gorm:"column:organizationId;not null" json:"organizationId"`
	// [IPrismaInviteProjectAccess]
	ProjectAccess datatypes.JSON `gorm:"column:projectAccess;default:'[]';not null" json:"projectAccess"`
	ProjectScope  string         `gorm:"column:projectScope;default:'all';not null" json:"projectScope"`
	ExpiresAt     DateTime       `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Role          string         `gorm:"column:role;not null" json:"role"`
	CreatedBy     *User          `gorm:"foreignKey:CreatedByID;belongsTo"`
	Organization  *Organization  `gorm:"belongsTo"`
}

func (Invite) TableName() string {
	return "invites"
}

func (m *Invite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Invite) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Project struct {
	ID                         string             `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name                       string             `gorm:"column:name;not null" json:"name"`
	OrganizationID             string             `gorm:"column:organizationId;not null" json:"organizationId"`
	EventsCount                int                `gorm:"column:eventsCount;size:32;default:0;not null" json:"eventsCount"`
	Types                      ScalarList[string] `gorm:"column:types;type:text[]" json:"types"`
	Domain                     *string            `gorm:"column:domain" json:"domain"`
	Cors                       ScalarList[string] `gorm:"column:cors;type:text[]" json:"cors"`
	CrossDomain                bool               `gorm:"column:crossDomain;default:false;not null" json:"crossDomain"`
	AllowUnsafeRevenueTracking bool               `gorm:"column:allowUnsafeRevenueTracking;default:false;not null" json:"allowUnsafeRevenueTracking"`
	AiStoreContent             *bool              `gorm:"column:aiStoreContent;default:true;not null" json:"aiStoreContent"`
	// [IPrismaProjectFilters]
	Filters               datatypes.JSON     `gorm:"column:filters;default:'[]';not null" json:"filters"`
	FirstEventAt          *DateTime          `gorm:"column:firstEventAt;type:timestamp(3)" json:"firstEventAt"`
	NoDataNotifiedAt      *DateTime          `gorm:"column:noDataNotifiedAt;type:timestamp(3)" json:"noDataNotifiedAt"`
	DataStoppedNotifiedAt *DateTime          `gorm:"column:dataStoppedNotifiedAt;type:timestamp(3)" json:"dataStoppedNotifiedAt"`
	DeleteAt              *DateTime          `gorm:"column:deleteAt;type:timestamp(3)" json:"deleteAt"`
	CreatedAt             DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization          *Organization      `gorm:"belongsTo"`
	Clients               []Client           `gorm:"foreignKey:ProjectID;constraint:clients_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Reports               []Report           `gorm:"foreignKey:ProjectID;constraint:reports_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Dashboards            []Dashboard        `gorm:"foreignKey:ProjectID;constraint:dashboards_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShareDashboards       []ShareDashboard   `gorm:"foreignKey:ProjectID;constraint:share_dashboards_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShareReports          []ShareReport      `gorm:"foreignKey:ProjectID;constraint:share_reports_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShareWidgets          []ShareWidget      `gorm:"foreignKey:ProjectID;constraint:share_widgets_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Meta                  []EventMeta        `gorm:"foreignKey:ProjectID;constraint:event_meta_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomEvents          []CustomEvent      `gorm:"foreignKey:ProjectID;constraint:custom_events_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	References            []Reference        `gorm:"foreignKey:ProjectID;constraint:references_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	MemberProjects        []MemberProject    `gorm:"foreignKey:ProjectID;constraint:member_projects_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	NotificationRules     []NotificationRule `gorm:"foreignKey:ProjectID;constraint:notification_rules_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Notifications         []Notification     `gorm:"foreignKey:ProjectID;constraint:notifications_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Imports               []Import           `gorm:"foreignKey:ProjectID;constraint:imports_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Cohorts               []Cohort           `gorm:"foreignKey:ProjectID;constraint:cohorts_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExportWatermarks      []ExportWatermark  `gorm:"foreignKey:ProjectID;constraint:export_watermarks_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Integrations          []Integration      `gorm:"foreignKey:ProjectID;constraint:integrations_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Conversations         []Conversation     `gorm:"foreignKey:ProjectID;constraint:conversations_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Share                 *ShareOverview     `gorm:"foreignKey:ProjectID;constraint:shares_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	GscConnection         *GscConnection     `gorm:"foreignKey:ProjectID;constraint:gsc_connections_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Project) TableName() string {
	return "projects"
}

func (m *Project) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type Cohort struct {
	ID          string  `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Name        string  `gorm:"column:name;not null" json:"name"`
	Description *string `gorm:"column:description" json:"description"`
	ProjectID   string  `gorm:"column:projectId;index:cohorts_projectId_idx;not null" json:"projectId"`
	// [IPrismaCohortDefinition]
	Definition     datatypes.JSON `gorm:"column:definition;default:'{}';not null" json:"definition"`
	IsStatic       bool           `gorm:"column:isStatic;default:false;not null" json:"isStatic"`
	ProfileCount   int            `gorm:"column:profileCount;size:32;default:0;not null" json:"profileCount"`
	LastComputedAt *DateTime      `gorm:"column:lastComputedAt;type:timestamp(3)" json:"lastComputedAt"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project        *Project       `gorm:"belongsTo"`
}

func (Cohort) TableName() string {
	return "cohorts"
}

func (m *Cohort) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Cohort) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Salt struct {
	Salt      string   `gorm:"column:salt;primaryKey" json:"salt"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Salt) TableName() string {
	return "salts"
}

func (m *Salt) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Salt) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Client struct {
	ID                  string        `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Name                string        `gorm:"column:name;not null" json:"name"`
	Secret              *string       `gorm:"column:secret" json:"secret"`
	Type                string        `gorm:"column:type;default:'write';not null" json:"type"`
	ProjectID           *string       `gorm:"column:projectId" json:"projectId"`
	OrganizationID      string        `gorm:"column:organizationId;not null" json:"organizationId"`
	IgnoreCorsAndSecret bool          `gorm:"column:ignoreCorsAndSecret;default:false;not null" json:"ignoreCorsAndSecret"`
	CreatedAt           DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project             *Project      `gorm:"belongsTo"`
	Organization        *Organization `gorm:"belongsTo"`
}

func (Client) TableName() string {
	return "clients"
}

func (m *Client) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Client) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Dashboard struct {
	ID             string          `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name           string          `gorm:"column:name;not null" json:"name"`
	OrganizationID string          `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID      string          `gorm:"column:projectId;not null" json:"projectId"`
	Filters        datatypes.JSON  `gorm:"column:filters;default:'[]';not null" json:"filters"`
	CreatedAt      DateTime        `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime        `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization   `gorm:"belongsTo"`
	Project        *Project        `gorm:"belongsTo"`
	Reports        []Report        `gorm:"foreignKey:DashboardID;constraint:reports_dashboardId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Share          *ShareDashboard `gorm:"foreignKey:DashboardID;constraint:share_dashboards_dashboardId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Dashboard) TableName() string {
	return "dashboards"
}

func (m *Dashboard) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type Report struct {
	ID            string         `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	Interval      string         `gorm:"column:interval;not null" json:"interval"`
	Range         *string        `gorm:"column:range;default:'30d';not null" json:"range"`
	ChartType     string         `gorm:"column:chartType;not null" json:"chartType"`
	LineType      *string        `gorm:"column:lineType;default:'monotone';not null" json:"lineType"`
	Breakdowns    datatypes.JSON `gorm:"column:breakdowns;not null" json:"breakdowns"`
	Events        datatypes.JSON `gorm:"column:events;not null" json:"events"`
	GlobalFilters datatypes.JSON `gorm:"column:globalFilters;default:'[]';not null" json:"globalFilters"`
	Formula       *string        `gorm:"column:formula" json:"formula"`
	Unit          *string        `gorm:"column:unit" json:"unit"`
	Metric        string         `gorm:"column:metric;default:'sum';not null" json:"metric"`
	ProjectID     string         `gorm:"column:projectId;not null" json:"projectId"`
	Previous      bool           `gorm:"column:previous;default:false;not null" json:"previous"`
	Criteria      *string        `gorm:"column:criteria" json:"criteria"`
	FunnelGroup   *string        `gorm:"column:funnelGroup" json:"funnelGroup"`
	FunnelWindow  *float64       `gorm:"column:funnelWindow;type:double precision" json:"funnelWindow"`
	// [IReportOptions]
	Options datatypes.JSON `gorm:"column:options" json:"options"`
	// How a dashboard tile shows the report: { view: 'chart' | 'table' }
	Display       datatypes.JSON     `gorm:"column:display" json:"display"`
	VisibleSeries ScalarList[string] `gorm:"column:visibleSeries;type:text[]" json:"visibleSeries"`
	StartDate     *string            `gorm:"column:startDate" json:"startDate"`
	EndDate       *string            `gorm:"column:endDate" json:"endDate"`
	DashboardID   string             `gorm:"column:dashboardId;not null" json:"dashboardId"`
	CreatedAt     DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project       *Project           `gorm:"belongsTo"`
	Dashboard     *Dashboard         `gorm:"belongsTo"`
	Layout        *ReportLayout      `gorm:"foreignKey:ReportID;constraint:report_layouts_reportId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Share         *ShareReport       `gorm:"foreignKey:ReportID;constraint:share_reports_reportId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Report) TableName() string {
	return "reports"
}

func (m *Report) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Report) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ReportLayout struct {
	ID          string   `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	ReportID    string   `gorm:"column:reportId;uniqueIndex:report_layouts_reportId_key;type:char(36);not null" json:"reportId"`
	GridColumns *int     `gorm:"column:gridColumns;size:32;default:12;not null" json:"gridColumns"`
	X           int      `gorm:"column:x;size:32;default:0;not null" json:"x"`
	Y           int      `gorm:"column:y;size:32;default:0;not null" json:"y"`
	W           *int     `gorm:"column:w;size:32;default:8;not null" json:"w"`
	H           *int     `gorm:"column:h;size:32;default:6;not null" json:"h"`
	MinW        *int     `gorm:"column:minW;size:32;default:4" json:"minW"`
	MinH        *int     `gorm:"column:minH;size:32;default:4" json:"minH"`
	MaxW        *int     `gorm:"column:maxW;size:32" json:"maxW"`
	MaxH        *int     `gorm:"column:maxH;size:32" json:"maxH"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Report      *Report  `gorm:"belongsTo"`
}

func (ReportLayout) TableName() string {
	return "report_layouts"
}

func (m *ReportLayout) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ReportLayout) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ShareOverview struct {
	ID             string        `gorm:"column:id;primaryKey;uniqueIndex:shares_id_key" json:"id"`
	ProjectID      string        `gorm:"column:projectId;uniqueIndex:shares_projectId_key;not null" json:"projectId"`
	OrganizationID string        `gorm:"column:organizationId;not null" json:"organizationId"`
	Public         bool          `gorm:"column:public;default:false;not null" json:"public"`
	Password       *string       `gorm:"column:password" json:"password"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project        *Project      `gorm:"belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (ShareOverview) TableName() string {
	return "shares"
}

func (m *ShareOverview) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ShareOverview) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ShareDashboard struct {
	ID             string        `gorm:"column:id;primaryKey;uniqueIndex:share_dashboards_id_key" json:"id"`
	DashboardID    string        `gorm:"column:dashboardId;uniqueIndex:share_dashboards_dashboardId_key;not null" json:"dashboardId"`
	OrganizationID string        `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID      string        `gorm:"column:projectId;not null" json:"projectId"`
	Public         bool          `gorm:"column:public;default:false;not null" json:"public"`
	Password       *string       `gorm:"column:password" json:"password"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Dashboard      *Dashboard    `gorm:"belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
	Project        *Project      `gorm:"belongsTo"`
}

func (ShareDashboard) TableName() string {
	return "share_dashboards"
}

func (m *ShareDashboard) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ShareDashboard) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ShareReport struct {
	ID             string        `gorm:"column:id;primaryKey;uniqueIndex:share_reports_id_key" json:"id"`
	ReportID       string        `gorm:"column:reportId;uniqueIndex:share_reports_reportId_key;type:char(36);not null" json:"reportId"`
	OrganizationID string        `gorm:"column:organizationId;not null" json:"organizationId"`
	ProjectID      string        `gorm:"column:projectId;not null" json:"projectId"`
	Public         bool          `gorm:"column:public;default:false;not null" json:"public"`
	Password       *string       `gorm:"column:password" json:"password"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Report         *Report       `gorm:"belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
	Project        *Project      `gorm:"belongsTo"`
}

func (ShareReport) TableName() string {
	return "share_reports"
}

func (m *ShareReport) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ShareReport) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ShareWidget struct {
	ID             string `gorm:"column:id;primaryKey;uniqueIndex:share_widgets_id_key" json:"id"`
	ProjectID      string `gorm:"column:projectId;not null" json:"projectId"`
	OrganizationID string `gorm:"column:organizationId;not null" json:"organizationId"`
	Public         *bool  `gorm:"column:public;default:true;not null" json:"public"`
	// [IPrismaWidgetOptions]
	Options      datatypes.JSON `gorm:"column:options;not null" json:"options"`
	CreatedAt    DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project      *Project       `gorm:"belongsTo"`
	Organization *Organization  `gorm:"belongsTo"`
}

func (ShareWidget) TableName() string {
	return "share_widgets"
}

func (m *ShareWidget) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ShareWidget) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EventMeta struct {
	ID         string   `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Name       string   `gorm:"column:name;uniqueIndex:event_meta_name_projectId_key;not null" json:"name"`
	Conversion *bool    `gorm:"column:conversion" json:"conversion"`
	Color      *string  `gorm:"column:color" json:"color"`
	Icon       *string  `gorm:"column:icon" json:"icon"`
	ProjectID  string   `gorm:"column:projectId;uniqueIndex:event_meta_name_projectId_key;not null" json:"projectId"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project    *Project `gorm:"belongsTo"`
}

func (EventMeta) TableName() string {
	return "event_meta"
}

func (m *EventMeta) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EventMeta) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CustomEvent struct {
	ID   string `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Name string `gorm:"column:name;uniqueIndex:custom_events_projectId_name_key;not null" json:"name"`
	// [IPrismaCustomEventItems]
	Items     datatypes.JSON `gorm:"column:items;not null" json:"items"`
	ProjectID string         `gorm:"column:projectId;uniqueIndex:custom_events_projectId_name_key;not null" json:"projectId"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project   *Project       `gorm:"belongsTo"`
}

func (CustomEvent) TableName() string {
	return "custom_events"
}

func (m *CustomEvent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CustomEvent) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Reference struct {
	ID          string   `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Title       string   `gorm:"column:title;not null" json:"title"`
	Description *string  `gorm:"column:description" json:"description"`
	Date        DateTime `gorm:"column:date;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"date"`
	ProjectID   string   `gorm:"column:projectId;not null" json:"projectId"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project     *Project `gorm:"belongsTo"`
}

func (Reference) TableName() string {
	return "references"
}

func (m *Reference) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Date.IsZero() {
		m.Date = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Reference) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type NotificationRule struct {
	ID          string `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	ProjectID   string `gorm:"column:projectId;not null" json:"projectId"`
	SendToApp   bool   `gorm:"column:sendToApp;default:false;not null" json:"sendToApp"`
	SendToEmail bool   `gorm:"column:sendToEmail;default:false;not null" json:"sendToEmail"`
	// [IPrismaNotificationRuleConfig]
	Config        datatypes.JSON `gorm:"column:config;not null" json:"config"`
	Template      *string        `gorm:"column:template" json:"template"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project       *Project       `gorm:"belongsTo"`
	Notifications []Notification `gorm:"foreignKey:NotificationRuleID;constraint:notifications_notificationRuleId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Integrations  []Integration  `gorm:"many2many:_IntegrationToNotificationRule;joinForeignKey:B;joinReferences:A"`
}

func (NotificationRule) TableName() string {
	return "notification_rules"
}

func (m *NotificationRule) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*NotificationRule) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Notification struct {
	ID                 string    `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	ProjectID          string    `gorm:"column:projectId;not null" json:"projectId"`
	Title              string    `gorm:"column:title;not null" json:"title"`
	Message            string    `gorm:"column:message;not null" json:"message"`
	IsReadAt           *DateTime `gorm:"column:isReadAt;type:timestamp(3)" json:"isReadAt"`
	CreatedAt          DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	SendToApp          bool      `gorm:"column:sendToApp;default:false;not null" json:"sendToApp"`
	SendToEmail        bool      `gorm:"column:sendToEmail;default:false;not null" json:"sendToEmail"`
	IntegrationID      *string   `gorm:"column:integrationId;type:char(36)" json:"integrationId"`
	NotificationRuleID *string   `gorm:"column:notificationRuleId;type:char(36)" json:"notificationRuleId"`
	// [IPrismaNotificationPayload]
	Payload          datatypes.JSON    `gorm:"column:payload" json:"payload"`
	Project          *Project          `gorm:"belongsTo"`
	Integration      *Integration      `gorm:"belongsTo"`
	NotificationRule *NotificationRule `gorm:"belongsTo"`
}

func (Notification) TableName() string {
	return "notifications"
}

func (m *Notification) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type Integration struct {
	ID   string `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
	// [IPrismaIntegrationConfig]
	Config            datatypes.JSON     `gorm:"column:config;not null" json:"config"`
	OrganizationID    string             `gorm:"column:organizationId;index:integrations_organizationId_idx;not null" json:"organizationId"`
	ProjectID         *string            `gorm:"column:projectId;index:integrations_projectId_idx" json:"projectId"`
	CreatedAt         DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Organization      *Organization      `gorm:"belongsTo"`
	Project           *Project           `gorm:"belongsTo"`
	Notifications     []Notification     `gorm:"foreignKey:IntegrationID;constraint:notifications_integrationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExportWatermarks  []ExportWatermark  `gorm:"foreignKey:IntegrationID;constraint:export_watermarks_integrationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	NotificationRules []NotificationRule `gorm:"many2many:_IntegrationToNotificationRule;joinForeignKey:A;joinReferences:B"`
}

func (Integration) TableName() string {
	return "integrations"
}

func (m *Integration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Integration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ExportWatermark struct {
	ID             string       `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	ProjectID      string       `gorm:"column:projectId;uniqueIndex:export_watermarks_projectId_integrationId_key;not null" json:"projectId"`
	IntegrationID  string       `gorm:"column:integrationId;uniqueIndex:export_watermarks_projectId_integrationId_key;index:export_watermarks_integrationId_idx;type:char(36);not null" json:"integrationId"`
	LastInsertedAt DateTime     `gorm:"column:lastInsertedAt;type:timestamp(3);not null" json:"lastInsertedAt"`
	LastEventID    string       `gorm:"column:lastEventId;default:'';not null" json:"lastEventId"`
	CreatedAt      DateTime     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime     `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project        *Project     `gorm:"belongsTo"`
	Integration    *Integration `gorm:"belongsTo"`
}

func (ExportWatermark) TableName() string {
	return "export_watermarks"
}

func (m *ExportWatermark) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ExportWatermark) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ResetPassword struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	AccountID string   `gorm:"column:accountId;not null" json:"accountId"`
	ExpiresAt DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Account   *Account `gorm:"belongsTo"`
}

func (ResetPassword) TableName() string {
	return "reset_password"
}

func (m *ResetPassword) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ResetPassword) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Import struct {
	ID            string  `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	ProjectID     string  `gorm:"column:projectId;not null" json:"projectId"`
	JobID         *string `gorm:"column:jobId" json:"jobId"`
	Status        string  `gorm:"column:status;not null" json:"status"`
	StatusMessage *string `gorm:"column:statusMessage" json:"statusMessage"`
	ErrorMessage  *string `gorm:"column:errorMessage" json:"errorMessage"`
	// [IPrismaImportConfig]
	Config          datatypes.JSON `gorm:"column:config;not null" json:"config"`
	TotalEvents     int            `gorm:"column:totalEvents;size:32;default:0;not null" json:"totalEvents"`
	ProcessedEvents int            `gorm:"column:processedEvents;size:32;default:0;not null" json:"processedEvents"`
	CurrentStep     *string        `gorm:"column:currentStep" json:"currentStep"`
	CurrentBatch    *string        `gorm:"column:currentBatch" json:"currentBatch"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	CompletedAt     *DateTime      `gorm:"column:completedAt;type:timestamp(3)" json:"completedAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project         *Project       `gorm:"belongsTo"`
}

func (Import) TableName() string {
	return "imports"
}

func (m *Import) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Import) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ProjectInsight struct {
	ID           string  `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	ProjectID    string  `gorm:"column:projectId;uniqueIndex:project_insights_projectId_moduleKey_dimensionKey_windowKin_key;index:project_insights_projectId_impactScore_idx;index:project_insights_projectId_moduleKey_windowKind_state_idx;index:project_insights_projectId_state_relevanceScore_idx;not null" json:"projectId"`
	ModuleKey    string  `gorm:"column:moduleKey;uniqueIndex:project_insights_projectId_moduleKey_dimensionKey_windowKin_key;index:project_insights_projectId_moduleKey_windowKind_state_idx;not null" json:"moduleKey"`
	DimensionKey string  `gorm:"column:dimensionKey;uniqueIndex:project_insights_projectId_moduleKey_dimensionKey_windowKin_key;not null" json:"dimensionKey"`
	WindowKind   string  `gorm:"column:windowKind;uniqueIndex:project_insights_projectId_moduleKey_dimensionKey_windowKin_key;index:project_insights_projectId_moduleKey_windowKind_state_idx;not null" json:"windowKind"`
	State        string  `gorm:"column:state;uniqueIndex:project_insights_projectId_moduleKey_dimensionKey_windowKin_key;index:project_insights_projectId_moduleKey_windowKind_state_idx;index:project_insights_projectId_state_relevanceScore_idx;default:'active';not null" json:"state"`
	Title        string  `gorm:"column:title;not null" json:"title"`
	Summary      *string `gorm:"column:summary" json:"summary"`
	DisplayName  string  `gorm:"column:displayName;default:'';not null" json:"displayName"`
	// [IPrismaProjectInsightPayload]
	Payload         datatypes.JSON `gorm:"column:payload;default:'{}';not null" json:"payload"`
	Direction       *string        `gorm:"column:direction" json:"direction"`
	ImpactScore     float64        `gorm:"column:impactScore;index:project_insights_projectId_impactScore_idx;type:double precision;default:0;not null" json:"impactScore"`
	SeverityBand    *string        `gorm:"column:severityBand" json:"severityBand"`
	RelevanceScore  *float64       `gorm:"column:relevanceScore;index:project_insights_projectId_state_relevanceScore_idx;type:double precision" json:"relevanceScore"`
	AiSummary       *string        `gorm:"column:aiSummary" json:"aiSummary"`
	AiCategory      *string        `gorm:"column:aiCategory" json:"aiCategory"`
	EmailWorthy     bool           `gorm:"column:emailWorthy;default:false;not null" json:"emailWorthy"`
	ReferenceWorthy bool           `gorm:"column:referenceWorthy;default:false;not null" json:"referenceWorthy"`
	EnrichedAt      *DateTime      `gorm:"column:enrichedAt;type:timestamp(3)" json:"enrichedAt"`
	EnrichVersion   *int           `gorm:"column:enrichVersion;size:32" json:"enrichVersion"`
	Version         *int           `gorm:"column:version;size:32;default:1;not null" json:"version"`
	ThreadID        string         `gorm:"column:threadId;type:char(36);default:gen_random_uuid();not null" json:"threadId"`
	WindowStart     *DateTime      `gorm:"column:windowStart;type:timestamp(3)" json:"windowStart"`
	WindowEnd       *DateTime      `gorm:"column:windowEnd;type:timestamp(3)" json:"windowEnd"`
	FirstDetectedAt DateTime       `gorm:"column:firstDetectedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"firstDetectedAt"`
	LastUpdatedAt   DateTime       `gorm:"column:lastUpdatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastUpdatedAt"`
	LastSeenAt      DateTime       `gorm:"column:lastSeenAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastSeenAt"`
	Events          []InsightEvent `gorm:"foreignKey:InsightID;constraint:insight_events_insightId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ProjectInsight) TableName() string {
	return "project_insights"
}

func (m *ProjectInsight) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.FirstDetectedAt.IsZero() {
		m.FirstDetectedAt = DateTime{Time: now}
	}
	if m.LastUpdatedAt.IsZero() {
		m.LastUpdatedAt = DateTime{Time: now}
	}
	if m.LastSeenAt.IsZero() {
		m.LastSeenAt = DateTime{Time: now}
	}
	return nil
}

func (*ProjectInsight) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastUpdatedAt") {
		tx.Statement.SetColumn("LastUpdatedAt", DateTime{Time: now})
	}
	return nil
}

type InsightEvent struct {
	ID         string          `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	InsightID  string          `gorm:"column:insightId;index:insight_events_insightId_createdAt_idx;type:char(36);not null" json:"insightId"`
	EventKind  string          `gorm:"column:eventKind;not null" json:"eventKind"`
	ChangeFrom datatypes.JSON  `gorm:"column:changeFrom" json:"changeFrom"`
	ChangeTo   datatypes.JSON  `gorm:"column:changeTo" json:"changeTo"`
	CreatedAt  DateTime        `gorm:"column:createdAt;index:insight_events_insightId_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Insight    *ProjectInsight `gorm:"foreignKey:InsightID;belongsTo"`
}

func (InsightEvent) TableName() string {
	return "insight_events"
}

func (m *InsightEvent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type GscConnection struct {
	ID                   string    `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	ProjectID            string    `gorm:"column:projectId;uniqueIndex:gsc_connections_projectId_key;not null" json:"projectId"`
	SiteURL              string    `gorm:"column:siteUrl;default:'';not null" json:"siteUrl"`
	AccessToken          string    `gorm:"column:accessToken;not null" json:"accessToken"`
	RefreshToken         string    `gorm:"column:refreshToken;not null" json:"refreshToken"`
	AccessTokenExpiresAt *DateTime `gorm:"column:accessTokenExpiresAt;type:timestamp(3)" json:"accessTokenExpiresAt"`
	LastSyncedAt         *DateTime `gorm:"column:lastSyncedAt;type:timestamp(3)" json:"lastSyncedAt"`
	LastSyncStatus       *string   `gorm:"column:lastSyncStatus" json:"lastSyncStatus"`
	LastSyncError        *string   `gorm:"column:lastSyncError" json:"lastSyncError"`
	BackfillStatus       *string   `gorm:"column:backfillStatus" json:"backfillStatus"`
	CreatedAt            DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Project              *Project  `gorm:"belongsTo"`
}

func (GscConnection) TableName() string {
	return "gsc_connections"
}

func (m *GscConnection) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GscConnection) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EmailUnsubscribe struct {
	ID        string   `gorm:"column:id;primaryKey;type:char(36);default:gen_random_uuid()" json:"id"`
	Email     string   `gorm:"column:email;uniqueIndex:email_unsubscribes_email_category_key;not null" json:"email"`
	Category  string   `gorm:"column:category;uniqueIndex:email_unsubscribes_email_category_key;not null" json:"category"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (EmailUnsubscribe) TableName() string {
	return "email_unsubscribes"
}

func (m *EmailUnsubscribe) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
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
