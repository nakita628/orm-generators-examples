use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One browser a person subscribed to Web Push, owned by notification. A
/// person holds one row per device; the endpoint names the browser, so a
/// browser that subscribes again replaces its own row.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WebPushSubscription")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// The person the pushes are for. A plain column: the row leaves with the
    /// person through notification's own subscriber on the user lifecycle.
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    /// The push service URL the browser handed out. Unique per browser.
    #[sea_orm(unique)]
    pub endpoint: String,
    /// The browser's P-256 public key and auth secret, base64url, as the
    /// browser reports them. Payloads are encrypted to these.
    pub p256dh: String,
    pub auth: String,
    #[sea_orm(column_name = "userAgent")]
    pub user_agent: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    /// When the push service last accepted a push for this browser.
    #[sea_orm(column_name = "lastSuccessAt")]
    pub last_success_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
        Ok(self)
    }
}
