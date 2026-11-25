use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a dashboard containing multiple charts.
/// Dashboards aggregate analytics insights at the workspace level.
///
/// @property id - Unique identifier for the dashboard
/// @property name - Display name of the dashboard
/// @property workspace - The workspace this dashboard belongs to
/// @property widgets - Charts on this dashboard
/// @property filters - The dashboard filters pinned to its top bar, in order: `{ id, field, operator, label }`,
/// at most 10 (ADR-0026). Each viewer's picks live in their own browser, not here.
/// @property createdBy - User who created the dashboard
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Dashboard")]
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
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub filters: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "Dashboard_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "Dashboard_createdBy_fkey"
    )]
    Creator,
    #[sea_orm(has_many = "super::dashboard_widget::Entity")]
    Widgets,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creator.def()
    }
}

impl Related<super::dashboard_widget::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Widgets.def()
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
