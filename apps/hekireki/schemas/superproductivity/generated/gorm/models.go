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

type User struct {
	ID                            int                          `gorm:"column:id;primaryKey;autoIncrement;size:32" json:"id"`
	Email                         string                       `gorm:"column:email;uniqueIndex:users_email_key;not null" json:"email"`
	PasswordHash                  *string                      `gorm:"column:password_hash" json:"password_hash"`
	IsVerified                    int                          `gorm:"column:is_verified;size:32;default:0;not null" json:"is_verified"`
	VerificationToken             *string                      `gorm:"column:verification_token;index:users_verification_token_idx" json:"verification_token"`
	VerificationTokenExpiresAt    *int64                       `gorm:"column:verification_token_expires_at" json:"verification_token_expires_at"`
	VerificationResendCount       int                          `gorm:"column:verification_resend_count;size:32;default:0;not null" json:"verification_resend_count"`
	ResetPasswordToken            *string                      `gorm:"column:reset_password_token;index:users_reset_password_token_idx" json:"reset_password_token"`
	ResetPasswordTokenExpiresAt   *int64                       `gorm:"column:reset_password_token_expires_at" json:"reset_password_token_expires_at"`
	PasskeyRecoveryToken          *string                      `gorm:"column:passkey_recovery_token;index:users_passkey_recovery_token_idx" json:"passkey_recovery_token"`
	PasskeyRecoveryTokenExpiresAt *int64                       `gorm:"column:passkey_recovery_token_expires_at" json:"passkey_recovery_token_expires_at"`
	LoginToken                    *string                      `gorm:"column:login_token;index:users_login_token_idx" json:"login_token"`
	LoginTokenExpiresAt           *int64                       `gorm:"column:login_token_expires_at" json:"login_token_expires_at"`
	FailedLoginAttempts           int                          `gorm:"column:failed_login_attempts;size:32;default:0;not null" json:"failed_login_attempts"`
	LockedUntil                   *int64                       `gorm:"column:locked_until" json:"locked_until"`
	TokenVersion                  int                          `gorm:"column:token_version;size:32;default:0;not null" json:"token_version"`
	TermsAcceptedAt               *int64                       `gorm:"column:terms_accepted_at" json:"terms_accepted_at"`
	CreatedAt                     DateTime                     `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	StorageQuotaBytes             *int64                       `gorm:"column:storage_quota_bytes;default:104857600;not null" json:"storage_quota_bytes"`
	StorageUsedBytes              int64                        `gorm:"column:storage_used_bytes;default:0;not null" json:"storage_used_bytes"`
	Operations                    []Operation                  `gorm:"foreignKey:UserID;constraint:operations_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Devices                       []SyncDevice                 `gorm:"foreignKey:UserID;constraint:sync_devices_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Passkeys                      []Passkey                    `gorm:"foreignKey:UserID;constraint:passkeys_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PendingPasskeyRegistrations   []PendingPasskeyRegistration `gorm:"foreignKey:UserID;constraint:pending_passkey_registrations_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CaptureGrants                 []CaptureGrant               `gorm:"foreignKey:UserID;constraint:capture_grants_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CaptureItems                  []CaptureItem                `gorm:"foreignKey:UserID;constraint:capture_items_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CaptureAuthorizations         []CaptureAuth                `gorm:"foreignKey:UserID;constraint:capture_auth_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConnectRequests               []ConnectRequest             `gorm:"foreignKey:UserID;constraint:connect_requests_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	SyncState                     *UserSyncState               `gorm:"foreignKey:UserID;constraint:user_sync_state_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CaptureKey                    *CaptureKey                  `gorm:"foreignKey:UserID;constraint:capture_keys_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	CaptureConnectCode            *CaptureConnectCode          `gorm:"foreignKey:UserID;constraint:capture_connect_codes_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "users"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Passkey struct {
	ID           string    `gorm:"column:id;primaryKey" json:"id"`
	CredentialID []byte    `gorm:"column:credential_id;uniqueIndex:passkeys_credential_id_key;not null" json:"credential_id"`
	PublicKey    []byte    `gorm:"column:public_key;not null" json:"public_key"`
	Counter      int64     `gorm:"column:counter;default:0;not null" json:"counter"`
	Transports   *string   `gorm:"column:transports" json:"transports"`
	CreatedAt    DateTime  `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	LastUsedAt   *DateTime `gorm:"column:last_used_at;type:timestamp(3)" json:"last_used_at"`
	UserID       int       `gorm:"column:user_id;index:passkeys_user_id_idx;size:32;not null" json:"user_id"`
	User         *User     `gorm:"belongsTo"`
}

func (Passkey) TableName() string {
	return "passkeys"
}

func (m *Passkey) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type PendingPasskeyRegistration struct {
	ID                         string   `gorm:"column:id;primaryKey" json:"id"`
	VerificationToken          string   `gorm:"column:verification_token;uniqueIndex:pending_passkey_registrations_verification_token_key;not null" json:"verification_token"`
	VerificationTokenExpiresAt int64    `gorm:"column:verification_token_expires_at;not null" json:"verification_token_expires_at"`
	CredentialID               []byte   `gorm:"column:credential_id;not null" json:"credential_id"`
	PublicKey                  []byte   `gorm:"column:public_key;not null" json:"public_key"`
	Counter                    int64    `gorm:"column:counter;default:0;not null" json:"counter"`
	Transports                 *string  `gorm:"column:transports" json:"transports"`
	CreatedAt                  DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UserID                     int      `gorm:"column:user_id;index:pending_passkey_registrations_user_id_idx;size:32;not null" json:"user_id"`
	User                       *User    `gorm:"belongsTo"`
}

func (PendingPasskeyRegistration) TableName() string {
	return "pending_passkey_registrations"
}

func (m *PendingPasskeyRegistration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type Operation struct {
	ID                  string             `gorm:"column:id;primaryKey" json:"id"`
	UserID              int                `gorm:"column:user_id;uniqueIndex:operations_user_id_server_seq_key;index:operations_user_id_entity_type_entity_id_server_seq_idx;index:operations_user_id_received_at_server_seq_idx;size:32;not null" json:"user_id"`
	ClientID            string             `gorm:"column:client_id;not null" json:"client_id"`
	ServerSeq           int                `gorm:"column:server_seq;uniqueIndex:operations_user_id_server_seq_key;index:operations_user_id_entity_type_entity_id_server_seq_idx;index:operations_user_id_received_at_server_seq_idx;size:32;not null" json:"server_seq"`
	ActionType          string             `gorm:"column:action_type;not null" json:"action_type"`
	OpType              string             `gorm:"column:op_type;not null" json:"op_type"`
	EntityType          string             `gorm:"column:entity_type;index:operations_user_id_entity_type_entity_id_server_seq_idx;not null" json:"entity_type"`
	EntityID            *string            `gorm:"column:entity_id;index:operations_user_id_entity_type_entity_id_server_seq_idx" json:"entity_id"`
	EntityIds           ScalarList[string] `gorm:"column:entity_ids;index:operations_entity_ids_gin;type:text[]" json:"entity_ids"`
	Payload             datatypes.JSON     `gorm:"column:payload;not null" json:"payload"`
	PayloadBytes        int64              `gorm:"column:payload_bytes;default:0;not null" json:"payload_bytes"`
	VectorClock         datatypes.JSON     `gorm:"column:vector_clock;not null" json:"vector_clock"`
	SchemaVersion       int                `gorm:"column:schema_version;size:32;not null" json:"schema_version"`
	ClientTimestamp     int64              `gorm:"column:client_timestamp;not null" json:"client_timestamp"`
	ReceivedAt          int64              `gorm:"column:received_at;index:operations_user_id_received_at_server_seq_idx;not null" json:"received_at"`
	IsPayloadEncrypted  bool               `gorm:"column:is_payload_encrypted;default:false;not null" json:"is_payload_encrypted"`
	SyncImportReason    *string            `gorm:"column:sync_import_reason" json:"sync_import_reason"`
	RepairBaseServerSeq *int               `gorm:"column:repair_base_server_seq;size:32" json:"repair_base_server_seq"`
	User                *User              `gorm:"belongsTo"`
}

func (Operation) TableName() string {
	return "operations"
}

type UserSyncState struct {
	UserID                     int            `gorm:"column:user_id;primaryKey;size:32" json:"user_id"`
	LastSeq                    int            `gorm:"column:last_seq;size:32;default:0;not null" json:"last_seq"`
	LastSnapshotSeq            *int           `gorm:"column:last_snapshot_seq;size:32" json:"last_snapshot_seq"`
	SnapshotData               []byte         `gorm:"column:snapshot_data" json:"snapshot_data"`
	SnapshotAt                 *int64         `gorm:"column:snapshot_at" json:"snapshot_at"`
	SnapshotSchemaVersion      *int           `gorm:"column:snapshot_schema_version;size:32;default:1" json:"snapshot_schema_version"`
	LatestFullStateSeq         *int           `gorm:"column:latest_full_state_seq;size:32" json:"latest_full_state_seq"`
	LatestFullStateVectorClock datatypes.JSON `gorm:"column:latest_full_state_vector_clock" json:"latest_full_state_vector_clock"`
	LatestStateReplacementSeq  *int           `gorm:"column:latest_state_replacement_seq;size:32" json:"latest_state_replacement_seq"`
	User                       *User          `gorm:"belongsTo"`
}

func (UserSyncState) TableName() string {
	return "user_sync_state"
}

type SyncDevice struct {
	ClientID   string  `gorm:"column:client_id;primaryKey" json:"client_id"`
	UserID     int     `gorm:"column:user_id;primaryKey;size:32" json:"user_id"`
	DeviceName *string `gorm:"column:device_name" json:"device_name"`
	UserAgent  *string `gorm:"column:user_agent" json:"user_agent"`
	// Bare semver the client last reported on download (`18.22.0`), NULL until
	// it does. Read by the checkpoint gate (#9962); never shown to users.
	AppVersion *string `gorm:"column:app_version" json:"app_version"`
	LastSeenAt int64   `gorm:"column:last_seen_at;not null" json:"last_seen_at"`
	// Reserved: written once at device registration, never read or updated.
	LastAckedSeq int   `gorm:"column:last_acked_seq;size:32;default:0;not null" json:"last_acked_seq"`
	CreatedAt    int64 `gorm:"column:created_at;autoCreateTime:false;not null" json:"created_at"`
	User         *User `gorm:"belongsTo"`
}

func (SyncDevice) TableName() string {
	return "sync_devices"
}

// Public half of the user's capture key. The row outlives switching the
// feature off: `disabledAt` stops a stale device re-registering that key.
type CaptureKey struct {
	UserID int `gorm:"column:user_id;primaryKey;size:32" json:"user_id"`
	// Base64 SPKI DER of an ECDH P-256 public key.
	PublicKey   string `gorm:"column:public_key;not null" json:"public_key"`
	Fingerprint string `gorm:"column:fingerprint;not null" json:"fingerprint"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime:false;not null" json:"created_at"`
	DisabledAt  *int64 `gorm:"column:disabled_at" json:"disabled_at"`
	User        *User  `gorm:"belongsTo"`
}

func (CaptureKey) TableName() string {
	return "capture_keys"
}

// Capture tokens and OAuth grants store only hashes of bearer credentials.
type CaptureGrant struct {
	ID               string  `gorm:"column:id;primaryKey" json:"id"`
	UserID           int     `gorm:"column:user_id;index:capture_grants_user_id_idx;size:32;not null" json:"user_id"`
	Kind             string  `gorm:"column:kind;not null" json:"kind"`
	ClientName       string  `gorm:"column:client_name;not null" json:"client_name"`
	SecretHash       string  `gorm:"column:secret_hash;not null" json:"secret_hash"`
	CreatedAt        int64   `gorm:"column:created_at;autoCreateTime:false;not null" json:"created_at"`
	ExpiresAt        int64   `gorm:"column:expires_at;not null" json:"expires_at"`
	LastUsedAt       *int64  `gorm:"column:last_used_at" json:"last_used_at"`
	RevokedAt        *int64  `gorm:"column:revoked_at" json:"revoked_at"`
	ClientID         *string `gorm:"column:client_id" json:"client_id"`
	Resource         *string `gorm:"column:resource" json:"resource"`
	RefreshHash      *string `gorm:"column:refresh_hash" json:"refresh_hash"`
	PrevRefreshHash  *string `gorm:"column:prev_refresh_hash" json:"prev_refresh_hash"`
	RotatedAt        *int64  `gorm:"column:rotated_at" json:"rotated_at"`
	RefreshExpiresAt *int64  `gorm:"column:refresh_expires_at" json:"refresh_expires_at"`
	User             *User   `gorm:"belongsTo"`
}

func (CaptureGrant) TableName() string {
	return "capture_grants"
}

// One sealed capture waiting for a device. No relation to its grant: items
// outlive a grant that is revoked without discarding them.
type CaptureItem struct {
	ID         string `gorm:"column:id;primaryKey" json:"id"`
	UserID     int    `gorm:"column:user_id;index:capture_items_user_id_created_at_idx;size:32;not null" json:"user_id"`
	GrantID    string `gorm:"column:grant_id;index:capture_items_grant_id_idx;not null" json:"grant_id"`
	Ciphertext string `gorm:"column:ciphertext;not null" json:"ciphertext"`
	// Task id the claiming device must use; replaced when another device takes over.
	TaskID        string  `gorm:"column:task_id;not null" json:"task_id"`
	CreatedAt     int64   `gorm:"column:created_at;index:capture_items_user_id_created_at_idx;autoCreateTime:false;not null" json:"created_at"`
	LeaseDeviceID *string `gorm:"column:lease_device_id" json:"lease_device_id"`
	LeaseUntil    *int64  `gorm:"column:lease_until" json:"lease_until"`
	User          *User   `gorm:"belongsTo"`
}

func (CaptureItem) TableName() string {
	return "capture_items"
}

type CaptureConnectCode struct {
	UserID       int    `gorm:"column:user_id;primaryKey;size:32" json:"user_id"`
	CodeHash     string `gorm:"column:code_hash;uniqueIndex:capture_connect_codes_code_hash_key;not null" json:"code_hash"`
	TokenVersion int    `gorm:"column:token_version;size:32;not null" json:"token_version"`
	ExpiresAt    int64  `gorm:"column:expires_at;not null" json:"expires_at"`
	User         *User  `gorm:"belongsTo"`
}

func (CaptureConnectCode) TableName() string {
	return "capture_connect_codes"
}

type CaptureAuth struct {
	ID            string  `gorm:"column:id;primaryKey" json:"id"`
	UserID        *int    `gorm:"column:user_id;index:capture_auth_user_id_idx;size:32" json:"user_id"`
	ClientID      string  `gorm:"column:client_id;not null" json:"client_id"`
	ClientName    string  `gorm:"column:client_name;not null" json:"client_name"`
	RedirectURI   string  `gorm:"column:redirect_uri;not null" json:"redirect_uri"`
	CodeChallenge string  `gorm:"column:code_challenge;not null" json:"code_challenge"`
	Resource      string  `gorm:"column:resource;not null" json:"resource"`
	State         string  `gorm:"column:state;not null" json:"state"`
	Scope         string  `gorm:"column:scope;not null" json:"scope"`
	CodeHash      *string `gorm:"column:code_hash;uniqueIndex:capture_auth_code_hash_key" json:"code_hash"`
	TokenVersion  *int    `gorm:"column:token_version;size:32" json:"token_version"`
	ExpiresAt     int64   `gorm:"column:expires_at;index:capture_auth_expires_at_idx;not null" json:"expires_at"`
	User          *User   `gorm:"belongsTo"`
}

func (CaptureAuth) TableName() string {
	return "capture_auth"
}

type ConnectRequest struct {
	ID             string  `gorm:"column:id;primaryKey" json:"id"`
	PollSecretHash string  `gorm:"column:poll_secret_hash;not null" json:"poll_secret_hash"`
	MatchCode      string  `gorm:"column:match_code;not null" json:"match_code"`
	DeviceName     string  `gorm:"column:device_name;not null" json:"device_name"`
	Platform       string  `gorm:"column:platform;not null" json:"platform"`
	Status         *string `gorm:"column:status;default:'pending';not null" json:"status"`
	UserID         *int    `gorm:"column:user_id;index:connect_requests_user_id_idx;size:32" json:"user_id"`
	TokenVersion   *int    `gorm:"column:token_version;size:32" json:"token_version"`
	ExpiresAt      int64   `gorm:"column:expires_at;index:connect_requests_expires_at_idx;not null" json:"expires_at"`
	User           *User   `gorm:"belongsTo"`
}

func (ConnectRequest) TableName() string {
	return "connect_requests"
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
