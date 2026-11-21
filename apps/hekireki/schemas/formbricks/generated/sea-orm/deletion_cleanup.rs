use super::deletion_cleanup_kind::DeletionCleanupKind;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
/// storage files that can't be removed inside the database transaction. A row is written in the same
/// transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
/// queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DeletionCleanup")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    pub kind: DeletionCleanupKind,
    /// Scoping and logs. Not foreign keys: the rows they name are usually gone already.
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "workspaceId")]
    pub workspace_id: String,
    #[sea_orm(column_name = "surveyId")]
    pub survey_id: String,
    /// Hub kinds: the organisation's feedback directories, read inside the delete's transaction.
    #[sea_orm(column_name = "tenantIds")]
    pub tenant_ids: Vec<String>,
    /// `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response.
    #[sea_orm(column_name = "responseIds")]
    pub response_ids: Vec<String>,
    /// `storageFiles`: the object keys or URLs to delete.
    #[sea_orm(column_name = "fileKeys")]
    pub file_keys: Vec<String>,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    /// Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup.
    #[sea_orm(column_name = "nextAttemptAt")]
    pub next_attempt_at: DateTime,
    /// The last failure: a short code and message, never record or file content.
    #[sea_orm(column_name = "lastError")]
    pub last_error: Option<String>,
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
        if insert && self.next_attempt_at.is_not_set() {
            self.next_attempt_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
