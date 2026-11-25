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
)

type DocumentAnnotation struct {
	ID          string            `gorm:"column:id;primaryKey" json:"id"`
	Title       string            `gorm:"column:title;not null" json:"title"`
	Content     datatypes.JSON    `gorm:"column:content;not null" json:"content"`
	Pages       ScalarList[int]   `gorm:"column:pages;type:integer[]" json:"pages"`
	DocumentID  string            `gorm:"column:documentId;index:DocumentAnnotation_documentId_idx;not null" json:"documentId"`
	TeamID      string            `gorm:"column:teamId;index:DocumentAnnotation_teamId_idx;not null" json:"teamId"`
	CreatedByID string            `gorm:"column:createdById;index:DocumentAnnotation_createdById_idx;not null" json:"createdById"`
	IsVisible   *bool             `gorm:"column:isVisible;default:true;not null" json:"isVisible"`
	CreatedAt   DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Document    *Document         `gorm:"belongsTo"`
	Team        *Team             `gorm:"belongsTo"`
	CreatedBy   *User             `gorm:"foreignKey:CreatedByID;belongsTo"`
	Images      []AnnotationImage `gorm:"foreignKey:AnnotationID;constraint:AnnotationImage_annotationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (DocumentAnnotation) TableName() string {
	return "DocumentAnnotation"
}

func (m *DocumentAnnotation) BeforeCreate(tx *gorm.DB) error {
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

func (*DocumentAnnotation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AnnotationImage struct {
	ID           string              `gorm:"column:id;primaryKey" json:"id"`
	Filename     string              `gorm:"column:filename;not null" json:"filename"`
	URL          string              `gorm:"column:url;not null" json:"url"`
	Size         *int                `gorm:"column:size;size:32" json:"size"`
	MimeType     *string             `gorm:"column:mimeType" json:"mimeType"`
	AnnotationID string              `gorm:"column:annotationId;index:AnnotationImage_annotationId_idx;not null" json:"annotationId"`
	CreatedAt    DateTime            `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Annotation   *DocumentAnnotation `gorm:"foreignKey:AnnotationID;belongsTo"`
}

func (AnnotationImage) TableName() string {
	return "AnnotationImage"
}

func (m *AnnotationImage) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Brand struct {
	ID                             string     `gorm:"column:id;primaryKey" json:"id"`
	Logo                           *string    `gorm:"column:logo" json:"logo"`
	HideLogo                       bool       `gorm:"column:hideLogo;default:false;not null" json:"hideLogo"`
	Banner                         *string    `gorm:"column:banner" json:"banner"`
	BrandColor                     *string    `gorm:"column:brandColor" json:"brandColor"`
	AccentColor                    *string    `gorm:"column:accentColor" json:"accentColor"`
	AccentButtonColor              *string    `gorm:"column:accentButtonColor" json:"accentButtonColor"`
	ApplyAccentColorToDataroomView bool       `gorm:"column:applyAccentColorToDataroomView;default:false;not null" json:"applyAccentColorToDataroomView"`
	WelcomeMessage                 *string    `gorm:"column:welcomeMessage" json:"welcomeMessage"`
	CtaLabel                       *string    `gorm:"column:ctaLabel" json:"ctaLabel"`
	CtaURL                         *string    `gorm:"column:ctaUrl" json:"ctaUrl"`
	PrivacyPolicyURL               *string    `gorm:"column:privacyPolicyUrl" json:"privacyPolicyUrl"`
	CardLayout                     *string    `gorm:"column:cardLayout;default:'LIST';not null" json:"cardLayout"`
	ShowFolderTree                 *bool      `gorm:"column:showFolderTree;default:true;not null" json:"showFolderTree"`
	ViewerLayoutPreset             *string    `gorm:"column:viewerLayoutPreset;default:'STANDARD';not null" json:"viewerLayoutPreset"`
	ViewerHeaderStyle              *string    `gorm:"column:viewerHeaderStyle;default:'DEFAULT';not null" json:"viewerHeaderStyle"`
	HideFolderIconsInMain          bool       `gorm:"column:hideFolderIconsInMain;default:false;not null" json:"hideFolderIconsInMain"`
	CustomLinkPreviewEnabled       bool       `gorm:"column:customLinkPreviewEnabled;default:false;not null" json:"customLinkPreviewEnabled"`
	LinkPreviewTitle               *string    `gorm:"column:linkPreviewTitle" json:"linkPreviewTitle"`
	LinkPreviewDescription         *string    `gorm:"column:linkPreviewDescription" json:"linkPreviewDescription"`
	LinkPreviewImage               *string    `gorm:"column:linkPreviewImage" json:"linkPreviewImage"`
	LinkPreviewFavicon             *string    `gorm:"column:linkPreviewFavicon" json:"linkPreviewFavicon"`
	DefaultLanguage                *string    `gorm:"column:defaultLanguage;default:'en';not null" json:"defaultLanguage"`
	Name                           *string    `gorm:"column:name;uniqueIndex:Brand_teamId_name_key,priority:2;default:'Default';not null" json:"name"`
	TeamID                         string     `gorm:"column:teamId;uniqueIndex:Brand_teamId_name_key,priority:1;index:Brand_teamId_idx;not null" json:"teamId"`
	CreatedAt                      DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                      DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team                           *Team      `gorm:"belongsTo"`
	DefaultForTeam                 []Team     `gorm:"foreignKey:DefaultBrandID;constraint:Team_defaultBrandId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Links                          []Link     `gorm:"foreignKey:BrandID;constraint:Link_brandId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Datarooms                      []Dataroom `gorm:"foreignKey:BrandID;constraint:Dataroom_brandId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Brand) TableName() string {
	return "Brand"
}

func (m *Brand) BeforeCreate(tx *gorm.DB) error {
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

func (*Brand) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Conversation struct {
	ID                    string                    `gorm:"column:id;primaryKey" json:"id"`
	Title                 *string                   `gorm:"column:title" json:"title"`
	IsEnabled             *bool                     `gorm:"column:isEnabled;default:true;not null" json:"isEnabled"`
	VisibilityMode        string                    `gorm:"column:visibilityMode;default:'PRIVATE';not null" json:"visibilityMode"`
	DataroomID            string                    `gorm:"column:dataroomId;index:Conversation_dataroomId_idx;not null" json:"dataroomId"`
	DataroomDocumentID    *string                   `gorm:"column:dataroomDocumentId;index:Conversation_dataroomDocumentId_idx" json:"dataroomDocumentId"`
	DocumentVersionNumber *int                      `gorm:"column:documentVersionNumber;size:32" json:"documentVersionNumber"`
	DocumentPageNumber    *int                      `gorm:"column:documentPageNumber;size:32" json:"documentPageNumber"`
	LinkID                *string                   `gorm:"column:linkId;index:Conversation_linkId_idx" json:"linkId"`
	ViewerGroupID         *string                   `gorm:"column:viewerGroupId;index:Conversation_viewerGroupId_idx" json:"viewerGroupId"`
	QuestionID            *string                   `gorm:"column:questionId;uniqueIndex:Conversation_questionId_recipientViewerId_key,priority:1;index:Conversation_questionId_idx" json:"questionId"`
	Origin                *string                   `gorm:"column:origin;default:'visitor';not null" json:"origin"`
	Category              *string                   `gorm:"column:category" json:"category"`
	Priority              *string                   `gorm:"column:priority" json:"priority"`
	RecipientViewerID     *string                   `gorm:"column:recipientViewerId;uniqueIndex:Conversation_questionId_recipientViewerId_key,priority:2;index:Conversation_recipientViewerId_idx" json:"recipientViewerId"`
	InitialViewID         *string                   `gorm:"column:initialViewId;index:Conversation_initialViewId_idx" json:"initialViewId"`
	TeamID                string                    `gorm:"column:teamId;index:Conversation_teamId_idx;not null" json:"teamId"`
	LastMessageAt         *DateTime                 `gorm:"column:lastMessageAt;type:timestamp(3)" json:"lastMessageAt"`
	CreatedAt             DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime                  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom              *Dataroom                 `gorm:"belongsTo"`
	DataroomDocument      *DataroomDocument         `gorm:"belongsTo"`
	Link                  *Link                     `gorm:"belongsTo"`
	ViewerGroup           *ViewerGroup              `gorm:"belongsTo"`
	Question              *DataroomQuestion         `gorm:"foreignKey:QuestionID;belongsTo"`
	RecipientViewer       *Viewer                   `gorm:"foreignKey:RecipientViewerID;belongsTo"`
	InitialView           *View                     `gorm:"foreignKey:InitialViewID;belongsTo"`
	Team                  *Team                     `gorm:"belongsTo"`
	Views                 []ConversationView        `gorm:"foreignKey:ConversationID;constraint:ConversationView_conversationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Participants          []ConversationParticipant `gorm:"foreignKey:ConversationID;constraint:ConversationParticipant_conversationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Messages              []Message                 `gorm:"foreignKey:ConversationID;constraint:Message_conversationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FaqItems              []DataroomFaqItem         `gorm:"foreignKey:SourceConversationID;constraint:DataroomFaqItem_sourceConversationId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Conversation) TableName() string {
	return "Conversation"
}

func (m *Conversation) BeforeCreate(tx *gorm.DB) error {
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

func (*Conversation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DataroomQuestion struct {
	ID                 string                       `gorm:"column:id;primaryKey" json:"id"`
	Content            string                       `gorm:"column:content;type:text;not null" json:"content"`
	Category           *string                      `gorm:"column:category" json:"category"`
	Priority           *string                      `gorm:"column:priority" json:"priority"`
	Status             *string                      `gorm:"column:status;index:DataroomQuestion_status_idx;default:'unanswered';not null" json:"status"`
	OrderIndex         *int                         `gorm:"column:orderIndex;size:32" json:"orderIndex"`
	ImportBatchID      *string                      `gorm:"column:importBatchId;index:DataroomQuestion_importBatchId_idx" json:"importBatchId"`
	DataroomID         string                       `gorm:"column:dataroomId;index:DataroomQuestion_dataroomId_idx;not null" json:"dataroomId"`
	TeamID             string                       `gorm:"column:teamId;index:DataroomQuestion_teamId_idx;not null" json:"teamId"`
	CreatedByUserID    *string                      `gorm:"column:createdByUserId;index:DataroomQuestion_createdByUserId_idx" json:"createdByUserId"`
	DataroomDocumentID *string                      `gorm:"column:dataroomDocumentId;index:DataroomQuestion_dataroomDocumentId_idx" json:"dataroomDocumentId"`
	DocumentPageNumber *int                         `gorm:"column:documentPageNumber;size:32" json:"documentPageNumber"`
	CreatedAt          DateTime                     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime                     `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom           *Dataroom                    `gorm:"belongsTo"`
	Team               *Team                        `gorm:"belongsTo"`
	CreatedByUser      *User                        `gorm:"foreignKey:CreatedByUserID;belongsTo"`
	DataroomDocument   *DataroomDocument            `gorm:"belongsTo"`
	Assignments        []DataroomQuestionAssignment `gorm:"foreignKey:QuestionID;constraint:DataroomQuestionAssignment_questionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Conversations      []Conversation               `gorm:"foreignKey:QuestionID;constraint:Conversation_questionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (DataroomQuestion) TableName() string {
	return "DataroomQuestion"
}

func (m *DataroomQuestion) BeforeCreate(tx *gorm.DB) error {
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

func (*DataroomQuestion) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DataroomQuestionAssignment struct {
	ID         string            `gorm:"column:id;primaryKey" json:"id"`
	QuestionID string            `gorm:"column:questionId;uniqueIndex:DataroomQuestionAssignment_questionId_linkId_key,priority:1;uniqueIndex:DataroomQuestionAssignment_questionId_groupId_key,priority:1;uniqueIndex:DataroomQuestionAssignment_questionId_viewerId_key,priority:1;uniqueIndex:DataroomQuestionAssignment_questionId_email_key,priority:1;index:DataroomQuestionAssignment_questionId_idx;not null" json:"questionId"`
	LinkID     *string           `gorm:"column:linkId;uniqueIndex:DataroomQuestionAssignment_questionId_linkId_key,priority:2;index:DataroomQuestionAssignment_linkId_idx" json:"linkId"`
	GroupID    *string           `gorm:"column:groupId;uniqueIndex:DataroomQuestionAssignment_questionId_groupId_key,priority:2;index:DataroomQuestionAssignment_groupId_idx" json:"groupId"`
	ViewerID   *string           `gorm:"column:viewerId;uniqueIndex:DataroomQuestionAssignment_questionId_viewerId_key,priority:2;index:DataroomQuestionAssignment_viewerId_idx" json:"viewerId"`
	Email      *string           `gorm:"column:email;uniqueIndex:DataroomQuestionAssignment_questionId_email_key,priority:2" json:"email"`
	CreatedAt  DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Question   *DataroomQuestion `gorm:"foreignKey:QuestionID;belongsTo"`
	Link       *Link             `gorm:"belongsTo"`
	Group      *ViewerGroup      `gorm:"foreignKey:GroupID;belongsTo"`
	Viewer     *Viewer           `gorm:"belongsTo"`
}

func (DataroomQuestionAssignment) TableName() string {
	return "DataroomQuestionAssignment"
}

func (m *DataroomQuestionAssignment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type ConversationParticipant struct {
	ID                   string        `gorm:"column:id;primaryKey" json:"id"`
	ConversationID       string        `gorm:"column:conversationId;uniqueIndex:ConversationParticipant_conversationId_viewerId_key,priority:1;uniqueIndex:ConversationParticipant_conversationId_userId_key,priority:1;index:ConversationParticipant_conversationId_idx;not null" json:"conversationId"`
	Role                 string        `gorm:"column:role;default:'PARTICIPANT';not null" json:"role"`
	ViewerID             *string       `gorm:"column:viewerId;uniqueIndex:ConversationParticipant_conversationId_viewerId_key,priority:2;index:ConversationParticipant_viewerId_idx" json:"viewerId"`
	UserID               *string       `gorm:"column:userId;uniqueIndex:ConversationParticipant_conversationId_userId_key,priority:2;index:ConversationParticipant_userId_idx" json:"userId"`
	ReceiveNotifications bool          `gorm:"column:receiveNotifications;default:false;not null" json:"receiveNotifications"`
	CreatedAt            DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Conversation         *Conversation `gorm:"belongsTo"`
	Viewer               *Viewer       `gorm:"belongsTo"`
	User                 *User         `gorm:"belongsTo"`
}

func (ConversationParticipant) TableName() string {
	return "ConversationParticipant"
}

func (m *ConversationParticipant) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Message struct {
	ID             string            `gorm:"column:id;primaryKey" json:"id"`
	Content        string            `gorm:"column:content;not null" json:"content"`
	ConversationID string            `gorm:"column:conversationId;index:Message_conversationId_idx;not null" json:"conversationId"`
	UserID         *string           `gorm:"column:userId;index:Message_userId_idx" json:"userId"`
	ViewerID       *string           `gorm:"column:viewerId;index:Message_viewerId_idx" json:"viewerId"`
	ViewID         *string           `gorm:"column:viewId;index:Message_viewId_idx" json:"viewId"`
	IsRead         bool              `gorm:"column:isRead;default:false;not null" json:"isRead"`
	IsInternalNote bool              `gorm:"column:isInternalNote;default:false;not null" json:"isInternalNote"`
	CreatedAt      DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Conversation   *Conversation     `gorm:"belongsTo"`
	User           *User             `gorm:"belongsTo"`
	Viewer         *Viewer           `gorm:"belongsTo"`
	View           *View             `gorm:"belongsTo"`
	FaqAsQuestion  []DataroomFaqItem `gorm:"foreignKey:QuestionMessageID;constraint:DataroomFaqItem_questionMessageId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	FaqAsAnswer    []DataroomFaqItem `gorm:"foreignKey:AnswerMessageID;constraint:DataroomFaqItem_answerMessageId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Message) TableName() string {
	return "Message"
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
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

func (*Message) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ConversationView struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	ConversationID string        `gorm:"column:conversationId;uniqueIndex:ConversationView_conversationId_viewId_key,priority:1;index:ConversationView_conversationId_idx;not null" json:"conversationId"`
	ViewID         string        `gorm:"column:viewId;uniqueIndex:ConversationView_conversationId_viewId_key,priority:2;index:ConversationView_viewId_idx;not null" json:"viewId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Conversation   *Conversation `gorm:"belongsTo"`
	View           *View         `gorm:"belongsTo"`
}

func (ConversationView) TableName() string {
	return "ConversationView"
}

func (m *ConversationView) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DataroomFaqItem struct {
	ID                    string             `gorm:"column:id;primaryKey" json:"id"`
	Title                 *string            `gorm:"column:title" json:"title"`
	EditedQuestion        string             `gorm:"column:editedQuestion;not null" json:"editedQuestion"`
	OriginalQuestion      *string            `gorm:"column:originalQuestion" json:"originalQuestion"`
	Answer                string             `gorm:"column:answer;not null" json:"answer"`
	Description           *string            `gorm:"column:description" json:"description"`
	DataroomID            string             `gorm:"column:dataroomId;index:DataroomFaqItem_dataroomId_idx;not null" json:"dataroomId"`
	LinkID                *string            `gorm:"column:linkId;index:DataroomFaqItem_linkId_idx" json:"linkId"`
	DataroomDocumentID    *string            `gorm:"column:dataroomDocumentId;index:DataroomFaqItem_dataroomDocumentId_idx" json:"dataroomDocumentId"`
	SourceConversationID  *string            `gorm:"column:sourceConversationId;index:DataroomFaqItem_sourceConversationId_idx" json:"sourceConversationId"`
	QuestionMessageID     *string            `gorm:"column:questionMessageId" json:"questionMessageId"`
	AnswerMessageID       *string            `gorm:"column:answerMessageId" json:"answerMessageId"`
	TeamID                string             `gorm:"column:teamId;index:DataroomFaqItem_teamId_idx;not null" json:"teamId"`
	PublishedByUserID     string             `gorm:"column:publishedByUserId;index:DataroomFaqItem_publishedByUserId_idx;not null" json:"publishedByUserId"`
	VisibilityMode        string             `gorm:"column:visibilityMode;index:DataroomFaqItem_visibilityMode_idx;default:'PUBLIC_DATAROOM';not null" json:"visibilityMode"`
	Status                string             `gorm:"column:status;index:DataroomFaqItem_status_idx;default:'PUBLISHED';not null" json:"status"`
	IsAnonymized          *bool              `gorm:"column:isAnonymized;default:true;not null" json:"isAnonymized"`
	ViewCount             int                `gorm:"column:viewCount;size:32;default:0;not null" json:"viewCount"`
	Tags                  ScalarList[string] `gorm:"column:tags;type:text[]" json:"tags"`
	DocumentPageNumber    *int               `gorm:"column:documentPageNumber;size:32" json:"documentPageNumber"`
	DocumentVersionNumber *int               `gorm:"column:documentVersionNumber;size:32" json:"documentVersionNumber"`
	CreatedAt             DateTime           `gorm:"column:createdAt;index:DataroomFaqItem_createdAt_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom              *Dataroom          `gorm:"belongsTo"`
	Link                  *Link              `gorm:"belongsTo"`
	DataroomDocument      *DataroomDocument  `gorm:"belongsTo"`
	SourceConversation    *Conversation      `gorm:"foreignKey:SourceConversationID;belongsTo"`
	QuestionMessage       *Message           `gorm:"foreignKey:QuestionMessageID;belongsTo"`
	AnswerMessage         *Message           `gorm:"foreignKey:AnswerMessageID;belongsTo"`
	Team                  *Team              `gorm:"belongsTo"`
	PublishedByUser       *User              `gorm:"foreignKey:PublishedByUserID;belongsTo"`
}

func (DataroomFaqItem) TableName() string {
	return "DataroomFaqItem"
}

func (m *DataroomFaqItem) BeforeCreate(tx *gorm.DB) error {
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

func (*DataroomFaqItem) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Dataroom struct {
	ID                                     string             `gorm:"column:id;primaryKey" json:"id"`
	PID                                    string             `gorm:"column:pId;uniqueIndex:Dataroom_pId_key;not null" json:"pId"`
	Name                                   string             `gorm:"column:name;not null" json:"name"`
	InternalName                           *string            `gorm:"column:internalName" json:"internalName"`
	Description                            *string            `gorm:"column:description" json:"description"`
	TeamID                                 string             `gorm:"column:teamId;index:Dataroom_teamId_idx;not null" json:"teamId"`
	BrandID                                *string            `gorm:"column:brandId;index:Dataroom_brandId_idx" json:"brandId"`
	ConversationsEnabled                   bool               `gorm:"column:conversationsEnabled;default:false;not null" json:"conversationsEnabled"`
	AgentsEnabled                          bool               `gorm:"column:agentsEnabled;default:false;not null" json:"agentsEnabled"`
	VectorStoreID                          *string            `gorm:"column:vectorStoreId" json:"vectorStoreId"`
	RequestListEnabled                     bool               `gorm:"column:requestListEnabled;default:false;not null" json:"requestListEnabled"`
	EnableChangeNotifications              bool               `gorm:"column:enableChangeNotifications;default:false;not null" json:"enableChangeNotifications"`
	EnableVisitorUploadChangeNotifications bool               `gorm:"column:enableVisitorUploadChangeNotifications;default:false;not null" json:"enableVisitorUploadChangeNotifications"`
	DefaultPermissionStrategy              string             `gorm:"column:defaultPermissionStrategy;default:'INHERIT_FROM_PARENT';not null" json:"defaultPermissionStrategy"`
	DefaultShowBanner                      *bool              `gorm:"column:defaultShowBanner;default:true;not null" json:"defaultShowBanner"`
	DefaultGroupPermissionStrategy         string             `gorm:"column:defaultGroupPermissionStrategy;default:'INHERIT_FROM_PARENT';not null" json:"defaultGroupPermissionStrategy"`
	DefaultRootItemAccess                  string             `gorm:"column:defaultRootItemAccess;default:'VIEW_ONLY';not null" json:"defaultRootItemAccess"`
	DefaultGroupRootItemAccess             string             `gorm:"column:defaultGroupRootItemAccess;default:'VIEW_ONLY';not null" json:"defaultGroupRootItemAccess"`
	AllowBulkDownload                      *bool              `gorm:"column:allowBulkDownload;default:true;not null" json:"allowBulkDownload"`
	ShowLastUpdated                        *bool              `gorm:"column:showLastUpdated;default:true;not null" json:"showLastUpdated"`
	IntroductionEnabled                    bool               `gorm:"column:introductionEnabled;default:false;not null" json:"introductionEnabled"`
	IntroductionContent                    datatypes.JSON     `gorm:"column:introductionContent" json:"introductionContent"`
	IsFrozen                               bool               `gorm:"column:isFrozen;default:false;not null" json:"isFrozen"`
	FrozenAt                               *DateTime          `gorm:"column:frozenAt;type:timestamp(3)" json:"frozenAt"`
	FrozenBy                               *string            `gorm:"column:frozenBy" json:"frozenBy"`
	FreezeArchiveURL                       *string            `gorm:"column:freezeArchiveUrl" json:"freezeArchiveUrl"`
	FreezeArchiveHash                      *string            `gorm:"column:freezeArchiveHash" json:"freezeArchiveHash"`
	CreatedAt                              DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                              DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team                                   *Team              `gorm:"belongsTo"`
	BaseBrand                              *Brand             `gorm:"foreignKey:BrandID;belongsTo"`
	FrozenByUser                           *User              `gorm:"foreignKey:FrozenBy;belongsTo"`
	Documents                              []DataroomDocument `gorm:"foreignKey:DataroomID;constraint:DataroomDocument_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Folders                                []DataroomFolder   `gorm:"foreignKey:DataroomID;constraint:DataroomFolder_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Links                                  []Link             `gorm:"foreignKey:DataroomID;constraint:Link_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Views                                  []View             `gorm:"foreignKey:DataroomID;constraint:View_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Viewers                                []Viewer           `gorm:"foreignKey:DataroomID;constraint:Viewer_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ViewerGroups                           []ViewerGroup      `gorm:"foreignKey:DataroomID;constraint:ViewerGroup_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserDatarooms                          []UserDataroom     `gorm:"foreignKey:DataroomID;constraint:UserDataroom_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PermissionGroups                       []PermissionGroup  `gorm:"foreignKey:DataroomID;constraint:PermissionGroup_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Conversations                          []Conversation     `gorm:"foreignKey:DataroomID;constraint:Conversation_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Questions                              []DataroomQuestion `gorm:"foreignKey:DataroomID;constraint:DataroomQuestion_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FaqItems                               []DataroomFaqItem  `gorm:"foreignKey:DataroomID;constraint:DataroomFaqItem_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Chats                                  []Chat             `gorm:"foreignKey:DataroomID;constraint:Chat_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UploadedDocuments                      []DocumentUpload   `gorm:"foreignKey:DataroomID;constraint:DocumentUpload_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tags                                   []TagItem          `gorm:"foreignKey:DataroomID;constraint:TagItem_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskLists                              []TaskList         `gorm:"foreignKey:DataroomID;constraint:TaskList_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tasks                                  []Task             `gorm:"foreignKey:DataroomID;constraint:Task_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Brand                                  *DataroomBrand     `gorm:"foreignKey:DataroomID;constraint:DataroomBrand_dataroomId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Dataroom) TableName() string {
	return "Dataroom"
}

func (m *Dataroom) BeforeCreate(tx *gorm.DB) error {
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

func (*Dataroom) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TaskList struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	OrderIndex *int      `gorm:"column:orderIndex;index:TaskList_dataroomId_orderIndex_idx,priority:2;size:32" json:"orderIndex"`
	DataroomID string    `gorm:"column:dataroomId;index:TaskList_dataroomId_idx;index:TaskList_dataroomId_orderIndex_idx,priority:1;not null" json:"dataroomId"`
	TeamID     string    `gorm:"column:teamId;index:TaskList_teamId_idx;not null" json:"teamId"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom   *Dataroom `gorm:"belongsTo"`
	Team       *Team     `gorm:"belongsTo"`
	Tasks      []Task    `gorm:"foreignKey:TaskListID;constraint:Task_taskListId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TaskList) TableName() string {
	return "TaskList"
}

func (m *TaskList) BeforeCreate(tx *gorm.DB) error {
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

func (*TaskList) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Task struct {
	ID                     string           `gorm:"column:id;primaryKey" json:"id"`
	Title                  string           `gorm:"column:title;not null" json:"title"`
	Description            *string          `gorm:"column:description" json:"description"`
	Type                   *string          `gorm:"column:type;default:'TODO';not null" json:"type"`
	Status                 *string          `gorm:"column:status;default:'OPEN';not null" json:"status"`
	Category               *string          `gorm:"column:category" json:"category"`
	DueDate                *DateTime        `gorm:"column:dueDate;type:timestamp(3)" json:"dueDate"`
	OrderIndex             *int             `gorm:"column:orderIndex;index:Task_taskListId_orderIndex_idx,priority:2;size:32" json:"orderIndex"`
	TaskListID             string           `gorm:"column:taskListId;index:Task_taskListId_idx;index:Task_taskListId_orderIndex_idx,priority:1;not null" json:"taskListId"`
	DataroomID             string           `gorm:"column:dataroomId;index:Task_dataroomId_idx;not null" json:"dataroomId"`
	TeamID                 string           `gorm:"column:teamId;index:Task_teamId_idx;not null" json:"teamId"`
	UploadFolderID         *string          `gorm:"column:uploadFolderId;index:Task_uploadFolderId_idx" json:"uploadFolderId"`
	CreatedByUserID        *string          `gorm:"column:createdByUserId" json:"createdByUserId"`
	LastActivityNotifiedAt *DateTime        `gorm:"column:lastActivityNotifiedAt;type:timestamp(3)" json:"lastActivityNotifiedAt"`
	CreatedAt              DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt              DateTime         `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	TaskList               *TaskList        `gorm:"belongsTo"`
	Dataroom               *Dataroom        `gorm:"belongsTo"`
	Team                   *Team            `gorm:"belongsTo"`
	UploadFolder           *DataroomFolder  `gorm:"foreignKey:UploadFolderID;belongsTo"`
	CreatedByUser          *User            `gorm:"foreignKey:CreatedByUserID;belongsTo"`
	Assignments            []TaskAssignment `gorm:"foreignKey:TaskID;constraint:TaskAssignment_taskId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Activities             []TaskActivity   `gorm:"foreignKey:TaskID;constraint:TaskActivity_taskId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Uploads                []DocumentUpload `gorm:"foreignKey:TaskID;constraint:DocumentUpload_taskId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Task) TableName() string {
	return "Task"
}

func (m *Task) BeforeCreate(tx *gorm.DB) error {
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

func (*Task) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TaskAssignment struct {
	ID             string       `gorm:"column:id;primaryKey" json:"id"`
	TaskID         string       `gorm:"column:taskId;uniqueIndex:TaskAssignment_taskId_viewerId_key,priority:1;uniqueIndex:TaskAssignment_taskId_groupId_key,priority:1;uniqueIndex:TaskAssignment_taskId_linkId_key,priority:1;uniqueIndex:TaskAssignment_taskId_email_key,priority:1;index:TaskAssignment_taskId_idx;not null" json:"taskId"`
	ViewerID       *string      `gorm:"column:viewerId;uniqueIndex:TaskAssignment_taskId_viewerId_key,priority:2;index:TaskAssignment_viewerId_idx" json:"viewerId"`
	GroupID        *string      `gorm:"column:groupId;uniqueIndex:TaskAssignment_taskId_groupId_key,priority:2;index:TaskAssignment_groupId_idx" json:"groupId"`
	LinkID         *string      `gorm:"column:linkId;uniqueIndex:TaskAssignment_taskId_linkId_key,priority:2;index:TaskAssignment_linkId_idx" json:"linkId"`
	Email          *string      `gorm:"column:email;uniqueIndex:TaskAssignment_taskId_email_key,priority:2" json:"email"`
	NotifiedAt     *DateTime    `gorm:"column:notifiedAt;type:timestamp(3)" json:"notifiedAt"`
	LastReminderAt *DateTime    `gorm:"column:lastReminderAt;type:timestamp(3)" json:"lastReminderAt"`
	CreatedAt      DateTime     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime     `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Task           *Task        `gorm:"belongsTo"`
	Viewer         *Viewer      `gorm:"belongsTo"`
	Group          *ViewerGroup `gorm:"foreignKey:GroupID;belongsTo"`
	Link           *Link        `gorm:"belongsTo"`
}

func (TaskAssignment) TableName() string {
	return "TaskAssignment"
}

func (m *TaskAssignment) BeforeCreate(tx *gorm.DB) error {
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

func (*TaskAssignment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TaskActivity struct {
	ID         string   `gorm:"column:id;primaryKey" json:"id"`
	TaskID     string   `gorm:"column:taskId;index:TaskActivity_taskId_idx;index:TaskActivity_taskId_createdAt_idx,priority:1;not null" json:"taskId"`
	Type       string   `gorm:"column:type;not null" json:"type"`
	Comment    *string  `gorm:"column:comment" json:"comment"`
	FromStatus *string  `gorm:"column:fromStatus" json:"fromStatus"`
	ToStatus   *string  `gorm:"column:toStatus" json:"toStatus"`
	UserID     *string  `gorm:"column:userId;index:TaskActivity_userId_idx" json:"userId"`
	ViewerID   *string  `gorm:"column:viewerId;index:TaskActivity_viewerId_idx" json:"viewerId"`
	ViewID     *string  `gorm:"column:viewId;index:TaskActivity_viewId_idx" json:"viewId"`
	CreatedAt  DateTime `gorm:"column:createdAt;index:TaskActivity_taskId_createdAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Task       *Task    `gorm:"belongsTo"`
	User       *User    `gorm:"belongsTo"`
	Viewer     *Viewer  `gorm:"belongsTo"`
	View       *View    `gorm:"belongsTo"`
}

func (TaskActivity) TableName() string {
	return "TaskActivity"
}

func (m *TaskActivity) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DataroomDocument struct {
	ID                string             `gorm:"column:id;primaryKey" json:"id"`
	DataroomID        string             `gorm:"column:dataroomId;uniqueIndex:DataroomDocument_dataroomId_documentId_key,priority:1;index:DataroomDocument_dataroomId_folderId_orderIndex_idx,priority:1;not null" json:"dataroomId"`
	DocumentID        string             `gorm:"column:documentId;uniqueIndex:DataroomDocument_dataroomId_documentId_key,priority:2;index:DataroomDocument_documentId_idx;not null" json:"documentId"`
	FolderID          *string            `gorm:"column:folderId;index:DataroomDocument_folderId_idx;index:DataroomDocument_dataroomId_folderId_orderIndex_idx,priority:2" json:"folderId"`
	OrderIndex        *int               `gorm:"column:orderIndex;index:DataroomDocument_dataroomId_folderId_orderIndex_idx,priority:3;size:32" json:"orderIndex"`
	HierarchicalIndex *string            `gorm:"column:hierarchicalIndex" json:"hierarchicalIndex"`
	VectorStoreFileID *string            `gorm:"column:vectorStoreFileId" json:"vectorStoreFileId"`
	CreatedAt         DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom          *Dataroom          `gorm:"belongsTo"`
	Document          *Document          `gorm:"belongsTo"`
	Folder            *DataroomFolder    `gorm:"foreignKey:FolderID;belongsTo"`
	Conversations     []Conversation     `gorm:"foreignKey:DataroomDocumentID;constraint:Conversation_dataroomDocumentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Questions         []DataroomQuestion `gorm:"foreignKey:DataroomDocumentID;constraint:DataroomQuestion_dataroomDocumentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	DataroomFAQs      []DataroomFaqItem  `gorm:"foreignKey:DataroomDocumentID;constraint:DataroomFaqItem_dataroomDocumentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	UploadedDocuments []DocumentUpload   `gorm:"foreignKey:DataroomDocumentID;constraint:DocumentUpload_dataroomDocumentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (DataroomDocument) TableName() string {
	return "DataroomDocument"
}

func (m *DataroomDocument) BeforeCreate(tx *gorm.DB) error {
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

func (*DataroomDocument) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DataroomFolder struct {
	ID                string             `gorm:"column:id;primaryKey" json:"id"`
	Name              string             `gorm:"column:name;not null" json:"name"`
	Path              string             `gorm:"column:path;uniqueIndex:DataroomFolder_dataroomId_path_key,priority:2;not null" json:"path"`
	ParentID          *string            `gorm:"column:parentId;index:DataroomFolder_parentId_idx;index:DataroomFolder_dataroomId_parentId_orderIndex_idx,priority:2" json:"parentId"`
	Icon              *string            `gorm:"column:icon" json:"icon"`
	Color             *string            `gorm:"column:color" json:"color"`
	DataroomID        string             `gorm:"column:dataroomId;uniqueIndex:DataroomFolder_dataroomId_path_key,priority:1;index:DataroomFolder_dataroomId_parentId_orderIndex_idx,priority:1;not null" json:"dataroomId"`
	OrderIndex        *int               `gorm:"column:orderIndex;index:DataroomFolder_dataroomId_parentId_orderIndex_idx,priority:3;size:32" json:"orderIndex"`
	HierarchicalIndex *string            `gorm:"column:hierarchicalIndex" json:"hierarchicalIndex"`
	CreatedAt         DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	ParentFolder      *DataroomFolder    `gorm:"foreignKey:ParentID;belongsTo"`
	Dataroom          *Dataroom          `gorm:"belongsTo"`
	Documents         []DataroomDocument `gorm:"foreignKey:FolderID;constraint:DataroomDocument_folderId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ChildFolders      []DataroomFolder   `gorm:"foreignKey:ParentID;constraint:DataroomFolder_parentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	UploadTasks       []Task             `gorm:"foreignKey:UploadFolderID;constraint:Task_uploadFolderId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (DataroomFolder) TableName() string {
	return "DataroomFolder"
}

func (m *DataroomFolder) BeforeCreate(tx *gorm.DB) error {
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

func (*DataroomFolder) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DataroomBrand struct {
	ID                             string    `gorm:"column:id;primaryKey" json:"id"`
	Logo                           *string   `gorm:"column:logo" json:"logo"`
	HideLogo                       *bool     `gorm:"column:hideLogo" json:"hideLogo"`
	Banner                         *string   `gorm:"column:banner" json:"banner"`
	BrandColor                     *string   `gorm:"column:brandColor" json:"brandColor"`
	AccentColor                    *string   `gorm:"column:accentColor" json:"accentColor"`
	AccentButtonColor              *string   `gorm:"column:accentButtonColor" json:"accentButtonColor"`
	ApplyAccentColorToDataroomView bool      `gorm:"column:applyAccentColorToDataroomView;default:false;not null" json:"applyAccentColorToDataroomView"`
	WelcomeMessage                 *string   `gorm:"column:welcomeMessage" json:"welcomeMessage"`
	CardLayout                     *string   `gorm:"column:cardLayout;default:'LIST';not null" json:"cardLayout"`
	ShowFolderTree                 *bool     `gorm:"column:showFolderTree;default:true;not null" json:"showFolderTree"`
	ViewerLayoutPreset             *string   `gorm:"column:viewerLayoutPreset;default:'STANDARD';not null" json:"viewerLayoutPreset"`
	ViewerHeaderStyle              *string   `gorm:"column:viewerHeaderStyle;default:'DEFAULT';not null" json:"viewerHeaderStyle"`
	HideFolderIconsInMain          bool      `gorm:"column:hideFolderIconsInMain;default:false;not null" json:"hideFolderIconsInMain"`
	CtaLabel                       *string   `gorm:"column:ctaLabel" json:"ctaLabel"`
	CtaURL                         *string   `gorm:"column:ctaUrl" json:"ctaUrl"`
	CustomLinkPreviewEnabled       bool      `gorm:"column:customLinkPreviewEnabled;default:false;not null" json:"customLinkPreviewEnabled"`
	LinkPreviewTitle               *string   `gorm:"column:linkPreviewTitle" json:"linkPreviewTitle"`
	LinkPreviewDescription         *string   `gorm:"column:linkPreviewDescription" json:"linkPreviewDescription"`
	LinkPreviewImage               *string   `gorm:"column:linkPreviewImage" json:"linkPreviewImage"`
	LinkPreviewFavicon             *string   `gorm:"column:linkPreviewFavicon" json:"linkPreviewFavicon"`
	DefaultLanguage                *string   `gorm:"column:defaultLanguage;default:'en';not null" json:"defaultLanguage"`
	DataroomID                     string    `gorm:"column:dataroomId;uniqueIndex:DataroomBrand_dataroomId_key;not null" json:"dataroomId"`
	CreatedAt                      DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                      DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom                       *Dataroom `gorm:"belongsTo"`
}

func (DataroomBrand) TableName() string {
	return "DataroomBrand"
}

func (m *DataroomBrand) BeforeCreate(tx *gorm.DB) error {
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

func (*DataroomBrand) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ViewerGroup struct {
	ID                  string                       `gorm:"column:id;primaryKey" json:"id"`
	Name                string                       `gorm:"column:name;not null" json:"name"`
	Domains             ScalarList[string]           `gorm:"column:domains;type:text[]" json:"domains"`
	AllowAll            bool                         `gorm:"column:allowAll;default:false;not null" json:"allowAll"`
	DataroomID          string                       `gorm:"column:dataroomId;index:ViewerGroup_dataroomId_idx;index:ViewerGroup_dataroomId_createdAt_idx,priority:1;not null" json:"dataroomId"`
	TeamID              string                       `gorm:"column:teamId;index:ViewerGroup_teamId_idx;not null" json:"teamId"`
	CreatedAt           DateTime                     `gorm:"column:createdAt;index:ViewerGroup_dataroomId_createdAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime                     `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom            *Dataroom                    `gorm:"belongsTo"`
	Team                *Team                        `gorm:"belongsTo"`
	Members             []ViewerGroupMembership      `gorm:"foreignKey:GroupID;constraint:ViewerGroupMembership_groupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Links               []Link                       `gorm:"foreignKey:GroupID;constraint:Link_groupId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AccessControls      []ViewerGroupAccessControls  `gorm:"foreignKey:GroupID;constraint:ViewerGroupAccessControls_groupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Views               []View                       `gorm:"foreignKey:GroupID;constraint:View_groupId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Conversations       []Conversation               `gorm:"foreignKey:ViewerGroupID;constraint:Conversation_viewerGroupId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Invitations         []ViewerInvitation           `gorm:"foreignKey:GroupID;constraint:ViewerInvitation_groupId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	QuestionAssignments []DataroomQuestionAssignment `gorm:"foreignKey:GroupID;constraint:DataroomQuestionAssignment_groupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskAssignments     []TaskAssignment             `gorm:"foreignKey:GroupID;constraint:TaskAssignment_groupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ViewerGroup) TableName() string {
	return "ViewerGroup"
}

func (m *ViewerGroup) BeforeCreate(tx *gorm.DB) error {
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

func (*ViewerGroup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ViewerGroupMembership struct {
	ID        string       `gorm:"column:id;primaryKey" json:"id"`
	ViewerID  string       `gorm:"column:viewerId;uniqueIndex:ViewerGroupMembership_viewerId_groupId_key,priority:1;index:ViewerGroupMembership_viewerId_idx;not null" json:"viewerId"`
	GroupID   string       `gorm:"column:groupId;uniqueIndex:ViewerGroupMembership_viewerId_groupId_key,priority:2;index:ViewerGroupMembership_groupId_idx;not null" json:"groupId"`
	CreatedAt DateTime     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime     `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Viewer    *Viewer      `gorm:"belongsTo"`
	Group     *ViewerGroup `gorm:"foreignKey:GroupID;belongsTo"`
}

func (ViewerGroupMembership) TableName() string {
	return "ViewerGroupMembership"
}

func (m *ViewerGroupMembership) BeforeCreate(tx *gorm.DB) error {
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

func (*ViewerGroupMembership) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ViewerGroupAccessControls struct {
	ID          string       `gorm:"column:id;primaryKey" json:"id"`
	GroupID     string       `gorm:"column:groupId;uniqueIndex:ViewerGroupAccessControls_groupId_itemId_key,priority:1;index:ViewerGroupAccessControls_groupId_idx;not null" json:"groupId"`
	ItemID      string       `gorm:"column:itemId;uniqueIndex:ViewerGroupAccessControls_groupId_itemId_key,priority:2;not null" json:"itemId"`
	ItemType    string       `gorm:"column:itemType;not null" json:"itemType"`
	CanView     *bool        `gorm:"column:canView;default:true;not null" json:"canView"`
	CanDownload bool         `gorm:"column:canDownload;default:false;not null" json:"canDownload"`
	CreatedAt   DateTime     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime     `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Group       *ViewerGroup `gorm:"foreignKey:GroupID;belongsTo"`
}

func (ViewerGroupAccessControls) TableName() string {
	return "ViewerGroupAccessControls"
}

func (m *ViewerGroupAccessControls) BeforeCreate(tx *gorm.DB) error {
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

func (*ViewerGroupAccessControls) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PermissionGroup struct {
	ID             string                          `gorm:"column:id;primaryKey" json:"id"`
	Name           string                          `gorm:"column:name;not null" json:"name"`
	Description    *string                         `gorm:"column:description" json:"description"`
	DataroomID     string                          `gorm:"column:dataroomId;index:PermissionGroup_dataroomId_idx;not null" json:"dataroomId"`
	TeamID         string                          `gorm:"column:teamId;index:PermissionGroup_teamId_idx;not null" json:"teamId"`
	CreatedAt      DateTime                        `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime                        `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom       *Dataroom                       `gorm:"belongsTo"`
	Team           *Team                           `gorm:"belongsTo"`
	Links          []Link                          `gorm:"foreignKey:PermissionGroupID;constraint:Link_permissionGroupId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	AccessControls []PermissionGroupAccessControls `gorm:"foreignKey:GroupID;constraint:PermissionGroupAccessControls_groupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (PermissionGroup) TableName() string {
	return "PermissionGroup"
}

func (m *PermissionGroup) BeforeCreate(tx *gorm.DB) error {
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

func (*PermissionGroup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PermissionGroupAccessControls struct {
	ID                  string           `gorm:"column:id;primaryKey" json:"id"`
	GroupID             string           `gorm:"column:groupId;uniqueIndex:PermissionGroupAccessControls_groupId_itemId_key,priority:1;index:PermissionGroupAccessControls_groupId_idx;not null" json:"groupId"`
	ItemID              string           `gorm:"column:itemId;uniqueIndex:PermissionGroupAccessControls_groupId_itemId_key,priority:2;not null" json:"itemId"`
	ItemType            string           `gorm:"column:itemType;not null" json:"itemType"`
	CanView             *bool            `gorm:"column:canView;default:true;not null" json:"canView"`
	CanDownload         bool             `gorm:"column:canDownload;default:false;not null" json:"canDownload"`
	CanDownloadOriginal bool             `gorm:"column:canDownloadOriginal;default:false;not null" json:"canDownloadOriginal"`
	CreatedAt           DateTime         `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime         `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Group               *PermissionGroup `gorm:"foreignKey:GroupID;belongsTo"`
}

func (PermissionGroupAccessControls) TableName() string {
	return "PermissionGroupAccessControls"
}

func (m *PermissionGroupAccessControls) BeforeCreate(tx *gorm.DB) error {
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

func (*PermissionGroupAccessControls) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Document struct {
	ID                   string                 `gorm:"column:id;primaryKey" json:"id"`
	Name                 string                 `gorm:"column:name;index:Document_teamId_name_idx,priority:2;not null" json:"name"`
	Description          *string                `gorm:"column:description" json:"description"`
	File                 string                 `gorm:"column:file;not null" json:"file"`
	OriginalFile         *string                `gorm:"column:originalFile" json:"originalFile"`
	Type                 *string                `gorm:"column:type" json:"type"`
	ContentType          *string                `gorm:"column:contentType" json:"contentType"`
	StorageType          string                 `gorm:"column:storageType;default:'VERCEL_BLOB';not null" json:"storageType"`
	NumPages             *int                   `gorm:"column:numPages;size:32" json:"numPages"`
	TeamID               string                 `gorm:"column:teamId;index:Document_teamId_idx;index:Document_teamId_folderId_idx,priority:1;index:Document_teamId_name_idx,priority:1;index:Document_teamId_hiddenInAllDocuments_idx,priority:1;not null" json:"teamId"`
	OwnerID              *string                `gorm:"column:ownerId;index:Document_ownerId_idx" json:"ownerId"`
	AssistantEnabled     bool                   `gorm:"column:assistantEnabled;default:false;not null" json:"assistantEnabled"`
	AdvancedExcelEnabled bool                   `gorm:"column:advancedExcelEnabled;default:false;not null" json:"advancedExcelEnabled"`
	AgentsEnabled        bool                   `gorm:"column:agentsEnabled;default:false;not null" json:"agentsEnabled"`
	DownloadOnly         bool                   `gorm:"column:downloadOnly;default:false;not null" json:"downloadOnly"`
	HiddenInAllDocuments bool                   `gorm:"column:hiddenInAllDocuments;index:Document_teamId_hiddenInAllDocuments_idx,priority:2;default:false;not null" json:"hiddenInAllDocuments"`
	CreatedAt            DateTime               `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime               `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	FolderID             *string                `gorm:"column:folderId;index:Document_folderId_idx;index:Document_teamId_folderId_idx,priority:2" json:"folderId"`
	IsExternalUpload     bool                   `gorm:"column:isExternalUpload;default:false;not null" json:"isExternalUpload"`
	Owner                *User                  `gorm:"foreignKey:OwnerID;belongsTo"`
	Team                 *Team                  `gorm:"belongsTo"`
	Folder               *Folder                `gorm:"belongsTo"`
	Links                []Link                 `gorm:"foreignKey:DocumentID;constraint:Link_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Views                []View                 `gorm:"foreignKey:DocumentID;constraint:View_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Versions             []DocumentVersion      `gorm:"foreignKey:DocumentID;constraint:DocumentVersion_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Chats                []Chat                 `gorm:"foreignKey:DocumentID;constraint:Chat_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Datarooms            []DataroomDocument     `gorm:"foreignKey:DocumentID;constraint:DataroomDocument_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags                 []TagItem              `gorm:"foreignKey:DocumentID;constraint:TagItem_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Annotations          []DocumentAnnotation   `gorm:"foreignKey:DocumentID;constraint:DocumentAnnotation_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RedactionJobs        []DocumentRedactionJob `gorm:"foreignKey:DocumentID;constraint:DocumentRedactionJob_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UploadedDocument     []DocumentUpload       `gorm:"foreignKey:DocumentID;constraint:DocumentUpload_documentId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Document) TableName() string {
	return "Document"
}

func (m *Document) BeforeCreate(tx *gorm.DB) error {
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

func (*Document) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DocumentVersion struct {
	ID                string                 `gorm:"column:id;primaryKey" json:"id"`
	VersionNumber     int                    `gorm:"column:versionNumber;uniqueIndex:DocumentVersion_versionNumber_documentId_key,priority:1;size:32;not null" json:"versionNumber"`
	DocumentID        string                 `gorm:"column:documentId;uniqueIndex:DocumentVersion_versionNumber_documentId_key,priority:2;index:DocumentVersion_documentId_idx;index:DocumentVersion_documentId_isPrimary_idx,priority:1;index:DocumentVersion_documentId_createdAt_idx,priority:1;index:DocumentVersion_documentId_isPrimary_createdAt_idx,priority:1;not null" json:"documentId"`
	File              string                 `gorm:"column:file;not null" json:"file"`
	OriginalFile      *string                `gorm:"column:originalFile" json:"originalFile"`
	Type              *string                `gorm:"column:type" json:"type"`
	ContentType       *string                `gorm:"column:contentType" json:"contentType"`
	FileSize          *int64                 `gorm:"column:fileSize" json:"fileSize"`
	StorageType       string                 `gorm:"column:storageType;default:'VERCEL_BLOB';not null" json:"storageType"`
	NumPages          *int                   `gorm:"column:numPages;size:32" json:"numPages"`
	IsPrimary         bool                   `gorm:"column:isPrimary;index:DocumentVersion_documentId_isPrimary_idx,priority:2;index:DocumentVersion_documentId_isPrimary_createdAt_idx,priority:2;default:false;not null" json:"isPrimary"`
	IsVertical        bool                   `gorm:"column:isVertical;default:false;not null" json:"isVertical"`
	FileID            *string                `gorm:"column:fileId" json:"fileId"`
	VectorStoreFileID *string                `gorm:"column:vectorStoreFileId" json:"vectorStoreFileId"`
	HasPages          bool                   `gorm:"column:hasPages;default:false;not null" json:"hasPages"`
	Length            *int                   `gorm:"column:length;size:32" json:"length"`
	CreatedAt         DateTime               `gorm:"column:createdAt;index:DocumentVersion_documentId_createdAt_idx,priority:2,sort:desc;index:DocumentVersion_documentId_isPrimary_createdAt_idx,priority:3;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime               `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Document          *Document              `gorm:"belongsTo"`
	Pages             []DocumentPage         `gorm:"foreignKey:VersionID;constraint:DocumentPage_versionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RedactionJobs     []DocumentRedactionJob `gorm:"foreignKey:DocumentVersionID;constraint:DocumentRedactionJob_documentVersionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (DocumentVersion) TableName() string {
	return "DocumentVersion"
}

func (m *DocumentVersion) BeforeCreate(tx *gorm.DB) error {
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

func (*DocumentVersion) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DocumentPage struct {
	ID            string             `gorm:"column:id;primaryKey" json:"id"`
	VersionID     string             `gorm:"column:versionId;uniqueIndex:DocumentPage_pageNumber_versionId_key,priority:2;index:DocumentPage_versionId_idx;not null" json:"versionId"`
	PageNumber    int                `gorm:"column:pageNumber;uniqueIndex:DocumentPage_pageNumber_versionId_key,priority:1;size:32;not null" json:"pageNumber"`
	EmbeddedLinks ScalarList[string] `gorm:"column:embeddedLinks;type:text[]" json:"embeddedLinks"`
	PageLinks     datatypes.JSON     `gorm:"column:pageLinks" json:"pageLinks"`
	Metadata      datatypes.JSON     `gorm:"column:metadata" json:"metadata"`
	File          string             `gorm:"column:file;not null" json:"file"`
	StorageType   string             `gorm:"column:storageType;default:'VERCEL_BLOB';not null" json:"storageType"`
	CreatedAt     DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Version       *DocumentVersion   `gorm:"foreignKey:VersionID;belongsTo"`
}

func (DocumentPage) TableName() string {
	return "DocumentPage"
}

func (m *DocumentPage) BeforeCreate(tx *gorm.DB) error {
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

func (*DocumentPage) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Folder struct {
	ID                   string     `gorm:"column:id;primaryKey" json:"id"`
	Name                 string     `gorm:"column:name;not null" json:"name"`
	Path                 string     `gorm:"column:path;uniqueIndex:Folder_teamId_path_key,priority:2;not null" json:"path"`
	ParentID             *string    `gorm:"column:parentId;index:Folder_parentId_idx" json:"parentId"`
	HiddenInAllDocuments bool       `gorm:"column:hiddenInAllDocuments;index:Folder_teamId_hiddenInAllDocuments_idx,priority:2;default:false;not null" json:"hiddenInAllDocuments"`
	Icon                 *string    `gorm:"column:icon" json:"icon"`
	Color                *string    `gorm:"column:color" json:"color"`
	TeamID               string     `gorm:"column:teamId;uniqueIndex:Folder_teamId_path_key,priority:1;index:Folder_teamId_hiddenInAllDocuments_idx,priority:1;not null" json:"teamId"`
	CreatedAt            DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	ParentFolder         *Folder    `gorm:"foreignKey:ParentID;belongsTo"`
	Team                 *Team      `gorm:"belongsTo"`
	Documents            []Document `gorm:"foreignKey:FolderID;constraint:Document_folderId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ChildFolders         []Folder   `gorm:"foreignKey:ParentID;constraint:Folder_parentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Folder) TableName() string {
	return "Folder"
}

func (m *Folder) BeforeCreate(tx *gorm.DB) error {
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

func (*Folder) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DocumentUpload struct {
	ID                 string            `gorm:"column:id;primaryKey" json:"id"`
	DocumentID         string            `gorm:"column:documentId;index:DocumentUpload_documentId_idx;not null" json:"documentId"`
	TeamID             string            `gorm:"column:teamId;index:DocumentUpload_teamId_idx;not null" json:"teamId"`
	ViewerID           *string           `gorm:"column:viewerId;index:DocumentUpload_viewerId_idx" json:"viewerId"`
	ViewID             *string           `gorm:"column:viewId;index:DocumentUpload_viewId_idx" json:"viewId"`
	LinkID             string            `gorm:"column:linkId;index:DocumentUpload_linkId_idx;not null" json:"linkId"`
	DataroomID         *string           `gorm:"column:dataroomId;index:DocumentUpload_dataroomId_idx" json:"dataroomId"`
	DataroomDocumentID *string           `gorm:"column:dataroomDocumentId;index:DocumentUpload_dataroomDocumentId_idx" json:"dataroomDocumentId"`
	TaskID             *string           `gorm:"column:taskId;index:DocumentUpload_taskId_idx" json:"taskId"`
	OriginalFilename   *string           `gorm:"column:originalFilename" json:"originalFilename"`
	FileSize           *int64            `gorm:"column:fileSize" json:"fileSize"`
	NumPages           *int              `gorm:"column:numPages;size:32" json:"numPages"`
	MimeType           *string           `gorm:"column:mimeType" json:"mimeType"`
	UploadedAt         DateTime          `gorm:"column:uploadedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"uploadedAt"`
	CreatedAt          DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt          DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Document           *Document         `gorm:"belongsTo"`
	Team               *Team             `gorm:"belongsTo"`
	Viewer             *Viewer           `gorm:"belongsTo"`
	View               *View             `gorm:"belongsTo"`
	Link               *Link             `gorm:"belongsTo"`
	Dataroom           *Dataroom         `gorm:"belongsTo"`
	DataroomDocument   *DataroomDocument `gorm:"belongsTo"`
	Task               *Task             `gorm:"belongsTo"`
}

func (DocumentUpload) TableName() string {
	return "DocumentUpload"
}

func (m *DocumentUpload) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.UploadedAt.IsZero() {
		m.UploadedAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*DocumentUpload) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type PendingDocumentUpload struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	UploadID    string    `gorm:"column:uploadId;uniqueIndex:PendingDocumentUpload_uploadId_key;not null" json:"uploadId"`
	TeamID      string    `gorm:"column:teamId;index:PendingDocumentUpload_teamId_expiresAt_idx,priority:1;not null" json:"teamId"`
	TokenID     string    `gorm:"column:tokenId;index:PendingDocumentUpload_tokenId_expiresAt_idx,priority:1;not null" json:"tokenId"`
	Key         string    `gorm:"column:key;not null" json:"key"`
	ContentType string    `gorm:"column:contentType;not null" json:"contentType"`
	StorageType string    `gorm:"column:storageType;default:'S3_PATH';not null" json:"storageType"`
	ExpiresAt   DateTime  `gorm:"column:expiresAt;index:PendingDocumentUpload_teamId_expiresAt_idx,priority:2;index:PendingDocumentUpload_tokenId_expiresAt_idx,priority:2;type:timestamp(3);not null" json:"expiresAt"`
	ConsumedAt  *DateTime `gorm:"column:consumedAt;type:timestamp(3)" json:"consumedAt"`
	CreatedAt   DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (PendingDocumentUpload) TableName() string {
	return "PendingDocumentUpload"
}

func (m *PendingDocumentUpload) BeforeCreate(tx *gorm.DB) error {
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
	ID                    string                 `gorm:"column:id;primaryKey" json:"id"`
	Name                  string                 `gorm:"column:name;not null" json:"name"`
	Slug                  string                 `gorm:"column:slug;uniqueIndex:Integration_slug_key;not null" json:"slug"`
	Description           *string                `gorm:"column:description" json:"description"`
	Readme                *string                `gorm:"column:readme" json:"readme"`
	Developer             string                 `gorm:"column:developer;not null" json:"developer"`
	Website               string                 `gorm:"column:website;not null" json:"website"`
	Logo                  *string                `gorm:"column:logo" json:"logo"`
	Screenshots           datatypes.JSON         `gorm:"column:screenshots" json:"screenshots"`
	Verified              bool                   `gorm:"column:verified;default:false;not null" json:"verified"`
	InstallURL            *string                `gorm:"column:installUrl" json:"installUrl"`
	Category              *string                `gorm:"column:category" json:"category"`
	ComingSoon            bool                   `gorm:"column:comingSoon;default:false;not null" json:"comingSoon"`
	CreatedAt             DateTime               `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime               `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	InstalledIntegrations []InstalledIntegration `gorm:"foreignKey:IntegrationID;constraint:InstalledIntegration_integrationId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	Credentials   datatypes.JSON `gorm:"column:credentials" json:"credentials"`
	Configuration datatypes.JSON `gorm:"column:configuration" json:"configuration"`
	Enabled       *bool          `gorm:"column:enabled;default:true;not null" json:"enabled"`
	IntegrationID string         `gorm:"column:integrationId;uniqueIndex:InstalledIntegration_teamId_integrationId_key,priority:2;index:InstalledIntegration_integrationId_idx;not null" json:"integrationId"`
	UserID        *string        `gorm:"column:userId" json:"userId"`
	TeamID        string         `gorm:"column:teamId;uniqueIndex:InstalledIntegration_teamId_integrationId_key,priority:1;index:InstalledIntegration_teamId_idx;not null" json:"teamId"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Integration   *Integration   `gorm:"belongsTo"`
	User          *User          `gorm:"belongsTo"`
	Team          *Team          `gorm:"belongsTo"`
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

type JacksonIndex struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Key      string `gorm:"column:key;index:_jackson_index_key_store,priority:1;not null" json:"key"`
	StoreKey string `gorm:"column:storeKey;index:_jackson_index_key_store,priority:2;not null" json:"storeKey"`
}

func (JacksonIndex) TableName() string {
	return "jackson_index"
}

type JacksonStore struct {
	Key        string    `gorm:"column:key;primaryKey" json:"key"`
	Value      string    `gorm:"column:value;not null" json:"value"`
	Iv         *string   `gorm:"column:iv" json:"iv"`
	Tag        *string   `gorm:"column:tag" json:"tag"`
	Namespace  *string   `gorm:"column:namespace;index:_jackson_store_namespace" json:"namespace"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(0);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
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
	Key       string `gorm:"column:key;primaryKey" json:"key"`
	ExpiresAt int64  `gorm:"column:expiresAt;index:_jackson_ttl_expires_at;not null" json:"expiresAt"`
}

func (JacksonTTL) TableName() string {
	return "jackson_ttl"
}

type Link struct {
	ID                         string             `gorm:"column:id;primaryKey" json:"id"`
	DocumentID                 *string            `gorm:"column:documentId;index:Link_documentId_idx;index:Link_documentId_isArchived_idx,priority:1" json:"documentId"`
	DataroomID                 *string            `gorm:"column:dataroomId" json:"dataroomId"`
	LinkType                   string             `gorm:"column:linkType;default:'DOCUMENT_LINK';not null" json:"linkType"`
	URL                        *string            `gorm:"column:url;uniqueIndex:Link_url_key" json:"url"`
	Name                       *string            `gorm:"column:name" json:"name"`
	Slug                       *string            `gorm:"column:slug;uniqueIndex:Link_domainSlug_slug_key,priority:2" json:"slug"`
	ExpiresAt                  *DateTime          `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	Password                   *string            `gorm:"column:password" json:"password"`
	AllowList                  ScalarList[string] `gorm:"column:allowList;type:text[]" json:"allowList"`
	DenyList                   ScalarList[string] `gorm:"column:denyList;type:text[]" json:"denyList"`
	EmailProtected             *bool              `gorm:"column:emailProtected;default:true;not null" json:"emailProtected"`
	EmailAuthenticated         bool               `gorm:"column:emailAuthenticated;default:false;not null" json:"emailAuthenticated"`
	AllowDownload              *bool              `gorm:"column:allowDownload;default:false" json:"allowDownload"`
	IsArchived                 bool               `gorm:"column:isArchived;index:Link_documentId_isArchived_idx,priority:2;default:false;not null" json:"isArchived"`
	DeletedAt                  *DateTime          `gorm:"column:deletedAt;index:Link_deletedAt_idx;type:timestamp(3)" json:"deletedAt"`
	DomainID                   *string            `gorm:"column:domainId" json:"domainId"`
	DomainSlug                 *string            `gorm:"column:domainSlug;uniqueIndex:Link_domainSlug_slug_key,priority:1" json:"domainSlug"`
	CreatedAt                  DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	EnableNotification         *bool              `gorm:"column:enableNotification;default:true" json:"enableNotification"`
	EnableFeedback             *bool              `gorm:"column:enableFeedback;default:false" json:"enableFeedback"`
	EnableQuestion             *bool              `gorm:"column:enableQuestion;default:false" json:"enableQuestion"`
	EnableScreenshotProtection *bool              `gorm:"column:enableScreenshotProtection;default:false" json:"enableScreenshotProtection"`
	EnableConfidentialView     *bool              `gorm:"column:enableConfidentialView;default:false" json:"enableConfidentialView"`
	EnableAgreement            *bool              `gorm:"column:enableAgreement;default:false" json:"enableAgreement"`
	AgreementID                *string            `gorm:"column:agreementId" json:"agreementId"`
	ShowBanner                 *bool              `gorm:"column:showBanner;default:false" json:"showBanner"`
	EnableWatermark            *bool              `gorm:"column:enableWatermark;default:false" json:"enableWatermark"`
	WatermarkConfig            datatypes.JSON     `gorm:"column:watermarkConfig" json:"watermarkConfig"`
	AudienceType               string             `gorm:"column:audienceType;default:'GENERAL';not null" json:"audienceType"`
	GroupID                    *string            `gorm:"column:groupId" json:"groupId"`
	PermissionGroupID          *string            `gorm:"column:permissionGroupId;index:Link_permissionGroupId_idx" json:"permissionGroupId"`
	MetaTitle                  *string            `gorm:"column:metaTitle" json:"metaTitle"`
	MetaDescription            *string            `gorm:"column:metaDescription" json:"metaDescription"`
	MetaImage                  *string            `gorm:"column:metaImage" json:"metaImage"`
	MetaFavicon                *string            `gorm:"column:metaFavicon" json:"metaFavicon"`
	EnableCustomMetatag        *bool              `gorm:"column:enableCustomMetatag;default:false" json:"enableCustomMetatag"`
	WelcomeMessage             *string            `gorm:"column:welcomeMessage" json:"welcomeMessage"`
	EnableConversation         bool               `gorm:"column:enableConversation;default:false;not null" json:"enableConversation"`
	EnableAIAgents             *bool              `gorm:"column:enableAIAgents;default:false" json:"enableAIAgents"`
	EnableUpload               *bool              `gorm:"column:enableUpload;default:false" json:"enableUpload"`
	IsFileRequestOnly          *bool              `gorm:"column:isFileRequestOnly;default:false" json:"isFileRequestOnly"`
	// Deprecated: replaced by `uploadFolderIds`. Retained in the DB only so old
	// rows can be inspected; no application code reads or writes this column.
	UploadFolderID      *string                      `gorm:"column:uploadFolderId" json:"uploadFolderId"`
	UploadFolderIds     ScalarList[string]           `gorm:"column:uploadFolderIds;type:text[]" json:"uploadFolderIds"`
	EnableIndexFile     *bool                        `gorm:"column:enableIndexFile;default:false" json:"enableIndexFile"`
	TeamID              *string                      `gorm:"column:teamId;index:Link_teamId_idx" json:"teamId"`
	BrandID             *string                      `gorm:"column:brandId;index:Link_brandId_idx" json:"brandId"`
	OwnerID             *string                      `gorm:"column:ownerId;index:Link_ownerId_idx" json:"ownerId"`
	Document            *Document                    `gorm:"belongsTo"`
	Dataroom            *Dataroom                    `gorm:"belongsTo"`
	Domain              *Domain                      `gorm:"belongsTo"`
	Agreement           *Agreement                   `gorm:"belongsTo"`
	Group               *ViewerGroup                 `gorm:"foreignKey:GroupID;belongsTo"`
	PermissionGroup     *PermissionGroup             `gorm:"belongsTo"`
	Team                *Team                        `gorm:"belongsTo"`
	Brand               *Brand                       `gorm:"belongsTo"`
	Owner               *User                        `gorm:"foreignKey:OwnerID;belongsTo"`
	Views               []View                       `gorm:"foreignKey:LinkID;constraint:View_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Conversations       []Conversation               `gorm:"foreignKey:LinkID;constraint:Conversation_linkId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	QuestionAssignments []DataroomQuestionAssignment `gorm:"foreignKey:LinkID;constraint:DataroomQuestionAssignment_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DataroomFaqItems    []DataroomFaqItem            `gorm:"foreignKey:LinkID;constraint:DataroomFaqItem_linkId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Chats               []Chat                       `gorm:"foreignKey:LinkID;constraint:Chat_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UploadedDocuments   []DocumentUpload             `gorm:"foreignKey:LinkID;constraint:DocumentUpload_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomFields        []CustomField                `gorm:"foreignKey:LinkID;constraint:CustomField_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags                []TagItem                    `gorm:"foreignKey:LinkID;constraint:TagItem_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ViewerInvitations   []ViewerInvitation           `gorm:"foreignKey:LinkID;constraint:ViewerInvitation_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	VisitorGroups       []LinkVisitorGroup           `gorm:"foreignKey:LinkID;constraint:LinkVisitorGroup_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskAssignments     []TaskAssignment             `gorm:"foreignKey:LinkID;constraint:TaskAssignment_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Feedback            *Feedback                    `gorm:"foreignKey:LinkID;constraint:Feedback_linkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Workflow            *Workflow                    `gorm:"foreignKey:EntryLinkID;constraint:Workflow_entryLinkId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Link) TableName() string {
	return "Link"
}

func (m *Link) BeforeCreate(tx *gorm.DB) error {
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

func (*Link) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type LinkPreset struct {
	ID                         string             `gorm:"column:id;primaryKey" json:"id"`
	Name                       string             `gorm:"column:name;not null" json:"name"`
	TeamID                     string             `gorm:"column:teamId;index:LinkPreset_teamId_idx;not null" json:"teamId"`
	PID                        *string            `gorm:"column:pId;uniqueIndex:LinkPreset_pId_key" json:"pId"`
	EnableCustomMetaTag        *bool              `gorm:"column:enableCustomMetaTag;default:false" json:"enableCustomMetaTag"`
	MetaTitle                  *string            `gorm:"column:metaTitle" json:"metaTitle"`
	MetaDescription            *string            `gorm:"column:metaDescription" json:"metaDescription"`
	MetaImage                  *string            `gorm:"column:metaImage" json:"metaImage"`
	MetaFavicon                *string            `gorm:"column:metaFavicon" json:"metaFavicon"`
	EnableNotification         *bool              `gorm:"column:enableNotification;default:false" json:"enableNotification"`
	EmailProtected             *bool              `gorm:"column:emailProtected;default:true" json:"emailProtected"`
	EmailAuthenticated         *bool              `gorm:"column:emailAuthenticated;default:false" json:"emailAuthenticated"`
	AllowDownload              *bool              `gorm:"column:allowDownload;default:false" json:"allowDownload"`
	EnableAllowList            *bool              `gorm:"column:enableAllowList;default:false" json:"enableAllowList"`
	AllowList                  ScalarList[string] `gorm:"column:allowList;type:text[]" json:"allowList"`
	EnableDenyList             *bool              `gorm:"column:enableDenyList;default:false" json:"enableDenyList"`
	DenyList                   ScalarList[string] `gorm:"column:denyList;type:text[]" json:"denyList"`
	ExpiresIn                  *int               `gorm:"column:expiresIn;size:32" json:"expiresIn"`
	EnableScreenshotProtection *bool              `gorm:"column:enableScreenshotProtection;default:false" json:"enableScreenshotProtection"`
	EnableConfidentialView     *bool              `gorm:"column:enableConfidentialView;default:false" json:"enableConfidentialView"`
	ExpiresAt                  *DateTime          `gorm:"column:expiresAt;type:timestamp(3)" json:"expiresAt"`
	EnablePassword             *bool              `gorm:"column:enablePassword;default:false" json:"enablePassword"`
	Password                   *string            `gorm:"column:password" json:"password"`
	EnableWatermark            *bool              `gorm:"column:enableWatermark;default:false" json:"enableWatermark"`
	WatermarkConfig            datatypes.JSON     `gorm:"column:watermarkConfig" json:"watermarkConfig"`
	EnableAgreement            *bool              `gorm:"column:enableAgreement;default:false" json:"enableAgreement"`
	AgreementID                *string            `gorm:"column:agreementId" json:"agreementId"`
	EnableCustomFields         *bool              `gorm:"column:enableCustomFields;default:false" json:"enableCustomFields"`
	CustomFields               datatypes.JSON     `gorm:"column:customFields" json:"customFields"`
	ShowBanner                 *bool              `gorm:"column:showBanner;default:false" json:"showBanner"`
	WelcomeMessage             *string            `gorm:"column:welcomeMessage" json:"welcomeMessage"`
	IsDefault                  bool               `gorm:"column:isDefault;default:false;not null" json:"isDefault"`
	CreatedAt                  DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                  DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team                       *Team              `gorm:"belongsTo"`
}

func (LinkPreset) TableName() string {
	return "LinkPreset"
}

func (m *LinkPreset) BeforeCreate(tx *gorm.DB) error {
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

func (*LinkPreset) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type VisitorGroup struct {
	ID        string             `gorm:"column:id;primaryKey" json:"id"`
	Name      string             `gorm:"column:name;not null" json:"name"`
	Emails    ScalarList[string] `gorm:"column:emails;type:text[]" json:"emails"`
	TeamID    string             `gorm:"column:teamId;index:VisitorGroup_teamId_idx;not null" json:"teamId"`
	CreatedAt DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team      *Team              `gorm:"belongsTo"`
	Links     []LinkVisitorGroup `gorm:"foreignKey:VisitorGroupID;constraint:LinkVisitorGroup_visitorGroupId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (VisitorGroup) TableName() string {
	return "VisitorGroup"
}

func (m *VisitorGroup) BeforeCreate(tx *gorm.DB) error {
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

func (*VisitorGroup) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type LinkVisitorGroup struct {
	ID             string        `gorm:"column:id;primaryKey" json:"id"`
	LinkID         string        `gorm:"column:linkId;uniqueIndex:LinkVisitorGroup_linkId_visitorGroupId_key,priority:1;index:LinkVisitorGroup_linkId_idx;not null" json:"linkId"`
	VisitorGroupID string        `gorm:"column:visitorGroupId;uniqueIndex:LinkVisitorGroup_linkId_visitorGroupId_key,priority:2;index:LinkVisitorGroup_visitorGroupId_idx;not null" json:"visitorGroupId"`
	CreatedAt      DateTime      `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Link           *Link         `gorm:"belongsTo"`
	VisitorGroup   *VisitorGroup `gorm:"belongsTo"`
}

func (LinkVisitorGroup) TableName() string {
	return "LinkVisitorGroup"
}

func (m *LinkVisitorGroup) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type CustomField struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Type        string   `gorm:"column:type;not null" json:"type"`
	Identifier  string   `gorm:"column:identifier;not null" json:"identifier"`
	Label       string   `gorm:"column:label;not null" json:"label"`
	Placeholder *string  `gorm:"column:placeholder" json:"placeholder"`
	Required    bool     `gorm:"column:required;default:false;not null" json:"required"`
	Disabled    bool     `gorm:"column:disabled;default:false;not null" json:"disabled"`
	LinkID      string   `gorm:"column:linkId;index:CustomField_linkId_idx;not null" json:"linkId"`
	OrderIndex  int      `gorm:"column:orderIndex;size:32;default:0;not null" json:"orderIndex"`
	Link        *Link    `gorm:"belongsTo"`
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

type CustomFieldResponse struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	Data      datatypes.JSON `gorm:"column:data;not null" json:"data"`
	ViewID    string         `gorm:"column:viewId;uniqueIndex:CustomFieldResponse_viewId_key;index:CustomFieldResponse_viewId_idx;not null" json:"viewId"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	View      *View          `gorm:"belongsTo"`
}

func (CustomFieldResponse) TableName() string {
	return "CustomFieldResponse"
}

func (m *CustomFieldResponse) BeforeCreate(tx *gorm.DB) error {
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

func (*CustomFieldResponse) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type NotificationPreference struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	UserID    string   `gorm:"column:userId;uniqueIndex:NotificationPreference_userId_teamId_type_key,priority:1;index:NotificationPreference_userId_teamId_idx,priority:1;not null" json:"userId"`
	TeamID    string   `gorm:"column:teamId;uniqueIndex:NotificationPreference_userId_teamId_type_key,priority:2;index:NotificationPreference_userId_teamId_idx,priority:2;not null" json:"teamId"`
	Type      string   `gorm:"column:type;uniqueIndex:NotificationPreference_userId_teamId_type_key,priority:3;not null" json:"type"`
	Frequency *string  `gorm:"column:frequency;default:'IMMEDIATE';not null" json:"frequency"`
	Scope     *string  `gorm:"column:scope;default:'ALL';not null" json:"scope"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User      *User    `gorm:"belongsTo"`
	Team      *Team    `gorm:"belongsTo"`
}

func (NotificationPreference) TableName() string {
	return "NotificationPreference"
}

func (m *NotificationPreference) BeforeCreate(tx *gorm.DB) error {
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

func (*NotificationPreference) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OAuthClient struct {
	ID               string             `gorm:"column:id;primaryKey" json:"id"`
	ClientID         string             `gorm:"column:clientId;uniqueIndex:OAuthClient_clientId_key;index:OAuthClient_clientId_idx;not null" json:"clientId"`
	ClientName       string             `gorm:"column:clientName;not null" json:"clientName"`
	ClientURI        *string            `gorm:"column:clientUri" json:"clientUri"`
	LogoURI          *string            `gorm:"column:logoUri" json:"logoUri"`
	RedirectUris     ScalarList[string] `gorm:"column:redirectUris;type:text[]" json:"redirectUris"`
	GrantTypes       ScalarList[string] `gorm:"column:grantTypes;type:text[]" json:"grantTypes"`
	Scopes           string             `gorm:"column:scopes;not null" json:"scopes"`
	TokenAuthMethod  string             `gorm:"column:tokenAuthMethod;not null" json:"tokenAuthMethod"`
	HashedSecret     *string            `gorm:"column:hashedSecret" json:"hashedSecret"`
	RegistrationType string             `gorm:"column:registrationType;not null" json:"registrationType"`
	MetadataURL      *string            `gorm:"column:metadataUrl" json:"metadataUrl"`
	CreatedAt        DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Tokens           []RestrictedToken  `gorm:"foreignKey:ClientID;constraint:RestrictedToken_clientId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (OAuthClient) TableName() string {
	return "OAuthClient"
}

func (m *OAuthClient) BeforeCreate(tx *gorm.DB) error {
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

func (*OAuthClient) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type OAuthClientApproval struct {
	ID                 string    `gorm:"column:id;primaryKey" json:"id"`
	UserID             string    `gorm:"column:userId;uniqueIndex:OAuthClientApproval_userId_clientId_key,priority:1;not null" json:"userId"`
	ClientID           string    `gorm:"column:clientId;uniqueIndex:OAuthClientApproval_userId_clientId_key,priority:2;index:OAuthClientApproval_clientId_idx;not null" json:"clientId"`
	ApprovedAt         DateTime  `gorm:"column:approvedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"approvedAt"`
	RevokedAt          *DateTime `gorm:"column:revokedAt;type:timestamp(3)" json:"revokedAt"`
	ClientMetadataHash *string   `gorm:"column:clientMetadataHash" json:"clientMetadataHash"`
}

func (OAuthClientApproval) TableName() string {
	return "OAuthClientApproval"
}

func (m *OAuthClientApproval) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.ApprovedAt.IsZero() {
		m.ApprovedAt = DateTime{Time: now}
	}
	return nil
}

type OAuthGrantTeam struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	GrantID   string   `gorm:"column:grantId;uniqueIndex:OAuthGrantTeam_grantId_key;index:OAuthGrantTeam_grantId_idx;not null" json:"grantId"`
	TeamID    string   `gorm:"column:teamId;not null" json:"teamId"`
	UserID    string   `gorm:"column:userId;not null" json:"userId"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (OAuthGrantTeam) TableName() string {
	return "OAuthGrantTeam"
}

func (m *OAuthGrantTeam) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type OAuthRecord struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	Kind       string         `gorm:"column:kind;index:OAuthRecord_kind_idx;not null" json:"kind"`
	Payload    datatypes.JSON `gorm:"column:payload;not null" json:"payload"`
	GrantID    *string        `gorm:"column:grantId;index:OAuthRecord_grantId_idx" json:"grantId"`
	UserCode   *string        `gorm:"column:userCode;index:OAuthRecord_userCode_idx" json:"userCode"`
	UID        *string        `gorm:"column:uid;index:OAuthRecord_uid_idx" json:"uid"`
	ExpiresAt  *DateTime      `gorm:"column:expiresAt;index:OAuthRecord_expiresAt_idx;type:timestamp(3)" json:"expiresAt"`
	ConsumedAt *DateTime      `gorm:"column:consumedAt;type:timestamp(3)" json:"consumedAt"`
	CreatedAt  DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
}

func (OAuthRecord) TableName() string {
	return "OAuthRecord"
}

func (m *OAuthRecord) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DocumentRedactionJob struct {
	ID                string              `gorm:"column:id;primaryKey" json:"id"`
	DocumentID        string              `gorm:"column:documentId;index:DocumentRedactionJob_documentId_idx;not null" json:"documentId"`
	DocumentVersionID string              `gorm:"column:documentVersionId;index:DocumentRedactionJob_documentVersionId_idx;not null" json:"documentVersionId"`
	TeamID            string              `gorm:"column:teamId;index:DocumentRedactionJob_teamId_idx;index:DocumentRedactionJob_teamId_status_idx,priority:1;not null" json:"teamId"`
	CreatedByID       string              `gorm:"column:createdById;not null" json:"createdById"`
	Status            *string             `gorm:"column:status;index:DocumentRedactionJob_teamId_status_idx,priority:2;default:'PENDING';not null" json:"status"`
	CustomTerms       ScalarList[string]  `gorm:"column:customTerms;type:text[]" json:"customTerms"`
	Reasons           ScalarList[string]  `gorm:"column:reasons;type:text[]" json:"reasons"`
	ResultVersionID   *string             `gorm:"column:resultVersionId" json:"resultVersionId"`
	Error             *string             `gorm:"column:error" json:"error"`
	DetectRunID       *string             `gorm:"column:detectRunId" json:"detectRunId"`
	ApplyRunID        *string             `gorm:"column:applyRunId" json:"applyRunId"`
	CreatedAt         DateTime            `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime            `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Document          *Document           `gorm:"belongsTo"`
	DocumentVersion   *DocumentVersion    `gorm:"belongsTo"`
	Team              *Team               `gorm:"belongsTo"`
	CreatedBy         *User               `gorm:"foreignKey:CreatedByID;belongsTo"`
	Redactions        []DocumentRedaction `gorm:"foreignKey:JobID;constraint:DocumentRedaction_jobId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (DocumentRedactionJob) TableName() string {
	return "DocumentRedactionJob"
}

func (m *DocumentRedactionJob) BeforeCreate(tx *gorm.DB) error {
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

func (*DocumentRedactionJob) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type DocumentRedaction struct {
	ID          string                `gorm:"column:id;primaryKey" json:"id"`
	JobID       string                `gorm:"column:jobId;index:DocumentRedaction_jobId_idx;index:DocumentRedaction_jobId_pageNumber_idx,priority:1;index:DocumentRedaction_jobId_status_idx,priority:1;not null" json:"jobId"`
	PageNumber  int                   `gorm:"column:pageNumber;index:DocumentRedaction_jobId_pageNumber_idx,priority:2;size:32;not null" json:"pageNumber"`
	X           float64               `gorm:"column:x;type:double precision;not null" json:"x"`
	Y           float64               `gorm:"column:y;type:double precision;not null" json:"y"`
	Width       float64               `gorm:"column:width;type:double precision;not null" json:"width"`
	Height      float64               `gorm:"column:height;type:double precision;not null" json:"height"`
	TextPreview *string               `gorm:"column:textPreview" json:"textPreview"`
	Category    *string               `gorm:"column:category" json:"category"`
	Confidence  *string               `gorm:"column:confidence" json:"confidence"`
	Reason      *string               `gorm:"column:reason" json:"reason"`
	Source      *string               `gorm:"column:source;default:'AI';not null" json:"source"`
	Status      *string               `gorm:"column:status;index:DocumentRedaction_jobId_status_idx,priority:2;default:'PENDING';not null" json:"status"`
	CreatedAt   DateTime              `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime              `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Job         *DocumentRedactionJob `gorm:"foreignKey:JobID;belongsTo"`
}

func (DocumentRedaction) TableName() string {
	return "DocumentRedaction"
}

func (m *DocumentRedaction) BeforeCreate(tx *gorm.DB) error {
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

func (*DocumentRedaction) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ReferralInvite struct {
	ID     string  `gorm:"column:id;primaryKey" json:"id"`
	UserID string  `gorm:"column:userId;index:ReferralInvite_userId_idx;not null" json:"userId"`
	TeamID *string `gorm:"column:teamId;index:ReferralInvite_teamId_idx" json:"teamId"`
	// *
	// * Whose cap and discount this invite counts against — the fund's team, or the
	// * sender alone. One non-null key rather than a pair of nullable ones: on a
	// * nullable teamId every personal invite is distinct to Postgres, so the two
	// * keys guarded different scopes and a member could not invite a company both
	// * for their fund and for themselves.
	ScopeID      string   `gorm:"column:scopeId;uniqueIndex:ReferralInvite_scopeId_email_key,priority:1;not null" json:"scopeId"`
	Email        string   `gorm:"column:email;uniqueIndex:ReferralInvite_scopeId_email_key,priority:2;not null" json:"email"`
	ReferralLink *string  `gorm:"column:referralLink" json:"referralLink"`
	SentAt       DateTime `gorm:"column:sentAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"sentAt"`
	User         *User    `gorm:"belongsTo"`
	Team         *Team    `gorm:"belongsTo"`
}

func (ReferralInvite) TableName() string {
	return "ReferralInvite"
}

func (m *ReferralInvite) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.SentAt.IsZero() {
		m.SentAt = DateTime{Time: now}
	}
	return nil
}

type Account struct {
	ID                string  `gorm:"column:id;primaryKey" json:"id"`
	UserID            string  `gorm:"column:userId;not null" json:"userId"`
	Type              string  `gorm:"column:type;not null" json:"type"`
	Provider          string  `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key,priority:1;not null" json:"provider"`
	ProviderAccountID string  `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key,priority:2;not null" json:"providerAccountId"`
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
	ID                        string                    `gorm:"column:id;primaryKey" json:"id"`
	Name                      *string                   `gorm:"column:name" json:"name"`
	Email                     *string                   `gorm:"column:email;uniqueIndex:User_email_key" json:"email"`
	EmailVerified             *DateTime                 `gorm:"column:emailVerified;type:timestamp(3)" json:"emailVerified"`
	Image                     *string                   `gorm:"column:image" json:"image"`
	CreatedAt                 DateTime                  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	ContactID                 *string                   `gorm:"column:contactId" json:"contactId"`
	Plan                      *string                   `gorm:"column:plan;default:'free';not null" json:"plan"`
	StripeID                  *string                   `gorm:"column:stripeId;uniqueIndex:User_stripeId_key" json:"stripeId"`
	SubscriptionID            *string                   `gorm:"column:subscriptionId;uniqueIndex:User_subscriptionId_key" json:"subscriptionId"`
	StartsAt                  *DateTime                 `gorm:"column:startsAt;type:timestamp(3)" json:"startsAt"`
	EndsAt                    *DateTime                 `gorm:"column:endsAt;type:timestamp(3)" json:"endsAt"`
	ReferralTrack             *string                   `gorm:"column:referralTrack" json:"referralTrack"`
	ReferralTrackSelectedAt   *DateTime                 `gorm:"column:referralTrackSelectedAt;type:timestamp(3)" json:"referralTrackSelectedAt"`
	ReferralProfile           datatypes.JSON            `gorm:"column:referralProfile" json:"referralProfile"`
	Accounts                  []Account                 `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions                  []Session                 `gorm:"foreignKey:UserID;constraint:Session_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Documents                 []Document                `gorm:"foreignKey:OwnerID;constraint:Document_ownerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Teams                     []UserTeam                `gorm:"foreignKey:UserID;constraint:UserTeam_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserDatarooms             []UserDataroom            `gorm:"foreignKey:UserID;constraint:UserDataroom_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Domains                   []Domain                  `gorm:"foreignKey:UserID;constraint:Domain_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Chats                     []Chat                    `gorm:"foreignKey:UserID;constraint:Chat_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RestrictedTokens          []RestrictedToken         `gorm:"foreignKey:UserID;constraint:RestrictedToken_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ParticipatedConversations []ConversationParticipant `gorm:"foreignKey:UserID;constraint:ConversationParticipant_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Messages                  []Message                 `gorm:"foreignKey:UserID;constraint:Message_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedDataroomQuestions  []DataroomQuestion        `gorm:"foreignKey:CreatedByUserID;constraint:DataroomQuestion_createdByUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	PublishedFaqItems         []DataroomFaqItem         `gorm:"foreignKey:PublishedByUserID;constraint:DataroomFaqItem_publishedByUserId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAnnotations        []DocumentAnnotation      `gorm:"foreignKey:CreatedByID;constraint:DocumentAnnotation_createdById_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedRedactionJobs      []DocumentRedactionJob    `gorm:"foreignKey:CreatedByID;constraint:DocumentRedactionJob_createdById_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	InstalledIntegrations     []InstalledIntegration    `gorm:"foreignKey:UserID;constraint:InstalledIntegration_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Links                     []Link                    `gorm:"foreignKey:OwnerID;constraint:Link_ownerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	FrozenDatarooms           []Dataroom                `gorm:"foreignKey:FrozenBy;constraint:Dataroom_frozenBy_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedTasks              []Task                    `gorm:"foreignKey:CreatedByUserID;constraint:Task_createdByUserId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TaskActivities            []TaskActivity            `gorm:"foreignKey:UserID;constraint:TaskActivity_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	NotificationPreferences   []NotificationPreference  `gorm:"foreignKey:UserID;constraint:NotificationPreference_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReferralInvites           []ReferralInvite          `gorm:"foreignKey:UserID;constraint:ReferralInvite_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type VerificationToken struct {
	Identifier string   `gorm:"column:identifier;uniqueIndex:VerificationToken_identifier_token_key,priority:1;not null" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:VerificationToken_token_key;uniqueIndex:VerificationToken_identifier_token_key,priority:2" json:"token"`
	Expires    DateTime `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

type Domain struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	Slug        string   `gorm:"column:slug;uniqueIndex:Domain_slug_key;not null" json:"slug"`
	UserID      *string  `gorm:"column:userId;index:Domain_userId_idx" json:"userId"`
	TeamID      string   `gorm:"column:teamId;index:Domain_teamId_idx;not null" json:"teamId"`
	Verified    bool     `gorm:"column:verified;default:false;not null" json:"verified"`
	IsDefault   bool     `gorm:"column:isDefault;default:false;not null" json:"isDefault"`
	RedirectURL *string  `gorm:"column:redirectUrl" json:"redirectUrl"`
	LastChecked DateTime `gorm:"column:lastChecked;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"lastChecked"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User        *User    `gorm:"belongsTo"`
	Team        *Team    `gorm:"belongsTo"`
	Links       []Link   `gorm:"foreignKey:DomainID;constraint:Link_domainId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
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

type View struct {
	ID                   string               `gorm:"column:id;primaryKey" json:"id"`
	LinkID               string               `gorm:"column:linkId;index:View_linkId_idx;not null" json:"linkId"`
	DocumentID           *string              `gorm:"column:documentId;index:View_documentId_idx;index:View_viewerId_documentId_idx,priority:2;index:View_documentId_isArchived_idx,priority:1;index:View_documentId_viewedAt_idx,priority:1" json:"documentId"`
	DataroomID           *string              `gorm:"column:dataroomId;index:View_dataroomId_idx" json:"dataroomId"`
	DataroomViewID       *string              `gorm:"column:dataroomViewId;index:View_dataroomViewId_idx" json:"dataroomViewId"`
	ViewerEmail          *string              `gorm:"column:viewerEmail;index:View_viewerEmail_idx" json:"viewerEmail"`
	ViewerName           *string              `gorm:"column:viewerName" json:"viewerName"`
	Verified             bool                 `gorm:"column:verified;default:false;not null" json:"verified"`
	ViewedAt             DateTime             `gorm:"column:viewedAt;index:View_viewedAt_idx,sort:desc;index:View_documentId_viewedAt_idx,priority:2,sort:desc;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"viewedAt"`
	DownloadedAt         *DateTime            `gorm:"column:downloadedAt;type:timestamp(3)" json:"downloadedAt"`
	DownloadType         *string              `gorm:"column:downloadType" json:"downloadType"`
	DownloadMetadata     datatypes.JSON       `gorm:"column:downloadMetadata" json:"downloadMetadata"`
	ViewType             string               `gorm:"column:viewType;default:'DOCUMENT_VIEW';not null" json:"viewType"`
	ViewerID             *string              `gorm:"column:viewerId;index:View_viewerId_idx;index:View_viewerId_documentId_idx,priority:1" json:"viewerId"`
	GroupID              *string              `gorm:"column:groupId;index:View_groupId_idx" json:"groupId"`
	IsArchived           bool                 `gorm:"column:isArchived;index:View_documentId_isArchived_idx,priority:2;default:false;not null" json:"isArchived"`
	TeamID               *string              `gorm:"column:teamId;index:View_teamId_idx" json:"teamId"`
	Link                 *Link                `gorm:"belongsTo"`
	Document             *Document            `gorm:"belongsTo"`
	Dataroom             *Dataroom            `gorm:"belongsTo"`
	Viewer               *Viewer              `gorm:"belongsTo"`
	Group                *ViewerGroup         `gorm:"foreignKey:GroupID;belongsTo"`
	Team                 *Team                `gorm:"belongsTo"`
	Reactions            []Reaction           `gorm:"foreignKey:ViewID;constraint:Reaction_viewId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConversationViews    []ConversationView   `gorm:"foreignKey:ViewID;constraint:ConversationView_viewId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Messages             []Message            `gorm:"foreignKey:ViewID;constraint:Message_viewId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	InitialConversations []Conversation       `gorm:"foreignKey:InitialViewID;constraint:Conversation_initialViewId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	UploadedDocuments    []DocumentUpload     `gorm:"foreignKey:ViewID;constraint:DocumentUpload_viewId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TaskActivities       []TaskActivity       `gorm:"foreignKey:ViewID;constraint:TaskActivity_viewId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Chats                []Chat               `gorm:"foreignKey:ViewID;constraint:Chat_viewId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	FeedbackResponse     *FeedbackResponse    `gorm:"foreignKey:ViewID;constraint:FeedbackResponse_viewId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	AgreementResponse    *AgreementResponse   `gorm:"foreignKey:ViewID;constraint:AgreementResponse_viewId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomFieldResponse  *CustomFieldResponse `gorm:"foreignKey:ViewID;constraint:CustomFieldResponse_viewId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (View) TableName() string {
	return "View"
}

func (m *View) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.ViewedAt.IsZero() {
		m.ViewedAt = DateTime{Time: now}
	}
	return nil
}

type Viewer struct {
	ID                        string                       `gorm:"column:id;primaryKey" json:"id"`
	Email                     string                       `gorm:"column:email;uniqueIndex:Viewer_teamId_email_key,priority:2;not null" json:"email"`
	Verified                  bool                         `gorm:"column:verified;default:false;not null" json:"verified"`
	InvitedAt                 *DateTime                    `gorm:"column:invitedAt;type:timestamp(3)" json:"invitedAt"`
	NotificationPreferences   datatypes.JSON               `gorm:"column:notificationPreferences" json:"notificationPreferences"`
	DataroomID                *string                      `gorm:"column:dataroomId;index:Viewer_dataroomId_idx" json:"dataroomId"`
	TeamID                    string                       `gorm:"column:teamId;uniqueIndex:Viewer_teamId_email_key,priority:1;index:Viewer_teamId_idx;not null" json:"teamId"`
	CreatedAt                 DateTime                     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                 DateTime                     `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Dataroom                  *Dataroom                    `gorm:"belongsTo"`
	Team                      *Team                        `gorm:"belongsTo"`
	Views                     []View                       `gorm:"foreignKey:ViewerID;constraint:View_viewerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Groups                    []ViewerGroupMembership      `gorm:"foreignKey:ViewerID;constraint:ViewerGroupMembership_viewerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invitations               []ViewerInvitation           `gorm:"foreignKey:ViewerID;constraint:ViewerInvitation_viewerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ParticipatedConversations []ConversationParticipant    `gorm:"foreignKey:ViewerID;constraint:ConversationParticipant_viewerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Messages                  []Message                    `gorm:"foreignKey:ViewerID;constraint:Message_viewerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	RecipientConversations    []Conversation               `gorm:"foreignKey:RecipientViewerID;constraint:Conversation_recipientViewerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	QuestionAssignments       []DataroomQuestionAssignment `gorm:"foreignKey:ViewerID;constraint:DataroomQuestionAssignment_viewerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UploadedDocuments         []DocumentUpload             `gorm:"foreignKey:ViewerID;constraint:DocumentUpload_viewerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TaskAssignments           []TaskAssignment             `gorm:"foreignKey:ViewerID;constraint:TaskAssignment_viewerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskActivities            []TaskActivity               `gorm:"foreignKey:ViewerID;constraint:TaskActivity_viewerId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Chats                     []Chat                       `gorm:"foreignKey:ViewerID;constraint:Chat_viewerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Viewer) TableName() string {
	return "Viewer"
}

func (m *Viewer) BeforeCreate(tx *gorm.DB) error {
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

func (*Viewer) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Reaction struct {
	ID         string   `gorm:"column:id;primaryKey" json:"id"`
	ViewID     string   `gorm:"column:viewId;index:Reaction_viewId_idx;index:Reaction_viewId_type_idx,priority:1;not null" json:"viewId"`
	PageNumber int      `gorm:"column:pageNumber;size:32;not null" json:"pageNumber"`
	Type       string   `gorm:"column:type;index:Reaction_viewId_type_idx,priority:2;not null" json:"type"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	View       *View    `gorm:"belongsTo"`
}

func (Reaction) TableName() string {
	return "Reaction"
}

func (m *Reaction) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Invitation struct {
	Email       string             `gorm:"column:email;uniqueIndex:Invitation_email_teamId_key,priority:1;not null" json:"email"`
	Expires     DateTime           `gorm:"column:expires;type:timestamp(3);not null" json:"expires"`
	TeamID      string             `gorm:"column:teamId;uniqueIndex:Invitation_email_teamId_key,priority:2;not null" json:"teamId"`
	CreatedAt   DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Token       string             `gorm:"column:token;primaryKey;uniqueIndex:Invitation_token_key" json:"token"`
	Role        string             `gorm:"column:role;default:'MEMBER';not null" json:"role"`
	DataroomIds ScalarList[string] `gorm:"column:dataroomIds;type:text[]" json:"dataroomIds"`
	Team        *Team              `gorm:"belongsTo"`
}

func (Invitation) TableName() string {
	return "Invitation"
}

func (m *Invitation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type SentEmail struct {
	ID         string   `gorm:"column:id;primaryKey" json:"id"`
	Type       string   `gorm:"column:type;not null" json:"type"`
	Recipient  string   `gorm:"column:recipient;not null" json:"recipient"`
	Marketing  bool     `gorm:"column:marketing;default:false;not null" json:"marketing"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	TeamID     string   `gorm:"column:teamId;index:SentEmail_teamId_idx;not null" json:"teamId"`
	DomainSlug *string  `gorm:"column:domainSlug" json:"domainSlug"`
	Team       *Team    `gorm:"belongsTo"`
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

type Chat struct {
	ID            string        `gorm:"column:id;primaryKey" json:"id"`
	Title         *string       `gorm:"column:title" json:"title"`
	TeamID        string        `gorm:"column:teamId;index:Chat_teamId_idx;not null" json:"teamId"`
	DocumentID    *string       `gorm:"column:documentId;index:Chat_documentId_idx" json:"documentId"`
	DataroomID    *string       `gorm:"column:dataroomId;index:Chat_dataroomId_idx" json:"dataroomId"`
	LinkID        *string       `gorm:"column:linkId;index:Chat_linkId_idx" json:"linkId"`
	ViewID        *string       `gorm:"column:viewId;index:Chat_viewId_idx" json:"viewId"`
	UserID        *string       `gorm:"column:userId;index:Chat_userId_idx" json:"userId"`
	ViewerID      *string       `gorm:"column:viewerId;index:Chat_viewerId_idx" json:"viewerId"`
	VectorStoreID *string       `gorm:"column:vectorStoreId" json:"vectorStoreId"`
	CreatedAt     DateTime      `gorm:"column:createdAt;index:Chat_createdAt_idx,sort:desc;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime      `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	LastMessageAt *DateTime     `gorm:"column:lastMessageAt;type:timestamp(3)" json:"lastMessageAt"`
	Team          *Team         `gorm:"belongsTo"`
	Document      *Document     `gorm:"belongsTo"`
	Dataroom      *Dataroom     `gorm:"belongsTo"`
	Link          *Link         `gorm:"belongsTo"`
	View          *View         `gorm:"belongsTo"`
	User          *User         `gorm:"belongsTo"`
	Viewer        *Viewer       `gorm:"belongsTo"`
	Messages      []ChatMessage `gorm:"foreignKey:ChatID;constraint:ChatMessage_chatId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Chat) TableName() string {
	return "Chat"
}

func (m *Chat) BeforeCreate(tx *gorm.DB) error {
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

func (*Chat) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ChatMessage struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	ChatID    string         `gorm:"column:chatId;index:ChatMessage_chatId_idx;index:ChatMessage_chatId_createdAt_idx,priority:1;not null" json:"chatId"`
	Role      string         `gorm:"column:role;not null" json:"role"`
	Content   string         `gorm:"column:content;type:text;not null" json:"content"`
	Metadata  datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	CreatedAt DateTime       `gorm:"column:createdAt;index:ChatMessage_chatId_createdAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Chat      *Chat          `gorm:"belongsTo"`
}

func (ChatMessage) TableName() string {
	return "ChatMessage"
}

func (m *ChatMessage) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Feedback struct {
	ID        string             `gorm:"column:id;primaryKey" json:"id"`
	LinkID    string             `gorm:"column:linkId;uniqueIndex:Feedback_linkId_key;index:Feedback_linkId_idx;not null" json:"linkId"`
	Data      datatypes.JSON     `gorm:"column:data;not null" json:"data"`
	CreatedAt DateTime           `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime           `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Link      *Link              `gorm:"belongsTo"`
	Responses []FeedbackResponse `gorm:"foreignKey:FeedbackID;constraint:FeedbackResponse_feedbackId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Feedback) TableName() string {
	return "Feedback"
}

func (m *Feedback) BeforeCreate(tx *gorm.DB) error {
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

func (*Feedback) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type FeedbackResponse struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	FeedbackID string         `gorm:"column:feedbackId;index:FeedbackResponse_feedbackId_idx;not null" json:"feedbackId"`
	Data       datatypes.JSON `gorm:"column:data;not null" json:"data"`
	ViewID     string         `gorm:"column:viewId;uniqueIndex:FeedbackResponse_viewId_key;index:FeedbackResponse_viewId_idx;not null" json:"viewId"`
	CreatedAt  DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Feedback   *Feedback      `gorm:"belongsTo"`
	View       *View          `gorm:"belongsTo"`
}

func (FeedbackResponse) TableName() string {
	return "FeedbackResponse"
}

func (m *FeedbackResponse) BeforeCreate(tx *gorm.DB) error {
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

func (*FeedbackResponse) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Agreement struct {
	ID                string              `gorm:"column:id;primaryKey" json:"id"`
	Name              string              `gorm:"column:name;not null" json:"name"`
	Content           string              `gorm:"column:content;not null" json:"content"`
	ContentType       *string             `gorm:"column:contentType;default:'LINK';not null" json:"contentType"`
	SigningProvider   *string             `gorm:"column:signingProvider;default:'LEGACY';not null" json:"signingProvider"`
	SigningExternalID *string             `gorm:"column:signingExternalId;uniqueIndex:Agreement_signingExternalId_key" json:"signingExternalId"`
	SigningEnvelopeID *string             `gorm:"column:signingEnvelopeId" json:"signingEnvelopeId"`
	SigningTemplateID *string             `gorm:"column:signingTemplateId" json:"signingTemplateId"`
	RequireName       *bool               `gorm:"column:requireName;default:true;not null" json:"requireName"`
	TeamID            string              `gorm:"column:teamId;index:Agreement_teamId_idx;not null" json:"teamId"`
	CreatedAt         DateTime            `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime            `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	DeletedAt         *DateTime           `gorm:"column:deletedAt;type:timestamp(3)" json:"deletedAt"`
	DeletedBy         *string             `gorm:"column:deletedBy" json:"deletedBy"`
	Team              *Team               `gorm:"belongsTo"`
	Links             []Link              `gorm:"foreignKey:AgreementID;constraint:Link_agreementId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Responses         []AgreementResponse `gorm:"foreignKey:AgreementID;constraint:AgreementResponse_agreementId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Agreement) TableName() string {
	return "Agreement"
}

func (m *Agreement) BeforeCreate(tx *gorm.DB) error {
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

func (*Agreement) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AgreementResponse struct {
	ID                    string     `gorm:"column:id;primaryKey" json:"id"`
	AgreementID           string     `gorm:"column:agreementId;index:AgreementResponse_agreementId_idx;not null" json:"agreementId"`
	ViewID                *string    `gorm:"column:viewId;uniqueIndex:AgreementResponse_viewId_key;index:AgreementResponse_viewId_idx" json:"viewId"`
	LinkID                *string    `gorm:"column:linkId;index:AgreementResponse_linkId_idx" json:"linkId"`
	SignerEmail           *string    `gorm:"column:signerEmail;index:AgreementResponse_signerEmail_idx" json:"signerEmail"`
	SignerName            *string    `gorm:"column:signerName" json:"signerName"`
	SigningStatus         *string    `gorm:"column:signingStatus;default:'PENDING';not null" json:"signingStatus"`
	SigningExternalID     *string    `gorm:"column:signingExternalId;uniqueIndex:AgreementResponse_signingExternalId_key" json:"signingExternalId"`
	SigningEnvelopeID     *string    `gorm:"column:signingEnvelopeId" json:"signingEnvelopeId"`
	SigningDocumentID     *int       `gorm:"column:signingDocumentId;size:32" json:"signingDocumentId"`
	SignedAt              *DateTime  `gorm:"column:signedAt;type:timestamp(3)" json:"signedAt"`
	CompletedAt           *DateTime  `gorm:"column:completedAt;type:timestamp(3)" json:"completedAt"`
	SignedFileKey         *string    `gorm:"column:signedFileKey" json:"signedFileKey"`
	SignedFileName        *string    `gorm:"column:signedFileName" json:"signedFileName"`
	SignedFileStorageType *string    `gorm:"column:signedFileStorageType" json:"signedFileStorageType"`
	CreatedAt             DateTime   `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt             DateTime   `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Agreement             *Agreement `gorm:"belongsTo"`
	View                  *View      `gorm:"belongsTo"`
}

func (AgreementResponse) TableName() string {
	return "AgreementResponse"
}

func (m *AgreementResponse) BeforeCreate(tx *gorm.DB) error {
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

func (*AgreementResponse) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type IncomingWebhook struct {
	ID                  string    `gorm:"column:id;primaryKey" json:"id"`
	ExternalID          string    `gorm:"column:externalId;uniqueIndex:IncomingWebhook_externalId_key;not null" json:"externalId"`
	Name                string    `gorm:"column:name;not null" json:"name"`
	Secret              *string   `gorm:"column:secret" json:"secret"`
	Source              *string   `gorm:"column:source" json:"source"`
	Actions             *string   `gorm:"column:actions" json:"actions"`
	ConsecutiveFailures int       `gorm:"column:consecutiveFailures;size:32;default:0;not null" json:"consecutiveFailures"`
	LastFailedAt        *DateTime `gorm:"column:lastFailedAt;type:timestamp(3)" json:"lastFailedAt"`
	DisabledAt          *DateTime `gorm:"column:disabledAt;type:timestamp(3)" json:"disabledAt"`
	TeamID              string    `gorm:"column:teamId;index:IncomingWebhook_teamId_idx;not null" json:"teamId"`
	CreatedAt           DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt           DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team                *Team     `gorm:"belongsTo"`
}

func (IncomingWebhook) TableName() string {
	return "IncomingWebhook"
}

func (m *IncomingWebhook) BeforeCreate(tx *gorm.DB) error {
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

func (*IncomingWebhook) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type RestrictedToken struct {
	ID          string       `gorm:"column:id;primaryKey" json:"id"`
	Name        string       `gorm:"column:name;not null" json:"name"`
	HashedKey   string       `gorm:"column:hashedKey;uniqueIndex:RestrictedToken_hashedKey_key;index:RestrictedToken_hashedKey_idx;not null" json:"hashedKey"`
	PartialKey  string       `gorm:"column:partialKey;not null" json:"partialKey"`
	Scopes      *string      `gorm:"column:scopes" json:"scopes"`
	Mode        *string      `gorm:"column:mode;default:'live';not null" json:"mode"`
	Source      *string      `gorm:"column:source;default:'dashboard';not null" json:"source"`
	SubjectType *string      `gorm:"column:subjectType;default:'user';not null" json:"subjectType"`
	ClientID    *string      `gorm:"column:clientId;index:RestrictedToken_clientId_idx" json:"clientId"`
	Expires     *DateTime    `gorm:"column:expires;type:timestamp(3)" json:"expires"`
	LastUsed    *DateTime    `gorm:"column:lastUsed;type:timestamp(3)" json:"lastUsed"`
	RateLimit   *int         `gorm:"column:rateLimit;size:32;default:60;not null" json:"rateLimit"`
	UserID      string       `gorm:"column:userId;index:RestrictedToken_userId_idx;not null" json:"userId"`
	TeamID      string       `gorm:"column:teamId;index:RestrictedToken_teamId_idx;not null" json:"teamId"`
	CreatedAt   DateTime     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime     `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	User        *User        `gorm:"belongsTo"`
	Team        *Team        `gorm:"belongsTo"`
	OauthClient *OAuthClient `gorm:"foreignKey:ClientID;belongsTo"`
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

type Webhook struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	PID       string         `gorm:"column:pId;uniqueIndex:Webhook_pId_key;not null" json:"pId"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	URL       string         `gorm:"column:url;not null" json:"url"`
	Secret    string         `gorm:"column:secret;not null" json:"secret"`
	Triggers  datatypes.JSON `gorm:"column:triggers;not null" json:"triggers"`
	TeamID    string         `gorm:"column:teamId;index:Webhook_teamId_idx;not null" json:"teamId"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team      *Team          `gorm:"belongsTo"`
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

type YearInReview struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	TeamID        string         `gorm:"column:teamId;index:YearInReview_teamId_idx;not null" json:"teamId"`
	Status        *string        `gorm:"column:status;index:YearInReview_status_attempts_idx,priority:1;default:'pending';not null" json:"status"`
	Attempts      int            `gorm:"column:attempts;index:YearInReview_status_attempts_idx,priority:2;size:32;default:0;not null" json:"attempts"`
	LastAttempted *DateTime      `gorm:"column:lastAttempted;type:timestamp(3)" json:"lastAttempted"`
	Error         *string        `gorm:"column:error" json:"error"`
	Stats         datatypes.JSON `gorm:"column:stats;not null" json:"stats"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
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
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*YearInReview) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Tag struct {
	ID          string    `gorm:"column:id;primaryKey;index:Tag_id_idx" json:"id"`
	Name        string    `gorm:"column:name;uniqueIndex:Tag_teamId_name_key,priority:2;index:Tag_name_idx;not null" json:"name"`
	Color       string    `gorm:"column:color;not null" json:"color"`
	Description *string   `gorm:"column:description" json:"description"`
	TeamID      string    `gorm:"column:teamId;uniqueIndex:Tag_teamId_name_key,priority:1;index:Tag_teamId_idx;not null" json:"teamId"`
	CreatedBy   *string   `gorm:"column:createdBy" json:"createdBy"`
	CreatedAt   DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Team        *Team     `gorm:"belongsTo"`
	Items       []TagItem `gorm:"foreignKey:TagID;constraint:TagItem_tagId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type TagItem struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	TagID      string    `gorm:"column:tagId;index:TagItem_tagId_linkId_idx,priority:1;index:TagItem_tagId_documentId_idx,priority:1;index:TagItem_tagId_dataroomId_idx,priority:1;not null" json:"tagId"`
	ItemType   string    `gorm:"column:itemType;not null" json:"itemType"`
	LinkID     *string   `gorm:"column:linkId;index:TagItem_tagId_linkId_idx,priority:2" json:"linkId"`
	DocumentID *string   `gorm:"column:documentId;index:TagItem_tagId_documentId_idx,priority:2" json:"documentId"`
	DataroomID *string   `gorm:"column:dataroomId;index:TagItem_tagId_dataroomId_idx,priority:2" json:"dataroomId"`
	TaggedBy   *string   `gorm:"column:taggedBy" json:"taggedBy"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime  `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Tag        *Tag      `gorm:"belongsTo"`
	Link       *Link     `gorm:"belongsTo"`
	Document   *Document `gorm:"belongsTo"`
	Dataroom   *Dataroom `gorm:"belongsTo"`
}

func (TagItem) TableName() string {
	return "TagItem"
}

func (m *TagItem) BeforeCreate(tx *gorm.DB) error {
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

func (*TagItem) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ViewerInvitation struct {
	ID            string       `gorm:"column:id;primaryKey" json:"id"`
	ViewerID      string       `gorm:"column:viewerId;index:ViewerInvitation_viewerId_idx;not null" json:"viewerId"`
	LinkID        string       `gorm:"column:linkId;index:ViewerInvitation_linkId_idx;not null" json:"linkId"`
	GroupID       *string      `gorm:"column:groupId;index:ViewerInvitation_groupId_idx" json:"groupId"`
	InvitedBy     string       `gorm:"column:invitedBy;not null" json:"invitedBy"`
	CustomMessage *string      `gorm:"column:customMessage" json:"customMessage"`
	SentAt        DateTime     `gorm:"column:sentAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"sentAt"`
	Status        string       `gorm:"column:status;default:'SENT';not null" json:"status"`
	CreatedAt     DateTime     `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Viewer        *Viewer      `gorm:"belongsTo"`
	Link          *Link        `gorm:"belongsTo"`
	Group         *ViewerGroup `gorm:"foreignKey:GroupID;belongsTo"`
}

func (ViewerInvitation) TableName() string {
	return "ViewerInvitation"
}

func (m *ViewerInvitation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.SentAt.IsZero() {
		m.SentAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Team struct {
	ID                       string                   `gorm:"column:id;primaryKey" json:"id"`
	Name                     string                   `gorm:"column:name;not null" json:"name"`
	Slug                     *string                  `gorm:"column:slug;uniqueIndex:Team_slug_key" json:"slug"`
	DefaultBrandID           *string                  `gorm:"column:defaultBrandId;index:Team_defaultBrandId_idx" json:"defaultBrandId"`
	Plan                     *string                  `gorm:"column:plan;default:'free';not null" json:"plan"`
	StripeID                 *string                  `gorm:"column:stripeId;uniqueIndex:Team_stripeId_key" json:"stripeId"`
	SubscriptionID           *string                  `gorm:"column:subscriptionId;uniqueIndex:Team_subscriptionId_key" json:"subscriptionId"`
	StartsAt                 *DateTime                `gorm:"column:startsAt;type:timestamp(3)" json:"startsAt"`
	EndsAt                   *DateTime                `gorm:"column:endsAt;type:timestamp(3)" json:"endsAt"`
	PausedAt                 *DateTime                `gorm:"column:pausedAt;type:timestamp(3)" json:"pausedAt"`
	PauseStartsAt            *DateTime                `gorm:"column:pauseStartsAt;type:timestamp(3)" json:"pauseStartsAt"`
	PauseEndsAt              *DateTime                `gorm:"column:pauseEndsAt;type:timestamp(3)" json:"pauseEndsAt"`
	CancelledAt              *DateTime                `gorm:"column:cancelledAt;type:timestamp(3)" json:"cancelledAt"`
	TrialStartedAt           *DateTime                `gorm:"column:trialStartedAt;type:timestamp(3)" json:"trialStartedAt"`
	TrialEndsAt              *DateTime                `gorm:"column:trialEndsAt;type:timestamp(3)" json:"trialEndsAt"`
	TrialReminderRunID       *string                  `gorm:"column:trialReminderRunId" json:"trialReminderRunId"`
	TrialExpiredRunID        *string                  `gorm:"column:trialExpiredRunId" json:"trialExpiredRunId"`
	Limits                   datatypes.JSON           `gorm:"column:limits" json:"limits"`
	EnableExcelAdvancedMode  bool                     `gorm:"column:enableExcelAdvancedMode;default:false;not null" json:"enableExcelAdvancedMode"`
	ReplicateDataroomFolders *bool                    `gorm:"column:replicateDataroomFolders;default:true;not null" json:"replicateDataroomFolders"`
	Timezone                 *string                  `gorm:"column:timezone;default:'Etc/UTC';not null" json:"timezone"`
	SsoEnabled               bool                     `gorm:"column:ssoEnabled;default:false;not null" json:"ssoEnabled"`
	SsoEmailDomain           *string                  `gorm:"column:ssoEmailDomain;uniqueIndex:Team_ssoEmailDomain_key" json:"ssoEmailDomain"`
	SsoEnforcedAt            *DateTime                `gorm:"column:ssoEnforcedAt;type:timestamp(3)" json:"ssoEnforcedAt"`
	SurveyData               datatypes.JSON           `gorm:"column:surveyData" json:"surveyData"`
	AgentsEnabled            bool                     `gorm:"column:agentsEnabled;default:false;not null" json:"agentsEnabled"`
	VectorStoreID            *string                  `gorm:"column:vectorStoreId" json:"vectorStoreId"`
	CreatedAt                DateTime                 `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt                DateTime                 `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	IgnoredDomains           ScalarList[string]       `gorm:"column:ignoredDomains;type:text[]" json:"ignoredDomains"`
	GlobalBlockList          ScalarList[string]       `gorm:"column:globalBlockList;type:text[]" json:"globalBlockList"`
	ReferralTrack            *string                  `gorm:"column:referralTrack" json:"referralTrack"`
	ReferralTrackSelectedAt  *DateTime                `gorm:"column:referralTrackSelectedAt;type:timestamp(3)" json:"referralTrackSelectedAt"`
	ReferralProfile          datatypes.JSON           `gorm:"column:referralProfile" json:"referralProfile"`
	DefaultBrand             *Brand                   `gorm:"foreignKey:DefaultBrandID;belongsTo"`
	Users                    []UserTeam               `gorm:"foreignKey:TeamID;constraint:UserTeam_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserDatarooms            []UserDataroom           `gorm:"foreignKey:TeamID;constraint:UserDataroom_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Documents                []Document               `gorm:"foreignKey:TeamID;constraint:Document_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Folders                  []Folder                 `gorm:"foreignKey:TeamID;constraint:Folder_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Domains                  []Domain                 `gorm:"foreignKey:TeamID;constraint:Domain_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invitations              []Invitation             `gorm:"foreignKey:TeamID;constraint:Invitation_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SentEmails               []SentEmail              `gorm:"foreignKey:TeamID;constraint:SentEmail_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Brands                   []Brand                  `gorm:"foreignKey:TeamID;constraint:Brand_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Datarooms                []Dataroom               `gorm:"foreignKey:TeamID;constraint:Dataroom_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Agreements               []Agreement              `gorm:"foreignKey:TeamID;constraint:Agreement_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ViewerGroups             []ViewerGroup            `gorm:"foreignKey:TeamID;constraint:ViewerGroup_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	VisitorGroups            []VisitorGroup           `gorm:"foreignKey:TeamID;constraint:VisitorGroup_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Viewers                  []Viewer                 `gorm:"foreignKey:TeamID;constraint:Viewer_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PermissionGroups         []PermissionGroup        `gorm:"foreignKey:TeamID;constraint:PermissionGroup_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	LinkPresets              []LinkPreset             `gorm:"foreignKey:TeamID;constraint:LinkPreset_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	IncomingWebhooks         []IncomingWebhook        `gorm:"foreignKey:TeamID;constraint:IncomingWebhook_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RestrictedTokens         []RestrictedToken        `gorm:"foreignKey:TeamID;constraint:RestrictedToken_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Webhooks                 []Webhook                `gorm:"foreignKey:TeamID;constraint:Webhook_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Conversations            []Conversation           `gorm:"foreignKey:TeamID;constraint:Conversation_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DataroomQuestions        []DataroomQuestion       `gorm:"foreignKey:TeamID;constraint:DataroomQuestion_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DataroomFaqItems         []DataroomFaqItem        `gorm:"foreignKey:TeamID;constraint:DataroomFaqItem_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UploadedDocuments        []DocumentUpload         `gorm:"foreignKey:TeamID;constraint:DocumentUpload_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tag                      []Tag                    `gorm:"foreignKey:TeamID;constraint:Tag_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Annotations              []DocumentAnnotation     `gorm:"foreignKey:TeamID;constraint:DocumentAnnotation_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	RedactionJobs            []DocumentRedactionJob   `gorm:"foreignKey:TeamID;constraint:DocumentRedactionJob_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	InstalledIntegrations    []InstalledIntegration   `gorm:"foreignKey:TeamID;constraint:InstalledIntegration_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Links                    []Link                   `gorm:"foreignKey:TeamID;constraint:Link_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Views                    []View                   `gorm:"foreignKey:TeamID;constraint:View_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Workflows                []Workflow               `gorm:"foreignKey:TeamID;constraint:Workflow_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskLists                []TaskList               `gorm:"foreignKey:TeamID;constraint:TaskList_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tasks                    []Task                   `gorm:"foreignKey:TeamID;constraint:Task_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Chats                    []Chat                   `gorm:"foreignKey:TeamID;constraint:Chat_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	NotificationPreferences  []NotificationPreference `gorm:"foreignKey:TeamID;constraint:NotificationPreference_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReferralInvites          []ReferralInvite         `gorm:"foreignKey:TeamID;constraint:ReferralInvite_teamId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type UserDataroom struct {
	UserID     string    `gorm:"column:userId;primaryKey;index:UserDataroom_userId_idx" json:"userId"`
	TeamID     string    `gorm:"column:teamId;index:UserDataroom_teamId_idx;not null" json:"teamId"`
	DataroomID string    `gorm:"column:dataroomId;primaryKey;index:UserDataroom_dataroomId_idx" json:"dataroomId"`
	CreatedAt  DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	User       *User     `gorm:"belongsTo"`
	Team       *Team     `gorm:"belongsTo"`
	Dataroom   *Dataroom `gorm:"belongsTo"`
}

func (UserDataroom) TableName() string {
	return "UserDataroom"
}

func (m *UserDataroom) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type UserTeam struct {
	Role                    string         `gorm:"column:role;default:'MEMBER';not null" json:"role"`
	Status                  *string        `gorm:"column:status;default:'ACTIVE';not null" json:"status"`
	UserID                  string         `gorm:"column:userId;primaryKey;index:UserTeam_userId_idx" json:"userId"`
	TeamID                  string         `gorm:"column:teamId;primaryKey;index:UserTeam_teamId_idx" json:"teamId"`
	BlockedAt               *DateTime      `gorm:"column:blockedAt;type:timestamp(3)" json:"blockedAt"`
	NotificationPreferences datatypes.JSON `gorm:"column:notificationPreferences" json:"notificationPreferences"`
	User                    *User          `gorm:"belongsTo"`
	Team                    *Team          `gorm:"belongsTo"`
}

func (UserTeam) TableName() string {
	return "UserTeam"
}

type Workflow struct {
	ID          string              `gorm:"column:id;primaryKey" json:"id"`
	Name        string              `gorm:"column:name;not null" json:"name"`
	Description *string             `gorm:"column:description" json:"description"`
	EntryLinkID string              `gorm:"column:entryLinkId;uniqueIndex:Workflow_entryLinkId_key;index:Workflow_entryLinkId_idx;not null" json:"entryLinkId"`
	TeamID      string              `gorm:"column:teamId;index:Workflow_teamId_idx;not null" json:"teamId"`
	IsActive    *bool               `gorm:"column:isActive;index:Workflow_isActive_idx;default:true;not null" json:"isActive"`
	CreatedAt   DateTime            `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime            `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	EntryLink   *Link               `gorm:"foreignKey:EntryLinkID;belongsTo"`
	Team        *Team               `gorm:"belongsTo"`
	Steps       []WorkflowStep      `gorm:"foreignKey:WorkflowID;constraint:WorkflowStep_workflowId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Executions  []WorkflowExecution `gorm:"foreignKey:WorkflowID;constraint:WorkflowExecution_workflowId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Workflow) TableName() string {
	return "Workflow"
}

func (m *Workflow) BeforeCreate(tx *gorm.DB) error {
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

func (*Workflow) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkflowStep struct {
	ID            string            `gorm:"column:id;primaryKey" json:"id"`
	WorkflowID    string            `gorm:"column:workflowId;uniqueIndex:WorkflowStep_workflowId_stepOrder_key,priority:1;index:WorkflowStep_workflowId_idx;not null" json:"workflowId"`
	Name          string            `gorm:"column:name;not null" json:"name"`
	StepOrder     int               `gorm:"column:stepOrder;uniqueIndex:WorkflowStep_workflowId_stepOrder_key,priority:2;size:32;not null" json:"stepOrder"`
	StepType      string            `gorm:"column:stepType;default:'ROUTER';not null" json:"stepType"`
	Conditions    datatypes.JSON    `gorm:"column:conditions;not null" json:"conditions"`
	Actions       datatypes.JSON    `gorm:"column:actions;not null" json:"actions"`
	CreatedAt     DateTime          `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime          `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	Workflow      *Workflow         `gorm:"belongsTo"`
	ExecutionLogs []WorkflowStepLog `gorm:"foreignKey:WorkflowStepID;constraint:WorkflowStepLog_workflowStepId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (WorkflowStep) TableName() string {
	return "WorkflowStep"
}

func (m *WorkflowStep) BeforeCreate(tx *gorm.DB) error {
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

func (*WorkflowStep) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type WorkflowExecution struct {
	ID           string            `gorm:"column:id;primaryKey" json:"id"`
	WorkflowID   string            `gorm:"column:workflowId;index:WorkflowExecution_workflowId_startedAt_idx,priority:1;not null" json:"workflowId"`
	VisitorEmail *string           `gorm:"column:visitorEmail;index:WorkflowExecution_visitorEmail_idx" json:"visitorEmail"`
	VisitorIP    *string           `gorm:"column:visitorIp" json:"visitorIp"`
	Status       string            `gorm:"column:status;index:WorkflowExecution_status_idx;not null" json:"status"`
	StartedAt    DateTime          `gorm:"column:startedAt;index:WorkflowExecution_workflowId_startedAt_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"startedAt"`
	CompletedAt  *DateTime         `gorm:"column:completedAt;type:timestamp(3)" json:"completedAt"`
	Result       datatypes.JSON    `gorm:"column:result" json:"result"`
	Metadata     datatypes.JSON    `gorm:"column:metadata" json:"metadata"`
	Workflow     *Workflow         `gorm:"belongsTo"`
	StepLogs     []WorkflowStepLog `gorm:"foreignKey:ExecutionID;constraint:WorkflowStepLog_executionId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (WorkflowExecution) TableName() string {
	return "WorkflowExecution"
}

func (m *WorkflowExecution) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.StartedAt.IsZero() {
		m.StartedAt = DateTime{Time: now}
	}
	return nil
}

type WorkflowStepLog struct {
	ID                string             `gorm:"column:id;primaryKey" json:"id"`
	ExecutionID       string             `gorm:"column:executionId;index:WorkflowStepLog_executionId_idx;not null" json:"executionId"`
	WorkflowStepID    string             `gorm:"column:workflowStepId;index:WorkflowStepLog_workflowStepId_idx;not null" json:"workflowStepId"`
	ConditionsMatched bool               `gorm:"column:conditionsMatched;not null" json:"conditionsMatched"`
	ConditionResults  datatypes.JSON     `gorm:"column:conditionResults" json:"conditionResults"`
	ActionsExecuted   datatypes.JSON     `gorm:"column:actionsExecuted" json:"actionsExecuted"`
	ExecutedAt        DateTime           `gorm:"column:executedAt;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"executedAt"`
	Duration          *int               `gorm:"column:duration;size:32" json:"duration"`
	Error             *string            `gorm:"column:error" json:"error"`
	Execution         *WorkflowExecution `gorm:"foreignKey:ExecutionID;belongsTo"`
	Step              *WorkflowStep      `gorm:"foreignKey:WorkflowStepID;belongsTo"`
}

func (WorkflowStepLog) TableName() string {
	return "WorkflowStepLog"
}

func (m *WorkflowStepLog) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.ExecutedAt.IsZero() {
		m.ExecutedAt = DateTime{Time: now}
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
