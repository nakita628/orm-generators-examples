use super::chart_type::ChartType;
use super::interval::Interval;
use super::metric::Metric;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "reports")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(default_expr = "Expr::cust(\"gen_random_uuid()\")")]
    pub id: Uuid,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    pub interval: Interval,
    #[sea_orm(column_type = "Text", default_value = "30d")]
    pub range: String,
    #[sea_orm(column_name = "chartType")]
    pub chart_type: ChartType,
    #[sea_orm(
        column_name = "lineType",
        column_type = "Text",
        default_value = "monotone"
    )]
    pub line_type: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub breakdowns: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub events: Json,
    #[sea_orm(
        column_name = "globalFilters",
        column_type = "JsonBinary",
        default_value = "[]"
    )]
    pub global_filters: Json,
    #[sea_orm(column_type = "Text")]
    pub formula: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub unit: Option<String>,
    #[sea_orm(default_value = "sum")]
    pub metric: Metric,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(default_value = false)]
    pub previous: bool,
    #[sea_orm(column_type = "Text")]
    pub criteria: Option<String>,
    #[sea_orm(column_name = "funnelGroup", column_type = "Text")]
    pub funnel_group: Option<String>,
    #[sea_orm(column_name = "funnelWindow")]
    pub funnel_window: Option<f64>,
    /// [IReportOptions]
    #[sea_orm(column_type = "JsonBinary")]
    pub options: Option<Json>,
    /// How a dashboard tile shows the report: { view: 'chart' | 'table' }
    #[sea_orm(column_type = "JsonBinary")]
    pub display: Option<Json>,
    #[sea_orm(
        column_name = "visibleSeries",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub visible_series: Vec<String>,
    #[sea_orm(column_name = "startDate", column_type = "Text")]
    pub start_date: Option<String>,
    #[sea_orm(column_name = "endDate", column_type = "Text")]
    pub end_date: Option<String>,
    #[sea_orm(column_name = "dashboardId", column_type = "Text")]
    pub dashboard_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "reports_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::dashboard::Entity",
        from = "Column::DashboardId",
        to = "super::dashboard::Column::Id",
        on_delete = "Cascade",
        fk_name = "reports_dashboardId_fkey"
    )]
    Dashboard,
    #[sea_orm(has_one = "super::report_layout::Entity")]
    Layout,
    #[sea_orm(has_one = "super::share_report::Entity")]
    Share,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::dashboard::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dashboard.def()
    }
}

impl Related<super::report_layout::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Layout.def()
    }
}

impl Related<super::share_report::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Share.def()
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
