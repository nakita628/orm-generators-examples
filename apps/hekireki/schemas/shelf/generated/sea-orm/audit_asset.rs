use super::audit_asset_status::AuditAssetStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuditAsset")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "auditSessionId")]
    pub audit_session_id: String,
    #[sea_orm(column_name = "assetId")]
    pub asset_id: String,
    #[sea_orm(default_value = true)]
    pub expected: bool,
    #[sea_orm(default_value = "PENDING")]
    pub status: AuditAssetStatus,
    #[sea_orm(column_name = "scannedById")]
    pub scanned_by_id: Option<String>,
    #[sea_orm(column_name = "scannedAt")]
    pub scanned_at: Option<DateTime>,
    pub metadata: Option<Json>,
    /// Units the audit expects to find, snapshotted when the row is created: the
    /// units in the audit's scope (placed at its locations, held in its kits, or
    /// held by its person; at most `Asset.quantity`), or `Asset.quantity` when
    /// none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
    /// which count as one.
    ///
    /// A snapshot, not a live read: stock that moves during the audit must not
    /// move the target the auditor is counting against.
    #[sea_orm(column_name = "expectedQuantity")]
    pub expected_quantity: Option<i32>,
    /// Units the auditor counted. Null until a quantity-tracked asset is
    /// counted, and always null for INDIVIDUAL assets. A FOUND row with a count
    /// below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
    /// was located), the gap is this number.
    ///
    /// Both columns are non-negative, enforced by CHECK constraints in raw SQL
    /// (migration `20261104120000_add_audit_asset_quantities`).
    #[sea_orm(column_name = "foundQuantity")]
    pub found_quantity: Option<i32>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::audit_session::Entity",
        from = "Column::AuditSessionId",
        to = "super::audit_session::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    AuditSession,
    #[sea_orm(
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ScannedById",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull"
    )]
    ScannedBy,
    #[sea_orm(has_many = "super::audit_scan::Entity")]
    Scans,
    #[sea_orm(has_many = "super::audit_image::Entity")]
    Images,
    #[sea_orm(has_many = "super::audit_note::Entity")]
    Notes,
}

impl Related<super::audit_session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditSession.def()
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

impl Related<super::audit_scan::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scans.def()
    }
}

impl Related<super::audit_image::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Images.def()
    }
}

impl Related<super::audit_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notes.def()
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
