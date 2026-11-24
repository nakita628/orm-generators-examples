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
	"gorm.io/gorm/schema"
)

var NamingStrategy = schema.NamingStrategy{SingularTable: true, NoLowerCase: true}

type Account struct {
	ID                string  `gorm:"column:id;primaryKey" json:"id"`
	UserID            int     `gorm:"column:userId;size:32;not null" json:"userId"`
	Type              string  `gorm:"column:type;not null" json:"type"`
	Provider          string  `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key;not null" json:"provider"`
	ProviderAccountID string  `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key;not null" json:"providerAccountId"`
	RefreshToken      *string `gorm:"column:refresh_token" json:"refresh_token"`
	AccessToken       *string `gorm:"column:access_token" json:"access_token"`
	ExpiresAt         *int    `gorm:"column:expires_at;size:32" json:"expires_at"`
	TokenType         *string `gorm:"column:token_type" json:"token_type"`
	Scope             *string `gorm:"column:scope" json:"scope"`
	IDToken           *string `gorm:"column:id_token" json:"id_token"`
	SessionState      *string `gorm:"column:session_state" json:"session_state"`
	User              *User   `gorm:"belongsTo"`
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

type User struct {
	ID                      int                   `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UUID                    string                `gorm:"column:uuid;uniqueIndex:User_uuid_key;type:char(36);not null" json:"uuid"`
	Name                    *string               `gorm:"column:name" json:"name"`
	Username                *string               `gorm:"column:username;uniqueIndex:User_username_key" json:"username"`
	Email                   *string               `gorm:"column:email;uniqueIndex:User_email_key" json:"email"`
	EmailVerified           *DateTime             `gorm:"column:emailVerified;type:timestamp(3)" json:"emailVerified"`
	UnverifiedNewEmail      *string               `gorm:"column:unverifiedNewEmail" json:"unverifiedNewEmail"`
	Image                   *string               `gorm:"column:image" json:"image"`
	Password                *string               `gorm:"column:password" json:"password"`
	Locale                  *string               `gorm:"column:locale;default:'en';not null" json:"locale"`
	ParentSubscriptionID    *int                  `gorm:"column:parentSubscriptionId;size:32" json:"parentSubscriptionId"`
	CollectionOrder         ScalarList[int]       `gorm:"column:collectionOrder;type:integer[]" json:"collectionOrder"`
	LinksRouteTo            string                `gorm:"column:linksRouteTo;default:'ORIGINAL';not null" json:"linksRouteTo"`
	AiTaggingMethod         string                `gorm:"column:aiTaggingMethod;default:'DISABLED';not null" json:"aiTaggingMethod"`
	AiPredefinedTags        ScalarList[string]    `gorm:"column:aiPredefinedTags;type:text[]" json:"aiPredefinedTags"`
	AiTagExistingLinks      bool                  `gorm:"column:aiTagExistingLinks;default:false;not null" json:"aiTagExistingLinks"`
	Theme                   string                `gorm:"column:theme;default:'dark';not null" json:"theme"`
	DismissedAnnouncementID *string               `gorm:"column:dismissedAnnouncementId" json:"dismissedAnnouncementId"`
	ReadableFontFamily      *string               `gorm:"column:readableFontFamily;default:'sans-serif'" json:"readableFontFamily"`
	ReadableFontSize        *string               `gorm:"column:readableFontSize;default:'20px'" json:"readableFontSize"`
	ReadableLineHeight      *string               `gorm:"column:readableLineHeight;default:'1.8'" json:"readableLineHeight"`
	ReadableLineWidth       *string               `gorm:"column:readableLineWidth;default:'normal'" json:"readableLineWidth"`
	PreventDuplicateLinks   bool                  `gorm:"column:preventDuplicateLinks;default:false;not null" json:"preventDuplicateLinks"`
	ArchiveAsScreenshot     *bool                 `gorm:"column:archiveAsScreenshot;default:true;not null" json:"archiveAsScreenshot"`
	ArchiveAsMonolith       *bool                 `gorm:"column:archiveAsMonolith;default:true;not null" json:"archiveAsMonolith"`
	ArchiveAsPDF            *bool                 `gorm:"column:archiveAsPDF;default:true;not null" json:"archiveAsPDF"`
	ArchiveAsReadable       *bool                 `gorm:"column:archiveAsReadable;default:true;not null" json:"archiveAsReadable"`
	ArchiveAsWaybackMachine bool                  `gorm:"column:archiveAsWaybackMachine;default:false;not null" json:"archiveAsWaybackMachine"`
	ReferredBy              *string               `gorm:"column:referredBy" json:"referredBy"`
	LastPickedAt            *DateTime             `gorm:"column:lastPickedAt;type:timestamp(3)" json:"lastPickedAt"`
	AcceptPromotionalEmails bool                  `gorm:"column:acceptPromotionalEmails;default:false;not null" json:"acceptPromotionalEmails"`
	ShareUsageAnalytics     *bool                 `gorm:"column:shareUsageAnalytics;default:true;not null" json:"shareUsageAnalytics"`
	TrialEndEmailSent       bool                  `gorm:"column:trialEndEmailSent;default:false;not null" json:"trialEndEmailSent"`
	SessionVersion          int                   `gorm:"column:sessionVersion;size:32;default:0;not null" json:"sessionVersion"`
	PasswordRetries         int                   `gorm:"column:passwordRetries;size:32;default:0;not null" json:"passwordRetries"`
	PasswordRetryAt         *DateTime             `gorm:"column:passwordRetryAt;type:timestamp(3)" json:"passwordRetryAt"`
	CreatedAt               DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime              `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	ParentSubscription      *Subscription         `gorm:"foreignKey:ParentSubscriptionID;belongsTo"`
	Accounts                []Account             `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Collections             []Collection          `gorm:"foreignKey:OwnerID;constraint:Collection_ownerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags                    []Tag                 `gorm:"foreignKey:OwnerID;constraint:Tag_ownerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedLinks            []Link                `gorm:"foreignKey:CreatedByID;constraint:Link_createdById_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedCollections      []Collection          `gorm:"foreignKey:CreatedByID;constraint:Collection_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Highlights              []Highlight           `gorm:"foreignKey:UserID;constraint:Highlight_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CollectionsJoined       []UsersAndCollections `gorm:"foreignKey:UserID;constraint:UsersAndCollections_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessTokens            []AccessToken         `gorm:"foreignKey:UserID;constraint:AccessToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DashboardSections       []DashboardSection    `gorm:"foreignKey:UserID;constraint:DashboardSection_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subscriptions           *Subscription         `gorm:"foreignKey:UserID;constraint:Subscription_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PinnedLinks             []Link                `gorm:"many2many:_PinnedLinks;joinForeignKey:B;joinReferences:A"`
}

func (User) TableName() string {
	return "User"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UUID == "" {
		m.UUID = uuid.NewString()
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
	Identifier string   `gorm:"column:identifier;uniqueIndex:VerificationToken_identifier_token_key;not null" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:VerificationToken_token_key;uniqueIndex:VerificationToken_identifier_token_key" json:"token"`
	Expires    DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

func (m *VerificationToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type EmailVerificationCode struct {
	Identifier string   `gorm:"column:identifier;primaryKey;uniqueIndex:EmailVerificationCode_identifier_key" json:"identifier"`
	Code       string   `gorm:"column:code;not null" json:"code"`
	Attempts   int      `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	Expires    DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (EmailVerificationCode) TableName() string {
	return "EmailVerificationCode"
}

func (m *EmailVerificationCode) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EmailVerificationCode) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PasswordResetToken struct {
	Identifier string   `gorm:"column:identifier;not null" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:PasswordResetToken_token_key" json:"token"`
	Expires    DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (PasswordResetToken) TableName() string {
	return "PasswordResetToken"
}

func (m *PasswordResetToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PasswordResetToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Collection struct {
	ID               int                   `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name             string                `gorm:"column:name;not null" json:"name"`
	Description      string                `gorm:"column:description;default:'';not null" json:"description"`
	Icon             *string               `gorm:"column:icon" json:"icon"`
	IconWeight       *string               `gorm:"column:iconWeight" json:"iconWeight"`
	Color            *string               `gorm:"column:color;default:'#0ea5e9';not null" json:"color"`
	ParentID         *int                  `gorm:"column:parentId;size:32" json:"parentId"`
	IsPublic         bool                  `gorm:"column:isPublic;default:false;not null" json:"isPublic"`
	OwnerID          int                   `gorm:"column:ownerId;index:Collection_ownerId_idx;size:32;not null" json:"ownerId"`
	CreatedByID      *int                  `gorm:"column:createdById;size:32" json:"createdById"`
	CreatedAt        DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime              `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Parent           *Collection           `gorm:"foreignKey:ParentID;belongsTo"`
	Owner            *User                 `gorm:"foreignKey:OwnerID;belongsTo"`
	CreatedBy        *User                 `gorm:"foreignKey:CreatedByID;belongsTo"`
	SubCollections   []Collection          `gorm:"foreignKey:ParentID;constraint:Collection_parentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Members          []UsersAndCollections `gorm:"foreignKey:CollectionID;constraint:UsersAndCollections_collectionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Links            []Link                `gorm:"foreignKey:CollectionID;constraint:Link_collectionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RssSubscriptions []RssSubscription     `gorm:"foreignKey:CollectionID;constraint:RssSubscription_collectionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DashboardSection []DashboardSection    `gorm:"foreignKey:CollectionID;constraint:DashboardSection_collectionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Collection) TableName() string {
	return "Collection"
}

func (m *Collection) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Collection) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type UsersAndCollections struct {
	UserID       int         `gorm:"column:userId;primaryKey;index:UsersAndCollections_userId_idx;size:32" json:"userId"`
	CollectionID int         `gorm:"column:collectionId;primaryKey;size:32" json:"collectionId"`
	CanCreate    bool        `gorm:"column:canCreate;not null" json:"canCreate"`
	CanUpdate    bool        `gorm:"column:canUpdate;not null" json:"canUpdate"`
	CanDelete    bool        `gorm:"column:canDelete;not null" json:"canDelete"`
	CreatedAt    DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime    `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User         *User       `gorm:"belongsTo"`
	Collection   *Collection `gorm:"belongsTo"`
}

func (UsersAndCollections) TableName() string {
	return "UsersAndCollections"
}

func (m *UsersAndCollections) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*UsersAndCollections) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Link struct {
	ID              int         `gorm:"column:id;primaryKey;autoIncrement;index:Link_collectionId_collectionOrder_id_idx;size:32" json:"id"`
	Name            string      `gorm:"column:name;default:'';not null" json:"name"`
	Type            *string     `gorm:"column:type;default:'url';not null" json:"type"`
	Description     string      `gorm:"column:description;default:'';not null" json:"description"`
	CreatedByID     *int        `gorm:"column:createdById;size:32" json:"createdById"`
	CollectionID    int         `gorm:"column:collectionId;index:Link_collectionId_idx;index:Link_collectionId_collectionOrder_id_idx;size:32;not null" json:"collectionId"`
	CollectionOrder float64     `gorm:"column:collectionOrder;index:Link_collectionId_collectionOrder_id_idx;type:double precision;not null" json:"collectionOrder"`
	Icon            *string     `gorm:"column:icon" json:"icon"`
	IconWeight      *string     `gorm:"column:iconWeight" json:"iconWeight"`
	Color           *string     `gorm:"column:color" json:"color"`
	URL             *string     `gorm:"column:url" json:"url"`
	TextContent     *string     `gorm:"column:textContent" json:"textContent"`
	Preview         *string     `gorm:"column:preview" json:"preview"`
	Image           *string     `gorm:"column:image" json:"image"`
	Pdf             *string     `gorm:"column:pdf" json:"pdf"`
	Readable        *string     `gorm:"column:readable" json:"readable"`
	Monolith        *string     `gorm:"column:monolith" json:"monolith"`
	ClientSide      bool        `gorm:"column:clientSide;default:false;not null" json:"clientSide"`
	AiTagged        bool        `gorm:"column:aiTagged;default:false;not null" json:"aiTagged"`
	MetaDescription *string     `gorm:"column:metaDescription" json:"metaDescription"`
	IndexVersion    *int        `gorm:"column:indexVersion;size:32" json:"indexVersion"`
	LastPreserved   *DateTime   `gorm:"column:lastPreserved;type:timestamp(3)" json:"lastPreserved"`
	ImportDate      *DateTime   `gorm:"column:importDate;type:timestamp(3)" json:"importDate"`
	CreatedAt       DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime    `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedBy       *User       `gorm:"foreignKey:CreatedByID;belongsTo"`
	Collection      *Collection `gorm:"belongsTo"`
	Highlight       []Highlight `gorm:"foreignKey:LinkID;constraint:Highlight_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PinnedBy        []User      `gorm:"many2many:_PinnedLinks;joinForeignKey:A;joinReferences:B"`
	Tags            []Tag       `gorm:"many2many:_LinkToTag;joinForeignKey:A;joinReferences:B"`
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

type Tag struct {
	ID                      int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name                    string   `gorm:"column:name;uniqueIndex:Tag_name_ownerId_key;not null" json:"name"`
	OwnerID                 int      `gorm:"column:ownerId;uniqueIndex:Tag_name_ownerId_key;index:Tag_ownerId_idx;size:32;not null" json:"ownerId"`
	ArchiveAsScreenshot     *bool    `gorm:"column:archiveAsScreenshot" json:"archiveAsScreenshot"`
	ArchiveAsMonolith       *bool    `gorm:"column:archiveAsMonolith" json:"archiveAsMonolith"`
	ArchiveAsPDF            *bool    `gorm:"column:archiveAsPDF" json:"archiveAsPDF"`
	ArchiveAsReadable       *bool    `gorm:"column:archiveAsReadable" json:"archiveAsReadable"`
	ArchiveAsWaybackMachine *bool    `gorm:"column:archiveAsWaybackMachine" json:"archiveAsWaybackMachine"`
	AiTag                   *bool    `gorm:"column:aiTag" json:"aiTag"`
	AiGenerated             bool     `gorm:"column:aiGenerated;default:false;not null" json:"aiGenerated"`
	CreatedAt               DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Owner                   *User    `gorm:"foreignKey:OwnerID;belongsTo"`
	Links                   []Link   `gorm:"many2many:_LinkToTag;joinForeignKey:B;joinReferences:A"`
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

type Subscription struct {
	ID                         int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Active                     bool           `gorm:"column:active;not null" json:"active"`
	Provider                   string         `gorm:"column:provider;not null" json:"provider"`
	CurrentPeriodStart         DateTime       `gorm:"column:currentPeriodStart;type:timestamp(3);not null" json:"currentPeriodStart"`
	CurrentPeriodEnd           DateTime       `gorm:"column:currentPeriodEnd;type:timestamp(3);not null" json:"currentPeriodEnd"`
	Quantity                   *int           `gorm:"column:quantity;size:32;default:1;not null" json:"quantity"`
	StoreOriginalTransactionID *string        `gorm:"column:storeOriginalTransactionId;uniqueIndex:Subscription_storeOriginalTransactionId_key" json:"storeOriginalTransactionId"`
	StoreProductID             *string        `gorm:"column:storeProductId" json:"storeProductId"`
	GooglePurchaseToken        *string        `gorm:"column:googlePurchaseToken;uniqueIndex:Subscription_googlePurchaseToken_key" json:"googlePurchaseToken"`
	StoreMetadata              datatypes.JSON `gorm:"column:storeMetadata" json:"storeMetadata"`
	StripeSubscriptionID       *string        `gorm:"column:stripeSubscriptionId;uniqueIndex:Subscription_stripeSubscriptionId_key" json:"stripeSubscriptionId"`
	UserID                     int            `gorm:"column:userId;uniqueIndex:Subscription_userId_key;size:32;not null" json:"userId"`
	CreatedAt                  DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User                       *User          `gorm:"belongsTo"`
	ChildUsers                 []User         `gorm:"foreignKey:ParentSubscriptionID;constraint:User_parentSubscriptionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
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

type AccessToken struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	UserID     int       `gorm:"column:userId;size:32;not null" json:"userId"`
	Token      string    `gorm:"column:token;uniqueIndex:AccessToken_token_key;not null" json:"token"`
	Revoked    bool      `gorm:"column:revoked;default:false;not null" json:"revoked"`
	IsSession  bool      `gorm:"column:isSession;default:false;not null" json:"isSession"`
	Expires    DateTime  `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	LastUsedAt *DateTime `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User       *User     `gorm:"belongsTo"`
}

func (AccessToken) TableName() string {
	return "AccessToken"
}

func (m *AccessToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AccessToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RssSubscription struct {
	ID            int         `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	URL           string      `gorm:"column:url;not null" json:"url"`
	Name          string      `gorm:"column:name;not null" json:"name"`
	LastBuildDate *DateTime   `gorm:"column:lastBuildDate;type:timestamp(3)" json:"lastBuildDate"`
	CollectionID  int         `gorm:"column:collectionId;size:32;not null" json:"collectionId"`
	OwnerID       int         `gorm:"column:ownerId;size:32;not null" json:"ownerId"`
	CreatedAt     DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime    `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Collection    *Collection `gorm:"belongsTo"`
}

func (RssSubscription) TableName() string {
	return "RssSubscription"
}

func (m *RssSubscription) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*RssSubscription) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Highlight struct {
	ID          int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Color       string   `gorm:"column:color;not null" json:"color"`
	Comment     *string  `gorm:"column:comment" json:"comment"`
	LinkID      int      `gorm:"column:linkId;size:32;not null" json:"linkId"`
	UserID      int      `gorm:"column:userId;size:32;not null" json:"userId"`
	StartOffset int      `gorm:"column:startOffset;size:32;not null" json:"startOffset"`
	EndOffset   int      `gorm:"column:endOffset;size:32;not null" json:"endOffset"`
	Text        string   `gorm:"column:text;not null" json:"text"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	Link        *Link    `gorm:"belongsTo"`
	User        *User    `gorm:"belongsTo"`
}

func (Highlight) TableName() string {
	return "Highlight"
}

func (m *Highlight) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Highlight) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DashboardSection struct {
	ID           int         `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID       int         `gorm:"column:userId;uniqueIndex:DashboardSection_userId_collectionId_key;size:32;not null" json:"userId"`
	CollectionID *int        `gorm:"column:collectionId;uniqueIndex:DashboardSection_userId_collectionId_key;size:32" json:"collectionId"`
	Type         string      `gorm:"column:type;not null" json:"type"`
	Order        int         `gorm:"column:order;size:32;not null" json:"order"`
	CreatedAt    DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime    `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	User         *User       `gorm:"belongsTo"`
	Collection   *Collection `gorm:"belongsTo"`
}

func (DashboardSection) TableName() string {
	return "DashboardSection"
}

func (m *DashboardSection) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DashboardSection) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AppMigration struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name       string    `gorm:"column:name;uniqueIndex:AppMigration_name_key;not null" json:"name"`
	Status     string    `gorm:"column:status;not null" json:"status"`
	FinishedAt *DateTime `gorm:"column:finishedAt;type:timestamp(3)" json:"finishedAt"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (AppMigration) TableName() string {
	return "AppMigration"
}

func (m *AppMigration) BeforeCreate(tx *gorm.DB) error {
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
