package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lucsky/cuid"
	"github.com/nrednav/cuid2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// A feedback button or a banner. Opens at most one Destination: a survey (`surveyId`) or a link
// (`linkUrl`), never both.
//
// @property type - Fixed at creation.
// @property content - Per-type JSON (`FeedbackButtonContent` / `BannerContent` in the v3 spec), texts keyed by locale code.
// @property styling - Per-type JSON (`FeedbackButtonStyling` / `BannerStyling` in the v3 spec).
// @property survey - The Destination survey, in the same workspace by construction of the composite key.
// @property showConditions - Show when (`ZShowConditions`, shared with surveys); null always shows.
// @property autoSelectLanguage - The browser's language picks the texts' language when the site sets none; null (never chosen, as on Survey) is off.
type Creative struct {
	ID                 string                  `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt          DateTime                `gorm:"column:created_at;index:Creative_workspaceId_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt          DateTime                `gorm:"column:updated_at;index:Creative_workspaceId_updated_at_idx,priority:2;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name               string                  `gorm:"column:name;index:Creative_workspaceId_name_idx,priority:2;not null" json:"name"`
	Type               string                  `gorm:"column:type;not null" json:"type"`
	Status             string                  `gorm:"column:status;default:'draft';not null" json:"status"`
	PublishedAt        *DateTime               `gorm:"column:publishedAt;type:timestamp(3)" json:"publishedAt"`
	ArchivedAt         *DateTime               `gorm:"column:archivedAt;type:timestamp(3)" json:"archivedAt"`
	WorkspaceID        string                  `gorm:"column:workspaceId;index:Creative_workspaceId_updated_at_idx,priority:1;index:Creative_workspaceId_created_at_idx,priority:1;index:Creative_workspaceId_name_idx,priority:1;not null" json:"workspaceId"`
	CreatedBy          *string                 `gorm:"column:createdBy" json:"createdBy"`
	Content            datatypes.JSON          `gorm:"column:content;not null" json:"content"`
	Styling            datatypes.JSON          `gorm:"column:styling;not null" json:"styling"`
	SurveyID           *string                 `gorm:"column:surveyId;index:Creative_surveyId_idx" json:"surveyId"`
	LinkURL            *string                 `gorm:"column:linkUrl" json:"linkUrl"`
	ShowOn             string                  `gorm:"column:showOn;default:'pageView';not null" json:"showOn"`
	ShowConditions     datatypes.JSON          `gorm:"column:showConditions" json:"showConditions"`
	DisplayPercentage  *float64                `gorm:"column:displayPercentage;type:double precision" json:"displayPercentage"`
	StartAt            *DateTime               `gorm:"column:startAt;type:timestamp(3)" json:"startAt"`
	EndAt              *DateTime               `gorm:"column:endAt;type:timestamp(3)" json:"endAt"`
	AutoSelectLanguage *bool                   `gorm:"column:autoSelectLanguage" json:"autoSelectLanguage"`
	Workspace          *Workspace              `gorm:"belongsTo"`
	Creator            *User                   `gorm:"foreignKey:CreatedBy;belongsTo"`
	Survey             *Survey                 `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
	Triggers           []CreativeTrigger       `gorm:"foreignKey:CreativeID;constraint:CreativeTrigger_creativeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Languages          []CreativeLanguage      `gorm:"foreignKey:CreativeID;constraint:CreativeLanguage_creativeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DailyStats         []CreativeDailyStat     `gorm:"foreignKey:CreativeID;constraint:CreativeDailyStat_creativeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DailyPageStats     []CreativeDailyPageStat `gorm:"foreignKey:CreativeID;constraint:CreativeDailyPageStat_creativeId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Creative) TableName() string {
	return "Creative"
}

func (m *Creative) BeforeCreate(tx *gorm.DB) error {
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

func (*Creative) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
// only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
// running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
// the creative deletes its counts.
type CreativeDailyStat struct {
	CreativeID string `gorm:"column:creativeId;primaryKey" json:"creativeId"`
	// UTC day.
	Date        Date      `gorm:"column:date;primaryKey;type:date" json:"date"`
	DeviceType  string    `gorm:"column:deviceType;primaryKey" json:"deviceType"`
	Impressions int       `gorm:"column:impressions;size:32;default:0;not null" json:"impressions"`
	Clicks      int       `gorm:"column:clicks;size:32;default:0;not null" json:"clicks"`
	Creative    *Creative `gorm:"belongsTo"`
}

func (CreativeDailyStat) TableName() string {
	return "CreativeDailyStat"
}

// A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
// `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
type CreativeDailyPageStat struct {
	CreativeID string `gorm:"column:creativeId;primaryKey" json:"creativeId"`
	// UTC day.
	Date Date `gorm:"column:date;primaryKey;type:date" json:"date"`
	// Host + path, without query string or hash, or `(other)`.
	Page        string    `gorm:"column:page;primaryKey" json:"page"`
	Impressions int       `gorm:"column:impressions;size:32;default:0;not null" json:"impressions"`
	Clicks      int       `gorm:"column:clicks;size:32;default:0;not null" json:"clicks"`
	Creative    *Creative `gorm:"belongsTo"`
}

func (CreativeDailyPageStat) TableName() string {
	return "CreativeDailyPageStat"
}

// Links a creative to the actions that show it (`showOn: action`). Mirrors `SurveyTrigger`.
type CreativeTrigger struct {
	ID            string       `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt     DateTime     `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt     DateTime     `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	CreativeID    string       `gorm:"column:creativeId;uniqueIndex:CreativeTrigger_creativeId_actionClassId_key,priority:1;not null" json:"creativeId"`
	ActionClassID string       `gorm:"column:actionClassId;uniqueIndex:CreativeTrigger_creativeId_actionClassId_key,priority:2;index:CreativeTrigger_actionClassId_idx;not null" json:"actionClassId"`
	Creative      *Creative    `gorm:"belongsTo"`
	ActionClass   *ActionClass `gorm:"belongsTo"`
}

func (CreativeTrigger) TableName() string {
	return "CreativeTrigger"
}

func (m *CreativeTrigger) BeforeCreate(tx *gorm.DB) error {
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

func (*CreativeTrigger) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// The languages a creative's texts are written in. Mirrors `SurveyLanguage`; exactly one is the default.
type CreativeLanguage struct {
	LanguageID string    `gorm:"column:languageId;primaryKey" json:"languageId"`
	CreativeID string    `gorm:"column:creativeId;primaryKey;index:CreativeLanguage_creativeId_idx" json:"creativeId"`
	Default    bool      `gorm:"column:default;default:false;not null" json:"default"`
	Enabled    *bool     `gorm:"column:enabled;default:true;not null" json:"enabled"`
	Language   *Language `gorm:"belongsTo"`
	Creative   *Creative `gorm:"belongsTo"`
}

func (CreativeLanguage) TableName() string {
	return "CreativeLanguage"
}

// Metadata for an outbound email attempt. Content and recipient addresses are never stored here.
// Identifiers are deliberately not foreign keys so erasure can scrub a row without retaining source data.
type OutboundEmailAudit struct {
	ID                    string    `gorm:"column:id;primaryKey;index:OutboundEmailAudit_environment_createdAt_id_idx,priority:3;index:OutboundEmailAudit_organizationId_createdAt_id_idx,priority:3;index:OutboundEmailAudit_surveyId_createdAt_id_idx,priority:3;index:OutboundEmailAudit_responseId_createdAt_id_idx,priority:3;index:OutboundEmailAudit_actorUserId_createdAt_id_idx,priority:3;index:OutboundEmailAudit_configuredByUserId_createdAt_id_idx,priority:3;index:OutboundEmailAudit_actorUserHash_createdAt_id_idx,priority:3;index:OutboundEmailAudit_configuredByUserHash_createdAt_id_idx,priority:3;index:OutboundEmailAudit_recipientMac_createdAt_id_idx,priority:3;index:OutboundEmailAudit_status_createdAt_id_idx,priority:3;index:OutboundEmailAudit_contentExpiresAt_id_idx,priority:2;index:OutboundEmailAudit_expiresAt_id_idx,priority:2;index:OutboundEmailAudit_erasureRequestedAt_id_idx,priority:2;index:OutboundEmailAudit_contentDeleteRetryAt_id_idx,priority:2" json:"id"`
	CreatedAt             DateTime  `gorm:"column:createdAt;index:OutboundEmailAudit_environment_createdAt_id_idx,priority:2;index:OutboundEmailAudit_organizationId_createdAt_id_idx,priority:2;index:OutboundEmailAudit_surveyId_createdAt_id_idx,priority:2;index:OutboundEmailAudit_responseId_createdAt_id_idx,priority:2;index:OutboundEmailAudit_actorUserId_createdAt_id_idx,priority:2;index:OutboundEmailAudit_configuredByUserId_createdAt_id_idx,priority:2;index:OutboundEmailAudit_actorUserHash_createdAt_id_idx,priority:2;index:OutboundEmailAudit_configuredByUserHash_createdAt_id_idx,priority:2;index:OutboundEmailAudit_recipientMac_createdAt_id_idx,priority:2;index:OutboundEmailAudit_status_createdAt_id_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	ExpiresAt             DateTime  `gorm:"column:expiresAt;index:OutboundEmailAudit_expiresAt_id_idx,priority:1;type:timestamp(3);not null" json:"expiresAt"`
	ContentExpiresAt      DateTime  `gorm:"column:contentExpiresAt;index:OutboundEmailAudit_contentExpiresAt_id_idx,priority:1;type:timestamp(3);not null" json:"contentExpiresAt"`
	Environment           string    `gorm:"column:environment;index:OutboundEmailAudit_environment_createdAt_id_idx,priority:1;not null" json:"environment"`
	EmailType             string    `gorm:"column:emailType;not null" json:"emailType"`
	Trigger               string    `gorm:"column:trigger;not null" json:"trigger"`
	ActorKind             string    `gorm:"column:actorKind;not null" json:"actorKind"`
	ActorUserID           *string   `gorm:"column:actorUserId;index:OutboundEmailAudit_actorUserId_createdAt_id_idx,priority:1" json:"actorUserId"`
	ConfiguredByUserID    *string   `gorm:"column:configuredByUserId;index:OutboundEmailAudit_configuredByUserId_createdAt_id_idx,priority:1" json:"configuredByUserId"`
	ActorUserHash         *string   `gorm:"column:actorUserHash;index:OutboundEmailAudit_actorUserHash_createdAt_id_idx,priority:1" json:"actorUserHash"`
	ConfiguredByUserHash  *string   `gorm:"column:configuredByUserHash;index:OutboundEmailAudit_configuredByUserHash_createdAt_id_idx,priority:1" json:"configuredByUserHash"`
	OrganizationID        *string   `gorm:"column:organizationId;index:OutboundEmailAudit_organizationId_createdAt_id_idx,priority:1" json:"organizationId"`
	SurveyID              *string   `gorm:"column:surveyId;index:OutboundEmailAudit_surveyId_createdAt_id_idx,priority:1" json:"surveyId"`
	ResponseID            *string   `gorm:"column:responseId;index:OutboundEmailAudit_responseId_createdAt_id_idx,priority:1" json:"responseId"`
	FollowUpID            *string   `gorm:"column:followUpId" json:"followUpId"`
	WorkflowID            *string   `gorm:"column:workflowId" json:"workflowId"`
	WorkflowRunID         *string   `gorm:"column:workflowRunId" json:"workflowRunId"`
	InviteID              *string   `gorm:"column:inviteId" json:"inviteId"`
	RecipientMac          *string   `gorm:"column:recipientMac;index:OutboundEmailAudit_recipientMac_createdAt_id_idx,priority:1" json:"recipientMac"`
	SesMessageHash        *string   `gorm:"column:sesMessageHash;index:OutboundEmailAudit_sesMessageHash_idx" json:"sesMessageHash"`
	SnapshotKey           *string   `gorm:"column:snapshotKey" json:"snapshotKey"`
	SnapshotSha256        *string   `gorm:"column:snapshotSha256" json:"snapshotSha256"`
	Status                *string   `gorm:"column:status;index:OutboundEmailAudit_status_createdAt_id_idx,priority:1;default:'prepared';not null" json:"status"`
	StatusRecordedAt      *DateTime `gorm:"column:statusRecordedAt;type:timestamp(3)" json:"statusRecordedAt"`
	AuditGap              bool      `gorm:"column:auditGap;default:false;not null" json:"auditGap"`
	ErasureRequestedAt    *DateTime `gorm:"column:erasureRequestedAt;index:OutboundEmailAudit_erasureRequestedAt_id_idx,priority:1;type:timestamp(3)" json:"erasureRequestedAt"`
	ErasedAt              *DateTime `gorm:"column:erasedAt;type:timestamp(3)" json:"erasedAt"`
	ContentDeleteRetryAt  *DateTime `gorm:"column:contentDeleteRetryAt;index:OutboundEmailAudit_contentDeleteRetryAt_id_idx,priority:1;type:timestamp(3)" json:"contentDeleteRetryAt"`
	ContentDeleteFailures int       `gorm:"column:contentDeleteFailures;size:32;default:0;not null" json:"contentDeleteFailures"`
}

func (OutboundEmailAudit) TableName() string {
	return "OutboundEmailAudit"
}

func (m *OutboundEmailAudit) BeforeCreate(tx *gorm.DB) error {
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

func (*OutboundEmailAudit) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Holds an account erasure request across an in-flight email send. The address is never stored.
// The user id is removed once deletion is confirmed; one-way recovery keys expire after 90 days.
type OutboundEmailAuditErasureIntent struct {
	ID           string    `gorm:"column:id;primaryKey;index:OutboundEmailAuditErasureIntent_checkedAt_id_idx,priority:2" json:"id"`
	UserID       *string   `gorm:"column:userId;index:OutboundEmailAuditErasureIntent_userId_idx" json:"userId"`
	UserHash     string    `gorm:"column:userHash;index:OutboundEmailAuditErasureIntent_userHash_activatedAt_idx,priority:1;not null" json:"userHash"`
	RecipientMac string    `gorm:"column:recipientMac;index:OutboundEmailAuditErasureIntent_recipientMac_activatedAt_idx,priority:1;not null" json:"recipientMac"`
	CreatedAt    DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ActivatedAt  *DateTime `gorm:"column:activatedAt;index:OutboundEmailAuditErasureIntent_recipientMac_activatedAt_idx,priority:2;index:OutboundEmailAuditErasureIntent_userHash_activatedAt_idx,priority:2;type:timestamp(3)" json:"activatedAt"`
	CheckedAt    *DateTime `gorm:"column:checkedAt;index:OutboundEmailAuditErasureIntent_checkedAt_id_idx,priority:1;type:timestamp(3)" json:"checkedAt"`
	ExpiresAt    DateTime  `gorm:"column:expiresAt;index:OutboundEmailAuditErasureIntent_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
}

func (OutboundEmailAuditErasureIntent) TableName() string {
	return "OutboundEmailAuditErasureIntent"
}

func (m *OutboundEmailAuditErasureIntent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Represents a webhook endpoint for receiving survey-related events.
// Webhooks can be configured to receive notifications about response creation, updates, and completion.
//
// @property id - Unique identifier for the webhook
// @property name - Optional display name for the webhook
// @property url - The endpoint URL where events will be sent
// @property source - Origin of the webhook (user, zapier, make, etc.)
// @property workspace - Associated workspace
// @property triggers - Types of events that trigger this webhook
// @property surveyIds - List of surveys this webhook is monitoring
type Webhook struct {
	ID          string             `gorm:"column:id;primaryKey" json:"id"`
	Name        *string            `gorm:"column:name" json:"name"`
	CreatedAt   DateTime           `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime           `gorm:"column:updated_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
	URL         string             `gorm:"column:url;not null" json:"url"`
	Source      string             `gorm:"column:source;default:'user';not null" json:"source"`
	WorkspaceID string             `gorm:"column:workspaceId;index:Webhook_workspaceId_idx;not null" json:"workspaceId"`
	Triggers    ScalarList[string] `gorm:"column:triggers;type:text[]" json:"triggers"`
	SurveyIds   ScalarList[string] `gorm:"column:surveyIds;type:text[]" json:"surveyIds"`
	Secret      *string            `gorm:"column:secret" json:"secret"`
	Workspace   *Workspace         `gorm:"belongsTo"`
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

// Represents an attribute value associated with a contact.
// Used to store custom properties and characteristics of contacts.
//
// @property id - Unique identifier for the attribute
// @property attributeKey - Reference to the attribute definition
// @property contact - The contact this attribute belongs to
// @property value - The string value of the attribute (used for string type + backwards compatibility)
// @property valueNumber - Native numeric storage for number type attributes
// @property valueDate - Native date storage for date type attributes
type ContactAttribute struct {
	ID             string               `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime             `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt      DateTime             `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	AttributeKeyID string               `gorm:"column:attributeKeyId;uniqueIndex:ContactAttribute_contactId_attributeKeyId_key,priority:2;index:ContactAttribute_attributeKeyId_value_idx,priority:1;index:ContactAttribute_attributeKeyId_valueNumber_idx,priority:1;index:ContactAttribute_attributeKeyId_valueDate_idx,priority:1;not null" json:"attributeKeyId"`
	ContactID      string               `gorm:"column:contactId;uniqueIndex:ContactAttribute_contactId_attributeKeyId_key,priority:1;not null" json:"contactId"`
	Value          string               `gorm:"column:value;index:ContactAttribute_attributeKeyId_value_idx,priority:2;not null" json:"value"`
	ValueNumber    *float64             `gorm:"column:valueNumber;index:ContactAttribute_attributeKeyId_valueNumber_idx,priority:2;type:double precision" json:"valueNumber"`
	ValueDate      *DateTime            `gorm:"column:valueDate;index:ContactAttribute_attributeKeyId_valueDate_idx,priority:2;type:timestamp(3)" json:"valueDate"`
	AttributeKey   *ContactAttributeKey `gorm:"foreignKey:AttributeKeyID;belongsTo"`
	Contact        *Contact             `gorm:"belongsTo"`
}

func (ContactAttribute) TableName() string {
	return "ContactAttribute"
}

func (m *ContactAttribute) BeforeCreate(tx *gorm.DB) error {
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

func (*ContactAttribute) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Defines the possible attributes that can be assigned to contacts.
// Acts as a schema for contact attributes within a workspace.
//
// @property id - Unique identifier for the attribute key
// @property isUnique - Whether the attribute must have unique values across contacts
// @property key - The attribute identifier used in the system
// @property name - Display name for the attribute
// @property type - Whether this is a default or custom attribute
// @property dataType - The data type of the attribute (string, number, date)
// @property workspace - The workspace this attribute belongs to
type ContactAttributeKey struct {
	ID               string                  `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt        DateTime                `gorm:"column:created_at;index:ContactAttributeKey_workspaceId_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt        DateTime                `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	IsUnique         bool                    `gorm:"column:isUnique;default:false;not null" json:"isUnique"`
	Key              string                  `gorm:"column:key;uniqueIndex:ContactAttributeKey_key_workspaceId_key,priority:1;not null" json:"key"`
	Name             *string                 `gorm:"column:name" json:"name"`
	Description      *string                 `gorm:"column:description" json:"description"`
	Type             string                  `gorm:"column:type;default:'custom';not null" json:"type"`
	DataType         string                  `gorm:"column:dataType;default:'string';not null" json:"dataType"`
	WorkspaceID      string                  `gorm:"column:workspaceId;uniqueIndex:ContactAttributeKey_key_workspaceId_key,priority:2;index:ContactAttributeKey_workspaceId_created_at_idx,priority:1;not null" json:"workspaceId"`
	Workspace        *Workspace              `gorm:"belongsTo"`
	Attributes       []ContactAttribute      `gorm:"foreignKey:AttributeKeyID;constraint:ContactAttribute_attributeKeyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AttributeFilters []SurveyAttributeFilter `gorm:"foreignKey:AttributeKeyID;constraint:SurveyAttributeFilter_attributeKeyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ContactAttributeKey) TableName() string {
	return "ContactAttributeKey"
}

func (m *ContactAttributeKey) BeforeCreate(tx *gorm.DB) error {
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

func (*ContactAttributeKey) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a person or user who can receive and respond to surveys.
// Contacts are workspace-specific and can have multiple attributes and responses.
//
// @property id - Unique identifier for the contact
// @property workspace - The workspace this contact belongs to
// @property responses - Survey responses from this contact
// @property attributes - Custom attributes associated with this contact
// @property displays - Record of surveys shown to this contact
type Contact struct {
	ID          string             `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime           `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime           `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	WorkspaceID string             `gorm:"column:workspaceId;index:Contact_workspaceId_idx;not null" json:"workspaceId"`
	Workspace   *Workspace         `gorm:"belongsTo"`
	Responses   []Response         `gorm:"foreignKey:ContactID;constraint:Response_contactId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attributes  []ContactAttribute `gorm:"foreignKey:ContactID;constraint:ContactAttribute_contactId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Displays    []Display          `gorm:"foreignKey:ContactID;constraint:Display_contactId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Contact) TableName() string {
	return "Contact"
}

func (m *Contact) BeforeCreate(tx *gorm.DB) error {
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

func (*Contact) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Stores a user's response to a survey, including their answers and metadata.
// Each response is linked to a specific survey and optionally to a contact.
//
// @property id - Unique identifier for the response
// @property finished - Whether the survey was completed
// @property survey - The associated survey
// @property contact - The optional contact who provided the response
// @property data - JSON object containing the actual response data
// @property variables - Custom variables used in the response
// @property ttc - Time to completion metrics
// @property meta - Additional metadata about the response
type Response struct {
	ID        string   `gorm:"column:id;primaryKey;index:Response_created_at_id_idx,priority:2;index:Response_surveyId_created_at_id_idx,priority:3;index:Response_contactId_created_at_id_idx,priority:3" json:"id"`
	CreatedAt DateTime `gorm:"column:created_at;index:Response_created_at_idx;index:Response_created_at_id_idx,priority:1;index:Response_surveyId_created_at_idx,priority:2;index:Response_surveyId_created_at_id_idx,priority:2;index:Response_contactId_created_at_idx,priority:2;index:Response_contactId_created_at_id_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
	Finished  bool     `gorm:"column:finished;index:Response_surveyId_finished_idx,priority:2;default:false;not null" json:"finished"`
	SurveyID  string   `gorm:"column:surveyId;uniqueIndex:Response_surveyId_singleUseId_key,priority:1;index:Response_surveyId_created_at_idx,priority:1;index:Response_surveyId_created_at_id_idx,priority:1;index:Response_surveyId_finished_idx,priority:1;not null" json:"surveyId"`
	ContactID *string  `gorm:"column:contactId;index:Response_contactId_created_at_idx,priority:1;index:Response_contactId_created_at_id_idx,priority:1" json:"contactId"`
	EndingID  *string  `gorm:"column:endingId" json:"endingId"`
	// [ResponseData]
	Data datatypes.JSON `gorm:"column:data;default:'{}';not null" json:"data"`
	// [ResponseVariables]
	Variables datatypes.JSON `gorm:"column:variables;default:'{}';not null" json:"variables"`
	// [ResponseTtc]
	Ttc datatypes.JSON `gorm:"column:ttc;default:'{}';not null" json:"ttc"`
	// [ResponseMeta]
	Meta datatypes.JSON `gorm:"column:meta;default:'{}';not null" json:"meta"`
	// Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
	// or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
	// from the client. Null means no ingest boundary has written it, which is not the same claim as
	// an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
	// and surfacing that in the UI (v2) would have nothing to read.
	// [ResponseIngestFlags]
	IngestFlags datatypes.JSON `gorm:"column:ingest_flags" json:"ingest_flags"`
	// What the Cal.com elements reported when the respondent booked, `{ [Question ID]: booking }`. `data`
	// keeps the string "booked" for those questions, so this column is purely additive: it is omitted from
	// every read without a `select` (`PRISMA_GLOBAL_OMIT`), and only the webhook body (as
	// `calBookingsByQuestionId`) and `/api/v3/responses` read it. Server-derived from the incoming `data`,
	// never accepted as a field of its own.
	// [ResponseCalBookings]
	CalBookings datatypes.JSON `gorm:"column:cal_bookings_by_question_id" json:"cal_bookings_by_question_id"`
	// A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
	// export by default and from Feedback Records, and nothing else: it is never deleted and every
	// count, list, webhook and integration still includes it (ENG-6780).
	Outcome *string `gorm:"column:outcome" json:"outcome"`
	// What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
	// from them, and per element its options, matrix rows and columns, by stable ids. Written by the
	// survey while they answer; null means unknown (older responses and bundles), read as before.
	// Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
	// webhook or integration.
	// [ResponseNotShown]
	NotShown datatypes.JSON `gorm:"column:not_shown" json:"not_shown"`
	// [ResponseContactAttributes]
	ContactAttributes datatypes.JSON      `gorm:"column:contactAttributes" json:"contactAttributes"`
	SingleUseID       *string             `gorm:"column:singleUseId;uniqueIndex:Response_surveyId_singleUseId_key,priority:2" json:"singleUseId"`
	Language          *string             `gorm:"column:language" json:"language"`
	DisplayID         *string             `gorm:"column:displayId;uniqueIndex:Response_displayId_key" json:"displayId"`
	Survey            *Survey             `gorm:"belongsTo"`
	Contact           *Contact            `gorm:"belongsTo"`
	Display           *Display            `gorm:"belongsTo"`
	Tags              []TagsOnResponses   `gorm:"foreignKey:ResponseID;constraint:TagsOnResponses_responseId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	QuotaLinks        []ResponseQuotaLink `gorm:"foreignKey:ResponseID;constraint:ResponseQuotaLink_responseId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkflowRuns      []WorkflowRun       `gorm:"foreignKey:ResponseID;constraint:WorkflowRun_responseId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Response) TableName() string {
	return "Response"
}

func (m *Response) BeforeCreate(tx *gorm.DB) error {
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

func (*Response) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a label that can be applied to survey responses.
// Used for categorizing and organizing responses within a workspace.
//
// @property id - Unique identifier for the tag
// @property name - Display name of the tag
// @property responses - Survey responses tagged with this label
// @property workspace - The workspace this tag belongs to
type Tag struct {
	ID          string            `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime          `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime          `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name        string            `gorm:"column:name;uniqueIndex:Tag_workspaceId_name_key,priority:2;not null" json:"name"`
	WorkspaceID string            `gorm:"column:workspaceId;uniqueIndex:Tag_workspaceId_name_key,priority:1;index:Tag_workspaceId_idx;not null" json:"workspaceId"`
	Workspace   *Workspace        `gorm:"belongsTo"`
	Responses   []TagsOnResponses `gorm:"foreignKey:TagID;constraint:TagsOnResponses_tagId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

// Junction table linking tags to responses.
// Enables many-to-many relationship between tags and responses.
//
// @property response - The tagged response
// @property tag - The tag applied to the response
type TagsOnResponses struct {
	ResponseID string    `gorm:"column:responseId;primaryKey" json:"responseId"`
	TagID      string    `gorm:"column:tagId;primaryKey" json:"tagId"`
	Response   *Response `gorm:"belongsTo"`
	Tag        *Tag      `gorm:"belongsTo"`
}

func (TagsOnResponses) TableName() string {
	return "TagsOnResponses"
}

// Records when a survey is shown to a user.
// Tracks survey display history and response status.
//
// @property id - Unique identifier for the display event
// @property survey - The survey that was displayed
// @property contact - The contact who saw the survey
// @property response - The associated response if one exists
// @property creativeId - The creative that opened the survey. No foreign key, so the record survives the creative's deletion.
type Display struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt  DateTime  `gorm:"column:created_at;index:Display_contactId_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt  DateTime  `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	SurveyID   string    `gorm:"column:surveyId;index:Display_surveyId_idx;not null" json:"surveyId"`
	ContactID  *string   `gorm:"column:contactId;index:Display_contactId_created_at_idx,priority:1" json:"contactId"`
	CreativeID *string   `gorm:"column:creativeId" json:"creativeId"`
	Survey     *Survey   `gorm:"belongsTo"`
	Contact    *Contact  `gorm:"belongsTo"`
	Response   *Response `gorm:"foreignKey:DisplayID;constraint:Response_displayId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Display) TableName() string {
	return "Display"
}

func (m *Display) BeforeCreate(tx *gorm.DB) error {
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

func (*Display) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Links surveys to specific trigger actions.
// Defines which user actions should cause a survey to be displayed.
// This is the connection table between Surveys and ActionClasses that determines
// when and under what conditions a survey should be triggered.
//
// @property id - Unique identifier for the trigger
// @property survey - The survey to be triggered
// @property actionClass - The action that triggers the survey
// @property createdAt - When the trigger was created
// @property updatedAt - When the trigger was last modified
type SurveyTrigger struct {
	ID            string       `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt     DateTime     `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt     DateTime     `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	SurveyID      string       `gorm:"column:surveyId;uniqueIndex:SurveyTrigger_surveyId_actionClassId_key,priority:1;not null" json:"surveyId"`
	ActionClassID string       `gorm:"column:actionClassId;uniqueIndex:SurveyTrigger_surveyId_actionClassId_key,priority:2;not null" json:"actionClassId"`
	Survey        *Survey      `gorm:"belongsTo"`
	ActionClass   *ActionClass `gorm:"belongsTo"`
}

func (SurveyTrigger) TableName() string {
	return "SurveyTrigger"
}

func (m *SurveyTrigger) BeforeCreate(tx *gorm.DB) error {
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

func (*SurveyTrigger) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Defines targeting rules for surveys based on contact attributes.
// Used to show surveys only to contacts matching specific criteria.
//
// @property id - Unique identifier for the filter
// @property attributeKey - The contact attribute to filter on
// @property survey - The survey being filtered
// @property condition - The comparison operator to use
// @property value - The value to compare against
type SurveyAttributeFilter struct {
	ID             string               `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime             `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt      DateTime             `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	AttributeKeyID string               `gorm:"column:attributeKeyId;uniqueIndex:SurveyAttributeFilter_surveyId_attributeKeyId_key,priority:2;index:SurveyAttributeFilter_attributeKeyId_idx;not null" json:"attributeKeyId"`
	SurveyID       string               `gorm:"column:surveyId;uniqueIndex:SurveyAttributeFilter_surveyId_attributeKeyId_key,priority:1;not null" json:"surveyId"`
	Condition      string               `gorm:"column:condition;not null" json:"condition"`
	Value          string               `gorm:"column:value;not null" json:"value"`
	AttributeKey   *ContactAttributeKey `gorm:"foreignKey:AttributeKeyID;belongsTo"`
	Survey         *Survey              `gorm:"belongsTo"`
}

func (SurveyAttributeFilter) TableName() string {
	return "SurveyAttributeFilter"
}

func (m *SurveyAttributeFilter) BeforeCreate(tx *gorm.DB) error {
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

func (*SurveyAttributeFilter) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a complete survey configuration including questions, styling, and display rules.
// Core model for the survey functionality in Formbricks.
//
// @property id - Unique identifier for the survey
// @property name - Display name of the survey
// @property type - Survey delivery method (link, web, website, app)
// @property status - Current state of the survey (draft, active, completed, etc)
// @property workspace - The workspace this survey belongs to
// @property questions - JSON array containing survey questions configuration
// @property displayOption - Rules for how often the survey can be shown
// @property triggers - Actions that can trigger this survey
// @property attributeFilters - Rules for targeting specific contacts
// @property customHeadScripts - Survey-specific custom HTML scripts (self-hosted only)
// @property customHeadScriptsMode - "add" (merge with workspace) or "replace" (override workspace)
type Survey struct {
	// [SurveyHiddenFields]
	HiddenFields datatypes.JSON `gorm:"column:hiddenFields;default:'{\"enabled\": false}';not null" json:"hiddenFields"`
	// [SurveyVariables]
	Variables            datatypes.JSON `gorm:"column:variables;default:'[]';not null" json:"variables"`
	OwnershipInitialized bool           `gorm:"column:ownershipInitialized;default:false;not null" json:"ownershipInitialized"`
	ID                   string         `gorm:"column:id;primaryKey;uniqueIndex:Survey_id_workspaceId_key,priority:1" json:"id"`
	CreatedAt            DateTime       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt            DateTime       `gorm:"column:updated_at;index:Survey_workspaceId_updated_at_idx,priority:2;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name                 string         `gorm:"column:name;not null" json:"name"`
	RedirectURL          *string        `gorm:"column:redirectUrl" json:"redirectUrl"`
	Type                 string         `gorm:"column:type;default:'app';not null" json:"type"`
	WorkspaceID          string         `gorm:"column:workspaceId;uniqueIndex:Survey_id_workspaceId_key,priority:2;index:Survey_workspaceId_updated_at_idx,priority:1;index:Survey_workspaceId_ownerId_idx,priority:1;not null" json:"workspaceId"`
	CreatedBy            *string        `gorm:"column:createdBy" json:"createdBy"`
	// ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
	// the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
	Visibility string  `gorm:"column:visibility;default:'workspace';not null" json:"visibility"`
	OwnerID    *string `gorm:"column:ownerId;index:Survey_workspaceId_ownerId_idx,priority:2" json:"ownerId"`
	// Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
	// with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
	// projection, decided from PostgreSQL facts until the projector acknowledges it.
	VisibilityVersion int `gorm:"column:visibilityVersion;size:32;default:0;not null" json:"visibilityVersion"`
	// The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was.
	VisibilityProjectedVersion int `gorm:"column:visibilityProjectedVersion;size:32;default:0;not null" json:"visibilityProjectedVersion"`
	// "visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
	// trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
	// in create/update data.
	VisibilityPending     bool      `gorm:"column:visibilityPending;default:false;not null" json:"visibilityPending"`
	VisibilityChangedAt   *DateTime `gorm:"column:visibilityChangedAt;type:timestamp(3)" json:"visibilityChangedAt"`
	VisibilityChangedByID *string   `gorm:"column:visibilityChangedById" json:"visibilityChangedById"`
	Status                string    `gorm:"column:status;index:Survey_status_publishOn_idx,priority:1;index:Survey_status_closeOn_idx,priority:1;default:'draft';not null" json:"status"`
	// [SurveyWelcomeCard]
	WelcomeCard datatypes.JSON `gorm:"column:welcomeCard;default:'{\"enabled\": false}';not null" json:"welcomeCard"`
	// [SurveyQuestions]
	Questions datatypes.JSON `gorm:"column:questions;default:'[]';not null" json:"questions"`
	// [SurveyBlocks]
	Blocks []datatypes.JSON `gorm:"column:blocks;serializer:json" json:"blocks"`
	// [SurveyEnding]
	Endings       []datatypes.JSON `gorm:"column:endings;serializer:json" json:"endings"`
	DisplayOption string           `gorm:"column:displayOption;default:'displayOnce';not null" json:"displayOption"`
	RecontactDays *int             `gorm:"column:recontactDays;size:32" json:"recontactDays"`
	DisplayLimit  *int             `gorm:"column:displayLimit;size:32" json:"displayLimit"`
	// [SurveyInlineTriggers]
	InlineTriggers     datatypes.JSON `gorm:"column:inlineTriggers" json:"inlineTriggers"`
	AutoClose          *int           `gorm:"column:autoClose;size:32" json:"autoClose"`
	KeepEndingCardOpen bool           `gorm:"column:keepEndingCardOpen;default:false;not null" json:"keepEndingCardOpen"`
	AutoComplete       *int           `gorm:"column:autoComplete;size:32" json:"autoComplete"`
	Delay              int            `gorm:"column:delay;size:32;default:0;not null" json:"delay"`
	PublishOn          *DateTime      `gorm:"column:publishOn;index:Survey_status_publishOn_idx,priority:2;type:timestamp(3)" json:"publishOn"`
	CloseOn            *DateTime      `gorm:"column:closeOn;index:Survey_status_closeOn_idx,priority:2;type:timestamp(3)" json:"closeOn"`
	ArchivedAt         *DateTime      `gorm:"column:archivedAt;index:Survey_archivedAt_idx;type:timestamp(3)" json:"archivedAt"`
	// [SurveyClosedMessage]
	SurveyClosedMessage datatypes.JSON `gorm:"column:surveyClosedMessage" json:"surveyClosedMessage"`
	SegmentID           *string        `gorm:"column:segmentId;index:Survey_segmentId_idx" json:"segmentId"`
	// [SurveyWorkspaceOverwrites]
	WorkspaceOverwrites datatypes.JSON `gorm:"column:workspaceOverwrites" json:"workspaceOverwrites"`
	// [SurveyStyling]
	Styling datatypes.JSON `gorm:"column:styling" json:"styling"`
	// [SurveySingleUse]
	SingleUse                   datatypes.JSON `gorm:"column:singleUse;default:'{\"enabled\": false, \"isEncrypted\": true}'" json:"singleUse"`
	IsVerifyEmailEnabled        bool           `gorm:"column:isVerifyEmailEnabled;default:false;not null" json:"isVerifyEmailEnabled"`
	IsBackButtonHidden          bool           `gorm:"column:isBackButtonHidden;default:false;not null" json:"isBackButtonHidden"`
	IsAutoProgressingEnabled    bool           `gorm:"column:isAutoProgressingEnabled;default:false;not null" json:"isAutoProgressingEnabled"`
	IsCaptureIPEnabled          bool           `gorm:"column:isCaptureIpEnabled;default:false;not null" json:"isCaptureIpEnabled"`
	IsAnonymizeResponsesEnabled bool           `gorm:"column:is_anonymize_responses_enabled;default:false;not null" json:"is_anonymize_responses_enabled"`
	Pin                         *string        `gorm:"column:pin" json:"pin"`
	DisplayPercentage           *float64       `gorm:"column:displayPercentage;type:decimal(65,30)" json:"displayPercentage"`
	ShowLanguageSwitch          *bool          `gorm:"column:showLanguageSwitch" json:"showLanguageSwitch"`
	// Open the survey in the respondent's browser language when it matches an enabled survey
	// language. Null on surveys that predate the setting, which reads as off.
	AutoSelectLanguage *bool `gorm:"column:autoSelectLanguage" json:"autoSelectLanguage"`
	// [SurveyRecaptcha]
	Recaptcha datatypes.JSON `gorm:"column:recaptcha;default:'{\"enabled\": false, \"threshold\":0.1}'" json:"recaptcha"`
	// [SurveyLinkMetadata]
	Metadata              datatypes.JSON `gorm:"column:metadata;default:'{}';not null" json:"metadata"`
	Slug                  *string        `gorm:"column:slug;uniqueIndex:Survey_slug_key" json:"slug"`
	CustomHeadScripts     *string        `gorm:"column:customHeadScripts" json:"customHeadScripts"`
	CustomHeadScriptsMode *string        `gorm:"column:customHeadScriptsMode;default:'add'" json:"customHeadScriptsMode"`
	// [CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
	// by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
	CustomCSS datatypes.JSON `gorm:"column:customCss" json:"customCss"`
	// [SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
	// (ENG-3873). null = no conditions, which is what every existing survey keeps.
	ShowConditions datatypes.JSON `gorm:"column:showConditions" json:"showConditions"`
	// [SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
	// survey on load. null = none, which is what every existing survey keeps.
	StartLogic             datatypes.JSON                    `gorm:"column:startLogic" json:"startLogic"`
	Workspace              *Workspace                        `gorm:"belongsTo"`
	Creator                *User                             `gorm:"foreignKey:CreatedBy;belongsTo"`
	Owner                  *User                             `gorm:"foreignKey:OwnerID;belongsTo"`
	Segment                *Segment                          `gorm:"belongsTo"`
	Responses              []Response                        `gorm:"foreignKey:SurveyID;constraint:Response_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Quotas                 []SurveyQuota                     `gorm:"foreignKey:SurveyID;constraint:SurveyQuota_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Triggers               []SurveyTrigger                   `gorm:"foreignKey:SurveyID;constraint:SurveyTrigger_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AttributeFilters       []SurveyAttributeFilter           `gorm:"foreignKey:SurveyID;constraint:SurveyAttributeFilter_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Displays               []Display                         `gorm:"foreignKey:SurveyID;constraint:Display_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Languages              []SurveyLanguage                  `gorm:"foreignKey:SurveyID;constraint:SurveyLanguage_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmbeddedData           []EmbeddedData                    `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;constraint:EmbeddedData_surveyId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmbeddedDataLinks      []SurveyEmbeddedData              `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;constraint:SurveyEmbeddedData_surveyId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FollowUps              []SurveyFollowUp                  `gorm:"foreignKey:SurveyID;constraint:SurveyFollowUp_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FeedbackSourceMappings []FeedbackSourceFormbricksMapping `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;constraint:FeedbackSourceFormbricksMapping_surveyId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RetentionExemptions    []RetentionExemption              `gorm:"foreignKey:SurveyID;constraint:RetentionExemption_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RetentionNotices       []RetentionNotice                 `gorm:"foreignKey:SurveyID;constraint:RetentionNotice_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Creatives whose Destination is this survey (ENG-3993).
	Creatives []Creative `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;constraint:Creative_surveyId_workspaceId_fkey,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	// Background attachment ZIP exports of this survey (ENG-6862).
	AttachmentExports []AttachmentExport      `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;constraint:AttachmentExport_surveyId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RetentionRestore  *RetentionSurveyRestore `gorm:"foreignKey:SurveyID;constraint:RetentionSurveyRestore_surveyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Survey) TableName() string {
	return "Survey"
}

func (m *Survey) BeforeCreate(tx *gorm.DB) error {
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

func (*Survey) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a quota configuration for a survey.
// Defines response limits and conditions for quota management.
//
// @property id - Unique identifier for the quota
// @property survey - The survey this quota belongs to
// @property name - Display name for the quota
// @property limit - Maximum number of responses allowed
// @property conditions - JSON object containing quota conditions
// @property action - Action to take when quota is reached
// @property endingCardId - Optional ending card to show when quota is reached
type SurveyQuota struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	SurveyID  string   `gorm:"column:surveyId;uniqueIndex:SurveyQuota_surveyId_name_key,priority:1;not null" json:"surveyId"`
	Name      string   `gorm:"column:name;uniqueIndex:SurveyQuota_surveyId_name_key,priority:2;not null" json:"name"`
	Limit     int      `gorm:"column:limit;size:32;not null" json:"limit"`
	// [SurveyQuotaLogic]
	Logic                   datatypes.JSON      `gorm:"column:logic;default:'{}';not null" json:"logic"`
	Action                  string              `gorm:"column:action;not null" json:"action"`
	EndingCardID            *string             `gorm:"column:endingCardId" json:"endingCardId"`
	CountPartialSubmissions bool                `gorm:"column:countPartialSubmissions;default:false;not null" json:"countPartialSubmissions"`
	Survey                  *Survey             `gorm:"belongsTo"`
	QuotaLinks              []ResponseQuotaLink `gorm:"foreignKey:QuotaID;constraint:ResponseQuotaLink_quotaId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (SurveyQuota) TableName() string {
	return "SurveyQuota"
}

func (m *SurveyQuota) BeforeCreate(tx *gorm.DB) error {
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

func (*SurveyQuota) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Junction table linking responses to quotas.
// Tracks which responses counted towards specific quotas and their status.
//
// @property response - The response that counted towards the quota
// @property quota - The quota this response counted towards
// @property status - Whether the response was screened in or out
type ResponseQuotaLink struct {
	ResponseID string       `gorm:"column:responseId;primaryKey" json:"responseId"`
	QuotaID    string       `gorm:"column:quotaId;primaryKey;index:ResponseQuotaLink_quotaId_status_idx,priority:1" json:"quotaId"`
	Status     string       `gorm:"column:status;index:ResponseQuotaLink_quotaId_status_idx,priority:2;not null" json:"status"`
	Response   *Response    `gorm:"belongsTo"`
	Quota      *SurveyQuota `gorm:"foreignKey:QuotaID;belongsTo"`
}

func (ResponseQuotaLink) TableName() string {
	return "ResponseQuotaLink"
}

// Defines follow-up actions for survey responses.
// Enables automated actions based on specific survey response conditions.
//
// @property id - Unique identifier for the follow-up
// @property survey - The associated survey
// @property name - Display name for the follow-up
// @property trigger - Conditions that activate the follow-up
// @property action - Actions to take when triggered
type SurveyFollowUp struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	SurveyID  string   `gorm:"column:surveyId;not null" json:"surveyId"`
	Name      string   `gorm:"column:name;not null" json:"name"`
	// [SurveyFollowUpTrigger]
	Trigger datatypes.JSON `gorm:"column:trigger;not null" json:"trigger"`
	// [SurveyFollowUpAction]
	Action datatypes.JSON `gorm:"column:action;not null" json:"action"`
	Survey *Survey        `gorm:"belongsTo"`
}

func (SurveyFollowUp) TableName() string {
	return "SurveyFollowUp"
}

func (m *SurveyFollowUp) BeforeCreate(tx *gorm.DB) error {
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

func (*SurveyFollowUp) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a user action that can trigger surveys.
// Used to define points in the user journey where surveys can be shown.
//
// @property id - Unique identifier for the action
// @property name - Display name of the action
// @property type - Whether this is a code or no-code action
// @property key - Unique identifier used in code implementation
// @property noCodeConfig - Configuration for no-code setup
// @property workspace - The workspace this action belongs to
type ActionClass struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:created_at;index:ActionClass_workspaceId_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name        string   `gorm:"column:name;uniqueIndex:ActionClass_name_workspaceId_key,priority:1;not null" json:"name"`
	Description *string  `gorm:"column:description" json:"description"`
	Type        string   `gorm:"column:type;not null" json:"type"`
	Key         *string  `gorm:"column:key;uniqueIndex:ActionClass_key_workspaceId_key,priority:1" json:"key"`
	// [ActionClassNoCodeConfig]
	NoCodeConfig     datatypes.JSON    `gorm:"column:noCodeConfig" json:"noCodeConfig"`
	WorkspaceID      string            `gorm:"column:workspaceId;uniqueIndex:ActionClass_key_workspaceId_key,priority:2;uniqueIndex:ActionClass_name_workspaceId_key,priority:2;index:ActionClass_workspaceId_created_at_idx,priority:1;not null" json:"workspaceId"`
	Workspace        *Workspace        `gorm:"belongsTo"`
	SurveyTriggers   []SurveyTrigger   `gorm:"foreignKey:ActionClassID;constraint:SurveyTrigger_actionClassId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreativeTriggers []CreativeTrigger `gorm:"foreignKey:ActionClassID;constraint:CreativeTrigger_actionClassId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ActionClass) TableName() string {
	return "ActionClass"
}

func (m *ActionClass) BeforeCreate(tx *gorm.DB) error {
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

func (*ActionClass) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents third-party service integrations.
// Enables data flow between Formbricks and external services.
//
// @property id - Unique identifier for the integration
// @property type - The service being integrated (Google Sheets, Notion, etc.)
// @property workspace - The workspace this integration belongs to
// @property config - Service-specific configuration details
type Integration struct {
	ID   string `gorm:"column:id;primaryKey" json:"id"`
	Type string `gorm:"column:type;uniqueIndex:Integration_type_workspaceId_key,priority:1;not null" json:"type"`
	// [IntegrationConfig]
	Config      datatypes.JSON `gorm:"column:config;not null" json:"config"`
	WorkspaceID string         `gorm:"column:workspaceId;uniqueIndex:Integration_type_workspaceId_key,priority:2;index:Integration_workspaceId_idx;not null" json:"workspaceId"`
	Workspace   *Workspace     `gorm:"belongsTo"`
}

func (Integration) TableName() string {
	return "Integration"
}

func (m *Integration) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Tracks database schema migrations.
// Used to manage and track the state of data structure changes.
//
// @property id - Unique identifier for the migration
// @property name - Name of the migration
// @property status - Current state of the migration (pending, applied, failed)
// @property startedAt - When the migration began
// @property finishedAt - When the migration completed
type DataMigration struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	StartedAt  DateTime  `gorm:"column:started_at;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"started_at"`
	FinishedAt *DateTime `gorm:"column:finished_at;type:timestamp(3)" json:"finished_at"`
	Name       string    `gorm:"column:name;uniqueIndex:DataMigration_name_key;not null" json:"name"`
	Status     string    `gorm:"column:status;not null" json:"status"`
}

func (DataMigration) TableName() string {
	return "DataMigration"
}

func (m *DataMigration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.StartedAt.IsZero() {
		m.StartedAt = DateTime{Time: now}
	}
	return nil
}

// Main grouping mechanism for resources in Formbricks.
// Each organization can have multiple workspaces to separate different applications or products.
//
// @property id - Unique identifier for the workspace
// @property name - Display name of the workspace
// @property organization - Reference to parent organization
// @property surveys - Collection of surveys in this workspace
// @property styling - Workspace-wide styling configuration
// @property config - Workspace-specific configuration
// @property recontactDays - Default recontact delay for surveys
// @property placement - Default widget placement for in-app surveys
type Workspace struct {
	ID                  string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt           DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt           DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name                string   `gorm:"column:name;uniqueIndex:Workspace_organizationId_name_key,priority:2;not null" json:"name"`
	LegacyEnvironmentID *string  `gorm:"column:legacyEnvironmentId;uniqueIndex:Workspace_legacyEnvironmentId_key" json:"legacyEnvironmentId"`
	OrganizationID      string   `gorm:"column:organizationId;uniqueIndex:Workspace_organizationId_name_key,priority:1;not null" json:"organizationId"`
	// [Styling]
	Styling datatypes.JSON `gorm:"column:styling;default:'{\"allowStyleOverwrite\":true}';not null" json:"styling"`
	// [WorkspaceConfig]
	Config              datatypes.JSON `gorm:"column:config;default:'{}';not null" json:"config"`
	RecontactDays       *int           `gorm:"column:recontactDays;size:32;default:7;not null" json:"recontactDays"`
	LinkSurveyBranding  *bool          `gorm:"column:linkSurveyBranding;default:true;not null" json:"linkSurveyBranding"`
	InAppSurveyBranding *bool          `gorm:"column:inAppSurveyBranding;default:true;not null" json:"inAppSurveyBranding"`
	Placement           string         `gorm:"column:placement;default:'bottomRight';not null" json:"placement"`
	ClickOutsideClose   *bool          `gorm:"column:clickOutsideClose;default:true;not null" json:"clickOutsideClose"`
	Overlay             string         `gorm:"column:overlay;default:'none';not null" json:"overlay"`
	// Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`.
	OverlayColor   *string `gorm:"column:overlayColor" json:"overlayColor"`
	OverlayOpacity *int    `gorm:"column:overlayOpacity;size:32" json:"overlayOpacity"`
	// [Logo]
	Logo              datatypes.JSON `gorm:"column:logo" json:"logo"`
	AppSetupCompleted bool           `gorm:"column:appSetupCompleted;default:false;not null" json:"appSetupCompleted"`
	CustomHeadScripts *string        `gorm:"column:customHeadScripts" json:"customHeadScripts"`
	// [CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service.
	CustomCSS datatypes.JSON `gorm:"column:customCss" json:"customCss"`
	// [CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
	// submitting its source again through the current processor and authorization checks.
	CustomCSSPrevious           datatypes.JSON               `gorm:"column:customCssPrevious" json:"customCssPrevious"`
	Organization                *Organization                `gorm:"belongsTo"`
	Languages                   []Language                   `gorm:"foreignKey:WorkspaceID;constraint:Language_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmbeddedData                []EmbeddedData               `gorm:"foreignKey:WorkspaceID;constraint:EmbeddedData_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkspaceTeams              []WorkspaceTeam              `gorm:"foreignKey:WorkspaceID;constraint:WorkspaceTeam_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FeedbackDirectoryWorkspaces []FeedbackDirectoryWorkspace `gorm:"foreignKey:WorkspaceID;constraint:FeedbackDirectoryWorkspace_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Charts                      []Chart                      `gorm:"foreignKey:WorkspaceID;constraint:Chart_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Dashboards                  []Dashboard                  `gorm:"foreignKey:WorkspaceID;constraint:Dashboard_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmbeddedFieldLabels         []EmbeddedFieldLabel         `gorm:"foreignKey:WorkspaceID;constraint:EmbeddedFieldLabel_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Creatives                   []Creative                   `gorm:"foreignKey:WorkspaceID;constraint:Creative_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AttachmentExports           []AttachmentExport           `gorm:"foreignKey:WorkspaceID;constraint:AttachmentExport_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Surveys                     []Survey                     `gorm:"foreignKey:WorkspaceID;constraint:Survey_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Contacts                    []Contact                    `gorm:"foreignKey:WorkspaceID;constraint:Contact_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ActionClasses               []ActionClass                `gorm:"foreignKey:WorkspaceID;constraint:ActionClass_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ContactAttributeKeys        []ContactAttributeKey        `gorm:"foreignKey:WorkspaceID;constraint:ContactAttributeKey_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks                    []Webhook                    `gorm:"foreignKey:WorkspaceID;constraint:Webhook_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags                        []Tag                        `gorm:"foreignKey:WorkspaceID;constraint:Tag_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Segments                    []Segment                    `gorm:"foreignKey:WorkspaceID;constraint:Segment_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Integrations                []Integration                `gorm:"foreignKey:WorkspaceID;constraint:Integration_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeyWorkspaces            []APIKeyWorkspace            `gorm:"foreignKey:WorkspaceID;constraint:ApiKeyWorkspace_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FeedbackSources             []FeedbackSource             `gorm:"foreignKey:WorkspaceID;constraint:FeedbackSource_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Workflows                   []Workflow                   `gorm:"foreignKey:WorkspaceID;constraint:Workflow_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkflowRuns                []WorkflowRun                `gorm:"foreignKey:WorkspaceID;constraint:WorkflowRun_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	McpGrants                   []McpGrantWorkspace          `gorm:"foreignKey:WorkspaceID;constraint:McpGrantWorkspace_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

// Represents the top-level organizational hierarchy in Formbricks.
// Self-hosting instances typically have a single organization, while Formbricks Cloud
// supports multiple organizations with multi-tenancy.
//
// @property id - Unique identifier for the organization
// @property name - Display name of the organization
// @property memberships - User memberships within the organization
// @property workspaces - Collection of workspaces owned by the organization
// @property billing - JSON field containing billing information
// @property whitelabel - Whitelabel configuration for the organization
// @property isAISmartToolsEnabled - Controls access to AI smart tools (e.g. translations) that never touch collected data
// @property displayTimeZone - Optional IANA time zone used for human-facing response timestamps (exports, integrations) and for the calendar days behind the analytics date filters (survey summary, charts, dashboards); null means UTC
type Organization struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name      string   `gorm:"column:name;not null" json:"name"`
	// [OrganizationWhitelabel]
	Whitelabel            datatypes.JSON         `gorm:"column:whitelabel;default:'{}';not null" json:"whitelabel"`
	IsAISmartToolsEnabled bool                   `gorm:"column:isAISmartToolsEnabled;default:false;not null" json:"isAISmartToolsEnabled"`
	DisplayTimeZone       *string                `gorm:"column:displayTimeZone" json:"displayTimeZone"`
	Memberships           []Membership           `gorm:"foreignKey:OrganizationID;constraint:Membership_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Workspaces            []Workspace            `gorm:"foreignKey:OrganizationID;constraint:Workspace_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invites               []Invite               `gorm:"foreignKey:OrganizationID;constraint:Invite_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Teams                 []Team                 `gorm:"foreignKey:OrganizationID;constraint:Team_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	APIKeys               []APIKey               `gorm:"foreignKey:OrganizationID;constraint:ApiKey_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FeedbackDirectories   []FeedbackDirectory    `gorm:"foreignKey:OrganizationID;constraint:FeedbackDirectory_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RetentionPolicies     []RetentionPolicy      `gorm:"foreignKey:OrganizationID;constraint:RetentionPolicy_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RetentionExemptions   []RetentionExemption   `gorm:"foreignKey:OrganizationID;constraint:RetentionExemption_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RetentionRuns         []RetentionRun         `gorm:"foreignKey:OrganizationID;constraint:RetentionRun_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RetentionNotices      []RetentionNotice      `gorm:"foreignKey:OrganizationID;constraint:RetentionNotice_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	McpGrants             []McpGrantOrganization `gorm:"foreignKey:OrganizationID;constraint:McpGrantOrganization_organizationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Billing               *OrganizationBilling   `gorm:"foreignKey:OrganizationID;constraint:OrganizationBilling_organization_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

// Stores billing and Stripe synchronization data for an organization.
//
// @property organizationId - Reference to the organization
// @property stripeCustomerId - Stripe customer ID for cloud billing
// @property limits - Plan limits used by feature and quota checks
// @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
// @property stripe - Stripe synchronization snapshot data
type OrganizationBilling struct {
	OrganizationID   string  `gorm:"column:organization_id;primaryKey" json:"organization_id"`
	StripeCustomerID *string `gorm:"column:stripe_customer_id;uniqueIndex:OrganizationBilling_stripe_customer_id_key" json:"stripe_customer_id"`
	// [OrganizationBillingPlanLimits]
	Limits           datatypes.JSON `gorm:"column:limits;not null" json:"limits"`
	UsageCycleAnchor *DateTime      `gorm:"column:usage_cycle_anchor;type:timestamp(3)" json:"usage_cycle_anchor"`
	// [OrganizationStripeBilling]
	Stripe       datatypes.JSON `gorm:"column:stripe" json:"stripe"`
	CreatedAt    DateTime       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt    DateTime       `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Organization *Organization  `gorm:"belongsTo"`
}

func (OrganizationBilling) TableName() string {
	return "OrganizationBilling"
}

func (m *OrganizationBilling) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

// Links users to organizations with specific roles.
// Manages organization membership and permissions.
// Core model for managing user access within organizations.
//
// @property organization - The organization the user belongs to
// @property user - The member user
// @property accepted - Whether the user has accepted the membership
// @property role - User's role within the organization (owner, manager, member, billing)
type Membership struct {
	UserID         string        `gorm:"column:userId;primaryKey" json:"userId"`
	OrganizationID string        `gorm:"column:organizationId;primaryKey;index:Membership_organizationId_idx" json:"organizationId"`
	Accepted       bool          `gorm:"column:accepted;default:false;not null" json:"accepted"`
	Role           string        `gorm:"column:role;default:'member';not null" json:"role"`
	Organization   *Organization `gorm:"belongsTo"`
	User           *User         `gorm:"belongsTo"`
}

func (Membership) TableName() string {
	return "Membership"
}

// Represents pending invitations to join an organization.
// Used to manage the process of adding new users to an organization.
// Once accepted, invites are converted into memberships.
//
// @property id - Unique identifier for the invite
// @property email - Email address of the invited user
// @property name - Optional display name for the invited user
// @property organization - The organization sending the invite
// @property creator - The user who created the invite
// @property acceptor - The user who accepted the invite (if accepted)
// @property expiresAt - When the invite becomes invalid
// @property role - Intended role for the invited user
// @property teamIds - Teams the user will be added to upon acceptance
type Invite struct {
	ID             string             `gorm:"column:id;primaryKey" json:"id"`
	Email          string             `gorm:"column:email;index:Invite_email_organizationId_idx,priority:1;not null" json:"email"`
	Name           *string            `gorm:"column:name" json:"name"`
	OrganizationID string             `gorm:"column:organizationId;index:Invite_email_organizationId_idx,priority:2;index:Invite_organizationId_idx;not null" json:"organizationId"`
	CreatorID      string             `gorm:"column:creatorId;not null" json:"creatorId"`
	AcceptorID     *string            `gorm:"column:acceptorId" json:"acceptorId"`
	CreatedAt      DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ExpiresAt      DateTime           `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	Role           string             `gorm:"column:role;default:'member';not null" json:"role"`
	TeamIds        ScalarList[string] `gorm:"column:teamIds;type:text[]" json:"teamIds"`
	Organization   *Organization      `gorm:"belongsTo"`
	Creator        *User              `gorm:"foreignKey:CreatorID;belongsTo"`
	Acceptor       *User              `gorm:"foreignKey:AcceptorID;belongsTo"`
}

func (Invite) TableName() string {
	return "Invite"
}

func (m *Invite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Represents enhanced API authentication keys with organization-level ownership.
// Used for authenticating API requests to Formbricks with more granular permissions.
//
// @property id - Unique identifier for the API key
// @property label - Optional descriptive name for the key
// @property hashedKey - Securely stored API key
// @property organization - The organization this key belongs to
// @property createdBy - User ID who created this key
// @property lastUsedAt - Timestamp of last usage
// @property apiKeyWorkspaces - Workspaces this key has access to
type APIKey struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	CreatedBy      *string   `gorm:"column:createdBy" json:"createdBy"`
	LastUsedAt     *DateTime `gorm:"column:lastUsedAt;type:timestamp(3)" json:"lastUsedAt"`
	Label          string    `gorm:"column:label;not null" json:"label"`
	HashedKey      string    `gorm:"column:hashedKey;not null" json:"hashedKey"`
	LookupHash     *string   `gorm:"column:lookupHash;uniqueIndex:ApiKey_lookupHash_key" json:"lookupHash"`
	OrganizationID string    `gorm:"column:organizationId;index:ApiKey_organizationId_idx;not null" json:"organizationId"`
	// [OrganizationAccess]
	OrganizationAccess datatypes.JSON    `gorm:"column:organizationAccess;default:'{}';not null" json:"organizationAccess"`
	Organization       *Organization     `gorm:"belongsTo"`
	APIKeyWorkspaces   []APIKeyWorkspace `gorm:"foreignKey:APIKeyID;constraint:ApiKeyWorkspace_apiKeyId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

// Links API keys to workspaces with specific permissions.
// Enables granular access control for API keys across workspaces.
//
// @property id - Unique identifier for the workspace access entry
// @property apiKey - The associated API key
// @property workspace - The workspace being accessed
// @property permission - Level of access granted
type APIKeyWorkspace struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	APIKeyID    string     `gorm:"column:apiKeyId;uniqueIndex:ApiKeyWorkspace_apiKeyId_workspaceId_key,priority:1;not null" json:"apiKeyId"`
	WorkspaceID string     `gorm:"column:workspaceId;uniqueIndex:ApiKeyWorkspace_apiKeyId_workspaceId_key,priority:2;index:ApiKeyWorkspace_workspaceId_idx;not null" json:"workspaceId"`
	Permission  string     `gorm:"column:permission;not null" json:"permission"`
	APIKey      *APIKey    `gorm:"belongsTo"`
	Workspace   *Workspace `gorm:"belongsTo"`
}

func (APIKeyWorkspace) TableName() string {
	return "ApiKeyWorkspace"
}

func (m *APIKeyWorkspace) BeforeCreate(tx *gorm.DB) error {
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

func (*APIKeyWorkspace) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Stores third-party authentication account information.
// Enables OAuth and other external authentication methods.
//
// @property id - Unique identifier for the account
// @property user - The Formbricks user who owns this account
// @property provider - The authentication provider (GitHub, Google, etc.)
// @property providerAccountId - User ID from the provider
// @property access_token - OAuth access token
// @property refresh_token - OAuth refresh token
type Account struct {
	ID                    string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt             DateTime  `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt             DateTime  `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	UserID                string    `gorm:"column:userId;index:Account_userId_idx;not null" json:"userId"`
	Type                  *string   `gorm:"column:type" json:"type"`
	Provider              string    `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key,priority:1;not null" json:"provider"`
	ProviderAccountID     string    `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key,priority:2;uniqueIndex:Account_issuer_providerAccountId_key,priority:2;not null" json:"providerAccountId"`
	AccessToken           *string   `gorm:"column:access_token;type:text" json:"access_token"`
	RefreshToken          *string   `gorm:"column:refresh_token;type:text" json:"refresh_token"`
	ExpiresAt             *int      `gorm:"column:expires_at;size:32" json:"expires_at"`
	ExtExpiresIn          *int      `gorm:"column:ext_expires_in;size:32" json:"ext_expires_in"`
	TokenType             *string   `gorm:"column:token_type" json:"token_type"`
	Scope                 *string   `gorm:"column:scope" json:"scope"`
	IDToken               *string   `gorm:"column:id_token;type:text" json:"id_token"`
	SessionState          *string   `gorm:"column:session_state" json:"session_state"`
	Password              *string   `gorm:"column:password" json:"password"`
	AccessTokenExpiresAt  *DateTime `gorm:"column:accessTokenExpiresAt;type:timestamp(3)" json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt *DateTime `gorm:"column:refreshTokenExpiresAt;type:timestamp(3)" json:"refreshTokenExpiresAt"`
	// Better Auth 1.7.0 (ENG-2343): the account's identity namespace. 1.7.0 keyed accounts on
	// (issuer, accountId) and filtered every lookup on it — findAccountByKey, findAccountOwnerByKey,
	// findCredentialAccount and updatePassword — so on 1.7.0 a null here locks the user out.
	// 1.7.3 dropped the field and 1.7.7 keys accounts on (provider, providerAccountId) again
	// (ENG-3880). We keep writing it (see `account.create.before` in better-auth-hooks.ts) so a
	// rollback to 1.7.0 still finds every row; drop it once that rollback is off the table.
	//
	// NULLABLE on purpose. A NOT NULL column would fail every INSERT from a 1.6 pod still serving
	// during a rolling deploy. The migration backfills every existing row; see the migration SQL for
	// the value derivation.
	Issuer *string `gorm:"column:issuer;uniqueIndex:Account_issuer_providerAccountId_key,priority:1" json:"issuer"`
	User   *User   `gorm:"belongsTo"`
}

func (Account) TableName() string {
	return "Account"
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

// Stores active authentication sessions for revocable server-side login state.
//
// @property sessionToken - Opaque token stored in the browser cookie
// @property user - The Formbricks user who owns this session
// @property expires - Hard expiry for the session
type Session struct {
	ID                 string              `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt          DateTime            `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt          DateTime            `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	SessionToken       string              `gorm:"column:sessionToken;uniqueIndex:Session_sessionToken_key;not null" json:"sessionToken"`
	UserID             string              `gorm:"column:userId;index:Session_userId_idx;not null" json:"userId"`
	Expires            DateTime            `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	IPAddress          *string             `gorm:"column:ipAddress" json:"ipAddress"`
	UserAgent          *string             `gorm:"column:userAgent" json:"userAgent"`
	User               *User               `gorm:"belongsTo"`
	OauthAccessTokens  []OauthAccessToken  `gorm:"foreignKey:SessionID;constraint:oauthAccessToken_sessionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	OauthRefreshTokens []OauthRefreshToken `gorm:"foreignKey:SessionID;constraint:oauthRefreshToken_sessionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Session) TableName() string {
	return "Session"
}

func (m *Session) BeforeCreate(tx *gorm.DB) error {
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

func (*Session) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Stores one-time verification tokens used by Auth.js adapter flows.
type VerificationToken struct {
	Identifier string   `gorm:"column:identifier;primaryKey;uniqueIndex:VerificationToken_identifier_token_key,priority:1" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:VerificationToken_identifier_token_key,priority:2" json:"token"`
	Expires    DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

// Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
// User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
// at rest by Better Auth using BETTER_AUTH_SECRET.
type TwoFactor struct {
	ID          string `gorm:"column:id;primaryKey" json:"id"`
	Secret      string `gorm:"column:secret;not null" json:"secret"`
	BackupCodes string `gorm:"column:backupCodes;not null" json:"backupCodes"`
	UserID      string `gorm:"column:userId;uniqueIndex:TwoFactor_userId_key;not null" json:"userId"`
	Verified    *bool  `gorm:"column:verified;default:true;not null" json:"verified"`
	// Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts.
	FailedVerificationCount int `gorm:"column:failedVerificationCount;size:32;default:0;not null" json:"failedVerificationCount"`
	// Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit.
	LockedUntil *DateTime `gorm:"column:lockedUntil;type:timestamp(3)" json:"lockedUntil"`
	User        *User     `gorm:"belongsTo"`
}

func (TwoFactor) TableName() string {
	return "TwoFactor"
}

func (m *TwoFactor) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Better Auth JWT signing keys for OAuth/OIDC access and ID tokens.
type Jwks struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	PublicKey  string    `gorm:"column:publicKey;not null" json:"publicKey"`
	PrivateKey string    `gorm:"column:privateKey;not null" json:"privateKey"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	ExpiresAt  *DateTime `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	// Better Auth 1.7 (ENG-2343): the JWS algorithm this key signs with, so a keyring can hold more
	// than one alg and `getLatestKeyByAlg` can pick the right key per resource.
	//
	// NOT optional upstream in practice, despite being declared `required: false`: `createJwk` always
	// puts `alg` in the row it writes (defaulting to EdDSA), so without this column every key mint is
	// rejected by Prisma as an unknown argument — which breaks JWT signing, and with it the whole MCP
	// OAuth flow, on any deployment that has yet to mint a key. Nullable because rows written before
	// 1.7 have no value; upstream reads treat `alg: null` as the configured default alg, so existing
	// keys keep working untouched.
	Alg *string `gorm:"column:alg" json:"alg"`
	// Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
	// the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
	// same always-written path as `alg` above.
	Crv *string `gorm:"column:crv" json:"crv"`
}

func (Jwks) TableName() string {
	return "jwks"
}

func (m *Jwks) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Better Auth OAuth 2.1 client registrations for MCP and future OAuth clients.
type OauthClient struct {
	ID                      string             `gorm:"column:id;primaryKey" json:"id"`
	ClientID                string             `gorm:"column:clientId;uniqueIndex:oauthClient_clientId_key;not null" json:"clientId"`
	ClientSecret            *string            `gorm:"column:clientSecret" json:"clientSecret"`
	Disabled                *bool              `gorm:"column:disabled;default:false" json:"disabled"`
	SkipConsent             *bool              `gorm:"column:skipConsent" json:"skipConsent"`
	EnableEndSession        *bool              `gorm:"column:enableEndSession" json:"enableEndSession"`
	SubjectType             *string            `gorm:"column:subjectType" json:"subjectType"`
	Scopes                  ScalarList[string] `gorm:"column:scopes;type:text[]" json:"scopes"`
	UserID                  *string            `gorm:"column:userId;index:oauthClient_userId_idx" json:"userId"`
	CreatedAt               *DateTime          `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false" json:"createdAt"`
	UpdatedAt               *DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	Name                    *string            `gorm:"column:name" json:"name"`
	URI                     *string            `gorm:"column:uri" json:"uri"`
	Icon                    *string            `gorm:"column:icon" json:"icon"`
	Contacts                ScalarList[string] `gorm:"column:contacts;type:text[]" json:"contacts"`
	Tos                     *string            `gorm:"column:tos" json:"tos"`
	Policy                  *string            `gorm:"column:policy" json:"policy"`
	SoftwareID              *string            `gorm:"column:softwareId" json:"softwareId"`
	SoftwareVersion         *string            `gorm:"column:softwareVersion" json:"softwareVersion"`
	SoftwareStatement       *string            `gorm:"column:softwareStatement" json:"softwareStatement"`
	RedirectUris            ScalarList[string] `gorm:"column:redirectUris;type:text[]" json:"redirectUris"`
	PostLogoutRedirectUris  ScalarList[string] `gorm:"column:postLogoutRedirectUris;type:text[]" json:"postLogoutRedirectUris"`
	TokenEndpointAuthMethod *string            `gorm:"column:tokenEndpointAuthMethod" json:"tokenEndpointAuthMethod"`
	GrantTypes              ScalarList[string] `gorm:"column:grantTypes;type:text[]" json:"grantTypes"`
	ResponseTypes           ScalarList[string] `gorm:"column:responseTypes;type:text[]" json:"responseTypes"`
	// Unwritten by Better Auth 1.7 (it dropped both from its schema), kept deliberately: they are
	// nullable, so keeping them makes the 1.7 migration purely additive and a rollback to 1.6 needs no
	// down-migration. Drop them in a separate cleanup after a bake period.
	Public      *bool          `gorm:"column:public" json:"public"`
	Type        *string        `gorm:"column:type" json:"type"`
	RequirePKCE *bool          `gorm:"column:requirePKCE" json:"requirePKCE"`
	ReferenceID *string        `gorm:"column:referenceId" json:"referenceId"`
	Metadata    datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	// Better Auth 1.7 (ENG-2343).
	ClientDiscoveryID                *string               `gorm:"column:clientDiscoveryId" json:"clientDiscoveryId"`
	ClientCredentialsScopes          ScalarList[string]    `gorm:"column:clientCredentialsScopes;type:text[]" json:"clientCredentialsScopes"`
	BackchannelLogoutURI             *string               `gorm:"column:backchannelLogoutUri" json:"backchannelLogoutUri"`
	BackchannelLogoutSessionRequired *bool                 `gorm:"column:backchannelLogoutSessionRequired" json:"backchannelLogoutSessionRequired"`
	ApplicationType                  *string               `gorm:"column:applicationType" json:"applicationType"`
	Jwks                             *string               `gorm:"column:jwks" json:"jwks"`
	JwksURI                          *string               `gorm:"column:jwksUri" json:"jwksUri"`
	DpopBoundAccessTokens            *bool                 `gorm:"column:dpopBoundAccessTokens;default:false" json:"dpopBoundAccessTokens"`
	User                             *User                 `gorm:"belongsTo"`
	AccessTokens                     []OauthAccessToken    `gorm:"foreignKey:ClientID;references:ClientID;constraint:oauthAccessToken_clientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RefreshTokens                    []OauthRefreshToken   `gorm:"foreignKey:ClientID;references:ClientID;constraint:oauthRefreshToken_clientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Consents                         []OauthConsent        `gorm:"foreignKey:ClientID;references:ClientID;constraint:oauthConsent_clientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Resources                        []OauthClientResource `gorm:"foreignKey:ClientID;references:ClientID;constraint:oauthClientResource_clientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	McpGrants                        []McpGrant            `gorm:"foreignKey:ClientID;references:ClientID;constraint:McpGrant_clientId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OauthClient) TableName() string {
	return "oauthClient"
}

func (m *OauthClient) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Better Auth opaque access token records. The token column stores Better Auth's hashed lookup value.
// JWT access tokens are verified through JWKS and bounded by the configured short lifetime.
type OauthAccessToken struct {
	ID          string             `gorm:"column:id;primaryKey" json:"id"`
	Token       string             `gorm:"column:token;uniqueIndex:oauthAccessToken_token_key;not null" json:"token"`
	ClientID    string             `gorm:"column:clientId;index:oauthAccessToken_clientId_idx;not null" json:"clientId"`
	SessionID   *string            `gorm:"column:sessionId;index:oauthAccessToken_sessionId_idx" json:"sessionId"`
	UserID      *string            `gorm:"column:userId;index:oauthAccessToken_userId_idx" json:"userId"`
	ReferenceID *string            `gorm:"column:referenceId" json:"referenceId"`
	RefreshID   *string            `gorm:"column:refreshId;index:oauthAccessToken_refreshId_idx" json:"refreshId"`
	ExpiresAt   DateTime           `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt   DateTime           `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	Scopes      ScalarList[string] `gorm:"column:scopes;type:text[]" json:"scopes"`
	// Better Auth 1.7 (ENG-2343). `resources` is the RFC 8707 resource set the grant approved — the
	// binding this whole upgrade exists to get (GHSA-p2fr-6hmx-4528).
	AuthorizationCodeID     *string            `gorm:"column:authorizationCodeId;index:oauthAccessToken_authorizationCodeId_idx" json:"authorizationCodeId"`
	Resources               ScalarList[string] `gorm:"column:resources;type:text[]" json:"resources"`
	RequestedUserInfoClaims ScalarList[string] `gorm:"column:requestedUserInfoClaims;type:text[]" json:"requestedUserInfoClaims"`
	Revoked                 *DateTime          `gorm:"column:revoked;type:timestamp(3)" json:"revoked"`
	Confirmation            datatypes.JSON     `gorm:"column:confirmation" json:"confirmation"`
	Client                  *OauthClient       `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
	Session                 *Session           `gorm:"belongsTo"`
	User                    *User              `gorm:"belongsTo"`
	Refresh                 *OauthRefreshToken `gorm:"foreignKey:RefreshID;belongsTo"`
}

func (OauthAccessToken) TableName() string {
	return "oauthAccessToken"
}

func (m *OauthAccessToken) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Better Auth OAuth refresh tokens used by delegated MCP clients. The token column stores a hashed lookup value.
type OauthRefreshToken struct {
	ID          string             `gorm:"column:id;primaryKey" json:"id"`
	Token       string             `gorm:"column:token;uniqueIndex:oauthRefreshToken_token_key;not null" json:"token"`
	ClientID    string             `gorm:"column:clientId;index:oauthRefreshToken_clientId_idx;not null" json:"clientId"`
	SessionID   *string            `gorm:"column:sessionId;index:oauthRefreshToken_sessionId_idx" json:"sessionId"`
	UserID      string             `gorm:"column:userId;index:oauthRefreshToken_userId_idx;not null" json:"userId"`
	ReferenceID *string            `gorm:"column:referenceId" json:"referenceId"`
	ExpiresAt   DateTime           `gorm:"column:expiresAt;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt   DateTime           `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	Revoked     *DateTime          `gorm:"column:revoked;type:timestamp(3)" json:"revoked"`
	AuthTime    *DateTime          `gorm:"column:authTime;type:timestamp(3)" json:"authTime"`
	Scopes      ScalarList[string] `gorm:"column:scopes;type:text[]" json:"scopes"`
	// Better Auth 1.7 (ENG-2343). `resources` carries the full original grant so a refresh can narrow
	// the resource set but never widen it.
	AuthorizationCodeID     *string            `gorm:"column:authorizationCodeId;index:oauthRefreshToken_authorizationCodeId_idx" json:"authorizationCodeId"`
	Resources               ScalarList[string] `gorm:"column:resources;type:text[]" json:"resources"`
	RequestedUserInfoClaims ScalarList[string] `gorm:"column:requestedUserInfoClaims;type:text[]" json:"requestedUserInfoClaims"`
	RotatedAt               *DateTime          `gorm:"column:rotatedAt;type:timestamp(3)" json:"rotatedAt"`
	RotationReplayResponse  *string            `gorm:"column:rotationReplayResponse" json:"rotationReplayResponse"`
	RotationReplayExpiresAt *DateTime          `gorm:"column:rotationReplayExpiresAt;type:timestamp(3)" json:"rotationReplayExpiresAt"`
	Confirmation            datatypes.JSON     `gorm:"column:confirmation" json:"confirmation"`
	Client                  *OauthClient       `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
	Session                 *Session           `gorm:"belongsTo"`
	User                    *User              `gorm:"belongsTo"`
	AccessTokens            []OauthAccessToken `gorm:"foreignKey:RefreshID;constraint:oauthAccessToken_refreshId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OauthRefreshToken) TableName() string {
	return "oauthRefreshToken"
}

func (m *OauthRefreshToken) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Better Auth OAuth user consent grants for delegated clients.
type OauthConsent struct {
	ID          string             `gorm:"column:id;primaryKey" json:"id"`
	ClientID    string             `gorm:"column:clientId;index:oauthConsent_clientId_idx;not null" json:"clientId"`
	UserID      *string            `gorm:"column:userId;index:oauthConsent_userId_idx" json:"userId"`
	ReferenceID *string            `gorm:"column:referenceId" json:"referenceId"`
	Scopes      ScalarList[string] `gorm:"column:scopes;type:text[]" json:"scopes"`
	CreatedAt   DateTime           `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	// Better Auth 1.7 (ENG-2343). /authorize compares the requested resources against these; without
	// them every pre-existing connection is forced back through the consent screen.
	Resources               ScalarList[string] `gorm:"column:resources;type:text[]" json:"resources"`
	RequestedUserInfoClaims ScalarList[string] `gorm:"column:requestedUserInfoClaims;type:text[]" json:"requestedUserInfoClaims"`
	Client                  *OauthClient       `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
	User                    *User              `gorm:"belongsTo"`
}

func (OauthConsent) TableName() string {
	return "oauthConsent"
}

func (m *OauthConsent) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
// its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
// a token is now bound to the resource it was approved for instead of to a static allow-list.
// `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
// What one user allows one MCP client to do (GLOSSARY.md "Grant", ADR-0041): its organization & workspace
// access and an access level per permission. One per OAuth connection, keyed by user and client rather
// than by `oauthConsent` row, because the provider can write duplicate consent rows for one connection
// and settles them later. Every MCP request reads it; without one the request is refused. The client's
// OAuth scopes stay on the consent and cap it.
type McpGrant struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	UserID    string   `gorm:"column:userId;uniqueIndex:McpGrant_userId_clientId_key,priority:1;not null" json:"userId"`
	ClientID  string   `gorm:"column:clientId;uniqueIndex:McpGrant_userId_clientId_key,priority:2;index:McpGrant_clientId_idx;not null" json:"clientId"`
	// Which organizations and workspaces the grant reaches (ENG-7153): All access, only the organizations
	// in `organizations` (and every workspace in them, later ones included), or only the workspaces in
	// `workspaces`. Enforced in SpiceDB through `can()` (ADR-0042); a value no row backs reaches nothing.
	Access string `gorm:"column:access;default:'all';not null" json:"access"`
	// The access level per permission of `apps/web/modules/mcp/permissions.ts`: `{ "surveys": "write" }`.
	// A permission that is absent has No access.
	Permissions datatypes.JSON `gorm:"column:permissions;default:'{}';not null" json:"permissions"`
	User        *User          `gorm:"belongsTo"`
	Client      *OauthClient   `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
	// The selections behind `organizations` and `workspaces` access (ENG-7153).
	Organizations []McpGrantOrganization `gorm:"foreignKey:GrantID;constraint:McpGrantOrganization_grantId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Workspaces    []McpGrantWorkspace    `gorm:"foreignKey:GrantID;constraint:McpGrantWorkspace_grantId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (McpGrant) TableName() string {
	return "McpGrant"
}

func (m *McpGrant) BeforeCreate(tx *gorm.DB) error {
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

func (*McpGrant) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// An organization a grant with `organizations` access reaches, workspaces created there later included.
type McpGrantOrganization struct {
	GrantID        string        `gorm:"column:grantId;primaryKey" json:"grantId"`
	OrganizationID string        `gorm:"column:organizationId;primaryKey;index:McpGrantOrganization_organizationId_idx" json:"organizationId"`
	Grant          *McpGrant     `gorm:"foreignKey:GrantID;belongsTo"`
	Organization   *Organization `gorm:"belongsTo"`
}

func (McpGrantOrganization) TableName() string {
	return "McpGrantOrganization"
}

// A workspace a grant with `workspaces` access reaches.
type McpGrantWorkspace struct {
	GrantID     string     `gorm:"column:grantId;primaryKey" json:"grantId"`
	WorkspaceID string     `gorm:"column:workspaceId;primaryKey;index:McpGrantWorkspace_workspaceId_idx" json:"workspaceId"`
	Grant       *McpGrant  `gorm:"foreignKey:GrantID;belongsTo"`
	Workspace   *Workspace `gorm:"belongsTo"`
}

func (McpGrantWorkspace) TableName() string {
	return "McpGrantWorkspace"
}

type OauthResource struct {
	ID                            string                `gorm:"column:id;primaryKey" json:"id"`
	Identifier                    string                `gorm:"column:identifier;uniqueIndex:oauthResource_identifier_key;not null" json:"identifier"`
	Name                          string                `gorm:"column:name;not null" json:"name"`
	AccessTokenTTL                *int                  `gorm:"column:accessTokenTtl;size:32" json:"accessTokenTtl"`
	RefreshTokenTTL               *int                  `gorm:"column:refreshTokenTtl;size:32" json:"refreshTokenTtl"`
	SigningAlgorithm              *string               `gorm:"column:signingAlgorithm" json:"signingAlgorithm"`
	SigningKeyID                  *string               `gorm:"column:signingKeyId" json:"signingKeyId"`
	AllowedScopes                 ScalarList[string]    `gorm:"column:allowedScopes;type:text[]" json:"allowedScopes"`
	CustomClaims                  datatypes.JSON        `gorm:"column:customClaims" json:"customClaims"`
	DpopBoundAccessTokensRequired *bool                 `gorm:"column:dpopBoundAccessTokensRequired;default:false" json:"dpopBoundAccessTokensRequired"`
	Disabled                      *bool                 `gorm:"column:disabled;default:false" json:"disabled"`
	CreatedAt                     *DateTime             `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false" json:"createdAt"`
	UpdatedAt                     *DateTime             `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false" json:"updatedAt"`
	PolicyVersion                 *int                  `gorm:"column:policyVersion;size:32;default:1" json:"policyVersion"`
	Metadata                      datatypes.JSON        `gorm:"column:metadata" json:"metadata"`
	Clients                       []OauthClientResource `gorm:"foreignKey:ResourceID;references:Identifier;constraint:oauthClientResource_resourceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OauthResource) TableName() string {
	return "oauthResource"
}

func (m *OauthResource) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
// `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
// endpoint, which is why the migration backfills one row per existing client.
// Both sides reference natural keys, matching the other oauth child tables.
type OauthClientResource struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	ClientID   string         `gorm:"column:clientId;uniqueIndex:oauthClientResource_clientId_resourceId_key,priority:1;not null" json:"clientId"`
	ResourceID string         `gorm:"column:resourceId;uniqueIndex:oauthClientResource_clientId_resourceId_key,priority:2;index:oauthClientResource_resourceId_idx;not null" json:"resourceId"`
	Metadata   datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt  *DateTime      `gorm:"column:createdAt;type:timestamp(3);autoCreateTime:false" json:"createdAt"`
	Client     *OauthClient   `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
	Resource   *OauthResource `gorm:"foreignKey:ResourceID;references:Identifier;belongsTo"`
}

func (OauthClientResource) TableName() string {
	return "oauthClientResource"
}

func (m *OauthClientResource) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
// assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
// table, so omitting it risks an adapter error on a code path we do not control.
type OauthClientAssertion struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	ExpiresAt DateTime `gorm:"column:expiresAt;index:oauthClientAssertion_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
}

func (OauthClientAssertion) TableName() string {
	return "oauthClientAssertion"
}

func (m *OauthClientAssertion) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Stores the active password reset token for a user.
// Tokens are opaque, hashed at rest, revocable, and single-use.
type PasswordResetToken struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	TokenHash string   `gorm:"column:token_hash;uniqueIndex:PasswordResetToken_token_hash_key;not null" json:"token_hash"`
	ExpiresAt DateTime `gorm:"column:expires_at;index:PasswordResetToken_expires_at_idx;type:timestamp(3);not null" json:"expires_at"`
	UserID    string   `gorm:"column:userId;uniqueIndex:PasswordResetToken_userId_key;not null" json:"userId"`
	User      *User    `gorm:"belongsTo"`
}

func (PasswordResetToken) TableName() string {
	return "PasswordResetToken"
}

func (m *PasswordResetToken) BeforeCreate(tx *gorm.DB) error {
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

func (*PasswordResetToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a user in the Formbricks system.
// Central model for user authentication and profile management.
//
// @property id - Unique identifier for the user
// @property name - Display name of the user
// @property email - User's email address
// @property twoFactorEnabled - Whether 2FA is active
// @property memberships - Organizations the user belongs to
// @property notificationSettings - User's notification preferences
type User struct {
	ID                        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt                 DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt                 DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name                      string   `gorm:"column:name;not null" json:"name"`
	Email                     string   `gorm:"column:email;uniqueIndex:User_email_key;not null" json:"email"`
	EmailVerified             bool     `gorm:"column:email_verified;default:false;not null" json:"email_verified"`
	TwoFactorSecret           *string  `gorm:"column:twoFactorSecret" json:"twoFactorSecret"`
	TwoFactorEnabled          bool     `gorm:"column:twoFactorEnabled;default:false;not null" json:"twoFactorEnabled"`
	BackupCodes               *string  `gorm:"column:backupCodes" json:"backupCodes"`
	Password                  *string  `gorm:"column:password" json:"password"`
	IdentityProvider          string   `gorm:"column:identityProvider;default:'email';not null" json:"identityProvider"`
	IdentityProviderAccountID *string  `gorm:"column:identityProviderAccountId" json:"identityProviderAccountId"`
	GroupID                   *string  `gorm:"column:groupId" json:"groupId"`
	// [UserNotificationSettings]
	NotificationSettings datatypes.JSON `gorm:"column:notificationSettings;default:'{}';not null" json:"notificationSettings"`
	// [Locale]
	Locale      *string   `gorm:"column:locale;default:'en-US';not null" json:"locale"`
	LastLoginAt *DateTime `gorm:"column:lastLoginAt;type:timestamp(3)" json:"lastLoginAt"`
	IsActive    *bool     `gorm:"column:isActive;default:true;not null" json:"isActive"`
	// When an admin last reactivated this user. Restarts the members retention clock (ENG-3713)
	// without touching `lastLoginAt`, which the member list, usage counts and APIs report as a real
	// sign-in.
	ReactivatedAt *DateTime `gorm:"column:reactivatedAt;type:timestamp(3)" json:"reactivatedAt"`
	// When the user last completed a sign-in (`lastActivePlugin`, after any second factor) or renewed a
	// session (Better Auth `databaseHooks.session.update`), on the database clock and never moved
	// backwards. The members retention clock reads it: sessions are
	// deleted on sign-out and expiry, so their own timestamps can't serve as a record of activity.
	LastActiveAt              *DateTime            `gorm:"column:lastActiveAt;type:timestamp(3)" json:"lastActiveAt"`
	IsBootstrapAdmin          *string              `gorm:"column:isBootstrapAdmin;uniqueIndex:User_isBootstrapAdmin_key" json:"isBootstrapAdmin"`
	Memberships               []Membership         `gorm:"foreignKey:UserID;constraint:Membership_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Accounts                  []Account            `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions                  []Session            `gorm:"foreignKey:UserID;constraint:Session_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OauthClients              []OauthClient        `gorm:"foreignKey:UserID;constraint:oauthClient_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OauthAccessTokens         []OauthAccessToken   `gorm:"foreignKey:UserID;constraint:oauthAccessToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OauthRefreshTokens        []OauthRefreshToken  `gorm:"foreignKey:UserID;constraint:oauthRefreshToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	OauthConsents             []OauthConsent       `gorm:"foreignKey:UserID;constraint:oauthConsent_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	McpGrants                 []McpGrant           `gorm:"foreignKey:UserID;constraint:McpGrant_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	InvitesCreated            []Invite             `gorm:"foreignKey:CreatorID;constraint:Invite_creatorId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	InvitesAccepted           []Invite             `gorm:"foreignKey:AcceptorID;constraint:Invite_acceptorId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Surveys                   []Survey             `gorm:"foreignKey:CreatedBy;constraint:Survey_createdBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	OwnedSurveys              []Survey             `gorm:"foreignKey:OwnerID;constraint:Survey_ownerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Charts                    []Chart              `gorm:"foreignKey:CreatedBy;constraint:Chart_createdBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Dashboards                []Dashboard          `gorm:"foreignKey:CreatedBy;constraint:Dashboard_createdBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Creatives                 []Creative           `gorm:"foreignKey:CreatedBy;constraint:Creative_createdBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	FeedbackSources           []FeedbackSource     `gorm:"foreignKey:CreatedBy;constraint:FeedbackSource_created_by_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TeamUsers                 []TeamUser           `gorm:"foreignKey:UserID;constraint:TeamUser_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Workflows                 []Workflow           `gorm:"foreignKey:CreatedBy;constraint:Workflow_createdBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	WorkflowVersionsPublished []WorkflowVersion    `gorm:"foreignKey:PublishedBy;constraint:WorkflowVersion_publishedBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	RetentionPoliciesUpdated  []RetentionPolicy    `gorm:"foreignKey:UpdatedByID;constraint:RetentionPolicy_updatedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ExemptionsCreated         []RetentionExemption `gorm:"foreignKey:CreatedByID;constraint:RetentionExemption_createdById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ExemptionsRevoked         []RetentionExemption `gorm:"foreignKey:RevokedByID;constraint:RetentionExemption_revokedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	RetentionNotices          []RetentionNotice    `gorm:"foreignKey:UserID;constraint:RetentionNotice_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AttachmentExports         []AttachmentExport   `gorm:"foreignKey:RequestedByID;constraint:AttachmentExport_requestedById_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TwoFactor                 *TwoFactor           `gorm:"foreignKey:UserID;constraint:TwoFactor_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PasswordResetToken        *PasswordResetToken  `gorm:"foreignKey:UserID;constraint:PasswordResetToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	return nil
}

func (*User) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Defines a segment of contacts based on attributes.
// Used for targeting surveys to specific user groups.
//
// @property id - Unique identifier for the segment
// @property title - Display name of the segment
// @property filters - Rules defining the segment
// @property isPrivate - Whether the segment is private
// @property workspace - The workspace this segment belongs to
type Segment struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Title       string   `gorm:"column:title;uniqueIndex:Segment_workspaceId_title_key,priority:2;not null" json:"title"`
	Description *string  `gorm:"column:description" json:"description"`
	IsPrivate   *bool    `gorm:"column:isPrivate;default:true;not null" json:"isPrivate"`
	// [SegmentFilter]
	Filters     datatypes.JSON `gorm:"column:filters;default:'[]';not null" json:"filters"`
	WorkspaceID string         `gorm:"column:workspaceId;uniqueIndex:Segment_workspaceId_title_key,priority:1;index:Segment_workspaceId_idx;not null" json:"workspaceId"`
	Workspace   *Workspace     `gorm:"belongsTo"`
	Surveys     []Survey       `gorm:"foreignKey:SegmentID;constraint:Survey_segmentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Segment) TableName() string {
	return "Segment"
}

func (m *Segment) BeforeCreate(tx *gorm.DB) error {
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

func (*Segment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a supported language in the system.
// Used for multilingual survey support.
//
// @property id - Unique identifier for the language
// @property code - Language code (e.g., 'en-US')
// @property alias - Optional friendly name
// @property workspace - The workspace this language is enabled for
type Language struct {
	ID                string             `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt         DateTime           `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt         DateTime           `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Code              string             `gorm:"column:code;uniqueIndex:Language_workspaceId_code_key,priority:2;not null" json:"code"`
	Alias             *string            `gorm:"column:alias" json:"alias"`
	WorkspaceID       string             `gorm:"column:workspaceId;uniqueIndex:Language_workspaceId_code_key,priority:1;not null" json:"workspaceId"`
	Workspace         *Workspace         `gorm:"belongsTo"`
	SurveyLanguages   []SurveyLanguage   `gorm:"foreignKey:LanguageID;constraint:SurveyLanguage_languageId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreativeLanguages []CreativeLanguage `gorm:"foreignKey:LanguageID;constraint:CreativeLanguage_languageId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Language) TableName() string {
	return "Language"
}

func (m *Language) BeforeCreate(tx *gorm.DB) error {
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

func (*Language) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Links surveys to their supported languages.
// Manages which languages are available for each survey.
//
// @property language - The supported language
// @property survey - The associated survey
// @property default - Whether this is the default language
// @property enabled - Whether this language is currently active
type SurveyLanguage struct {
	LanguageID string    `gorm:"column:languageId;primaryKey" json:"languageId"`
	SurveyID   string    `gorm:"column:surveyId;primaryKey;index:SurveyLanguage_surveyId_idx" json:"surveyId"`
	Default    bool      `gorm:"column:default;default:false;not null" json:"default"`
	Enabled    *bool     `gorm:"column:enabled;default:true;not null" json:"enabled"`
	Language   *Language `gorm:"belongsTo"`
	Survey     *Survey   `gorm:"belongsTo"`
}

func (SurveyLanguage) TableName() string {
	return "SurveyLanguage"
}

// Defines one Embedded Data field: the extra context stored on a response that isn't an
// answer to a question (brand, plan, page type, a computed score, ...).
//
// A field is either **local** — used by a single survey, which is how existing variables and
// hidden fields are migrated — or **shared**, meaning it is part of the workspace library and can
// be linked by many surveys. There is no stored flag for that: `surveyId` is the owner and `key`
// is the library name, so a local field has `surveyId` set and `key` null, and a shared field the
// other way round. Deriving it rather than storing a boolean removes a column that could drift
// out of step with the two fields it stands for.
//
// @property key - Library name, unique per workspace and immutable. Null for local fields
// @property name - Human-readable display label
// @property source - Where the value comes from (computed | ingested). `reserved` is part of the
// shared enum but never stored as a row: those fields are a static catalog in code
// @property dataType - Declared value type, used for coercion at ingest
// @property defaultValue - Initial value (computed) or fallback when nothing arrives (ingested)
// @property locked - Ingested only: ignore external writes and use the default instead
// @property survey - The owning survey for a local field; cascades so nothing is orphaned
type EmbeddedData struct {
	ID          string   `gorm:"column:id;primaryKey;uniqueIndex:EmbeddedData_id_workspaceId_key,priority:1" json:"id"`
	CreatedAt   DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Key         *string  `gorm:"column:key;uniqueIndex:EmbeddedData_workspaceId_key_key,priority:2" json:"key"`
	Name        string   `gorm:"column:name;not null" json:"name"`
	Description *string  `gorm:"column:description" json:"description"`
	Source      string   `gorm:"column:source;not null" json:"source"`
	DataType    string   `gorm:"column:dataType;default:'string';not null" json:"dataType"`
	// [EmbeddedDataDefaultValue]
	DefaultValue datatypes.JSON       `gorm:"column:defaultValue" json:"defaultValue"`
	Locked       bool                 `gorm:"column:locked;default:false;not null" json:"locked"`
	SurveyID     *string              `gorm:"column:surveyId;index:EmbeddedData_surveyId_idx" json:"surveyId"`
	WorkspaceID  string               `gorm:"column:workspaceId;uniqueIndex:EmbeddedData_workspaceId_key_key,priority:1;uniqueIndex:EmbeddedData_id_workspaceId_key,priority:2;not null" json:"workspaceId"`
	Survey       *Survey              `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
	Workspace    *Workspace           `gorm:"belongsTo"`
	SurveyLinks  []SurveyEmbeddedData `gorm:"foreignKey:EmbeddedDataID,WorkspaceID;references:ID,WorkspaceID;constraint:SurveyEmbeddedData_embeddedDataId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (EmbeddedData) TableName() string {
	return "EmbeddedData"
}

func (m *EmbeddedData) BeforeCreate(tx *gorm.DB) error {
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

func (*EmbeddedData) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Links a survey to an Embedded Data field and records the key that field's value is stored
// under inside that survey.
//
// `workspaceId` is part of both foreign keys on purpose: `surveyId` and `embeddedDataId` alone are
// independent, so nothing would stop a survey in one workspace linking a field defined in another.
// Referencing `(id, workspaceId)` on both sides makes a cross-workspace pair unrepresentable, and
// puts the workspace scope every query needs on the row instead of behind a join through Survey.
//
// @property storageKey - The key the value lives under: a cuid for computed fields, the field
// name for ingested ones. Migrated fields keep their original key, which is what lets existing
// recall tokens, logic conditions and stored responses keep resolving untouched.
//
// @property order - Position within this survey's field list, dense from 0. Lives on the link
// rather than on `EmbeddedData` because a shared definition sits at a different position in each
// survey that links it. Deliberately has no default: every writer knows the position it is writing,
// so a missing one should be a type error rather than a row that silently sorts to the front.
type SurveyEmbeddedData struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	WorkspaceID    string        `gorm:"column:workspaceId;not null" json:"workspaceId"`
	SurveyID       string        `gorm:"column:surveyId;uniqueIndex:SurveyEmbeddedData_surveyId_embeddedDataId_key,priority:1;uniqueIndex:SurveyEmbeddedData_surveyId_storageKey_key,priority:1;index:SurveyEmbeddedData_surveyId_order_idx,priority:1;not null" json:"surveyId"`
	EmbeddedDataID string        `gorm:"column:embeddedDataId;uniqueIndex:SurveyEmbeddedData_surveyId_embeddedDataId_key,priority:2;index:SurveyEmbeddedData_embeddedDataId_idx;not null" json:"embeddedDataId"`
	StorageKey     string        `gorm:"column:storageKey;uniqueIndex:SurveyEmbeddedData_surveyId_storageKey_key,priority:2;not null" json:"storageKey"`
	Order          int           `gorm:"column:order;index:SurveyEmbeddedData_surveyId_order_idx,priority:2;size:32;not null" json:"order"`
	Survey         *Survey       `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
	EmbeddedData   *EmbeddedData `gorm:"foreignKey:EmbeddedDataID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
}

func (SurveyEmbeddedData) TableName() string {
	return "SurveyEmbeddedData"
}

func (m *SurveyEmbeddedData) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Represents a team within an organization.
// Enables group-based access control and collaboration.
//
// @property id - Unique identifier for the team
// @property name - Display name of the team
// @property organization - The parent organization
// @property teamUsers - Users who are part of this team
// @property workspaceTeams - Workspaces this team has access to
type Team struct {
	ID             string          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime        `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt      DateTime        `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name           string          `gorm:"column:name;uniqueIndex:Team_organizationId_name_key,priority:2;not null" json:"name"`
	OrganizationID string          `gorm:"column:organizationId;uniqueIndex:Team_organizationId_name_key,priority:1;not null" json:"organizationId"`
	Organization   *Organization   `gorm:"belongsTo"`
	TeamUsers      []TeamUser      `gorm:"foreignKey:TeamID;constraint:TeamUser_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	WorkspaceTeams []WorkspaceTeam `gorm:"foreignKey:TeamID;constraint:WorkspaceTeam_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Team) TableName() string {
	return "Team"
}

func (m *Team) BeforeCreate(tx *gorm.DB) error {
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

func (*Team) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Links users to teams with specific roles.
// Manages team membership and permissions.
//
// @property team - The associated team
// @property user - The team member
// @property role - User's role within the team
type TeamUser struct {
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	TeamID    string   `gorm:"column:teamId;primaryKey" json:"teamId"`
	UserID    string   `gorm:"column:userId;primaryKey;index:TeamUser_userId_idx" json:"userId"`
	Role      string   `gorm:"column:role;not null" json:"role"`
	Team      *Team    `gorm:"belongsTo"`
	User      *User    `gorm:"belongsTo"`
}

func (TeamUser) TableName() string {
	return "TeamUser"
}

func (m *TeamUser) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*TeamUser) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Defines team access to specific workspaces.
// Manages workspace-level permissions for teams.
//
// @property workspace - The accessed workspace
// @property team - The team receiving access
// @property permission - Level of access granted
type WorkspaceTeam struct {
	CreatedAt   DateTime   `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime   `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	WorkspaceID string     `gorm:"column:workspaceId;primaryKey" json:"workspaceId"`
	TeamID      string     `gorm:"column:teamId;primaryKey;index:WorkspaceTeam_teamId_idx" json:"teamId"`
	Permission  string     `gorm:"column:permission;default:'read';not null" json:"permission"`
	Workspace   *Workspace `gorm:"belongsTo"`
	Team        *Team      `gorm:"belongsTo"`
}

func (WorkspaceTeam) TableName() string {
	return "WorkspaceTeam"
}

func (m *WorkspaceTeam) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkspaceTeam) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a chart/visualization that can be used in multiple dashboards.
// Charts are reusable components that query analytics data.
//
// @property id - Unique identifier for the chart
// @property name - Display name of the chart
// @property type - Type of visualization (bar, line, pie, etc.)
// @property workspace - The workspace this chart belongs to
// @property query - Cube.js query configuration (JSON)
// @property config - Chart-specific configuration (colors, labels, etc.)
// @property createdBy - User who created the chart
// @property dashboards - Dashboards that use this chart
type Chart struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:created_at;index:Chart_workspaceId_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name        string   `gorm:"column:name;uniqueIndex:Chart_workspaceId_name_key,priority:2;not null" json:"name"`
	Type        string   `gorm:"column:type;not null" json:"type"`
	WorkspaceID string   `gorm:"column:workspaceId;uniqueIndex:Chart_workspaceId_name_key,priority:1;index:Chart_workspaceId_created_at_idx,priority:1;not null" json:"workspaceId"`
	// [ChartQuery] - Cube.js query configuration
	Query datatypes.JSON `gorm:"column:query;default:'{}';not null" json:"query"`
	// [ChartConfig] - Visualization configuration (colors, labels, formatting)
	Config              datatypes.JSON     `gorm:"column:config;default:'{}';not null" json:"config"`
	CreatedBy           *string            `gorm:"column:createdBy" json:"createdBy"`
	FeedbackDirectoryID string             `gorm:"column:feedbackDirectoryId;index:Chart_feedbackDirectoryId_idx;not null" json:"feedbackDirectoryId"`
	Workspace           *Workspace         `gorm:"belongsTo"`
	Creator             *User              `gorm:"foreignKey:CreatedBy;belongsTo"`
	FeedbackDirectory   *FeedbackDirectory `gorm:"belongsTo"`
	Widgets             []DashboardWidget  `gorm:"foreignKey:ChartID;constraint:DashboardWidget_chartId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Chart) TableName() string {
	return "Chart"
}

func (m *Chart) BeforeCreate(tx *gorm.DB) error {
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

func (*Chart) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A workspace's display name for an imported (embedded-data) key of its Feedback Datasets, shown in
// place of the key in chart legends, headers and the chart builder's field pickers. One per key and
// workspace; clearing the name deletes the row, and the key shows as collected again. Built-in fields
// are never renamed here: they keep their translated names.
//
// @property key - The embedded-data key exactly as collected, case included
// @property label - The name shown for the key across the workspace
type EmbeddedFieldLabel struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime   `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime   `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	WorkspaceID string     `gorm:"column:workspaceId;uniqueIndex:EmbeddedFieldLabel_workspaceId_key_key,priority:1;not null" json:"workspaceId"`
	Key         string     `gorm:"column:key;uniqueIndex:EmbeddedFieldLabel_workspaceId_key_key,priority:2;not null" json:"key"`
	Label       string     `gorm:"column:label;not null" json:"label"`
	Workspace   *Workspace `gorm:"belongsTo"`
}

func (EmbeddedFieldLabel) TableName() string {
	return "EmbeddedFieldLabel"
}

func (m *EmbeddedFieldLabel) BeforeCreate(tx *gorm.DB) error {
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

func (*EmbeddedFieldLabel) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a dashboard containing multiple charts.
// Dashboards aggregate analytics insights at the workspace level.
//
// @property id - Unique identifier for the dashboard
// @property name - Display name of the dashboard
// @property workspace - The workspace this dashboard belongs to
// @property widgets - Charts on this dashboard
// @property filters - The dashboard filters pinned to its top bar, in order: `{ id, field, operator, label }`,
// at most 10 (ADR-0026). Each viewer's picks live in their own browser, not here.
// @property createdBy - User who created the dashboard
type Dashboard struct {
	ID          string            `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime          `gorm:"column:created_at;index:Dashboard_workspaceId_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime          `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name        string            `gorm:"column:name;uniqueIndex:Dashboard_workspaceId_name_key,priority:2;not null" json:"name"`
	WorkspaceID string            `gorm:"column:workspaceId;uniqueIndex:Dashboard_workspaceId_name_key,priority:1;index:Dashboard_workspaceId_created_at_idx,priority:1;not null" json:"workspaceId"`
	CreatedBy   *string           `gorm:"column:createdBy" json:"createdBy"`
	Filters     datatypes.JSON    `gorm:"column:filters;default:'[]';not null" json:"filters"`
	Workspace   *Workspace        `gorm:"belongsTo"`
	Creator     *User             `gorm:"foreignKey:CreatedBy;belongsTo"`
	Widgets     []DashboardWidget `gorm:"foreignKey:DashboardID;constraint:DashboardWidget_dashboardId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Dashboard) TableName() string {
	return "Dashboard"
}

func (m *Dashboard) BeforeCreate(tx *gorm.DB) error {
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

func (*Dashboard) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Represents a chart widget on a dashboard.
// Widgets are positioned using a grid layout system.
//
// @property id - Unique identifier for the widget
// @property dashboard - The dashboard this widget belongs to
// @property chart - The chart displayed in this widget (chart name is used for display)
// @property layout - Grid layout configuration (x, y, width, height)
// @property order - Display order within the dashboard
type DashboardWidget struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	DashboardID string   `gorm:"column:dashboardId;uniqueIndex:DashboardWidget_dashboardId_chartId_key,priority:1;index:DashboardWidget_dashboardId_order_idx,priority:1;not null" json:"dashboardId"`
	ChartID     string   `gorm:"column:chartId;uniqueIndex:DashboardWidget_dashboardId_chartId_key,priority:2;not null" json:"chartId"`
	// [WidgetLayout] - Grid layout: { x, y, w, h }
	Layout    datatypes.JSON `gorm:"column:layout;default:'{\"x\":0,\"y\":0,\"w\":4,\"h\":3}';not null" json:"layout"`
	Order     int            `gorm:"column:order;index:DashboardWidget_dashboardId_order_idx,priority:2;size:32;default:0;not null" json:"order"`
	Dashboard *Dashboard     `gorm:"belongsTo"`
	Chart     *Chart         `gorm:"belongsTo"`
}

func (DashboardWidget) TableName() string {
	return "DashboardWidget"
}

func (m *DashboardWidget) BeforeCreate(tx *gorm.DB) error {
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

func (*DashboardWidget) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
// storage files that can't be removed inside the database transaction. A row is written in the same
// transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
// queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
type DeletionCleanup struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:created_at;index:DeletionCleanup_organizationId_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Kind      string   `gorm:"column:kind;not null" json:"kind"`
	// Scoping and logs. Not foreign keys: the rows they name are usually gone already.
	OrganizationID string `gorm:"column:organizationId;index:DeletionCleanup_organizationId_created_at_idx,priority:1;not null" json:"organizationId"`
	WorkspaceID    string `gorm:"column:workspaceId;not null" json:"workspaceId"`
	SurveyID       string `gorm:"column:surveyId;not null" json:"surveyId"`
	// Hub kinds: the organisation's feedback directories, read inside the delete's transaction.
	TenantIds ScalarList[string] `gorm:"column:tenantIds;type:text[]" json:"tenantIds"`
	// `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response.
	ResponseIds ScalarList[string] `gorm:"column:responseIds;type:text[]" json:"responseIds"`
	// `storageFiles`: the object keys or URLs to delete.
	FileKeys ScalarList[string] `gorm:"column:fileKeys;type:text[]" json:"fileKeys"`
	Attempts int                `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	// Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup.
	NextAttemptAt DateTime `gorm:"column:nextAttemptAt;index:DeletionCleanup_nextAttemptAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"nextAttemptAt"`
	// The last failure: a short code and message, never record or file content.
	LastError *string `gorm:"column:lastError" json:"lastError"`
}

func (DeletionCleanup) TableName() string {
	return "DeletionCleanup"
}

func (m *DeletionCleanup) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	if m.NextAttemptAt.IsZero() {
		m.NextAttemptAt = DateTime{Time: now}
	}
	return nil
}

func (*DeletionCleanup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A survey's response attachments built into one ZIP in object storage by a background job and
// downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
// Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
// survey deletes it with the survey's storage prefix.
//
// A partial unique index in the migration allows one queued or running export per requester and
// survey, so a double click returns the export already in flight.
type AttachmentExport struct {
	ID            string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt     DateTime `gorm:"column:created_at;index:AttachmentExport_surveyId_requestedById_created_at_idx,priority:3;index:AttachmentExport_status_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt     DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	SurveyID      string   `gorm:"column:surveyId;index:AttachmentExport_surveyId_requestedById_created_at_idx,priority:1;not null" json:"surveyId"`
	WorkspaceID   string   `gorm:"column:workspaceId;index:AttachmentExport_workspaceId_idx;not null" json:"workspaceId"`
	RequestedByID *string  `gorm:"column:requestedById;index:AttachmentExport_surveyId_requestedById_created_at_idx,priority:2" json:"requestedById"`
	Status        string   `gorm:"column:status;index:AttachmentExport_status_created_at_idx,priority:1;default:'queued';not null" json:"status"`
	// The response filter at request time, a `TResponseFilterCriteria`; null exports every response.
	FilterCriteria datatypes.JSON `gorm:"column:filterCriteria" json:"filterCriteria"`
	// IANA zone the archive's folder names are stamped in: the org's display zone at request time.
	TimeZone string `gorm:"column:timeZone;not null" json:"timeZone"`
	// Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
	// older pod dropped as unknown during a rolling upgrade.
	DispatchCount int `gorm:"column:dispatchCount;size:32;default:0;not null" json:"dispatchCount"`
	// The pre-flight's counts, shown while the archive builds.
	FileCount     int `gorm:"column:fileCount;size:32;not null" json:"fileCount"`
	ResponseCount int `gorm:"column:responseCount;size:32;not null" json:"responseCount"`
	// What the archive actually holds, once built.
	IncludedFileCount *int    `gorm:"column:includedFileCount;size:32" json:"includedFileCount"`
	TotalBytes        *int64  `gorm:"column:totalBytes" json:"totalBytes"`
	Truncated         bool    `gorm:"column:truncated;default:false;not null" json:"truncated"`
	StorageKey        *string `gorm:"column:storageKey" json:"storageKey"`
	// A short code (`stale`, `storage_upload_failed`, …), never a message or file name.
	ErrorCode   *string    `gorm:"column:errorCode" json:"errorCode"`
	StartedAt   *DateTime  `gorm:"column:startedAt;type:timestamp(3)" json:"startedAt"`
	FinishedAt  *DateTime  `gorm:"column:finishedAt;type:timestamp(3)" json:"finishedAt"`
	ExpiresAt   *DateTime  `gorm:"column:expiresAt;index:AttachmentExport_expiresAt_idx;type:timestamp(3)" json:"expiresAt"`
	Survey      *Survey    `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
	Workspace   *Workspace `gorm:"belongsTo"`
	RequestedBy *User      `gorm:"foreignKey:RequestedByID;belongsTo"`
}

func (AttachmentExport) TableName() string {
	return "AttachmentExport"
}

func (m *AttachmentExport) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AttachmentExport) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// A feedback source — connects external data into Unify Feedback for record creation.
// Naming note: "FeedbackSource" is the Unify Feedback domain term; the word "Connector"
// is reserved for the Workflows domain (see docs/development/standards/practices/naming-conventions.mdx).
//
// @property id - Unique identifier for the feedback source
// @property name - Display name for the feedback source
// @property type - Type of feedback source (formbricks_survey, csv, …)
// @property status - Current state of the feedback source (active, paused, error)
// @property workspace - The workspace this feedback source belongs to
// @property feedbackDirectory - Hub directory the feedback source writes into
// @property formbricksMappings - Element mappings for Formbricks feedback sources
// @property fieldMappings - Field mappings for other feedback source types
type FeedbackSource struct {
	ID        string   `gorm:"column:id;primaryKey;uniqueIndex:FeedbackSource_id_workspaceId_key,priority:1" json:"id"`
	CreatedAt DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name      string   `gorm:"column:name;uniqueIndex:FeedbackSource_workspaceId_name_key,priority:2;not null" json:"name"`
	Type      string   `gorm:"column:type;index:FeedbackSource_type_idx;not null" json:"type"`
	Status    string   `gorm:"column:status;default:'active';not null" json:"status"`
	// Defaults to completedOnly so a new source agrees with the live ingestion path, which only
	// fires on responseFinished. Existing rows take the default; nothing is re-imported.
	ImportMode                 string                            `gorm:"column:import_mode;default:'completedOnly';not null" json:"import_mode"`
	ElementScope               string                            `gorm:"column:elementScope;default:'specific';not null" json:"elementScope"`
	WorkspaceID                string                            `gorm:"column:workspaceId;uniqueIndex:FeedbackSource_id_workspaceId_key,priority:2;uniqueIndex:FeedbackSource_workspaceId_name_key,priority:1;index:FeedbackSource_feedbackDirectoryId_workspaceId_idx,priority:2;not null" json:"workspaceId"`
	FeedbackDirectoryID        string                            `gorm:"column:feedbackDirectoryId;index:FeedbackSource_feedbackDirectoryId_workspaceId_idx,priority:1;not null" json:"feedbackDirectoryId"`
	LastSyncAt                 *DateTime                         `gorm:"column:last_sync_at;type:timestamp(3)" json:"last_sync_at"`
	CreatedBy                  *string                           `gorm:"column:created_by" json:"created_by"`
	Workspace                  *Workspace                        `gorm:"belongsTo"`
	FeedbackDirectory          *FeedbackDirectory                `gorm:"belongsTo"`
	Creator                    *User                             `gorm:"foreignKey:CreatedBy;belongsTo"`
	FeedbackDirectoryWorkspace *FeedbackDirectoryWorkspace       `gorm:"foreignKey:FeedbackDirectoryID,WorkspaceID;references:FeedbackDirectoryID,WorkspaceID;belongsTo"`
	FormbricksMappings         []FeedbackSourceFormbricksMapping `gorm:"foreignKey:FeedbackSourceID,WorkspaceID;references:ID,WorkspaceID;constraint:FeedbackSourceFormbricksMapping_feedbackSourceId_workspaceId_fk,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FieldMappings              []FeedbackSourceFieldMapping      `gorm:"foreignKey:FeedbackSourceID,WorkspaceID;references:ID,WorkspaceID;constraint:FeedbackSourceFieldMapping_feedbackSourceId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (FeedbackSource) TableName() string {
	return "FeedbackSource"
}

func (m *FeedbackSource) BeforeCreate(tx *gorm.DB) error {
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

func (*FeedbackSource) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Maps survey elements to Hub FeedbackRecords for Formbricks feedback sources.
// Each row represents one element that will create FeedbackRecords when answered.
//
// @property id - Unique identifier for the mapping
// @property feedbackSource - The parent feedback source
// @property survey - The survey containing the element
// @property elementId - The element ID within the survey (from blocks[].elements[].id)
// @property hubFieldType - The field_type to use in Hub (text, nps, rating, etc.)
// @property customFieldLabel - Optional override for the element headline as field_label in Hub
type FeedbackSourceFormbricksMapping struct {
	ID               string          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt        DateTime        `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	FeedbackSourceID string          `gorm:"column:feedback_source_id;uniqueIndex:FeedbackSourceFormbricksMapping_workspaceId_feedbackSourceId_su,priority:2;not null" json:"feedback_source_id"`
	WorkspaceID      string          `gorm:"column:workspaceId;uniqueIndex:FeedbackSourceFormbricksMapping_workspaceId_feedbackSourceId_su,priority:1;index:FeedbackSourceFormbricksMapping_workspaceId_surveyId_idx,priority:1;not null" json:"workspaceId"`
	SurveyID         string          `gorm:"column:surveyId;uniqueIndex:FeedbackSourceFormbricksMapping_workspaceId_feedbackSourceId_su,priority:3;index:FeedbackSourceFormbricksMapping_workspaceId_surveyId_idx,priority:2;index:FeedbackSourceFormbricksMapping_surveyId_idx;not null" json:"surveyId"`
	ElementID        string          `gorm:"column:elementId;uniqueIndex:FeedbackSourceFormbricksMapping_workspaceId_feedbackSourceId_su,priority:4;not null" json:"elementId"`
	HubFieldType     string          `gorm:"column:hubFieldType;not null" json:"hubFieldType"`
	CustomFieldLabel *string         `gorm:"column:custom_field_label" json:"custom_field_label"`
	FeedbackSource   *FeedbackSource `gorm:"foreignKey:FeedbackSourceID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
	Survey           *Survey         `gorm:"foreignKey:SurveyID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
}

func (FeedbackSourceFormbricksMapping) TableName() string {
	return "FeedbackSourceFormbricksMapping"
}

func (m *FeedbackSourceFormbricksMapping) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
// Maps source fields to Hub FeedbackRecord fields.
//
// @property id - Unique identifier for the mapping
// @property feedbackSource - The parent feedback source
// @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
// @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
// @property staticValue - If set, use this value instead of reading from sourceFieldId
type FeedbackSourceFieldMapping struct {
	ID               string          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt        DateTime        `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	FeedbackSourceID string          `gorm:"column:feedback_source_id;uniqueIndex:FeedbackSourceFieldMapping_workspaceId_feedbackSourceId_sourceF,priority:2;not null" json:"feedback_source_id"`
	WorkspaceID      string          `gorm:"column:workspaceId;uniqueIndex:FeedbackSourceFieldMapping_workspaceId_feedbackSourceId_sourceF,priority:1;not null" json:"workspaceId"`
	SourceFieldID    string          `gorm:"column:source_field_id;uniqueIndex:FeedbackSourceFieldMapping_workspaceId_feedbackSourceId_sourceF,priority:3;not null" json:"source_field_id"`
	TargetFieldID    string          `gorm:"column:target_field_id;uniqueIndex:FeedbackSourceFieldMapping_workspaceId_feedbackSourceId_sourceF,priority:4;not null" json:"target_field_id"`
	StaticValue      *string         `gorm:"column:static_value" json:"static_value"`
	FeedbackSource   *FeedbackSource `gorm:"foreignKey:FeedbackSourceID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
}

func (FeedbackSourceFieldMapping) TableName() string {
	return "FeedbackSourceFieldMapping"
}

func (m *FeedbackSourceFieldMapping) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Represents a feedback directory (Hub tenant) owned by an organization.
// Directories group feedback data and are assigned to workspaces for access control.
//
// @property id - Unique identifier for the directory
// @property name - Display name of the directory
// @property isArchived - Soft delete flag
// @property organization - The parent organization
// @property workspaces - Workspaces assigned to this directory
type FeedbackDirectory struct {
	ID              string                       `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt       DateTime                     `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt       DateTime                     `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name            string                       `gorm:"column:name;uniqueIndex:FeedbackDirectory_organizationId_name_key,priority:2;not null" json:"name"`
	IsArchived      bool                         `gorm:"column:isArchived;default:false;not null" json:"isArchived"`
	OrganizationID  string                       `gorm:"column:organizationId;uniqueIndex:FeedbackDirectory_organizationId_name_key,priority:1;not null" json:"organizationId"`
	Organization    *Organization                `gorm:"belongsTo"`
	Workspaces      []FeedbackDirectoryWorkspace `gorm:"foreignKey:FeedbackDirectoryID;constraint:FeedbackDirectoryWorkspace_feedbackDirectoryId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FeedbackSources []FeedbackSource             `gorm:"foreignKey:FeedbackDirectoryID;constraint:FeedbackSource_feedbackDirectoryId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Charts          []Chart                      `gorm:"foreignKey:FeedbackDirectoryID;constraint:Chart_feedbackDirectoryId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (FeedbackDirectory) TableName() string {
	return "FeedbackDirectory"
}

func (m *FeedbackDirectory) BeforeCreate(tx *gorm.DB) error {
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

func (*FeedbackDirectory) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Links feedback directories to workspaces (workspaces).
// Manages which workspaces can access a given directory.
//
// @property feedbackDirectory - The directory being accessed
// @property workspace - The workspace receiving access
type FeedbackDirectoryWorkspace struct {
	CreatedAt           DateTime           `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt           DateTime           `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	FeedbackDirectoryID string             `gorm:"column:feedbackDirectoryId;primaryKey" json:"feedbackDirectoryId"`
	WorkspaceID         string             `gorm:"column:workspaceId;primaryKey;index:FeedbackDirectoryWorkspace_workspaceId_idx" json:"workspaceId"`
	FeedbackDirectory   *FeedbackDirectory `gorm:"belongsTo"`
	Workspace           *Workspace         `gorm:"belongsTo"`
	FeedbackSources     []FeedbackSource   `gorm:"foreignKey:FeedbackDirectoryID,WorkspaceID;references:FeedbackDirectoryID,WorkspaceID;constraint:FeedbackSource_feedbackDirectoryId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:NO ACTION"`
}

func (FeedbackDirectoryWorkspace) TableName() string {
	return "FeedbackDirectoryWorkspace"
}

func (m *FeedbackDirectoryWorkspace) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*FeedbackDirectoryWorkspace) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// ENG-3282: per-scope readiness of a SpiceDB projection. A row with `readyAt` set means the scope was
// backfilled and audited clean, so the evaluator may enforce it. Written only by the backfill CLI.
type AuthzedProjectionScopeState struct {
	Scope     string    `gorm:"column:scope;primaryKey" json:"scope"`
	ReadyAt   *DateTime `gorm:"column:readyAt;type:timestamp(3)" json:"readyAt"`
	ReadyBy   *string   `gorm:"column:readyBy" json:"readyBy"`
	UpdatedAt DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AuthzedProjectionScopeState) TableName() string {
	return "AuthzedProjectionScopeState"
}

func (m *AuthzedProjectionScopeState) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AuthzedProjectionScopeState) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Durable PostgreSQL queue for rebuilding SpiceDB relationships after an authorization source changes.
// Source-table triggers insert these rows in the same transaction as the mutation. Identifiers are
// operationally sensitive and must never be included in logs, metrics, or CLI output.
//
// This model declares NO indexes on purpose. All three live in
// `migration/20260818120000_add_authzed_projection_outbox/migration.sql`, because each one is
// partial (`WHERE "processedAt" IS NULL AND ...`) and Prisma's `@@index` has no `where`. Declaring
// approximate non-partial copies here would be worse than declaring none: `prisma db push` would
// see a name match with a different definition, drop and recreate them without the predicate, and
// the migration's `CREATE INDEX IF NOT EXISTS` would then skip them — leaving a silently wrong
// shape that carries a week of delivered rows on every hot-path lookup. Deployments are unaffected
// either way (`prisma migrate deploy` never reads this file); a dev who runs `db push` restores
// them by rerunning the migration, which is the same contract its triggers already live under.
type AuthzedProjectionOutbox struct {
	ID           string  `gorm:"column:id;primaryKey" json:"id"`
	TargetType   string  `gorm:"column:targetType;not null" json:"targetType"`
	PrimaryID    string  `gorm:"column:primaryId;not null" json:"primaryId"`
	SecondaryID  *string `gorm:"column:secondaryId" json:"secondaryId"`
	IsRevocation bool    `gorm:"column:isRevocation;default:false;not null" json:"isRevocation"`
	// Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried.
	Attempts int `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	// Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
	// that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
	PermanentFailures int       `gorm:"column:permanentFailures;size:32;default:0;not null" json:"permanentFailures"`
	AvailableAt       DateTime  `gorm:"column:availableAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"availableAt"`
	LeasedAt          *DateTime `gorm:"column:leasedAt;type:timestamp(3)" json:"leasedAt"`
	LeaseExpiresAt    *DateTime `gorm:"column:leaseExpiresAt;type:timestamp(3)" json:"leaseExpiresAt"`
	LeaseOwner        *string   `gorm:"column:leaseOwner" json:"leaseOwner"`
	ProcessedAt       *DateTime `gorm:"column:processedAt;type:timestamp(3)" json:"processedAt"`
	DeadLetteredAt    *DateTime `gorm:"column:deadLetteredAt;type:timestamp(3)" json:"deadLetteredAt"`
	LastAttemptAt     *DateTime `gorm:"column:lastAttemptAt;type:timestamp(3)" json:"lastAttemptAt"`
	LastErrorCode     *string   `gorm:"column:lastErrorCode" json:"lastErrorCode"`
	CreatedAt         DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (AuthzedProjectionOutbox) TableName() string {
	return "AuthzedProjectionOutbox"
}

func (m *AuthzedProjectionOutbox) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.AvailableAt.IsZero() {
		m.AvailableAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AuthzedProjectionOutbox) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OnlineUpgradeState struct {
	Release              string             `gorm:"column:release;primaryKey" json:"release"`
	TargetImage          string             `gorm:"column:targetImage;not null" json:"targetImage"`
	Phase                string             `gorm:"column:phase;not null" json:"phase"`
	CompletedPhases      ScalarList[string] `gorm:"column:completedPhases;type:text[]" json:"completedPhases"`
	SurveyCursor         *string            `gorm:"column:surveyCursor" json:"surveyCursor"`
	SurveyCount          int64              `gorm:"column:surveyCount;default:0;not null" json:"surveyCount"`
	BaselineCompletedAt  *DateTime          `gorm:"column:baselineCompletedAt;type:timestamp(3)" json:"baselineCompletedAt"`
	ForwardOnlyAt        *DateTime          `gorm:"column:forwardOnlyAt;type:timestamp(3)" json:"forwardOnlyAt"`
	QueuePausedAt        *DateTime          `gorm:"column:queuePausedAt;type:timestamp(3)" json:"queuePausedAt"`
	QueuePauseDeadlineAt *DateTime          `gorm:"column:queuePauseDeadlineAt;type:timestamp(3)" json:"queuePauseDeadlineAt"`
	QueueResumedAt       *DateTime          `gorm:"column:queueResumedAt;type:timestamp(3)" json:"queueResumedAt"`
	ActivatedAt          *DateTime          `gorm:"column:activatedAt;type:timestamp(3)" json:"activatedAt"`
	UpdatedAt            DateTime           `gorm:"column:updatedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updatedAt"`
}

func (OnlineUpgradeState) TableName() string {
	return "OnlineUpgradeState"
}

func (m *OnlineUpgradeState) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*OnlineUpgradeState) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// One policy per organisation and kind of data: a notice `warnDays` before the action, and the action
// `periodDays` after the target's clock. The action follows from the kind of data: responses are
// deleted, surveys archived (the archive purge deletes them 30 days later, as it does every archived
// survey), members deactivated. The dates themselves are never stored; they are computed by the shared
// helper in `apps/web/modules/ee/data-retention/lib/schedule.ts`, so the UI, the emails and the sweep
// always agree.
//
// The allowed ranges are validated by the API. The migration repeats the ones that guard against
// acting early or re-notifying (a notice of at least 14 days and shorter than the period, an enabled
// policy always has `enabledAt`) and the structural ones (conditions on the surveys policy only) as
// CHECK constraints, which Prisma cannot express.
type RetentionPolicy struct {
	ID             string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt      DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt      DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	OrganizationID string   `gorm:"column:organizationId;uniqueIndex:RetentionPolicy_organizationId_entity_key,priority:1;not null" json:"organizationId"`
	Entity         string   `gorm:"column:entity;uniqueIndex:RetentionPolicy_organizationId_entity_key,priority:2;not null" json:"entity"`
	Enabled        bool     `gorm:"column:enabled;default:false;not null" json:"enabled"`
	// When the current configuration took effect: the policy was switched on or unpaused, or changed
	// in a way that could bring an action earlier (a shorter period or warning, different
	// conditions). Lengthening does not move it. The warning always runs in full from here, so no
	// change can delete data on the night it is made (ENG-3614). Also the clock for members with no
	// recorded sign-in.
	EnabledAt *DateTime `gorm:"column:enabledAt;type:timestamp(3)" json:"enabledAt"`
	// How long before the action the notice goes out. The action never runs before the notice has
	// been out for this long, so the warning always runs in full.
	WarnDays *int `gorm:"column:warnDays;size:32;default:60;not null" json:"warnDays"`
	// The period, counted from the target's clock, after which the policy acts: responses are
	// deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
	PeriodDays int `gorm:"column:periodDays;size:32;not null" json:"periodDays"`
	// Surveys policy only.
	Conditions ScalarList[string] `gorm:"column:conditions;type:text[]" json:"conditions"`
	// The last night the sweep's time ran out before it reached this policy's organisation, licensed.
	// The organisation was due and the sweep was running, so it counts as reached for the gap that
	// restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
	DeferredAt   *DateTime     `gorm:"column:deferredAt;type:timestamp(3)" json:"deferredAt"`
	UpdatedByID  *string       `gorm:"column:updatedById;index:RetentionPolicy_updatedById_idx" json:"updatedById"`
	Organization *Organization `gorm:"belongsTo"`
	UpdatedBy    *User         `gorm:"foreignKey:UpdatedByID;belongsTo"`
}

func (RetentionPolicy) TableName() string {
	return "RetentionPolicy"
}

func (m *RetentionPolicy) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*RetentionPolicy) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Holds one survey back from one policy until `until`. Keyed by organisation + entity rather than a
// policy id, so removing and re-adding a policy doesn't orphan its exemptions.
//
// Active means `revokedAt IS NULL AND until > now()`, so expiry needs no job. "One active exemption
// per survey and policy" is the partial unique index `RetentionExemption_surveyId_entity_active_key`
// (`WHERE "revokedAt" IS NULL`) in
// `migration/20261007120000_add_data_retention_tables/migration.sql`; Prisma's `@@unique` has no
// `where`, so it is deliberately not declared here (see "Indexes Prisma cannot express" in the
// README). An expired row still holds that slot until it is revoked, so creating a new exemption
// must close an expired one in the same transaction.
//
// `surveyId` is nullable and a CHECK requires exactly one target, so a later `workflowId` is one more
// column in that CHECK. The survey's organisation can't be tied to `organizationId` in the database
// (a survey carries its workspace, not its organisation): writers derive it from the survey.
type RetentionExemption struct {
	ID             string   `gorm:"column:id;primaryKey;index:RetentionExemption_organizationId_created_at_id_idx,priority:3" json:"id"`
	CreatedAt      DateTime `gorm:"column:created_at;index:RetentionExemption_organizationId_created_at_id_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	OrganizationID string   `gorm:"column:organizationId;index:RetentionExemption_organizationId_created_at_id_idx,priority:1;not null" json:"organizationId"`
	Entity         string   `gorm:"column:entity;not null" json:"entity"`
	SurveyID       *string  `gorm:"column:surveyId;index:RetentionExemption_surveyId_idx" json:"surveyId"`
	Until          DateTime `gorm:"column:until;type:timestamp(3);not null" json:"until"`
	Source         string   `gorm:"column:source;default:'manual';not null" json:"source"`
	// Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`).
	Reason       *string       `gorm:"column:reason" json:"reason"`
	CreatedByID  *string       `gorm:"column:createdById;index:RetentionExemption_createdById_idx" json:"createdById"`
	RevokedAt    *DateTime     `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	RevokedByID  *string       `gorm:"column:revokedById;index:RetentionExemption_revokedById_idx" json:"revokedById"`
	Organization *Organization `gorm:"belongsTo"`
	Survey       *Survey       `gorm:"belongsTo"`
	CreatedBy    *User         `gorm:"foreignKey:CreatedByID;belongsTo"`
	RevokedBy    *User         `gorm:"foreignKey:RevokedByID;belongsTo"`
}

func (RetentionExemption) TableName() string {
	return "RetentionExemption"
}

func (m *RetentionExemption) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
// while the run is in progress, or if it never finished.
type RetentionRun struct {
	ID             string    `gorm:"column:id;primaryKey;index:RetentionRun_organizationId_startedAt_id_idx,priority:3" json:"id"`
	OrganizationID string    `gorm:"column:organizationId;index:RetentionRun_organizationId_startedAt_id_idx,priority:1;not null" json:"organizationId"`
	Entity         string    `gorm:"column:entity;not null" json:"entity"`
	Kind           string    `gorm:"column:kind;default:'sweep';not null" json:"kind"`
	StartedAt      DateTime  `gorm:"column:startedAt;index:RetentionRun_organizationId_startedAt_id_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"startedAt"`
	FinishedAt     *DateTime `gorm:"column:finishedAt;type:timestamp(3)" json:"finishedAt"`
	NotifiedCount  int       `gorm:"column:notifiedCount;size:32;default:0;not null" json:"notifiedCount"`
	// Archived surveys, or deactivated members.
	ArchivedCount int `gorm:"column:archivedCount;size:32;default:0;not null" json:"archivedCount"`
	// Deleted responses, or, on a purge run, deleted surveys.
	DeletedCount int `gorm:"column:deletedCount;size:32;default:0;not null" json:"deletedCount"`
	SkippedCount int `gorm:"column:skippedCount;size:32;default:0;not null" json:"skippedCount"`
	// False when the run changed nothing. History hides those by default.
	HasChanges bool `gorm:"column:hasChanges;default:false;not null" json:"hasChanges"`
	// Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
	// of rescanning the same first targets every night. Null when the scan reached the end.
	ScanCursor   *string            `gorm:"column:scanCursor" json:"scanCursor"`
	Organization *Organization      `gorm:"belongsTo"`
	Items        []RetentionRunItem `gorm:"foreignKey:RunID;constraint:RetentionRunItem_runId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RetentionRun) TableName() string {
	return "RetentionRun"
}

func (m *RetentionRun) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.StartedAt.IsZero() {
		m.StartedAt = DateTime{Time: now}
	}
	return nil
}

// What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
// never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
// `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
type RetentionRunItem struct {
	ID         string  `gorm:"column:id;primaryKey;index:RetentionRunItem_runId_id_idx,priority:2" json:"id"`
	RunID      string  `gorm:"column:runId;index:RetentionRunItem_runId_id_idx,priority:1;not null" json:"runId"`
	TargetType string  `gorm:"column:targetType;not null" json:"targetType"`
	TargetID   string  `gorm:"column:targetId;index:RetentionRunItem_targetId_idx;not null" json:"targetId"`
	TargetName *string `gorm:"column:targetName" json:"targetName"`
	Action     string  `gorm:"column:action;not null" json:"action"`
	// 1, except for grouped response deletions, where it is the number of responses deleted.
	Count *int `gorm:"column:count;size:32;default:1;not null" json:"count"`
	// Who was emailed, for `notified` rows.
	Recipient  *string       `gorm:"column:recipient" json:"recipient"`
	SkipReason *string       `gorm:"column:skipReason" json:"skipReason"`
	Run        *RetentionRun `gorm:"foreignKey:RunID;belongsTo"`
}

func (RetentionRunItem) TableName() string {
	return "RetentionRunItem"
}

func (m *RetentionRunItem) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	return nil
}

// The notice for one target under one policy, so each notice goes out once: a survey or a member, or
// the one-time responses reminder for a survey. It counts only once delivered, and is void if it was
// claimed before the policy's `enabledAt`, before the target's latest exemption ended or (surveys and
// members) before the target's clock: then a later cycle claims it again and sends a new one.
//
// A CHECK in the migration ties each entity to its target: `surveys` and `responses` → `surveyId`,
// `members` → `userId`; another makes a sent email imply a delivery time. The survey notices that still
// hold a clock are the partial index `RetentionNotice_surveys_clock_pending_idx` (`WHERE "entity" =
// 'surveys' AND "clockAt" IS NOT NULL`) in
// `migration/20261012120000_retention_notice_void_on_response_delete/migration.sql`, with the triggers
// that clear `clockAt`; Prisma's `@@index` has no `where`, so it is deliberately not declared here (see
// "Indexes Prisma cannot express" in the README).
type RetentionNotice struct {
	ID             string  `gorm:"column:id;primaryKey" json:"id"`
	OrganizationID string  `gorm:"column:organizationId;uniqueIndex:RetentionNotice_userId_organizationId_entity_key,priority:2;index:RetentionNotice_organizationId_idx;not null" json:"organizationId"`
	Entity         string  `gorm:"column:entity;uniqueIndex:RetentionNotice_surveyId_entity_key,priority:2;uniqueIndex:RetentionNotice_userId_organizationId_entity_key,priority:3;not null" json:"entity"`
	SurveyID       *string `gorm:"column:surveyId;uniqueIndex:RetentionNotice_surveyId_entity_key,priority:1" json:"surveyId"`
	UserID         *string `gorm:"column:userId;uniqueIndex:RetentionNotice_userId_organizationId_entity_key,priority:1" json:"userId"`
	// When the notice was claimed for sending. A claim that never got delivered can be taken over.
	SentAt DateTime `gorm:"column:sentAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"sentAt"`
	// When the notice was handed to the mail transport, or recorded without one (no SMTP). The
	// warning only counts from a delivered notice, so a crash between claim and send can't let the
	// action run without it.
	DeliveredAt *DateTime `gorm:"column:deliveredAt;type:timestamp(3)" json:"deliveredAt"`
	// Whether an email actually went out. False when SMTP isn't configured.
	EmailSent bool `gorm:"column:emailSent;default:false;not null" json:"emailSent"`
	// Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
	// can't mark a newer claim delivered.
	ClaimToken *string `gorm:"column:claimToken" json:"claimToken"`
	// The target's clock the notice was computed for (surveys and members). A notice counts only while
	// the target's clock is still this one, so activity between the sweep's read and its claim voids
	// it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
	// for a survey notice voided for good: a person deleted or moved one of its responses, or the
	// responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
	ClockAt      *DateTime     `gorm:"column:clockAt;type:timestamp(3)" json:"clockAt"`
	Organization *Organization `gorm:"belongsTo"`
	Survey       *Survey       `gorm:"belongsTo"`
	User         *User         `gorm:"belongsTo"`
}

func (RetentionNotice) TableName() string {
	return "RetentionNotice"
}

func (m *RetentionNotice) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.SentAt.IsZero() {
		m.SentAt = DateTime{Time: now}
	}
	return nil
}

// When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
// and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
// so retention adds nothing to the survey row.
type RetentionSurveyRestore struct {
	SurveyID   string   `gorm:"column:surveyId;primaryKey" json:"surveyId"`
	RestoredAt DateTime `gorm:"column:restoredAt;type:timestamp(3);not null" json:"restoredAt"`
	Survey     *Survey  `gorm:"belongsTo"`
}

func (RetentionSurveyRestore) TableName() string {
	return "RetentionSurveyRestore"
}

type Workflow struct {
	ID          string   `gorm:"column:id;primaryKey;uniqueIndex:Workflow_id_workspaceId_key,priority:1" json:"id"`
	CreatedAt   DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime `gorm:"column:updated_at;index:Workflow_workspaceId_updated_at_idx,priority:2;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Name        string   `gorm:"column:name;not null" json:"name"`
	Description *string  `gorm:"column:description" json:"description"`
	Status      string   `gorm:"column:status;index:Workflow_workspaceId_status_idx,priority:2;default:'draft';not null" json:"status"`
	WorkspaceID string   `gorm:"column:workspaceId;uniqueIndex:Workflow_id_workspaceId_key,priority:2;index:Workflow_workspaceId_status_idx,priority:1;index:Workflow_workspaceId_updated_at_idx,priority:1;not null" json:"workspaceId"`
	CreatedBy   *string  `gorm:"column:createdBy" json:"createdBy"`
	// [WorkflowDefinition]
	Definition datatypes.JSON `gorm:"column:definition;not null" json:"definition"`
	// Webhook Secret URL token (ADRs 0048, 0057): SHA-256 hash to check calls, and an encrypted copy so
	// the editor can show the URL again. Never part of the definition, its versions or its copies.
	WebhookTokenHash      *string           `gorm:"column:webhookTokenHash" json:"webhookTokenHash"`
	WebhookTokenEncrypted *string           `gorm:"column:webhookTokenEncrypted" json:"webhookTokenEncrypted"`
	Workspace             *Workspace        `gorm:"belongsTo"`
	Creator               *User             `gorm:"foreignKey:CreatedBy;belongsTo"`
	Versions              []WorkflowVersion `gorm:"foreignKey:WorkflowID,WorkspaceID;references:ID,WorkspaceID;constraint:WorkflowVersion_workflowId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Runs                  []WorkflowRun     `gorm:"foreignKey:WorkflowID,WorkspaceID;references:ID,WorkspaceID;constraint:WorkflowRun_workflowId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Secrets               []WorkflowSecret  `gorm:"foreignKey:WorkflowID,WorkspaceID;references:ID,WorkspaceID;constraint:WorkflowSecret_workflowId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Samples               []WorkflowSample  `gorm:"foreignKey:WorkflowID,WorkspaceID;references:ID,WorkspaceID;constraint:WorkflowSample_workflowId_workspaceId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Workflow) TableName() string {
	return "Workflow"
}

func (m *Workflow) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Workflow) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkflowVersion struct {
	ID          string `gorm:"column:id;primaryKey" json:"id"`
	WorkflowID  string `gorm:"column:workflowId;uniqueIndex:WorkflowVersion_workflowId_version_key,priority:1;index:WorkflowVersion_workflowId_workspaceId_idx,priority:1;not null" json:"workflowId"`
	WorkspaceID string `gorm:"column:workspaceId;index:WorkflowVersion_workflowId_workspaceId_idx,priority:2;not null" json:"workspaceId"`
	Version     int    `gorm:"column:version;uniqueIndex:WorkflowVersion_workflowId_version_key,priority:2;size:32;not null" json:"version"`
	// [WorkflowExecutableDefinition]
	Definition  datatypes.JSON `gorm:"column:definition;not null" json:"definition"`
	PublishedAt DateTime       `gorm:"column:publishedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"publishedAt"`
	PublishedBy *string        `gorm:"column:publishedBy;index:WorkflowVersion_publishedBy_idx" json:"publishedBy"`
	Workflow    *Workflow      `gorm:"foreignKey:WorkflowID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
	Publisher   *User          `gorm:"foreignKey:PublishedBy;belongsTo"`
	Runs        []WorkflowRun  `gorm:"foreignKey:WorkflowVersionID;constraint:WorkflowRun_workflowVersionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (WorkflowVersion) TableName() string {
	return "WorkflowVersion"
}

func (m *WorkflowVersion) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.PublishedAt.IsZero() {
		m.PublishedAt = DateTime{Time: now}
	}
	return nil
}

type WorkflowRun struct {
	ID                string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt         DateTime  `gorm:"column:created_at;index:WorkflowRun_workspaceId_created_at_idx,priority:2;index:WorkflowRun_workflowId_workspaceId_created_at_idx,priority:3;index:WorkflowRun_status_created_at_idx,priority:2;index:WorkflowRun_isDryRun_created_at_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt         DateTime  `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	WorkflowID        string    `gorm:"column:workflowId;uniqueIndex:WorkflowRun_workflowId_idempotencyKey_key,priority:1;index:WorkflowRun_workflowId_workspaceId_created_at_idx,priority:1;not null" json:"workflowId"`
	WorkspaceID       string    `gorm:"column:workspaceId;index:WorkflowRun_workspaceId_created_at_idx,priority:1;index:WorkflowRun_workflowId_workspaceId_created_at_idx,priority:2;not null" json:"workspaceId"`
	WorkflowVersionID *string   `gorm:"column:workflowVersionId" json:"workflowVersionId"`
	ResponseID        *string   `gorm:"column:responseId;index:WorkflowRun_responseId_idx" json:"responseId"`
	Status            string    `gorm:"column:status;index:WorkflowRun_status_created_at_idx,priority:1;default:'queued';not null" json:"status"`
	TriggerType       string    `gorm:"column:triggerType;not null" json:"triggerType"`
	SurveyID          *string   `gorm:"column:surveyId" json:"surveyId"`
	IsDryRun          bool      `gorm:"column:isDryRun;index:WorkflowRun_isDryRun_idx;index:WorkflowRun_isDryRun_created_at_idx,priority:1;default:false;not null" json:"isDryRun"`
	IdempotencyKey    *string   `gorm:"column:idempotencyKey;uniqueIndex:WorkflowRun_workflowId_idempotencyKey_key,priority:2" json:"idempotencyKey"`
	Attempt           int       `gorm:"column:attempt;size:32;default:0;not null" json:"attempt"`
	NextAttemptAt     *DateTime `gorm:"column:nextAttemptAt;index:WorkflowRun_nextAttemptAt_idx;type:timestamp(3)" json:"nextAttemptAt"`
	LastErrorAt       *DateTime `gorm:"column:lastErrorAt;type:timestamp(3)" json:"lastErrorAt"`
	// [WorkflowTriggerRunPayload]
	TriggerPayload datatypes.JSON `gorm:"column:triggerPayload;not null" json:"triggerPayload"`
	// [WorkflowRunData]
	Data         datatypes.JSON `gorm:"column:data;default:'{}';not null" json:"data"`
	Error        *string        `gorm:"column:error" json:"error"`
	DispatchedAt *DateTime      `gorm:"column:dispatchedAt;type:timestamp(3)" json:"dispatchedAt"`
	StartedAt    *DateTime      `gorm:"column:startedAt;type:timestamp(3)" json:"startedAt"`
	FinishedAt   *DateTime      `gorm:"column:finishedAt;index:WorkflowRun_contentDeletedAt_finishedAt_idx,priority:2;type:timestamp(3)" json:"finishedAt"`
	// The caller's IP for a run started by a Webhook Secret URL call; shown in run detail.
	CallerIP *string `gorm:"column:callerIp" json:"callerIp"`
	// When the run retention job cleared the run's content (ADR 0044): payload contents, run data, step
	// input, output and errors, the run error and the caller IP. Null while the content is kept.
	ContentDeletedAt *DateTime `gorm:"column:contentDeletedAt;index:WorkflowRun_contentDeletedAt_finishedAt_idx,priority:1;type:timestamp(3)" json:"contentDeletedAt"`
	// The retention period in days that applied when the content was cleared ("Details deleted after N days").
	ContentRetentionDays *int `gorm:"column:contentRetentionDays;size:32" json:"contentRetentionDays"`
	// The draft definition a Test run executed (ADR 0049), so its graph can be drawn after the draft
	// changes. Null for live runs (they point at `workflowVersionId`) and for Test runs saved before it
	// existed. Configuration like a version's definition, not run content: retention leaves it alone.
	// [WorkflowExecutableDefinition]
	TestRunDefinition datatypes.JSON   `gorm:"column:testRunDefinition" json:"testRunDefinition"`
	Workflow          *Workflow        `gorm:"foreignKey:WorkflowID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
	Workspace         *Workspace       `gorm:"belongsTo"`
	WorkflowVersion   *WorkflowVersion `gorm:"belongsTo"`
	Response          *Response        `gorm:"belongsTo"`
	Logs              []WorkflowRunLog `gorm:"foreignKey:RunID;constraint:WorkflowRunLog_runId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (WorkflowRun) TableName() string {
	return "WorkflowRun"
}

func (m *WorkflowRun) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkflowRun) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Per-call data kept beside a workflow (ADR 0048): one row per source, `trigger` for the Webhook
// trigger's last payload (plus the last call's time and IP), or an HTTP request step id for its last
// Test request response. Written on every call, so it lives here and never bumps `Workflow.updatedAt`.
type WorkflowSample struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime `gorm:"column:updated_at;index:WorkflowSample_updated_at_idx;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	WorkflowID  string   `gorm:"column:workflowId;uniqueIndex:WorkflowSample_workflowId_source_key,priority:1;index:WorkflowSample_workflowId_workspaceId_idx,priority:1;not null" json:"workflowId"`
	WorkspaceID string   `gorm:"column:workspaceId;index:WorkflowSample_workflowId_workspaceId_idx,priority:2;not null" json:"workspaceId"`
	Source      string   `gorm:"column:source;uniqueIndex:WorkflowSample_workflowId_source_key,priority:2;not null" json:"source"`
	// [WorkflowSampleData]
	Sample       datatypes.JSON `gorm:"column:sample;not null" json:"sample"`
	LastCalledAt *DateTime      `gorm:"column:lastCalledAt;type:timestamp(3)" json:"lastCalledAt"`
	LastCallerIP *string        `gorm:"column:lastCallerIp" json:"lastCallerIp"`
	Workflow     *Workflow      `gorm:"foreignKey:WorkflowID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
}

func (WorkflowSample) TableName() string {
	return "WorkflowSample"
}

func (m *WorkflowSample) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkflowSample) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkflowRunLog struct {
	ID       string `gorm:"column:id;primaryKey" json:"id"`
	RunID    string `gorm:"column:runId;uniqueIndex:WorkflowRunLog_runId_stepId_key,priority:1;index:WorkflowRunLog_runId_sequence_idx,priority:1;not null" json:"runId"`
	Sequence int    `gorm:"column:sequence;index:WorkflowRunLog_runId_sequence_idx,priority:2;size:32;not null" json:"sequence"`
	StepID   string `gorm:"column:stepId;uniqueIndex:WorkflowRunLog_runId_stepId_key,priority:2;not null" json:"stepId"`
	StepType string `gorm:"column:stepType;not null" json:"stepType"`
	Status   string `gorm:"column:status;index:WorkflowRunLog_status_startedAt_idx,priority:1;default:'pending';not null" json:"status"`
	// [WorkflowRunLogInput]
	Input datatypes.JSON `gorm:"column:input;default:'{}';not null" json:"input"`
	// [WorkflowRunLogOutput]
	Output     datatypes.JSON `gorm:"column:output;default:'{}';not null" json:"output"`
	Error      *string        `gorm:"column:error" json:"error"`
	StartedAt  *DateTime      `gorm:"column:startedAt;index:WorkflowRunLog_status_startedAt_idx,priority:2;type:timestamp(3)" json:"startedAt"`
	FinishedAt *DateTime      `gorm:"column:finishedAt;type:timestamp(3)" json:"finishedAt"`
	Run        *WorkflowRun   `gorm:"foreignKey:RunID;belongsTo"`
}

func (WorkflowRunLog) TableName() string {
	return "WorkflowRunLog"
}

func (m *WorkflowRunLog) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	return nil
}

// Encrypted value of an HTTP request secret header (ENG-7168, ADR 0043). The definition and every version
// store only `{name, secret: true, secretId}`; Replace overwrites this row in place, so the enabled
// version uses the new value at once. Rows nothing references are deleted on Enable; all are deleted
// with the workflow. A value is only sent to the origin it was set for (ADR 0051).
type WorkflowSecret struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt   DateTime `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	WorkflowID  string   `gorm:"column:workflowId;index:WorkflowSecret_workflowId_workspaceId_idx,priority:1;not null" json:"workflowId"`
	WorkspaceID string   `gorm:"column:workspaceId;index:WorkflowSecret_workflowId_workspaceId_idx,priority:2;not null" json:"workspaceId"`
	// `symmetricEncrypt` ciphertext (AES-256-GCM). Never logged, never returned.
	Value string `gorm:"column:value;not null" json:"value"`
	// Origin (scheme, host, port) of the step's URL when the value was set or replaced; the value is only
	// sent there (ADR 0051). Null for a value set before host binding: held to the step's saved URL
	// until it is set again.
	BoundOrigin *string   `gorm:"column:boundOrigin" json:"boundOrigin"`
	Workflow    *Workflow `gorm:"foreignKey:WorkflowID,WorkspaceID;references:ID,WorkspaceID;belongsTo"`
}

func (WorkflowSecret) TableName() string {
	return "WorkflowSecret"
}

func (m *WorkflowSecret) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid2.Generate()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*WorkflowSecret) BeforeUpdate(tx *gorm.DB) error {
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
