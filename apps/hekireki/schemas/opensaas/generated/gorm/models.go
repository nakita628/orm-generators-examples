package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                     string               `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt              DateTime             `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Email                  *string              `gorm:"column:email;uniqueIndex:User_email_key" json:"email"`
	Username               *string              `gorm:"column:username;uniqueIndex:User_username_key" json:"username"`
	IsAdmin                bool                 `gorm:"column:isAdmin;default:false;not null" json:"isAdmin"`
	PaymentProcessorUserID *string              `gorm:"column:paymentProcessorUserId;uniqueIndex:User_paymentProcessorUserId_key" json:"paymentProcessorUserId"`
	SubscriptionStatus     *string              `gorm:"column:subscriptionStatus" json:"subscriptionStatus"`
	SubscriptionPlan       *string              `gorm:"column:subscriptionPlan" json:"subscriptionPlan"`
	CurrentPeriodEnd       *DateTime            `gorm:"column:currentPeriodEnd;type:timestamp(3)" json:"currentPeriodEnd"`
	CancelAtPeriodEnd      bool                 `gorm:"column:cancelAtPeriodEnd;default:false;not null" json:"cancelAtPeriodEnd"`
	DatePaid               *DateTime            `gorm:"column:datePaid;type:timestamp(3)" json:"datePaid"`
	Credits                *int                 `gorm:"column:credits;size:32;default:3;not null" json:"credits"`
	GptResponses           []GptResponse        `gorm:"foreignKey:UserID;constraint:GptResponse_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ContactFormMessages    []ContactFormMessage `gorm:"foreignKey:UserID;constraint:ContactFormMessage_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Tasks                  []Task               `gorm:"foreignKey:UserID;constraint:Task_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Files                  []File               `gorm:"foreignKey:UserID;constraint:File_userId_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
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
	return nil
}

type ProcessedWebhookEvent struct {
	ID                      string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt               DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	PaymentProcessorEventID string   `gorm:"column:paymentProcessorEventId;uniqueIndex:ProcessedWebhookEvent_paymentProcessorEventId_key;not null" json:"paymentProcessorEventId"`
	EventType               string   `gorm:"column:eventType;not null" json:"eventType"`
}

func (ProcessedWebhookEvent) TableName() string {
	return "ProcessedWebhookEvent"
}

func (m *ProcessedWebhookEvent) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type GptResponse struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UpdatedAt DateTime `gorm:"column:updatedAt;type:timestamp(3);autoUpdateTime:false;not null" json:"updatedAt"`
	UserID    string   `gorm:"column:userId;not null" json:"userId"`
	Content   string   `gorm:"column:content;not null" json:"content"`
	User      *User    `gorm:"belongsTo"`
}

func (GptResponse) TableName() string {
	return "GptResponse"
}

func (m *GptResponse) BeforeCreate(tx *gorm.DB) error {
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

func (*GptResponse) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type Task struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UserID      string   `gorm:"column:userId;not null" json:"userId"`
	Description string   `gorm:"column:description;not null" json:"description"`
	Time        *string  `gorm:"column:time;default:'1';not null" json:"time"`
	IsDone      bool     `gorm:"column:isDone;default:false;not null" json:"isDone"`
	User        *User    `gorm:"belongsTo"`
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
	return nil
}

type File struct {
	ID        string   `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UserID    string   `gorm:"column:userId;not null" json:"userId"`
	Name      string   `gorm:"column:name;not null" json:"name"`
	Type      string   `gorm:"column:type;not null" json:"type"`
	S3Key     string   `gorm:"column:s3Key;not null" json:"s3Key"`
	User      *User    `gorm:"belongsTo"`
}

func (File) TableName() string {
	return "File"
}

func (m *File) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DailyStats struct {
	ID                        int              `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Date                      DateTime         `gorm:"column:date;uniqueIndex:DailyStats_date_key;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"date"`
	TotalViews                int              `gorm:"column:totalViews;size:32;default:0;not null" json:"totalViews"`
	PrevDayViewsChangePercent *string          `gorm:"column:prevDayViewsChangePercent;default:'0';not null" json:"prevDayViewsChangePercent"`
	UserCount                 int              `gorm:"column:userCount;size:32;default:0;not null" json:"userCount"`
	PaidUserCount             int              `gorm:"column:paidUserCount;size:32;default:0;not null" json:"paidUserCount"`
	UserDelta                 int              `gorm:"column:userDelta;size:32;default:0;not null" json:"userDelta"`
	PaidUserDelta             int              `gorm:"column:paidUserDelta;size:32;default:0;not null" json:"paidUserDelta"`
	TotalRevenue              float64          `gorm:"column:totalRevenue;type:double precision;default:0;not null" json:"totalRevenue"`
	TotalProfit               float64          `gorm:"column:totalProfit;type:double precision;default:0;not null" json:"totalProfit"`
	Sources                   []PageViewSource `gorm:"foreignKey:DailyStatsID;constraint:PageViewSource_dailyStatsId_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (DailyStats) TableName() string {
	return "DailyStats"
}

func (m *DailyStats) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Date.IsZero() {
		m.Date = DateTime{Time: now}
	}
	return nil
}

type PageViewSource struct {
	Date         DateTime    `gorm:"column:date;primaryKey;type:timestamp(3);default:CURRENT_TIMESTAMP" json:"date"`
	Name         string      `gorm:"column:name;primaryKey" json:"name"`
	DailyStatsID *int        `gorm:"column:dailyStatsId;size:32" json:"dailyStatsId"`
	Visitors     int         `gorm:"column:visitors;size:32;not null" json:"visitors"`
	DailyStats   *DailyStats `gorm:"belongsTo"`
}

func (PageViewSource) TableName() string {
	return "PageViewSource"
}

func (m *PageViewSource) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.Date.IsZero() {
		m.Date = DateTime{Time: now}
	}
	return nil
}

type Logs struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	CreatedAt DateTime `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	Message   string   `gorm:"column:message;not null" json:"message"`
	Level     string   `gorm:"column:level;not null" json:"level"`
}

func (Logs) TableName() string {
	return "Logs"
}

func (m *Logs) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type ContactFormMessage struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt DateTime  `gorm:"column:createdAt;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"createdAt"`
	UserID    string    `gorm:"column:userId;not null" json:"userId"`
	Content   string    `gorm:"column:content;not null" json:"content"`
	IsRead    bool      `gorm:"column:isRead;default:false;not null" json:"isRead"`
	RepliedAt *DateTime `gorm:"column:repliedAt;type:timestamp(3)" json:"repliedAt"`
	User      *User     `gorm:"belongsTo"`
}

func (ContactFormMessage) TableName() string {
	return "ContactFormMessage"
}

func (m *ContactFormMessage) BeforeCreate(tx *gorm.DB) error {
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
