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
	"gorm.io/gorm/schema"
)

var NamingStrategy = schema.NamingStrategy{SingularTable: true, NoLowerCase: true}

type Image struct {
	ID           string        `gorm:"column:id;primaryKey" json:"id"`
	ContentType  string        `gorm:"column:contentType;not null" json:"contentType"`
	AltText      *string       `gorm:"column:altText" json:"altText"`
	Blob         []byte        `gorm:"column:blob;not null" json:"blob"`
	CreatedAt    DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OwnerOrgID   string        `gorm:"column:ownerOrgId;index:Image_ownerOrgId_idx;not null" json:"ownerOrgId"`
	UserID       string        `gorm:"column:userId;index:Image_userId_idx;not null" json:"userId"`
	OwnerOrg     *Organization `gorm:"foreignKey:OwnerOrgID;belongsTo"`
	User         *User         `gorm:"belongsTo"`
	Location     *Location     `gorm:"foreignKey:ImageID;constraint:Location_imageId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Organization *Organization `gorm:"foreignKey:ImageID;constraint:Organization_imageId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Image) TableName() string {
	return "Image"
}

func (m *Image) BeforeCreate(tx *gorm.DB) error {
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

func (*Image) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type User struct {
	ID                         string                   `gorm:"column:id;primaryKey" json:"id"`
	Email                      string                   `gorm:"column:email;uniqueIndex:User_email_key;uniqueIndex:User_email_username_key;not null" json:"email"`
	Username                   string                   `gorm:"column:username;uniqueIndex:User_username_key;uniqueIndex:User_email_username_key;not null" json:"username"`
	FirstName                  *string                  `gorm:"column:firstName;index:User_firstName_lastName_idx" json:"firstName"`
	LastName                   *string                  `gorm:"column:lastName;index:User_firstName_lastName_idx" json:"lastName"`
	DisplayName                *string                  `gorm:"column:displayName" json:"displayName"`
	ProfilePicture             *string                  `gorm:"column:profilePicture" json:"profilePicture"`
	UsedFreeTrial              bool                     `gorm:"column:usedFreeTrial;default:false;not null" json:"usedFreeTrial"`
	Onboarded                  bool                     `gorm:"column:onboarded;default:false;not null" json:"onboarded"`
	CustomerID                 *string                  `gorm:"column:customerId;uniqueIndex:User_customerId_key" json:"customerId"`
	Sso                        bool                     `gorm:"column:sso;default:false;not null" json:"sso"`
	CreatedWithInvite          bool                     `gorm:"column:createdWithInvite;default:false;not null" json:"createdWithInvite"`
	SkipSubscriptionCheck      bool                     `gorm:"column:skipSubscriptionCheck;default:false;not null" json:"skipSubscriptionCheck"`
	HasUnpaidInvoice           bool                     `gorm:"column:hasUnpaidInvoice;default:false;not null" json:"hasUnpaidInvoice"`
	WarnForNoPaymentMethod     bool                     `gorm:"column:warnForNoPaymentMethod;default:false;not null" json:"warnForNoPaymentMethod"`
	LastSelectedOrganizationID *string                  `gorm:"column:lastSelectedOrganizationId;index:User_lastSelectedOrganizationId_idx" json:"lastSelectedOrganizationId"`
	CreatedAt                  DateTime                 `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime                 `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt                  *DateTime                `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	LastMobileActiveAt         *DateTime                `gorm:"column:lastMobileActiveAt;index:User_lastMobileActiveAt_idx;type:timestamp(3)" json:"lastMobileActiveAt"`
	DateFormat                 *string                  `gorm:"column:dateFormat" json:"dateFormat"`
	TimeFormat                 *string                  `gorm:"column:timeFormat" json:"timeFormat"`
	WeekStart                  *string                  `gorm:"column:weekStart" json:"weekStart"`
	TimeZone                   *string                  `gorm:"column:timeZone" json:"timeZone"`
	TierID                     string                   `gorm:"column:tierId;index:User_tierId_idx;default:'free';not null" json:"tierId"`
	ReferralSource             *string                  `gorm:"column:referralSource" json:"referralSource"`
	LastSelectedOrganization   *Organization            `gorm:"foreignKey:LastSelectedOrganizationID;belongsTo"`
	Tier                       *Tier                    `gorm:"belongsTo"`
	Assets                     []Asset                  `gorm:"foreignKey:UserID;constraint:Asset_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Categories                 []Category               `gorm:"foreignKey:UserID;constraint:Category_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Notes                      []Note                   `gorm:"foreignKey:UserID;constraint:Note_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	BookingNotes               []BookingNote            `gorm:"foreignKey:UserID;constraint:BookingNote_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AuditNotes                 []AuditNote              `gorm:"foreignKey:UserID;constraint:AuditNote_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	LocationNotes              []LocationNote           `gorm:"foreignKey:UserID;constraint:LocationNote_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	QrCodes                    []Qr                     `gorm:"foreignKey:UserID;constraint:Qr_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Scans                      []Scan                   `gorm:"foreignKey:UserID;constraint:Scan_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tags                       []Tag                    `gorm:"foreignKey:UserID;constraint:Tag_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Locations                  []Location               `gorm:"foreignKey:UserID;constraint:Location_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Images                     []Image                  `gorm:"foreignKey:UserID;constraint:Image_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Organizations              []Organization           `gorm:"foreignKey:UserID;constraint:Organization_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomFields               []CustomField            `gorm:"foreignKey:UserID;constraint:CustomField_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SentInvites                []Invite                 `gorm:"foreignKey:InviterID;constraint:Invite_inviterId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReceivedInvites            []Invite                 `gorm:"foreignKey:InviteeUserID;constraint:Invite_inviteeUserId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TeamMembers                []TeamMember             `gorm:"foreignKey:UserID;constraint:TeamMember_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserOrganizations          []UserOrganization       `gorm:"foreignKey:UserID;constraint:UserOrganization_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Bookings                   []Booking                `gorm:"foreignKey:CreatorID;constraint:Booking_creatorId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Custodies                  []Booking                `gorm:"foreignKey:CustodianUserID;constraint:Booking_custodianUserId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedKits                []Kit                    `gorm:"foreignKey:CreatedByID;constraint:Kit_createdById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AssetReminders             []AssetReminder          `gorm:"foreignKey:CreatedByID;constraint:AssetReminder_createdById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AuditsCreated              []AuditSession           `gorm:"foreignKey:CreatedByID;constraint:AuditSession_createdById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AssetImports               []AssetImport            `gorm:"foreignKey:CreatedByID;constraint:AssetImport_createdById_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuditAssignments           []AuditAssignment        `gorm:"foreignKey:UserID;constraint:AuditAssignment_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuditAssetsScanned         []AuditAsset             `gorm:"foreignKey:ScannedByID;constraint:AuditAsset_scannedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AuditScans                 []AuditScan              `gorm:"foreignKey:ScannedByID;constraint:AuditScan_scannedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AuditImagesUploaded        []AuditImage             `gorm:"foreignKey:UploadedByID;constraint:AuditImage_uploadedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AssetIndexSettings         []AssetIndexSettings     `gorm:"foreignKey:UserID;constraint:AssetIndexSettings_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PartialCheckins            []PartialBookingCheckin  `gorm:"foreignKey:CheckedInByID;constraint:PartialBookingCheckin_checkedInById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PartialCheckouts           []PartialBookingCheckout `gorm:"foreignKey:CheckedOutByID;constraint:PartialBookingCheckout_checkedOutById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	MobileAuthCodes            []MobileAuthCode         `gorm:"foreignKey:UserID;constraint:MobileAuthCode_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PushTokens                 []PushToken              `gorm:"foreignKey:UserID;constraint:PushToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedUpdates             []Update                 `gorm:"foreignKey:CreatedByID;constraint:Update_createdById_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UpdateReads                []UserUpdateRead         `gorm:"foreignKey:UserID;constraint:UserUpdateRead_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetFilterPresets         []AssetFilterPreset      `gorm:"foreignKey:OwnerID;constraint:AssetFilterPreset_ownerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetModels                []AssetModel             `gorm:"foreignKey:UserID;constraint:AssetModel_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConsumptionLogs            []ConsumptionLog         `gorm:"foreignKey:UserID;constraint:ConsumptionLog_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ScimExternalIds            []UserScimExternalID     `gorm:"foreignKey:UserID;constraint:UserScimExternalId_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleChangesReceived        []RoleChangeLog          `gorm:"foreignKey:UserID;constraint:RoleChangeLog_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	RoleChangesPerformed       []RoleChangeLog          `gorm:"foreignKey:ChangedByID;constraint:RoleChangeLog_changedById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	TeamMemberNotesAuthored    []TeamMemberNote         `gorm:"foreignKey:UserID;constraint:TeamMemberNote_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Contact                    *UserContact             `gorm:"foreignKey:UserID;constraint:UserContact_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BusinessIntel              *UserBusinessIntel       `gorm:"foreignKey:UserID;constraint:UserBusinessIntel_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	NotificationPreference     *NotificationPreference  `gorm:"foreignKey:UserID;constraint:NotificationPreference_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomTierLimit            *CustomTierLimit         `gorm:"foreignKey:UserID;constraint:CustomTierLimit_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Roles                      []Role                   `gorm:"many2many:_RoleToUser;joinForeignKey:B;joinReferences:A"`
}

func (User) TableName() string {
	return "User"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.Username == "" {
		m.Username = cuid.New()
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

type UserContact struct {
	ID            string   `gorm:"column:id;primaryKey" json:"id"`
	Phone         *string  `gorm:"column:phone;index:UserContact_phone_idx" json:"phone"`
	Street        *string  `gorm:"column:street" json:"street"`
	City          *string  `gorm:"column:city;index:UserContact_city_stateProvince_idx;index:UserContact_city_countryRegion_idx" json:"city"`
	StateProvince *string  `gorm:"column:stateProvince;index:UserContact_city_stateProvince_idx" json:"stateProvince"`
	ZipPostalCode *string  `gorm:"column:zipPostalCode;index:UserContact_zipPostalCode_idx" json:"zipPostalCode"`
	CountryRegion *string  `gorm:"column:countryRegion;index:UserContact_countryRegion_idx;index:UserContact_city_countryRegion_idx" json:"countryRegion"`
	UserID        string   `gorm:"column:userId;uniqueIndex:UserContact_userId_key;index:UserContact_userId_idx;not null" json:"userId"`
	CreatedAt     DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User          *User    `gorm:"belongsTo"`
}

func (UserContact) TableName() string {
	return "UserContact"
}

func (m *UserContact) BeforeCreate(tx *gorm.DB) error {
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

func (*UserContact) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type UserBusinessIntel struct {
	ID                   string   `gorm:"column:id;primaryKey" json:"id"`
	HowDidYouHearAboutUs *string  `gorm:"column:howDidYouHearAboutUs" json:"howDidYouHearAboutUs"`
	JobTitle             *string  `gorm:"column:jobTitle;index:UserBusinessIntel_jobTitle_idx" json:"jobTitle"`
	TeamSize             *string  `gorm:"column:teamSize;index:UserBusinessIntel_teamSize_idx" json:"teamSize"`
	CompanyName          *string  `gorm:"column:companyName;index:UserBusinessIntel_companyName_idx" json:"companyName"`
	PrimaryUseCase       *string  `gorm:"column:primaryUseCase" json:"primaryUseCase"`
	CurrentSolution      *string  `gorm:"column:currentSolution" json:"currentSolution"`
	Timeline             *string  `gorm:"column:timeline" json:"timeline"`
	SignupPlan           *string  `gorm:"column:signupPlan" json:"signupPlan"`
	SignupTrial          *bool    `gorm:"column:signupTrial" json:"signupTrial"`
	UtmSource            *string  `gorm:"column:utmSource;index:UserBusinessIntel_utmSource_idx" json:"utmSource"`
	UtmMedium            *string  `gorm:"column:utmMedium" json:"utmMedium"`
	UtmCampaign          *string  `gorm:"column:utmCampaign;index:UserBusinessIntel_utmCampaign_idx" json:"utmCampaign"`
	UtmContent           *string  `gorm:"column:utmContent" json:"utmContent"`
	UserID               string   `gorm:"column:userId;uniqueIndex:UserBusinessIntel_userId_key;index:UserBusinessIntel_userId_idx;not null" json:"userId"`
	CreatedAt            DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User                 *User    `gorm:"belongsTo"`
}

func (UserBusinessIntel) TableName() string {
	return "UserBusinessIntel"
}

func (m *UserBusinessIntel) BeforeCreate(tx *gorm.DB) error {
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

func (*UserBusinessIntel) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Asset struct {
	ID                   string                  `gorm:"column:id;primaryKey;index:Asset_organizationId_createdAt_id_idx" json:"id"`
	Title                string                  `gorm:"column:title;index:Asset_title_description_idx;index:Asset_organizationId_compound_idx;not null" json:"title"`
	Description          *string                 `gorm:"column:description;index:Asset_title_description_idx" json:"description"`
	MainImage            *string                 `gorm:"column:mainImage" json:"mainImage"`
	ThumbnailImage       *string                 `gorm:"column:thumbnailImage" json:"thumbnailImage"`
	MainImageExpiration  *DateTime               `gorm:"column:mainImageExpiration;type:timestamp(3)" json:"mainImageExpiration"`
	Status               string                  `gorm:"column:status;index:Asset_organizationId_compound_idx;index:Asset_status_organizationId_idx;default:'AVAILABLE';not null" json:"status"`
	Valuation            *float64                `gorm:"column:value;index:Asset_valuation_organizationId_idx;type:double precision" json:"value"`
	AvailableToBook      *bool                   `gorm:"column:availableToBook;index:Asset_organizationId_compound_idx;default:true;not null" json:"availableToBook"`
	SequentialID         *string                 `gorm:"column:sequentialId;uniqueIndex:Asset_organizationId_sequentialId_key;index:Asset_sequentialId_trgm_idx;index:Asset_sequentialId_idx" json:"sequentialId"`
	ArchivedAt           *DateTime               `gorm:"column:archivedAt;index:Asset_organizationId_archivedAt_idx;type:timestamptz(3)" json:"archivedAt"`
	Type                 string                  `gorm:"column:type;default:'INDIVIDUAL';not null" json:"type"`
	AssetModelID         *string                 `gorm:"column:assetModelId;index:Asset_assetModelId_organizationId_idx" json:"assetModelId"`
	Quantity             *int                    `gorm:"column:quantity;size:32" json:"quantity"`
	MinQuantity          *int                    `gorm:"column:minQuantity;size:32" json:"minQuantity"`
	ConsumptionType      *string                 `gorm:"column:consumptionType" json:"consumptionType"`
	UnitOfMeasure        *string                 `gorm:"column:unitOfMeasure" json:"unitOfMeasure"`
	LowStockNotifiedAt   *DateTime               `gorm:"column:lowStockNotifiedAt;type:timestamp(3)" json:"lowStockNotifiedAt"`
	StockLedgerStartedAt *DateTime               `gorm:"column:stockLedgerStartedAt;type:timestamp(3)" json:"stockLedgerStartedAt"`
	CreatedAt            DateTime                `gorm:"column:createdAt;index:Asset_createdAt_organizationId_idx;index:Asset_organizationId_createdAt_id_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime                `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID               string                  `gorm:"column:userId;index:Asset_userId_idx;not null" json:"userId"`
	OrganizationID       string                  `gorm:"column:organizationId;uniqueIndex:Asset_organizationId_sequentialId_key;index:Asset_organizationId_compound_idx;index:Asset_status_organizationId_idx;index:Asset_organizationId_archivedAt_idx;index:Asset_createdAt_organizationId_idx;index:Asset_organizationId_createdAt_id_idx;index:Asset_valuation_organizationId_idx;index:Asset_categoryId_organizationId_idx;index:Asset_assetModelId_organizationId_idx;not null" json:"organizationId"`
	CategoryID           *string                 `gorm:"column:categoryId;index:Asset_categoryId_organizationId_idx" json:"categoryId"`
	PreferredBarcodeID   *string                 `gorm:"column:preferredBarcodeId;uniqueIndex:Asset_preferredBarcodeId_key" json:"preferredBarcodeId"`
	AssetModel           *AssetModel             `gorm:"belongsTo"`
	User                 *User                   `gorm:"belongsTo"`
	Organization         *Organization           `gorm:"belongsTo"`
	Category             *Category               `gorm:"belongsTo"`
	PreferredBarcode     *Barcode                `gorm:"foreignKey:PreferredBarcodeID;belongsTo"`
	Custody              []Custody               `gorm:"foreignKey:AssetID;constraint:Custody_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Notes                []Note                  `gorm:"foreignKey:AssetID;constraint:Note_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	QrCodes              []Qr                    `gorm:"foreignKey:AssetID;constraint:Qr_assetId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Barcodes             []Barcode               `gorm:"foreignKey:AssetID;constraint:Barcode_assetId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Reports              []ReportFound           `gorm:"foreignKey:AssetID;constraint:ReportFound_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomFields         []AssetCustomFieldValue `gorm:"foreignKey:AssetID;constraint:AssetCustomFieldValue_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BookingAssets        []BookingAsset          `gorm:"foreignKey:AssetID;constraint:BookingAsset_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConsumptionLogs      []ConsumptionLog        `gorm:"foreignKey:AssetID;constraint:ConsumptionLog_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Reminders            []AssetReminder         `gorm:"foreignKey:AssetID;constraint:AssetReminder_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuditEntries         []AuditAsset            `gorm:"foreignKey:AssetID;constraint:AuditAsset_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuditScans           []AuditScan             `gorm:"foreignKey:AssetID;constraint:AuditScan_assetId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Scans                []Scan                  `gorm:"foreignKey:AssetID;constraint:Scan_assetId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AssetKits            []AssetKit              `gorm:"foreignKey:AssetID;constraint:AssetKit_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetLocations       []AssetLocation         `gorm:"foreignKey:AssetID;constraint:AssetLocation_assetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags                 []Tag                   `gorm:"many2many:_AssetToTag;joinForeignKey:A;joinReferences:B"`
}

func (Asset) TableName() string {
	return "Asset"
}

func (m *Asset) BeforeCreate(tx *gorm.DB) error {
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

func (*Asset) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AssetFilterPreset struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string        `gorm:"column:organizationId;uniqueIndex:asset_filter_presets_owner_name_unique;index:asset_filter_presets_owner_lookup_idx;not null" json:"organizationId"`
	OwnerID        string        `gorm:"column:ownerId;uniqueIndex:asset_filter_presets_owner_name_unique;index:asset_filter_presets_owner_lookup_idx;not null" json:"ownerId"`
	Name           string        `gorm:"column:name;uniqueIndex:asset_filter_presets_owner_name_unique;not null" json:"name"`
	Query          string        `gorm:"column:query;not null" json:"query"`
	Starred        bool          `gorm:"column:starred;default:false;not null" json:"starred"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization `gorm:"belongsTo"`
	Owner          *User         `gorm:"foreignKey:OwnerID;belongsTo"`
}

func (AssetFilterPreset) TableName() string {
	return "AssetFilterPreset"
}

func (m *AssetFilterPreset) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetFilterPreset) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AssetIndexSettings struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	UserID         string         `gorm:"column:userId;uniqueIndex:AssetIndexSettings_userId_organizationId_key;not null" json:"userId"`
	OrganizationID string         `gorm:"column:organizationId;uniqueIndex:AssetIndexSettings_userId_organizationId_key;index:AssetIndexSettings_organizationId_idx;not null" json:"organizationId"`
	Mode           string         `gorm:"column:mode;default:'SIMPLE';not null" json:"mode"`
	Columns        datatypes.JSON `gorm:"column:columns;default:'[ {\"name\": \"id\", \"visible\": true, \"position\": 0}, {\"name\": \"status\", \"visible\": true, \"position\": 1}, {\"name\": \"description\", \"visible\": true, \"position\": 2}, {\"name\": \"valuation\", \"visible\": true, \"position\": 3}, {\"name\": \"createdAt\", \"visible\": true, \"position\": 4}, {\"name\": \"category\", \"visible\": true, \"position\": 5}, {\"name\": \"tags\", \"visible\": true, \"position\": 6}, {\"name\": \"location\", \"visible\": true, \"position\": 7}, {\"name\": \"kit\", \"visible\": true, \"position\": 8}, {\"name\": \"custody\", \"visible\": true, \"position\": 9} ]';not null" json:"columns"`
	FreezeColumn   *bool          `gorm:"column:freezeColumn;default:true;not null" json:"freezeColumn"`
	ShowAssetImage *bool          `gorm:"column:showAssetImage;default:true;not null" json:"showAssetImage"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User           *User          `gorm:"belongsTo"`
	Organization   *Organization  `gorm:"belongsTo"`
}

func (AssetIndexSettings) TableName() string {
	return "AssetIndexSettings"
}

func (m *AssetIndexSettings) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetIndexSettings) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A saved layout for a generated PDF, kept per team workspace under a name:
// which columns, rows and header lines print, and in what order the rows
// come. A row is a layout, not a generated file; the person generating a PDF
// picks a template by name.
type PdfTemplate struct {
	ID   string `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;uniqueIndex:PdfTemplate_organizationId_kind_name_key;not null" json:"name"`
	Kind string `gorm:"column:kind;uniqueIndex:PdfTemplate_organizationId_kind_name_key;default:'BOOKING';not null" json:"kind"`
	// The layout's switches. Read through `parseBookingTemplateConfig`,
	// which fills any key a stored row lacks from the Checklist template's
	// values.
	Config datatypes.JSON `gorm:"column:config;default:'{}';not null" json:"config"`
	// The template the PDF dialog offers first. One per workspace and kind,
	// enforced by a partial unique index in raw SQL (migration
	// 20261103120000_add_pdf_templates). Names are also unique per workspace
	// and kind ignoring case, through an expression index in the same
	// migration.
	IsDefault      bool          `gorm:"column:isDefault;default:false;not null" json:"isDefault"`
	OrganizationID string        `gorm:"column:organizationId;uniqueIndex:PdfTemplate_organizationId_kind_name_key;not null" json:"organizationId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (PdfTemplate) TableName() string {
	return "PdfTemplate"
}

func (m *PdfTemplate) BeforeCreate(tx *gorm.DB) error {
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

func (*PdfTemplate) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Category struct {
	ID                 string        `gorm:"column:id;primaryKey" json:"id"`
	Name               string        `gorm:"column:name;index:Category_name_trgm_idx;not null" json:"name"`
	Description        *string       `gorm:"column:description" json:"description"`
	Color              string        `gorm:"column:color;not null" json:"color"`
	CreatedAt          DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID             string        `gorm:"column:userId;index:Category_userId_idx;not null" json:"userId"`
	OrganizationID     string        `gorm:"column:organizationId;index:Category_organizationId_idx;not null" json:"organizationId"`
	User               *User         `gorm:"belongsTo"`
	Organization       *Organization `gorm:"belongsTo"`
	Assets             []Asset       `gorm:"foreignKey:CategoryID;constraint:Asset_categoryId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Kits               []Kit         `gorm:"foreignKey:CategoryID;constraint:Kit_categoryId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AssetModelDefaults []AssetModel  `gorm:"foreignKey:DefaultCategoryID;constraint:AssetModel_defaultCategoryId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	CustomFields       []CustomField `gorm:"many2many:_CategoryToCustomField;joinForeignKey:A;joinReferences:B"`
}

func (Category) TableName() string {
	return "Category"
}

func (m *Category) BeforeCreate(tx *gorm.DB) error {
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

func (*Category) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// *
// * A template/grouping entity for assets.
// * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
// * and provides default field values when creating new assets from the model.
type AssetModel struct {
	ID          string  `gorm:"column:id;primaryKey" json:"id"`
	Name        string  `gorm:"column:name;index:AssetModel_organizationId_name_idx;index:AssetModel_name_trgm_idx;not null" json:"name"`
	Description *string `gorm:"column:description" json:"description"`
	// Public URL (bucket `files`) of the model's cover image. Rendered by any
	// asset of this model that has no image of its own. Public — and therefore
	// never-expiring and cacheable — like Location.imageUrl; see
	// resolveAssetImage() for the cascade.
	Image *string `gorm:"column:image" json:"image"`
	// Public URL of the 108px thumbnail generated alongside `image`. Used at
	// list sizes so inheriting assets never render the full-size original.
	ThumbnailImage    *string       `gorm:"column:thumbnailImage" json:"thumbnailImage"`
	DefaultCategoryID *string       `gorm:"column:defaultCategoryId" json:"defaultCategoryId"`
	DefaultValuation  *float64      `gorm:"column:defaultValuation;type:double precision" json:"defaultValuation"`
	OrganizationID    string        `gorm:"column:organizationId;index:AssetModel_organizationId_name_idx;not null" json:"organizationId"`
	UserID            string        `gorm:"column:userId;index:AssetModel_userId_idx;not null" json:"userId"`
	CreatedAt         DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DefaultCategory   *Category     `gorm:"foreignKey:DefaultCategoryID;belongsTo"`
	Organization      *Organization `gorm:"belongsTo"`
	CreatedBy         *User         `gorm:"foreignKey:UserID;belongsTo"`
	// Default custom field values. Like the category and valuation defaults,
	// each one fills the matching field only when it is empty, at the moment
	// an asset joins this model. Nothing stays linked afterwards.
	CustomFieldDefaults []AssetModelCustomFieldDefault `gorm:"foreignKey:AssetModelID;constraint:AssetModelCustomFieldDefault_assetModelId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Assets              []Asset                        `gorm:"foreignKey:AssetModelID;constraint:Asset_assetModelId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ModelRequests       []BookingModelRequest          `gorm:"foreignKey:AssetModelID;constraint:BookingModelRequest_assetModelId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AssetModel) TableName() string {
	return "AssetModel"
}

func (m *AssetModel) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetModel) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Book-by-Model: intent row attached to a booking for a quantity of
// assets of a given AssetModel, without yet picking specific assets.
// Concrete `BookingAsset` rows are only created at
// scan-to-assign time; when the last unit is assigned, the request
// row is deleted. `quantity` is the number of units still needing
// assignment (starts at the full requested total, decrements on scan).
//
// Keeping model requests separate from `BookingAsset` means the ~40
// code paths that assume `BookingAsset.assetId` points to a real
// asset stay untouched — model-only bookings become concrete the
// moment they're checked out.
type BookingModelRequest struct {
	ID           string `gorm:"column:id;primaryKey" json:"id"`
	BookingID    string `gorm:"column:bookingId;uniqueIndex:BookingModelRequest_bookingId_assetModelId_key;not null" json:"bookingId"`
	AssetModelID string `gorm:"column:assetModelId;uniqueIndex:BookingModelRequest_bookingId_assetModelId_key;index:BookingModelRequest_assetModelId_idx;not null" json:"assetModelId"`
	// Total units originally reserved. Mutated only by explicit
	// operator edits (upsert) — never decremented on materialisation.
	// This preserves the audit trail of "how many were planned".
	Quantity int `gorm:"column:quantity;size:32;not null" json:"quantity"`
	// Units already assigned to concrete `BookingAsset` rows via
	// `materializeModelRequestForAsset`. Scans increment it.
	//
	// It DOES decrease: `removeAssets` (booking service) decrements it by the
	// number of removed `BookingAsset` rows whose `bookingModelRequestId`
	// points at this request, floored at 0, and clears `fulfilledAt` when the
	// result drops below `quantity`.
	//
	// Counted from that column, never from a shared `assetModelId`: an asset
	// added directly alongside a reservation shares the model without having
	// discharged anything, so grouping by model re-opened reservations that
	// were still correctly satisfied.
	FulfilledQuantity int `gorm:"column:fulfilledQuantity;size:32;default:0;not null" json:"fulfilledQuantity"`
	// Timestamp of the moment the reservation became fully fulfilled — set
	// when `fulfilledQuantity` catches up to `quantity`, and preserved (not
	// re-stamped) if an unchanged quantity is saved again. Cleared back to
	// null when outstanding units reappear: the operator edits `quantity`
	// upwards, or `removeAssets` drops `fulfilledQuantity` below it.
	//
	// The invariant that makes this the primary "is this outstanding?" filter:
	// non-null IFF `fulfilledQuantity === quantity`, so `fulfilledAt IS NULL`
	// means units remain to materialise. `fulfilledQuantity` can never exceed
	// `quantity` (the claim in `materializeModelRequestForAsset` is a single
	// statement guarded by `WHERE "fulfilledQuantity" < "quantity"`; upsert
	// rejects a shrink below it), so any decrement re-opens the row.
	FulfilledAt *DateTime   `gorm:"column:fulfilledAt;type:timestamp(3)" json:"fulfilledAt"`
	CreatedAt   DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime    `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Booking     *Booking    `gorm:"belongsTo"`
	AssetModel  *AssetModel `gorm:"belongsTo"`
}

func (BookingModelRequest) TableName() string {
	return "BookingModelRequest"
}

func (m *BookingModelRequest) BeforeCreate(tx *gorm.DB) error {
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

func (*BookingModelRequest) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Tag struct {
	ID             string             `gorm:"column:id;primaryKey" json:"id"`
	Name           string             `gorm:"column:name;index:Tag_name_trgm_idx;not null" json:"name"`
	Description    *string            `gorm:"column:description" json:"description"`
	Color          *string            `gorm:"column:color" json:"color"`
	UseFor         ScalarList[string] `gorm:"column:useFor;type:text[]" json:"useFor"`
	UserID         string             `gorm:"column:userId;index:Tag_userId_idx;not null" json:"userId"`
	OrganizationID string             `gorm:"column:organizationId;index:Tag_organizationId_idx;not null" json:"organizationId"`
	CreatedAt      DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User           *User              `gorm:"belongsTo"`
	Organization   *Organization      `gorm:"belongsTo"`
	Assets         []Asset            `gorm:"many2many:_AssetToTag;joinForeignKey:B;joinReferences:A"`
	Bookings       []Booking          `gorm:"many2many:_BookingToTag;joinForeignKey:B;joinReferences:A"`
}

func (Tag) TableName() string {
	return "Tag"
}

func (m *Tag) BeforeCreate(tx *gorm.DB) error {
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

func (*Tag) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Note struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Content   string   `gorm:"column:content;not null" json:"content"`
	Type      string   `gorm:"column:type;default:'COMMENT';not null" json:"type"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID    *string  `gorm:"column:userId;index:Note_userId_idx" json:"userId"`
	AssetID   string   `gorm:"column:assetId;index:Note_assetId_idx;not null" json:"assetId"`
	User      *User    `gorm:"belongsTo"`
	Asset     *Asset   `gorm:"belongsTo"`
}

func (Note) TableName() string {
	return "Note"
}

func (m *Note) BeforeCreate(tx *gorm.DB) error {
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

func (*Note) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BookingNote struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Content   string   `gorm:"column:content;not null" json:"content"`
	Type      string   `gorm:"column:type;default:'COMMENT';not null" json:"type"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID    *string  `gorm:"column:userId;index:BookingNote_userId_idx" json:"userId"`
	BookingID string   `gorm:"column:bookingId;index:BookingNote_bookingId_idx;not null" json:"bookingId"`
	User      *User    `gorm:"belongsTo"`
	Booking   *Booking `gorm:"belongsTo"`
}

func (BookingNote) TableName() string {
	return "BookingNote"
}

func (m *BookingNote) BeforeCreate(tx *gorm.DB) error {
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

func (*BookingNote) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type LocationNote struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	Content    string    `gorm:"column:content;not null" json:"content"`
	Type       string    `gorm:"column:type;default:'COMMENT';not null" json:"type"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID     *string   `gorm:"column:userId;index:LocationNote_userId_idx" json:"userId"`
	LocationID string    `gorm:"column:locationId;index:LocationNote_locationId_idx;not null" json:"locationId"`
	User       *User     `gorm:"belongsTo"`
	Location   *Location `gorm:"belongsTo"`
}

func (LocationNote) TableName() string {
	return "LocationNote"
}

func (m *LocationNote) BeforeCreate(tx *gorm.DB) error {
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

func (*LocationNote) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Qr struct {
	ID              string        `gorm:"column:id;primaryKey;index:Qr_id_trgm_idx" json:"id"`
	Version         int           `gorm:"column:version;size:32;default:0;not null" json:"version"`
	ErrorCorrection string        `gorm:"column:errorCorrection;default:'L';not null" json:"errorCorrection"`
	AssetID         *string       `gorm:"column:assetId;index:Qr_assetId_idx" json:"assetId"`
	KitID           *string       `gorm:"column:kitId;index:Qr_kitId_idx" json:"kitId"`
	UserID          *string       `gorm:"column:userId;index:Qr_userId_idx" json:"userId"`
	OrganizationID  *string       `gorm:"column:organizationId;index:Qr_organizationId_idx" json:"organizationId"`
	BatchID         *string       `gorm:"column:batchId;index:Qr_batchId_idx" json:"batchId"`
	CreatedAt       DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Asset           *Asset        `gorm:"belongsTo"`
	Kit             *Kit          `gorm:"belongsTo"`
	User            *User         `gorm:"belongsTo"`
	Organization    *Organization `gorm:"belongsTo"`
	Batch           *PrintBatch   `gorm:"foreignKey:BatchID;belongsTo"`
	Scans           []Scan        `gorm:"foreignKey:QrID;constraint:Scan_qrId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Qr) TableName() string {
	return "Qr"
}

func (m *Qr) BeforeCreate(tx *gorm.DB) error {
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

func (*Qr) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Barcode struct {
	ID               string        `gorm:"column:id;primaryKey" json:"id"`
	Value            string        `gorm:"column:value;uniqueIndex:Barcode_organizationId_value_key;index:Barcode_organizationId_value_idx;index:Barcode_value_trgm_idx;not null" json:"value"`
	Type             string        `gorm:"column:type;default:'Code128';not null" json:"type"`
	AssetID          *string       `gorm:"column:assetId;index:Barcode_assetId_idx" json:"assetId"`
	KitID            *string       `gorm:"column:kitId;index:Barcode_kitId_idx" json:"kitId"`
	OrganizationID   string        `gorm:"column:organizationId;uniqueIndex:Barcode_organizationId_value_key;index:Barcode_organizationId_value_idx;index:Barcode_organizationId_idx;not null" json:"organizationId"`
	CreatedAt        DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Asset            *Asset        `gorm:"belongsTo"`
	Kit              *Kit          `gorm:"belongsTo"`
	Organization     *Organization `gorm:"belongsTo"`
	Scans            []Scan        `gorm:"foreignKey:BarcodeID;constraint:Scan_barcodeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PreferredByAsset *Asset        `gorm:"foreignKey:PreferredBarcodeID;constraint:Asset_preferredBarcodeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Barcode) TableName() string {
	return "Barcode"
}

func (m *Barcode) BeforeCreate(tx *gorm.DB) error {
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

func (*Barcode) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PrintBatch struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Name      string   `gorm:"column:name;uniqueIndex:PrintBatch_name_key;not null" json:"name"`
	Printed   bool     `gorm:"column:printed;default:false;not null" json:"printed"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	QrCodes   []Qr     `gorm:"foreignKey:BatchID;constraint:Qr_batchId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (PrintBatch) TableName() string {
	return "PrintBatch"
}

func (m *PrintBatch) BeforeCreate(tx *gorm.DB) error {
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

func (*PrintBatch) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ReportFound struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Email     string   `gorm:"column:email;not null" json:"email"`
	Content   string   `gorm:"column:content;not null" json:"content"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	AssetID   *string  `gorm:"column:assetId;index:ReportFound_assetId_idx" json:"assetId"`
	KitID     *string  `gorm:"column:kitId;index:ReportFound_kitId_idx" json:"kitId"`
	Asset     *Asset   `gorm:"belongsTo"`
	Kit       *Kit     `gorm:"belongsTo"`
}

func (ReportFound) TableName() string {
	return "ReportFound"
}

func (m *ReportFound) BeforeCreate(tx *gorm.DB) error {
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

func (*ReportFound) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Scan struct {
	ID                string        `gorm:"column:id;primaryKey" json:"id"`
	Latitude          *string       `gorm:"column:latitude" json:"latitude"`
	Longitude         *string       `gorm:"column:longitude" json:"longitude"`
	UserAgent         *string       `gorm:"column:userAgent" json:"userAgent"`
	UserID            *string       `gorm:"column:userId;index:Scan_userId_idx" json:"userId"`
	QrID              *string       `gorm:"column:qrId;index:Scan_qrId_idx" json:"qrId"`
	RawQrID           *string       `gorm:"column:rawQrId;index:Scan_rawQrId_createdAt_idx" json:"rawQrId"`
	CodeType          string        `gorm:"column:codeType;default:'QR';not null" json:"codeType"`
	Code              *string       `gorm:"column:code" json:"code"`
	BarcodeID         *string       `gorm:"column:barcodeId;index:Scan_barcodeId_idx" json:"barcodeId"`
	AssetID           *string       `gorm:"column:assetId;index:Scan_assetId_createdAt_idx" json:"assetId"`
	KitID             *string       `gorm:"column:kitId;index:Scan_kitId_createdAt_idx" json:"kitId"`
	OrganizationID    string        `gorm:"column:organizationId;index:Scan_organizationId_idx;not null" json:"organizationId"`
	Source            *string       `gorm:"column:source" json:"source"`
	ManuallyGenerated bool          `gorm:"column:manuallyGenerated;default:false;not null" json:"manuallyGenerated"`
	CreatedAt         DateTime      `gorm:"column:createdAt;index:Scan_rawQrId_createdAt_idx;index:Scan_assetId_createdAt_idx;index:Scan_kitId_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User              *User         `gorm:"belongsTo"`
	Qr                *Qr           `gorm:"belongsTo"`
	Barcode           *Barcode      `gorm:"belongsTo"`
	Asset             *Asset        `gorm:"belongsTo"`
	Kit               *Kit          `gorm:"belongsTo"`
	Organization      *Organization `gorm:"belongsTo"`
}

func (Scan) TableName() string {
	return "Scan"
}

func (m *Scan) BeforeCreate(tx *gorm.DB) error {
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

func (*Scan) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Location struct {
	ID              string           `gorm:"column:id;primaryKey" json:"id"`
	Name            string           `gorm:"column:name;index:Location_name_trgm_idx;not null" json:"name"`
	Description     *string          `gorm:"column:description;index:Location_description_trgm_idx" json:"description"`
	Address         *string          `gorm:"column:address;index:Location_address_trgm_idx" json:"address"`
	Latitude        *float64         `gorm:"column:latitude;type:double precision" json:"latitude"`
	Longitude       *float64         `gorm:"column:longitude;type:double precision" json:"longitude"`
	ImageURL        *string          `gorm:"column:imageUrl" json:"imageUrl"`
	ThumbnailURL    *string          `gorm:"column:thumbnailUrl" json:"thumbnailUrl"`
	ImageID         *string          `gorm:"column:imageId;uniqueIndex:Location_imageId_key" json:"imageId"`
	CreatedAt       DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime         `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID          string           `gorm:"column:userId;index:Location_userId_idx;not null" json:"userId"`
	OrganizationID  string           `gorm:"column:organizationId;index:Location_organizationId_idx;index:Location_organizationId_parentId_idx;not null" json:"organizationId"`
	ParentID        *string          `gorm:"column:parentId;index:Location_organizationId_parentId_idx" json:"parentId"`
	Image           *Image           `gorm:"belongsTo"`
	User            *User            `gorm:"belongsTo"`
	Organization    *Organization    `gorm:"belongsTo"`
	Parent          *Location        `gorm:"foreignKey:ParentID;belongsTo"`
	AssetLocations  []AssetLocation  `gorm:"foreignKey:LocationID;constraint:AssetLocation_locationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Kits            []Kit            `gorm:"foreignKey:LocationID;constraint:Kit_locationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Notes           []LocationNote   `gorm:"foreignKey:LocationID;constraint:LocationNote_locationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Custodies       []Custody        `gorm:"foreignKey:LocationID;constraint:Custody_locationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ConsumptionLogs []ConsumptionLog `gorm:"foreignKey:LocationID;constraint:ConsumptionLog_locationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Children        []Location       `gorm:"foreignKey:ParentID;constraint:Location_parentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Location) TableName() string {
	return "Location"
}

func (m *Location) BeforeCreate(tx *gorm.DB) error {
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

func (*Location) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Role struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Name      string   `gorm:"column:name;uniqueIndex:Role_name_key;default:'USER';not null" json:"name"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Users     []User   `gorm:"many2many:_RoleToUser;joinForeignKey:A;joinReferences:B"`
}

func (Role) TableName() string {
	return "Role"
}

func (m *Role) BeforeCreate(tx *gorm.DB) error {
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

func (*Role) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TeamMember struct {
	ID                          string            `gorm:"column:id;primaryKey" json:"id"`
	Name                        string            `gorm:"column:name;index:TeamMember_name_idx;not null" json:"name"`
	OrganizationID              string            `gorm:"column:organizationId;index:TeamMember_organizationId_idx;not null" json:"organizationId"`
	UserID                      *string           `gorm:"column:userId;index:TeamMember_userId_idx" json:"userId"`
	CreatedAt                   DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                   DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt                   *DateTime         `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	Organization                *Organization     `gorm:"belongsTo"`
	User                        *User             `gorm:"belongsTo"`
	Custodies                   []Custody         `gorm:"foreignKey:TeamMemberID;constraint:Custody_teamMemberId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ReceivedInvites             []Invite          `gorm:"foreignKey:TeamMemberID;constraint:Invite_teamMemberId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	KitCustodies                []KitCustody      `gorm:"foreignKey:CustodianID;constraint:KitCustody_custodianId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Bookings                    []Booking         `gorm:"foreignKey:CustodianTeamMemberID;constraint:Booking_custodianTeamMemberId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConsumptionLogs             []ConsumptionLog  `gorm:"foreignKey:CustodianID;constraint:ConsumptionLog_custodianId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TeamMemberNotes             []TeamMemberNote  `gorm:"foreignKey:TeamMemberID;constraint:TeamMemberNote_teamMemberId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetReminders              []AssetReminder   `gorm:"many2many:_AssetReminderToTeamMember;joinForeignKey:B;joinReferences:A"`
	AlwaysNotifyBookingSettings []BookingSettings `gorm:"many2many:_BookingSettingsAlwaysNotify;joinForeignKey:B;joinReferences:A"`
	BookingNotifications        []Booking         `gorm:"many2many:_BookingNotificationRecipients;joinForeignKey:B;joinReferences:A"`
}

func (TeamMember) TableName() string {
	return "TeamMember"
}

func (m *TeamMember) BeforeCreate(tx *gorm.DB) error {
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

func (*TeamMember) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Custody struct {
	ID            string      `gorm:"column:id;primaryKey" json:"id"`
	TeamMemberID  string      `gorm:"column:teamMemberId;index:Custody_assetId_teamMemberId_idx;index:Custody_teamMemberId_idx;not null" json:"teamMemberId"`
	AssetID       string      `gorm:"column:assetId;index:Custody_assetId_teamMemberId_idx;not null" json:"assetId"`
	KitCustodyID  *string     `gorm:"column:kitCustodyId;index:Custody_kitCustodyId_idx" json:"kitCustodyId"`
	Quantity      *int        `gorm:"column:quantity;size:32;default:1;not null" json:"quantity"`
	LocationID    *string     `gorm:"column:locationId;index:Custody_locationId_idx" json:"locationId"`
	SourceUnknown bool        `gorm:"column:sourceUnknown;default:false;not null" json:"sourceUnknown"`
	CreatedAt     DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime    `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Custodian     *TeamMember `gorm:"foreignKey:TeamMemberID;belongsTo"`
	Asset         *Asset      `gorm:"belongsTo"`
	KitCustody    *KitCustody `gorm:"belongsTo"`
	Location      *Location   `gorm:"belongsTo"`
}

func (Custody) TableName() string {
	return "Custody"
}

func (m *Custody) BeforeCreate(tx *gorm.DB) error {
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

func (*Custody) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Organization struct {
	ID                        string               `gorm:"column:id;primaryKey" json:"id"`
	Name                      *string              `gorm:"column:name;default:'Personal';not null" json:"name"`
	Type                      string               `gorm:"column:type;default:'PERSONAL';not null" json:"type"`
	UserID                    string               `gorm:"column:userId;index:Organization_userId_idx;not null" json:"userId"`
	Currency                  string               `gorm:"column:currency;default:'USD';not null" json:"currency"`
	ImageID                   *string              `gorm:"column:imageId;uniqueIndex:Organization_imageId_key" json:"imageId"`
	EnabledSso                bool                 `gorm:"column:enabledSso;default:false;not null" json:"enabledSso"`
	SsoDetailsID              *string              `gorm:"column:ssoDetailsId;index:Organization_ssoDetailsId_idx" json:"ssoDetailsId"`
	SelfServiceCanSeeCustody  bool                 `gorm:"column:selfServiceCanSeeCustody;default:false;not null" json:"selfServiceCanSeeCustody"`
	SelfServiceCanSeeBookings bool                 `gorm:"column:selfServiceCanSeeBookings;default:false;not null" json:"selfServiceCanSeeBookings"`
	BaseUserCanSeeCustody     bool                 `gorm:"column:baseUserCanSeeCustody;default:false;not null" json:"baseUserCanSeeCustody"`
	BaseUserCanSeeBookings    bool                 `gorm:"column:baseUserCanSeeBookings;default:false;not null" json:"baseUserCanSeeBookings"`
	ManagersCanEditAssets     bool                 `gorm:"column:managersCanEditAssets;default:false;not null" json:"managersCanEditAssets"`
	BarcodesEnabled           bool                 `gorm:"column:barcodesEnabled;default:false;not null" json:"barcodesEnabled"`
	BarcodesEnabledAt         *DateTime            `gorm:"column:barcodesEnabledAt;type:timestamp(3)" json:"barcodesEnabledAt"`
	UsedBarcodeTrial          bool                 `gorm:"column:usedBarcodeTrial;default:false;not null" json:"usedBarcodeTrial"`
	AuditsEnabled             bool                 `gorm:"column:auditsEnabled;default:false;not null" json:"auditsEnabled"`
	AuditsEnabledAt           *DateTime            `gorm:"column:auditsEnabledAt;type:timestamp(3)" json:"auditsEnabledAt"`
	UsedAuditTrial            bool                 `gorm:"column:usedAuditTrial;default:false;not null" json:"usedAuditTrial"`
	WorkspaceDisabled         bool                 `gorm:"column:workspaceDisabled;default:false;not null" json:"workspaceDisabled"`
	CreatedAt                 DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime             `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	HasSequentialIdsMigrated  bool                 `gorm:"column:hasSequentialIdsMigrated;default:false;not null" json:"hasSequentialIdsMigrated"`
	QrIDDisplayPreference     string               `gorm:"column:qrIdDisplayPreference;default:'QR_ID';not null" json:"qrIdDisplayPreference"`
	ShowShelfBranding         *bool                `gorm:"column:showShelfBranding;default:true;not null" json:"showShelfBranding"`
	ShowQrCodesOnPdfs         *bool                `gorm:"column:showQrCodesOnPdfs;default:true;not null" json:"showQrCodesOnPdfs"`
	CustomEmailFooter         *string              `gorm:"column:customEmailFooter;type:varchar(500)" json:"customEmailFooter"`
	Owner                     *User                `gorm:"foreignKey:UserID;belongsTo"`
	Image                     *Image               `gorm:"belongsTo"`
	SsoDetails                *SsoDetails          `gorm:"belongsTo"`
	Members                   []TeamMember         `gorm:"foreignKey:OrganizationID;constraint:TeamMember_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Assets                    []Asset              `gorm:"foreignKey:OrganizationID;constraint:Asset_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetKits                 []AssetKit           `gorm:"foreignKey:OrganizationID;constraint:AssetKit_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetLocations            []AssetLocation      `gorm:"foreignKey:OrganizationID;constraint:AssetLocation_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Scans                     []Scan               `gorm:"foreignKey:OrganizationID;constraint:Scan_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Locations                 []Location           `gorm:"foreignKey:OrganizationID;constraint:Location_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Categories                []Category           `gorm:"foreignKey:OrganizationID;constraint:Category_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags                      []Tag                `gorm:"foreignKey:OrganizationID;constraint:Tag_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	QrCodes                   []Qr                 `gorm:"foreignKey:OrganizationID;constraint:Qr_organizationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Barcodes                  []Barcode            `gorm:"foreignKey:OrganizationID;constraint:Barcode_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites                   []Invite             `gorm:"foreignKey:OrganizationID;constraint:Invite_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserOrganizations         []UserOrganization   `gorm:"foreignKey:OrganizationID;constraint:UserOrganization_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomFields              []CustomField        `gorm:"foreignKey:OrganizationID;constraint:CustomField_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Images                    []Image              `gorm:"foreignKey:OwnerOrgID;constraint:Image_ownerOrgId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Bookings                  []Booking            `gorm:"foreignKey:OrganizationID;constraint:Booking_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Kits                      []Kit                `gorm:"foreignKey:OrganizationID;constraint:Kit_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetModels               []AssetModel         `gorm:"foreignKey:OrganizationID;constraint:AssetModel_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetIndexSettings        []AssetIndexSettings `gorm:"foreignKey:OrganizationID;constraint:AssetIndexSettings_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetReminders            []AssetReminder      `gorm:"foreignKey:OrganizationID;constraint:AssetReminder_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AssetFilterPresets        []AssetFilterPreset  `gorm:"foreignKey:OrganizationID;constraint:AssetFilterPreset_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuditSessions             []AuditSession       `gorm:"foreignKey:OrganizationID;constraint:AuditSession_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuditImages               []AuditImage         `gorm:"foreignKey:OrganizationID;constraint:AuditImage_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetImports              []AssetImport        `gorm:"foreignKey:OrganizationID;constraint:AssetImport_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	LinkPreviews              []LinkPreview        `gorm:"foreignKey:OrganizationID;constraint:LinkPreview_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PdfTemplates              []PdfTemplate        `gorm:"foreignKey:OrganizationID;constraint:PdfTemplate_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UsersWithLastSelected     []User               `gorm:"foreignKey:LastSelectedOrganizationID;constraint:User_lastSelectedOrganizationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ScimTokens                []ScimToken          `gorm:"foreignKey:OrganizationID;constraint:ScimToken_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ScimExternalIds           []UserScimExternalID `gorm:"foreignKey:OrganizationID;constraint:UserScimExternalId_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleChangeLogs            []RoleChangeLog      `gorm:"foreignKey:OrganizationID;constraint:RoleChangeLog_organizationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	TeamMemberNotes           []TeamMemberNote     `gorm:"foreignKey:OrganizationID;constraint:TeamMemberNote_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ActivityEvents            []ActivityEvent      `gorm:"foreignKey:OrganizationID;constraint:ActivityEvent_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkingHours              *WorkingHours        `gorm:"foreignKey:OrganizationID;constraint:WorkingHours_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BookingSettings           *BookingSettings     `gorm:"foreignKey:OrganizationID;constraint:BookingSettings_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type UserOrganization struct {
	ID             string             `gorm:"column:id;primaryKey" json:"id"`
	UserID         string             `gorm:"column:userId;uniqueIndex:UserOrganization_userId_organizationId_key;not null" json:"userId"`
	OrganizationID string             `gorm:"column:organizationId;uniqueIndex:UserOrganization_userId_organizationId_key;index:UserOrganization_organizationId_idx;not null" json:"organizationId"`
	Roles          ScalarList[string] `gorm:"column:roles;type:text[]" json:"roles"`
	CreatedAt      DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	// Secret, unguessable token for this member's iCal calendar-feed subscription
	// in this workspace. Null until the user generates a feed; regenerating it
	// rotates the value (invalidating any previously-shared subscription URL).
	// `@unique` doubles as the lookup index for the public feed route.
	CalendarTokenID *string       `gorm:"column:calendarTokenId;uniqueIndex:UserOrganization_calendarTokenId_key" json:"calendarTokenId"`
	User            *User         `gorm:"belongsTo"`
	Organization    *Organization `gorm:"belongsTo"`
}

func (UserOrganization) TableName() string {
	return "UserOrganization"
}

func (m *UserOrganization) BeforeCreate(tx *gorm.DB) error {
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

func (*UserOrganization) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SsoDetails struct {
	ID                 string         `gorm:"column:id;primaryKey" json:"id"`
	Domain             string         `gorm:"column:domain;not null" json:"domain"`
	BaseUserGroupID    *string        `gorm:"column:baseUserGroupId" json:"baseUserGroupId"`
	SelfServiceGroupID *string        `gorm:"column:selfServiceGroupId" json:"selfServiceGroupId"`
	AdminGroupID       *string        `gorm:"column:adminGroupId" json:"adminGroupId"`
	ManagerGroupID     *string        `gorm:"column:managerGroupId" json:"managerGroupId"`
	RequireSsoLogin    *bool          `gorm:"column:requireSsoLogin;default:true;not null" json:"requireSsoLogin"`
	CreatedAt          DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organizations      []Organization `gorm:"foreignKey:SsoDetailsID;constraint:Organization_ssoDetailsId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (SsoDetails) TableName() string {
	return "SsoDetails"
}

func (m *SsoDetails) BeforeCreate(tx *gorm.DB) error {
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

func (*SsoDetails) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ScimToken struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	TokenHash      string        `gorm:"column:tokenHash;uniqueIndex:ScimToken_tokenHash_key;not null" json:"tokenHash"`
	Label          string        `gorm:"column:label;not null" json:"label"`
	OrganizationID string        `gorm:"column:organizationId;index:ScimToken_organizationId_idx;not null" json:"organizationId"`
	CreatedByID    string        `gorm:"column:createdById;not null" json:"createdById"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	LastUsedAt     *DateTime     `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (ScimToken) TableName() string {
	return "ScimToken"
}

func (m *ScimToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// *
// * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
// * object within a specific organization.  A single Shelf user can appear in
// * multiple organizations, each managed by a different identity provider, so the
// * external ID must be scoped to the (user, organization) pair.
type UserScimExternalID struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	UserID         string        `gorm:"column:userId;uniqueIndex:UserScimExternalId_userId_organizationId_key;not null" json:"userId"`
	OrganizationID string        `gorm:"column:organizationId;uniqueIndex:UserScimExternalId_userId_organizationId_key;uniqueIndex:UserScimExternalId_organizationId_scimExternalId_key;index:UserScimExternalId_organizationId_idx;not null" json:"organizationId"`
	ScimExternalID string        `gorm:"column:scimExternalId;uniqueIndex:UserScimExternalId_organizationId_scimExternalId_key;not null" json:"scimExternalId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User           *User         `gorm:"belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (UserScimExternalID) TableName() string {
	return "UserScimExternalId"
}

func (m *UserScimExternalID) BeforeCreate(tx *gorm.DB) error {
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

func (*UserScimExternalID) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Tier struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name;not null" json:"name"`
	TierLimitID *string    `gorm:"column:tierLimitId;uniqueIndex:Tier_tierLimitId_key" json:"tierLimitId"`
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	TierLimit   *TierLimit `gorm:"belongsTo"`
	Subscribers []User     `gorm:"foreignKey:TierID;constraint:User_tierId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Tier) TableName() string {
	return "Tier"
}

func (m *Tier) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Tier) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TierLimit struct {
	ID                       string   `gorm:"column:id;primaryKey" json:"id"`
	CanImportAssets          bool     `gorm:"column:canImportAssets;default:false;not null" json:"canImportAssets"`
	CanExportAssets          bool     `gorm:"column:canExportAssets;default:false;not null" json:"canExportAssets"`
	CanImportNRM             bool     `gorm:"column:canImportNRM;default:false;not null" json:"canImportNRM"`
	CanHideShelfBranding     bool     `gorm:"column:canHideShelfBranding;default:false;not null" json:"canHideShelfBranding"`
	CanUseReports            bool     `gorm:"column:canUseReports;default:false;not null" json:"canUseReports"`
	CanUseRecurringReminders bool     `gorm:"column:canUseRecurringReminders;default:false;not null" json:"canUseRecurringReminders"`
	CanUsePushNotifications  bool     `gorm:"column:canUsePushNotifications;default:false;not null" json:"canUsePushNotifications"`
	MaxCustomFields          int      `gorm:"column:maxCustomFields;size:32;default:0;not null" json:"maxCustomFields"`
	MaxOrganizations         *int     `gorm:"column:maxOrganizations;size:32;default:1;not null" json:"maxOrganizations"`
	CreatedAt                DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Tier                     *Tier    `gorm:"foreignKey:TierLimitID;constraint:Tier_tierLimitId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (TierLimit) TableName() string {
	return "TierLimit"
}

func (m *TierLimit) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TierLimit) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CustomTierLimit struct {
	ID                       string   `gorm:"column:id;primaryKey" json:"id"`
	UserID                   *string  `gorm:"column:userId;uniqueIndex:CustomTierLimit_userId_key" json:"userId"`
	CanImportAssets          *bool    `gorm:"column:canImportAssets;default:true;not null" json:"canImportAssets"`
	CanExportAssets          *bool    `gorm:"column:canExportAssets;default:true;not null" json:"canExportAssets"`
	CanImportNRM             *bool    `gorm:"column:canImportNRM;default:true;not null" json:"canImportNRM"`
	CanHideShelfBranding     *bool    `gorm:"column:canHideShelfBranding;default:true;not null" json:"canHideShelfBranding"`
	CanUseReports            *bool    `gorm:"column:canUseReports;default:true;not null" json:"canUseReports"`
	CanUseRecurringReminders *bool    `gorm:"column:canUseRecurringReminders;default:true;not null" json:"canUseRecurringReminders"`
	CanUsePushNotifications  *bool    `gorm:"column:canUsePushNotifications;default:true;not null" json:"canUsePushNotifications"`
	MaxCustomFields          *int     `gorm:"column:maxCustomFields;size:32;default:1000;not null" json:"maxCustomFields"`
	MaxOrganizations         *int     `gorm:"column:maxOrganizations;size:32;default:1;not null" json:"maxOrganizations"`
	IsEnterprise             bool     `gorm:"column:isEnterprise;default:false;not null" json:"isEnterprise"`
	CreatedAt                DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User                     *User    `gorm:"belongsTo"`
}

func (CustomTierLimit) TableName() string {
	return "CustomTierLimit"
}

func (m *CustomTierLimit) BeforeCreate(tx *gorm.DB) error {
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

func (*CustomTierLimit) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CustomField struct {
	ID       string             `gorm:"column:id;primaryKey" json:"id"`
	Name     string             `gorm:"column:name;not null" json:"name"`
	HelpText *string            `gorm:"column:helpText" json:"helpText"`
	Required bool               `gorm:"column:required;default:false;not null" json:"required"`
	Active   *bool              `gorm:"column:active;default:true;not null" json:"active"`
	Type     string             `gorm:"column:type;default:'TEXT';not null" json:"type"`
	Options  ScalarList[string] `gorm:"column:options;type:text[]" json:"options"`
	// How a TEXT field's values are read. LINK requires a web link in every
	// value; any other type stays PLAIN (enforced by a CHECK constraint).
	Format                  string                  `gorm:"column:format;default:'PLAIN';not null" json:"format"`
	OrganizationID          string                  `gorm:"column:organizationId;index:CustomField_organizationId_idx;index:CustomField_organizationId_deletedAt_idx;not null" json:"organizationId"`
	UserID                  string                  `gorm:"column:userId;index:CustomField_userId_idx;not null" json:"userId"`
	CreatedAt               DateTime                `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt               DateTime                `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt               *DateTime               `gorm:"column:deletedAt;index:CustomField_organizationId_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	Organization            *Organization           `gorm:"belongsTo"`
	CreatedBy               *User                   `gorm:"foreignKey:UserID;belongsTo"`
	AssetCustomFieldsValues []AssetCustomFieldValue `gorm:"foreignKey:CustomFieldID;constraint:AssetCustomFieldValue_customFieldId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Categories              []Category              `gorm:"many2many:_CategoryToCustomField;joinForeignKey:B;joinReferences:A"`
}

func (CustomField) TableName() string {
	return "CustomField"
}

func (m *CustomField) BeforeCreate(tx *gorm.DB) error {
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

func (*CustomField) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AssetCustomFieldValue struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	Value         datatypes.JSON `gorm:"column:value;not null" json:"value"`
	AssetID       string         `gorm:"column:assetId;index:AssetCustomFieldValue_lookup_idx;not null" json:"assetId"`
	CustomFieldID string         `gorm:"column:customFieldId;index:AssetCustomFieldValue_lookup_idx;index:AssetCustomFieldValue_customFieldId_idx;not null" json:"customFieldId"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Asset         *Asset         `gorm:"belongsTo"`
	CustomField   *CustomField   `gorm:"belongsTo"`
}

func (AssetCustomFieldValue) TableName() string {
	return "AssetCustomFieldValue"
}

func (m *AssetCustomFieldValue) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetCustomFieldValue) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A custom field value an AssetModel gives the assets that join it.
//
// `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
// asset is a plain copy. Rows cascade away with their model or field.
type AssetModelCustomFieldDefault struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	Value         datatypes.JSON `gorm:"column:value;not null" json:"value"`
	AssetModelID  string         `gorm:"column:assetModelId;uniqueIndex:AssetModelCustomFieldDefault_assetModelId_customFieldId_key;not null" json:"assetModelId"`
	CustomFieldID string         `gorm:"column:customFieldId;uniqueIndex:AssetModelCustomFieldDefault_assetModelId_customFieldId_key;index:AssetModelCustomFieldDefault_customFieldId_idx;not null" json:"customFieldId"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	AssetModel    *AssetModel    `gorm:"belongsTo"`
}

func (AssetModelCustomFieldDefault) TableName() string {
	return "AssetModelCustomFieldDefault"
}

func (m *AssetModelCustomFieldDefault) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetModelCustomFieldDefault) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Invite struct {
	ID                string             `gorm:"column:id;primaryKey" json:"id"`
	InviterID         string             `gorm:"column:inviterId;index:Invite_inviterId_idx;not null" json:"inviterId"`
	OrganizationID    string             `gorm:"column:organizationId;index:Invite_organizationId_idx;not null" json:"organizationId"`
	InviteeUserID     *string            `gorm:"column:inviteeUserId;index:Invite_inviteeUserId_idx" json:"inviteeUserId"`
	TeamMemberID      string             `gorm:"column:teamMemberId;index:Invite_teamMemberId_idx;not null" json:"teamMemberId"`
	InviteeEmail      string             `gorm:"column:inviteeEmail;not null" json:"inviteeEmail"`
	Status            string             `gorm:"column:status;default:'PENDING';not null" json:"status"`
	InviteCode        string             `gorm:"column:inviteCode;not null" json:"inviteCode"`
	Roles             ScalarList[string] `gorm:"column:roles;type:text[]" json:"roles"`
	InviteMessage     *string            `gorm:"column:inviteMessage;type:varchar(1000)" json:"inviteMessage"`
	ExpiresAt         DateTime           `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt         DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Inviter           *User              `gorm:"foreignKey:InviterID;belongsTo"`
	Organization      *Organization      `gorm:"belongsTo"`
	InviteeUser       *User              `gorm:"foreignKey:InviteeUserID;belongsTo"`
	InviteeTeamMember *TeamMember        `gorm:"foreignKey:TeamMemberID;belongsTo"`
}

func (Invite) TableName() string {
	return "Invite"
}

func (m *Invite) BeforeCreate(tx *gorm.DB) error {
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

func (*Invite) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Announcement struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Name      string   `gorm:"column:name;not null" json:"name"`
	Content   string   `gorm:"column:content;not null" json:"content"`
	Link      *string  `gorm:"column:link" json:"link"`
	LinkText  *string  `gorm:"column:linkText" json:"linkText"`
	Published bool     `gorm:"column:published;default:false;not null" json:"published"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Announcement) TableName() string {
	return "Announcement"
}

func (m *Announcement) BeforeCreate(tx *gorm.DB) error {
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

func (*Announcement) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// *
// * Audit trail for all quantity changes on quantity-tracked assets.
// * Every event records who performed it, who received/returned items,
// * which booking it relates to, and the category of change.
// * Quantity is always positive — direction is determined by category.
type ConsumptionLog struct {
	ID             string      `gorm:"column:id;primaryKey" json:"id"`
	AssetID        string      `gorm:"column:assetId;index:ConsumptionLog_assetId_createdAt_idx;not null" json:"assetId"`
	Category       string      `gorm:"column:category;not null" json:"category"`
	Quantity       int         `gorm:"column:quantity;size:32;not null" json:"quantity"`
	Note           *string     `gorm:"column:note" json:"note"`
	UserID         string      `gorm:"column:userId;index:ConsumptionLog_userId_idx;not null" json:"userId"`
	BookingID      *string     `gorm:"column:bookingId;index:ConsumptionLog_bookingId_idx" json:"bookingId"`
	CustodianID    *string     `gorm:"column:custodianId;index:ConsumptionLog_custodianId_idx" json:"custodianId"`
	BookingAssetID *string     `gorm:"column:bookingAssetId;index:ConsumptionLog_bookingAssetId_idx" json:"bookingAssetId"`
	LocationID     *string     `gorm:"column:locationId;index:ConsumptionLog_locationId_idx" json:"locationId"`
	StockChange    *int        `gorm:"column:stockChange;size:32" json:"stockChange"`
	CreatedAt      DateTime    `gorm:"column:createdAt;index:ConsumptionLog_assetId_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Asset          *Asset      `gorm:"belongsTo"`
	PerformedBy    *User       `gorm:"foreignKey:UserID;belongsTo"`
	Booking        *Booking    `gorm:"belongsTo"`
	Custodian      *TeamMember `gorm:"foreignKey:CustodianID;belongsTo"`
	Location       *Location   `gorm:"belongsTo"`
}

func (ConsumptionLog) TableName() string {
	return "ConsumptionLog"
}

func (m *ConsumptionLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// *
// * Explicit pivot table between Booking and Asset.
// * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
type BookingAsset struct {
	ID                    string    `gorm:"column:id;primaryKey" json:"id"`
	BookingID             string    `gorm:"column:bookingId;not null" json:"bookingId"`
	AssetID               string    `gorm:"column:assetId;index:BookingAsset_assetId_idx;not null" json:"assetId"`
	Quantity              *int      `gorm:"column:quantity;size:32;default:1;not null" json:"quantity"`
	BookingModelRequestID *string   `gorm:"column:bookingModelRequestId" json:"bookingModelRequestId"`
	AssetKitID            *string   `gorm:"column:assetKitId;index:BookingAsset_assetKitId_idx" json:"assetKitId"`
	SourceKitID           *string   `gorm:"column:sourceKitId;index:BookingAsset_sourceKitId_idx" json:"sourceKitId"`
	CheckedOutAt          *DateTime `gorm:"column:checkedOutAt;type:timestamptz(3)" json:"checkedOutAt"`
	CheckedOutByID        *string   `gorm:"column:checkedOutById" json:"checkedOutById"`
	CheckedInAt           *DateTime `gorm:"column:checkedInAt;type:timestamptz(3)" json:"checkedInAt"`
	CheckedInByID         *string   `gorm:"column:checkedInById" json:"checkedInById"`
	CheckedOutQuantity    int       `gorm:"column:checkedOutQuantity;size:32;default:0;not null" json:"checkedOutQuantity"`
	SourceLocationID      *string   `gorm:"column:sourceLocationId;index:BookingAsset_sourceLocationId_idx" json:"sourceLocationId"`
	Booking               *Booking  `gorm:"belongsTo"`
	Asset                 *Asset    `gorm:"belongsTo"`
}

func (BookingAsset) TableName() string {
	return "BookingAsset"
}

func (m *BookingAsset) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type Booking struct {
	ID                       string                   `gorm:"column:id;primaryKey;index:Booking_organizationId_from_id_idx" json:"id"`
	Name                     string                   `gorm:"column:name;not null" json:"name"`
	Status                   string                   `gorm:"column:status;index:Booking_status_idx;default:'DRAFT';not null" json:"status"`
	Description              *string                  `gorm:"column:description;default:''" json:"description"`
	ActiveSchedulerReference *string                  `gorm:"column:activeSchedulerReference" json:"activeSchedulerReference"`
	CreatorID                string                   `gorm:"column:creatorId;index:Booking_creatorId_idx;not null" json:"creatorId"`
	CustodianUserID          *string                  `gorm:"column:custodianUserId;index:Booking_custodianUserId_idx" json:"custodianUserId"`
	CustodianTeamMemberID    *string                  `gorm:"column:custodianTeamMemberId;index:Booking_custodianTeamMemberId_idx" json:"custodianTeamMemberId"`
	OrganizationID           string                   `gorm:"column:organizationId;index:Booking_organizationId_idx;index:Booking_organizationId_from_id_idx;not null" json:"organizationId"`
	CreatedAt                DateTime                 `gorm:"column:createdAt;type:timestamptz(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime                 `gorm:"column:updatedAt;type:timestamptz(3);autoUpdateTime:false;not null" json:"updatedAt"`
	From                     DateTime                 `gorm:"column:from;index:Booking_organizationId_from_id_idx;type:timestamptz(3);not null" json:"from"`
	To                       DateTime                 `gorm:"column:to;type:timestamptz(3);not null" json:"to"`
	OriginalFrom             *DateTime                `gorm:"column:originalFrom;type:timestamptz(3)" json:"originalFrom"`
	OriginalTo               *DateTime                `gorm:"column:originalTo;type:timestamptz(3)" json:"originalTo"`
	AutoArchivedAt           *DateTime                `gorm:"column:autoArchivedAt;type:timestamptz(3)" json:"autoArchivedAt"`
	ArchivedWithoutCheckin   bool                     `gorm:"column:archivedWithoutCheckin;default:false;not null" json:"archivedWithoutCheckin"`
	CancellationReason       *string                  `gorm:"column:cancellationReason" json:"cancellationReason"`
	Creator                  *User                    `gorm:"foreignKey:CreatorID;belongsTo"`
	CustodianUser            *User                    `gorm:"foreignKey:CustodianUserID;belongsTo"`
	CustodianTeamMember      *TeamMember              `gorm:"foreignKey:CustodianTeamMemberID;belongsTo"`
	Organization             *Organization            `gorm:"belongsTo"`
	BookingAssets            []BookingAsset           `gorm:"foreignKey:BookingID;constraint:BookingAsset_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConsumptionLogs          []ConsumptionLog         `gorm:"foreignKey:BookingID;constraint:ConsumptionLog_bookingId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Notes                    []BookingNote            `gorm:"foreignKey:BookingID;constraint:BookingNote_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ModelRequests            []BookingModelRequest    `gorm:"foreignKey:BookingID;constraint:BookingModelRequest_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PartialCheckins          []PartialBookingCheckin  `gorm:"foreignKey:BookingID;constraint:PartialBookingCheckin_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PartialCheckouts         []PartialBookingCheckout `gorm:"foreignKey:BookingID;constraint:PartialBookingCheckout_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags                     []Tag                    `gorm:"many2many:_BookingToTag;joinForeignKey:A;joinReferences:B"`
	NotificationRecipients   []TeamMember             `gorm:"many2many:_BookingNotificationRecipients;joinForeignKey:A;joinReferences:B"`
}

func (Booking) TableName() string {
	return "Booking"
}

func (m *Booking) BeforeCreate(tx *gorm.DB) error {
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

func (*Booking) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BookingSettings struct {
	ID                                    string        `gorm:"column:id;primaryKey" json:"id"`
	BufferStartTime                       int           `gorm:"column:bufferStartTime;size:32;default:0;not null" json:"bufferStartTime"`
	TagsRequired                          bool          `gorm:"column:tagsRequired;default:false;not null" json:"tagsRequired"`
	MaxBookingLength                      *int          `gorm:"column:maxBookingLength;size:32" json:"maxBookingLength"`
	MaxBookingLengthSkipClosedDays        bool          `gorm:"column:maxBookingLengthSkipClosedDays;default:false;not null" json:"maxBookingLengthSkipClosedDays"`
	AutoArchiveBookings                   bool          `gorm:"column:autoArchiveBookings;default:false;not null" json:"autoArchiveBookings"`
	AutoArchiveDays                       *int          `gorm:"column:autoArchiveDays;size:32;default:2;not null" json:"autoArchiveDays"`
	AutoArchiveExpiredReservations        bool          `gorm:"column:autoArchiveExpiredReservations;default:false;not null" json:"autoArchiveExpiredReservations"`
	RequireExplicitCheckinForAdmin        bool          `gorm:"column:requireExplicitCheckinForAdmin;default:false;not null" json:"requireExplicitCheckinForAdmin"`
	RequireExplicitCheckinForSelfService  bool          `gorm:"column:requireExplicitCheckinForSelfService;default:false;not null" json:"requireExplicitCheckinForSelfService"`
	RequireExplicitCheckoutForAdmin       bool          `gorm:"column:requireExplicitCheckoutForAdmin;default:false;not null" json:"requireExplicitCheckoutForAdmin"`
	RequireExplicitCheckoutForSelfService bool          `gorm:"column:requireExplicitCheckoutForSelfService;default:false;not null" json:"requireExplicitCheckoutForSelfService"`
	CountKitsAsSingleUnit                 bool          `gorm:"column:countKitsAsSingleUnit;default:false;not null" json:"countKitsAsSingleUnit"`
	NotifyBookingCreator                  *bool         `gorm:"column:notifyBookingCreator;default:true;not null" json:"notifyBookingCreator"`
	NotifyAdminsOnNewBooking              *bool         `gorm:"column:notifyAdminsOnNewBooking;default:true;not null" json:"notifyAdminsOnNewBooking"`
	OrganizationID                        string        `gorm:"column:organizationId;uniqueIndex:BookingSettings_organizationId_key;index:BookingSettings_organizationId_idx;not null" json:"organizationId"`
	CreatedAt                             DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                             DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization                          *Organization `gorm:"belongsTo"`
	AlwaysNotifyTeamMembers               []TeamMember  `gorm:"many2many:_BookingSettingsAlwaysNotify;joinForeignKey:A;joinReferences:B"`
}

func (BookingSettings) TableName() string {
	return "BookingSettings"
}

func (m *BookingSettings) BeforeCreate(tx *gorm.DB) error {
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

func (*BookingSettings) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A progressive check-in SESSION, mirroring PartialBookingCheckout.
//
// `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
// QUANTITY_TRACKED only once its remaining units hit zero (partial returns
// live in `ConsumptionLog`). Completing a booking outright writes no row here,
// so the same warning applies — absence is not evidence of anything.
// `BookingAsset.checkedInAt` is the per-slice fact.
type PartialBookingCheckin struct {
	ID               string             `gorm:"column:id;primaryKey" json:"id"`
	AssetIds         ScalarList[string] `gorm:"column:assetIds;type:text[]" json:"assetIds"`
	CheckinCount     int                `gorm:"column:checkinCount;size:32;not null" json:"checkinCount"`
	CheckinTimestamp DateTime           `gorm:"column:checkinTimestamp;index:PartialBookingCheckin_checkinTimestamp_idx;index:PartialBookingCheckin_bookingId_checkinTimestamp_idx;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"checkinTimestamp"`
	BookingID        string             `gorm:"column:bookingId;index:PartialBookingCheckin_bookingId_idx;index:PartialBookingCheckin_bookingId_checkinTimestamp_idx;not null" json:"bookingId"`
	CheckedInByID    string             `gorm:"column:checkedInById;index:PartialBookingCheckin_checkedInById_idx;not null" json:"checkedInById"`
	CreatedAt        DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Booking          *Booking           `gorm:"belongsTo"`
	CheckedInBy      *User              `gorm:"foreignKey:CheckedInByID;belongsTo"`
}

func (PartialBookingCheckin) TableName() string {
	return "PartialBookingCheckin"
}

func (m *PartialBookingCheckin) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CheckinTimestamp.IsZero() {
		m.CheckinTimestamp = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PartialBookingCheckin) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A progressive checkout SESSION: one scan batch, covering N assets via the
// positional `assetIds` / `quantities` / `bookingAssetIds` arrays.
//
// Records sessions and the units they claimed — NOT whether an asset is out.
// The all-at-once checkout (`checkoutBooking`, the Check out button) flips
// statuses and writes NO row here, so **absence proves nothing** and must never
// be read as "was not checked out". Reading it that way is what let a single
// later scan mark every button-checked-out asset on a booking as
// "never checked out". `BookingAsset.checkedOutAt` answers that question.
type PartialBookingCheckout struct {
	ID                string             `gorm:"column:id;primaryKey" json:"id"`
	AssetIds          ScalarList[string] `gorm:"column:assetIds;type:text[]" json:"assetIds"`
	Quantities        ScalarList[int]    `gorm:"column:quantities;type:integer[]" json:"quantities"`
	BookingAssetIds   ScalarList[string] `gorm:"column:bookingAssetIds;type:text[]" json:"bookingAssetIds"`
	CheckoutCount     int                `gorm:"column:checkoutCount;size:32;not null" json:"checkoutCount"`
	CheckoutTimestamp DateTime           `gorm:"column:checkoutTimestamp;index:PartialBookingCheckout_checkoutTimestamp_idx;index:PartialBookingCheckout_bookingId_checkoutTimestamp_idx;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"checkoutTimestamp"`
	BookingID         string             `gorm:"column:bookingId;index:PartialBookingCheckout_bookingId_idx;index:PartialBookingCheckout_bookingId_checkoutTimestamp_idx;not null" json:"bookingId"`
	CheckedOutByID    string             `gorm:"column:checkedOutById;index:PartialBookingCheckout_checkedOutById_idx;not null" json:"checkedOutById"`
	CreatedAt         DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Booking           *Booking           `gorm:"belongsTo"`
	CheckedOutBy      *User              `gorm:"foreignKey:CheckedOutByID;belongsTo"`
}

func (PartialBookingCheckout) TableName() string {
	return "PartialBookingCheckout"
}

func (m *PartialBookingCheckout) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CheckoutTimestamp.IsZero() {
		m.CheckoutTimestamp = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PartialBookingCheckout) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Native-app SSO web-delegated login — single-use session-handoff codes.
// @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
type MobileAuthCode struct {
	ID                string    `gorm:"column:id;primaryKey" json:"id"`
	UserID            string    `gorm:"column:userId;index:MobileAuthCode_userId_idx;not null" json:"userId"`
	CodeHash          string    `gorm:"column:codeHash;uniqueIndex:MobileAuthCode_codeHash_key;not null" json:"codeHash"`
	CodeChallenge     *string   `gorm:"column:codeChallenge" json:"codeChallenge"`
	SessionCiphertext *string   `gorm:"column:sessionCiphertext" json:"sessionCiphertext"`
	ConsumedAt        *DateTime `gorm:"column:consumedAt;type:timestamp(3)" json:"consumedAt"`
	ExpiresAt         DateTime  `gorm:"column:expiresAt;index:MobileAuthCode_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt         DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	User              *User     `gorm:"belongsTo"`
}

func (MobileAuthCode) TableName() string {
	return "MobileAuthCode"
}

func (m *MobileAuthCode) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// One Expo push token per installed companion app. A token identifies a
// device, not a person: when someone else signs in on the same phone, the
// row moves to them. Pushes go only to rows where `disabledAt` is null.
// @see apps/webapp/app/modules/push/service.server.ts
type PushToken struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	Token      string    `gorm:"column:token;uniqueIndex:PushToken_token_key;not null" json:"token"`
	Platform   string    `gorm:"column:platform;not null" json:"platform"`
	DeviceName *string   `gorm:"column:deviceName" json:"deviceName"`
	AppVersion *string   `gorm:"column:appVersion" json:"appVersion"`
	LastSeenAt DateTime  `gorm:"column:lastSeenAt;type:timestamp(3);not null" json:"lastSeenAt"`
	DisabledAt *DateTime `gorm:"column:disabledAt;type:timestamp(3)" json:"disabledAt"`
	UserID     string    `gorm:"column:userId;index:PushToken_userId_idx;not null" json:"userId"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User       *User     `gorm:"belongsTo"`
}

func (PushToken) TableName() string {
	return "PushToken"
}

func (m *PushToken) BeforeCreate(tx *gorm.DB) error {
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

func (*PushToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A person's push notification switches. A missing row means every switch
// is on, so no backfill is needed. The three category columns are named
// after `PushCategory` in apps/webapp/app/modules/push/types.ts.
type NotificationPreference struct {
	UserID            string   `gorm:"column:userId;primaryKey" json:"userId"`
	PushEnabled       *bool    `gorm:"column:pushEnabled;default:true;not null" json:"pushEnabled"`
	Bookings          *bool    `gorm:"column:bookings;default:true;not null" json:"bookings"`
	Audits            *bool    `gorm:"column:audits;default:true;not null" json:"audits"`
	StockAndReminders *bool    `gorm:"column:stockAndReminders;default:true;not null" json:"stockAndReminders"`
	UpdatedAt         DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User              *User    `gorm:"belongsTo"`
}

func (NotificationPreference) TableName() string {
	return "NotificationPreference"
}

func (m *NotificationPreference) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*NotificationPreference) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Kit struct {
	ID              string        `gorm:"column:id;primaryKey" json:"id"`
	Name            string        `gorm:"column:name;index:Kit_categoryId_organizationId_name_idx;not null" json:"name"`
	Description     *string       `gorm:"column:description" json:"description"`
	Status          string        `gorm:"column:status;index:Kit_categoryId_organizationId_status_idx;default:'AVAILABLE';not null" json:"status"`
	Image           *string       `gorm:"column:image" json:"image"`
	ImageExpiration *DateTime     `gorm:"column:imageExpiration;type:timestamp(3)" json:"imageExpiration"`
	OrganizationID  string        `gorm:"column:organizationId;index:Kit_organizationId_idx;index:Kit_categoryId_organizationId_idx;index:Kit_categoryId_organizationId_createdAt_idx;index:Kit_categoryId_organizationId_name_idx;index:Kit_categoryId_organizationId_status_idx;not null" json:"organizationId"`
	CreatedByID     string        `gorm:"column:createdById;index:Kit_createdById_idx;not null" json:"createdById"`
	CategoryID      *string       `gorm:"column:categoryId;index:Kit_categoryId_organizationId_idx;index:Kit_categoryId_organizationId_createdAt_idx;index:Kit_categoryId_organizationId_name_idx;index:Kit_categoryId_organizationId_status_idx" json:"categoryId"`
	CreatedAt       DateTime      `gorm:"column:createdAt;index:Kit_categoryId_organizationId_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	LocationID      *string       `gorm:"column:locationId" json:"locationId"`
	Organization    *Organization `gorm:"belongsTo"`
	CreatedBy       *User         `gorm:"foreignKey:CreatedByID;belongsTo"`
	Category        *Category     `gorm:"belongsTo"`
	Location        *Location     `gorm:"belongsTo"`
	AssetKits       []AssetKit    `gorm:"foreignKey:KitID;constraint:AssetKit_kitId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	QrCodes         []Qr          `gorm:"foreignKey:KitID;constraint:Qr_kitId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Barcodes        []Barcode     `gorm:"foreignKey:KitID;constraint:Barcode_kitId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Reports         []ReportFound `gorm:"foreignKey:KitID;constraint:ReportFound_kitId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Scans           []Scan        `gorm:"foreignKey:KitID;constraint:Scan_kitId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Custody         *KitCustody   `gorm:"foreignKey:KitID;constraint:KitCustody_kitId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Kit) TableName() string {
	return "Kit"
}

func (m *Kit) BeforeCreate(tx *gorm.DB) error {
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

func (*Kit) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type KitCustody struct {
	ID               string      `gorm:"column:id;primaryKey" json:"id"`
	CustodianID      string      `gorm:"column:custodianId;index:KitCustody_custodianId_idx;not null" json:"custodianId"`
	KitID            string      `gorm:"column:kitId;uniqueIndex:KitCustody_kitId_key;not null" json:"kitId"`
	CreatedAt        DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime    `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Custodian        *TeamMember `gorm:"foreignKey:CustodianID;belongsTo"`
	Kit              *Kit        `gorm:"belongsTo"`
	InheritedCustody []Custody   `gorm:"foreignKey:KitCustodyID;constraint:Custody_kitCustodyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (KitCustody) TableName() string {
	return "KitCustody"
}

func (m *KitCustody) BeforeCreate(tx *gorm.DB) error {
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

func (*KitCustody) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Pivot linking an `Asset` to the `Kit` it belongs to.
//
// `@@unique([assetId])` enforces "at most one kit per asset" today.
// The structure already allows multi-row-per-asset; if multi-kit
// allocation lands later, drop the `assetId` unique constraint and
// add a type-aware single-row trigger + sum-within-total CONSTRAINT
// TRIGGER (parallel to the Custody / BookingAsset pivots).
type AssetKit struct {
	ID             string          `gorm:"column:id;primaryKey" json:"id"`
	AssetID        string          `gorm:"column:assetId;uniqueIndex:AssetKit_assetId_kitId_key;not null" json:"assetId"`
	KitID          string          `gorm:"column:kitId;uniqueIndex:AssetKit_assetId_kitId_key;index:AssetKit_kitId_idx;not null" json:"kitId"`
	OrganizationID string          `gorm:"column:organizationId;index:AssetKit_organizationId_idx;not null" json:"organizationId"`
	Quantity       *int            `gorm:"column:quantity;size:32;default:1;not null" json:"quantity"`
	CreatedAt      DateTime        `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime        `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Asset          *Asset          `gorm:"belongsTo"`
	Kit            *Kit            `gorm:"belongsTo"`
	Organization   *Organization   `gorm:"belongsTo"`
	AssetLocations []AssetLocation `gorm:"foreignKey:AssetKitID;constraint:AssetLocation_assetKitId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AssetKit) TableName() string {
	return "AssetKit"
}

func (m *AssetKit) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetKit) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AssetLocation struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	AssetID        string        `gorm:"column:assetId;index:AssetLocation_assetId_idx;not null" json:"assetId"`
	LocationID     string        `gorm:"column:locationId;index:AssetLocation_locationId_idx;not null" json:"locationId"`
	OrganizationID string        `gorm:"column:organizationId;index:AssetLocation_organizationId_idx;not null" json:"organizationId"`
	Quantity       *int          `gorm:"column:quantity;size:32;default:1;not null" json:"quantity"`
	AssetKitID     *string       `gorm:"column:assetKitId;index:AssetLocation_assetKitId_idx" json:"assetKitId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Asset          *Asset        `gorm:"belongsTo"`
	Location       *Location     `gorm:"belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
	AssetKit       *AssetKit     `gorm:"belongsTo"`
}

func (AssetLocation) TableName() string {
	return "AssetLocation"
}

func (m *AssetLocation) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetLocation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AssetReminder struct {
	ID                       string        `gorm:"column:id;primaryKey" json:"id"`
	Name                     string        `gorm:"column:name;index:AssetReminder_name_message_idx;not null" json:"name"`
	Message                  string        `gorm:"column:message;index:AssetReminder_name_message_idx;not null" json:"message"`
	AlertDateTime            DateTime      `gorm:"column:alertDateTime;index:AssetReminder_assetId_alertDateTime_idx;index:AssetReminder_organizationId_alertDateTime_assetId_idx;index:AssetReminder_alertDateTime_activeSchedulerReference_idx;type:timestamp(3);not null" json:"alertDateTime"`
	ActiveSchedulerReference *string       `gorm:"column:activeSchedulerReference;index:AssetReminder_alertDateTime_activeSchedulerReference_idx" json:"activeSchedulerReference"`
	RecurrenceUnit           *string       `gorm:"column:recurrenceUnit" json:"recurrenceUnit"`
	RecurrenceInterval       *int          `gorm:"column:recurrenceInterval;size:32" json:"recurrenceInterval"`
	RecurrenceTimezone       *string       `gorm:"column:recurrenceTimezone" json:"recurrenceTimezone"`
	RecurrenceEndsAt         *DateTime     `gorm:"column:recurrenceEndsAt;type:timestamp(3)" json:"recurrenceEndsAt"`
	OrganizationID           string        `gorm:"column:organizationId;index:AssetReminder_organizationId_alertDateTime_assetId_idx;not null" json:"organizationId"`
	AssetID                  string        `gorm:"column:assetId;index:AssetReminder_assetId_alertDateTime_idx;index:AssetReminder_organizationId_alertDateTime_assetId_idx;not null" json:"assetId"`
	CreatedByID              string        `gorm:"column:createdById;index:AssetReminder_createdById_idx;not null" json:"createdById"`
	CreatedAt                DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization             *Organization `gorm:"belongsTo"`
	Asset                    *Asset        `gorm:"belongsTo"`
	CreatedBy                *User         `gorm:"foreignKey:CreatedByID;belongsTo"`
	TeamMembers              []TeamMember  `gorm:"many2many:_AssetReminderToTeamMember;joinForeignKey:A;joinReferences:B"`
}

func (AssetReminder) TableName() string {
	return "AssetReminder"
}

func (m *AssetReminder) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetReminder) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkingHours struct {
	ID             string                 `gorm:"column:id;primaryKey" json:"id"`
	Enabled        bool                   `gorm:"column:enabled;default:false;not null" json:"enabled"`
	WeeklySchedule datatypes.JSON         `gorm:"column:weeklySchedule;default:'{\"0\":{\"isOpen\":false},\"1\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"2\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"3\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"4\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"5\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"6\":{\"isOpen\":false}}';not null" json:"weeklySchedule"`
	OrganizationID string                 `gorm:"column:organizationId;uniqueIndex:WorkingHours_organizationId_key;index:WorkingHours_organizationId_idx;not null" json:"organizationId"`
	CreatedAt      DateTime               `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime               `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization   *Organization          `gorm:"belongsTo"`
	Overrides      []WorkingHoursOverride `gorm:"foreignKey:WorkingHoursID;constraint:WorkingHoursOverride_workingHoursId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (WorkingHours) TableName() string {
	return "WorkingHours"
}

func (m *WorkingHours) BeforeCreate(tx *gorm.DB) error {
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

func (*WorkingHours) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkingHoursOverride struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	Date           Date          `gorm:"column:date;uniqueIndex:WorkingHoursOverride_workingHoursId_date_key;index:WorkingHoursOverride_workingHoursId_date_idx;index:WorkingHoursOverride_date_isOpen_idx;type:date;not null" json:"date"`
	IsOpen         bool          `gorm:"column:isOpen;index:WorkingHoursOverride_date_isOpen_idx;default:false;not null" json:"isOpen"`
	OpenTime       *string       `gorm:"column:openTime" json:"openTime"`
	CloseTime      *string       `gorm:"column:closeTime" json:"closeTime"`
	Reason         *string       `gorm:"column:reason" json:"reason"`
	WorkingHoursID string        `gorm:"column:workingHoursId;uniqueIndex:WorkingHoursOverride_workingHoursId_date_key;index:WorkingHoursOverride_workingHoursId_date_idx;not null" json:"workingHoursId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	WorkingHours   *WorkingHours `gorm:"belongsTo"`
}

func (WorkingHoursOverride) TableName() string {
	return "WorkingHoursOverride"
}

func (m *WorkingHoursOverride) BeforeCreate(tx *gorm.DB) error {
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

func (*WorkingHoursOverride) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Update struct {
	ID          string             `gorm:"column:id;primaryKey" json:"id"`
	Title       string             `gorm:"column:title;not null" json:"title"`
	Content     string             `gorm:"column:content;not null" json:"content"`
	URL         *string            `gorm:"column:url" json:"url"`
	ImageURL    *string            `gorm:"column:imageUrl" json:"imageUrl"`
	PublishDate DateTime           `gorm:"column:publishDate;index:Update_status_publishDate_idx;index:Update_publishDate_idx;type:timestamp(3);not null" json:"publishDate"`
	Status      string             `gorm:"column:status;index:Update_status_publishDate_idx;default:'DRAFT';not null" json:"status"`
	TargetRoles ScalarList[string] `gorm:"column:targetRoles;type:text[]" json:"targetRoles"`
	ClickCount  int                `gorm:"column:clickCount;size:32;default:0;not null" json:"clickCount"`
	ViewCount   int                `gorm:"column:viewCount;size:32;default:0;not null" json:"viewCount"`
	CreatedByID string             `gorm:"column:createdById;index:Update_createdById_idx;not null" json:"createdById"`
	CreatedAt   DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedBy   *User              `gorm:"foreignKey:CreatedByID;belongsTo"`
	UserReads   []UserUpdateRead   `gorm:"foreignKey:UpdateID;constraint:UserUpdateRead_updateId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Update) TableName() string {
	return "Update"
}

func (m *Update) BeforeCreate(tx *gorm.DB) error {
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

func (*Update) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type UserUpdateRead struct {
	ID       string   `gorm:"column:id;primaryKey" json:"id"`
	UserID   string   `gorm:"column:userId;uniqueIndex:UserUpdateRead_userId_updateId_key;index:UserUpdateRead_userId_idx;not null" json:"userId"`
	UpdateID string   `gorm:"column:updateId;uniqueIndex:UserUpdateRead_userId_updateId_key;index:UserUpdateRead_updateId_idx;not null" json:"updateId"`
	ReadAt   DateTime `gorm:"column:readAt;index:UserUpdateRead_readAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"readAt"`
	User     *User    `gorm:"belongsTo"`
	Update   *Update  `gorm:"belongsTo"`
}

func (UserUpdateRead) TableName() string {
	return "UserUpdateRead"
}

func (m *UserUpdateRead) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.ReadAt.IsZero() {
		m.ReadAt = DateTime{Time: now}
	}
	return nil
}

type AuditNote struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	Content        string        `gorm:"column:content;not null" json:"content"`
	Type           string        `gorm:"column:type;default:'COMMENT';not null" json:"type"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID         *string       `gorm:"column:userId;index:AuditNote_userId_idx" json:"userId"`
	AuditSessionID string        `gorm:"column:auditSessionId;index:AuditNote_auditSessionId_idx;not null" json:"auditSessionId"`
	AuditAssetID   *string       `gorm:"column:auditAssetId;index:AuditNote_auditAssetId_idx" json:"auditAssetId"`
	User           *User         `gorm:"belongsTo"`
	AuditSession   *AuditSession `gorm:"belongsTo"`
	AuditAsset     *AuditAsset   `gorm:"belongsTo"`
}

func (AuditNote) TableName() string {
	return "AuditNote"
}

func (m *AuditNote) BeforeCreate(tx *gorm.DB) error {
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

func (*AuditNote) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TeamMemberNote struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	Content        string        `gorm:"column:content;not null" json:"content"`
	Type           string        `gorm:"column:type;default:'COMMENT';not null" json:"type"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID         *string       `gorm:"column:userId;index:TeamMemberNote_userId_idx" json:"userId"`
	TeamMemberID   string        `gorm:"column:teamMemberId;index:TeamMemberNote_teamMemberId_idx;index:TeamMemberNote_teamMemberId_organizationId_idx;not null" json:"teamMemberId"`
	OrganizationID string        `gorm:"column:organizationId;index:TeamMemberNote_organizationId_idx;index:TeamMemberNote_teamMemberId_organizationId_idx;not null" json:"organizationId"`
	User           *User         `gorm:"belongsTo"`
	TeamMember     *TeamMember   `gorm:"belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (TeamMemberNote) TableName() string {
	return "TeamMemberNote"
}

func (m *TeamMemberNote) BeforeCreate(tx *gorm.DB) error {
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

func (*TeamMemberNote) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AuditSession struct {
	ID                       string            `gorm:"column:id;primaryKey" json:"id"`
	Name                     string            `gorm:"column:name;not null" json:"name"`
	Description              *string           `gorm:"column:description" json:"description"`
	TargetID                 *string           `gorm:"column:targetId" json:"targetId"`
	Status                   string            `gorm:"column:status;index:AuditSession_organizationId_status_idx;index:AuditSession_status_createdAt_idx;default:'PENDING';not null" json:"status"`
	ScopeMeta                datatypes.JSON    `gorm:"column:scopeMeta" json:"scopeMeta"`
	ExpectedAssetCount       int               `gorm:"column:expectedAssetCount;size:32;default:0;not null" json:"expectedAssetCount"`
	FoundAssetCount          int               `gorm:"column:foundAssetCount;size:32;default:0;not null" json:"foundAssetCount"`
	MissingAssetCount        int               `gorm:"column:missingAssetCount;size:32;default:0;not null" json:"missingAssetCount"`
	UnexpectedAssetCount     int               `gorm:"column:unexpectedAssetCount;size:32;default:0;not null" json:"unexpectedAssetCount"`
	StartedAt                *DateTime         `gorm:"column:startedAt;type:timestamp(3)" json:"startedAt"`
	DueDate                  *DateTime         `gorm:"column:dueDate;type:timestamp(3)" json:"dueDate"`
	CompletedAt              *DateTime         `gorm:"column:completedAt;type:timestamp(3)" json:"completedAt"`
	CancelledAt              *DateTime         `gorm:"column:cancelledAt;type:timestamp(3)" json:"cancelledAt"`
	ActiveSchedulerReference *string           `gorm:"column:activeSchedulerReference" json:"activeSchedulerReference"`
	CreatedByID              string            `gorm:"column:createdById;index:AuditSession_createdById_idx;not null" json:"createdById"`
	OrganizationID           string            `gorm:"column:organizationId;index:AuditSession_organizationId_status_idx;not null" json:"organizationId"`
	CreatedAt                DateTime          `gorm:"column:createdAt;index:AuditSession_status_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedBy                *User             `gorm:"foreignKey:CreatedByID;belongsTo"`
	Organization             *Organization     `gorm:"belongsTo"`
	Assignments              []AuditAssignment `gorm:"foreignKey:AuditSessionID;constraint:AuditAssignment_auditSessionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Assets                   []AuditAsset      `gorm:"foreignKey:AuditSessionID;constraint:AuditAsset_auditSessionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Scans                    []AuditScan       `gorm:"foreignKey:AuditSessionID;constraint:AuditScan_auditSessionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Notes                    []AuditNote       `gorm:"foreignKey:AuditSessionID;constraint:AuditNote_auditSessionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Images                   []AuditImage      `gorm:"foreignKey:AuditSessionID;constraint:AuditImage_auditSessionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AuditSession) TableName() string {
	return "AuditSession"
}

func (m *AuditSession) BeforeCreate(tx *gorm.DB) error {
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

func (*AuditSession) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One CSV content import. Holds the parsed rows from preview until the import
// completes, so the job imports exactly the file the user previewed.
// A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
// active import per workspace. Prisma cannot express it; it lives in migration
// 20261015120100_asset_import_one_active_per_workspace.
type AssetImport struct {
	ID        string  `gorm:"column:id;primaryKey" json:"id"`
	Status    string  `gorm:"column:status;index:AssetImport_organizationId_status_idx;index:AssetImport_createdById_organizationId_status_idx;default:'DRAFT';not null" json:"status"`
	Phase     string  `gorm:"column:phase;default:'PREPARE';not null" json:"phase"`
	FileName  *string `gorm:"column:fileName" json:"fileName"`
	TotalRows int     `gorm:"column:totalRows;size:32;not null" json:"totalRows"`
	// Rows committed so far. Advanced in the same transaction as each chunk.
	ProcessedRows int `gorm:"column:processedRows;size:32;default:0;not null" json:"processedRows"`
	// Parsed rows (`StoredImportRow[]`). Cleared when the import completes.
	Rows datatypes.JSON `gorm:"column:rows" json:"rows"`
	// Name to id maps written by the PREPARE step (`ResolvedImportEntities`).
	Resolved       datatypes.JSON `gorm:"column:resolved" json:"resolved"`
	CanUseBarcodes bool           `gorm:"column:canUseBarcodes;default:false;not null" json:"canUseBarcodes"`
	// User-facing reason the import failed. Null unless status is FAILED.
	ErrorMessage *string `gorm:"column:errorMessage" json:"errorMessage"`
	// Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500.
	Warnings datatypes.JSON `gorm:"column:warnings" json:"warnings"`
	// Exact number of warnings, including any beyond the stored cap.
	WarningCount int `gorm:"column:warningCount;size:32;default:0;not null" json:"warningCount"`
	// A step holds the import until this time. Null when no step is running.
	LockedUntil    *DateTime     `gorm:"column:lockedUntil;type:timestamp(3)" json:"lockedUntil"`
	ConfirmedAt    *DateTime     `gorm:"column:confirmedAt;type:timestamp(3)" json:"confirmedAt"`
	StartedAt      *DateTime     `gorm:"column:startedAt;type:timestamp(3)" json:"startedAt"`
	CompletedAt    *DateTime     `gorm:"column:completedAt;type:timestamp(3)" json:"completedAt"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	CreatedByID    string        `gorm:"column:createdById;index:AssetImport_createdById_organizationId_status_idx;not null" json:"createdById"`
	OrganizationID string        `gorm:"column:organizationId;index:AssetImport_organizationId_status_idx;index:AssetImport_createdById_organizationId_status_idx;not null" json:"organizationId"`
	CreatedBy      *User         `gorm:"foreignKey:CreatedByID;belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (AssetImport) TableName() string {
	return "AssetImport"
}

func (m *AssetImport) BeforeCreate(tx *gorm.DB) error {
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

func (*AssetImport) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// What Shelf knows about one web link used in a workspace's custom fields,
// keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
// the link preview job (fetch columns) and people (customTitle).
type LinkPreview struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// `LinkKey` from @shelf/links, never a raw field value. At most 2048
	// characters (CHECK constraint in the migration).
	URL       string  `gorm:"column:url;uniqueIndex:LinkPreview_organizationId_url_key;not null" json:"url"`
	PageTitle *string `gorm:"column:pageTitle" json:"pageTitle"`
	// Null until the first fetch finishes.
	Outcome   *string   `gorm:"column:outcome" json:"outcome"`
	Attempts  int       `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	FetchedAt *DateTime `gorm:"column:fetchedAt;type:timestamp(3)" json:"fetchedAt"`
	// When the link may be fetched next; null means never again. A claim moves
	// it a few minutes ahead, which is also the lease on a fetch in flight.
	NextFetchAt *DateTime `gorm:"column:nextFetchAt;type:timestamp(3);default:CURRENT_TIMESTAMP" json:"nextFetchAt"`
	// The name a person gave the link in this workspace; wins over pageTitle.
	CustomTitle    *string       `gorm:"column:customTitle" json:"customTitle"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OrganizationID string        `gorm:"column:organizationId;uniqueIndex:LinkPreview_organizationId_url_key;not null" json:"organizationId"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (LinkPreview) TableName() string {
	return "LinkPreview"
}

func (m *LinkPreview) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.NextFetchAt == nil {
		m.NextFetchAt = &DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*LinkPreview) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AuditAssignment struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	AuditSessionID string        `gorm:"column:auditSessionId;uniqueIndex:AuditAssignment_auditSessionId_userId_key;not null" json:"auditSessionId"`
	UserID         string        `gorm:"column:userId;uniqueIndex:AuditAssignment_auditSessionId_userId_key;index:AuditAssignment_userId_idx;not null" json:"userId"`
	Role           *string       `gorm:"column:role" json:"role"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	AuditSession   *AuditSession `gorm:"belongsTo"`
	User           *User         `gorm:"belongsTo"`
}

func (AuditAssignment) TableName() string {
	return "AuditAssignment"
}

func (m *AuditAssignment) BeforeCreate(tx *gorm.DB) error {
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

func (*AuditAssignment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AuditAsset struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	AuditSessionID string         `gorm:"column:auditSessionId;uniqueIndex:AuditAsset_auditSessionId_assetId_key;not null" json:"auditSessionId"`
	AssetID        string         `gorm:"column:assetId;uniqueIndex:AuditAsset_auditSessionId_assetId_key;index:AuditAsset_assetId_idx;not null" json:"assetId"`
	Expected       *bool          `gorm:"column:expected;default:true;not null" json:"expected"`
	Status         string         `gorm:"column:status;index:AuditAsset_status_idx;default:'PENDING';not null" json:"status"`
	ScannedByID    *string        `gorm:"column:scannedById;index:AuditAsset_scannedById_idx" json:"scannedById"`
	ScannedAt      *DateTime      `gorm:"column:scannedAt;type:timestamp(3)" json:"scannedAt"`
	Metadata       datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	// Units the audit expects to find, snapshotted when the row is created: the
	// units in the audit's scope (placed at its locations, held in its kits, or
	// held by its person; at most `Asset.quantity`), or `Asset.quantity` when
	// none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
	// which count as one.
	//
	// A snapshot, not a live read: stock that moves during the audit must not
	// move the target the auditor is counting against.
	ExpectedQuantity *int `gorm:"column:expectedQuantity;size:32" json:"expectedQuantity"`
	// Units the auditor counted. Null until a quantity-tracked asset is
	// counted, and always null for INDIVIDUAL assets. A FOUND row with a count
	// below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
	// was located), the gap is this number.
	//
	// Both columns are non-negative, enforced by CHECK constraints in raw SQL
	// (migration `20261104120000_add_audit_asset_quantities`).
	FoundQuantity *int          `gorm:"column:foundQuantity;size:32" json:"foundQuantity"`
	CreatedAt     DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	AuditSession  *AuditSession `gorm:"belongsTo"`
	Asset         *Asset        `gorm:"belongsTo"`
	ScannedBy     *User         `gorm:"foreignKey:ScannedByID;belongsTo"`
	Scans         []AuditScan   `gorm:"foreignKey:AuditAssetID;constraint:AuditScan_auditAssetId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Images        []AuditImage  `gorm:"foreignKey:AuditAssetID;constraint:AuditImage_auditAssetId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Notes         []AuditNote   `gorm:"foreignKey:AuditAssetID;constraint:AuditNote_auditAssetId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (AuditAsset) TableName() string {
	return "AuditAsset"
}

func (m *AuditAsset) BeforeCreate(tx *gorm.DB) error {
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

func (*AuditAsset) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AuditScan struct {
	ID             string  `gorm:"column:id;primaryKey" json:"id"`
	AuditSessionID string  `gorm:"column:auditSessionId;index:AuditScan_auditSessionId_scannedAt_idx;not null" json:"auditSessionId"`
	AuditAssetID   *string `gorm:"column:auditAssetId;index:AuditScan_auditAssetId_idx" json:"auditAssetId"`
	AssetID        *string `gorm:"column:assetId;index:AuditScan_assetId_idx" json:"assetId"`
	ScannedByID    *string `gorm:"column:scannedById" json:"scannedById"`
	Code           *string `gorm:"column:code" json:"code"`
	// What was scanned, recorded BY VALUE at scan time.
	//
	// `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset
	// leaves this row pointing at nothing: no name, and no way to tell whether
	// it had been expected. An audit is a historical record, and a record whose
	// entries can silently lose their meaning is not one — so the two
	// descriptive facts are snapshotted here rather than only referenced.
	//
	// Readers prefer the LIVE asset when it still exists (a rename should show
	// the current name) and fall back to these.
	//
	// Null means the snapshot could not be taken: scans whose asset had already
	// been deleted before these columns were backfilled. Nothing can recover
	// those, and they degrade to the scanned `code`, which also survives. Every
	// other row carries a snapshot, so null is never merely "old".
	AssetTitle   *string        `gorm:"column:assetTitle" json:"assetTitle"`
	WasExpected  *bool          `gorm:"column:wasExpected" json:"wasExpected"`
	Metadata     datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	ScannedAt    DateTime       `gorm:"column:scannedAt;index:AuditScan_auditSessionId_scannedAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"scannedAt"`
	CreatedAt    DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	AuditSession *AuditSession  `gorm:"belongsTo"`
	AuditAsset   *AuditAsset    `gorm:"belongsTo"`
	Asset        *Asset         `gorm:"belongsTo"`
	ScannedBy    *User          `gorm:"foreignKey:ScannedByID;belongsTo"`
}

func (AuditScan) TableName() string {
	return "AuditScan"
}

func (m *AuditScan) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.ScannedAt.IsZero() {
		m.ScannedAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type AuditImage struct {
	ID             string  `gorm:"column:id;primaryKey" json:"id"`
	ImageURL       string  `gorm:"column:imageUrl;not null" json:"imageUrl"`
	ThumbnailURL   *string `gorm:"column:thumbnailUrl" json:"thumbnailUrl"`
	Description    *string `gorm:"column:description" json:"description"`
	AuditSessionID string  `gorm:"column:auditSessionId;index:AuditImage_auditSessionId_idx;not null" json:"auditSessionId"`
	AuditAssetID   *string `gorm:"column:auditAssetId;index:AuditImage_auditAssetId_idx" json:"auditAssetId"`
	UploadedByID   *string `gorm:"column:uploadedById;index:AuditImage_uploadedById_idx" json:"uploadedById"`
	OrganizationID string  `gorm:"column:organizationId;index:AuditImage_organizationId_idx;not null" json:"organizationId"`
	// When the photo was taken, from its EXIF data (or the phone app's copy of
	// it). Null when the photo carried no usable date. EXIF can be edited, so
	// this is the best signal available, not proof.
	CapturedAt *DateTime `gorm:"column:capturedAt;type:timestamp(3)" json:"capturedAt"`
	// How the photo reached Shelf. Null only for photos uploaded before this
	// field existed; a phone upload with no or an unknown source is PHONE.
	CaptureSource *string       `gorm:"column:captureSource" json:"captureSource"`
	CreatedAt     DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	AuditSession  *AuditSession `gorm:"belongsTo"`
	AuditAsset    *AuditAsset   `gorm:"belongsTo"`
	UploadedBy    *User         `gorm:"foreignKey:UploadedByID;belongsTo"`
	Organization  *Organization `gorm:"belongsTo"`
}

func (AuditImage) TableName() string {
	return "AuditImage"
}

func (m *AuditImage) BeforeCreate(tx *gorm.DB) error {
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

func (*AuditImage) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RoleChangeLog struct {
	ID           string `gorm:"column:id;primaryKey" json:"id"`
	PreviousRole string `gorm:"column:previousRole;not null" json:"previousRole"`
	NewRole      string `gorm:"column:newRole;not null" json:"newRole"`
	// How the change was made. SSO rows name the member themself as `changedBy`.
	Source         string        `gorm:"column:source;default:'MANUAL';not null" json:"source"`
	CreatedAt      DateTime      `gorm:"column:createdAt;index:RoleChangeLog_organizationId_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UserID         string        `gorm:"column:userId;index:RoleChangeLog_userId_organizationId_idx;not null" json:"userId"`
	ChangedByID    string        `gorm:"column:changedById;not null" json:"changedById"`
	OrganizationID string        `gorm:"column:organizationId;index:RoleChangeLog_userId_organizationId_idx;index:RoleChangeLog_organizationId_createdAt_idx;not null" json:"organizationId"`
	User           *User         `gorm:"foreignKey:UserID;belongsTo"`
	ChangedBy      *User         `gorm:"foreignKey:ChangedByID;belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (RoleChangeLog) TableName() string {
	return "RoleChangeLog"
}

func (m *RoleChangeLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type ActivityEvent struct {
	ID             string         `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string         `gorm:"column:organizationId;index:ActivityEvent_organizationId_occurredAt_idx;index:ActivityEvent_organizationId_action_occurredAt_idx;index:ActivityEvent_organizationId_entityType_entityId_occurredAt_idx;not null" json:"organizationId"`
	OccurredAt     DateTime       `gorm:"column:occurredAt;index:ActivityEvent_organizationId_occurredAt_idx;index:ActivityEvent_organizationId_action_occurredAt_idx;index:ActivityEvent_organizationId_entityType_entityId_occurredAt_idx;index:ActivityEvent_actorUserId_occurredAt_idx;index:ActivityEvent_assetId_occurredAt_idx;index:ActivityEvent_bookingId_occurredAt_idx;index:ActivityEvent_auditSessionId_occurredAt_idx;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"occurredAt"`
	ActorUserID    *string        `gorm:"column:actorUserId;index:ActivityEvent_actorUserId_occurredAt_idx" json:"actorUserId"`
	ActorSnapshot  datatypes.JSON `gorm:"column:actorSnapshot" json:"actorSnapshot"`
	Action         string         `gorm:"column:action;index:ActivityEvent_organizationId_action_occurredAt_idx;not null" json:"action"`
	EntityType     string         `gorm:"column:entityType;index:ActivityEvent_organizationId_entityType_entityId_occurredAt_idx;not null" json:"entityType"`
	EntityID       string         `gorm:"column:entityId;index:ActivityEvent_organizationId_entityType_entityId_occurredAt_idx;not null" json:"entityId"`
	AssetID        *string        `gorm:"column:assetId;index:ActivityEvent_assetId_occurredAt_idx" json:"assetId"`
	BookingID      *string        `gorm:"column:bookingId;index:ActivityEvent_bookingId_occurredAt_idx" json:"bookingId"`
	AuditSessionID *string        `gorm:"column:auditSessionId;index:ActivityEvent_auditSessionId_occurredAt_idx" json:"auditSessionId"`
	AuditAssetID   *string        `gorm:"column:auditAssetId" json:"auditAssetId"`
	KitID          *string        `gorm:"column:kitId" json:"kitId"`
	LocationID     *string        `gorm:"column:locationId" json:"locationId"`
	TeamMemberID   *string        `gorm:"column:teamMemberId" json:"teamMemberId"`
	TargetUserID   *string        `gorm:"column:targetUserId" json:"targetUserId"`
	Field          *string        `gorm:"column:field" json:"field"`
	FromValue      datatypes.JSON `gorm:"column:fromValue" json:"fromValue"`
	ToValue        datatypes.JSON `gorm:"column:toValue" json:"toValue"`
	Meta           datatypes.JSON `gorm:"column:meta" json:"meta"`
	Organization   *Organization  `gorm:"belongsTo"`
}

func (ActivityEvent) TableName() string {
	return "ActivityEvent"
}

func (m *ActivityEvent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.OccurredAt.IsZero() {
		m.OccurredAt = DateTime{Time: now}
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
