use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "computed_update_outbox")]
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
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
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
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::computed_update_outbox_seed::Entity")]
    Seeds,
}

impl Related<super::computed_update_outbox_seed::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Seeds.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.next_run_at.is_not_set() {
            self.next_run_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
