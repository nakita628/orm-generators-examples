use super::dashboard_scope::DashboardScope;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Dashboard")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// The project that owns the board; only there can it be edited.
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    pub name: String,
    /// Legacy board sort position. The Dashboards v2 app orders boards by a
    /// member's favourites and never writes this; kept because the legacy
    /// Analytics reports section and the REST reorder endpoint still write it,
    /// and a rolling deploy's older image still reads it.
    #[sea_orm(default_value = 0)]
    pub order: i32,
    /// Free text under the board's name; null until a member writes one.
    /// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    pub description: Option<String>,
    /// The member who created it; null for boards that predate the column.
    /// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: Option<String>,
    /// Who sees it. Only the author changes it; a board with no author stays PROJECT.
    /// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    #[sea_orm(default_value = "PROJECT")]
    pub scope: DashboardScope,
    /// The owning project's organization, stamped when the board is first set to
    /// ORGANIZATION; null before. No relation, per the database rule.
    /// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id"
    )]
    Project,
    #[sea_orm(has_many = "super::custom_graph::Entity")]
    Graphs,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::custom_graph::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Graphs.def()
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
