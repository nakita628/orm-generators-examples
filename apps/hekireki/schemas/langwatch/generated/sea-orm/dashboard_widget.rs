use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
/// unplaced. Widgets made before the split were copied from CustomGraph with their
/// ids. No relation, per the database rule.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DashboardWidget")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// The DashboardsV2 board it sits on; null while unplaced.
    #[sea_orm(column_name = "dashboardId", column_type = "Text")]
    pub dashboard_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    /// The widget's srcdoc definition.
    #[sea_orm(column_type = "JsonBinary")]
    pub definition: Json,
    #[sea_orm(column_name = "gridColumn", default_value = 0)]
    pub grid_column: i32,
    #[sea_orm(column_name = "gridRow", default_value = 0)]
    pub grid_row: i32,
    #[sea_orm(column_name = "colSpan", default_value = 1)]
    pub col_span: i32,
    #[sea_orm(column_name = "rowSpan", default_value = 1)]
    pub row_span: i32,
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
pub enum Relation {}

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
