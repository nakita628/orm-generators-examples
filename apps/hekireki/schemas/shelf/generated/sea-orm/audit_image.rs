use super::audit_image_capture_source::AuditImageCaptureSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuditImage")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "imageUrl", column_type = "Text")]
    pub image_url: String,
    #[sea_orm(column_name = "thumbnailUrl", column_type = "Text")]
    pub thumbnail_url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "auditSessionId", column_type = "Text")]
    pub audit_session_id: String,
    #[sea_orm(column_name = "auditAssetId", column_type = "Text")]
    pub audit_asset_id: Option<String>,
    #[sea_orm(column_name = "uploadedById", column_type = "Text")]
    pub uploaded_by_id: Option<String>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// When the photo was taken, from its EXIF data (or the phone app's copy of
    /// it). Null when the photo carried no usable date. EXIF can be edited, so
    /// this is the best signal available, not proof.
    #[sea_orm(column_name = "capturedAt", column_type = "custom(\"timestamp(3)\")")]
    pub captured_at: Option<DateTime>,
    /// How the photo reached Shelf. Null only for photos uploaded before this
    /// field existed; a phone upload with no or an unknown source is PHONE.
    #[sea_orm(column_name = "captureSource")]
    pub capture_source: Option<AuditImageCaptureSource>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::audit_session::Entity",
        from = "Column::AuditSessionId",
        to = "super::audit_session::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "AuditImage_auditSessionId_fkey"
    )]
    AuditSession,
    #[sea_orm(
        belongs_to = "super::audit_asset::Entity",
        from = "Column::AuditAssetId",
        to = "super::audit_asset::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "AuditImage_auditAssetId_fkey"
    )]
    AuditAsset,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UploadedById",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "AuditImage_uploadedById_fkey"
    )]
    UploadedBy,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "AuditImage_organizationId_fkey"
    )]
    Organization,
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

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UploadedBy.def()
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
