use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Pivot linking an `Asset` to the `Kit` it belongs to.
///
/// `@@unique([assetId])` enforces "at most one kit per asset" today.
/// The structure already allows multi-row-per-asset; if multi-kit
/// allocation lands later, drop the `assetId` unique constraint and
/// add a type-aware single-row trigger + sum-within-total CONSTRAINT
/// TRIGGER (parallel to the Custody / BookingAsset pivots).
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AssetKit")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "assetId", column_type = "Text")]
    pub asset_id: String,
    #[sea_orm(column_name = "kitId", column_type = "Text")]
    pub kit_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(default_value = 1)]
    pub quantity: i32,
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
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "AssetKit_assetId_fkey"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::kit::Entity",
        from = "Column::KitId",
        to = "super::kit::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "AssetKit_kitId_fkey"
    )]
    Kit,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "AssetKit_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::asset_location::Entity")]
    AssetLocations,
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

impl Related<super::asset_location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetLocations.def()
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
