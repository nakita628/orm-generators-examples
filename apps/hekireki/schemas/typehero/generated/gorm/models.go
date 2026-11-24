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
	"gorm.io/gorm/schema"
)

var NamingStrategy = schema.NamingStrategy{SingularTable: true, NoLowerCase: true}

type Account struct {
	ID                string  `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	UserID            string  `gorm:"column:userId;index:Account_userId_idx;type:varchar(191);not null" json:"userId"`
	Type              string  `gorm:"column:type;type:varchar(191);not null" json:"type"`
	Provider          string  `gorm:"column:provider;uniqueIndex:Account_provider_providerAccountId_key;type:varchar(191);not null" json:"provider"`
	ProviderAccountID string  `gorm:"column:providerAccountId;uniqueIndex:Account_provider_providerAccountId_key;type:varchar(191);not null" json:"providerAccountId"`
	RefreshToken      *string `gorm:"column:refresh_token;type:varchar(191)" json:"refresh_token"`
	AccessToken       *string `gorm:"column:access_token;type:varchar(191)" json:"access_token"`
	ExpiresAt         *int    `gorm:"column:expires_at;size:32" json:"expires_at"`
	TokenType         *string `gorm:"column:token_type;type:varchar(191)" json:"token_type"`
	Scope             *string `gorm:"column:scope;type:varchar(191)" json:"scope"`
	IDToken           *string `gorm:"column:id_token;type:varchar(191)" json:"id_token"`
	SessionState      *string `gorm:"column:session_state;type:varchar(191)" json:"session_state"`
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
	ID           string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	SessionToken string   `gorm:"column:sessionToken;uniqueIndex:Session_sessionToken_key;type:varchar(191);not null" json:"sessionToken"`
	UserID       string   `gorm:"column:userId;index:Session_userId_idx;type:varchar(191);not null" json:"userId"`
	Expires      DateTime `gorm:"column:expires;not null" json:"expires"`
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

type Role struct {
	ID   string `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Role string `gorm:"column:role;uniqueIndex:Role_role_key;not null" json:"role"`
	User []User `gorm:"many2many:_RoleToUser;joinForeignKey:A;joinReferences:B"`
}

func (Role) TableName() string {
	return "Role"
}

func (m *Role) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

// This is a table that contains all the user links they can add to profile
type UserLink struct {
	ID        string   `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	URL       string   `gorm:"column:url;type:varchar(256);not null" json:"url"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	User      []User   `gorm:"many2many:_UserToUserLink;joinForeignKey:B;joinReferences:A"`
}

func (UserLink) TableName() string {
	return "UserLink"
}

func (m *UserLink) BeforeCreate(tx *gorm.DB) error {
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

func (*UserLink) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type User struct {
	ID            string    `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Name          string    `gorm:"column:name;type:varchar(191);not null" json:"name"`
	Email         string    `gorm:"column:email;uniqueIndex:User_email_key;type:varchar(191);not null" json:"email"`
	EmailVerified *DateTime `gorm:"column:emailVerified" json:"emailVerified"`
	Image         *string   `gorm:"column:image;type:varchar(191)" json:"image"`
	// this is the bio that shows up on the profile page
	Bio                  string           `gorm:"column:bio;type:varchar(1000);default:'';not null" json:"bio"`
	Status               string           `gorm:"column:status;default:'ACTIVE';not null" json:"status"`
	BanReason            *string          `gorm:"column:banReason;type:text" json:"banReason"`
	CreatedAt            DateTime         `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt            DateTime         `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Accounts             []Account        `gorm:"foreignKey:UserID;constraint:Account_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sessions             []Session        `gorm:"foreignKey:UserID;constraint:Session_userId_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Challenge            []Challenge      `gorm:"foreignKey:UserID;constraint:Challenge_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Vote                 []Vote           `gorm:"foreignKey:UserID;constraint:Vote_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Submission           []Submission     `gorm:"foreignKey:UserID;constraint:Submission_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Bookmark             []Bookmark       `gorm:"foreignKey:UserID;constraint:Bookmark_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SharedSolution       []SharedSolution `gorm:"foreignKey:UserID;constraint:SharedSolution_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comment              []Comment        `gorm:"foreignKey:UserID;constraint:Comment_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Reports              []Report         `gorm:"foreignKey:ReporterID;constraint:Report_reporterId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ReportsAbout         []Report         `gorm:"foreignKey:UserID;constraint:Report_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Moderated            []Report         `gorm:"foreignKey:ModeratorID;constraint:Report_moderatorId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ImageUpload          []ImageUpload    `gorm:"foreignKey:UserID;constraint:ImageUpload_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	BetaTokens           []BetaTokens     `gorm:"foreignKey:UserID;constraint:BetaTokens_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Donations            []Donation       `gorm:"foreignKey:UserID;constraint:Donation_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ShortUrls            []ShortURL       `gorm:"foreignKey:UserID;constraint:ShortURL_userId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	ToUserNotification   []Notification   `gorm:"foreignKey:ToUserID;constraint:Notification_toUserId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	FromUserNotification []Notification   `gorm:"foreignKey:FromUserID;constraint:Notification_fromUserId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Roles                []Role           `gorm:"many2many:_RoleToUser;joinForeignKey:B;joinReferences:A"`
	// this is the twitter URL that shows up on the profile page
	UserLinks []UserLink `gorm:"many2many:_UserToUserLink;joinForeignKey:A;joinReferences:B"`
	Tracks    []Track    `gorm:"many2many:_TrackToUser;joinForeignKey:B;joinReferences:A"`
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
	Identifier string   `gorm:"column:identifier;uniqueIndex:VerificationToken_identifier_token_key;type:varchar(191);not null" json:"identifier"`
	Token      string   `gorm:"column:token;primaryKey;uniqueIndex:VerificationToken_token_key;uniqueIndex:VerificationToken_identifier_token_key;type:varchar(191)" json:"token"`
	Expires    DateTime `gorm:"column:expires;not null" json:"expires"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

type Tag struct {
	ID         string      `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	Tag        string      `gorm:"column:tag;uniqueIndex:Tag_tag_key;not null" json:"tag"`
	Challenges []Challenge `gorm:"many2many:_ChallengeToTag;joinForeignKey:B;joinReferences:A"`
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

type Challenge struct {
	ID         int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	CreatedAt  DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Difficulty string   `gorm:"column:difficulty;not null" json:"difficulty"`
	Name       string   `gorm:"column:name;uniqueIndex:Challenge_name_key;type:varchar(191);not null" json:"name"`
	Slug       string   `gorm:"column:slug;uniqueIndex:Challenge_slug_key;type:varchar(191);not null" json:"slug"`
	// this is the description to describe the challenge on description panel
	Description string `gorm:"column:description;type:text;not null" json:"description"`
	// this is a short blue thats shown on the explore card
	ShortDescription string `gorm:"column:shortDescription;type:varchar(191);not null" json:"shortDescription"`
	// Starter code for the challenge
	Code string `gorm:"column:code;type:text;not null" json:"code"`
	// Test cases that need to be passed in order have the submission be accepted.
	Tests          string           `gorm:"column:tests;type:text;not null" json:"tests"`
	Status         string           `gorm:"column:status;default:'PENDING';not null" json:"status"`
	Tsconfig       datatypes.JSON   `gorm:"column:tsconfig" json:"tsconfig"`
	UserID         string           `gorm:"column:userId;index:Challenge_userId_idx;type:varchar(191);not null" json:"userId"`
	User           *User            `gorm:"belongsTo"`
	Vote           []Vote           `gorm:"foreignKey:ChallengeID;constraint:Vote_challengeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Submission     []Submission     `gorm:"foreignKey:ChallengeID;constraint:Submission_challengeId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Comment        []Comment        `gorm:"foreignKey:RootChallengeID;constraint:Comment_rootChallengeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Bookmark       []Bookmark       `gorm:"foreignKey:ChallengeID;constraint:Bookmark_challengeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	SharedSolution []SharedSolution `gorm:"foreignKey:ChallengeID;constraint:SharedSolution_challengeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Report         []Report         `gorm:"foreignKey:ChallengeID;constraint:Report_challengeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TrackChallenge []TrackChallenge `gorm:"foreignKey:ChallengeID;constraint:TrackChallenge_challengeId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Notification   []Notification   `gorm:"foreignKey:ChallengeID;constraint:Notification_challengeId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tags           []Tag            `gorm:"many2many:_ChallengeToTag;joinForeignKey:A;joinReferences:B"`
}

func (Challenge) TableName() string {
	return "Challenge"
}

func (m *Challenge) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Challenge) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Vote struct {
	ID               int             `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID           string          `gorm:"column:userId;index:Vote_userId_idx;type:varchar(191);not null" json:"userId"`
	RootType         string          `gorm:"column:rootType;not null" json:"rootType"`
	ChallengeID      *int            `gorm:"column:challengeId;index:Vote_challengeId_idx;size:32" json:"challengeId"`
	CommentID        *int            `gorm:"column:commentId;index:Vote_commentId_idx;size:32" json:"commentId"`
	SubmissionID     *int            `gorm:"column:submissionId;index:Vote_submissionId_idx;size:32" json:"submissionId"`
	SharedSolutionID *int            `gorm:"column:sharedSolutionId;index:Vote_sharedSolutionId_idx;size:32" json:"sharedSolutionId"`
	User             *User           `gorm:"belongsTo"`
	Challenge        *Challenge      `gorm:"belongsTo"`
	Comment          *Comment        `gorm:"belongsTo"`
	Submission       *Submission     `gorm:"belongsTo"`
	SharedSolution   *SharedSolution `gorm:"belongsTo"`
}

func (Vote) TableName() string {
	return "Vote"
}

// This is a public solution created by a user for a challenge
type SharedSolution struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	CreatedAt       DateTime       `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	IsPinned        bool           `gorm:"column:isPinned;default:false;not null" json:"isPinned"`
	Title           string         `gorm:"column:title;type:varchar(191);not null" json:"title"`
	Description     string         `gorm:"column:description;type:text;not null" json:"description"`
	UserID          *string        `gorm:"column:userId;index:SharedSolution_userId_idx;type:varchar(191)" json:"userId"`
	ChallengeID     *int           `gorm:"column:challengeId;index:SharedSolution_challengeId_idx;size:32" json:"challengeId"`
	User            *User          `gorm:"belongsTo"`
	Challenge       *Challenge     `gorm:"belongsTo"`
	Vote            []Vote         `gorm:"foreignKey:SharedSolutionID;constraint:Vote_sharedSolutionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	SolutionComment []Comment      `gorm:"foreignKey:RootSolutionID;constraint:Comment_rootSolutionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Report          []Report       `gorm:"foreignKey:SolutionID;constraint:Report_solutionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Notification    []Notification `gorm:"foreignKey:SharedSolutionID;constraint:Notification_sharedSolutionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (SharedSolution) TableName() string {
	return "SharedSolution"
}

func (m *SharedSolution) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

// Every attempt a user makes to a challenge is stored here
type Submission struct {
	ID           int        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Code         string     `gorm:"column:code;type:text;not null" json:"code"`
	CreatedAt    DateTime   `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	IsSuccessful bool       `gorm:"column:isSuccessful;not null" json:"isSuccessful"`
	UserID       string     `gorm:"column:userId;index:Submission_userId_idx;type:varchar(191);not null" json:"userId"`
	ChallengeID  int        `gorm:"column:challengeId;index:Submission_challengeId_idx;size:32;not null" json:"challengeId"`
	User         *User      `gorm:"belongsTo"`
	Challenge    *Challenge `gorm:"belongsTo"`
	Vote         []Vote     `gorm:"foreignKey:SubmissionID;constraint:Vote_submissionId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Submission) TableName() string {
	return "Submission"
}

func (m *Submission) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Bookmark struct {
	ID          int        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID      string     `gorm:"column:userId;uniqueIndex:Bookmark_userId_challengeId_key;index:Bookmark_userId_idx;type:varchar(191);not null" json:"userId"`
	ChallengeID *int       `gorm:"column:challengeId;uniqueIndex:Bookmark_userId_challengeId_key;index:Bookmark_challengeId_idx;size:32" json:"challengeId"`
	CreatedAt   *DateTime  `gorm:"column:createdAt;autoCreateTime:false" json:"createdAt"`
	User        *User      `gorm:"belongsTo"`
	Challenge   *Challenge `gorm:"belongsTo"`
}

func (Bookmark) TableName() string {
	return "Bookmark"
}

func (m *Bookmark) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt == nil {
		m.CreatedAt = &DateTime{Time: now}
	}
	return nil
}

type ReportIssue struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	ReportID int     `gorm:"column:reportId;index:ReportIssue_reportId_idx;size:32;not null" json:"reportId"`
	Type     string  `gorm:"column:type;not null" json:"type"`
	Report   *Report `gorm:"belongsTo"`
}

func (ReportIssue) TableName() string {
	return "ReportIssue"
}

type Report struct {
	ID          int             `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Text        string          `gorm:"column:text;type:text;not null" json:"text"`
	Type        string          `gorm:"column:type;not null" json:"type"`
	Status      string          `gorm:"column:status;not null" json:"status"`
	ReporterID  string          `gorm:"column:reporterId;index:Report_reporterId_idx;type:varchar(191);not null" json:"reporterId"`
	UserID      *string         `gorm:"column:userId;index:Report_userId_idx;type:varchar(191)" json:"userId"`
	ChallengeID *int            `gorm:"column:challengeId;index:Report_challengeId_idx;size:32" json:"challengeId"`
	CommentID   *int            `gorm:"column:commentId;index:Report_commentId_idx;size:32" json:"commentId"`
	SolutionID  *int            `gorm:"column:solutionId;index:Report_solutionId_idx;size:32" json:"solutionId"`
	ModeratorID *string         `gorm:"column:moderatorId;index:Report_moderatorId_idx;type:varchar(191)" json:"moderatorId"`
	CreatedAt   DateTime        `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt   DateTime        `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	Reporter    *User           `gorm:"foreignKey:ReporterID;belongsTo"`
	User        *User           `gorm:"foreignKey:UserID;belongsTo"`
	Challenge   *Challenge      `gorm:"belongsTo"`
	Comment     *Comment        `gorm:"belongsTo"`
	Solution    *SharedSolution `gorm:"foreignKey:SolutionID;belongsTo"`
	Moderator   *User           `gorm:"foreignKey:ModeratorID;belongsTo"`
	Issues      []ReportIssue   `gorm:"foreignKey:ReportID;constraint:ReportIssue_reportId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Report) TableName() string {
	return "Report"
}

func (m *Report) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Report) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Comment struct {
	ID              int             `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Text            string          `gorm:"column:text;type:text;not null" json:"text"`
	Visible         *bool           `gorm:"column:visible;default:true;not null" json:"visible"`
	ParentID        *int            `gorm:"column:parentId;index:Comment_parentId_idx;size:32" json:"parentId"`
	UserID          string          `gorm:"column:userId;index:Comment_userId_idx;type:varchar(191);not null" json:"userId"`
	RootType        string          `gorm:"column:rootType;default:'CHALLENGE';not null" json:"rootType"`
	RootChallengeID *int            `gorm:"column:rootChallengeId;index:Comment_rootChallengeId_idx;size:32" json:"rootChallengeId"`
	RootSolutionID  *int            `gorm:"column:rootSolutionId;index:Comment_rootSolutionId_idx;size:32" json:"rootSolutionId"`
	CreatedAt       DateTime        `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt       DateTime        `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	ParentComment   *Comment        `gorm:"foreignKey:ParentID;belongsTo"`
	User            *User           `gorm:"belongsTo"`
	RootChallenge   *Challenge      `gorm:"foreignKey:RootChallengeID;belongsTo"`
	RootSolution    *SharedSolution `gorm:"foreignKey:RootSolutionID;belongsTo"`
	Replies         []Comment       `gorm:"foreignKey:ParentID;constraint:Comment_parentId_fkey,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Report          []Report        `gorm:"foreignKey:CommentID;constraint:Report_commentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Vote            []Vote          `gorm:"foreignKey:CommentID;constraint:Vote_commentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	Notification    []Notification  `gorm:"foreignKey:CommentID;constraint:Notification_commentId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Comment) TableName() string {
	return "Comment"
}

func (m *Comment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Comment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ImageUpload struct {
	// the id of the db row
	ID string `gorm:"column:id;primaryKey;type:varchar(191)" json:"id"`
	// the user id fk
	UserID string `gorm:"column:userId;index:ImageUpload_userId_idx;type:varchar(191);not null" json:"userId"`
	// the url to the s3 uploaded assets
	URL       string   `gorm:"column:url;type:varchar(256);not null" json:"url"`
	CreatedAt DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	// The user that this is related to
	User *User `gorm:"belongsTo"`
}

func (ImageUpload) TableName() string {
	return "ImageUpload"
}

func (m *ImageUpload) BeforeCreate(tx *gorm.DB) error {
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

func (*ImageUpload) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type TrackChallenge struct {
	ID          int        `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	OrderID     int        `gorm:"column:orderId;size:32;not null" json:"orderId"`
	ChallengeID int        `gorm:"column:challengeId;index:TrackChallenge_challengeId_idx;size:32;not null" json:"challengeId"`
	TrackID     *int       `gorm:"column:trackId;index:TrackChallenge_trackId_idx;size:32" json:"trackId"`
	Challenge   *Challenge `gorm:"belongsTo"`
	Track       *Track     `gorm:"belongsTo"`
}

func (TrackChallenge) TableName() string {
	return "TrackChallenge"
}

type Track struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Name        string `gorm:"column:name;uniqueIndex:Track_name_key;type:varchar(191);not null" json:"name"`
	Slug        string `gorm:"column:slug;uniqueIndex:Track_slug_key;type:varchar(191);not null" json:"slug"`
	Description string `gorm:"column:description;type:text;not null" json:"description"`
	// determines if a track can be seen on tracks page
	Visible *bool `gorm:"column:visible;default:true;not null" json:"visible"`
	// use this to display coming soon UI
	IsComingSoon    bool             `gorm:"column:isComingSoon;default:false;not null" json:"isComingSoon"`
	TrackChallenges []TrackChallenge `gorm:"foreignKey:TrackID;constraint:TrackChallenge_trackId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	EnrolledUsers   []User           `gorm:"many2many:_TrackToUser;joinForeignKey:A;joinReferences:B"`
}

func (Track) TableName() string {
	return "Track"
}

type BetaTokens struct {
	ID     int     `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Token  string  `gorm:"column:token;uniqueIndex:BetaTokens_token_key;type:varchar(191);not null" json:"token"`
	UserID *string `gorm:"column:userId;index:BetaTokens_userId_idx;type:varchar(191)" json:"userId"`
	User   *User   `gorm:"belongsTo"`
}

func (BetaTokens) TableName() string {
	return "BetaTokens"
}

type Donation struct {
	ID                      int       `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	UserID                  string    `gorm:"column:userId;index:Donation_userId_idx;type:varchar(191);not null" json:"userId"`
	Amount                  int       `gorm:"column:amount;size:32;not null" json:"amount"`
	StripeCheckoutSessionID string    `gorm:"column:stripeCheckoutSessionId;uniqueIndex:Donation_stripeCheckoutSessionId_key;type:varchar(191);not null" json:"stripeCheckoutSessionId"`
	CreatedAt               *DateTime `gorm:"column:createdAt;autoCreateTime:false" json:"createdAt"`
	User                    *User     `gorm:"belongsTo"`
}

func (Donation) TableName() string {
	return "Donation"
}

func (m *Donation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt == nil {
		m.CreatedAt = &DateTime{Time: now}
	}
	return nil
}

type Notification struct {
	ID         int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	CreatedAt  DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt  DateTime `gorm:"column:updatedAt;autoUpdateTime:false;not null" json:"updatedAt"`
	IsRead     bool     `gorm:"column:isRead;default:false;not null" json:"isRead"`
	Type       string   `gorm:"column:type;not null" json:"type"`
	ToUserID   string   `gorm:"column:toUserId;index:Notification_toUserId_idx;type:varchar(191);not null" json:"toUserId"`
	FromUserID string   `gorm:"column:fromUserId;index:Notification_fromUserId_idx;type:varchar(191);not null" json:"fromUserId"`
	CommentID  *int     `gorm:"column:commentId;index:Notification_commentId_idx;size:32" json:"commentId"`
	// could be comment text or solution title to avoid having to join to grab this data
	Content *string `gorm:"column:content;type:varchar(191)" json:"content"`
	// can be used to replace the blurb (eg: mention you in a comment)
	Blurb *string `gorm:"column:blurb;type:varchar(191)" json:"blurb"`
	// use this instead of building the url on the fly (avoids a lot of joins and perf hits to build a string)
	URL              string          `gorm:"column:url;type:varchar(191);not null" json:"url"`
	ChallengeID      *int            `gorm:"column:challengeId;index:Notification_challengeId_idx;size:32" json:"challengeId"`
	SharedSolutionID *int            `gorm:"column:sharedSolutionId;index:Notification_sharedSolutionId_idx;size:32" json:"sharedSolutionId"`
	ToUser           *User           `gorm:"foreignKey:ToUserID;belongsTo"`
	FromUser         *User           `gorm:"foreignKey:FromUserID;belongsTo"`
	Comment          *Comment        `gorm:"belongsTo"`
	Challenge        *Challenge      `gorm:"belongsTo"`
	SharedSolution   *SharedSolution `gorm:"belongsTo"`
}

func (Notification) TableName() string {
	return "Notification"
}

func (m *Notification) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*Notification) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ShortURL struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	// the original full url that we want to shorten e.g. https://www.url.com
	OriginalURL string `gorm:"column:originalUrl;type:text;not null" json:"originalUrl"`
	// slug for the short url e.g. url.com/[shortUrlSlug]
	ShortURLSlug string   `gorm:"column:shortUrlSlug;uniqueIndex:ShortURL_shortUrlSlug_key;index:ShortURL_shortUrlSlug_idx;type:varchar(191);not null" json:"shortUrlSlug"`
	CreatedAt    DateTime `gorm:"column:createdAt;autoCreateTime:false;not null" json:"createdAt"`
	// the time the short url will expire
	ExpiresAt *DateTime `gorm:"column:expiresAt" json:"expiresAt"`
	// the number of times the short url has redirected to the original url
	Clicks int     `gorm:"column:clicks;size:32;default:0;not null" json:"clicks"`
	UserID *string `gorm:"column:userId;index:ShortURL_userId_idx;type:varchar(191)" json:"userId"`
	User   *User   `gorm:"belongsTo"`
}

func (ShortURL) TableName() string {
	return "ShortURL"
}

func (m *ShortURL) BeforeCreate(tx *gorm.DB) error {
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
