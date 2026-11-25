use super::data_db_connection_health_state::DataDbConnectionHealthState;
use super::data_db_connection_status::DataDbConnectionStatus;
use super::data_db_provider::DataDbProvider;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "data_db_connection")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(default_value = "postgres")]
    pub provider: DataDbProvider,
    #[sea_orm(column_type = "Text")]
    pub encrypted_url: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub url_fingerprint: String,
    #[sea_orm(column_type = "Text")]
    pub display_host: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub display_database: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub internal_schema: String,
    #[sea_orm(default_value = "pending")]
    pub status: DataDbConnectionStatus,
    #[sea_orm(column_type = "Text")]
    pub schema_version: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub capabilities: Option<Json>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_validated_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub last_error: Option<String>,
    #[sea_orm(default_value = "healthy")]
    pub health_state: DataDbConnectionHealthState,
    #[sea_orm(column_type = "Text")]
    pub health_reason: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub health_changed_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_health_check_at: Option<DateTime>,
    #[sea_orm(default_value = 0)]
    pub consecutive_health_failures: i32,
    #[sea_orm(column_type = "Text")]
    pub created_by: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_time: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_modified_time: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::space_data_db_binding::Entity")]
    SpaceBindings,
    #[sea_orm(
        has_many = "super::space_data_db_migration_job::Entity",
        from = "Column::Id",
        to = "super::space_data_db_migration_job::Column::SourceConnectionId"
    )]
    SourceMigrationJobs,
    #[sea_orm(
        has_many = "super::space_data_db_migration_job::Entity",
        from = "Column::Id",
        to = "super::space_data_db_migration_job::Column::TargetConnectionId"
    )]
    TargetMigrationJobs,
}

impl Related<super::space_data_db_binding::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SpaceBindings.def()
    }
}

impl Related<super::space_data_db_migration_job::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SourceMigrationJobs.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_time.is_not_set() {
            self.created_time = Set(now.naive_utc());
        }
        if !self.last_modified_time.is_set() {
            self.last_modified_time = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
