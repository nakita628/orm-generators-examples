use super::prisma_date_time::PrismaDateTime;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "scheduled_jobs")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub prompt: String,
    #[sea_orm(column_type = "Text")]
    pub tools: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub schedule: String,
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    #[sea_orm(column_name = "lastRunAt")]
    pub last_run_at: Option<PrismaDateTime>,
    #[sea_orm(column_name = "nextRunAt")]
    pub next_run_at: Option<PrismaDateTime>,
    #[sea_orm(column_name = "createdAt", default_expr = "Expr::current_timestamp()")]
    pub created_at: PrismaDateTime,
    #[sea_orm(column_name = "updatedAt", default_expr = "Expr::current_timestamp()")]
    pub updated_at: PrismaDateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::scheduled_job_runs::Entity")]
    Runs,
}

impl Related<super::scheduled_job_runs::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Runs.def()
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
            self.created_at = Set(now.into());
        }
        if insert && self.updated_at.is_not_set() {
            self.updated_at = Set(now.into());
        }
        Ok(self)
    }
}
