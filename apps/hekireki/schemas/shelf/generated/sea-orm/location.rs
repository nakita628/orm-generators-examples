use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Location")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub name: String,
    pub description: Option<String>,
    pub address: Option<String>,
    pub latitude: Option<f64>,
    pub longitude: Option<f64>,
    #[sea_orm(column_name = "imageUrl")]
    pub image_url: Option<String>,
    #[sea_orm(column_name = "thumbnailUrl")]
    pub thumbnail_url: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "imageId")]
    pub image_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "parentId")]
    pub parent_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::image::Entity",
        from = "Column::ImageId",
        to = "super::image::Column::Id"
    )]
    Image,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::location::Entity",
        from = "Column::ParentId",
        to = "super::location::Column::Id",
        on_delete = "SetNull"
    )]
    Parent,
    #[sea_orm(has_many = "super::asset_location::Entity")]
    AssetLocations,
    #[sea_orm(has_many = "super::kit::Entity")]
    Kits,
    #[sea_orm(has_many = "super::location_note::Entity")]
    Notes,
    #[sea_orm(has_many = "super::custody::Entity")]
    Custodies,
    #[sea_orm(has_many = "super::consumption_log::Entity")]
    ConsumptionLogs,
    #[sea_orm(
        has_many = "super::location::Entity",
        from = "Column::Id",
        to = "super::location::Column::ParentId"
    )]
    Children,
}

impl Related<super::image::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Image.def()
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

impl Related<super::location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Parent.def()
    }
}

impl Related<super::asset_location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetLocations.def()
    }
}

impl Related<super::kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Kits.def()
    }
}

impl Related<super::location_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notes.def()
    }
}

impl Related<super::custody::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Custodies.def()
    }
}

impl Related<super::consumption_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ConsumptionLogs.def()
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
