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

type APIUsage struct {
	SpaceID           string    `gorm:"column:space_id;primaryKey" json:"space_id"`
	Resource          string    `gorm:"column:resource;primaryKey" json:"resource"`
	Period            Date      `gorm:"column:period;primaryKey;type:date" json:"period"`
	Count             int       `gorm:"column:count;size:32;default:0;not null" json:"count"`
	NearLimitWarnedAt *DateTime `gorm:"column:near_limit_warned_at;type:timestamp(3)" json:"near_limit_warned_at"`
	AtLimitWarnedAt   *DateTime `gorm:"column:at_limit_warned_at;type:timestamp(3)" json:"at_limit_warned_at"`
	UpdatedAt         DateTime  `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Space             *Space    `gorm:"belongsTo"`
}

func (APIUsage) TableName() string {
	return "api_usage"
}

func (m *APIUsage) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*APIUsage) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Subscription struct {
	ID                   string    `gorm:"column:id;primaryKey" json:"id"`
	PriceID              string    `gorm:"column:price_id;not null" json:"price_id"`
	Quantity             *int      `gorm:"column:quantity;size:32;default:1;not null" json:"quantity"`
	SubscriptionItemID   string    `gorm:"column:subscription_item_id;not null" json:"subscription_item_id"`
	Amount               int       `gorm:"column:amount;size:32;not null" json:"amount"`
	DiscountPercentOff   *float64  `gorm:"column:discount_percent_off;type:double precision" json:"discount_percent_off"`
	DiscountAmountOff    *int      `gorm:"column:discount_amount_off;size:32" json:"discount_amount_off"`
	Status               string    `gorm:"column:status;not null" json:"status"`
	Active               bool      `gorm:"column:active;not null" json:"active"`
	Currency             string    `gorm:"column:currency;not null" json:"currency"`
	Interval             string    `gorm:"column:interval;not null" json:"interval"`
	CreatedAt            DateTime  `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	PeriodStart          DateTime  `gorm:"column:period_start;type:timestamp(3);not null" json:"period_start"`
	PeriodEnd            DateTime  `gorm:"column:period_end;type:timestamp(3);not null" json:"period_end"`
	CancelAtPeriodEnd    bool      `gorm:"column:cancel_at_period_end;default:false;not null" json:"cancel_at_period_end"`
	UserID               string    `gorm:"column:user_id;index:subscriptions_user_id_idx;not null" json:"user_id"`
	SpaceID              string    `gorm:"column:space_id;index:subscriptions_space_id_idx;not null" json:"space_id"`
	APIAddonItemID       *string   `gorm:"column:api_addon_item_id" json:"api_addon_item_id"`
	APIAddonPriceID      *string   `gorm:"column:api_addon_price_id" json:"api_addon_price_id"`
	APIAddonAmount       *int      `gorm:"column:api_addon_amount;size:32" json:"api_addon_amount"`
	APIAddonTier         *string   `gorm:"column:api_addon_tier" json:"api_addon_tier"`
	APIPollRetainedTier  *string   `gorm:"column:api_poll_retained_tier" json:"api_poll_retained_tier"`
	APIPollRetainedUntil *DateTime `gorm:"column:api_poll_retained_until;type:timestamp(3)" json:"api_poll_retained_until"`
	User                 *User     `gorm:"belongsTo"`
	Space                *Space    `gorm:"belongsTo"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}

func (m *Subscription) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type PaymentMethod struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	UserID    string         `gorm:"column:user_id;not null" json:"user_id"`
	Type      string         `gorm:"column:type;not null" json:"type"`
	Data      datatypes.JSON `gorm:"column:data;not null" json:"data"`
	CreatedAt DateTime       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime       `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	User      *User          `gorm:"belongsTo"`
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}

func (m *PaymentMethod) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*PaymentMethod) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type NonprofitApplication struct {
	ID               string   `gorm:"column:id;primaryKey" json:"id"`
	SpaceID          string   `gorm:"column:space_id;index:nonprofit_applications_space_id_created_at_idx;not null" json:"space_id"`
	UserID           string   `gorm:"column:user_id;not null" json:"user_id"`
	OrganizationName string   `gorm:"column:organization_name;not null" json:"organization_name"`
	Website          string   `gorm:"column:website;not null" json:"website"`
	EmailDomain      string   `gorm:"column:email_domain;not null" json:"email_domain"`
	Status           string   `gorm:"column:status;not null" json:"status"`
	Reason           *string  `gorm:"column:reason" json:"reason"`
	Model            *string  `gorm:"column:model" json:"model"`
	CreatedAt        DateTime `gorm:"column:created_at;index:nonprofit_applications_space_id_created_at_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	Space            *Space   `gorm:"belongsTo"`
	User             *User    `gorm:"belongsTo"`
}

func (NonprofitApplication) TableName() string {
	return "nonprofit_applications"
}

func (m *NonprofitApplication) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// An email waiting to be sent by the queue, one row per recipient. The
// email is built when it is sent, from the row its subject points at.
type QueuedEmail struct {
	ID string `gorm:"column:id;primaryKey" json:"id"`
	// Which handler builds the email. A string rather than an enum so a new
	// kind needs no migration; the claim only takes kinds the build knows.
	Kind string `gorm:"column:kind;not null" json:"kind"`
	// The row the email is about, interpreted by the kind. Not a foreign key:
	// kinds point at different tables, and a subject deleted before the send
	// makes its handler skip the email.
	SubjectID string `gorm:"column:subject_id;not null" json:"subject_id"`
	// Groups the emails one action queued, so the action can send its own
	// first batch at once and its progress can be read back.
	BatchID string `gorm:"column:batch_id;index:queued_emails_batch_id_idx;not null" json:"batch_id"`
	// The account whose action queued the email. Banning it skips whatever
	// is still pending.
	UserID string `gorm:"column:user_id;index:queued_emails_user_id_status_idx;not null" json:"user_id"`
	Status string `gorm:"column:status;index:queued_emails_status_created_at_idx;index:queued_emails_status_updated_at_idx;index:queued_emails_user_id_status_idx;default:'pending';not null" json:"status"`
	// Not claimed before this time. Unset sends at the next run.
	SendAfter *DateTime `gorm:"column:send_after;type:timestamp(3)" json:"send_after"`
	Attempts  int       `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	ClaimedAt *DateTime `gorm:"column:claimed_at;type:timestamp(3)" json:"claimed_at"`
	SentAt    *DateTime `gorm:"column:sent_at;type:timestamp(3)" json:"sent_at"`
	LastError *string   `gorm:"column:last_error" json:"last_error"`
	CreatedAt DateTime  `gorm:"column:created_at;index:queued_emails_status_created_at_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime  `gorm:"column:updated_at;index:queued_emails_status_updated_at_idx;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	User      *User     `gorm:"belongsTo"`
}

func (QueuedEmail) TableName() string {
	return "queued_emails"
}

func (m *QueuedEmail) BeforeCreate(tx *gorm.DB) error {
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

func (*QueuedEmail) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EventType struct {
	ID              string           `gorm:"column:id;primaryKey" json:"id"`
	SpaceID         string           `gorm:"column:space_id;index:event_types_space_id_idx;not null" json:"space_id"`
	HostID          string           `gorm:"column:host_id;index:event_types_host_id_idx;not null" json:"host_id"`
	Name            string           `gorm:"column:name;not null" json:"name"`
	Duration        int              `gorm:"column:duration_minutes;size:32;not null" json:"duration_minutes"`
	Capacity        *int             `gorm:"column:capacity;size:32" json:"capacity"`
	Description     *string          `gorm:"column:description" json:"description"`
	Location        datatypes.JSON   `gorm:"column:location" json:"location"`
	CreatedAt       DateTime         `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt       DateTime         `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Deleted         bool             `gorm:"column:deleted;default:false;not null" json:"deleted"`
	DeletedAt       *DateTime        `gorm:"column:deleted_at;type:timestamp(3)" json:"deleted_at"`
	Space           *Space           `gorm:"belongsTo"`
	Host            *User            `gorm:"foreignKey:HostID;belongsTo"`
	Slots           []SheetSlot      `gorm:"foreignKey:EventTypeID;constraint:sheet_slots_event_type_id_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ScheduledEvents []ScheduledEvent `gorm:"foreignKey:EventTypeID;constraint:scheduled_events_event_type_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (EventType) TableName() string {
	return "event_types"
}

func (m *EventType) BeforeCreate(tx *gorm.DB) error {
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

func (*EventType) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ScheduledEvent struct {
	ID           string         `gorm:"column:id;primaryKey" json:"id"`
	UserID       string         `gorm:"column:user_id;index:scheduled_events_user_id_idx;not null" json:"user_id"`
	SpaceID      string         `gorm:"column:space_id;index:scheduled_events_space_id_idx;index:scheduled_events_space_id_status_start_idx;not null" json:"space_id"`
	EventTypeID  *string        `gorm:"column:event_type_id;index:scheduled_events_event_type_id_idx" json:"event_type_id"`
	SheetSlotID  *string        `gorm:"column:sheet_slot_id;uniqueIndex:scheduled_events_sheet_slot_id_key" json:"sheet_slot_id"`
	Title        string         `gorm:"column:title;not null" json:"title"`
	Description  *string        `gorm:"column:description" json:"description"`
	Location     datatypes.JSON `gorm:"column:location" json:"location"`
	Conferencing datatypes.JSON `gorm:"column:conferencing" json:"conferencing"`
	Capacity     *int           `gorm:"column:capacity;size:32" json:"capacity"`
	CreatedAt    DateTime       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt    DateTime       `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Status       string         `gorm:"column:status;index:scheduled_events_space_id_status_start_idx;default:'confirmed';not null" json:"status"`
	// Venue / anchor timezone for in-person or hybrid events. Null means no fixed venue (e.g. remote-only).
	TimeZone         *string                `gorm:"column:time_zone" json:"time_zone"`
	Start            DateTime               `gorm:"column:start;index:scheduled_events_space_id_status_start_idx;type:timestamp(3);not null" json:"start"`
	End              DateTime               `gorm:"column:end;type:timestamp(3);not null" json:"end"`
	AllDay           bool                   `gorm:"column:all_day;default:false;not null" json:"all_day"`
	DeletedAt        *DateTime              `gorm:"column:deleted_at;type:timestamp(3)" json:"deleted_at"`
	Sequence         int                    `gorm:"column:sequence;size:32;default:0;not null" json:"sequence"`
	UID              string                 `gorm:"column:uid;uniqueIndex:scheduled_events_uid_key;not null" json:"uid"`
	HideAttendees    bool                   `gorm:"column:hide_attendees;default:false;not null" json:"hide_attendees"`
	User             *User                  `gorm:"belongsTo"`
	Space            *Space                 `gorm:"belongsTo"`
	EventType        *EventType             `gorm:"belongsTo"`
	SheetSlot        *SheetSlot             `gorm:"belongsTo"`
	RescheduledDates []RescheduledEventDate `gorm:"foreignKey:ScheduledEventID;constraint:rescheduled_event_dates_scheduled_event_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites          []ScheduledEventInvite `gorm:"foreignKey:ScheduledEventID;constraint:scheduled_event_invites_scheduled_event_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Polls            []Poll                 `gorm:"foreignKey:ScheduledEventID;constraint:polls_scheduled_event_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (ScheduledEvent) TableName() string {
	return "scheduled_events"
}

func (m *ScheduledEvent) BeforeCreate(tx *gorm.DB) error {
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

func (*ScheduledEvent) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RescheduledEventDate struct {
	ID               string          `gorm:"column:id;primaryKey" json:"id"`
	ScheduledEventID string          `gorm:"column:scheduled_event_id;index:rescheduled_event_dates_scheduled_event_id_idx;not null" json:"scheduled_event_id"`
	Start            DateTime        `gorm:"column:start;type:timestamp(3);not null" json:"start"`
	End              DateTime        `gorm:"column:end;type:timestamp(3);not null" json:"end"`
	AllDay           bool            `gorm:"column:all_day;default:false;not null" json:"all_day"`
	CreatedAt        DateTime        `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	ScheduledEvent   *ScheduledEvent `gorm:"belongsTo"`
}

func (RescheduledEventDate) TableName() string {
	return "rescheduled_event_dates"
}

func (m *RescheduledEventDate) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type ScheduledEventInvite struct {
	ID               string          `gorm:"column:id;primaryKey" json:"id"`
	UID              string          `gorm:"column:uid;uniqueIndex:scheduled_event_invites_uid_key;not null" json:"uid"`
	ScheduledEventID string          `gorm:"column:scheduled_event_id;uniqueIndex:scheduled_event_invites_scheduled_event_id_invitee_email_key;not null" json:"scheduled_event_id"`
	InviteeName      string          `gorm:"column:invitee_name;not null" json:"invitee_name"`
	InviteeEmail     string          `gorm:"column:invitee_email;uniqueIndex:scheduled_event_invites_scheduled_event_id_invitee_email_key;index:scheduled_event_invites_invitee_email_idx;not null" json:"invitee_email"`
	InviteeID        *string         `gorm:"column:invitee_id;index:scheduled_event_invites_invitee_id_idx" json:"invitee_id"`
	InviteeTimeZone  *string         `gorm:"column:invitee_time_zone" json:"invitee_time_zone"`
	InviteeLocale    *string         `gorm:"column:invitee_locale" json:"invitee_locale"`
	Status           string          `gorm:"column:status;default:'pending';not null" json:"status"`
	CreatedAt        DateTime        `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt        DateTime        `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	ScheduledEvent   *ScheduledEvent `gorm:"belongsTo"`
	User             *User           `gorm:"foreignKey:InviteeID;belongsTo"`
}

func (ScheduledEventInvite) TableName() string {
	return "scheduled_event_invites"
}

func (m *ScheduledEventInvite) BeforeCreate(tx *gorm.DB) error {
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

func (*ScheduledEventInvite) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type InstanceSettings struct {
	ID                      int            `gorm:"column:id;primaryKey;size:32" json:"id"`
	DisableUserRegistration bool           `gorm:"column:disable_user_registration;default:false;not null" json:"disable_user_registration"`
	InstanceID              string         `gorm:"column:instance_id;uniqueIndex:instance_settings_instance_id_key;type:char(36);default:gen_random_uuid();not null" json:"instance_id"`
	AppName                 *string        `gorm:"column:app_name" json:"app_name"`
	PrimaryColor            *string        `gorm:"column:primary_color" json:"primary_color"`
	PrimaryColorDark        *string        `gorm:"column:primary_color_dark" json:"primary_color_dark"`
	Logo                    *string        `gorm:"column:logo" json:"logo"`
	LogoDark                *string        `gorm:"column:logo_dark" json:"logo_dark"`
	LogoIcon                *string        `gorm:"column:logo_icon" json:"logo_icon"`
	HideAttribution         *bool          `gorm:"column:hide_attribution" json:"hide_attribution"`
	FooterLinks             datatypes.JSON `gorm:"column:footer_links;default:'[]';not null" json:"footer_links"`
	CreatedAt               DateTime       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt               DateTime       `gorm:"column:updated_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
}

func (InstanceSettings) TableName() string {
	return "instance_settings"
}

func (m *InstanceSettings) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*InstanceSettings) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Credential struct {
	ID                      string                   `gorm:"column:id;primaryKey" json:"id"`
	UserID                  string                   `gorm:"column:user_id;uniqueIndex:credentials_user_id_provider_provider_account_id_key;not null" json:"user_id"`
	Provider                string                   `gorm:"column:provider;uniqueIndex:credentials_user_id_provider_provider_account_id_key;not null" json:"provider"`
	ProviderAccountID       string                   `gorm:"column:provider_account_id;uniqueIndex:credentials_user_id_provider_provider_account_id_key;not null" json:"provider_account_id"`
	Type                    string                   `gorm:"column:type;not null" json:"type"`
	Secret                  string                   `gorm:"column:secret;not null" json:"secret"`
	Scopes                  ScalarList[string]       `gorm:"column:scopes;type:text[]" json:"scopes"`
	ExpiresAt               *DateTime                `gorm:"column:expires_at;index:credential_expiry_idx;type:timestamp(3)" json:"expires_at"`
	CreatedAt               DateTime                 `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt               DateTime                 `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	User                    *User                    `gorm:"belongsTo"`
	CalendarConnections     []CalendarConnection     `gorm:"foreignKey:CredentialID;constraint:calendar_connections_credential_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConferencingConnections []ConferencingConnection `gorm:"foreignKey:CredentialID;constraint:conferencing_connections_credential_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Credential) TableName() string {
	return "credentials"
}

func (m *Credential) BeforeCreate(tx *gorm.DB) error {
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

func (*Credential) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CalendarConnection struct {
	ID                string             `gorm:"column:id;primaryKey" json:"id"`
	UserID            string             `gorm:"column:user_id;uniqueIndex:calendar_connections_user_id_provider_provider_account_id_key;not null" json:"user_id"`
	Provider          string             `gorm:"column:provider;uniqueIndex:calendar_connections_user_id_provider_provider_account_id_key;not null" json:"provider"`
	IntegrationID     string             `gorm:"column:integration_id;not null" json:"integration_id"`
	ProviderAccountID string             `gorm:"column:provider_account_id;uniqueIndex:calendar_connections_user_id_provider_provider_account_id_key;not null" json:"provider_account_id"`
	Email             string             `gorm:"column:email;not null" json:"email"`
	DisplayName       *string            `gorm:"column:display_name" json:"display_name"`
	CreatedAt         DateTime           `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt         DateTime           `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	CredentialID      string             `gorm:"column:credential_id;not null" json:"credential_id"`
	Credential        *Credential        `gorm:"belongsTo"`
	User              *User              `gorm:"belongsTo"`
	ProviderCalendars []ProviderCalendar `gorm:"foreignKey:CalendarConnectionID;constraint:provider_calendars_calendar_connection_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (CalendarConnection) TableName() string {
	return "calendar_connections"
}

func (m *CalendarConnection) BeforeCreate(tx *gorm.DB) error {
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

func (*CalendarConnection) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ProviderCalendar struct {
	ID                     string              `gorm:"column:id;primaryKey" json:"id"`
	CalendarConnectionID   string              `gorm:"column:calendar_connection_id;uniqueIndex:provider_calendars_calendar_connection_id_provider_calendar_key;index:connection_selected_idx;not null" json:"calendar_connection_id"`
	ProviderCalendarID     string              `gorm:"column:provider_calendar_id;uniqueIndex:provider_calendars_calendar_connection_id_provider_calendar_key;not null" json:"provider_calendar_id"`
	Name                   string              `gorm:"column:name;not null" json:"name"`
	IsPrimary              bool                `gorm:"column:primary;index:primary_calendar_idx;default:false;not null" json:"primary"`
	TimeZone               *string             `gorm:"column:time_zone" json:"time_zone"`
	IsSelected             bool                `gorm:"column:selected;index:connection_selected_idx;default:false;not null" json:"selected"`
	IsDeleted              bool                `gorm:"column:deleted;default:false;not null" json:"deleted"`
	IsWritable             bool                `gorm:"column:writable;default:false;not null" json:"writable"`
	ProviderData           datatypes.JSON      `gorm:"column:provider_data" json:"provider_data"`
	LastSyncedAt           DateTime            `gorm:"column:last_synced_at;index:sync_time_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"last_synced_at"`
	CreatedAt              DateTime            `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt              DateTime            `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	CalendarConnection     *CalendarConnection `gorm:"belongsTo"`
	UserDefaultDestination []User              `gorm:"foreignKey:DefaultDestinationCalendarID;constraint:users_default_destination_calendar_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (ProviderCalendar) TableName() string {
	return "provider_calendars"
}

func (m *ProviderCalendar) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.LastSyncedAt.IsZero() {
		m.LastSyncedAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ProviderCalendar) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ConferencingConnection struct {
	ID                string      `gorm:"column:id;primaryKey" json:"id"`
	UserID            string      `gorm:"column:user_id;uniqueIndex:conferencing_connections_user_id_provider_provider_account__key;not null" json:"user_id"`
	Provider          string      `gorm:"column:provider;uniqueIndex:conferencing_connections_user_id_provider_provider_account__key;not null" json:"provider"`
	IntegrationID     string      `gorm:"column:integration_id;not null" json:"integration_id"`
	ProviderAccountID string      `gorm:"column:provider_account_id;uniqueIndex:conferencing_connections_user_id_provider_provider_account__key;not null" json:"provider_account_id"`
	Email             string      `gorm:"column:email;not null" json:"email"`
	DisplayName       *string     `gorm:"column:display_name" json:"display_name"`
	CreatedAt         DateTime    `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt         DateTime    `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	CredentialID      string      `gorm:"column:credential_id;not null" json:"credential_id"`
	Credential        *Credential `gorm:"belongsTo"`
	User              *User       `gorm:"belongsTo"`
}

func (ConferencingConnection) TableName() string {
	return "conferencing_connections"
}

func (m *ConferencingConnection) BeforeCreate(tx *gorm.DB) error {
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

func (*ConferencingConnection) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type License struct {
	ID              string              `gorm:"column:id;primaryKey" json:"id"`
	LicenseKey      string              `gorm:"column:license_key;uniqueIndex:licenses_license_key_key;not null" json:"license_key"`
	IdempotencyKey  *string             `gorm:"column:idempotency_key;uniqueIndex:licenses_idempotency_key_key" json:"idempotency_key"`
	Version         *int                `gorm:"column:version;size:32" json:"version"`
	Type            string              `gorm:"column:type;not null" json:"type"`
	Seats           *int                `gorm:"column:seats;size:32" json:"seats"`
	IssuedAt        DateTime            `gorm:"column:issued_at;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"issued_at"`
	ExpiresAt       *DateTime           `gorm:"column:expires_at;type:timestamp(3)" json:"expires_at"`
	LicenseeEmail   *string             `gorm:"column:licensee_email" json:"licensee_email"`
	LicenseeName    *string             `gorm:"column:licensee_name" json:"licensee_name"`
	WhiteLabelAddon bool                `gorm:"column:white_label_addon;default:false;not null" json:"white_label_addon"`
	Status          string              `gorm:"column:status;default:'ACTIVE';not null" json:"status"`
	Validations     []LicenseValidation `gorm:"foreignKey:LicenseID;constraint:license_validations_license_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (License) TableName() string {
	return "licenses"
}

func (m *License) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.IssuedAt.IsZero() {
		m.IssuedAt = DateTime{Time: now}
	}
	return nil
}

type LicenseValidation struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	LicenseID   string   `gorm:"column:license_id;not null" json:"license_id"`
	IPAddress   *string  `gorm:"column:ip_address" json:"ip_address"`
	Fingerprint *string  `gorm:"column:fingerprint" json:"fingerprint"`
	ValidatedAt DateTime `gorm:"column:validated_at;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"validated_at"`
	UserAgent   *string  `gorm:"column:user_agent" json:"user_agent"`
	License     *License `gorm:"belongsTo"`
}

func (LicenseValidation) TableName() string {
	return "license_validations"
}

func (m *LicenseValidation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.ValidatedAt.IsZero() {
		m.ValidatedAt = DateTime{Time: now}
	}
	return nil
}

type InstanceLicense struct {
	ID              string    `gorm:"column:id;primaryKey" json:"id"`
	LicenseKey      string    `gorm:"column:license_key;uniqueIndex:instance_licenses_license_key_key;not null" json:"license_key"`
	Version         *int      `gorm:"column:version;size:32" json:"version"`
	Type            string    `gorm:"column:type;not null" json:"type"`
	Seats           *int      `gorm:"column:seats;size:32" json:"seats"`
	IssuedAt        DateTime  `gorm:"column:issued_at;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"issued_at"`
	ExpiresAt       *DateTime `gorm:"column:expires_at;type:timestamp(3)" json:"expires_at"`
	LicenseeEmail   *string   `gorm:"column:licensee_email" json:"licensee_email"`
	LicenseeName    *string   `gorm:"column:licensee_name" json:"licensee_name"`
	WhiteLabelAddon bool      `gorm:"column:white_label_addon;default:false;not null" json:"white_label_addon"`
	Status          string    `gorm:"column:status;default:'ACTIVE';not null" json:"status"`
}

func (InstanceLicense) TableName() string {
	return "instance_licenses"
}

func (m *InstanceLicense) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.IssuedAt.IsZero() {
		m.IssuedAt = DateTime{Time: now}
	}
	return nil
}

type UserNotificationPreferences struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	UserID    string         `gorm:"column:user_id;uniqueIndex:user_notification_preferences_user_id_key;not null" json:"user_id"`
	Prefs     datatypes.JSON `gorm:"column:prefs;default:'{}';not null" json:"prefs"`
	CreatedAt DateTime       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime       `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	User      *User          `gorm:"belongsTo"`
}

func (UserNotificationPreferences) TableName() string {
	return "user_notification_preferences"
}

func (m *UserNotificationPreferences) BeforeCreate(tx *gorm.DB) error {
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

func (*UserNotificationPreferences) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Poll struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime  `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime  `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Deadline    *DateTime `gorm:"column:deadline;type:timestamp(3)" json:"deadline"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	Description *string   `gorm:"column:description" json:"description"`
	Location    *string   `gorm:"column:location" json:"location"`
	// What the organizer asked for: a provider to mint a link with at
	// finalize ("zoom" | "meet" | "teams" | "webex"), or a link they pasted
	// themselves.
	Conferencing            datatypes.JSON `gorm:"column:conferencing" json:"conferencing"`
	UserID                  *string        `gorm:"column:user_id;index:polls_user_id_idx" json:"user_id"`
	TimeZone                *string        `gorm:"column:time_zone" json:"time_zone"`
	Status                  string         `gorm:"column:status;default:'open';not null" json:"status"`
	ClosedReason            *string        `gorm:"column:closed_reason" json:"closed_reason"`
	Kind                    string         `gorm:"column:kind;default:'date';not null" json:"kind"`
	Deleted                 bool           `gorm:"column:deleted;default:false;not null" json:"deleted"`
	DeletedAt               *DateTime      `gorm:"column:deleted_at;type:timestamp(3)" json:"deleted_at"`
	EventID                 *string        `gorm:"column:event_id;uniqueIndex:polls_event_id_key" json:"event_id"`
	ScheduledEventID        *string        `gorm:"column:scheduled_event_id" json:"scheduled_event_id"`
	HideParticipants        bool           `gorm:"column:hide_participants;default:false;not null" json:"hide_participants"`
	HideScores              bool           `gorm:"column:hide_scores;default:false;not null" json:"hide_scores"`
	DisableComments         *bool          `gorm:"column:disable_comments;default:true;not null" json:"disable_comments"`
	RequireParticipantEmail bool           `gorm:"column:require_participant_email;default:false;not null" json:"require_participant_email"`
	// Whether participants may cast the tentative "if need be" vote. Positive
	// flag so a future vote-type expansion adds columns rather than inverting a
	// negative. Turning it off only blocks new tentative votes; existing ones
	// are kept and stay visible.
	AllowTentativeVotes *bool           `gorm:"column:allow_tentative_votes;default:true;not null" json:"allow_tentative_votes"`
	Muted               bool            `gorm:"column:muted;default:false;not null" json:"muted"`
	SpaceID             *string         `gorm:"column:space_id;index:polls_space_id_idx" json:"space_id"`
	User                *User           `gorm:"belongsTo"`
	ScheduledEvent      *ScheduledEvent `gorm:"belongsTo"`
	Space               *Space          `gorm:"belongsTo"`
	Options             []Option        `gorm:"foreignKey:PollID;constraint:options_poll_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Participants        []Participant   `gorm:"foreignKey:PollID;constraint:participants_poll_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments            []Comment       `gorm:"foreignKey:PollID;constraint:comments_poll_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Votes               []Vote          `gorm:"foreignKey:PollID;constraint:votes_poll_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites             []PollInvite    `gorm:"foreignKey:PollID;constraint:poll_invites_poll_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Activities          []PollActivity  `gorm:"foreignKey:PollID;constraint:poll_activities_poll_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Poll) TableName() string {
	return "polls"
}

func (m *Poll) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Poll) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Participant struct {
	ID    string  `gorm:"column:id;primaryKey" json:"id"`
	Name  string  `gorm:"column:name;not null" json:"name"`
	Email *string `gorm:"column:email" json:"email"`
	Note  *string `gorm:"column:note" json:"note"`
	// Per response edit link: grants edit access to this response without a session. App minted at creation, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet — url safe with no linkifier edge cases, unlike base64url); never a DB default so a code path that forgets to mint fails loudly. A response created through an emailed invite takes the invite's token. Rows that predate minting were backfilled with gen_random_uuid() hex
	Token     string      `gorm:"column:token;uniqueIndex:participants_token_key;not null" json:"token"`
	UserID    *string     `gorm:"column:user_id;index:participants_user_id_idx" json:"user_id"`
	PollID    string      `gorm:"column:poll_id;index:participants_poll_id_idx;not null" json:"poll_id"`
	Locale    *string     `gorm:"column:locale" json:"locale"`
	TimeZone  *string     `gorm:"column:time_zone" json:"time_zone"`
	CreatedAt DateTime    `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt *DateTime   `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false" json:"updated_at"`
	Poll      *Poll       `gorm:"belongsTo"`
	User      *User       `gorm:"belongsTo"`
	Votes     []Vote      `gorm:"foreignKey:ParticipantID;constraint:votes_participant_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invite    *PollInvite `gorm:"foreignKey:ParticipantID;constraint:poll_invites_participant_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Participant) TableName() string {
	return "participants"
}

func (m *Participant) BeforeCreate(tx *gorm.DB) error {
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

func (*Participant) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type PollInvite struct {
	ID     string `gorm:"column:id;primaryKey" json:"id"`
	PollID string `gorm:"column:poll_id;uniqueIndex:poll_invites_poll_id_email_key;not null" json:"poll_id"`
	Email  string `gorm:"column:email;uniqueIndex:poll_invites_poll_id_email_key;not null" json:"email"`
	// Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access
	Token      string    `gorm:"column:token;uniqueIndex:poll_invites_token_key;not null" json:"token"`
	OpenedAt   *DateTime `gorm:"column:opened_at;type:timestamp(3)" json:"opened_at"`
	RemindedAt *DateTime `gorm:"column:reminded_at;type:timestamp(3)" json:"reminded_at"`
	// Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token
	RevokedAt *DateTime `gorm:"column:revoked_at;type:timestamp(3)" json:"revoked_at"`
	// Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite
	ParticipantID *string      `gorm:"column:participant_id;uniqueIndex:poll_invites_participant_id_key" json:"participant_id"`
	CreatedAt     DateTime     `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt     DateTime     `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Poll          *Poll        `gorm:"belongsTo"`
	Participant   *Participant `gorm:"belongsTo"`
}

func (PollInvite) TableName() string {
	return "poll_invites"
}

func (m *PollInvite) BeforeCreate(tx *gorm.DB) error {
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

func (*PollInvite) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PollActivity struct {
	ID     string `gorm:"column:id;primaryKey" json:"id"`
	PollID string `gorm:"column:poll_id;index:poll_activities_poll_id_created_at_idx;index:poll_activities_poll_id_participant_id_idx;index:poll_activities_poll_id_option_id_idx;index:poll_activities_poll_id_invite_id_idx;not null" json:"poll_id"`
	// Validated by a zod union in the feature, not a DB enum: the activity vocabulary grows without migrations
	Type string `gorm:"column:type;not null" json:"type"`
	// Subject ids are soft references (no FKs): activity must outlive its subjects — including account deletion — so refs persist as identity and the payload carries the display snapshot; whether a ref still resolves is the renderer's question
	UserID        *string        `gorm:"column:user_id;index:poll_activities_user_id_idx" json:"user_id"`
	ParticipantID *string        `gorm:"column:participant_id;index:poll_activities_poll_id_participant_id_idx" json:"participant_id"`
	InviteID      *string        `gorm:"column:invite_id;index:poll_activities_poll_id_invite_id_idx" json:"invite_id"`
	OptionID      *string        `gorm:"column:option_id;index:poll_activities_poll_id_option_id_idx" json:"option_id"`
	Payload       datatypes.JSON `gorm:"column:payload;default:'{}';not null" json:"payload"`
	CreatedAt     DateTime       `gorm:"column:created_at;index:poll_activities_poll_id_created_at_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	Poll          *Poll          `gorm:"belongsTo"`
}

func (PollActivity) TableName() string {
	return "poll_activities"
}

func (m *PollActivity) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Option struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	StartTime DateTime `gorm:"column:start_time;type:timestamp(0);not null" json:"start_time"`
	Duration  int      `gorm:"column:duration_minutes;size:32;default:0;not null" json:"duration_minutes"`
	PollID    string   `gorm:"column:poll_id;index:options_poll_id_idx;not null" json:"poll_id"`
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	Poll      *Poll    `gorm:"belongsTo"`
	Votes     []Vote   `gorm:"foreignKey:OptionID;constraint:votes_option_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Option) TableName() string {
	return "options"
}

func (m *Option) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Vote struct {
	ID            string       `gorm:"column:id;primaryKey" json:"id"`
	ParticipantID string       `gorm:"column:participant_id;index:votes_participant_id_idx;not null" json:"participant_id"`
	OptionID      string       `gorm:"column:option_id;index:votes_option_id_idx;not null" json:"option_id"`
	PollID        string       `gorm:"column:poll_id;index:votes_poll_id_idx;not null" json:"poll_id"`
	Type          string       `gorm:"column:type;default:'yes';not null" json:"type"`
	CreatedAt     DateTime     `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt     *DateTime    `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false" json:"updated_at"`
	Participant   *Participant `gorm:"belongsTo"`
	Option        *Option      `gorm:"belongsTo"`
	Poll          *Poll        `gorm:"belongsTo"`
}

func (Vote) TableName() string {
	return "votes"
}

func (m *Vote) BeforeCreate(tx *gorm.DB) error {
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

func (*Vote) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type Comment struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	Content    string    `gorm:"column:content;not null" json:"content"`
	PollID     string    `gorm:"column:poll_id;index:comments_poll_id_idx;not null" json:"poll_id"`
	AuthorName string    `gorm:"column:author_name;not null" json:"author_name"`
	UserID     *string   `gorm:"column:user_id;index:comments_user_id_idx" json:"user_id"`
	CreatedAt  DateTime  `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt  *DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false" json:"updated_at"`
	Poll       *Poll     `gorm:"belongsTo"`
	User       *User     `gorm:"belongsTo"`
}

func (Comment) TableName() string {
	return "comments"
}

func (m *Comment) BeforeCreate(tx *gorm.DB) error {
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

func (*Comment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type RegisteredInstance struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	InstanceID  string   `gorm:"column:instance_id;uniqueIndex:registered_instances_instance_id_key;not null" json:"instance_id"`
	Version     string   `gorm:"column:version;not null" json:"version"`
	FirstSeenAt DateTime `gorm:"column:first_seen_at;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"first_seen_at"`
	LastSeenAt  DateTime `gorm:"column:last_seen_at;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"last_seen_at"`
}

func (RegisteredInstance) TableName() string {
	return "registered_instances"
}

func (m *RegisteredInstance) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.FirstSeenAt.IsZero() {
		m.FirstSeenAt = DateTime{Time: now}
	}
	if m.LastSeenAt.IsZero() {
		m.LastSeenAt = DateTime{Time: now}
	}
	return nil
}

func (*RegisteredInstance) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastSeenAt") {
		tx.Statement.SetColumn("LastSeenAt", DateTime{Time: now})
	}
	return nil
}

type Sheet struct {
	ID          string      `gorm:"column:id;primaryKey" json:"id"`
	SpaceID     string      `gorm:"column:space_id;index:sheets_space_id_idx;not null" json:"space_id"`
	HostID      string      `gorm:"column:host_id;index:sheets_host_id_idx;not null" json:"host_id"`
	Title       string      `gorm:"column:title;not null" json:"title"`
	Description *string     `gorm:"column:description" json:"description"`
	URLID       string      `gorm:"column:url_id;uniqueIndex:sheets_url_id_key;not null" json:"url_id"`
	CreatedAt   DateTime    `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime    `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Deleted     bool        `gorm:"column:deleted;default:false;not null" json:"deleted"`
	DeletedAt   *DateTime   `gorm:"column:deleted_at;type:timestamp(3)" json:"deleted_at"`
	Space       *Space      `gorm:"belongsTo"`
	Host        *User       `gorm:"foreignKey:HostID;belongsTo"`
	Slots       []SheetSlot `gorm:"foreignKey:SheetID;constraint:sheet_slots_sheet_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Sheet) TableName() string {
	return "sheets"
}

func (m *Sheet) BeforeCreate(tx *gorm.DB) error {
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

func (*Sheet) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SheetSlot struct {
	ID             string          `gorm:"column:id;primaryKey" json:"id"`
	SheetID        string          `gorm:"column:sheet_id;index:sheet_slots_sheet_id_start_time_idx;not null" json:"sheet_id"`
	EventTypeID    string          `gorm:"column:event_type_id;index:sheet_slots_event_type_id_idx;not null" json:"event_type_id"`
	StartTime      DateTime        `gorm:"column:start_time;index:sheet_slots_sheet_id_start_time_idx;type:timestamp(3);not null" json:"start_time"`
	CreatedAt      DateTime        `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt      DateTime        `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Sheet          *Sheet          `gorm:"belongsTo"`
	EventType      *EventType      `gorm:"belongsTo"`
	ScheduledEvent *ScheduledEvent `gorm:"foreignKey:SheetSlotID;constraint:scheduled_events_sheet_slot_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (SheetSlot) TableName() string {
	return "sheet_slots"
}

func (m *SheetSlot) BeforeCreate(tx *gorm.DB) error {
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

func (*SheetSlot) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Space struct {
	ID                         string                 `gorm:"column:id;primaryKey" json:"id"`
	Name                       string                 `gorm:"column:name;not null" json:"name"`
	Image                      *string                `gorm:"column:image" json:"image"`
	OwnerID                    string                 `gorm:"column:owner_id;index:spaces_owner_id_idx;not null" json:"owner_id"`
	Tier                       string                 `gorm:"column:tier;default:'hobby';not null" json:"tier"`
	SpaceType                  *string                `gorm:"column:space_type" json:"space_type"`
	Industry                   *string                `gorm:"column:industry" json:"industry"`
	PrimaryColor               *string                `gorm:"column:primary_color" json:"primary_color"`
	ShowBranding               bool                   `gorm:"column:show_branding;default:false;not null" json:"show_branding"`
	HideAttribution            bool                   `gorm:"column:hide_attribution;default:false;not null" json:"hide_attribution"`
	Shared                     bool                   `gorm:"column:shared;default:false;not null" json:"shared"`
	NonprofitDiscountGrantedAt *DateTime              `gorm:"column:nonprofit_discount_granted_at;type:timestamp(3)" json:"nonprofit_discount_granted_at"`
	APIPollQuotaOverride       *int                   `gorm:"column:api_poll_quota_override;size:32" json:"api_poll_quota_override"`
	CreatedAt                  DateTime               `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt                  DateTime               `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Owner                      *User                  `gorm:"foreignKey:OwnerID;belongsTo"`
	Polls                      []Poll                 `gorm:"foreignKey:SpaceID;constraint:polls_space_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ScheduledEvents            []ScheduledEvent       `gorm:"foreignKey:SpaceID;constraint:scheduled_events_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EventTypes                 []EventType            `gorm:"foreignKey:SpaceID;constraint:event_types_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sheets                     []Sheet                `gorm:"foreignKey:SpaceID;constraint:sheets_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subscriptions              []Subscription         `gorm:"foreignKey:SpaceID;constraint:subscriptions_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Members                    []SpaceMember          `gorm:"foreignKey:SpaceID;constraint:space_members_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	MemberInvites              []SpaceMemberInvite    `gorm:"foreignKey:SpaceID;constraint:space_member_invites_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeys                    []SpaceAPIKey          `gorm:"foreignKey:SpaceID;constraint:space_api_keys_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	NonprofitApplications      []NonprofitApplication `gorm:"foreignKey:SpaceID;constraint:nonprofit_applications_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks                   []SpaceWebhook         `gorm:"foreignKey:SpaceID;constraint:space_webhooks_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIUsage                   []APIUsage             `gorm:"foreignKey:SpaceID;constraint:api_usage_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Space) TableName() string {
	return "spaces"
}

func (m *Space) BeforeCreate(tx *gorm.DB) error {
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

func (*Space) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SpaceMember struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	SpaceID        string   `gorm:"column:space_id;uniqueIndex:space_members_space_id_user_id_key;not null" json:"space_id"`
	UserID         string   `gorm:"column:user_id;uniqueIndex:space_members_space_id_user_id_key;index:space_members_user_id_idx;not null" json:"user_id"`
	CreatedAt      DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt      DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Role           string   `gorm:"column:role;default:'MEMBER';not null" json:"role"`
	LastSelectedAt DateTime `gorm:"column:last_selected_at;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"last_selected_at"`
	Space          *Space   `gorm:"belongsTo"`
	User           *User    `gorm:"belongsTo"`
}

func (SpaceMember) TableName() string {
	return "space_members"
}

func (m *SpaceMember) BeforeCreate(tx *gorm.DB) error {
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
	if m.LastSelectedAt.IsZero() {
		m.LastSelectedAt = DateTime{Time: now}
	}
	return nil
}

func (*SpaceMember) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SpaceMemberInvite struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	SpaceID   string   `gorm:"column:space_id;uniqueIndex:space_member_invites_space_id_email_key;not null" json:"space_id"`
	Email     string   `gorm:"column:email;uniqueIndex:space_member_invites_space_id_email_key;not null" json:"email"`
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Role      string   `gorm:"column:role;default:'MEMBER';not null" json:"role"`
	InviterID string   `gorm:"column:inviter_id;not null" json:"inviter_id"`
	InvitedBy *User    `gorm:"foreignKey:InviterID;belongsTo"`
	Space     *Space   `gorm:"belongsTo"`
}

func (SpaceMemberInvite) TableName() string {
	return "space_member_invites"
}

func (m *SpaceMemberInvite) BeforeCreate(tx *gorm.DB) error {
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

func (*SpaceMemberInvite) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SpaceAPIKey struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	SpaceID    string    `gorm:"column:space_id;index:space_api_key_space_idx;not null" json:"space_id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Prefix     string    `gorm:"column:prefix;uniqueIndex:space_api_keys_prefix_key;not null" json:"prefix"`
	HashedKey  string    `gorm:"column:hashed_key;type:text;not null" json:"hashed_key"`
	LastUsedAt *DateTime `gorm:"column:last_used_at;type:timestamp(3)" json:"last_used_at"`
	ExpiresAt  *DateTime `gorm:"column:expires_at;index:space_api_key_expires_idx;type:timestamp(3)" json:"expires_at"`
	RevokedAt  *DateTime `gorm:"column:revoked_at;index:space_api_key_revoked_idx;type:timestamp(3)" json:"revoked_at"`
	CreatedAt  DateTime  `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt  DateTime  `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Space      *Space    `gorm:"belongsTo"`
}

func (SpaceAPIKey) TableName() string {
	return "space_api_keys"
}

func (m *SpaceAPIKey) BeforeCreate(tx *gorm.DB) error {
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

func (*SpaceAPIKey) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SpaceWebhook struct {
	ID      string `gorm:"column:id;primaryKey" json:"id"`
	SpaceID string `gorm:"column:space_id;index:space_webhooks_space_id_idx;not null" json:"space_id"`
	URL     string `gorm:"column:url;not null" json:"url"`
	// AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation
	Secret string `gorm:"column:secret;not null" json:"secret"`
	// Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations
	Events  ScalarList[string] `gorm:"column:events;type:text[]" json:"events"`
	Enabled *bool              `gorm:"column:enabled;default:true;not null" json:"enabled"`
	// The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape
	Version string `gorm:"column:version;not null" json:"version"`
	// Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history
	Cursor          DateTime  `gorm:"column:cursor;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"cursor"`
	LastDeliveredAt *DateTime `gorm:"column:last_delivered_at;type:timestamp(3)" json:"last_delivered_at"`
	LastFailedAt    *DateTime `gorm:"column:last_failed_at;type:timestamp(3)" json:"last_failed_at"`
	// Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap
	ConsecutiveFailures int               `gorm:"column:consecutive_failures;size:32;default:0;not null" json:"consecutive_failures"`
	CreatedAt           DateTime          `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt           DateTime          `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Space               *Space            `gorm:"belongsTo"`
	Deliveries          []WebhookDelivery `gorm:"foreignKey:WebhookID;constraint:webhook_deliveries_webhook_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (SpaceWebhook) TableName() string {
	return "space_webhooks"
}

func (m *SpaceWebhook) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.Cursor.IsZero() {
		m.Cursor = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*SpaceWebhook) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WebhookDelivery struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	WebhookID string `gorm:"column:webhook_id;uniqueIndex:webhook_deliveries_webhook_id_activity_id_key;not null" json:"webhook_id"`
	// Soft reference to the PollActivity row this delivery was fanned out from; with webhookId it makes fan-out idempotent
	ActivityID string `gorm:"column:activity_id;uniqueIndex:webhook_deliveries_webhook_id_activity_id_key;not null" json:"activity_id"`
	EventType  string `gorm:"column:event_type;not null" json:"event_type"`
	// Frozen at fan-out: retries resend exactly what the first attempt sent
	Payload            datatypes.JSON `gorm:"column:payload;not null" json:"payload"`
	Status             string         `gorm:"column:status;index:webhook_deliveries_status_next_attempt_at_idx;default:'pending';not null" json:"status"`
	Attempts           int            `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	NextAttemptAt      DateTime       `gorm:"column:next_attempt_at;index:webhook_deliveries_status_next_attempt_at_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"next_attempt_at"`
	LastResponseStatus *int           `gorm:"column:last_response_status;size:32" json:"last_response_status"`
	LastError          *string        `gorm:"column:last_error" json:"last_error"`
	CreatedAt          DateTime       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt          DateTime       `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Webhook            *SpaceWebhook  `gorm:"foreignKey:WebhookID;belongsTo"`
}

func (WebhookDelivery) TableName() string {
	return "webhook_deliveries"
}

func (m *WebhookDelivery) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.NextAttemptAt.IsZero() {
		m.NextAttemptAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WebhookDelivery) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Account struct {
	ID                    string    `gorm:"column:id;primaryKey" json:"id"`
	UserID                string    `gorm:"column:user_id;index:accounts_user_id_idx;not null" json:"user_id"`
	Provider              string    `gorm:"column:provider;uniqueIndex:accounts_provider_provider_account_id_key;not null" json:"provider"`
	ProviderAccountID     string    `gorm:"column:provider_account_id;uniqueIndex:accounts_provider_provider_account_id_key;not null" json:"provider_account_id"`
	RefreshToken          *string   `gorm:"column:refresh_token;type:text" json:"refresh_token"`
	AccessToken           *string   `gorm:"column:access_token;type:text" json:"access_token"`
	ExpiresAt             *int      `gorm:"column:expires_at;size:32" json:"expires_at"`
	Scope                 *string   `gorm:"column:scope" json:"scope"`
	IDToken               *string   `gorm:"column:id_token;type:text" json:"id_token"`
	AccessTokenExpiresAt  *DateTime `gorm:"column:access_token_expires_at;type:timestamp(3)" json:"access_token_expires_at"`
	RefreshTokenExpiresAt *DateTime `gorm:"column:refresh_token_expires_at;type:timestamp(3)" json:"refresh_token_expires_at"`
	Password              *string   `gorm:"column:password" json:"password"`
	CreatedAt             DateTime  `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt             DateTime  `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	User                  *User     `gorm:"belongsTo"`
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

func (*Account) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type User struct {
	ID                           string                       `gorm:"column:id;primaryKey" json:"id"`
	Name                         string                       `gorm:"column:name;not null" json:"name"`
	Email                        string                       `gorm:"column:email;uniqueIndex:users_email_key;not null" json:"email"`
	EmailVerified                *bool                        `gorm:"column:email_verified" json:"email_verified"`
	Image                        *string                      `gorm:"column:image" json:"image"`
	TimeZone                     *string                      `gorm:"column:time_zone" json:"time_zone"`
	WeekStart                    *int                         `gorm:"column:week_start;size:32" json:"week_start"`
	TimeFormat                   *string                      `gorm:"column:time_format" json:"time_format"`
	Locale                       *string                      `gorm:"column:locale" json:"locale"`
	CreatedAt                    DateTime                     `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt                    *DateTime                    `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false" json:"updated_at"`
	CustomerID                   *string                      `gorm:"column:customer_id" json:"customer_id"`
	Banned                       bool                         `gorm:"column:banned;default:false;not null" json:"banned"`
	BannedAt                     *DateTime                    `gorm:"column:banned_at;type:timestamp(3)" json:"banned_at"`
	BanReason                    *string                      `gorm:"column:ban_reason" json:"ban_reason"`
	BanExpires                   *DateTime                    `gorm:"column:ban_expires;type:timestamp(3)" json:"ban_expires"`
	Role                         string                       `gorm:"column:role;default:'user';not null" json:"role"`
	LastLoginMethod              *string                      `gorm:"column:last_login_method" json:"last_login_method"`
	IsAnonymous                  bool                         `gorm:"column:anonymous;index:users_anonymous_last_seen_at_idx;default:false;not null" json:"anonymous"`
	LastSeenAt                   DateTime                     `gorm:"column:last_seen_at;index:users_anonymous_last_seen_at_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"last_seen_at"`
	DeletedAt                    *DateTime                    `gorm:"column:deleted_at;index:users_deleted_at_idx;type:timestamp(3)" json:"deleted_at"`
	JobTitle                     *string                      `gorm:"column:job_title" json:"job_title"`
	ReviewRequestedAt            *DateTime                    `gorm:"column:review_requested_at;type:timestamp(3)" json:"review_requested_at"`
	DefaultDestinationCalendarID *string                      `gorm:"column:default_destination_calendar_id" json:"default_destination_calendar_id"`
	DefaultDestinationCalendar   *ProviderCalendar            `gorm:"foreignKey:DefaultDestinationCalendarID;belongsTo"`
	Comments                     []Comment                    `gorm:"foreignKey:UserID;constraint:comments_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Polls                        []Poll                       `gorm:"foreignKey:UserID;constraint:polls_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	QueuedEmails                 []QueuedEmail                `gorm:"foreignKey:UserID;constraint:queued_emails_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Accounts                     []Account                    `gorm:"foreignKey:UserID;constraint:accounts_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Participants                 []Participant                `gorm:"foreignKey:UserID;constraint:participants_user_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PaymentMethods               []PaymentMethod              `gorm:"foreignKey:UserID;constraint:payment_methods_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SpaceMemberInvites           []SpaceMemberInvite          `gorm:"foreignKey:InviterID;constraint:space_member_invites_inviter_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	NonprofitApplications        []NonprofitApplication       `gorm:"foreignKey:UserID;constraint:nonprofit_applications_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subscriptions                []Subscription               `gorm:"foreignKey:UserID;constraint:subscriptions_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Spaces                       []Space                      `gorm:"foreignKey:OwnerID;constraint:spaces_owner_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	MemberOf                     []SpaceMember                `gorm:"foreignKey:UserID;constraint:space_members_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ScheduledEvents              []ScheduledEvent             `gorm:"foreignKey:UserID;constraint:scheduled_events_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ScheduledEventInvites        []ScheduledEventInvite       `gorm:"foreignKey:InviteeID;constraint:scheduled_event_invites_invitee_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	HostedEventTypes             []EventType                  `gorm:"foreignKey:HostID;constraint:event_types_host_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	HostedSheets                 []Sheet                      `gorm:"foreignKey:HostID;constraint:sheets_host_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CalendarConnections          []CalendarConnection         `gorm:"foreignKey:UserID;constraint:calendar_connections_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConferencingConnections      []ConferencingConnection     `gorm:"foreignKey:UserID;constraint:conferencing_connections_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Credentials                  []Credential                 `gorm:"foreignKey:UserID;constraint:credentials_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions                     []Session                    `gorm:"foreignKey:UserID;constraint:sessions_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ImpersonatedSessions         []Session                    `gorm:"foreignKey:ImpersonatedBy;constraint:sessions_impersonated_by_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	NotificationPreferences      *UserNotificationPreferences `gorm:"foreignKey:UserID;constraint:user_notification_preferences_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	if m.LastSeenAt.IsZero() {
		m.LastSeenAt = DateTime{Time: now}
	}
	return nil
}

func (*User) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type Session struct {
	ID                 string   `gorm:"column:id;primaryKey" json:"id"`
	ExpiresAt          DateTime `gorm:"column:expires_at;type:timestamp(3);not null" json:"expires_at"`
	Token              string   `gorm:"column:token;uniqueIndex:sessions_token_key;not null" json:"token"`
	CreatedAt          DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt          DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	IPAddress          *string  `gorm:"column:ip_address" json:"ip_address"`
	UserAgent          *string  `gorm:"column:user_agent" json:"user_agent"`
	UserID             string   `gorm:"column:user_id;not null" json:"user_id"`
	ImpersonatedBy     *string  `gorm:"column:impersonated_by" json:"impersonated_by"`
	User               *User    `gorm:"foreignKey:UserID;belongsTo"`
	ImpersonatedByUser *User    `gorm:"foreignKey:ImpersonatedBy;belongsTo"`
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
	Identifier string   `gorm:"column:identifier;uniqueIndex:verifications_identifier_key;index:verifications_identifier_idx;not null" json:"identifier"`
	Value      string   `gorm:"column:value;not null" json:"value"`
	ExpiresAt  DateTime `gorm:"column:expires_at;type:timestamp(3);not null" json:"expires_at"`
	CreatedAt  DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt  DateTime `gorm:"column:updated_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
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
