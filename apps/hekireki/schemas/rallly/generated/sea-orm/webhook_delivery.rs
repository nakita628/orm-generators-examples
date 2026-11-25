use super::webhook_delivery_status::WebhookDeliveryStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "webhook_deliveries")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub webhook_id: String,
    /// Soft reference to the PollActivity row this delivery was fanned out from; with webhookId it makes fan-out idempotent
    #[sea_orm(column_type = "Text")]
    pub activity_id: String,
    #[sea_orm(column_type = "Text")]
    pub event_type: String,
    /// Frozen at fan-out: retries resend exactly what the first attempt sent
    #[sea_orm(column_type = "JsonBinary")]
    pub payload: Json,
    #[sea_orm(default_value = "pending")]
    pub status: WebhookDeliveryStatus,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub next_attempt_at: DateTime,
    pub last_response_status: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub last_error: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::space_webhook::Entity",
        from = "Column::WebhookId",
        to = "super::space_webhook::Column::Id",
        on_delete = "Cascade",
        fk_name = "webhook_deliveries_webhook_id_fkey"
    )]
    Webhook,
}

impl Related<super::space_webhook::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Webhook.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.next_attempt_at.is_not_set() {
            self.next_attempt_at = Set(now.naive_utc());
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
