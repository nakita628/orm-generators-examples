use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "space_webhooks")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub space_id: String,
    pub url: String,
    /// AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation
    pub secret: String,
    /// Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations
    pub events: Vec<String>,
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    /// The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape
    pub version: String,
    /// Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history
    pub cursor: DateTime,
    pub last_delivered_at: Option<DateTime>,
    pub last_failed_at: Option<DateTime>,
    /// Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap
    #[sea_orm(default_value = 0)]
    pub consecutive_failures: i32,
    pub created_at: DateTime,
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
    #[sea_orm(has_many = "super::webhook_delivery::Entity")]
    Deliveries,
}

impl Related<super::space::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Space.def()
    }
}

impl Related<super::webhook_delivery::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Deliveries.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.cursor.is_not_set() {
            self.cursor = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
