use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Custody")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "teamMemberId", column_type = "Text")]
    pub team_member_id: String,
    #[sea_orm(column_name = "assetId", column_type = "Text")]
    pub asset_id: String,
    #[sea_orm(column_name = "kitCustodyId", column_type = "Text")]
    pub kit_custody_id: Option<String>,
    #[sea_orm(default_value = 1)]
    pub quantity: i32,
    #[sea_orm(column_name = "locationId", column_type = "Text")]
    pub location_id: Option<String>,
    #[sea_orm(column_name = "sourceUnknown", default_value = false)]
    pub source_unknown: bool,
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
        belongs_to = "super::team_member::Entity",
        from = "Column::TeamMemberId",
        to = "super::team_member::Column::Id",
        fk_name = "Custody_teamMemberId_fkey"
    )]
    Custodian,
    #[sea_orm(
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Custody_assetId_fkey"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::kit_custody::Entity",
        from = "Column::KitCustodyId",
        to = "super::kit_custody::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Custody_kitCustodyId_fkey"
    )]
    KitCustody,
    #[sea_orm(
        belongs_to = "super::location::Entity",
        from = "Column::LocationId",
        to = "super::location::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "Custody_locationId_fkey"
    )]
    Location,
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Custodian.def()
    }
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Asset.def()
    }
}

impl Related<super::kit_custody::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::KitCustody.def()
    }
}

impl Related<super::location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Location.def()
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
