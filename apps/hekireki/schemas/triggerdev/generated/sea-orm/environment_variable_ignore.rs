use super::environment_variable_ignore_reason::EnvironmentVariableIgnoreReason;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A reason the environment variables "Needs attention" page shows a key in one environment, which
/// the project's team chose to ignore. An ignored EMPTY goes when the value is next saved.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "EnvironmentVariableIgnore")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "environmentId", column_type = "Text")]
    pub environment_id: String,
    #[sea_orm(column_name = "variableId", column_type = "Text")]
    pub variable_id: String,
    pub reason: EnvironmentVariableIgnoreReason,
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "EnvironmentVariableIgnore_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::EnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "EnvironmentVariableIgnore_environmentId_fkey"
    )]
    Environment,
    #[sea_orm(
        belongs_to = "super::environment_variable::Entity",
        from = "Column::VariableId",
        to = "super::environment_variable::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "EnvironmentVariableIgnore_variableId_fkey"
    )]
    Variable,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "EnvironmentVariableIgnore_createdById_fkey"
    )]
    CreatedBy,
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

impl Related<super::environment_variable::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Variable.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
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
        Ok(self)
    }
}
