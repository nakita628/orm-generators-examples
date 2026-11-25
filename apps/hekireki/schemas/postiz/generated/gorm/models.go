package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/lucsky/cuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Organization struct {
	ID                  string               `gorm:"column:id;primaryKey" json:"id"`
	Name                string               `gorm:"column:name;not null" json:"name"`
	Description         *string              `gorm:"column:description" json:"description"`
	APIKey              *string              `gorm:"column:apiKey;index:Organization_apiKey_idx" json:"apiKey"`
	PaymentID           *string              `gorm:"column:paymentId;index:Organization_paymentId_idx" json:"paymentId"`
	StreakSince         *DateTime            `gorm:"column:streakSince;index:Organization_streakSince_idx;type:timestamp(3)" json:"streakSince"`
	CreatedAt           DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime             `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt           *DateTime            `gorm:"column:deletedAt;index:Organization_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	AllowTrial          bool                 `gorm:"column:allowTrial;default:false;not null" json:"allowTrial"`
	IsTrailing          bool                 `gorm:"column:isTrailing;default:false;not null" json:"isTrailing"`
	Shortlink           string               `gorm:"column:shortlink;default:'ASK';not null" json:"shortlink"`
	AutoPost            []AutoPost           `gorm:"foreignKey:OrganizationID;constraint:AutoPost_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Comments            []Comments           `gorm:"foreignKey:OrganizationID;constraint:Comments_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Credits             []Credits            `gorm:"foreignKey:OrganizationID;constraint:Credits_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Clippings           []Clipping           `gorm:"foreignKey:OrganizationID;constraint:Clipping_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Customers           []Customer           `gorm:"foreignKey:OrgID;constraint:Customer_orgId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Errors              []Errors             `gorm:"foreignKey:OrganizationID;constraint:Errors_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Github              []GitHub             `gorm:"foreignKey:OrganizationID;constraint:GitHub_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Integration         []Integration        `gorm:"foreignKey:OrganizationID;constraint:Integration_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Media               []Media              `gorm:"foreignKey:OrganizationID;constraint:Media_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	BuyerOrganization   []MessagesGroup      `gorm:"foreignKey:BuyerOrganizationID;constraint:MessagesGroup_buyerOrganizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Notifications       []Notifications      `gorm:"foreignKey:OrganizationID;constraint:Notifications_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Plugs               []Plugs              `gorm:"foreignKey:OrganizationID;constraint:Plugs_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Post                []Post               `gorm:"foreignKey:OrganizationID;constraint:Post_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SubmittedPost       []Post               `gorm:"foreignKey:SubmittedForOrganizationID;constraint:Post_submittedForOrganizationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Sets                []Sets               `gorm:"foreignKey:OrganizationID;constraint:Sets_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Signatures          []Signatures         `gorm:"foreignKey:OrganizationID;constraint:Signatures_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Tags                []Tags               `gorm:"foreignKey:OrgID;constraint:Tags_orgId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ThirdParty          []ThirdParty         `gorm:"foreignKey:OrganizationID;constraint:ThirdParty_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UsedCodes           []UsedCodes          `gorm:"foreignKey:OrgID;constraint:UsedCodes_orgId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Users               []UserOrganization   `gorm:"foreignKey:OrganizationID;constraint:UserOrganization_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Webhooks            []Webhooks           `gorm:"foreignKey:OrganizationID;constraint:Webhooks_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OauthApp            []OAuthApp           `gorm:"foreignKey:OrganizationID;constraint:OAuthApp_organizationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	OauthAuthorizations []OAuthAuthorization `gorm:"foreignKey:OrganizationID;constraint:OAuthAuthorization_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Subscription        *Subscription        `gorm:"foreignKey:OrganizationID;constraint:Subscription_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Organization) TableName() string {
	return "Organization"
}

func (m *Organization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
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

type Tags struct {
	ID           string        `gorm:"column:id;primaryKey" json:"id"`
	Name         string        `gorm:"column:name;not null" json:"name"`
	Color        string        `gorm:"column:color;not null" json:"color"`
	OrgID        string        `gorm:"column:orgId;index:Tags_orgId_idx;not null" json:"orgId"`
	DeletedAt    *DateTime     `gorm:"column:deletedAt;index:Tags_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	CreatedAt    DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization *Organization `gorm:"foreignKey:OrgID;belongsTo"`
	Posts        []TagsPosts   `gorm:"foreignKey:TagID;constraint:TagsPosts_tagId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Tags) TableName() string {
	return "Tags"
}

func (m *Tags) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Tags) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TagsPosts struct {
	PostID    string   `gorm:"column:postId;primaryKey;uniqueIndex:TagsPosts_postId_tagId_key,priority:1" json:"postId"`
	TagID     string   `gorm:"column:tagId;primaryKey;uniqueIndex:TagsPosts_postId_tagId_key,priority:2" json:"tagId"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Post      *Post    `gorm:"belongsTo"`
	Tag       *Tags    `gorm:"foreignKey:TagID;belongsTo"`
}

func (TagsPosts) TableName() string {
	return "TagsPosts"
}

func (m *TagsPosts) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TagsPosts) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type User struct {
	ID                    string               `gorm:"column:id;primaryKey" json:"id"`
	Email                 string               `gorm:"column:email;uniqueIndex:User_email_providerName_key,priority:1;not null" json:"email"`
	Password              *string              `gorm:"column:password" json:"password"`
	ProviderName          string               `gorm:"column:providerName;uniqueIndex:User_email_providerName_key,priority:2;index:User_providerId_providerName_idx,priority:2;not null" json:"providerName"`
	Name                  *string              `gorm:"column:name" json:"name"`
	LastName              *string              `gorm:"column:lastName" json:"lastName"`
	IsSuperAdmin          bool                 `gorm:"column:isSuperAdmin;default:false;not null" json:"isSuperAdmin"`
	Bio                   *string              `gorm:"column:bio" json:"bio"`
	Audience              int                  `gorm:"column:audience;size:32;default:0;not null" json:"audience"`
	PictureID             *string              `gorm:"column:pictureId;index:User_pictureId_idx" json:"pictureId"`
	ProviderID            *string              `gorm:"column:providerId;index:User_providerId_providerName_idx,priority:1" json:"providerId"`
	Timezone              int                  `gorm:"column:timezone;size:32;not null" json:"timezone"`
	TimezoneName          *string              `gorm:"column:timezoneName" json:"timezoneName"`
	CreatedAt             DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime             `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	LastReadNotifications DateTime             `gorm:"column:lastReadNotifications;index:User_lastReadNotifications_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastReadNotifications"`
	InviteID              *string              `gorm:"column:inviteId;index:User_inviteId_idx" json:"inviteId"`
	Activated             *bool                `gorm:"column:activated;default:true;not null" json:"activated"`
	Account               *string              `gorm:"column:account;index:User_account_idx" json:"account"`
	ConnectedAccount      bool                 `gorm:"column:connectedAccount;default:false;not null" json:"connectedAccount"`
	LastOnline            DateTime             `gorm:"column:lastOnline;index:User_lastOnline_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastOnline"`
	IP                    *string              `gorm:"column:ip" json:"ip"`
	Agent                 *string              `gorm:"column:agent" json:"agent"`
	DeletedAt             *DateTime            `gorm:"column:deletedAt;index:User_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	SendSuccessEmails     *bool                `gorm:"column:sendSuccessEmails;default:true;not null" json:"sendSuccessEmails"`
	SendFailureEmails     *bool                `gorm:"column:sendFailureEmails;default:true;not null" json:"sendFailureEmails"`
	SendStreakEmails      *bool                `gorm:"column:sendStreakEmails;default:true;not null" json:"sendStreakEmails"`
	Picture               *Media               `gorm:"foreignKey:PictureID;belongsTo"`
	Comments              []Comments           `gorm:"foreignKey:UserID;constraint:Comments_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Items                 []ItemUser           `gorm:"foreignKey:UserID;constraint:ItemUser_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	GroupBuyer            []MessagesGroup      `gorm:"foreignKey:BuyerID;constraint:MessagesGroup_buyerId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	GroupSeller           []MessagesGroup      `gorm:"foreignKey:SellerID;constraint:MessagesGroup_sellerId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OrderBuyer            []Orders             `gorm:"foreignKey:BuyerID;constraint:Orders_buyerId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OrderSeller           []Orders             `gorm:"foreignKey:SellerID;constraint:Orders_sellerId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PayoutProblems        []PayoutProblems     `gorm:"foreignKey:UserID;constraint:PayoutProblems_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Organizations         []UserOrganization   `gorm:"foreignKey:UserID;constraint:UserOrganization_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OauthAuthorizations   []OAuthAuthorization `gorm:"foreignKey:UserID;constraint:OAuthAuthorization_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Agencies              *SocialMediaAgency   `gorm:"foreignKey:UserID;constraint:SocialMediaAgency_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (User) TableName() string {
	return "User"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	if m.LastReadNotifications.IsZero() {
		m.LastReadNotifications = DateTime{Time: now}
	}
	if m.LastOnline.IsZero() {
		m.LastOnline = DateTime{Time: now}
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

type UsedCodes struct {
	ID           string        `gorm:"column:id;primaryKey" json:"id"`
	Code         string        `gorm:"column:code;index:UsedCodes_code_idx;not null" json:"code"`
	OrgID        string        `gorm:"column:orgId;not null" json:"orgId"`
	CreatedAt    DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization *Organization `gorm:"foreignKey:OrgID;belongsTo"`
}

func (UsedCodes) TableName() string {
	return "UsedCodes"
}

func (m *UsedCodes) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*UsedCodes) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type UserOrganization struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	UserID         string        `gorm:"column:userId;uniqueIndex:UserOrganization_userId_organizationId_key,priority:1;not null" json:"userId"`
	OrganizationID string        `gorm:"column:organizationId;uniqueIndex:UserOrganization_userId_organizationId_key,priority:2;index:UserOrganization_organizationId_idx;not null" json:"organizationId"`
	Disabled       bool          `gorm:"column:disabled;index:UserOrganization_disabled_idx;default:false;not null" json:"disabled"`
	Role           string        `gorm:"column:role;default:'USER';not null" json:"role"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization `gorm:"belongsTo"`
	User           *User         `gorm:"belongsTo"`
}

func (UserOrganization) TableName() string {
	return "UserOrganization"
}

func (m *UserOrganization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*UserOrganization) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type GitHub struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	Login          *string       `gorm:"column:login;index:GitHub_login_idx" json:"login"`
	Name           *string       `gorm:"column:name" json:"name"`
	Token          string        `gorm:"column:token;not null" json:"token"`
	JobID          *string       `gorm:"column:jobId" json:"jobId"`
	OrganizationID string        `gorm:"column:organizationId;index:GitHub_organizationId_idx;not null" json:"organizationId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (GitHub) TableName() string {
	return "GitHub"
}

func (m *GitHub) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*GitHub) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Trending struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	TrendingList string   `gorm:"column:trendingList;not null" json:"trendingList"`
	Language     *string  `gorm:"column:language;uniqueIndex:Trending_language_key" json:"language"`
	Hash         string   `gorm:"column:hash;index:Trending_hash_idx;not null" json:"hash"`
	Date         DateTime `gorm:"column:date;type:timestamp(3);not null" json:"date"`
	CreatedAt    DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Trending) TableName() string {
	return "Trending"
}

func (m *Trending) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Trending) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TrendingLog struct {
	ID       string   `gorm:"column:id;primaryKey" json:"id"`
	Language *string  `gorm:"column:language" json:"language"`
	Date     DateTime `gorm:"column:date;type:timestamp(3);not null" json:"date"`
}

func (TrendingLog) TableName() string {
	return "TrendingLog"
}

func (m *TrendingLog) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type ItemUser struct {
	ID     string `gorm:"column:id;primaryKey" json:"id"`
	UserID string `gorm:"column:userId;uniqueIndex:ItemUser_userId_key_key,priority:1;index:ItemUser_userId_idx;not null" json:"userId"`
	Key    string `gorm:"column:key;uniqueIndex:ItemUser_userId_key_key,priority:2;index:ItemUser_key_idx;not null" json:"key"`
	User   *User  `gorm:"belongsTo"`
}

func (ItemUser) TableName() string {
	return "ItemUser"
}

func (m *ItemUser) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type Star struct {
	ID         string   `gorm:"column:id;primaryKey" json:"id"`
	Stars      int      `gorm:"column:stars;size:32;not null" json:"stars"`
	TotalStars int      `gorm:"column:totalStars;size:32;not null" json:"totalStars"`
	Forks      int      `gorm:"column:forks;size:32;not null" json:"forks"`
	TotalForks int      `gorm:"column:totalForks;size:32;not null" json:"totalForks"`
	Login      string   `gorm:"column:login;uniqueIndex:Star_login_date_key,priority:1;not null" json:"login"`
	Date       Date     `gorm:"column:date;uniqueIndex:Star_login_date_key,priority:2;type:date;default:CURRENT_TIMESTAMP;not null" json:"date"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Star) TableName() string {
	return "Star"
}

func (m *Star) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.Date.IsZero() {
		m.Date = Date{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Star) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Media struct {
	ID                 string              `gorm:"column:id;primaryKey" json:"id"`
	Name               string              `gorm:"column:name;not null" json:"name"`
	OriginalName       *string             `gorm:"column:originalName" json:"originalName"`
	Path               string              `gorm:"column:path;not null" json:"path"`
	OrganizationID     string              `gorm:"column:organizationId;index:Media_organizationId_createdAt_deletedAt_status_idx,priority:1;not null" json:"organizationId"`
	CreatedAt          DateTime            `gorm:"column:createdAt;index:Media_organizationId_createdAt_deletedAt_status_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime            `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt          *DateTime           `gorm:"column:deletedAt;index:Media_organizationId_createdAt_deletedAt_status_idx,priority:3;type:timestamp(3)" json:"deletedAt"`
	FileSize           int                 `gorm:"column:fileSize;size:32;default:0;not null" json:"fileSize"`
	Type               *string             `gorm:"column:type;default:'image';not null" json:"type"`
	Thumbnail          *string             `gorm:"column:thumbnail" json:"thumbnail"`
	Alt                *string             `gorm:"column:alt" json:"alt"`
	ThumbnailTimestamp *int                `gorm:"column:thumbnailTimestamp;size:32" json:"thumbnailTimestamp"`
	Status             *string             `gorm:"column:status;index:Media_organizationId_createdAt_deletedAt_status_idx,priority:4;default:'ready';not null" json:"status"`
	ProcessingError    *string             `gorm:"column:processingError" json:"processingError"`
	Organization       *Organization       `gorm:"belongsTo"`
	Agencies           []SocialMediaAgency `gorm:"foreignKey:LogoID;constraint:SocialMediaAgency_logoId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserPicture        []User              `gorm:"foreignKey:PictureID;constraint:User_pictureId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	OauthApps          []OAuthApp          `gorm:"foreignKey:PictureID;constraint:OAuthApp_pictureId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Media) TableName() string {
	return "Media"
}

func (m *Media) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Media) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SocialMediaAgency struct {
	ID               string                   `gorm:"column:id;primaryKey;index:SocialMediaAgency_id_idx" json:"id"`
	UserID           string                   `gorm:"column:userId;uniqueIndex:SocialMediaAgency_userId_key;index:SocialMediaAgency_userId_idx;not null" json:"userId"`
	Name             string                   `gorm:"column:name;not null" json:"name"`
	LogoID           *string                  `gorm:"column:logoId" json:"logoId"`
	Website          *string                  `gorm:"column:website" json:"website"`
	Slug             *string                  `gorm:"column:slug" json:"slug"`
	Facebook         *string                  `gorm:"column:facebook" json:"facebook"`
	Instagram        *string                  `gorm:"column:instagram" json:"instagram"`
	Twitter          *string                  `gorm:"column:twitter" json:"twitter"`
	LinkedIn         *string                  `gorm:"column:linkedIn" json:"linkedIn"`
	Youtube          *string                  `gorm:"column:youtube" json:"youtube"`
	Tiktok           *string                  `gorm:"column:tiktok" json:"tiktok"`
	OtherSocialMedia *string                  `gorm:"column:otherSocialMedia" json:"otherSocialMedia"`
	ShortDescription string                   `gorm:"column:shortDescription;not null" json:"shortDescription"`
	Description      string                   `gorm:"column:description;not null" json:"description"`
	Approved         bool                     `gorm:"column:approved;default:false;not null" json:"approved"`
	CreatedAt        DateTime                 `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime                 `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt        *DateTime                `gorm:"column:deletedAt;index:SocialMediaAgency_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	Logo             *Media                   `gorm:"foreignKey:LogoID;belongsTo"`
	User             *User                    `gorm:"belongsTo"`
	Niches           []SocialMediaAgencyNiche `gorm:"foreignKey:AgencyID;constraint:SocialMediaAgencyNiche_agencyId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (SocialMediaAgency) TableName() string {
	return "SocialMediaAgency"
}

func (m *SocialMediaAgency) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SocialMediaAgency) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SocialMediaAgencyNiche struct {
	AgencyID string             `gorm:"column:agencyId;primaryKey" json:"agencyId"`
	Niche    string             `gorm:"column:niche;primaryKey" json:"niche"`
	Agency   *SocialMediaAgency `gorm:"foreignKey:AgencyID;belongsTo"`
}

func (SocialMediaAgencyNiche) TableName() string {
	return "SocialMediaAgencyNiche"
}

type Credits struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string        `gorm:"column:organizationId;index:Credits_organizationId_idx;not null" json:"organizationId"`
	Credits        int           `gorm:"column:credits;size:32;not null" json:"credits"`
	CreatedAt      DateTime      `gorm:"column:createdAt;index:Credits_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Type           *string       `gorm:"column:type;default:'ai_images';not null" json:"type"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (Credits) TableName() string {
	return "Credits"
}

func (m *Credits) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Credits) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Clipping struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string         `gorm:"column:organizationId;index:Clipping_organizationId_idx;not null" json:"organizationId"`
	URL            string         `gorm:"column:url;not null" json:"url"`
	Status         *string        `gorm:"column:status;default:'analysing';not null" json:"status"`
	Error          *string        `gorm:"column:error" json:"error"`
	Title          *string        `gorm:"column:title" json:"title"`
	Thumbnail      *string        `gorm:"column:thumbnail" json:"thumbnail"`
	Duration       *int           `gorm:"column:duration;size:32" json:"duration"`
	MaxClips       *int           `gorm:"column:maxClips;size:32;default:5;not null" json:"maxClips"`
	Fit            *string        `gorm:"column:fit;default:'blur';not null" json:"fit"`
	Integrations   string         `gorm:"column:integrations;not null" json:"integrations"`
	CreditsID      *string        `gorm:"column:creditsId" json:"creditsId"`
	DeletedAt      *DateTime      `gorm:"column:deletedAt;index:Clipping_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization  `gorm:"belongsTo"`
	Clips          []ClippingClip `gorm:"foreignKey:ClippingID;constraint:ClippingClip_clippingId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Clipping) TableName() string {
	return "Clipping"
}

func (m *Clipping) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Clipping) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ClippingClip struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	ClippingID string    `gorm:"column:clippingId;index:ClippingClip_clippingId_idx;not null" json:"clippingId"`
	Title      string    `gorm:"column:title;not null" json:"title"`
	Content    string    `gorm:"column:content;not null" json:"content"`
	Start      float64   `gorm:"column:start;type:double precision;not null" json:"start"`
	End        float64   `gorm:"column:end;type:double precision;not null" json:"end"`
	TrimStart  *int      `gorm:"column:trimStart;size:32" json:"trimStart"`
	Status     *string   `gorm:"column:status;default:'pending';not null" json:"status"`
	Error      *string   `gorm:"column:error" json:"error"`
	MediaID    *string   `gorm:"column:mediaId" json:"mediaId"`
	Path       *string   `gorm:"column:path" json:"path"`
	Thumbnail  *string   `gorm:"column:thumbnail" json:"thumbnail"`
	DraftedAt  *DateTime `gorm:"column:draftedAt;type:timestamp(3)" json:"draftedAt"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Clipping   *Clipping `gorm:"belongsTo"`
}

func (ClippingClip) TableName() string {
	return "ClippingClip"
}

func (m *ClippingClip) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ClippingClip) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Subscription struct {
	ID               string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID   string        `gorm:"column:organizationId;uniqueIndex:Subscription_organizationId_key;index:Subscription_organizationId_idx;not null" json:"organizationId"`
	SubscriptionTier string        `gorm:"column:subscriptionTier;not null" json:"subscriptionTier"`
	Provider         *string       `gorm:"column:provider;default:'stripe';not null" json:"provider"`
	Identifier       *string       `gorm:"column:identifier" json:"identifier"`
	CancelAt         *DateTime     `gorm:"column:cancelAt;type:timestamp(3)" json:"cancelAt"`
	Period           string        `gorm:"column:period;not null" json:"period"`
	TotalChannels    int           `gorm:"column:totalChannels;size:32;not null" json:"totalChannels"`
	IsLifetime       bool          `gorm:"column:isLifetime;default:false;not null" json:"isLifetime"`
	CreatedAt        DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt        *DateTime     `gorm:"column:deletedAt;index:Subscription_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	Organization     *Organization `gorm:"belongsTo"`
}

func (Subscription) TableName() string {
	return "Subscription"
}

func (m *Subscription) BeforeCreate(tx *gorm.DB) error {
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

func (*Subscription) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Customer struct {
	ID           string        `gorm:"column:id;primaryKey" json:"id"`
	Name         string        `gorm:"column:name;uniqueIndex:Customer_orgId_name_deletedAt_key,priority:2;not null" json:"name"`
	OrgID        string        `gorm:"column:orgId;uniqueIndex:Customer_orgId_name_deletedAt_key,priority:1;not null" json:"orgId"`
	CreatedAt    DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt    *DateTime     `gorm:"column:deletedAt;uniqueIndex:Customer_orgId_name_deletedAt_key,priority:3;type:timestamp(3)" json:"deletedAt"`
	Organization *Organization `gorm:"foreignKey:OrgID;belongsTo"`
	Integrations []Integration `gorm:"foreignKey:CustomerID;constraint:Integration_customerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Customer) TableName() string {
	return "Customer"
}

func (m *Customer) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
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

type Integration struct {
	ID                    string                 `gorm:"column:id;primaryKey" json:"id"`
	InternalID            string                 `gorm:"column:internalId;uniqueIndex:Integration_organizationId_internalId_key,priority:2;not null" json:"internalId"`
	OrganizationID        string                 `gorm:"column:organizationId;uniqueIndex:Integration_organizationId_internalId_key,priority:1;index:Integration_organizationId_idx;not null" json:"organizationId"`
	Name                  string                 `gorm:"column:name;not null" json:"name"`
	CustomName            *string                `gorm:"column:customName" json:"customName"`
	Picture               *string                `gorm:"column:picture" json:"picture"`
	ProviderIdentifier    string                 `gorm:"column:providerIdentifier;index:Integration_providerIdentifier_idx;not null" json:"providerIdentifier"`
	Type                  string                 `gorm:"column:type;not null" json:"type"`
	Token                 string                 `gorm:"column:token;not null" json:"token"`
	Disabled              bool                   `gorm:"column:disabled;index:Integration_disabled_idx;default:false;not null" json:"disabled"`
	TokenExpiration       *DateTime              `gorm:"column:tokenExpiration;type:timestamp(3)" json:"tokenExpiration"`
	RefreshToken          *string                `gorm:"column:refreshToken" json:"refreshToken"`
	Profile               *string                `gorm:"column:profile" json:"profile"`
	DeletedAt             *DateTime              `gorm:"column:deletedAt;index:Integration_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	CreatedAt             DateTime               `gorm:"column:createdAt;index:Integration_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             *DateTime              `gorm:"column:updatedAt;index:Integration_updatedAt_idx;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	InBetweenSteps        bool                   `gorm:"column:inBetweenSteps;index:Integration_inBetweenSteps_idx;default:false;not null" json:"inBetweenSteps"`
	RefreshNeeded         bool                   `gorm:"column:refreshNeeded;index:Integration_refreshNeeded_idx;default:false;not null" json:"refreshNeeded"`
	PostingTimes          *string                `gorm:"column:postingTimes;default:'[{\"time\":120}, {\"time\":400}, {\"time\":700}]';not null" json:"postingTimes"`
	CustomInstanceDetails *string                `gorm:"column:customInstanceDetails" json:"customInstanceDetails"`
	CustomerID            *string                `gorm:"column:customerId;index:Integration_customerId_idx" json:"customerId"`
	RootInternalID        *string                `gorm:"column:rootInternalId;index:Integration_rootInternalId_idx" json:"rootInternalId"`
	AdditionalSettings    *string                `gorm:"column:additionalSettings;default:'[]'" json:"additionalSettings"`
	Customer              *Customer              `gorm:"belongsTo"`
	Organization          *Organization          `gorm:"belongsTo"`
	ExisingPlugData       []ExisingPlugData      `gorm:"foreignKey:IntegrationID;constraint:ExisingPlugData_integrationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Webhooks              []IntegrationsWebhooks `gorm:"foreignKey:IntegrationID;constraint:IntegrationsWebhooks_integrationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OrderItems            []OrderItems           `gorm:"foreignKey:IntegrationID;constraint:OrderItems_integrationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Plugs                 []Plugs                `gorm:"foreignKey:IntegrationID;constraint:Plugs_integrationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Posts                 []Post                 `gorm:"foreignKey:IntegrationID;constraint:Post_integrationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
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
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	return nil
}

func (*Integration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type Signatures struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string        `gorm:"column:organizationId;index:Signatures_organizationId_idx;not null" json:"organizationId"`
	Content        string        `gorm:"column:content;not null" json:"content"`
	AutoAdd        bool          `gorm:"column:autoAdd;not null" json:"autoAdd"`
	CreatedAt      DateTime      `gorm:"column:createdAt;index:Signatures_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt      *DateTime     `gorm:"column:deletedAt;index:Signatures_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (Signatures) TableName() string {
	return "Signatures"
}

func (m *Signatures) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Signatures) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Comments struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	Content        string        `gorm:"column:content;not null" json:"content"`
	OrganizationID string        `gorm:"column:organizationId;index:Comments_organizationId_idx;not null" json:"organizationId"`
	PostID         string        `gorm:"column:postId;index:Comments_postId_idx;not null" json:"postId"`
	UserID         *string       `gorm:"column:userId;index:Comments_userId_idx" json:"userId"`
	DisplayName    *string       `gorm:"column:displayName" json:"displayName"`
	ParentID       *string       `gorm:"column:parentId;index:Comments_parentId_idx" json:"parentId"`
	AnchorStart    *int          `gorm:"column:anchorStart;size:32" json:"anchorStart"`
	AnchorEnd      *int          `gorm:"column:anchorEnd;size:32" json:"anchorEnd"`
	AnchorQuote    *string       `gorm:"column:anchorQuote" json:"anchorQuote"`
	ResolvedAt     *DateTime     `gorm:"column:resolvedAt;index:Comments_resolvedAt_idx;type:timestamp(3)" json:"resolvedAt"`
	CreatedAt      DateTime      `gorm:"column:createdAt;index:Comments_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt      *DateTime     `gorm:"column:deletedAt;index:Comments_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	Organization   *Organization `gorm:"belongsTo"`
	Post           *Post         `gorm:"belongsTo"`
	User           *User         `gorm:"belongsTo"`
	Parent         *Comments     `gorm:"foreignKey:ParentID;belongsTo"`
	Replies        []Comments    `gorm:"foreignKey:ParentID;constraint:Comments_parentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Comments) TableName() string {
	return "Comments"
}

func (m *Comments) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Comments) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Post struct {
	ID                         string           `gorm:"column:id;primaryKey" json:"id"`
	State                      string           `gorm:"column:state;index:Post_publishDate_state_deletedAt_parentPostId_integrationId_idx,priority:2;index:Post_state_idx;index:Post_organizationId_publishDate_state_deletedAt_idx,priority:3;default:'QUEUE';not null" json:"state"`
	PublishDate                DateTime         `gorm:"column:publishDate;index:Post_publishDate_idx;index:Post_publishDate_state_deletedAt_parentPostId_integrationId_idx,priority:1;index:Post_organizationId_publishDate_state_deletedAt_idx,priority:2;type:timestamp(3);not null" json:"publishDate"`
	OrganizationID             string           `gorm:"column:organizationId;index:Post_publishDate_state_deletedAt_parentPostId_integrationId_idx,priority:6;index:Post_organizationId_idx;index:Post_organizationId_publishDate_state_deletedAt_idx,priority:1;index:Post_organizationId_intervalInDays_idx,priority:1;not null" json:"organizationId"`
	IntegrationID              string           `gorm:"column:integrationId;index:Post_publishDate_state_deletedAt_parentPostId_integrationId_idx,priority:5;index:Post_integrationId_idx;not null" json:"integrationId"`
	Content                    string           `gorm:"column:content;not null" json:"content"`
	Delay                      int              `gorm:"column:delay;size:32;default:0;not null" json:"delay"`
	Group                      string           `gorm:"column:group;index:Post_group_idx;not null" json:"group"`
	BatchID                    *string          `gorm:"column:batchId;index:Post_batchId_idx" json:"batchId"`
	Title                      *string          `gorm:"column:title" json:"title"`
	Description                *string          `gorm:"column:description" json:"description"`
	ParentPostID               *string          `gorm:"column:parentPostId;index:Post_publishDate_state_deletedAt_parentPostId_integrationId_idx,priority:4;index:Post_parentPostId_idx" json:"parentPostId"`
	ReleaseID                  *string          `gorm:"column:releaseId" json:"releaseId"`
	ReleaseURL                 *string          `gorm:"column:releaseURL" json:"releaseURL"`
	Settings                   *string          `gorm:"column:settings" json:"settings"`
	Image                      *string          `gorm:"column:image" json:"image"`
	SubmittedForOrderID        *string          `gorm:"column:submittedForOrderId;index:Post_submittedForOrderId_idx" json:"submittedForOrderId"`
	SubmittedForOrganizationID *string          `gorm:"column:submittedForOrganizationId" json:"submittedForOrganizationId"`
	ApprovedSubmitForOrder     string           `gorm:"column:approvedSubmitForOrder;default:'NO';not null" json:"approvedSubmitForOrder"`
	CreationMethod             string           `gorm:"column:creationMethod;index:Post_publishDate_state_deletedAt_parentPostId_integrationId_idx,priority:7;default:'UNKNOWN';not null" json:"creationMethod"`
	LastMessageID              *string          `gorm:"column:lastMessageId;index:Post_lastMessageId_idx" json:"lastMessageId"`
	IntervalInDays             *int             `gorm:"column:intervalInDays;index:Post_organizationId_intervalInDays_idx,priority:2;index:Post_intervalInDays_idx;size:32" json:"intervalInDays"`
	Error                      *string          `gorm:"column:error" json:"error"`
	CreatedAt                  DateTime         `gorm:"column:createdAt;index:Post_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime         `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt                  *DateTime        `gorm:"column:deletedAt;index:Post_publishDate_state_deletedAt_parentPostId_integrationId_idx,priority:3;index:Post_organizationId_publishDate_state_deletedAt_idx,priority:4;type:timestamp(3)" json:"deletedAt"`
	Integration                *Integration     `gorm:"belongsTo"`
	LastMessage                *Messages        `gorm:"foreignKey:LastMessageID;belongsTo"`
	Organization               *Organization    `gorm:"foreignKey:OrganizationID;belongsTo"`
	ParentPost                 *Post            `gorm:"foreignKey:ParentPostID;belongsTo"`
	SubmittedForOrder          *Orders          `gorm:"foreignKey:SubmittedForOrderID;belongsTo"`
	SubmittedForOrganization   *Organization    `gorm:"foreignKey:SubmittedForOrganizationID;belongsTo"`
	Comments                   []Comments       `gorm:"foreignKey:PostID;constraint:Comments_postId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Errors                     []Errors         `gorm:"foreignKey:PostID;constraint:Errors_postId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PayoutProblems             []PayoutProblems `gorm:"foreignKey:PostID;constraint:PayoutProblems_postId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ChildrenPost               []Post           `gorm:"foreignKey:ParentPostID;constraint:Post_parentPostId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tags                       []TagsPosts      `gorm:"foreignKey:PostID;constraint:TagsPosts_postId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Post) TableName() string {
	return "Post"
}

func (m *Post) BeforeCreate(tx *gorm.DB) error {
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

func (*Post) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Notifications struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string        `gorm:"column:organizationId;index:Notifications_organizationId_idx;index:Notifications_organizationId_createdAt_deletedAt_idx,priority:1;not null" json:"organizationId"`
	Content        string        `gorm:"column:content;not null" json:"content"`
	Link           *string       `gorm:"column:link" json:"link"`
	CreatedAt      DateTime      `gorm:"column:createdAt;index:Notifications_createdAt_idx;index:Notifications_organizationId_createdAt_deletedAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt      *DateTime     `gorm:"column:deletedAt;index:Notifications_organizationId_createdAt_deletedAt_idx,priority:3;type:timestamp(3)" json:"deletedAt"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (Notifications) TableName() string {
	return "Notifications"
}

func (m *Notifications) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Notifications) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type MessagesGroup struct {
	ID                  string        `gorm:"column:id;primaryKey" json:"id"`
	BuyerOrganizationID string        `gorm:"column:buyerOrganizationId;index:MessagesGroup_buyerOrganizationId_idx;not null" json:"buyerOrganizationId"`
	BuyerID             string        `gorm:"column:buyerId;uniqueIndex:MessagesGroup_buyerId_sellerId_key,priority:1;not null" json:"buyerId"`
	SellerID            string        `gorm:"column:sellerId;uniqueIndex:MessagesGroup_buyerId_sellerId_key,priority:2;not null" json:"sellerId"`
	CreatedAt           DateTime      `gorm:"column:createdAt;index:MessagesGroup_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime      `gorm:"column:updatedAt;index:MessagesGroup_updatedAt_idx;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Buyer               *User         `gorm:"foreignKey:BuyerID;belongsTo"`
	BuyerOrganization   *Organization `gorm:"foreignKey:BuyerOrganizationID;belongsTo"`
	Seller              *User         `gorm:"foreignKey:SellerID;belongsTo"`
	Messages            []Messages    `gorm:"foreignKey:GroupID;constraint:Messages_groupId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Orders              []Orders      `gorm:"foreignKey:MessageGroupID;constraint:Orders_messageGroupId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (MessagesGroup) TableName() string {
	return "MessagesGroup"
}

func (m *MessagesGroup) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*MessagesGroup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PayoutProblems struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Status    string   `gorm:"column:status;not null" json:"status"`
	OrderID   string   `gorm:"column:orderId;not null" json:"orderId"`
	UserID    string   `gorm:"column:userId;not null" json:"userId"`
	PostID    *string  `gorm:"column:postId" json:"postId"`
	Amount    int      `gorm:"column:amount;size:32;not null" json:"amount"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Order     *Orders  `gorm:"foreignKey:OrderID;belongsTo"`
	Post      *Post    `gorm:"belongsTo"`
	User      *User    `gorm:"belongsTo"`
}

func (PayoutProblems) TableName() string {
	return "PayoutProblems"
}

func (m *PayoutProblems) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PayoutProblems) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Orders struct {
	ID             string           `gorm:"column:id;primaryKey" json:"id"`
	BuyerID        string           `gorm:"column:buyerId;index:Orders_buyerId_idx;not null" json:"buyerId"`
	SellerID       string           `gorm:"column:sellerId;index:Orders_sellerId_idx;not null" json:"sellerId"`
	Status         string           `gorm:"column:status;not null" json:"status"`
	MessageGroupID string           `gorm:"column:messageGroupId;index:Orders_messageGroupId_idx;not null" json:"messageGroupId"`
	CaptureID      *string          `gorm:"column:captureId" json:"captureId"`
	CreatedAt      DateTime         `gorm:"column:createdAt;index:Orders_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime         `gorm:"column:updatedAt;index:Orders_updatedAt_idx;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Buyer          *User            `gorm:"foreignKey:BuyerID;belongsTo"`
	MessageGroup   *MessagesGroup   `gorm:"foreignKey:MessageGroupID;belongsTo"`
	Seller         *User            `gorm:"foreignKey:SellerID;belongsTo"`
	OrdersItems    []OrderItems     `gorm:"foreignKey:OrderID;constraint:OrderItems_orderId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PayoutProblems []PayoutProblems `gorm:"foreignKey:OrderID;constraint:PayoutProblems_orderId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Posts          []Post           `gorm:"foreignKey:SubmittedForOrderID;constraint:Post_submittedForOrderId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Orders) TableName() string {
	return "Orders"
}

func (m *Orders) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Orders) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrderItems struct {
	ID            string       `gorm:"column:id;primaryKey" json:"id"`
	OrderID       string       `gorm:"column:orderId;index:OrderItems_orderId_idx;not null" json:"orderId"`
	IntegrationID string       `gorm:"column:integrationId;index:OrderItems_integrationId_idx;not null" json:"integrationId"`
	Quantity      int          `gorm:"column:quantity;size:32;not null" json:"quantity"`
	Price         int          `gorm:"column:price;size:32;not null" json:"price"`
	Integration   *Integration `gorm:"belongsTo"`
	Order         *Orders      `gorm:"foreignKey:OrderID;belongsTo"`
}

func (OrderItems) TableName() string {
	return "OrderItems"
}

func (m *OrderItems) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type Messages struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	From      string         `gorm:"column:from;not null" json:"from"`
	Content   *string        `gorm:"column:content" json:"content"`
	GroupID   string         `gorm:"column:groupId;index:Messages_groupId_idx;not null" json:"groupId"`
	Special   *string        `gorm:"column:special" json:"special"`
	CreatedAt DateTime       `gorm:"column:createdAt;index:Messages_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt *DateTime      `gorm:"column:deletedAt;index:Messages_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	Group     *MessagesGroup `gorm:"foreignKey:GroupID;belongsTo"`
	Posts     []Post         `gorm:"foreignKey:LastMessageID;constraint:Post_lastMessageId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Messages) TableName() string {
	return "Messages"
}

func (m *Messages) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Messages) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Plugs struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string        `gorm:"column:organizationId;index:Plugs_organizationId_idx;not null" json:"organizationId"`
	PlugFunction   string        `gorm:"column:plugFunction;uniqueIndex:Plugs_plugFunction_integrationId_key,priority:1;not null" json:"plugFunction"`
	Data           string        `gorm:"column:data;not null" json:"data"`
	IntegrationID  string        `gorm:"column:integrationId;uniqueIndex:Plugs_plugFunction_integrationId_key,priority:2;not null" json:"integrationId"`
	Activated      *bool         `gorm:"column:activated;default:true;not null" json:"activated"`
	Integration    *Integration  `gorm:"belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (Plugs) TableName() string {
	return "Plugs"
}

func (m *Plugs) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type ExisingPlugData struct {
	ID            string       `gorm:"column:id;primaryKey" json:"id"`
	IntegrationID string       `gorm:"column:integrationId;uniqueIndex:ExisingPlugData_integrationId_methodName_value_key,priority:1;not null" json:"integrationId"`
	MethodName    string       `gorm:"column:methodName;uniqueIndex:ExisingPlugData_integrationId_methodName_value_key,priority:2;not null" json:"methodName"`
	Value         string       `gorm:"column:value;uniqueIndex:ExisingPlugData_integrationId_methodName_value_key,priority:3;not null" json:"value"`
	Integration   *Integration `gorm:"belongsTo"`
}

func (ExisingPlugData) TableName() string {
	return "ExisingPlugData"
}

func (m *ExisingPlugData) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type PopularPosts struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Category  string   `gorm:"column:category;not null" json:"category"`
	Topic     string   `gorm:"column:topic;not null" json:"topic"`
	Content   string   `gorm:"column:content;not null" json:"content"`
	Hook      string   `gorm:"column:hook;not null" json:"hook"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (PopularPosts) TableName() string {
	return "PopularPosts"
}

func (m *PopularPosts) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PopularPosts) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type IntegrationsWebhooks struct {
	IntegrationID string       `gorm:"column:integrationId;primaryKey;uniqueIndex:IntegrationsWebhooks_integrationId_webhookId_key,priority:1;index:IntegrationsWebhooks_integrationId_idx" json:"integrationId"`
	WebhookID     string       `gorm:"column:webhookId;primaryKey;uniqueIndex:IntegrationsWebhooks_integrationId_webhookId_key,priority:2;index:IntegrationsWebhooks_webhookId_idx" json:"webhookId"`
	Integration   *Integration `gorm:"belongsTo"`
	Webhook       *Webhooks    `gorm:"foreignKey:WebhookID;belongsTo"`
}

func (IntegrationsWebhooks) TableName() string {
	return "IntegrationsWebhooks"
}

type Webhooks struct {
	ID             string                 `gorm:"column:id;primaryKey" json:"id"`
	Name           string                 `gorm:"column:name;not null" json:"name"`
	OrganizationID string                 `gorm:"column:organizationId;index:Webhooks_organizationId_idx;not null" json:"organizationId"`
	URL            string                 `gorm:"column:url;not null" json:"url"`
	DeletedAt      *DateTime              `gorm:"column:deletedAt;index:Webhooks_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	CreatedAt      DateTime               `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime               `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization          `gorm:"belongsTo"`
	Integrations   []IntegrationsWebhooks `gorm:"foreignKey:WebhookID;constraint:IntegrationsWebhooks_webhookId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Webhooks) TableName() string {
	return "Webhooks"
}

func (m *Webhooks) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Webhooks) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AutoPost struct {
	ID              string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID  string        `gorm:"column:organizationId;not null" json:"organizationId"`
	Title           string        `gorm:"column:title;not null" json:"title"`
	Content         *string       `gorm:"column:content" json:"content"`
	OnSlot          bool          `gorm:"column:onSlot;not null" json:"onSlot"`
	SyncLast        bool          `gorm:"column:syncLast;not null" json:"syncLast"`
	URL             string        `gorm:"column:url;not null" json:"url"`
	LastURL         string        `gorm:"column:lastUrl;not null" json:"lastUrl"`
	Active          bool          `gorm:"column:active;not null" json:"active"`
	AddPicture      bool          `gorm:"column:addPicture;not null" json:"addPicture"`
	GenerateContent bool          `gorm:"column:generateContent;not null" json:"generateContent"`
	Integrations    string        `gorm:"column:integrations;not null" json:"integrations"`
	DeletedAt       *DateTime     `gorm:"column:deletedAt;index:AutoPost_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	CreatedAt       DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization    *Organization `gorm:"belongsTo"`
}

func (AutoPost) TableName() string {
	return "AutoPost"
}

func (m *AutoPost) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AutoPost) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Sets struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string        `gorm:"column:organizationId;index:Sets_organizationId_idx;not null" json:"organizationId"`
	Name           string        `gorm:"column:name;not null" json:"name"`
	Content        string        `gorm:"column:content;not null" json:"content"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (Sets) TableName() string {
	return "Sets"
}

func (m *Sets) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Sets) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ThirdParty struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string        `gorm:"column:organizationId;uniqueIndex:ThirdParty_organizationId_internalId_key,priority:1;index:ThirdParty_organizationId_idx;not null" json:"organizationId"`
	Identifier     string        `gorm:"column:identifier;not null" json:"identifier"`
	Name           string        `gorm:"column:name;not null" json:"name"`
	InternalID     string        `gorm:"column:internalId;uniqueIndex:ThirdParty_organizationId_internalId_key,priority:2;not null" json:"internalId"`
	APIKey         string        `gorm:"column:apiKey;not null" json:"apiKey"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt      *DateTime     `gorm:"column:deletedAt;index:ThirdParty_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (ThirdParty) TableName() string {
	return "ThirdParty"
}

func (m *ThirdParty) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ThirdParty) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Errors struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	Message        string        `gorm:"column:message;not null" json:"message"`
	Platform       string        `gorm:"column:platform;index:Errors_createdAt_platform_idx,priority:2;not null" json:"platform"`
	OrganizationID string        `gorm:"column:organizationId;index:Errors_organizationId_idx;not null" json:"organizationId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;index:Errors_createdAt_idx;index:Errors_createdAt_platform_idx,priority:1;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	PostID         string        `gorm:"column:postId;not null" json:"postId"`
	Body           *string       `gorm:"column:body;default:'{}';not null" json:"body"`
	Organization   *Organization `gorm:"belongsTo"`
	Post           *Post         `gorm:"belongsTo"`
}

func (Errors) TableName() string {
	return "Errors"
}

func (m *Errors) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Errors) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Mentions struct {
	Name      string   `gorm:"column:name;primaryKey" json:"name"`
	Username  string   `gorm:"column:username;primaryKey" json:"username"`
	Platform  string   `gorm:"column:platform;primaryKey" json:"platform"`
	CreatedAt DateTime `gorm:"column:createdAt;index:Mentions_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Image     string   `gorm:"column:image;primaryKey" json:"image"`
}

func (Mentions) TableName() string {
	return "Mentions"
}

func (m *Mentions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Mentions) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type MastraMessages struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	ThreadID   string    `gorm:"column:thread_id;index:public_mastra_messages_thread_id_createdat_idx,priority:1;index:mastra_messages_thread_id_createdat_idx,priority:1;not null" json:"thread_id"`
	Content    string    `gorm:"column:content;not null" json:"content"`
	Role       string    `gorm:"column:role;not null" json:"role"`
	Type       string    `gorm:"column:type;not null" json:"type"`
	CreatedAt  DateTime  `gorm:"column:createdAt;index:public_mastra_messages_thread_id_createdat_idx,priority:2,sort:desc;index:mastra_messages_thread_id_createdat_idx,priority:2,sort:desc;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	ResourceID *string   `gorm:"column:resourceId" json:"resourceId"`
	CreatedAtZ *DateTime `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraMessages) TableName() string {
	return "mastra_messages"
}

func (m *MastraMessages) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraResources struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	WorkingMemory *string        `gorm:"column:workingMemory" json:"workingMemory"`
	Metadata      datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ    *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ    *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraResources) TableName() string {
	return "mastra_resources"
}

func (m *MastraResources) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraScorers struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	ScorerID             string         `gorm:"column:scorerId;not null" json:"scorerId"`
	TraceID              *string        `gorm:"column:traceId;index:public_mastra_scores_trace_id_span_id_created_at_idx,priority:1;index:mastra_scores_trace_id_span_id_created_at_idx,priority:1" json:"traceId"`
	RunID                string         `gorm:"column:runId;not null" json:"runId"`
	Scorer               datatypes.JSON `gorm:"column:scorer;not null" json:"scorer"`
	PreprocessStepResult datatypes.JSON `gorm:"column:preprocessStepResult" json:"preprocessStepResult"`
	ExtractStepResult    datatypes.JSON `gorm:"column:extractStepResult" json:"extractStepResult"`
	AnalyzeStepResult    datatypes.JSON `gorm:"column:analyzeStepResult" json:"analyzeStepResult"`
	Score                float64        `gorm:"column:score;type:double precision;not null" json:"score"`
	Reason               *string        `gorm:"column:reason" json:"reason"`
	Metadata             datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	PreprocessPrompt     *string        `gorm:"column:preprocessPrompt" json:"preprocessPrompt"`
	ExtractPrompt        *string        `gorm:"column:extractPrompt" json:"extractPrompt"`
	GenerateScorePrompt  *string        `gorm:"column:generateScorePrompt" json:"generateScorePrompt"`
	GenerateReasonPrompt *string        `gorm:"column:generateReasonPrompt" json:"generateReasonPrompt"`
	AnalyzePrompt        *string        `gorm:"column:analyzePrompt" json:"analyzePrompt"`
	ReasonPrompt         *string        `gorm:"column:reasonPrompt" json:"reasonPrompt"`
	Input                datatypes.JSON `gorm:"column:input;not null" json:"input"`
	Output               datatypes.JSON `gorm:"column:output;not null" json:"output"`
	AdditionalContext    datatypes.JSON `gorm:"column:additionalContext" json:"additionalContext"`
	RuntimeContext       datatypes.JSON `gorm:"column:runtimeContext" json:"runtimeContext"`
	EntityType           *string        `gorm:"column:entityType" json:"entityType"`
	Entity               datatypes.JSON `gorm:"column:entity" json:"entity"`
	EntityID             *string        `gorm:"column:entityId" json:"entityId"`
	Source               string         `gorm:"column:source;not null" json:"source"`
	ResourceID           *string        `gorm:"column:resourceId" json:"resourceId"`
	ThreadID             *string        `gorm:"column:threadId" json:"threadId"`
	CreatedAt            DateTime       `gorm:"column:createdAt;index:public_mastra_scores_trace_id_span_id_created_at_idx,priority:3,sort:desc;index:mastra_scores_trace_id_span_id_created_at_idx,priority:3,sort:desc;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ           *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ           *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
	SpanID               *string        `gorm:"column:spanId;index:public_mastra_scores_trace_id_span_id_created_at_idx,priority:2;index:mastra_scores_trace_id_span_id_created_at_idx,priority:2" json:"spanId"`
	RequestContext       datatypes.JSON `gorm:"column:requestContext" json:"requestContext"`
	OrganizationID       *string        `gorm:"column:organizationId" json:"organizationId"`
	ProjectID            *string        `gorm:"column:projectId" json:"projectId"`
	BatchID              *string        `gorm:"column:batchId" json:"batchId"`
	DatasetID            *string        `gorm:"column:datasetId" json:"datasetId"`
	DatasetItemID        *string        `gorm:"column:datasetItemId" json:"datasetItemId"`
}

func (MastraScorers) TableName() string {
	return "mastra_scorers"
}

func (m *MastraScorers) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraThreads struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	ResourceID string    `gorm:"column:resourceId;index:public_mastra_threads_resourceid_createdat_idx,priority:1;index:mastra_threads_resourceid_createdat_idx,priority:1;not null" json:"resourceId"`
	Title      string    `gorm:"column:title;not null" json:"title"`
	Metadata   *string   `gorm:"column:metadata" json:"metadata"`
	CreatedAt  DateTime  `gorm:"column:createdAt;index:public_mastra_threads_resourceid_createdat_idx,priority:2,sort:desc;index:mastra_threads_resourceid_createdat_idx,priority:2,sort:desc;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ *DateTime `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ *DateTime `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraThreads) TableName() string {
	return "mastra_threads"
}

func (m *MastraThreads) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraTraces struct {
	ID           string         `gorm:"column:id;primaryKey" json:"id"`
	ParentSpanID *string        `gorm:"column:parentSpanId" json:"parentSpanId"`
	Name         string         `gorm:"column:name;index:public_mastra_traces_name_starttime_idx,priority:1;not null" json:"name"`
	TraceID      string         `gorm:"column:traceId;not null" json:"traceId"`
	Scope        string         `gorm:"column:scope;not null" json:"scope"`
	Kind         int            `gorm:"column:kind;size:32;not null" json:"kind"`
	Attributes   datatypes.JSON `gorm:"column:attributes" json:"attributes"`
	Status       datatypes.JSON `gorm:"column:status" json:"status"`
	Events       datatypes.JSON `gorm:"column:events" json:"events"`
	Links        datatypes.JSON `gorm:"column:links" json:"links"`
	Other        *string        `gorm:"column:other" json:"other"`
	StartTime    int64          `gorm:"column:startTime;index:public_mastra_traces_name_starttime_idx,priority:2,sort:desc;not null" json:"startTime"`
	EndTime      int64          `gorm:"column:endTime;not null" json:"endTime"`
	CreatedAt    DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ   *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraTraces) TableName() string {
	return "mastra_traces"
}

func (m *MastraTraces) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraWorkflowSnapshot struct {
	WorkflowName string    `gorm:"column:workflow_name;primaryKey;uniqueIndex:public_mastra_workflow_snapshot_workflow_name_run_id_key,priority:1;index:mastra_workflow_snapshot_name_createdat_idx,priority:1" json:"workflow_name"`
	RunID        string    `gorm:"column:run_id;primaryKey;uniqueIndex:public_mastra_workflow_snapshot_workflow_name_run_id_key,priority:2" json:"run_id"`
	ResourceID   *string   `gorm:"column:resourceId" json:"resourceId"`
	Snapshot     string    `gorm:"column:snapshot;not null" json:"snapshot"`
	CreatedAt    DateTime  `gorm:"column:createdAt;index:mastra_workflow_snapshot_name_createdat_idx,priority:2,sort:desc;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime  `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ   *DateTime `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ   *DateTime `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraWorkflowSnapshot) TableName() string {
	return "mastra_workflow_snapshot"
}

func (m *MastraWorkflowSnapshot) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraAgentVersions struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	AgentID              string         `gorm:"column:agentId;not null" json:"agentId"`
	VersionNumber        int            `gorm:"column:versionNumber;size:32;not null" json:"versionNumber"`
	Name                 string         `gorm:"column:name;not null" json:"name"`
	Description          *string        `gorm:"column:description" json:"description"`
	Instructions         string         `gorm:"column:instructions;not null" json:"instructions"`
	Model                datatypes.JSON `gorm:"column:model;not null" json:"model"`
	Tools                datatypes.JSON `gorm:"column:tools" json:"tools"`
	DefaultOptions       datatypes.JSON `gorm:"column:defaultOptions" json:"defaultOptions"`
	Workflows            datatypes.JSON `gorm:"column:workflows" json:"workflows"`
	Agents               datatypes.JSON `gorm:"column:agents" json:"agents"`
	IntegrationTools     datatypes.JSON `gorm:"column:integrationTools" json:"integrationTools"`
	InputProcessors      datatypes.JSON `gorm:"column:inputProcessors" json:"inputProcessors"`
	OutputProcessors     datatypes.JSON `gorm:"column:outputProcessors" json:"outputProcessors"`
	Memory               datatypes.JSON `gorm:"column:memory" json:"memory"`
	Scorers              datatypes.JSON `gorm:"column:scorers" json:"scorers"`
	McpClients           datatypes.JSON `gorm:"column:mcpClients" json:"mcpClients"`
	RequestContextSchema datatypes.JSON `gorm:"column:requestContextSchema" json:"requestContextSchema"`
	Workspace            datatypes.JSON `gorm:"column:workspace" json:"workspace"`
	Skills               datatypes.JSON `gorm:"column:skills" json:"skills"`
	SkillsFormat         *string        `gorm:"column:skillsFormat" json:"skillsFormat"`
	ChangedFields        datatypes.JSON `gorm:"column:changedFields" json:"changedFields"`
	ChangeMessage        *string        `gorm:"column:changeMessage" json:"changeMessage"`
	CreatedAt            DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ           *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	Durable              datatypes.JSON `gorm:"column:durable" json:"durable"`
	Browser              datatypes.JSON `gorm:"column:browser" json:"browser"`
	ToolProviders        datatypes.JSON `gorm:"column:toolProviders" json:"toolProviders"`
}

func (MastraAgentVersions) TableName() string {
	return "mastra_agent_versions"
}

func (m *MastraAgentVersions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraAgents struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Status          string         `gorm:"column:status;not null" json:"status"`
	ActiveVersionID *string        `gorm:"column:activeVersionId" json:"activeVersionId"`
	AuthorID        *string        `gorm:"column:authorId" json:"authorId"`
	Metadata        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ      *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ      *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
	Visibility      *string        `gorm:"column:visibility" json:"visibility"`
	FavoriteCount   *int           `gorm:"column:favoriteCount;size:32" json:"favoriteCount"`
}

func (MastraAgents) TableName() string {
	return "mastra_agents"
}

func (m *MastraAgents) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraBackgroundTasks struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	ToolCallID     string         `gorm:"column:tool_call_id;index:mastra_bg_tasks_tool_call_idx;not null" json:"tool_call_id"`
	ToolName       string         `gorm:"column:tool_name;not null" json:"tool_name"`
	AgentID        string         `gorm:"column:agent_id;index:mastra_bg_tasks_agent_status_idx,priority:1;not null" json:"agent_id"`
	RunID          string         `gorm:"column:run_id;not null" json:"run_id"`
	ThreadID       *string        `gorm:"column:thread_id;index:mastra_bg_tasks_thread_idx,priority:1" json:"thread_id"`
	ResourceID     *string        `gorm:"column:resource_id" json:"resource_id"`
	Status         string         `gorm:"column:status;index:mastra_bg_tasks_agent_status_idx,priority:2;index:mastra_bg_tasks_status_created_at_idx,priority:1;not null" json:"status"`
	Args           datatypes.JSON `gorm:"column:args;not null" json:"args"`
	Result         datatypes.JSON `gorm:"column:result" json:"result"`
	Error          datatypes.JSON `gorm:"column:error" json:"error"`
	SuspendPayload datatypes.JSON `gorm:"column:suspend_payload" json:"suspend_payload"`
	RetryCount     int            `gorm:"column:retry_count;size:32;not null" json:"retry_count"`
	MaxRetries     int            `gorm:"column:max_retries;size:32;not null" json:"max_retries"`
	TimeoutMs      int            `gorm:"column:timeout_ms;size:32;not null" json:"timeout_ms"`
	CreatedAt      DateTime       `gorm:"column:createdAt;index:mastra_bg_tasks_status_created_at_idx,priority:2;index:mastra_bg_tasks_thread_idx,priority:2;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	StartedAt      *DateTime      `gorm:"column:startedAt;type:timestamp(6)" json:"startedAt"`
	SuspendedAt    *DateTime      `gorm:"column:suspendedAt;type:timestamp(6)" json:"suspendedAt"`
	CompletedAt    *DateTime      `gorm:"column:completedAt;type:timestamp(6)" json:"completedAt"`
	CreatedAtZ     *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	StartedAtZ     *DateTime      `gorm:"column:startedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"startedAtZ"`
	SuspendedAtZ   *DateTime      `gorm:"column:suspendedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"suspendedAtZ"`
	CompletedAtZ   *DateTime      `gorm:"column:completedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"completedAtZ"`
}

func (MastraBackgroundTasks) TableName() string {
	return "mastra_background_tasks"
}

func (m *MastraBackgroundTasks) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.StartedAtZ == nil {
		m.StartedAtZ = &DateTime{Time: now}
	}
	if m.SuspendedAtZ == nil {
		m.SuspendedAtZ = &DateTime{Time: now}
	}
	if m.CompletedAtZ == nil {
		m.CompletedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraChannelConfig struct {
	Platform   string         `gorm:"column:platform;primaryKey" json:"platform"`
	Data       datatypes.JSON `gorm:"column:data;not null" json:"data"`
	UpdatedAt  DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	UpdatedAtZ *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraChannelConfig) TableName() string {
	return "mastra_channel_config"
}

func (m *MastraChannelConfig) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraChannelInstallations struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	Platform   string         `gorm:"column:platform;index:idx_channel_installations_platform_agent,priority:1;not null" json:"platform"`
	AgentID    string         `gorm:"column:agentId;index:idx_channel_installations_platform_agent,priority:2;not null" json:"agentId"`
	Status     string         `gorm:"column:status;not null" json:"status"`
	WebhookID  *string        `gorm:"column:webhookId;uniqueIndex:idx_channel_installations_webhook" json:"webhookId"`
	Data       datatypes.JSON `gorm:"column:data;not null" json:"data"`
	ConfigHash *string        `gorm:"column:configHash" json:"configHash"`
	Error      *string        `gorm:"column:error" json:"error"`
	CreatedAt  DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraChannelInstallations) TableName() string {
	return "mastra_channel_installations"
}

func (m *MastraChannelInstallations) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraDatasetItems struct {
	ID                 string         `gorm:"column:id;primaryKey" json:"id"`
	DatasetID          string         `gorm:"column:datasetId;index:idx_dataset_items_dataset_validto,priority:1;index:idx_dataset_items_dataset_validto_deleted,priority:1;index:idx_dataset_items_dataset_version,priority:1;index:idx_dataset_items_external_id_history,priority:1;not null" json:"datasetId"`
	DatasetVersion     int            `gorm:"column:datasetVersion;primaryKey;index:idx_dataset_items_dataset_version,priority:2;index:idx_dataset_items_external_id_history,priority:3;size:32" json:"datasetVersion"`
	ValidTo            *int           `gorm:"column:validTo;index:idx_dataset_items_dataset_validto,priority:2;index:idx_dataset_items_dataset_validto_deleted,priority:2;size:32" json:"validTo"`
	IsDeleted          bool           `gorm:"column:isDeleted;index:idx_dataset_items_dataset_validto_deleted,priority:3;not null" json:"isDeleted"`
	Input              datatypes.JSON `gorm:"column:input;not null" json:"input"`
	GroundTruth        datatypes.JSON `gorm:"column:groundTruth" json:"groundTruth"`
	RequestContext     datatypes.JSON `gorm:"column:requestContext" json:"requestContext"`
	Metadata           datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	Source             datatypes.JSON `gorm:"column:source" json:"source"`
	CreatedAt          DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ         *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ         *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
	ExpectedTrajectory datatypes.JSON `gorm:"column:expectedTrajectory" json:"expectedTrajectory"`
	OrganizationID     *string        `gorm:"column:organizationId;index:idx_dataset_items_org_project,priority:1" json:"organizationId"`
	ProjectID          *string        `gorm:"column:projectId;index:idx_dataset_items_org_project,priority:2" json:"projectId"`
	ToolMocks          datatypes.JSON `gorm:"column:toolMocks" json:"toolMocks"`
	UnmockedToolPolicy *string        `gorm:"column:unmockedToolPolicy" json:"unmockedToolPolicy"`
	ScorerIds          datatypes.JSON `gorm:"column:scorerIds" json:"scorerIds"`
	ExternalID         *string        `gorm:"column:externalId;index:idx_dataset_items_external_id_history,priority:2" json:"externalId"`
}

func (MastraDatasetItems) TableName() string {
	return "mastra_dataset_items"
}

func (m *MastraDatasetItems) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraDatasetVersions struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	DatasetID  string    `gorm:"column:datasetId;uniqueIndex:idx_dataset_versions_dataset_version_unique,priority:1;index:idx_dataset_versions_dataset_version,priority:1;not null" json:"datasetId"`
	Version    int       `gorm:"column:version;uniqueIndex:idx_dataset_versions_dataset_version_unique,priority:2;index:idx_dataset_versions_dataset_version,priority:2;size:32;not null" json:"version"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ *DateTime `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraDatasetVersions) TableName() string {
	return "mastra_dataset_versions"
}

func (m *MastraDatasetVersions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraDatasets struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	Name                 string         `gorm:"column:name;not null" json:"name"`
	Description          *string        `gorm:"column:description" json:"description"`
	Metadata             datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	InputSchema          datatypes.JSON `gorm:"column:inputSchema" json:"inputSchema"`
	GroundTruthSchema    datatypes.JSON `gorm:"column:groundTruthSchema" json:"groundTruthSchema"`
	RequestContextSchema datatypes.JSON `gorm:"column:requestContextSchema" json:"requestContextSchema"`
	Tags                 datatypes.JSON `gorm:"column:tags" json:"tags"`
	TargetType           *string        `gorm:"column:targetType" json:"targetType"`
	TargetIds            datatypes.JSON `gorm:"column:targetIds" json:"targetIds"`
	ScorerIds            datatypes.JSON `gorm:"column:scorerIds" json:"scorerIds"`
	Version              int            `gorm:"column:version;size:32;not null" json:"version"`
	CreatedAt            DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ           *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ           *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
	OrganizationID       *string        `gorm:"column:organizationId;index:idx_datasets_org_project,priority:1" json:"organizationId"`
	ProjectID            *string        `gorm:"column:projectId;index:idx_datasets_org_project,priority:2" json:"projectId"`
	CandidateKey         *string        `gorm:"column:candidateKey;index:idx_datasets_candidate,priority:1" json:"candidateKey"`
	CandidateID          *string        `gorm:"column:candidateId;index:idx_datasets_candidate,priority:2" json:"candidateId"`
}

func (MastraDatasets) TableName() string {
	return "mastra_datasets"
}

func (m *MastraDatasets) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraExperimentResults struct {
	ID                 string         `gorm:"column:id;primaryKey" json:"id"`
	ExperimentID       string         `gorm:"column:experimentId;uniqueIndex:idx_experiment_results_exp_item_attempt,priority:1;index:idx_experiment_results_experimentid;not null" json:"experimentId"`
	ItemID             string         `gorm:"column:itemId;uniqueIndex:idx_experiment_results_exp_item_attempt,priority:2;not null" json:"itemId"`
	ItemDatasetVersion *int           `gorm:"column:itemDatasetVersion;size:32" json:"itemDatasetVersion"`
	Input              datatypes.JSON `gorm:"column:input;not null" json:"input"`
	Output             datatypes.JSON `gorm:"column:output" json:"output"`
	GroundTruth        datatypes.JSON `gorm:"column:groundTruth" json:"groundTruth"`
	Error              datatypes.JSON `gorm:"column:error" json:"error"`
	StartedAt          DateTime       `gorm:"column:startedAt;type:timestamp(6);not null" json:"startedAt"`
	CompletedAt        DateTime       `gorm:"column:completedAt;type:timestamp(6);not null" json:"completedAt"`
	RetryCount         int            `gorm:"column:retryCount;size:32;not null" json:"retryCount"`
	TraceID            *string        `gorm:"column:traceId" json:"traceId"`
	Status             *string        `gorm:"column:status" json:"status"`
	Tags               datatypes.JSON `gorm:"column:tags;index:idx_experiment_results_tags_gin,type:gin" json:"tags"`
	CreatedAt          DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	StartedAtZ         *DateTime      `gorm:"column:startedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"startedAtZ"`
	CompletedAtZ       *DateTime      `gorm:"column:completedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"completedAtZ"`
	CreatedAtZ         *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	Comment            *string        `gorm:"column:comment" json:"comment"`
	ToolMockReport     datatypes.JSON `gorm:"column:toolMockReport" json:"toolMockReport"`
	Metadata           datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	OrganizationID     *string        `gorm:"column:organizationId;index:idx_experiment_results_org_project,priority:1" json:"organizationId"`
	ProjectID          *string        `gorm:"column:projectId;index:idx_experiment_results_org_project,priority:2" json:"projectId"`
	Attempt            *int           `gorm:"column:attempt;uniqueIndex:idx_experiment_results_exp_item_attempt,priority:3;size:32" json:"attempt"`
}

func (MastraExperimentResults) TableName() string {
	return "mastra_experiment_results"
}

func (m *MastraExperimentResults) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.StartedAtZ == nil {
		m.StartedAtZ = &DateTime{Time: now}
	}
	if m.CompletedAtZ == nil {
		m.CompletedAtZ = &DateTime{Time: now}
	}
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraExperiments struct {
	ID                string         `gorm:"column:id;primaryKey" json:"id"`
	Name              *string        `gorm:"column:name" json:"name"`
	Description       *string        `gorm:"column:description" json:"description"`
	Metadata          datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	DatasetID         *string        `gorm:"column:datasetId;index:idx_experiments_datasetid" json:"datasetId"`
	DatasetVersion    *int           `gorm:"column:datasetVersion;size:32" json:"datasetVersion"`
	TargetType        string         `gorm:"column:targetType;not null" json:"targetType"`
	TargetID          string         `gorm:"column:targetId;not null" json:"targetId"`
	Status            string         `gorm:"column:status;not null" json:"status"`
	TotalItems        int            `gorm:"column:totalItems;size:32;not null" json:"totalItems"`
	SucceededCount    int            `gorm:"column:succeededCount;size:32;not null" json:"succeededCount"`
	FailedCount       int            `gorm:"column:failedCount;size:32;not null" json:"failedCount"`
	SkippedCount      int            `gorm:"column:skippedCount;size:32;not null" json:"skippedCount"`
	StartedAt         *DateTime      `gorm:"column:startedAt;type:timestamp(6)" json:"startedAt"`
	CompletedAt       *DateTime      `gorm:"column:completedAt;type:timestamp(6)" json:"completedAt"`
	AgentVersion      *string        `gorm:"column:agentVersion" json:"agentVersion"`
	CreatedAt         DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	StartedAtZ        *DateTime      `gorm:"column:startedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"startedAtZ"`
	CompletedAtZ      *DateTime      `gorm:"column:completedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"completedAtZ"`
	CreatedAtZ        *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ        *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
	OrganizationID    *string        `gorm:"column:organizationId;index:idx_experiments_org_project,priority:1" json:"organizationId"`
	ProjectID         *string        `gorm:"column:projectId;index:idx_experiments_org_project,priority:2" json:"projectId"`
	Provenance        datatypes.JSON `gorm:"column:provenance" json:"provenance"`
	RunnerAttestation datatypes.JSON `gorm:"column:runnerAttestation" json:"runnerAttestation"`
	ExperimentSetID   *string        `gorm:"column:experimentSetId;index:idx_experiments_grouping,priority:1" json:"experimentSetId"`
	ComparisonID      *string        `gorm:"column:comparisonId;index:idx_experiments_grouping,priority:2" json:"comparisonId"`
	VariantID         *string        `gorm:"column:variantId;index:idx_experiments_grouping,priority:3" json:"variantId"`
	TrialIndex        *int           `gorm:"column:trialIndex;index:idx_experiments_grouping,priority:4;size:32" json:"trialIndex"`
	ScorerIds         datatypes.JSON `gorm:"column:scorerIds" json:"scorerIds"`
}

func (MastraExperiments) TableName() string {
	return "mastra_experiments"
}

func (m *MastraExperiments) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.StartedAtZ == nil {
		m.StartedAtZ = &DateTime{Time: now}
	}
	if m.CompletedAtZ == nil {
		m.CompletedAtZ = &DateTime{Time: now}
	}
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraFavorites struct {
	UserID     string    `gorm:"column:userId;primaryKey" json:"userId"`
	EntityType string    `gorm:"column:entityType;primaryKey;index:idx_favorites_entity,priority:1" json:"entityType"`
	EntityID   string    `gorm:"column:entityId;primaryKey;index:idx_favorites_entity,priority:2" json:"entityId"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ *DateTime `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraFavorites) TableName() string {
	return "mastra_favorites"
}

func (m *MastraFavorites) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraKnowledgeActivity struct {
	ID             string         `gorm:"column:id;primaryKey;index:idx_knowledge_activity_latest,sort:desc" json:"id"`
	Action         string         `gorm:"column:action;not null" json:"action"`
	RecordType     string         `gorm:"column:recordType;not null" json:"recordType"`
	RecordID       string         `gorm:"column:recordId;not null" json:"recordId"`
	Scope          datatypes.JSON `gorm:"column:scope;not null" json:"scope"`
	ScopeKey       string         `gorm:"column:scopeKey;not null" json:"scopeKey"`
	SourceThreadID *string        `gorm:"column:sourceThreadId" json:"sourceThreadId"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ     *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraKnowledgeActivity) TableName() string {
	return "mastra_knowledge_activity"
}

func (m *MastraKnowledgeActivity) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraKnowledgeCursors struct {
	SourceThreadID  string    `gorm:"column:sourceThreadId;primaryKey" json:"sourceThreadId"`
	Agent           string    `gorm:"column:agent;primaryKey" json:"agent"`
	LastKnowledgeID string    `gorm:"column:lastKnowledgeId;not null" json:"lastKnowledgeId"`
	UpdatedAt       DateTime  `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	UpdatedAtZ      *DateTime `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraKnowledgeCursors) TableName() string {
	return "mastra_knowledge_cursors"
}

func (m *MastraKnowledgeCursors) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraKnowledgeMentions struct {
	SourceType string `gorm:"column:sourceType;primaryKey;index:idx_knowledge_mentions_record,priority:2" json:"sourceType"`
	SourceID   string `gorm:"column:sourceId;primaryKey;index:idx_knowledge_mentions_record,priority:3" json:"sourceId"`
	RecordID   string `gorm:"column:recordId;primaryKey;index:idx_knowledge_mentions_record,priority:1" json:"recordId"`
}

func (MastraKnowledgeMentions) TableName() string {
	return "mastra_knowledge_mentions"
}

type MastraKnowledgeNodes struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	Type          string         `gorm:"column:type;uniqueIndex:idx_knowledge_nodes_identity,priority:1;index:idx_knowledge_nodes_scope,priority:2;not null" json:"type"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	CanonicalName string         `gorm:"column:canonicalName;uniqueIndex:idx_knowledge_nodes_identity,priority:3;not null" json:"canonicalName"`
	Kind          *string        `gorm:"column:kind" json:"kind"`
	Content       *string        `gorm:"column:content" json:"content"`
	Description   *string        `gorm:"column:description" json:"description"`
	Scope         datatypes.JSON `gorm:"column:scope;not null" json:"scope"`
	ScopeKey      string         `gorm:"column:scopeKey;uniqueIndex:idx_knowledge_nodes_identity,priority:2;index:idx_knowledge_nodes_scope,priority:1;not null" json:"scopeKey"`
	Version       int            `gorm:"column:version;size:32;not null" json:"version"`
	MergedInto    *string        `gorm:"column:mergedInto" json:"mergedInto"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ    *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ    *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraKnowledgeNodes) TableName() string {
	return "mastra_knowledge_nodes"
}

func (m *MastraKnowledgeNodes) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraKnowledgeRecords struct {
	ID             string         `gorm:"column:id;primaryKey;index:idx_knowledge_records_node_latest,priority:2,sort:desc;index:idx_knowledge_records_thread_latest,priority:2,sort:desc" json:"id"`
	Node           string         `gorm:"column:node;index:idx_knowledge_records_node_latest,priority:1;not null" json:"node"`
	Text           string         `gorm:"column:text;not null" json:"text"`
	Scope          datatypes.JSON `gorm:"column:scope;not null" json:"scope"`
	ScopeKey       string         `gorm:"column:scopeKey;not null" json:"scopeKey"`
	SourceThreadID string         `gorm:"column:sourceThreadId;index:idx_knowledge_records_thread_latest,priority:1;not null" json:"sourceThreadId"`
	CapturedAt     DateTime       `gorm:"column:capturedAt;type:timestamp(6);not null" json:"capturedAt"`
	When           *DateTime      `gorm:"column:when;type:timestamp(6)" json:"when"`
	MaxScope       *string        `gorm:"column:maxScope" json:"maxScope"`
	Metadata       datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	DeletedAt      *DateTime      `gorm:"column:deletedAt;type:timestamp(6)" json:"deletedAt"`
	DeletedBy      *string        `gorm:"column:deletedBy" json:"deletedBy"`
	CapturedAtZ    *DateTime      `gorm:"column:capturedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"capturedAtZ"`
	WhenZ          *DateTime      `gorm:"column:whenZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"whenZ"`
	DeletedAtZ     *DateTime      `gorm:"column:deletedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"deletedAtZ"`
}

func (MastraKnowledgeRecords) TableName() string {
	return "mastra_knowledge_records"
}

func (m *MastraKnowledgeRecords) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CapturedAtZ == nil {
		m.CapturedAtZ = &DateTime{Time: now}
	}
	if m.WhenZ == nil {
		m.WhenZ = &DateTime{Time: now}
	}
	if m.DeletedAtZ == nil {
		m.DeletedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraKnowledgeSemanticOutbox struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	IdempotencyKey string         `gorm:"column:idempotencyKey;uniqueIndex:idx_knowledge_outbox_idempotency;not null" json:"idempotencyKey"`
	DocumentID     string         `gorm:"column:documentId;not null" json:"documentId"`
	DocumentType   string         `gorm:"column:documentType;not null" json:"documentType"`
	Operation      string         `gorm:"column:operation;not null" json:"operation"`
	Scope          datatypes.JSON `gorm:"column:scope;not null" json:"scope"`
	ScopeKey       string         `gorm:"column:scopeKey;not null" json:"scopeKey"`
	Status         string         `gorm:"column:status;index:idx_knowledge_outbox_claim,priority:1;not null" json:"status"`
	Attempts       int            `gorm:"column:attempts;size:32;not null" json:"attempts"`
	AvailableAt    DateTime       `gorm:"column:availableAt;index:idx_knowledge_outbox_claim,priority:2;type:timestamp(6);not null" json:"availableAt"`
	ClaimedAt      *DateTime      `gorm:"column:claimedAt;type:timestamp(6)" json:"claimedAt"`
	ClaimedBy      *string        `gorm:"column:claimedBy" json:"claimedBy"`
	CreatedAt      DateTime       `gorm:"column:createdAt;index:idx_knowledge_outbox_claim,priority:3;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CompletedAt    *DateTime      `gorm:"column:completedAt;type:timestamp(6)" json:"completedAt"`
	AvailableAtZ   *DateTime      `gorm:"column:availableAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"availableAtZ"`
	ClaimedAtZ     *DateTime      `gorm:"column:claimedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"claimedAtZ"`
	CreatedAtZ     *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	CompletedAtZ   *DateTime      `gorm:"column:completedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"completedAtZ"`
}

func (MastraKnowledgeSemanticOutbox) TableName() string {
	return "mastra_knowledge_semantic_outbox"
}

func (m *MastraKnowledgeSemanticOutbox) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.AvailableAtZ == nil {
		m.AvailableAtZ = &DateTime{Time: now}
	}
	if m.ClaimedAtZ == nil {
		m.ClaimedAtZ = &DateTime{Time: now}
	}
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.CompletedAtZ == nil {
		m.CompletedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraMcpClientVersions struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	McpClientID   string         `gorm:"column:mcpClientId;uniqueIndex:idx_mcp_client_versions_client_version,priority:1;not null" json:"mcpClientId"`
	VersionNumber int            `gorm:"column:versionNumber;uniqueIndex:idx_mcp_client_versions_client_version,priority:2;size:32;not null" json:"versionNumber"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	Description   *string        `gorm:"column:description" json:"description"`
	Servers       datatypes.JSON `gorm:"column:servers;not null" json:"servers"`
	ChangedFields datatypes.JSON `gorm:"column:changedFields" json:"changedFields"`
	ChangeMessage *string        `gorm:"column:changeMessage" json:"changeMessage"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ    *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraMcpClientVersions) TableName() string {
	return "mastra_mcp_client_versions"
}

func (m *MastraMcpClientVersions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraMcpClients struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Status          string         `gorm:"column:status;not null" json:"status"`
	ActiveVersionID *string        `gorm:"column:activeVersionId" json:"activeVersionId"`
	AuthorID        *string        `gorm:"column:authorId" json:"authorId"`
	Metadata        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ      *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ      *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraMcpClients) TableName() string {
	return "mastra_mcp_clients"
}

func (m *MastraMcpClients) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraMcpServerVersions struct {
	ID               string         `gorm:"column:id;primaryKey" json:"id"`
	McpServerID      string         `gorm:"column:mcpServerId;uniqueIndex:idx_mcp_server_versions_server_version,priority:1;not null" json:"mcpServerId"`
	VersionNumber    int            `gorm:"column:versionNumber;uniqueIndex:idx_mcp_server_versions_server_version,priority:2;size:32;not null" json:"versionNumber"`
	Name             string         `gorm:"column:name;not null" json:"name"`
	Version          string         `gorm:"column:version;not null" json:"version"`
	Description      *string        `gorm:"column:description" json:"description"`
	Instructions     *string        `gorm:"column:instructions" json:"instructions"`
	Repository       datatypes.JSON `gorm:"column:repository" json:"repository"`
	ReleaseDate      *string        `gorm:"column:releaseDate" json:"releaseDate"`
	IsLatest         *bool          `gorm:"column:isLatest" json:"isLatest"`
	PackageCanonical *string        `gorm:"column:packageCanonical" json:"packageCanonical"`
	Tools            datatypes.JSON `gorm:"column:tools" json:"tools"`
	Agents           datatypes.JSON `gorm:"column:agents" json:"agents"`
	Workflows        datatypes.JSON `gorm:"column:workflows" json:"workflows"`
	ChangedFields    datatypes.JSON `gorm:"column:changedFields" json:"changedFields"`
	ChangeMessage    *string        `gorm:"column:changeMessage" json:"changeMessage"`
	CreatedAt        DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ       *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraMcpServerVersions) TableName() string {
	return "mastra_mcp_server_versions"
}

func (m *MastraMcpServerVersions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraMcpServers struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Status          string         `gorm:"column:status;not null" json:"status"`
	ActiveVersionID *string        `gorm:"column:activeVersionId" json:"activeVersionId"`
	AuthorID        *string        `gorm:"column:authorId" json:"authorId"`
	Metadata        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ      *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ      *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraMcpServers) TableName() string {
	return "mastra_mcp_servers"
}

func (m *MastraMcpServers) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraObservationalMemory struct {
	ID                              string         `gorm:"column:id;primaryKey" json:"id"`
	LookupKey                       string         `gorm:"column:lookupKey;index:idx_om_lookup_key;not null" json:"lookupKey"`
	Scope                           string         `gorm:"column:scope;not null" json:"scope"`
	ResourceID                      *string        `gorm:"column:resourceId" json:"resourceId"`
	ThreadID                        *string        `gorm:"column:threadId" json:"threadId"`
	ActiveObservations              string         `gorm:"column:activeObservations;not null" json:"activeObservations"`
	ActiveObservationsPendingUpdate *string        `gorm:"column:activeObservationsPendingUpdate" json:"activeObservationsPendingUpdate"`
	OriginType                      string         `gorm:"column:originType;not null" json:"originType"`
	Config                          string         `gorm:"column:config;not null" json:"config"`
	GenerationCount                 int            `gorm:"column:generationCount;size:32;not null" json:"generationCount"`
	LastObservedAt                  *DateTime      `gorm:"column:lastObservedAt;type:timestamp(6)" json:"lastObservedAt"`
	LastReflectionAt                *DateTime      `gorm:"column:lastReflectionAt;type:timestamp(6)" json:"lastReflectionAt"`
	PendingMessageTokens            int            `gorm:"column:pendingMessageTokens;size:32;not null" json:"pendingMessageTokens"`
	TotalTokensObserved             int            `gorm:"column:totalTokensObserved;size:32;not null" json:"totalTokensObserved"`
	ObservationTokenCount           int            `gorm:"column:observationTokenCount;size:32;not null" json:"observationTokenCount"`
	IsObserving                     bool           `gorm:"column:isObserving;not null" json:"isObserving"`
	IsReflecting                    bool           `gorm:"column:isReflecting;not null" json:"isReflecting"`
	ObservedMessageIds              datatypes.JSON `gorm:"column:observedMessageIds" json:"observedMessageIds"`
	ObservedTimezone                *string        `gorm:"column:observedTimezone" json:"observedTimezone"`
	BufferedObservations            *string        `gorm:"column:bufferedObservations" json:"bufferedObservations"`
	BufferedObservationTokens       *int           `gorm:"column:bufferedObservationTokens;size:32" json:"bufferedObservationTokens"`
	BufferedMessageIds              datatypes.JSON `gorm:"column:bufferedMessageIds" json:"bufferedMessageIds"`
	BufferedReflection              *string        `gorm:"column:bufferedReflection" json:"bufferedReflection"`
	BufferedReflectionTokens        *int           `gorm:"column:bufferedReflectionTokens;size:32" json:"bufferedReflectionTokens"`
	BufferedReflectionInputTokens   *int           `gorm:"column:bufferedReflectionInputTokens;size:32" json:"bufferedReflectionInputTokens"`
	ReflectedObservationLineCount   *int           `gorm:"column:reflectedObservationLineCount;size:32" json:"reflectedObservationLineCount"`
	BufferedObservationChunks       datatypes.JSON `gorm:"column:bufferedObservationChunks" json:"bufferedObservationChunks"`
	IsBufferingObservation          bool           `gorm:"column:isBufferingObservation;not null" json:"isBufferingObservation"`
	IsBufferingReflection           bool           `gorm:"column:isBufferingReflection;not null" json:"isBufferingReflection"`
	LastBufferedAtTokens            int            `gorm:"column:lastBufferedAtTokens;size:32;not null" json:"lastBufferedAtTokens"`
	LastBufferedAtTime              *DateTime      `gorm:"column:lastBufferedAtTime;type:timestamp(6)" json:"lastBufferedAtTime"`
	Metadata                        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt                       DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                       DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	LastObservedAtZ                 *DateTime      `gorm:"column:lastObservedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"lastObservedAtZ"`
	LastReflectionAtZ               *DateTime      `gorm:"column:lastReflectionAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"lastReflectionAtZ"`
	LastBufferedAtTimeZ             *DateTime      `gorm:"column:lastBufferedAtTimeZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"lastBufferedAtTimeZ"`
	CreatedAtZ                      *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ                      *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraObservationalMemory) TableName() string {
	return "mastra_observational_memory"
}

func (m *MastraObservationalMemory) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.LastObservedAtZ == nil {
		m.LastObservedAtZ = &DateTime{Time: now}
	}
	if m.LastReflectionAtZ == nil {
		m.LastReflectionAtZ = &DateTime{Time: now}
	}
	if m.LastBufferedAtTimeZ == nil {
		m.LastBufferedAtTimeZ = &DateTime{Time: now}
	}
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraPromptBlockVersions struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	BlockID              string         `gorm:"column:blockId;uniqueIndex:idx_prompt_block_versions_block_version,priority:1;not null" json:"blockId"`
	VersionNumber        int            `gorm:"column:versionNumber;uniqueIndex:idx_prompt_block_versions_block_version,priority:2;size:32;not null" json:"versionNumber"`
	Name                 string         `gorm:"column:name;not null" json:"name"`
	Description          *string        `gorm:"column:description" json:"description"`
	Content              string         `gorm:"column:content;not null" json:"content"`
	Rules                datatypes.JSON `gorm:"column:rules" json:"rules"`
	RequestContextSchema datatypes.JSON `gorm:"column:requestContextSchema" json:"requestContextSchema"`
	ChangedFields        datatypes.JSON `gorm:"column:changedFields" json:"changedFields"`
	ChangeMessage        *string        `gorm:"column:changeMessage" json:"changeMessage"`
	CreatedAt            DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ           *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraPromptBlockVersions) TableName() string {
	return "mastra_prompt_block_versions"
}

func (m *MastraPromptBlockVersions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraPromptBlocks struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Status          string         `gorm:"column:status;not null" json:"status"`
	ActiveVersionID *string        `gorm:"column:activeVersionId" json:"activeVersionId"`
	AuthorID        *string        `gorm:"column:authorId" json:"authorId"`
	Metadata        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ      *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ      *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraPromptBlocks) TableName() string {
	return "mastra_prompt_blocks"
}

func (m *MastraPromptBlocks) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraScheduleTriggers struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	ScheduleID      string         `gorm:"column:schedule_id;index:idx_mastra_schedule_triggers_schedule_fire,priority:1;not null" json:"schedule_id"`
	RunID           *string        `gorm:"column:run_id" json:"run_id"`
	ScheduledFireAt int64          `gorm:"column:scheduled_fire_at;not null" json:"scheduled_fire_at"`
	ActualFireAt    int64          `gorm:"column:actual_fire_at;index:idx_mastra_schedule_triggers_schedule_fire,priority:2,sort:desc;not null" json:"actual_fire_at"`
	Outcome         string         `gorm:"column:outcome;not null" json:"outcome"`
	Error           *string        `gorm:"column:error" json:"error"`
	TriggerKind     string         `gorm:"column:trigger_kind;not null" json:"trigger_kind"`
	ParentTriggerID *string        `gorm:"column:parent_trigger_id" json:"parent_trigger_id"`
	Metadata        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
}

func (MastraScheduleTriggers) TableName() string {
	return "mastra_schedule_triggers"
}

type MastraSchedules struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	Target     datatypes.JSON `gorm:"column:target;not null" json:"target"`
	Cron       string         `gorm:"column:cron;not null" json:"cron"`
	Timezone   *string        `gorm:"column:timezone" json:"timezone"`
	Status     string         `gorm:"column:status;index:idx_mastra_schedules_status_next_fire,priority:1;not null" json:"status"`
	NextFireAt int64          `gorm:"column:next_fire_at;index:idx_mastra_schedules_status_next_fire,priority:2;not null" json:"next_fire_at"`
	LastFireAt *int64         `gorm:"column:last_fire_at" json:"last_fire_at"`
	LastRunID  *string        `gorm:"column:last_run_id" json:"last_run_id"`
	CreatedAt  int64          `gorm:"column:created_at;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt  int64          `gorm:"column:updated_at;autoUpdateTime:false;not null" json:"updated_at"`
	Metadata   datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	OwnerType  *string        `gorm:"column:owner_type" json:"owner_type"`
	OwnerID    *string        `gorm:"column:owner_id" json:"owner_id"`
}

func (MastraSchedules) TableName() string {
	return "mastra_schedules"
}

type MastraScorerDefinitionVersions struct {
	ID                 string         `gorm:"column:id;primaryKey" json:"id"`
	ScorerDefinitionID string         `gorm:"column:scorerDefinitionId;uniqueIndex:idx_scorer_definition_versions_def_version,priority:1;not null" json:"scorerDefinitionId"`
	VersionNumber      int            `gorm:"column:versionNumber;uniqueIndex:idx_scorer_definition_versions_def_version,priority:2;size:32;not null" json:"versionNumber"`
	Name               string         `gorm:"column:name;not null" json:"name"`
	Description        *string        `gorm:"column:description" json:"description"`
	Type               string         `gorm:"column:type;not null" json:"type"`
	Model              datatypes.JSON `gorm:"column:model" json:"model"`
	Instructions       *string        `gorm:"column:instructions" json:"instructions"`
	ScoreRange         datatypes.JSON `gorm:"column:scoreRange" json:"scoreRange"`
	PresetConfig       datatypes.JSON `gorm:"column:presetConfig" json:"presetConfig"`
	DefaultSampling    datatypes.JSON `gorm:"column:defaultSampling" json:"defaultSampling"`
	ChangedFields      datatypes.JSON `gorm:"column:changedFields" json:"changedFields"`
	ChangeMessage      *string        `gorm:"column:changeMessage" json:"changeMessage"`
	CreatedAt          DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ         *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraScorerDefinitionVersions) TableName() string {
	return "mastra_scorer_definition_versions"
}

func (m *MastraScorerDefinitionVersions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraScorerDefinitions struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Status          string         `gorm:"column:status;not null" json:"status"`
	ActiveVersionID *string        `gorm:"column:activeVersionId" json:"activeVersionId"`
	AuthorID        *string        `gorm:"column:authorId" json:"authorId"`
	Metadata        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ      *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ      *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
	OrganizationID  *string        `gorm:"column:organizationId" json:"organizationId"`
	ProjectID       *string        `gorm:"column:projectId" json:"projectId"`
}

func (MastraScorerDefinitions) TableName() string {
	return "mastra_scorer_definitions"
}

func (m *MastraScorerDefinitions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraSkillBlobs struct {
	Hash       string    `gorm:"column:hash;primaryKey" json:"hash"`
	Content    string    `gorm:"column:content;not null" json:"content"`
	Size       int       `gorm:"column:size;size:32;not null" json:"size"`
	MimeType   *string   `gorm:"column:mimeType" json:"mimeType"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ *DateTime `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraSkillBlobs) TableName() string {
	return "mastra_skill_blobs"
}

func (m *MastraSkillBlobs) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraSkillVersions struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	SkillID       string         `gorm:"column:skillId;uniqueIndex:idx_skill_versions_skill_version,priority:1;not null" json:"skillId"`
	VersionNumber int            `gorm:"column:versionNumber;uniqueIndex:idx_skill_versions_skill_version,priority:2;size:32;not null" json:"versionNumber"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	Description   string         `gorm:"column:description;not null" json:"description"`
	Instructions  string         `gorm:"column:instructions;not null" json:"instructions"`
	License       *string        `gorm:"column:license" json:"license"`
	Compatibility datatypes.JSON `gorm:"column:compatibility" json:"compatibility"`
	Source        datatypes.JSON `gorm:"column:source" json:"source"`
	References    datatypes.JSON `gorm:"column:references" json:"references"`
	Scripts       datatypes.JSON `gorm:"column:scripts" json:"scripts"`
	Assets        datatypes.JSON `gorm:"column:assets" json:"assets"`
	Metadata      datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	Tree          datatypes.JSON `gorm:"column:tree" json:"tree"`
	ChangedFields datatypes.JSON `gorm:"column:changedFields" json:"changedFields"`
	ChangeMessage *string        `gorm:"column:changeMessage" json:"changeMessage"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ    *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	Files         datatypes.JSON `gorm:"column:files" json:"files"`
}

func (MastraSkillVersions) TableName() string {
	return "mastra_skill_versions"
}

func (m *MastraSkillVersions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraSkills struct {
	ID              string    `gorm:"column:id;primaryKey" json:"id"`
	Status          string    `gorm:"column:status;not null" json:"status"`
	ActiveVersionID *string   `gorm:"column:activeVersionId" json:"activeVersionId"`
	AuthorID        *string   `gorm:"column:authorId" json:"authorId"`
	CreatedAt       DateTime  `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime  `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ      *DateTime `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ      *DateTime `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
	Visibility      *string   `gorm:"column:visibility" json:"visibility"`
	FavoriteCount   *int      `gorm:"column:favoriteCount;size:32" json:"favoriteCount"`
}

func (MastraSkills) TableName() string {
	return "mastra_skills"
}

func (m *MastraSkills) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraThreadState struct {
	ThreadID   string         `gorm:"column:threadId;primaryKey" json:"threadId"`
	Type       string         `gorm:"column:type;primaryKey" json:"type"`
	Value      datatypes.JSON `gorm:"column:value;not null" json:"value"`
	CreatedAt  DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraThreadState) TableName() string {
	return "mastra_thread_state"
}

func (m *MastraThreadState) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraToolProviderConnections struct {
	AuthorID     string    `gorm:"column:authorId;primaryKey;index:idx_tool_provider_connections_author,priority:1" json:"authorId"`
	ProviderID   string    `gorm:"column:providerId;primaryKey;index:idx_tool_provider_connections_author,priority:2" json:"providerId"`
	ConnectionID string    `gorm:"column:connectionId;primaryKey" json:"connectionId"`
	Toolkit      string    `gorm:"column:toolkit;index:idx_tool_provider_connections_author,priority:3;not null" json:"toolkit"`
	Label        *string   `gorm:"column:label" json:"label"`
	Scope        string    `gorm:"column:scope;not null" json:"scope"`
	CreatedAt    DateTime  `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime  `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ   *DateTime `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ   *DateTime `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraToolProviderConnections) TableName() string {
	return "mastra_tool_provider_connections"
}

func (m *MastraToolProviderConnections) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraWorkflowDefinitions struct {
	ID                   string         `gorm:"column:id;primaryKey" json:"id"`
	Description          *string        `gorm:"column:description" json:"description"`
	Metadata             datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	InputSchema          datatypes.JSON `gorm:"column:inputSchema;not null" json:"inputSchema"`
	OutputSchema         datatypes.JSON `gorm:"column:outputSchema;not null" json:"outputSchema"`
	StateSchema          datatypes.JSON `gorm:"column:stateSchema" json:"stateSchema"`
	RequestContextSchema datatypes.JSON `gorm:"column:requestContextSchema" json:"requestContextSchema"`
	Graph                datatypes.JSON `gorm:"column:graph;not null" json:"graph"`
	Schedule             datatypes.JSON `gorm:"column:schedule" json:"schedule"`
	Status               string         `gorm:"column:status;index:idx_workflow_definitions_status;not null" json:"status"`
	Source               string         `gorm:"column:source;not null" json:"source"`
	AuthorID             *string        `gorm:"column:authorId" json:"authorId"`
	CreatedAt            DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ           *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ           *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraWorkflowDefinitions) TableName() string {
	return "mastra_workflow_definitions"
}

func (m *MastraWorkflowDefinitions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraWorkspaceVersions struct {
	ID               string         `gorm:"column:id;primaryKey" json:"id"`
	WorkspaceID      string         `gorm:"column:workspaceId;uniqueIndex:idx_workspace_versions_workspace_version,priority:1;not null" json:"workspaceId"`
	VersionNumber    int            `gorm:"column:versionNumber;uniqueIndex:idx_workspace_versions_workspace_version,priority:2;size:32;not null" json:"versionNumber"`
	Name             string         `gorm:"column:name;not null" json:"name"`
	Description      *string        `gorm:"column:description" json:"description"`
	Filesystem       datatypes.JSON `gorm:"column:filesystem" json:"filesystem"`
	Sandbox          datatypes.JSON `gorm:"column:sandbox" json:"sandbox"`
	Mounts           datatypes.JSON `gorm:"column:mounts" json:"mounts"`
	Search           datatypes.JSON `gorm:"column:search" json:"search"`
	Skills           datatypes.JSON `gorm:"column:skills" json:"skills"`
	Tools            datatypes.JSON `gorm:"column:tools" json:"tools"`
	AutoSync         *bool          `gorm:"column:autoSync" json:"autoSync"`
	OperationTimeout *int           `gorm:"column:operationTimeout;size:32" json:"operationTimeout"`
	ChangedFields    datatypes.JSON `gorm:"column:changedFields" json:"changedFields"`
	ChangeMessage    *string        `gorm:"column:changeMessage" json:"changeMessage"`
	CreatedAt        DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	CreatedAtZ       *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
}

func (MastraWorkspaceVersions) TableName() string {
	return "mastra_workspace_versions"
}

func (m *MastraWorkspaceVersions) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type MastraWorkspaces struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Status          string         `gorm:"column:status;not null" json:"status"`
	ActiveVersionID *string        `gorm:"column:activeVersionId" json:"activeVersionId"`
	AuthorID        *string        `gorm:"column:authorId" json:"authorId"`
	Metadata        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(6);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(6);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedAtZ      *DateTime      `gorm:"column:createdAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"createdAtZ"`
	UpdatedAtZ      *DateTime      `gorm:"column:updatedAtZ;type:timestamptz(6);default:CURRENT_TIMESTAMP" json:"updatedAtZ"`
}

func (MastraWorkspaces) TableName() string {
	return "mastra_workspaces"
}

func (m *MastraWorkspaces) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAtZ == nil {
		m.CreatedAtZ = &DateTime{Time: now}
	}
	if m.UpdatedAtZ == nil {
		m.UpdatedAtZ = &DateTime{Time: now}
	}
	return nil
}

type OAuthApp struct {
	ID                       string                         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID           *string                        `gorm:"column:organizationId;uniqueIndex:OAuthApp_organizationId_deletedAt_key,priority:1;index:OAuthApp_organizationId_idx" json:"organizationId"`
	Name                     string                         `gorm:"column:name;not null" json:"name"`
	Description              *string                        `gorm:"column:description" json:"description"`
	PictureID                *string                        `gorm:"column:pictureId" json:"pictureId"`
	RedirectURL              string                         `gorm:"column:redirectUrl;not null" json:"redirectUrl"`
	RedirectUris             *string                        `gorm:"column:redirectUris" json:"redirectUris"`
	ClientID                 string                         `gorm:"column:clientId;uniqueIndex:OAuthApp_clientId_key;index:OAuthApp_clientId_idx;not null" json:"clientId"`
	ClientSecret             *string                        `gorm:"column:clientSecret" json:"clientSecret"`
	Dynamic                  bool                           `gorm:"column:dynamic;index:OAuthApp_dynamic_idx;default:false;not null" json:"dynamic"`
	TokenEndpointAuthMethod  *string                        `gorm:"column:tokenEndpointAuthMethod" json:"tokenEndpointAuthMethod"`
	CreatedAt                DateTime                       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime                       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt                *DateTime                      `gorm:"column:deletedAt;uniqueIndex:OAuthApp_organizationId_deletedAt_key,priority:2;index:OAuthApp_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	Organization             *Organization                  `gorm:"belongsTo"`
	Picture                  *Media                         `gorm:"foreignKey:PictureID;belongsTo"`
	Authorizations           []OAuthAuthorization           `gorm:"foreignKey:OauthAppID;constraint:OAuthAuthorization_oauthAppId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SelfHostedAuthorizations []OAuthSelfHostedAuthorization `gorm:"foreignKey:OauthAppID;constraint:OAuthSelfHostedAuthorization_oauthAppId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (OAuthApp) TableName() string {
	return "OAuthApp"
}

func (m *OAuthApp) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OAuthApp) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OAuthAuthorization struct {
	ID                  string        `gorm:"column:id;primaryKey" json:"id"`
	OauthAppID          string        `gorm:"column:oauthAppId;uniqueIndex:OAuthAuthorization_oauthAppId_userId_organizationId_key,priority:1;index:OAuthAuthorization_oauthAppId_idx;not null" json:"oauthAppId"`
	UserID              string        `gorm:"column:userId;uniqueIndex:OAuthAuthorization_oauthAppId_userId_organizationId_key,priority:2;index:OAuthAuthorization_userId_idx;not null" json:"userId"`
	OrganizationID      string        `gorm:"column:organizationId;uniqueIndex:OAuthAuthorization_oauthAppId_userId_organizationId_key,priority:3;index:OAuthAuthorization_organizationId_idx;not null" json:"organizationId"`
	AccessToken         *string       `gorm:"column:accessToken;index:OAuthAuthorization_accessToken_idx" json:"accessToken"`
	AuthorizationCode   *string       `gorm:"column:authorizationCode;index:OAuthAuthorization_authorizationCode_idx" json:"authorizationCode"`
	CodeExpiresAt       *DateTime     `gorm:"column:codeExpiresAt;type:timestamp(3)" json:"codeExpiresAt"`
	CodeChallenge       *string       `gorm:"column:codeChallenge" json:"codeChallenge"`
	CodeChallengeMethod *string       `gorm:"column:codeChallengeMethod" json:"codeChallengeMethod"`
	RedirectURI         *string       `gorm:"column:redirectUri" json:"redirectUri"`
	RevokedAt           *DateTime     `gorm:"column:revokedAt;index:OAuthAuthorization_revokedAt_idx;type:timestamp(3)" json:"revokedAt"`
	CreatedAt           DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OauthApp            *OAuthApp     `gorm:"foreignKey:OauthAppID;belongsTo"`
	User                *User         `gorm:"belongsTo"`
	Organization        *Organization `gorm:"belongsTo"`
}

func (OAuthAuthorization) TableName() string {
	return "OAuthAuthorization"
}

func (m *OAuthAuthorization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OAuthAuthorization) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OAuthSelfHostedAuthorization struct {
	ID                  string    `gorm:"column:id;primaryKey" json:"id"`
	OauthAppID          string    `gorm:"column:oauthAppId;index:OAuthSelfHostedAuthorization_oauthAppId_idx;not null" json:"oauthAppId"`
	McpURL              string    `gorm:"column:mcpUrl;not null" json:"mcpUrl"`
	APIKey              string    `gorm:"column:apiKey;not null" json:"apiKey"`
	Email               *string   `gorm:"column:email" json:"email"`
	AccessToken         *string   `gorm:"column:accessToken;index:OAuthSelfHostedAuthorization_accessToken_idx" json:"accessToken"`
	AuthorizationCode   *string   `gorm:"column:authorizationCode;index:OAuthSelfHostedAuthorization_authorizationCode_idx" json:"authorizationCode"`
	CodeExpiresAt       *DateTime `gorm:"column:codeExpiresAt;type:timestamp(3)" json:"codeExpiresAt"`
	CodeChallenge       *string   `gorm:"column:codeChallenge" json:"codeChallenge"`
	CodeChallengeMethod *string   `gorm:"column:codeChallengeMethod" json:"codeChallengeMethod"`
	RedirectURI         *string   `gorm:"column:redirectUri" json:"redirectUri"`
	RevokedAt           *DateTime `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	CreatedAt           DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OauthApp            *OAuthApp `gorm:"foreignKey:OauthAppID;belongsTo"`
}

func (OAuthSelfHostedAuthorization) TableName() string {
	return "OAuthSelfHostedAuthorization"
}

func (m *OAuthSelfHostedAuthorization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OAuthSelfHostedAuthorization) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Announcement struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	Title       string   `gorm:"column:title;not null" json:"title"`
	Description string   `gorm:"column:description;not null" json:"description"`
	Color       string   `gorm:"column:color;default:'INFO';not null" json:"color"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (Announcement) TableName() string {
	return "Announcement"
}

func (m *Announcement) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
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
	return dateTime.UTC().Truncate(time.Millisecond), nil
}

func (dateTime *DateTime) Scan(src any) error {
	read, err := readPrismaTime(src)
	dateTime.Time = read
	return err
}

type Date struct{ time.Time }

func (Date) GormDataType() string {
	return "time"
}

func (date Date) Value() (driver.Value, error) {
	return date.UTC().Format("2006-01-02"), nil
}

func (date *Date) Scan(src any) error {
	read, err := readPrismaTime(src)
	if err != nil || read.IsZero() {
		date.Time = time.Time{}
		return err
	}
	date.Time = time.Date(read.Year(), read.Month(), read.Day(), 0, 0, 0, 0, time.UTC)
	return nil
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
