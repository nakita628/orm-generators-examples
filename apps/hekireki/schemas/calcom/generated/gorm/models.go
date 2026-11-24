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

type Host struct {
	UserID           int           `gorm:"column:userId;primaryKey;index:Host_userId_idx;size:32" json:"userId"`
	EventTypeID      int           `gorm:"column:eventTypeId;primaryKey;index:Host_eventTypeId_idx;size:32" json:"eventTypeId"`
	IsFixed          bool          `gorm:"column:isFixed;default:false;not null" json:"isFixed"`
	Priority         *int          `gorm:"column:priority;size:32" json:"priority"`
	Weight           *int          `gorm:"column:weight;size:32" json:"weight"`
	WeightAdjustment *int          `gorm:"column:weightAdjustment;size:32" json:"weightAdjustment"`
	ScheduleID       *int          `gorm:"column:scheduleId;index:Host_scheduleId_idx;size:32" json:"scheduleId"`
	CreatedAt        DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	GroupID          *string       `gorm:"column:groupId" json:"groupId"`
	MemberID         *int          `gorm:"column:memberId;index:Host_memberId_idx;size:32" json:"memberId"`
	User             *User         `gorm:"belongsTo"`
	EventType        *EventType    `gorm:"belongsTo"`
	Schedule         *Schedule     `gorm:"belongsTo"`
	Group            *HostGroup    `gorm:"foreignKey:GroupID;belongsTo"`
	Member           *Membership   `gorm:"foreignKey:MemberID;belongsTo"`
	Location         *HostLocation `gorm:"foreignKey:UserID,EventTypeID;references:UserID,EventTypeID;constraint:HostLocation_userId_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Host) TableName() string {
	return "Host"
}

func (m *Host) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type HostGroup struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name;index:HostGroup_name_idx;not null" json:"name"`
	EventTypeID *int       `gorm:"column:eventTypeId;index:HostGroup_eventTypeId_idx;size:32" json:"eventTypeId"`
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	EventType   *EventType `gorm:"belongsTo"`
	Hosts       []Host     `gorm:"foreignKey:GroupID;constraint:Host_groupId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (HostGroup) TableName() string {
	return "HostGroup"
}

func (m *HostGroup) BeforeCreate(tx *gorm.DB) error {
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

func (*HostGroup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type HostLocation struct {
	ID           string      `gorm:"column:id;primaryKey" json:"id"`
	UserID       int         `gorm:"column:userId;uniqueIndex:HostLocation_userId_eventTypeId_key;size:32;not null" json:"userId"`
	EventTypeID  int         `gorm:"column:eventTypeId;uniqueIndex:HostLocation_userId_eventTypeId_key;index:HostLocation_eventTypeId_idx;size:32;not null" json:"eventTypeId"`
	Type         string      `gorm:"column:type;not null" json:"type"`
	CredentialID *int        `gorm:"column:credentialId;index:HostLocation_credentialId_idx;size:32" json:"credentialId"`
	Link         *string     `gorm:"column:link" json:"link"`
	Address      *string     `gorm:"column:address" json:"address"`
	PhoneNumber  *string     `gorm:"column:phoneNumber" json:"phoneNumber"`
	CreatedAt    DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime    `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Host         *Host       `gorm:"foreignKey:UserID,EventTypeID;references:UserID,EventTypeID;belongsTo"`
	Credential   *Credential `gorm:"belongsTo"`
}

func (HostLocation) TableName() string {
	return "HostLocation"
}

func (m *HostLocation) BeforeCreate(tx *gorm.DB) error {
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

func (*HostLocation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CalVideoSettings struct {
	EventTypeID                          int        `gorm:"column:eventTypeId;primaryKey;size:32" json:"eventTypeId"`
	DisableRecordingForOrganizer         bool       `gorm:"column:disableRecordingForOrganizer;default:false;not null" json:"disableRecordingForOrganizer"`
	DisableRecordingForGuests            bool       `gorm:"column:disableRecordingForGuests;default:false;not null" json:"disableRecordingForGuests"`
	EnableAutomaticTranscription         bool       `gorm:"column:enableAutomaticTranscription;default:false;not null" json:"enableAutomaticTranscription"`
	EnableAutomaticRecordingForOrganizer bool       `gorm:"column:enableAutomaticRecordingForOrganizer;default:false;not null" json:"enableAutomaticRecordingForOrganizer"`
	RedirectURLOnExit                    *string    `gorm:"column:redirectUrlOnExit" json:"redirectUrlOnExit"`
	DisableTranscriptionForGuests        bool       `gorm:"column:disableTranscriptionForGuests;default:false;not null" json:"disableTranscriptionForGuests"`
	DisableTranscriptionForOrganizer     bool       `gorm:"column:disableTranscriptionForOrganizer;default:false;not null" json:"disableTranscriptionForOrganizer"`
	RequireEmailForGuests                bool       `gorm:"column:requireEmailForGuests;default:false;not null" json:"requireEmailForGuests"`
	CreatedAt                            DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                            DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	EventType                            *EventType `gorm:"belongsTo"`
}

func (CalVideoSettings) TableName() string {
	return "CalVideoSettings"
}

func (m *CalVideoSettings) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CalVideoSettings) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type VideoCallGuest struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	BookingUID string    `gorm:"column:bookingUid;uniqueIndex:VideoCallGuest_bookingUid_email_key;index:VideoCallGuest_bookingUid_idx;not null" json:"bookingUid"`
	Email      string    `gorm:"column:email;uniqueIndex:VideoCallGuest_bookingUid_email_key;index:VideoCallGuest_email_idx;not null" json:"email"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	JoinedAt   *DateTime `gorm:"column:joinedAt;type:timestamp(3)" json:"joinedAt"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (VideoCallGuest) TableName() string {
	return "VideoCallGuest"
}

func (m *VideoCallGuest) BeforeCreate(tx *gorm.DB) error {
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

func (*VideoCallGuest) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EventType struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	// @zod.string.min(1)
	Title string `gorm:"column:title;not null" json:"title"`
	// @zod.import(["import { eventTypeSlug } from '../../zod-utils'"]).custom.use(eventTypeSlug)
	Slug              string  `gorm:"column:slug;uniqueIndex:EventType_userId_slug_key;uniqueIndex:EventType_teamId_slug_key;not null" json:"slug"`
	Description       *string `gorm:"column:description" json:"description"`
	InterfaceLanguage *string `gorm:"column:interfaceLanguage" json:"interfaceLanguage"`
	Position          int     `gorm:"column:position;size:32;default:0;not null" json:"position"`
	// @zod.import(["import { eventTypeLocations } from '../../zod-utils'"]).custom.use(eventTypeLocations)
	Locations datatypes.JSON `gorm:"column:locations" json:"locations"`
	// @zod.number.min(1)
	Length                         int     `gorm:"column:length;size:32;not null" json:"length"`
	OffsetStart                    int     `gorm:"column:offsetStart;size:32;default:0;not null" json:"offsetStart"`
	Hidden                         bool    `gorm:"column:hidden;default:false;not null" json:"hidden"`
	UserID                         *int    `gorm:"column:userId;uniqueIndex:EventType_userId_slug_key;uniqueIndex:EventType_userId_parentId_key;index:EventType_userId_idx;size:32" json:"userId"`
	ProfileID                      *int    `gorm:"column:profileId;index:EventType_profileId_idx;size:32" json:"profileId"`
	TeamID                         *int    `gorm:"column:teamId;uniqueIndex:EventType_teamId_slug_key;index:EventType_teamId_idx;index:EventType_parentId_teamId_idx;size:32" json:"teamId"`
	UseEventLevelSelectedCalendars bool    `gorm:"column:useEventLevelSelectedCalendars;default:false;not null" json:"useEventLevelSelectedCalendars"`
	EventName                      *string `gorm:"column:eventName" json:"eventName"`
	ParentID                       *int    `gorm:"column:parentId;uniqueIndex:EventType_userId_parentId_key;index:EventType_parentId_idx;index:EventType_parentId_teamId_idx;size:32" json:"parentId"`
	// @zod.import(["import { eventTypeBookingFields } from '../../zod-utils'"]).custom.use(eventTypeBookingFields)
	BookingFields datatypes.JSON `gorm:"column:bookingFields" json:"bookingFields"`
	TimeZone      *string        `gorm:"column:timeZone" json:"timeZone"`
	PeriodType    string         `gorm:"column:periodType;default:'unlimited';not null" json:"periodType"`
	// @zod.import(["import { coerceToDate } from '../../zod-utils'"]).custom.use(coerceToDate)
	PeriodStartDate *DateTime `gorm:"column:periodStartDate;type:timestamp(3)" json:"periodStartDate"`
	// @zod.import(["import { coerceToDate } from '../../zod-utils'"]).custom.use(coerceToDate)
	PeriodEndDate                           *DateTime `gorm:"column:periodEndDate;type:timestamp(3)" json:"periodEndDate"`
	PeriodDays                              *int      `gorm:"column:periodDays;size:32" json:"periodDays"`
	PeriodCountCalendarDays                 *bool     `gorm:"column:periodCountCalendarDays" json:"periodCountCalendarDays"`
	LockTimeZoneToggleOnBookingPage         bool      `gorm:"column:lockTimeZoneToggleOnBookingPage;default:false;not null" json:"lockTimeZoneToggleOnBookingPage"`
	LockedTimeZone                          *string   `gorm:"column:lockedTimeZone" json:"lockedTimeZone"`
	RequiresConfirmation                    bool      `gorm:"column:requiresConfirmation;default:false;not null" json:"requiresConfirmation"`
	RequiresConfirmationWillBlockSlot       bool      `gorm:"column:requiresConfirmationWillBlockSlot;default:false;not null" json:"requiresConfirmationWillBlockSlot"`
	RequiresConfirmationForFreeEmail        bool      `gorm:"column:requiresConfirmationForFreeEmail;default:false;not null" json:"requiresConfirmationForFreeEmail"`
	RequiresBookerEmailVerification         bool      `gorm:"column:requiresBookerEmailVerification;default:false;not null" json:"requiresBookerEmailVerification"`
	CanSendCalVideoTranscriptionEmails      *bool     `gorm:"column:canSendCalVideoTranscriptionEmails;default:true;not null" json:"canSendCalVideoTranscriptionEmails"`
	AutoTranslateDescriptionEnabled         bool      `gorm:"column:autoTranslateDescriptionEnabled;default:false;not null" json:"autoTranslateDescriptionEnabled"`
	AutoTranslateInstantMeetingTitleEnabled *bool     `gorm:"column:autoTranslateInstantMeetingTitleEnabled;default:true;not null" json:"autoTranslateInstantMeetingTitleEnabled"`
	// @zod.import(["import { recurringEventType } from '../../zod-utils'"]).custom.use(recurringEventType)
	RecurringEvent           datatypes.JSON `gorm:"column:recurringEvent" json:"recurringEvent"`
	DisableGuests            bool           `gorm:"column:disableGuests;default:false;not null" json:"disableGuests"`
	HideCalendarNotes        bool           `gorm:"column:hideCalendarNotes;default:false;not null" json:"hideCalendarNotes"`
	HideCalendarEventDetails bool           `gorm:"column:hideCalendarEventDetails;default:false;not null" json:"hideCalendarEventDetails"`
	// @zod.number.min(0)
	MinimumBookingNotice       *int  `gorm:"column:minimumBookingNotice;size:32;default:120;not null" json:"minimumBookingNotice"`
	BeforeEventBuffer          int   `gorm:"column:beforeEventBuffer;size:32;default:0;not null" json:"beforeEventBuffer"`
	AfterEventBuffer           int   `gorm:"column:afterEventBuffer;size:32;default:0;not null" json:"afterEventBuffer"`
	SeatsPerTimeSlot           *int  `gorm:"column:seatsPerTimeSlot;size:32" json:"seatsPerTimeSlot"`
	OnlyShowFirstAvailableSlot bool  `gorm:"column:onlyShowFirstAvailableSlot;default:false;not null" json:"onlyShowFirstAvailableSlot"`
	ShowOptimizedSlots         *bool `gorm:"column:showOptimizedSlots;default:false" json:"showOptimizedSlots"`
	DisableCancelling          *bool `gorm:"column:disableCancelling;default:false" json:"disableCancelling"`
	DisableRescheduling        *bool `gorm:"column:disableRescheduling;default:false" json:"disableRescheduling"`
	// @zod.number.min(0)
	MinimumRescheduleNotice            *int    `gorm:"column:minimumRescheduleNotice;size:32" json:"minimumRescheduleNotice"`
	SeatsShowAttendees                 *bool   `gorm:"column:seatsShowAttendees;default:false" json:"seatsShowAttendees"`
	SeatsShowAvailabilityCount         *bool   `gorm:"column:seatsShowAvailabilityCount;default:true" json:"seatsShowAvailabilityCount"`
	SchedulingType                     *string `gorm:"column:schedulingType" json:"schedulingType"`
	ScheduleID                         *int    `gorm:"column:scheduleId;index:EventType_scheduleId_idx;size:32" json:"scheduleId"`
	AllowReschedulingCancelledBookings *bool   `gorm:"column:allowReschedulingCancelledBookings;default:false" json:"allowReschedulingCancelledBookings"`
	Price                              int     `gorm:"column:price;size:32;default:0;not null" json:"price"`
	Currency                           *string `gorm:"column:currency;default:'usd';not null" json:"currency"`
	SlotInterval                       *int    `gorm:"column:slotInterval;size:32" json:"slotInterval"`
	// @zod.import(["import { EventTypeMetaDataSchema } from '../../zod-utils'"]).custom.use(EventTypeMetaDataSchema)
	Metadata datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	// @zod.import(["import { successRedirectUrl } from '../../zod-utils'"]).custom.use(successRedirectUrl)
	SuccessRedirectURL           *string `gorm:"column:successRedirectUrl" json:"successRedirectUrl"`
	ForwardParamsSuccessRedirect *bool   `gorm:"column:forwardParamsSuccessRedirect;default:true" json:"forwardParamsSuccessRedirect"`
	// @zod.import(["import { intervalLimitsType } from '../../zod-utils'"]).custom.use(intervalLimitsType)
	BookingLimits datatypes.JSON `gorm:"column:bookingLimits" json:"bookingLimits"`
	// @zod.import(["import { intervalLimitsType } from '../../zod-utils'"]).custom.use(intervalLimitsType)
	DurationLimits                          datatypes.JSON     `gorm:"column:durationLimits" json:"durationLimits"`
	IsInstantEvent                          bool               `gorm:"column:isInstantEvent;default:false;not null" json:"isInstantEvent"`
	InstantMeetingExpiryTimeOffsetInSeconds *int               `gorm:"column:instantMeetingExpiryTimeOffsetInSeconds;size:32;default:90;not null" json:"instantMeetingExpiryTimeOffsetInSeconds"`
	InstantMeetingScheduleID                *int               `gorm:"column:instantMeetingScheduleId;size:32" json:"instantMeetingScheduleId"`
	InstantMeetingParameters                ScalarList[string] `gorm:"column:instantMeetingParameters;type:text[]" json:"instantMeetingParameters"`
	AssignAllTeamMembers                    bool               `gorm:"column:assignAllTeamMembers;default:false;not null" json:"assignAllTeamMembers"`
	AssignRRMembersUsingSegment             bool               `gorm:"column:assignRRMembersUsingSegment;default:false;not null" json:"assignRRMembersUsingSegment"`
	// @zod.import(["import { rrSegmentQueryValueSchema } from '../../zod-utils'"]).custom.use(rrSegmentQueryValueSchema)
	RrSegmentQueryValue                      datatypes.JSON `gorm:"column:rrSegmentQueryValue" json:"rrSegmentQueryValue"`
	UseEventTypeDestinationCalendarEmail     bool           `gorm:"column:useEventTypeDestinationCalendarEmail;default:false;not null" json:"useEventTypeDestinationCalendarEmail"`
	IsRRWeightsEnabled                       bool           `gorm:"column:isRRWeightsEnabled;default:false;not null" json:"isRRWeightsEnabled"`
	MaxLeadThreshold                         *int           `gorm:"column:maxLeadThreshold;size:32" json:"maxLeadThreshold"`
	IncludeNoShowInRRCalculation             bool           `gorm:"column:includeNoShowInRRCalculation;default:false;not null" json:"includeNoShowInRRCalculation"`
	AllowReschedulingPastBookings            bool           `gorm:"column:allowReschedulingPastBookings;default:false;not null" json:"allowReschedulingPastBookings"`
	HideOrganizerEmail                       bool           `gorm:"column:hideOrganizerEmail;default:false;not null" json:"hideOrganizerEmail"`
	MaxActiveBookingsPerBooker               *int           `gorm:"column:maxActiveBookingsPerBooker;size:32" json:"maxActiveBookingsPerBooker"`
	MaxActiveBookingPerBookerOfferReschedule bool           `gorm:"column:maxActiveBookingPerBookerOfferReschedule;default:false;not null" json:"maxActiveBookingPerBookerOfferReschedule"`
	// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
	CustomReplyToEmail *string `gorm:"column:customReplyToEmail" json:"customReplyToEmail"`
	// @zod.import(["import { eventTypeColor } from '../../zod-utils'"]).custom.use(eventTypeColor)
	EventTypeColor                   datatypes.JSON         `gorm:"column:eventTypeColor" json:"eventTypeColor"`
	RescheduleWithSameRoundRobinHost bool                   `gorm:"column:rescheduleWithSameRoundRobinHost;default:false;not null" json:"rescheduleWithSameRoundRobinHost"`
	SecondaryEmailID                 *int                   `gorm:"column:secondaryEmailId;index:EventType_secondaryEmailId_idx;size:32" json:"secondaryEmailId"`
	UseBookerTimezone                bool                   `gorm:"column:useBookerTimezone;default:false;not null" json:"useBookerTimezone"`
	RestrictionScheduleID            *int                   `gorm:"column:restrictionScheduleId;index:EventType_restrictionScheduleId_idx;size:32" json:"restrictionScheduleId"`
	BookingRequiresAuthentication    bool                   `gorm:"column:bookingRequiresAuthentication;default:false;not null" json:"bookingRequiresAuthentication"`
	RrHostSubsetEnabled              bool                   `gorm:"column:rrHostSubsetEnabled;default:false;not null" json:"rrHostSubsetEnabled"`
	RequiresCancellationReason       *string                `gorm:"column:requiresCancellationReason;default:'MANDATORY_HOST_ONLY'" json:"requiresCancellationReason"`
	EnablePerHostLocations           bool                   `gorm:"column:enablePerHostLocations;default:false;not null" json:"enablePerHostLocations"`
	CreatedAt                        *DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false" json:"createdAt"`
	UpdatedAt                        *DateTime              `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	Owner                            *User                  `gorm:"foreignKey:UserID;belongsTo"`
	Profile                          *Profile               `gorm:"belongsTo"`
	Team                             *Team                  `gorm:"belongsTo"`
	Parent                           *EventType             `gorm:"foreignKey:ParentID;belongsTo"`
	Schedule                         *Schedule              `gorm:"foreignKey:ScheduleID;belongsTo"`
	InstantMeetingSchedule           *Schedule              `gorm:"foreignKey:InstantMeetingScheduleID;belongsTo"`
	SecondaryEmail                   *SecondaryEmail        `gorm:"belongsTo"`
	RestrictionSchedule              *Schedule              `gorm:"foreignKey:RestrictionScheduleID;belongsTo"`
	Hosts                            []Host                 `gorm:"foreignKey:EventTypeID;constraint:Host_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	HashedLink                       []HashedLink           `gorm:"foreignKey:EventTypeID;constraint:HashedLink_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Bookings                         []Booking              `gorm:"foreignKey:EventTypeID;constraint:Booking_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Availability                     []Availability         `gorm:"foreignKey:EventTypeID;constraint:Availability_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Webhooks                         []Webhook              `gorm:"foreignKey:EventTypeID;constraint:Webhook_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomInputs                     []EventTypeCustomInput `gorm:"foreignKey:EventTypeID;constraint:EventTypeCustomInput_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Children                         []EventType            `gorm:"foreignKey:ParentID;constraint:EventType_parentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FieldTranslations                []EventTypeTranslation `gorm:"foreignKey:EventTypeID;constraint:EventTypeTranslation_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SelectedCalendars                []SelectedCalendar     `gorm:"foreignKey:EventTypeID;constraint:SelectedCalendar_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	HostGroups                       []HostGroup            `gorm:"foreignKey:EventTypeID;constraint:HostGroup_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DestinationCalendar              *DestinationCalendar   `gorm:"foreignKey:EventTypeID;constraint:DestinationCalendar_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CalVideoSettings                 *CalVideoSettings      `gorm:"foreignKey:EventTypeID;constraint:CalVideoSettings_eventTypeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Users                            []User                 `gorm:"many2many:_user_eventtype;joinForeignKey:A;joinReferences:B"`
}

func (EventType) TableName() string {
	return "EventType"
}

func (m *EventType) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt == nil {
		m.CreatedAt = &DateTime{Time: now}
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	return nil
}

func (*EventType) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type Credential struct {
	ID                        int                        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Type                      string                     `gorm:"column:type;not null" json:"type"`
	Key                       datatypes.JSON             `gorm:"column:key;not null" json:"key"`
	EncryptedKey              *string                    `gorm:"column:encryptedKey" json:"encryptedKey"`
	UserID                    *int                       `gorm:"column:userId;index:Credential_userId_delegationCredentialId_idx;size:32" json:"userId"`
	TeamID                    *int                       `gorm:"column:teamId;index:Credential_teamId_idx;size:32" json:"teamId"`
	AppID                     *string                    `gorm:"column:appId;index:Credential_appId_idx" json:"appId"`
	SubscriptionID            *string                    `gorm:"column:subscriptionId;index:Credential_subscriptionId_idx" json:"subscriptionId"`
	PaymentStatus             *string                    `gorm:"column:paymentStatus" json:"paymentStatus"`
	BillingCycleStart         *int                       `gorm:"column:billingCycleStart;size:32" json:"billingCycleStart"`
	Invalid                   *bool                      `gorm:"column:invalid;index:Credential_invalid_idx;default:false" json:"invalid"`
	DelegationCredentialID    *string                    `gorm:"column:delegationCredentialId;index:Credential_userId_delegationCredentialId_idx" json:"delegationCredentialId"`
	User                      *User                      `gorm:"belongsTo"`
	Team                      *Team                      `gorm:"belongsTo"`
	App                       *App                       `gorm:"foreignKey:AppID;references:Slug;belongsTo"`
	DelegationCredential      *DelegationCredential      `gorm:"belongsTo"`
	DestinationCalendars      []DestinationCalendar      `gorm:"foreignKey:CredentialID;constraint:DestinationCalendar_credentialId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SelectedCalendars         []SelectedCalendar         `gorm:"foreignKey:CredentialID;constraint:SelectedCalendar_credentialId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CalendarCache             []CalendarCache            `gorm:"foreignKey:CredentialID;constraint:CalendarCache_credentialId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	References                []BookingReference         `gorm:"foreignKey:CredentialID;constraint:BookingReference_credentialId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	IntegrationAttributeSyncs []IntegrationAttributeSync `gorm:"foreignKey:CredentialID;constraint:IntegrationAttributeSync_credentialId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	HostLocations             []HostLocation             `gorm:"foreignKey:CredentialID;constraint:HostLocation_credentialId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Credential) TableName() string {
	return "Credential"
}

type DestinationCalendar struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Integration string `gorm:"column:integration;not null" json:"integration"`
	ExternalID  string `gorm:"column:externalId;not null" json:"externalId"`
	// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
	PrimaryEmail           *string               `gorm:"column:primaryEmail" json:"primaryEmail"`
	UserID                 *int                  `gorm:"column:userId;uniqueIndex:DestinationCalendar_userId_key;index:DestinationCalendar_userId_idx;size:32" json:"userId"`
	EventTypeID            *int                  `gorm:"column:eventTypeId;uniqueIndex:DestinationCalendar_eventTypeId_key;index:DestinationCalendar_eventTypeId_idx;size:32" json:"eventTypeId"`
	CredentialID           *int                  `gorm:"column:credentialId;index:DestinationCalendar_credentialId_idx;size:32" json:"credentialId"`
	CreatedAt              *DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false" json:"createdAt"`
	UpdatedAt              *DateTime             `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	DelegationCredentialID *string               `gorm:"column:delegationCredentialId" json:"delegationCredentialId"`
	CustomCalendarReminder *int                  `gorm:"column:customCalendarReminder;size:32" json:"customCalendarReminder"`
	User                   *User                 `gorm:"belongsTo"`
	EventType              *EventType            `gorm:"belongsTo"`
	Credential             *Credential           `gorm:"belongsTo"`
	DelegationCredential   *DelegationCredential `gorm:"belongsTo"`
	Booking                []Booking             `gorm:"foreignKey:DestinationCalendarID;constraint:Booking_destinationCalendarId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (DestinationCalendar) TableName() string {
	return "DestinationCalendar"
}

func (m *DestinationCalendar) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt == nil {
		m.CreatedAt = &DateTime{Time: now}
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	return nil
}

func (*DestinationCalendar) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type UserPassword struct {
	Hash   string `gorm:"column:hash;not null" json:"hash"`
	UserID int    `gorm:"column:userId;primaryKey;uniqueIndex:UserPassword_userId_key;size:32" json:"userId"`
	User   *User  `gorm:"belongsTo"`
}

func (UserPassword) TableName() string {
	return "UserPassword"
}

type TravelSchedule struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID       int       `gorm:"column:userId;size:32;not null" json:"userId"`
	TimeZone     string    `gorm:"column:timeZone;not null" json:"timeZone"`
	StartDate    DateTime  `gorm:"column:startDate;index:TravelSchedule_startDate_idx;type:timestamp(3);not null" json:"startDate"`
	EndDate      *DateTime `gorm:"column:endDate;index:TravelSchedule_endDate_idx;type:timestamp(3)" json:"endDate"`
	PrevTimeZone *string   `gorm:"column:prevTimeZone" json:"prevTimeZone"`
	User         *User     `gorm:"belongsTo"`
}

func (TravelSchedule) TableName() string {
	return "TravelSchedule"
}

type User struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UUID     string  `gorm:"column:uuid;uniqueIndex:users_uuid_key;type:char(36);not null" json:"uuid"`
	Username *string `gorm:"column:username;uniqueIndex:users_email_username_key;uniqueIndex:users_username_organizationId_key;index:users_username_idx" json:"username"`
	Name     *string `gorm:"column:name" json:"name"`
	// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
	Email                           string    `gorm:"column:email;uniqueIndex:users_email_key;uniqueIndex:users_email_username_key;not null" json:"email"`
	EmailVerified                   *DateTime `gorm:"column:emailVerified;index:users_emailVerified_idx;type:timestamp(3)" json:"emailVerified"`
	Bio                             *string   `gorm:"column:bio" json:"bio"`
	AvatarURL                       *string   `gorm:"column:avatarUrl" json:"avatarUrl"`
	TimeZone                        *string   `gorm:"column:timeZone;default:'Europe/London';not null" json:"timeZone"`
	WeekStart                       *string   `gorm:"column:weekStart;default:'Sunday';not null" json:"weekStart"`
	BufferTime                      int       `gorm:"column:bufferTime;size:32;default:0;not null" json:"bufferTime"`
	HideBranding                    bool      `gorm:"column:hideBranding;default:false;not null" json:"hideBranding"`
	Theme                           *string   `gorm:"column:theme" json:"theme"`
	AppTheme                        *string   `gorm:"column:appTheme" json:"appTheme"`
	CreatedDate                     DateTime  `gorm:"column:created;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created"`
	TrialEndsAt                     *DateTime `gorm:"column:trialEndsAt;type:timestamp(3)" json:"trialEndsAt"`
	LastActiveAt                    *DateTime `gorm:"column:lastActiveAt;type:timestamp(3)" json:"lastActiveAt"`
	DefaultScheduleID               *int      `gorm:"column:defaultScheduleId;size:32" json:"defaultScheduleId"`
	CompletedOnboarding             bool      `gorm:"column:completedOnboarding;default:false;not null" json:"completedOnboarding"`
	Locale                          *string   `gorm:"column:locale" json:"locale"`
	TimeFormat                      *int      `gorm:"column:timeFormat;size:32;default:12" json:"timeFormat"`
	TwoFactorSecret                 *string   `gorm:"column:twoFactorSecret" json:"twoFactorSecret"`
	TwoFactorEnabled                bool      `gorm:"column:twoFactorEnabled;default:false;not null" json:"twoFactorEnabled"`
	BackupCodes                     *string   `gorm:"column:backupCodes" json:"backupCodes"`
	IdentityProvider                string    `gorm:"column:identityProvider;index:users_identityProvider_idx;default:'CAL';not null" json:"identityProvider"`
	IdentityProviderID              *string   `gorm:"column:identityProviderId;index:users_identityProviderId_idx" json:"identityProviderId"`
	InvitedTo                       *int      `gorm:"column:invitedTo;size:32" json:"invitedTo"`
	BrandColor                      *string   `gorm:"column:brandColor" json:"brandColor"`
	DarkBrandColor                  *string   `gorm:"column:darkBrandColor" json:"darkBrandColor"`
	AllowDynamicBooking             *bool     `gorm:"column:allowDynamicBooking;default:true" json:"allowDynamicBooking"`
	AllowSEOIndexing                *bool     `gorm:"column:allowSEOIndexing;default:true" json:"allowSEOIndexing"`
	ReceiveMonthlyDigestEmail       *bool     `gorm:"column:receiveMonthlyDigestEmail;default:true" json:"receiveMonthlyDigestEmail"`
	RequiresBookerEmailVerification *bool     `gorm:"column:requiresBookerEmailVerification;default:false" json:"requiresBookerEmailVerification"`
	// @zod.import(["import { userMetadata } from '../../zod-utils'"]).custom.use(userMetadata)
	Metadata                       datatypes.JSON                `gorm:"column:metadata" json:"metadata"`
	Verified                       *bool                         `gorm:"column:verified;default:false" json:"verified"`
	Role                           string                        `gorm:"column:role;default:'USER';not null" json:"role"`
	OrganizationID                 *int                          `gorm:"column:organizationId;uniqueIndex:users_username_organizationId_key;size:32" json:"organizationId"`
	Locked                         bool                          `gorm:"column:locked;default:false;not null" json:"locked"`
	MovedToProfileID               *int                          `gorm:"column:movedToProfileId;uniqueIndex:users_movedToProfileId_key;size:32" json:"movedToProfileId"`
	IsPlatformManaged              bool                          `gorm:"column:isPlatformManaged;default:false;not null" json:"isPlatformManaged"`
	SmsLockState                   string                        `gorm:"column:smsLockState;default:'UNLOCKED';not null" json:"smsLockState"`
	SmsLockReviewedByAdmin         bool                          `gorm:"column:smsLockReviewedByAdmin;default:false;not null" json:"smsLockReviewedByAdmin"`
	ReferralLinkID                 *string                       `gorm:"column:referralLinkId" json:"referralLinkId"`
	CreationSource                 *string                       `gorm:"column:creationSource" json:"creationSource"`
	AutoOptInFeatures              bool                          `gorm:"column:autoOptInFeatures;default:false;not null" json:"autoOptInFeatures"`
	Organization                   *Team                         `gorm:"foreignKey:OrganizationID;belongsTo"`
	MovedToProfile                 *Profile                      `gorm:"foreignKey:MovedToProfileID;belongsTo"`
	TravelSchedules                []TravelSchedule              `gorm:"foreignKey:UserID;constraint:TravelSchedule_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Credentials                    []Credential                  `gorm:"foreignKey:UserID;constraint:Credential_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Teams                          []Membership                  `gorm:"foreignKey:UserID;constraint:Membership_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Bookings                       []Booking                     `gorm:"foreignKey:UserID;constraint:Booking_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Schedules                      []Schedule                    `gorm:"foreignKey:UserID;constraint:Schedule_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SelectedCalendars              []SelectedCalendar            `gorm:"foreignKey:UserID;constraint:SelectedCalendar_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Availability                   []Availability                `gorm:"foreignKey:UserID;constraint:Availability_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks                       []Webhook                     `gorm:"foreignKey:UserID;constraint:Webhook_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeys                        []APIKey                      `gorm:"foreignKey:UserID;constraint:ApiKey_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Accounts                       []Account                     `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions                       []Session                     `gorm:"foreignKey:UserID;constraint:Session_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Feedback                       []Feedback                    `gorm:"foreignKey:UserID;constraint:Feedback_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OwnedEventTypes                []EventType                   `gorm:"foreignKey:UserID;constraint:EventType_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	VerifiedNumbers                []VerifiedNumber              `gorm:"foreignKey:UserID;constraint:VerifiedNumber_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	VerifiedEmails                 []VerifiedEmail               `gorm:"foreignKey:UserID;constraint:VerifiedEmail_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Hosts                          []Host                        `gorm:"foreignKey:UserID;constraint:Host_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessCodes                    []AccessCode                  `gorm:"foreignKey:UserID;constraint:AccessCode_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BookingRedirects               []OutOfOfficeEntry            `gorm:"foreignKey:UserID;constraint:OutOfOfficeEntry_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BookingRedirectsTo             []OutOfOfficeEntry            `gorm:"foreignKey:ToUserID;constraint:OutOfOfficeEntry_toUserId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessToken                    []AccessToken                 `gorm:"foreignKey:UserID;constraint:AccessToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RefreshToken                   []RefreshToken                `gorm:"foreignKey:UserID;constraint:RefreshToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlatformAuthorizationToken     []PlatformAuthorizationToken  `gorm:"foreignKey:UserID;constraint:PlatformAuthorizationToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Profiles                       []Profile                     `gorm:"foreignKey:UserID;constraint:Profile_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SecondaryEmails                []SecondaryEmail              `gorm:"foreignKey:UserID;constraint:SecondaryEmail_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OutOfOfficeReasons             []OutOfOfficeReason           `gorm:"foreignKey:UserID;constraint:OutOfOfficeReason_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	NotificationsSubscriptions     []NotificationsSubscriptions  `gorm:"foreignKey:UserID;constraint:NotificationsSubscriptions_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Features                       []UserFeatures                `gorm:"foreignKey:UserID;constraint:UserFeatures_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReassignedBookings             []Booking                     `gorm:"foreignKey:ReassignByID;constraint:Booking_reassignById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAttributeToUsers        []AttributeToUser             `gorm:"foreignKey:CreatedByID;constraint:AttributeToUser_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	UpdatedAttributeToUsers        []AttributeToUser             `gorm:"foreignKey:UpdatedByID;constraint:AttributeToUser_updatedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedTranslations            []EventTypeTranslation        `gorm:"foreignKey:CreatedBy;constraint:EventTypeTranslation_createdBy_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UpdatedTranslations            []EventTypeTranslation        `gorm:"foreignKey:UpdatedBy;constraint:EventTypeTranslation_updatedBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ReportedBookings               []BookingReport               `gorm:"foreignKey:ReportedByID;constraint:BookingReport_reportedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	WrongAssignmentReports         []WrongAssignmentReport       `gorm:"foreignKey:ReportedByID;constraint:WrongAssignmentReport_reportedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ReviewedWrongAssignmentReports []WrongAssignmentReport       `gorm:"foreignKey:ReviewedByID;constraint:WrongAssignmentReport_reviewedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	BookingInternalNote            []BookingInternalNote         `gorm:"foreignKey:CreatedByID;constraint:BookingInternalNote_createdById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedOrganizationOnboardings []OrganizationOnboarding      `gorm:"foreignKey:CreatedByID;constraint:OrganizationOnboarding_createdById_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FilterSegments                 []FilterSegment               `gorm:"foreignKey:UserID;constraint:FilterSegment_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FilterSegmentPreferences       []UserFilterSegmentPreference `gorm:"foreignKey:UserID;constraint:UserFilterSegmentPreference_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CalAiPhoneNumbers              []CalAiPhoneNumber            `gorm:"foreignKey:UserID;constraint:CalAiPhoneNumber_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Agents                         []Agent                       `gorm:"foreignKey:UserID;constraint:Agent_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OAuthClients                   []OAuthClient                 `gorm:"foreignKey:UserID;constraint:OAuthClient_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Password                       *UserPassword                 `gorm:"foreignKey:UserID;constraint:UserPassword_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DestinationCalendar            *DestinationCalendar          `gorm:"foreignKey:UserID;constraint:DestinationCalendar_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreditBalance                  *CreditBalance                `gorm:"foreignKey:UserID;constraint:CreditBalance_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	HolidaySettings                *UserHolidaySettings          `gorm:"foreignKey:UserID;constraint:UserHolidaySettings_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EventTypes                     []EventType                   `gorm:"many2many:_user_eventtype;joinForeignKey:B;joinReferences:A"`
	PlatformOAuthClients           []PlatformOAuthClient         `gorm:"many2many:_PlatformOAuthClientToUser;joinForeignKey:B;joinReferences:A"`
}

func (User) TableName() string {
	return "users"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UUID == "" {
		m.UUID = uuid.NewString()
	}
	if m.CreatedDate.IsZero() {
		m.CreatedDate = DateTime{Time: now}
	}
	return nil
}

type NotificationsSubscriptions struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID       int    `gorm:"column:userId;index:NotificationsSubscriptions_userId_subscription_idx;size:32;not null" json:"userId"`
	Subscription string `gorm:"column:subscription;index:NotificationsSubscriptions_userId_subscription_idx;not null" json:"subscription"`
	User         *User  `gorm:"belongsTo"`
}

func (NotificationsSubscriptions) TableName() string {
	return "NotificationsSubscriptions"
}

type Profile struct {
	ID             int         `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UID            string      `gorm:"column:uid;index:Profile_uid_idx;not null" json:"uid"`
	UserID         int         `gorm:"column:userId;uniqueIndex:Profile_userId_organizationId_key;index:Profile_userId_idx;size:32;not null" json:"userId"`
	OrganizationID int         `gorm:"column:organizationId;uniqueIndex:Profile_userId_organizationId_key;uniqueIndex:Profile_username_organizationId_key;index:Profile_organizationId_idx;size:32;not null" json:"organizationId"`
	Username       string      `gorm:"column:username;uniqueIndex:Profile_username_organizationId_key;not null" json:"username"`
	CreatedAt      DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime    `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User           *User       `gorm:"belongsTo"`
	Organization   *Team       `gorm:"foreignKey:OrganizationID;belongsTo"`
	EventTypes     []EventType `gorm:"foreignKey:ProfileID;constraint:EventType_profileId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	MovedFromUser  *User       `gorm:"foreignKey:MovedToProfileID;constraint:users_movedToProfileId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Profile) TableName() string {
	return "Profile"
}

func (m *Profile) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Profile) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Team struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	// @zod.string.min(1)
	Name string `gorm:"column:name;not null" json:"name"`
	// @zod.string.min(1)
	Slug                *string  `gorm:"column:slug;uniqueIndex:Team_slug_parentId_key" json:"slug"`
	LogoURL             *string  `gorm:"column:logoUrl" json:"logoUrl"`
	CalVideoLogo        *string  `gorm:"column:calVideoLogo" json:"calVideoLogo"`
	AppLogo             *string  `gorm:"column:appLogo" json:"appLogo"`
	AppIconLogo         *string  `gorm:"column:appIconLogo" json:"appIconLogo"`
	Bio                 *string  `gorm:"column:bio" json:"bio"`
	HideBranding        bool     `gorm:"column:hideBranding;default:false;not null" json:"hideBranding"`
	HideTeamProfileLink bool     `gorm:"column:hideTeamProfileLink;default:false;not null" json:"hideTeamProfileLink"`
	IsPrivate           bool     `gorm:"column:isPrivate;default:false;not null" json:"isPrivate"`
	HideBookATeamMember bool     `gorm:"column:hideBookATeamMember;default:false;not null" json:"hideBookATeamMember"`
	CreatedAt           DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	// @zod.import(["import { teamMetadataSchema } from '../../zod-utils'"]).custom.use(teamMetadataSchema)
	Metadata               datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	Theme                  *string        `gorm:"column:theme" json:"theme"`
	RrResetInterval        *string        `gorm:"column:rrResetInterval;default:'MONTH'" json:"rrResetInterval"`
	RrTimestampBasis       string         `gorm:"column:rrTimestampBasis;default:'CREATED_AT';not null" json:"rrTimestampBasis"`
	BrandColor             *string        `gorm:"column:brandColor" json:"brandColor"`
	DarkBrandColor         *string        `gorm:"column:darkBrandColor" json:"darkBrandColor"`
	BannerURL              *string        `gorm:"column:bannerUrl" json:"bannerUrl"`
	ParentID               *int           `gorm:"column:parentId;uniqueIndex:Team_slug_parentId_key;index:Team_parentId_idx;size:32" json:"parentId"`
	TimeFormat             *int           `gorm:"column:timeFormat;size:32" json:"timeFormat"`
	TimeZone               *string        `gorm:"column:timeZone;default:'Europe/London';not null" json:"timeZone"`
	WeekStart              *string        `gorm:"column:weekStart;default:'Sunday';not null" json:"weekStart"`
	IsOrganization         bool           `gorm:"column:isOrganization;default:false;not null" json:"isOrganization"`
	PendingPayment         bool           `gorm:"column:pendingPayment;default:false;not null" json:"pendingPayment"`
	IsPlatform             bool           `gorm:"column:isPlatform;default:false;not null" json:"isPlatform"`
	CreatedByOAuthClientID *string        `gorm:"column:createdByOAuthClientId" json:"createdByOAuthClientId"`
	SmsLockState           string         `gorm:"column:smsLockState;default:'UNLOCKED';not null" json:"smsLockState"`
	SmsLockReviewedByAdmin bool           `gorm:"column:smsLockReviewedByAdmin;default:false;not null" json:"smsLockReviewedByAdmin"`
	// @zod.import(["import { intervalLimitsType } from '../../zod-utils'"]).custom.use(intervalLimitsType)
	BookingLimits                datatypes.JSON             `gorm:"column:bookingLimits" json:"bookingLimits"`
	IncludeManagedEventsInLimits bool                       `gorm:"column:includeManagedEventsInLimits;default:false;not null" json:"includeManagedEventsInLimits"`
	AutoOptInFeatures            bool                       `gorm:"column:autoOptInFeatures;default:false;not null" json:"autoOptInFeatures"`
	Parent                       *Team                      `gorm:"foreignKey:ParentID;belongsTo"`
	CreatedByOAuthClient         *PlatformOAuthClient       `gorm:"foreignKey:CreatedByOAuthClientID;belongsTo"`
	Members                      []Membership               `gorm:"foreignKey:TeamID;constraint:Membership_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EventTypes                   []EventType                `gorm:"foreignKey:TeamID;constraint:EventType_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	VerifiedNumbers              []VerifiedNumber           `gorm:"foreignKey:TeamID;constraint:VerifiedNumber_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	VerifiedEmails               []VerifiedEmail            `gorm:"foreignKey:TeamID;constraint:VerifiedEmail_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Children                     []Team                     `gorm:"foreignKey:ParentID;constraint:Team_parentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrgUsers                     []User                     `gorm:"foreignKey:OrganizationID;constraint:users_organizationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	InviteTokens                 []VerificationToken        `gorm:"foreignKey:TeamID;constraint:VerificationToken_teamId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Webhooks                     []Webhook                  `gorm:"foreignKey:TeamID;constraint:Webhook_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeys                      []APIKey                   `gorm:"foreignKey:TeamID;constraint:ApiKey_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Credentials                  []Credential               `gorm:"foreignKey:TeamID;constraint:Credential_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessCodes                  []AccessCode               `gorm:"foreignKey:TeamID;constraint:AccessCode_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	InstantMeetingTokens         []InstantMeetingToken      `gorm:"foreignKey:TeamID;constraint:InstantMeetingToken_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrgProfiles                  []Profile                  `gorm:"foreignKey:OrganizationID;constraint:Profile_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DsyncTeamGroupMapping        []DSyncTeamGroupMapping    `gorm:"foreignKey:TeamID;constraint:DSyncTeamGroupMapping_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlatformOAuthClient          []PlatformOAuthClient      `gorm:"foreignKey:OrganizationID;constraint:PlatformOAuthClient_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attributes                   []Attribute                `gorm:"foreignKey:TeamID;constraint:Attribute_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DelegationCredentials        []DelegationCredential     `gorm:"foreignKey:OrganizationID;constraint:DelegationCredential_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Roles                        []Role                     `gorm:"foreignKey:TeamID;constraint:Role_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CalAiPhoneNumbers            []CalAiPhoneNumber         `gorm:"foreignKey:TeamID;constraint:CalAiPhoneNumber_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Agents                       []Agent                    `gorm:"foreignKey:TeamID;constraint:Agent_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BookingReports               []BookingReport            `gorm:"foreignKey:OrganizationID;constraint:BookingReport_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	WrongAssignmentReports       []WrongAssignmentReport    `gorm:"foreignKey:TeamID;constraint:WrongAssignmentReport_teamId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Features                     []TeamFeatures             `gorm:"foreignKey:TeamID;constraint:TeamFeatures_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	InternalNotePresets          []InternalNotePreset       `gorm:"foreignKey:TeamID;constraint:InternalNotePreset_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ManagedOrganizations         []ManagedOrganization      `gorm:"foreignKey:ManagerOrganizationID;constraint:ManagedOrganization_managerOrganizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FilterSegments               []FilterSegment            `gorm:"foreignKey:TeamID;constraint:FilterSegment_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	IntegrationAttributeSyncs    []IntegrationAttributeSync `gorm:"foreignKey:OrganizationID;constraint:IntegrationAttributeSync_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SeatChangeLogs               []SeatChangeLog            `gorm:"foreignKey:TeamID;constraint:SeatChangeLog_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	MonthlyProrations            []MonthlyProration         `gorm:"foreignKey:TeamID;constraint:MonthlyProration_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationSettings         *OrganizationSettings      `gorm:"foreignKey:OrganizationID;constraint:OrganizationSettings_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlatformBilling              *PlatformBilling           `gorm:"foreignKey:ID;constraint:PlatformBilling_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreditBalance                *CreditBalance             `gorm:"foreignKey:TeamID;constraint:CreditBalance_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationOnboarding       *OrganizationOnboarding    `gorm:"foreignKey:OrganizationID;constraint:OrganizationOnboarding_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ManagedOrganization          *ManagedOrganization       `gorm:"foreignKey:ManagedOrganizationID;constraint:ManagedOrganization_managedOrganizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TeamBilling                  *TeamBilling               `gorm:"foreignKey:TeamID;constraint:TeamBilling_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationBilling          *OrganizationBilling       `gorm:"foreignKey:TeamID;constraint:OrganizationBilling_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type CreditBalance struct {
	ID                string              `gorm:"column:id;primaryKey" json:"id"`
	TeamID            *int                `gorm:"column:teamId;uniqueIndex:CreditBalance_teamId_key;size:32" json:"teamId"`
	UserID            *int                `gorm:"column:userId;uniqueIndex:CreditBalance_userId_key;size:32" json:"userId"`
	AdditionalCredits int                 `gorm:"column:additionalCredits;size:32;default:0;not null" json:"additionalCredits"`
	LimitReachedAt    *DateTime           `gorm:"column:limitReachedAt;type:timestamp(3)" json:"limitReachedAt"`
	WarningSentAt     *DateTime           `gorm:"column:warningSentAt;type:timestamp(3)" json:"warningSentAt"`
	Team              *Team               `gorm:"belongsTo"`
	User              *User               `gorm:"belongsTo"`
	ExpenseLogs       []CreditExpenseLog  `gorm:"foreignKey:CreditBalanceID;constraint:CreditExpenseLog_creditBalanceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PurchaseLogs      []CreditPurchaseLog `gorm:"foreignKey:CreditBalanceID;constraint:CreditPurchaseLog_creditBalanceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (CreditBalance) TableName() string {
	return "CreditBalance"
}

func (m *CreditBalance) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type CreditPurchaseLog struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	CreditBalanceID string         `gorm:"column:creditBalanceId;not null" json:"creditBalanceId"`
	Credits         int            `gorm:"column:credits;size:32;not null" json:"credits"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	CreditBalance   *CreditBalance `gorm:"belongsTo"`
}

func (CreditPurchaseLog) TableName() string {
	return "CreditPurchaseLog"
}

func (m *CreditPurchaseLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type CreditExpenseLog struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	CreditBalanceID string         `gorm:"column:creditBalanceId;not null" json:"creditBalanceId"`
	BookingUID      *string        `gorm:"column:bookingUid" json:"bookingUid"`
	Credits         *int           `gorm:"column:credits;size:32" json:"credits"`
	CreditType      string         `gorm:"column:creditType;not null" json:"creditType"`
	Date            DateTime       `gorm:"column:date;type:timestamp(3);not null" json:"date"`
	SmsSid          *string        `gorm:"column:smsSid" json:"smsSid"`
	SmsSegments     *int           `gorm:"column:smsSegments;size:32" json:"smsSegments"`
	PhoneNumber     *string        `gorm:"column:phoneNumber" json:"phoneNumber"`
	Email           *string        `gorm:"column:email" json:"email"`
	CallDuration    *int           `gorm:"column:callDuration;size:32" json:"callDuration"`
	CreditFor       *string        `gorm:"column:creditFor" json:"creditFor"`
	ExternalRef     *string        `gorm:"column:externalRef;uniqueIndex:CreditExpenseLog_externalRef_key" json:"externalRef"`
	CreditBalance   *CreditBalance `gorm:"belongsTo"`
	Booking         *Booking       `gorm:"references:UID;belongsTo"`
}

func (CreditExpenseLog) TableName() string {
	return "CreditExpenseLog"
}

func (m *CreditExpenseLog) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type OrganizationSettings struct {
	ID                                    int        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	OrganizationID                        int        `gorm:"column:organizationId;uniqueIndex:OrganizationSettings_organizationId_key;size:32;not null" json:"organizationId"`
	IsOrganizationConfigured              bool       `gorm:"column:isOrganizationConfigured;default:false;not null" json:"isOrganizationConfigured"`
	IsOrganizationVerified                bool       `gorm:"column:isOrganizationVerified;default:false;not null" json:"isOrganizationVerified"`
	OrgAutoAcceptEmail                    string     `gorm:"column:orgAutoAcceptEmail;not null" json:"orgAutoAcceptEmail"`
	LockEventTypeCreationForUsers         bool       `gorm:"column:lockEventTypeCreationForUsers;default:false;not null" json:"lockEventTypeCreationForUsers"`
	AdminGetsNoSlotsNotification          bool       `gorm:"column:adminGetsNoSlotsNotification;default:false;not null" json:"adminGetsNoSlotsNotification"`
	IsAdminReviewed                       bool       `gorm:"column:isAdminReviewed;default:false;not null" json:"isAdminReviewed"`
	IsAdminAPIEnabled                     bool       `gorm:"column:isAdminAPIEnabled;default:false;not null" json:"isAdminAPIEnabled"`
	AllowSEOIndexing                      bool       `gorm:"column:allowSEOIndexing;default:false;not null" json:"allowSEOIndexing"`
	OrgProfileRedirectsToVerifiedDomain   bool       `gorm:"column:orgProfileRedirectsToVerifiedDomain;default:false;not null" json:"orgProfileRedirectsToVerifiedDomain"`
	DisablePhoneOnlySMSNotifications      bool       `gorm:"column:disablePhoneOnlySMSNotifications;default:false;not null" json:"disablePhoneOnlySMSNotifications"`
	DisableAutofillOnBookingPage          bool       `gorm:"column:disableAutofillOnBookingPage;default:false;not null" json:"disableAutofillOnBookingPage"`
	OrgAutoJoinOnSignup                   *bool      `gorm:"column:orgAutoJoinOnSignup;default:true;not null" json:"orgAutoJoinOnSignup"`
	DisableAttendeeConfirmationEmail      bool       `gorm:"column:disableAttendeeConfirmationEmail;default:false;not null" json:"disableAttendeeConfirmationEmail"`
	DisableAttendeeCancellationEmail      bool       `gorm:"column:disableAttendeeCancellationEmail;default:false;not null" json:"disableAttendeeCancellationEmail"`
	DisableAttendeeRescheduledEmail       bool       `gorm:"column:disableAttendeeRescheduledEmail;default:false;not null" json:"disableAttendeeRescheduledEmail"`
	DisableAttendeeRequestEmail           bool       `gorm:"column:disableAttendeeRequestEmail;default:false;not null" json:"disableAttendeeRequestEmail"`
	DisableAttendeeReassignedEmail        bool       `gorm:"column:disableAttendeeReassignedEmail;default:false;not null" json:"disableAttendeeReassignedEmail"`
	DisableAttendeeAwaitingPaymentEmail   bool       `gorm:"column:disableAttendeeAwaitingPaymentEmail;default:false;not null" json:"disableAttendeeAwaitingPaymentEmail"`
	DisableAttendeeRescheduleRequestEmail bool       `gorm:"column:disableAttendeeRescheduleRequestEmail;default:false;not null" json:"disableAttendeeRescheduleRequestEmail"`
	DisableAttendeeLocationChangeEmail    bool       `gorm:"column:disableAttendeeLocationChangeEmail;default:false;not null" json:"disableAttendeeLocationChangeEmail"`
	DisableAttendeeNewEventEmail          bool       `gorm:"column:disableAttendeeNewEventEmail;default:false;not null" json:"disableAttendeeNewEventEmail"`
	Organization                          *Team      `gorm:"foreignKey:OrganizationID;belongsTo"`
	DSyncData                             *DSyncData `gorm:"foreignKey:OrganizationID;references:OrganizationID;constraint:DSyncData_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OrganizationSettings) TableName() string {
	return "OrganizationSettings"
}

type Membership struct {
	ID              int               `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	TeamID          int               `gorm:"column:teamId;uniqueIndex:Membership_userId_teamId_key;index:Membership_teamId_idx;size:32;not null" json:"teamId"`
	UserID          int               `gorm:"column:userId;uniqueIndex:Membership_userId_teamId_key;index:Membership_userId_idx;size:32;not null" json:"userId"`
	Accepted        bool              `gorm:"column:accepted;index:Membership_accepted_idx;default:false;not null" json:"accepted"`
	Role            string            `gorm:"column:role;index:Membership_role_idx;not null" json:"role"`
	CustomRoleID    *string           `gorm:"column:customRoleId;index:Membership_customRoleId_idx" json:"customRoleId"`
	CreatedAt       *DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false" json:"createdAt"`
	UpdatedAt       *DateTime         `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	CustomRole      *Role             `gorm:"foreignKey:CustomRoleID;belongsTo"`
	Team            *Team             `gorm:"belongsTo"`
	User            *User             `gorm:"belongsTo"`
	AttributeToUser []AttributeToUser `gorm:"foreignKey:MemberID;constraint:AttributeToUser_memberId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Host            []Host            `gorm:"foreignKey:MemberID;constraint:Host_memberId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Membership) TableName() string {
	return "Membership"
}

func (m *Membership) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt == nil {
		m.CreatedAt = &DateTime{Time: now}
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	return nil
}

func (*Membership) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type VerificationToken struct {
	ID               int             `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Identifier       string          `gorm:"column:identifier;uniqueIndex:VerificationToken_identifier_token_key;not null" json:"identifier"`
	Token            string          `gorm:"column:token;uniqueIndex:VerificationToken_token_key;uniqueIndex:VerificationToken_identifier_token_key;index:VerificationToken_token_idx;not null" json:"token"`
	Expires          DateTime        `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	ExpiresInDays    *int            `gorm:"column:expiresInDays;size:32" json:"expiresInDays"`
	CreatedAt        DateTime        `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime        `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	TeamID           *int            `gorm:"column:teamId;index:VerificationToken_teamId_idx;size:32" json:"teamId"`
	SecondaryEmailID *int            `gorm:"column:secondaryEmailId;index:VerificationToken_secondaryEmailId_idx;size:32" json:"secondaryEmailId"`
	Team             *Team           `gorm:"belongsTo"`
	SecondaryEmail   *SecondaryEmail `gorm:"belongsTo"`
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

type InstantMeetingToken struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Token     string   `gorm:"column:token;uniqueIndex:InstantMeetingToken_token_key;index:InstantMeetingToken_token_idx;not null" json:"token"`
	Expires   DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	TeamID    int      `gorm:"column:teamId;size:32;not null" json:"teamId"`
	BookingID *int     `gorm:"column:bookingId;uniqueIndex:InstantMeetingToken_bookingId_key;size:32" json:"bookingId"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team      *Team    `gorm:"belongsTo"`
	Booking   *Booking `gorm:"belongsTo"`
}

func (InstantMeetingToken) TableName() string {
	return "InstantMeetingToken"
}

func (m *InstantMeetingToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*InstantMeetingToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BookingReference struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	// @zod.string.min(1)
	Type string `gorm:"column:type;index:BookingReference_type_idx;not null" json:"type"`
	// @zod.string.min(1)
	UID                        string                `gorm:"column:uid;index:BookingReference_uid_idx;not null" json:"uid"`
	MeetingID                  *string               `gorm:"column:meetingId" json:"meetingId"`
	ThirdPartyRecurringEventID *string               `gorm:"column:thirdPartyRecurringEventId" json:"thirdPartyRecurringEventId"`
	MeetingPassword            *string               `gorm:"column:meetingPassword" json:"meetingPassword"`
	MeetingURL                 *string               `gorm:"column:meetingUrl" json:"meetingUrl"`
	BookingID                  *int                  `gorm:"column:bookingId;index:BookingReference_bookingId_idx;size:32" json:"bookingId"`
	ExternalCalendarID         *string               `gorm:"column:externalCalendarId" json:"externalCalendarId"`
	Deleted                    *bool                 `gorm:"column:deleted" json:"deleted"`
	CredentialID               *int                  `gorm:"column:credentialId;size:32" json:"credentialId"`
	DelegationCredentialID     *string               `gorm:"column:delegationCredentialId" json:"delegationCredentialId"`
	Booking                    *Booking              `gorm:"belongsTo"`
	Credential                 *Credential           `gorm:"belongsTo"`
	DelegationCredential       *DelegationCredential `gorm:"belongsTo"`
}

func (BookingReference) TableName() string {
	return "BookingReference"
}

type Attendee struct {
	ID          int          `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Email       string       `gorm:"column:email;index:Attendee_email_idx;index:Attendee_email_bookingId_idx;not null" json:"email"`
	Name        string       `gorm:"column:name;not null" json:"name"`
	TimeZone    string       `gorm:"column:timeZone;not null" json:"timeZone"`
	PhoneNumber *string      `gorm:"column:phoneNumber" json:"phoneNumber"`
	Locale      *string      `gorm:"column:locale;default:'en'" json:"locale"`
	BookingID   *int         `gorm:"column:bookingId;index:Attendee_bookingId_idx;index:Attendee_email_bookingId_idx;size:32" json:"bookingId"`
	NoShow      *bool        `gorm:"column:noShow;default:false" json:"noShow"`
	Booking     *Booking     `gorm:"belongsTo"`
	BookingSeat *BookingSeat `gorm:"foreignKey:AttendeeID;constraint:BookingSeat_attendeeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Attendee) TableName() string {
	return "Attendee"
}

type Booking struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UID            string  `gorm:"column:uid;uniqueIndex:Booking_uid_key;index:Booking_uid_idx;not null" json:"uid"`
	IdempotencyKey *string `gorm:"column:idempotencyKey;uniqueIndex:Booking_idempotencyKey_key" json:"idempotencyKey"`
	UserID         *int    `gorm:"column:userId;index:Booking_userId_idx;index:Booking_userId_endTime_idx;index:Booking_userId_status_startTime_idx;index:Booking_userId_createdAt_idx;size:32" json:"userId"`
	// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
	UserPrimaryEmail *string        `gorm:"column:userPrimaryEmail" json:"userPrimaryEmail"`
	EventTypeID      *int           `gorm:"column:eventTypeId;index:Booking_eventTypeId_idx;index:Booking_eventTypeId_status_idx;size:32" json:"eventTypeId"`
	Title            string         `gorm:"column:title;not null" json:"title"`
	Description      *string        `gorm:"column:description" json:"description"`
	CustomInputs     datatypes.JSON `gorm:"column:customInputs" json:"customInputs"`
	// @zod.import(["import { bookingResponses } from '../../zod-utils'"]).custom.use(bookingResponses)
	Responses             datatypes.JSON     `gorm:"column:responses" json:"responses"`
	StartTime             DateTime           `gorm:"column:startTime;index:Booking_startTime_endTime_status_idx;index:Booking_userId_status_startTime_idx;type:timestamp(3);not null" json:"startTime"`
	EndTime               DateTime           `gorm:"column:endTime;index:Booking_startTime_endTime_status_idx;index:Booking_userId_endTime_idx;type:timestamp(3);not null" json:"endTime"`
	Location              *string            `gorm:"column:location" json:"location"`
	CreatedAt             DateTime           `gorm:"column:createdAt;index:Booking_userId_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             *DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	Status                string             `gorm:"column:status;index:Booking_status_idx;index:Booking_startTime_endTime_status_idx;index:Booking_userId_status_startTime_idx;index:Booking_eventTypeId_status_idx;default:'accepted';not null" json:"status"`
	Paid                  bool               `gorm:"column:paid;default:false;not null" json:"paid"`
	DestinationCalendarID *int               `gorm:"column:destinationCalendarId;index:Booking_destinationCalendarId_idx;size:32" json:"destinationCalendarId"`
	CancellationReason    *string            `gorm:"column:cancellationReason" json:"cancellationReason"`
	RejectionReason       *string            `gorm:"column:rejectionReason" json:"rejectionReason"`
	ReassignReason        *string            `gorm:"column:reassignReason" json:"reassignReason"`
	ReassignByID          *int               `gorm:"column:reassignById;size:32" json:"reassignById"`
	DynamicEventSlugRef   *string            `gorm:"column:dynamicEventSlugRef" json:"dynamicEventSlugRef"`
	DynamicGroupSlugRef   *string            `gorm:"column:dynamicGroupSlugRef" json:"dynamicGroupSlugRef"`
	Rescheduled           *bool              `gorm:"column:rescheduled" json:"rescheduled"`
	FromReschedule        *string            `gorm:"column:fromReschedule;index:Booking_fromReschedule_idx" json:"fromReschedule"`
	RecurringEventID      *string            `gorm:"column:recurringEventId;index:Booking_recurringEventId_idx" json:"recurringEventId"`
	SmsReminderNumber     *string            `gorm:"column:smsReminderNumber" json:"smsReminderNumber"`
	ScheduledJobs         ScalarList[string] `gorm:"column:scheduledJobs;type:text[]" json:"scheduledJobs"`
	// @zod.import(["import { bookingMetadataSchema } from '../../zod-utils'"]).custom.use(bookingMetadataSchema)
	Metadata        datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	IsRecorded      bool           `gorm:"column:isRecorded;default:false;not null" json:"isRecorded"`
	ICalUID         *string        `gorm:"column:iCalUID;default:''" json:"iCalUID"`
	ICalSequence    int            `gorm:"column:iCalSequence;size:32;default:0;not null" json:"iCalSequence"`
	Rating          *int           `gorm:"column:rating;size:32" json:"rating"`
	RatingFeedback  *string        `gorm:"column:ratingFeedback" json:"ratingFeedback"`
	NoShowHost      *bool          `gorm:"column:noShowHost;default:false" json:"noShowHost"`
	OneTimePassword *string        `gorm:"column:oneTimePassword;uniqueIndex:Booking_oneTimePassword_key" json:"oneTimePassword"`
	// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
	CancelledBy *string `gorm:"column:cancelledBy" json:"cancelledBy"`
	// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
	RescheduledBy          *string                    `gorm:"column:rescheduledBy" json:"rescheduledBy"`
	CreationSource         *string                    `gorm:"column:creationSource" json:"creationSource"`
	User                   *User                      `gorm:"foreignKey:UserID;belongsTo"`
	EventType              *EventType                 `gorm:"belongsTo"`
	DestinationCalendar    *DestinationCalendar       `gorm:"belongsTo"`
	ReassignBy             *User                      `gorm:"foreignKey:ReassignByID;belongsTo"`
	References             []BookingReference         `gorm:"foreignKey:BookingID;constraint:BookingReference_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attendees              []Attendee                 `gorm:"foreignKey:BookingID;constraint:Attendee_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Payment                []Payment                  `gorm:"foreignKey:BookingID;constraint:Payment_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SeatsReferences        []BookingSeat              `gorm:"foreignKey:BookingID;constraint:BookingSeat_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ScheduledTriggers      []WebhookScheduledTriggers `gorm:"foreignKey:BookingID;constraint:WebhookScheduledTriggers_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AssignmentReason       []AssignmentReason         `gorm:"foreignKey:BookingID;constraint:AssignmentReason_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	InternalNote           []BookingInternalNote      `gorm:"foreignKey:BookingID;constraint:BookingInternalNote_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExpenseLogs            []CreditExpenseLog         `gorm:"foreignKey:BookingUID;references:UID;constraint:CreditExpenseLog_bookingUid_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	WrongAssignmentReports []WrongAssignmentReport    `gorm:"foreignKey:BookingUID;references:UID;constraint:WrongAssignmentReport_bookingUid_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	InstantMeetingToken    *InstantMeetingToken       `gorm:"foreignKey:BookingID;constraint:InstantMeetingToken_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tracking               *Tracking                  `gorm:"foreignKey:BookingID;constraint:Tracking_bookingId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Report                 *BookingReport             `gorm:"foreignKey:BookingUID;references:UID;constraint:BookingReport_bookingUid_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Booking) TableName() string {
	return "Booking"
}

func (m *Booking) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	if m.OneTimePassword == nil {
		generated := uuid.NewString()
		m.OneTimePassword = &generated
	}
	return nil
}

func (*Booking) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type Tracking struct {
	ID          int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	BookingID   int      `gorm:"column:bookingId;uniqueIndex:Tracking_bookingId_key;size:32;not null" json:"bookingId"`
	UtmSource   *string  `gorm:"column:utm_source" json:"utm_source"`
	UtmMedium   *string  `gorm:"column:utm_medium" json:"utm_medium"`
	UtmCampaign *string  `gorm:"column:utm_campaign" json:"utm_campaign"`
	UtmTerm     *string  `gorm:"column:utm_term" json:"utm_term"`
	UtmContent  *string  `gorm:"column:utm_content" json:"utm_content"`
	Booking     *Booking `gorm:"belongsTo"`
}

func (Tracking) TableName() string {
	return "Tracking"
}

type Schedule struct {
	ID                   int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID               int            `gorm:"column:userId;index:Schedule_userId_idx;size:32;not null" json:"userId"`
	Name                 string         `gorm:"column:name;not null" json:"name"`
	TimeZone             *string        `gorm:"column:timeZone" json:"timeZone"`
	User                 *User          `gorm:"belongsTo"`
	EventType            []EventType    `gorm:"foreignKey:ScheduleID;constraint:EventType_scheduleId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	InstantMeetingEvents []EventType    `gorm:"foreignKey:InstantMeetingScheduleID;constraint:EventType_instantMeetingScheduleId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	RestrictionSchedule  []EventType    `gorm:"foreignKey:RestrictionScheduleID;constraint:EventType_restrictionScheduleId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Availability         []Availability `gorm:"foreignKey:ScheduleID;constraint:Availability_scheduleId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Host                 []Host         `gorm:"foreignKey:ScheduleID;constraint:Host_scheduleId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Schedule) TableName() string {
	return "Schedule"
}

type Availability struct {
	ID          int             `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID      *int            `gorm:"column:userId;index:Availability_userId_idx;size:32" json:"userId"`
	EventTypeID *int            `gorm:"column:eventTypeId;index:Availability_eventTypeId_idx;size:32" json:"eventTypeId"`
	Days        ScalarList[int] `gorm:"column:days;type:integer[]" json:"days"`
	StartTime   TimeOfDay       `gorm:"column:startTime;type:time;not null" json:"startTime"`
	EndTime     TimeOfDay       `gorm:"column:endTime;type:time;not null" json:"endTime"`
	Date        *Date           `gorm:"column:date;type:date" json:"date"`
	ScheduleID  *int            `gorm:"column:scheduleId;index:Availability_scheduleId_idx;size:32" json:"scheduleId"`
	User        *User           `gorm:"belongsTo"`
	EventType   *EventType      `gorm:"belongsTo"`
	Schedule    *Schedule       `gorm:"belongsTo"`
}

func (Availability) TableName() string {
	return "Availability"
}

type SelectedCalendar struct {
	ID                       string                `gorm:"column:id;primaryKey" json:"id"`
	UserID                   int                   `gorm:"column:userId;uniqueIndex:SelectedCalendar_userId_integration_externalId_eventTypeId_key;index:SelectedCalendar_userId_idx;size:32;not null" json:"userId"`
	Integration              string                `gorm:"column:integration;uniqueIndex:SelectedCalendar_userId_integration_externalId_eventTypeId_key;index:SelectedCalendar_watch_idx;index:SelectedCalendar_unwatch_idx;not null" json:"integration"`
	ExternalID               string                `gorm:"column:externalId;uniqueIndex:SelectedCalendar_userId_integration_externalId_eventTypeId_key;index:SelectedCalendar_externalId_idx;not null" json:"externalId"`
	CredentialID             *int                  `gorm:"column:credentialId;index:SelectedCalendar_credentialId_idx;size:32" json:"credentialId"`
	CreatedAt                *DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false" json:"createdAt"`
	UpdatedAt                *DateTime             `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	GoogleChannelID          *string               `gorm:"column:googleChannelId;uniqueIndex:SelectedCalendar_googleChannelId_eventTypeId_key" json:"googleChannelId"`
	GoogleChannelKind        *string               `gorm:"column:googleChannelKind" json:"googleChannelKind"`
	GoogleChannelResourceID  *string               `gorm:"column:googleChannelResourceId" json:"googleChannelResourceId"`
	GoogleChannelResourceURI *string               `gorm:"column:googleChannelResourceUri" json:"googleChannelResourceUri"`
	GoogleChannelExpiration  *string               `gorm:"column:googleChannelExpiration;index:SelectedCalendar_watch_idx;index:SelectedCalendar_unwatch_idx" json:"googleChannelExpiration"`
	ChannelID                *string               `gorm:"column:channelId;index:SelectedCalendar_channelId_idx" json:"channelId"`
	ChannelKind              *string               `gorm:"column:channelKind" json:"channelKind"`
	ChannelResourceID        *string               `gorm:"column:channelResourceId" json:"channelResourceId"`
	ChannelResourceURI       *string               `gorm:"column:channelResourceUri" json:"channelResourceUri"`
	ChannelExpiration        *DateTime             `gorm:"column:channelExpiration;type:timestamp(3)" json:"channelExpiration"`
	SyncSubscribedAt         *DateTime             `gorm:"column:syncSubscribedAt;type:timestamp(3)" json:"syncSubscribedAt"`
	SyncSubscribedErrorAt    *DateTime             `gorm:"column:syncSubscribedErrorAt;type:timestamp(3)" json:"syncSubscribedErrorAt"`
	SyncSubscribedErrorCount int                   `gorm:"column:syncSubscribedErrorCount;size:32;default:0;not null" json:"syncSubscribedErrorCount"`
	SyncToken                *string               `gorm:"column:syncToken" json:"syncToken"`
	SyncedAt                 *DateTime             `gorm:"column:syncedAt;type:timestamp(3)" json:"syncedAt"`
	SyncErrorAt              *DateTime             `gorm:"column:syncErrorAt;type:timestamp(3)" json:"syncErrorAt"`
	SyncErrorCount           *int                  `gorm:"column:syncErrorCount;size:32;default:0" json:"syncErrorCount"`
	DelegationCredentialID   *string               `gorm:"column:delegationCredentialId" json:"delegationCredentialId"`
	Error                    *string               `gorm:"column:error;index:SelectedCalendar_watch_idx;index:SelectedCalendar_unwatch_idx" json:"error"`
	LastErrorAt              *DateTime             `gorm:"column:lastErrorAt;type:timestamp(3)" json:"lastErrorAt"`
	WatchAttempts            int                   `gorm:"column:watchAttempts;index:SelectedCalendar_watch_idx;size:32;default:0;not null" json:"watchAttempts"`
	UnwatchAttempts          int                   `gorm:"column:unwatchAttempts;index:SelectedCalendar_unwatch_idx;size:32;default:0;not null" json:"unwatchAttempts"`
	MaxAttempts              *int                  `gorm:"column:maxAttempts;index:SelectedCalendar_watch_idx;index:SelectedCalendar_unwatch_idx;size:32;default:3;not null" json:"maxAttempts"`
	EventTypeID              *int                  `gorm:"column:eventTypeId;uniqueIndex:SelectedCalendar_userId_integration_externalId_eventTypeId_key;uniqueIndex:SelectedCalendar_googleChannelId_eventTypeId_key;index:SelectedCalendar_eventTypeId_idx;size:32" json:"eventTypeId"`
	User                     *User                 `gorm:"belongsTo"`
	Credential               *Credential           `gorm:"belongsTo"`
	DelegationCredential     *DelegationCredential `gorm:"belongsTo"`
	EventType                *EventType            `gorm:"belongsTo"`
	CalendarCacheEvents      []CalendarCacheEvent  `gorm:"foreignKey:SelectedCalendarID;constraint:CalendarCacheEvent_selectedCalendarId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (SelectedCalendar) TableName() string {
	return "SelectedCalendar"
}

func (m *SelectedCalendar) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt == nil {
		m.CreatedAt = &DateTime{Time: now}
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	return nil
}

func (*SelectedCalendar) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type EventTypeCustomInput struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	EventTypeID int    `gorm:"column:eventTypeId;index:EventTypeCustomInput_eventTypeId_idx;size:32;not null" json:"eventTypeId"`
	Label       string `gorm:"column:label;not null" json:"label"`
	Type        string `gorm:"column:type;not null" json:"type"`
	// @zod.import(["import { customInputOptionSchema } from '../../zod-utils'"]).custom.use(customInputOptionSchema)
	Options     datatypes.JSON `gorm:"column:options" json:"options"`
	Required    bool           `gorm:"column:required;not null" json:"required"`
	Placeholder string         `gorm:"column:placeholder;default:'';not null" json:"placeholder"`
	EventType   *EventType     `gorm:"belongsTo"`
}

func (EventTypeCustomInput) TableName() string {
	return "EventTypeCustomInput"
}

type ResetPasswordRequest struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Email     string   `gorm:"column:email;not null" json:"email"`
	Expires   DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
}

func (ResetPasswordRequest) TableName() string {
	return "ResetPasswordRequest"
}

func (m *ResetPasswordRequest) BeforeCreate(tx *gorm.DB) error {
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

func (*ResetPasswordRequest) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ReminderMail struct {
	ID             int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	ReferenceID    int      `gorm:"column:referenceId;index:ReminderMail_referenceId_idx;size:32;not null" json:"referenceId"`
	ReminderType   string   `gorm:"column:reminderType;index:ReminderMail_reminderType_idx;not null" json:"reminderType"`
	ElapsedMinutes int      `gorm:"column:elapsedMinutes;size:32;not null" json:"elapsedMinutes"`
	CreatedAt      DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (ReminderMail) TableName() string {
	return "ReminderMail"
}

func (m *ReminderMail) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Payment struct {
	ID            int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UID           string         `gorm:"column:uid;uniqueIndex:Payment_uid_key;not null" json:"uid"`
	AppID         *string        `gorm:"column:appId" json:"appId"`
	BookingID     int            `gorm:"column:bookingId;index:Payment_bookingId_idx;size:32;not null" json:"bookingId"`
	Amount        int            `gorm:"column:amount;size:32;not null" json:"amount"`
	Fee           int            `gorm:"column:fee;size:32;not null" json:"fee"`
	Currency      string         `gorm:"column:currency;not null" json:"currency"`
	Success       bool           `gorm:"column:success;not null" json:"success"`
	Refunded      bool           `gorm:"column:refunded;not null" json:"refunded"`
	Data          datatypes.JSON `gorm:"column:data;not null" json:"data"`
	ExternalID    string         `gorm:"column:externalId;uniqueIndex:Payment_externalId_key;index:Payment_externalId_idx;not null" json:"externalId"`
	PaymentOption *string        `gorm:"column:paymentOption;default:'ON_BOOKING'" json:"paymentOption"`
	App           *App           `gorm:"foreignKey:AppID;references:Slug;belongsTo"`
	Booking       *Booking       `gorm:"belongsTo"`
}

func (Payment) TableName() string {
	return "Payment"
}

type Webhook struct {
	ID                    string  `gorm:"column:id;primaryKey" json:"id"`
	UserID                *int    `gorm:"column:userId;uniqueIndex:Webhook_userId_subscriberUrl_key;size:32" json:"userId"`
	TeamID                *int    `gorm:"column:teamId;size:32" json:"teamId"`
	EventTypeID           *int    `gorm:"column:eventTypeId;size:32" json:"eventTypeId"`
	PlatformOAuthClientID *string `gorm:"column:platformOAuthClientId;uniqueIndex:Webhook_platformOAuthClientId_subscriberUrl_key" json:"platformOAuthClientId"`
	// @zod.string.url()
	SubscriberURL       string                     `gorm:"column:subscriberUrl;uniqueIndex:Webhook_userId_subscriberUrl_key;uniqueIndex:Webhook_platformOAuthClientId_subscriberUrl_key;not null" json:"subscriberUrl"`
	PayloadTemplate     *string                    `gorm:"column:payloadTemplate" json:"payloadTemplate"`
	CreatedAt           DateTime                   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Active              *bool                      `gorm:"column:active;index:Webhook_active_idx;default:true;not null" json:"active"`
	EventTriggers       ScalarList[string]         `gorm:"column:eventTriggers;type:text[]" json:"eventTriggers"`
	AppID               *string                    `gorm:"column:appId" json:"appId"`
	Secret              *string                    `gorm:"column:secret" json:"secret"`
	Platform            bool                       `gorm:"column:platform;default:false;not null" json:"platform"`
	Time                *int                       `gorm:"column:time;size:32" json:"time"`
	TimeUnit            *string                    `gorm:"column:timeUnit" json:"timeUnit"`
	Version             *string                    `gorm:"column:version;default:'2021-10-20';not null" json:"version"`
	User                *User                      `gorm:"belongsTo"`
	Team                *Team                      `gorm:"belongsTo"`
	EventType           *EventType                 `gorm:"belongsTo"`
	PlatformOAuthClient *PlatformOAuthClient       `gorm:"belongsTo"`
	App                 *App                       `gorm:"foreignKey:AppID;references:Slug;belongsTo"`
	ScheduledTriggers   []WebhookScheduledTriggers `gorm:"foreignKey:WebhookID;constraint:WebhookScheduledTriggers_webhookId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Webhook) TableName() string {
	return "Webhook"
}

func (m *Webhook) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type APIKey struct {
	ID         string      `gorm:"column:id;primaryKey" json:"id"`
	UserID     int         `gorm:"column:userId;index:ApiKey_userId_idx;size:32;not null" json:"userId"`
	TeamID     *int        `gorm:"column:teamId;size:32" json:"teamId"`
	Note       *string     `gorm:"column:note" json:"note"`
	CreatedAt  DateTime    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ExpiresAt  *DateTime   `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	LastUsedAt *DateTime   `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	HashedKey  string      `gorm:"column:hashedKey;uniqueIndex:ApiKey_hashedKey_key;not null" json:"hashedKey"`
	AppID      *string     `gorm:"column:appId" json:"appId"`
	User       *User       `gorm:"belongsTo"`
	Team       *Team       `gorm:"belongsTo"`
	App        *App        `gorm:"foreignKey:AppID;references:Slug;belongsTo"`
	RateLimits []RateLimit `gorm:"foreignKey:APIKeyID;constraint:RateLimit_apiKeyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (APIKey) TableName() string {
	return "ApiKey"
}

func (m *APIKey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type RateLimit struct {
	ID            string   `gorm:"column:id;primaryKey" json:"id"`
	Name          string   `gorm:"column:name;not null" json:"name"`
	APIKeyID      string   `gorm:"column:apiKeyId;index:RateLimit_apiKeyId_idx;not null" json:"apiKeyId"`
	TTL           int      `gorm:"column:ttl;size:32;not null" json:"ttl"`
	Limit         int      `gorm:"column:limit;size:32;not null" json:"limit"`
	BlockDuration int      `gorm:"column:blockDuration;size:32;not null" json:"blockDuration"`
	CreatedAt     DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	APIKey        *APIKey  `gorm:"belongsTo"`
}

func (RateLimit) TableName() string {
	return "RateLimit"
}

func (m *RateLimit) BeforeCreate(tx *gorm.DB) error {
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

func (*RateLimit) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type HashedLink struct {
	ID            int        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Link          string     `gorm:"column:link;uniqueIndex:HashedLink_link_key;not null" json:"link"`
	EventTypeID   int        `gorm:"column:eventTypeId;index:HashedLink_eventTypeId_idx;size:32;not null" json:"eventTypeId"`
	ExpiresAt     *DateTime  `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	MaxUsageCount *int       `gorm:"column:maxUsageCount;size:32;default:1;not null" json:"maxUsageCount"`
	UsageCount    int        `gorm:"column:usageCount;size:32;default:0;not null" json:"usageCount"`
	EventType     *EventType `gorm:"belongsTo"`
}

func (HashedLink) TableName() string {
	return "HashedLink"
}

type Account struct {
	ID                string  `gorm:"column:id;primaryKey" json:"id"`
	UserID            int     `gorm:"column:userId;index:Account_userId_idx;size:32;not null" json:"userId"`
	Type              string  `gorm:"column:type;index:Account_type_idx;not null" json:"type"`
	Provider          string  `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key;not null" json:"provider"`
	ProviderAccountID string  `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key;not null" json:"providerAccountId"`
	ProviderEmail     *string `gorm:"column:providerEmail" json:"providerEmail"`
	RefreshToken      *string `gorm:"column:refresh_token;type:text" json:"refresh_token"`
	AccessToken       *string `gorm:"column:access_token;type:text" json:"access_token"`
	ExpiresAt         *int    `gorm:"column:expires_at;size:32" json:"expires_at"`
	TokenType         *string `gorm:"column:token_type" json:"token_type"`
	Scope             *string `gorm:"column:scope" json:"scope"`
	IDToken           *string `gorm:"column:id_token;type:text" json:"id_token"`
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

type Session struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	SessionToken string   `gorm:"column:sessionToken;uniqueIndex:Session_sessionToken_key;not null" json:"sessionToken"`
	UserID       int      `gorm:"column:userId;index:Session_userId_idx;size:32;not null" json:"userId"`
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

type App struct {
	Slug        string             `gorm:"column:slug;primaryKey" json:"slug"`
	DirName     string             `gorm:"column:dirName;uniqueIndex:App_dirName_key;not null" json:"dirName"`
	Keys        datatypes.JSON     `gorm:"column:keys" json:"keys"`
	Categories  ScalarList[string] `gorm:"column:categories;type:text[]" json:"categories"`
	CreatedAt   DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Enabled     bool               `gorm:"column:enabled;index:App_enabled_idx;default:false;not null" json:"enabled"`
	Credentials []Credential       `gorm:"foreignKey:AppID;references:Slug;constraint:Credential_appId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Payments    []Payment          `gorm:"foreignKey:AppID;references:Slug;constraint:Payment_appId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhook     []Webhook          `gorm:"foreignKey:AppID;references:Slug;constraint:Webhook_appId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKey      []APIKey           `gorm:"foreignKey:AppID;references:Slug;constraint:ApiKey_appId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (App) TableName() string {
	return "App"
}

func (m *App) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*App) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Feedback struct {
	ID      int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Date    DateTime `gorm:"column:date;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"date"`
	UserID  int      `gorm:"column:userId;index:Feedback_userId_idx;size:32;not null" json:"userId"`
	Rating  string   `gorm:"column:rating;index:Feedback_rating_idx;not null" json:"rating"`
	Comment *string  `gorm:"column:comment" json:"comment"`
	User    *User    `gorm:"belongsTo"`
}

func (Feedback) TableName() string {
	return "Feedback"
}

func (m *Feedback) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Date.IsZero() {
		m.Date = DateTime{Time: now}
	}
	return nil
}

type Deployment struct {
	// This is a single row table, so we use a fixed id
	ID   int     `gorm:"column:id;primaryKey;size:32" json:"id"`
	Logo *string `gorm:"column:logo" json:"logo"`
	// @zod.import(["import { DeploymentTheme } from '../../zod-utils'"]).custom.use(DeploymentTheme)
	Theme                   datatypes.JSON `gorm:"column:theme" json:"theme"`
	LicenseKey              *string        `gorm:"column:licenseKey" json:"licenseKey"`
	SignatureTokenEncrypted *string        `gorm:"column:signatureTokenEncrypted" json:"signatureTokenEncrypted"`
	AgreedLicenseAt         *DateTime      `gorm:"column:agreedLicenseAt;type:timestamp(3)" json:"agreedLicenseAt"`
}

func (Deployment) TableName() string {
	return "Deployment"
}

type WebhookScheduledTriggers struct {
	ID            int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	JobName       *string   `gorm:"column:jobName" json:"jobName"`
	SubscriberURL string    `gorm:"column:subscriberUrl;not null" json:"subscriberUrl"`
	Payload       string    `gorm:"column:payload;not null" json:"payload"`
	StartAfter    DateTime  `gorm:"column:startAfter;type:timestamp(3);not null" json:"startAfter"`
	RetryCount    int       `gorm:"column:retryCount;size:32;default:0;not null" json:"retryCount"`
	CreatedAt     *DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false" json:"createdAt"`
	AppID         *string   `gorm:"column:appId" json:"appId"`
	WebhookID     *string   `gorm:"column:webhookId" json:"webhookId"`
	BookingID     *int      `gorm:"column:bookingId;size:32" json:"bookingId"`
	Webhook       *Webhook  `gorm:"belongsTo"`
	Booking       *Booking  `gorm:"belongsTo"`
}

func (WebhookScheduledTriggers) TableName() string {
	return "WebhookScheduledTriggers"
}

func (m *WebhookScheduledTriggers) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt == nil {
		m.CreatedAt = &DateTime{Time: now}
	}
	return nil
}

type BookingSeat struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	ReferenceUID string `gorm:"column:referenceUid;uniqueIndex:BookingSeat_referenceUid_key;not null" json:"referenceUid"`
	BookingID    int    `gorm:"column:bookingId;index:BookingSeat_bookingId_idx;size:32;not null" json:"bookingId"`
	AttendeeID   int    `gorm:"column:attendeeId;uniqueIndex:BookingSeat_attendeeId_key;index:BookingSeat_attendeeId_idx;size:32;not null" json:"attendeeId"`
	// @zod.import(["import { bookingSeatDataSchema } from '../../zod-utils'"]).custom.use(bookingSeatDataSchema)
	Data     datatypes.JSON `gorm:"column:data" json:"data"`
	Metadata datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	Booking  *Booking       `gorm:"belongsTo"`
	Attendee *Attendee      `gorm:"belongsTo"`
}

func (BookingSeat) TableName() string {
	return "BookingSeat"
}

type VerifiedNumber struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID      *int   `gorm:"column:userId;index:VerifiedNumber_userId_idx;size:32" json:"userId"`
	TeamID      *int   `gorm:"column:teamId;index:VerifiedNumber_teamId_idx;size:32" json:"teamId"`
	PhoneNumber string `gorm:"column:phoneNumber;not null" json:"phoneNumber"`
	User        *User  `gorm:"belongsTo"`
	Team        *Team  `gorm:"belongsTo"`
}

func (VerifiedNumber) TableName() string {
	return "VerifiedNumber"
}

type VerifiedEmail struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID *int   `gorm:"column:userId;index:VerifiedEmail_userId_idx;size:32" json:"userId"`
	TeamID *int   `gorm:"column:teamId;index:VerifiedEmail_teamId_idx;size:32" json:"teamId"`
	Email  string `gorm:"column:email;not null" json:"email"`
	User   *User  `gorm:"belongsTo"`
	Team   *Team  `gorm:"belongsTo"`
}

func (VerifiedEmail) TableName() string {
	return "VerifiedEmail"
}

type Feature struct {
	Slug        string         `gorm:"column:slug;primaryKey" json:"slug"`
	Enabled     bool           `gorm:"column:enabled;index:Feature_enabled_idx;default:false;not null" json:"enabled"`
	Description *string        `gorm:"column:description" json:"description"`
	Type        *string        `gorm:"column:type;default:'RELEASE'" json:"type"`
	Stale       *bool          `gorm:"column:stale;index:Feature_stale_idx;default:false" json:"stale"`
	LastUsedAt  *DateTime      `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	CreatedAt   *DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false" json:"createdAt"`
	UpdatedAt   *DateTime      `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false" json:"updatedAt"`
	UpdatedBy   *int           `gorm:"column:updatedBy;size:32" json:"updatedBy"`
	Users       []UserFeatures `gorm:"foreignKey:FeatureID;references:Slug;constraint:UserFeatures_featureId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Teams       []TeamFeatures `gorm:"foreignKey:FeatureID;references:Slug;constraint:TeamFeatures_featureId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Feature) TableName() string {
	return "Feature"
}

func (m *Feature) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt == nil {
		m.CreatedAt = &DateTime{Time: now}
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	return nil
}

func (*Feature) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type UserFeatures struct {
	UserID     int      `gorm:"column:userId;primaryKey;index:UserFeatures_userId_featureId_idx;size:32" json:"userId"`
	FeatureID  string   `gorm:"column:featureId;primaryKey;index:UserFeatures_userId_featureId_idx" json:"featureId"`
	Enabled    bool     `gorm:"column:enabled;not null" json:"enabled"`
	AssignedAt DateTime `gorm:"column:assignedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"assignedAt"`
	AssignedBy string   `gorm:"column:assignedBy;not null" json:"assignedBy"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User       *User    `gorm:"belongsTo"`
	Feature    *Feature `gorm:"foreignKey:FeatureID;references:Slug;belongsTo"`
}

func (UserFeatures) TableName() string {
	return "UserFeatures"
}

func (m *UserFeatures) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.AssignedAt.IsZero() {
		m.AssignedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*UserFeatures) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TeamFeatures struct {
	TeamID     int      `gorm:"column:teamId;primaryKey;index:TeamFeatures_teamId_featureId_idx;size:32" json:"teamId"`
	FeatureID  string   `gorm:"column:featureId;primaryKey;index:TeamFeatures_teamId_featureId_idx" json:"featureId"`
	Enabled    bool     `gorm:"column:enabled;not null" json:"enabled"`
	AssignedAt DateTime `gorm:"column:assignedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"assignedAt"`
	AssignedBy string   `gorm:"column:assignedBy;not null" json:"assignedBy"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team       *Team    `gorm:"belongsTo"`
	Feature    *Feature `gorm:"foreignKey:FeatureID;references:Slug;belongsTo"`
}

func (TeamFeatures) TableName() string {
	return "TeamFeatures"
}

func (m *TeamFeatures) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.AssignedAt.IsZero() {
		m.AssignedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TeamFeatures) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SelectedSlots struct {
	ID               int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	EventTypeID      int      `gorm:"column:eventTypeId;size:32;not null" json:"eventTypeId"`
	UserID           int      `gorm:"column:userId;uniqueIndex:SelectedSlots_userId_slotUtcStartDate_slotUtcEndDate_uid_key;size:32;not null" json:"userId"`
	SlotUtcStartDate DateTime `gorm:"column:slotUtcStartDate;uniqueIndex:SelectedSlots_userId_slotUtcStartDate_slotUtcEndDate_uid_key;type:timestamp(3);not null" json:"slotUtcStartDate"`
	SlotUtcEndDate   DateTime `gorm:"column:slotUtcEndDate;uniqueIndex:SelectedSlots_userId_slotUtcStartDate_slotUtcEndDate_uid_key;type:timestamp(3);not null" json:"slotUtcEndDate"`
	UID              string   `gorm:"column:uid;uniqueIndex:SelectedSlots_userId_slotUtcStartDate_slotUtcEndDate_uid_key;not null" json:"uid"`
	ReleaseAt        DateTime `gorm:"column:releaseAt;type:timestamp(3);not null" json:"releaseAt"`
	IsSeat           bool     `gorm:"column:isSeat;default:false;not null" json:"isSeat"`
}

func (SelectedSlots) TableName() string {
	return "SelectedSlots"
}

type OAuthClient struct {
	ClientID        string       `gorm:"column:clientId;primaryKey" json:"clientId"`
	RedirectURI     string       `gorm:"column:redirectUri;not null" json:"redirectUri"`
	ClientSecret    *string      `gorm:"column:clientSecret" json:"clientSecret"`
	ClientType      string       `gorm:"column:clientType;default:'confidential';not null" json:"clientType"`
	Name            string       `gorm:"column:name;not null" json:"name"`
	Purpose         *string      `gorm:"column:purpose" json:"purpose"`
	Logo            *string      `gorm:"column:logo" json:"logo"`
	WebsiteURL      *string      `gorm:"column:websiteUrl" json:"websiteUrl"`
	RejectionReason *string      `gorm:"column:rejectionReason" json:"rejectionReason"`
	IsTrusted       bool         `gorm:"column:isTrusted;default:false;not null" json:"isTrusted"`
	Status          string       `gorm:"column:status;default:'approved';not null" json:"status"`
	UserID          *int         `gorm:"column:userId;index:OAuthClient_userId_idx;size:32" json:"userId"`
	CreatedAt       DateTime     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	User            *User        `gorm:"belongsTo"`
	AccessCodes     []AccessCode `gorm:"foreignKey:ClientID;references:ClientID;constraint:AccessCode_clientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OAuthClient) TableName() string {
	return "OAuthClient"
}

func (m *OAuthClient) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type AccessCode struct {
	ID                  int                `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Code                string             `gorm:"column:code;not null" json:"code"`
	ClientID            *string            `gorm:"column:clientId" json:"clientId"`
	ExpiresAt           DateTime           `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	Scopes              ScalarList[string] `gorm:"column:scopes;type:text[]" json:"scopes"`
	UserID              *int               `gorm:"column:userId;size:32" json:"userId"`
	TeamID              *int               `gorm:"column:teamId;size:32" json:"teamId"`
	CodeChallenge       *string            `gorm:"column:codeChallenge" json:"codeChallenge"`
	CodeChallengeMethod *string            `gorm:"column:codeChallengeMethod;default:'S256'" json:"codeChallengeMethod"`
	Client              *OAuthClient       `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
	User                *User              `gorm:"belongsTo"`
	Team                *Team              `gorm:"belongsTo"`
}

func (AccessCode) TableName() string {
	return "AccessCode"
}

type BookingDenormalized struct {
	ID             int       `gorm:"column:id;primaryKey;size:32" json:"id"`
	UID            string    `gorm:"column:uid;not null" json:"uid"`
	EventTypeID    *int      `gorm:"column:eventTypeId;index:BookingDenormalized_eventTypeId_idx;size:32" json:"eventTypeId"`
	Title          string    `gorm:"column:title;not null" json:"title"`
	Description    *string   `gorm:"column:description" json:"description"`
	StartTime      DateTime  `gorm:"column:startTime;index:BookingDenormalized_startTime_idx;index:BookingDenormalized_startTime_endTime_idx;type:timestamp(3);not null" json:"startTime"`
	EndTime        DateTime  `gorm:"column:endTime;index:BookingDenormalized_endTime_idx;index:BookingDenormalized_startTime_endTime_idx;type:timestamp(3);not null" json:"endTime"`
	CreatedAt      DateTime  `gorm:"column:createdAt;index:BookingDenormalized_createdAt_idx;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      *DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	Location       *string   `gorm:"column:location" json:"location"`
	Paid           bool      `gorm:"column:paid;not null" json:"paid"`
	Status         string    `gorm:"column:status;index:BookingDenormalized_status_idx;not null" json:"status"`
	Rescheduled    *bool     `gorm:"column:rescheduled" json:"rescheduled"`
	UserID         *int      `gorm:"column:userId;index:BookingDenormalized_userId_idx;index:BookingDenormalized_userId_isTeamBooking_idx;size:32" json:"userId"`
	TeamID         *int      `gorm:"column:teamId;index:BookingDenormalized_teamId_idx;index:BookingDenormalized_teamId_isTeamBooking_idx;size:32" json:"teamId"`
	EventLength    *int      `gorm:"column:eventLength;size:32" json:"eventLength"`
	EventParentID  *int      `gorm:"column:eventParentId;index:BookingDenormalized_eventParentId_idx;size:32" json:"eventParentId"`
	UserEmail      *string   `gorm:"column:userEmail" json:"userEmail"`
	UserName       *string   `gorm:"column:userName" json:"userName"`
	UserUsername   *string   `gorm:"column:userUsername" json:"userUsername"`
	RatingFeedback *string   `gorm:"column:ratingFeedback" json:"ratingFeedback"`
	Rating         *int      `gorm:"column:rating;size:32" json:"rating"`
	NoShowHost     *bool     `gorm:"column:noShowHost" json:"noShowHost"`
	IsTeamBooking  bool      `gorm:"column:isTeamBooking;index:BookingDenormalized_teamId_isTeamBooking_idx;index:BookingDenormalized_userId_isTeamBooking_idx;not null" json:"isTeamBooking"`
}

func (BookingDenormalized) TableName() string {
	return "BookingDenormalized"
}

type CalendarCache struct {
	ID           *string        `gorm:"column:id" json:"id"`
	Key          string         `gorm:"column:key;primaryKey;uniqueIndex:CalendarCache_credentialId_key_key;index:CalendarCache_userId_key_idx" json:"key"`
	Value        datatypes.JSON `gorm:"column:value;not null" json:"value"`
	ExpiresAt    DateTime       `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	UpdatedAt    DateTime       `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
	CredentialID int            `gorm:"column:credentialId;primaryKey;uniqueIndex:CalendarCache_credentialId_key_key;size:32" json:"credentialId"`
	UserID       *int           `gorm:"column:userId;index:CalendarCache_userId_key_idx;size:32" json:"userId"`
	Credential   *Credential    `gorm:"belongsTo"`
}

func (CalendarCache) TableName() string {
	return "CalendarCache"
}

func (m *CalendarCache) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == nil {
		generated := uuid.NewString()
		m.ID = &generated
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CalendarCache) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TempOrgRedirect struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	From      string   `gorm:"column:from;uniqueIndex:TempOrgRedirect_from_type_fromOrgId_key;not null" json:"from"`
	FromOrgID int      `gorm:"column:fromOrgId;uniqueIndex:TempOrgRedirect_from_type_fromOrgId_key;size:32;not null" json:"fromOrgId"`
	Type      string   `gorm:"column:type;uniqueIndex:TempOrgRedirect_from_type_fromOrgId_key;not null" json:"type"`
	ToURL     string   `gorm:"column:toUrl;not null" json:"toUrl"`
	Enabled   *bool    `gorm:"column:enabled;default:true;not null" json:"enabled"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (TempOrgRedirect) TableName() string {
	return "TempOrgRedirect"
}

func (m *TempOrgRedirect) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TempOrgRedirect) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Avatar struct {
	TeamID    int    `gorm:"column:teamId;uniqueIndex:avatars_teamId_userId_isBanner_key;size:32;default:0;not null" json:"teamId"`
	UserID    int    `gorm:"column:userId;uniqueIndex:avatars_teamId_userId_isBanner_key;size:32;default:0;not null" json:"userId"`
	Data      string `gorm:"column:data;not null" json:"data"`
	ObjectKey string `gorm:"column:objectKey;primaryKey;uniqueIndex:avatars_objectKey_key" json:"objectKey"`
	IsBanner  bool   `gorm:"column:isBanner;uniqueIndex:avatars_teamId_userId_isBanner_key;default:false;not null" json:"isBanner"`
}

func (Avatar) TableName() string {
	return "avatars"
}

type OutOfOfficeEntry struct {
	ID               int                `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UUID             string             `gorm:"column:uuid;uniqueIndex:OutOfOfficeEntry_uuid_key;index:OutOfOfficeEntry_uuid_idx;not null" json:"uuid"`
	Start            DateTime           `gorm:"column:start;index:OutOfOfficeEntry_start_end_idx;type:timestamp(3);not null" json:"start"`
	End              DateTime           `gorm:"column:end;index:OutOfOfficeEntry_start_end_idx;type:timestamp(3);not null" json:"end"`
	Notes            *string            `gorm:"column:notes" json:"notes"`
	ShowNotePublicly bool               `gorm:"column:showNotePublicly;default:false;not null" json:"showNotePublicly"`
	UserID           int                `gorm:"column:userId;index:OutOfOfficeEntry_userId_idx;size:32;not null" json:"userId"`
	ToUserID         *int               `gorm:"column:toUserId;index:OutOfOfficeEntry_toUserId_idx;size:32" json:"toUserId"`
	ReasonID         *int               `gorm:"column:reasonId;size:32" json:"reasonId"`
	CreatedAt        DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User             *User              `gorm:"foreignKey:UserID;belongsTo"`
	ToUser           *User              `gorm:"foreignKey:ToUserID;belongsTo"`
	Reason           *OutOfOfficeReason `gorm:"foreignKey:ReasonID;belongsTo"`
}

func (OutOfOfficeEntry) TableName() string {
	return "OutOfOfficeEntry"
}

func (m *OutOfOfficeEntry) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OutOfOfficeEntry) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OutOfOfficeReason struct {
	ID      int                `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Emoji   string             `gorm:"column:emoji;not null" json:"emoji"`
	Reason  string             `gorm:"column:reason;uniqueIndex:OutOfOfficeReason_reason_key;not null" json:"reason"`
	Enabled *bool              `gorm:"column:enabled;default:true;not null" json:"enabled"`
	UserID  *int               `gorm:"column:userId;size:32" json:"userId"`
	User    *User              `gorm:"belongsTo"`
	Entries []OutOfOfficeEntry `gorm:"foreignKey:ReasonID;constraint:OutOfOfficeEntry_reasonId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (OutOfOfficeReason) TableName() string {
	return "OutOfOfficeReason"
}

type UserHolidaySettings struct {
	ID          int                `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID      int                `gorm:"column:userId;uniqueIndex:UserHolidaySettings_userId_key;size:32;not null" json:"userId"`
	CountryCode *string            `gorm:"column:countryCode" json:"countryCode"`
	DisabledIds ScalarList[string] `gorm:"column:disabledIds;type:text[]" json:"disabledIds"`
	CreatedAt   DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User        *User              `gorm:"belongsTo"`
}

func (UserHolidaySettings) TableName() string {
	return "UserHolidaySettings"
}

func (m *UserHolidaySettings) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*UserHolidaySettings) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type HolidayCache struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CountryCode string   `gorm:"column:countryCode;uniqueIndex:HolidayCache_countryCode_eventId_key;index:HolidayCache_countryCode_year_idx;index:HolidayCache_countryCode_date_idx;not null" json:"countryCode"`
	CalendarID  string   `gorm:"column:calendarId;not null" json:"calendarId"`
	EventID     string   `gorm:"column:eventId;uniqueIndex:HolidayCache_countryCode_eventId_key;not null" json:"eventId"`
	Name        string   `gorm:"column:name;not null" json:"name"`
	Date        DateTime `gorm:"column:date;index:HolidayCache_countryCode_date_idx;type:timestamp(3);not null" json:"date"`
	Year        int      `gorm:"column:year;index:HolidayCache_countryCode_year_idx;size:32;not null" json:"year"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (HolidayCache) TableName() string {
	return "HolidayCache"
}

func (m *HolidayCache) BeforeCreate(tx *gorm.DB) error {
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

func (*HolidayCache) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PlatformOAuthClient struct {
	ID                           string                       `gorm:"column:id;primaryKey" json:"id"`
	Name                         string                       `gorm:"column:name;not null" json:"name"`
	Secret                       string                       `gorm:"column:secret;not null" json:"secret"`
	Permissions                  int                          `gorm:"column:permissions;size:32;not null" json:"permissions"`
	Logo                         *string                      `gorm:"column:logo" json:"logo"`
	RedirectUris                 ScalarList[string]           `gorm:"column:redirectUris;type:text[]" json:"redirectUris"`
	OrganizationID               int                          `gorm:"column:organizationId;size:32;not null" json:"organizationId"`
	BookingRedirectURI           *string                      `gorm:"column:bookingRedirectUri" json:"bookingRedirectUri"`
	BookingCancelRedirectURI     *string                      `gorm:"column:bookingCancelRedirectUri" json:"bookingCancelRedirectUri"`
	BookingRescheduleRedirectURI *string                      `gorm:"column:bookingRescheduleRedirectUri" json:"bookingRescheduleRedirectUri"`
	AreEmailsEnabled             bool                         `gorm:"column:areEmailsEnabled;default:false;not null" json:"areEmailsEnabled"`
	AreDefaultEventTypesEnabled  *bool                        `gorm:"column:areDefaultEventTypesEnabled;default:true;not null" json:"areDefaultEventTypesEnabled"`
	AreCalendarEventsEnabled     *bool                        `gorm:"column:areCalendarEventsEnabled;default:true;not null" json:"areCalendarEventsEnabled"`
	CreatedAt                    DateTime                     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Organization                 *Team                        `gorm:"foreignKey:OrganizationID;belongsTo"`
	Teams                        []Team                       `gorm:"foreignKey:CreatedByOAuthClientID;constraint:Team_createdByOAuthClientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessTokens                 []AccessToken                `gorm:"foreignKey:PlatformOAuthClientID;constraint:AccessToken_platformOAuthClientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RefreshToken                 []RefreshToken               `gorm:"foreignKey:PlatformOAuthClientID;constraint:RefreshToken_platformOAuthClientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AuthorizationTokens          []PlatformAuthorizationToken `gorm:"foreignKey:PlatformOAuthClientID;constraint:PlatformAuthorizationToken_platformOAuthClientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhook                      []Webhook                    `gorm:"foreignKey:PlatformOAuthClientID;constraint:Webhook_platformOAuthClientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Users                        []User                       `gorm:"many2many:_PlatformOAuthClientToUser;joinForeignKey:A;joinReferences:B"`
}

func (PlatformOAuthClient) TableName() string {
	return "PlatformOAuthClient"
}

func (m *PlatformOAuthClient) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type PlatformAuthorizationToken struct {
	ID                    string               `gorm:"column:id;primaryKey" json:"id"`
	PlatformOAuthClientID string               `gorm:"column:platformOAuthClientId;uniqueIndex:PlatformAuthorizationToken_userId_platformOAuthClientId_key;not null" json:"platformOAuthClientId"`
	UserID                int                  `gorm:"column:userId;uniqueIndex:PlatformAuthorizationToken_userId_platformOAuthClientId_key;size:32;not null" json:"userId"`
	CreatedAt             DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Owner                 *User                `gorm:"foreignKey:UserID;belongsTo"`
	Client                *PlatformOAuthClient `gorm:"foreignKey:PlatformOAuthClientID;belongsTo"`
}

func (PlatformAuthorizationToken) TableName() string {
	return "PlatformAuthorizationToken"
}

func (m *PlatformAuthorizationToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type AccessToken struct {
	ID                    int                  `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Secret                string               `gorm:"column:secret;uniqueIndex:AccessToken_secret_key;not null" json:"secret"`
	CreatedAt             DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ExpiresAt             DateTime             `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	PlatformOAuthClientID string               `gorm:"column:platformOAuthClientId;not null" json:"platformOAuthClientId"`
	UserID                int                  `gorm:"column:userId;size:32;not null" json:"userId"`
	Owner                 *User                `gorm:"foreignKey:UserID;belongsTo"`
	Client                *PlatformOAuthClient `gorm:"foreignKey:PlatformOAuthClientID;belongsTo"`
}

func (AccessToken) TableName() string {
	return "AccessToken"
}

func (m *AccessToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type RefreshToken struct {
	ID                    int                  `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Secret                string               `gorm:"column:secret;uniqueIndex:RefreshToken_secret_key;not null" json:"secret"`
	CreatedAt             DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ExpiresAt             DateTime             `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	PlatformOAuthClientID string               `gorm:"column:platformOAuthClientId;not null" json:"platformOAuthClientId"`
	UserID                int                  `gorm:"column:userId;size:32;not null" json:"userId"`
	Owner                 *User                `gorm:"foreignKey:UserID;belongsTo"`
	Client                *PlatformOAuthClient `gorm:"foreignKey:PlatformOAuthClientID;belongsTo"`
}

func (RefreshToken) TableName() string {
	return "RefreshToken"
}

func (m *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DSyncData struct {
	ID                      int                     `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	DirectoryID             string                  `gorm:"column:directoryId;uniqueIndex:DSyncData_directoryId_key;not null" json:"directoryId"`
	Tenant                  string                  `gorm:"column:tenant;not null" json:"tenant"`
	OrganizationID          *int                    `gorm:"column:organizationId;uniqueIndex:DSyncData_organizationId_key;size:32" json:"organizationId"`
	Org                     *OrganizationSettings   `gorm:"foreignKey:OrganizationID;references:OrganizationID;belongsTo"`
	TeamGroupMapping        []DSyncTeamGroupMapping `gorm:"foreignKey:DirectoryID;references:DirectoryID;constraint:DSyncTeamGroupMapping_directoryId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAttributeToUsers []AttributeToUser       `gorm:"foreignKey:CreatedByDSyncID;references:DirectoryID;constraint:AttributeToUser_createdByDSyncId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	UpdatedAttributeToUsers []AttributeToUser       `gorm:"foreignKey:UpdatedByDSyncID;references:DirectoryID;constraint:AttributeToUser_updatedByDSyncId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (DSyncData) TableName() string {
	return "DSyncData"
}

type DSyncTeamGroupMapping struct {
	ID             int        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	OrganizationID int        `gorm:"column:organizationId;size:32;not null" json:"organizationId"`
	TeamID         int        `gorm:"column:teamId;uniqueIndex:DSyncTeamGroupMapping_teamId_groupName_key;size:32;not null" json:"teamId"`
	DirectoryID    string     `gorm:"column:directoryId;not null" json:"directoryId"`
	GroupName      string     `gorm:"column:groupName;uniqueIndex:DSyncTeamGroupMapping_teamId_groupName_key;not null" json:"groupName"`
	Team           *Team      `gorm:"belongsTo"`
	Directory      *DSyncData `gorm:"foreignKey:DirectoryID;references:DirectoryID;belongsTo"`
}

func (DSyncTeamGroupMapping) TableName() string {
	return "DSyncTeamGroupMapping"
}

type SecondaryEmail struct {
	ID                 int                 `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID             int                 `gorm:"column:userId;uniqueIndex:SecondaryEmail_userId_email_key;index:SecondaryEmail_userId_idx;size:32;not null" json:"userId"`
	Email              string              `gorm:"column:email;uniqueIndex:SecondaryEmail_email_key;uniqueIndex:SecondaryEmail_userId_email_key;index:SecondaryEmail_email_emailVerified_idx;not null" json:"email"`
	EmailVerified      *DateTime           `gorm:"column:emailVerified;index:SecondaryEmail_email_emailVerified_idx;type:timestamp(3)" json:"emailVerified"`
	User               *User               `gorm:"belongsTo"`
	VerificationTokens []VerificationToken `gorm:"foreignKey:SecondaryEmailID;constraint:VerificationToken_secondaryEmailId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	EventTypes         []EventType         `gorm:"foreignKey:SecondaryEmailID;constraint:EventType_secondaryEmailId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (SecondaryEmail) TableName() string {
	return "SecondaryEmail"
}

type Task struct {
	ID                  string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt           DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	ScheduledAt         DateTime  `gorm:"column:scheduledAt;index:Task_scheduledAt_succeededAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"scheduledAt"`
	SucceededAt         *DateTime `gorm:"column:succeededAt;index:Task_succeededAt_idx;index:Task_scheduledAt_succeededAt_idx;type:timestamp(3)" json:"succeededAt"`
	Type                string    `gorm:"column:type;uniqueIndex:Task_referenceUid_type_key;not null" json:"type"`
	Payload             string    `gorm:"column:payload;not null" json:"payload"`
	Attempts            int       `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	MaxAttempts         *int      `gorm:"column:maxAttempts;size:32;default:3;not null" json:"maxAttempts"`
	LastError           *string   `gorm:"column:lastError" json:"lastError"`
	LastFailedAttemptAt *DateTime `gorm:"column:lastFailedAttemptAt;type:timestamp(3)" json:"lastFailedAttemptAt"`
	ReferenceUID        *string   `gorm:"column:referenceUid;uniqueIndex:Task_referenceUid_type_key" json:"referenceUid"`
}

func (Task) TableName() string {
	return "Task"
}

func (m *Task) BeforeCreate(tx *gorm.DB) error {
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
	if m.ScheduledAt.IsZero() {
		m.ScheduledAt = DateTime{Time: now}
	}
	return nil
}

func (*Task) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ManagedOrganization struct {
	ManagedOrganizationID int      `gorm:"column:managedOrganizationId;primaryKey;uniqueIndex:ManagedOrganization_managedOrganizationId_key;uniqueIndex:ManagedOrganization_managerOrganizationId_managedOrganizati_key;size:32" json:"managedOrganizationId"`
	ManagerOrganizationID int      `gorm:"column:managerOrganizationId;uniqueIndex:ManagedOrganization_managerOrganizationId_managedOrganizati_key;index:ManagedOrganization_managerOrganizationId_idx;size:32;not null" json:"managerOrganizationId"`
	CreatedAt             DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ManagedOrganization   *Team    `gorm:"foreignKey:ManagedOrganizationID;belongsTo"`
	ManagerOrganization   *Team    `gorm:"foreignKey:ManagerOrganizationID;belongsTo"`
}

func (ManagedOrganization) TableName() string {
	return "ManagedOrganization"
}

func (m *ManagedOrganization) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type PlatformBilling struct {
	ID                int               `gorm:"column:id;primaryKey;size:32" json:"id"`
	CustomerID        string            `gorm:"column:customerId;not null" json:"customerId"`
	SubscriptionID    *string           `gorm:"column:subscriptionId" json:"subscriptionId"`
	PriceID           *string           `gorm:"column:priceId" json:"priceId"`
	Plan              *string           `gorm:"column:plan;default:'none';not null" json:"plan"`
	BillingCycleStart *int              `gorm:"column:billingCycleStart;size:32" json:"billingCycleStart"`
	BillingCycleEnd   *int              `gorm:"column:billingCycleEnd;size:32" json:"billingCycleEnd"`
	Overdue           *bool             `gorm:"column:overdue;default:false" json:"overdue"`
	ManagerBillingID  *int              `gorm:"column:managerBillingId;size:32" json:"managerBillingId"`
	ManagerBilling    *PlatformBilling  `gorm:"foreignKey:ManagerBillingID;belongsTo"`
	Team              *Team             `gorm:"foreignKey:ID;belongsTo"`
	ManagedBillings   []PlatformBilling `gorm:"foreignKey:ManagerBillingID;constraint:PlatformBilling_managerBillingId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (PlatformBilling) TableName() string {
	return "PlatformBilling"
}

type AttributeOption struct {
	ID            string             `gorm:"column:id;primaryKey" json:"id"`
	AttributeID   string             `gorm:"column:attributeId;not null" json:"attributeId"`
	Value         string             `gorm:"column:value;not null" json:"value"`
	Slug          string             `gorm:"column:slug;not null" json:"slug"`
	IsGroup       bool               `gorm:"column:isGroup;default:false;not null" json:"isGroup"`
	Contains      ScalarList[string] `gorm:"column:contains;type:text[]" json:"contains"`
	Attribute     *Attribute         `gorm:"belongsTo"`
	AssignedUsers []AttributeToUser  `gorm:"foreignKey:AttributeOptionID;constraint:AttributeToUser_attributeOptionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AttributeOption) TableName() string {
	return "AttributeOption"
}

func (m *AttributeOption) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

type Attribute struct {
	ID                         string                      `gorm:"column:id;primaryKey" json:"id"`
	TeamID                     int                         `gorm:"column:teamId;uniqueIndex:Attribute_teamId_slug_key;index:Attribute_teamId_idx;size:32;not null" json:"teamId"`
	Type                       string                      `gorm:"column:type;not null" json:"type"`
	Name                       string                      `gorm:"column:name;not null" json:"name"`
	Slug                       string                      `gorm:"column:slug;uniqueIndex:Attribute_teamId_slug_key;not null" json:"slug"`
	Enabled                    *bool                       `gorm:"column:enabled;default:true;not null" json:"enabled"`
	UsersCanEditRelation       bool                        `gorm:"column:usersCanEditRelation;default:false;not null" json:"usersCanEditRelation"`
	CreatedAt                  DateTime                    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime                    `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	IsWeightsEnabled           bool                        `gorm:"column:isWeightsEnabled;default:false;not null" json:"isWeightsEnabled"`
	IsLocked                   bool                        `gorm:"column:isLocked;default:false;not null" json:"isLocked"`
	Team                       *Team                       `gorm:"belongsTo"`
	Options                    []AttributeOption           `gorm:"foreignKey:AttributeID;constraint:AttributeOption_attributeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AttributeSyncFieldMappings []AttributeSyncFieldMapping `gorm:"foreignKey:AttributeID;constraint:AttributeSyncFieldMapping_attributeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Attribute) TableName() string {
	return "Attribute"
}

func (m *Attribute) BeforeCreate(tx *gorm.DB) error {
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

func (*Attribute) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AttributeToUser struct {
	ID                string           `gorm:"column:id;primaryKey" json:"id"`
	MemberID          int              `gorm:"column:memberId;uniqueIndex:AttributeToUser_memberId_attributeOptionId_key;size:32;not null" json:"memberId"`
	AttributeOptionID string           `gorm:"column:attributeOptionId;uniqueIndex:AttributeToUser_memberId_attributeOptionId_key;not null" json:"attributeOptionId"`
	Weight            *int             `gorm:"column:weight;size:32" json:"weight"`
	CreatedAt         DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	CreatedByID       *int             `gorm:"column:createdById;size:32" json:"createdById"`
	CreatedByDSyncID  *string          `gorm:"column:createdByDSyncId" json:"createdByDSyncId"`
	UpdatedAt         *DateTime        `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	UpdatedByID       *int             `gorm:"column:updatedById;size:32" json:"updatedById"`
	UpdatedByDSyncID  *string          `gorm:"column:updatedByDSyncId" json:"updatedByDSyncId"`
	Member            *Membership      `gorm:"foreignKey:MemberID;belongsTo"`
	AttributeOption   *AttributeOption `gorm:"belongsTo"`
	CreatedBy         *User            `gorm:"foreignKey:CreatedByID;belongsTo"`
	CreatedByDSync    *DSyncData       `gorm:"foreignKey:CreatedByDSyncID;references:DirectoryID;belongsTo"`
	UpdatedBy         *User            `gorm:"foreignKey:UpdatedByID;belongsTo"`
	UpdatedByDSync    *DSyncData       `gorm:"foreignKey:UpdatedByDSyncID;references:DirectoryID;belongsTo"`
}

func (AttributeToUser) TableName() string {
	return "AttributeToUser"
}

func (m *AttributeToUser) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = &DateTime{Time: now}
	}
	return nil
}

func (*AttributeToUser) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", &DateTime{Time: now})
	}
	return nil
}

type AssignmentReason struct {
	ID           int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	CreatedAt    DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	BookingID    int      `gorm:"column:bookingId;index:AssignmentReason_bookingId_idx;size:32;not null" json:"bookingId"`
	ReasonEnum   string   `gorm:"column:reasonEnum;not null" json:"reasonEnum"`
	ReasonString string   `gorm:"column:reasonString;not null" json:"reasonString"`
	Booking      *Booking `gorm:"belongsTo"`
}

func (AssignmentReason) TableName() string {
	return "AssignmentReason"
}

func (m *AssignmentReason) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DelegationCredential struct {
	ID                  string `gorm:"column:id;primaryKey" json:"id"`
	WorkspacePlatformID int    `gorm:"column:workspacePlatformId;size:32;not null" json:"workspacePlatformId"`
	// @zod.import(["import { serviceAccountKeySchema } from '../../zod-utils'"]).custom.use(serviceAccountKeySchema)
	ServiceAccountKey   datatypes.JSON        `gorm:"column:serviceAccountKey;not null" json:"serviceAccountKey"`
	Enabled             bool                  `gorm:"column:enabled;index:DelegationCredential_enabled_idx;default:false;not null" json:"enabled"`
	LastEnabledAt       *DateTime             `gorm:"column:lastEnabledAt;type:timestamp(3)" json:"lastEnabledAt"`
	LastDisabledAt      *DateTime             `gorm:"column:lastDisabledAt;type:timestamp(3)" json:"lastDisabledAt"`
	OrganizationID      int                   `gorm:"column:organizationId;uniqueIndex:DelegationCredential_organizationId_domain_key;size:32;not null" json:"organizationId"`
	Domain              string                `gorm:"column:domain;uniqueIndex:DelegationCredential_organizationId_domain_key;not null" json:"domain"`
	CreatedAt           DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime              `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	WorkspacePlatform   *WorkspacePlatform    `gorm:"belongsTo"`
	Organization        *Team                 `gorm:"foreignKey:OrganizationID;belongsTo"`
	SelectedCalendars   []SelectedCalendar    `gorm:"foreignKey:DelegationCredentialID;constraint:SelectedCalendar_delegationCredentialId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DestinationCalendar []DestinationCalendar `gorm:"foreignKey:DelegationCredentialID;constraint:DestinationCalendar_delegationCredentialId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BookingReferences   []BookingReference    `gorm:"foreignKey:DelegationCredentialID;constraint:BookingReference_delegationCredentialId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Credentials         []Credential          `gorm:"foreignKey:DelegationCredentialID;constraint:Credential_delegationCredentialId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (DelegationCredential) TableName() string {
	return "DelegationCredential"
}

func (m *DelegationCredential) BeforeCreate(tx *gorm.DB) error {
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

func (*DelegationCredential) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkspacePlatform struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	// @zod.string.min(1)
	Slug string `gorm:"column:slug;uniqueIndex:WorkspacePlatform_slug_key;not null" json:"slug"`
	// @zod.string.min(1)
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	// @zod.import(["import { serviceAccountKeySchema } from '../../zod-utils'"]).custom.use(serviceAccountKeySchema)
	DefaultServiceAccountKey datatypes.JSON         `gorm:"column:defaultServiceAccountKey;not null" json:"defaultServiceAccountKey"`
	CreatedAt                DateTime               `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime               `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Enabled                  bool                   `gorm:"column:enabled;default:false;not null" json:"enabled"`
	DelegationCredentials    []DelegationCredential `gorm:"foreignKey:WorkspacePlatformID;constraint:DelegationCredential_workspacePlatformId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (WorkspacePlatform) TableName() string {
	return "WorkspacePlatform"
}

func (m *WorkspacePlatform) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkspacePlatform) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type EventTypeTranslation struct {
	UID            string     `gorm:"column:uid;primaryKey" json:"uid"`
	EventTypeID    int        `gorm:"column:eventTypeId;uniqueIndex:EventTypeTranslation_eventTypeId_field_targetLocale_key;index:EventTypeTranslation_eventTypeId_field_targetLocale_idx;size:32;not null" json:"eventTypeId"`
	Field          string     `gorm:"column:field;uniqueIndex:EventTypeTranslation_eventTypeId_field_targetLocale_key;index:EventTypeTranslation_eventTypeId_field_targetLocale_idx;not null" json:"field"`
	SourceLocale   string     `gorm:"column:sourceLocale;not null" json:"sourceLocale"`
	TargetLocale   string     `gorm:"column:targetLocale;uniqueIndex:EventTypeTranslation_eventTypeId_field_targetLocale_key;index:EventTypeTranslation_eventTypeId_field_targetLocale_idx;not null" json:"targetLocale"`
	TranslatedText string     `gorm:"column:translatedText;type:text;not null" json:"translatedText"`
	CreatedAt      DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	CreatedBy      int        `gorm:"column:createdBy;size:32;not null" json:"createdBy"`
	UpdatedAt      DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UpdatedBy      *int       `gorm:"column:updatedBy;size:32" json:"updatedBy"`
	EventType      *EventType `gorm:"belongsTo"`
	Creator        *User      `gorm:"foreignKey:CreatedBy;belongsTo"`
	Updater        *User      `gorm:"foreignKey:UpdatedBy;belongsTo"`
}

func (EventTypeTranslation) TableName() string {
	return "EventTypeTranslation"
}

func (m *EventTypeTranslation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UID == "" {
		m.UID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*EventTypeTranslation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Watchlist struct {
	ID                   string           `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	Type                 string           `gorm:"column:type;uniqueIndex:Watchlist_type_value_organizationId_key;index:Watchlist_isGlobal_action_organizationId_type_value_idx;not null" json:"type"`
	Value                string           `gorm:"column:value;uniqueIndex:Watchlist_type_value_organizationId_key;index:Watchlist_isGlobal_action_organizationId_type_value_idx;not null" json:"value"`
	Description          *string          `gorm:"column:description" json:"description"`
	IsGlobal             bool             `gorm:"column:isGlobal;index:Watchlist_isGlobal_action_organizationId_type_value_idx;default:false;not null" json:"isGlobal"`
	OrganizationID       *int             `gorm:"column:organizationId;uniqueIndex:Watchlist_type_value_organizationId_key;index:Watchlist_isGlobal_action_organizationId_type_value_idx;size:32" json:"organizationId"`
	Action               string           `gorm:"column:action;index:Watchlist_isGlobal_action_organizationId_type_value_idx;default:'REPORT';not null" json:"action"`
	Source               string           `gorm:"column:source;default:'MANUAL';not null" json:"source"`
	LastUpdatedAt        DateTime         `gorm:"column:lastUpdatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastUpdatedAt"`
	BookingReports       []BookingReport  `gorm:"foreignKey:WatchlistID;constraint:BookingReport_watchlistId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	GlobalBookingReports []BookingReport  `gorm:"foreignKey:GlobalWatchlistID;constraint:BookingReport_globalWatchlistId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Audits               []WatchlistAudit `gorm:"foreignKey:WatchlistID;constraint:WatchlistAudit_watchlistId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Watchlist) TableName() string {
	return "Watchlist"
}

func (m *Watchlist) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.LastUpdatedAt.IsZero() {
		m.LastUpdatedAt = DateTime{Time: now}
	}
	return nil
}

type WatchlistAudit struct {
	ID              string     `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	Type            string     `gorm:"column:type;not null" json:"type"`
	Value           string     `gorm:"column:value;not null" json:"value"`
	Description     *string    `gorm:"column:description" json:"description"`
	Action          string     `gorm:"column:action;default:'REPORT';not null" json:"action"`
	ChangedAt       DateTime   `gorm:"column:changedAt;index:WatchlistAudit_watchlistId_changedAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"changedAt"`
	ChangedByUserID *int       `gorm:"column:changedByUserId;size:32" json:"changedByUserId"`
	WatchlistID     *string    `gorm:"column:watchlistId;index:WatchlistAudit_watchlistId_changedAt_idx;type:char(36)" json:"watchlistId"`
	Watchlist       *Watchlist `gorm:"belongsTo"`
}

func (WatchlistAudit) TableName() string {
	return "WatchlistAudit"
}

func (m *WatchlistAudit) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.Must(uuid.NewV7()).String()
	}
	if m.ChangedAt.IsZero() {
		m.ChangedAt = DateTime{Time: now}
	}
	return nil
}

type WatchlistEventAudit struct {
	ID          string   `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	WatchlistID string   `gorm:"column:watchlistId;type:char(36);not null" json:"watchlistId"`
	EventTypeID int      `gorm:"column:eventTypeId;size:32;not null" json:"eventTypeId"`
	ActionTaken string   `gorm:"column:actionTaken;not null" json:"actionTaken"`
	Timestamp   DateTime `gorm:"column:timestamp;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"timestamp"`
}

func (WatchlistEventAudit) TableName() string {
	return "WatchlistEventAudit"
}

func (m *WatchlistEventAudit) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.Must(uuid.NewV7()).String()
	}
	if m.Timestamp.IsZero() {
		m.Timestamp = DateTime{Time: now}
	}
	return nil
}

type BookingReport struct {
	ID                string     `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	BookingUID        string     `gorm:"column:bookingUid;uniqueIndex:BookingReport_bookingUid_key;not null" json:"bookingUid"`
	BookerEmail       string     `gorm:"column:bookerEmail;index:BookingReport_bookerEmail_idx;not null" json:"bookerEmail"`
	ReportedByID      *int       `gorm:"column:reportedById;index:BookingReport_reportedById_idx;size:32" json:"reportedById"`
	OrganizationID    *int       `gorm:"column:organizationId;index:BookingReport_organizationId_idx;size:32" json:"organizationId"`
	Reason            string     `gorm:"column:reason;not null" json:"reason"`
	Description       *string    `gorm:"column:description" json:"description"`
	Cancelled         bool       `gorm:"column:cancelled;default:false;not null" json:"cancelled"`
	CreatedAt         DateTime   `gorm:"column:createdAt;index:BookingReport_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Status            string     `gorm:"column:status;default:'PENDING';not null" json:"status"`
	SystemStatus      string     `gorm:"column:systemStatus;index:BookingReport_systemStatus_idx;default:'PENDING';not null" json:"systemStatus"`
	WatchlistID       *string    `gorm:"column:watchlistId;index:BookingReport_watchlistId_idx;type:char(36)" json:"watchlistId"`
	GlobalWatchlistID *string    `gorm:"column:globalWatchlistId;index:BookingReport_globalWatchlistId_idx;type:char(36)" json:"globalWatchlistId"`
	Booking           *Booking   `gorm:"references:UID;belongsTo"`
	ReportedBy        *User      `gorm:"foreignKey:ReportedByID;belongsTo"`
	Organization      *Team      `gorm:"foreignKey:OrganizationID;belongsTo"`
	Watchlist         *Watchlist `gorm:"foreignKey:WatchlistID;belongsTo"`
	GlobalWatchlist   *Watchlist `gorm:"foreignKey:GlobalWatchlistID;belongsTo"`
}

func (BookingReport) TableName() string {
	return "BookingReport"
}

func (m *BookingReport) BeforeCreate(tx *gorm.DB) error {
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

func (*BookingReport) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WrongAssignmentReport struct {
	ID              string    `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	BookingUID      string    `gorm:"column:bookingUid;uniqueIndex:WrongAssignmentReport_bookingUid_key;not null" json:"bookingUid"`
	ReportedByID    *int      `gorm:"column:reportedById;index:WrongAssignmentReport_reportedById_idx;size:32" json:"reportedById"`
	CorrectAssignee *string   `gorm:"column:correctAssignee" json:"correctAssignee"`
	AdditionalNotes string    `gorm:"column:additionalNotes;not null" json:"additionalNotes"`
	TeamID          *int      `gorm:"column:teamId;index:WrongAssignmentReport_teamId_idx;index:WrongAssignmentReport_teamId_status_createdAt_idx;size:32" json:"teamId"`
	Status          string    `gorm:"column:status;index:WrongAssignmentReport_status_idx;index:WrongAssignmentReport_teamId_status_createdAt_idx;default:'PENDING';not null" json:"status"`
	ReviewedByID    *int      `gorm:"column:reviewedById;index:WrongAssignmentReport_reviewedById_idx;size:32" json:"reviewedById"`
	ReviewedAt      *DateTime `gorm:"column:reviewedAt;type:timestamp(3)" json:"reviewedAt"`
	CreatedAt       DateTime  `gorm:"column:createdAt;index:WrongAssignmentReport_createdAt_idx;index:WrongAssignmentReport_teamId_status_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Booking         *Booking  `gorm:"references:UID;belongsTo"`
	ReportedBy      *User     `gorm:"foreignKey:ReportedByID;belongsTo"`
	Team            *Team     `gorm:"belongsTo"`
	ReviewedBy      *User     `gorm:"foreignKey:ReviewedByID;belongsTo"`
}

func (WrongAssignmentReport) TableName() string {
	return "WrongAssignmentReport"
}

func (m *WrongAssignmentReport) BeforeCreate(tx *gorm.DB) error {
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

func (*WrongAssignmentReport) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganizationOnboarding struct {
	ID                       string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedByID              int      `gorm:"column:createdById;size:32;not null" json:"createdById"`
	CreatedAt                DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	OrgOwnerEmail            string   `gorm:"column:orgOwnerEmail;uniqueIndex:OrganizationOnboarding_orgOwnerEmail_key;index:OrganizationOnboarding_orgOwnerEmail_idx;not null" json:"orgOwnerEmail"`
	Error                    *string  `gorm:"column:error" json:"error"`
	UpdatedAt                DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	OrganizationID           *int     `gorm:"column:organizationId;uniqueIndex:OrganizationOnboarding_organizationId_key;size:32" json:"organizationId"`
	BillingPeriod            string   `gorm:"column:billingPeriod;not null" json:"billingPeriod"`
	PricePerSeat             float64  `gorm:"column:pricePerSeat;type:double precision;not null" json:"pricePerSeat"`
	Seats                    int      `gorm:"column:seats;size:32;not null" json:"seats"`
	IsPlatform               bool     `gorm:"column:isPlatform;default:false;not null" json:"isPlatform"`
	Name                     string   `gorm:"column:name;not null" json:"name"`
	Slug                     string   `gorm:"column:slug;not null" json:"slug"`
	Logo                     *string  `gorm:"column:logo" json:"logo"`
	Bio                      *string  `gorm:"column:bio" json:"bio"`
	BrandColor               *string  `gorm:"column:brandColor" json:"brandColor"`
	BannerURL                *string  `gorm:"column:bannerUrl" json:"bannerUrl"`
	IsDomainConfigured       bool     `gorm:"column:isDomainConfigured;default:false;not null" json:"isDomainConfigured"`
	StripeCustomerID         *string  `gorm:"column:stripeCustomerId;uniqueIndex:OrganizationOnboarding_stripeCustomerId_key;index:OrganizationOnboarding_stripeCustomerId_idx" json:"stripeCustomerId"`
	StripeSubscriptionID     *string  `gorm:"column:stripeSubscriptionId" json:"stripeSubscriptionId"`
	StripeSubscriptionItemID *string  `gorm:"column:stripeSubscriptionItemId" json:"stripeSubscriptionItemId"`
	// @zod.import(["import { orgOnboardingInvitedMembersSchema } from '../../zod-utils'"]).custom.use(orgOnboardingInvitedMembersSchema)
	InvitedMembers datatypes.JSON `gorm:"column:invitedMembers;default:'[]';not null" json:"invitedMembers"`
	// @zod.import(["import { orgOnboardingTeamsSchema } from '../../zod-utils'"]).custom.use(orgOnboardingTeamsSchema)
	Teams        datatypes.JSON `gorm:"column:teams;default:'[]';not null" json:"teams"`
	IsComplete   bool           `gorm:"column:isComplete;default:false;not null" json:"isComplete"`
	CreatedBy    *User          `gorm:"foreignKey:CreatedByID;belongsTo"`
	Organization *Team          `gorm:"foreignKey:OrganizationID;belongsTo"`
}

func (OrganizationOnboarding) TableName() string {
	return "OrganizationOnboarding"
}

func (m *OrganizationOnboarding) BeforeCreate(tx *gorm.DB) error {
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

func (*OrganizationOnboarding) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type InternalNotePreset struct {
	ID                  int                   `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name                string                `gorm:"column:name;uniqueIndex:InternalNotePreset_teamId_name_key;not null" json:"name"`
	CancellationReason  *string               `gorm:"column:cancellationReason" json:"cancellationReason"`
	TeamID              int                   `gorm:"column:teamId;uniqueIndex:InternalNotePreset_teamId_name_key;index:InternalNotePreset_teamId_idx;size:32;not null" json:"teamId"`
	CreatedAt           DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Team                *Team                 `gorm:"belongsTo"`
	BookingInternalNote []BookingInternalNote `gorm:"foreignKey:NotePresetID;constraint:BookingInternalNote_notePresetId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (InternalNotePreset) TableName() string {
	return "InternalNotePreset"
}

func (m *InternalNotePreset) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type FilterSegment struct {
	ID               int                           `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name             string                        `gorm:"column:name;not null" json:"name"`
	TableIdentifier  string                        `gorm:"column:tableIdentifier;index:FilterSegment_scope_userId_tableIdentifier_idx;index:FilterSegment_scope_teamId_tableIdentifier_idx;not null" json:"tableIdentifier"`
	Scope            string                        `gorm:"column:scope;index:FilterSegment_scope_userId_tableIdentifier_idx;index:FilterSegment_scope_teamId_tableIdentifier_idx;not null" json:"scope"`
	ActiveFilters    datatypes.JSON                `gorm:"column:activeFilters" json:"activeFilters"`
	Sorting          datatypes.JSON                `gorm:"column:sorting" json:"sorting"`
	ColumnVisibility datatypes.JSON                `gorm:"column:columnVisibility" json:"columnVisibility"`
	ColumnSizing     datatypes.JSON                `gorm:"column:columnSizing" json:"columnSizing"`
	PerPage          int                           `gorm:"column:perPage;size:32;not null" json:"perPage"`
	SearchTerm       *string                       `gorm:"column:searchTerm;type:text" json:"searchTerm"`
	CreatedAt        DateTime                      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime                      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID           int                           `gorm:"column:userId;index:FilterSegment_scope_userId_tableIdentifier_idx;size:32;not null" json:"userId"`
	TeamID           *int                          `gorm:"column:teamId;index:FilterSegment_scope_teamId_tableIdentifier_idx;size:32" json:"teamId"`
	User             *User                         `gorm:"belongsTo"`
	Team             *Team                         `gorm:"belongsTo"`
	UserPreferences  []UserFilterSegmentPreference `gorm:"foreignKey:SegmentID;constraint:UserFilterSegmentPreference_segmentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (FilterSegment) TableName() string {
	return "FilterSegment"
}

func (m *FilterSegment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*FilterSegment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type UserFilterSegmentPreference struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID          int            `gorm:"column:userId;uniqueIndex:UserFilterSegmentPreference_userId_tableIdentifier_key;index:UserFilterSegmentPreference_userId_idx;size:32;not null" json:"userId"`
	TableIdentifier string         `gorm:"column:tableIdentifier;uniqueIndex:UserFilterSegmentPreference_userId_tableIdentifier_key;not null" json:"tableIdentifier"`
	SegmentID       *int           `gorm:"column:segmentId;index:UserFilterSegmentPreference_segmentId_idx;size:32" json:"segmentId"`
	SystemSegmentID *string        `gorm:"column:systemSegmentId" json:"systemSegmentId"`
	CreatedAt       DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User            *User          `gorm:"belongsTo"`
	Segment         *FilterSegment `gorm:"foreignKey:SegmentID;belongsTo"`
}

func (UserFilterSegmentPreference) TableName() string {
	return "UserFilterSegmentPreference"
}

func (m *UserFilterSegmentPreference) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*UserFilterSegmentPreference) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type BookingInternalNote struct {
	ID           int                 `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	NotePresetID *int                `gorm:"column:notePresetId;uniqueIndex:BookingInternalNote_bookingId_notePresetId_key;size:32" json:"notePresetId"`
	Text         *string             `gorm:"column:text" json:"text"`
	BookingID    int                 `gorm:"column:bookingId;uniqueIndex:BookingInternalNote_bookingId_notePresetId_key;index:BookingInternalNote_bookingId_idx;size:32;not null" json:"bookingId"`
	CreatedByID  int                 `gorm:"column:createdById;size:32;not null" json:"createdById"`
	CreatedAt    DateTime            `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	NotePreset   *InternalNotePreset `gorm:"foreignKey:NotePresetID;belongsTo"`
	Booking      *Booking            `gorm:"belongsTo"`
	CreatedBy    *User               `gorm:"foreignKey:CreatedByID;belongsTo"`
}

func (BookingInternalNote) TableName() string {
	return "BookingInternalNote"
}

func (m *BookingInternalNote) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Role struct {
	ID          string           `gorm:"column:id;primaryKey" json:"id"`
	Name        string           `gorm:"column:name;uniqueIndex:Role_name_teamId_key;not null" json:"name"`
	Color       *string          `gorm:"column:color" json:"color"`
	Description *string          `gorm:"column:description" json:"description"`
	TeamID      *int             `gorm:"column:teamId;uniqueIndex:Role_name_teamId_key;index:Role_teamId_idx;size:32" json:"teamId"`
	CreatedAt   DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime         `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Type        string           `gorm:"column:type;default:'CUSTOM';not null" json:"type"`
	Team        *Team            `gorm:"belongsTo"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID;constraint:RolePermission_roleId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Memberships []Membership     `gorm:"foreignKey:CustomRoleID;constraint:Membership_customRoleId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
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

type RolePermission struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	RoleID    string   `gorm:"column:roleId;uniqueIndex:RolePermission_roleId_resource_action_key;index:RolePermission_roleId_idx;not null" json:"roleId"`
	Resource  string   `gorm:"column:resource;uniqueIndex:RolePermission_roleId_resource_action_key;not null" json:"resource"`
	Action    string   `gorm:"column:action;uniqueIndex:RolePermission_roleId_resource_action_key;index:RolePermission_action_idx;not null" json:"action"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Role      *Role    `gorm:"belongsTo"`
}

func (RolePermission) TableName() string {
	return "RolePermission"
}

func (m *RolePermission) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type AuditActor struct {
	ID            string         `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	Type          string         `gorm:"column:type;not null" json:"type"`
	UserUUID      *string        `gorm:"column:userUuid;uniqueIndex:AuditActor_userUuid_key;index:AuditActor_userUuid_idx;type:char(36)" json:"userUuid"`
	AttendeeID    *int           `gorm:"column:attendeeId;uniqueIndex:AuditActor_attendeeId_key;index:AuditActor_attendeeId_idx;size:32" json:"attendeeId"`
	CredentialID  *int           `gorm:"column:credentialId;uniqueIndex:AuditActor_credentialId_key;index:AuditActor_credentialId_idx;size:32" json:"credentialId"`
	Email         *string        `gorm:"column:email;uniqueIndex:AuditActor_email_key;index:AuditActor_email_idx" json:"email"`
	Phone         *string        `gorm:"column:phone;uniqueIndex:AuditActor_phone_key" json:"phone"`
	Name          *string        `gorm:"column:name" json:"name"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	BookingAudits []BookingAudit `gorm:"foreignKey:ActorID;constraint:BookingAudit_actorId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (AuditActor) TableName() string {
	return "AuditActor"
}

func (m *AuditActor) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.Must(uuid.NewV7()).String()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type BookingAudit struct {
	ID          string         `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	BookingUID  string         `gorm:"column:bookingUid;index:BookingAudit_bookingUid_idx;not null" json:"bookingUid"`
	ActorID     string         `gorm:"column:actorId;index:BookingAudit_actorId_idx;type:char(36);not null" json:"actorId"`
	Type        string         `gorm:"column:type;not null" json:"type"`
	Action      string         `gorm:"column:action;not null" json:"action"`
	Timestamp   DateTime       `gorm:"column:timestamp;index:BookingAudit_timestamp_idx;type:timestamp(3);not null" json:"timestamp"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Source      string         `gorm:"column:source;not null" json:"source"`
	OperationID string         `gorm:"column:operationId;index:BookingAudit_operationId_idx;not null" json:"operationId"`
	Data        datatypes.JSON `gorm:"column:data" json:"data"`
	Context     datatypes.JSON `gorm:"column:context" json:"context"`
	Actor       *AuditActor    `gorm:"foreignKey:ActorID;belongsTo"`
}

func (BookingAudit) TableName() string {
	return "BookingAudit"
}

func (m *BookingAudit) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.Must(uuid.NewV7()).String()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*BookingAudit) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Agent struct {
	ID                   string             `gorm:"column:id;primaryKey" json:"id"`
	Name                 string             `gorm:"column:name;not null" json:"name"`
	UserID               *int               `gorm:"column:userId;index:Agent_userId_idx;size:32" json:"userId"`
	TeamID               *int               `gorm:"column:teamId;index:Agent_teamId_idx;size:32" json:"teamId"`
	ProviderAgentID      string             `gorm:"column:providerAgentId;uniqueIndex:Agent_providerAgentId_key;not null" json:"providerAgentId"`
	InboundEventTypeID   *int               `gorm:"column:inboundEventTypeId;index:Agent_inboundEventTypeId_idx;size:32" json:"inboundEventTypeId"`
	OutboundEventTypeID  *int               `gorm:"column:outboundEventTypeId;index:Agent_outboundEventTypeId_idx;size:32" json:"outboundEventTypeId"`
	Enabled              *bool              `gorm:"column:enabled;default:true;not null" json:"enabled"`
	CreatedAt            DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User                 *User              `gorm:"belongsTo"`
	Team                 *Team              `gorm:"belongsTo"`
	InboundPhoneNumbers  []CalAiPhoneNumber `gorm:"foreignKey:InboundAgentID;constraint:CalAiPhoneNumber_inboundAgentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	OutboundPhoneNumbers []CalAiPhoneNumber `gorm:"foreignKey:OutboundAgentID;constraint:CalAiPhoneNumber_outboundAgentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Agent) TableName() string {
	return "Agent"
}

func (m *Agent) BeforeCreate(tx *gorm.DB) error {
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

func (*Agent) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CalAiPhoneNumber struct {
	ID                    int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID                *int     `gorm:"column:userId;index:CalAiPhoneNumber_userId_idx;size:32" json:"userId"`
	TeamID                *int     `gorm:"column:teamId;index:CalAiPhoneNumber_teamId_idx;size:32" json:"teamId"`
	PhoneNumber           string   `gorm:"column:phoneNumber;uniqueIndex:CalAiPhoneNumber_phoneNumber_key;not null" json:"phoneNumber"`
	Provider              string   `gorm:"column:provider;not null" json:"provider"`
	ProviderPhoneNumberID *string  `gorm:"column:providerPhoneNumberId;uniqueIndex:CalAiPhoneNumber_providerPhoneNumberId_key" json:"providerPhoneNumberId"`
	CreatedAt             DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	StripeCustomerID      *string  `gorm:"column:stripeCustomerId" json:"stripeCustomerId"`
	StripeSubscriptionID  *string  `gorm:"column:stripeSubscriptionId;uniqueIndex:CalAiPhoneNumber_stripeSubscriptionId_key" json:"stripeSubscriptionId"`
	SubscriptionStatus    *string  `gorm:"column:subscriptionStatus" json:"subscriptionStatus"`
	InboundAgentID        *string  `gorm:"column:inboundAgentId;index:CalAiPhoneNumber_inboundAgentId_idx" json:"inboundAgentId"`
	OutboundAgentID       *string  `gorm:"column:outboundAgentId;index:CalAiPhoneNumber_outboundAgentId_idx" json:"outboundAgentId"`
	User                  *User    `gorm:"belongsTo"`
	Team                  *Team    `gorm:"belongsTo"`
	InboundAgent          *Agent   `gorm:"foreignKey:InboundAgentID;belongsTo"`
	OutboundAgent         *Agent   `gorm:"foreignKey:OutboundAgentID;belongsTo"`
}

func (CalAiPhoneNumber) TableName() string {
	return "CalAiPhoneNumber"
}

func (m *CalAiPhoneNumber) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*CalAiPhoneNumber) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TeamBilling struct {
	ID                       string             `gorm:"column:id;primaryKey" json:"id"`
	TeamID                   int                `gorm:"column:teamId;uniqueIndex:TeamBilling_teamId_key;size:32;not null" json:"teamId"`
	SubscriptionID           string             `gorm:"column:subscriptionId;uniqueIndex:TeamBilling_subscriptionId_key;not null" json:"subscriptionId"`
	SubscriptionItemID       string             `gorm:"column:subscriptionItemId;not null" json:"subscriptionItemId"`
	CustomerID               string             `gorm:"column:customerId;not null" json:"customerId"`
	Status                   string             `gorm:"column:status;not null" json:"status"`
	PlanName                 string             `gorm:"column:planName;not null" json:"planName"`
	SubscriptionStart        *DateTime          `gorm:"column:subscriptionStart;type:timestamp(3)" json:"subscriptionStart"`
	SubscriptionTrialEnd     *DateTime          `gorm:"column:subscriptionTrialEnd;type:timestamp(3)" json:"subscriptionTrialEnd"`
	SubscriptionEnd          *DateTime          `gorm:"column:subscriptionEnd;type:timestamp(3)" json:"subscriptionEnd"`
	BillingPeriod            *string            `gorm:"column:billingPeriod" json:"billingPeriod"`
	BillingMode              string             `gorm:"column:billingMode;default:'SEATS';not null" json:"billingMode"`
	PricePerSeat             *int               `gorm:"column:pricePerSeat;size:32" json:"pricePerSeat"`
	PaidSeats                *int               `gorm:"column:paidSeats;size:32" json:"paidSeats"`
	HighWaterMark            *int               `gorm:"column:highWaterMark;size:32" json:"highWaterMark"`
	HighWaterMarkPeriodStart *DateTime          `gorm:"column:highWaterMarkPeriodStart;type:timestamp(3)" json:"highWaterMarkPeriodStart"`
	CreatedAt                DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team                     *Team              `gorm:"belongsTo"`
	MonthlyProrations        []MonthlyProration `gorm:"foreignKey:TeamBillingID;constraint:MonthlyProration_teamBillingId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	SeatChangeLogs           []SeatChangeLog    `gorm:"foreignKey:TeamBillingID;constraint:SeatChangeLog_teamBillingId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (TeamBilling) TableName() string {
	return "TeamBilling"
}

func (m *TeamBilling) BeforeCreate(tx *gorm.DB) error {
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

func (*TeamBilling) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OrganizationBilling struct {
	ID                       string             `gorm:"column:id;primaryKey" json:"id"`
	TeamID                   int                `gorm:"column:teamId;uniqueIndex:OrganizationBilling_teamId_key;size:32;not null" json:"teamId"`
	SubscriptionID           string             `gorm:"column:subscriptionId;uniqueIndex:OrganizationBilling_subscriptionId_key;not null" json:"subscriptionId"`
	SubscriptionItemID       string             `gorm:"column:subscriptionItemId;not null" json:"subscriptionItemId"`
	CustomerID               string             `gorm:"column:customerId;not null" json:"customerId"`
	Status                   string             `gorm:"column:status;not null" json:"status"`
	PlanName                 string             `gorm:"column:planName;not null" json:"planName"`
	SubscriptionStart        *DateTime          `gorm:"column:subscriptionStart;type:timestamp(3)" json:"subscriptionStart"`
	SubscriptionTrialEnd     *DateTime          `gorm:"column:subscriptionTrialEnd;type:timestamp(3)" json:"subscriptionTrialEnd"`
	SubscriptionEnd          *DateTime          `gorm:"column:subscriptionEnd;type:timestamp(3)" json:"subscriptionEnd"`
	BillingPeriod            *string            `gorm:"column:billingPeriod" json:"billingPeriod"`
	BillingMode              string             `gorm:"column:billingMode;default:'SEATS';not null" json:"billingMode"`
	PricePerSeat             *int               `gorm:"column:pricePerSeat;size:32" json:"pricePerSeat"`
	PaidSeats                *int               `gorm:"column:paidSeats;size:32" json:"paidSeats"`
	HighWaterMark            *int               `gorm:"column:highWaterMark;size:32" json:"highWaterMark"`
	HighWaterMarkPeriodStart *DateTime          `gorm:"column:highWaterMarkPeriodStart;type:timestamp(3)" json:"highWaterMarkPeriodStart"`
	CreatedAt                DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team                     *Team              `gorm:"belongsTo"`
	MonthlyProrations        []MonthlyProration `gorm:"foreignKey:OrganizationBillingID;constraint:MonthlyProration_organizationBillingId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	SeatChangeLogs           []SeatChangeLog    `gorm:"foreignKey:OrganizationBillingID;constraint:SeatChangeLog_organizationBillingId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (OrganizationBilling) TableName() string {
	return "OrganizationBilling"
}

func (m *OrganizationBilling) BeforeCreate(tx *gorm.DB) error {
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

func (*OrganizationBilling) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type SeatChangeLog struct {
	ID                     string               `gorm:"column:id;primaryKey" json:"id"`
	TeamID                 int                  `gorm:"column:teamId;uniqueIndex:SeatChangeLog_teamId_operationId_key;index:SeatChangeLog_teamId_monthKey_idx;index:SeatChangeLog_teamId_processedInProrationId_idx;size:32;not null" json:"teamId"`
	ChangeType             string               `gorm:"column:changeType;not null" json:"changeType"`
	SeatCount              int                  `gorm:"column:seatCount;size:32;not null" json:"seatCount"`
	UserID                 *int                 `gorm:"column:userId;size:32" json:"userId"`
	TriggeredBy            *int                 `gorm:"column:triggeredBy;size:32" json:"triggeredBy"`
	ChangeDate             DateTime             `gorm:"column:changeDate;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"changeDate"`
	MonthKey               string               `gorm:"column:monthKey;index:SeatChangeLog_teamId_monthKey_idx;index:SeatChangeLog_monthKey_idx;not null" json:"monthKey"`
	OperationID            *string              `gorm:"column:operationId;uniqueIndex:SeatChangeLog_teamId_operationId_key" json:"operationId"`
	ProcessedInProrationID *string              `gorm:"column:processedInProrationId;index:SeatChangeLog_teamId_processedInProrationId_idx" json:"processedInProrationId"`
	Metadata               datatypes.JSON       `gorm:"column:metadata" json:"metadata"`
	CreatedAt              DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	TeamBillingID          *string              `gorm:"column:teamBillingId" json:"teamBillingId"`
	OrganizationBillingID  *string              `gorm:"column:organizationBillingId" json:"organizationBillingId"`
	Team                   *Team                `gorm:"belongsTo"`
	Proration              *MonthlyProration    `gorm:"foreignKey:ProcessedInProrationID;belongsTo"`
	TeamBilling            *TeamBilling         `gorm:"belongsTo"`
	OrganizationBilling    *OrganizationBilling `gorm:"belongsTo"`
}

func (SeatChangeLog) TableName() string {
	return "SeatChangeLog"
}

func (m *SeatChangeLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.ChangeDate.IsZero() {
		m.ChangeDate = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type MonthlyProration struct {
	ID                    string               `gorm:"column:id;primaryKey" json:"id"`
	TeamID                int                  `gorm:"column:teamId;uniqueIndex:MonthlyProration_teamId_monthKey_key;index:MonthlyProration_teamId_idx;size:32;not null" json:"teamId"`
	MonthKey              string               `gorm:"column:monthKey;uniqueIndex:MonthlyProration_teamId_monthKey_key;index:MonthlyProration_monthKey_status_idx;not null" json:"monthKey"`
	PeriodStart           DateTime             `gorm:"column:periodStart;type:timestamp(3);not null" json:"periodStart"`
	PeriodEnd             DateTime             `gorm:"column:periodEnd;type:timestamp(3);not null" json:"periodEnd"`
	SeatsAtStart          int                  `gorm:"column:seatsAtStart;size:32;not null" json:"seatsAtStart"`
	SeatsAdded            int                  `gorm:"column:seatsAdded;size:32;not null" json:"seatsAdded"`
	SeatsRemoved          int                  `gorm:"column:seatsRemoved;size:32;not null" json:"seatsRemoved"`
	NetSeatIncrease       int                  `gorm:"column:netSeatIncrease;size:32;not null" json:"netSeatIncrease"`
	SeatsAtEnd            int                  `gorm:"column:seatsAtEnd;size:32;not null" json:"seatsAtEnd"`
	SubscriptionID        string               `gorm:"column:subscriptionId;not null" json:"subscriptionId"`
	SubscriptionItemID    string               `gorm:"column:subscriptionItemId;not null" json:"subscriptionItemId"`
	CustomerID            string               `gorm:"column:customerId;not null" json:"customerId"`
	SubscriptionStart     DateTime             `gorm:"column:subscriptionStart;type:timestamp(3);not null" json:"subscriptionStart"`
	SubscriptionEnd       DateTime             `gorm:"column:subscriptionEnd;type:timestamp(3);not null" json:"subscriptionEnd"`
	RemainingDays         int                  `gorm:"column:remainingDays;size:32;not null" json:"remainingDays"`
	PricePerSeat          int                  `gorm:"column:pricePerSeat;size:32;not null" json:"pricePerSeat"`
	ProratedAmount        int                  `gorm:"column:proratedAmount;size:32;not null" json:"proratedAmount"`
	InvoiceItemID         *string              `gorm:"column:invoiceItemId" json:"invoiceItemId"`
	InvoiceID             *string              `gorm:"column:invoiceId" json:"invoiceId"`
	InvoiceURL            *string              `gorm:"column:invoiceUrl" json:"invoiceUrl"`
	Status                string               `gorm:"column:status;index:MonthlyProration_monthKey_status_idx;index:MonthlyProration_status_idx;default:'PENDING';not null" json:"status"`
	ChargedAt             *DateTime            `gorm:"column:chargedAt;type:timestamp(3)" json:"chargedAt"`
	FailedAt              *DateTime            `gorm:"column:failedAt;type:timestamp(3)" json:"failedAt"`
	FailureReason         *string              `gorm:"column:failureReason" json:"failureReason"`
	RetryCount            int                  `gorm:"column:retryCount;size:32;default:0;not null" json:"retryCount"`
	Metadata              datatypes.JSON       `gorm:"column:metadata" json:"metadata"`
	CreatedAt             DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime             `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	TeamBillingID         *string              `gorm:"column:teamBillingId" json:"teamBillingId"`
	OrganizationBillingID *string              `gorm:"column:organizationBillingId" json:"organizationBillingId"`
	Team                  *Team                `gorm:"belongsTo"`
	TeamBilling           *TeamBilling         `gorm:"belongsTo"`
	OrganizationBilling   *OrganizationBilling `gorm:"belongsTo"`
	SeatChanges           []SeatChangeLog      `gorm:"foreignKey:ProcessedInProrationID;constraint:SeatChangeLog_processedInProrationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (MonthlyProration) TableName() string {
	return "MonthlyProration"
}

func (m *MonthlyProration) BeforeCreate(tx *gorm.DB) error {
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

func (*MonthlyProration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type CalendarCacheEvent struct {
	ID                 string            `gorm:"column:id;primaryKey" json:"id"`
	SelectedCalendarID string            `gorm:"column:selectedCalendarId;uniqueIndex:CalendarCacheEvent_selectedCalendarId_externalId_key;index:CalendarCacheEvent_selectedCalendarId_iCalUID_idx;not null" json:"selectedCalendarId"`
	ExternalID         string            `gorm:"column:externalId;uniqueIndex:CalendarCacheEvent_selectedCalendarId_externalId_key;not null" json:"externalId"`
	ExternalEtag       string            `gorm:"column:externalEtag;not null" json:"externalEtag"`
	ICalUID            *string           `gorm:"column:iCalUID;index:CalendarCacheEvent_selectedCalendarId_iCalUID_idx" json:"iCalUID"`
	ICalSequence       int               `gorm:"column:iCalSequence;size:32;default:0;not null" json:"iCalSequence"`
	Summary            *string           `gorm:"column:summary" json:"summary"`
	Description        *string           `gorm:"column:description" json:"description"`
	Location           *string           `gorm:"column:location" json:"location"`
	Start              DateTime          `gorm:"column:start;index:CalendarCacheEvent_start_end_status_idx;type:timestamp(3);not null" json:"start"`
	End                DateTime          `gorm:"column:end;index:CalendarCacheEvent_start_end_status_idx;type:timestamp(3);not null" json:"end"`
	IsAllDay           bool              `gorm:"column:isAllDay;default:false;not null" json:"isAllDay"`
	TimeZone           *string           `gorm:"column:timeZone" json:"timeZone"`
	Status             string            `gorm:"column:status;index:CalendarCacheEvent_start_end_status_idx;default:'confirmed';not null" json:"status"`
	RecurringEventID   *string           `gorm:"column:recurringEventId" json:"recurringEventId"`
	OriginalStartTime  *DateTime         `gorm:"column:originalStartTime;type:timestamp(3)" json:"originalStartTime"`
	CreatedAt          DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	ExternalCreatedAt  *DateTime         `gorm:"column:externalCreatedAt;type:timestamp(3)" json:"externalCreatedAt"`
	ExternalUpdatedAt  *DateTime         `gorm:"column:externalUpdatedAt;type:timestamp(3)" json:"externalUpdatedAt"`
	SelectedCalendar   *SelectedCalendar `gorm:"belongsTo"`
}

func (CalendarCacheEvent) TableName() string {
	return "CalendarCacheEvent"
}

func (m *CalendarCacheEvent) BeforeCreate(tx *gorm.DB) error {
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

func (*CalendarCacheEvent) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type IntegrationAttributeSync struct {
	ID                string                      `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID    int                         `gorm:"column:organizationId;index:IntegrationAttributeSync_organizationId_idx;size:32;not null" json:"organizationId"`
	Name              string                      `gorm:"column:name;not null" json:"name"`
	Integration       string                      `gorm:"column:integration;not null" json:"integration"`
	CredentialID      *int                        `gorm:"column:credentialId;index:IntegrationAttributeSync_credentialId_idx;size:32" json:"credentialId"`
	Enabled           bool                        `gorm:"column:enabled;not null" json:"enabled"`
	CreatedAt         DateTime                    `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime                    `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Organization      *Team                       `gorm:"foreignKey:OrganizationID;belongsTo"`
	Credential        *Credential                 `gorm:"belongsTo"`
	SyncFieldMappings []AttributeSyncFieldMapping `gorm:"foreignKey:IntegrationAttributeSyncID;constraint:AttributeSyncFieldMapping_integrationAttributeSyncId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AttributeSyncRule *AttributeSyncRule          `gorm:"foreignKey:IntegrationAttributeSyncID;constraint:AttributeSyncRule_integrationAttributeSyncId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (IntegrationAttributeSync) TableName() string {
	return "IntegrationAttributeSync"
}

func (m *IntegrationAttributeSync) BeforeCreate(tx *gorm.DB) error {
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

func (*IntegrationAttributeSync) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AttributeSyncRule struct {
	ID                         string                    `gorm:"column:id;primaryKey" json:"id"`
	IntegrationAttributeSyncID string                    `gorm:"column:integrationAttributeSyncId;uniqueIndex:AttributeSyncRule_integrationAttributeSyncId_key;not null" json:"integrationAttributeSyncId"`
	Rule                       datatypes.JSON            `gorm:"column:rule;not null" json:"rule"`
	CreatedAt                  DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime                  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	IntegrationAttributeSync   *IntegrationAttributeSync `gorm:"belongsTo"`
}

func (AttributeSyncRule) TableName() string {
	return "AttributeSyncRule"
}

func (m *AttributeSyncRule) BeforeCreate(tx *gorm.DB) error {
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

func (*AttributeSyncRule) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AttributeSyncFieldMapping struct {
	ID                         string                    `gorm:"column:id;primaryKey" json:"id"`
	IntegrationFieldName       string                    `gorm:"column:integrationFieldName;not null" json:"integrationFieldName"`
	AttributeID                string                    `gorm:"column:attributeId;uniqueIndex:AttributeSyncFieldMapping_integrationAttributeSyncId_attrib_key;not null" json:"attributeId"`
	Enabled                    bool                      `gorm:"column:enabled;not null" json:"enabled"`
	IntegrationAttributeSyncID string                    `gorm:"column:integrationAttributeSyncId;uniqueIndex:AttributeSyncFieldMapping_integrationAttributeSyncId_attrib_key;index:AttributeSyncFieldMapping_integrationAttributeSyncId_idx;not null" json:"integrationAttributeSyncId"`
	CreatedAt                  DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime                  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Attribute                  *Attribute                `gorm:"belongsTo"`
	IntegrationAttributeSync   *IntegrationAttributeSync `gorm:"belongsTo"`
}

func (AttributeSyncFieldMapping) TableName() string {
	return "AttributeSyncFieldMapping"
}

func (m *AttributeSyncFieldMapping) BeforeCreate(tx *gorm.DB) error {
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

func (*AttributeSyncFieldMapping) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// BookingTimeStatus is a database view: read it, do not write to it, and do not pass it to
// AutoMigrate, which would create a table of that name.
type BookingTimeStatus struct {
	ID             int       `gorm:"column:id;primaryKey;uniqueIndex:BookingTimeStatus_id_key;size:32" json:"id"`
	UID            *string   `gorm:"column:uid" json:"uid"`
	EventTypeID    *int      `gorm:"column:eventTypeId;size:32" json:"eventTypeId"`
	Title          *string   `gorm:"column:title" json:"title"`
	Description    *string   `gorm:"column:description" json:"description"`
	StartTime      *DateTime `gorm:"column:startTime;type:timestamp(3)" json:"startTime"`
	EndTime        *DateTime `gorm:"column:endTime;type:timestamp(3)" json:"endTime"`
	CreatedAt      *DateTime `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false" json:"createdAt"`
	Location       *string   `gorm:"column:location" json:"location"`
	Paid           *bool     `gorm:"column:paid" json:"paid"`
	Status         *string   `gorm:"column:status" json:"status"`
	Rescheduled    *bool     `gorm:"column:rescheduled" json:"rescheduled"`
	UserID         *int      `gorm:"column:userId;size:32" json:"userId"`
	TeamID         *int      `gorm:"column:teamId;size:32" json:"teamId"`
	EventLength    *int      `gorm:"column:eventLength;size:32" json:"eventLength"`
	TimeStatus     *string   `gorm:"column:timeStatus" json:"timeStatus"`
	EventParentID  *int      `gorm:"column:eventParentId;size:32" json:"eventParentId"`
	UserEmail      *string   `gorm:"column:userEmail" json:"userEmail"`
	Username       *string   `gorm:"column:username" json:"username"`
	RatingFeedback *string   `gorm:"column:ratingFeedback" json:"ratingFeedback"`
	Rating         *int      `gorm:"column:rating;size:32" json:"rating"`
	NoShowHost     *bool     `gorm:"column:noShowHost" json:"noShowHost"`
	IsTeamBooking  bool      `gorm:"column:isTeamBooking;not null" json:"isTeamBooking"`
}

func (BookingTimeStatus) TableName() string {
	return "BookingTimeStatus"
}

// BookingTimeStatusDenormalized is a database view: read it, do not write to it, and do not pass it to
// AutoMigrate, which would create a table of that name.
type BookingTimeStatusDenormalized struct {
	ID             int       `gorm:"column:id;primaryKey;uniqueIndex:BookingTimeStatusDenormalized_id_key;size:32" json:"id"`
	UID            string    `gorm:"column:uid;not null" json:"uid"`
	EventTypeID    *int      `gorm:"column:eventTypeId;size:32" json:"eventTypeId"`
	Title          string    `gorm:"column:title;not null" json:"title"`
	Description    *string   `gorm:"column:description" json:"description"`
	StartTime      DateTime  `gorm:"column:startTime;type:timestamp(3);not null" json:"startTime"`
	EndTime        DateTime  `gorm:"column:endTime;type:timestamp(3);not null" json:"endTime"`
	CreatedAt      DateTime  `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      *DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	Location       *string   `gorm:"column:location" json:"location"`
	Paid           bool      `gorm:"column:paid;not null" json:"paid"`
	Status         string    `gorm:"column:status;not null" json:"status"`
	Rescheduled    *bool     `gorm:"column:rescheduled" json:"rescheduled"`
	UserID         *int      `gorm:"column:userId;size:32" json:"userId"`
	TeamID         *int      `gorm:"column:teamId;size:32" json:"teamId"`
	EventLength    *int      `gorm:"column:eventLength;size:32" json:"eventLength"`
	EventParentID  *int      `gorm:"column:eventParentId;size:32" json:"eventParentId"`
	UserEmail      *string   `gorm:"column:userEmail" json:"userEmail"`
	UserName       *string   `gorm:"column:userName" json:"userName"`
	UserUsername   *string   `gorm:"column:userUsername" json:"userUsername"`
	RatingFeedback *string   `gorm:"column:ratingFeedback" json:"ratingFeedback"`
	Rating         *int      `gorm:"column:rating;size:32" json:"rating"`
	NoShowHost     *bool     `gorm:"column:noShowHost" json:"noShowHost"`
	IsTeamBooking  bool      `gorm:"column:isTeamBooking;not null" json:"isTeamBooking"`
	TimeStatus     *string   `gorm:"column:timeStatus" json:"timeStatus"`
}

func (BookingTimeStatusDenormalized) TableName() string {
	return "BookingTimeStatusDenormalized"
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

type TimeOfDay struct{ time.Time }

func (TimeOfDay) GormDataType() string {
	return "time"
}

func (clock TimeOfDay) Value() (driver.Value, error) {
	return clock.UTC().Truncate(time.Millisecond).Format("15:04:05.000"), nil
}

func (clock *TimeOfDay) Scan(src any) error {
	switch value := src.(type) {
	case nil:
		clock.Time = time.Time{}
		return nil
	case []byte:
		return clock.Scan(string(value))
	case string:
		if end := strings.IndexAny(value, "+-Z"); end > 0 {
			value = value[:end]
		}
		read, err := time.Parse("15:04:05.999999999", value)
		if err != nil {
			return fmt.Errorf("a time column held %q", value)
		}
		clock.Time = time.Date(1970, 1, 1, read.Hour(), read.Minute(), read.Second(), read.Nanosecond(), time.UTC).Truncate(time.Millisecond)
		return nil
	}
	read, err := readPrismaTime(src)
	clock.Time = time.Date(1970, 1, 1, read.Hour(), read.Minute(), read.Second(), read.Nanosecond(), time.UTC)
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
