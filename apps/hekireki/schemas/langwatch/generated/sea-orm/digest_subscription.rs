use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A person's opt-out of the digest. No row, or no `unsubscribedAt`: subscribed.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DigestSubscription")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(
        column_name = "unsubscribedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub unsubscribed_at: Option<DateTime>,
    /// The send whose unsubscribe link or one-click header the person used.
    #[sea_orm(column_name = "unsubscribedFromSendId", column_type = "Text")]
    pub unsubscribed_from_send_id: Option<String>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
