use super::dashboard_widget_chart_type::DashboardWidgetChartType;
use super::dashboard_widget_views::DashboardWidgetViews;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "dashboard_widgets")]
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
    pub project_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: String,
    pub view: DashboardWidgetViews,
    #[sea_orm(column_type = "JsonBinary")]
    pub dimensions: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub metrics: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub filters: Json,
    pub chart_type: DashboardWidgetChartType,
    #[sea_orm(column_type = "JsonBinary")]
    pub chart_config: Json,
    #[sea_orm(default_value = 1)]
    pub min_version: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "dashboard_widgets_created_by_fkey"
    )]
    CreatedByUser,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UpdatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "dashboard_widgets_updated_by_fkey"
    )]
    UpdatedByUser,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "dashboard_widgets_project_id_fkey"
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
