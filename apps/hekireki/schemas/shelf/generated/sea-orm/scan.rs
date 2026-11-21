use super::scan_code_type::ScanCodeType;
use super::scan_source::ScanSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Scan")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub latitude: Option<String>,
    pub longitude: Option<String>,
    #[sea_orm(column_name = "userAgent")]
    pub user_agent: Option<String>,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "qrId")]
    pub qr_id: Option<String>,
    #[sea_orm(column_name = "rawQrId")]
    pub raw_qr_id: Option<String>,
    #[sea_orm(column_name = "codeType", default_value = "QR")]
    pub code_type: ScanCodeType,
    pub code: Option<String>,
    #[sea_orm(column_name = "barcodeId")]
    pub barcode_id: Option<String>,
    #[sea_orm(column_name = "assetId")]
    pub asset_id: Option<String>,
    #[sea_orm(column_name = "kitId")]
    pub kit_id: Option<String>,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    pub source: Option<ScanSource>,
    #[sea_orm(column_name = "manuallyGenerated", default_value = false)]
    pub manually_generated: bool,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::qr::Entity",
        from = "Column::QrId",
        to = "super::qr::Column::Id",
        on_delete = "SetNull"
    )]
    Qr,
    #[sea_orm(
        belongs_to = "super::barcode::Entity",
        from = "Column::BarcodeId",
        to = "super::barcode::Column::Id",
        on_delete = "SetNull"
    )]
    Barcode,
    #[sea_orm(
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_delete = "SetNull"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::kit::Entity",
        from = "Column::KitId",
        to = "super::kit::Column::Id",
        on_delete = "SetNull"
    )]
    Kit,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade"
    )]
    Organization,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::qr::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Qr.def()
    }
}

impl Related<super::barcode::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Barcode.def()
    }
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Asset.def()
    }
}

impl Related<super::kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Kit.def()
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
