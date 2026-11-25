use super::bulk_action_notification_type::BulkActionNotificationType;
use super::bulk_action_status::BulkActionStatus;
use super::bulk_action_type::BulkActionType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Bulk actions, like canceling and replaying runs
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BulkActionGroup")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "environmentId", column_type = "Text")]
    pub environment_id: Option<String>,
    pub r#type: BulkActionType,
    /// When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed.
    #[sea_orm(default_value = "PENDING")]
    pub status: BulkActionStatus,
    /// The query that will be executed to get the runs to process (if null, we are passing run ids directly)
    #[sea_orm(column_name = "queryName", column_type = "Text")]
    pub query_name: Option<String>,
    /// The params that will be passed to the query
    #[sea_orm(column_type = "JsonBinary")]
    pub params: Option<Json>,
    /// Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions.
    #[sea_orm(column_name = "dedupeKey", column_type = "Text")]
    pub dedupe_key: Option<String>,
    /// The cursor that will be passed to the query (null for the first page)
    #[sea_orm(column_type = "JsonBinary")]
    pub cursor: Option<Json>,
    /// The number of runs that have been processed successfully
    #[sea_orm(column_name = "successCount", default_value = 0)]
    pub success_count: i32,
    /// The number of runs that have been failed
    #[sea_orm(column_name = "failureCount", default_value = 0)]
    pub failure_count: i32,
    /// The total number of runs that will be processed
    #[sea_orm(column_name = "totalCount", default_value = 0)]
    pub total_count: i32,
    #[sea_orm(column_name = "completionNotification", default_value = "NONE")]
    pub completion_notification: BulkActionNotificationType,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    /// The name of the bulk action
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    /// The time the bulk action was completed
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BulkActionGroup_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::EnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BulkActionGroup_environmentId_fkey"
    )]
    Environment,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "BulkActionGroup_userId_fkey"
    )]
    User,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Environment.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
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
