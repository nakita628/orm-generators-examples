use super::project_alert_status::ProjectAlertStatus;
use super::project_alert_type::ProjectAlertType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ProjectAlert")]
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
    pub environment_id: String,
    #[sea_orm(column_name = "channelId", column_type = "Text")]
    pub channel_id: String,
    #[sea_orm(default_value = "PENDING")]
    pub status: ProjectAlertStatus,
    pub r#type: ProjectAlertType,
    #[sea_orm(column_name = "taskRunAttemptId", column_type = "Text")]
    pub task_run_attempt_id: Option<String>,
    #[sea_orm(column_name = "taskRunId", column_type = "Text")]
    pub task_run_id: Option<String>,
    #[sea_orm(column_name = "workerDeploymentId", column_type = "Text")]
    pub worker_deployment_id: Option<String>,
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
        fk_name = "ProjectAlert_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::EnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "ProjectAlert_environmentId_fkey"
    )]
    Environment,
    #[sea_orm(
        belongs_to = "super::project_alert_channel::Entity",
        from = "Column::ChannelId",
        to = "super::project_alert_channel::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "ProjectAlert_channelId_fkey"
    )]
    Channel,
    #[sea_orm(
        belongs_to = "super::worker_deployment::Entity",
        from = "Column::WorkerDeploymentId",
        to = "super::worker_deployment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "ProjectAlert_workerDeploymentId_fkey"
    )]
    WorkerDeployment,
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

impl Related<super::project_alert_channel::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Channel.def()
    }
}

impl Related<super::worker_deployment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkerDeployment.def()
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
