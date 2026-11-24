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

type Customer struct {
	ID                     string    `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name                   *string   `gorm:"column:name;index:Customer_email_name_idx;type:varchar(191)" json:"name"`
	Email                  *string   `gorm:"column:email;index:Customer_projectId_email_idx;index:Customer_email_name_idx;type:varchar(191)" json:"email"`
	NormalizedEmail        *string   `gorm:"column:normalizedEmail;index:Customer_projectId_normalizedEmail_idx;type:varchar(191)" json:"normalizedEmail"`
	Avatar                 *string   `gorm:"column:avatar;type:text" json:"avatar"`
	ExternalID             *string   `gorm:"column:externalId;uniqueIndex:Customer_projectId_externalId_key;type:varchar(191)" json:"externalId"`
	StripeCustomerID       *string   `gorm:"column:stripeCustomerId;uniqueIndex:Customer_stripeCustomerId_key;type:varchar(191)" json:"stripeCustomerId"`
	LinkID                 *string   `gorm:"column:linkId;index:Customer_linkId_idx;type:varchar(191)" json:"linkId"`
	ClickID                *string   `gorm:"column:clickId;type:varchar(191)" json:"clickId"`
	ClickedAt              *DateTime `gorm:"column:clickedAt" json:"clickedAt"`
	Country                *string   `gorm:"column:country;index:Customer_country_idx;type:varchar(191)" json:"country"`
	Sales                  int       `gorm:"column:sales;size:32;default:0;not null" json:"sales"`
	SaleAmount             int64     `gorm:"column:saleAmount;index:Customer_projectId_saleAmount_idx;default:0;not null" json:"saleAmount"`
	FirstSaleAt            *DateTime `gorm:"column:firstSaleAt;index:Customer_projectId_firstSaleAt_idx" json:"firstSaleAt"`
	SubscriptionCanceledAt *DateTime `gorm:"column:subscriptionCanceledAt;index:Customer_projectId_subscriptionCanceledAt_idx" json:"subscriptionCanceledAt"`
	ProjectID              string    `gorm:"column:projectId;uniqueIndex:Customer_projectId_externalId_key;index:Customer_projectId_email_idx;index:Customer_projectId_normalizedEmail_idx;index:Customer_projectId_createdAt_idx;index:Customer_projectId_saleAmount_idx;index:Customer_projectId_firstSaleAt_idx;index:Customer_projectId_subscriptionCanceledAt_idx;type:varchar(191);not null" json:"projectId"`
	ProjectConnectID       *string   `gorm:"column:projectConnectId;type:varchar(191)" json:"projectConnectId"`
	CreatedAt              DateTime  `gorm:"column:createdAt;index:Customer_projectId_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt              DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project                *Project  `gorm:"belongsTo"`
	Link                   *Link     `gorm:"belongsTo"`
}

func (Customer) TableName() string {
	return "Customer"
}

func (m *Customer) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Customer) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Dashboard struct {
	ID              string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	LinkID          *string  `gorm:"column:linkId;uniqueIndex:Dashboard_linkId_key;type:varchar(191)" json:"linkId"`
	FolderID        *string  `gorm:"column:folderId;uniqueIndex:Dashboard_folderId_key;type:varchar(191)" json:"folderId"`
	ProjectID       *string  `gorm:"column:projectId;index:Dashboard_projectId_idx;type:varchar(191)" json:"projectId"`
	UserID          *string  `gorm:"column:userId;index:Dashboard_userId_idx;type:varchar(191)" json:"userId"`
	DoIndex         bool     `gorm:"column:doIndex;default:false;not null" json:"doIndex"`
	Password        *string  `gorm:"column:password;type:varchar(191)" json:"password"`
	ShowConversions bool     `gorm:"column:showConversions;default:false;not null" json:"showConversions"`
	CreatedAt       DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Link            *Link    `gorm:"belongsTo"`
	Folder          *Folder  `gorm:"belongsTo"`
	Project         *Project `gorm:"belongsTo"`
	User            *User    `gorm:"belongsTo"`
}

func (Dashboard) TableName() string {
	return "Dashboard"
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

type Domain struct {
	ID                      string            `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Slug                    string            `gorm:"column:slug;uniqueIndex:Domain_slug_key;type:varchar(191);not null" json:"slug"`
	Verified                bool              `gorm:"column:verified;default:false;not null" json:"verified"`
	Placeholder             *string           `gorm:"column:placeholder;type:varchar(191)" json:"placeholder"`
	ExpiredURL              *string           `gorm:"column:expiredUrl;type:text" json:"expiredUrl"`
	NotFoundURL             *string           `gorm:"column:notFoundUrl;type:text" json:"notFoundUrl"`
	Primary                 bool              `gorm:"column:primary;default:false;not null" json:"primary"`
	Archived                bool              `gorm:"column:archived;default:false;not null" json:"archived"`
	LastChecked             DateTime          `gorm:"column:lastChecked;index:Domain_lastChecked_idx;not null" json:"lastChecked"`
	Logo                    *string           `gorm:"column:logo;type:varchar(191)" json:"logo"`
	AppleAppSiteAssociation datatypes.JSON    `gorm:"column:appleAppSiteAssociation" json:"appleAppSiteAssociation"`
	AssetLinks              datatypes.JSON    `gorm:"column:assetLinks" json:"assetLinks"`
	DeepviewData            datatypes.JSON    `gorm:"column:deepviewData;default:'{}'" json:"deepviewData"`
	LinkRetentionDays       *int              `gorm:"column:linkRetentionDays;size:32" json:"linkRetentionDays"`
	CreatedAt               DateTime          `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime          `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	ProjectID               *string           `gorm:"column:projectId;index:Domain_projectId_idx;type:varchar(191)" json:"projectId"`
	Project                 *Project          `gorm:"belongsTo"`
	Links                   []Link            `gorm:"foreignKey:Domain;references:Slug;constraint:Link_domain_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	RegisteredDomain        *RegisteredDomain `gorm:"foreignKey:DomainID;constraint:RegisteredDomain_domainId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Domain) TableName() string {
	return "Domain"
}

func (m *Domain) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.LastChecked.IsZero() {
		m.LastChecked = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Domain) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RegisteredDomain struct {
	ID                    string    `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Slug                  string    `gorm:"column:slug;type:varchar(191);not null" json:"slug"`
	ProjectID             string    `gorm:"column:projectId;index:RegisteredDomain_projectId_idx;type:varchar(191);not null" json:"projectId"`
	DomainID              *string   `gorm:"column:domainId;uniqueIndex:RegisteredDomain_domainId_key;type:varchar(191)" json:"domainId"`
	AutoRenewalDisabledAt *DateTime `gorm:"column:autoRenewalDisabledAt" json:"autoRenewalDisabledAt"`
	RenewalFee            *int      `gorm:"column:renewalFee;size:32;default:1200;not null" json:"renewalFee"`
	ExpiresAt             DateTime  `gorm:"column:expiresAt;index:RegisteredDomain_expiresAt_idx;not null" json:"expiresAt"`
	CreatedAt             DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project               *Project  `gorm:"belongsTo"`
	Domain                *Domain   `gorm:"belongsTo"`
}

func (RegisteredDomain) TableName() string {
	return "RegisteredDomain"
}

func (m *RegisteredDomain) BeforeCreate(tx *gorm.DB) error {
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

func (*RegisteredDomain) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DefaultDomains struct {
	ID          string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Dubsh       *bool    `gorm:"column:dubsh;default:true;not null" json:"dubsh"`
	Chatgpt     *bool    `gorm:"column:chatgpt;default:true;not null" json:"chatgpt"`
	Sptifi      *bool    `gorm:"column:sptifi;default:true;not null" json:"sptifi"`
	Gitnew      *bool    `gorm:"column:gitnew;default:true;not null" json:"gitnew"`
	Callink     *bool    `gorm:"column:callink;default:true;not null" json:"callink"`
	Amznid      *bool    `gorm:"column:amznid;default:true;not null" json:"amznid"`
	Ggllink     *bool    `gorm:"column:ggllink;default:true;not null" json:"ggllink"`
	Figpage     *bool    `gorm:"column:figpage;default:true;not null" json:"figpage"`
	Loooooooong bool     `gorm:"column:loooooooong;default:false;not null" json:"loooooooong"`
	ProjectID   string   `gorm:"column:projectId;uniqueIndex:DefaultDomains_projectId_key;type:varchar(191);not null" json:"projectId"`
	Project     *Project `gorm:"belongsTo"`
}

func (DefaultDomains) TableName() string {
	return "DefaultDomains"
}

func (m *DefaultDomains) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type Folder struct {
	ID             string                `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name           string                `gorm:"column:name;uniqueIndex:Folder_name_projectId_key;type:varchar(191);not null" json:"name"`
	Description    *string               `gorm:"column:description;type:text" json:"description"`
	ProjectID      string                `gorm:"column:projectId;uniqueIndex:Folder_name_projectId_key;index:Folder_projectId_idx;type:varchar(191);not null" json:"projectId"`
	Type           string                `gorm:"column:type;default:'default';not null" json:"type"`
	AccessLevel    *string               `gorm:"column:accessLevel" json:"accessLevel"`
	CreatedAt      DateTime              `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime              `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project        *Project              `gorm:"belongsTo"`
	Links          []Link                `gorm:"foreignKey:FolderID;constraint:Link_folderId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Users          []FolderUser          `gorm:"foreignKey:FolderID;constraint:FolderUser_folderId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessRequests []FolderAccessRequest `gorm:"foreignKey:FolderID;constraint:FolderAccessRequest_folderId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks       []FolderWebhook       `gorm:"foreignKey:FolderID;constraint:FolderWebhook_folderId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Dashboard      *Dashboard            `gorm:"foreignKey:FolderID;constraint:Dashboard_folderId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Folder) TableName() string {
	return "Folder"
}

func (m *Folder) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type FolderUser struct {
	ID        string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	FolderID  string   `gorm:"column:folderId;uniqueIndex:FolderUser_folderId_userId_key;type:varchar(191);not null" json:"folderId"`
	UserID    string   `gorm:"column:userId;uniqueIndex:FolderUser_folderId_userId_key;index:FolderUser_userId_idx;type:varchar(191);not null" json:"userId"`
	Role      *string  `gorm:"column:role" json:"role"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Folder    *Folder  `gorm:"belongsTo"`
	User      *User    `gorm:"belongsTo"`
}

func (FolderUser) TableName() string {
	return "FolderUser"
}

func (m *FolderUser) BeforeCreate(tx *gorm.DB) error {
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

func (*FolderUser) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type FolderAccessRequest struct {
	ID        string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	FolderID  string   `gorm:"column:folderId;uniqueIndex:FolderAccessRequest_folderId_userId_key;type:varchar(191);not null" json:"folderId"`
	UserID    string   `gorm:"column:userId;uniqueIndex:FolderAccessRequest_folderId_userId_key;index:FolderAccessRequest_userId_idx;type:varchar(191);not null" json:"userId"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Folder    *Folder  `gorm:"belongsTo"`
	User      *User    `gorm:"belongsTo"`
}

func (FolderAccessRequest) TableName() string {
	return "FolderAccessRequest"
}

func (m *FolderAccessRequest) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Integration struct {
	ID            string                 `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	UserID        *string                `gorm:"column:userId;index:Integration_userId_idx;type:varchar(191)" json:"userId"`
	ProjectID     string                 `gorm:"column:projectId;index:Integration_projectId_idx;type:varchar(191);not null" json:"projectId"`
	Name          string                 `gorm:"column:name;type:varchar(191);not null" json:"name"`
	Slug          string                 `gorm:"column:slug;uniqueIndex:Integration_slug_key;type:varchar(191);not null" json:"slug"`
	Description   *string                `gorm:"column:description;type:varchar(191)" json:"description"`
	Readme        *string                `gorm:"column:readme;type:text" json:"readme"`
	Developer     string                 `gorm:"column:developer;type:varchar(191);not null" json:"developer"`
	Website       string                 `gorm:"column:website;type:varchar(191);not null" json:"website"`
	Logo          *string                `gorm:"column:logo;type:varchar(191)" json:"logo"`
	Screenshots   datatypes.JSON         `gorm:"column:screenshots" json:"screenshots"`
	Verified      bool                   `gorm:"column:verified;default:false;not null" json:"verified"`
	InstallURL    *string                `gorm:"column:installUrl;type:text" json:"installUrl"`
	GuideURL      *string                `gorm:"column:guideUrl;type:text" json:"guideUrl"`
	Category      *string                `gorm:"column:category;type:varchar(191)" json:"category"`
	ComingSoon    bool                   `gorm:"column:comingSoon;default:false;not null" json:"comingSoon"`
	CreatedAt     DateTime               `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime               `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	User          *User                  `gorm:"belongsTo"`
	Project       *Project               `gorm:"belongsTo"`
	Installations []InstalledIntegration `gorm:"foreignKey:IntegrationID;constraint:InstalledIntegration_integrationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OAuthApp      *OAuthApp              `gorm:"foreignKey:IntegrationID;constraint:OAuthApp_integrationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Integration) TableName() string {
	return "Integration"
}

func (m *Integration) BeforeCreate(tx *gorm.DB) error {
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

func (*Integration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type InstalledIntegration struct {
	ID            string              `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	UserID        string              `gorm:"column:userId;uniqueIndex:InstalledIntegration_userId_integrationId_projectId_key;type:varchar(191);not null" json:"userId"`
	IntegrationID string              `gorm:"column:integrationId;uniqueIndex:InstalledIntegration_userId_integrationId_projectId_key;index:InstalledIntegration_integrationId_idx;type:varchar(191);not null" json:"integrationId"`
	ProjectID     string              `gorm:"column:projectId;uniqueIndex:InstalledIntegration_userId_integrationId_projectId_key;index:InstalledIntegration_projectId_idx;type:varchar(191);not null" json:"projectId"`
	CreatedAt     DateTime            `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime            `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Credentials   datatypes.JSON      `gorm:"column:credentials" json:"credentials"`
	Settings      datatypes.JSON      `gorm:"column:settings" json:"settings"`
	User          *User               `gorm:"belongsTo"`
	Integration   *Integration        `gorm:"belongsTo"`
	Project       *Project            `gorm:"belongsTo"`
	RefreshTokens []OAuthRefreshToken `gorm:"foreignKey:InstallationID;constraint:OAuthRefreshToken_installationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessTokens  []RestrictedToken   `gorm:"foreignKey:InstallationID;constraint:RestrictedToken_installationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks      []Webhook           `gorm:"foreignKey:InstallationID;constraint:Webhook_installationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (InstalledIntegration) TableName() string {
	return "InstalledIntegration"
}

func (m *InstalledIntegration) BeforeCreate(tx *gorm.DB) error {
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

func (*InstalledIntegration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Invoice struct {
	ID                   string         `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	WorkspaceID          string         `gorm:"column:workspaceId;index:Invoice_workspaceId_idx;type:varchar(191);not null" json:"workspaceId"`
	Number               *string        `gorm:"column:number;uniqueIndex:Invoice_number_key;type:varchar(191)" json:"number"`
	Status               string         `gorm:"column:status;default:'processing';not null" json:"status"`
	Type                 string         `gorm:"column:type;default:'domainRenewal';not null" json:"type"`
	PaymentMethod        *string        `gorm:"column:paymentMethod" json:"paymentMethod"`
	Amount               int            `gorm:"column:amount;size:32;default:0;not null" json:"amount"`
	Fee                  int            `gorm:"column:fee;size:32;default:0;not null" json:"fee"`
	RefundedAmount       int            `gorm:"column:refundedAmount;size:32;default:0;not null" json:"refundedAmount"`
	Total                int            `gorm:"column:total;size:32;default:0;not null" json:"total"`
	ReceiptURL           *string        `gorm:"column:receiptUrl;type:text" json:"receiptUrl"`
	RefundReason         *string        `gorm:"column:refundReason;type:text" json:"refundReason"`
	FailedReason         *string        `gorm:"column:failedReason;type:text" json:"failedReason"`
	RegisteredDomains    datatypes.JSON `gorm:"column:registeredDomains" json:"registeredDomains"`
	StripeChargeMetadata datatypes.JSON `gorm:"column:stripeChargeMetadata" json:"stripeChargeMetadata"`
	FailedAttempts       int            `gorm:"column:failedAttempts;size:32;default:0;not null" json:"failedAttempts"`
	CreatedAt            DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	PaidAt               *DateTime      `gorm:"column:paidAt" json:"paidAt"`
	Workspace            *Project       `gorm:"foreignKey:WorkspaceID;belongsTo"`
}

func (Invoice) TableName() string {
	return "Invoice"
}

func (m *Invoice) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type JacksonIndex struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Key      string `gorm:"column:key;index:_jackson_index_key_store;type:varchar(250);not null" json:"key"`
	StoreKey string `gorm:"column:storeKey;index:_jackson_index_key_store;type:varchar(250);not null" json:"storeKey"`
}

func (JacksonIndex) TableName() string {
	return "jackson_index"
}

type JacksonStore struct {
	Key        string    `gorm:"column:key;primaryKey;type:varchar(250)" json:"key"`
	Value      string    `gorm:"column:value;type:text;not null" json:"value"`
	Iv         *string   `gorm:"column:iv;type:varchar(64)" json:"iv"`
	Tag        *string   `gorm:"column:tag;type:varchar(64)" json:"tag"`
	Namespace  *string   `gorm:"column:namespace;index:_jackson_store_namespace;type:varchar(64)" json:"namespace"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(0);autoCreateTime:false;not null" json:"createdAt"`
	ModifiedAt *DateTime `gorm:"column:modifiedAt;type:timestamp(0)" json:"modifiedAt"`
}

func (JacksonStore) TableName() string {
	return "jackson_store"
}

func (m *JacksonStore) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type JacksonTTL struct {
	Key       string `gorm:"column:key;primaryKey;type:varchar(250)" json:"key"`
	ExpiresAt int64  `gorm:"column:expiresAt;index:_jackson_ttl_expires_at;not null" json:"expiresAt"`
}

func (JacksonTTL) TableName() string {
	return "jackson_ttl"
}

type Job struct {
	ID          string         `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name        string         `gorm:"column:name;type:varchar(191);not null" json:"name"`
	Payload     datatypes.JSON `gorm:"column:payload;not null" json:"payload"`
	Options     datatypes.JSON `gorm:"column:options" json:"options"`
	ScheduledAt DateTime       `gorm:"column:scheduledAt;index:Job_scheduledAt_attempts_idx;not null" json:"scheduledAt"`
	Attempts    int            `gorm:"column:attempts;index:Job_scheduledAt_attempts_idx;size:32;default:0;not null" json:"attempts"`
	LastError   *string        `gorm:"column:lastError;type:text" json:"lastError"`
	CreatedAt   DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Job) TableName() string {
	return "Job"
}

func (m *Job) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ScheduledAt.IsZero() {
		m.ScheduledAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Job) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Link struct {
	ID                             string         `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Domain                         string         `gorm:"column:domain;uniqueIndex:Link_domain_key_key;index:Link_domain_createdAt_idx;type:varchar(191);not null" json:"domain"`
	Key                            string         `gorm:"column:key;uniqueIndex:Link_domain_key_key;type:varchar(191);not null" json:"key"`
	URL                            string         `gorm:"column:url;index:Link_projectId_url_idx;type:text;not null" json:"url"`
	ShortLink                      string         `gorm:"column:shortLink;uniqueIndex:Link_shortLink_key;type:varchar(400);not null" json:"shortLink"`
	Archived                       bool           `gorm:"column:archived;index:Link_projectId_folderId_archived_createdAt_idx;default:false;not null" json:"archived"`
	ExpiresAt                      *DateTime      `gorm:"column:expiresAt" json:"expiresAt"`
	ExpiredURL                     *string        `gorm:"column:expiredUrl;type:text" json:"expiredUrl"`
	DisabledAt                     *DateTime      `gorm:"column:disabledAt" json:"disabledAt"`
	Password                       *string        `gorm:"column:password;type:varchar(191)" json:"password"`
	TrackConversion                bool           `gorm:"column:trackConversion;default:false;not null" json:"trackConversion"`
	Proxy                          bool           `gorm:"column:proxy;default:false;not null" json:"proxy"`
	Title                          *string        `gorm:"column:title;type:varchar(191)" json:"title"`
	Description                    *string        `gorm:"column:description;type:varchar(280)" json:"description"`
	Image                          *string        `gorm:"column:image;type:text" json:"image"`
	Video                          *string        `gorm:"column:video;type:text" json:"video"`
	UtmSource                      *string        `gorm:"column:utm_source;type:varchar(255)" json:"utm_source"`
	UtmMedium                      *string        `gorm:"column:utm_medium;type:varchar(255)" json:"utm_medium"`
	UtmCampaign                    *string        `gorm:"column:utm_campaign;type:varchar(255)" json:"utm_campaign"`
	UtmTerm                        *string        `gorm:"column:utm_term;type:varchar(255)" json:"utm_term"`
	UtmContent                     *string        `gorm:"column:utm_content;type:varchar(255)" json:"utm_content"`
	Rewrite                        bool           `gorm:"column:rewrite;default:false;not null" json:"rewrite"`
	LinkRetentionCleanupDisabledAt *DateTime      `gorm:"column:linkRetentionCleanupDisabledAt" json:"linkRetentionCleanupDisabledAt"`
	DoIndex                        bool           `gorm:"column:doIndex;default:false;not null" json:"doIndex"`
	Ios                            *string        `gorm:"column:ios;type:text" json:"ios"`
	Android                        *string        `gorm:"column:android;type:text" json:"android"`
	Geo                            datatypes.JSON `gorm:"column:geo" json:"geo"`
	TestVariants                   datatypes.JSON `gorm:"column:testVariants" json:"testVariants"`
	TestStartedAt                  *DateTime      `gorm:"column:testStartedAt" json:"testStartedAt"`
	TestCompletedAt                *DateTime      `gorm:"column:testCompletedAt" json:"testCompletedAt"`
	UserID                         *string        `gorm:"column:userId;index:Link_userId_idx;type:varchar(191)" json:"userId"`
	ProjectID                      *string        `gorm:"column:projectId;uniqueIndex:Link_projectId_externalId_key;index:Link_projectId_tenantId_idx;index:Link_projectId_url_idx;index:Link_projectId_folderId_archived_createdAt_idx;type:varchar(191)" json:"projectId"`
	FolderID                       *string        `gorm:"column:folderId;index:Link_projectId_folderId_archived_createdAt_idx;index:Link_folderId_idx;type:varchar(191)" json:"folderId"`
	ExternalID                     *string        `gorm:"column:externalId;uniqueIndex:Link_projectId_externalId_key;type:varchar(191)" json:"externalId"`
	TenantID                       *string        `gorm:"column:tenantId;index:Link_projectId_tenantId_idx;type:varchar(191)" json:"tenantId"`
	PublicStats                    bool           `gorm:"column:publicStats;default:false;not null" json:"publicStats"`
	Clicks                         int            `gorm:"column:clicks;size:32;default:0;not null" json:"clicks"`
	Leads                          int            `gorm:"column:leads;size:32;default:0;not null" json:"leads"`
	Conversions                    int            `gorm:"column:conversions;size:32;default:0;not null" json:"conversions"`
	Sales                          int            `gorm:"column:sales;size:32;default:0;not null" json:"sales"`
	SaleAmount                     int64          `gorm:"column:saleAmount;default:0;not null" json:"saleAmount"`
	LastClicked                    *DateTime      `gorm:"column:lastClicked" json:"lastClicked"`
	LastLeadAt                     *DateTime      `gorm:"column:lastLeadAt" json:"lastLeadAt"`
	LastConversionAt               *DateTime      `gorm:"column:lastConversionAt" json:"lastConversionAt"`
	CreatedAt                      DateTime       `gorm:"column:createdAt;index:Link_projectId_folderId_archived_createdAt_idx;index:Link_domain_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                      DateTime       `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Comments                       *string        `gorm:"column:comments;type:text" json:"comments"`
	User                           *User          `gorm:"belongsTo"`
	Project                        *Project       `gorm:"belongsTo"`
	Folder                         *Folder        `gorm:"belongsTo"`
	ShortDomain                    *Domain        `gorm:"foreignKey:Domain;references:Slug;belongsTo"`
	Tags                           []LinkTag      `gorm:"foreignKey:LinkID;constraint:LinkTag_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks                       []LinkWebhook  `gorm:"foreignKey:LinkID;constraint:LinkWebhook_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Customers                      []Customer     `gorm:"foreignKey:LinkID;constraint:Customer_linkId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Dashboard                      *Dashboard     `gorm:"foreignKey:LinkID;constraint:Dashboard_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Link) TableName() string {
	return "Link"
}

func (m *Link) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Link) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type YearInReview struct {
	ID           string         `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Year         int            `gorm:"column:year;size:32;not null" json:"year"`
	TotalLinks   int            `gorm:"column:totalLinks;size:32;not null" json:"totalLinks"`
	TotalClicks  int            `gorm:"column:totalClicks;size:32;not null" json:"totalClicks"`
	TopLinks     datatypes.JSON `gorm:"column:topLinks;not null" json:"topLinks"`
	TopCountries datatypes.JSON `gorm:"column:topCountries;not null" json:"topCountries"`
	WorkspaceID  string         `gorm:"column:workspaceId;index:YearInReview_workspaceId_idx;type:varchar(191);not null" json:"workspaceId"`
	CreatedAt    DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	SentAt       *DateTime      `gorm:"column:sentAt" json:"sentAt"`
	Workspace    *Project       `gorm:"foreignKey:WorkspaceID;belongsTo"`
}

func (YearInReview) TableName() string {
	return "YearInReview"
}

func (m *YearInReview) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type UserNotificationPreferences struct {
	ID       string `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	UserID   string `gorm:"column:userId;uniqueIndex:UserNotificationPreferences_userId_key;type:varchar(191);not null" json:"userId"`
	DubLinks *bool  `gorm:"column:dubLinks;default:true;not null" json:"dubLinks"`
	User     *User  `gorm:"belongsTo"`
}

func (UserNotificationPreferences) TableName() string {
	return "UserNotificationPreferences"
}

func (m *UserNotificationPreferences) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type NotificationPreference struct {
	ID                         string        `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	ProjectUserID              string        `gorm:"column:projectUserId;uniqueIndex:NotificationPreference_projectUserId_key;type:varchar(191);not null" json:"projectUserId"`
	LinkUsageSummary           *bool         `gorm:"column:linkUsageSummary;default:true;not null" json:"linkUsageSummary"`
	DomainConfigurationUpdates *bool         `gorm:"column:domainConfigurationUpdates;default:true;not null" json:"domainConfigurationUpdates"`
	ProjectUser                *ProjectUsers `gorm:"foreignKey:ProjectUserID;belongsTo"`
}

func (NotificationPreference) TableName() string {
	return "NotificationPreference"
}

func (m *NotificationPreference) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type OAuthApp struct {
	ID                  string         `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	IntegrationID       string         `gorm:"column:integrationId;uniqueIndex:OAuthApp_integrationId_key;type:varchar(191);not null" json:"integrationId"`
	ClientID            string         `gorm:"column:clientId;uniqueIndex:OAuthApp_clientId_key;type:varchar(191);not null" json:"clientId"`
	HashedClientSecret  string         `gorm:"column:hashedClientSecret;type:varchar(191);not null" json:"hashedClientSecret"`
	PartialClientSecret string         `gorm:"column:partialClientSecret;type:varchar(191);not null" json:"partialClientSecret"`
	RedirectUris        datatypes.JSON `gorm:"column:redirectUris;not null" json:"redirectUris"`
	Pkce                bool           `gorm:"column:pkce;default:false;not null" json:"pkce"`
	Integration         *Integration   `gorm:"belongsTo"`
	OAuthCodes          []OAuthCode    `gorm:"foreignKey:ClientID;references:ClientID;constraint:OAuthCode_clientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OAuthApp) TableName() string {
	return "OAuthApp"
}

type OAuthCode struct {
	ID                  string    `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	ClientID            string    `gorm:"column:clientId;index:OAuthCode_clientId_idx;type:varchar(191);not null" json:"clientId"`
	UserID              string    `gorm:"column:userId;index:OAuthCode_userId_idx;type:varchar(191);not null" json:"userId"`
	ProjectID           string    `gorm:"column:projectId;index:OAuthCode_projectId_idx;type:varchar(191);not null" json:"projectId"`
	Code                string    `gorm:"column:code;uniqueIndex:OAuthCode_code_key;type:varchar(191);not null" json:"code"`
	Scopes              *string   `gorm:"column:scopes;type:text" json:"scopes"`
	RedirectURI         string    `gorm:"column:redirectUri;type:varchar(191);not null" json:"redirectUri"`
	CodeChallenge       *string   `gorm:"column:codeChallenge;type:varchar(191)" json:"codeChallenge"`
	CodeChallengeMethod *string   `gorm:"column:codeChallengeMethod;type:varchar(191)" json:"codeChallengeMethod"`
	ExpiresAt           DateTime  `gorm:"column:expiresAt;not null" json:"expiresAt"`
	CreatedAt           DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	OAuthApp            *OAuthApp `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
	User                *User     `gorm:"belongsTo"`
	Project             *Project  `gorm:"belongsTo"`
}

func (OAuthCode) TableName() string {
	return "OAuthCode"
}

func (m *OAuthCode) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type OAuthRefreshToken struct {
	ID                   string                `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	InstallationID       string                `gorm:"column:installationId;index:OAuthRefreshToken_installationId_idx;type:varchar(191);not null" json:"installationId"`
	AccessTokenID        string                `gorm:"column:accessTokenId;index:OAuthRefreshToken_accessTokenId_idx;type:varchar(191);not null" json:"accessTokenId"`
	HashedRefreshToken   string                `gorm:"column:hashedRefreshToken;uniqueIndex:OAuthRefreshToken_hashedRefreshToken_key;type:varchar(191);not null" json:"hashedRefreshToken"`
	ExpiresAt            DateTime              `gorm:"column:expiresAt;not null" json:"expiresAt"`
	CreatedAt            DateTime              `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	AccessToken          *RestrictedToken      `gorm:"foreignKey:AccessTokenID;belongsTo"`
	InstalledIntegration *InstalledIntegration `gorm:"foreignKey:InstallationID;belongsTo"`
}

func (OAuthRefreshToken) TableName() string {
	return "OAuthRefreshToken"
}

func (m *OAuthRefreshToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type User struct {
	ID                      string                       `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name                    *string                      `gorm:"column:name;type:varchar(191)" json:"name"`
	Email                   *string                      `gorm:"column:email;uniqueIndex:User_email_key;type:varchar(191)" json:"email"`
	EmailVerified           *DateTime                    `gorm:"column:emailVerified" json:"emailVerified"`
	Image                   *string                      `gorm:"column:image;type:varchar(191)" json:"image"`
	IsMachine               bool                         `gorm:"column:isMachine;default:false;not null" json:"isMachine"`
	CreatedAt               DateTime                     `gorm:"column:createdAt;index:User_createdAt_idx;autoCreateTime:false;not null" json:"createdAt"`
	PasswordHash            *string                      `gorm:"column:passwordHash;type:varchar(191)" json:"passwordHash"`
	InvalidLoginAttempts    int                          `gorm:"column:invalidLoginAttempts;size:32;default:0;not null" json:"invalidLoginAttempts"`
	LockedAt                *DateTime                    `gorm:"column:lockedAt" json:"lockedAt"`
	DefaultWorkspace        *string                      `gorm:"column:defaultWorkspace;index:User_defaultWorkspace_idx;type:varchar(191)" json:"defaultWorkspace"`
	Source                  *string                      `gorm:"column:source;index:User_source_idx;type:varchar(191)" json:"source"`
	SentMail                bool                         `gorm:"column:sentMail;index:User_sentMail_idx;default:false;not null" json:"sentMail"`
	Accounts                []Account                    `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions                []Session                    `gorm:"foreignKey:UserID;constraint:Session_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Projects                []ProjectUsers               `gorm:"foreignKey:UserID;constraint:ProjectUsers_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Links                   []Link                       `gorm:"foreignKey:UserID;constraint:Link_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Dashboards              []Dashboard                  `gorm:"foreignKey:UserID;constraint:Dashboard_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tokens                  []Token                      `gorm:"foreignKey:UserID;constraint:Token_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RestrictedTokens        []RestrictedToken            `gorm:"foreignKey:UserID;constraint:RestrictedToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OAuthCodes              []OAuthCode                  `gorm:"foreignKey:UserID;constraint:OAuthCode_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Integrations            []Integration                `gorm:"foreignKey:UserID;constraint:Integration_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	InstalledIntegrations   []InstalledIntegration       `gorm:"foreignKey:UserID;constraint:InstalledIntegration_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Folders                 []FolderUser                 `gorm:"foreignKey:UserID;constraint:FolderUser_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FolderAccessRequests    []FolderAccessRequest        `gorm:"foreignKey:UserID;constraint:FolderAccessRequest_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UtmTemplates            []UtmTemplate                `gorm:"foreignKey:UserID;constraint:UtmTemplate_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	NotificationPreferences *UserNotificationPreferences `gorm:"foreignKey:UserID;constraint:UserNotificationPreferences_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	return nil
}

type Account struct {
	ID                    string  `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	UserID                string  `gorm:"column:userId;index:Account_userId_idx;type:varchar(191);not null" json:"userId"`
	Type                  string  `gorm:"column:type;type:varchar(191);not null" json:"type"`
	Provider              string  `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key;type:varchar(191);not null" json:"provider"`
	ProviderAccountID     string  `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key;type:varchar(191);not null" json:"providerAccountId"`
	RefreshToken          *string `gorm:"column:refresh_token;type:text" json:"refresh_token"`
	RefreshTokenExpiresIn *int    `gorm:"column:refresh_token_expires_in;size:32" json:"refresh_token_expires_in"`
	AccessToken           *string `gorm:"column:access_token;type:text" json:"access_token"`
	ExpiresAt             *int    `gorm:"column:expires_at;size:32" json:"expires_at"`
	TokenType             *string `gorm:"column:token_type;type:varchar(191)" json:"token_type"`
	Scope                 *string `gorm:"column:scope;type:varchar(191)" json:"scope"`
	IDToken               *string `gorm:"column:id_token;type:text" json:"id_token"`
	SessionState          *string `gorm:"column:session_state;type:varchar(191)" json:"session_state"`
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
	ID           string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	SessionToken string   `gorm:"column:sessionToken;uniqueIndex:Session_sessionToken_key;type:varchar(191);not null" json:"sessionToken"`
	UserID       string   `gorm:"column:userId;index:Session_userId_idx;type:varchar(191);not null" json:"userId"`
	Expires      DateTime `gorm:"column:expires;not null" json:"expires"`
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

type Tag struct {
	ID        string    `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name      string    `gorm:"column:name;uniqueIndex:Tag_name_projectId_key;type:varchar(191);not null" json:"name"`
	Color     *string   `gorm:"column:color;type:varchar(191);default:'blue';not null" json:"color"`
	CreatedAt DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	ProjectID *string   `gorm:"column:projectId;uniqueIndex:Tag_name_projectId_key;index:Tag_projectId_idx;type:varchar(191)" json:"projectId"`
	Project   *Project  `gorm:"belongsTo"`
	Links     []LinkTag `gorm:"foreignKey:TagID;constraint:LinkTag_tagId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Tag) TableName() string {
	return "Tag"
}

func (m *Tag) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Tag) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type LinkTag struct {
	ID        string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	LinkID    string   `gorm:"column:linkId;uniqueIndex:LinkTag_linkId_tagId_key;type:varchar(191);not null" json:"linkId"`
	TagID     string   `gorm:"column:tagId;uniqueIndex:LinkTag_linkId_tagId_key;index:LinkTag_tagId_idx;type:varchar(191);not null" json:"tagId"`
	Link      *Link    `gorm:"belongsTo"`
	Tag       *Tag     `gorm:"belongsTo"`
}

func (LinkTag) TableName() string {
	return "LinkTag"
}

func (m *LinkTag) BeforeCreate(tx *gorm.DB) error {
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

func (*LinkTag) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Token struct {
	ID         string    `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(191);not null" json:"name"`
	HashedKey  string    `gorm:"column:hashedKey;uniqueIndex:Token_hashedKey_key;type:varchar(191);not null" json:"hashedKey"`
	PartialKey string    `gorm:"column:partialKey;type:varchar(191);not null" json:"partialKey"`
	Expires    *DateTime `gorm:"column:expires" json:"expires"`
	LastUsed   *DateTime `gorm:"column:lastUsed" json:"lastUsed"`
	CreatedAt  DateTime  `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	UserID     string    `gorm:"column:userId;index:Token_userId_idx;type:varchar(191);not null" json:"userId"`
	User       *User     `gorm:"belongsTo"`
}

func (Token) TableName() string {
	return "Token"
}

func (m *Token) BeforeCreate(tx *gorm.DB) error {
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

func (*Token) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RestrictedToken struct {
	ID                   string                `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name                 string                `gorm:"column:name;type:varchar(191);not null" json:"name"`
	HashedKey            string                `gorm:"column:hashedKey;uniqueIndex:RestrictedToken_hashedKey_key;type:varchar(191);not null" json:"hashedKey"`
	PartialKey           string                `gorm:"column:partialKey;type:varchar(191);not null" json:"partialKey"`
	Scopes               *string               `gorm:"column:scopes;type:text" json:"scopes"`
	Expires              *DateTime             `gorm:"column:expires" json:"expires"`
	LastUsed             *DateTime             `gorm:"column:lastUsed" json:"lastUsed"`
	CreatedAt            DateTime              `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime              `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	UserID               string                `gorm:"column:userId;index:RestrictedToken_userId_idx;type:varchar(191);not null" json:"userId"`
	ProjectID            string                `gorm:"column:projectId;index:RestrictedToken_projectId_idx;type:varchar(191);not null" json:"projectId"`
	InstallationID       *string               `gorm:"column:installationId;index:RestrictedToken_installationId_idx;type:varchar(191)" json:"installationId"`
	User                 *User                 `gorm:"belongsTo"`
	Project              *Project              `gorm:"belongsTo"`
	InstalledIntegration *InstalledIntegration `gorm:"foreignKey:InstallationID;belongsTo"`
	RefreshTokens        []OAuthRefreshToken   `gorm:"foreignKey:AccessTokenID;constraint:OAuthRefreshToken_accessTokenId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RestrictedToken) TableName() string {
	return "RestrictedToken"
}

func (m *RestrictedToken) BeforeCreate(tx *gorm.DB) error {
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

func (*RestrictedToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type VerificationToken struct {
	Identifier string   `gorm:"column:identifier;uniqueIndex:VerificationToken_identifier_token_key;type:varchar(191);not null" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:VerificationToken_token_key;uniqueIndex:VerificationToken_identifier_token_key;type:varchar(191)" json:"token"`
	Expires    DateTime `gorm:"column:expires;not null" json:"expires"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

type EmailVerificationToken struct {
	Identifier string   `gorm:"column:identifier;uniqueIndex:EmailVerificationToken_identifier_token_key;type:varchar(191);not null" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:EmailVerificationToken_token_key;uniqueIndex:EmailVerificationToken_identifier_token_key;type:varchar(191)" json:"token"`
	Expires    DateTime `gorm:"column:expires;not null" json:"expires"`
}

func (EmailVerificationToken) TableName() string {
	return "EmailVerificationToken"
}

type PasswordResetToken struct {
	Identifier string   `gorm:"column:identifier;uniqueIndex:PasswordResetToken_identifier_token_key;type:varchar(191);not null" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:PasswordResetToken_token_key;uniqueIndex:PasswordResetToken_identifier_token_key;type:varchar(191)" json:"token"`
	Expires    DateTime `gorm:"column:expires;not null" json:"expires"`
}

func (PasswordResetToken) TableName() string {
	return "PasswordResetToken"
}

type UtmTemplate struct {
	ID          string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name        string   `gorm:"column:name;uniqueIndex:UtmTemplate_projectId_name_key;type:varchar(191);not null" json:"name"`
	UtmSource   *string  `gorm:"column:utm_source;type:varchar(255)" json:"utm_source"`
	UtmMedium   *string  `gorm:"column:utm_medium;type:varchar(255)" json:"utm_medium"`
	UtmCampaign *string  `gorm:"column:utm_campaign;type:varchar(255)" json:"utm_campaign"`
	UtmTerm     *string  `gorm:"column:utm_term;type:varchar(255)" json:"utm_term"`
	UtmContent  *string  `gorm:"column:utm_content;type:varchar(255)" json:"utm_content"`
	Ref         *string  `gorm:"column:ref;type:varchar(255)" json:"ref"`
	UserID      *string  `gorm:"column:userId;index:UtmTemplate_userId_idx;type:varchar(191)" json:"userId"`
	ProjectID   *string  `gorm:"column:projectId;uniqueIndex:UtmTemplate_projectId_name_key;type:varchar(191)" json:"projectId"`
	CreatedAt   DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	User        *User    `gorm:"belongsTo"`
	Project     *Project `gorm:"belongsTo"`
}

func (UtmTemplate) TableName() string {
	return "UtmTemplate"
}

func (m *UtmTemplate) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*UtmTemplate) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Webhook struct {
	ID                   string                `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	ProjectID            string                `gorm:"column:projectId;index:Webhook_projectId_idx;type:varchar(191);not null" json:"projectId"`
	InstallationID       *string               `gorm:"column:installationId;index:Webhook_installationId_idx;type:varchar(191)" json:"installationId"`
	Receiver             string                `gorm:"column:receiver;default:'user';not null" json:"receiver"`
	Name                 string                `gorm:"column:name;type:varchar(191);not null" json:"name"`
	URL                  string                `gorm:"column:url;type:text;not null" json:"url"`
	Secret               string                `gorm:"column:secret;type:varchar(191);not null" json:"secret"`
	Triggers             datatypes.JSON        `gorm:"column:triggers;not null" json:"triggers"`
	LinkScope            *string               `gorm:"column:linkScope" json:"linkScope"`
	ConsecutiveFailures  int                   `gorm:"column:consecutiveFailures;size:32;default:0;not null" json:"consecutiveFailures"`
	LastFailedAt         *DateTime             `gorm:"column:lastFailedAt" json:"lastFailedAt"`
	DisabledAt           *DateTime             `gorm:"column:disabledAt" json:"disabledAt"`
	CreatedAt            DateTime              `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime              `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Project              *Project              `gorm:"belongsTo"`
	InstalledIntegration *InstalledIntegration `gorm:"foreignKey:InstallationID;belongsTo"`
	Links                []LinkWebhook         `gorm:"foreignKey:WebhookID;constraint:LinkWebhook_webhookId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Folders              []FolderWebhook       `gorm:"foreignKey:WebhookID;constraint:FolderWebhook_webhookId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Webhook) TableName() string {
	return "Webhook"
}

func (m *Webhook) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type LinkWebhook struct {
	ID        string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	LinkID    string   `gorm:"column:linkId;uniqueIndex:LinkWebhook_linkId_webhookId_key;type:varchar(191);not null" json:"linkId"`
	WebhookID string   `gorm:"column:webhookId;uniqueIndex:LinkWebhook_linkId_webhookId_key;index:LinkWebhook_webhookId_idx;type:varchar(191);not null" json:"webhookId"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Link      *Link    `gorm:"belongsTo"`
	Webhook   *Webhook `gorm:"belongsTo"`
}

func (LinkWebhook) TableName() string {
	return "LinkWebhook"
}

func (m *LinkWebhook) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type FolderWebhook struct {
	ID        string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	FolderID  string   `gorm:"column:folderId;uniqueIndex:FolderWebhook_folderId_webhookId_key;type:varchar(191);not null" json:"folderId"`
	WebhookID string   `gorm:"column:webhookId;uniqueIndex:FolderWebhook_folderId_webhookId_key;index:FolderWebhook_webhookId_idx;type:varchar(191);not null" json:"webhookId"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Folder    *Folder  `gorm:"belongsTo"`
	Webhook   *Webhook `gorm:"belongsTo"`
}

func (FolderWebhook) TableName() string {
	return "FolderWebhook"
}

func (m *FolderWebhook) BeforeCreate(tx *gorm.DB) error {
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
	ID                        string                 `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name                      string                 `gorm:"column:name;type:varchar(191);not null" json:"name"`
	Slug                      string                 `gorm:"column:slug;uniqueIndex:Project_slug_key;type:varchar(191);not null" json:"slug"`
	Logo                      *string                `gorm:"column:logo;type:varchar(191)" json:"logo"`
	InviteCode                *string                `gorm:"column:inviteCode;uniqueIndex:Project_inviteCode_key;type:varchar(191)" json:"inviteCode"`
	Plan                      *string                `gorm:"column:plan;type:varchar(191);default:'free';not null" json:"plan"`
	PlanTier                  *int                   `gorm:"column:planTier;size:32;default:1;not null" json:"planTier"`
	PlanPeriod                *string                `gorm:"column:planPeriod" json:"planPeriod"`
	StripeID                  *string                `gorm:"column:stripeId;uniqueIndex:Project_stripeId_key;type:varchar(191)" json:"stripeId"`
	BillingCycleStart         int                    `gorm:"column:billingCycleStart;size:32;not null" json:"billingCycleStart"`
	BillingCycleEndsAt        *DateTime              `gorm:"column:billingCycleEndsAt" json:"billingCycleEndsAt"`
	TrialEndsAt               *DateTime              `gorm:"column:trialEndsAt;index:Project_trialEndsAt_idx" json:"trialEndsAt"`
	SubscriptionCanceledAt    *DateTime              `gorm:"column:subscriptionCanceledAt" json:"subscriptionCanceledAt"`
	PaymentFailedAt           *DateTime              `gorm:"column:paymentFailedAt" json:"paymentFailedAt"`
	InvoicePrefix             *string                `gorm:"column:invoicePrefix;uniqueIndex:Project_invoicePrefix_key;type:varchar(191)" json:"invoicePrefix"`
	StripeConnectID           *string                `gorm:"column:stripeConnectId;uniqueIndex:Project_stripeConnectId_stripeConnectMode_key;type:varchar(191)" json:"stripeConnectId"`
	StripeConnectMode         *string                `gorm:"column:stripeConnectMode;uniqueIndex:Project_stripeConnectId_stripeConnectMode_key" json:"stripeConnectMode"`
	ShopifyStoreID            *string                `gorm:"column:shopifyStoreId;uniqueIndex:Project_shopifyStoreId_key;type:varchar(191)" json:"shopifyStoreId"`
	TotalLinks                int                    `gorm:"column:totalLinks;size:32;default:0;not null" json:"totalLinks"`
	TotalClicks               int                    `gorm:"column:totalClicks;size:32;default:0;not null" json:"totalClicks"`
	Usage                     int                    `gorm:"column:usage;size:32;default:0;not null" json:"usage"`
	UsageLimit                *int                   `gorm:"column:usageLimit;size:32;default:1000;not null" json:"usageLimit"`
	LinksUsage                int                    `gorm:"column:linksUsage;size:32;default:0;not null" json:"linksUsage"`
	LinksLimit                *int                   `gorm:"column:linksLimit;size:32;default:25;not null" json:"linksLimit"`
	DomainsLimit              *int                   `gorm:"column:domainsLimit;size:32;default:3;not null" json:"domainsLimit"`
	TagsLimit                 *int                   `gorm:"column:tagsLimit;size:32;default:5;not null" json:"tagsLimit"`
	FoldersUsage              int                    `gorm:"column:foldersUsage;size:32;default:0;not null" json:"foldersUsage"`
	FoldersLimit              int                    `gorm:"column:foldersLimit;size:32;default:0;not null" json:"foldersLimit"`
	UsersLimit                *int                   `gorm:"column:usersLimit;size:32;default:1;not null" json:"usersLimit"`
	AiUsage                   int                    `gorm:"column:aiUsage;size:32;default:0;not null" json:"aiUsage"`
	AiLimit                   *int                   `gorm:"column:aiLimit;size:32;default:10;not null" json:"aiLimit"`
	Store                     datatypes.JSON         `gorm:"column:store" json:"store"`
	SiteVisitTrackingSettings datatypes.JSON         `gorm:"column:siteVisitTrackingSettings" json:"siteVisitTrackingSettings"`
	AllowedHostnames          datatypes.JSON         `gorm:"column:allowedHostnames" json:"allowedHostnames"`
	PublishableKey            *string                `gorm:"column:publishableKey;uniqueIndex:Project_publishableKey_key;type:varchar(191)" json:"publishableKey"`
	ConversionEnabled         bool                   `gorm:"column:conversionEnabled;default:false;not null" json:"conversionEnabled"`
	WebhookEnabled            bool                   `gorm:"column:webhookEnabled;default:false;not null" json:"webhookEnabled"`
	DotLinkClaimed            bool                   `gorm:"column:dotLinkClaimed;default:false;not null" json:"dotLinkClaimed"`
	SsoEmailDomain            *string                `gorm:"column:ssoEmailDomain;uniqueIndex:Project_ssoEmailDomain_key;type:varchar(191)" json:"ssoEmailDomain"`
	SsoEnforcedAt             *DateTime              `gorm:"column:ssoEnforcedAt" json:"ssoEnforcedAt"`
	CreatedAt                 DateTime               `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime               `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	UsageLastChecked          DateTime               `gorm:"column:usageLastChecked;index:Project_usageLastChecked_idx;not null" json:"usageLastChecked"`
	DisabledAt                *DateTime              `gorm:"column:disabledAt" json:"disabledAt"`
	Users                     []ProjectUsers         `gorm:"foreignKey:ProjectID;constraint:ProjectUsers_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites                   []ProjectInvite        `gorm:"foreignKey:ProjectID;constraint:ProjectInvite_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SentEmails                []SentEmail            `gorm:"foreignKey:ProjectID;constraint:SentEmail_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Links                     []Link                 `gorm:"foreignKey:ProjectID;constraint:Link_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Domains                   []Domain               `gorm:"foreignKey:ProjectID;constraint:Domain_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags                      []Tag                  `gorm:"foreignKey:ProjectID;constraint:Tag_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invoices                  []Invoice              `gorm:"foreignKey:WorkspaceID;constraint:Invoice_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Customers                 []Customer             `gorm:"foreignKey:ProjectID;constraint:Customer_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DefaultDomains            []DefaultDomains       `gorm:"foreignKey:ProjectID;constraint:DefaultDomains_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RestrictedTokens          []RestrictedToken      `gorm:"foreignKey:ProjectID;constraint:RestrictedToken_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OAuthCodes                []OAuthCode            `gorm:"foreignKey:ProjectID;constraint:OAuthCode_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Integrations              []Integration          `gorm:"foreignKey:ProjectID;constraint:Integration_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	InstalledIntegrations     []InstalledIntegration `gorm:"foreignKey:ProjectID;constraint:InstalledIntegration_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks                  []Webhook              `gorm:"foreignKey:ProjectID;constraint:Webhook_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Folders                   []Folder               `gorm:"foreignKey:ProjectID;constraint:Folder_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	RegisteredDomains         []RegisteredDomain     `gorm:"foreignKey:ProjectID;constraint:RegisteredDomain_projectId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Dashboards                []Dashboard            `gorm:"foreignKey:ProjectID;constraint:Dashboard_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UtmTemplates              []UtmTemplate          `gorm:"foreignKey:ProjectID;constraint:UtmTemplate_projectId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	YearInReviews             []YearInReview         `gorm:"foreignKey:WorkspaceID;constraint:YearInReview_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Project) TableName() string {
	return "Project"
}

func (m *Project) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	if m.UsageLastChecked.IsZero() {
		m.UsageLastChecked = DateTime{Time: now}
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

type ProjectInvite struct {
	Email     string   `gorm:"column:email;primaryKey;uniqueIndex:ProjectInvite_email_projectId_key;type:varchar(191)" json:"email"`
	Expires   DateTime `gorm:"column:expires;not null" json:"expires"`
	ProjectID string   `gorm:"column:projectId;primaryKey;uniqueIndex:ProjectInvite_email_projectId_key;index:ProjectInvite_projectId_idx;type:varchar(191)" json:"projectId"`
	Role      string   `gorm:"column:role;default:'member';not null" json:"role"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	Project   *Project `gorm:"belongsTo"`
}

func (ProjectInvite) TableName() string {
	return "ProjectInvite"
}

func (m *ProjectInvite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type ProjectUsers struct {
	ID                     string                  `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Role                   string                  `gorm:"column:role;default:'member';not null" json:"role"`
	UserID                 string                  `gorm:"column:userId;uniqueIndex:ProjectUsers_userId_projectId_key;type:varchar(191);not null" json:"userId"`
	ProjectID              string                  `gorm:"column:projectId;uniqueIndex:ProjectUsers_userId_projectId_key;index:ProjectUsers_projectId_idx;type:varchar(191);not null" json:"projectId"`
	WorkspacePreferences   datatypes.JSON          `gorm:"column:workspacePreferences" json:"workspacePreferences"`
	DefaultFolderID        *string                 `gorm:"column:defaultFolderId;type:varchar(191)" json:"defaultFolderId"`
	CreatedAt              DateTime                `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt              DateTime                `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	User                   *User                   `gorm:"belongsTo"`
	Project                *Project                `gorm:"belongsTo"`
	NotificationPreference *NotificationPreference `gorm:"foreignKey:ProjectUserID;constraint:NotificationPreference_projectUserId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ProjectUsers) TableName() string {
	return "ProjectUsers"
}

func (m *ProjectUsers) BeforeCreate(tx *gorm.DB) error {
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

func (*ProjectUsers) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SentEmail struct {
	ID        string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Type      string   `gorm:"column:type;type:varchar(191);not null" json:"type"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	ProjectID *string  `gorm:"column:projectId;index:SentEmail_projectId_idx;type:varchar(191)" json:"projectId"`
	Project   *Project `gorm:"belongsTo"`
}

func (SentEmail) TableName() string {
	return "SentEmail"
}

func (m *SentEmail) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
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
	return dateTime.UTC().Truncate(time.Millisecond).Format("2006-01-02 15:04:05.000"), nil
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
		return time.Date(value.Year(), value.Month(), value.Day(), value.Hour(), value.Minute(), value.Second(), value.Nanosecond(), time.UTC).Truncate(time.Millisecond), nil
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
