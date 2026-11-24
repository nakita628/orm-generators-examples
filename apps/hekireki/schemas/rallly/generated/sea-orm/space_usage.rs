use super::space_usage_resource::SpaceUsageResource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Monthly counters for metered actions charged to a space, independent of
/// the records they produce, so deleting a poll does not give back what it
/// spent.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "space_usage")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub space_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    pub resource: SpaceUsageResource,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Date")]
    pub period: Date,
    #[sea_orm(default_value = 0)]
    pub count: i32,
    pub near_limit_warned_at: Option<DateTime>,
    pub at_limit_warned_at: Option<DateTime>,
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::space::Entity",
        from = "Column::SpaceId",
        to = "super::space::Column::Id",
        on_delete = "Cascade"
    )]
    Space,
}

impl Related<super::space::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Space.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, _insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
