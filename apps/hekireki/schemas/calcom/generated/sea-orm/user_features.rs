use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "UserFeatures")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "featureId", column_type = "Text")]
    pub feature_id: String,
    pub enabled: bool,
    #[sea_orm(
        column_name = "assignedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub assigned_at: DateTime,
    #[sea_orm(column_name = "assignedBy", column_type = "Text")]
    pub assigned_by: String,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "UserFeatures_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::feature::Entity",
        from = "Column::FeatureId",
        to = "super::feature::Column::Slug",
        on_delete = "Cascade",
        fk_name = "UserFeatures_featureId_fkey"
    )]
    Feature,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::feature::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Feature.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.assigned_at.is_not_set() {
            self.assigned_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
