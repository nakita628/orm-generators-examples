use super::monitor_severity::MonitorSeverity;
use super::monitor_status::MonitorStatus;
use super::monitor_threshold_operator::MonitorThresholdOperator;
use super::monitor_view::MonitorView;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "monitors")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub updated_by: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
    pub view: MonitorView,
    #[sea_orm(column_type = "JsonBinary")]
    pub filters: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub metric: Json,
    pub window_ms: i64,
    pub cadence_ms: i64,
    pub threshold_operator: MonitorThresholdOperator,
    #[sea_orm(column_type = "Decimal(Some((65, 30)))")]
    pub alert_threshold: Decimal,
    #[sea_orm(column_type = "Decimal(Some((65, 30)))")]
    pub warning_threshold: Option<Decimal>,
    #[sea_orm(default_value = "UNKNOWN")]
    pub severity: MonitorSeverity,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub severity_changed_at: Option<DateTime>,
    #[sea_orm(column_type = "JsonBinary")]
    pub no_data: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub renotify: Json,
    #[sea_orm(default_value = "ACTIVE")]
    pub status: MonitorStatus,
    pub scheduler_batch_id: i64,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub next_run_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_published_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_claimed_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_completed_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub tags: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub trigger_ids: Vec<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub alerted_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "monitors_created_by_fkey"
    )]
    CreatedByUser,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UpdatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "monitors_updated_by_fkey"
    )]
    UpdatedByUser,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "monitors_project_id_fkey"
    )]
    Project,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedByUser.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
