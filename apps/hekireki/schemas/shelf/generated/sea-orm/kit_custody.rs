use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "KitCustody")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "custodianId")]
    pub custodian_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "kitId")]
    pub kit_id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::team_member::Entity",
        from = "Column::CustodianId",
        to = "super::team_member::Column::Id"
    )]
    Custodian,
    #[sea_orm(
        belongs_to = "super::kit::Entity",
        from = "Column::KitId",
        to = "super::kit::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Kit,
    #[sea_orm(has_many = "super::custody::Entity")]
    InheritedCustody,
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Custodian.def()
    }
}

impl Related<super::kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Kit.def()
    }
}

impl Related<super::custody::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InheritedCustody.def()
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
