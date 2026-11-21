use super::data_migration_status::DataMigrationStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Tracks database schema migrations.
/// Used to manage and track the state of data structure changes.
///
/// @property id - Unique identifier for the migration
/// @property name - Name of the migration
/// @property status - Current state of the migration (pending, applied, failed)
/// @property startedAt - When the migration began
/// @property finishedAt - When the migration completed
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DataMigration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub started_at: DateTime,
    pub finished_at: Option<DateTime>,
    #[sea_orm(unique)]
    pub name: String,
    pub status: DataMigrationStatus,
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
        if insert && self.started_at.is_not_set() {
            self.started_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
