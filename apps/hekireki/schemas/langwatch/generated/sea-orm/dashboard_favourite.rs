use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A member's star, in their own order within a project: a board or a From
/// LangWatch template. Personal: one member's stars never show to another. No
/// foreign key (per the database rule): the dashboard service removes a board's
/// favourites when it deletes the board, so a star never outlives its dashboard.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DashboardFavourite")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    /// A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
    /// the star is a template.
    #[sea_orm(column_name = "dashboardId", column_type = "Text")]
    pub dashboard_id: Option<String>,
    /// The starred From LangWatch template id; null when the star is a board.
    #[sea_orm(column_name = "templateId", column_type = "Text")]
    pub template_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// The member's drag order among their own stars in the project.
    pub position: i32,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
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
        Ok(self)
    }
}
