use super::dashboard_scope::DashboardScope;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
/// made before the split were copied here with their ids, so stars and insight
/// pointers still name them. No relation, per the database rule.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DashboardsV2")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The project that owns the board; only there can it be edited.
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    /// Free text under the board's name; null until a member writes one.
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    /// The member who created it; null for boards that predate authorship.
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: Option<String>,
    /// Who sees it. Only the author changes it; a board with no author stays PROJECT.
    #[sea_orm(default_value = "PROJECT")]
    pub scope: DashboardScope,
    /// The owning project's organization, stamped when the board is first set to
    /// ORGANIZATION; null before.
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    /// The product whose sidebar lists it: llm-ops, me, gateway or governance (ADR-002,
    /// decision 8). Every board made before homes existed is llm-ops.
    #[sea_orm(
        column_name = "homeProduct",
        column_type = "Text",
        default_value = "llm-ops"
    )]
    pub home_product: String,
    /// Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
    /// 2026-10-28); dropped in a later migration, once no serving release reads it.
    #[sea_orm(column_name = "excludedOrigins", column_type = "JsonBinary")]
    pub excluded_origins: Option<Json>,
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
