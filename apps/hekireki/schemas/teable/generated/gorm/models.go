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

type Space struct {
	ID                         string                    `gorm:"column:id;primaryKey" json:"id"`
	Name                       string                    `gorm:"column:name;not null" json:"name"`
	Avatar                     *string                   `gorm:"column:avatar" json:"avatar"`
	Credit                     *int                      `gorm:"column:credit;size:32" json:"credit"`
	DeletedTime                *DateTime                 `gorm:"column:deleted_time;type:timestamp(3)" json:"deleted_time"`
	CreatedTime                DateTime                  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy                  string                    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy             *string                   `gorm:"column:last_modified_by" json:"last_modified_by"`
	LastModifiedTime           *DateTime                 `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	IsTemplate                 *bool                     `gorm:"column:is_template" json:"is_template"`
	BaseGroup                  []Base                    `gorm:"foreignKey:SpaceID;constraint:base_space_id_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DataDbMigrationJobs        []SpaceDataDbMigrationJob `gorm:"foreignKey:SpaceID;constraint:space_data_db_migration_job_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BaseDataDbMoveJobsAsSource []BaseDataDbMoveJob       `gorm:"foreignKey:SourceSpaceID;constraint:base_data_db_move_job_source_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BaseDataDbMoveJobsAsTarget []BaseDataDbMoveJob       `gorm:"foreignKey:TargetSpaceID;constraint:base_data_db_move_job_target_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	DataDbBinding              *SpaceDataDbBinding       `gorm:"foreignKey:SpaceID;constraint:space_data_db_binding_space_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Space) TableName() string {
	return "space"
}

func (m *Space) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Space) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type DataDbConnection struct {
	ID                        string                    `gorm:"column:id;primaryKey" json:"id"`
	Provider                  string                    `gorm:"column:provider;default:'postgres';not null" json:"provider"`
	EncryptedURL              string                    `gorm:"column:encrypted_url;not null" json:"encrypted_url"`
	URLFingerprint            string                    `gorm:"column:url_fingerprint;uniqueIndex:data_db_connection_url_fingerprint_key;not null" json:"url_fingerprint"`
	DisplayHost               *string                   `gorm:"column:display_host" json:"display_host"`
	DisplayDatabase           *string                   `gorm:"column:display_database" json:"display_database"`
	InternalSchema            string                    `gorm:"column:internal_schema;not null" json:"internal_schema"`
	Status                    string                    `gorm:"column:status;default:'pending';not null" json:"status"`
	SchemaVersion             *string                   `gorm:"column:schema_version" json:"schema_version"`
	Capabilities              datatypes.JSON            `gorm:"column:capabilities" json:"capabilities"`
	LastValidatedAt           *DateTime                 `gorm:"column:last_validated_at;type:timestamp(3)" json:"last_validated_at"`
	LastError                 *string                   `gorm:"column:last_error" json:"last_error"`
	HealthState               string                    `gorm:"column:health_state;default:'healthy';not null" json:"health_state"`
	HealthReason              *string                   `gorm:"column:health_reason" json:"health_reason"`
	HealthChangedAt           *DateTime                 `gorm:"column:health_changed_at;type:timestamp(3)" json:"health_changed_at"`
	LastHealthCheckAt         *DateTime                 `gorm:"column:last_health_check_at;type:timestamp(3)" json:"last_health_check_at"`
	ConsecutiveHealthFailures int                       `gorm:"column:consecutive_health_failures;size:32;default:0;not null" json:"consecutive_health_failures"`
	CreatedBy                 string                    `gorm:"column:created_by;not null" json:"created_by"`
	CreatedTime               DateTime                  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime          *DateTime                 `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	SpaceBindings             []SpaceDataDbBinding      `gorm:"foreignKey:DataDbConnectionID;constraint:space_data_db_binding_data_db_connection_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	SourceMigrationJobs       []SpaceDataDbMigrationJob `gorm:"foreignKey:SourceConnectionID;constraint:space_data_db_migration_job_source_connection_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
	TargetMigrationJobs       []SpaceDataDbMigrationJob `gorm:"foreignKey:TargetConnectionID;constraint:space_data_db_migration_job_target_connection_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (DataDbConnection) TableName() string {
	return "data_db_connection"
}

func (m *DataDbConnection) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*DataDbConnection) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type SpaceDataDbBinding struct {
	ID                   string            `gorm:"column:id;primaryKey" json:"id"`
	SpaceID              string            `gorm:"column:space_id;uniqueIndex:space_data_db_binding_space_id_key;not null" json:"space_id"`
	DataDbConnectionID   *string           `gorm:"column:data_db_connection_id;index:space_data_db_binding_data_db_connection_id_idx" json:"data_db_connection_id"`
	Mode                 string            `gorm:"column:mode;default:'default';not null" json:"mode"`
	State                string            `gorm:"column:state;default:'ready';not null" json:"state"`
	CreatedBy            string            `gorm:"column:created_by;not null" json:"created_by"`
	CreatedTime          DateTime          `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime     *DateTime         `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastHistoryFlushedAt *DateTime         `gorm:"column:last_history_flushed_at;type:timestamp(3)" json:"last_history_flushed_at"`
	LastRemovalFlushedAt *DateTime         `gorm:"column:last_removal_flushed_at;type:timestamp(3)" json:"last_removal_flushed_at"`
	Space                *Space            `gorm:"belongsTo"`
	DataDbConnection     *DataDbConnection `gorm:"belongsTo"`
}

func (SpaceDataDbBinding) TableName() string {
	return "space_data_db_binding"
}

func (m *SpaceDataDbBinding) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*SpaceDataDbBinding) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type SpaceDataDbMigrationJob struct {
	ID                   string            `gorm:"column:id;primaryKey" json:"id"`
	SpaceID              string            `gorm:"column:space_id;index:space_data_db_migration_job_space_id_state_idx,priority:1;not null" json:"space_id"`
	SourceConnectionID   *string           `gorm:"column:source_connection_id;index:space_data_db_migration_job_source_connection_id_idx" json:"source_connection_id"`
	TargetConnectionID   *string           `gorm:"column:target_connection_id;index:space_data_db_migration_job_target_connection_id_idx" json:"target_connection_id"`
	TargetMode           *string           `gorm:"column:target_mode;default:'migrate-space';not null" json:"target_mode"`
	SwitchOnCompletion   bool              `gorm:"column:switch_on_completion;default:false;not null" json:"switch_on_completion"`
	State                string            `gorm:"column:state;index:space_data_db_migration_job_space_id_state_idx,priority:2;default:'pending';not null" json:"state"`
	TargetURLFingerprint string            `gorm:"column:target_url_fingerprint;not null" json:"target_url_fingerprint"`
	TargetInternalSchema string            `gorm:"column:target_internal_schema;not null" json:"target_internal_schema"`
	Inventory            datatypes.JSON    `gorm:"column:inventory" json:"inventory"`
	CopyStats            datatypes.JSON    `gorm:"column:copy_stats" json:"copy_stats"`
	ValidationStats      datatypes.JSON    `gorm:"column:validation_stats" json:"validation_stats"`
	LastError            *string           `gorm:"column:last_error" json:"last_error"`
	StartedAt            *DateTime         `gorm:"column:started_at;type:timestamp(3)" json:"started_at"`
	CompletedAt          *DateTime         `gorm:"column:completed_at;type:timestamp(3)" json:"completed_at"`
	CreatedBy            string            `gorm:"column:created_by;not null" json:"created_by"`
	CreatedTime          DateTime          `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime     *DateTime         `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	Space                *Space            `gorm:"belongsTo"`
	SourceConnection     *DataDbConnection `gorm:"foreignKey:SourceConnectionID;belongsTo"`
	TargetConnection     *DataDbConnection `gorm:"foreignKey:TargetConnectionID;belongsTo"`
}

func (SpaceDataDbMigrationJob) TableName() string {
	return "space_data_db_migration_job"
}

func (m *SpaceDataDbMigrationJob) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*SpaceDataDbMigrationJob) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type BaseDataDbMoveJob struct {
	ID                 string         `gorm:"column:id;primaryKey" json:"id"`
	BaseID             string         `gorm:"column:base_id;index:base_data_db_move_job_base_id_state_idx,priority:1;not null" json:"base_id"`
	SourceSpaceID      string         `gorm:"column:source_space_id;index:base_data_db_move_job_source_space_id_state_idx,priority:1;not null" json:"source_space_id"`
	TargetSpaceID      string         `gorm:"column:target_space_id;index:base_data_db_move_job_target_space_id_state_idx,priority:1;not null" json:"target_space_id"`
	SourceConnectionID *string        `gorm:"column:source_connection_id" json:"source_connection_id"`
	TargetConnectionID *string        `gorm:"column:target_connection_id" json:"target_connection_id"`
	State              string         `gorm:"column:state;index:base_data_db_move_job_base_id_state_idx,priority:2;index:base_data_db_move_job_state_idx;index:base_data_db_move_job_source_space_id_state_idx,priority:2;index:base_data_db_move_job_target_space_id_state_idx,priority:2;default:'pending';not null" json:"state"`
	Inventory          datatypes.JSON `gorm:"column:inventory" json:"inventory"`
	CopyStats          datatypes.JSON `gorm:"column:copy_stats" json:"copy_stats"`
	ValidationStats    datatypes.JSON `gorm:"column:validation_stats" json:"validation_stats"`
	LastError          *string        `gorm:"column:last_error" json:"last_error"`
	StartedAt          *DateTime      `gorm:"column:started_at;type:timestamp(3)" json:"started_at"`
	CompletedAt        *DateTime      `gorm:"column:completed_at;type:timestamp(3)" json:"completed_at"`
	CreatedBy          string         `gorm:"column:created_by;not null" json:"created_by"`
	CreatedTime        DateTime       `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime   *DateTime      `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	Base               *Base          `gorm:"belongsTo"`
	SourceSpace        *Space         `gorm:"foreignKey:SourceSpaceID;belongsTo"`
	TargetSpace        *Space         `gorm:"foreignKey:TargetSpaceID;belongsTo"`
}

func (BaseDataDbMoveJob) TableName() string {
	return "base_data_db_move_job"
}

func (m *BaseDataDbMoveJob) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*BaseDataDbMoveJob) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type PinResource struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	Type        string   `gorm:"column:type;not null" json:"type"`
	ResourceID  string   `gorm:"column:resource_id;uniqueIndex:pin_resource_created_by_resource_id_key,priority:2;not null" json:"resource_id"`
	CreatedTime DateTime `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy   string   `gorm:"column:created_by;uniqueIndex:pin_resource_created_by_resource_id_key,priority:1;not null" json:"created_by"`
	Order       float64  `gorm:"column:order;index:pin_resource_order_idx;type:double precision;not null" json:"order"`
}

func (PinResource) TableName() string {
	return "pin_resource"
}

func (m *PinResource) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type Base struct {
	ID               string              `gorm:"column:id;primaryKey" json:"id"`
	SpaceID          string              `gorm:"column:space_id;index:base_space_id_idx;not null" json:"space_id"`
	Name             *string             `gorm:"column:name" json:"name"`
	Order            float64             `gorm:"column:order;index:base_order_idx;type:double precision;not null" json:"order"`
	Icon             *string             `gorm:"column:icon" json:"icon"`
	SchemaPass       *string             `gorm:"column:schema_pass" json:"schema_pass"`
	V2Enabled        bool                `gorm:"column:v2_enabled;default:false;not null" json:"v2_enabled"`
	ProvisionState   string              `gorm:"column:provision_state;default:'ready';not null" json:"provision_state"`
	DeletedTime      *DateTime           `gorm:"column:deleted_time;type:timestamp(3)" json:"deleted_time"`
	CreatedTime      DateTime            `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string              `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy   *string             `gorm:"column:last_modified_by" json:"last_modified_by"`
	LastModifiedTime *DateTime           `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	Space            *Space              `gorm:"belongsTo"`
	Tables           []TableMeta         `gorm:"foreignKey:BaseID;constraint:table_meta_base_id_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	DataDbMoveJobs   []BaseDataDbMoveJob `gorm:"foreignKey:BaseID;constraint:base_data_db_move_job_base_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserOrders       []UserBaseOrder     `gorm:"foreignKey:BaseID;constraint:user_base_order_base_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Base) TableName() string {
	return "base"
}

func (m *Base) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Base) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type TableMeta struct {
	ID                  string              `gorm:"column:id;primaryKey" json:"id"`
	BaseID              string              `gorm:"column:base_id;index:table_meta_base_id_deleted_time_idx,priority:1;not null" json:"base_id"`
	Name                string              `gorm:"column:name;not null" json:"name"`
	Description         *string             `gorm:"column:description" json:"description"`
	Icon                *string             `gorm:"column:icon" json:"icon"`
	DbTableName         string              `gorm:"column:db_table_name;index:table_meta_db_table_name_idx;not null" json:"db_table_name"`
	DbViewName          *string             `gorm:"column:db_view_name" json:"db_view_name"`
	SearchIndex         datatypes.JSON      `gorm:"column:search_index;type:jsonb" json:"search_index"`
	SearchIndexPolicy   datatypes.JSON      `gorm:"column:search_index_policy;type:jsonb" json:"search_index_policy"`
	QueryIndexPolicy    datatypes.JSON      `gorm:"column:query_index_policy;type:jsonb" json:"query_index_policy"`
	ProvisionState      string              `gorm:"column:provision_state;default:'ready';not null" json:"provision_state"`
	ActorStorageVersion int                 `gorm:"column:actor_storage_version;size:32;default:0;not null" json:"actor_storage_version"`
	Version             int                 `gorm:"column:version;size:32;not null" json:"version"`
	Order               float64             `gorm:"column:order;index:table_meta_order_idx;type:double precision;not null" json:"order"`
	CreatedTime         DateTime            `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime    *DateTime           `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	DeletedTime         *DateTime           `gorm:"column:deleted_time;index:table_meta_base_id_deleted_time_idx,priority:2;type:timestamp(3)" json:"deleted_time"`
	CreatedBy           string              `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy      *string             `gorm:"column:last_modified_by" json:"last_modified_by"`
	Base                *Base               `gorm:"belongsTo"`
	Fields              []Field             `gorm:"foreignKey:TableID;constraint:field_table_id_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Views               []View              `gorm:"foreignKey:TableID;constraint:view_table_id_fkey,OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PluginPanel         []PluginPanel       `gorm:"foreignKey:TableID;constraint:plugin_panel_table_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	PluginContextMenu   []PluginContextMenu `gorm:"foreignKey:TableID;constraint:plugin_context_menu_table_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TableMeta) TableName() string {
	return "table_meta"
}

func (m *TableMeta) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*TableMeta) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type TableQueryObservationShard struct {
	SpaceID           *string        `gorm:"column:space_id" json:"space_id"`
	BaseID            string         `gorm:"column:base_id;index:table_query_observation_shard_base_start_idx,priority:1;not null" json:"base_id"`
	TableID           string         `gorm:"column:table_id;primaryKey;index:table_query_observation_shard_table_start_idx,priority:1;index:table_query_observation_shard_search_activity_idx,priority:2" json:"table_id"`
	QueryKind         string         `gorm:"column:query_kind;primaryKey;index:table_query_observation_shard_search_activity_idx,priority:1" json:"query_kind"`
	ShapeHash         string         `gorm:"column:shape_hash;primaryKey" json:"shape_hash"`
	WindowStart       DateTime       `gorm:"column:window_start;primaryKey;index:table_query_observation_shard_window_start_idx;index:table_query_observation_shard_table_start_idx,priority:2;index:table_query_observation_shard_search_activity_idx,priority:3;index:table_query_observation_shard_base_start_idx,priority:2;type:timestamptz(6)" json:"window_start"`
	WriterID          string         `gorm:"column:writer_id;primaryKey" json:"writer_id"`
	WindowSizeSeconds int            `gorm:"column:window_size_seconds;size:32;not null" json:"window_size_seconds"`
	RequestCount      int            `gorm:"column:request_count;size:32;not null" json:"request_count"`
	SlowCount         int            `gorm:"column:slow_count;size:32;not null" json:"slow_count"`
	TimeoutCount      int            `gorm:"column:timeout_count;size:32;not null" json:"timeout_count"`
	DbErrorCount      int            `gorm:"column:db_error_count;size:32;not null" json:"db_error_count"`
	TotalDurationMs   float64        `gorm:"column:total_duration_ms;type:double precision;not null" json:"total_duration_ms"`
	MaxDurationMs     float64        `gorm:"column:max_duration_ms;type:double precision;not null" json:"max_duration_ms"`
	TotalDbDurationMs *float64       `gorm:"column:total_db_duration_ms;type:double precision" json:"total_db_duration_ms"`
	MaxDbDurationMs   *float64       `gorm:"column:max_db_duration_ms;type:double precision" json:"max_db_duration_ms"`
	Shape             datatypes.JSON `gorm:"column:shape;not null" json:"shape"`
	SQLDiagnostics    datatypes.JSON `gorm:"column:sql_diagnostics" json:"sql_diagnostics"`
	CreatedTime       DateTime       `gorm:"column:created_time;type:timestamptz(6);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime  *DateTime      `gorm:"column:last_modified_time;type:timestamptz(6)" json:"last_modified_time"`
}

func (TableQueryObservationShard) TableName() string {
	return "table_query_observation_shard"
}

func (m *TableQueryObservationShard) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type SchemaOperation struct {
	ID               string         `gorm:"column:id;primaryKey" json:"id"`
	Type             string         `gorm:"column:type;not null" json:"type"`
	Status           string         `gorm:"column:status;index:schema_operation_status_next_run_at_idx,priority:1;index:schema_operation_resource_status_idx,priority:3;index:schema_operation_base_status_idx,priority:2;index:schema_operation_table_status_idx,priority:2;not null" json:"status"`
	Phase            string         `gorm:"column:phase;not null" json:"phase"`
	ResourceType     string         `gorm:"column:resource_type;index:schema_operation_resource_status_idx,priority:1;not null" json:"resource_type"`
	ResourceID       string         `gorm:"column:resource_id;index:schema_operation_resource_status_idx,priority:2;not null" json:"resource_id"`
	BaseID           *string        `gorm:"column:base_id;index:schema_operation_base_status_idx,priority:1" json:"base_id"`
	TableID          *string        `gorm:"column:table_id;index:schema_operation_table_status_idx,priority:1" json:"table_id"`
	IdempotencyKey   string         `gorm:"column:idempotency_key;uniqueIndex:schema_operation_idempotency_key_key;not null" json:"idempotency_key"`
	Payload          datatypes.JSON `gorm:"column:payload" json:"payload"`
	Result           datatypes.JSON `gorm:"column:result" json:"result"`
	Attempts         int            `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	MaxAttempts      *int           `gorm:"column:max_attempts;size:32;default:8;not null" json:"max_attempts"`
	NextRunAt        DateTime       `gorm:"column:next_run_at;index:schema_operation_status_next_run_at_idx,priority:2;type:timestamptz(6);default:CURRENT_TIMESTAMP;not null" json:"next_run_at"`
	LockedAt         *DateTime      `gorm:"column:locked_at;type:timestamptz(6)" json:"locked_at"`
	LockedBy         *string        `gorm:"column:locked_by" json:"locked_by"`
	LastError        *string        `gorm:"column:last_error" json:"last_error"`
	CreatedTime      DateTime       `gorm:"column:created_time;type:timestamptz(6);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string         `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime *DateTime      `gorm:"column:last_modified_time;type:timestamptz(6)" json:"last_modified_time"`
	LastModifiedBy   *string        `gorm:"column:last_modified_by" json:"last_modified_by"`
}

func (SchemaOperation) TableName() string {
	return "schema_operation"
}

func (m *SchemaOperation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.NextRunAt.IsZero() {
		m.NextRunAt = DateTime{Time: now}
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*SchemaOperation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Field struct {
	ID                   string                     `gorm:"column:id;primaryKey" json:"id"`
	Name                 string                     `gorm:"column:name;not null" json:"name"`
	Description          *string                    `gorm:"column:description" json:"description"`
	Options              *string                    `gorm:"column:options" json:"options"`
	Meta                 *string                    `gorm:"column:meta" json:"meta"`
	AiConfig             *string                    `gorm:"column:ai_config" json:"ai_config"`
	Type                 string                     `gorm:"column:type;not null" json:"type"`
	CellValueType        string                     `gorm:"column:cell_value_type;not null" json:"cell_value_type"`
	IsMultipleCellValue  *bool                      `gorm:"column:is_multiple_cell_value" json:"is_multiple_cell_value"`
	DbFieldType          string                     `gorm:"column:db_field_type;not null" json:"db_field_type"`
	DbFieldName          string                     `gorm:"column:db_field_name;not null" json:"db_field_name"`
	ProvisionState       string                     `gorm:"column:provision_state;default:'ready';not null" json:"provision_state"`
	NotNull              *bool                      `gorm:"column:not_null" json:"not_null"`
	Unique               *bool                      `gorm:"column:unique" json:"unique"`
	IsPrimary            *bool                      `gorm:"column:is_primary" json:"is_primary"`
	IsComputed           *bool                      `gorm:"column:is_computed" json:"is_computed"`
	IsLookup             *bool                      `gorm:"column:is_lookup" json:"is_lookup"`
	IsConditionalLookup  *bool                      `gorm:"column:is_conditional_lookup" json:"is_conditional_lookup"`
	IsPending            *bool                      `gorm:"column:is_pending" json:"is_pending"`
	HasError             *bool                      `gorm:"column:has_error" json:"has_error"`
	NeedsPeriodicRefresh bool                       `gorm:"column:needs_periodic_refresh;default:false;not null" json:"needs_periodic_refresh"`
	RefreshCadence       *string                    `gorm:"column:refresh_cadence;type:varchar(16)" json:"refresh_cadence"`
	LookupLinkedFieldID  *string                    `gorm:"column:lookup_linked_field_id;index:field_lookup_linked_field_id_idx" json:"lookup_linked_field_id"`
	LookupOptions        *string                    `gorm:"column:lookup_options" json:"lookup_options"`
	TableID              string                     `gorm:"column:table_id;index:field_table_id_deleted_time_idx,priority:1;not null" json:"table_id"`
	Order                float64                    `gorm:"column:order;type:double precision;not null" json:"order"`
	Version              int                        `gorm:"column:version;size:32;not null" json:"version"`
	CreatedTime          DateTime                   `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime     *DateTime                  `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	DeletedTime          *DateTime                  `gorm:"column:deleted_time;index:field_table_id_deleted_time_idx,priority:2;type:timestamp(3)" json:"deleted_time"`
	CreatedBy            string                     `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy       *string                    `gorm:"column:last_modified_by" json:"last_modified_by"`
	Table                *TableMeta                 `gorm:"foreignKey:TableID;belongsTo"`
	PeriodicRefreshState *FieldPeriodicRefreshState `gorm:"foreignKey:FieldID;constraint:field_periodic_refresh_state_field_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Field) TableName() string {
	return "field"
}

func (m *Field) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Field) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type DomainEventOutbox struct {
	ID                string         `gorm:"column:id;primaryKey" json:"id"`
	BaseID            string         `gorm:"column:base_id;index:domain_event_outbox_base_id_idx;not null" json:"base_id"`
	TableID           *string        `gorm:"column:table_id" json:"table_id"`
	MessageName       string         `gorm:"column:message_name;not null" json:"message_name"`
	SchemaVersion     int            `gorm:"column:schema_version;size:32;not null" json:"schema_version"`
	AggregateID       *string        `gorm:"column:aggregate_id" json:"aggregate_id"`
	Payload           datatypes.JSON `gorm:"column:payload;not null" json:"payload"`
	PayloadBytes      int            `gorm:"column:payload_bytes;size:32;not null" json:"payload_bytes"`
	CatalogGeneration int            `gorm:"column:catalog_generation;size:32;not null" json:"catalog_generation"`
	RequiredConsumers datatypes.JSON `gorm:"column:required_consumers;not null" json:"required_consumers"`
	BindingID         *string        `gorm:"column:binding_id" json:"binding_id"`
	StorageEpoch      *int           `gorm:"column:storage_epoch;size:32" json:"storage_epoch"`
	Unpublished       *bool          `gorm:"column:unpublished;default:true;not null" json:"unpublished"`
	Settled           *string        `gorm:"column:settled" json:"settled"`
	SettledAt         *DateTime      `gorm:"column:settled_at;type:timestamp(3)" json:"settled_at"`
	CreatedAt         DateTime       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
}

func (DomainEventOutbox) TableName() string {
	return "domain_event_outbox"
}

func (m *DomainEventOutbox) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DomainEventDelivery struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`
	EventID        string    `gorm:"column:event_id;uniqueIndex:domain_event_delivery_event_id_consumer_id_key,priority:1;not null" json:"event_id"`
	ConsumerID     string    `gorm:"column:consumer_id;uniqueIndex:domain_event_delivery_event_id_consumer_id_key,priority:2;not null" json:"consumer_id"`
	Status         string    `gorm:"column:status;not null" json:"status"`
	Attempts       int       `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	MaxAttempts    *int      `gorm:"column:max_attempts;size:32;default:12;not null" json:"max_attempts"`
	LeaseToken     *string   `gorm:"column:lease_token" json:"lease_token"`
	LeaseExpiresAt *DateTime `gorm:"column:lease_expires_at;type:timestamptz(3)" json:"lease_expires_at"`
	NextAttemptAt  DateTime  `gorm:"column:next_attempt_at;type:timestamptz(3);default:CURRENT_TIMESTAMP;not null" json:"next_attempt_at"`
	LastError      *string   `gorm:"column:last_error" json:"last_error"`
	CreatedAt      DateTime  `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
}

func (DomainEventDelivery) TableName() string {
	return "domain_event_delivery"
}

func (m *DomainEventDelivery) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.NextAttemptAt.IsZero() {
		m.NextAttemptAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type DomainEventInbox struct {
	ConsumerID string   `gorm:"column:consumer_id;primaryKey" json:"consumer_id"`
	EventID    string   `gorm:"column:event_id;primaryKey;index:domain_event_inbox_event_id_idx" json:"event_id"`
	CreatedAt  DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
}

func (DomainEventInbox) TableName() string {
	return "domain_event_inbox"
}

func (m *DomainEventInbox) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type FieldPeriodicRefreshState struct {
	FieldID                string    `gorm:"column:field_id;primaryKey" json:"field_id"`
	TableID                string    `gorm:"column:table_id;index:field_periodic_refresh_state_table_id_idx;not null" json:"table_id"`
	BaseID                 string    `gorm:"column:base_id;index:field_periodic_refresh_state_base_id_enabled_idx,priority:1;not null" json:"base_id"`
	Enabled                *bool     `gorm:"column:enabled;index:field_periodic_refresh_state_base_id_enabled_idx,priority:2;index:field_periodic_refresh_state_enabled_local_date_idx,priority:1;default:true;not null" json:"enabled"`
	TimeZone               *string   `gorm:"column:time_zone;type:varchar(64);default:'UTC';not null" json:"time_zone"`
	ContainsNowPrecision   bool      `gorm:"column:contains_now_precision;default:false;not null" json:"contains_now_precision"`
	LastRefreshedAt        *DateTime `gorm:"column:last_refreshed_at;type:timestamptz(6)" json:"last_refreshed_at"`
	LastRefreshedLocalDate *string   `gorm:"column:last_refreshed_local_date;index:field_periodic_refresh_state_enabled_local_date_idx,priority:2;type:varchar(16)" json:"last_refreshed_local_date"`
	LastRangeHash          *string   `gorm:"column:last_range_hash;type:varchar(128)" json:"last_range_hash"`
	LastStatus             *string   `gorm:"column:last_status;type:varchar(32);default:'idle';not null" json:"last_status"`
	LastError              *string   `gorm:"column:last_error" json:"last_error"`
	PendingTaskID          *string   `gorm:"column:pending_task_id;type:varchar(64)" json:"pending_task_id"`
	PendingLocalDate       *string   `gorm:"column:pending_local_date;type:varchar(16)" json:"pending_local_date"`
	PendingRangeHash       *string   `gorm:"column:pending_range_hash;type:varchar(128)" json:"pending_range_hash"`
	PendingStartedAt       *DateTime `gorm:"column:pending_started_at;type:timestamptz(6)" json:"pending_started_at"`
	CreatedTime            DateTime  `gorm:"column:created_time;type:timestamptz(6);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime       *DateTime `gorm:"column:last_modified_time;type:timestamptz(6)" json:"last_modified_time"`
	Field                  *Field    `gorm:"belongsTo"`
}

func (FieldPeriodicRefreshState) TableName() string {
	return "field_periodic_refresh_state"
}

func (m *FieldPeriodicRefreshState) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*FieldPeriodicRefreshState) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type ComputedUpdateOutbox struct {
	ID                       string                     `gorm:"column:id;primaryKey" json:"id"`
	BaseID                   string                     `gorm:"column:base_id;index:computed_update_outbox_base_id_seed_table_id_idx,priority:1;not null" json:"base_id"`
	WorkerLane               *string                    `gorm:"column:worker_lane;index:computed_update_outbox_worker_lane_status_next_run_at_idx,priority:1;default:'production';not null" json:"worker_lane"`
	SeedTableID              string                     `gorm:"column:seed_table_id;index:computed_update_outbox_base_id_seed_table_id_idx,priority:2;not null" json:"seed_table_id"`
	SeedRecordIds            datatypes.JSON             `gorm:"column:seed_record_ids" json:"seed_record_ids"`
	ChangeType               string                     `gorm:"column:change_type;not null" json:"change_type"`
	Steps                    datatypes.JSON             `gorm:"column:steps" json:"steps"`
	Edges                    datatypes.JSON             `gorm:"column:edges" json:"edges"`
	Status                   string                     `gorm:"column:status;index:computed_update_outbox_status_next_run_at_idx,priority:1;index:computed_update_outbox_worker_lane_status_next_run_at_idx,priority:2;not null" json:"status"`
	Attempts                 int                        `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	MaxAttempts              *int                       `gorm:"column:max_attempts;size:32;default:8;not null" json:"max_attempts"`
	NextRunAt                DateTime                   `gorm:"column:next_run_at;index:computed_update_outbox_status_next_run_at_idx,priority:2;index:computed_update_outbox_worker_lane_status_next_run_at_idx,priority:3;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"next_run_at"`
	LockedAt                 *DateTime                  `gorm:"column:locked_at;type:timestamp(3)" json:"locked_at"`
	LockedBy                 *string                    `gorm:"column:locked_by" json:"locked_by"`
	LastError                *string                    `gorm:"column:last_error" json:"last_error"`
	EstimatedComplexity      int                        `gorm:"column:estimated_complexity;size:32;default:0;not null" json:"estimated_complexity"`
	PlanHash                 string                     `gorm:"column:plan_hash;index:computed_update_outbox_plan_hash_idx;not null" json:"plan_hash"`
	DirtyStats               datatypes.JSON             `gorm:"column:dirty_stats" json:"dirty_stats"`
	RunID                    string                     `gorm:"column:run_id;index:computed_update_outbox_run_id_idx;not null" json:"run_id"`
	OriginRunIds             ScalarList[string]         `gorm:"column:origin_run_ids;type:text[]" json:"origin_run_ids"`
	RunTotalSteps            int                        `gorm:"column:run_total_steps;size:32;default:0;not null" json:"run_total_steps"`
	RunCompletedStepsBefore  int                        `gorm:"column:run_completed_steps_before;size:32;default:0;not null" json:"run_completed_steps_before"`
	AffectedTableIds         ScalarList[string]         `gorm:"column:affected_table_ids;type:text[]" json:"affected_table_ids"`
	AffectedFieldIds         ScalarList[string]         `gorm:"column:affected_field_ids;type:text[]" json:"affected_field_ids"`
	SyncMaxLevel             *int                       `gorm:"column:sync_max_level;size:32" json:"sync_max_level"`
	SourceChangedAt          *DateTime                  `gorm:"column:source_changed_at;type:timestamp(3)" json:"source_changed_at"`
	StageDepth               int                        `gorm:"column:stage_depth;size:32;default:0;not null" json:"stage_depth"`
	PredecessorTaskID        *string                    `gorm:"column:predecessor_task_id" json:"predecessor_task_id"`
	DeadLetterRestoreVersion *int                       `gorm:"column:dead_letter_restore_version;size:32" json:"dead_letter_restore_version"`
	CreatedAt                DateTime                   `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt                DateTime                   `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	Seeds                    []ComputedUpdateOutboxSeed `gorm:"foreignKey:TaskID;constraint:computed_update_outbox_seed_task_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ComputedUpdateOutbox) TableName() string {
	return "computed_update_outbox"
}

func (m *ComputedUpdateOutbox) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.NextRunAt.IsZero() {
		m.NextRunAt = DateTime{Time: now}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ComputedUpdateOutbox) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

type ComputedUpdateOutboxSeed struct {
	ID       string                `gorm:"column:id;primaryKey" json:"id"`
	TaskID   string                `gorm:"column:task_id;uniqueIndex:computed_update_outbox_seed_task_id_table_id_record_id_key,priority:1;index:computed_update_outbox_seed_task_id_idx;not null" json:"task_id"`
	TableID  string                `gorm:"column:table_id;uniqueIndex:computed_update_outbox_seed_task_id_table_id_record_id_key,priority:2;not null" json:"table_id"`
	RecordID string                `gorm:"column:record_id;uniqueIndex:computed_update_outbox_seed_task_id_table_id_record_id_key,priority:3;not null" json:"record_id"`
	Task     *ComputedUpdateOutbox `gorm:"foreignKey:TaskID;belongsTo"`
}

func (ComputedUpdateOutboxSeed) TableName() string {
	return "computed_update_outbox_seed"
}

func (m *ComputedUpdateOutboxSeed) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

// Actual output-value evidence; independent from continuation progress.
type ComputedUpdateChangeFrontier struct {
	ScopeID  string `gorm:"column:scope_id;primaryKey" json:"scope_id"`
	Kind     string `gorm:"column:kind;primaryKey" json:"kind"`
	TableID  string `gorm:"column:table_id;primaryKey" json:"table_id"`
	RecordID string `gorm:"column:record_id;primaryKey" json:"record_id"`
	FieldID  string `gorm:"column:field_id;primaryKey" json:"field_id"`
}

func (ComputedUpdateChangeFrontier) TableName() string {
	return "computed_update_change_frontier"
}

// Durable per-stage state for budget-staged computed updates, keyed by the
// continuation chain's root task id (its scope). Rows are written once per
// record and never copied between continuation tasks:
// - kind 'excluded': processed-target exclusion ledger for partial batches;
// - kind 'frontier': seq-ordered queue of sources whose outgoing propagation
// is not finished (self-referential generations + migrated explicit seeds);
// - kind 'consumed': retired frontier sources preserved for deferred edge
// chunks, handed to the continuation as seeds at stage completion.
// Cleared when the chain's stage completes or the chain dead-letters.
type ComputedUpdateStageLedger struct {
	ScopeID  string `gorm:"column:scope_id;primaryKey;index:computed_update_stage_ledger_scope_id_kind_seq_idx,priority:1" json:"scope_id"`
	Kind     string `gorm:"column:kind;primaryKey;index:computed_update_stage_ledger_scope_id_kind_seq_idx,priority:2" json:"kind"`
	TableID  string `gorm:"column:table_id;primaryKey" json:"table_id"`
	RecordID string `gorm:"column:record_id;primaryKey" json:"record_id"`
	Seq      int64  `gorm:"column:seq;index:computed_update_stage_ledger_scope_id_kind_seq_idx,priority:3;default:0;not null" json:"seq"`
}

func (ComputedUpdateStageLedger) TableName() string {
	return "computed_update_stage_ledger"
}

type ComputedUpdateDeadLetter struct {
	ID                      string                         `gorm:"column:id;primaryKey" json:"id"`
	BaseID                  string                         `gorm:"column:base_id;index:computed_update_dead_letter_base_id_seed_table_id_idx,priority:1;not null" json:"base_id"`
	WorkerLane              *string                        `gorm:"column:worker_lane;default:'production';not null" json:"worker_lane"`
	SeedTableID             string                         `gorm:"column:seed_table_id;index:computed_update_dead_letter_base_id_seed_table_id_idx,priority:2;not null" json:"seed_table_id"`
	SeedRecordIds           datatypes.JSON                 `gorm:"column:seed_record_ids" json:"seed_record_ids"`
	ChangeType              string                         `gorm:"column:change_type;not null" json:"change_type"`
	Steps                   datatypes.JSON                 `gorm:"column:steps" json:"steps"`
	Edges                   datatypes.JSON                 `gorm:"column:edges" json:"edges"`
	Status                  string                         `gorm:"column:status;not null" json:"status"`
	Attempts                int                            `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	MaxAttempts             *int                           `gorm:"column:max_attempts;size:32;default:8;not null" json:"max_attempts"`
	NextRunAt               DateTime                       `gorm:"column:next_run_at;type:timestamp(3);not null" json:"next_run_at"`
	LockedAt                *DateTime                      `gorm:"column:locked_at;type:timestamp(3)" json:"locked_at"`
	LockedBy                *string                        `gorm:"column:locked_by" json:"locked_by"`
	LastError               *string                        `gorm:"column:last_error" json:"last_error"`
	EstimatedComplexity     int                            `gorm:"column:estimated_complexity;size:32;default:0;not null" json:"estimated_complexity"`
	PlanHash                string                         `gorm:"column:plan_hash;index:computed_update_dead_letter_plan_hash_idx;not null" json:"plan_hash"`
	DirtyStats              datatypes.JSON                 `gorm:"column:dirty_stats" json:"dirty_stats"`
	RunID                   string                         `gorm:"column:run_id;index:computed_update_dead_letter_run_id_idx;not null" json:"run_id"`
	OriginRunIds            ScalarList[string]             `gorm:"column:origin_run_ids;type:text[]" json:"origin_run_ids"`
	RunTotalSteps           int                            `gorm:"column:run_total_steps;size:32;default:0;not null" json:"run_total_steps"`
	RunCompletedStepsBefore int                            `gorm:"column:run_completed_steps_before;size:32;default:0;not null" json:"run_completed_steps_before"`
	AffectedTableIds        ScalarList[string]             `gorm:"column:affected_table_ids;type:text[]" json:"affected_table_ids"`
	AffectedFieldIds        ScalarList[string]             `gorm:"column:affected_field_ids;type:text[]" json:"affected_field_ids"`
	SyncMaxLevel            *int                           `gorm:"column:sync_max_level;size:32" json:"sync_max_level"`
	SourceChangedAt         *DateTime                      `gorm:"column:source_changed_at;type:timestamp(3)" json:"source_changed_at"`
	StageDepth              int                            `gorm:"column:stage_depth;size:32;default:0;not null" json:"stage_depth"`
	PredecessorTaskID       *string                        `gorm:"column:predecessor_task_id" json:"predecessor_task_id"`
	TraceData               datatypes.JSON                 `gorm:"column:trace_data" json:"trace_data"`
	FailedAt                DateTime                       `gorm:"column:failed_at;type:timestamp(3);not null" json:"failed_at"`
	CreatedAt               DateTime                       `gorm:"column:created_at;type:timestamp(3);autoCreateTime:false;not null" json:"created_at"`
	UpdatedAt               DateTime                       `gorm:"column:updated_at;type:timestamp(3);autoUpdateTime:false;not null" json:"updated_at"`
	DeleteAuthorizedTxid    *int64                         `gorm:"column:delete_authorized_txid" json:"delete_authorized_txid"`
	Seeds                   []ComputedUpdateDeadLetterSeed `gorm:"foreignKey:TaskID;constraint:computed_update_dead_letter_seed_task_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ComputedUpdateDeadLetter) TableName() string {
	return "computed_update_dead_letter"
}

type ComputedUpdateDeadLetterSeed struct {
	TaskID   string                    `gorm:"column:task_id;primaryKey" json:"task_id"`
	TableID  string                    `gorm:"column:table_id;primaryKey" json:"table_id"`
	RecordID string                    `gorm:"column:record_id;primaryKey" json:"record_id"`
	Task     *ComputedUpdateDeadLetter `gorm:"foreignKey:TaskID;belongsTo"`
}

func (ComputedUpdateDeadLetterSeed) TableName() string {
	return "computed_update_dead_letter_seed"
}

// Completion ledger for computed-update outbox tasks (lineage / latency
// tracing). Successful tasks are hard-deleted from the outbox on markDone,
// so this table keeps a retention-bounded trail per completed task: the
// run/stage chain (runId, stageDepth, predecessorTaskId) and the timeline
// from source mutation to converged write (sourceChangedAt → completedAt).
// Observability only — the worker never reads it; rows are pruned after a
// retention window.
type ComputedUpdateRunHistory struct {
	TaskID           string             `gorm:"column:task_id;primaryKey" json:"task_id"`
	BaseID           string             `gorm:"column:base_id;index:computed_update_run_history_base_id_completed_at_idx,priority:1;not null" json:"base_id"`
	SeedTableID      string             `gorm:"column:seed_table_id;not null" json:"seed_table_id"`
	ChangeType       string             `gorm:"column:change_type;not null" json:"change_type"`
	RunID            string             `gorm:"column:run_id;index:computed_update_run_history_run_id_idx;not null" json:"run_id"`
	OriginRunIds     ScalarList[string] `gorm:"column:origin_run_ids;index:computed_update_run_history_origin_run_ids_gin,type:gin;type:text[]" json:"origin_run_ids"`
	Steps            datatypes.JSON     `gorm:"column:steps" json:"steps"`
	Edges            datatypes.JSON     `gorm:"column:edges" json:"edges"`
	AffectedTableIds ScalarList[string] `gorm:"column:affected_table_ids;type:text[]" json:"affected_table_ids"`
	AffectedFieldIds ScalarList[string] `gorm:"column:affected_field_ids;type:text[]" json:"affected_field_ids"`
	// Source (changed) field ids; populated for seed tasks, where
	// affected_field_ids carries the trigger fields rather than targets.
	SourceFieldIds          ScalarList[string] `gorm:"column:source_field_ids;type:text[]" json:"source_field_ids"`
	SeedRecordCount         int                `gorm:"column:seed_record_count;size:32;default:0;not null" json:"seed_record_count"`
	StageDepth              int                `gorm:"column:stage_depth;size:32;default:0;not null" json:"stage_depth"`
	PredecessorTaskID       *string            `gorm:"column:predecessor_task_id" json:"predecessor_task_id"`
	RunTotalSteps           int                `gorm:"column:run_total_steps;size:32;default:0;not null" json:"run_total_steps"`
	RunCompletedStepsBefore int                `gorm:"column:run_completed_steps_before;size:32;default:0;not null" json:"run_completed_steps_before"`
	SyncMaxLevel            *int               `gorm:"column:sync_max_level;size:32" json:"sync_max_level"`
	EstimatedComplexity     int                `gorm:"column:estimated_complexity;size:32;default:0;not null" json:"estimated_complexity"`
	Attempts                int                `gorm:"column:attempts;size:32;default:0;not null" json:"attempts"`
	Outcome                 string             `gorm:"column:outcome;not null" json:"outcome"`
	SourceChangedAt         *DateTime          `gorm:"column:source_changed_at;type:timestamp(3)" json:"source_changed_at"`
	EnqueuedAt              DateTime           `gorm:"column:enqueued_at;type:timestamp(3);not null" json:"enqueued_at"`
	StartedAt               *DateTime          `gorm:"column:started_at;type:timestamp(3)" json:"started_at"`
	CompletedAt             DateTime           `gorm:"column:completed_at;index:computed_update_run_history_base_id_completed_at_idx,priority:2;index:computed_update_run_history_completed_at_idx;type:timestamp(3);not null" json:"completed_at"`
	DurationMs              int                `gorm:"column:duration_ms;size:32;default:0;not null" json:"duration_ms"`
}

func (ComputedUpdateRunHistory) TableName() string {
	return "computed_update_run_history"
}

type ComputedUpdatePauseScope struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	ScopeType   string    `gorm:"column:scope_type;uniqueIndex:computed_update_pause_scope_scope_type_scope_id_key,priority:1;not null" json:"scope_type"`
	ScopeID     string    `gorm:"column:scope_id;uniqueIndex:computed_update_pause_scope_scope_type_scope_id_key,priority:2;not null" json:"scope_id"`
	PausedAt    DateTime  `gorm:"column:paused_at;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"paused_at"`
	PausedBy    *string   `gorm:"column:paused_by" json:"paused_by"`
	ResumeAt    *DateTime `gorm:"column:resume_at;index:computed_update_pause_scope_resume_at_idx;type:timestamp(3)" json:"resume_at"`
	Reason      *string   `gorm:"column:reason" json:"reason"`
	WritePolicy *string   `gorm:"column:write_policy;default:'allow_bounded';not null" json:"write_policy"`
	UpdatedAt   DateTime  `gorm:"column:updated_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
	UpdatedBy   *string   `gorm:"column:updated_by" json:"updated_by"`
}

func (ComputedUpdatePauseScope) TableName() string {
	return "computed_update_pause_scope"
}

func (m *ComputedUpdatePauseScope) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.PausedAt.IsZero() {
		m.PausedAt = DateTime{Time: now}
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

func (*ComputedUpdatePauseScope) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("UpdatedAt") {
		tx.Statement.SetColumn("UpdatedAt", DateTime{Time: now})
	}
	return nil
}

// Projection of async computed activity per field (not field schema meta).
type ComputedFieldActivity struct {
	FieldID               string         `gorm:"column:field_id;primaryKey" json:"field_id"`
	TableID               string         `gorm:"column:table_id;index:computed_field_activity_table_id_status_idx,priority:1;not null" json:"table_id"`
	BaseID                string         `gorm:"column:base_id;index:computed_field_activity_base_id_status_idx,priority:1;not null" json:"base_id"`
	Status                string         `gorm:"column:status;index:computed_field_activity_table_id_status_idx,priority:2;index:computed_field_activity_base_id_status_idx,priority:2;not null" json:"status"`
	ActiveTaskCount       int            `gorm:"column:active_task_count;size:32;default:0;not null" json:"active_task_count"`
	ProcessingTaskCount   int            `gorm:"column:processing_task_count;size:32;default:0;not null" json:"processing_task_count"`
	Generation            int64          `gorm:"column:generation;default:0;not null" json:"generation"`
	EstimatedComplexity   int64          `gorm:"column:estimated_complexity;default:0;not null" json:"estimated_complexity"`
	EstimatedDirtyRecords int64          `gorm:"column:estimated_dirty_records;default:0;not null" json:"estimated_dirty_records"`
	HasAllTargetRecords   bool           `gorm:"column:has_all_target_records;default:false;not null" json:"has_all_target_records"`
	QueuedAt              *DateTime      `gorm:"column:queued_at;type:timestamp(3)" json:"queued_at"`
	StartedAt             *DateTime      `gorm:"column:started_at;type:timestamp(3)" json:"started_at"`
	LastCompletedAt       *DateTime      `gorm:"column:last_completed_at;type:timestamp(3)" json:"last_completed_at"`
	LastDurationMs        *int           `gorm:"column:last_duration_ms;size:32" json:"last_duration_ms"`
	LastError             datatypes.JSON `gorm:"column:last_error" json:"last_error"`
	Extensions            datatypes.JSON `gorm:"column:extensions" json:"extensions"`
	UpdatedAt             DateTime       `gorm:"column:updated_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
}

func (ComputedFieldActivity) TableName() string {
	return "computed_field_activity"
}

func (m *ComputedFieldActivity) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

// Projection of async computed activity per table.
type ComputedTableActivity struct {
	TableID               string         `gorm:"column:table_id;primaryKey" json:"table_id"`
	BaseID                string         `gorm:"column:base_id;index:computed_table_activity_base_id_status_idx,priority:1;not null" json:"base_id"`
	Status                string         `gorm:"column:status;index:computed_table_activity_base_id_status_idx,priority:2;not null" json:"status"`
	CalculatingFieldCount int            `gorm:"column:calculating_field_count;size:32;default:0;not null" json:"calculating_field_count"`
	QueuedFieldCount      int            `gorm:"column:queued_field_count;size:32;default:0;not null" json:"queued_field_count"`
	EstimatedComplexity   int64          `gorm:"column:estimated_complexity;default:0;not null" json:"estimated_complexity"`
	RecentCompletions     datatypes.JSON `gorm:"column:recent_completions;default:'[]';not null" json:"recent_completions"`
	Generation            int64          `gorm:"column:generation;default:0;not null" json:"generation"`
	UpdatedAt             DateTime       `gorm:"column:updated_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoUpdateTime:false;not null" json:"updated_at"`
}

func (ComputedTableActivity) TableName() string {
	return "computed_table_activity"
}

func (m *ComputedTableActivity) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = DateTime{Time: now}
	}
	return nil
}

// Task→field refset for activity refcounting (prevents merge double-count).
type ComputedTaskFieldRef struct {
	TaskID        string   `gorm:"column:task_id;primaryKey" json:"task_id"`
	FieldID       string   `gorm:"column:field_id;primaryKey;index:computed_task_field_ref_field_id_idx" json:"field_id"`
	TableID       string   `gorm:"column:table_id;index:computed_task_field_ref_table_id_idx;not null" json:"table_id"`
	BaseID        string   `gorm:"column:base_id;not null" json:"base_id"`
	WasProcessing bool     `gorm:"column:was_processing;default:false;not null" json:"was_processing"`
	CreatedAt     DateTime `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP;autoCreateTime:false;not null" json:"created_at"`
}

func (ComputedTaskFieldRef) TableName() string {
	return "computed_task_field_ref"
}

func (m *ComputedTaskFieldRef) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedAt.IsZero() {
		m.CreatedAt = DateTime{Time: now}
	}
	return nil
}

type View struct {
	ID               string     `gorm:"column:id;primaryKey" json:"id"`
	Name             string     `gorm:"column:name;not null" json:"name"`
	Description      *string    `gorm:"column:description" json:"description"`
	TableID          string     `gorm:"column:table_id;index:view_table_id_deleted_time_idx,priority:1;not null" json:"table_id"`
	Type             string     `gorm:"column:type;not null" json:"type"`
	Sort             *string    `gorm:"column:sort" json:"sort"`
	Filter           *string    `gorm:"column:filter" json:"filter"`
	Group            *string    `gorm:"column:group" json:"group"`
	Options          *string    `gorm:"column:options" json:"options"`
	Order            float64    `gorm:"column:order;index:view_order_idx;type:double precision;not null" json:"order"`
	Version          int        `gorm:"column:version;size:32;not null" json:"version"`
	ColumnMeta       string     `gorm:"column:column_meta;not null" json:"column_meta"`
	IsLocked         *bool      `gorm:"column:is_locked" json:"is_locked"`
	EnableShare      *bool      `gorm:"column:enable_share" json:"enable_share"`
	ShareID          *string    `gorm:"column:share_id;uniqueIndex:view_share_id_key" json:"share_id"`
	ShareMeta        *string    `gorm:"column:share_meta" json:"share_meta"`
	CreatedTime      DateTime   `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime  `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	DeletedTime      *DateTime  `gorm:"column:deleted_time;index:view_table_id_deleted_time_idx,priority:2;type:timestamp(3)" json:"deleted_time"`
	CreatedBy        string     `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy   *string    `gorm:"column:last_modified_by" json:"last_modified_by"`
	Table            *TableMeta `gorm:"foreignKey:TableID;belongsTo"`
}

func (View) TableName() string {
	return "view"
}

func (m *View) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*View) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Ops struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	Collection  string   `gorm:"column:collection;uniqueIndex:ops_collection_doc_id_version_key,priority:1;index:ops_collection_created_time_idx,priority:1;not null" json:"collection"`
	DocID       string   `gorm:"column:doc_id;uniqueIndex:ops_collection_doc_id_version_key,priority:2;not null" json:"doc_id"`
	DocType     string   `gorm:"column:doc_type;not null" json:"doc_type"`
	Version     int      `gorm:"column:version;uniqueIndex:ops_collection_doc_id_version_key,priority:3;size:32;not null" json:"version"`
	Operation   string   `gorm:"column:operation;not null" json:"operation"`
	CreatedTime DateTime `gorm:"column:created_time;index:ops_collection_created_time_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy   string   `gorm:"column:created_by;not null" json:"created_by"`
}

func (Ops) TableName() string {
	return "ops"
}

func (m *Ops) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type Reference struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	FromFieldID string   `gorm:"column:from_field_id;uniqueIndex:reference_to_field_id_from_field_id_key,priority:2;index:reference_from_field_id_idx;not null" json:"from_field_id"`
	ToFieldID   string   `gorm:"column:to_field_id;uniqueIndex:reference_to_field_id_from_field_id_key,priority:1;index:reference_to_field_id_idx;not null" json:"to_field_id"`
	CreatedTime DateTime `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
}

func (Reference) TableName() string {
	return "reference"
}

func (m *Reference) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type User struct {
	ID                   string          `gorm:"column:id;primaryKey" json:"id"`
	Name                 string          `gorm:"column:name;not null" json:"name"`
	Password             *string         `gorm:"column:password" json:"password"`
	Salt                 *string         `gorm:"column:salt" json:"salt"`
	Phone                *string         `gorm:"column:phone;uniqueIndex:users_phone_key" json:"phone"`
	Email                string          `gorm:"column:email;uniqueIndex:users_email_key;not null" json:"email"`
	Avatar               *string         `gorm:"column:avatar" json:"avatar"`
	IsSystem             *bool           `gorm:"column:is_system" json:"is_system"`
	IsAdmin              *bool           `gorm:"column:is_admin" json:"is_admin"`
	IsTrialUsed          *bool           `gorm:"column:is_trial_used" json:"is_trial_used"`
	Lang                 *string         `gorm:"column:lang" json:"lang"`
	NotifyMeta           *string         `gorm:"column:notify_meta" json:"notify_meta"`
	LastSignTime         *DateTime       `gorm:"column:last_sign_time;type:timestamp(3)" json:"last_sign_time"`
	DeactivatedTime      *DateTime       `gorm:"column:deactivated_time;type:timestamp(3)" json:"deactivated_time"`
	CreatedTime          DateTime        `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	DeletedTime          *DateTime       `gorm:"column:deleted_time;type:timestamp(3)" json:"deleted_time"`
	LastModifiedTime     *DateTime       `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	PermanentDeletedTime *DateTime       `gorm:"column:permanent_deleted_time;type:timestamp(3)" json:"permanent_deleted_time"`
	RefMeta              *string         `gorm:"column:ref_meta" json:"ref_meta"`
	Accounts             []Account       `gorm:"foreignKey:UserID;constraint:account_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	BaseOrders           []UserBaseOrder `gorm:"foreignKey:UserID;constraint:user_base_order_user_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "users"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*User) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Account struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	UserID       string   `gorm:"column:user_id;not null" json:"user_id"`
	Type         string   `gorm:"column:type;not null" json:"type"`
	Provider     string   `gorm:"column:provider;uniqueIndex:account_provider_provider_id_connection_id_key,priority:1;not null" json:"provider"`
	ProviderID   string   `gorm:"column:provider_id;uniqueIndex:account_provider_provider_id_connection_id_key,priority:2;not null" json:"provider_id"`
	ConnectionID string   `gorm:"column:connection_id;uniqueIndex:account_provider_provider_id_connection_id_key,priority:3;default:'';not null" json:"connection_id"`
	CreatedTime  DateTime `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	User         *User    `gorm:"belongsTo"`
}

func (Account) TableName() string {
	return "account"
}

func (m *Account) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type Attachments struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`
	Token          string    `gorm:"column:token;uniqueIndex:attachments_token_key;not null" json:"token"`
	Hash           string    `gorm:"column:hash;not null" json:"hash"`
	Size           int64     `gorm:"column:size;not null" json:"size"`
	Mimetype       string    `gorm:"column:mimetype;not null" json:"mimetype"`
	Path           string    `gorm:"column:path;not null" json:"path"`
	Width          *int      `gorm:"column:width;size:32" json:"width"`
	Height         *int      `gorm:"column:height;size:32" json:"height"`
	DeletedTime    *DateTime `gorm:"column:deleted_time;type:timestamp(3)" json:"deleted_time"`
	CreatedTime    DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy      string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
	ThumbnailPath  *string   `gorm:"column:thumbnail_path" json:"thumbnail_path"`
}

func (Attachments) TableName() string {
	return "attachments"
}

func (m *Attachments) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type AttachmentsTable struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	AttachmentID     string    `gorm:"column:attachment_id;index:attachments_table_attachment_id_idx;not null" json:"attachment_id"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Token            string    `gorm:"column:token;not null" json:"token"`
	TableID          string    `gorm:"column:table_id;index:attachments_table_table_id_record_id_idx,priority:1;index:attachments_table_table_id_field_id_idx,priority:1;not null" json:"table_id"`
	RecordID         string    `gorm:"column:record_id;index:attachments_table_table_id_record_id_idx,priority:2;not null" json:"record_id"`
	FieldID          string    `gorm:"column:field_id;index:attachments_table_table_id_field_id_idx,priority:2;not null" json:"field_id"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
}

func (AttachmentsTable) TableName() string {
	return "attachments_table"
}

func (m *AttachmentsTable) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*AttachmentsTable) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Collaborator struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	RoleName         string    `gorm:"column:role_name;not null" json:"role_name"`
	ResourceType     string    `gorm:"column:resource_type;uniqueIndex:collaborator_resource_type_resource_id_principal_id_princip_key,priority:1;not null" json:"resource_type"`
	ResourceID       string    `gorm:"column:resource_id;uniqueIndex:collaborator_resource_type_resource_id_principal_id_princip_key,priority:2;index:collaborator_resource_id_idx;not null" json:"resource_id"`
	PrincipalID      string    `gorm:"column:principal_id;uniqueIndex:collaborator_resource_type_resource_id_principal_id_princip_key,priority:3;index:collaborator_principal_id_idx;not null" json:"principal_id"`
	PrincipalType    string    `gorm:"column:principal_type;uniqueIndex:collaborator_resource_type_resource_id_principal_id_princip_key,priority:4;not null" json:"principal_type"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
}

func (Collaborator) TableName() string {
	return "collaborator"
}

func (m *Collaborator) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Collaborator) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Invitation struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	BaseID           *string   `gorm:"column:base_id;index:invitation_base_id_idx" json:"base_id"`
	SpaceID          *string   `gorm:"column:space_id;index:invitation_space_id_idx" json:"space_id"`
	Type             string    `gorm:"column:type;not null" json:"type"`
	Role             string    `gorm:"column:role;not null" json:"role"`
	InvitationCode   string    `gorm:"column:invitation_code;not null" json:"invitation_code"`
	ExpiredTime      *DateTime `gorm:"column:expired_time;type:timestamp(3)" json:"expired_time"`
	CreatedBy        string    `gorm:"column:create_by;not null" json:"create_by"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
	DeletedTime      *DateTime `gorm:"column:deleted_time;type:timestamp(3)" json:"deleted_time"`
}

func (Invitation) TableName() string {
	return "invitation"
}

func (m *Invitation) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Invitation) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type InvitationRecord struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	InvitationID string   `gorm:"column:invitation_id;index:invitation_record_invitation_id_idx;not null" json:"invitation_id"`
	BaseID       *string  `gorm:"column:base_id;index:invitation_record_base_id_idx" json:"base_id"`
	SpaceID      *string  `gorm:"column:space_id;index:invitation_record_space_id_idx" json:"space_id"`
	Type         string   `gorm:"column:type;not null" json:"type"`
	Inviter      string   `gorm:"column:inviter;not null" json:"inviter"`
	Accepter     string   `gorm:"column:accepter;not null" json:"accepter"`
	CreatedTime  DateTime `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
}

func (InvitationRecord) TableName() string {
	return "invitation_record"
}

func (m *InvitationRecord) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type Notification struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	FromUserID  string   `gorm:"column:from_user_id;not null" json:"from_user_id"`
	ToUserID    string   `gorm:"column:to_user_id;index:notification_to_user_id_is_read_created_time_idx,priority:1;not null" json:"to_user_id"`
	Type        string   `gorm:"column:type;not null" json:"type"`
	Message     string   `gorm:"column:message;not null" json:"message"`
	MessageI18n *string  `gorm:"column:message_i18n" json:"message_i18n"`
	Severity    *string  `gorm:"column:severity;default:'info';not null" json:"severity"`
	URLPath     *string  `gorm:"column:url_path" json:"url_path"`
	IsRead      bool     `gorm:"column:is_read;index:notification_to_user_id_is_read_created_time_idx,priority:2;default:false;not null" json:"is_read"`
	CreatedTime DateTime `gorm:"column:created_time;index:notification_to_user_id_is_read_created_time_idx,priority:3;index:notification_created_time_idx;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy   string   `gorm:"column:created_by;not null" json:"created_by"`
}

func (Notification) TableName() string {
	return "notification"
}

func (m *Notification) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type AccessToken struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Description      *string   `gorm:"column:description" json:"description"`
	UserID           string    `gorm:"column:user_id;index:access_token_user_id_idx;not null" json:"user_id"`
	Scopes           string    `gorm:"column:scopes;not null" json:"scopes"`
	SpaceIds         *string   `gorm:"column:space_ids" json:"space_ids"`
	BaseIds          *string   `gorm:"column:base_ids" json:"base_ids"`
	Sign             string    `gorm:"column:sign;not null" json:"sign"`
	ClientID         *string   `gorm:"column:client_id;index:access_token_client_id_idx" json:"client_id"`
	HasFullAccess    *bool     `gorm:"column:has_full_access" json:"has_full_access"`
	ExpiredTime      DateTime  `gorm:"column:expired_time;type:timestamp(3);not null" json:"expired_time"`
	LastUsedTime     *DateTime `gorm:"column:last_used_time;type:timestamp(3)" json:"last_used_time"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
}

func (AccessToken) TableName() string {
	return "access_token"
}

func (m *AccessToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*AccessToken) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Setting struct {
	Name             string    `gorm:"column:name;primaryKey;uniqueIndex:setting_name_key" json:"name"`
	Content          *string   `gorm:"column:content" json:"content"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
}

func (Setting) TableName() string {
	return "setting"
}

func (m *Setting) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Setting) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type OAuthApp struct {
	ID                  string               `gorm:"column:id;primaryKey" json:"id"`
	Name                string               `gorm:"column:name;not null" json:"name"`
	Logo                *string              `gorm:"column:logo" json:"logo"`
	Homepage            string               `gorm:"column:homepage;not null" json:"homepage"`
	Description         *string              `gorm:"column:description" json:"description"`
	ClientID            string               `gorm:"column:client_id;uniqueIndex:oauth_app_client_id_key;not null" json:"client_id"`
	RedirectUris        *string              `gorm:"column:redirect_uris" json:"redirect_uris"`
	Scopes              *string              `gorm:"column:scopes" json:"scopes"`
	AllowDeviceFlow     bool                 `gorm:"column:allow_device_flow;default:false;not null" json:"allow_device_flow"`
	DynamicRegistration bool                 `gorm:"column:dynamic_registration;default:false;not null" json:"dynamic_registration"`
	CreatedTime         DateTime             `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime    *DateTime            `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	CreatedBy           string               `gorm:"column:created_by;not null" json:"created_by"`
	Secrets             []OAuthAppSecret     `gorm:"foreignKey:ClientID;references:ClientID;constraint:oauth_app_secret_client_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Authorized          []OAuthAppAuthorized `gorm:"foreignKey:ClientID;references:ClientID;constraint:oauth_app_authorized_client_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tokens              []OAuthAppToken      `gorm:"foreignKey:ClientID;references:ClientID;constraint:oauth_app_token_client_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OAuthApp) TableName() string {
	return "oauth_app"
}

func (m *OAuthApp) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*OAuthApp) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type OAuthAppAuthorized struct {
	ID             string    `gorm:"column:id;primaryKey" json:"id"`
	ClientID       string    `gorm:"column:client_id;uniqueIndex:oauth_app_authorized_client_id_user_id_key,priority:1;not null" json:"client_id"`
	UserID         string    `gorm:"column:user_id;uniqueIndex:oauth_app_authorized_client_id_user_id_key,priority:2;not null" json:"user_id"`
	AuthorizedTime DateTime  `gorm:"column:authorized_time;type:timestamp(3);not null" json:"authorized_time"`
	OAuthApp       *OAuthApp `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
}

func (OAuthAppAuthorized) TableName() string {
	return "oauth_app_authorized"
}

func (m *OAuthAppAuthorized) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = cuid.New()
	}
	return nil
}

type OAuthAppSecret struct {
	ID           string          `gorm:"column:id;primaryKey" json:"id"`
	ClientID     string          `gorm:"column:client_id;not null" json:"client_id"`
	Secret       string          `gorm:"column:secret;uniqueIndex:oauth_app_secret_secret_key;not null" json:"secret"`
	MaskedSecret string          `gorm:"column:masked_secret;not null" json:"masked_secret"`
	CreatedTime  DateTime        `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy    string          `gorm:"column:created_by;not null" json:"created_by"`
	LastUsedTime *DateTime       `gorm:"column:last_used_time;type:timestamp(3)" json:"last_used_time"`
	OAuthApp     *OAuthApp       `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
	Tokens       []OAuthAppToken `gorm:"foreignKey:AppSecretID;constraint:oauth_app_token_app_secret_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OAuthAppSecret) TableName() string {
	return "oauth_app_secret"
}

func (m *OAuthAppSecret) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type OAuthAppToken struct {
	ID               string          `gorm:"column:id;primaryKey" json:"id"`
	ClientID         string          `gorm:"column:client_id;index:oauth_app_token_client_id_idx;not null" json:"client_id"`
	AppSecretID      *string         `gorm:"column:app_secret_id" json:"app_secret_id"`
	RefreshTokenSign string          `gorm:"column:refresh_token_sign;uniqueIndex:oauth_app_token_refresh_token_sign_key;not null" json:"refresh_token_sign"`
	ExpiredTime      DateTime        `gorm:"column:expired_time;type:timestamp(3);not null" json:"expired_time"`
	CreatedTime      DateTime        `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string          `gorm:"column:created_by;not null" json:"created_by"`
	OAuthAppSecret   *OAuthAppSecret `gorm:"foreignKey:AppSecretID;belongsTo"`
	OAuthApp         *OAuthApp       `gorm:"foreignKey:ClientID;references:ClientID;belongsTo"`
}

func (OAuthAppToken) TableName() string {
	return "oauth_app_token"
}

func (m *OAuthAppToken) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type RecordHistory struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	TableID     string   `gorm:"column:table_id;index:record_history_table_id_record_id_created_time_idx,priority:1;index:record_history_table_id_created_time_idx,priority:1;not null" json:"table_id"`
	RecordID    string   `gorm:"column:record_id;index:record_history_table_id_record_id_created_time_idx,priority:2;not null" json:"record_id"`
	FieldID     string   `gorm:"column:field_id;not null" json:"field_id"`
	Before      string   `gorm:"column:before;not null" json:"before"`
	After       string   `gorm:"column:after;not null" json:"after"`
	CreatedTime DateTime `gorm:"column:created_time;index:record_history_table_id_record_id_created_time_idx,priority:3;index:record_history_table_id_created_time_idx,priority:2;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy   string   `gorm:"column:created_by;not null" json:"created_by"`
	ActorType   *string  `gorm:"column:actor_type" json:"actor_type"`
	ActorID     *string  `gorm:"column:actor_id" json:"actor_id"`
}

func (RecordHistory) TableName() string {
	return "record_history"
}

func (m *RecordHistory) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type Trash struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	ResourceType string   `gorm:"column:resource_type;uniqueIndex:trash_resource_type_resource_id_key,priority:1;not null" json:"resource_type"`
	ResourceID   string   `gorm:"column:resource_id;uniqueIndex:trash_resource_type_resource_id_key,priority:2;not null" json:"resource_id"`
	ParentID     *string  `gorm:"column:parent_id" json:"parent_id"`
	DeletedTime  DateTime `gorm:"column:deleted_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"deleted_time"`
	DeletedBy    string   `gorm:"column:deleted_by;not null" json:"deleted_by"`
}

func (Trash) TableName() string {
	return "trash"
}

func (m *Trash) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.DeletedTime.IsZero() {
		m.DeletedTime = DateTime{Time: now}
	}
	return nil
}

type TableTrash struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	TableID      string   `gorm:"column:table_id;index:table_trash_table_id_idx;not null" json:"table_id"`
	ResourceType string   `gorm:"column:resource_type;not null" json:"resource_type"`
	Snapshot     string   `gorm:"column:snapshot;not null" json:"snapshot"`
	CreatedTime  DateTime `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy    string   `gorm:"column:created_by;not null" json:"created_by"`
}

func (TableTrash) TableName() string {
	return "table_trash"
}

func (m *TableTrash) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type RecordTrash struct {
	ID                     string    `gorm:"column:id;primaryKey" json:"id"`
	TableID                string    `gorm:"column:table_id;index:record_trash_table_id_record_id_idx,priority:1;not null" json:"table_id"`
	RecordID               string    `gorm:"column:record_id;index:record_trash_table_id_record_id_idx,priority:2;not null" json:"record_id"`
	Snapshot               string    `gorm:"column:snapshot;not null" json:"snapshot"`
	CreatedTime            DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy              string    `gorm:"column:created_by;not null" json:"created_by"`
	Reason                 *string   `gorm:"column:reason;default:'deleted';not null" json:"reason"`
	RecordCreatedTime      *DateTime `gorm:"column:record_created_time;type:timestamp(3)" json:"record_created_time"`
	RecordCreatedBy        *string   `gorm:"column:record_created_by" json:"record_created_by"`
	RecordLastModifiedTime *DateTime `gorm:"column:record_last_modified_time;type:timestamp(3)" json:"record_last_modified_time"`
	RecordLastModifiedBy   *string   `gorm:"column:record_last_modified_by" json:"record_last_modified_by"`
	OperationID            *string   `gorm:"column:operation_id" json:"operation_id"`
}

func (RecordTrash) TableName() string {
	return "record_trash"
}

func (m *RecordTrash) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type RecordRemovalTombstone struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	TableID     string   `gorm:"column:table_id;index:record_removal_tombstone_table_id_record_id_idx,priority:1;not null" json:"table_id"`
	RecordID    string   `gorm:"column:record_id;index:record_removal_tombstone_table_id_record_id_idx,priority:2;not null" json:"record_id"`
	Type        string   `gorm:"column:type;not null" json:"type"`
	CreatedTime DateTime `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
}

func (RecordRemovalTombstone) TableName() string {
	return "record_removal_tombstone"
}

func (m *RecordRemovalTombstone) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type Plugin struct {
	ID               string          `gorm:"column:id;primaryKey" json:"id"`
	Name             string          `gorm:"column:name;not null" json:"name"`
	Description      *string         `gorm:"column:description" json:"description"`
	DetailDesc       *string         `gorm:"column:detail_desc" json:"detail_desc"`
	Logo             string          `gorm:"column:logo;not null" json:"logo"`
	HelpURL          *string         `gorm:"column:help_url" json:"help_url"`
	Status           string          `gorm:"column:status;not null" json:"status"`
	Positions        string          `gorm:"column:positions;not null" json:"positions"`
	URL              *string         `gorm:"column:url" json:"url"`
	Secret           string          `gorm:"column:secret;uniqueIndex:plugin_secret_key;not null" json:"secret"`
	MaskedSecret     string          `gorm:"column:masked_secret;not null" json:"masked_secret"`
	I18n             *string         `gorm:"column:i18n" json:"i18n"`
	Config           *string         `gorm:"column:config" json:"config"`
	PluginUser       *string         `gorm:"column:plugin_user" json:"plugin_user"`
	CreatedTime      DateTime        `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime       `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	CreatedBy        string          `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy   *string         `gorm:"column:last_modified_by" json:"last_modified_by"`
	PluginInstall    []PluginInstall `gorm:"foreignKey:PluginID;constraint:plugin_install_plugin_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Plugin) TableName() string {
	return "plugin"
}

func (m *Plugin) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Plugin) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type PluginInstall struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	PluginID         string    `gorm:"column:plugin_id;not null" json:"plugin_id"`
	BaseID           string    `gorm:"column:base_id;index:plugin_install_base_id_idx;not null" json:"base_id"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	PositionID       string    `gorm:"column:position_id;index:plugin_install_position_id_idx;not null" json:"position_id"`
	Position         string    `gorm:"column:position;not null" json:"position"`
	Storage          *string   `gorm:"column:storage" json:"storage"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
	Plugin           *Plugin   `gorm:"belongsTo"`
}

func (PluginInstall) TableName() string {
	return "plugin_install"
}

func (m *PluginInstall) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*PluginInstall) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Dashboard struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	BaseID           string    `gorm:"column:base_id;index:dashboard_base_id_idx;not null" json:"base_id"`
	Layout           *string   `gorm:"column:layout" json:"layout"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
	DeletedTime      *DateTime `gorm:"column:deleted_time;type:timestamp(3)" json:"deleted_time"`
}

func (Dashboard) TableName() string {
	return "dashboard"
}

func (m *Dashboard) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Dashboard) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Comment struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	TableID          string    `gorm:"column:table_id;index:comment_table_id_record_id_idx,priority:1;not null" json:"table_id"`
	RecordID         string    `gorm:"column:record_id;index:comment_table_id_record_id_idx,priority:2;not null" json:"record_id"`
	QuoteID          *string   `gorm:"column:quote_Id" json:"quote_Id"`
	Content          *string   `gorm:"column:content" json:"content"`
	Reaction         *string   `gorm:"column:reaction" json:"reaction"`
	DeletedTime      *DateTime `gorm:"column:deleted_time;type:timestamp(3)" json:"deleted_time"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
}

func (Comment) TableName() string {
	return "comment"
}

func (m *Comment) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Comment) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type CommentSubscription struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	TableID     string   `gorm:"column:table_id;uniqueIndex:comment_subscription_table_id_record_id_key,priority:1;index:comment_subscription_table_id_record_id_idx,priority:1;not null" json:"table_id"`
	RecordID    string   `gorm:"column:record_id;uniqueIndex:comment_subscription_table_id_record_id_key,priority:2;index:comment_subscription_table_id_record_id_idx,priority:2;not null" json:"record_id"`
	CreatedBy   string   `gorm:"column:created_by;not null" json:"created_by"`
	CreatedTime DateTime `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
}

func (CommentSubscription) TableName() string {
	return "comment_subscription"
}

func (m *CommentSubscription) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type Integration struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	ResourceID       string    `gorm:"column:resource_id;uniqueIndex:integration_resource_id_key;index:integration_resource_id_idx;not null" json:"resource_id"`
	Config           string    `gorm:"column:config;not null" json:"config"`
	Type             string    `gorm:"column:type;not null" json:"type"`
	Enable           *bool     `gorm:"column:enable" json:"enable"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
}

func (Integration) TableName() string {
	return "integration"
}

func (m *Integration) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Integration) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type PluginPanel struct {
	ID               string     `gorm:"column:id;primaryKey" json:"id"`
	Name             string     `gorm:"column:name;not null" json:"name"`
	TableID          string     `gorm:"column:table_id;not null" json:"table_id"`
	Layout           *string    `gorm:"column:layout" json:"layout"`
	CreatedBy        string     `gorm:"column:created_by;not null" json:"created_by"`
	CreatedTime      DateTime   `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime  `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string    `gorm:"column:last_modified_by" json:"last_modified_by"`
	Table            *TableMeta `gorm:"foreignKey:TableID;belongsTo"`
}

func (PluginPanel) TableName() string {
	return "plugin_panel"
}

func (m *PluginPanel) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*PluginPanel) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type PluginContextMenu struct {
	ID               string     `gorm:"column:id;primaryKey" json:"id"`
	TableID          string     `gorm:"column:table_id;not null" json:"table_id"`
	PluginInstallID  string     `gorm:"column:plugin_install_id;uniqueIndex:plugin_context_menu_plugin_install_id_key;not null" json:"plugin_install_id"`
	Order            float64    `gorm:"column:order;type:double precision;not null" json:"order"`
	CreatedTime      DateTime   `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string     `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime *DateTime  `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string    `gorm:"column:last_modified_by" json:"last_modified_by"`
	Table            *TableMeta `gorm:"foreignKey:TableID;belongsTo"`
}

func (PluginContextMenu) TableName() string {
	return "plugin_context_menu"
}

func (m *PluginContextMenu) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*PluginContextMenu) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type UserLastVisit struct {
	ID               string   `gorm:"column:id;primaryKey" json:"id"`
	UserID           string   `gorm:"column:user_id;uniqueIndex:user_last_visit_user_id_resource_type_resource_id_key,priority:1;index:user_last_visit_user_id_resource_type_parent_resource_id_idx,priority:1;not null" json:"user_id"`
	ResourceType     string   `gorm:"column:resource_type;uniqueIndex:user_last_visit_user_id_resource_type_resource_id_key,priority:2;index:user_last_visit_user_id_resource_type_parent_resource_id_idx,priority:2;not null" json:"resource_type"`
	ResourceID       string   `gorm:"column:resource_id;uniqueIndex:user_last_visit_user_id_resource_type_resource_id_key,priority:3;not null" json:"resource_id"`
	ParentResourceID string   `gorm:"column:parent_resource_id;index:user_last_visit_user_id_resource_type_parent_resource_id_idx,priority:3;not null" json:"parent_resource_id"`
	LastVisitTime    DateTime `gorm:"column:last_visit_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"last_visit_time"`
}

func (UserLastVisit) TableName() string {
	return "user_last_visit"
}

func (m *UserLastVisit) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.LastVisitTime.IsZero() {
		m.LastVisitTime = DateTime{Time: now}
	}
	return nil
}

type UserBaseOrder struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	UserID           string    `gorm:"column:user_id;uniqueIndex:user_base_order_user_id_base_id_key,priority:1;index:user_base_order_user_id_order_idx,priority:1;not null" json:"user_id"`
	BaseID           string    `gorm:"column:base_id;uniqueIndex:user_base_order_user_id_base_id_key,priority:2;index:user_base_order_base_id_idx;not null" json:"base_id"`
	Order            float64   `gorm:"column:order;index:user_base_order_user_id_order_idx,priority:2;type:double precision;not null" json:"order"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	User             *User     `gorm:"belongsTo"`
	Base             *Base     `gorm:"belongsTo"`
}

func (UserBaseOrder) TableName() string {
	return "user_base_order"
}

func (m *UserBaseOrder) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*UserBaseOrder) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Template struct {
	ID                  string             `gorm:"column:id;primaryKey" json:"id"`
	BaseID              *string            `gorm:"column:base_id;uniqueIndex:template_base_id_key" json:"base_id"`
	Cover               *string            `gorm:"column:cover" json:"cover"`
	Name                *string            `gorm:"column:name" json:"name"`
	Description         *string            `gorm:"column:description" json:"description"`
	MarkdownDescription *string            `gorm:"column:markdown_description" json:"markdown_description"`
	CategoryID          ScalarList[string] `gorm:"column:category_id;type:text[]" json:"category_id"`
	CreatedTime         DateTime           `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy           string             `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime    *DateTime          `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy      *string            `gorm:"column:last_modified_by" json:"last_modified_by"`
	IsSystem            *bool              `gorm:"column:is_system" json:"is_system"`
	IsPublished         *bool              `gorm:"column:is_published" json:"is_published"`
	Featured            *bool              `gorm:"column:featured" json:"featured"`
	Snapshot            *string            `gorm:"column:snapshot" json:"snapshot"`
	Order               float64            `gorm:"column:order;type:double precision;not null" json:"order"`
	UsageCount          int                `gorm:"column:usage_count;size:32;default:0;not null" json:"usage_count"`
	PublishInfo         datatypes.JSON     `gorm:"column:publish_info" json:"publish_info"`
	VisitCount          int                `gorm:"column:visit_count;size:32;default:0;not null" json:"visit_count"`
	Kind                *string            `gorm:"column:kind" json:"kind"`
	Version             *int               `gorm:"column:version;size:32;default:1;not null" json:"version"`
	Manifest            datatypes.JSON     `gorm:"column:manifest" json:"manifest"`
	Review              datatypes.JSON     `gorm:"column:review" json:"review"`
	History             datatypes.JSON     `gorm:"column:history" json:"history"`
	ParentTemplateID    *string            `gorm:"column:parent_template_id" json:"parent_template_id"`
	Slug                *string            `gorm:"column:slug;uniqueIndex:template_slug_key" json:"slug"`
}

func (Template) TableName() string {
	return "template"
}

func (m *Template) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Template) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type TemplateCategory struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	Name             string    `gorm:"column:name;uniqueIndex:template_category_name_key;not null" json:"name"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
	Order            float64   `gorm:"column:order;type:double precision;not null" json:"order"`
	ParentID         *string   `gorm:"column:parent_id;index:template_category_parent_id_idx" json:"parent_id"`
	Kind             *string   `gorm:"column:kind;default:'industry';not null" json:"kind"`
	Key              *string   `gorm:"column:key;uniqueIndex:template_category_key_key" json:"key"`
}

func (TemplateCategory) TableName() string {
	return "template_category"
}

func (m *TemplateCategory) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*TemplateCategory) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type Task struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	Type             string    `gorm:"column:type;index:task_type_status_idx,priority:1;not null" json:"type"`
	Status           string    `gorm:"column:status;index:task_type_status_idx,priority:2;not null" json:"status"`
	Snapshot         *string   `gorm:"column:snapshot" json:"snapshot"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
	Runs             []TaskRun `gorm:"foreignKey:TaskID;constraint:task_run_task_id_fkey,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Task) TableName() string {
	return "task"
}

func (m *Task) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*Task) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type TaskRun struct {
	ID               string             `gorm:"column:id;primaryKey" json:"id"`
	TaskID           string             `gorm:"column:task_id;index:task_run_task_id_status_idx,priority:1;not null" json:"task_id"`
	BaseID           string             `gorm:"column:base_id;index:task_run_status_base_id_created_time_idx,priority:2;not null" json:"base_id"`
	Status           string             `gorm:"column:status;index:task_run_task_id_status_idx,priority:2;index:task_run_status_base_id_created_time_idx,priority:1;index:task_run_status_last_modified_time_idx,priority:1;not null" json:"status"`
	Snapshot         string             `gorm:"column:snapshot;not null" json:"snapshot"`
	DependsOnRunIds  ScalarList[string] `gorm:"column:depends_on_run_ids;type:text[]" json:"depends_on_run_ids"`
	Spent            *int               `gorm:"column:spent;size:32" json:"spent"`
	Log              *string            `gorm:"column:log" json:"log"`
	ErrorMsg         *string            `gorm:"column:error_msg" json:"error_msg"`
	StartedTime      *DateTime          `gorm:"column:started_time;type:timestamp(3)" json:"started_time"`
	CreatedTime      DateTime           `gorm:"column:created_time;index:task_run_status_base_id_created_time_idx,priority:3;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	LastModifiedTime *DateTime          `gorm:"column:last_modified_time;index:task_run_status_last_modified_time_idx,priority:2;type:timestamp(3)" json:"last_modified_time"`
	Task             *Task              `gorm:"belongsTo"`
}

func (TaskRun) TableName() string {
	return "task_run"
}

func (m *TaskRun) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*TaskRun) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type TaskReference struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	FromFieldID string   `gorm:"column:from_field_id;uniqueIndex:task_reference_to_field_id_from_field_id_key,priority:2;index:task_reference_from_field_id_idx;not null" json:"from_field_id"`
	ToFieldID   string   `gorm:"column:to_field_id;uniqueIndex:task_reference_to_field_id_from_field_id_key,priority:1;index:task_reference_to_field_id_idx;not null" json:"to_field_id"`
	CreatedTime DateTime `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
}

func (TaskReference) TableName() string {
	return "task_reference"
}

func (m *TaskReference) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type Waitlist struct {
	Email       string    `gorm:"column:email;primaryKey;uniqueIndex:waitlist_email_key" json:"email"`
	Invite      *bool     `gorm:"column:invite" json:"invite"`
	InviteTime  *DateTime `gorm:"column:invite_time;type:timestamp(3)" json:"invite_time"`
	CreatedTime DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
}

func (Waitlist) TableName() string {
	return "waitlist"
}

func (m *Waitlist) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	return nil
}

type BaseNode struct {
	ID               string     `gorm:"column:id;primaryKey" json:"id"`
	ParentID         *string    `gorm:"column:parent_id" json:"parent_id"`
	BaseID           string     `gorm:"column:base_id;uniqueIndex:base_node_base_id_resource_type_resource_id_key,priority:1;not null" json:"base_id"`
	ResourceType     string     `gorm:"column:resource_type;uniqueIndex:base_node_base_id_resource_type_resource_id_key,priority:2;not null" json:"resource_type"`
	ResourceID       string     `gorm:"column:resource_id;uniqueIndex:base_node_base_id_resource_type_resource_id_key,priority:3;not null" json:"resource_id"`
	Order            float64    `gorm:"column:order;type:double precision;not null" json:"order"`
	CreatedTime      DateTime   `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string     `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime *DateTime  `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string    `gorm:"column:last_modified_by" json:"last_modified_by"`
	Parent           *BaseNode  `gorm:"foreignKey:ParentID;belongsTo"`
	Children         []BaseNode `gorm:"foreignKey:ParentID;constraint:base_node_parent_id_fkey,OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (BaseNode) TableName() string {
	return "base_node"
}

func (m *BaseNode) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*BaseNode) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type BaseNodeFolder struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	BaseID           string    `gorm:"column:base_id;index:base_node_folder_base_id_idx;not null" json:"base_id"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"last_modified_by"`
	DeletedTime      *DateTime `gorm:"column:deleted_time;type:timestamp(3)" json:"deleted_time"`
}

func (BaseNodeFolder) TableName() string {
	return "base_node_folder"
}

func (m *BaseNodeFolder) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*BaseNodeFolder) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type BaseShare struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	BaseID           string    `gorm:"column:base_id;uniqueIndex:base_share_base_id_node_id_key,priority:1;index:base_share_base_id_idx;not null" json:"base_id"`
	ShareID          string    `gorm:"column:share_id;uniqueIndex:base_share_share_id_key;not null" json:"share_id"`
	Password         *string   `gorm:"column:password" json:"password"`
	NodeID           *string   `gorm:"column:node_id;uniqueIndex:base_share_base_id_node_id_key,priority:2" json:"node_id"`
	AllowSave        *bool     `gorm:"column:allow_save" json:"allow_save"`
	AllowCopy        *bool     `gorm:"column:allow_copy" json:"allow_copy"`
	AllowEdit        *bool     `gorm:"column:allow_edit" json:"allow_edit"`
	Enabled          *bool     `gorm:"column:enabled;default:true;not null" json:"enabled"`
	CreatedTime      DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy        string    `gorm:"column:created_by;not null" json:"created_by"`
	LastModifiedTime *DateTime `gorm:"column:last_modified_time;type:timestamp(3)" json:"last_modified_time"`
}

func (BaseShare) TableName() string {
	return "base_share"
}

func (m *BaseShare) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
	}
	if m.LastModifiedTime == nil {
		m.LastModifiedTime = &DateTime{Time: now}
	}
	return nil
}

func (*BaseShare) BeforeUpdate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if !tx.Statement.Changed("LastModifiedTime") {
		tx.Statement.SetColumn("LastModifiedTime", &DateTime{Time: now})
	}
	return nil
}

type ShortLink struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	Code        string    `gorm:"column:code;uniqueIndex:short_link_code_key;type:varchar(32);not null" json:"code"`
	Type        string    `gorm:"column:type;uniqueIndex:short_link_type_resource_id_key,priority:1;type:varchar(32);not null" json:"type"`
	ResourceID  string    `gorm:"column:resource_id;uniqueIndex:short_link_type_resource_id_key,priority:2;not null" json:"resource_id"`
	CreatedTime DateTime  `gorm:"column:created_time;type:timestamp(3);default:CURRENT_TIMESTAMP;not null" json:"created_time"`
	CreatedBy   string    `gorm:"column:created_by;not null" json:"created_by"`
	DeletedTime *DateTime `gorm:"column:deleted_time;index:short_link_deleted_time_idx;type:timestamp(3)" json:"deleted_time"`
}

func (ShortLink) TableName() string {
	return "short_link"
}

func (m *ShortLink) BeforeCreate(tx *gorm.DB) error {
	now := tx.NowFunc().UTC().Truncate(time.Millisecond)
	if m.ID == "" {
		m.ID = cuid.New()
	}
	if m.CreatedTime.IsZero() {
		m.CreatedTime = DateTime{Time: now}
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
