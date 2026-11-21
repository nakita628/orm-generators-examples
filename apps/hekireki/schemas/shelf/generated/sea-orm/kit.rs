use super::kit_status::KitStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Kit")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub name: String,
    pub description: Option<String>,
    #[sea_orm(default_value = "AVAILABLE")]
    pub status: KitStatus,
    pub image: Option<String>,
    #[sea_orm(column_name = "imageExpiration")]
    pub image_expiration: Option<DateTime>,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: String,
    #[sea_orm(column_name = "categoryId")]
    pub category_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "locationId")]
    pub location_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id"
    )]
    CreatedBy,
    #[sea_orm(
        belongs_to = "super::category::Entity",
        from = "Column::CategoryId",
        to = "super::category::Column::Id"
    )]
    Category,
    #[sea_orm(
        belongs_to = "super::location::Entity",
        from = "Column::LocationId",
        to = "super::location::Column::Id"
    )]
    Location,
    #[sea_orm(has_many = "super::asset_kit::Entity")]
    AssetKits,
    #[sea_orm(has_many = "super::qr::Entity")]
    QrCodes,
    #[sea_orm(has_many = "super::barcode::Entity")]
    Barcodes,
    #[sea_orm(has_many = "super::report_found::Entity")]
    Reports,
    #[sea_orm(has_many = "super::scan::Entity")]
    Scans,
    #[sea_orm(has_one = "super::kit_custody::Entity")]
    Custody,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::category::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Category.def()
    }
}

impl Related<super::location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Location.def()
    }
}

impl Related<super::asset_kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetKits.def()
    }
}

impl Related<super::qr::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::QrCodes.def()
    }
}

impl Related<super::barcode::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Barcodes.def()
    }
}

impl Related<super::report_found::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Reports.def()
    }
}

impl Related<super::scan::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scans.def()
    }
}

impl Related<super::kit_custody::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Custody.def()
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
