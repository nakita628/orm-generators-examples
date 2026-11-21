use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A person's push notification switches. A missing row means every switch
/// is on, so no backfill is needed. The three category columns are named
/// after `PushCategory` in apps/webapp/app/modules/push/types.ts.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NotificationPreference")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    #[sea_orm(column_name = "pushEnabled", default_value = true)]
    pub push_enabled: bool,
    #[sea_orm(default_value = true)]
    pub bookings: bool,
    #[sea_orm(default_value = true)]
    pub audits: bool,
    #[sea_orm(column_name = "stockAndReminders", default_value = true)]
    pub stock_and_reminders: bool,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
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
