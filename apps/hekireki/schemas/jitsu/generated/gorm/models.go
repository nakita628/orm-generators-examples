package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lucsky/cuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EmailLog struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	Time       DateTime       `gorm:"column:time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"time"`
	Email      datatypes.JSON `gorm:"column:email;not null" json:"email"`
	Error      *string        `gorm:"column:error" json:"error"`
	PreviewURL *string        `gorm:"column:previewUrl" json:"previewUrl"`
	MessageID  *string        `gorm:"column:messageId" json:"messageId"`
	Status     string         `gorm:"column:status;not null" json:"status"`
}

func (EmailLog) TableName() string {
	return "EmailLog"
}

func (m *EmailLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.Time.IsZero() {
		m.Time = DateTime{Time: now}
	}
	return nil
}

type UserPreferences struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	UserID      string         `gorm:"column:userId;not null" json:"userId"`
	WorkspaceID *string        `gorm:"column:workspaceId" json:"workspaceId"`
	Scope       *string        `gorm:"column:scope" json:"scope"`
	Preferences datatypes.JSON `gorm:"column:preferences;not null" json:"preferences"`
	User        *UserProfile   `gorm:"foreignKey:UserID;belongsTo"`
	Workspace   *Workspace     `gorm:"belongsTo"`
}

func (UserPreferences) TableName() string {
	return "UserPreferences"
}

func (m *UserPreferences) BeforeCreate(tx *gorm.DB) error {
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

func (*UserPreferences) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type UserProfile struct {
	ID    string `gorm:"column:id;primaryKey" json:"id"`
	Name  string `gorm:"column:name;not null" json:"name"`
	Email string `gorm:"column:email;not null" json:"email"`
	Admin *bool  `gorm:"column:admin;default:false" json:"admin"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt               DateTime                  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	LoginProvider           string                    `gorm:"column:loginProvider;uniqueIndex:UserProfile_loginProvider_externalId_key,priority:1;not null" json:"loginProvider"`
	ExternalUsername        *string                   `gorm:"column:externalUsername" json:"externalUsername"`
	ExternalID              string                    `gorm:"column:externalId;uniqueIndex:UserProfile_loginProvider_externalId_key,priority:2;index:UserProfile_externalId_idx;not null" json:"externalId"`
	Tokens                  []UserAPIToken            `gorm:"foreignKey:UserID;constraint:UserApiToken_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	WorkspaceAccess         []WorkspaceAccess         `gorm:"foreignKey:UserID;constraint:WorkspaceAccess_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UserPreferences         []UserPreferences         `gorm:"foreignKey:UserID;constraint:UserPreferences_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	WorkspaceUserProperties []WorkspaceUserProperties `gorm:"foreignKey:UserID;constraint:WorkspaceUserProperties_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Password                *UserPassword             `gorm:"foreignKey:UserID;constraint:UserPassword_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (UserProfile) TableName() string {
	return "UserProfile"
}

func (m *UserProfile) BeforeCreate(tx *gorm.DB) error {
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

func (*UserProfile) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type UserPassword struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt         DateTime     `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	UserID            string       `gorm:"column:userId;uniqueIndex:UserPassword_userId_key;not null" json:"userId"`
	Hash              string       `gorm:"column:hash;not null" json:"hash"`
	ChangeAtNextLogin bool         `gorm:"column:changeAtNextLogin;default:false;not null" json:"changeAtNextLogin"`
	User              *UserProfile `gorm:"foreignKey:UserID;belongsTo"`
}

func (UserPassword) TableName() string {
	return "UserPassword"
}

func (m *UserPassword) BeforeCreate(tx *gorm.DB) error {
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

func (*UserPassword) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Workspace struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt               DateTime                  `gorm:"column:updatedAt;index:Workspace_updatedAt_idx,sort:desc;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Name                    string                    `gorm:"column:name;not null" json:"name"`
	Deleted                 bool                      `gorm:"column:deleted;default:false;not null" json:"deleted"`
	Slug                    *string                   `gorm:"column:slug;uniqueIndex:Workspace_slug_key;index:Workspace_slug_idx" json:"slug"`
	FeaturesEnabled         ScalarList[string]        `gorm:"column:featuresEnabled;type:text[]" json:"featuresEnabled"`
	WorkspaceAccess         []WorkspaceAccess         `gorm:"foreignKey:WorkspaceID;constraint:WorkspaceAccess_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	WorkspaceOptions        []WorkspaceOptions        `gorm:"foreignKey:WorkspaceID;constraint:WorkspaceOptions_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ConfigurationObject     []ConfigurationObject     `gorm:"foreignKey:WorkspaceID;constraint:ConfigurationObject_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ConfigurationObjectLink []ConfigurationObjectLink `gorm:"foreignKey:WorkspaceID;constraint:ConfigurationObjectLink_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Preferences             []UserPreferences         `gorm:"foreignKey:WorkspaceID;constraint:UserPreferences_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	WorkspaceUserProperties []WorkspaceUserProperties `gorm:"foreignKey:WorkspaceID;constraint:WorkspaceUserProperties_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ProfileBuilders         []ProfileBuilder          `gorm:"foreignKey:WorkspaceID;constraint:ProfileBuilder_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OidcLoginGroups         []OidcLoginGroup          `gorm:"foreignKey:WorkspaceID;constraint:OidcLoginGroup_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Workspace) TableName() string {
	return "Workspace"
}

func (m *Workspace) BeforeCreate(tx *gorm.DB) error {
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

func (*Workspace) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkspaceOptions struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	WorkspaceID string         `gorm:"column:workspaceId;index:WorkspaceOptions_workspaceId_namespace_idx,priority:1;not null" json:"workspaceId"`
	Namespace   string         `gorm:"column:namespace;index:WorkspaceOptions_workspaceId_namespace_idx,priority:2;not null" json:"namespace"`
	Value       datatypes.JSON `gorm:"column:value;not null" json:"value"`
	Workspace   *Workspace     `gorm:"belongsTo"`
}

func (WorkspaceOptions) TableName() string {
	return "WorkspaceOptions"
}

func (m *WorkspaceOptions) BeforeCreate(tx *gorm.DB) error {
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

func (*WorkspaceOptions) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkspaceAccess struct {
	UserID      string `gorm:"column:userId;primaryKey" json:"userId"`
	WorkspaceID string `gorm:"column:workspaceId;primaryKey" json:"workspaceId"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt DateTime     `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Role      string       `gorm:"column:role;default:'owner';not null" json:"role"`
	Workspace *Workspace   `gorm:"belongsTo"`
	User      *UserProfile `gorm:"foreignKey:UserID;belongsTo"`
}

func (WorkspaceAccess) TableName() string {
	return "WorkspaceAccess"
}

func (m *WorkspaceAccess) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkspaceAccess) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkspaceUserProperties struct {
	WorkspaceID string `gorm:"column:workspaceId;primaryKey" json:"workspaceId"`
	UserID      string `gorm:"column:userId;primaryKey" json:"userId"`
	// @zod.custom(z.coerce.date())
	LastUsed  DateTime     `gorm:"column:lastUsed;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastUsed"`
	Workspace *Workspace   `gorm:"belongsTo"`
	User      *UserProfile `gorm:"foreignKey:UserID;belongsTo"`
}

func (WorkspaceUserProperties) TableName() string {
	return "WorkspaceUserProperties"
}

func (m *WorkspaceUserProperties) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.LastUsed.IsZero() {
		m.LastUsed = DateTime{Time: now}
	}
	return nil
}

type UserAPIToken struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	// @zod.custom(z.coerce.date())
	LastUsed *DateTime `gorm:"column:lastUsed;type:timestamp(3)" json:"lastUsed"`
	Hint     string    `gorm:"column:hint;not null" json:"hint"`
	Hash     string    `gorm:"column:hash;not null" json:"hash"`
	UserID   string    `gorm:"column:userId;not null" json:"userId"`
	Type     *string   `gorm:"column:type" json:"type"`
	Name     *string   `gorm:"column:name" json:"name"`
	// @zod.custom(z.coerce.date())
	ExpiresAt     *DateTime          `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	WorkspaceID   *string            `gorm:"column:workspaceId" json:"workspaceId"`
	Role          *string            `gorm:"column:role" json:"role"`
	OauthClientID *string            `gorm:"column:oauthClientId" json:"oauthClientId"`
	User          *UserProfile       `gorm:"foreignKey:UserID;belongsTo"`
	OauthClient   *OAuthClient       `gorm:"foreignKey:OauthClientID;belongsTo"`
	AccessTokens  []OAuthAccessToken `gorm:"foreignKey:RefreshTokenID;constraint:OAuthAccessToken_refreshTokenId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (UserAPIToken) TableName() string {
	return "UserApiToken"
}

func (m *UserAPIToken) BeforeCreate(tx *gorm.DB) error {
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

func (*UserAPIToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OAuthClient struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt        DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	ClientSecretHash string             `gorm:"column:clientSecretHash;not null" json:"clientSecretHash"`
	Name             string             `gorm:"column:name;not null" json:"name"`
	RedirectUris     ScalarList[string] `gorm:"column:redirectUris;type:text[]" json:"redirectUris"`
	Tokens           []UserAPIToken     `gorm:"foreignKey:OauthClientID;constraint:UserApiToken_oauthClientId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (OAuthClient) TableName() string {
	return "OAuthClient"
}

func (m *OAuthClient) BeforeCreate(tx *gorm.DB) error {
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

func (*OAuthClient) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type KvStoreEntry struct {
	Key   string         `gorm:"column:key;primaryKey;index:kv_key_pattern_idx,expression:\"key\" text_pattern_ops" json:"key"`
	Value datatypes.JSON `gorm:"column:value;not null" json:"value"`
	// @zod.custom(z.coerce.date())
	ExpiresAt *DateTime `gorm:"column:expires_at;index:kv_expires_at_idx;type:timestamp(3)" json:"expires_at"`
}

func (KvStoreEntry) TableName() string {
	return "kv"
}

type OAuthAccessToken struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	LastUsed *DateTime `gorm:"column:lastUsed;type:timestamp(3)" json:"lastUsed"`
	Hash     string    `gorm:"column:hash;not null" json:"hash"`
	// @zod.custom(z.coerce.date())
	ExpiresAt      DateTime      `gorm:"column:expiresAt;index:OAuthAccessToken_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
	RefreshTokenID string        `gorm:"column:refreshTokenId;index:OAuthAccessToken_refreshTokenId_idx;not null" json:"refreshTokenId"`
	RefreshToken   *UserAPIToken `gorm:"foreignKey:RefreshTokenID;belongsTo"`
}

func (OAuthAccessToken) TableName() string {
	return "OAuthAccessToken"
}

func (m *OAuthAccessToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type InvitationToken struct {
	ID          string `gorm:"column:id;primaryKey" json:"id"`
	WorkspaceID string `gorm:"column:workspaceId;not null" json:"workspaceId"`
	Email       string `gorm:"column:email;not null" json:"email"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Token     string   `gorm:"column:token;not null" json:"token"`
	UsedBy    *string  `gorm:"column:usedBy" json:"usedBy"`
	Role      string   `gorm:"column:role;default:'owner';not null" json:"role"`
}

func (InvitationToken) TableName() string {
	return "InvitationToken"
}

func (m *InvitationToken) BeforeCreate(tx *gorm.DB) error {
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

func (*InvitationToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ConfigurationObject struct {
	ID      string `gorm:"column:id;primaryKey" json:"id"`
	Deleted *bool  `gorm:"column:deleted;default:false" json:"deleted"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt              DateTime                  `gorm:"column:updatedAt;index:ConfigurationObject_updatedAt_idx,sort:desc;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Type                   string                    `gorm:"column:type;not null" json:"type"`
	WorkspaceID            string                    `gorm:"column:workspaceId;not null" json:"workspaceId"`
	Config                 datatypes.JSON            `gorm:"column:config" json:"config"`
	Workspace              *Workspace                `gorm:"belongsTo"`
	FromLinks              []ConfigurationObjectLink `gorm:"foreignKey:ToID;constraint:ConfigurationObjectLink_toId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ToLinks                []ConfigurationObjectLink `gorm:"foreignKey:FromID;constraint:ConfigurationObjectLink_fromId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ProfileBuilder         []ProfileBuilder          `gorm:"foreignKey:DestinationID;constraint:ProfileBuilder_destinationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProfileBuilderFunction []ProfileBuilderFunction  `gorm:"foreignKey:FunctionID;constraint:ProfileBuilderFunction_functionId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (ConfigurationObject) TableName() string {
	return "ConfigurationObject"
}

func (m *ConfigurationObject) BeforeCreate(tx *gorm.DB) error {
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

func (*ConfigurationObject) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ConfigurationObjectLink struct {
	ID      string `gorm:"column:id;primaryKey" json:"id"`
	Deleted *bool  `gorm:"column:deleted;default:false" json:"deleted"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt   DateTime             `gorm:"column:updatedAt;index:ConfigurationObjectLink_updatedAt_idx,sort:desc;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Type        *string              `gorm:"column:type;default:'push'" json:"type"`
	Data        datatypes.JSON       `gorm:"column:data" json:"data"`
	WorkspaceID string               `gorm:"column:workspaceId;index:ConfigurationObjectLink_workspaceId_idx;not null" json:"workspaceId"`
	FromID      string               `gorm:"column:fromId;not null" json:"fromId"`
	ToID        string               `gorm:"column:toId;not null" json:"toId"`
	Workspace   *Workspace           `gorm:"belongsTo"`
	From        *ConfigurationObject `gorm:"foreignKey:FromID;belongsTo"`
	To          *ConfigurationObject `gorm:"foreignKey:ToID;belongsTo"`
}

func (ConfigurationObjectLink) TableName() string {
	return "ConfigurationObjectLink"
}

func (m *ConfigurationObjectLink) BeforeCreate(tx *gorm.DB) error {
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

func (*ConfigurationObjectLink) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ConnectorPackage struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	PackageID   string         `gorm:"column:packageId;not null" json:"packageId"`
	PackageType *string        `gorm:"column:packageType;default:'airbyte';not null" json:"packageType"`
	Meta        datatypes.JSON `gorm:"column:meta" json:"meta"`
	LogoSvg     []byte         `gorm:"column:logoSvg" json:"logoSvg"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (ConnectorPackage) TableName() string {
	return "ConnectorPackage"
}

func (m *ConnectorPackage) BeforeCreate(tx *gorm.DB) error {
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

func (*ConnectorPackage) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AuditLog struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	Timestamp      DateTime       `gorm:"column:timestamp;index:AuditLog_timestamp_idx;index:AuditLog_workspaceId_timestamp_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
	Type           string         `gorm:"column:type;index:AuditLog_type_idx;not null" json:"type"`
	Severity       *string        `gorm:"column:severity;default:'info'" json:"severity"`
	UserID         *string        `gorm:"column:userId" json:"userId"`
	WorkspaceID    *string        `gorm:"column:workspaceId;index:AuditLog_workspaceId_timestamp_idx,priority:1" json:"workspaceId"`
	ObjectID       *string        `gorm:"column:objectId" json:"objectId"`
	Changes        datatypes.JSON `gorm:"column:changes" json:"changes"`
	AuthType       *string        `gorm:"column:authType" json:"authType"`
	TokenID        *string        `gorm:"column:tokenId" json:"tokenId"`
	IP             *string        `gorm:"column:ip" json:"ip"`
	RequestHeaders datatypes.JSON `gorm:"column:requestHeaders" json:"requestHeaders"`
}

func (AuditLog) TableName() string {
	return "AuditLog"
}

func (m *AuditLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type GlobalProps struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	Value     datatypes.JSON `gorm:"column:value" json:"value"`
}

func (GlobalProps) TableName() string {
	return "GlobalProps"
}

func (m *GlobalProps) BeforeCreate(tx *gorm.DB) error {
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

func (*GlobalProps) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SourceSpec struct {
	Package string         `gorm:"column:package;primaryKey" json:"package"`
	Version string         `gorm:"column:version;primaryKey" json:"version"`
	Specs   datatypes.JSON `gorm:"column:specs" json:"specs"`
	// @zod.custom(z.coerce.date())
	Timestamp DateTime `gorm:"column:timestamp;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
	Error     *string  `gorm:"column:error" json:"error"`
}

func (SourceSpec) TableName() string {
	return "source_spec"
}

func (m *SourceSpec) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type SourceCheck struct {
	Package     string  `gorm:"column:package;not null" json:"package"`
	Version     string  `gorm:"column:version;not null" json:"version"`
	Key         string  `gorm:"column:key;primaryKey" json:"key"`
	Status      string  `gorm:"column:status;not null" json:"status"`
	Description *string `gorm:"column:description" json:"description"`
	// @zod.custom(z.coerce.date())
	Timestamp DateTime `gorm:"column:timestamp;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
}

func (SourceCheck) TableName() string {
	return "source_check"
}

func (m *SourceCheck) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type SourceCatalog struct {
	Package string         `gorm:"column:package;primaryKey" json:"package"`
	Version string         `gorm:"column:version;primaryKey" json:"version"`
	Key     string         `gorm:"column:key;primaryKey" json:"key"`
	Catalog datatypes.JSON `gorm:"column:catalog" json:"catalog"`
	// @zod.custom(z.coerce.date())
	Timestamp   DateTime `gorm:"column:timestamp;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
	Status      string   `gorm:"column:status;not null" json:"status"`
	Description *string  `gorm:"column:description" json:"description"`
}

func (SourceCatalog) TableName() string {
	return "source_catalog"
}

func (m *SourceCatalog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type ReverseSyncControl struct {
	WorkspaceID         string  `gorm:"column:workspace_id;primaryKey;uniqueIndex:reverse_sync_control_workspace_id_sync_id_run_order_key,priority:1" json:"workspace_id"`
	SyncID              string  `gorm:"column:sync_id;primaryKey;uniqueIndex:reverse_sync_control_workspace_id_sync_id_run_order_key,priority:2;index:reverse_sync_control_sync_id_idx" json:"sync_id"`
	RunID               string  `gorm:"column:run_id;primaryKey" json:"run_id"`
	RunOrder            int64   `gorm:"column:run_order;uniqueIndex:reverse_sync_control_workspace_id_sync_id_run_order_key,priority:3;default:0;not null" json:"run_order"`
	Detached            bool    `gorm:"column:detached;default:false;not null" json:"detached"`
	Revision            string  `gorm:"column:revision;not null" json:"revision"`
	TargetHash          string  `gorm:"column:target_hash;not null" json:"target_hash"`
	Phase               *string `gorm:"column:phase;default:'new';not null" json:"phase"`
	Mode                string  `gorm:"column:mode;not null" json:"mode"`
	Extraction          string  `gorm:"column:extraction;not null" json:"extraction"`
	ArtifactHead        []byte  `gorm:"column:artifact_head" json:"artifact_head"`
	BaseSequence        int64   `gorm:"column:base_sequence;default:0;not null" json:"base_sequence"`
	NextSequence        int64   `gorm:"column:next_sequence;default:0;not null" json:"next_sequence"`
	CheckpointSequence  int64   `gorm:"column:checkpoint_sequence;default:0;not null" json:"checkpoint_sequence"`
	FinishSequence      *int64  `gorm:"column:finish_sequence" json:"finish_sequence"`
	Store               []byte  `gorm:"column:store" json:"store"`
	ProviderState       []byte  `gorm:"column:provider_state" json:"provider_state"`
	FinishResult        []byte  `gorm:"column:finish_result" json:"finish_result"`
	CommittedGeneration *string `gorm:"column:committed_generation" json:"committed_generation"`
}

func (ReverseSyncControl) TableName() string {
	return "reverse_sync_control"
}

type ReverseSyncTargetOwner struct {
	TargetHash  string `gorm:"column:target_hash;primaryKey" json:"target_hash"`
	WorkspaceID string `gorm:"column:workspace_id;uniqueIndex:reverse_sync_target_owner_workspace_id_sync_id_key,priority:1;not null" json:"workspace_id"`
	SyncID      string `gorm:"column:sync_id;uniqueIndex:reverse_sync_target_owner_workspace_id_sync_id_key,priority:2;not null" json:"sync_id"`
}

func (ReverseSyncTargetOwner) TableName() string {
	return "reverse_sync_target_owner"
}

type SourceState struct {
	SyncID string         `gorm:"column:sync_id;primaryKey" json:"sync_id"`
	Stream string         `gorm:"column:stream;primaryKey" json:"stream"`
	State  datatypes.JSON `gorm:"column:state;not null" json:"state"`
	// @zod.custom(z.coerce.date())
	Timestamp DateTime `gorm:"column:timestamp;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
}

func (SourceState) TableName() string {
	return "source_state"
}

func (m *SourceState) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type SourceTask struct {
	SyncID  string `gorm:"column:sync_id;index:source_task_sync_id_idx;index:source_task_sync_id_started_at_idx,priority:1;not null" json:"sync_id"`
	TaskID  string `gorm:"column:task_id;primaryKey" json:"task_id"`
	Package string `gorm:"column:package;not null" json:"package"`
	Version string `gorm:"column:version;not null" json:"version"`
	// @zod.custom(z.coerce.date())
	StartedAt DateTime `gorm:"column:started_at;index:source_task_started_at_idx;index:source_task_sync_id_started_at_idx,priority:2,sort:desc;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"started_at"`
	// @zod.custom(z.coerce.date())
	UpdatedAt   DateTime       `gorm:"column:updated_at;type:timestamptz(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
	StartedBy   datatypes.JSON `gorm:"column:started_by" json:"started_by"`
	Status      string         `gorm:"column:status;not null" json:"status"`
	Description *string        `gorm:"column:description" json:"description"`
	Metrics     datatypes.JSON `gorm:"column:metrics" json:"metrics"`
	Error       *string        `gorm:"column:error" json:"error"`
}

func (SourceTask) TableName() string {
	return "source_task"
}

func (m *SourceTask) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.StartedAt.IsZero() {
		m.StartedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

type TaskLog struct {
	ID      string `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	Level   string `gorm:"column:level;not null" json:"level"`
	Logger  string `gorm:"column:logger;not null" json:"logger"`
	Message string `gorm:"column:message;not null" json:"message"`
	SyncID  string `gorm:"column:sync_id;index:task_log_sync_id_idx;not null" json:"sync_id"`
	TaskID  string `gorm:"column:task_id;index:task_log_task_id_idx;not null" json:"task_id"`
	// @zod.custom(z.coerce.date())
	Timestamp DateTime `gorm:"column:timestamp;index:task_log_timestamp_idx;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
}

func (TaskLog) TableName() string {
	return "task_log"
}

func (m *TaskLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type ProfileBuilder struct {
	ID      string `gorm:"column:id;primaryKey" json:"id"`
	Name    string `gorm:"column:name;not null" json:"name"`
	Deleted *bool  `gorm:"column:deleted;default:false" json:"deleted"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt                      DateTime                 `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	WorkspaceID                    string                   `gorm:"column:workspaceId;not null" json:"workspaceId"`
	IntermediateStorageCredentials datatypes.JSON           `gorm:"column:intermediateStorageCredentials;not null" json:"intermediateStorageCredentials"`
	DestinationID                  *string                  `gorm:"column:destinationId" json:"destinationId"`
	ConnectionOptions              datatypes.JSON           `gorm:"column:connectionOptions" json:"connectionOptions"`
	Version                        int                      `gorm:"column:version;size:32;default:0;not null" json:"version"`
	Workspace                      *Workspace               `gorm:"belongsTo"`
	Destination                    *ConfigurationObject     `gorm:"foreignKey:DestinationID;belongsTo"`
	Functions                      []ProfileBuilderFunction `gorm:"foreignKey:ProfileBuilderID;constraint:ProfileBuilderFunction_profileBuilderId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (ProfileBuilder) TableName() string {
	return "ProfileBuilder"
}

func (m *ProfileBuilder) BeforeCreate(tx *gorm.DB) error {
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

func (*ProfileBuilder) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ProfileBuilderFunction struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt        DateTime             `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	ProfileBuilderID string               `gorm:"column:profileBuilderId;not null" json:"profileBuilderId"`
	FunctionID       string               `gorm:"column:functionId;not null" json:"functionId"`
	ProfileBuilder   *ProfileBuilder      `gorm:"belongsTo"`
	Function         *ConfigurationObject `gorm:"foreignKey:FunctionID;belongsTo"`
}

func (ProfileBuilderFunction) TableName() string {
	return "ProfileBuilderFunction"
}

func (m *ProfileBuilderFunction) BeforeCreate(tx *gorm.DB) error {
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

func (*ProfileBuilderFunction) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ProfileBuilderState struct {
	ProfileBuilderID      string    `gorm:"column:profileBuilderId;primaryKey" json:"profileBuilderId"`
	ProfileBuilderVersion int       `gorm:"column:profileBuilderVersion;primaryKey;size:32" json:"profileBuilderVersion"`
	TotalInstances        int       `gorm:"column:totalInstances;primaryKey;size:32" json:"totalInstances"`
	InstanceIndex         int       `gorm:"column:instanceIndex;primaryKey;size:32" json:"instanceIndex"`
	StartedAt             DateTime  `gorm:"column:startedAt;type:timestamp(3);not null" json:"startedAt"`
	UpdatedAt             DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	LastTimestamp         *DateTime `gorm:"column:lastTimestamp;type:timestamp(3)" json:"lastTimestamp"`
	ProcessedUsers        int       `gorm:"column:processedUsers;size:32;not null" json:"processedUsers"`
	TotalUsers            int       `gorm:"column:totalUsers;size:32;not null" json:"totalUsers"`
	ErrorUsers            int       `gorm:"column:errorUsers;size:32;not null" json:"errorUsers"`
	Speed                 float64   `gorm:"column:speed;type:double precision;not null" json:"speed"`
}

func (ProfileBuilderState) TableName() string {
	return "ProfileBuilderState"
}

type ProfileBuilderState2 struct {
	ProfileBuilderID string         `gorm:"column:profileBuilderId;primaryKey" json:"profileBuilderId"`
	UpdatedAt        DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	FullRebuildInfo  datatypes.JSON `gorm:"column:fullRebuildInfo" json:"fullRebuildInfo"`
	QueuesInfo       datatypes.JSON `gorm:"column:queuesInfo" json:"queuesInfo"`
	Metrics          datatypes.JSON `gorm:"column:metrics" json:"metrics"`
}

func (ProfileBuilderState2) TableName() string {
	return "ProfileBuilderState2"
}

type StatusChange struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	WorkspaceID string `gorm:"column:workspaceId;not null" json:"workspaceId"`
	ActorID     string `gorm:"column:actorId;index:StatusChange_actorId_type_tableName_idx,priority:1;not null" json:"actorId"`
	Type        string `gorm:"column:type;index:StatusChange_actorId_type_tableName_idx,priority:2;not null" json:"type"`
	TableName_  string `gorm:"column:tableName;index:StatusChange_actorId_type_tableName_idx,priority:3;not null" json:"tableName"`
	// @zod.custom(z.coerce.date())
	Timestamp   DateTime `gorm:"column:timestamp;index:StatusChange_timestamp_idx;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
	Status      string   `gorm:"column:status;not null" json:"status"`
	StartedAt   DateTime `gorm:"column:startedAt;index:StatusChange_startedAt_idx;type:timestamp(3);not null" json:"startedAt"`
	Counts      int      `gorm:"column:counts;size:32;default:0;not null" json:"counts"`
	Description *string  `gorm:"column:description" json:"description"`
	QueueSize   int      `gorm:"column:queueSize;size:32;default:0;not null" json:"queueSize"`
}

func (StatusChange) TableName() string {
	return "StatusChange"
}

func (m *StatusChange) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type NotificationState struct {
	WorkspaceID string `gorm:"column:workspaceId;not null" json:"workspaceId"`
	ChannelID   string `gorm:"column:channelId;primaryKey" json:"channelId"`
	ActorID     string `gorm:"column:actorId;primaryKey" json:"actorId"`
	Type        string `gorm:"column:type;primaryKey" json:"type"`
	TableName_  string `gorm:"column:tableName;primaryKey" json:"tableName"`
	// @zod.custom(z.coerce.date())
	FlappingSince *DateTime `gorm:"column:flappingSince;type:timestamptz(3)" json:"flappingSince"`
	// @zod.custom(z.coerce.date())
	LastNotification DateTime `gorm:"column:lastNotification;type:timestamptz(3);not null" json:"lastNotification"`
	StatusChangeID   int64    `gorm:"column:statusChangeId;not null" json:"statusChangeId"`
	Error            *string  `gorm:"column:error" json:"error"`
}

func (NotificationState) TableName() string {
	return "NotificationState"
}

type Notification struct {
	ID             string `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	WorkspaceID    string `gorm:"column:workspaceId;not null" json:"workspaceId"`
	ActorID        string `gorm:"column:actorId;index:Notification_actorId_type_tableName_idx,priority:1;not null" json:"actorId"`
	Type           string `gorm:"column:type;index:Notification_actorId_type_tableName_idx,priority:2;not null" json:"type"`
	TableName_     string `gorm:"column:tableName;index:Notification_actorId_type_tableName_idx,priority:3;not null" json:"tableName"`
	ChannelID      string `gorm:"column:channelId;not null" json:"channelId"`
	StatusChangeID int64  `gorm:"column:statusChangeId;not null" json:"statusChangeId"`
	// @zod.custom(z.coerce.date())
	Timestamp DateTime `gorm:"column:timestamp;index:Notification_timestamp_idx;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
	Status    string   `gorm:"column:status;not null" json:"status"`
	Error     *string  `gorm:"column:error" json:"error"`
}

func (Notification) TableName() string {
	return "Notification"
}

func (m *Notification) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type OidcProvider struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt             DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Name                  string             `gorm:"column:name;not null" json:"name"`
	Issuer                string             `gorm:"column:issuer;uniqueIndex:OidcProvider_issuer_clientId_key,priority:1;not null" json:"issuer"`
	ClientID              string             `gorm:"column:clientId;uniqueIndex:OidcProvider_issuer_clientId_key,priority:2;not null" json:"clientId"`
	ClientSecret          string             `gorm:"column:clientSecret;not null" json:"clientSecret"`
	AuthorizationEndpoint *string            `gorm:"column:authorizationEndpoint" json:"authorizationEndpoint"`
	TokenEndpoint         *string            `gorm:"column:tokenEndpoint" json:"tokenEndpoint"`
	UserinfoEndpoint      *string            `gorm:"column:userinfoEndpoint" json:"userinfoEndpoint"`
	JwksURI               *string            `gorm:"column:jwksUri" json:"jwksUri"`
	IntrospectionEndpoint *string            `gorm:"column:introspectionEndpoint" json:"introspectionEndpoint"`
	Scopes                ScalarList[string] `gorm:"column:scopes;type:text[]" json:"scopes"`
	Audience              *string            `gorm:"column:audience" json:"audience"`
	Prompt                *string            `gorm:"column:prompt" json:"prompt"`
	UsernameClaim         *string            `gorm:"column:usernameClaim;default:'email';not null" json:"usernameClaim"`
	EmailClaim            *string            `gorm:"column:emailClaim;default:'email';not null" json:"emailClaim"`
	NameClaim             *string            `gorm:"column:nameClaim;default:'name';not null" json:"nameClaim"`
	GroupClaim            *string            `gorm:"column:groupClaim;default:'groups';not null" json:"groupClaim"`
	Enabled               *bool              `gorm:"column:enabled;default:true;not null" json:"enabled"`
	AutoDiscovery         *bool              `gorm:"column:autoDiscovery;default:true;not null" json:"autoDiscovery"`
	OidcLoginGroups       []OidcLoginGroup   `gorm:"foreignKey:OidcProviderID;constraint:OidcLoginGroup_oidcProviderId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (OidcProvider) TableName() string {
	return "OidcProvider"
}

func (m *OidcProvider) BeforeCreate(tx *gorm.DB) error {
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

func (*OidcProvider) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OidcLoginGroup struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// @zod.custom(z.coerce.date())
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.custom(z.coerce.date())
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	WorkspaceID    string        `gorm:"column:workspaceId;uniqueIndex:OidcLoginGroup_workspaceId_oidcProviderId_key,priority:1;index:OidcLoginGroup_workspaceId_idx;not null" json:"workspaceId"`
	OidcProviderID string        `gorm:"column:oidcProviderId;uniqueIndex:OidcLoginGroup_workspaceId_oidcProviderId_key,priority:2;index:OidcLoginGroup_oidcProviderId_idx;not null" json:"oidcProviderId"`
	GroupValue     *string       `gorm:"column:groupValue" json:"groupValue"`
	AllowAllUsers  bool          `gorm:"column:allowAllUsers;default:false;not null" json:"allowAllUsers"`
	Workspace      *Workspace    `gorm:"belongsTo"`
	OidcProvider   *OidcProvider `gorm:"belongsTo"`
}

func (OidcLoginGroup) TableName() string {
	return "OidcLoginGroup"
}

func (m *OidcLoginGroup) BeforeCreate(tx *gorm.DB) error {
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

func (*OidcLoginGroup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ReprocessingJobs struct {
	ID                  string                `gorm:"column:id;primaryKey;type:varchar(64)" json:"id"`
	Config              datatypes.JSON        `gorm:"column:config;not null" json:"config"`
	Status              string                `gorm:"column:status;type:varchar(32);not null" json:"status"`
	CreatedAt           DateTime              `gorm:"column:created_at;index:reprocessing_jobs_created_at_idx,sort:desc;type:timestamptz(6);autoCreateTime:false;not null" json:"created_at"`
	StartedAt           *DateTime             `gorm:"column:started_at;type:timestamptz(6)" json:"started_at"`
	CompletedAt         *DateTime             `gorm:"column:completed_at;type:timestamptz(6)" json:"completed_at"`
	K8sJobName          *string               `gorm:"column:k8s_job_name;type:varchar(255)" json:"k8s_job_name"`
	TotalFiles          *int                  `gorm:"column:total_files;size:32;default:0" json:"total_files"`
	TotalWorkers        *int                  `gorm:"column:total_workers;size:32;default:0" json:"total_workers"`
	Error               *string               `gorm:"column:error" json:"error"`
	ReprocessingWorkers []ReprocessingWorkers `gorm:"foreignKey:JobID;constraint:reprocessing_workers_job_id_fkey,OnUpdate:NO ACTION,OnDelete:CASCADE"`
}

func (ReprocessingJobs) TableName() string {
	return "reprocessing_jobs"
}

func (m *ReprocessingJobs) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type ReprocessingWorkers struct {
	JobID          string            `gorm:"column:job_id;primaryKey;index:reprocessing_workers_job_id_idx;index:reprocessing_workers_status_idx,priority:1;type:varchar(64)" json:"job_id"`
	WorkerIndex    int               `gorm:"column:worker_index;primaryKey;size:32" json:"worker_index"`
	Status         string            `gorm:"column:status;index:reprocessing_workers_status_idx,priority:2;type:varchar(32);not null" json:"status"`
	StartedAt      *DateTime         `gorm:"column:started_at;type:timestamptz(6)" json:"started_at"`
	UpdatedAt      DateTime          `gorm:"column:updated_at;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
	CompletedAt    *DateTime         `gorm:"column:completed_at;type:timestamptz(6)" json:"completed_at"`
	CurrentFile    *string           `gorm:"column:current_file" json:"current_file"`
	CurrentLine    *int64            `gorm:"column:current_line;default:0" json:"current_line"`
	CurrentTs      *string           `gorm:"column:current_ts" json:"current_ts"`
	AssignedFiles  *int              `gorm:"column:assigned_files;size:32;default:0" json:"assigned_files"`
	ProcessedFiles *int              `gorm:"column:processed_files;size:32;default:0" json:"processed_files"`
	TotalLines     *int64            `gorm:"column:total_lines;default:0" json:"total_lines"`
	SuccessCount   *int64            `gorm:"column:success_count;default:0" json:"success_count"`
	ErrorCount     *int64            `gorm:"column:error_count;default:0" json:"error_count"`
	SkippedCount   *int64            `gorm:"column:skipped_count;default:0" json:"skipped_count"`
	TotalBytes     *int64            `gorm:"column:total_bytes;default:0" json:"total_bytes"`
	ProcessedBytes *int64            `gorm:"column:processed_bytes;default:0" json:"processed_bytes"`
	Error          *string           `gorm:"column:error" json:"error"`
	Job            *ReprocessingJobs `gorm:"foreignKey:JobID;belongsTo"`
}

func (ReprocessingWorkers) TableName() string {
	return "reprocessing_workers"
}

func (m *ReprocessingWorkers) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

type FunctionsServer struct {
	WorkspaceID      string             `gorm:"column:workspaceId;primaryKey" json:"workspaceId"`
	Class            string             `gorm:"column:class;primaryKey" json:"class"`
	DeploymentID     string             `gorm:"column:deploymentId;not null" json:"deploymentId"`
	Connections      ScalarList[string] `gorm:"column:connections;type:text[]" json:"connections"`
	EmptyConnections ScalarList[string] `gorm:"column:emptyConnections;type:text[]" json:"emptyConnections"`
	ProfileBuilders  ScalarList[string] `gorm:"column:profileBuilders;type:text[]" json:"profileBuilders"`
	CreatedAt        DateTime           `gorm:"column:createdAt;type:timestamptz(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        *DateTime          `gorm:"column:updatedAt;type:timestamptz(3);autoUpdateTime:false" json:"updatedAt"`
}

func (FunctionsServer) TableName() string {
	return "FunctionsServer"
}

func (m *FunctionsServer) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type RateLimitCounter struct {
	BucketKey   string   `gorm:"column:bucketKey;primaryKey" json:"bucketKey"`
	WindowStart DateTime `gorm:"column:windowStart;primaryKey;type:timestamptz(3)" json:"windowStart"`
	Count       int      `gorm:"column:count;size:32;default:0;not null" json:"count"`
	ExpiresAt   DateTime `gorm:"column:expiresAt;index:RateLimitCounter_expiresAt_idx;type:timestamptz(3);not null" json:"expiresAt"`
}

func (RateLimitCounter) TableName() string {
	return "RateLimitCounter"
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
