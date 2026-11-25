use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuditScan")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "auditSessionId", column_type = "Text")]
    pub audit_session_id: String,
    #[sea_orm(column_name = "auditAssetId", column_type = "Text")]
    pub audit_asset_id: Option<String>,
    #[sea_orm(column_name = "assetId", column_type = "Text")]
    pub asset_id: Option<String>,
    #[sea_orm(column_name = "scannedById", column_type = "Text")]
    pub scanned_by_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub code: Option<String>,
    /// What was scanned, recorded BY VALUE at scan time.
    ///
    /// `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset
    /// leaves this row pointing at nothing: no name, and no way to tell whether
    /// it had been expected. An audit is a historical record, and a record whose
    /// entries can silently lose their meaning is not one — so the two
    /// descriptive facts are snapshotted here rather than only referenced.
    ///
    /// Readers prefer the LIVE asset when it still exists (a rename should show
    /// the current name) and fall back to these.
    ///
    /// Null means the snapshot could not be taken: scans whose asset had already
    /// been deleted before these columns were backfilled. Nothing can recover
    /// those, and they degrade to the scanned `code`, which also survives. Every
    /// other row carries a snapshot, so null is never merely "old".
    #[sea_orm(column_name = "assetTitle", column_type = "Text")]
    pub asset_title: Option<String>,
    #[sea_orm(column_name = "wasExpected")]
    pub was_expected: Option<bool>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(
        column_name = "scannedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub scanned_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::audit_session::Entity",
        from = "Column::AuditSessionId",
        to = "super::audit_session::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "AuditScan_auditSessionId_fkey"
    )]
    AuditSession,
    #[sea_orm(
        belongs_to = "super::audit_asset::Entity",
        from = "Column::AuditAssetId",
        to = "super::audit_asset::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "AuditScan_auditAssetId_fkey"
    )]
    AuditAsset,
    #[sea_orm(
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "AuditScan_assetId_fkey"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ScannedById",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "AuditScan_scannedById_fkey"
    )]
    ScannedBy,
}

impl Related<super::audit_session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditSession.def()
    }
}

impl Related<super::audit_asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditAsset.def()
    }
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Asset.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScannedBy.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.scanned_at.is_not_set() {
            self.scanned_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
