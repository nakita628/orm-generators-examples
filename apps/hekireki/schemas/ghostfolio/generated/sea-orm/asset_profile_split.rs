use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A stock split of an asset profile. The numerator is the number of shares
/// held after the split, the denominator the number of shares held before, for
/// example 4 and 1 for a 4:1 split or 1 and 10 for a 1:10 reverse split.
///
/// The ratio is stored as two integers instead of the resulting factor, so that
/// it stays exact, for example 1/3 for a 1:3 reverse split.
///
/// Activities before the date are adjusted by the ratio: the quantity is
/// multiplied by it and the unit price divided by it, so that the total value
/// of an activity stays the same.
///
/// Market data is already split-adjusted by the data providers and must not be
/// adjusted again.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AssetProfileSplit")]
pub struct Model {
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    pub date: DateTime,
    pub denominator: i32,
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub numerator: i32,
    #[sea_orm(column_name = "symbolProfileId")]
    pub symbol_profile_id: String,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::symbol_profile::Entity",
        from = "Column::SymbolProfileId",
        to = "super::symbol_profile::Column::Id",
        on_delete = "Cascade"
    )]
    SymbolProfile,
}

impl Related<super::symbol_profile::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SymbolProfile.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
