package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Account struct {
	ID                    string  `gorm:"column:id;primaryKey" json:"id"`
	UserID                string  `gorm:"column:userId;not null" json:"userId"`
	Type                  string  `gorm:"column:type;not null" json:"type"`
	Provider              string  `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key;not null" json:"provider"`
	ProviderAccountID     string  `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key;not null" json:"providerAccountId"`
	RefreshToken          *string `gorm:"column:refresh_token" json:"refresh_token"`
	AccessToken           *string `gorm:"column:access_token" json:"access_token"`
	ExpiresAt             *int    `gorm:"column:expires_at;size:32" json:"expires_at"`
	TokenType             *string `gorm:"column:token_type" json:"token_type"`
	Scope                 *string `gorm:"column:scope" json:"scope"`
	IDToken               *string `gorm:"column:id_token" json:"id_token"`
	SessionState          *string `gorm:"column:session_state" json:"session_state"`
	OauthTokenSecret      *string `gorm:"column:oauth_token_secret" json:"oauth_token_secret"`
	OauthToken            *string `gorm:"column:oauth_token" json:"oauth_token"`
	RefreshTokenExpiresIn *int    `gorm:"column:refresh_token_expires_in;size:32" json:"refresh_token_expires_in"`
	User                  *User   `gorm:"belongsTo"`
}

func (Account) TableName() string {
	return "Account"
}

func (m *Account) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type Session struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	SessionToken string   `gorm:"column:sessionToken;uniqueIndex:Session_sessionToken_key;not null" json:"sessionToken"`
	UserID       string   `gorm:"column:userId;not null" json:"userId"`
	Expires      DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	User         *User    `gorm:"belongsTo"`
}

func (Session) TableName() string {
	return "Session"
}

func (m *Session) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type User struct {
	ID                          string                    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                   DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                   DateTime                  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	LastActivityAt              DateTime                  `gorm:"column:lastActivityAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastActivityAt"`
	Name                        *string                   `gorm:"column:name" json:"name"`
	Email                       *string                   `gorm:"column:email;uniqueIndex:User_email_key" json:"email"`
	EmailVerified               *DateTime                 `gorm:"column:emailVerified;type:timestamp(3)" json:"emailVerified"`
	Image                       *string                   `gorm:"column:image" json:"image"`
	Company                     *string                   `gorm:"column:company" json:"company"`
	OnboardingCategories        datatypes.JSON            `gorm:"column:onboardingCategories;not null" json:"onboardingCategories"`
	Referral                    *string                   `gorm:"column:referral" json:"referral"`
	GraphNavigation             *string                   `gorm:"column:graphNavigation" json:"graphNavigation"`
	PreferredAppAppearance      *string                   `gorm:"column:preferredAppAppearance" json:"preferredAppAppearance"`
	DisplayedInAppNotifications datatypes.JSON            `gorm:"column:displayedInAppNotifications" json:"displayedInAppNotifications"`
	GroupTitlesAutoGeneration   datatypes.JSON            `gorm:"column:groupTitlesAutoGeneration" json:"groupTitlesAutoGeneration"`
	PreferredLanguage           *string                   `gorm:"column:preferredLanguage" json:"preferredLanguage"`
	TermsAcceptedAt             *DateTime                 `gorm:"column:termsAcceptedAt;type:timestamp(3)" json:"termsAcceptedAt"`
	Accounts                    []Account                 `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APITokens                   []APIToken                `gorm:"foreignKey:OwnerID;constraint:ApiToken_ownerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CollaboratorsOnTypebots     []CollaboratorsOnTypebots `gorm:"foreignKey:UserID;constraint:CollaboratorsOnTypebots_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Workspaces                  []MemberInWorkspace       `gorm:"foreignKey:UserID;constraint:MemberInWorkspace_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions                    []Session                 `gorm:"foreignKey:UserID;constraint:Session_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BannedIps                   []BannedIP                `gorm:"foreignKey:UserID;constraint:BannedIp_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Credentials                 []UserCredentials         `gorm:"foreignKey:UserID;constraint:UserCredentials_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	if m.LastActivityAt.IsZero() {
		m.LastActivityAt = DateTime{Time: now}
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

type APIToken struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Token     string   `gorm:"column:token;uniqueIndex:ApiToken_token_key;not null" json:"token"`
	Name      string   `gorm:"column:name;not null" json:"name"`
	OwnerID   string   `gorm:"column:ownerId;not null" json:"ownerId"`
	Owner     *User    `gorm:"foreignKey:OwnerID;belongsTo"`
}

func (APIToken) TableName() string {
	return "ApiToken"
}

func (m *APIToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Workspace struct {
	ID                            string                `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                     DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                     DateTime              `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Name                          string                `gorm:"column:name;not null" json:"name"`
	Icon                          *string               `gorm:"column:icon" json:"icon"`
	Plan                          string                `gorm:"column:plan;default:'FREE';not null" json:"plan"`
	StripeID                      *string               `gorm:"column:stripeId;uniqueIndex:Workspace_stripeId_key" json:"stripeId"`
	AdditionalChatsIndex          int                   `gorm:"column:additionalChatsIndex;size:32;default:0;not null" json:"additionalChatsIndex"`
	AdditionalStorageIndex        int                   `gorm:"column:additionalStorageIndex;size:32;default:0;not null" json:"additionalStorageIndex"`
	ChatsLimitFirstEmailSentAt    *DateTime             `gorm:"column:chatsLimitFirstEmailSentAt;type:timestamp(3)" json:"chatsLimitFirstEmailSentAt"`
	StorageLimitFirstEmailSentAt  *DateTime             `gorm:"column:storageLimitFirstEmailSentAt;type:timestamp(3)" json:"storageLimitFirstEmailSentAt"`
	ChatsLimitSecondEmailSentAt   *DateTime             `gorm:"column:chatsLimitSecondEmailSentAt;type:timestamp(3)" json:"chatsLimitSecondEmailSentAt"`
	StorageLimitSecondEmailSentAt *DateTime             `gorm:"column:storageLimitSecondEmailSentAt;type:timestamp(3)" json:"storageLimitSecondEmailSentAt"`
	CustomChatsLimit              *int                  `gorm:"column:customChatsLimit;size:32" json:"customChatsLimit"`
	CustomStorageLimit            *int                  `gorm:"column:customStorageLimit;size:32" json:"customStorageLimit"`
	CustomSeatsLimit              *int                  `gorm:"column:customSeatsLimit;size:32" json:"customSeatsLimit"`
	IsQuarantined                 bool                  `gorm:"column:isQuarantined;default:false;not null" json:"isQuarantined"`
	IsSuspended                   bool                  `gorm:"column:isSuspended;default:false;not null" json:"isSuspended"`
	IsPastDue                     bool                  `gorm:"column:isPastDue;default:false;not null" json:"isPastDue"`
	IsVerified                    *bool                 `gorm:"column:isVerified" json:"isVerified"`
	ChatsHardLimit                *int                  `gorm:"column:chatsHardLimit;size:32" json:"chatsHardLimit"`
	Settings                      datatypes.JSON        `gorm:"column:settings" json:"settings"`
	LastActivityAt                *DateTime             `gorm:"column:lastActivityAt;type:timestamp(3)" json:"lastActivityAt"`
	Credentials                   []Credentials         `gorm:"foreignKey:WorkspaceID;constraint:Credentials_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomDomains                 []CustomDomain        `gorm:"foreignKey:WorkspaceID;constraint:CustomDomain_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Folders                       []DashboardFolder     `gorm:"foreignKey:WorkspaceID;constraint:DashboardFolder_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Members                       []MemberInWorkspace   `gorm:"foreignKey:WorkspaceID;constraint:MemberInWorkspace_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Typebots                      []Typebot             `gorm:"foreignKey:WorkspaceID;constraint:Typebot_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invitations                   []WorkspaceInvitation `gorm:"foreignKey:WorkspaceID;constraint:WorkspaceInvitation_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ThemeTemplates                []ThemeTemplate       `gorm:"foreignKey:WorkspaceID;constraint:ThemeTemplate_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Spaces                        []Space               `gorm:"foreignKey:WorkspaceID;constraint:Space_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ClaimableCustomPlan           *ClaimableCustomPlan  `gorm:"foreignKey:WorkspaceID;constraint:ClaimableCustomPlan_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type Space struct {
	ID          string     `gorm:"column:id;primaryKey;uniqueIndex:Space_id_workspaceId_key" json:"id"`
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime   `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Name        string     `gorm:"column:name;not null" json:"name"`
	Icon        *string    `gorm:"column:icon" json:"icon"`
	WorkspaceID string     `gorm:"column:workspaceId;uniqueIndex:Space_id_workspaceId_key;not null" json:"workspaceId"`
	Workspace   *Workspace `gorm:"belongsTo"`
	Typebots    []Typebot  `gorm:"foreignKey:SpaceID;constraint:Typebot_spaceId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Space) TableName() string {
	return "Space"
}

func (m *Space) BeforeCreate(tx *gorm.DB) error {
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

func (*Space) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type MemberInWorkspace struct {
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime   `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	UserID      string     `gorm:"column:userId;primaryKey;uniqueIndex:MemberInWorkspace_userId_workspaceId_key" json:"userId"`
	WorkspaceID string     `gorm:"column:workspaceId;primaryKey;uniqueIndex:MemberInWorkspace_userId_workspaceId_key" json:"workspaceId"`
	Role        string     `gorm:"column:role;not null" json:"role"`
	User        *User      `gorm:"belongsTo"`
	Workspace   *Workspace `gorm:"belongsTo"`
}

func (MemberInWorkspace) TableName() string {
	return "MemberInWorkspace"
}

func (m *MemberInWorkspace) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*MemberInWorkspace) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkspaceInvitation struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime   `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Email       string     `gorm:"column:email;not null" json:"email"`
	WorkspaceID string     `gorm:"column:workspaceId;not null" json:"workspaceId"`
	Type        string     `gorm:"column:type;not null" json:"type"`
	Workspace   *Workspace `gorm:"belongsTo"`
}

func (WorkspaceInvitation) TableName() string {
	return "WorkspaceInvitation"
}

func (m *WorkspaceInvitation) BeforeCreate(tx *gorm.DB) error {
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

func (*WorkspaceInvitation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CustomDomain struct {
	Name        string     `gorm:"column:name;primaryKey" json:"name"`
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	WorkspaceID string     `gorm:"column:workspaceId;not null" json:"workspaceId"`
	Workspace   *Workspace `gorm:"belongsTo"`
}

func (CustomDomain) TableName() string {
	return "CustomDomain"
}

func (m *CustomDomain) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Credentials struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	WorkspaceID string     `gorm:"column:workspaceId;not null" json:"workspaceId"`
	Data        string     `gorm:"column:data;not null" json:"data"`
	Name        string     `gorm:"column:name;not null" json:"name"`
	Type        string     `gorm:"column:type;not null" json:"type"`
	Iv          string     `gorm:"column:iv;not null" json:"iv"`
	Workspace   *Workspace `gorm:"belongsTo"`
}

func (Credentials) TableName() string {
	return "Credentials"
}

func (m *Credentials) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type UserCredentials struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UserID    string   `gorm:"column:userId;not null" json:"userId"`
	Data      string   `gorm:"column:data;type:text;not null" json:"data"`
	Name      string   `gorm:"column:name;not null" json:"name"`
	Type      string   `gorm:"column:type;not null" json:"type"`
	Iv        string   `gorm:"column:iv;not null" json:"iv"`
	User      *User    `gorm:"belongsTo"`
}

func (UserCredentials) TableName() string {
	return "UserCredentials"
}

func (m *UserCredentials) BeforeCreate(tx *gorm.DB) error {
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
	Identifier     string   `gorm:"column:identifier;uniqueIndex:VerificationToken_identifier_token_key;not null" json:"identifier"`
	Token          string   `gorm:"column:token;primaryKey;uniqueIndex:VerificationToken_token_key;uniqueIndex:VerificationToken_identifier_token_key" json:"token"`
	Value          *string  `gorm:"column:value" json:"value"`
	FailedAttempts int      `gorm:"column:failedAttempts;size:32;default:0;not null" json:"failedAttempts"`
	Expires        DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

type DashboardFolder struct {
	ID             string            `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime          `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Name           string            `gorm:"column:name;not null" json:"name"`
	ParentFolderID *string           `gorm:"column:parentFolderId" json:"parentFolderId"`
	WorkspaceID    string            `gorm:"column:workspaceId;not null" json:"workspaceId"`
	ParentFolder   *DashboardFolder  `gorm:"foreignKey:ParentFolderID;belongsTo"`
	Workspace      *Workspace        `gorm:"belongsTo"`
	ChildrenFolder []DashboardFolder `gorm:"foreignKey:ParentFolderID;constraint:DashboardFolder_parentFolderId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Typebots       []Typebot         `gorm:"foreignKey:FolderID;constraint:Typebot_folderId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (DashboardFolder) TableName() string {
	return "DashboardFolder"
}

func (m *DashboardFolder) BeforeCreate(tx *gorm.DB) error {
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

func (*DashboardFolder) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Typebot struct {
	ID                      string                    `gorm:"column:id;primaryKey" json:"id"`
	Version                 *string                   `gorm:"column:version" json:"version"`
	CreatedAt               DateTime                  `gorm:"column:createdAt;index:Typebot_isArchived_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime                  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Icon                    *string                   `gorm:"column:icon" json:"icon"`
	Name                    string                    `gorm:"column:name;not null" json:"name"`
	FolderID                *string                   `gorm:"column:folderId" json:"folderId"`
	Groups                  datatypes.JSON            `gorm:"column:groups;not null" json:"groups"`
	Events                  datatypes.JSON            `gorm:"column:events" json:"events"`
	Variables               datatypes.JSON            `gorm:"column:variables;not null" json:"variables"`
	Edges                   datatypes.JSON            `gorm:"column:edges;not null" json:"edges"`
	Theme                   datatypes.JSON            `gorm:"column:theme;not null" json:"theme"`
	SelectedThemeTemplateID *string                   `gorm:"column:selectedThemeTemplateId" json:"selectedThemeTemplateId"`
	Settings                datatypes.JSON            `gorm:"column:settings;not null" json:"settings"`
	PublicID                *string                   `gorm:"column:publicId;uniqueIndex:Typebot_publicId_key" json:"publicId"`
	CustomDomain            *string                   `gorm:"column:customDomain;uniqueIndex:Typebot_customDomain_key" json:"customDomain"`
	WorkspaceID             string                    `gorm:"column:workspaceId;index:Typebot_workspaceId_idx;not null" json:"workspaceId"`
	ResultsTablePreferences datatypes.JSON            `gorm:"column:resultsTablePreferences" json:"resultsTablePreferences"`
	IsArchived              bool                      `gorm:"column:isArchived;index:Typebot_isArchived_createdAt_idx;default:false;not null" json:"isArchived"`
	IsClosed                bool                      `gorm:"column:isClosed;default:false;not null" json:"isClosed"`
	WhatsAppCredentialsID   *string                   `gorm:"column:whatsAppCredentialsId" json:"whatsAppCredentialsId"`
	RiskLevel               *int                      `gorm:"column:riskLevel;size:32" json:"riskLevel"`
	SpaceID                 *string                   `gorm:"column:spaceId" json:"spaceId"`
	Folder                  *DashboardFolder          `gorm:"foreignKey:FolderID;belongsTo"`
	Workspace               *Workspace                `gorm:"belongsTo"`
	Space                   *Space                    `gorm:"belongsTo"`
	Collaborators           []CollaboratorsOnTypebots `gorm:"foreignKey:TypebotID;constraint:CollaboratorsOnTypebots_typebotId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invitations             []Invitation              `gorm:"foreignKey:TypebotID;constraint:Invitation_typebotId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Results                 []Result                  `gorm:"foreignKey:TypebotID;constraint:Result_typebotId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks                []Webhook                 `gorm:"foreignKey:TypebotID;constraint:Webhook_typebotId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BannedIps               []BannedIP                `gorm:"foreignKey:ResponsibleTypebotID;constraint:BannedIp_responsibleTypebotId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PublishedTypebot        *PublicTypebot            `gorm:"foreignKey:TypebotID;constraint:PublicTypebot_typebotId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Typebot) TableName() string {
	return "Typebot"
}

func (m *Typebot) BeforeCreate(tx *gorm.DB) error {
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

func (*Typebot) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Invitation struct {
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Email     string   `gorm:"column:email;primaryKey;uniqueIndex:Invitation_email_typebotId_key" json:"email"`
	TypebotID string   `gorm:"column:typebotId;primaryKey;uniqueIndex:Invitation_email_typebotId_key" json:"typebotId"`
	Type      string   `gorm:"column:type;not null" json:"type"`
	Typebot   *Typebot `gorm:"belongsTo"`
}

func (Invitation) TableName() string {
	return "Invitation"
}

func (m *Invitation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Invitation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CollaboratorsOnTypebots struct {
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	UserID    string   `gorm:"column:userId;primaryKey;uniqueIndex:CollaboratorsOnTypebots_userId_typebotId_key" json:"userId"`
	TypebotID string   `gorm:"column:typebotId;primaryKey;uniqueIndex:CollaboratorsOnTypebots_userId_typebotId_key" json:"typebotId"`
	Type      string   `gorm:"column:type;not null" json:"type"`
	Typebot   *Typebot `gorm:"belongsTo"`
	User      *User    `gorm:"belongsTo"`
}

func (CollaboratorsOnTypebots) TableName() string {
	return "CollaboratorsOnTypebots"
}

func (m *CollaboratorsOnTypebots) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CollaboratorsOnTypebots) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PublicTypebot struct {
	ID             string                `gorm:"column:id;primaryKey" json:"id"`
	Version        *string               `gorm:"column:version" json:"version"`
	CreatedAt      DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime              `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	LastActivityAt *DateTime             `gorm:"column:lastActivityAt;index:PublicTypebot_lastActivityAt_idx;type:timestamp(3)" json:"lastActivityAt"`
	TypebotID      string                `gorm:"column:typebotId;uniqueIndex:PublicTypebot_typebotId_key;not null" json:"typebotId"`
	Groups         datatypes.JSON        `gorm:"column:groups;not null" json:"groups"`
	Events         datatypes.JSON        `gorm:"column:events" json:"events"`
	Variables      datatypes.JSON        `gorm:"column:variables;not null" json:"variables"`
	Edges          datatypes.JSON        `gorm:"column:edges;not null" json:"edges"`
	Theme          datatypes.JSON        `gorm:"column:theme;not null" json:"theme"`
	Settings       datatypes.JSON        `gorm:"column:settings;not null" json:"settings"`
	Typebot        *Typebot              `gorm:"belongsTo"`
	MediaIdsCache  []RuntimeMediaIDCache `gorm:"foreignKey:PublicTypebotID;constraint:RuntimeMediaIdCache_publicTypebotId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (PublicTypebot) TableName() string {
	return "PublicTypebot"
}

func (m *PublicTypebot) BeforeCreate(tx *gorm.DB) error {
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

type RuntimeMediaIDCache struct {
	Provider        string         `gorm:"column:provider;primaryKey;uniqueIndex:RuntimeMediaIdCache_publicTypebotId_provider_url_key" json:"provider"`
	URL             string         `gorm:"column:url;primaryKey;uniqueIndex:RuntimeMediaIdCache_publicTypebotId_provider_url_key" json:"url"`
	MediaID         string         `gorm:"column:mediaId;not null" json:"mediaId"`
	ExpiresAt       *DateTime      `gorm:"column:expiresAt;index:RuntimeMediaIdCache_expiresAt_idx;type:timestamp(3)" json:"expiresAt"`
	PublicTypebotID string         `gorm:"column:publicTypebotId;primaryKey;uniqueIndex:RuntimeMediaIdCache_publicTypebotId_provider_url_key" json:"publicTypebotId"`
	PublicTypebot   *PublicTypebot `gorm:"belongsTo"`
}

func (RuntimeMediaIDCache) TableName() string {
	return "RuntimeMediaIdCache"
}

type Result struct {
	ID                 string                   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt          DateTime                 `gorm:"column:createdAt;index:Result_typebotId_hasStarted_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	TypebotID          string                   `gorm:"column:typebotId;index:Result_typebotId_hasStarted_createdAt_idx;index:Result_typebotId_isCompleted_idx;not null" json:"typebotId"`
	Variables          datatypes.JSON           `gorm:"column:variables;not null" json:"variables"`
	IsCompleted        bool                     `gorm:"column:isCompleted;index:Result_typebotId_isCompleted_idx;not null" json:"isCompleted"`
	HasStarted         *bool                    `gorm:"column:hasStarted;index:Result_typebotId_hasStarted_createdAt_idx" json:"hasStarted"`
	IsArchived         *bool                    `gorm:"column:isArchived;default:false" json:"isArchived"`
	LastChatSessionID  *string                  `gorm:"column:lastChatSessionId" json:"lastChatSessionId"`
	Typebot            *Typebot                 `gorm:"belongsTo"`
	Answers            []Answer                 `gorm:"foreignKey:ResultID;constraint:Answer_resultId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AnswersV2          []AnswerV2               `gorm:"foreignKey:ResultID;constraint:AnswerV2_resultId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Logs               []Log                    `gorm:"foreignKey:ResultID;constraint:Log_resultId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Edges              []VisitedEdge            `gorm:"foreignKey:ResultID;constraint:VisitedEdge_resultId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SetVariableHistory []SetVariableHistoryItem `gorm:"foreignKey:ResultID;constraint:SetVariableHistoryItem_resultId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Result) TableName() string {
	return "Result"
}

func (m *Result) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SetVariableHistoryItem struct {
	ResultID   string         `gorm:"column:resultId;primaryKey;uniqueIndex:SetVariableHistoryItem_resultId_index_key" json:"resultId"`
	Index      int            `gorm:"column:index;primaryKey;uniqueIndex:SetVariableHistoryItem_resultId_index_key;size:32" json:"index"`
	VariableID string         `gorm:"column:variableId;not null" json:"variableId"`
	BlockID    string         `gorm:"column:blockId;not null" json:"blockId"`
	BlockIndex *int           `gorm:"column:blockIndex;size:32" json:"blockIndex"`
	Value      datatypes.JSON `gorm:"column:value;not null" json:"value"`
	Result     *Result        `gorm:"belongsTo"`
}

func (SetVariableHistoryItem) TableName() string {
	return "SetVariableHistoryItem"
}

type VisitedEdge struct {
	ResultID string  `gorm:"column:resultId;primaryKey;uniqueIndex:VisitedEdge_resultId_index_key" json:"resultId"`
	EdgeID   string  `gorm:"column:edgeId;not null" json:"edgeId"`
	Index    int     `gorm:"column:index;primaryKey;uniqueIndex:VisitedEdge_resultId_index_key;size:32" json:"index"`
	Result   *Result `gorm:"belongsTo"`
}

func (VisitedEdge) TableName() string {
	return "VisitedEdge"
}

type Log struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ResultID    string   `gorm:"column:resultId;index:Log_resultId_idx;not null" json:"resultId"`
	Status      string   `gorm:"column:status;not null" json:"status"`
	Context     *string  `gorm:"column:context" json:"context"`
	Description string   `gorm:"column:description;not null" json:"description"`
	Details     *string  `gorm:"column:details" json:"details"`
	Result      *Result  `gorm:"belongsTo"`
}

func (Log) TableName() string {
	return "Log"
}

func (m *Log) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Answer struct {
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ResultID   string   `gorm:"column:resultId;primaryKey;uniqueIndex:Answer_resultId_blockId_groupId_key" json:"resultId"`
	BlockID    string   `gorm:"column:blockId;primaryKey;uniqueIndex:Answer_resultId_blockId_groupId_key" json:"blockId"`
	GroupID    string   `gorm:"column:groupId;primaryKey;uniqueIndex:Answer_resultId_blockId_groupId_key" json:"groupId"`
	VariableID *string  `gorm:"column:variableId" json:"variableId"`
	Content    string   `gorm:"column:content;not null" json:"content"`
	Result     *Result  `gorm:"belongsTo"`
}

func (Answer) TableName() string {
	return "Answer"
}

func (m *Answer) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Answer) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("CreatedAt") {
		tx.Statement.SetColumn("CreatedAt", DateTime{Time: now})
	}
	return nil
}

type AnswerV2 struct {
	ID               int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	CreatedAt        DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	BlockID          string         `gorm:"column:blockId;index:AnswerV2_blockId_idx;not null" json:"blockId"`
	Content          string         `gorm:"column:content;not null" json:"content"`
	AttachedFileUrls datatypes.JSON `gorm:"column:attachedFileUrls" json:"attachedFileUrls"`
	ResultID         string         `gorm:"column:resultId;not null" json:"resultId"`
	Result           *Result        `gorm:"belongsTo"`
}

func (AnswerV2) TableName() string {
	return "AnswerV2"
}

func (m *AnswerV2) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Coupon struct {
	UserPropertiesToUpdate datatypes.JSON `gorm:"column:userPropertiesToUpdate;not null" json:"userPropertiesToUpdate"`
	Code                   string         `gorm:"column:code;primaryKey" json:"code"`
	DateRedeemed           *DateTime      `gorm:"column:dateRedeemed;type:timestamp(3)" json:"dateRedeemed"`
}

func (Coupon) TableName() string {
	return "Coupon"
}

type Webhook struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	URL         *string        `gorm:"column:url" json:"url"`
	Method      string         `gorm:"column:method;not null" json:"method"`
	QueryParams datatypes.JSON `gorm:"column:queryParams;not null" json:"queryParams"`
	Headers     datatypes.JSON `gorm:"column:headers;not null" json:"headers"`
	Body        *string        `gorm:"column:body" json:"body"`
	TypebotID   string         `gorm:"column:typebotId;not null" json:"typebotId"`
	Typebot     *Typebot       `gorm:"belongsTo"`
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

type ClaimableCustomPlan struct {
	ID           string     `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt    DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ClaimedAt    *DateTime  `gorm:"column:claimedAt;type:timestamp(3)" json:"claimedAt"`
	Name         string     `gorm:"column:name;not null" json:"name"`
	Description  *string    `gorm:"column:description" json:"description"`
	Price        int        `gorm:"column:price;size:32;not null" json:"price"`
	Currency     string     `gorm:"column:currency;not null" json:"currency"`
	WorkspaceID  string     `gorm:"column:workspaceId;uniqueIndex:ClaimableCustomPlan_workspaceId_key;not null" json:"workspaceId"`
	ChatsLimit   int        `gorm:"column:chatsLimit;size:32;not null" json:"chatsLimit"`
	StorageLimit int        `gorm:"column:storageLimit;size:32;not null" json:"storageLimit"`
	SeatsLimit   int        `gorm:"column:seatsLimit;size:32;not null" json:"seatsLimit"`
	IsYearly     bool       `gorm:"column:isYearly;default:false;not null" json:"isYearly"`
	CompanyName  *string    `gorm:"column:companyName" json:"companyName"`
	VatType      *string    `gorm:"column:vatType" json:"vatType"`
	VatValue     *string    `gorm:"column:vatValue" json:"vatValue"`
	Workspace    *Workspace `gorm:"belongsTo"`
}

func (ClaimableCustomPlan) TableName() string {
	return "ClaimableCustomPlan"
}

func (m *ClaimableCustomPlan) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type ChatSession struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt  DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	State      datatypes.JSON `gorm:"column:state;not null" json:"state"`
	IsReplying *bool          `gorm:"column:isReplying" json:"isReplying"`
}

func (ChatSession) TableName() string {
	return "ChatSession"
}

func (m *ChatSession) BeforeCreate(tx *gorm.DB) error {
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

func (*ChatSession) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ThemeTemplate struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	Theme       datatypes.JSON `gorm:"column:theme;not null" json:"theme"`
	WorkspaceID string         `gorm:"column:workspaceId;not null" json:"workspaceId"`
	Workspace   *Workspace     `gorm:"belongsTo"`
}

func (ThemeTemplate) TableName() string {
	return "ThemeTemplate"
}

func (m *ThemeTemplate) BeforeCreate(tx *gorm.DB) error {
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

func (*ThemeTemplate) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BannedIP struct {
	ID                   string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt            DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	IP                   string   `gorm:"column:ip;uniqueIndex:BannedIp_ip_key;not null" json:"ip"`
	ResponsibleTypebotID string   `gorm:"column:responsibleTypebotId;not null" json:"responsibleTypebotId"`
	UserID               string   `gorm:"column:userId;not null" json:"userId"`
	ResponsibleTypebot   *Typebot `gorm:"foreignKey:ResponsibleTypebotID;belongsTo"`
	User                 *User    `gorm:"belongsTo"`
}

func (BannedIP) TableName() string {
	return "BannedIp"
}

func (m *BannedIP) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SuppressedEmail struct {
	ID                          string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                   DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                   DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Email                       string    `gorm:"column:email;uniqueIndex:SuppressedEmail_email_key;not null" json:"email"`
	TransientGeneralBounceCount int       `gorm:"column:transientGeneralBounceCount;size:32;default:0;not null" json:"transientGeneralBounceCount"`
	SuppressedAt                *DateTime `gorm:"column:suppressedAt;type:timestamp(3)" json:"suppressedAt"`
	LastWebhookID               *string   `gorm:"column:lastWebhookId" json:"lastWebhookId"`
}

func (SuppressedEmail) TableName() string {
	return "SuppressedEmail"
}

func (m *SuppressedEmail) BeforeCreate(tx *gorm.DB) error {
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

func (*SuppressedEmail) BeforeUpdate(tx *gorm.DB) error {
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
