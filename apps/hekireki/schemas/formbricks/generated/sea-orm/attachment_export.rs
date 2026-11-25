use super::attachment_export_status::AttachmentExportStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A survey's response attachments built into one ZIP in object storage by a background job and
/// downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
/// Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
/// survey deletes it with the survey's storage prefix.
///
/// A partial unique index in the migration allows one queued or running export per requester and
/// survey, so a double click returns the export already in flight.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AttachmentExport")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: String,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "requestedById", column_type = "Text")]
    pub requested_by_id: Option<String>,
    #[sea_orm(default_value = "queued")]
    pub status: AttachmentExportStatus,
    /// The response filter at request time, a `TResponseFilterCriteria`; null exports every response.
    #[sea_orm(column_name = "filterCriteria", column_type = "JsonBinary")]
    pub filter_criteria: Option<Json>,
    /// IANA zone the archive's folder names are stamped in: the org's display zone at request time.
    #[sea_orm(column_name = "timeZone", column_type = "Text")]
    pub time_zone: String,
    /// Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
    /// older pod dropped as unknown during a rolling upgrade.
    #[sea_orm(column_name = "dispatchCount", default_value = 0)]
    pub dispatch_count: i32,
    /// The pre-flight's counts, shown while the archive builds.
    #[sea_orm(column_name = "fileCount")]
    pub file_count: i32,
    #[sea_orm(column_name = "responseCount")]
    pub response_count: i32,
    /// What the archive actually holds, once built.
    #[sea_orm(column_name = "includedFileCount")]
    pub included_file_count: Option<i32>,
    #[sea_orm(column_name = "totalBytes")]
    pub total_bytes: Option<i64>,
    #[sea_orm(default_value = false)]
    pub truncated: bool,
    #[sea_orm(column_name = "storageKey", column_type = "Text")]
    pub storage_key: Option<String>,
    /// A short code (`stale`, `storage_upload_failed`, …), never a message or file name.
    #[sea_orm(column_name = "errorCode", column_type = "Text")]
    pub error_code: Option<String>,
    #[sea_orm(column_name = "startedAt", column_type = "custom(\"timestamp(3)\")")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_name = "finishedAt", column_type = "custom(\"timestamp(3)\")")]
    pub finished_at: Option<DateTime>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "(Column::SurveyId, Column::WorkspaceId)",
        to = "(super::survey::Column::Id, super::survey::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "AttachmentExport_surveyId_workspaceId_fkey"
    )]
    Survey,
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "AttachmentExport_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::RequestedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "AttachmentExport_requestedById_fkey"
    )]
    RequestedBy,
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
    }
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RequestedBy.def()
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
