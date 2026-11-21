use super::data_store_kind::DataStoreKind;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
/// Multiple organizations can share a single data store row via organizationIds.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganizationDataStore")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// Human-readable unique key (e.g. "hipaa-clickhouse-us-east")
    #[sea_orm(unique)]
    pub key: String,
    /// Organization IDs that use this data store
    #[sea_orm(column_name = "organizationIds")]
    pub organization_ids: Vec<String>,
    pub kind: DataStoreKind,
    /// Versioned config JSON. Structure is discriminated by the top-level `version` field.
    pub config: Json,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
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
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
