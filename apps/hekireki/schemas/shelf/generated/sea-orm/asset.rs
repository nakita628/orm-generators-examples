use super::asset_status::AssetStatus;
use super::asset_type::AssetType;
use super::consumption_type::ConsumptionType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Asset")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "mainImage", column_type = "Text")]
    pub main_image: Option<String>,
    #[sea_orm(column_name = "thumbnailImage", column_type = "Text")]
    pub thumbnail_image: Option<String>,
    #[sea_orm(
        column_name = "mainImageExpiration",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub main_image_expiration: Option<DateTime>,
    #[sea_orm(default_value = "AVAILABLE")]
    pub status: AssetStatus,
    #[sea_orm(column_name = "value")]
    pub valuation: Option<f64>,
    #[sea_orm(column_name = "availableToBook", default_value = true)]
    pub available_to_book: bool,
    #[sea_orm(column_name = "sequentialId", column_type = "Text")]
    pub sequential_id: Option<String>,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamptz(3)\")")]
    pub archived_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(default_value = "INDIVIDUAL")]
    pub r#type: AssetType,
    #[sea_orm(column_name = "assetModelId", column_type = "Text")]
    pub asset_model_id: Option<String>,
    pub quantity: Option<i32>,
    #[sea_orm(column_name = "minQuantity")]
    pub min_quantity: Option<i32>,
    #[sea_orm(column_name = "consumptionType")]
    pub consumption_type: Option<ConsumptionType>,
    #[sea_orm(column_name = "unitOfMeasure", column_type = "Text")]
    pub unit_of_measure: Option<String>,
    #[sea_orm(
        column_name = "lowStockNotifiedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub low_stock_notified_at: Option<DateTime>,
    #[sea_orm(
        column_name = "stockLedgerStartedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub stock_ledger_started_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "categoryId", column_type = "Text")]
    pub category_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "preferredBarcodeId", column_type = "Text")]
    pub preferred_barcode_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::asset_model::Entity",
        from = "Column::AssetModelId",
        to = "super::asset_model::Column::Id",
        fk_name = "Asset_assetModelId_fkey"
    )]
    AssetModel,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Asset_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Asset_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::category::Entity",
        from = "Column::CategoryId",
        to = "super::category::Column::Id",
        fk_name = "Asset_categoryId_fkey"
    )]
    Category,
    #[sea_orm(
        belongs_to = "super::barcode::Entity",
        from = "Column::PreferredBarcodeId",
        to = "super::barcode::Column::Id",
        on_delete = "SetNull",
        fk_name = "Asset_preferredBarcodeId_fkey"
    )]
    PreferredBarcode,
    #[sea_orm(has_many = "super::custody::Entity")]
    Custody,
    #[sea_orm(has_many = "super::note::Entity")]
    Notes,
    #[sea_orm(has_many = "super::qr::Entity")]
    QrCodes,
    #[sea_orm(
        has_many = "super::barcode::Entity",
        from = "Column::Id",
        to = "super::barcode::Column::AssetId"
    )]
    Barcodes,
    #[sea_orm(has_many = "super::report_found::Entity")]
    Reports,
    #[sea_orm(has_many = "super::asset_custom_field_value::Entity")]
    CustomFields,
    #[sea_orm(has_many = "super::booking_asset::Entity")]
    BookingAssets,
    #[sea_orm(has_many = "super::consumption_log::Entity")]
    ConsumptionLogs,
    #[sea_orm(has_many = "super::asset_reminder::Entity")]
    Reminders,
    #[sea_orm(has_many = "super::audit_asset::Entity")]
    AuditEntries,
    #[sea_orm(has_many = "super::audit_scan::Entity")]
    AuditScans,
    #[sea_orm(has_many = "super::scan::Entity")]
    Scans,
    #[sea_orm(has_many = "super::asset_kit::Entity")]
    AssetKits,
    #[sea_orm(has_many = "super::asset_location::Entity")]
    AssetLocations,
}

impl Related<super::asset_model::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetModel.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::category::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Category.def()
    }
}

impl Related<super::barcode::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PreferredBarcode.def()
    }
}

impl Related<super::custody::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Custody.def()
    }
}

impl Related<super::note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notes.def()
    }
}

impl Related<super::qr::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::QrCodes.def()
    }
}

impl Related<super::report_found::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Reports.def()
    }
}

impl Related<super::asset_custom_field_value::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomFields.def()
    }
}

impl Related<super::booking_asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BookingAssets.def()
    }
}

impl Related<super::consumption_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ConsumptionLogs.def()
    }
}

impl Related<super::asset_reminder::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Reminders.def()
    }
}

impl Related<super::audit_asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditEntries.def()
    }
}

impl Related<super::audit_scan::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditScans.def()
    }
}

impl Related<super::scan::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scans.def()
    }
}

impl Related<super::asset_kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetKits.def()
    }
}

impl Related<super::asset_location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetLocations.def()
    }
}

impl Related<super::tag::Entity> for Entity {
    fn to() -> RelationDef {
        super::asset_to_tag::Relation::Tag.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::asset_to_tag::Relation::Asset.def().rev())
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
