use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "domain_event_outbox")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub base_id: String,
    pub table_id: Option<String>,
    pub message_name: String,
    pub schema_version: i32,
    pub aggregate_id: Option<String>,
    pub payload: Json,
    pub payload_bytes: i32,
    pub catalog_generation: i32,
    pub required_consumers: Json,
    pub binding_id: Option<String>,
    pub storage_epoch: Option<i32>,
    #[sea_orm(default_value = true)]
    pub unpublished: bool,
    pub settled: Option<String>,
    pub settled_at: Option<DateTime>,
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
