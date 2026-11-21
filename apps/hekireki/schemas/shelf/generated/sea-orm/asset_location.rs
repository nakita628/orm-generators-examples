use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AssetLocation")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "assetId")]
    pub asset_id: String,
    #[sea_orm(column_name = "locationId")]
    pub location_id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(default_value = 1)]
    pub quantity: i32,
    #[sea_orm(column_name = "assetKitId")]
    pub asset_kit_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::location::Entity",
        from = "Column::LocationId",
        to = "super::location::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Location,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::asset_kit::Entity",
        from = "Column::AssetKitId",
        to = "super::asset_kit::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    AssetKit,
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Asset.def()
    }
}

impl Related<super::location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Location.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::asset_kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetKit.def()
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
