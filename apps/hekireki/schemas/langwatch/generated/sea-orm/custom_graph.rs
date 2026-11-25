use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "CustomGraph")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    /// The chart's definition, in the shape its `kind` names.
    #[sea_orm(column_type = "JsonBinary")]
    pub graph: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub filters: Option<Json>,
    /// Discriminator for the chart shapes that share this table:
    /// - "builder" (default) - the chart builder's own `graph` payload
    /// - "workbench_sql"     - a saved LangWatchQL workbench chart, whose
    /// `graph` holds the versioned { sql, parameters, vegaLiteSpec }
    /// definition (see server/analytics/saved-workbench-charts)
    /// - "dashboard_srcdoc"  - frozen legacy copies of Dashboards v2 widgets, which
    /// now live in DashboardWidget; nothing new writes this kind
    /// Each reader filters by its own kind, so neither sees the other's rows and
    /// a builder chart can never be interpreted as a workbench definition.
    #[sea_orm(column_type = "Text", default_value = "builder")]
    pub kind: String,
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
    #[sea_orm(column_name = "dashboardId", column_type = "Text")]
    pub dashboard_id: Option<String>,
    #[sea_orm(column_name = "gridColumn", default_value = 0)]
    pub grid_column: i32,
    #[sea_orm(column_name = "gridRow", default_value = 0)]
    pub grid_row: i32,
    #[sea_orm(column_name = "colSpan", default_value = 1)]
    pub col_span: i32,
    #[sea_orm(column_name = "rowSpan", default_value = 1)]
    pub row_span: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "CustomGraph_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::dashboard::Entity",
        from = "Column::DashboardId",
        to = "super::dashboard::Column::Id",
        on_delete = "Cascade",
        fk_name = "CustomGraph_dashboardId_fkey"
    )]
    Dashboard,
    #[sea_orm(has_many = "super::trigger_sent::Entity")]
    TriggerSent,
    #[sea_orm(has_one = "super::trigger::Entity")]
    Trigger,
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

impl Related<super::trigger_sent::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TriggerSent.def()
    }
}

impl Related<super::trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Trigger.def()
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
