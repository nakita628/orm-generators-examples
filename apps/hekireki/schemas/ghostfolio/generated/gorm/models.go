package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var NamingStrategy = schema.NamingStrategy{SingularTable: true, NoLowerCase: true}

type Access struct {
	Alias         *string            `gorm:"column:alias;index:Access_alias_idx" json:"alias"`
	CreatedAt     DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ExpiresAt     DateTime           `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	GranteeUserID *string            `gorm:"column:granteeUserId;index:Access_granteeUserId_idx" json:"granteeUserId"`
	ID            string             `gorm:"column:id;primaryKey" json:"id"`
	LastUsedAt    *DateTime          `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	Scopes        ScalarList[string] `gorm:"column:scopes;type:text[]" json:"scopes"`
	Settings      datatypes.JSON     `gorm:"column:settings;default:'{}';not null" json:"settings"`
	Type          string             `gorm:"column:type;default:'PRIVATE';not null" json:"type"`
	UpdatedAt     DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID        string             `gorm:"column:userId;index:Access_userId_idx;not null" json:"userId"`
	GranteeUser   *User              `gorm:"foreignKey:GranteeUserID;belongsTo"`
	User          *User              `gorm:"foreignKey:UserID;belongsTo"`
}

func (Access) TableName() string {
	return "Access"
}

func (m *Access) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Access) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Account struct {
	Comment    *string          `gorm:"column:comment" json:"comment"`
	CreatedAt  DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Currency   *string          `gorm:"column:currency;index:Account_currency_idx" json:"currency"`
	ID         string           `gorm:"column:id;primaryKey;index:Account_id_idx" json:"id"`
	Name       *string          `gorm:"column:name;index:Account_name_idx" json:"name"`
	PlatformID *string          `gorm:"column:platformId" json:"platformId"`
	UpdatedAt  DateTime         `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID     string           `gorm:"column:userId;primaryKey;index:Account_userId_idx" json:"userId"`
	Platform   *Platform        `gorm:"belongsTo"`
	User       *User            `gorm:"belongsTo"`
	Activities []Order          `gorm:"foreignKey:AccountID,AccountUserID;references:ID,UserID;constraint:Order_accountId_accountUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Balances   []AccountBalance `gorm:"foreignKey:AccountID,UserID;references:ID,UserID;constraint:AccountBalance_accountId_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags       []TagsOnAccounts `gorm:"foreignKey:AccountID,UserID;references:ID,UserID;constraint:TagsOnAccounts_accountId_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Account) TableName() string {
	return "Account"
}

func (m *Account) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
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

type AccountBalance struct {
	AccountID string   `gorm:"column:accountId;uniqueIndex:AccountBalance_accountId_userId_date_key,priority:1;index:AccountBalance_accountId_idx;not null" json:"accountId"`
	Comment   *string  `gorm:"column:comment" json:"comment"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Date      DateTime `gorm:"column:date;uniqueIndex:AccountBalance_accountId_userId_date_key,priority:3;index:AccountBalance_date_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"date"`
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID    string   `gorm:"column:userId;uniqueIndex:AccountBalance_accountId_userId_date_key,priority:2;not null" json:"userId"`
	Value     float64  `gorm:"column:value;type:double precision;not null" json:"value"`
	Account   *Account `gorm:"foreignKey:AccountID,UserID;references:ID,UserID;belongsTo"`
}

func (AccountBalance) TableName() string {
	return "AccountBalance"
}

func (m *AccountBalance) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.Date.IsZero() {
		m.Date = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AccountBalance) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Analytics struct {
	ActivityCount                       int      `gorm:"column:activityCount;size:32;default:0;not null" json:"activityCount"`
	Country                             *string  `gorm:"column:country" json:"country"`
	DataProviderGhostfolioDailyRequests int      `gorm:"column:dataProviderGhostfolioDailyRequests;size:32;default:0;not null" json:"dataProviderGhostfolioDailyRequests"`
	LastRequestAt                       DateTime `gorm:"column:lastRequestAt;index:Analytics_lastRequestAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastRequestAt"`
	LastRequestClient                   *string  `gorm:"column:lastRequestClient" json:"lastRequestClient"`
	UpdatedAt                           DateTime `gorm:"column:updatedAt;index:Analytics_updatedAt_idx;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID                              string   `gorm:"column:userId;primaryKey" json:"userId"`
	User                                *User    `gorm:"belongsTo"`
}

func (Analytics) TableName() string {
	return "Analytics"
}

func (m *Analytics) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.LastRequestAt.IsZero() {
		m.LastRequestAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Analytics) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type APIKey struct {
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	HashedKey string   `gorm:"column:hashedKey;uniqueIndex:ApiKey_hashedKey_key;not null" json:"hashedKey"`
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID    string   `gorm:"column:userId;index:ApiKey_userId_idx;not null" json:"userId"`
	User      *User    `gorm:"belongsTo"`
}

func (APIKey) TableName() string {
	return "ApiKey"
}

func (m *APIKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
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

type AssetProfileOverrides struct {
	AssetClass      *string        `gorm:"column:assetClass" json:"assetClass"`
	AssetSubClass   *string        `gorm:"column:assetSubClass" json:"assetSubClass"`
	Countries       datatypes.JSON `gorm:"column:countries;default:'[]'" json:"countries"`
	Holdings        datatypes.JSON `gorm:"column:holdings;default:'[]'" json:"holdings"`
	Name            *string        `gorm:"column:name" json:"name"`
	Sectors         datatypes.JSON `gorm:"column:sectors;default:'[]'" json:"sectors"`
	SymbolProfileID string         `gorm:"column:symbolProfileId;primaryKey" json:"symbolProfileId"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	URL             *string        `gorm:"column:url" json:"url"`
	SymbolProfile   *SymbolProfile `gorm:"belongsTo"`
}

func (AssetProfileOverrides) TableName() string {
	return "SymbolProfileOverrides"
}

func (m *AssetProfileOverrides) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AssetProfileOverrides) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AssetProfileResolution struct {
	CreatedAt        DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Currency         string   `gorm:"column:currency;not null" json:"currency"`
	DataSourceOrigin string   `gorm:"column:dataSourceOrigin;uniqueIndex:AssetProfileResolution_dataSourceOrigin_symbolOrigin_key,priority:1;not null" json:"dataSourceOrigin"`
	DataSourceTarget string   `gorm:"column:dataSourceTarget;not null" json:"dataSourceTarget"`
	ID               string   `gorm:"column:id;primaryKey" json:"id"`
	RequestCount     *int     `gorm:"column:requestCount;size:32;default:1;not null" json:"requestCount"`
	SymbolOrigin     string   `gorm:"column:symbolOrigin;uniqueIndex:AssetProfileResolution_dataSourceOrigin_symbolOrigin_key,priority:2;not null" json:"symbolOrigin"`
	SymbolTarget     string   `gorm:"column:symbolTarget;not null" json:"symbolTarget"`
	UpdatedAt        DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AssetProfileResolution) TableName() string {
	return "AssetProfileResolution"
}

func (m *AssetProfileResolution) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AssetProfileResolution) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A stock split of an asset profile. The numerator is the number of shares
// held after the split, the denominator the number of shares held before, for
// example 4 and 1 for a 4:1 split or 1 and 10 for a 1:10 reverse split.
//
// The ratio is stored as two integers instead of the resulting factor, so that
// it stays exact, for example 1/3 for a 1:3 reverse split.
//
// Activities before the date are adjusted by the ratio: the quantity is
// multiplied by it and the unit price divided by it, so that the total value
// of an activity stays the same.
//
// Market data is already split-adjusted by the data providers and must not be
// adjusted again.
type AssetProfileSplit struct {
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Date            DateTime       `gorm:"column:date;uniqueIndex:AssetProfileSplit_symbolProfileId_date_key,priority:2;index:AssetProfileSplit_date_idx;type:timestamp(3);not null" json:"date"`
	Denominator     int            `gorm:"column:denominator;size:32;not null" json:"denominator"`
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Numerator       int            `gorm:"column:numerator;size:32;not null" json:"numerator"`
	SymbolProfileID string         `gorm:"column:symbolProfileId;uniqueIndex:AssetProfileSplit_symbolProfileId_date_key,priority:1;not null" json:"symbolProfileId"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	SymbolProfile   *SymbolProfile `gorm:"belongsTo"`
}

func (AssetProfileSplit) TableName() string {
	return "AssetProfileSplit"
}

func (m *AssetProfileSplit) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AssetProfileSplit) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AuthDevice struct {
	CreatedAt           DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	CredentialID        []byte   `gorm:"column:credentialId;not null" json:"credentialId"`
	CredentialPublicKey []byte   `gorm:"column:credentialPublicKey;not null" json:"credentialPublicKey"`
	Counter             int      `gorm:"column:counter;size:32;not null" json:"counter"`
	ID                  string   `gorm:"column:id;primaryKey" json:"id"`
	UpdatedAt           DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID              string   `gorm:"column:userId;index:AuthDevice_userId_idx;not null" json:"userId"`
	User                *User    `gorm:"belongsTo"`
}

func (AuthDevice) TableName() string {
	return "AuthDevice"
}

func (m *AuthDevice) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AuthDevice) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type MarketData struct {
	CreatedAt        DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	DataSource       string   `gorm:"column:dataSource;uniqueIndex:MarketData_dataSource_symbol_date_key,priority:1;not null" json:"dataSource"`
	Date             DateTime `gorm:"column:date;uniqueIndex:MarketData_dataSource_symbol_date_key,priority:3;index:MarketData_date_idx;type:timestamp(3);not null" json:"date"`
	ID               string   `gorm:"column:id;primaryKey" json:"id"`
	IsCarriedForward bool     `gorm:"column:isCarriedForward;default:false;not null" json:"isCarriedForward"`
	MarketPrice      float64  `gorm:"column:marketPrice;type:double precision;not null" json:"marketPrice"`
	State            string   `gorm:"column:state;index:MarketData_state_idx;default:'CLOSE';not null" json:"state"`
	Symbol           string   `gorm:"column:symbol;uniqueIndex:MarketData_dataSource_symbol_date_key,priority:2;index:MarketData_symbol_idx;not null" json:"symbol"`
}

func (MarketData) TableName() string {
	return "MarketData"
}

func (m *MarketData) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type Order struct {
	AccountID       *string        `gorm:"column:accountId;index:Order_accountId_idx" json:"accountId"`
	AccountUserID   *string        `gorm:"column:accountUserId" json:"accountUserId"`
	Comment         *string        `gorm:"column:comment" json:"comment"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Currency        *string        `gorm:"column:currency" json:"currency"`
	Date            DateTime       `gorm:"column:date;index:Order_date_idx;index:Order_symbolProfileId_date_idx,priority:2;index:Order_userId_date_idx,priority:2;index:Order_userId_type_date_idx,priority:3;type:timestamp(3);not null" json:"date"`
	Fee             float64        `gorm:"column:fee;type:double precision;not null" json:"fee"`
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	Quantity        float64        `gorm:"column:quantity;type:double precision;not null" json:"quantity"`
	SymbolProfileID string         `gorm:"column:symbolProfileId;index:Order_symbolProfileId_date_idx,priority:1;not null" json:"symbolProfileId"`
	Type            string         `gorm:"column:type;index:Order_userId_type_date_idx,priority:2;not null" json:"type"`
	UnitPrice       float64        `gorm:"column:unitPrice;type:double precision;not null" json:"unitPrice"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID          string         `gorm:"column:userId;index:Order_userId_date_idx,priority:1;index:Order_userId_type_date_idx,priority:1;not null" json:"userId"`
	Account         *Account       `gorm:"foreignKey:AccountID,AccountUserID;references:ID,UserID;belongsTo"`
	User            *User          `gorm:"belongsTo"`
	SymbolProfile   *SymbolProfile `gorm:"belongsTo"`
	Tags            []Tag          `gorm:"many2many:_OrderToTag;joinForeignKey:A;joinReferences:B"`
}

func (Order) TableName() string {
	return "Order"
}

func (m *Order) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Order) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Platform struct {
	ID       string    `gorm:"column:id;primaryKey" json:"id"`
	Name     *string   `gorm:"column:name;index:Platform_name_idx" json:"name"`
	URL      string    `gorm:"column:url;uniqueIndex:Platform_url_key;not null" json:"url"`
	Accounts []Account `gorm:"foreignKey:PlatformID;constraint:Account_platformId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Platform) TableName() string {
	return "Platform"
}

func (m *Platform) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type Property struct {
	Key   string `gorm:"column:key;primaryKey" json:"key"`
	Value string `gorm:"column:value;not null" json:"value"`
}

func (Property) TableName() string {
	return "Property"
}

type Settings struct {
	Settings  datatypes.JSON `gorm:"column:settings" json:"settings"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID    string         `gorm:"column:userId;primaryKey" json:"userId"`
	User      *User          `gorm:"belongsTo"`
}

func (Settings) TableName() string {
	return "Settings"
}

func (m *Settings) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Settings) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SymbolProfile struct {
	AssetClass             *string                `gorm:"column:assetClass;index:SymbolProfile_assetClass_idx" json:"assetClass"`
	AssetSubClass          *string                `gorm:"column:assetSubClass" json:"assetSubClass"`
	Comment                *string                `gorm:"column:comment" json:"comment"`
	Countries              datatypes.JSON         `gorm:"column:countries" json:"countries"`
	CreatedAt              DateTime               `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Currency               string                 `gorm:"column:currency;index:SymbolProfile_currency_idx;not null" json:"currency"`
	Cusip                  *string                `gorm:"column:cusip;index:SymbolProfile_cusip_idx" json:"cusip"`
	DataGatheringFrequency string                 `gorm:"column:dataGatheringFrequency;index:SymbolProfile_dataGatheringFrequency_idx;default:'DAILY';not null" json:"dataGatheringFrequency"`
	DataSource             string                 `gorm:"column:dataSource;uniqueIndex:SymbolProfile_dataSource_symbol_key,priority:1;index:SymbolProfile_dataSource_idx;not null" json:"dataSource"`
	Figi                   *string                `gorm:"column:figi" json:"figi"`
	FigiComposite          *string                `gorm:"column:figiComposite" json:"figiComposite"`
	FigiShareClass         *string                `gorm:"column:figiShareClass" json:"figiShareClass"`
	Holdings               datatypes.JSON         `gorm:"column:holdings;default:'[]'" json:"holdings"`
	ID                     string                 `gorm:"column:id;primaryKey" json:"id"`
	IsActive               *bool                  `gorm:"column:isActive;index:SymbolProfile_isActive_idx;default:true;not null" json:"isActive"`
	Isin                   *string                `gorm:"column:isin;index:SymbolProfile_isin_idx" json:"isin"`
	Name                   *string                `gorm:"column:name;index:SymbolProfile_name_idx" json:"name"`
	ScraperConfiguration   datatypes.JSON         `gorm:"column:scraperConfiguration" json:"scraperConfiguration"`
	Sectors                datatypes.JSON         `gorm:"column:sectors" json:"sectors"`
	Symbol                 string                 `gorm:"column:symbol;uniqueIndex:SymbolProfile_dataSource_symbol_key,priority:2;index:SymbolProfile_symbol_idx;not null" json:"symbol"`
	SymbolMapping          datatypes.JSON         `gorm:"column:symbolMapping" json:"symbolMapping"`
	UpdatedAt              DateTime               `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	URL                    *string                `gorm:"column:url" json:"url"`
	UserID                 *string                `gorm:"column:userId" json:"userId"`
	User                   *User                  `gorm:"belongsTo"`
	Activities             []Order                `gorm:"foreignKey:SymbolProfileID;constraint:Order_symbolProfileId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AssetProfileSplits     []AssetProfileSplit    `gorm:"foreignKey:SymbolProfileID;constraint:AssetProfileSplit_symbolProfileId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssetProfileOverrides  *AssetProfileOverrides `gorm:"foreignKey:SymbolProfileID;constraint:SymbolProfileOverrides_symbolProfileId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	WatchedBy              []User                 `gorm:"many2many:_UserWatchlist;joinForeignKey:A;joinReferences:B"`
}

func (SymbolProfile) TableName() string {
	return "SymbolProfile"
}

func (m *SymbolProfile) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SymbolProfile) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Subscription struct {
	CreatedAt               DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ExpiresAt               DateTime `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	ID                      string   `gorm:"column:id;primaryKey" json:"id"`
	Price                   *float64 `gorm:"column:price;type:double precision" json:"price"`
	StripeCheckoutSessionID *string  `gorm:"column:stripeCheckoutSessionId;uniqueIndex:Subscription_stripeCheckoutSessionId_key" json:"stripeCheckoutSessionId"`
	UpdatedAt               DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID                  string   `gorm:"column:userId;index:Subscription_userId_idx;not null" json:"userId"`
	User                    *User    `gorm:"belongsTo"`
}

func (Subscription) TableName() string {
	return "Subscription"
}

func (m *Subscription) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
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

type Tag struct {
	ID         string           `gorm:"column:id;primaryKey" json:"id"`
	Name       string           `gorm:"column:name;uniqueIndex:Tag_name_userId_key,priority:1;index:Tag_name_idx;not null" json:"name"`
	UserID     *string          `gorm:"column:userId;uniqueIndex:Tag_name_userId_key,priority:2" json:"userId"`
	User       *User            `gorm:"belongsTo"`
	Accounts   []TagsOnAccounts `gorm:"foreignKey:TagID;constraint:TagsOnAccounts_tagId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Activities []Order          `gorm:"many2many:_OrderToTag;joinForeignKey:B;joinReferences:A"`
}

func (Tag) TableName() string {
	return "Tag"
}

func (m *Tag) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type TagsOnAccounts struct {
	AccountID string   `gorm:"column:accountId;primaryKey" json:"accountId"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	TagID     string   `gorm:"column:tagId;primaryKey;index:TagsOnAccounts_tagId_idx" json:"tagId"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID    string   `gorm:"column:userId;primaryKey" json:"userId"`
	Account   *Account `gorm:"foreignKey:AccountID,UserID;references:ID,UserID;belongsTo"`
	Tag       *Tag     `gorm:"belongsTo"`
}

func (TagsOnAccounts) TableName() string {
	return "TagsOnAccounts"
}

func (m *TagsOnAccounts) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TagsOnAccounts) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type User struct {
	AccessToken   *string         `gorm:"column:accessToken;index:User_accessToken_idx" json:"accessToken"`
	AuthChallenge *string         `gorm:"column:authChallenge" json:"authChallenge"`
	CreatedAt     DateTime        `gorm:"column:createdAt;index:User_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ID            string          `gorm:"column:id;primaryKey" json:"id"`
	Provider      string          `gorm:"column:provider;index:User_provider_idx;default:'ANONYMOUS';not null" json:"provider"`
	Role          string          `gorm:"column:role;index:User_role_idx;default:'USER';not null" json:"role"`
	ThirdPartyID  *string         `gorm:"column:thirdPartyId;index:User_thirdPartyId_idx" json:"thirdPartyId"`
	UpdatedAt     DateTime        `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	AccessesGet   []Access        `gorm:"foreignKey:GranteeUserID;constraint:Access_granteeUserId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessesGive  []Access        `gorm:"foreignKey:UserID;constraint:Access_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Accounts      []Account       `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Activities    []Order         `gorm:"foreignKey:UserID;constraint:Order_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeys       []APIKey        `gorm:"foreignKey:UserID;constraint:ApiKey_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuthDevices   []AuthDevice    `gorm:"foreignKey:UserID;constraint:AuthDevice_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subscriptions []Subscription  `gorm:"foreignKey:UserID;constraint:Subscription_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags          []Tag           `gorm:"foreignKey:UserID;constraint:Tag_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SymbolProfile []SymbolProfile `gorm:"foreignKey:UserID;constraint:SymbolProfile_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Analytics     *Analytics      `gorm:"foreignKey:UserID;constraint:Analytics_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Settings      *Settings       `gorm:"foreignKey:UserID;constraint:Settings_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Watchlist     []SymbolProfile `gorm:"many2many:_UserWatchlist;joinForeignKey:B;joinReferences:A"`
}

func (User) TableName() string {
	return "User"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.ID == "" {
		m.ID = uuid.NewString()
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
