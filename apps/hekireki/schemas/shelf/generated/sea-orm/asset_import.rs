use super::asset_import_phase::AssetImportPhase;
use super::asset_import_status::AssetImportStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One CSV content import. Holds the parsed rows from preview until the import
/// completes, so the job imports exactly the file the user previewed.
/// A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
/// active import per workspace. Prisma cannot express it; it lives in migration
/// 20261015120100_asset_import_one_active_per_workspace.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AssetImport")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(default_value = "DRAFT")]
    pub status: AssetImportStatus,
    #[sea_orm(default_value = "PREPARE")]
    pub phase: AssetImportPhase,
    #[sea_orm(column_name = "fileName")]
    pub file_name: Option<String>,
    #[sea_orm(column_name = "totalRows")]
    pub total_rows: i32,
    /// Rows committed so far. Advanced in the same transaction as each chunk.
    #[sea_orm(column_name = "processedRows", default_value = 0)]
    pub processed_rows: i32,
    /// Parsed rows (`StoredImportRow[]`). Cleared when the import completes.
    pub rows: Option<Json>,
    /// Name to id maps written by the PREPARE step (`ResolvedImportEntities`).
    pub resolved: Option<Json>,
    #[sea_orm(column_name = "canUseBarcodes", default_value = false)]
    pub can_use_barcodes: bool,
    /// User-facing reason the import failed. Null unless status is FAILED.
    #[sea_orm(column_name = "errorMessage")]
    pub error_message: Option<String>,
    /// Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500.
    pub warnings: Option<Json>,
    /// Exact number of warnings, including any beyond the stored cap.
    #[sea_orm(column_name = "warningCount", default_value = 0)]
    pub warning_count: i32,
    /// A step holds the import until this time. Null when no step is running.
    #[sea_orm(column_name = "lockedUntil")]
    pub locked_until: Option<DateTime>,
    #[sea_orm(column_name = "confirmedAt")]
    pub confirmed_at: Option<DateTime>,
    #[sea_orm(column_name = "startedAt")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_name = "completedAt")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    CreatedBy,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
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
