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
	ID                string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	UserID            string   `gorm:"column:userId;type:varchar(191);not null" json:"userId"`
	Type              string   `gorm:"column:type;type:varchar(191);not null" json:"type"`
	Provider          string   `gorm:"column:provider;uniqueIndex:accounts_provider_providerAccountId_key,priority:1;type:varchar(191);not null" json:"provider"`
	ProviderAccountID string   `gorm:"column:providerAccountId;uniqueIndex:accounts_provider_providerAccountId_key,priority:2;type:varchar(191);not null" json:"providerAccountId"`
	RefreshToken      *string  `gorm:"column:refresh_token;type:text" json:"refresh_token"`
	AccessToken       *string  `gorm:"column:access_token;type:text" json:"access_token"`
	ExpiresAt         *int     `gorm:"column:expires_at;size:32" json:"expires_at"`
	TokenType         *string  `gorm:"column:token_type;type:varchar(191)" json:"token_type"`
	Scope             *string  `gorm:"column:scope;type:varchar(191)" json:"scope"`
	IDToken           *string  `gorm:"column:id_token;type:text" json:"id_token"`
	SessionState      *string  `gorm:"column:session_state;type:varchar(191)" json:"session_state"`
	CreatedAt         DateTime `gorm:"column:created_at;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt         DateTime `gorm:"column:updated_at;autoUpdateTime:false;not null" json:"updated_at"`
	User              *User    `gorm:"belongsTo"`
}

func (Account) TableName() string {
	return "accounts"
}

func (m *Account) BeforeCreate(tx *gorm.DB) error {
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

type Session struct {
	ID           string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	SessionToken string   `gorm:"column:sessionToken;uniqueIndex:sessions_sessionToken_key;type:varchar(191);not null" json:"sessionToken"`
	UserID       string   `gorm:"column:userId;type:varchar(191);not null" json:"userId"`
	Expires      DateTime `gorm:"column:expires;not null" json:"expires"`
	User         *User    `gorm:"belongsTo"`
}

func (Session) TableName() string {
	return "sessions"
}

func (m *Session) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type User struct {
	ID                     string    `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name                   *string   `gorm:"column:name;type:varchar(191)" json:"name"`
	Email                  *string   `gorm:"column:email;uniqueIndex:users_email_key;type:varchar(191)" json:"email"`
	EmailVerified          *DateTime `gorm:"column:emailVerified" json:"emailVerified"`
	Image                  *string   `gorm:"column:image;type:varchar(191)" json:"image"`
	CreatedAt              DateTime  `gorm:"column:created_at;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt              DateTime  `gorm:"column:updated_at;autoUpdateTime:false;not null" json:"updated_at"`
	StripeCustomerID       *string   `gorm:"column:stripe_customer_id;uniqueIndex:users_stripe_customer_id_key;type:varchar(191)" json:"stripe_customer_id"`
	StripeSubscriptionID   *string   `gorm:"column:stripe_subscription_id;uniqueIndex:users_stripe_subscription_id_key;type:varchar(191)" json:"stripe_subscription_id"`
	StripePriceID          *string   `gorm:"column:stripe_price_id;type:varchar(191)" json:"stripe_price_id"`
	StripeCurrentPeriodEnd *DateTime `gorm:"column:stripe_current_period_end" json:"stripe_current_period_end"`
	Accounts               []Account `gorm:"foreignKey:UserID;constraint:accounts_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions               []Session `gorm:"foreignKey:UserID;constraint:sessions_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Post                   []Post    `gorm:"foreignKey:AuthorID;constraint:posts_authorId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (User) TableName() string {
	return "users"
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
	return nil
}

type VerificationToken struct {
	Identifier string   `gorm:"column:identifier;uniqueIndex:verification_tokens_identifier_token_key,priority:1;type:varchar(191);not null" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:verification_tokens_token_key;uniqueIndex:verification_tokens_identifier_token_key,priority:2;type:varchar(191)" json:"token"`
	Expires    DateTime `gorm:"column:expires;not null" json:"expires"`
}

func (VerificationToken) TableName() string {
	return "verification_tokens"
}

type Post struct {
	ID        string         `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Title     string         `gorm:"column:title;type:varchar(191);not null" json:"title"`
	Content   datatypes.JSON `gorm:"column:content" json:"content"`
	Published bool           `gorm:"column:published;default:false;not null" json:"published"`
	CreatedAt DateTime       `gorm:"column:created_at;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime       `gorm:"column:updated_at;autoUpdateTime:false;not null" json:"updated_at"`
	AuthorID  string         `gorm:"column:authorId;type:varchar(191);not null" json:"authorId"`
	Author    *User          `gorm:"foreignKey:AuthorID;belongsTo"`
}

func (Post) TableName() string {
	return "posts"
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
