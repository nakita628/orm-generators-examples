use super::import_status::ImportStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "imports")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(default_expr = "Expr::cust(\"gen_random_uuid()\")")]
    pub id: Uuid,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "jobId", column_type = "Text")]
    pub job_id: Option<String>,
    pub status: ImportStatus,
    #[sea_orm(column_name = "statusMessage", column_type = "Text")]
    pub status_message: Option<String>,
    #[sea_orm(column_name = "errorMessage", column_type = "Text")]
    pub error_message: Option<String>,
    /// [IPrismaImportConfig]
    #[sea_orm(column_type = "JsonBinary")]
    pub config: Json,
    #[sea_orm(column_name = "totalEvents", default_value = 0)]
    pub total_events: i32,
    #[sea_orm(column_name = "processedEvents", default_value = 0)]
    pub processed_events: i32,
    #[sea_orm(column_name = "currentStep", column_type = "Text")]
    pub current_step: Option<String>,
    #[sea_orm(column_name = "currentBatch", column_type = "Text")]
    pub current_batch: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "imports_projectId_fkey"
    )]
    Project,
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
