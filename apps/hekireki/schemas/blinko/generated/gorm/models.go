package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Accounts struct {
	ID               int                 `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name             string              `gorm:"column:name;default:'';not null" json:"name"`
	Nickname         string              `gorm:"column:nickname;default:'';not null" json:"nickname"`
	Password         string              `gorm:"column:password;default:'';not null" json:"password"`
	Image            string              `gorm:"column:image;default:'';not null" json:"image"`
	APIToken         string              `gorm:"column:apiToken;default:'';not null" json:"apiToken"`
	Description      string              `gorm:"column:description;type:text;default:'';not null" json:"description"`
	Note             int                 `gorm:"column:note;size:32;default:0;not null" json:"note"`
	Role             string              `gorm:"column:role;default:'';not null" json:"role"`
	LoginType        string              `gorm:"column:loginType;default:'';not null" json:"loginType"`
	LinkAccountID    *int                `gorm:"column:linkAccountId;size:32" json:"linkAccountId"`
	CreatedAt        DateTime            `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt        DateTime            `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Notes            []Notes             `gorm:"foreignKey:AccountID;constraint:notes_accountId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Configs          []Config            `gorm:"foreignKey:UserID;constraint:config_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tags             []Tag               `gorm:"foreignKey:AccountID;constraint:tag_accountId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments         []Comments          `gorm:"foreignKey:AccountID;constraint:comments_accountId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Attachments      []Attachments       `gorm:"foreignKey:AccountID;constraint:attachments_accountId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Follows          []Follows           `gorm:"foreignKey:AccountID;constraint:follows_accountId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Notifications    []Notifications     `gorm:"foreignKey:AccountID;constraint:notifications_accountId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Conversations    []Conversation      `gorm:"foreignKey:AccountID;constraint:conversation_accountId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SharedNotes      []NoteInternalShare `gorm:"foreignKey:AccountID;constraint:noteInternalShare_accountId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	AiScheduledTasks []AiScheduledTask   `gorm:"foreignKey:AccountID;constraint:aiScheduledTask_accountId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Accounts) TableName() string {
	return "accounts"
}

func (m *Accounts) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Accounts) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Attachments struct {
	ID            int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	IsShare       bool           `gorm:"column:isShare;default:false;not null" json:"isShare"`
	SharePassword string         `gorm:"column:sharePassword;default:'';not null" json:"sharePassword"`
	Name          string         `gorm:"column:name;default:'';not null" json:"name"`
	Path          string         `gorm:"column:path;default:'';not null" json:"path"`
	Size          float64        `gorm:"column:size;type:decimal;default:0;not null" json:"size"`
	Type          string         `gorm:"column:type;default:'';not null" json:"type"`
	NoteID        *int           `gorm:"column:noteId;size:32" json:"noteId"`
	AccountID     *int           `gorm:"column:accountId;size:32" json:"accountId"`
	SortOrder     int            `gorm:"column:sortOrder;size:32;default:0;not null" json:"sortOrder"`
	CreatedAt     DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt     DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	PerfixPath    *string        `gorm:"column:perfixPath;default:''" json:"perfixPath"`
	Depth         *int           `gorm:"column:depth;size:32" json:"depth"`
	Metadata      datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	Note          *Notes         `gorm:"foreignKey:NoteID;belongsTo"`
	Account       *Accounts      `gorm:"foreignKey:AccountID;belongsTo"`
}

func (Attachments) TableName() string {
	return "attachments"
}

func (m *Attachments) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Attachments) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Config struct {
	ID     int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Key    string         `gorm:"column:key;default:'';not null" json:"key"`
	Config datatypes.JSON `gorm:"column:config" json:"config"`
	UserID *int           `gorm:"column:userId;size:32" json:"userId"`
	User   *Accounts      `gorm:"foreignKey:UserID;belongsTo"`
}

func (Config) TableName() string {
	return "config"
}

type Notes struct {
	ID                int                 `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Type              int                 `gorm:"column:type;size:32;default:0;not null" json:"type"`
	Content           string              `gorm:"column:content;default:'';not null" json:"content"`
	IsArchived        bool                `gorm:"column:isArchived;default:false;not null" json:"isArchived"`
	IsRecycle         bool                `gorm:"column:isRecycle;default:false;not null" json:"isRecycle"`
	IsShare           bool                `gorm:"column:isShare;default:false;not null" json:"isShare"`
	IsTop             bool                `gorm:"column:isTop;default:false;not null" json:"isTop"`
	IsReviewed        bool                `gorm:"column:isReviewed;default:false;not null" json:"isReviewed"`
	SharePassword     string              `gorm:"column:sharePassword;default:'';not null" json:"sharePassword"`
	ShareEncryptedURL *string             `gorm:"column:shareEncryptedUrl" json:"shareEncryptedUrl"`
	ShareExpiryDate   *DateTime           `gorm:"column:shareExpiryDate;type:timestamptz(6)" json:"shareExpiryDate"`
	ShareMaxView      *int                `gorm:"column:shareMaxView;size:32;default:0" json:"shareMaxView"`
	ShareViewCount    *int                `gorm:"column:shareViewCount;size:32;default:0" json:"shareViewCount"`
	Metadata          datatypes.JSON      `gorm:"column:metadata" json:"metadata"`
	AccountID         *int                `gorm:"column:accountId;size:32" json:"accountId"`
	SortOrder         int                 `gorm:"column:sortOrder;size:32;default:0;not null" json:"sortOrder"`
	CreatedAt         DateTime            `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt         DateTime            `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Account           *Accounts           `gorm:"foreignKey:AccountID;belongsTo"`
	Attachments       []Attachments       `gorm:"foreignKey:NoteID;constraint:attachments_noteId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tags              []TagsToNote        `gorm:"foreignKey:NoteID;constraint:tagsToNote_noteId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ReferencedBy      []NoteReference     `gorm:"foreignKey:ToNoteID;constraint:noteReference_toNoteId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	References        []NoteReference     `gorm:"foreignKey:FromNoteID;constraint:noteReference_fromNoteId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Comments          []Comments          `gorm:"foreignKey:NoteID;constraint:comments_noteId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Histories         []NoteHistory       `gorm:"foreignKey:NoteID;constraint:noteHistory_noteId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	InternalShares    []NoteInternalShare `gorm:"foreignKey:NoteID;constraint:noteInternalShare_noteId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Notes) TableName() string {
	return "notes"
}

func (m *Notes) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Notes) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Comments struct {
	ID        int        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Content   string     `gorm:"column:content;type:text;not null" json:"content"`
	AccountID *int       `gorm:"column:accountId;size:32" json:"accountId"`
	GuestName *string    `gorm:"column:guestName" json:"guestName"`
	GuestIP   *string    `gorm:"column:guestIP" json:"guestIP"`
	GuestUA   *string    `gorm:"column:guestUA" json:"guestUA"`
	NoteID    int        `gorm:"column:noteId;size:32;not null" json:"noteId"`
	ParentID  *int       `gorm:"column:parentId;size:32" json:"parentId"`
	CreatedAt DateTime   `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime   `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Note      *Notes     `gorm:"foreignKey:NoteID;belongsTo"`
	Account   *Accounts  `gorm:"foreignKey:AccountID;belongsTo"`
	Parent    *Comments  `gorm:"foreignKey:ParentID;belongsTo"`
	Replies   []Comments `gorm:"foreignKey:ParentID;constraint:comments_parentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Comments) TableName() string {
	return "comments"
}

func (m *Comments) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type Tag struct {
	ID         int          `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name       string       `gorm:"column:name;default:'';not null" json:"name"`
	Icon       string       `gorm:"column:icon;default:'';not null" json:"icon"`
	Parent     int          `gorm:"column:parent;size:32;default:0;not null" json:"parent"`
	AccountID  *int         `gorm:"column:accountId;size:32" json:"accountId"`
	CreatedAt  DateTime     `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime     `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	SortOrder  int          `gorm:"column:sortOrder;size:32;default:0;not null" json:"sortOrder"`
	Account    *Accounts    `gorm:"foreignKey:AccountID;belongsTo"`
	TagsToNote []TagsToNote `gorm:"foreignKey:TagID;constraint:tagsToNote_tagId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Tag) TableName() string {
	return "tag"
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

type TagsToNote struct {
	ID     int    `gorm:"column:id;size:32;not null" json:"id"`
	NoteID int    `gorm:"column:noteId;primaryKey;size:32;default:0" json:"noteId"`
	TagID  int    `gorm:"column:tagId;primaryKey;size:32;default:0" json:"tagId"`
	Note   *Notes `gorm:"foreignKey:NoteID;belongsTo"`
	Tag    *Tag   `gorm:"belongsTo"`
}

func (TagsToNote) TableName() string {
	return "tagsToNote"
}

type NoteReference struct {
	ID         int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	FromNoteID int      `gorm:"column:fromNoteId;uniqueIndex:noteReference_fromNoteId_toNoteId_key;size:32;not null" json:"fromNoteId"`
	ToNoteID   int      `gorm:"column:toNoteId;uniqueIndex:noteReference_fromNoteId_toNoteId_key;size:32;not null" json:"toNoteId"`
	CreatedAt  DateTime `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	FromNote   *Notes   `gorm:"foreignKey:FromNoteID;belongsTo"`
	ToNote     *Notes   `gorm:"foreignKey:ToNoteID;belongsTo"`
}

func (NoteReference) TableName() string {
	return "noteReference"
}

func (m *NoteReference) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Follows struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	SiteName    *string   `gorm:"column:siteName" json:"siteName"`
	SiteURL     string    `gorm:"column:siteUrl;not null" json:"siteUrl"`
	SiteAvatar  *string   `gorm:"column:siteAvatar" json:"siteAvatar"`
	Description *string   `gorm:"column:description;type:text" json:"description"`
	CreatedAt   DateTime  `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime  `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	FollowType  *string   `gorm:"column:followType;default:'following';not null" json:"followType"`
	AccountID   int       `gorm:"column:accountId;size:32;not null" json:"accountId"`
	Account     *Accounts `gorm:"foreignKey:AccountID;belongsTo"`
}

func (Follows) TableName() string {
	return "follows"
}

func (m *Follows) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Follows) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Notifications struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Type      string         `gorm:"column:type;not null" json:"type"`
	Title     string         `gorm:"column:title;not null" json:"title"`
	Content   string         `gorm:"column:content;type:text;not null" json:"content"`
	Metadata  datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	IsRead    bool           `gorm:"column:isRead;default:false;not null" json:"isRead"`
	AccountID int            `gorm:"column:accountId;size:32;not null" json:"accountId"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Account   *Accounts      `gorm:"foreignKey:AccountID;belongsTo"`
}

func (Notifications) TableName() string {
	return "notifications"
}

func (m *Notifications) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type Cache struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Key       string         `gorm:"column:key;uniqueIndex:cache_key_key;not null" json:"key"`
	Value     datatypes.JSON `gorm:"column:value;not null" json:"value"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Cache) TableName() string {
	return "cache"
}

func (m *Cache) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Cache) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Plugin struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Metadata  datatypes.JSON `gorm:"column:metadata;not null" json:"metadata"`
	Path      string         `gorm:"column:path;not null" json:"path"`
	IsUse     *bool          `gorm:"column:isUse;default:true;not null" json:"isUse"`
	IsDev     bool           `gorm:"column:isDev;default:false;not null" json:"isDev"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Plugin) TableName() string {
	return "plugin"
}

func (m *Plugin) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Plugin) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Conversation struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Title     string    `gorm:"column:title;type:varchar(255);default:'';not null" json:"title"`
	IsShare   bool      `gorm:"column:isShare;default:false;not null" json:"isShare"`
	AccountID int       `gorm:"column:accountId;index:conversation_accountId_idx;size:32;not null" json:"accountId"`
	CreatedAt DateTime  `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime  `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Account   *Accounts `gorm:"foreignKey:AccountID;belongsTo"`
	Messages  []Message `gorm:"foreignKey:ConversationID;constraint:message_conversationId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Conversation) TableName() string {
	return "conversation"
}

func (m *Conversation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type Message struct {
	ID             int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Content        string         `gorm:"column:content;type:text;not null" json:"content"`
	Role           string         `gorm:"column:role;type:varchar(50);not null" json:"role"`
	ConversationID int            `gorm:"column:conversationId;index:message_conversationId_idx;size:32;not null" json:"conversationId"`
	CreatedAt      DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt      DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Metadata       datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	Conversation   *Conversation  `gorm:"belongsTo"`
}

func (Message) TableName() string {
	return "message"
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
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

type NoteHistory struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	NoteID    int            `gorm:"column:noteId;index:noteHistory_noteId_idx;size:32;not null" json:"noteId"`
	Content   string         `gorm:"column:content;type:text;not null" json:"content"`
	Metadata  datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	Version   int            `gorm:"column:version;size:32;not null" json:"version"`
	AccountID *int           `gorm:"column:accountId;index:noteHistory_accountId_idx;size:32" json:"accountId"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Note      *Notes         `gorm:"foreignKey:NoteID;belongsTo"`
}

func (NoteHistory) TableName() string {
	return "noteHistory"
}

func (m *NoteHistory) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type NoteInternalShare struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	NoteID    int       `gorm:"column:noteId;uniqueIndex:noteInternalShare_noteId_accountId_key;index:noteInternalShare_noteId_idx;size:32;not null" json:"noteId"`
	AccountID int       `gorm:"column:accountId;uniqueIndex:noteInternalShare_noteId_accountId_key;index:noteInternalShare_accountId_idx;size:32;not null" json:"accountId"`
	CanEdit   *bool     `gorm:"column:canEdit;default:true;not null" json:"canEdit"`
	CreatedAt DateTime  `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime  `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Note      *Notes    `gorm:"foreignKey:NoteID;belongsTo"`
	Account   *Accounts `gorm:"foreignKey:AccountID;belongsTo"`
}

func (NoteInternalShare) TableName() string {
	return "noteInternalShare"
}

func (m *NoteInternalShare) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*NoteInternalShare) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Session struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	Sid       string   `gorm:"column:sid;uniqueIndex:session_sid_key;not null" json:"sid"`
	Data      string   `gorm:"column:data;type:text;not null" json:"data"`
	ExpiresAt DateTime `gorm:"column:expiresAt;index:session_expiresAt_idx;type:timestamp(3);not null" json:"expiresAt"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Session) TableName() string {
	return "session"
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

type AiProviders struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Title     string         `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Provider  string         `gorm:"column:provider;type:varchar(50);not null" json:"provider"`
	BaseURL   *string        `gorm:"column:baseURL;type:varchar(500)" json:"baseURL"`
	APIKey    *string        `gorm:"column:apiKey;type:varchar(500)" json:"apiKey"`
	Config    datatypes.JSON `gorm:"column:config" json:"config"`
	SortOrder int            `gorm:"column:sortOrder;size:32;default:0;not null" json:"sortOrder"`
	CreatedAt DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Models    []AiModels     `gorm:"foreignKey:ProviderID;constraint:aiModels_providerId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AiProviders) TableName() string {
	return "aiProviders"
}

func (m *AiProviders) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AiProviders) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AiModels struct {
	ID           int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	ProviderID   int            `gorm:"column:providerId;index:aiModels_providerId_idx;size:32;not null" json:"providerId"`
	Title        string         `gorm:"column:title;type:varchar(255);not null" json:"title"`
	ModelKey     string         `gorm:"column:modelKey;type:varchar(255);not null" json:"modelKey"`
	Capabilities datatypes.JSON `gorm:"column:capabilities;not null" json:"capabilities"`
	Config       datatypes.JSON `gorm:"column:config" json:"config"`
	SortOrder    int            `gorm:"column:sortOrder;size:32;default:0;not null" json:"sortOrder"`
	CreatedAt    DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt    DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Provider     *AiProviders   `gorm:"foreignKey:ProviderID;belongsTo"`
}

func (AiModels) TableName() string {
	return "aiModels"
}

func (m *AiModels) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AiModels) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type AiScheduledTask struct {
	ID         int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name       string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Prompt     string         `gorm:"column:prompt;type:text;not null" json:"prompt"`
	Schedule   string         `gorm:"column:schedule;type:varchar(100);not null" json:"schedule"`
	IsEnabled  *bool          `gorm:"column:isEnabled;default:true;not null" json:"isEnabled"`
	LastRun    *DateTime      `gorm:"column:lastRun;type:timestamptz(6)" json:"lastRun"`
	LastResult datatypes.JSON `gorm:"column:lastResult" json:"lastResult"`
	AccountID  int            `gorm:"column:accountId;index:aiScheduledTask_accountId_idx;size:32;not null" json:"accountId"`
	CreatedAt  DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
	Account    *Accounts      `gorm:"foreignKey:AccountID;belongsTo"`
}

func (AiScheduledTask) TableName() string {
	return "aiScheduledTask"
}

func (m *AiScheduledTask) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*AiScheduledTask) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type McpServers struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name        string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description *string        `gorm:"column:description;type:text" json:"description"`
	Type        string         `gorm:"column:type;type:varchar(50);not null" json:"type"`
	Command     *string        `gorm:"column:command;type:text" json:"command"`
	Args        datatypes.JSON `gorm:"column:args" json:"args"`
	URL         *string        `gorm:"column:url;type:varchar(500)" json:"url"`
	Env         datatypes.JSON `gorm:"column:env" json:"env"`
	Headers     datatypes.JSON `gorm:"column:headers" json:"headers"`
	IsEnabled   *bool          `gorm:"column:isEnabled;default:true;not null" json:"isEnabled"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (McpServers) TableName() string {
	return "mcpServers"
}

func (m *McpServers) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*McpServers) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Fonts struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name        string         `gorm:"column:name;uniqueIndex:fonts_name_key;type:varchar(255);not null" json:"name"`
	DisplayName string         `gorm:"column:displayName;type:varchar(255);not null" json:"displayName"`
	URL         *string        `gorm:"column:url;type:text" json:"url"`
	FileData    []byte         `gorm:"column:fileData" json:"fileData"`
	IsLocal     bool           `gorm:"column:isLocal;default:false;not null" json:"isLocal"`
	IsSystem    bool           `gorm:"column:isSystem;default:false;not null" json:"isSystem"`
	Weights     datatypes.JSON `gorm:"column:weights;not null" json:"weights"`
	Category    *string        `gorm:"column:category;type:varchar(50);default:'sans-serif';not null" json:"category"`
	SortOrder   int            `gorm:"column:sortOrder;size:32;default:0;not null" json:"sortOrder"`
	CreatedAt   DateTime       `gorm:"column:createdAt;type:timestamptz(6);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime       `gorm:"column:updatedAt;type:timestamptz(6);autoUpdateTime:false;not null" json:"updatedAt"`
}

func (Fonts) TableName() string {
	return "fonts"
}

func (m *Fonts) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Fonts) BeforeUpdate(tx *gorm.DB) error {
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
