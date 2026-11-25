use super::checkpoint_type::CheckpointType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Checkpoint")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    pub r#type: CheckpointType,
    #[sea_orm(column_type = "Text")]
    pub location: String,
    #[sea_orm(column_name = "imageRef", column_type = "Text")]
    pub image_ref: String,
    #[sea_orm(column_type = "Text")]
    pub reason: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub metadata: Option<String>,
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_name = "attemptId", column_type = "Text")]
    pub attempt_id: String,
    #[sea_orm(column_name = "attemptNumber")]
    pub attempt_number: Option<i32>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "runtimeEnvironmentId", column_type = "Text")]
    pub runtime_environment_id: String,
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
        belongs_to = "super::task_run::Entity",
        from = "Column::RunId",
        to = "super::task_run::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Checkpoint_runId_fkey"
    )]
    Run,
    #[sea_orm(
        belongs_to = "super::task_run_attempt::Entity",
        from = "Column::AttemptId",
        to = "super::task_run_attempt::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Checkpoint_attemptId_fkey"
    )]
    Attempt,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Checkpoint_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::RuntimeEnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Checkpoint_runtimeEnvironmentId_fkey"
    )]
    RuntimeEnvironment,
    #[sea_orm(has_many = "super::checkpoint_restore_event::Entity")]
    Events,
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Run.def()
    }
}

impl Related<super::task_run_attempt::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Attempt.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RuntimeEnvironment.def()
    }
}

impl Related<super::checkpoint_restore_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Events.def()
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
