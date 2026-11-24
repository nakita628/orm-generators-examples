package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID                   int                    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name                 *string                `gorm:"column:name" json:"name"`
	Email                string                 `gorm:"column:email;uniqueIndex:User_email_key;index:User_email_idx;not null" json:"email"`
	EmailVerified        *DateTime              `gorm:"column:emailVerified;type:timestamp(3)" json:"emailVerified"`
	Password             *string                `gorm:"column:password" json:"password"`
	Source               *string                `gorm:"column:source" json:"source"`
	Signature            *string                `gorm:"column:signature" json:"signature"`
	CreatedAt            DateTime               `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime               `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	LastSignedIn         DateTime               `gorm:"column:lastSignedIn;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastSignedIn"`
	Roles                ScalarList[string]     `gorm:"column:roles;type:text[]" json:"roles"`
	IdentityProvider     string                 `gorm:"column:identityProvider;default:'DOCUMENSO';not null" json:"identityProvider"`
	AvatarImageID        *string                `gorm:"column:avatarImageId" json:"avatarImageId"`
	Disabled             bool                   `gorm:"column:disabled;default:false;not null" json:"disabled"`
	TwoFactorSecret      *string                `gorm:"column:twoFactorSecret" json:"twoFactorSecret"`
	TwoFactorEnabled     bool                   `gorm:"column:twoFactorEnabled;default:false;not null" json:"twoFactorEnabled"`
	TwoFactorBackupCodes *string                `gorm:"column:twoFactorBackupCodes" json:"twoFactorBackupCodes"`
	AvatarImage          *AvatarImage           `gorm:"belongsTo"`
	Accounts             []Account              `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions             []Session              `gorm:"foreignKey:UserID;constraint:Session_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PasswordResetTokens  []PasswordResetToken   `gorm:"foreignKey:UserID;constraint:PasswordResetToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OwnedOrganisations   []Organisation         `gorm:"foreignKey:OwnerUserID;constraint:Organisation_ownerUserId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganisationMember   []OrganisationMember   `gorm:"foreignKey:UserID;constraint:OrganisationMember_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Folders              []Folder               `gorm:"foreignKey:UserID;constraint:Folder_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Envelopes            []Envelope             `gorm:"foreignKey:UserID;constraint:Envelope_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	VerificationTokens   []VerificationToken    `gorm:"foreignKey:UserID;constraint:VerificationToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APITokens            []APIToken             `gorm:"foreignKey:UserID;constraint:ApiToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SecurityAuditLogs    []UserSecurityAuditLog `gorm:"foreignKey:UserID;constraint:UserSecurityAuditLog_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks             []Webhook              `gorm:"foreignKey:UserID;constraint:Webhook_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SiteSettings         []SiteSettings         `gorm:"foreignKey:LastModifiedByUserID;constraint:SiteSettings_lastModifiedByUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Passkeys             []Passkey              `gorm:"foreignKey:UserID;constraint:Passkey_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "User"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	if m.LastSignedIn.IsZero() {
		m.LastSignedIn = DateTime{Time: now}
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

type TeamProfile struct {
	ID      string  `gorm:"column:id;primaryKey" json:"id"`
	Enabled bool    `gorm:"column:enabled;default:false;not null" json:"enabled"`
	TeamID  int     `gorm:"column:teamId;uniqueIndex:TeamProfile_teamId_key;size:32;not null" json:"teamId"`
	Bio     *string `gorm:"column:bio" json:"bio"`
	Team    *Team   `gorm:"belongsTo"`
}

func (TeamProfile) TableName() string {
	return "TeamProfile"
}

func (m *TeamProfile) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type UserSecurityAuditLog struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID    int      `gorm:"column:userId;size:32;not null" json:"userId"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Type      string   `gorm:"column:type;not null" json:"type"`
	UserAgent *string  `gorm:"column:userAgent" json:"userAgent"`
	IPAddress *string  `gorm:"column:ipAddress" json:"ipAddress"`
	User      *User    `gorm:"belongsTo"`
}

func (UserSecurityAuditLog) TableName() string {
	return "UserSecurityAuditLog"
}

func (m *UserSecurityAuditLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type PasswordResetToken struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Token     string   `gorm:"column:token;uniqueIndex:PasswordResetToken_token_key;not null" json:"token"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Expiry    DateTime `gorm:"column:expiry;type:timestamp(3);not null" json:"expiry"`
	UserID    int      `gorm:"column:userId;size:32;not null" json:"userId"`
	User      *User    `gorm:"belongsTo"`
}

func (PasswordResetToken) TableName() string {
	return "PasswordResetToken"
}

func (m *PasswordResetToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Passkey struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	UserID     int       `gorm:"column:userId;size:32;not null" json:"userId"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	LastUsedAt *DateTime `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	// @zod.custom.use(z.instanceof(Uint8Array))
	CredentialID []byte `gorm:"column:credentialId;not null" json:"credentialId"`
	// @zod.custom.use(z.instanceof(Uint8Array))
	CredentialPublicKey  []byte             `gorm:"column:credentialPublicKey;not null" json:"credentialPublicKey"`
	Counter              int64              `gorm:"column:counter;not null" json:"counter"`
	CredentialDeviceType string             `gorm:"column:credentialDeviceType;not null" json:"credentialDeviceType"`
	CredentialBackedUp   bool               `gorm:"column:credentialBackedUp;not null" json:"credentialBackedUp"`
	Transports           ScalarList[string] `gorm:"column:transports;type:text[]" json:"transports"`
	User                 *User              `gorm:"belongsTo"`
}

func (Passkey) TableName() string {
	return "Passkey"
}

func (m *Passkey) BeforeCreate(tx *gorm.DB) error {
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

type AnonymousVerificationToken struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Token     string   `gorm:"column:token;uniqueIndex:AnonymousVerificationToken_token_key;not null" json:"token"`
	ExpiresAt DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (AnonymousVerificationToken) TableName() string {
	return "AnonymousVerificationToken"
}

func (m *AnonymousVerificationToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type VerificationToken struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	SecondaryID string         `gorm:"column:secondaryId;uniqueIndex:VerificationToken_secondaryId_key;not null" json:"secondaryId"`
	Identifier  string         `gorm:"column:identifier;not null" json:"identifier"`
	Token       string         `gorm:"column:token;uniqueIndex:VerificationToken_token_key;not null" json:"token"`
	Completed   bool           `gorm:"column:completed;default:false;not null" json:"completed"`
	Expires     DateTime       `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Metadata    datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	UserID      int            `gorm:"column:userId;size:32;not null" json:"userId"`
	User        *User          `gorm:"belongsTo"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

func (m *VerificationToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.SecondaryID == "" {
		m.SecondaryID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Webhook struct {
	ID            string             `gorm:"column:id;primaryKey" json:"id"`
	WebhookURL    string             `gorm:"column:webhookUrl;not null" json:"webhookUrl"`
	EventTriggers ScalarList[string] `gorm:"column:eventTriggers;type:text[]" json:"eventTriggers"`
	Secret        *string            `gorm:"column:secret" json:"secret"`
	Enabled       *bool              `gorm:"column:enabled;default:true;not null" json:"enabled"`
	CreatedAt     DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	UserID        int                `gorm:"column:userId;size:32;not null" json:"userId"`
	TeamID        int                `gorm:"column:teamId;size:32;not null" json:"teamId"`
	User          *User              `gorm:"belongsTo"`
	Team          *Team              `gorm:"belongsTo"`
	WebhookCalls  []WebhookCall      `gorm:"foreignKey:WebhookID;constraint:WebhookCall_webhookId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Webhook) TableName() string {
	return "Webhook"
}

func (m *Webhook) BeforeCreate(tx *gorm.DB) error {
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

func (*Webhook) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WebhookCall struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Status          string         `gorm:"column:status;not null" json:"status"`
	URL             string         `gorm:"column:url;not null" json:"url"`
	Event           string         `gorm:"column:event;not null" json:"event"`
	RequestBody     datatypes.JSON `gorm:"column:requestBody;not null" json:"requestBody"`
	ResponseCode    int            `gorm:"column:responseCode;size:32;not null" json:"responseCode"`
	ResponseHeaders datatypes.JSON `gorm:"column:responseHeaders" json:"responseHeaders"`
	ResponseBody    datatypes.JSON `gorm:"column:responseBody" json:"responseBody"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	WebhookID       string         `gorm:"column:webhookId;not null" json:"webhookId"`
	Webhook         *Webhook       `gorm:"belongsTo"`
}

func (WebhookCall) TableName() string {
	return "WebhookCall"
}

func (m *WebhookCall) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type APIToken struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Token      string    `gorm:"column:token;uniqueIndex:ApiToken_token_key;not null" json:"token"`
	Algorithm  string    `gorm:"column:algorithm;default:'SHA512';not null" json:"algorithm"`
	Expires    *DateTime `gorm:"column:expires;type:timestamp(3)" json:"expires"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	LastUsedAt *DateTime `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	UserID     *int      `gorm:"column:userId;size:32" json:"userId"`
	TeamID     int       `gorm:"column:teamId;size:32;not null" json:"teamId"`
	User       *User     `gorm:"belongsTo"`
	Team       *Team     `gorm:"belongsTo"`
}

func (APIToken) TableName() string {
	return "ApiToken"
}

func (m *APIToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Subscription struct {
	ID                int           `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Status            string        `gorm:"column:status;default:'INACTIVE';not null" json:"status"`
	PlanID            string        `gorm:"column:planId;uniqueIndex:Subscription_planId_key;not null" json:"planId"`
	PriceID           string        `gorm:"column:priceId;not null" json:"priceId"`
	PeriodEnd         *DateTime     `gorm:"column:periodEnd;type:timestamp(3)" json:"periodEnd"`
	CreatedAt         DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	CancelAtPeriodEnd bool          `gorm:"column:cancelAtPeriodEnd;default:false;not null" json:"cancelAtPeriodEnd"`
	CustomerID        string        `gorm:"column:customerId;not null" json:"customerId"`
	OrganisationID    string        `gorm:"column:organisationId;uniqueIndex:Subscription_organisationId_key;index:Subscription_organisationId_idx;not null" json:"organisationId"`
	Organisation      *Organisation `gorm:"belongsTo"`
}

func (Subscription) TableName() string {
	return "Subscription"
}

func (m *Subscription) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

// @zod.import(["import { ZClaimFlagsSchema, ZRateLimitArraySchema } from '@documenso/lib/types/subscription';"])
type SubscriptionClaim struct {
	ID                        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                 DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Name                      string   `gorm:"column:name;not null" json:"name"`
	Locked                    bool     `gorm:"column:locked;default:false;not null" json:"locked"`
	TeamCount                 int      `gorm:"column:teamCount;size:32;not null" json:"teamCount"`
	MemberCount               int      `gorm:"column:memberCount;size:32;not null" json:"memberCount"`
	EnvelopeItemCount         int      `gorm:"column:envelopeItemCount;size:32;not null" json:"envelopeItemCount"`
	RecipientCount            int      `gorm:"column:recipientCount;size:32;not null" json:"recipientCount"`
	EnvelopeContentCount      int      `gorm:"column:envelopeContentCount;size:32;not null" json:"envelopeContentCount"`
	EnvelopeContentImageCount int      `gorm:"column:envelopeContentImageCount;size:32;not null" json:"envelopeContentImageCount"`
	// [ClaimFlags] @zod.custom.use(ZClaimFlagsSchema)
	Flags datatypes.JSON `gorm:"column:flags;not null" json:"flags"`
	// [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema)
	DocumentRateLimits datatypes.JSON `gorm:"column:documentRateLimits;not null" json:"documentRateLimits"`
	DocumentQuota      *int           `gorm:"column:documentQuota;size:32" json:"documentQuota"`
	// [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema)
	EmailRateLimits datatypes.JSON `gorm:"column:emailRateLimits;not null" json:"emailRateLimits"`
	EmailQuota      *int           `gorm:"column:emailQuota;size:32" json:"emailQuota"`
	// [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema)
	APIRateLimits    datatypes.JSON  `gorm:"column:apiRateLimits;not null" json:"apiRateLimits"`
	APIQuota         *int            `gorm:"column:apiQuota;size:32" json:"apiQuota"`
	EmailTransportID *string         `gorm:"column:emailTransportId" json:"emailTransportId"`
	EmailTransport   *EmailTransport `gorm:"belongsTo"`
}

func (SubscriptionClaim) TableName() string {
	return "SubscriptionClaim"
}

func (m *SubscriptionClaim) BeforeCreate(tx *gorm.DB) error {
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

func (*SubscriptionClaim) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// @zod.import(["import { ZClaimFlagsSchema, ZRateLimitArraySchema } from '@documenso/lib/types/subscription';"])
type OrganisationClaim struct {
	ID                          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                   DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OriginalSubscriptionClaimID *string  `gorm:"column:originalSubscriptionClaimId" json:"originalSubscriptionClaimId"`
	TeamCount                   int      `gorm:"column:teamCount;size:32;not null" json:"teamCount"`
	MemberCount                 int      `gorm:"column:memberCount;size:32;not null" json:"memberCount"`
	EnvelopeItemCount           int      `gorm:"column:envelopeItemCount;size:32;not null" json:"envelopeItemCount"`
	RecipientCount              int      `gorm:"column:recipientCount;size:32;not null" json:"recipientCount"`
	EnvelopeContentCount        int      `gorm:"column:envelopeContentCount;size:32;not null" json:"envelopeContentCount"`
	EnvelopeContentImageCount   int      `gorm:"column:envelopeContentImageCount;size:32;not null" json:"envelopeContentImageCount"`
	// [ClaimFlags] @zod.custom.use(ZClaimFlagsSchema)
	Flags datatypes.JSON `gorm:"column:flags;not null" json:"flags"`
	// [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema)
	DocumentRateLimits datatypes.JSON `gorm:"column:documentRateLimits;not null" json:"documentRateLimits"`
	DocumentQuota      *int           `gorm:"column:documentQuota;size:32" json:"documentQuota"`
	// [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema)
	EmailRateLimits datatypes.JSON `gorm:"column:emailRateLimits;not null" json:"emailRateLimits"`
	EmailQuota      *int           `gorm:"column:emailQuota;size:32" json:"emailQuota"`
	// [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema)
	APIRateLimits    datatypes.JSON  `gorm:"column:apiRateLimits;not null" json:"apiRateLimits"`
	APIQuota         *int            `gorm:"column:apiQuota;size:32" json:"apiQuota"`
	EmailTransportID *string         `gorm:"column:emailTransportId" json:"emailTransportId"`
	EmailTransport   *EmailTransport `gorm:"belongsTo"`
	Organisation     *Organisation   `gorm:"foreignKey:OrganisationClaimID;constraint:Organisation_organisationClaimId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (OrganisationClaim) TableName() string {
	return "OrganisationClaim"
}

func (m *OrganisationClaim) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganisationClaim) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganisationMonthlyStat struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OrganisationID string   `gorm:"column:organisationId;uniqueIndex:OrganisationMonthlyStat_organisationId_period_key;index:OrganisationMonthlyStat_organisationId_idx;not null" json:"organisationId"`
	// UTC calendar month in `YYYY-MM` form, e.g. "2026-05".
	Period        string        `gorm:"column:period;uniqueIndex:OrganisationMonthlyStat_organisationId_period_key;not null" json:"period"`
	DocumentCount int           `gorm:"column:documentCount;size:32;default:0;not null" json:"documentCount"`
	EmailCount    int           `gorm:"column:emailCount;size:32;default:0;not null" json:"emailCount"`
	APICount      int           `gorm:"column:apiCount;size:32;default:0;not null" json:"apiCount"`
	EmailReports  int           `gorm:"column:emailReports;size:32;default:0;not null" json:"emailReports"`
	Organisation  *Organisation `gorm:"belongsTo"`
}

func (OrganisationMonthlyStat) TableName() string {
	return "OrganisationMonthlyStat"
}

func (m *OrganisationMonthlyStat) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganisationMonthlyStat) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Account struct {
	ID                string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt         DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UserID            int      `gorm:"column:userId;size:32;not null" json:"userId"`
	Type              string   `gorm:"column:type;not null" json:"type"`
	Provider          string   `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key;not null" json:"provider"`
	ProviderAccountID string   `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key;not null" json:"providerAccountId"`
	RefreshToken      *string  `gorm:"column:refresh_token;type:text" json:"refresh_token"`
	AccessToken       *string  `gorm:"column:access_token;type:text" json:"access_token"`
	ExpiresAt         *int     `gorm:"column:expires_at;size:32" json:"expires_at"`
	CreatedAt1        *int     `gorm:"column:created_at;size:32" json:"created_at"`
	ExtExpiresIn      *int     `gorm:"column:ext_expires_in;size:32" json:"ext_expires_in"`
	TokenType         *string  `gorm:"column:token_type" json:"token_type"`
	Scope             *string  `gorm:"column:scope" json:"scope"`
	IDToken           *string  `gorm:"column:id_token;type:text" json:"id_token"`
	SessionState      *string  `gorm:"column:session_state" json:"session_state"`
	Password          *string  `gorm:"column:password" json:"password"`
	User              *User    `gorm:"belongsTo"`
}

func (Account) TableName() string {
	return "Account"
}

func (m *Account) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Session struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	SessionToken string   `gorm:"column:sessionToken;uniqueIndex:Session_sessionToken_key;index:Session_sessionToken_idx;not null" json:"sessionToken"`
	UserID       int      `gorm:"column:userId;index:Session_userId_idx;size:32;not null" json:"userId"`
	IPAddress    *string  `gorm:"column:ipAddress" json:"ipAddress"`
	UserAgent    *string  `gorm:"column:userAgent" json:"userAgent"`
	ExpiresAt    DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt    DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	IsEmbed      bool     `gorm:"column:isEmbed;default:false;not null" json:"isEmbed"`
	User         *User    `gorm:"belongsTo"`
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

type AuthHandoff struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Type      string   `gorm:"column:type;not null" json:"type"`
	Payload   string   `gorm:"column:payload;not null" json:"payload"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ExpiresAt DateTime `gorm:"column:expiresAt;index:AuthHandoff_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
}

func (AuthHandoff) TableName() string {
	return "AuthHandoff"
}

func (m *AuthHandoff) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Folder struct {
	ID         string     `gorm:"column:id;primaryKey" json:"id"`
	Name       string     `gorm:"column:name;not null" json:"name"`
	UserID     int        `gorm:"column:userId;index:Folder_userId_idx;size:32;not null" json:"userId"`
	TeamID     int        `gorm:"column:teamId;index:Folder_teamId_idx;size:32;not null" json:"teamId"`
	Pinned     bool       `gorm:"column:pinned;default:false;not null" json:"pinned"`
	ParentID   *string    `gorm:"column:parentId;index:Folder_parentId_idx" json:"parentId"`
	CreatedAt  DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime   `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Visibility string     `gorm:"column:visibility;default:'EVERYONE';not null" json:"visibility"`
	Type       string     `gorm:"column:type;index:Folder_type_idx;not null" json:"type"`
	User       *User      `gorm:"belongsTo"`
	Team       *Team      `gorm:"belongsTo"`
	Parent     *Folder    `gorm:"foreignKey:ParentID;belongsTo"`
	Envelopes  []Envelope `gorm:"foreignKey:FolderID;constraint:Envelope_folderId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Subfolders []Folder   `gorm:"foreignKey:ParentID;constraint:Folder_parentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Folder) TableName() string {
	return "Folder"
}

func (m *Folder) BeforeCreate(tx *gorm.DB) error {
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

func (*Folder) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// @zod.import(["import { ZDocumentAuthOptionsSchema } from '@documenso/lib/types/document-auth';", "import { ZDocumentFormValuesSchema } from '@documenso/lib/types/document-form-values';", "import { ZSignatureLevelSchema } from '@documenso/lib/types/signature-level';"])
type Envelope struct {
	ID          string `gorm:"column:id;primaryKey" json:"id"`
	SecondaryID string `gorm:"column:secondaryId;uniqueIndex:Envelope_secondaryId_key;not null" json:"secondaryId"`
	// @zod.string.describe("A custom external ID you can use to identify the document.")
	ExternalID  *string   `gorm:"column:externalId" json:"externalId"`
	Type        string    `gorm:"column:type;index:Envelope_type_idx;not null" json:"type"`
	CreatedAt   DateTime  `gorm:"column:createdAt;index:Envelope_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	CompletedAt *DateTime `gorm:"column:completedAt;type:timestamp(3)" json:"completedAt"`
	DeletedAt   *DateTime `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	Status      string    `gorm:"column:status;index:Envelope_status_idx;default:'DRAFT';not null" json:"status"`
	Source      string    `gorm:"column:source;not null" json:"source"`
	// @zod.string.describe("The token for viewing the document using the QR code on the certificate.")
	QrToken *string `gorm:"column:qrToken" json:"qrToken"`
	// [SignatureLevel] @zod.custom.use(ZSignatureLevelSchema)
	SignatureLevel          string `gorm:"column:signatureLevel;not null" json:"signatureLevel"`
	InternalVersion         int    `gorm:"column:internalVersion;size:32;not null" json:"internalVersion"`
	UseLegacyFieldInsertion bool   `gorm:"column:useLegacyFieldInsertion;default:false;not null" json:"useLegacyFieldInsertion"`
	// [DocumentAuthOptions] @zod.custom.use(ZDocumentAuthOptionsSchema)
	AuthOptions datatypes.JSON `gorm:"column:authOptions" json:"authOptions"`
	// [DocumentFormValues] @zod.custom.use(ZDocumentFormValuesSchema)
	FormValues        datatypes.JSON `gorm:"column:formValues" json:"formValues"`
	Visibility        string         `gorm:"column:visibility;default:'EVERYONE';not null" json:"visibility"`
	TemplateType      string         `gorm:"column:templateType;default:'PRIVATE';not null" json:"templateType"`
	PublicTitle       string         `gorm:"column:publicTitle;default:'';not null" json:"publicTitle"`
	PublicDescription string         `gorm:"column:publicDescription;default:'';not null" json:"publicDescription"`
	TemplateID        *int           `gorm:"column:templateId;size:32" json:"templateId"`
	// @zod.number.describe("The ID of the user that created this document.")
	UserID              int                  `gorm:"column:userId;index:Envelope_userId_idx;size:32;not null" json:"userId"`
	TeamID              int                  `gorm:"column:teamId;index:Envelope_teamId_idx;size:32;not null" json:"teamId"`
	FolderID            *string              `gorm:"column:folderId;index:Envelope_folderId_idx" json:"folderId"`
	DocumentMetaID      string               `gorm:"column:documentMetaId;uniqueIndex:Envelope_documentMetaId_key;not null" json:"documentMetaId"`
	User                *User                `gorm:"belongsTo"`
	Team                *Team                `gorm:"belongsTo"`
	Folder              *Folder              `gorm:"belongsTo"`
	DocumentMeta        *DocumentMeta        `gorm:"belongsTo"`
	EnvelopeItems       []EnvelopeItem       `gorm:"foreignKey:EnvelopeID;constraint:EnvelopeItem_envelopeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Recipients          []Recipient          `gorm:"foreignKey:EnvelopeID;constraint:Recipient_envelopeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Fields              []Field              `gorm:"foreignKey:EnvelopeID;constraint:Field_envelopeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Contents            []EnvelopeContent    `gorm:"foreignKey:EnvelopeID;constraint:EnvelopeContent_envelopeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShareLinks          []DocumentShareLink  `gorm:"foreignKey:EnvelopeID;constraint:DocumentShareLink_envelopeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuditLogs           []DocumentAuditLog   `gorm:"foreignKey:EnvelopeID;constraint:DocumentAuditLog_envelopeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	EnvelopeAttachments []EnvelopeAttachment `gorm:"foreignKey:EnvelopeID;constraint:EnvelopeAttachment_envelopeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DirectLink          *TemplateDirectLink  `gorm:"foreignKey:EnvelopeID;constraint:TemplateDirectLink_envelopeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Envelope) TableName() string {
	return "Envelope"
}

func (m *Envelope) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Envelope) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EnvelopeItem struct {
	ID             string            `gorm:"column:id;primaryKey" json:"id"`
	Title          string            `gorm:"column:title;not null" json:"title"`
	Order          int               `gorm:"column:order;size:32;not null" json:"order"`
	DocumentDataID string            `gorm:"column:documentDataId;uniqueIndex:EnvelopeItem_documentDataId_key;not null" json:"documentDataId"`
	EnvelopeID     string            `gorm:"column:envelopeId;index:EnvelopeItem_envelopeId_idx;not null" json:"envelopeId"`
	DocumentData   *DocumentData     `gorm:"belongsTo"`
	Envelope       *Envelope         `gorm:"belongsTo"`
	Field          []Field           `gorm:"foreignKey:EnvelopeItemID;constraint:Field_envelopeItemId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Contents       []EnvelopeContent `gorm:"foreignKey:EnvelopeItemID;constraint:EnvelopeContent_envelopeItemId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (EnvelopeItem) TableName() string {
	return "EnvelopeItem"
}

type DocumentAuditLog struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	EnvelopeID *string        `gorm:"column:envelopeId;index:DocumentAuditLog_envelopeId_idx" json:"envelopeId"`
	CreatedAt  DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Type       string         `gorm:"column:type;not null" json:"type"`
	Data       datatypes.JSON `gorm:"column:data;not null" json:"data"`
	Name       *string        `gorm:"column:name" json:"name"`
	Email      *string        `gorm:"column:email" json:"email"`
	UserID     *int           `gorm:"column:userId;size:32" json:"userId"`
	UserAgent  *string        `gorm:"column:userAgent" json:"userAgent"`
	IPAddress  *string        `gorm:"column:ipAddress" json:"ipAddress"`
	Envelope   *Envelope      `gorm:"belongsTo"`
}

func (DocumentAuditLog) TableName() string {
	return "DocumentAuditLog"
}

func (m *DocumentAuditLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DocumentData struct {
	ID           string        `gorm:"column:id;primaryKey" json:"id"`
	Type         string        `gorm:"column:type;not null" json:"type"`
	Data         string        `gorm:"column:data;not null" json:"data"`
	InitialData  string        `gorm:"column:initialData;not null" json:"initialData"`
	EnvelopeItem *EnvelopeItem `gorm:"foreignKey:DocumentDataID;constraint:EnvelopeItem_documentDataId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (DocumentData) TableName() string {
	return "DocumentData"
}

func (m *DocumentData) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';"])
type DocumentMeta struct {
	ID                     string  `gorm:"column:id;primaryKey" json:"id"`
	Subject                *string `gorm:"column:subject" json:"subject"`
	Message                *string `gorm:"column:message" json:"message"`
	Timezone               *string `gorm:"column:timezone;type:text;default:'Etc/UTC'" json:"timezone"`
	DateFormat             *string `gorm:"column:dateFormat;type:text;default:'yyyy-MM-dd hh:mm a'" json:"dateFormat"`
	RedirectURL            *string `gorm:"column:redirectUrl" json:"redirectUrl"`
	SigningOrder           string  `gorm:"column:signingOrder;default:'PARALLEL';not null" json:"signingOrder"`
	AllowDictateNextSigner bool    `gorm:"column:allowDictateNextSigner;default:false;not null" json:"allowDictateNextSigner"`
	TypedSignatureEnabled  *bool   `gorm:"column:typedSignatureEnabled;default:true;not null" json:"typedSignatureEnabled"`
	UploadSignatureEnabled *bool   `gorm:"column:uploadSignatureEnabled;default:true;not null" json:"uploadSignatureEnabled"`
	DrawSignatureEnabled   *bool   `gorm:"column:drawSignatureEnabled;default:true;not null" json:"drawSignatureEnabled"`
	Language               *string `gorm:"column:language;default:'en';not null" json:"language"`
	DistributionMethod     string  `gorm:"column:distributionMethod;default:'EMAIL';not null" json:"distributionMethod"`
	// [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema)
	EmailSettings datatypes.JSON `gorm:"column:emailSettings" json:"emailSettings"`
	EmailReplyTo  *string        `gorm:"column:emailReplyTo" json:"emailReplyTo"`
	EmailID       *string        `gorm:"column:emailId" json:"emailId"`
	// [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
	EnvelopeExpirationPeriod datatypes.JSON `gorm:"column:envelopeExpirationPeriod" json:"envelopeExpirationPeriod"`
	// [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
	ReminderSettings datatypes.JSON `gorm:"column:reminderSettings" json:"reminderSettings"`
	Envelope         *Envelope      `gorm:"foreignKey:DocumentMetaID;constraint:Envelope_documentMetaId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (DocumentMeta) TableName() string {
	return "DocumentMeta"
}

func (m *DocumentMeta) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// @zod.import(["import { ZEnvelopeAttachmentTypeSchema } from '@documenso/lib/types/envelope-attachment';"])
type EnvelopeAttachment struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// [EnvelopeAttachmentType] @zod.custom.use(ZEnvelopeAttachmentTypeSchema)
	Type       string    `gorm:"column:type;not null" json:"type"`
	Label      string    `gorm:"column:label;not null" json:"label"`
	Data       string    `gorm:"column:data;not null" json:"data"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	EnvelopeID string    `gorm:"column:envelopeId;index:EnvelopeAttachment_envelopeId_idx;not null" json:"envelopeId"`
	Envelope   *Envelope `gorm:"belongsTo"`
}

func (EnvelopeAttachment) TableName() string {
	return "EnvelopeAttachment"
}

func (m *EnvelopeAttachment) BeforeCreate(tx *gorm.DB) error {
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

func (*EnvelopeAttachment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// @zod.import(["import { ZRecipientAuthOptionsSchema } from '@documenso/lib/types/document-auth';"])
type Recipient struct {
	ID                   int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	EnvelopeID           string    `gorm:"column:envelopeId;index:Recipient_envelopeId_idx;index:Recipient_email_documentDeletedAt_envelopeId_idx;index:Recipient_email_envelopeId_idx;index:Recipient_email_signingStatus_envelopeId_role_idx;not null" json:"envelopeId"`
	Email                string    `gorm:"column:email;index:Recipient_email_idx;index:Recipient_email_documentDeletedAt_envelopeId_idx;index:Recipient_email_envelopeId_idx;index:Recipient_email_signingStatus_envelopeId_role_idx;index:Recipient_email_trgm_idx;type:varchar(255);not null" json:"email"`
	Name                 string    `gorm:"column:name;index:Recipient_name_trgm_idx;type:varchar(255);default:'';not null" json:"name"`
	Token                string    `gorm:"column:token;index:Recipient_token_idx;not null" json:"token"`
	DocumentDeletedAt    *DateTime `gorm:"column:documentDeletedAt;index:Recipient_email_documentDeletedAt_envelopeId_idx;type:timestamp(3)" json:"documentDeletedAt"`
	Expired              *DateTime `gorm:"column:expired;type:timestamp(3)" json:"expired"`
	ExpiresAt            *DateTime `gorm:"column:expiresAt;index:Recipient_expiresAt_idx;type:timestamp(3)" json:"expiresAt"`
	ExpirationNotifiedAt *DateTime `gorm:"column:expirationNotifiedAt;type:timestamp(3)" json:"expirationNotifiedAt"`
	SentAt               *DateTime `gorm:"column:sentAt;type:timestamp(3)" json:"sentAt"`
	SignedAt             *DateTime `gorm:"column:signedAt;index:Recipient_signedAt_idx;type:timestamp(3)" json:"signedAt"`
	LastReminderSentAt   *DateTime `gorm:"column:lastReminderSentAt;type:timestamp(3)" json:"lastReminderSentAt"`
	NextReminderAt       *DateTime `gorm:"column:nextReminderAt;index:Recipient_nextReminderAt_idx;type:timestamp(3)" json:"nextReminderAt"`
	ReminderCount        int       `gorm:"column:reminderCount;size:32;default:0;not null" json:"reminderCount"`
	// [RecipientAuthOptions] @zod.custom.use(ZRecipientAuthOptionsSchema)
	AuthOptions datatypes.JSON `gorm:"column:authOptions" json:"authOptions"`
	// @zod.number.describe("The order in which the recipient should sign the document. Only works if the document is set to sequential signing.")
	SigningOrder    *int           `gorm:"column:signingOrder;size:32" json:"signingOrder"`
	RejectionReason *string        `gorm:"column:rejectionReason" json:"rejectionReason"`
	Role            string         `gorm:"column:role;index:Recipient_email_signingStatus_envelopeId_role_idx;default:'SIGNER';not null" json:"role"`
	ReadStatus      string         `gorm:"column:readStatus;default:'NOT_OPENED';not null" json:"readStatus"`
	SigningStatus   string         `gorm:"column:signingStatus;index:Recipient_email_signingStatus_envelopeId_role_idx;default:'NOT_SIGNED';not null" json:"signingStatus"`
	SendStatus      string         `gorm:"column:sendStatus;default:'NOT_SENT';not null" json:"sendStatus"`
	Envelope        *Envelope      `gorm:"belongsTo"`
	Fields          []Field        `gorm:"foreignKey:RecipientID;constraint:Field_recipientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Signatures      []Signature    `gorm:"foreignKey:RecipientID;constraint:Signature_recipientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CscCredential   *CscCredential `gorm:"foreignKey:RecipientID;constraint:CscCredential_recipientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CscSession      *CscSession    `gorm:"foreignKey:RecipientID;constraint:CscSession_recipientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Recipient) TableName() string {
	return "Recipient"
}

// @zod.import(["import { ZFieldMetaNotOptionalSchema } from '@documenso/lib/types/field-meta';"])
type Field struct {
	ID             int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	SecondaryID    string `gorm:"column:secondaryId;uniqueIndex:Field_secondaryId_key;not null" json:"secondaryId"`
	EnvelopeID     string `gorm:"column:envelopeId;index:Field_envelopeId_idx;not null" json:"envelopeId"`
	EnvelopeItemID string `gorm:"column:envelopeItemId;index:Field_envelopeItemId_idx;not null" json:"envelopeItemId"`
	RecipientID    int    `gorm:"column:recipientId;index:Field_recipientId_idx;size:32;not null" json:"recipientId"`
	Type           string `gorm:"column:type;not null" json:"type"`
	// @zod.number.describe("The page number of the field on the document. Starts from 1.")
	Page       int      `gorm:"column:page;size:32;not null" json:"page"`
	PositionX  float64  `gorm:"column:positionX;type:decimal(65,30);default:0;not null" json:"positionX"`
	PositionY  float64  `gorm:"column:positionY;type:decimal(65,30);default:0;not null" json:"positionY"`
	Width      *float64 `gorm:"column:width;type:decimal(65,30);default:-1;not null" json:"width"`
	Height     *float64 `gorm:"column:height;type:decimal(65,30);default:-1;not null" json:"height"`
	CustomText string   `gorm:"column:customText;not null" json:"customText"`
	Inserted   bool     `gorm:"column:inserted;not null" json:"inserted"`
	// [FieldMeta] @zod.custom.use(ZFieldMetaNotOptionalSchema)
	FieldMeta    datatypes.JSON `gorm:"column:fieldMeta" json:"fieldMeta"`
	EnvelopeItem *EnvelopeItem  `gorm:"belongsTo"`
	Envelope     *Envelope      `gorm:"belongsTo"`
	Recipient    *Recipient     `gorm:"belongsTo"`
	Signature    *Signature     `gorm:"foreignKey:FieldID;constraint:Signature_fieldId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Field) TableName() string {
	return "Field"
}

func (m *Field) BeforeCreate(_ *gorm.DB) error {
	if m.SecondaryID == "" {
		m.SecondaryID = cuid.New()
	}
	return nil
}

// @zod.import(["import { ZEnvelopeContentMetaSchema } from '@documenso/lib/types/envelope-content-meta';"])
type EnvelopeContent struct {
	ID             string `gorm:"column:id;primaryKey" json:"id"`
	EnvelopeID     string `gorm:"column:envelopeId;index:EnvelopeContent_envelopeId_idx;not null" json:"envelopeId"`
	EnvelopeItemID string `gorm:"column:envelopeItemId;index:EnvelopeContent_envelopeItemId_idx;not null" json:"envelopeItemId"`
	// [EnvelopeContentMeta] @zod.custom.use(ZEnvelopeContentMetaSchema)
	ContentMeta datatypes.JSON `gorm:"column:contentMeta;not null" json:"contentMeta"`
	// Any number of contents may point at the same data content, e.g. a
	// document created from a template shares the template's images.
	DataContentID *string       `gorm:"column:dataContentId" json:"dataContentId"`
	EnvelopeItem  *EnvelopeItem `gorm:"belongsTo"`
	Envelope      *Envelope     `gorm:"belongsTo"`
	DataContent   *DataContent  `gorm:"belongsTo"`
}

func (EnvelopeContent) TableName() string {
	return "EnvelopeContent"
}

// Generic stored data (e.g. an uploaded image) which other models reference.
// The kind of data and its kind specific details live in `metadata`.
// @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
type DataContent struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// How `data` is stored, e.g. an S3 path or inline bytes.
	Type string `gorm:"column:type;not null" json:"type"`
	Data string `gorm:"column:data;not null" json:"data"`
	// [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema)
	Metadata         datatypes.JSON    `gorm:"column:metadata;not null" json:"metadata"`
	EnvelopeContents []EnvelopeContent `gorm:"foreignKey:DataContentID;constraint:EnvelopeContent_dataContentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (DataContent) TableName() string {
	return "DataContent"
}

type Signature struct {
	ID                     int        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Created                DateTime   `gorm:"column:created;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created"`
	RecipientID            int        `gorm:"column:recipientId;index:Signature_recipientId_idx;size:32;not null" json:"recipientId"`
	FieldID                int        `gorm:"column:fieldId;uniqueIndex:Signature_fieldId_key;size:32;not null" json:"fieldId"`
	SignatureImageAsBase64 *string    `gorm:"column:signatureImageAsBase64" json:"signatureImageAsBase64"`
	TypedSignature         *string    `gorm:"column:typedSignature" json:"typedSignature"`
	Recipient              *Recipient `gorm:"belongsTo"`
	Field                  *Field     `gorm:"belongsTo"`
}

func (Signature) TableName() string {
	return "Signature"
}

func (m *Signature) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Created.IsZero() {
		m.Created = DateTime{Time: now}
	}
	return nil
}

// Per-recipient cached credential metadata for the configured Cloud Signature
// Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
// symmetric-encrypted service-scope access token. One row per recipient on a
// CSC envelope; absent for SES recipients.
type CscCredential struct {
	ID           string `gorm:"column:id;primaryKey" json:"id"`
	ProviderID   string `gorm:"column:providerId;not null" json:"providerId"`
	CredentialID string `gorm:"column:credentialId;not null" json:"credentialId"`
	// TSP-validated certificate chain (length-prefixed). Nullable until the
	// service-scope OAuth callback successfully retrieves and validates it.
	CertCache []byte `gorm:"column:certCache" json:"certCache"`
	// Algorithm metadata derived from the CSC `credentials/info` response.
	// Held as flat scalars so each is independently queryable.
	SignatureAlgorithm string  `gorm:"column:signatureAlgorithm;not null" json:"signatureAlgorithm"`
	KeyType            string  `gorm:"column:keyType;not null" json:"keyType"`
	DigestAlgorithm    string  `gorm:"column:digestAlgorithm;not null" json:"digestAlgorithm"`
	KeyLenBits         *int    `gorm:"column:keyLenBits;size:32" json:"keyLenBits"`
	SignAlgoParams     *string `gorm:"column:signAlgoParams" json:"signAlgoParams"`
	// Service-scope access token, symmetric-encrypted at the application layer
	// before persistence. Nullable until the service-scope OAuth callback runs.
	ServiceTokenCiphertext []byte     `gorm:"column:serviceTokenCiphertext" json:"serviceTokenCiphertext"`
	ServiceTokenExpiresAt  *DateTime  `gorm:"column:serviceTokenExpiresAt;type:timestamp(3)" json:"serviceTokenExpiresAt"`
	RecipientID            int        `gorm:"column:recipientId;uniqueIndex:CscCredential_recipientId_key;size:32;not null" json:"recipientId"`
	CreatedAt              DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt              DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Recipient              *Recipient `gorm:"belongsTo"`
}

func (CscCredential) TableName() string {
	return "CscCredential"
}

func (m *CscCredential) BeforeCreate(tx *gorm.DB) error {
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

func (*CscCredential) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Per-recipient transient sign-time state held across the credential-scope
// OAuth round-trip. The `recipientId @unique` constraint caps each recipient
// to at most one in-flight session — re-clicking Sign UPSERTs the row.
// Cleaned up transitively when the Recipient is deleted.
// @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
type CscSession struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
	// flows through the Recipient cascade.
	EnvelopeID string `gorm:"column:envelopeId;not null" json:"envelopeId"`
	// Pinned at prep time; stable across overlay re-renders so the captured
	// signedAttrs digest matches the bytes the TSP signs.
	SigningTime DateTime `gorm:"column:signingTime;type:timestamp(3);not null" json:"signingTime"`
	// Contract between prep and sign: an array of
	// `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
	// `documentDataId` pins the immutable bytes that produced the captured
	// digest so sign-time can re-derive the same hash against the same content.
	// [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
	ItemsJSON datatypes.JSON `gorm:"column:itemsJson;not null" json:"itemsJson"`
	// Populated by the credential-scope OAuth callback. Both nullable at row
	// creation; set together when the SAD is exchanged.
	EncryptedSad []byte     `gorm:"column:encryptedSad" json:"encryptedSad"`
	SadExpiresAt *DateTime  `gorm:"column:sadExpiresAt;type:timestamp(3)" json:"sadExpiresAt"`
	CreatedAt    DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	RecipientID  int        `gorm:"column:recipientId;uniqueIndex:CscSession_recipientId_key;size:32;not null" json:"recipientId"`
	Recipient    *Recipient `gorm:"belongsTo"`
}

func (CscSession) TableName() string {
	return "CscSession"
}

func (m *CscSession) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DocumentShareLink struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Email      string    `gorm:"column:email;uniqueIndex:DocumentShareLink_envelopeId_email_key;not null" json:"email"`
	Slug       string    `gorm:"column:slug;uniqueIndex:DocumentShareLink_slug_key;not null" json:"slug"`
	EnvelopeID string    `gorm:"column:envelopeId;uniqueIndex:DocumentShareLink_envelopeId_email_key;not null" json:"envelopeId"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Envelope   *Envelope `gorm:"belongsTo"`
}

func (DocumentShareLink) TableName() string {
	return "DocumentShareLink"
}

func (m *DocumentShareLink) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DocumentShareLink) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Organisation struct {
	ID                                 string                            `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                          DateTime                          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                          DateTime                          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Type                               string                            `gorm:"column:type;not null" json:"type"`
	Name                               string                            `gorm:"column:name;index:Organisation_name_idx;not null" json:"name"`
	URL                                string                            `gorm:"column:url;uniqueIndex:Organisation_url_key;not null" json:"url"`
	AvatarImageID                      *string                           `gorm:"column:avatarImageId" json:"avatarImageId"`
	CustomerID                         *string                           `gorm:"column:customerId;uniqueIndex:Organisation_customerId_key" json:"customerId"`
	OrganisationClaimID                string                            `gorm:"column:organisationClaimId;uniqueIndex:Organisation_organisationClaimId_key;not null" json:"organisationClaimId"`
	OwnerUserID                        int                               `gorm:"column:ownerUserId;index:Organisation_ownerUserId_idx;size:32;not null" json:"ownerUserId"`
	OrganisationGlobalSettingsID       string                            `gorm:"column:organisationGlobalSettingsId;uniqueIndex:Organisation_organisationGlobalSettingsId_key;not null" json:"organisationGlobalSettingsId"`
	OrganisationAuthenticationPortalID string                            `gorm:"column:organisationAuthenticationPortalId;uniqueIndex:Organisation_organisationAuthenticationPortalId_key;not null" json:"organisationAuthenticationPortalId"`
	OrganisationClaim                  *OrganisationClaim                `gorm:"belongsTo"`
	AvatarImage                        *AvatarImage                      `gorm:"belongsTo"`
	Owner                              *User                             `gorm:"foreignKey:OwnerUserID;belongsTo"`
	OrganisationGlobalSettings         *OrganisationGlobalSettings       `gorm:"belongsTo"`
	OrganisationAuthenticationPortal   *OrganisationAuthenticationPortal `gorm:"belongsTo"`
	Members                            []OrganisationMember              `gorm:"foreignKey:OrganisationID;constraint:OrganisationMember_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites                            []OrganisationMemberInvite        `gorm:"foreignKey:OrganisationID;constraint:OrganisationMemberInvite_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Groups                             []OrganisationGroup               `gorm:"foreignKey:OrganisationID;constraint:OrganisationGroup_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Teams                              []Team                            `gorm:"foreignKey:OrganisationID;constraint:Team_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmailDomains                       []EmailDomain                     `gorm:"foreignKey:OrganisationID;constraint:EmailDomain_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganisationEmails                 []OrganisationEmail               `gorm:"foreignKey:OrganisationID;constraint:OrganisationEmail_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	MonthlyStats                       []OrganisationMonthlyStat         `gorm:"foreignKey:OrganisationID;constraint:OrganisationMonthlyStat_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmbedSignatures                    []EmbedSignature                  `gorm:"foreignKey:OrganisationID;constraint:EmbedSignature_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subscription                       *Subscription                     `gorm:"foreignKey:OrganisationID;constraint:Subscription_organisationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Organisation) TableName() string {
	return "Organisation"
}

func (m *Organisation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Organisation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganisationMember struct {
	ID                       string                    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime                  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID                   int                       `gorm:"column:userId;uniqueIndex:OrganisationMember_userId_organisationId_key;size:32;not null" json:"userId"`
	OrganisationID           string                    `gorm:"column:organisationId;uniqueIndex:OrganisationMember_userId_organisationId_key;index:OrganisationMember_organisationId_idx;not null" json:"organisationId"`
	User                     *User                     `gorm:"belongsTo"`
	Organisation             *Organisation             `gorm:"belongsTo"`
	OrganisationGroupMembers []OrganisationGroupMember `gorm:"foreignKey:OrganisationMemberID;constraint:OrganisationGroupMember_organisationMemberId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OrganisationMember) TableName() string {
	return "OrganisationMember"
}

func (m *OrganisationMember) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganisationMember) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganisationMemberInvite struct {
	ID               string        `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt        DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Email            string        `gorm:"column:email;not null" json:"email"`
	Token            string        `gorm:"column:token;uniqueIndex:OrganisationMemberInvite_token_key;not null" json:"token"`
	Status           string        `gorm:"column:status;default:'PENDING';not null" json:"status"`
	OrganisationID   string        `gorm:"column:organisationId;not null" json:"organisationId"`
	OrganisationRole string        `gorm:"column:organisationRole;not null" json:"organisationRole"`
	Organisation     *Organisation `gorm:"belongsTo"`
}

func (OrganisationMemberInvite) TableName() string {
	return "OrganisationMemberInvite"
}

func (m *OrganisationMemberInvite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type OrganisationGroup struct {
	ID                       string                    `gorm:"column:id;primaryKey" json:"id"`
	Name                     *string                   `gorm:"column:name" json:"name"`
	Type                     string                    `gorm:"column:type;not null" json:"type"`
	OrganisationRole         string                    `gorm:"column:organisationRole;not null" json:"organisationRole"`
	OrganisationID           string                    `gorm:"column:organisationId;index:OrganisationGroup_organisationId_idx;not null" json:"organisationId"`
	Organisation             *Organisation             `gorm:"belongsTo"`
	OrganisationGroupMembers []OrganisationGroupMember `gorm:"foreignKey:GroupID;constraint:OrganisationGroupMember_groupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TeamGroups               []TeamGroup               `gorm:"foreignKey:OrganisationGroupID;constraint:TeamGroup_organisationGroupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OrganisationGroup) TableName() string {
	return "OrganisationGroup"
}

type OrganisationGroupMember struct {
	ID                   string              `gorm:"column:id;primaryKey" json:"id"`
	GroupID              string              `gorm:"column:groupId;uniqueIndex:OrganisationGroupMember_organisationMemberId_groupId_key;index:OrganisationGroupMember_groupId_idx;not null" json:"groupId"`
	OrganisationMemberID string              `gorm:"column:organisationMemberId;uniqueIndex:OrganisationGroupMember_organisationMemberId_groupId_key;index:OrganisationGroupMember_organisationMemberId_idx;not null" json:"organisationMemberId"`
	Group                *OrganisationGroup  `gorm:"foreignKey:GroupID;belongsTo"`
	OrganisationMember   *OrganisationMember `gorm:"belongsTo"`
}

func (OrganisationGroupMember) TableName() string {
	return "OrganisationGroupMember"
}

type TeamGroup struct {
	ID                  string             `gorm:"column:id;primaryKey" json:"id"`
	OrganisationGroupID string             `gorm:"column:organisationGroupId;uniqueIndex:TeamGroup_teamId_organisationGroupId_key;index:TeamGroup_organisationGroupId_idx;not null" json:"organisationGroupId"`
	TeamRole            string             `gorm:"column:teamRole;not null" json:"teamRole"`
	TeamID              int                `gorm:"column:teamId;uniqueIndex:TeamGroup_teamId_organisationGroupId_key;index:TeamGroup_teamId_idx;size:32;not null" json:"teamId"`
	OrganisationGroup   *OrganisationGroup `gorm:"belongsTo"`
	Team                *Team              `gorm:"belongsTo"`
}

func (TeamGroup) TableName() string {
	return "TeamGroup"
}

// @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
type OrganisationGlobalSettings struct {
	ID                        string  `gorm:"column:id;primaryKey" json:"id"`
	DocumentVisibility        string  `gorm:"column:documentVisibility;default:'EVERYONE';not null" json:"documentVisibility"`
	DocumentLanguage          *string `gorm:"column:documentLanguage;default:'en';not null" json:"documentLanguage"`
	IncludeSenderDetails      *bool   `gorm:"column:includeSenderDetails;default:true;not null" json:"includeSenderDetails"`
	IncludeSigningCertificate *bool   `gorm:"column:includeSigningCertificate;default:true;not null" json:"includeSigningCertificate"`
	IncludeAuditLog           bool    `gorm:"column:includeAuditLog;default:false;not null" json:"includeAuditLog"`
	DocumentTimezone          *string `gorm:"column:documentTimezone" json:"documentTimezone"`
	DocumentDateFormat        *string `gorm:"column:documentDateFormat;default:'yyyy-MM-dd hh:mm a';not null" json:"documentDateFormat"`
	DelegateDocumentOwnership bool    `gorm:"column:delegateDocumentOwnership;default:false;not null" json:"delegateDocumentOwnership"`
	TypedSignatureEnabled     *bool   `gorm:"column:typedSignatureEnabled;default:true;not null" json:"typedSignatureEnabled"`
	UploadSignatureEnabled    *bool   `gorm:"column:uploadSignatureEnabled;default:true;not null" json:"uploadSignatureEnabled"`
	DrawSignatureEnabled      *bool   `gorm:"column:drawSignatureEnabled;default:true;not null" json:"drawSignatureEnabled"`
	// [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
	DefaultRecipients datatypes.JSON `gorm:"column:defaultRecipients" json:"defaultRecipients"`
	EmailID           *string        `gorm:"column:emailId" json:"emailId"`
	EmailReplyTo      *string        `gorm:"column:emailReplyTo" json:"emailReplyTo"`
	// [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema)
	EmailDocumentSettings  datatypes.JSON `gorm:"column:emailDocumentSettings;not null" json:"emailDocumentSettings"`
	BrandingEnabled        bool           `gorm:"column:brandingEnabled;default:false;not null" json:"brandingEnabled"`
	BrandingLogo           string         `gorm:"column:brandingLogo;default:'';not null" json:"brandingLogo"`
	BrandingURL            string         `gorm:"column:brandingUrl;default:'';not null" json:"brandingUrl"`
	BrandingCompanyDetails string         `gorm:"column:brandingCompanyDetails;default:'';not null" json:"brandingCompanyDetails"`
	// [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
	BrandingColors datatypes.JSON `gorm:"column:brandingColors" json:"brandingColors"`
	BrandingCSS    string         `gorm:"column:brandingCss;default:'';not null" json:"brandingCss"`
	// [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
	EnvelopeExpirationPeriod datatypes.JSON `gorm:"column:envelopeExpirationPeriod" json:"envelopeExpirationPeriod"`
	// [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
	ReminderSettings  datatypes.JSON     `gorm:"column:reminderSettings" json:"reminderSettings"`
	AiFeaturesEnabled bool               `gorm:"column:aiFeaturesEnabled;default:false;not null" json:"aiFeaturesEnabled"`
	Email             *OrganisationEmail `gorm:"foreignKey:EmailID;belongsTo"`
	Organisation      *Organisation      `gorm:"foreignKey:OrganisationGlobalSettingsID;constraint:Organisation_organisationGlobalSettingsId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (OrganisationGlobalSettings) TableName() string {
	return "OrganisationGlobalSettings"
}

// @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
type TeamGlobalSettings struct {
	ID                        string  `gorm:"column:id;primaryKey" json:"id"`
	DocumentVisibility        *string `gorm:"column:documentVisibility" json:"documentVisibility"`
	DocumentLanguage          *string `gorm:"column:documentLanguage" json:"documentLanguage"`
	DocumentTimezone          *string `gorm:"column:documentTimezone" json:"documentTimezone"`
	DocumentDateFormat        *string `gorm:"column:documentDateFormat" json:"documentDateFormat"`
	DelegateDocumentOwnership *bool   `gorm:"column:delegateDocumentOwnership" json:"delegateDocumentOwnership"`
	IncludeSenderDetails      *bool   `gorm:"column:includeSenderDetails" json:"includeSenderDetails"`
	IncludeSigningCertificate *bool   `gorm:"column:includeSigningCertificate" json:"includeSigningCertificate"`
	IncludeAuditLog           *bool   `gorm:"column:includeAuditLog" json:"includeAuditLog"`
	TypedSignatureEnabled     *bool   `gorm:"column:typedSignatureEnabled" json:"typedSignatureEnabled"`
	UploadSignatureEnabled    *bool   `gorm:"column:uploadSignatureEnabled" json:"uploadSignatureEnabled"`
	DrawSignatureEnabled      *bool   `gorm:"column:drawSignatureEnabled" json:"drawSignatureEnabled"`
	// [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
	DefaultRecipients datatypes.JSON `gorm:"column:defaultRecipients" json:"defaultRecipients"`
	EmailID           *string        `gorm:"column:emailId" json:"emailId"`
	EmailReplyTo      *string        `gorm:"column:emailReplyTo" json:"emailReplyTo"`
	// [DocumentEmailSettingsNullable] @zod.custom.use(ZDocumentEmailSettingsSchema)
	EmailDocumentSettings  datatypes.JSON `gorm:"column:emailDocumentSettings" json:"emailDocumentSettings"`
	BrandingEnabled        *bool          `gorm:"column:brandingEnabled" json:"brandingEnabled"`
	BrandingLogo           *string        `gorm:"column:brandingLogo" json:"brandingLogo"`
	BrandingURL            *string        `gorm:"column:brandingUrl" json:"brandingUrl"`
	BrandingCompanyDetails *string        `gorm:"column:brandingCompanyDetails" json:"brandingCompanyDetails"`
	// [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
	BrandingColors datatypes.JSON `gorm:"column:brandingColors" json:"brandingColors"`
	BrandingCSS    *string        `gorm:"column:brandingCss" json:"brandingCss"`
	// [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
	EnvelopeExpirationPeriod datatypes.JSON `gorm:"column:envelopeExpirationPeriod" json:"envelopeExpirationPeriod"`
	// [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
	ReminderSettings  datatypes.JSON     `gorm:"column:reminderSettings" json:"reminderSettings"`
	AiFeaturesEnabled *bool              `gorm:"column:aiFeaturesEnabled" json:"aiFeaturesEnabled"`
	Email             *OrganisationEmail `gorm:"foreignKey:EmailID;belongsTo"`
	Team              *Team              `gorm:"foreignKey:TeamGlobalSettingsID;constraint:Team_teamGlobalSettingsId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TeamGlobalSettings) TableName() string {
	return "TeamGlobalSettings"
}

type Team struct {
	ID                   int                    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name                 string                 `gorm:"column:name;index:Team_name_idx;not null" json:"name"`
	URL                  string                 `gorm:"column:url;uniqueIndex:Team_url_key;not null" json:"url"`
	CreatedAt            DateTime               `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	AvatarImageID        *string                `gorm:"column:avatarImageId" json:"avatarImageId"`
	OrganisationID       string                 `gorm:"column:organisationId;index:Team_organisationId_idx;not null" json:"organisationId"`
	TeamGlobalSettingsID string                 `gorm:"column:teamGlobalSettingsId;uniqueIndex:Team_teamGlobalSettingsId_key;not null" json:"teamGlobalSettingsId"`
	AvatarImage          *AvatarImage           `gorm:"belongsTo"`
	Organisation         *Organisation          `gorm:"belongsTo"`
	TeamGlobalSettings   *TeamGlobalSettings    `gorm:"belongsTo"`
	Envelopes            []Envelope             `gorm:"foreignKey:TeamID;constraint:Envelope_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Folders              []Folder               `gorm:"foreignKey:TeamID;constraint:Folder_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APITokens            []APIToken             `gorm:"foreignKey:TeamID;constraint:ApiToken_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks             []Webhook              `gorm:"foreignKey:TeamID;constraint:Webhook_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TeamGroups           []TeamGroup            `gorm:"foreignKey:TeamID;constraint:TeamGroup_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TeamEmail            *TeamEmail             `gorm:"foreignKey:TeamID;constraint:TeamEmail_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmailVerification    *TeamEmailVerification `gorm:"foreignKey:TeamID;constraint:TeamEmailVerification_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Profile              *TeamProfile           `gorm:"foreignKey:TeamID;constraint:TeamProfile_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Team) TableName() string {
	return "Team"
}

func (m *Team) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type TeamEmail struct {
	TeamID    int      `gorm:"column:teamId;primaryKey;size:32" json:"teamId"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Name      string   `gorm:"column:name;not null" json:"name"`
	Email     string   `gorm:"column:email;uniqueIndex:TeamEmail_email_key;not null" json:"email"`
	Team      *Team    `gorm:"belongsTo"`
}

func (TeamEmail) TableName() string {
	return "TeamEmail"
}

func (m *TeamEmail) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type TeamEmailVerification struct {
	TeamID    int      `gorm:"column:teamId;primaryKey;size:32" json:"teamId"`
	Name      string   `gorm:"column:name;not null" json:"name"`
	Email     string   `gorm:"column:email;not null" json:"email"`
	Token     string   `gorm:"column:token;uniqueIndex:TeamEmailVerification_token_key;not null" json:"token"`
	Completed bool     `gorm:"column:completed;default:false;not null" json:"completed"`
	ExpiresAt DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Team      *Team    `gorm:"belongsTo"`
}

func (TeamEmailVerification) TableName() string {
	return "TeamEmailVerification"
}

func (m *TeamEmailVerification) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type TemplateDirectLink struct {
	ID                        string    `gorm:"column:id;primaryKey" json:"id"`
	EnvelopeID                string    `gorm:"column:envelopeId;uniqueIndex:TemplateDirectLink_envelopeId_key;not null" json:"envelopeId"`
	Token                     string    `gorm:"column:token;uniqueIndex:TemplateDirectLink_token_key;not null" json:"token"`
	CreatedAt                 DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Enabled                   bool      `gorm:"column:enabled;not null" json:"enabled"`
	DirectTemplateRecipientID int       `gorm:"column:directTemplateRecipientId;size:32;not null" json:"directTemplateRecipientId"`
	Envelope                  *Envelope `gorm:"belongsTo"`
}

func (TemplateDirectLink) TableName() string {
	return "TemplateDirectLink"
}

func (m *TemplateDirectLink) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SiteSettings struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	Enabled              bool           `gorm:"column:enabled;default:false;not null" json:"enabled"`
	Data                 datatypes.JSON `gorm:"column:data;not null" json:"data"`
	LastModifiedByUserID *int           `gorm:"column:lastModifiedByUserId;size:32" json:"lastModifiedByUserId"`
	LastModifiedAt       DateTime       `gorm:"column:lastModifiedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastModifiedAt"`
	LastModifiedByUser   *User          `gorm:"foreignKey:LastModifiedByUserID;belongsTo"`
}

func (SiteSettings) TableName() string {
	return "SiteSettings"
}

func (m *SiteSettings) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.LastModifiedAt.IsZero() {
		m.LastModifiedAt = DateTime{Time: now}
	}
	return nil
}

type BackgroundJob struct {
	ID            string              `gorm:"column:id;primaryKey" json:"id"`
	Status        string              `gorm:"column:status;default:'PENDING';not null" json:"status"`
	Payload       datatypes.JSON      `gorm:"column:payload" json:"payload"`
	Retried       int                 `gorm:"column:retried;size:32;default:0;not null" json:"retried"`
	MaxRetries    *int                `gorm:"column:maxRetries;size:32;default:3;not null" json:"maxRetries"`
	JobID         string              `gorm:"column:jobId;not null" json:"jobId"`
	Name          string              `gorm:"column:name;not null" json:"name"`
	Version       string              `gorm:"column:version;not null" json:"version"`
	SubmittedAt   DateTime            `gorm:"column:submittedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"submittedAt"`
	UpdatedAt     DateTime            `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	CompletedAt   *DateTime           `gorm:"column:completedAt;type:timestamp(3)" json:"completedAt"`
	LastRetriedAt *DateTime           `gorm:"column:lastRetriedAt;type:timestamp(3)" json:"lastRetriedAt"`
	Tasks         []BackgroundJobTask `gorm:"foreignKey:JobID;constraint:BackgroundJobTask_jobId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (BackgroundJob) TableName() string {
	return "BackgroundJob"
}

func (m *BackgroundJob) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.SubmittedAt.IsZero() {
		m.SubmittedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BackgroundJob) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BackgroundJobTask struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	Status        string         `gorm:"column:status;default:'PENDING';not null" json:"status"`
	Result        datatypes.JSON `gorm:"column:result" json:"result"`
	Retried       int            `gorm:"column:retried;size:32;default:0;not null" json:"retried"`
	MaxRetries    *int           `gorm:"column:maxRetries;size:32;default:3;not null" json:"maxRetries"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	CompletedAt   *DateTime      `gorm:"column:completedAt;type:timestamp(3)" json:"completedAt"`
	JobID         string         `gorm:"column:jobId;not null" json:"jobId"`
	BackgroundJob *BackgroundJob `gorm:"foreignKey:JobID;belongsTo"`
}

func (BackgroundJobTask) TableName() string {
	return "BackgroundJobTask"
}

func (m *BackgroundJobTask) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BackgroundJobTask) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AvatarImage struct {
	ID           string         `gorm:"column:id;primaryKey" json:"id"`
	Bytes        string         `gorm:"column:bytes;not null" json:"bytes"`
	Team         []Team         `gorm:"foreignKey:AvatarImageID;constraint:Team_avatarImageId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	User         []User         `gorm:"foreignKey:AvatarImageID;constraint:User_avatarImageId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Organisation []Organisation `gorm:"foreignKey:AvatarImageID;constraint:Organisation_avatarImageId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (AvatarImage) TableName() string {
	return "AvatarImage"
}

func (m *AvatarImage) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type EmailDomain struct {
	ID             string              `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime            `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime            `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Status         string              `gorm:"column:status;default:'PENDING';not null" json:"status"`
	Selector       string              `gorm:"column:selector;uniqueIndex:EmailDomain_selector_key;not null" json:"selector"`
	Domain         string              `gorm:"column:domain;uniqueIndex:EmailDomain_domain_key;not null" json:"domain"`
	PublicKey      string              `gorm:"column:publicKey;not null" json:"publicKey"`
	PrivateKey     string              `gorm:"column:privateKey;not null" json:"privateKey"`
	LastVerifiedAt *DateTime           `gorm:"column:lastVerifiedAt;type:timestamp(3)" json:"lastVerifiedAt"`
	OrganisationID string              `gorm:"column:organisationId;not null" json:"organisationId"`
	Organisation   *Organisation       `gorm:"belongsTo"`
	Emails         []OrganisationEmail `gorm:"foreignKey:EmailDomainID;constraint:OrganisationEmail_emailDomainId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (EmailDomain) TableName() string {
	return "EmailDomain"
}

func (m *EmailDomain) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EmailDomain) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganisationEmail struct {
	ID                         string                       `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                  DateTime                     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime                     `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Email                      string                       `gorm:"column:email;uniqueIndex:OrganisationEmail_email_key;not null" json:"email"`
	EmailName                  string                       `gorm:"column:emailName;not null" json:"emailName"`
	EmailDomainID              string                       `gorm:"column:emailDomainId;not null" json:"emailDomainId"`
	OrganisationID             string                       `gorm:"column:organisationId;not null" json:"organisationId"`
	EmailDomain                *EmailDomain                 `gorm:"belongsTo"`
	Organisation               *Organisation                `gorm:"belongsTo"`
	OrganisationGlobalSettings []OrganisationGlobalSettings `gorm:"foreignKey:EmailID;constraint:OrganisationGlobalSettings_emailId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TeamGlobalSettings         []TeamGlobalSettings         `gorm:"foreignKey:EmailID;constraint:TeamGlobalSettings_emailId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (OrganisationEmail) TableName() string {
	return "OrganisationEmail"
}

func (m *OrganisationEmail) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OrganisationEmail) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EmailTransport struct {
	ID                 string              `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt          DateTime            `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime            `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Name               string              `gorm:"column:name;not null" json:"name"`
	Type               string              `gorm:"column:type;not null" json:"type"`
	FromName           string              `gorm:"column:fromName;not null" json:"fromName"`
	FromAddress        string              `gorm:"column:fromAddress;not null" json:"fromAddress"`
	Config             string              `gorm:"column:config;not null" json:"config"`
	SubscriptionClaims []SubscriptionClaim `gorm:"foreignKey:EmailTransportID;constraint:SubscriptionClaim_emailTransportId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	OrganisationClaims []OrganisationClaim `gorm:"foreignKey:EmailTransportID;constraint:OrganisationClaim_emailTransportId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (EmailTransport) TableName() string {
	return "EmailTransport"
}

func (m *EmailTransport) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EmailTransport) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganisationAuthenticationPortal struct {
	ID                         string             `gorm:"column:id;primaryKey" json:"id"`
	Enabled                    bool               `gorm:"column:enabled;default:false;not null" json:"enabled"`
	ClientID                   string             `gorm:"column:clientId;default:'';not null" json:"clientId"`
	ClientSecret               string             `gorm:"column:clientSecret;default:'';not null" json:"clientSecret"`
	WellKnownURL               string             `gorm:"column:wellKnownUrl;default:'';not null" json:"wellKnownUrl"`
	DefaultOrganisationRole    string             `gorm:"column:defaultOrganisationRole;default:'MEMBER';not null" json:"defaultOrganisationRole"`
	AutoProvisionUsers         *bool              `gorm:"column:autoProvisionUsers;default:true;not null" json:"autoProvisionUsers"`
	AllowedDomains             ScalarList[string] `gorm:"column:allowedDomains;type:text[]" json:"allowedDomains"`
	AllowPersonalOrganisations bool               `gorm:"column:allowPersonalOrganisations;default:false;not null" json:"allowPersonalOrganisations"`
	Organisation               *Organisation      `gorm:"foreignKey:OrganisationAuthenticationPortalID;constraint:Organisation_organisationAuthenticationPortalId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (OrganisationAuthenticationPortal) TableName() string {
	return "OrganisationAuthenticationPortal"
}

type Counter struct {
	ID    string `gorm:"column:id;primaryKey" json:"id"`
	Value int    `gorm:"column:value;size:32;not null" json:"value"`
}

func (Counter) TableName() string {
	return "Counter"
}

// `key` is sha256(organisationId:lower(recipientEmail):clientUserId) so neither
// the recipient email nor the embedder's identifier is stored in plaintext.
type EmbedSignature struct {
	Key            string   `gorm:"column:key;primaryKey" json:"key"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OrganisationID string   `gorm:"column:organisationId;index:EmbedSignature_organisationId_idx;not null" json:"organisationId"`
	// Last writer. Not a relation so the row survives document deletion and
	// remains useful for abuse tracing.
	RecipientID  int           `gorm:"column:recipientId;index:EmbedSignature_recipientId_idx;size:32;not null" json:"recipientId"`
	Signature    string        `gorm:"column:signature;not null" json:"signature"`
	Organisation *Organisation `gorm:"belongsTo"`
}

func (EmbedSignature) TableName() string {
	return "EmbedSignature"
}

func (m *EmbedSignature) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EmbedSignature) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RateLimit struct {
	Key       string   `gorm:"column:key;primaryKey" json:"key"`
	Action    string   `gorm:"column:action;primaryKey" json:"action"`
	Bucket    DateTime `gorm:"column:bucket;primaryKey;type:timestamp(3)" json:"bucket"`
	Count     *int     `gorm:"column:count;size:32;default:1;not null" json:"count"`
	CreatedAt DateTime `gorm:"column:createdAt;index:RateLimit_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (RateLimit) TableName() string {
	return "RateLimit"
}

func (m *RateLimit) BeforeCreate(tx *gorm.DB) error {
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
