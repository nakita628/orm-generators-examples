use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "computed_update_dead_letter")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub base_id: String,
    #[sea_orm(column_type = "Text", default_value = "production")]
    pub worker_lane: String,
    #[sea_orm(column_type = "Text")]
    pub seed_table_id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub seed_record_ids: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub change_type: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub steps: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub edges: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    #[sea_orm(default_value = 8)]
    pub max_attempts: i32,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub next_run_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub locked_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub locked_by: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub last_error: Option<String>,
    #[sea_orm(default_value = 0)]
    pub estimated_complexity: i32,
    #[sea_orm(column_type = "Text")]
    pub plan_hash: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub dirty_stats: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub origin_run_ids: Vec<String>,
    #[sea_orm(default_value = 0)]
    pub run_total_steps: i32,
    #[sea_orm(default_value = 0)]
    pub run_completed_steps_before: i32,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub affected_table_ids: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub affected_field_ids: Vec<String>,
    pub sync_max_level: Option<i32>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub source_changed_at: Option<DateTime>,
    #[sea_orm(default_value = 0)]
    pub stage_depth: i32,
    #[sea_orm(column_type = "Text")]
    pub predecessor_task_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub trace_data: Option<Json>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub failed_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    pub delete_authorized_txid: Option<i64>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::computed_update_dead_letter_seed::Entity")]
    Seeds,
}

impl Related<super::computed_update_dead_letter_seed::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Seeds.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
