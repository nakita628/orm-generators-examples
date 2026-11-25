use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a chart widget on a dashboard.
/// Widgets are positioned using a grid layout system.
///
/// @property id - Unique identifier for the widget
/// @property dashboard - The dashboard this widget belongs to
/// @property chart - The chart displayed in this widget (chart name is used for display)
/// @property layout - Grid layout configuration (x, y, width, height)
/// @property order - Display order within the dashboard
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DashboardWidget")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "dashboardId", column_type = "Text")]
    pub dashboard_id: String,
    #[sea_orm(column_name = "chartId", column_type = "Text")]
    pub chart_id: String,
    /// [WidgetLayout] - Grid layout: { x, y, w, h }
    #[sea_orm(
        column_type = "JsonBinary",
        default_value = "{\"x\":0,\"y\":0,\"w\":4,\"h\":3}"
    )]
    pub layout: Json,
    #[sea_orm(default_value = 0)]
    pub order: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::dashboard::Entity",
        from = "Column::DashboardId",
        to = "super::dashboard::Column::Id",
        on_delete = "Cascade",
        fk_name = "DashboardWidget_dashboardId_fkey"
    )]
    Dashboard,
    #[sea_orm(
        belongs_to = "super::chart::Entity",
        from = "Column::ChartId",
        to = "super::chart::Column::Id",
        on_delete = "Cascade",
        fk_name = "DashboardWidget_chartId_fkey"
    )]
    Chart,
}

impl Related<super::dashboard::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dashboard.def()
    }
}

impl Related<super::chart::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Chart.def()
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
