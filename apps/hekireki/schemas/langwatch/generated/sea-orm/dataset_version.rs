use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DatasetVersion")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "datasetId", column_type = "Text")]
    pub dataset_id: String,
    pub version: i32,
    /// A name given by hand; a named version is never thinned by retention.
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    /// "autosave", "manual", "experiment_run", "restore" or "initial".
    #[sea_orm(column_type = "Text")]
    pub trigger: String,
    /// SHA-256 over the columns and the rows in order; equal content saves no new version.
    #[sea_orm(column_name = "contentHash", column_type = "Text")]
    pub content_hash: String,
    #[sea_orm(column_name = "columnTypes", column_type = "JsonBinary")]
    pub column_types: Json,
    #[sea_orm(column_name = "rowCount")]
    pub row_count: i32,
    #[sea_orm(column_name = "sizeBytes")]
    pub size_bytes: i64,
    #[sea_orm(column_name = "chunkCount")]
    pub chunk_count: i32,
    #[sea_orm(column_name = "restoredFromVersion")]
    pub restored_from_version: Option<i32>,
    #[sea_orm(column_name = "authorId", column_type = "Text")]
    pub author_id: Option<String>,
    /// "user", "api_key" or "system": who saved the version when authorId is null.
    #[sea_orm(column_name = "authorKind", column_type = "Text")]
    pub author_kind: Option<String>,
    /// Set when an experiment pins this version; a pinned version is never thinned.
    #[sea_orm(column_name = "pinnedAt", column_type = "custom(\"timestamp(3)\")")]
    pub pinned_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
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
